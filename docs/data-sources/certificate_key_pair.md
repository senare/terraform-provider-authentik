---
page_title: "authentik_certificate_key_pair Data Source - terraform-provider-authentik"
subcategory: "System"
description: |-
  Get certificate-key pairs by name
---

# authentik_certificate_key_pair (Data Source)

Get certificate-key pairs by name

## Example Usage

```terraform
# To get the the ID and other info about a certificate

data "authentik_certificate_key_pair" "generated" {
  name = "authentik Self-signed Certificate"
}

# Then use `data.authentik_certificate_key_pair.generated.id`
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `fetch_certificate` (Boolean) If set to true, certificate data will be fetched. Defaults to `true`.
- `fetch_key` (Boolean) If set to true, private key data will be fetched. Defaults to `true`.
- `key_data` (String, Sensitive) Generated.

### Read-Only

- `certificate_data` (String) Generated.
- `expiry` (String) Generated.
- `fingerprint1` (String) SHA1-hashed certificate fingerprint Generated.
- `fingerprint256` (String) SHA256-hashed certificate fingerprint Generated.
- `id` (String) Generated.
- `subject` (String) Generated.


