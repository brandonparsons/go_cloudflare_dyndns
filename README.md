Cloudflare Dynamic DNS
====================

Gets the current IP and updates Cloudflare DNS entries, respecting configurable service levels.


Install
-------

- `go get`
- `go run cmd/runner/dns_updater.go`
- `go install cmd/runner/dns_updater.go`

Usage
-----

`crontab`:

`30 2 * * *  /usr/local/bin/dns_updater --api-key=YOUR_API_KEY_HERE --api-email=YOUR_ACCOUNT_EMAIL_HERE --basedomain brandonparsons.me --domains pi,vault,static --servicelevels 0,1,1 >/dev/null 2>&1`

Options
-------

- api-key: String, required
- api-email: String, required
- basedomain: String, required: Main domain you have registered on cloudflare
- domains: Comma-separated list of strings, required. Which subdomains should be updated?
- servicelevels: Comma-separated list of 0|1, required. What service level should each domain have (read in the same order as `domains`.  0 = grey gloud, 1 = yellow cloud)
- force: Boolean flag, optional.  Whether or not to update even if the cached IP matches current.
