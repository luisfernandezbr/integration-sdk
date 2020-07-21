// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/graphql"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// TeamTable is the default table name
	TeamTable datamodel.ModelNameType = "customer_team"

	// TeamModelName is the model name
	TeamModelName datamodel.ModelNameType = "customer.Team"
)

const (
	// TeamModelActiveColumn is the column json value active
	TeamModelActiveColumn = "active"
	// TeamModelChildrenIdsColumn is the column json value children_ids
	TeamModelChildrenIdsColumn = "children_ids"
	// TeamModelCreatedAtColumn is the column json value created_ts
	TeamModelCreatedAtColumn = "created_ts"
	// TeamModelCustomerIDColumn is the column json value customer_id
	TeamModelCustomerIDColumn = "customer_id"
	// TeamModelDeletedColumn is the column json value deleted
	TeamModelDeletedColumn = "deleted"
	// TeamModelDeletedDateColumn is the column json value deleted_date
	TeamModelDeletedDateColumn = "deleted_date"
	// TeamModelDeletedDateEpochColumn is the column json value epoch
	TeamModelDeletedDateEpochColumn = "epoch"
	// TeamModelDeletedDateOffsetColumn is the column json value offset
	TeamModelDeletedDateOffsetColumn = "offset"
	// TeamModelDeletedDateRfc3339Column is the column json value rfc3339
	TeamModelDeletedDateRfc3339Column = "rfc3339"
	// TeamModelDescriptionColumn is the column json value description
	TeamModelDescriptionColumn = "description"
	// TeamModelIDColumn is the column json value id
	TeamModelIDColumn = "id"
	// TeamModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	TeamModelIntegrationInstanceIDColumn = "integration_instance_id"
	// TeamModelLeafColumn is the column json value leaf
	TeamModelLeafColumn = "leaf"
	// TeamModelNameColumn is the column json value name
	TeamModelNameColumn = "name"
	// TeamModelParentIdsColumn is the column json value parent_ids
	TeamModelParentIdsColumn = "parent_ids"
	// TeamModelProjectIdsColumn is the column json value project_ids
	TeamModelProjectIdsColumn = "project_ids"
	// TeamModelRefIDColumn is the column json value ref_id
	TeamModelRefIDColumn = "ref_id"
	// TeamModelRefTypeColumn is the column json value ref_type
	TeamModelRefTypeColumn = "ref_type"
	// TeamModelRepoIdsColumn is the column json value repo_ids
	TeamModelRepoIdsColumn = "repo_ids"
	// TeamModelUpdatedAtColumn is the column json value updated_ts
	TeamModelUpdatedAtColumn = "updated_ts"
)

// TeamDeletedDate represents the object structure for deleted_date
type TeamDeletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toTeamDeletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *TeamDeletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *TeamDeletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toTeamDeletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toTeamDeletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toTeamDeletedDateObject(o.Rfc3339, false),
	}
}

func (o *TeamDeletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *TeamDeletedDate) FromMap(kv map[string]interface{}) {

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

// Team a team is a grouping of one or more users
type Team struct {
	// Active whether the team is tracked in pinpoint
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids" codec:"children_ids" bson:"children_ids" yaml:"children_ids" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted delete flag for a team. true === deleted
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// DeletedDate when the user profile was soft deleted
	DeletedDate TeamDeletedDate `json:"deleted_date" codec:"deleted_date" bson:"deleted_date" yaml:"deleted_date" faker:"-"`
	// Description the description of the team
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Leaf True when team has no children_ids
	Leaf bool `json:"leaf" codec:"leaf" bson:"leaf" yaml:"leaf" faker:"-"`
	// Name the name of the team
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"team"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids" codec:"parent_ids" bson:"parent_ids" yaml:"parent_ids" faker:"-"`
	// ProjectIds project ids for the team
	ProjectIds []string `json:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoIds repo ids for the team
	RepoIds []string `json:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Team)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Team)(nil)

func toTeamObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Team:
		return v.ToMap()

	case TeamDeletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Team
func (o *Team) String() string {
	return fmt.Sprintf("customer.Team<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Team) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Team) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Team) GetTableName() string {
	return TeamTable.String()
}

// GetModelName returns the name of the model
func (o *Team) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

// NewTeamID provides a template for generating an ID field for Team
func NewTeamID(Name string, customerID string) string {
	return hash.Values(Name, customerID, datetime.EpochNow(), randomString(64))
}

func (o *Team) setDefaults(frommap bool) {
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
	}
	if o.RepoIds == nil {
		o.RepoIds = make([]string, 0)
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.Name, o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Team) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Team) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Team) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Team) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Team) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Team) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Team) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Team) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Team) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = TeamModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Team) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Team) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Team
func (o *Team) Clone() datamodel.Model {
	c := new(Team)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Team) Anon() datamodel.Model {
	c := new(Team)
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
func (o *Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Team) UnmarshalJSON(data []byte) error {
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
func (o *Team) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Team) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Team objects are equal
func (o *Team) IsEqual(other *Team) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Team) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toTeamObject(o.Active, false),
		"children_ids":            toTeamObject(o.ChildrenIds, false),
		"created_ts":              toTeamObject(o.CreatedAt, false),
		"customer_id":             toTeamObject(o.CustomerID, false),
		"deleted":                 toTeamObject(o.Deleted, false),
		"deleted_date":            toTeamObject(o.DeletedDate, false),
		"description":             toTeamObject(o.Description, false),
		"id":                      toTeamObject(o.ID, false),
		"integration_instance_id": toTeamObject(o.IntegrationInstanceID, true),
		"leaf":                    toTeamObject(o.Leaf, false),
		"name":                    toTeamObject(o.Name, false),
		"parent_ids":              toTeamObject(o.ParentIds, false),
		"project_ids":             toTeamObject(o.ProjectIds, false),
		"ref_id":                  toTeamObject(o.RefID, false),
		"ref_type":                toTeamObject(o.RefType, false),
		"repo_ids":                toTeamObject(o.RepoIds, false),
		"updated_ts":              toTeamObject(o.UpdatedAt, false),
		"hashcode":                toTeamObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Team) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["children_ids"]; ok {
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for children_ids field")
				}
			}
			o.ChildrenIds = na
		}
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
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
	if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = false
			} else {
				o.Deleted = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(TeamDeletedDate); ok {
			// struct
			o.DeletedDate = sv
		} else if sp, ok := val.(*TeamDeletedDate); ok {
			// struct pointer
			o.DeletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
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
	if val, ok := kv["leaf"].(bool); ok {
		o.Leaf = val
	} else {
		if val, ok := kv["leaf"]; ok {
			if val == nil {
				o.Leaf = false
			} else {
				o.Leaf = number.ToBoolAny(val)
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
	if val, ok := kv["parent_ids"]; ok {
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for parent_ids field")
				}
			}
			o.ParentIds = na
		}
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	if val, ok := kv["project_ids"]; ok {
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for project_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for project_ids field")
				}
			}
			o.ProjectIds = na
		}
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
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
	if val, ok := kv["repo_ids"]; ok {
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
								panic("unsupported type for repo_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for repo_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for repo_ids field")
				}
			}
			o.RepoIds = na
		}
	}
	if o.RepoIds == nil {
		o.RepoIds = make([]string, 0)
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
func (o *Team) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.ChildrenIds)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.DeletedDate)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Leaf)
	args = append(args, o.Name)
	args = append(args, o.ParentIds)
	args = append(args, o.ProjectIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoIds)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getTeamQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tchildren_ids\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tdeleted\n")
	// object with fields
	sb.WriteString("\t\t\tdeleted_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tdescription\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tleaf\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tparent_ids\n")
	// scalar
	sb.WriteString("\t\t\tproject_ids\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\trepo_ids\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// TeamPageInfo is a grapqhl PageInfo
type TeamPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// TeamConnection is a grapqhl connection
type TeamConnection struct {
	Edges      []*TeamEdge  `json:"edges,omitempty"`
	PageInfo   TeamPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64       `json:"totalCount,omitempty"`
}

// TeamEdge is a grapqhl edge
type TeamEdge struct {
	Node *Team `json:"node,omitempty"`
}

// QueryManyTeamNode is a grapqhl query many node
type QueryManyTeamNode struct {
	Object *TeamConnection `json:"Teams,omitempty"`
}

// QueryManyTeamData is a grapqhl query many data node
type QueryManyTeamData struct {
	Data *QueryManyTeamNode `json:"customer,omitempty"`
}

// QueryOneTeamNode is a grapqhl query one node
type QueryOneTeamNode struct {
	Object *Team `json:"Team,omitempty"`
}

// QueryOneTeamData is a grapqhl query one data node
type QueryOneTeamData struct {
	Data *QueryOneTeamNode `json:"customer,omitempty"`
}

// TeamQuery is query struct
type TeamQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// TeamOrder is the order direction
type TeamOrder *string

var (
	// TeamOrderDesc is descending
	TeamOrderDesc TeamOrder = pstrings.Pointer("DESC")
	// TeamOrderAsc is ascending
	TeamOrderAsc TeamOrder = pstrings.Pointer("ASC")
)

// TeamQueryInput is query input struct
type TeamQueryInput struct {
	First   *int64     `json:"first,omitempty"`
	Last    *int64     `json:"last,omitempty"`
	Before  *string    `json:"before,omitempty"`
	After   *string    `json:"after,omitempty"`
	Query   *TeamQuery `json:"query,omitempty"`
	OrderBy *string    `json:"orderBy,omitempty"`
	Order   TeamOrder  `json:"order,omitempty"`
}

// NewTeamQuery is a convenience for building a *TeamQuery
func NewTeamQuery(params ...interface{}) *TeamQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &TeamQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &TeamQueryInput{
		Query: q,
	}
}

// FindTeam will query an Team by id
func FindTeam(client graphql.Client, id string) (*Team, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query TeamQuery($id: ID) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tTeam(_id: $id) {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneTeamData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindTeams will query for any Teams matching the query
func FindTeams(client graphql.Client, input *TeamQueryInput) (*TeamConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
		if input.OrderBy != nil {
			variables["orderBy"] = *input.OrderBy
		}
		if input.Order != nil {
			variables["order"] = *input.Order
		}
	}
	var sb strings.Builder
	sb.WriteString("query TeamQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: CustomerTeamColumnEnum) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tTeams(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyTeamData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindTeamsPaginatedCallback is a callback function for handling each page
type FindTeamsPaginatedCallback func(conn *TeamConnection) (bool, error)

// FindTeamsPaginated will query for any Teams matching the query and return each page callback
func FindTeamsPaginated(client graphql.Client, query *TeamQuery, pageSize int64, callback FindTeamsPaginatedCallback) error {
	input := &TeamQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindTeams(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// HydratedQuery returns all fields, and one level deep of relations.
const TeamHydratedQuery = `query TeamQuery($id: ID) {
	customer {
		Team(_id: $id) {
			active
			children_ids
			children {
				edges {
					node {
			active
			children_ids
			created_ts
			customer_id
			deleted
			deleted_date {
			epoch
			offset
			rfc3339
			}
			description
			_id
			integration_instance_id
			leaf
			name
			parent_ids
			project_ids
			ref_id
			ref_type
			repo_ids
			updated_ts
					}
				}
			}
			created_ts
			customer_id
			customer {
			_id
			customer_id
			created_date
			updated_date
			name
			active
			apikey_id
			work_ids
			codequality_ids
			sourcecode_ids
			auth {
			provider
			}
			customer_information_id
			crm_id
			}
			deleted
			deleted_date {
			epoch
			offset
			rfc3339
			}
			description
			_id
			integration_instance_id
			integration_instance {
			active
			config
			created_by_profile_id
			created_by_user_id
			created_date {
			epoch
			offset
			rfc3339
			}
			created_ts
			customer_id
			deleted_by_profile_id
			deleted_by_user_id
			deleted_date {
			epoch
			offset
			rfc3339
			}
			error_message
			errored
			export_acknowledged
			_id
			integration_id
			integration_instance_id
			interval
			job_id
			last_export_completed_date {
			epoch
			offset
			rfc3339
			}
			last_export_requested_date {
			epoch
			offset
			rfc3339
			}
			last_historical_completed_date {
			epoch
			offset
			rfc3339
			}
			last_historical_requested_date {
			epoch
			offset
			rfc3339
			}
			last_processing_date {
			epoch
			offset
			rfc3339
			}
			location
			name
			paused
			processed
			ref_id
			ref_type
			requires_historical
			state
			throttled
			throttled_until {
			epoch
			offset
			rfc3339
			}
			updated_ts
			webhooks {
			enabled
			error_message
			errored
			ref_id
			url
			}
			}
			leaf
			name
			parent_ids
			parents {
				edges {
					node {
			active
			children_ids
			created_ts
			customer_id
			deleted
			deleted_date {
			epoch
			offset
			rfc3339
			}
			description
			_id
			integration_instance_id
			leaf
			name
			parent_ids
			project_ids
			ref_id
			ref_type
			repo_ids
			updated_ts
					}
				}
			}
			project_ids
			ref_id
			ref_type
			repo_ids
			updated_ts
		}
	}
}
`

// UpdateTeamNode is a grapqhl update node
type UpdateTeamNode struct {
	Object *Team `json:"updateTeam,omitempty"`
}

// UpdateTeamData is a grapqhl update data node
type UpdateTeamData struct {
	Data *UpdateTeamNode `json:"customer,omitempty"`
}

// ExecTeamUpdateMutation returns a graphql update mutation result for Team
func ExecTeamUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*Team, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation TeamUpdateMutation($id: String, $input: UpdateCustomerTeamInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateTeam(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getTeamQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateTeamData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecTeamSilentUpdateMutation returns a graphql update mutation result for Team
func ExecTeamSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation TeamUpdateMutation($id: String, $input: UpdateCustomerTeamInput, $upsert: Boolean) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tupdateTeam(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateTeamData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecTeamDeleteMutation executes a graphql delete mutation for Team
func ExecTeamDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation TeamDeleteMutation($id: String!) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tdeleteTeam(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateTeam(client graphql.Client, model Team) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CreateTeam($input: CreateCustomerTeamInput!) {\n")
	sb.WriteString("\tcustomer {\n")
	sb.WriteString("\t\tcreateTeam(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateTeamData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// TeamPartial is a partial struct for upsert mutations for Team
type TeamPartial struct {
	// Active whether the team is tracked in pinpoint
	Active *bool `json:"active,omitempty"`
	// ChildrenIds the children_ids for this team
	ChildrenIds []string `json:"children_ids,omitempty"`
	// Deleted delete flag for a team. true === deleted
	Deleted *bool `json:"deleted,omitempty"`
	// DeletedDate when the user profile was soft deleted
	DeletedDate *TeamDeletedDate `json:"deleted_date,omitempty"`
	// Description the description of the team
	Description *string `json:"description,omitempty"`
	// Leaf True when team has no children_ids
	Leaf *bool `json:"leaf,omitempty"`
	// Name the name of the team
	Name *string `json:"name,omitempty"`
	// ParentIds the parent_ids for this team
	ParentIds []string `json:"parent_ids,omitempty"`
}

var _ datamodel.PartialModel = (*TeamPartial)(nil)

// GetModelName returns the name of the model
func (o *TeamPartial) GetModelName() datamodel.ModelNameType {
	return TeamModelName
}

// ToMap returns the object as a map
func (o *TeamPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":       toTeamObject(o.Active, true),
		"children_ids": toTeamObject(o.ChildrenIds, true),
		"deleted":      toTeamObject(o.Deleted, true),
		"deleted_date": toTeamObject(o.DeletedDate, true),
		"description":  toTeamObject(o.Description, true),
		"leaf":         toTeamObject(o.Leaf, true),
		"name":         toTeamObject(o.Name, true),
		"parent_ids":   toTeamObject(o.ParentIds, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "children_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "deleted_date" {
				if dt, ok := v.(*TeamDeletedDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}

			if k == "parent_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *TeamPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *TeamPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *TeamPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *TeamPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *TeamPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["children_ids"]; ok {
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for children_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for children_ids field")
				}
			}
			o.ChildrenIds = na
		}
	}
	if o.ChildrenIds == nil {
		o.ChildrenIds = make([]string, 0)
	}
	if val, ok := kv["deleted"].(*bool); ok {
		o.Deleted = val
	} else if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = &val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Deleted = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.DeletedDate == nil {
		o.DeletedDate = &TeamDeletedDate{}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(TeamDeletedDate); ok {
			// struct
			o.DeletedDate = &sv
		} else if sp, ok := val.(*TeamDeletedDate); ok {
			// struct pointer
			o.DeletedDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
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
	if val, ok := kv["leaf"].(*bool); ok {
		o.Leaf = val
	} else if val, ok := kv["leaf"].(bool); ok {
		o.Leaf = &val
	} else {
		if val, ok := kv["leaf"]; ok {
			if val == nil {
				o.Leaf = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Leaf = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["name"].(*string); ok {
		o.Name = val
	} else if val, ok := kv["name"].(string); ok {
		o.Name = &val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Name = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["parent_ids"]; ok {
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for parent_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for parent_ids field")
				}
			}
			o.ParentIds = na
		}
	}
	if o.ParentIds == nil {
		o.ParentIds = make([]string, 0)
	}
	o.setDefaults(false)
}
