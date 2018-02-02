---
layout: "opentelekomcloud"
page_title: "OpenTelekomCloud: opentelekomcloud_evs_v2"
sidebar_current: "docs-opentelekomcloud-datasource-evs-v2"
description: |-
  Get information on an OpenTelekomCloud EVS.
---

# opentelekomcloud_evs_v2

opentelekomcloud_evs_v2 provides details about a specific EVS.

This resource can prove useful when a module accepts a evs id as an input variable and needs to, for example, determine the Status of that EVS.

## Example Usage
//tobe done later
The following example shows how one might accept a EVS id as a variable and use this data source to obtain the data necessary to create a subnet within it.

```hcl

variable "evs_id" {}

data "opentelekomcloud_evs_v2" "evs" {
  name = "${var.evs_id}"
}

```

## Argument Reference
//to be done later
The arguments of this data source act as filters for querying the available EVSs in the current region. The given filters must match exactly one EVS whose data will be exported as attributes.

* `id` - (Optional) The id of the specific EVS to retrieve.

* `name` - (Optional)  A unique name for the EVS. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).

* `status` - (Optional) The current status of the desired EVS. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.

* `availability_zone` - (Optional) A unique name for the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).

* `volume_type` - (Optional) The cidr block of the desired VPC.



## Attributes Reference
//to be done later
The following attributes are exported:

* `id` - ID of the VPC.

* `name` -  See Argument Reference above.

* `availability_zone` - See Argument Reference above.

* `created_at` - See Argument Reference above.

* `volume_type` - The list of route information with destination and nexthop fields.

* `size` - Specifies whether the cross-tenant sharing is supported.

* `snapshot_id` - See Argument Reference above.

* `links` - See Argument Reference above.

* `status` - See Argument Reference above.

* `size` - See Argument Reference above.
