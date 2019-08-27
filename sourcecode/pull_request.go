// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// PullRequestTopic is the default topic name
	PullRequestTopic datamodel.TopicNameType = "sourcecode_PullRequest_topic"

	// PullRequestStream is the default stream name
	PullRequestStream datamodel.TopicNameType = "sourcecode_PullRequest_stream"

	// PullRequestTable is the default table name
	PullRequestTable datamodel.TopicNameType = "sourcecode_pullrequest"

	// PullRequestModelName is the model name
	PullRequestModelName datamodel.ModelNameType = "sourcecode.PullRequest"
)

const (
	// PullRequestBranchIDColumn is the branch_id column name
	PullRequestBranchIDColumn = "branch_id"
	// PullRequestClosedByRefIDColumn is the closed_by_ref_id column name
	PullRequestClosedByRefIDColumn = "closed_by_ref_id"
	// PullRequestClosedDateColumn is the closed_date column name
	PullRequestClosedDateColumn = "closed_date"
	// PullRequestClosedDateColumnEpochColumn is the epoch column property of the ClosedDate name
	PullRequestClosedDateColumnEpochColumn = "closed_date->epoch"
	// PullRequestClosedDateColumnOffsetColumn is the offset column property of the ClosedDate name
	PullRequestClosedDateColumnOffsetColumn = "closed_date->offset"
	// PullRequestClosedDateColumnRfc3339Column is the rfc3339 column property of the ClosedDate name
	PullRequestClosedDateColumnRfc3339Column = "closed_date->rfc3339"
	// PullRequestCreatedByRefIDColumn is the created_by_ref_id column name
	PullRequestCreatedByRefIDColumn = "created_by_ref_id"
	// PullRequestCreatedDateColumn is the created_date column name
	PullRequestCreatedDateColumn = "created_date"
	// PullRequestCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	PullRequestCreatedDateColumnEpochColumn = "created_date->epoch"
	// PullRequestCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	PullRequestCreatedDateColumnOffsetColumn = "created_date->offset"
	// PullRequestCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	PullRequestCreatedDateColumnRfc3339Column = "created_date->rfc3339"
	// PullRequestCustomerIDColumn is the customer_id column name
	PullRequestCustomerIDColumn = "customer_id"
	// PullRequestDescriptionColumn is the description column name
	PullRequestDescriptionColumn = "description"
	// PullRequestIDColumn is the id column name
	PullRequestIDColumn = "id"
	// PullRequestMergeShaColumn is the merge_sha column name
	PullRequestMergeShaColumn = "merge_sha"
	// PullRequestMergedByRefIDColumn is the merged_by_ref_id column name
	PullRequestMergedByRefIDColumn = "merged_by_ref_id"
	// PullRequestMergedDateColumn is the merged_date column name
	PullRequestMergedDateColumn = "merged_date"
	// PullRequestMergedDateColumnEpochColumn is the epoch column property of the MergedDate name
	PullRequestMergedDateColumnEpochColumn = "merged_date->epoch"
	// PullRequestMergedDateColumnOffsetColumn is the offset column property of the MergedDate name
	PullRequestMergedDateColumnOffsetColumn = "merged_date->offset"
	// PullRequestMergedDateColumnRfc3339Column is the rfc3339 column property of the MergedDate name
	PullRequestMergedDateColumnRfc3339Column = "merged_date->rfc3339"
	// PullRequestRefIDColumn is the ref_id column name
	PullRequestRefIDColumn = "ref_id"
	// PullRequestRefTypeColumn is the ref_type column name
	PullRequestRefTypeColumn = "ref_type"
	// PullRequestRepoIDColumn is the repo_id column name
	PullRequestRepoIDColumn = "repo_id"
	// PullRequestStatusColumn is the status column name
	PullRequestStatusColumn = "status"
	// PullRequestTitleColumn is the title column name
	PullRequestTitleColumn = "title"
	// PullRequestUpdatedDateColumn is the updated_date column name
	PullRequestUpdatedDateColumn = "updated_date"
	// PullRequestUpdatedDateColumnEpochColumn is the epoch column property of the UpdatedDate name
	PullRequestUpdatedDateColumnEpochColumn = "updated_date->epoch"
	// PullRequestUpdatedDateColumnOffsetColumn is the offset column property of the UpdatedDate name
	PullRequestUpdatedDateColumnOffsetColumn = "updated_date->offset"
	// PullRequestUpdatedDateColumnRfc3339Column is the rfc3339 column property of the UpdatedDate name
	PullRequestUpdatedDateColumnRfc3339Column = "updated_date->rfc3339"
	// PullRequestUpdatedAtColumn is the updated_ts column name
	PullRequestUpdatedAtColumn = "updated_ts"
	// PullRequestURLColumn is the url column name
	PullRequestURLColumn = "url"
)

// PullRequestClosedDate represents the object structure for closed_date
type PullRequestClosedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestClosedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestClosedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestClosedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestClosedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestClosedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestClosedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestClosedDateObject(o.Rfc3339, isavro, false, "string"),
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestCreatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCreatedDateObject(o.Rfc3339, isavro, false, "string"),
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestMergedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestMergedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestMergedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestMergedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestMergedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestMergedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestMergedDateObject(o.Rfc3339, isavro, false, "string"),
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestUpdatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestUpdatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequestUpdatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *PullRequestUpdatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestUpdatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toPullRequestUpdatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestUpdatedDateObject(o.Rfc3339, isavro, false, "string"),
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
	// BranchID the unique id for the branch
	BranchID string `json:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// ClosedByRefID the id of user who closed the pull request
	ClosedByRefID string `json:"closed_by_ref_id" bson:"closed_by_ref_id" yaml:"closed_by_ref_id" faker:"-"`
	// ClosedDate the timestamp in UTC that the pull request was closed
	ClosedDate PullRequestClosedDate `json:"closed_date" bson:"closed_date" yaml:"closed_date" faker:"-"`
	// CreatedByRefID the user ref_id in the source system
	CreatedByRefID string `json:"created_by_ref_id" bson:"created_by_ref_id" yaml:"created_by_ref_id" faker:"-"`
	// CreatedDate the timestamp in UTC that the pull request was created
	CreatedDate PullRequestCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the pull request
	Description string `json:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// MergeSha the sha of the merge commit
	MergeSha string `json:"merge_sha" bson:"merge_sha" yaml:"merge_sha" faker:"-"`
	// MergedByRefID the id of user who merged the pull request
	MergedByRefID string `json:"merged_by_ref_id" bson:"merged_by_ref_id" yaml:"merged_by_ref_id" faker:"-"`
	// MergedDate the timestamp in UTC that the pull request was merged
	MergedDate PullRequestMergedDate `json:"merged_date" bson:"merged_date" yaml:"merged_date" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Status the status of the pull request
	Status PullRequestStatus `json:"status" bson:"status" yaml:"status" faker:"-"`
	// Title the title of the pull request
	Title string `json:"title" bson:"title" yaml:"title" faker:"commit_message"`
	// UpdatedDate the timestamp in UTC that the pull request was closed
	UpdatedDate PullRequestUpdatedDate `json:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the pull request home page
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequest)(nil)

func toPullRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *PullRequest:
		return v.ToMap(isavro)

	case PullRequestClosedDate:
		return v.ToMap(isavro)

	case PullRequestCreatedDate:
		return v.ToMap(isavro)

	case PullRequestMergedDate:
		return v.ToMap(isavro)

	case PullRequestStatus:
		return v.String()

	case PullRequestUpdatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
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

// GetModelName returns the name of the model
func (o *PullRequest) GetModelName() datamodel.ModelNameType {
	return PullRequestModelName
}

func (o *PullRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequest", o.CustomerID, o.RefType, o.GetRefID())
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
	var dt interface{} = o.UpdatedAt
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
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
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *PullRequest) GetStateKey() string {
	key := "repo_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *PullRequest) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *PullRequest) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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

var cachedCodecPullRequest *goavro.Codec
var cachedCodecPullRequestLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *PullRequest) GetAvroCodec() *goavro.Codec {
	cachedCodecPullRequestLock.Lock()
	if cachedCodecPullRequest == nil {
		c, err := GetPullRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullRequest = c
	}
	cachedCodecPullRequestLock.Unlock()
	return cachedCodecPullRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *PullRequest) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
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
func (o *PullRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"branch_id":         toPullRequestObject(o.BranchID, isavro, false, "string"),
		"closed_by_ref_id":  toPullRequestObject(o.ClosedByRefID, isavro, false, "string"),
		"closed_date":       toPullRequestObject(o.ClosedDate, isavro, false, "closed_date"),
		"created_by_ref_id": toPullRequestObject(o.CreatedByRefID, isavro, false, "string"),
		"created_date":      toPullRequestObject(o.CreatedDate, isavro, false, "created_date"),
		"customer_id":       toPullRequestObject(o.CustomerID, isavro, false, "string"),
		"description":       toPullRequestObject(o.Description, isavro, false, "string"),
		"id":                toPullRequestObject(o.ID, isavro, false, "string"),
		"merge_sha":         toPullRequestObject(o.MergeSha, isavro, false, "string"),
		"merged_by_ref_id":  toPullRequestObject(o.MergedByRefID, isavro, false, "string"),
		"merged_date":       toPullRequestObject(o.MergedDate, isavro, false, "merged_date"),
		"ref_id":            toPullRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":          toPullRequestObject(o.RefType, isavro, false, "string"),
		"repo_id":           toPullRequestObject(o.RepoID, isavro, false, "string"),
		"status":            toPullRequestObject(o.Status, isavro, false, "status"),
		"title":             toPullRequestObject(o.Title, isavro, false, "string"),
		"updated_date":      toPullRequestObject(o.UpdatedDate, isavro, false, "updated_date"),
		"updated_ts":        toPullRequestObject(o.UpdatedAt, isavro, false, "long"),
		"url":               toPullRequestObject(o.URL, isavro, false, "string"),
		"hashcode":          toPullRequestObject(o.Hashcode, isavro, false, "string"),
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
	args = append(args, o.ClosedByRefID)
	args = append(args, o.ClosedDate)
	args = append(args, o.CreatedByRefID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
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

// GetPullRequestAvroSchemaSpec creates the avro schema specification for PullRequest
func GetPullRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "PullRequest",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "branch_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "closed_by_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "closed_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the pull request was closed", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "closed_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "created_by_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the pull request was created", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "created_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merge_sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merged_by_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "merged_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the pull request was merged", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "merged_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "status",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "status",
					"symbols": []interface{}{"OPEN", "CLOSED", "MERGED", "SUPERSEDED", "LOCKED"},
				},
			},
			map[string]interface{}{
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "updated_date",
				"type": map[string]interface{}{"doc": "the timestamp in UTC that the pull request was closed", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "updated_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
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

// GetPullRequestAvroSchema creates the avro schema for PullRequest
func GetPullRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullRequestAvroSchemaSpec())
}

// TransformPullRequestFunc is a function for transforming PullRequest during processing
type TransformPullRequestFunc func(input *PullRequest) (*PullRequest, error)

// NewPullRequestPipe creates a pipe for processing PullRequest items
func NewPullRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewPullRequestInputStream(input, errors)
	var stream chan PullRequest
	if len(transforms) > 0 {
		stream = make(chan PullRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewPullRequestOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewPullRequestInputStreamDir creates a channel for reading PullRequest as JSON newlines from a directory of files
func NewPullRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/pull_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request")
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewPullRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewPullRequestInputStreamFile creates an channel for reading PullRequest as JSON newlines from filename
func NewPullRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan PullRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewPullRequestInputStream(f, errors, transforms...)
}

// NewPullRequestInputStream creates an channel for reading PullRequest as JSON newlines from stream
func NewPullRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequest, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item PullRequest
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewPullRequestOutputStreamDir will output json newlines from channel and save in dir
func NewPullRequestOutputStreamDir(dir string, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/pull_request\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewPullRequestOutputStream(gz, ch, errors, transforms...)
}

// NewPullRequestOutputStream will output json newlines from channel to the stream
func NewPullRequestOutputStream(stream io.WriteCloser, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// PullRequestSendEvent is an event detail for sending data
type PullRequestSendEvent struct {
	PullRequest *PullRequest
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*PullRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *PullRequestSendEvent) Key() string {
	if e.key == "" {
		return e.PullRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *PullRequestSendEvent) Object() datamodel.Model {
	return e.PullRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *PullRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *PullRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// PullRequestSendEventOpts is a function handler for setting opts
type PullRequestSendEventOpts func(o *PullRequestSendEvent)

// WithPullRequestSendEventKey sets the key value to a value different than the object ID
func WithPullRequestSendEventKey(key string) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		o.key = key
	}
}

// WithPullRequestSendEventTimestamp sets the timestamp value
func WithPullRequestSendEventTimestamp(tv time.Time) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		o.time = tv
	}
}

// WithPullRequestSendEventHeader sets the timestamp value
func WithPullRequestSendEventHeader(key, value string) PullRequestSendEventOpts {
	return func(o *PullRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewPullRequestSendEvent returns a new PullRequestSendEvent instance
func NewPullRequestSendEvent(o *PullRequest, opts ...PullRequestSendEventOpts) *PullRequestSendEvent {
	res := &PullRequestSendEvent{
		PullRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewPullRequestProducer will stream data from the channel
func NewPullRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*PullRequest); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.PullRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewPullRequestConsumer will stream data from the topic into the provided channel
func NewPullRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object PullRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.PullRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into sourcecode.PullRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.PullRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &PullRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object PullRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// PullRequestReceiveEvent is an event detail for receiving data
type PullRequestReceiveEvent struct {
	PullRequest *PullRequest
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*PullRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *PullRequestReceiveEvent) Object() datamodel.Model {
	return e.PullRequest
}

// Message returns the underlying message data for the event
func (e *PullRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *PullRequestReceiveEvent) EOF() bool {
	return e.eof
}

// PullRequestProducer implements the datamodel.ModelEventProducer
type PullRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*PullRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *PullRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *PullRequestProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *PullRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *PullRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewPullRequestProducerChannel returns a channel which can be used for producing Model events
func NewPullRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewPullRequestProducerChannelSize(producer, 0, errors)
}

// NewPullRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewPullRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &PullRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewPullRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// PullRequestConsumer implements the datamodel.ModelEventConsumer
type PullRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*PullRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *PullRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *PullRequestConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *PullRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullRequestConsumer{
		ch:       ch,
		callback: NewPullRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewPullRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewPullRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &PullRequestConsumer{
		ch:       ch,
		callback: NewPullRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
