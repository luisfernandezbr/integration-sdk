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
	// KanbanTopic is the default topic name
	KanbanTopic datamodel.TopicNameType = "work_Kanban"

	// KanbanTable is the default table name
	KanbanTable datamodel.ModelNameType = "work_kanban"

	// KanbanModelName is the model name
	KanbanModelName datamodel.ModelNameType = "work.Kanban"
)

const (
	// KanbanModelActiveColumn is the column json value active
	KanbanModelActiveColumn = "active"
	// KanbanModelBoardIDColumn is the column json value board_id
	KanbanModelBoardIDColumn = "board_id"
	// KanbanModelColumnsColumn is the column json value columns
	KanbanModelColumnsColumn = "columns"
	// KanbanModelColumnsIssueIdsColumn is the column json value issue_ids
	KanbanModelColumnsIssueIdsColumn = "issue_ids"
	// KanbanModelColumnsNameColumn is the column json value name
	KanbanModelColumnsNameColumn = "name"
	// KanbanModelCustomerIDColumn is the column json value customer_id
	KanbanModelCustomerIDColumn = "customer_id"
	// KanbanModelDeletedColumn is the column json value deleted
	KanbanModelDeletedColumn = "deleted"
	// KanbanModelIDColumn is the column json value id
	KanbanModelIDColumn = "id"
	// KanbanModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	KanbanModelIntegrationInstanceIDColumn = "integration_instance_id"
	// KanbanModelIssueIdsColumn is the column json value issue_ids
	KanbanModelIssueIdsColumn = "issue_ids"
	// KanbanModelNameColumn is the column json value name
	KanbanModelNameColumn = "name"
	// KanbanModelProjectIdsColumn is the column json value project_ids
	KanbanModelProjectIdsColumn = "project_ids"
	// KanbanModelRefIDColumn is the column json value ref_id
	KanbanModelRefIDColumn = "ref_id"
	// KanbanModelRefTypeColumn is the column json value ref_type
	KanbanModelRefTypeColumn = "ref_type"
	// KanbanModelUpdatedDateColumn is the column json value updated_date
	KanbanModelUpdatedDateColumn = "updated_date"
	// KanbanModelUpdatedDateEpochColumn is the column json value epoch
	KanbanModelUpdatedDateEpochColumn = "epoch"
	// KanbanModelUpdatedDateOffsetColumn is the column json value offset
	KanbanModelUpdatedDateOffsetColumn = "offset"
	// KanbanModelUpdatedDateRfc3339Column is the column json value rfc3339
	KanbanModelUpdatedDateRfc3339Column = "rfc3339"
	// KanbanModelUpdatedAtColumn is the column json value updated_ts
	KanbanModelUpdatedAtColumn = "updated_ts"
	// KanbanModelURLColumn is the column json value url
	KanbanModelURLColumn = "url"
)

// KanbanColumns represents the object structure for columns
type KanbanColumns struct {
	// IssueIds the issue ids for the column
	IssueIds []string `json:"issue_ids" codec:"issue_ids" bson:"issue_ids" yaml:"issue_ids" faker:"-"`
	// Name the name of the column
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
}

func toKanbanColumnsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *KanbanColumns:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *KanbanColumns) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// IssueIds the issue ids for the column
		"issue_ids": toKanbanColumnsObject(o.IssueIds, false),
		// Name the name of the column
		"name": toKanbanColumnsObject(o.Name, false),
	}
}

func (o *KanbanColumns) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *KanbanColumns) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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
	o.setDefaults(false)
}

// KanbanUpdatedDate represents the object structure for updated_date
type KanbanUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toKanbanUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *KanbanUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *KanbanUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toKanbanUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toKanbanUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toKanbanUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *KanbanUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *KanbanUpdatedDate) FromMap(kv map[string]interface{}) {

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

// Kanban a kanban specific board
type Kanban struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// BoardID the board that this kanban board is related to
	BoardID string `json:"board_id" codec:"board_id" bson:"board_id" yaml:"board_id" faker:"-"`
	// Columns the columns for the board in the order that they are displayed and the issue ids for each
	Columns []KanbanColumns `json:"columns" codec:"columns" bson:"columns" yaml:"columns" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// IssueIds the issue ids for the board regardless of their column location
	IssueIds []string `json:"issue_ids" codec:"issue_ids" bson:"issue_ids" yaml:"issue_ids" faker:"-"`
	// Name the name of the board
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// ProjectIds ids of the projects used in this board
	ProjectIds []string `json:"project_ids" codec:"project_ids" bson:"project_ids" yaml:"project_ids" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedDate the timestamp that the kanban was updated
	UpdatedDate KanbanUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the board
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Kanban)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Kanban)(nil)

func toKanbanObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Kanban:
		return v.ToMap()

	case []KanbanColumns:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case KanbanUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Kanban
func (o *Kanban) String() string {
	return fmt.Sprintf("work.Kanban<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Kanban) GetTopicName() datamodel.TopicNameType {
	return KanbanTopic
}

// GetStreamName returns the name of the stream
func (o *Kanban) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Kanban) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Kanban) GetModelName() datamodel.ModelNameType {
	return KanbanModelName
}

// NewKanbanID provides a template for generating an ID field for Kanban
func NewKanbanID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *Kanban) setDefaults(frommap bool) {
	if o.Columns == nil {
		o.Columns = make([]KanbanColumns, 0)
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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
func (o *Kanban) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Kanban) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Kanban) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Kanban")
}

// GetRefID returns the RefID for the object
func (o *Kanban) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Kanban) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Kanban) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Kanban) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Kanban) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Kanban) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = KanbanModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Kanban) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Kanban) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Kanban) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Kanban) GetRefType() string {

	return o.RefType

}

// SetRefType will return the ref_type
func (o *Kanban) SetRefType(t string) {

	o.RefType = id

}

// Clone returns an exact copy of Kanban
func (o *Kanban) Clone() datamodel.Model {
	c := new(Kanban)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Kanban) Anon() datamodel.Model {
	c := new(Kanban)
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
func (o *Kanban) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Kanban) UnmarshalJSON(data []byte) error {
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
func (o *Kanban) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Kanban) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Kanban objects are equal
func (o *Kanban) IsEqual(other *Kanban) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Kanban) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toKanbanObject(o.Active, false),
		"board_id":                toKanbanObject(o.BoardID, false),
		"columns":                 toKanbanObject(o.Columns, false),
		"customer_id":             toKanbanObject(o.CustomerID, false),
		"deleted":                 toKanbanObject(o.Deleted, false),
		"id":                      toKanbanObject(o.ID, false),
		"integration_instance_id": toKanbanObject(o.IntegrationInstanceID, true),
		"issue_ids":               toKanbanObject(o.IssueIds, false),
		"name":                    toKanbanObject(o.Name, false),
		"project_ids":             toKanbanObject(o.ProjectIds, false),
		"ref_id":                  toKanbanObject(o.RefID, false),
		"ref_type":                toKanbanObject(o.RefType, false),
		"updated_date":            toKanbanObject(o.UpdatedDate, false),
		"updated_ts":              toKanbanObject(o.UpdatedAt, false),
		"url":                     toKanbanObject(o.URL, true),
		"hashcode":                toKanbanObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Kanban) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["board_id"].(string); ok {
		o.BoardID = val
	} else {
		if val, ok := kv["board_id"]; ok {
			if val == nil {
				o.BoardID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.BoardID = fmt.Sprintf("%v", val)
			}
		}
	}

	if o == nil {

		o.Columns = make([]KanbanColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]KanbanColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*KanbanColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(KanbanColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm KanbanColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av KanbanColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(KanbanColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm KanbanColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := KanbanColumns{}
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
							var fm KanbanColumns
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
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(KanbanUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*KanbanUpdatedDate); ok {
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

// Hash will return a hashcode for the object
func (o *Kanban) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.BoardID)
	args = append(args, o.Columns)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.IssueIds)
	args = append(args, o.Name)
	args = append(args, o.ProjectIds)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedDate)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Kanban) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Kanban) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// KanbanPartial is a partial struct for upsert mutations for Kanban
type KanbanPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// BoardID the board that this kanban board is related to
	BoardID *string `json:"board_id,omitempty"`
	// Columns the columns for the board in the order that they are displayed and the issue ids for each
	Columns []KanbanColumns `json:"columns,omitempty"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted *bool `json:"deleted,omitempty"`
	// IssueIds the issue ids for the board regardless of their column location
	IssueIds []string `json:"issue_ids,omitempty"`
	// Name the name of the board
	Name *string `json:"name,omitempty"`
	// ProjectIds ids of the projects used in this board
	ProjectIds []string `json:"project_ids,omitempty"`
	// UpdatedDate the timestamp that the kanban was updated
	UpdatedDate *KanbanUpdatedDate `json:"updated_date,omitempty"`
	// URL the url to the board
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*KanbanPartial)(nil)

// GetModelName returns the name of the model
func (o *KanbanPartial) GetModelName() datamodel.ModelNameType {
	return KanbanModelName
}

// ToMap returns the object as a map
func (o *KanbanPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":       toKanbanObject(o.Active, true),
		"board_id":     toKanbanObject(o.BoardID, true),
		"columns":      toKanbanObject(o.Columns, true),
		"deleted":      toKanbanObject(o.Deleted, true),
		"issue_ids":    toKanbanObject(o.IssueIds, true),
		"name":         toKanbanObject(o.Name, true),
		"project_ids":  toKanbanObject(o.ProjectIds, true),
		"updated_date": toKanbanObject(o.UpdatedDate, true),
		"url":          toKanbanObject(o.URL, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "columns" {
				if arr, ok := v.([]KanbanColumns); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "issue_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "project_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "updated_date" {
				if dt, ok := v.(*KanbanUpdatedDate); ok {
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
func (o *KanbanPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *KanbanPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *KanbanPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *KanbanPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *KanbanPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["board_id"].(*string); ok {
		o.BoardID = val
	} else if val, ok := kv["board_id"].(string); ok {
		o.BoardID = &val
	} else {
		if val, ok := kv["board_id"]; ok {
			if val == nil {
				o.BoardID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.BoardID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o == nil {

		o.Columns = make([]KanbanColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]KanbanColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*KanbanColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(KanbanColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm KanbanColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av KanbanColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(KanbanColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm KanbanColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := KanbanColumns{}
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
							var fm KanbanColumns
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Columns = append(o.Columns, fm)
						}
					}
				}
			}
		}
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
	if val, ok := kv["issue_ids"]; ok {
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for issue_ids field")
				}
			}
			o.IssueIds = na
		}
	}
	if o.IssueIds == nil {
		o.IssueIds = make([]string, 0)
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

	if o.UpdatedDate == nil {
		o.UpdatedDate = &KanbanUpdatedDate{}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(KanbanUpdatedDate); ok {
			// struct
			o.UpdatedDate = &sv
		} else if sp, ok := val.(*KanbanUpdatedDate); ok {
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
	o.setDefaults(false)
}
