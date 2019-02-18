package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secret": {
				Type: schema.TypeString,
				Computed: true,
			},
			"username": {
				Type: schema.TypeString,
				Computed: true,
			},
			"password": {
				Type: schema.TypeString,
				Computed: true,
			},
			"notes": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSecretRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*LastPassProvider)

	name := d.Get("name").(string)

	for _, account := range client.vault.Accounts {
		if account.Name == name {
			d.SetId(account.Id)
			d.Set("name", account.Name)
			d.Set("username", account.Username)
			d.Set("password", account.Password)
			d.Set("notes", account.Notes)
			return nil
		}
	}

	// Resource does not exist
	d.SetId("")
	return nil
}
