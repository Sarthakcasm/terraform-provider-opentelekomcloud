---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_subnet_v1"
sidebar_current: "docs-opentelekomcloud-resource-subnet-v1"
description: |-
  Manages a V1 Subnet resource within OpenTelekomCloud.
---

# opentelekomcloud_subnet_v1

Manages a Subnet resource within OpenTelekomCloud.

## Example Usage

```hcl
    resource "opentelekomcloud_vpc_v1" "vpc_v1" {
      name = "test_vpc"
      cidr = "$192.168.0.0/16"
    }

    variable "subnet_name" {
      default = "opentelekomcloud_subnet"
    }

    variable "subnet_cidr" {
      default = "192.168.0.0/16"
    }

    variable "subnet_gateway_ip" {
      default = "192.168.0.1"
    }

    variable "subnet_availability_zone" {
      default = "eu-de-02"
    }


    resource "opentelekomcloud_subnet_v1" "subnet_v1" {
      name = "${var.subnet_name}"
      cidr = "${var.subnet_cidr}"
      gateway_ip = "${var.subnet_gateway_ip}"
      vpc_id = "${opentelekomcloud_vpc_v1.vpc_v1.id}"
      availability_zone = "${var.subnet_availability_zone}"

    }

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the Subnet. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).

* `cidr` - (Required) The cidr block for desired Subnet.

* `gateway_ip` - (Required) Specifies the subnet gateway address.

* `availability_zone` -(Required) Specifies the ID of the AZ to which the subnet belongs.

* `vpc_id` -(Required) Specifies the ID of the VPC to which the subnet belongs.

* `primary_dns` -(Optional) Specifies the IP address of DNS server 1 on the subnet.

* `secondary_dns` -(Optional) Specifies the IP address of DNS server 2 on the subnet.

* `dhcp_enable` -(Optional) Specifies whether the DHCP function is enabled for the subnet.



## Attributes Reference

The following attributes are exported:

* `id` -  ID of the Subnet.

* `name` -  See Argument Reference above.

* `cidr` - See Argument Reference above.

* `status` -(Optional) Specifies the status of the subnet.The value can be ACTIVE, DOWN, UNKNOW, or ERROR.

* `gateway_ip` - See Argument Reference above.

* `availability_zone` - See Argument Reference above.

* `vpc_id` - See Argument Reference above.

* `primary_dns` - See Argument Reference above.

* `secondary_dns` - See Argument Reference above.

* `dhcp_enable` - See Argument Reference above.

## Import

Subnets can be imported using the `id`, e.g.

```
$ terraform import opentelekomcloud_subnet_v1.subnet_v1 7117d38e-4c8f-4624-a505-bd96b97d024c
```
