package opentelekomcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"
)

// PASS
func TestAccOTCSubnetV1_basic(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "terraform_provider_test"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "cidr", "192.168.0.0/16"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "gateway_ip", "192.168.0.1"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "vpc_id", "8f794f06-2275-4d82-9f5a-6d68fbe21a75"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "availability_zone", "eu-de-02"),
				),
			},
		},
	})
}

func TestAccOTCSubnetV1_update(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "terraform_provider_test"),
				),
			},
			resource.TestStep{
				Config: testAccSubnetV1_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_subnet_v1.subnet_1", "name", "terraform_provider_test1"),
				),
			},
		},
	})
}

// PASS
func TestAccOTCSubnetV1_timeout(t *testing.T) {
	var subnet subnets.Subnet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCSubnetV1Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSubnetV1_timeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCSubnetV1Exists("opentelekomcloud_subnet_v1.subnet_1", &subnet),
				),
			},
		},
	})
}

func testAccCheckOTCSubnetV1Destroy(s *terraform.State) error {
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

func testAccCheckOTCSubnetV1Exists(n string, vpc *subnets.Subnet) resource.TestCheckFunc {
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
			return fmt.Errorf("Error creating OpenTelekomCloud vpc client: %s", err)
		}

		found, err := subnets.Get(subnetClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("subnet not found")
		}

		*vpc = *found

		return nil
	}
}

const testAccSubnetV1_basic = `
resource "opentelekomcloud_subnet_v1" "subnet_1" {
  cidr = "192.168.0.0/16"
  name = "terraform_provider_test"
  gateway_ip = "192.168.0.1"
  vpc_id ="8f794f06-2275-4d82-9f5a-6d68fbe21a75"
  availability_zone ="eu-de-02"
}
`

const testAccSubnetV1_update = `
resource "opentelekomcloud_subnet_v1" "selected" {
  cidr = "192.168.0.0/16"
  name = "terraform_provider_test1"
  gateway_ip = "192.168.0.1"
  vpc_id ="8f794f06-2275-4d82-9f5a-6d68fbe21a75"
  availability_zone ="eu-de-02"
}
`
const testAccSubnetV1_timeout = `
resource "opentelekomcloud_subnet_v1" "selected" {
  cidr = "192.168.0.0/16"
  name = "terraform_provider_test1"
  gateway_ip = "192.168.0.1"
  vpc_id ="8f794f06-2275-4d82-9f5a-6d68fbe21a75"
  availability_zone ="eu-de-02"

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
