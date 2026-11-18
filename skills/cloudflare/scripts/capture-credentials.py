#!/usr/bin/env python3
"""Save newly issued Cloudflare credentials without logging their values."""

# Adapted from Tyler Gannon's Mission Control bucket-setup helper, commit
# 5944e9822f4c35f630345f46f6db4fac6a6e7922.

import argparse
import configparser
import contextlib
import fcntl
import json
import os
from pathlib import Path
import re
import stat
import subprocess
import sys
import tempfile
from urllib.parse import urlsplit


MAX_CAPTURE = 1024 * 1024


class CaptureError(Exception):
    pass


def capture_safari(path_pattern):
    javascript = (
        "if (location.origin !== 'https://dash.cloudflare.com' || "
        "!new RegExp(" + json.dumps("^" + path_pattern + "$") + ").test(location.pathname)) "
        "{ throw new Error('Open the expected Cloudflare token page first'); } "
        "JSON.stringify({url: location.href, text: document.body.innerText});"
    )
    applescript = (
        'tell application "Safari"\n'
        'if (count of windows) is 0 then error "No Safari window"\n'
        'return do JavaScript ' + json.dumps(javascript)
        + ' in current tab of front window\nend tell'
    )
    try:
        result = subprocess.run(
            ["osascript", "-e", applescript], capture_output=True,
            text=True, timeout=20, check=False,
        )
    except (OSError, subprocess.TimeoutExpired):
        raise CaptureError("Safari capture could not run; use browser-tool capture with --source stdin.")
    if result.returncode:
        raise CaptureError(
            "Safari did not allow capture or the success page is not active. "
            "Use browser-tool capture with --source stdin; do not change browser security settings."
        )
    return result.stdout


def read_json(raw):
    if len(raw) > MAX_CAPTURE:
        raise CaptureError("The capture exceeds the size limit.")
    try:
        return json.loads(raw)
    except ValueError:
        raise CaptureError("Expected JSON input.") from None


def read_capture(raw, path_pattern):
    capture = read_json(raw)
    if not isinstance(capture, dict) or not isinstance(capture.get("url"), str):
        raise CaptureError("The capture needs its source URL.")
    url = urlsplit(capture["url"])
    if (url.scheme != "https" or url.netloc != "dash.cloudflare.com"
            or not re.fullmatch(path_pattern, url.path)):
        raise CaptureError("The capture URL does not match the expected Cloudflare token page.")
    return capture


def extract_r2(capture):
    access = capture.get("access_key_id")
    secret = capture.get("secret_access_key")
    if access is None and secret is None and isinstance(capture.get("text"), str):
        text = capture["text"]
        access_match = re.search(r"Access Key ID[\s:]+([a-fA-F0-9]{32})\b", text)
        secret_match = re.search(r"Secret Access Key[\s:]+([a-fA-F0-9]{64})\b", text)
        access = access_match[1] if access_match else None
        secret = secret_match[1] if secret_match else None
    if (not isinstance(access, str) or not re.fullmatch(r"[a-fA-F0-9]{32}", access)
            or not isinstance(secret, str) or not re.fullmatch(r"[a-fA-F0-9]{64}", secret)):
        raise CaptureError("The generated S3 Access Key ID and Secret Access Key were not found.")
    return {"access_key_id": access, "secret_access_key": secret}


def extract_api_token(capture, account):
    prefix = "cfat_" if account else "cfut_"
    pattern = prefix + r"[A-Za-z0-9_-]{41,128}"
    value = capture.get("api_token")
    if value is None and isinstance(capture.get("text"), str):
        candidates = set(re.findall(
            r"(?<![A-Za-z0-9_-])" + pattern + r"(?![A-Za-z0-9_-])", capture["text"],
        ))
        if len(candidates) == 1:
            value = candidates.pop()
    if not isinstance(value, str) or not re.fullmatch(pattern, value):
        raise CaptureError("Expected one newly issued Cloudflare API token of the selected type.")
    return value


def save_profile(path, profile, credentials, replace):
    if not re.fullmatch(r"[A-Za-z0-9][A-Za-z0-9_.-]{0,127}", profile):
        raise CaptureError("Use a profile name containing letters, digits, dots, underscores, or hyphens.")
    path = path.expanduser().absolute()
    path.parent.mkdir(mode=0o700, parents=True, exist_ok=True)
    flags = os.O_RDWR | os.O_CREAT | getattr(os, "O_NOFOLLOW", 0)
    lock_path = path.with_name(path.name + ".mission-control.lock")
    with os.fdopen(os.open(lock_path, flags, 0o600), "r+") as lock:
        fcntl.flock(lock, fcntl.LOCK_EX)
        original = ""
        try:
            fd = os.open(path, os.O_RDONLY | getattr(os, "O_NOFOLLOW", 0))
        except FileNotFoundError:
            fd = None
        if fd is not None:
            with os.fdopen(fd, "r") as source:
                info = os.fstat(source.fileno())
                if not stat.S_ISREG(info.st_mode) or info.st_uid != os.getuid():
                    raise CaptureError("The credential destination must be a regular file owned by this user.")
                original = source.read()
        parser = configparser.RawConfigParser()
        try:
            parser.read_string(original)
        except configparser.Error:
            raise CaptureError("The existing credential file is not valid INI and was preserved.")
        if parser.has_section(profile):
            same = (parser.get(profile, "aws_access_key_id", fallback="") == credentials["access_key_id"]
                    and parser.get(profile, "aws_secret_access_key", fallback="") == credentials["secret_access_key"])
            if same:
                os.chmod(path, 0o600)
                return path, "unchanged"
            if not replace:
                raise CaptureError("That profile already has different credentials; use --replace only for an intended replacement.")
        block = (f"[{profile}]\naws_access_key_id = {credentials['access_key_id']}\n"
                 f"aws_secret_access_key = {credentials['secret_access_key']}\n")
        if parser.has_section(profile):
            section = re.compile(r"(?m)^\[" + re.escape(profile) + r"\][^\n]*\n.*?(?=^\[|\Z)", re.S)
            updated, count = section.subn(lambda _: block + "\n", original, count=1)
            if count != 1:
                raise CaptureError("The existing profile could not be located safely and was preserved.")
        else:
            updated = original.rstrip("\n") + ("\n\n" if original else "") + block
        fd, temporary = tempfile.mkstemp(prefix=".r2-credentials-", dir=path.parent)
        try:
            with os.fdopen(fd, "w") as target:
                os.fchmod(target.fileno(), 0o600)
                target.write(updated)
                target.flush()
                os.fsync(target.fileno())
            os.replace(temporary, path)
        finally:
            with contextlib.suppress(FileNotFoundError):
                os.unlink(temporary)
        return path, "saved"


def save_token(path, value, dry_run):
    if dry_run:
        print(json.dumps({"status": "ready", "destination": "token file"}))
        return
    path = path.expanduser().absolute()
    path.parent.mkdir(mode=0o700, parents=True, exist_ok=True)
    fd, temporary = tempfile.mkstemp(prefix=".cloudflare-token-", dir=path.parent)
    try:
        with os.fdopen(fd, "w") as target:
            os.fchmod(target.fileno(), 0o600)
            target.write(value + "\n")
            target.flush()
            os.fsync(target.fileno())
        try:
            os.link(temporary, path)
        except FileExistsError:
            raise CaptureError("The token file already exists. Choose a new destination.") from None
    finally:
        os.unlink(temporary)
    print(json.dumps({"status": "saved", "token_file": str(path)}))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--account", help="expected Cloudflare account ID, required for R2 and account tokens")
    parser.add_argument("--source", choices=["safari", "stdin"], default="stdin")
    destination = parser.add_mutually_exclusive_group(required=True)
    destination.add_argument("--profile", help="AWS profile to save; secret values are never printed")
    destination.add_argument("--clipboard", action="store_true", help="explicitly copy S3 credential JSON instead of saving a profile")
    destination.add_argument("--api-token-file", type=Path, help="save a new Cloudflare bearer token to a private file")
    destination.add_argument("--tunnel-token-file", type=Path, help="save a Tunnel connector token from cf JSON output")
    parser.add_argument("--credentials-file", type=Path, default=Path.home() / ".aws" / "credentials")
    parser.add_argument("--replace", action="store_true", help="allow replacing this profile's different credentials")
    parser.add_argument("--dry-run", action="store_true", help="validate the capture without writing credentials or the clipboard")
    args = parser.parse_args()
    if args.account is not None and not re.fullmatch(r"[a-f0-9]{32}", args.account):
        parser.error("--account must be a 32-character Cloudflare account ID")
    if args.replace and not args.profile:
        parser.error("--replace applies only to AWS profiles")
    if args.tunnel_token_file:
        if args.source != "stdin":
            parser.error("Tunnel token capture expects JSON on stdin")
        value = read_json(sys.stdin.read(MAX_CAPTURE + 1))
        if not isinstance(value, str) or not re.fullmatch(r"eyJ[A-Za-z0-9_+/-]+=*", value):
            raise CaptureError("Expected the JSON string returned by cf tunnels token get.")
        save_token(args.tunnel_token_file, value, args.dry_run)
        return
    if args.api_token_file:
        scope = args.account or "profile"
        path_pattern = rf"/{scope}/(?!r2/)(?:[^/]+/)*api-tokens(?:/.*)?"
    else:
        if not args.account:
            parser.error("--account is required for R2")
        path_pattern = f"/{args.account}/r2/api-tokens/success"
    raw = capture_safari(path_pattern) if args.source == "safari" else sys.stdin.read(MAX_CAPTURE + 1)
    capture = read_capture(raw, path_pattern)
    if args.api_token_file:
        value = extract_api_token(capture, args.account)
        save_token(args.api_token_file, value, args.dry_run)
        return
    credentials = extract_r2(capture)
    if args.dry_run:
        print(json.dumps({"status": "ready", "destination": "clipboard" if args.clipboard else "AWS profile", "profile": args.profile}))
    elif args.clipboard:
        subprocess.run(["pbcopy"], input=json.dumps(credentials), text=True, check=True, capture_output=True)
        print(json.dumps({"status": "copied", "warning": "Clipboard history applications may retain these credentials."}))
    else:
        path, status = save_profile(args.credentials_file, args.profile, credentials, args.replace)
        print(json.dumps({"status": status, "profile": args.profile, "credentials_file": str(path)}))


if __name__ == "__main__":
    try:
        main()
    except CaptureError as error:
        print(str(error), file=sys.stderr)
        sys.exit(1)
    except (OSError, ValueError, subprocess.SubprocessError):
        print("Credential capture/storage failed. No credential values were logged.", file=sys.stderr)
        sys.exit(1)
