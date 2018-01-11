package subnets

import "github.com/gophercloud/gophercloud"

const resourcePath = "subnets"

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func updateURL(c *gophercloud.ServiceClient, vpcid, id string) string {
	return c.ServiceURL("vpcs", vpcid, resourcePath, id)
}