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
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// PullRequestTopic is the default topic name
	PullRequestTopic datamodel.TopicNameType = "sourcecode_PullRequest_topic"

	// PullRequestModelName is the model name
	PullRequestModelName datamodel.ModelNameType = "sourcecode.PullRequest"
)

const (
	// PullRequestBranchIDColumn is the branch_id column name
	PullRequestBranchIDColumn = "BranchID"
	// PullRequestBranchNameColumn is the branch_name column name
	PullRequestBranchNameColumn = "BranchName"
	// PullRequestClosedByRefIDColumn is the closed_by_ref_id column name
	PullRequestClosedByRefIDColumn = "ClosedByRefID"
	// PullRequestClosedDateColumn is the closed_date column name
	PullRequestClosedDateColumn = "ClosedDate"
	// PullRequestClosedDateColumnEpochColumn is the epoch column property of the ClosedDate name
	PullRequestClosedDateColumnEpochColumn = "ClosedDate.Epoch"
	// PullRequestClosedDateColumnOffsetColumn is the offset column property of the ClosedDate name
	PullRequestClosedDateColumnOffsetColumn = "ClosedDate.Offset"
	// PullRequestClosedDateColumnRfc3339Column is the rfc3339 column property of the ClosedDate name
	PullRequestClosedDateColumnRfc3339Column = "ClosedDate.Rfc3339"
	// PullRequestCommitIdsColumn is the commit_ids column name
	PullRequestCommitIdsColumn = "CommitIds"
	// PullRequestCommitShasColumn is the commit_shas column name
	PullRequestCommitShasColumn = "CommitShas"
	// PullRequestCreatedByRefIDColumn is the created_by_ref_id column name
	PullRequestCreatedByRefIDColumn = "CreatedByRefID"
	// PullRequestCreatedDateColumn is the created_date column name
	PullRequestCreatedDateColumn = "CreatedDate"
	// PullRequestCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	PullRequestCreatedDateColumnEpochColumn = "CreatedDate.Epoch"
	// PullRequestCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	PullRequestCreatedDateColumnOffsetColumn = "CreatedDate.Offset"
	// PullRequestCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	PullRequestCreatedDateColumnRfc3339Column = "CreatedDate.Rfc3339"
	// PullRequestCustomerIDColumn is the customer_id column name
	PullRequestCustomerIDColumn = "CustomerID"
	// PullRequestDescriptionColumn is the description column name
	PullRequestDescriptionColumn = "Description"
	// PullRequestIDColumn is the id column name
	PullRequestIDColumn = "ID"
	// PullRequestMergeCommitIDColumn is the merge_commit_id column name
	PullRequestMergeCommitIDColumn = "MergeCommitID"
	// PullRequestMergeShaColumn is the merge_sha column name
	PullRequestMergeShaColumn = "MergeSha"
	// PullRequestMergedByRefIDColumn is the merged_by_ref_id column name
	PullRequestMergedByRefIDColumn = "MergedByRefID"
	// PullRequestMergedDateColumn is the merged_date column name
	PullRequestMergedDateColumn = "MergedDate"
	// PullRequestMergedDateColumnEpochColumn is the epoch column property of the MergedDate name
	PullRequestMergedDateColumnEpochColumn = "MergedDate.Epoch"
	// PullRequestMergedDateColumnOffsetColumn is the offset column property of the MergedDate name
	PullRequestMergedDateColumnOffsetColumn = "MergedDate.Offset"
	// PullRequestMergedDateColumnRfc3339Column is the rfc3339 column property of the MergedDate name
	PullRequestMergedDateColumnRfc3339Column = "MergedDate.Rfc3339"
	// PullRequestRefIDColumn is the ref_id column name
	PullRequestRefIDColumn = "RefID"
	// PullRequestRefTypeColumn is the ref_type column name
	PullRequestRefTypeColumn = "RefType"
	// PullRequestRepoIDColumn is the repo_id column name
	PullRequestRepoIDColumn = "RepoID"
	// PullRequestStatusColumn is the status column name
	PullRequestStatusColumn = "Status"
	// PullRequestTitleColumn is the title column name
	PullRequestTitleColumn = "Title"
	// PullRequestUpdatedDateColumn is the updated_date column name
	PullRequestUpdatedDateColumn = "UpdatedDate"
	// PullRequestUpdatedDateColumnEpochColumn is the epoch column property of the UpdatedDate name
	PullRequestUpdatedDateColumnEpochColumn = "UpdatedDate.Epoch"
	// PullRequestUpdatedDateColumnOffsetColumn is the offset column property of the UpdatedDate name
	PullRequestUpdatedDateColumnOffsetColumn = "UpdatedDate.Offset"
	// PullRequestUpdatedDateColumnRfc3339Column is the rfc3339 column property of the UpdatedDate name
	PullRequestUpdatedDateColumnRfc3339Column = "UpdatedDate.Rfc3339"
	// PullRequestUpdatedAtColumn is the updated_ts column name
	PullRequestUpdatedAtColumn = "UpdatedAt"
	// PullRequestURLColumn is the url column name
	PullRequestURLColumn = "URL"
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequestStatus is the enumeration type for status
type PullRequestStatus int32

// UnmarshalJSON unmarshals the enum value
func (v PullRequestStatus) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "OPEN":
		v = 0
	case "CLOSED":
		v = 1
	case "MERGED":
		v = 2
	case "SUPERSEDED":
		v = 3
	case "LOCKED":
		v = 4
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

const (
	// StatusOpen is the enumeration value for open
	PullRequestStatusOpen PullRequestStatus = 0
	// StatusClosed is the enumeration value for closed
	PullRequestStatusClosed PullRequestStatus = 1
	// StatusMerged is the enumeration value for merged
	PullRequestStatusMerged PullRequestStatus = 2
	// StatusSuperseded is the enumeration value for superseded
	PullRequestStatusSuperseded PullRequestStatus = 3
	// StatusLocked is the enumeration value for locked
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
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// PullRequest the pull request for a given repo
type PullRequest struct {
	// BranchID id for the branch based on branch_name, could be already deleted
	BranchID string `json:"branch_id" codec:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// BranchName branch name of the pr, could be  deleted
	BranchName string `json:"branch_name" codec:"branch_name" bson:"branch_name" yaml:"branch_name" faker:"-"`
	// ClosedByRefID the id of user who closed the pull request
	ClosedByRefID string `json:"closed_by_ref_id" codec:"closed_by_ref_id" bson:"closed_by_ref_id" yaml:"closed_by_ref_id" faker:"-"`
	// ClosedDate the timestamp in UTC that the pull request was closed
	ClosedDate PullRequestClosedDate `json:"closed_date" codec:"closed_date" bson:"closed_date" yaml:"closed_date" faker:"-"`
	// CommitIds list of commit ids added in this pr
	CommitIds []string `json:"commit_ids" codec:"commit_ids" bson:"commit_ids" yaml:"commit_ids" faker:"-"`
	// CommitShas list of commit shas added in this pr
	CommitShas []string `json:"commit_shas" codec:"commit_shas" bson:"commit_shas" yaml:"commit_shas" faker:"-"`
	// CreatedByRefID the user ref_id in the source system
	CreatedByRefID string `json:"created_by_ref_id" codec:"created_by_ref_id" bson:"created_by_ref_id" yaml:"created_by_ref_id" faker:"-"`
	// CreatedDate the timestamp in UTC that the pull request was created
	CreatedDate PullRequestCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the pull request
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
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
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	return PullRequestTopic
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
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequest) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedDate
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	case PullRequestCreatedDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for PullRequest")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "created_date",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
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

// IsEqual returns true if the two PullRequest objects are equal
func (o *PullRequest) IsEqual(other *PullRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"branch_id":         toPullRequestObject(o.BranchID, false),
		"branch_name":       toPullRequestObject(o.BranchName, false),
		"closed_by_ref_id":  toPullRequestObject(o.ClosedByRefID, false),
		"closed_date":       toPullRequestObject(o.ClosedDate, false),
		"commit_ids":        toPullRequestObject(o.CommitIds, false),
		"commit_shas":       toPullRequestObject(o.CommitShas, false),
		"created_by_ref_id": toPullRequestObject(o.CreatedByRefID, false),
		"created_date":      toPullRequestObject(o.CreatedDate, false),
		"customer_id":       toPullRequestObject(o.CustomerID, false),
		"description":       toPullRequestObject(o.Description, false),
		"id":                toPullRequestObject(o.ID, false),
		"merge_commit_id":   toPullRequestObject(o.MergeCommitID, false),
		"merge_sha":         toPullRequestObject(o.MergeSha, false),
		"merged_by_ref_id":  toPullRequestObject(o.MergedByRefID, false),
		"merged_date":       toPullRequestObject(o.MergedDate, false),
		"ref_id":            toPullRequestObject(o.RefID, false),
		"ref_type":          toPullRequestObject(o.RefType, false),
		"repo_id":           toPullRequestObject(o.RepoID, false),

		"status":       o.Status.String(),
		"title":        toPullRequestObject(o.Title, false),
		"updated_date": toPullRequestObject(o.UpdatedDate, false),
		"updated_ts":   toPullRequestObject(o.UpdatedAt, false),
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

	if val, ok := kv["branch_id"].(string); ok {
		o.BranchID = val
	} else {
		if val, ok := kv["branch_id"]; ok {
			if val == nil {
				o.BranchID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Description = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
	args = append(args, o.ID)
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
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *PullRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}
