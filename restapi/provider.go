package restapi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"net/url"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kafka schema registry endpoint. Example: http://localhost:8000",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"schemaregistry_subject": resourceSubject(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("uri").(string)
	if _, err := url.ParseRequestURI(endpoint); err != nil {
		return endpoint, diag.Diagnostics{
			{
				Severity:      0,
				Summary:       err.Error(),
				Detail:        "",
				AttributePath: nil,
			},
		}
	}

	return endpoint, nil
}
