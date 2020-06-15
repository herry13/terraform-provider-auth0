---
layout: "auth0"
page_title: "Auth0: auth0_hook_secret"
description: |-
  Hooks feature integrated secret management to securely store secrets while making them conveniently available in code.
---

# auth0_hook_secret

Hooks feature integrated secret management to securely store secrets while making them conveniently available in code.

## Example Usage

```
resource "auth0_hook" "my_hook" {
  name = "My Pre User Registration Hook"
  script = <<EOF
function (user, context, callback) { 
  callback(null, { user }); 
}
EOF
  trigger_id = "pre-user-registration"
  enabled = true
}

resource "auth0_hook_secret" "my_secret" {
  hook_id = auth0_hook.my_hook.id
  name = "secret_name"
  value = "secret_value"
}
```

## Argument Reference

The following arguments are supported:

* `hook_id` - (Required) Hook ID which the secret is associated with
* `name` - (Required) Secret name
* `script` - (Required) Secret value