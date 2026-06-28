#!/usr/bin/env node

const RESEND_API = 'https://api.resend.com';

const args = process.argv.slice(2);
const command = args.shift();

main().catch((error) => {
	console.error(error instanceof Error ? error.message : String(error));
	process.exit(1);
});

async function main() {
	if (!command || ['-h', '--help', 'help'].includes(command)) {
		console.log('Usage: resend-email.mjs find --recipient EMAIL [--limit N] | link --id EMAIL_ID');
		return;
	}

	const key = process.env.RESEND_API_KEY;
	if (!key) throw new Error('RESEND_API_KEY is required');

	if (command === 'find') {
		const recipient = requiredArg('--recipient').toLowerCase();
		const limit = argValue('--limit') ?? '100';
		const data = await resend(`/emails?limit=${encodeURIComponent(limit)}`, key);
		const emails = Array.isArray(data?.data) ? data.data : [];
		const match = emails.find((email) => toValues(email.to).some((to) => to.includes(recipient)));
		if (!match) throw new Error(`No Resend email found for ${recipient}`);
		printJson(safeEmailSummary(match));
		return;
	}

	if (command === 'link') {
		const id = requiredArg('--id');
		const email = await resend(`/emails/${encodeURIComponent(id)}`, key);
		const links = confirmationLinks(email);
		if (!links.length) throw new Error(`No confirmation link found for Resend email ${id}`);
		printJson({
			...safeEmailSummary(email),
			linkHosts: [...new Set(links.map((link) => new URL(link).host))].sort(),
			confirmationLinks: links,
			tokenHashes: [...new Set(links.map(tokenHash).filter(Boolean))],
		});
		return;
	}

	throw new Error(`Unknown command: ${command}`);
}

async function resend(path, key) {
	const response = await fetch(`${RESEND_API}${path}`, {
		headers: { Authorization: `Bearer ${key}` },
	});
	if (!response.ok) throw new Error(`Resend ${path} failed: HTTP ${response.status}`);
	return response.json();
}

function safeEmailSummary(email) {
	return {
		id: email.id,
		to: email.to,
		from: email.from,
		subject: email.subject,
		created_at: email.created_at,
		last_event: email.last_event,
	};
}

function confirmationLinks(email) {
	const body = decodeHtml(`${email.html ?? ''}\n${email.text ?? ''}`);
	return [...new Set([...body.matchAll(/https?:\/\/[^"\s<>]+/g)].map((match) => match[0].replace(/\]$/, '')))]
		.filter((link) => link.includes('/auth/confirm-signup') || link.includes('token_hash='));
}

function tokenHash(link) {
	try {
		return new URL(link).searchParams.get('token_hash');
	} catch {
		return null;
	}
}

function decodeHtml(value) {
	return value
		.replaceAll('&amp;', '&')
		.replaceAll('&quot;', '"')
		.replaceAll('&#39;', "'")
		.replaceAll('&lt;', '<')
		.replaceAll('&gt;', '>');
}

function toValues(value) {
	return (Array.isArray(value) ? value : [value]).filter(Boolean).map((item) => String(item).toLowerCase());
}

function requiredArg(name) {
	const value = argValue(name);
	if (!value) throw new Error(`${name} is required`);
	return value;
}

function argValue(name) {
	const index = args.indexOf(name);
	return index === -1 ? null : args[index + 1];
}

function printJson(value) {
	console.log(JSON.stringify(value, null, 2));
}
