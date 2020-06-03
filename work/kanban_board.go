// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// KanbanBoardTopic is the default topic name
	KanbanBoardTopic datamodel.TopicNameType = "work_KanbanBoard"

	// KanbanBoardTable is the default table name
	KanbanBoardTable datamodel.ModelNameType = "work_kanbanboard"

	// KanbanBoardModelName is the model name
	KanbanBoardModelName datamodel.ModelNameType = "work.KanbanBoard"
)

const (
	// KanbanBoardModelColumnsColumn is the column json value columns
	KanbanBoardModelColumnsColumn = "columns"
	// KanbanBoardModelColumnsNameColumn is the column json value name
	KanbanBoardModelColumnsNameColumn = "name"
	// KanbanBoardModelColumnsStatusIdsColumn is the column json value status_ids
	KanbanBoardModelColumnsStatusIdsColumn = "status_ids"
	// KanbanBoardModelCustomerIDColumn is the column json value customer_id
	KanbanBoardModelCustomerIDColumn = "customer_id"
	// KanbanBoardModelIDColumn is the column json value id
	KanbanBoardModelIDColumn = "id"
	// KanbanBoardModelNameColumn is the column json value name
	KanbanBoardModelNameColumn = "name"
	// KanbanBoardModelProjectIdsColumn is the column json value project_ids
	KanbanBoardModelProjectIdsColumn = "project_ids"
	// KanbanBoardModelRefIDColumn is the column json value ref_id
	KanbanBoardModelRefIDColumn = "ref_id"
	// KanbanBoardModelRefTypeColumn is the column json value ref_type
	KanbanBoardModelRefTypeColumn = "ref_type"
	// KanbanBoardModelUpdatedAtColumn is the column json value updated_ts
	KanbanBoardModelUpdatedAtColumn = "updated_ts"
)

// KanbanBoardColumns represents the object structure for columns
type KanbanBoardColumns struct {
	// Name name of the column
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// StatusIds ids of the statuses for this column
	StatusIds []string `json:"status_ids" codec:"status_ids" bson:"status_ids" yaml:"status_ids" faker:"-"`
}

func toKanbanBoardColumnsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *KanbanBoardColumns:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *KanbanBoardColumns) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Name name of the column
		"name": toKanbanBoardColumnsObject(o.Name, false),
		// StatusIds ids of the statuses for this column
		"status_ids": toKanbanBoardColumnsObject(o.StatusIds, false),
	}
}

func (o *KanbanBoardColumns) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *KanbanBoardColumns) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	if val, ok := kv["status_ids"]; ok {
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
								panic("unsupported type for status_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for status_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for status_ids field")
				}
			}
			o.StatusIds = na
		}
	}
	if o.StatusIds == nil {
		o.StatusIds = make([]string, 0)
	}
	o.setDefaults(false)
}

// KanbanBoard kanban board details
type KanbanBoard struct {
	// Columns the columns for this board
	Columns []KanbanBoardColumns `json:"columns" codec:"columns" bson:"columns" yaml:"columns" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the board
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// ProjectIds ids of the projects used in this board
	ProjectIds []string `json:"project_ids" codec:"project_ids" bson:"project_ids" yaml:"project_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*KanbanBoard)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*KanbanBoard)(nil)

func toKanbanBoardObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *KanbanBoard:
		return v.ToMap()

	case []KanbanBoardColumns:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of KanbanBoard
func (o *KanbanBoard) String() string {
	return fmt.Sprintf("work.KanbanBoard<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *KanbanBoard) GetTopicName() datamodel.TopicNameType {
	return KanbanBoardTopic
}

// GetStreamName returns the name of the stream
func (o *KanbanBoard) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *KanbanBoard) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *KanbanBoard) GetModelName() datamodel.ModelNameType {
	return KanbanBoardModelName
}

// NewKanbanBoardID provides a template for generating an ID field for KanbanBoard
func NewKanbanBoardID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *KanbanBoard) setDefaults(frommap bool) {
	if o.Columns == nil {
		o.Columns = make([]KanbanBoardColumns, 0)
	}
	if o.ProjectIds == nil {
		o.ProjectIds = make([]string, 0)
	}

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *KanbanBoard) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *KanbanBoard) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *KanbanBoard) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for KanbanBoard")
}

// GetRefID returns the RefID for the object
func (o *KanbanBoard) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *KanbanBoard) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *KanbanBoard) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *KanbanBoard) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *KanbanBoard) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *KanbanBoard) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = KanbanBoardModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *KanbanBoard) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     128,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetCustomerID will return the customer_id
func (o *KanbanBoard) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of KanbanBoard
func (o *KanbanBoard) Clone() datamodel.Model {
	c := new(KanbanBoard)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *KanbanBoard) Anon() datamodel.Model {
	c := new(KanbanBoard)
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
func (o *KanbanBoard) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *KanbanBoard) UnmarshalJSON(data []byte) error {
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
func (o *KanbanBoard) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *KanbanBoard) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two KanbanBoard objects are equal
func (o *KanbanBoard) IsEqual(other *KanbanBoard) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *KanbanBoard) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"columns":     toKanbanBoardObject(o.Columns, false),
		"customer_id": toKanbanBoardObject(o.CustomerID, false),
		"id":          toKanbanBoardObject(o.ID, false),
		"name":        toKanbanBoardObject(o.Name, false),
		"project_ids": toKanbanBoardObject(o.ProjectIds, false),
		"ref_id":      toKanbanBoardObject(o.RefID, false),
		"ref_type":    toKanbanBoardObject(o.RefType, false),
		"updated_ts":  toKanbanBoardObject(o.UpdatedAt, false),
		"hashcode":    toKanbanBoardObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *KanbanBoard) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if o == nil {

		o.Columns = make([]KanbanBoardColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]KanbanBoardColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*KanbanBoardColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(KanbanBoardColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm KanbanBoardColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av KanbanBoardColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(KanbanBoardColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm KanbanBoardColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := KanbanBoardColumns{}
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm KanbanBoardColumns
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Columns = append(o.Columns, fm)
						}
					}
				}
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
func (o *KanbanBoard) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Columns)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.ProjectIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
