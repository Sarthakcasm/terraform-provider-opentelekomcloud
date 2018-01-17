package opentelekomcloud

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"time"
)

func TestAccOTCVpcSubnetV1DataSource_basic(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOTCVpcSubnetV1Config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCVpcSubnetV1DataSourceID("data.opentelekomcloud_subnet_v1.by_id"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "opentelekomcloud_subnet"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "cidr", "192.168.0.0/16"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "gateway_ip", "192.168.0.1"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "availability_zone", "eu-de-02"),
				),
			},
		},
	})
}

func testAccCheckOTCVpcSubnetV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find security group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Security group data source ID not set")
		}

		return nil
	}
}

const testAccDataSourceOTCVpcSubnetV1Config  = `
resource "opentelekomcloud_vpc_v1" "vpc_1" {
	name = "test_vpc"
	cidr= "192.168.0.0/16"
}

resource "opentelekomcloud_subnet_v1" "subnet_1" {
  name = "opentelekomcloud_subnet"
  cidr = "opentelekomcloud_subnet"
  gateway_ip = "192.168.0.1"
  vpc_id = "${opentelekomcloud_vpc_v1.vpc_1.id}"
  availability_zone = "eu-de-02"
 }

data "opentelekomcloud_subnet_v1" "by_id" {
  id = "${opentelekomcloud_subnet_v1.subnet_1.id}"
}

data "opentelekomcloud_subnet_v1" "by_cidr" {
  cidr = "${opentelekomcloud_subnet_v1.subnet_1.cidr}"
}

data "opentelekomcloud_subnet_v1" "by_name" {
	name = "${opentelekomcloud_subnet_v1.subnet_1.name}"
}

data "opentelekomcloud_subnet_v1" "by_vpc_id" {
	vpc_id = "${opentelekomcloud_subnet_v1.subnet_1.vpc_id}"
}
`
