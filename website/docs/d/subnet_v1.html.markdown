---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_subnet_v1"
sidebar_current: "docs-opentelekomcloud-datasource-subnet-v1"
description: |-
  Get information on an OpenTelekomCloud VPC Subnet.
---

# opentelekomcloud_subnet_v1

opentelekomcloud_subnet_v1 provides details about a specific Subnet.

This resource can prove useful when a module accepts a subnet id as an input variable and needs to, for example, determine the CIDR block of that Subnet.

## Example Usage

The following example shows how one might accept a Subnet id as a variable and use this data source to obtain the vpc_id on which subnet exits .

```hcl

variable "subnet_name" {
  default =""
}

data "opentelekomcloud_subnet_v1" "subnet" {
  name = "${var.subnet_name}"
}

```

## Argument Reference

The arguments of this data source act as filters for querying the available Subnetss in the current region. The given filters must match exactly one Subnet whose data will be exported as attributes.

* `name` - (Optional) A unique name for the Subnet. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).

* `id` - (Optional) The id of the specific Subnet to retrieve.

* `cidr` - (Optional) The cidr block for desired Subnet.

* `gateway_ip` - (Optional) Specifies the subnet gateway address.

* `availability_zone` -(Optional) Specifies the ID of the AZ to which the subnet belongs.

* `vpc_id` -(Optional) Specifies the ID of the VPC to which the subnet belongs.

* `status` -(Optional) Specifies the status of the subnet.The value can be ACTIVE, DOWN, UNKNOW, or ERROR.






## Attributes Reference

The following attributes are exported:

* `id` - ID of the Subnet.

* `name` -  See Argument Reference above.

* `status` - See Argument Reference above.

* `cidr` - See Argument Reference above.

* `vpc_id` - See Argument Reference above.

* `primary_dns` - Specifies the IP address of DNS server 1 on the subnet.

* `secondary_dns` - Specifies the IP address of DNS server 2 on the subnet.

* `availability_zone` - Specifies the ID of the AZ to which the subnet belongs.

* `dhcp_enable` - Specifies whether the DHCP function is enabled for the subnet.


