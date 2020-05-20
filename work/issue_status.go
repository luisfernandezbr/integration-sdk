// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
)

const (

	// IssueStatusTable is the default table name
	IssueStatusTable datamodel.ModelNameType = "work_issuestatus"

	// IssueStatusModelName is the model name
	IssueStatusModelName datamodel.ModelNameType = "work.IssueStatus"
)

const (
	// IssueStatusModelCustomerIDColumn is the column json value customer_id
	IssueStatusModelCustomerIDColumn = "customer_id"
	// IssueStatusModelDescriptionColumn is the column json value description
	IssueStatusModelDescriptionColumn = "description"
	// IssueStatusModelIconURLColumn is the column json value icon_url
	IssueStatusModelIconURLColumn = "icon_url"
	// IssueStatusModelIDColumn is the column json value id
	IssueStatusModelIDColumn = "id"
	// IssueStatusModelNameColumn is the column json value name
	IssueStatusModelNameColumn = "name"
	// IssueStatusModelRefIDColumn is the column json value ref_id
	IssueStatusModelRefIDColumn = "ref_id"
	// IssueStatusModelRefTypeColumn is the column json value ref_type
	IssueStatusModelRefTypeColumn = "ref_type"
)

// IssueStatus the status type of an issue
type IssueStatus struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the status
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// IconURL an optional url for the icon for the status
	IconURL *string `json:"icon_url,omitempty" codec:"icon_url,omitempty" bson:"icon_url" yaml:"icon_url,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the status
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IssueStatus)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IssueStatus)(nil)

func toIssueStatusObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueStatus:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of IssueStatus
func (o *IssueStatus) String() string {
	return fmt.Sprintf("work.IssueStatus<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IssueStatus) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IssueStatus) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IssueStatus) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IssueStatus) GetModelName() datamodel.ModelNameType {
	return IssueStatusModelName
}

// NewIssueStatusID provides a template for generating an ID field for IssueStatus
func NewIssueStatusID(customerID string, refType string, refID string) string {
	return hash.Values(customerID, refType, refID)
}

func (o *IssueStatus) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefType, o.RefID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IssueStatus) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IssueStatus) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IssueStatus) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IssueStatus) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IssueStatus) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IssueStatus) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IssueStatus) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IssueStatus) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IssueStatus) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IssueStatusModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IssueStatus) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IssueStatus) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IssueStatus
func (o *IssueStatus) Clone() datamodel.Model {
	c := new(IssueStatus)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IssueStatus) Anon() datamodel.Model {
	c := new(IssueStatus)
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
func (o *IssueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssueStatus) UnmarshalJSON(data []byte) error {
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
func (o *IssueStatus) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IssueStatus) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IssueStatus objects are equal
func (o *IssueStatus) IsEqual(other *IssueStatus) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IssueStatus) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toIssueStatusObject(o.CustomerID, false),
		"description": toIssueStatusObject(o.Description, false),
		"icon_url":    toIssueStatusObject(o.IconURL, true),
		"id":          toIssueStatusObject(o.ID, false),
		"name":        toIssueStatusObject(o.Name, false),
		"ref_id":      toIssueStatusObject(o.RefID, false),
		"ref_type":    toIssueStatusObject(o.RefType, false),
		"hashcode":    toIssueStatusObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueStatus) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	if val, ok := kv["icon_url"].(*string); ok {
		o.IconURL = val
	} else if val, ok := kv["icon_url"].(string); ok {
		o.IconURL = &val
	} else {
		if val, ok := kv["icon_url"]; ok {
			if val == nil {
				o.IconURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IconURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IssueStatus) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.IconURL)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
