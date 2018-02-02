# Configure the OpenStack Provider
provider "opentelekomcloud" {
  user_name   = "lizhonghua"
  domain_name = "OTC00000000001000010501"
  password    = "slob@123"
  auth_url    = "https://iam.eu-de.otc.t-systems.com/v3"
  region      = "eu-de"
  tenant_id = "87a56a48977e42068f70ad3280c50f0e"
//  tenant_name  ="eu-de_Nordea"
}

variable "name"{
  default="ecs-2118-sunway-jumper-nordea"
}

resource"opentelekomcloud_evs_v2" "evs_v1"{

  availability_zone="eu-de-02"
  description="test_holo1"
  size=10
  name="EVS_getMade"

  volume_type="SSD"
}

data "opentelekomcloud_evs_v2" "selected" {
  //status = "normal"
  name = "${var.name}"

}


output "myOutput" {
  value = "${data.opentelekomcloud_evs_v2.selected.links}"
}






