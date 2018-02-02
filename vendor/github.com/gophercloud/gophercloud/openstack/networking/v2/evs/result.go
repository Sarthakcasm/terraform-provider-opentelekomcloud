package evs

import (
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/gophercloud/gophercloud"

)

type Link struct {
//	href string `json:"href"`
//	rel string `json:"rel"`
Href string `json:"href,omitempty"`
Rel string `json:"rel,omitempty"`

}

type Attachment struct{
	Server_id string `json:"server_id,omitempty"`
	Attachment_id string `json:"attachment_id,omitempty"`
	Volume_id string `json:"volume_id,omitempty"`
}

type EVS struct{

	ID string `json:"id"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	Availability_zone string `json:"availability_zone,omitempty"`

	Created_at string `json:"created_at,omitempty"`

	Volume_type string `json:"volume_type,omitempty"`

	Size int `json:"size,omitempty"`

	//Shareable string `json:"shareable,omitempty"`

	Source_volid string `json:"source_volid,omitempty"`

	Snapshot_id string `json:"snapshot_id,omitempty"`

	Description string `json:"description,omitempty"`

	Os_vol_tenant_attr string `json:"os-vol-tenant-attr:tenant_id,omitempty"`

	Bootable string `json:"bootable,omitempty"`

	Message string `json:"message,omitempty"`

	Code string `json:"code,omitempty"`

	Os_vol_host_attr string `json:"os-vol-host-attr:host,omitempty"`

	Links []Link `json:"links,omitempty"`

	Attachments []Attachment `json:"attachments,omitempty"`

	//Multiattach string `json:"multiattach,omitempty"`


}

func (r EVSPage) NextPageURL() (string, error) {
	var s struct {
		Links []gophercloud.Link `json:"volumes_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return gophercloud.ExtractNextURL(s.Links)
}

func (r EVSPage) IsEmpty() (bool, error) {
	is, err := ExtractEVS(r)
	return len(is) == 0, err
}

func (r commonResult) Extract() (*EVS, error) {
	var s struct {
		Evs *EVS `json:"volume"`
	}
	err := r.ExtractInto(&s)
	return s.Evs, err
}

func (r commonResult) EvsExtract() (*EVS, error) {
	var s struct {
		Evs *EVS `json:"job_id"`
	}
	err := r.ExtractInto(&s)
	return s.Evs, err
}

/*func (r commonResult)ExtractSingleEVS() (*EVS, error){
	var s struct {
		Evs []EVS `json:"volume"`
	}
	err := (r.(EVSPage)).ExtractInto(&s)
	fmt.Println("Error",err)
	return s.Evs, err
}
*/



func ExtractEVS(r pagination.Page) ([]EVS, error) {

	var s struct {
		Evs []EVS `json:"volumes"`
	}
	err := (r.(EVSPage)).ExtractInto(&s)
	return s.Evs, err
}

type GetResult struct {
	commonResult
}

type EVSPage struct {
	pagination.LinkedPageBase
}

type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Evs.
type UpdateResult struct {
	commonResult
}

type commonResult struct {
	gophercloud.Result
}

type DeleteResult struct {
	gophercloud.ErrResult
}