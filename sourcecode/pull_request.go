// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// PullRequestTable is the default table name
	PullRequestTable datamodel.ModelNameType = "sourcecode_pullrequest"

	// PullRequestModelName is the model name
	PullRequestModelName datamodel.ModelNameType = "sourcecode.PullRequest"
)

const (
	// PullRequestModelActiveColumn is the column json value active
	PullRequestModelActiveColumn = "active"
	// PullRequestModelBranchIDColumn is the column json value branch_id
	PullRequestModelBranchIDColumn = "branch_id"
	// PullRequestModelBranchNameColumn is the column json value branch_name
	PullRequestModelBranchNameColumn = "branch_name"
	// PullRequestModelClosedByRefIDColumn is the column json value closed_by_ref_id
	PullRequestModelClosedByRefIDColumn = "closed_by_ref_id"
	// PullRequestModelClosedDateColumn is the column json value closed_date
	PullRequestModelClosedDateColumn = "closed_date"
	// PullRequestModelClosedDateEpochColumn is the column json value epoch
	PullRequestModelClosedDateEpochColumn = "epoch"
	// PullRequestModelClosedDateOffsetColumn is the column json value offset
	PullRequestModelClosedDateOffsetColumn = "offset"
	// PullRequestModelClosedDateRfc3339Column is the column json value rfc3339
	PullRequestModelClosedDateRfc3339Column = "rfc3339"
	// PullRequestModelCommitIdsColumn is the column json value commit_ids
	PullRequestModelCommitIdsColumn = "commit_ids"
	// PullRequestModelCommitShasColumn is the column json value commit_shas
	PullRequestModelCommitShasColumn = "commit_shas"
	// PullRequestModelCreatedByRefIDColumn is the column json value created_by_ref_id
	PullRequestModelCreatedByRefIDColumn = "created_by_ref_id"
	// PullRequestModelCreatedDateColumn is the column json value created_date
	PullRequestModelCreatedDateColumn = "created_date"
	// PullRequestModelCreatedDateEpochColumn is the column json value epoch
	PullRequestModelCreatedDateEpochColumn = "epoch"
	// PullRequestModelCreatedDateOffsetColumn is the column json value offset
	PullRequestModelCreatedDateOffsetColumn = "offset"
	// PullRequestModelCreatedDateRfc3339Column is the column json value rfc3339
	PullRequestModelCreatedDateRfc3339Column = "rfc3339"
	// PullRequestModelCustomerIDColumn is the column json value customer_id
	PullRequestModelCustomerIDColumn = "customer_id"
	// PullRequestModelDescriptionColumn is the column json value description
	PullRequestModelDescriptionColumn = "description"
	// PullRequestModelDraftColumn is the column json value draft
	PullRequestModelDraftColumn = "draft"
	// PullRequestModelIDColumn is the column json value id
	PullRequestModelIDColumn = "id"
	// PullRequestModelIdentifierColumn is the column json value identifier
	PullRequestModelIdentifierColumn = "identifier"
	// PullRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PullRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PullRequestModelMergeCommitIDColumn is the column json value merge_commit_id
	PullRequestModelMergeCommitIDColumn = "merge_commit_id"
	// PullRequestModelMergeShaColumn is the column json value merge_sha
	PullRequestModelMergeShaColumn = "merge_sha"
	// PullRequestModelMergedByRefIDColumn is the column json value merged_by_ref_id
	PullRequestModelMergedByRefIDColumn = "merged_by_ref_id"
	// PullRequestModelMergedDateColumn is the column json value merged_date
	PullRequestModelMergedDateColumn = "merged_date"
	// PullRequestModelMergedDateEpochColumn is the column json value epoch
	PullRequestModelMergedDateEpochColumn = "epoch"
	// PullRequestModelMergedDateOffsetColumn is the column json value offset
	PullRequestModelMergedDateOffsetColumn = "offset"
	// PullRequestModelMergedDateRfc3339Column is the column json value rfc3339
	PullRequestModelMergedDateRfc3339Column = "rfc3339"
	// PullRequestModelRefIDColumn is the column json value ref_id
	PullRequestModelRefIDColumn = "ref_id"
	// PullRequestModelRefTypeColumn is the column json value ref_type
	PullRequestModelRefTypeColumn = "ref_type"
	// PullRequestModelRepoIDColumn is the column json value repo_id
	PullRequestModelRepoIDColumn = "repo_id"
	// PullRequestModelStatusColumn is the column json value status
	PullRequestModelStatusColumn = "status"
	// PullRequestModelTitleColumn is the column json value title
	PullRequestModelTitleColumn = "title"
	// PullRequestModelUpdatedDateColumn is the column json value updated_date
	PullRequestModelUpdatedDateColumn = "updated_date"
	// PullRequestModelUpdatedDateEpochColumn is the column json value epoch
	PullRequestModelUpdatedDateEpochColumn = "epoch"
	// PullRequestModelUpdatedDateOffsetColumn is the column json value offset
	PullRequestModelUpdatedDateOffsetColumn = "offset"
	// PullRequestModelUpdatedDateRfc3339Column is the column json value rfc3339
	PullRequestModelUpdatedDateRfc3339Column = "rfc3339"
	// PullRequestModelURLColumn is the column json value url
	PullRequestModelURLColumn = "url"
)

// PullRequestClosedDate represents the object structure for closed_date
type PullRequestClosedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestClosedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestClosedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestClosedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestClosedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestClosedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestClosedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestClosedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestClosedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequestCreatedDate represents the object structure for created_date
type PullRequestCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCreatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequestMergedDate represents the object structure for merged_date
type PullRequestMergedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestMergedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestMergedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestMergedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestMergedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestMergedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestMergedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestMergedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestMergedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequestStatus is the enumeration type for status
type PullRequestStatus int32

// toPullRequestStatusPointer is the enumeration pointer type for status
func toPullRequestStatusPointer(v int32) *PullRequestStatus {
	nv := PullRequestStatus(v)
	return &nv
}

// toPullRequestStatusEnum is the enumeration pointer wrapper for status
func toPullRequestStatusEnum(v *PullRequestStatus) string {
	if v == nil {
		return toPullRequestStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *PullRequestStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = PullRequestStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "OPEN":
			*v = PullRequestStatus(0)
		case "CLOSED":
			*v = PullRequestStatus(1)
		case "MERGED":
			*v = PullRequestStatus(2)
		case "SUPERSEDED":
			*v = PullRequestStatus(3)
		case "LOCKED":
			*v = PullRequestStatus(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *PullRequestStatus) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "OPEN":
		*v = 0
	case "CLOSED":
		*v = 1
	case "MERGED":
		*v = 2
	case "SUPERSEDED":
		*v = 3
	case "LOCKED":
		*v = 4
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v PullRequestStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("OPEN")
	case 1:
		return json.Marshal("CLOSED")
	case 2:
		return json.Marshal("MERGED")
	case 3:
		return json.Marshal("SUPERSEDED")
	case 4:
		return json.Marshal("LOCKED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v PullRequestStatus) String() string {
	switch int32(v) {
	case 0:
		return "OPEN"
	case 1:
		return "CLOSED"
	case 2:
		return "MERGED"
	case 3:
		return "SUPERSEDED"
	case 4:
		return "LOCKED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *PullRequestStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case PullRequestStatus:
		*v = val
	case int32:
		*v = PullRequestStatus(int32(val))
	case int:
		*v = PullRequestStatus(int32(val))
	case string:
		switch val {
		case "OPEN":
			*v = PullRequestStatus(0)
		case "CLOSED":
			*v = PullRequestStatus(1)
		case "MERGED":
			*v = PullRequestStatus(2)
		case "SUPERSEDED":
			*v = PullRequestStatus(3)
		case "LOCKED":
			*v = PullRequestStatus(4)
		}
	}
	return nil
}

const (
	// PullRequestStatusOpen is the enumeration value for open
	PullRequestStatusOpen PullRequestStatus = 0
	// PullRequestStatusClosed is the enumeration value for closed
	PullRequestStatusClosed PullRequestStatus = 1
	// PullRequestStatusMerged is the enumeration value for merged
	PullRequestStatusMerged PullRequestStatus = 2
	// PullRequestStatusSuperseded is the enumeration value for superseded
	PullRequestStatusSuperseded PullRequestStatus = 3
	// PullRequestStatusLocked is the enumeration value for locked
	PullRequestStatusLocked PullRequestStatus = 4
)

// PullRequestUpdatedDate represents the object structure for updated_date
type PullRequestUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestUpdatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequest the pull request for a given repo
type PullRequest struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// BranchID id for the branch based on branch_name, could be already deleted
	BranchID string `json:"branch_id" codec:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// BranchName branch name of the pull request
	BranchName string `json:"branch_name" codec:"branch_name" bson:"branch_name" yaml:"branch_name" faker:"-"`
	// ClosedByRefID the id of user who closed the pull request
	ClosedByRefID string `json:"closed_by_ref_id" codec:"closed_by_ref_id" bson:"closed_by_ref_id" yaml:"closed_by_ref_id" faker:"-"`
	// ClosedDate the timestamp in UTC that the pull request was closed
	ClosedDate PullRequestClosedDate `json:"closed_date" codec:"closed_date" bson:"closed_date" yaml:"closed_date" faker:"-"`
	// CommitIds list of commit ids added in this pull request
	CommitIds []string `json:"commit_ids" codec:"commit_ids" bson:"commit_ids" yaml:"commit_ids" faker:"-"`
	// CommitShas list of commit shas added in this pull request
	CommitShas []string `json:"commit_shas" codec:"commit_shas" bson:"commit_shas" yaml:"commit_shas" faker:"-"`
	// CreatedByRefID the user ref_id in the source system
	CreatedByRefID string `json:"created_by_ref_id" codec:"created_by_ref_id" bson:"created_by_ref_id" yaml:"created_by_ref_id" faker:"-"`
	// CreatedDate the timestamp in UTC that the pull request was created
	CreatedDate PullRequestCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the pull request
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// Draft if the pull request is in draft mode or not
	Draft bool `json:"draft" codec:"draft" bson:"draft" yaml:"draft" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier a human friendly identifier when displaying this pull request
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// MergeCommitID the id of the merge commit
	MergeCommitID string `json:"merge_commit_id" codec:"merge_commit_id" bson:"merge_commit_id" yaml:"merge_commit_id" faker:"-"`
	// MergeSha the sha of the merge commit
	MergeSha string `json:"merge_sha" codec:"merge_sha" bson:"merge_sha" yaml:"merge_sha" faker:"-"`
	// MergedByRefID the id of user who merged the pull request
	MergedByRefID string `json:"merged_by_ref_id" codec:"merged_by_ref_id" bson:"merged_by_ref_id" yaml:"merged_by_ref_id" faker:"-"`
	// MergedDate the timestamp in UTC that the pull request was merged
	MergedDate PullRequestMergedDate `json:"merged_date" codec:"merged_date" bson:"merged_date" yaml:"merged_date" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Status the status of the pull request
	Status PullRequestStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// Title the title of the pull request
	Title string `json:"title" codec:"title" bson:"title" yaml:"title" faker:"commit_message"`
	// UpdatedDate the timestamp in UTC that the pull request was closed
	UpdatedDate PullRequestUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// URL the url to the pull request home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequest)(nil)

func toPullRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequest:
		return v.ToMap()

	case PullRequestClosedDate:
		return v.ToMap()

	case PullRequestCreatedDate:
		return v.ToMap()

	case PullRequestMergedDate:
		return v.ToMap()

	case PullRequestStatus:
		return v.String()

	case PullRequestUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequest
func (o *PullRequest) String() string {
	return fmt.Sprintf("sourcecode.PullRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *PullRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequest) GetModelName() datamodel.ModelNameType {
	return PullRequestModelName
}

// NewPullRequestID provides a template for generating an ID field for PullRequest
func NewPullRequestID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequest) setDefaults(frommap bool) {
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType, o.RepoID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *PullRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequest
func (o *PullRequest) Clone() datamodel.Model {
	c := new(PullRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequest) Anon() datamodel.Model {
	c := new(PullRequest)
	if err := faker.FakeData(c); err != nil {
		panic("couldn't create anon version of object: " + err.Error())
	}
	kv := c.ToMap()
	for k, v := range o.ToMap() {
		if _, ok := kv[k]; !ok {
			kv[k] = v
		}
	}
	c.FromMap(kv)
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *PullRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *PullRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two PullRequest objects are equal
func (o *PullRequest) IsEqual(other *PullRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toPullRequestObject(o.Active, false),
		"branch_id":               toPullRequestObject(o.BranchID, false),
		"branch_name":             toPullRequestObject(o.BranchName, false),
		"closed_by_ref_id":        toPullRequestObject(o.ClosedByRefID, false),
		"closed_date":             toPullRequestObject(o.ClosedDate, false),
		"commit_ids":              toPullRequestObject(o.CommitIds, false),
		"commit_shas":             toPullRequestObject(o.CommitShas, false),
		"created_by_ref_id":       toPullRequestObject(o.CreatedByRefID, false),
		"created_date":            toPullRequestObject(o.CreatedDate, false),
		"customer_id":             toPullRequestObject(o.CustomerID, false),
		"description":             toPullRequestObject(o.Description, false),
		"draft":                   toPullRequestObject(o.Draft, false),
		"id":                      toPullRequestObject(o.ID, false),
		"identifier":              toPullRequestObject(o.Identifier, false),
		"integration_instance_id": toPullRequestObject(o.IntegrationInstanceID, true),
		"merge_commit_id":         toPullRequestObject(o.MergeCommitID, false),
		"merge_sha":               toPullRequestObject(o.MergeSha, false),
		"merged_by_ref_id":        toPullRequestObject(o.MergedByRefID, false),
		"merged_date":             toPullRequestObject(o.MergedDate, false),
		"ref_id":                  toPullRequestObject(o.RefID, false),
		"ref_type":                toPullRequestObject(o.RefType, false),
		"repo_id":                 toPullRequestObject(o.RepoID, false),

		"status":       o.Status.String(),
		"title":        toPullRequestObject(o.Title, false),
		"updated_date": toPullRequestObject(o.UpdatedDate, false),
		"url":          toPullRequestObject(o.URL, false),
		"hashcode":     toPullRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = false
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["branch_id"].(string); ok {
		o.BranchID = val
	} else {
		if val, ok := kv["branch_id"]; ok {
			if val == nil {
				o.BranchID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.BranchID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["branch_name"].(string); ok {
		o.BranchName = val
	} else {
		if val, ok := kv["branch_name"]; ok {
			if val == nil {
				o.BranchName = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.BranchName = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["closed_by_ref_id"].(string); ok {
		o.ClosedByRefID = val
	} else {
		if val, ok := kv["closed_by_ref_id"]; ok {
			if val == nil {
				o.ClosedByRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ClosedByRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["closed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ClosedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestClosedDate); ok {
			// struct
			o.ClosedDate = sv
		} else if sp, ok := val.(*PullRequestClosedDate); ok {
			// struct pointer
			o.ClosedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ClosedDate.Epoch = dt.Epoch
			o.ClosedDate.Rfc3339 = dt.Rfc3339
			o.ClosedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ClosedDate.Epoch = dt.Epoch
			o.ClosedDate.Rfc3339 = dt.Rfc3339
			o.ClosedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ClosedDate.Epoch = dt.Epoch
				o.ClosedDate.Rfc3339 = dt.Rfc3339
				o.ClosedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ClosedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["commit_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_ids field")
				}
			}
			o.CommitIds = na
		}
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}
	if val, ok := kv["commit_shas"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_shas field")
				}
			}
			o.CommitShas = na
		}
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}
	if val, ok := kv["created_by_ref_id"].(string); ok {
		o.CreatedByRefID = val
	} else {
		if val, ok := kv["created_by_ref_id"]; ok {
			if val == nil {
				o.CreatedByRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CreatedByRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["draft"].(bool); ok {
		o.Draft = val
	} else {
		if val, ok := kv["draft"]; ok {
			if val == nil {
				o.Draft = false
			} else {
				o.Draft = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Identifier = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["merge_commit_id"].(string); ok {
		o.MergeCommitID = val
	} else {
		if val, ok := kv["merge_commit_id"]; ok {
			if val == nil {
				o.MergeCommitID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.MergeCommitID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["merge_sha"].(string); ok {
		o.MergeSha = val
	} else {
		if val, ok := kv["merge_sha"]; ok {
			if val == nil {
				o.MergeSha = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.MergeSha = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["merged_by_ref_id"].(string); ok {
		o.MergedByRefID = val
	} else {
		if val, ok := kv["merged_by_ref_id"]; ok {
			if val == nil {
				o.MergedByRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.MergedByRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merged_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.MergedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestMergedDate); ok {
			// struct
			o.MergedDate = sv
		} else if sp, ok := val.(*PullRequestMergedDate); ok {
			// struct pointer
			o.MergedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.MergedDate.Epoch = dt.Epoch
			o.MergedDate.Rfc3339 = dt.Rfc3339
			o.MergedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.MergedDate.Epoch = dt.Epoch
			o.MergedDate.Rfc3339 = dt.Rfc3339
			o.MergedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.MergedDate.Epoch = dt.Epoch
				o.MergedDate.Rfc3339 = dt.Rfc3339
				o.MergedDate.Offset = dt.Offset
			}
		}
	} else {
		o.MergedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RepoID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["status"].(PullRequestStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["sourcecode.status"].(string)
			switch ev {
			case "open", "OPEN":
				o.Status = 0
			case "closed", "CLOSED":
				o.Status = 1
			case "merged", "MERGED":
				o.Status = 2
			case "superseded", "SUPERSEDED":
				o.Status = 3
			case "locked", "LOCKED":
				o.Status = 4
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "open", "OPEN":
				o.Status = 0
			case "closed", "CLOSED":
				o.Status = 1
			case "merged", "MERGED":
				o.Status = 2
			case "superseded", "SUPERSEDED":
				o.Status = 3
			case "locked", "LOCKED":
				o.Status = 4
			}
		}
	}
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		if val, ok := kv["title"]; ok {
			if val == nil {
				o.Title = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Title = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*PullRequestUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *PullRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.BranchID)
	args = append(args, o.BranchName)
	args = append(args, o.ClosedByRefID)
	args = append(args, o.ClosedDate)
	args = append(args, o.CommitIds)
	args = append(args, o.CommitShas)
	args = append(args, o.CreatedByRefID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.Draft)
	args = append(args, o.ID)
	args = append(args, o.Identifier)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.MergeCommitID)
	args = append(args, o.MergeSha)
	args = append(args, o.MergedByRefID)
	args = append(args, o.MergedDate)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Status)
	args = append(args, o.Title)
	args = append(args, o.UpdatedDate)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// PullRequestPartial is a partial struct for upsert mutations for PullRequest
type PullRequestPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// BranchID id for the branch based on branch_name, could be already deleted
	BranchID *string `json:"branch_id,omitempty"`
	// BranchName branch name of the pull request
	BranchName *string `json:"branch_name,omitempty"`
	// ClosedByRefID the id of user who closed the pull request
	ClosedByRefID *string `json:"closed_by_ref_id,omitempty"`
	// ClosedDate the timestamp in UTC that the pull request was closed
	ClosedDate *PullRequestClosedDate `json:"closed_date,omitempty"`
	// CommitIds list of commit ids added in this pull request
	CommitIds []string `json:"commit_ids,omitempty"`
	// CommitShas list of commit shas added in this pull request
	CommitShas []string `json:"commit_shas,omitempty"`
	// CreatedByRefID the user ref_id in the source system
	CreatedByRefID *string `json:"created_by_ref_id,omitempty"`
	// CreatedDate the timestamp in UTC that the pull request was created
	CreatedDate *PullRequestCreatedDate `json:"created_date,omitempty"`
	// Description the description of the pull request
	Description *string `json:"description,omitempty"`
	// Draft if the pull request is in draft mode or not
	Draft *bool `json:"draft,omitempty"`
	// Identifier a human friendly identifier when displaying this pull request
	Identifier *string `json:"identifier,omitempty"`
	// MergeCommitID the id of the merge commit
	MergeCommitID *string `json:"merge_commit_id,omitempty"`
	// MergeSha the sha of the merge commit
	MergeSha *string `json:"merge_sha,omitempty"`
	// MergedByRefID the id of user who merged the pull request
	MergedByRefID *string `json:"merged_by_ref_id,omitempty"`
	// MergedDate the timestamp in UTC that the pull request was merged
	MergedDate *PullRequestMergedDate `json:"merged_date,omitempty"`
	// RepoID the unique id for the repo
	RepoID *string `json:"repo_id,omitempty"`
	// Status the status of the pull request
	Status *PullRequestStatus `json:"status,omitempty"`
	// Title the title of the pull request
	Title *string `json:"title,omitempty"`
	// UpdatedDate the timestamp in UTC that the pull request was closed
	UpdatedDate *PullRequestUpdatedDate `json:"updated_date,omitempty"`
	// URL the url to the pull request home page
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*PullRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *PullRequestPartial) GetModelName() datamodel.ModelNameType {
	return PullRequestModelName
}

// ToMap returns the object as a map
func (o *PullRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":            toPullRequestObject(o.Active, true),
		"branch_id":         toPullRequestObject(o.BranchID, true),
		"branch_name":       toPullRequestObject(o.BranchName, true),
		"closed_by_ref_id":  toPullRequestObject(o.ClosedByRefID, true),
		"closed_date":       toPullRequestObject(o.ClosedDate, true),
		"commit_ids":        toPullRequestObject(o.CommitIds, true),
		"commit_shas":       toPullRequestObject(o.CommitShas, true),
		"created_by_ref_id": toPullRequestObject(o.CreatedByRefID, true),
		"created_date":      toPullRequestObject(o.CreatedDate, true),
		"description":       toPullRequestObject(o.Description, true),
		"draft":             toPullRequestObject(o.Draft, true),
		"identifier":        toPullRequestObject(o.Identifier, true),
		"merge_commit_id":   toPullRequestObject(o.MergeCommitID, true),
		"merge_sha":         toPullRequestObject(o.MergeSha, true),
		"merged_by_ref_id":  toPullRequestObject(o.MergedByRefID, true),
		"merged_date":       toPullRequestObject(o.MergedDate, true),
		"repo_id":           toPullRequestObject(o.RepoID, true),

		"status":       toPullRequestStatusEnum(o.Status),
		"title":        toPullRequestObject(o.Title, true),
		"updated_date": toPullRequestObject(o.UpdatedDate, true),
		"url":          toPullRequestObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "closed_date" {
				if dt, ok := v.(*PullRequestClosedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "commit_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "commit_shas" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "created_date" {
				if dt, ok := v.(*PullRequestCreatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "merged_date" {
				if dt, ok := v.(*PullRequestMergedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "updated_date" {
				if dt, ok := v.(*PullRequestUpdatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *PullRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *PullRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *PullRequestPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["branch_id"].(*string); ok {
		o.BranchID = val
	} else if val, ok := kv["branch_id"].(string); ok {
		o.BranchID = &val
	} else {
		if val, ok := kv["branch_id"]; ok {
			if val == nil {
				o.BranchID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.BranchID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["branch_name"].(*string); ok {
		o.BranchName = val
	} else if val, ok := kv["branch_name"].(string); ok {
		o.BranchName = &val
	} else {
		if val, ok := kv["branch_name"]; ok {
			if val == nil {
				o.BranchName = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.BranchName = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["closed_by_ref_id"].(*string); ok {
		o.ClosedByRefID = val
	} else if val, ok := kv["closed_by_ref_id"].(string); ok {
		o.ClosedByRefID = &val
	} else {
		if val, ok := kv["closed_by_ref_id"]; ok {
			if val == nil {
				o.ClosedByRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ClosedByRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.ClosedDate == nil {
		o.ClosedDate = &PullRequestClosedDate{}
	}

	if val, ok := kv["closed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ClosedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestClosedDate); ok {
			// struct
			o.ClosedDate = &sv
		} else if sp, ok := val.(*PullRequestClosedDate); ok {
			// struct pointer
			o.ClosedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ClosedDate.Epoch = dt.Epoch
			o.ClosedDate.Rfc3339 = dt.Rfc3339
			o.ClosedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.ClosedDate.Epoch = dt.Epoch
			o.ClosedDate.Rfc3339 = dt.Rfc3339
			o.ClosedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ClosedDate.Epoch = dt.Epoch
				o.ClosedDate.Rfc3339 = dt.Rfc3339
				o.ClosedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ClosedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["commit_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_ids field")
				}
			}
			o.CommitIds = na
		}
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}
	if val, ok := kv["commit_shas"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_shas field")
				}
			}
			o.CommitShas = na
		}
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}
	if val, ok := kv["created_by_ref_id"].(*string); ok {
		o.CreatedByRefID = val
	} else if val, ok := kv["created_by_ref_id"].(string); ok {
		o.CreatedByRefID = &val
	} else {
		if val, ok := kv["created_by_ref_id"]; ok {
			if val == nil {
				o.CreatedByRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CreatedDate == nil {
		o.CreatedDate = &PullRequestCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*PullRequestCreatedDate); ok {
			// struct pointer
			o.CreatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["draft"].(*bool); ok {
		o.Draft = val
	} else if val, ok := kv["draft"].(bool); ok {
		o.Draft = &val
	} else {
		if val, ok := kv["draft"]; ok {
			if val == nil {
				o.Draft = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Draft = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["identifier"].(*string); ok {
		o.Identifier = val
	} else if val, ok := kv["identifier"].(string); ok {
		o.Identifier = &val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Identifier = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["merge_commit_id"].(*string); ok {
		o.MergeCommitID = val
	} else if val, ok := kv["merge_commit_id"].(string); ok {
		o.MergeCommitID = &val
	} else {
		if val, ok := kv["merge_commit_id"]; ok {
			if val == nil {
				o.MergeCommitID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.MergeCommitID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["merge_sha"].(*string); ok {
		o.MergeSha = val
	} else if val, ok := kv["merge_sha"].(string); ok {
		o.MergeSha = &val
	} else {
		if val, ok := kv["merge_sha"]; ok {
			if val == nil {
				o.MergeSha = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.MergeSha = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["merged_by_ref_id"].(*string); ok {
		o.MergedByRefID = val
	} else if val, ok := kv["merged_by_ref_id"].(string); ok {
		o.MergedByRefID = &val
	} else {
		if val, ok := kv["merged_by_ref_id"]; ok {
			if val == nil {
				o.MergedByRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.MergedByRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.MergedDate == nil {
		o.MergedDate = &PullRequestMergedDate{}
	}

	if val, ok := kv["merged_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.MergedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestMergedDate); ok {
			// struct
			o.MergedDate = &sv
		} else if sp, ok := val.(*PullRequestMergedDate); ok {
			// struct pointer
			o.MergedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.MergedDate.Epoch = dt.Epoch
			o.MergedDate.Rfc3339 = dt.Rfc3339
			o.MergedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.MergedDate.Epoch = dt.Epoch
			o.MergedDate.Rfc3339 = dt.Rfc3339
			o.MergedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.MergedDate.Epoch = dt.Epoch
				o.MergedDate.Rfc3339 = dt.Rfc3339
				o.MergedDate.Offset = dt.Offset
			}
		}
	} else {
		o.MergedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["repo_id"].(*string); ok {
		o.RepoID = val
	} else if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = &val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["status"].(*PullRequestStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(PullRequestStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toPullRequestStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["PullRequestStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "open", "OPEN":
						o.Status = toPullRequestStatusPointer(0)
					case "closed", "CLOSED":
						o.Status = toPullRequestStatusPointer(1)
					case "merged", "MERGED":
						o.Status = toPullRequestStatusPointer(2)
					case "superseded", "SUPERSEDED":
						o.Status = toPullRequestStatusPointer(3)
					case "locked", "LOCKED":
						o.Status = toPullRequestStatusPointer(4)
					}
				}
			}
		}
	}
	if val, ok := kv["title"].(*string); ok {
		o.Title = val
	} else if val, ok := kv["title"].(string); ok {
		o.Title = &val
	} else {
		if val, ok := kv["title"]; ok {
			if val == nil {
				o.Title = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Title = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.UpdatedDate == nil {
		o.UpdatedDate = &PullRequestUpdatedDate{}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestUpdatedDate); ok {
			// struct
			o.UpdatedDate = &sv
		} else if sp, ok := val.(*PullRequestUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.UpdatedDate.Epoch = dt.Epoch
				o.UpdatedDate.Rfc3339 = dt.Rfc3339
				o.UpdatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["url"].(*string); ok {
		o.URL = val
	} else if val, ok := kv["url"].(string); ok {
		o.URL = &val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.URL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
