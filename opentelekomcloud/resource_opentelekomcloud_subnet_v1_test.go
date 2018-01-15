package opentelekomcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"
)

// PASS
func TestAccSubnetV1_basic(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
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

func TestAccSubnetV1_update(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "opentelekomcloud_subnet"),

				),
			},
			resource.TestStep{
				Config: testAccSubnetV1_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "opentelekomcloud_subnet_1"),

				),
			},
		},
	})
}

// PASS
func TestAccSubnetV1_timeout(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_timeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
				),
			},
		},
	})
}

func testAccCheckSubnetV1Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	subnetClient, err := config.subnetV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud vpc client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_subnet_v1" {
			continue
		}

		_, err := subnets.Get(subnetClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Subnet still exists")
		}
	}

	return nil
}

func testAccCheckSubnetV1Exists(n string, subnet *subnets.Subnet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		subnetClient, err := config.subnetV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating OpenTelekomCloud Vpc client: %s", err)
		}

		found, err := subnets.Get(subnetClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Subnet not found")
		}

		*subnet = *found

		return nil
	}
}

const testAccSubnetV1_basic = `
resource "opentelekomcloud_vpc_v1" "vpc_1" {
  name = "vpc_1"
  cidr = "192.168.0.0/16"
}

resource "opentelekomcloud_subnet_v1" "subnet_1" {
  name = "opentelekomcloud_subnet"
  cidr = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id = "${opentelekomcloud_vpc_v1.vpc_v1.id}"
  availability_zone = "eu-de-02"

}
`

const testAccSubnetV1_update = `
resource "opentelekomcloud_vpc_v1" "vpc_1" {
  name = "vpc_1"
  cidr = "192.168.0.0/16"
}

resource "opentelekomcloud_subnet_v1" "subnet_1" {
  name = "opentelekomcloud_subnet_1"
  cidr = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id = "${opentelekomcloud_vpc_v1.vpc_v1.id}"
  availability_zone = "eu-de-02"

}
`

const testAccSubnetV1_timeout = `
rresource "opentelekomcloud_vpc_v1" "vpc_1" {
  name = "vpc_1"
  cidr = "192.168.0.0/16"
}

resource "opentelekomcloud_subnet_v1" "subnet_1" {
  name = "opentelekomcloud_subnet_1"
  cidr = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id = "${opentelekomcloud_vpc_v1.vpc_v1.id}"
  availability_zone = "eu-de-02"

}

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
