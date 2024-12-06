data "hetznerdns_nameservers" "primary" {
  type = "authoritative"
}

resource "hetznerdns_record" "mydomain_de-NS" {
  for_each = toset(data.hetznerdns_nameservers.primary.ns.*.name)

  zone_id = hetznerdns_zone.de.id
  name    = "@"
  type    = "NS"
  value   = each.value
}
