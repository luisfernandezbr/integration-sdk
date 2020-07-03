// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
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

	// PullRequestCommentTable is the default table name
	PullRequestCommentTable datamodel.ModelNameType = "sourcecode_pullrequestcomment"

	// PullRequestCommentModelName is the model name
	PullRequestCommentModelName datamodel.ModelNameType = "sourcecode.PullRequestComment"
)

const (
	// PullRequestCommentModelBodyColumn is the column json value body
	PullRequestCommentModelBodyColumn = "body"
	// PullRequestCommentModelCreatedDateColumn is the column json value created_date
	PullRequestCommentModelCreatedDateColumn = "created_date"
	// PullRequestCommentModelCreatedDateEpochColumn is the column json value epoch
	PullRequestCommentModelCreatedDateEpochColumn = "epoch"
	// PullRequestCommentModelCreatedDateOffsetColumn is the column json value offset
	PullRequestCommentModelCreatedDateOffsetColumn = "offset"
	// PullRequestCommentModelCreatedDateRfc3339Column is the column json value rfc3339
	PullRequestCommentModelCreatedDateRfc3339Column = "rfc3339"
	// PullRequestCommentModelCustomerIDColumn is the column json value customer_id
	PullRequestCommentModelCustomerIDColumn = "customer_id"
	// PullRequestCommentModelIDColumn is the column json value id
	PullRequestCommentModelIDColumn = "id"
	// PullRequestCommentModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PullRequestCommentModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PullRequestCommentModelPullRequestIDColumn is the column json value pull_request_id
	PullRequestCommentModelPullRequestIDColumn = "pull_request_id"
	// PullRequestCommentModelRefIDColumn is the column json value ref_id
	PullRequestCommentModelRefIDColumn = "ref_id"
	// PullRequestCommentModelRefTypeColumn is the column json value ref_type
	PullRequestCommentModelRefTypeColumn = "ref_type"
	// PullRequestCommentModelRepoIDColumn is the column json value repo_id
	PullRequestCommentModelRepoIDColumn = "repo_id"
	// PullRequestCommentModelUpdatedDateColumn is the column json value updated_date
	PullRequestCommentModelUpdatedDateColumn = "updated_date"
	// PullRequestCommentModelUpdatedDateEpochColumn is the column json value epoch
	PullRequestCommentModelUpdatedDateEpochColumn = "epoch"
	// PullRequestCommentModelUpdatedDateOffsetColumn is the column json value offset
	PullRequestCommentModelUpdatedDateOffsetColumn = "offset"
	// PullRequestCommentModelUpdatedDateRfc3339Column is the column json value rfc3339
	PullRequestCommentModelUpdatedDateRfc3339Column = "rfc3339"
	// PullRequestCommentModelURLColumn is the column json value url
	PullRequestCommentModelURLColumn = "url"
	// PullRequestCommentModelUserRefIDColumn is the column json value user_ref_id
	PullRequestCommentModelUserRefIDColumn = "user_ref_id"
)

// PullRequestCommentCreatedDate represents the object structure for created_date
type PullRequestCommentCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommentCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestCommentCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestCommentCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommentCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestCommentUpdatedDate represents the object structure for updated_date
type PullRequestCommentUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommentUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestCommentUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommentUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommentUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommentUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestCommentUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommentUpdatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestComment the comment for a given pull request
type PullRequestComment struct {
	// Body the body of the comment
	Body string `json:"body" codec:"body" bson:"body" yaml:"body" faker:"commit_message"`
	// CreatedDate the timestamp in UTC that the comment was created
	CreatedDate PullRequestCommentCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// PullRequestID the pull request this comment is associated with
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedDate the timestamp in UTC that the comment was closed
	UpdatedDate PullRequestCommentUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// URL the URL to the source system for this comment
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestComment)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequestComment)(nil)

func toPullRequestCommentObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestComment:
		return v.ToMap()

	case PullRequestCommentCreatedDate:
		return v.ToMap()

	case PullRequestCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequestComment
func (o *PullRequestComment) String() string {
	return fmt.Sprintf("sourcecode.PullRequestComment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestComment) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *PullRequestComment) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequestComment) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestComment) GetModelName() datamodel.ModelNameType {
	return PullRequestCommentModelName
}

// NewPullRequestCommentID provides a template for generating an ID field for PullRequestComment
func NewPullRequestCommentID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestComment) setDefaults(frommap bool) {

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
func (o *PullRequestComment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestComment) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestComment) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequestComment) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestComment) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequestComment) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestComment) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestComment) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestComment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestCommentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestComment) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *PullRequestComment) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestComment
func (o *PullRequestComment) Clone() datamodel.Model {
	c := new(PullRequestComment)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestComment) Anon() datamodel.Model {
	c := new(PullRequestComment)
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
func (o *PullRequestComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestComment) UnmarshalJSON(data []byte) error {
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
func (o *PullRequestComment) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *PullRequestComment) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two PullRequestComment objects are equal
func (o *PullRequestComment) IsEqual(other *PullRequestComment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestComment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"body":                    toPullRequestCommentObject(o.Body, false),
		"created_date":            toPullRequestCommentObject(o.CreatedDate, false),
		"customer_id":             toPullRequestCommentObject(o.CustomerID, false),
		"id":                      toPullRequestCommentObject(o.ID, false),
		"integration_instance_id": toPullRequestCommentObject(o.IntegrationInstanceID, true),
		"pull_request_id":         toPullRequestCommentObject(o.PullRequestID, false),
		"ref_id":                  toPullRequestCommentObject(o.RefID, false),
		"ref_type":                toPullRequestCommentObject(o.RefType, false),
		"repo_id":                 toPullRequestCommentObject(o.RepoID, false),
		"updated_date":            toPullRequestCommentObject(o.UpdatedDate, false),
		"url":                     toPullRequestCommentObject(o.URL, false),
		"user_ref_id":             toPullRequestCommentObject(o.UserRefID, false),
		"hashcode":                toPullRequestCommentObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestComment) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["body"].(string); ok {
		o.Body = val
	} else {
		if val, ok := kv["body"]; ok {
			if val == nil {
				o.Body = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Body = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommentCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestCommentCreatedDate); ok {
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

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommentUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*PullRequestCommentUpdatedDate); ok {
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
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
			if val == nil {
				o.UserRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UserRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *PullRequestComment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Body)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedDate)
	args = append(args, o.URL)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// PullRequestCommentPartial is a partial struct for upsert mutations for PullRequestComment
type PullRequestCommentPartial struct {
	// Body the body of the comment
	Body *string `json:"body,omitempty"`
	// CreatedDate the timestamp in UTC that the comment was created
	CreatedDate *PullRequestCommentCreatedDate `json:"created_date,omitempty"`
	// PullRequestID the pull request this comment is associated with
	PullRequestID *string `json:"pull_request_id,omitempty"`
	// RepoID the unique id for the repo
	RepoID *string `json:"repo_id,omitempty"`
	// UpdatedDate the timestamp in UTC that the comment was closed
	UpdatedDate *PullRequestCommentUpdatedDate `json:"updated_date,omitempty"`
	// URL the URL to the source system for this comment
	URL *string `json:"url,omitempty"`
	// UserRefID the user ref_id in the source system
	UserRefID *string `json:"user_ref_id,omitempty"`
}

var _ datamodel.PartialModel = (*PullRequestCommentPartial)(nil)

// GetModelName returns the name of the model
func (o *PullRequestCommentPartial) GetModelName() datamodel.ModelNameType {
	return PullRequestCommentModelName
}

// ToMap returns the object as a map
func (o *PullRequestCommentPartial) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"body":            toPullRequestCommentObject(o.Body, true),
		"created_date":    toPullRequestCommentObject(o.CreatedDate, true),
		"pull_request_id": toPullRequestCommentObject(o.PullRequestID, true),
		"repo_id":         toPullRequestCommentObject(o.RepoID, true),
		"updated_date":    toPullRequestCommentObject(o.UpdatedDate, true),
		"url":             toPullRequestCommentObject(o.URL, true),
		"user_ref_id":     toPullRequestCommentObject(o.UserRefID, true),
	}
}

// Stringify returns the object in JSON format as a string
func (o *PullRequestCommentPartial) Stringify() string {
	return pjson.Stringify(o)
}
