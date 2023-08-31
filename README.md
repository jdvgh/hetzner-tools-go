# Hetzner-Tools-Go

![version](https://img.shields.io/github/v/release/jdvgh/hetzner-tools-go?include_prereleases)
[![Apache License v2](https://img.shields.io/github/license/jdvgh/hetzner-tools-go)](http://www.apache.org/licenses/)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jdvgh/hetzner-tools-go)](https://github.com/jdvgh/hetzner-tools-go)

This is just a personal project for helper functions to talk to the Hetzner-API [Hetzner](https://www.hetzner.com/).

# Motivation

My personal use case is that I want to have my personal domain point to a new IP without having to do any manual steps e.g.:

1. Deploy a 1-server-k3s-cluster by using [terraform-hcloud-kube-hetzner](https://github.com/kube-hetzner/terraform-hcloud-kube-hetzner)
1. Retrieve the public IPv4 of the server
1. Retrieve all DNS-Records (of type A) for the respective Zone
1. Update all my DNS-Record of the DNS-Zone to point to the new IP

# Underneath

It mainly uses

- Hetzner DNS API [Hetzner DNS API](https://dns.hetzner.com/api-docs)
- Hetzner HCloud Go API [Hetzner HCLoud Go](https://github.com/hetznercloud/hcloud-go)
- Cobra-CLI to generate the CLI [Cobra-CLI](https://github.com/spf13/cobra-cli/blob/main/README.md)
- My personal bootstrapping project (which uses terraform-hcloud-kube-hetzner - as mentioned above in the underlyings) [hetzner-k3s-bootstrap](https://github.com/jdvgh/hetzner-k3s-bootstrap)

# How to run

## Get Credentials

Log into your Hetzner Account and get a DNS API key and an HCloud API Key.
You can either specify them in a local `.env` file, or via the CLI arguments e.g.:

- DNS API Key Retrieval [Hetzner DNS API Key Docs](https://docs.hetzner.com/dns-console/dns/general/api-access-token/)
- Hcloud API Key [Hcloud API Key Docs](https://docs.hetzner.com/cloud/api/getting-started/generating-api-token/)

### Using .env file

```.env
HCLOUD_API_TOKEN="YOUR_HCLOUD_API_TOKEN_HERE"
HETZNER_DNS_API_TOKEN="YOUR_HETZNER_DNS_API_TOKEN_HERE"
HETZNER_DNS_ZONE_ID="example.com"
```

### Using CLI arguments

```
go run ./hetzner-tools-go/main.go  dns list --dns-api-token="YOUR_HCLOUD_API_TOKEN_HERE" --hcloud-dns-zone-id="example.com"

```

## Clone the repo and use go run

1. Change the directory to the root of the repository
1. Run the wanted command e.g. `go run hetzner-tools-go/main.go server list`

## Installing the command

```
go install github.com/jdvgh/hetzner-tools-go/hetzner-tools-go@main
```
