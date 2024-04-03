# Tailflare

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/aaanh/tailflare/ci.yaml)

Sync your **Tail**scale devices to Cloud**flare** DNS.

The functionality is based on this documentation on Tailscale: https://tailscale.com/kb/1054/dns?q=subdomain#using-a-public-dns-subdomain

It is basically taking the Tailscale IP addresses and put them under a subdomain A record on the DNS provider, which is Cloudflare in our case.
