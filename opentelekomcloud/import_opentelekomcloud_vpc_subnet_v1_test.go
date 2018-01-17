package opentelekomcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// PASS
func TestAccSubnetV1_importBasic(t *testing.T) {
	resourceName := "opentelekomcloud_subnet_v1.subnet_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
