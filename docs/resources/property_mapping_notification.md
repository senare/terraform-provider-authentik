---
page_title: "authentik_property_mapping_notification Resource - terraform-provider-authentik"
subcategory: "Customization"
description: |-
  
---

# authentik_property_mapping_notification (Resource)



## Example Usage

```terraform
# Create a custom Notification transport mapping

resource "authentik_property_mapping_notification" "name" {
  name       = "custom-field"
  expression = "return {\"foo\": context['foo']}"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `expression` (String)
- `name` (String)

### Read-Only

- `id` (String) The ID of this resource.


