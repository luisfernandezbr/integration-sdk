// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// ExportTriggerModelName is the model name
	ExportTriggerModelName datamodel.ModelNameType = "agent.ExportTrigger"
)

// ExportTrigger used to trigger an agent.ExportRequest
type ExportTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical bool `json:"reprocess_historical" codec:"reprocess_historical" bson:"reprocess_historical" yaml:"reprocess_historical" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID This UUID of the agent to trigger
	UUID *string `json:"uuid,omitempty" codec:"uuid,omitempty" bson:"uuid" yaml:"uuid,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportTrigger)(nil)

func toExportTriggerObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportTrigger:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportTrigger
func (o *ExportTrigger) String() string {
	return fmt.Sprintf("agent.ExportTrigger<%s>", o.ID)
}

// GetModelName returns the name of the model
func (o *ExportTrigger) GetModelName() datamodel.ModelNameType {
	return ExportTriggerModelName
}

// NewExportTriggerID provides a template for generating an ID field for ExportTrigger
func NewExportTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("ExportTrigger", customerID, refType, refID)
}

func (o *ExportTrigger) setDefaults(frommap bool) {
	if o.UUID == nil {
		o.UUID = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportTrigger) GetID() string {
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *ExportTrigger) GetRefID() string {
	return o.RefID
}

// GetCustomerID will return the customer_id
func (o *ExportTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ExportTrigger
func (o *ExportTrigger) Clone() datamodel.Model {
	c := new(ExportTrigger)
	c.FromMap(o.ToMap())
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportTrigger) UnmarshalJSON(data []byte) error {
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
func (o *ExportTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ExportTrigger objects are equal
func (o *ExportTrigger) IsEqual(other *ExportTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportTrigger) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":          toExportTriggerObject(o.CustomerID, false),
		"id":                   toExportTriggerObject(o.ID, false),
		"ref_id":               toExportTriggerObject(o.RefID, false),
		"ref_type":             toExportTriggerObject(o.RefType, false),
		"reprocess_historical": toExportTriggerObject(o.ReprocessHistorical, false),
		"updated_ts":           toExportTriggerObject(o.UpdatedAt, false),
		"uuid":                 toExportTriggerObject(o.UUID, true),
		"hashcode":             toExportTriggerObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportTrigger) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = number.ToBoolAny(nil)
			} else {
				o.ReprocessHistorical = number.ToBoolAny(val)
			}
		}
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

	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ExportTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
