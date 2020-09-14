// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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

	// IssueCommentTable is the default table name
	IssueCommentTable datamodel.ModelNameType = "work_issuecomment"

	// IssueCommentModelName is the model name
	IssueCommentModelName datamodel.ModelNameType = "work.IssueComment"
)

const (
	// IssueCommentModelActiveColumn is the column json value active
	IssueCommentModelActiveColumn = "active"
	// IssueCommentModelBodyColumn is the column json value body
	IssueCommentModelBodyColumn = "body"
	// IssueCommentModelCreatedDateColumn is the column json value created_date
	IssueCommentModelCreatedDateColumn = "created_date"
	// IssueCommentModelCreatedDateEpochColumn is the column json value epoch
	IssueCommentModelCreatedDateEpochColumn = "epoch"
	// IssueCommentModelCreatedDateOffsetColumn is the column json value offset
	IssueCommentModelCreatedDateOffsetColumn = "offset"
	// IssueCommentModelCreatedDateRfc3339Column is the column json value rfc3339
	IssueCommentModelCreatedDateRfc3339Column = "rfc3339"
	// IssueCommentModelCustomerIDColumn is the column json value customer_id
	IssueCommentModelCustomerIDColumn = "customer_id"
	// IssueCommentModelIDColumn is the column json value id
	IssueCommentModelIDColumn = "id"
	// IssueCommentModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IssueCommentModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IssueCommentModelIssueIDColumn is the column json value issue_id
	IssueCommentModelIssueIDColumn = "issue_id"
	// IssueCommentModelProjectIDColumn is the column json value project_id
	IssueCommentModelProjectIDColumn = "project_id"
	// IssueCommentModelRefIDColumn is the column json value ref_id
	IssueCommentModelRefIDColumn = "ref_id"
	// IssueCommentModelRefTypeColumn is the column json value ref_type
	IssueCommentModelRefTypeColumn = "ref_type"
	// IssueCommentModelUpdatedDateColumn is the column json value updated_date
	IssueCommentModelUpdatedDateColumn = "updated_date"
	// IssueCommentModelUpdatedDateEpochColumn is the column json value epoch
	IssueCommentModelUpdatedDateEpochColumn = "epoch"
	// IssueCommentModelUpdatedDateOffsetColumn is the column json value offset
	IssueCommentModelUpdatedDateOffsetColumn = "offset"
	// IssueCommentModelUpdatedDateRfc3339Column is the column json value rfc3339
	IssueCommentModelUpdatedDateRfc3339Column = "rfc3339"
	// IssueCommentModelURLColumn is the column json value url
	IssueCommentModelURLColumn = "url"
	// IssueCommentModelUserRefIDColumn is the column json value user_ref_id
	IssueCommentModelUserRefIDColumn = "user_ref_id"
)

// IssueCommentCreatedDate represents the object structure for created_date
type IssueCommentCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueCommentCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueCommentCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IssueCommentCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueCommentCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueCommentCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueCommentCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IssueCommentCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueCommentCreatedDate) FromMap(kv map[string]interface{}) {

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

// IssueCommentUpdatedDate represents the object structure for updated_date
type IssueCommentUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIssueCommentUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IssueCommentUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIssueCommentUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIssueCommentUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIssueCommentUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *IssueCommentUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueCommentUpdatedDate) FromMap(kv map[string]interface{}) {

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

// IssueComment the comment on issue
type IssueComment struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Body the body of the comment
	Body string `json:"body" codec:"body" bson:"body" yaml:"body" faker:"commit_message"`
	// CreatedDate the timestamp that the comment was created
	CreatedDate IssueCommentCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IssueID the issue this comment is associated with
	IssueID string `json:"issue_id" codec:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// ProjectID unique project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedDate the timestamp that the comment was updated
	UpdatedDate IssueCommentUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// URL the URL to the source system for this comment
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// UserRefID the user ref_id in the source system of the user that made the comment
	UserRefID string `json:"user_ref_id" codec:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IssueComment)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IssueComment)(nil)

func toIssueCommentObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueComment:
		return v.ToMap()

	case IssueCommentCreatedDate:
		return v.ToMap()

	case IssueCommentUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of IssueComment
func (o *IssueComment) String() string {
	return fmt.Sprintf("work.IssueComment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IssueComment) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IssueComment) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IssueComment) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IssueComment) GetModelName() datamodel.ModelNameType {
	return IssueCommentModelName
}

// NewIssueCommentID provides a template for generating an ID field for IssueComment
func NewIssueCommentID(customerID string, refID string, refType string, ProjectID string) string {
	return hash.Values(customerID, refID, refType, ProjectID)
}

func (o *IssueComment) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType, o.ProjectID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IssueComment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IssueComment) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IssueComment) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IssueComment) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IssueComment) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IssueComment) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IssueComment) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IssueComment) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IssueComment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IssueCommentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IssueComment) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IssueComment) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *IssueComment) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *IssueComment) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *IssueComment) SetRefType(t string) {

	o.RefType = id

}

// Clone returns an exact copy of IssueComment
func (o *IssueComment) Clone() datamodel.Model {
	c := new(IssueComment)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IssueComment) Anon() datamodel.Model {
	c := new(IssueComment)
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
func (o *IssueComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssueComment) UnmarshalJSON(data []byte) error {
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
func (o *IssueComment) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IssueComment) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IssueComment objects are equal
func (o *IssueComment) IsEqual(other *IssueComment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IssueComment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toIssueCommentObject(o.Active, false),
		"body":                    toIssueCommentObject(o.Body, false),
		"created_date":            toIssueCommentObject(o.CreatedDate, false),
		"customer_id":             toIssueCommentObject(o.CustomerID, false),
		"id":                      toIssueCommentObject(o.ID, false),
		"integration_instance_id": toIssueCommentObject(o.IntegrationInstanceID, true),
		"issue_id":                toIssueCommentObject(o.IssueID, false),
		"project_id":              toIssueCommentObject(o.ProjectID, false),
		"ref_id":                  toIssueCommentObject(o.RefID, false),
		"ref_type":                toIssueCommentObject(o.RefType, false),
		"updated_date":            toIssueCommentObject(o.UpdatedDate, false),
		"url":                     toIssueCommentObject(o.URL, false),
		"user_ref_id":             toIssueCommentObject(o.UserRefID, false),
		"hashcode":                toIssueCommentObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueComment) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(IssueCommentCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IssueCommentCreatedDate); ok {
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
	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		if val, ok := kv["issue_id"]; ok {
			if val == nil {
				o.IssueID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IssueID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ProjectID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueCommentUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*IssueCommentUpdatedDate); ok {
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
func (o *IssueComment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Body)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IssueID)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedDate)
	args = append(args, o.URL)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *IssueComment) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *IssueComment) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// IssueCommentPartial is a partial struct for upsert mutations for IssueComment
type IssueCommentPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// Body the body of the comment
	Body *string `json:"body,omitempty"`
	// CreatedDate the timestamp that the comment was created
	CreatedDate *IssueCommentCreatedDate `json:"created_date,omitempty"`
	// IssueID the issue this comment is associated with
	IssueID *string `json:"issue_id,omitempty"`
	// ProjectID unique project id
	ProjectID *string `json:"project_id,omitempty"`
	// UpdatedDate the timestamp that the comment was updated
	UpdatedDate *IssueCommentUpdatedDate `json:"updated_date,omitempty"`
	// URL the URL to the source system for this comment
	URL *string `json:"url,omitempty"`
	// UserRefID the user ref_id in the source system of the user that made the comment
	UserRefID *string `json:"user_ref_id,omitempty"`
}

var _ datamodel.PartialModel = (*IssueCommentPartial)(nil)

// GetModelName returns the name of the model
func (o *IssueCommentPartial) GetModelName() datamodel.ModelNameType {
	return IssueCommentModelName
}

// ToMap returns the object as a map
func (o *IssueCommentPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":       toIssueCommentObject(o.Active, true),
		"body":         toIssueCommentObject(o.Body, true),
		"created_date": toIssueCommentObject(o.CreatedDate, true),
		"issue_id":     toIssueCommentObject(o.IssueID, true),
		"project_id":   toIssueCommentObject(o.ProjectID, true),
		"updated_date": toIssueCommentObject(o.UpdatedDate, true),
		"url":          toIssueCommentObject(o.URL, true),
		"user_ref_id":  toIssueCommentObject(o.UserRefID, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "created_date" {
				if dt, ok := v.(*IssueCommentCreatedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "updated_date" {
				if dt, ok := v.(*IssueCommentUpdatedDate); ok {
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
func (o *IssueCommentPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IssueCommentPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssueCommentPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IssueCommentPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IssueCommentPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["body"].(*string); ok {
		o.Body = val
	} else if val, ok := kv["body"].(string); ok {
		o.Body = &val
	} else {
		if val, ok := kv["body"]; ok {
			if val == nil {
				o.Body = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Body = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.CreatedDate == nil {
		o.CreatedDate = &IssueCommentCreatedDate{}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueCommentCreatedDate); ok {
			// struct
			o.CreatedDate = &sv
		} else if sp, ok := val.(*IssueCommentCreatedDate); ok {
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

	if val, ok := kv["issue_id"].(*string); ok {
		o.IssueID = val
	} else if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = &val
	} else {
		if val, ok := kv["issue_id"]; ok {
			if val == nil {
				o.IssueID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IssueID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["project_id"].(*string); ok {
		o.ProjectID = val
	} else if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = &val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ProjectID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.UpdatedDate == nil {
		o.UpdatedDate = &IssueCommentUpdatedDate{}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(IssueCommentUpdatedDate); ok {
			// struct
			o.UpdatedDate = &sv
		} else if sp, ok := val.(*IssueCommentUpdatedDate); ok {
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
	if val, ok := kv["user_ref_id"].(*string); ok {
		o.UserRefID = val
	} else if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = &val
	} else {
		if val, ok := kv["user_ref_id"]; ok {
			if val == nil {
				o.UserRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UserRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
