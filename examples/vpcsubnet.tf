resource "opentelekomcloud_vpc_v1" "vpc_v1" {
  name = "opentelekomcloud_vpc"
  cidr = "192.168.0.0/16"
}

resource "opentelekomcloud_subnet_v1" "subnet_1" {
  name = "${var.project}-subnet${format("%02d", count.index+1)}"
  cidr = "${var.subnet_cidr}"
  gateway_ip = "${var.subnet_gateway_ip}"
  vpc_id = "${opentelekomcloud_vpc_v1.vpc_v1.id}"
  availability_zone = "${var.subnet_availability_zone}"

}