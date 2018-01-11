package opentelekomcloud

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"time"
)

func TestAccOTCSubnetV1DataSource_basic(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	rInt := rand.Intn(50)
	cidr := fmt.Sprintf("172.16.%d.0/24", rInt)
	name := fmt.Sprintf("terraform-testacc-subnet-data-source-%d", rInt)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOTCSubnetV1Config(name, cidr),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceOTCSubnetV1Check("data.opentelekomcloud_subnet_v1.by_id", name, cidr),
					testAccDataSourceOTCSubnetV1Check("data.opentelekomcloud_subnet_v1.by_name", name, cidr),
					testAccDataSourceOTCSubnetV1Check("data.opentelekomcloud_subnet_v1.by_vpc_id", name, cidr),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_subnet_v1.by_id", "gateway_ip", "192.168.0.1"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_subnet_v1.by_id", "availability_zone", "eu-de-02"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_subnet_v1.by_id", "status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_subnet_v1.by_id", "primary_dns", "114.114.114.114"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_subnet_v1.by_id", "secondary_dns", "114.114.115.115"),
				),
			},
		},
	})
}

func testAccDataSourceOTCSubnetV1Check(n, name, cidr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", n)
		}

		subnetRs, ok := s.RootModule().Resources["opentelekomcloud_subnet_v1.subnet_1"]
		if !ok {
			return fmt.Errorf("can't find opentelekomcloud_vpc_v1.subnet_1 in state")
		}

		attr := rs.Primary.Attributes

		if attr["id"] != subnetRs.Primary.Attributes["id"] {
			return fmt.Errorf(
				"id is %s; want %s",
				attr["id"],
				subnetRs.Primary.Attributes["id"],
			)
		}

		if attr["cidr"] != cidr {
			return fmt.Errorf("bad subnet cidr %s, expected: %s", attr["cidr"], cidr)
		}
		if attr["name"] != name {
			return fmt.Errorf("bad subnet name %s", attr["name"])
		}

		return nil
	}
}

func testAccDataSourceOTCSubnetV1Config(name, cidr string) string {
	return fmt.Sprintf(`
rresource "opentelekomcloud_subnet_v1" "subnet_1" {
  cidr = "192.168.0.0/16"
  name = "test_subnet1"
  gateway_ip = "192.168.0.1"
  vpc_id ="8f794f06-2275-4d82-9f5a-6d68fbe21a75"
  availability_zone ="eu-de-02"
}
}

data "opentelekomcloud_subnet_v1" "by_id" {
  id = "${opentelekomcloud_subnet_v1.subnet_1.id}"
}

data "opentelekomcloud_subnet_v1" "by_name" {
  cidr = "${opentelekomcloud_subnet_v1.subnet_1.name}"
}

data "opentelekomcloud_subnet_v1" "by_vpc_id" {
	name = "${opentelekomcloud_subnet_v1.subnet_1.vpc_id}"
}
`, name, cidr)
}
