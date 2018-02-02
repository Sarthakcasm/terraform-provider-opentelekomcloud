---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_evs_v2"
sidebar_current: "docs-opentelekomcloud-resource-evs-v2"
description: |-
  Manages a V2 EVS resource within OpenTelekomCloud.
---

# opentelekomcloud_evs_v2

Manages a EVS resource within OpenTelekomCloud.

## Example Usage

```hcl

variable "evs_name" {
  default = "opentelekomcloud_vpc"
}

variable "evs_availability_zone" {
  default = "eu-de-02"
}

variable "evs_description" {
  default = "test_holo"
}

variable "evs_size" {
  default = "10"
}

variable "evs_volume_type" {
  default = "SSD"
}

resource "opentelekomcloud_vpc_v1" "vpc_v1" {
  name = "${var.evs_name}"
  availability_zone = "${var.evs_availability_zone}"
  description = "${var.description}"
  size = "${var.evs_size}"
  volume_type = "${var.evs_volume_type}"
}

```

## Argument Reference
//decription of attributes is pending.
The following arguments are supported:

* `name` - (Required) The range of available subnets in the VPC. The value ranges from 10.0.0.0/8 to 10.255.255.0/24, 172.16.0.0/12 to 172.31.255.0/24, or 192.168.0.0/16 to 192.168.255.0/24.

* `availability_zone` - (Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to create a VPC. If omitted, the region argument of the provider is used. Changing this creates a new VPC.

* `description` - (Required) The name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-). Changing this updates the name of the existing VPC.

* `size` - (Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to create a VPC. If omitted, the region argument of the provider is used. Changing this creates a new VPC.

* `volume_type` - (Required) The name of the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-). Changing this updates the name of the existing VPC.


## Attributes Reference
//description of attributes is pending.
The following attributes are exported:

* `id` -  ID of the VPC.

* `availability_zone` - See Argument Reference above.

* `description` - The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.

* `name` -  See Argument Reference above.

* `region` -  See Argument Reference above.

* `size` - Specifies whether the cross-tenant sharing is supported.

* `volume_type` - See Argument Reference above.

## Import

EVSs can be imported using the `id`, e.g.

```
$ terraform import opentelekomcloud_evs_v2.evs_v2 7117d38e-4c8f-4624-a505-bd96b97d024c
```
