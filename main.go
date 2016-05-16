package main

import (
	// "github.com/DreamHostData/terraform-provider-writefile/writefile"
	"./writefile"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: writefile.Provider,
	})
}
