// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/graphql"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (

	// CostCenterTable is the default table name
	CostCenterTable datamodel.ModelNameType = "customer_costcenter"

	// CostCenterModelName is the model name
	CostCenterModelName datamodel.ModelNameType = "customer.CostCenter"
)

// CostCenter a cost center represents information about users and their cost
type CostCenter struct {
	// Active whether the cost center is tracked in pinpoint
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Cost the cost value of the cost center
	Cost float64 `json:"cost" codec:"cost" bson:"cost" yaml:"cost" faker:"salary"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description for the cost center
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the cost center
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"costcenter"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CostCenter)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CostCenter)(nil)

func toCostCenterObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CostCenter:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CostCenter
func (o *CostCenter) String() string {
	return fmt.Sprintf("customer.CostCenter<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CostCenter) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CostCenter) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CostCenter) GetTableName() string {
	return CostCenterTable.String()
}

// GetModelName returns the name of the model
func (o *CostCenter) GetModelName() datamodel.ModelNameType {
	return CostCenterModelName
}

// NewCostCenterID provides a template for generating an ID field for CostCenter
func NewCostCenterID(customerID string) string {
	return hash.Values(customerID, randomString(64))
}

func (o *CostCenter) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CostCenter) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CostCenter) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CostCenter) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CostCenter) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CostCenter) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CostCenter) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CostCenter) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CostCenter) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *CostCenter) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *CostCenter) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CostCenter
func (o *CostCenter) Clone() datamodel.Model {
	c := new(CostCenter)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CostCenter) Anon() datamodel.Model {
	c := new(CostCenter)
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
func (o *CostCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CostCenter) UnmarshalJSON(data []byte) error {
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
func (o *CostCenter) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CostCenter objects are equal
func (o *CostCenter) IsEqual(other *CostCenter) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CostCenter) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":      toCostCenterObject(o.Active, false),
		"cost":        toCostCenterObject(o.Cost, false),
		"created_ts":  toCostCenterObject(o.CreatedAt, false),
		"customer_id": toCostCenterObject(o.CustomerID, false),
		"description": toCostCenterObject(o.Description, false),
		"id":          toCostCenterObject(o.ID, false),
		"name":        toCostCenterObject(o.Name, false),
		"ref_id":      toCostCenterObject(o.RefID, false),
		"ref_type":    toCostCenterObject(o.RefType, false),
		"updated_ts":  toCostCenterObject(o.UpdatedAt, false),
		"hashcode":    toCostCenterObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CostCenter) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["cost"].(float64); ok {
		o.Cost = val
	} else {
		if val, ok := kv["cost"]; ok {
			if val == nil {
				o.Cost = 0.0
			} else {
				o.Cost = number.ToFloat64Any(val)
			}
		}
	}

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.CreatedAt = number.ToInt64Any(val)
			}
		}
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

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *CostCenter) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Cost)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CostCenter) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// ExecCostCenterUpdateMutation returns a graphql update mutation for CostCenter
func ExecCostCenterUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*CostCenter, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CostCenterUpdateMutation($id: String, $input: UpdateCustomerCostCenterInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateCostCenter(_id: $id, input: $input, upsert: $upsert) {\n")

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tcost\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tdescription\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	kv := make(map[string]interface{})
	if err := client.Mutate(sb.String(), variables, &kv); err != nil {
		return nil, err
	}
	if sdata, ok := kv["customer"].(map[string]interface{}); ok {
		if mdata, ok := sdata["updateCostCenter"].(map[string]interface{}); ok {
			var object CostCenter
			object.FromMap(mdata)
			return &object, nil
		}
		return nil, fmt.Errorf("missing expected updateCostCenter key from customer")
	}
	return nil, fmt.Errorf("missing expected customer key from data")
}
