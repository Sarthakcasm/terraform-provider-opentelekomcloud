package opentelekomcloud


import (

	"fmt"
	//"github.com/gophercloud/gophercloud/openstack/networking/v1/vpcs"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/evs"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"time"

	"github.com/gophercloud/gophercloud"
	//"github.com/hashicorp/terraform/helper/resource"
	//"strings"
//	"github.com/hashicorp/terraform/helper/resource"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/resource"
)


func resourceElasticVolumeServicesV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticVolumeServicesV2Create, //providers.go
		Read:   resourceElasticVolumeServicesV2Read,
		Update: resourceElasticVolumeServicesV2Update,
		Delete: resourceElasticVolumeServicesV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
//req.Parameters		backup_id,count,availability_zone,description,size,name,imageRef,volume_type
		Schema: map[string]*schema.Schema{ //request and response parameters
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

				/*"backup_id": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},*/
			/*	"count": &schema.Schema{
					Type:     schema.TypeInt,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},*/
				"availability_zone": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},
				"description": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},
				"size": &schema.Schema{
					Type:     schema.TypeInt,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					ValidateFunc: validateName,
				},
			/*	"imageref_type": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},*/
				"volume_type": &schema.Schema{
					Type:     schema.TypeString,
					Required:     true,
					ForceNew:     false,
					//ValidateFunc: validateName,
				},
				"status": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},
				"created_at": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},

			/*	"shareable": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},*/
				"source_volid": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},
				"snapshot_id": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},

				"bootable": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},
				"message": &schema.Schema{
						Type:     schema.TypeString,
						Optional: true,
				},
				"multiattach": &schema.Schema{
						Type:     schema.TypeBool,
						Optional: true,
				},
				"encrypted": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				},
		},
	}
}


func resourceElasticVolumeServicesV2Create(d *schema.ResourceData, meta interface{}) error {

	log.Printf("In Create")
	config := meta.(*Config)
	log.Printf("In Create,Config",config)
	log.Printf("In Create,d",d)
	evsClient, err := config.evsV2Client(GetRegion(d, config))
	log.Printf("In Create evs client",evsClient)
	log.Printf("[DEBUG] Value of evsClient: %#v", evsClient)

	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud evs client: %s", err)
	}


	createOpts := evs.CreateOpts{

		Name : d.Get("name").(string),
		Availability_zone : d.Get("availability_zone").(string),
		Description : d.Get("description").(string),
		Size : d.Get("size").(int),
		Volume_type : d.Get("volume_type").(string),
		}

	log.Printf("[DEBUG] Create Options: %#v", createOpts)

	n:= evs.Create(evsClient, createOpts)

	log.Printf("After create",n)
	log.Printf("After create error ",err)
	if err != nil {
		return fmt.Errorf("[Create]Error creating OpenTelekomCloud EVS: %s", err)
	}
	log.Printf("[INFO] Evs ID: %s", n)
////	d.SetId(n.ID)

	//log.Printf("[INFO] Evs ID: %s", n.ID)

////log.Printf("[DEBUG] Waiting for OpenTelekomCloud Evs (%s) to become available", n.ID)

/*	stateConf := &resource.StateChangeConf{
		Pending:    []string{"Creating"},
		Target:     []string{"Available"},
		Refresh:    waitForEvsActive(evsClient, n.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}*/

	////_, err = stateConf.WaitForState()
	////d.SetId(n.ID)

	return resourceElasticVolumeServicesV2Read(d, meta)

}

func resourceElasticVolumeServicesV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)


	evsClient, err := config.evsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud Evs client: %s", err)
	}

	n, err := evs.Get(evsClient, d.Id()).Extract()
	if err != nil {
		if _, ok := err.(gophercloud.ErrDefault404); ok {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error retrieving OpenTelekomCloud Evs: %s", err)
	}

	log.Printf("[DEBUG] Retrieved Evs %s: %+v", d.Id(), n)

	d.Set("name", n.Name)
	d.Set("availability_zone", n.Availability_zone)
	d.Set("status", n.Status)
	d.Set("id", n.ID)
	d.Set("created_at", n.Created_at)
	d.Set("size", n.Size)
	//d.Set("shareable", n.Shareable)
	d.Set("source_volid", n.Source_volid)
	d.Set("snapshot_id", n.Snapshot_id)
	d.Set("description", n.Description)
	d.Set("os-vol-tenant-attr:tenant_id", n.Os_vol_tenant_attr)
	d.Set("bootable", n.Bootable)

	return nil
}

func resourceElasticVolumeServicesV2Update(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	evsClient, err := config.evsV2Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud evsClient: %s", err)
	}

	var update bool
	var updateOpts evs.UpdateOpts

	if d.HasChange("name") {
		update = true
		updateOpts.Name = d.Get("name").(string)
	}
	if d.HasChange("description") {
		update = true
		updateOpts.Description = d.Get("description").(string)
	}

	log.Printf("[DEBUG] Updating Evs %s with options: %+v", d.Id(), updateOpts)

	if update {
		log.Printf("[DEBUG] Updating Evs %s with options: %#v", d.Id(), updateOpts)
		_, err = evs.Update(evsClient, d.Id(), updateOpts).Extract()
		if err != nil {
			return fmt.Errorf("Error updating OpenTelekomCloud Evs: %s", err)
		}
	}
	return resourceElasticVolumeServicesV2Read(d, meta)
}

func resourceElasticVolumeServicesV2Delete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Destroy evs: %s", d.Id())

	config := meta.(*Config)
	evsClient, err := config.evsV2Client(GetRegion(d, config))

	if err != nil {
		return fmt.Errorf("Error creating OpenTelekomCloud evsClient: %s", err)
	}
	evsId := d.Id()
	err1 :=evs.Delete(evsClient,evsId).ExtractErr()
	if(err1 != nil){
		log.Printf("[DEBUG] Value if error: %#v", err1)
	}

	/*stateConf := &resource.StateChangeConf{
		Pending:    []string{"Available"},
		Target:     []string{"Deleted"},
		Refresh:    waitForEvsDelete(evsClient, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error deleting OpenTelekomCloud Evs: %s", err)
	}
*/
	d.SetId("")
	return nil
}

func waitForEvsActive(evsClient *gophercloud.ServiceClient, evsId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := evs.Get(evsClient, evsId).Extract()
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG] OpenTelekomCloud EVS Client: %+v", n)
		if n.Status == "DOWN" || n.Status == "OK" {
			return n, "Available", nil
		}

		return n, n.Status, nil
	}
}

func waitForEvsDelete(evsClient *gophercloud.ServiceClient, evsId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG] Attempting to delete OpenTelekomCloud vpc %s.\n", evsId)

		r, err := evs.Get(evsClient, evsId).Extract()
		log.Printf("[DEBUG] Value after extract: %#v", r)
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted OpenTelekomCloud vpc %s", evsId)
				return r, "DELETED", nil
			}
			return r, "Available", err
		}

		err = evs.Delete(evsClient, evsId).ExtractErr()
		log.Printf("[DEBUG] Value if error: %#v", err)

		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				log.Printf("[DEBUG] Successfully deleted OpenTelekomCloud vpc %s", evsId)
				return r, "DELETED", nil
			}
			if errCode, ok := err.(gophercloud.ErrUnexpectedResponseCode); ok {
				if errCode.Actual == 409 {
					return r, "Available", nil
				}
			}
			return r, "Available", err
		}

		log.Printf("[DEBUG] OpenTelekomCloud evs %s still available.\n", evsId)
		return r, "Available", nil
	}
}

