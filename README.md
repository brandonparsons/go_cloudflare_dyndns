Cloudflare Dynamic DNS
====================

Gets the current IP and updates Cloudflare DNS entries.  Hardcoded to update the following domains: `pi` and `vault`.

Usage
-----

`crontab`:

30 2 * * *  /usr/local/bin/dns_updater --api-key=<YOUR API KEY HERE> --api-email=<YOUR ACCOUNT EMAIL HERE> >/dev/null 2>&1

Options
-------

- api-key: String, required
- api-email: String, required
- force: Boolean flag, optional.  Whether or not to update even if the cached IP matches current.
