package opentelekomcloud

import (
"testing"

"github.com/hashicorp/terraform/helper/resource"
)

// PASS
func TestAccOTCEvsV2_importBasic(t *testing.T) {
	resourceName := "opentelekomcloud_evs_v2.evs_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCEvsV2Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccEvsV2_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
