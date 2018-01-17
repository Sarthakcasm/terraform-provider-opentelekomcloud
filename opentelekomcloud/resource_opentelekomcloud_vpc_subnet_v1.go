package opentelekomcloud

import (
	"fmt"
	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"time"

	"github.com/gophercloud/gophercloud"
	"github.com/hashicorp/terraform/helper/resource"
)

func resourceVpcSubnetV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpcSubnetV1Create, //providers.go
		Read:   resourceVpcSubnetRead,
		Update: resourceVpcSubnetV1Update,
		Delete: resourceVpcSubnetV1Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{ //request and response parameters
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"cidr": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
			},
			"gateway_ip": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     false,
			},
			"dhcp_enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional:true,
				Default:      true,
				ValidateFunc: validateTrueOnly,
				ForceNew: false,
			},
			"primary_dns": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: false,
				Optional:true,
			},
			"secondary_dns": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: false,
				Optional:true,
		    },
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: false,
				Required:true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				DiffSuppressFunc:suppressAsvpcDiff,
				Required:true,
			},
		},

	}
}

func resourceVpcSubnetV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	subnetClient, err := config.subnetV1Client(GetRegion(d, config))

	log.Printf("[DEBUG] Value of vpcClient: %#v", subnetClient)

	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud vpc client: %s", err)
	}

	createOpts := subnets.CreateOpts{
		Name: d.Get("name").(string),
		CIDR: d.Get("cidr").(string),
		AvailabilityZone:d.Get("availability_zone").(string),
		GatewayIP:d.Get("gateway_ip").(string),
		VPC_ID:d.Get("vpc_id").(string),
	}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)
	n, err := subnets.Create(subnetClient, createOpts).Extract()

	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud VPC: %s", err)
	}
	d.SetId(n.ID)

	log.Printf("[INFO] Vpc ID: %s", n.ID)

	log.Printf("[DEBUG] Waiting for OpenTelekomCloud Vpc (%s) to become available", n.ID)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"CREATING"},
		Target:     []string{"ACTIVE"},
		Refresh:    waitForVpcSubnetActive(subnetClient, n.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	d.SetId(n.ID)

	return resourceVpcSubnetRead(d, meta)

}

func resourceVpcSubnetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	subnetClient, err := config.subnetV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud Vpc client: %s", err)
	}

	n, err := subnets.Get(subnetClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(gophercloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving OpenTelekomCloud Vpc: %s", err)
	}

	log.Printf("[DEBUG] Retrieved Vpc %s: %+v", d.Id(), n)

	d.Set("name", n.Name)
	d.Set("cidr", n.CIDR)
	d.Set("gateway_ip", n.GatewayIP)
	d.Set("dhcp_enable", n.EnableDHCP)
	d.Set("primary_dns", n.PRIMARY_DNS)
	d.Set("secondary_dns", n.SECONDARY_DNS)
	d.Set("availability_zone", n.AvailabilityZone)
	d.Set("vpc_id", n.VPC_ID)
	d.Set("region", GetRegion(d, config))

	return nil
}

func resourceVpcSubnetV1Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	vpcClient, err := config.vpcV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud Subnet: %s", err)
	}

	var update bool
	var updateOpts subnets.UpdateOpts

	updateOpts.Name = d.Get("name").(string)

	if d.HasChange("name") {
		update = true

	}
	if d.HasChange("dhcp_enable") {
		update = true
		updateOpts.EnableDHCP = d.Get("dhcp_enable").(bool)
	}
	if d.HasChange("primary_dns") {
		update = true
		updateOpts.PRIMARY_DNS = d.Get("primary_dns").(string)
	}
	if d.HasChange("secondary_dns") {
		update = true
		updateOpts.SECONDARY_DNS = d.Get("secondary_dns").(string)
	}

	log.Printf("[DEBUG] Updating Subnet %s with options: %+v", d.Id(), updateOpts)

	vpc_id:=d.Get("vpc_id").(string)

	log.Printf("[DEBUG] Subnet_id %+v", d.Id())

	if update {
		log.Printf("[DEBUG] Updating subnet %s with options: %#v", d.Id(), updateOpts)
		_, err = subnets.Update(vpcClient,vpc_id, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating OpenTelekomCloud Subnet: %s", err)
		}
	}
	return resourceVpcSubnetRead(d, meta)
}

func resourceVpcSubnetV1Delete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Destroy subnet: %s", d.Id())

	config := meta.(*Config)
	subnetClient, err := config.subnetV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud vpc client: %s", err)
	}
	vpc_id:=d.Get("vpc_id").(string)


	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ACTIVE"},
		Target:     []string{"DELETED"},
		Refresh:    waitForVpcSubnetDelete(subnetClient,vpc_id, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting OpenTelekomCloud Subnet: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForVpcSubnetActive(subnetClient *gophercloud.ServiceClient, vpcId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := subnets.Get(subnetClient, vpcId).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] OpenTelekomCloud VPC Client: %+v", n)
		if n.Status == "DOWN" || n.Status == "OK" {
			return n, "ACTIVE", nil
		}

		return n, n.Status, nil
	}
}

func waitForVpcSubnetDelete(subnetClient *gophercloud.ServiceClient, vpcId string,subnetId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete OpenTelekomCloud subnet %s.\n", subnetId)

		r, err := subnets.Get(subnetClient, subnetId).Extract()
		log.Printf("[DEBUG] Value after extract: %#v", r)
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted OpenTelekomCloud subnet %s", subnetId)
				return r, "DELETED", nil
			}
			return r, "ACTIVE", err
		}
		err = subnets.Delete(subnetClient, vpcId,subnetId).ExtractErr()
		log.Printf("[DEBUG] Value if error: %#v", err)

		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted OpenTelekomCloud subnet %s", subnetId)
				return r, "DELETED", nil
			}
			if errCode, ok := err.(gophercloud.ErrUnexpectedResponseCode); ok {
				if errCode.Actual == 409 {
					return r, "ACTIVE", nil
				}
			}
			return r, "ACTIVE", err
		}

		log.Printf("[DEBUG] OpenTelekomCloud subnet %s still active.\n", subnetId)
		return r, "ACTIVE", nil
	}
}
