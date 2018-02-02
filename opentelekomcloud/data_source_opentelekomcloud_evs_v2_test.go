package opentelekomcloud

import (
	"fmt"
	//"math/rand"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	//"time"
)

func TestAccOTCEvsV2DataSource_basic(t *testing.T) {

	//backup_id := fmt.Sprint("null")
	//count := fmt.Sprint("1")
	availability_zone := fmt.Sprintf("eu-de-02")
	description := fmt.Sprint("test_heloo")
	size := fmt.Sprint("10")
	name := fmt.Sprintf("EVS_getMade")
//	imageRef := fmt.Sprint("null")
	volume_type := fmt.Sprintf("SSD")


	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOTCEvsV2Config(availability_zone,description,size,name,volume_type),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceOTCEvsV2Check("data.opentelekomcloud_evs_v2.by_id" ),
					testAccDataSourceOTCEvsV2Check("data.opentelekomcloud_evs_v2.by_name"),
					testAccDataSourceOTCEvsV2Check("data.opentelekomcloud_evs_v2.by_status"),
					testAccDataSourceOTCEvsV2Check("data.opentelekomcloud_evs_v2.by_availability_zone"),

					testAccDataSourceOTCEvsV2Check("data.opentelekomcloud_evs_v2.by_volume_type"),

					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_evs_v2.by_id", "shareable", "false"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_evs_v2.by_id", "status", "in-use"),
					),
			},
		},
	})

}


// TestCheckFunc is the callback type used with acceptance tests to check
// the state of a resource. The state passed in is the latest state known,
// or in the case of being after a destroy, it is the last known state when
// it was created.


func testAccDataSourceOTCEvsV2Check(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", n)
		}

		evsRs, ok := s.RootModule().Resources["opentelekomcloud_evs_v2.evs_1"]
		if !ok {
			return fmt.Errorf("can't find opentelekomcloud_evs_v2.evs_1 in state")
		}

		attr := rs.Primary.Attributes

		if attr["id"] != evsRs.Primary.Attributes["id"] {
			return fmt.Errorf(
				"id is %s; want %s",
				attr["id"],
				evsRs.Primary.Attributes["id"],
			)
		}

		if attr["name"] != evsRs.Primary.Attributes["name"] {
			return fmt.Errorf(
				"name is %s; want %s",
				attr["name"],
				evsRs.Primary.Attributes["name"],
			)
		}

		if attr["status"] != evsRs.Primary.Attributes["status"] {
			return fmt.Errorf(
				"status is %s; want %s",
				attr["status"],
				evsRs.Primary.Attributes["status"],
			)
		}

		if attr["availability_zone"] != evsRs.Primary.Attributes["availability_zone"] {
			return fmt.Errorf(
				"availability_zone is %s; want %s",
				attr["availability_zone"],
				evsRs.Primary.Attributes["availability_zone"],
			)
		}

		if attr["volume_type"] != evsRs.Primary.Attributes["volume_type"] {
			return fmt.Errorf(
				"volume_type is %s; want %s",
				attr["volume_type"],
				evsRs.Primary.Attributes["volume_type"],
			)
		}
		return nil
	}
}

func testAccDataSourceOTCEvsV2Config(availability_zone,description,size,name,volume_type string) string {
	return fmt.Sprintf(`
resource "opentelekomcloud_evs_v2" "evs_1" {
availability_zone = "%s",
description = "%s",
size = "%d",
name = "%s",
volume_type = "%s"

}

data "opentelekomcloud_evs_v2" "by_id" {
  id = "${opentelekomcloud_evs_v2.evs_1.id}"
}

data "opentelekomcloud_evs_v2" "by_name" {
  name = "${opentelekomcloud_evs_v2.evs_1.name}"
}

data "opentelekomcloud_evs_v2" "by_status" {
	status = "${opentelekomcloud_evs_v2.evs_1.status}"
}

data "opentelekomcloud_evs_v2" "by_availability_zone" {
	availability_zone = "${opentelekomcloud_evs_v2.evs_1.availability_zone}"
}

data "opentelekomcloud_evs_v2" "by_volume_type" {
	volume_type = "${opentelekomcloud_evs_v2.evs_1.volume_type}"
}

`,availability_zone,description,size,name,volume_type)
}
