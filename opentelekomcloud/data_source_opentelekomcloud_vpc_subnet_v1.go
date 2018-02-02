package opentelekomcloud

import (
	"fmt"
	"log"

	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceVpcSubnetV1() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVpcSubnetV1Read,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
			Type:     schema.TypeString,
				Optional: true,
		},
			"cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
		},
			"dhcp_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional:true,
			},
			"primary_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional:true,
			},
			"secondary_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional:true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

		},
	}
}

func dataSourceVpcSubnetV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	subnetClient, err := config.subnetV1Client(GetRegion(d, config))

	listOpts := subnets.ListOpts{
		ID:d.Get("id").(string),
		Name:d.Get("name").(string),
		CIDR:d.Get("cidr").(string),
		Status:d.Get("status").(string),
		GatewayIP:d.Get("gateway_ip").(string),
		EnableDHCP:d.Get("dhcp_enable").(bool),
		PRIMARY_DNS:d.Get("primary_dns").(string),
		SECONDARY_DNS:d.Get("secondary_dns").(string),
		AvailabilityZone:d.Get("availability_zone").(string),
		VPC_ID: d.Get("vpc_id").(string),

	}

	refinedSubnets ,err := subnets.List(subnetClient,listOpts)
	log.Printf("[DEBUG] Value of allVpcs: %#v", refinedSubnets)
	if err != nil {
		return fmt.Errorf("Unable to retrieve subnets: %s", err)
	}

	if len(refinedSubnets) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(refinedSubnets) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	Subnets := refinedSubnets[0]

	log.Printf("[DEBUG] Retrieved Vpcs using given filter %s: %+v", Subnets.ID, Subnets)
	d.SetId(Subnets.ID)

	d.Set("name", Subnets.Name)
	d.Set("cidr", Subnets.CIDR)
	d.Set("status", Subnets.Status)
	d.Set("gateway_ip", Subnets.GatewayIP)
	d.Set("dhcp_enable", Subnets.EnableDHCP)
	d.Set("primary_dns", Subnets.PRIMARY_DNS)
	d.Set("secondary_dns", Subnets.SECONDARY_DNS)
	d.Set("availability_zone", Subnets.AvailabilityZone)
	d.Set("vpc_id", Subnets.VPC_ID)
	d.Set("region", GetRegion(d, config))

	return nil
}
