---
name: preview-resend-verification
description: >
  Verify preview signup or auth emails through Resend. Use when a user asks to
  bring up a PR preview, sign up with a disposable address, retrieve a
  confirmation email from Resend, extract a confirmation link, or debug preview
  auth-email delivery.
---

# Preview Resend Verification

Use this for preview auth-email checks. Load `proof-of-work` first for
repo/worklog discipline, use the browser the user requested, and load any
repo-local auth-email skill if inspecting templates, SMTP, or auth config.

Prefer the existing preview. Do not delete/recreate a PR, remote branch, or
tracking branch unless the preview environment is actually broken and the user
has asked for that disruptive fallback.

## Preview URL

Find the current preview URL from the target repo's PR comments/checks/deploy
provider. For Vercel repos, `gh pr view "$PR" --json comments,statusCheckRollup`
usually contains the preview host.

Open the preview in the requested browser. A healthy unauthenticated preview
usually redirects to a login or signup route.

## Trigger Signup

Use a disposable test recipient under a domain the team controls. Submit the
preview signup form and wait for the product's "check your email" state. Do not
commit or report the password.

## Check Resend

Use the bundled helper through the target repo's secret manager so
`RESEND_API_KEY` is present but never printed:

```bash
recipient="preview-test@example.com"

node skills/preview-resend-verification/scripts/resend-email.mjs \
  find --recipient "$recipient"
```

The output includes the Resend email id, subject, timestamp, recipient, sender,
and `last_event`. `last_event=delivered` means the recipient SMTP accepted the
message; it does not prove a mailbox UI displayed it.

To extract the confirmation link:

```bash
email_id="RESEND_EMAIL_ID"

node skills/preview-resend-verification/scripts/resend-email.mjs \
  link --id "$email_id"
```

Treat confirmation links and token values as sensitive. Do not paste them into
PR bodies, worklogs, or final reports.

## Preview Confirmation

Some auth providers send links using a production site URL. To activate a
preview account safely, preserve the token/query parameters and replace only
the host with the preview host when the repo's auth callback supports that
flow. Only open the confirmation URL when the task requires account activation;
otherwise Resend delivery plus extracted-link metadata is enough evidence.

## Record Proof

Record in the session worklog:

- preview URL opened,
- generated recipient, not the password or token,
- signup page reached the expected email-sent state,
- Resend email id, subject, `created_at`, and `last_event`,
- whether the confirmation link host matched the preview host or required a
  host replacement.

source: agents
