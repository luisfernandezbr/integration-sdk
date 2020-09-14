// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// PullRequestReviewRequestTable is the default table name
	PullRequestReviewRequestTable datamodel.ModelNameType = "sourcecode_pullrequestreviewrequest"

	// PullRequestReviewRequestModelName is the model name
	PullRequestReviewRequestModelName datamodel.ModelNameType = "sourcecode.PullRequestReviewRequest"
)

const (
	// PullRequestReviewRequestModelActiveColumn is the column json value active
	PullRequestReviewRequestModelActiveColumn = "active"
	// PullRequestReviewRequestModelCreatedDateColumn is the column json value created_date
	PullRequestReviewRequestModelCreatedDateColumn = "created_date"
	// PullRequestReviewRequestModelCreatedDateEpochColumn is the column json value epoch
	PullRequestReviewRequestModelCreatedDateEpochColumn = "epoch"
	// PullRequestReviewRequestModelCreatedDateOffsetColumn is the column json value offset
	PullRequestReviewRequestModelCreatedDateOffsetColumn = "offset"
	// PullRequestReviewRequestModelCreatedDateRfc3339Column is the column json value rfc3339
	PullRequestReviewRequestModelCreatedDateRfc3339Column = "rfc3339"
	// PullRequestReviewRequestModelCustomerIDColumn is the column json value customer_id
	PullRequestReviewRequestModelCustomerIDColumn = "customer_id"
	// PullRequestReviewRequestModelIDColumn is the column json value id
	PullRequestReviewRequestModelIDColumn = "id"
	// PullRequestReviewRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PullRequestReviewRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PullRequestReviewRequestModelPullRequestIDColumn is the column json value pull_request_id
	PullRequestReviewRequestModelPullRequestIDColumn = "pull_request_id"
	// PullRequestReviewRequestModelRefIDColumn is the column json value ref_id
	PullRequestReviewRequestModelRefIDColumn = "ref_id"
	// PullRequestReviewRequestModelRefTypeColumn is the column json value ref_type
	PullRequestReviewRequestModelRefTypeColumn = "ref_type"
	// PullRequestReviewRequestModelRepoIDColumn is the column json value repo_id
	PullRequestReviewRequestModelRepoIDColumn = "repo_id"
	// PullRequestReviewRequestModelRequestedReviewerRefIDColumn is the column json value requested_reviewer_ref_id
	PullRequestReviewRequestModelRequestedReviewerRefIDColumn = "requested_reviewer_ref_id"
	// PullRequestReviewRequestModelSenderRefIDColumn is the column json value sender_ref_id
	PullRequestReviewRequestModelSenderRefIDColumn = "sender_ref_id"
	// PullRequestReviewRequestModelURLColumn is the column json value url
	PullRequestReviewRequestModelURLColumn = "url"
)

// PullRequestReviewRequestCreatedDate represents the object structure for created_date
type PullRequestReviewRequestCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestReviewRequestCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestReviewRequestCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestReviewRequestCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestReviewRequestCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestReviewRequestCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestReviewRequestCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestReviewRequestCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReviewRequestCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestReviewRequest a request for a user to review a given pull request
type PullRequestReviewRequest struct {
	// Active indicates that this model is displayed in a source system, false if the request has been cacelled
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// CreatedDate the timestamp in UTC that the request was created
	CreatedDate PullRequestReviewRequestCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// PullRequestID the pull request this request is associated with
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// RequestedReviewerRefID the source system id of the user who's review is desired
	RequestedReviewerRefID string `json:"requested_reviewer_ref_id" codec:"requested_reviewer_ref_id" bson:"requested_reviewer_ref_id" yaml:"requested_reviewer_ref_id" faker:"-"`
	// SenderRefID the source system id of the user who made the request
	SenderRefID string `json:"sender_ref_id" codec:"sender_ref_id" bson:"sender_ref_id" yaml:"sender_ref_id" faker:"-"`
	// URL the URL to the source system for this review request
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestReviewRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequestReviewRequest)(nil)

func toPullRequestReviewRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestReviewRequest:
		return v.ToMap()

	case PullRequestReviewRequestCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequestReviewRequest
func (o *PullRequestReviewRequest) String() string {
	return fmt.Sprintf("sourcecode.PullRequestReviewRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestReviewRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *PullRequestReviewRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequestReviewRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestReviewRequest) GetModelName() datamodel.ModelNameType {
	return PullRequestReviewRequestModelName
}

// NewPullRequestReviewRequestID provides a template for generating an ID field for PullRequestReviewRequest
func NewPullRequestReviewRequestID(customerID string, refType string, PullRequestID string, RequestedReviewerRefID string) string {
	return hash.Values(customerID, refType, PullRequestID, RequestedReviewerRefID)
}

func (o *PullRequestReviewRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefType, o.PullRequestID, o.RequestedReviewerRefID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequestReviewRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestReviewRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestReviewRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequestReviewRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestReviewRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequestReviewRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestReviewRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestReviewRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestReviewRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestReviewRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestReviewRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *PullRequestReviewRequest) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *PullRequestReviewRequest) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *PullRequestReviewRequest) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *PullRequestReviewRequest) SetRefType(t string) {

	o.RefType = t

}

// Clone returns an exact copy of PullRequestReviewRequest
func (o *PullRequestReviewRequest) Clone() datamodel.Model {
	c := new(PullRequestReviewRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestReviewRequest) Anon() datamodel.Model {
	c := new(PullRequestReviewRequest)
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
func (o *PullRequestReviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestReviewRequest) UnmarshalJSON(data []byte) error {
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
func (o *PullRequestReviewRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *PullRequestReviewRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two PullRequestReviewRequest objects are equal
func (o *PullRequestReviewRequest) IsEqual(other *PullRequestReviewRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestReviewRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                    toPullRequestReviewRequestObject(o.Active, false),
		"created_date":              toPullRequestReviewRequestObject(o.CreatedDate, false),
		"customer_id":               toPullRequestReviewRequestObject(o.CustomerID, false),
		"id":                        toPullRequestReviewRequestObject(o.ID, false),
		"integration_instance_id":   toPullRequestReviewRequestObject(o.IntegrationInstanceID, true),
		"pull_request_id":           toPullRequestReviewRequestObject(o.PullRequestID, false),
		"ref_id":                    toPullRequestReviewRequestObject(o.RefID, false),
		"ref_type":                  toPullRequestReviewRequestObject(o.RefType, false),
		"repo_id":                   toPullRequestReviewRequestObject(o.RepoID, false),
		"requested_reviewer_ref_id": toPullRequestReviewRequestObject(o.RequestedReviewerRefID, false),
		"sender_ref_id":             toPullRequestReviewRequestObject(o.SenderRefID, false),
		"url":                       toPullRequestReviewRequestObject(o.URL, false),
		"hashcode":                  toPullRequestReviewRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReviewRequest) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestReviewRequestCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestReviewRequestCreatedDate); ok {
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
	if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.PullRequestID = fmt.Sprintf("%v", val)
			}
		}
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
	if val, ok := kv["requested_reviewer_ref_id"].(string); ok {
		o.RequestedReviewerRefID = val
	} else {
		if val, ok := kv["requested_reviewer_ref_id"]; ok {
			if val == nil {
				o.RequestedReviewerRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RequestedReviewerRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["sender_ref_id"].(string); ok {
		o.SenderRefID = val
	} else {
		if val, ok := kv["sender_ref_id"]; ok {
			if val == nil {
				o.SenderRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SenderRefID = fmt.Sprintf("%v", val)
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
func (o *PullRequestReviewRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.RequestedReviewerRefID)
	args = append(args, o.SenderRefID)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *PullRequestReviewRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *PullRequestReviewRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// PullRequestReviewRequestPartial is a partial struct for upsert mutations for PullRequestReviewRequest
type PullRequestReviewRequestPartial struct {
	// Active indicates that this model is displayed in a source system, false if the request has been cacelled
	Active *bool `json:"active,omitempty"`
	// CreatedDate the timestamp in UTC that the request was created
	CreatedDate *PullRequestReviewRequestCreatedDate `json:"created_date,omitempty"`
	// PullRequestID the pull request this request is associated with
	PullRequestID *string `json:"pull_request_id,omitempty"`
	// RepoID the unique id for the repo
	RepoID *string `json:"repo_id,omitempty"`
	// RequestedReviewerRefID the source system id of the user who's review is desired
	RequestedReviewerRefID *string `json:"requested_reviewer_ref_id,omitempty"`
	// SenderRefID the source system id of the user who made the request
	SenderRefID *string `json:"sender_ref_id,omitempty"`
	// URL the URL to the source system for this review request
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*PullRequestReviewRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *PullRequestReviewRequestPartial) GetModelName() datamodel.ModelNameType {
	return PullRequestReviewRequestModelName
}

// ToMap returns the object as a map
func (o *PullRequestReviewRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":                    toPullRequestReviewRequestObject(o.Active, true),
		"created_date":              toPullRequestReviewRequestObject(o.CreatedDate, true),
		"pull_request_id":           toPullRequestReviewRequestObject(o.PullRequestID, true),
		"repo_id":                   toPullRequestReviewRequestObject(o.RepoID, true),
		"requested_reviewer_ref_id": toPullRequestReviewRequestObject(o.RequestedReviewerRefID, true),
		"sender_ref_id":             toPullRequestReviewRequestObject(o.SenderRefID, true),
		"url":                       toPullRequestReviewRequestObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*PullRequestReviewRequestCreatedDate); ok {
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
func (o *PullRequestReviewRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequestReviewRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestReviewRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *PullRequestReviewRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReviewRequestPartial) FromMap(kv map[string]interface{}) {
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

	if o.CreatedDate == nil {
		o.CreatedDate = &PullRequestReviewRequestCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestReviewRequestCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*PullRequestReviewRequestCreatedDate); ok {
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

	if val, ok := kv["pull_request_id"].(*string); ok {
		o.PullRequestID = val
	} else if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = &val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PullRequestID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
	if val, ok := kv["requested_reviewer_ref_id"].(*string); ok {
		o.RequestedReviewerRefID = val
	} else if val, ok := kv["requested_reviewer_ref_id"].(string); ok {
		o.RequestedReviewerRefID = &val
	} else {
		if val, ok := kv["requested_reviewer_ref_id"]; ok {
			if val == nil {
				o.RequestedReviewerRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RequestedReviewerRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["sender_ref_id"].(*string); ok {
		o.SenderRefID = val
	} else if val, ok := kv["sender_ref_id"].(string); ok {
		o.SenderRefID = &val
	} else {
		if val, ok := kv["sender_ref_id"]; ok {
			if val == nil {
				o.SenderRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.SenderRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
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
