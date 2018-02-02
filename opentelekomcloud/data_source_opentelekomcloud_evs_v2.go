package opentelekomcloud

import ("github.com/hashicorp/terraform/helper/schema"
"github.com/gophercloud/gophercloud/openstack/networking/v2/evs"
"fmt"
"log"
	//"github.com/aws/aws-sdk-go/aws/client/metadata"
	//"github.com/aws/aws-sdk-go/aws/client/metadata"
)

func dataSourceElasticVolumeServicesEvsV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceElasticVolumeServicesEvsV2Read,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"volume_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			/*"shareable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},*/

			"snapshot_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},


		/*	"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},*/
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"rel": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"attachments": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"attachment_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

					},
				},
			},
		},
	}
}

func dataSourceElasticVolumeServicesEvsV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	evsClient, err := config.evsV2Client(GetRegion(d, config))
	log.Printf("Inside List")
	listOpts := evs.ListOpts{
		ID:     d.Get("id").(string),
		Name:   d.Get("name").(string),
		Status: d.Get("status").(string),
		Availability_zone:   d.Get("availability_zone").(string),
		Volume_type: d.Get("volume_type").(string),
	}
	log.Printf("Before List")
	refinedEvs, err := evs.List(evsClient,listOpts)
	log.Printf("After List")
	log.Printf("[DEBUG] Value of allEvs: %#v", refinedEvs)
	if err != nil {
		return fmt.Errorf("Unable to retrieve evs: %s", err)
	}

	if len(refinedEvs) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	}

	if len(refinedEvs) > 1 {
		return fmt.Errorf("Your query returned more than one result." +
			" Please try a more specific search criteria")
	}

	Evs := refinedEvs[0]

	var s []map[string]interface{}
	for _, links := range Evs.Links {
			mapping := map[string]interface{}{
				"href": links.Href,
				"rel" : links.Rel,
			}
			s = append(s, mapping)
	}

	var v []map[string]interface{}
	for _, attachments := range Evs.Attachments {
		mapping := map[string]interface{}{
			"server_id": attachments.Server_id,
			"attachment_id" : attachments.Attachment_id,
			"volume_id" : attachments.Volume_id,
			}
		v = append(v, mapping)
	}


	log.Printf("[DEBUG] Retrieved Evs using given filter %s: %+v", Evs.ID, Evs)
	d.SetId(Evs.ID)

	d.Set("name", Evs.Name)
	d.Set("availability_zone", Evs.Availability_zone)
	d.Set("status", Evs.Status)
	d.Set("id", Evs.ID)
	d.Set("size", Evs.Size)
	d.Set("snapshot_id", Evs.Snapshot_id)
	d.Set("description",Evs.Description)
	//d.Set("shareable",Evs.Shareable)
	d.Set("volume_type",Evs.Volume_type)
	d.Set("region", GetRegion(d, config))
	if err := d.Set("links", s); err != nil {
		return err
	}

	if err := d.Set("attachments", v); err != nil {
		return err
	}

	return nil
}
