package main

import (
	"github.com/Sarthakcasm/terraform-provider-opentelekomcloud/opentelekomcloud" // TODO: Revert path when merge
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opentelekomcloud.Provider})
}
