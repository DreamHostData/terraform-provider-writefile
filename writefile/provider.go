package writefile

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"writefile_file": fileResource(),
		},
		ConfigureFunc: writefileConfigure,
	}
}

func writefileConfigure(data *schema.ResourceData) (interface{}, error) {
	config := &writefileConfig{}
	return config, nil
}

type writefileConfig struct {
}
