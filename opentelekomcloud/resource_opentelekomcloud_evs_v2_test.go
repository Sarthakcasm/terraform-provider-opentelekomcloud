package opentelekomcloud

import (
"fmt"
"testing"

"github.com/hashicorp/terraform/helper/resource"
"github.com/hashicorp/terraform/terraform"

//"github.com/gophercloud/gophercloud/openstack/networking/v1/vpcs"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/evs"
)

// PASS
func TestAccOTCEvsV2_basic(t *testing.T) {
	var evs evs.EVS

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCEvsV2Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccEvsV2_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCEvsV2Exists("opentelekomcloud_evs_v2.evs_1", &evs),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "name", "EVS_getMade"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "size", "10"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "volume_type", "SSD"),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "description", "test_holo"),
				),
			},
		},
	})
}

func TestAccOTCEvsV2_update(t *testing.T) {
	var evs evs.EVS

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCEvsV2Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccEvsV2_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCEvsV2Exists("opentelekomcloud_evs_v2.evs_1", &evs),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "name", "EVS_getMade"),
				),
			},
			resource.TestStep{
				Config: testAccEvsV2_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCEvsV2Exists("opentelekomcloud_evs_v2.evs_1", &evs),
					resource.TestCheckResourceAttr(
						"opentelekomcloud_evs_v2.evs_1", "name", "EVS_getMade1"),
				),
			},
		},
	})
}

// PASS
func TestAccOTCEvsV2_timeout(t *testing.T) {
	var evs evs.EVS

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOTCEvsV2Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccEvsV2_timeout,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOTCEvsV2Exists("opentelekomcloud_evs_v2.evs_1", &evs),
				),
			},
		},
	})
}

func testAccCheckOTCEvsV2Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	evsClient, err := config.evsV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud evs client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_evs_v2" {
			continue
		}

		_, err := evs.Get(evsClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Evs still exists")
		}
	}

	return nil
}

func testAccCheckOTCEvsV2Exists(n string, evss *evs.EVS) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		evsClient, err := config.evsV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating OpenTelekomCloud evs client: %s", err)
		}
		
		found, err := evs.Get(evsClient, rs.Primary.ID).Extract()

		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("evs not found")
		}

		*evss = *found

		return nil
	}
}

const testAccEvsV2_basic = `
resource "opentelekomcloud_evs_v2" "evs_1" {
	availability_zone="eu-de-02"
    description="test_holo"
    size=10
    name="EVS_getMade"
    volume_type="SSD"
}
`

const testAccEvsV2_update = `
resource "opentelekomcloud_evs_v2" "evs_1" {
   availability_zone="eu-de-02"
  description="test_holo"
  size=10
  name="EVS_getMade1"

  volume_type="SSD"
}
`
const testAccEvsV2_timeout = `
resource "opentelekomcloud_evs_v2" "evs_1" {
	availability_zone="eu-de-02"
  description="test_holo"
  size=10
  name="EVS_getMade"

  volume_type="SSD"

  timeouts {
    create = "5m"
    delete = "5m"
  }
}
`
