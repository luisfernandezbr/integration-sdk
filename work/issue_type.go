// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// IssueTypeTable is the default table name
	IssueTypeTable datamodel.ModelNameType = "work_issuetype"

	// IssueTypeModelName is the model name
	IssueTypeModelName datamodel.ModelNameType = "work.IssueType"
)

const (
	// IssueTypeModelCustomerIDColumn is the column json value customer_id
	IssueTypeModelCustomerIDColumn = "customer_id"
	// IssueTypeModelDescriptionColumn is the column json value description
	IssueTypeModelDescriptionColumn = "description"
	// IssueTypeModelIconURLColumn is the column json value icon_url
	IssueTypeModelIconURLColumn = "icon_url"
	// IssueTypeModelIDColumn is the column json value id
	IssueTypeModelIDColumn = "id"
	// IssueTypeModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IssueTypeModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IssueTypeModelMappedTypeColumn is the column json value mapped_type
	IssueTypeModelMappedTypeColumn = "mapped_type"
	// IssueTypeModelNameColumn is the column json value name
	IssueTypeModelNameColumn = "name"
	// IssueTypeModelRefIDColumn is the column json value ref_id
	IssueTypeModelRefIDColumn = "ref_id"
	// IssueTypeModelRefTypeColumn is the column json value ref_type
	IssueTypeModelRefTypeColumn = "ref_type"
)

// IssueTypeMappedType is the enumeration type for mapped_type
type IssueTypeMappedType int32

// UnmarshalBSONValue for unmarshaling value
func (v *IssueTypeMappedType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IssueTypeMappedType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "UNKNOWN":
			*v = IssueTypeMappedType(0)
		case "FEATURE":
			*v = IssueTypeMappedType(1)
		case "BUG":
			*v = IssueTypeMappedType(2)
		case "ENHANCEMENT":
			*v = IssueTypeMappedType(3)
		case "EPIC":
			*v = IssueTypeMappedType(4)
		case "STORY":
			*v = IssueTypeMappedType(5)
		case "TASK":
			*v = IssueTypeMappedType(6)
		case "SUBTASK":
			*v = IssueTypeMappedType(7)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IssueTypeMappedType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "UNKNOWN":
		v = 0
	case "FEATURE":
		v = 1
	case "BUG":
		v = 2
	case "ENHANCEMENT":
		v = 3
	case "EPIC":
		v = 4
	case "STORY":
		v = 5
	case "TASK":
		v = 6
	case "SUBTASK":
		v = 7
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IssueTypeMappedType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("UNKNOWN")
	case 1:
		return json.Marshal("FEATURE")
	case 2:
		return json.Marshal("BUG")
	case 3:
		return json.Marshal("ENHANCEMENT")
	case 4:
		return json.Marshal("EPIC")
	case 5:
		return json.Marshal("STORY")
	case 6:
		return json.Marshal("TASK")
	case 7:
		return json.Marshal("SUBTASK")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for MappedType
func (v IssueTypeMappedType) String() string {
	switch int32(v) {
	case 0:
		return "UNKNOWN"
	case 1:
		return "FEATURE"
	case 2:
		return "BUG"
	case 3:
		return "ENHANCEMENT"
	case 4:
		return "EPIC"
	case 5:
		return "STORY"
	case 6:
		return "TASK"
	case 7:
		return "SUBTASK"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IssueTypeMappedType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IssueTypeMappedType:
		*v = val
	case int32:
		*v = IssueTypeMappedType(int32(val))
	case int:
		*v = IssueTypeMappedType(int32(val))
	case string:
		switch val {
		case "UNKNOWN":
			*v = IssueTypeMappedType(0)
		case "FEATURE":
			*v = IssueTypeMappedType(1)
		case "BUG":
			*v = IssueTypeMappedType(2)
		case "ENHANCEMENT":
			*v = IssueTypeMappedType(3)
		case "EPIC":
			*v = IssueTypeMappedType(4)
		case "STORY":
			*v = IssueTypeMappedType(5)
		case "TASK":
			*v = IssueTypeMappedType(6)
		case "SUBTASK":
			*v = IssueTypeMappedType(7)
		}
	}
	return nil
}

const (
	// IssueTypeMappedTypeUnknown is the enumeration value for unknown
	IssueTypeMappedTypeUnknown IssueTypeMappedType = 0
	// IssueTypeMappedTypeFeature is the enumeration value for feature
	IssueTypeMappedTypeFeature IssueTypeMappedType = 1
	// IssueTypeMappedTypeBug is the enumeration value for bug
	IssueTypeMappedTypeBug IssueTypeMappedType = 2
	// IssueTypeMappedTypeEnhancement is the enumeration value for enhancement
	IssueTypeMappedTypeEnhancement IssueTypeMappedType = 3
	// IssueTypeMappedTypeEpic is the enumeration value for epic
	IssueTypeMappedTypeEpic IssueTypeMappedType = 4
	// IssueTypeMappedTypeStory is the enumeration value for story
	IssueTypeMappedTypeStory IssueTypeMappedType = 5
	// IssueTypeMappedTypeTask is the enumeration value for task
	IssueTypeMappedTypeTask IssueTypeMappedType = 6
	// IssueTypeMappedTypeSubtask is the enumeration value for subtask
	IssueTypeMappedTypeSubtask IssueTypeMappedType = 7
)

// IssueType the issue type of an issue
type IssueType struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the issue
	Description *string `json:"description,omitempty" codec:"description,omitempty" bson:"description" yaml:"description,omitempty" faker:"-"`
	// IconURL an optional url for the icon for the issue
	IconURL *string `json:"icon_url,omitempty" codec:"icon_url,omitempty" bson:"icon_url" yaml:"icon_url,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// MappedType a default mapped type is known by the integration
	MappedType IssueTypeMappedType `json:"mapped_type" codec:"mapped_type" bson:"mapped_type" yaml:"mapped_type" faker:"-"`
	// Name the name of the issue
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IssueType)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IssueType)(nil)

func toIssueTypeObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IssueType:
		return v.ToMap()

	case IssueTypeMappedType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of IssueType
func (o *IssueType) String() string {
	return fmt.Sprintf("work.IssueType<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IssueType) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IssueType) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IssueType) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *IssueType) GetModelName() datamodel.ModelNameType {
	return IssueTypeModelName
}

// NewIssueTypeID provides a template for generating an ID field for IssueType
func NewIssueTypeID(customerID string, refType string, refID string) string {
	return hash.Values(customerID, refType, refID)
}

func (o *IssueType) setDefaults(frommap bool) {

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
func (o *IssueType) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IssueType) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IssueType) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IssueType) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IssueType) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IssueType) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IssueType) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IssueType) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IssueType) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IssueTypeModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IssueType) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IssueType) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of IssueType
func (o *IssueType) Clone() datamodel.Model {
	c := new(IssueType)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IssueType) Anon() datamodel.Model {
	c := new(IssueType)
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
func (o *IssueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IssueType) UnmarshalJSON(data []byte) error {
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
func (o *IssueType) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IssueType) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IssueType objects are equal
func (o *IssueType) IsEqual(other *IssueType) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IssueType) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toIssueTypeObject(o.CustomerID, false),
		"description":             toIssueTypeObject(o.Description, true),
		"icon_url":                toIssueTypeObject(o.IconURL, true),
		"id":                      toIssueTypeObject(o.ID, false),
		"integration_instance_id": toIssueTypeObject(o.IntegrationInstanceID, true),

		"mapped_type": o.MappedType.String(),
		"name":        toIssueTypeObject(o.Name, false),
		"ref_id":      toIssueTypeObject(o.RefID, false),
		"ref_type":    toIssueTypeObject(o.RefType, false),
		"hashcode":    toIssueTypeObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IssueType) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["mapped_type"].(IssueTypeMappedType); ok {
		o.MappedType = val
	} else {
		if em, ok := kv["mapped_type"].(map[string]interface{}); ok {

			ev := em["work.mapped_type"].(string)
			switch ev {
			case "unknown", "UNKNOWN":
				o.MappedType = 0
			case "feature", "FEATURE":
				o.MappedType = 1
			case "bug", "BUG":
				o.MappedType = 2
			case "enhancement", "ENHANCEMENT":
				o.MappedType = 3
			case "epic", "EPIC":
				o.MappedType = 4
			case "story", "STORY":
				o.MappedType = 5
			case "task", "TASK":
				o.MappedType = 6
			case "subtask", "SUBTASK":
				o.MappedType = 7
			}
		}
		if em, ok := kv["mapped_type"].(string); ok {
			switch em {
			case "unknown", "UNKNOWN":
				o.MappedType = 0
			case "feature", "FEATURE":
				o.MappedType = 1
			case "bug", "BUG":
				o.MappedType = 2
			case "enhancement", "ENHANCEMENT":
				o.MappedType = 3
			case "epic", "EPIC":
				o.MappedType = 4
			case "story", "STORY":
				o.MappedType = 5
			case "task", "TASK":
				o.MappedType = 6
			case "subtask", "SUBTASK":
				o.MappedType = 7
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
func (o *IssueType) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.IconURL)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.MappedType)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
