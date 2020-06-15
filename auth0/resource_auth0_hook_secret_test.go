package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHookSecret(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccHookSecretCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "hook_id", "hook-id"),
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "name", "name"),
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "value", "1"),
				),
			},
			{
				Config: testAccHookSecretUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "hook_id", "hook-id"),
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "name", "name"),
					resource.TestCheckResourceAttr("auth0_hook_secret.my_secret", "value", "2"),
				),
			},
		},
	})
}

const testAccHookSecretCreate = `

resource "auth0_hook_secret" "my_secret" {
  hook_id = "hook-id"
  name = "name"
  value = "1"
}
`

const testAccHookSecretUpdate = `

resource "auth0_hook_secret" "my_secret" {
	hook_id = "hook-id"
	name = "name"
	value = "2"
  }
`
