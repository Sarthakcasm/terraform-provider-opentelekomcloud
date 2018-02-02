package evs


import "github.com/gophercloud/gophercloud"

//const resourcePath = "/cloudvolumes/detail"

func EvsListURL(c *gophercloud.ServiceClient) string {

		resourcePath := "cloudvolumes/detail"
		return c.ServiceURL(resourcePath)

}

// for creating EVS
func EvsURLCreateEvs(c *gophercloud.ServiceClient) string {
	resourcePath := "cloudvolumes"
	return c.ServiceURL(resourcePath)
}
//for update EVS
func EvsURLupdate(c *gophercloud.ServiceClient,id string) string {
	resourcePath := "cloudvolumes"
	return c.ServiceURL(resourcePath,id)
}

func EVSresourceURL(c *gophercloud.ServiceClient, id string) string {
	resourcePath := "volumes"
	return c.ServiceURL(resourcePath, id)
}

func EVSDeleteURL(c *gophercloud.ServiceClient, id string) string {
	resourcePath := "cloudvolumes"
	return c.ServiceURL(resourcePath, id)
}

// for EVS size updation
func EvsURLSizeUpdate(c *gophercloud.ServiceClient,id  string ) string {
	resourcePathSizeUpdate := "cloudvolumes"
	action := "action"
	return c.ServiceURL(resourcePathSizeUpdate,id,action)
}