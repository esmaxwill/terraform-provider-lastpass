package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/mattn/lastpass-go"
	"log"
	"os"
)

type LastPassProvider struct {
	logger   *log.Logger
	username string
	password string
	vault    *lastpass.Vault
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "LastPass Username",
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("LASTPASS_USERNAME"); v != "" {
						return v, nil
					}
					return nil, nil
				},
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "LastPass Password",
				DefaultFunc: func() (interface{}, error) {
					if v := os.Getenv("LASTPASS_PASSWORD"); v != "" {
						return v, nil
					}
					return nil, nil
				},
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"lastpass_secret": dataSourceSecret(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	provider := new(LastPassProvider)

	provider.username = username
	provider.password = password

	vault, err := lastpass.CreateVault(username, password)
	provider.vault = vault

	return provider, err
}
