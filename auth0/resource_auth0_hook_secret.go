package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func newHookSecret() *schema.Resource {
	return &schema.Resource{

		Create: createHookSecret,
		Read:   readHookSecret,
		Update: updateHookSecret,
		Delete: deleteHookSecret,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"hook_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Hook ID which this hook secret is associated with",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of this hook secret",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Code to be executed when this hook runs",
			},
		},
	}
}

func createHookSecret(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	hookID := d.Get("hook_id").(string)
	name := d.Get("name").(string)
	if exist, err := hookSecretIsExist(hookID, name, api); err != nil {
		return err
	} else if !exist {
		value := d.Get("value").(string)
		secrets := make(management.HookSecrets)
		secrets[name] = value
		if err := api.Hook.CreateSecrets(hookID, &secrets); err != nil {
			return err
		}
	}
	return readHook(d, m)
}

func readHookSecret(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	hookID := d.Get("hook_id").(string)
	name := d.Get("name").(string)
	if exist, err := hookSecretIsExist(hookID, name, api); err != nil {
		return err
	} else if exist {
		value := d.Get("value")
		d.Set("value", value)
	} else {
		d.SetId("")
	}
	return nil
}

func updateHookSecret(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	hookID := d.Get("hook_id").(string)
	name := d.Get("name").(string)
	if exist, err := hookSecretIsExist(hookID, name, api); err != nil {
		return err
	} else if exist {
		value := d.Get("value").(string)
		secrets := make(management.HookSecrets)
		secrets[name] = value
		if err := api.Hook.UpdateSecrets(hookID, &secrets); err != nil {
			return err
		}
	}
	return readHook(d, m)
}

func deleteHookSecret(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	hookID := d.Get("hook_id").(string)
	name := d.Get("name").(string)
	if exist, err := hookSecretIsExist(hookID, name, api); err != nil {
		return err
	} else if exist {
		return api.Hook.RemoveSecrets(hookID, name)
	}
	return nil
}

func hookSecretIsExist(hookID, secretName string, api *management.Management) (bool, error) {
	hookSecrets, err := api.Hook.Secrets(hookID)
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				return false, nil
			}
		}
		return false, err
	}
	for _, key := range hookSecrets.Keys() {
		if secretName == key {
			return true, nil
		}
	}
	return false, nil
}
