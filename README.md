# Terraform Provider for Hetzner DNS

[![Tests](https://github.com/germanbrew/terraform-provider-hetznerdns/actions/workflows/build.yaml/badge.svg)](https://github.com/germanbrew/terraform-provider-hetznerdns/actions/workflows/build.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/germanbrew/terraform-provider-hetznerdns)
![GitHub](https://img.shields.io/github/license/germanbrew/terraform-provider-hetznerdns)

**This provider is on published on the Terraform and OpenTofu registry**.

> This project has been forked from [timohirt/terraform-provider-hetznerdns](https://github.com/timohirt/terraform-provider-hetznerdns), which is no longer maintained.

You can find resources and data sources
[documentation](https://registry.terraform.io/providers/germanbrew/hetznerdns/latest/docs)
there or [here](docs).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) > v1.0
- [Go](https://golang.org/) 1.16 (to build the provider plugin)

## Installing and Using this Plugin

You most likely want to download the provider from [Terraform
Registry](https://registry.terraform.io/providers/germanbrew/hetznerdns/latest/docs).

### Migration Guide

If you previously used the `timohirt/hetznerdns` provider, you can easily replace the provider in your terraform state by following these steps:

1. In your `terraform` -> `required_providers` config, replace the provider config:
  ```diff
  hetznerdns = {
  -  source  = "timohirt/hetznerdns"
  +  source = "germanbrew/hetznerdns"
  -  version = "2.2.0"
  +  version = "3.0.0"  # Replace with latest version
  }
  ```
2. Install the new provider and replace it in the state:
  ```sh
    terraform init
    terraform state replace-provider timohirt/hetznerdns germanbrew/hetznerdns
  ```

### Using Provider from Terraform Registry (TF >= 1.0)

This provider is published and available there. If you want to use it, just
add the following to your `terraform.tf`:

```terraform
terraform {
  required_providers {
    hetznerdns = {
      source = "germanbrew/hetznerdns"
      version = "3.0.0"
    }
  }
  required_version = ">= 1.0"
}
```

Then run `terraform init` to download the provider.

## Authentication

Once installed, you have three options to provide the required API token that
is used to authenticate at the Hetzner DNS API.

### Enter API Token when needed

You can enter it every time you run `terraform`.

### Configure the Provider to take the API Token from a Variable

Add the following to your `terraform.tf`:

```terraform
variable "hetznerdns_token" {}

provider "hetznerdns" {
  apitoken = var.hetznerdns_token
}
```

Now, assign your API token to `hetznerdns_token` in `terraform.tfvars`:

```terraform
hetznerdns_token = "kkd993i3kkmm4m4m4"
```

You don't have to enter the API token anymore.

### Inject the API Token via the Environment

Assign the API token to `HETZNER_DNS_API_TOKEN` env variable.

```
export HETZNER_DNS_API_TOKEN=<your api token>
```

The provider uses this token and you don't have to enter it
anymore.

### Example Usage

```terraform
# Specify a zone for a domain (example.com)
resource "hetznerdns_zone" "example_com" {
  name = "example.com"
  ttl  = 60
}

# Handle root (example.com)
resource "hetznerdns_record" "example_com_root" {
  zone_id = hetznerdns_zone.example_com.id
  name    = "@"
  value   = hcloud_server.server_name.ipv4_address
  type    = "A"
  # You only need to set a TTL if it's different from the zone's TTL above
  ttl     = 300
}

# Handle wildcard subdomain (*.example.com)
resource "hetznerdns_record" "all_example_com" {
  zone_id = hetznerdns_zone.example_com.id
  name    = "*"
  value   = hcloud_server.server_name.ipv4_address
  type    = "A"
}

# Handle specific subdomain (books.example.com)
resource "hetznerdns_record" "books_example_com" {
  zone_id = hetznerdns_zone.example_com.id
  name    = "books"
  value   = hcloud_server.server_name.ipv4_address
  type    = "A"
}

# Handle email (MX record with priority 10)
resource "hetznerdns_record" "example_com_email" {
  zone_id = hetznerdns_zone.example_com.id
  name    = "@"
  value   = "10 mail.example.com"
  type    = "MX"
}

# SPF record
resource "hetznerdns_record" "example_com_spf" {
  zone_id = hetznerdns_zone.example_com.id
  name    = "@"
  # The entire value needs to be enclosed in quotes in the zone file, if it contains a space or a quote. For Terraform, you need to escape these "inner" quotes:
  value   = "\"v=spf1 ip4:1.2.3.4 -all\""
  # Or let `jsonencode()` take care of the escaping:
  value   = jsonencode("v=spf1 ip4:1.2.3.4 -all")
  type    = "TXT"
}
```
