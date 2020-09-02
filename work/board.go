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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// BoardTopic is the default topic name
	BoardTopic datamodel.TopicNameType = "work_Board"

	// BoardTable is the default table name
	BoardTable datamodel.ModelNameType = "work_board"

	// BoardModelName is the model name
	BoardModelName datamodel.ModelNameType = "work.Board"
)

const (
	// BoardModelActiveColumn is the column json value active
	BoardModelActiveColumn = "active"
	// BoardModelBacklogIssueIdsColumn is the column json value backlog_issue_ids
	BoardModelBacklogIssueIdsColumn = "backlog_issue_ids"
	// BoardModelColumnsColumn is the column json value columns
	BoardModelColumnsColumn = "columns"
	// BoardModelColumnsNameColumn is the column json value name
	BoardModelColumnsNameColumn = "name"
	// BoardModelCustomerIDColumn is the column json value customer_id
	BoardModelCustomerIDColumn = "customer_id"
	// BoardModelDeletedColumn is the column json value deleted
	BoardModelDeletedColumn = "deleted"
	// BoardModelIDColumn is the column json value id
	BoardModelIDColumn = "id"
	// BoardModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	BoardModelIntegrationInstanceIDColumn = "integration_instance_id"
	// BoardModelNameColumn is the column json value name
	BoardModelNameColumn = "name"
	// BoardModelRefIDColumn is the column json value ref_id
	BoardModelRefIDColumn = "ref_id"
	// BoardModelRefTypeColumn is the column json value ref_type
	BoardModelRefTypeColumn = "ref_type"
	// BoardModelTypeColumn is the column json value type
	BoardModelTypeColumn = "type"
	// BoardModelUpdatedDateColumn is the column json value updated_date
	BoardModelUpdatedDateColumn = "updated_date"
	// BoardModelUpdatedDateEpochColumn is the column json value epoch
	BoardModelUpdatedDateEpochColumn = "epoch"
	// BoardModelUpdatedDateOffsetColumn is the column json value offset
	BoardModelUpdatedDateOffsetColumn = "offset"
	// BoardModelUpdatedDateRfc3339Column is the column json value rfc3339
	BoardModelUpdatedDateRfc3339Column = "rfc3339"
	// BoardModelUpdatedAtColumn is the column json value updated_ts
	BoardModelUpdatedAtColumn = "updated_ts"
)

// BoardColumns represents the object structure for columns
type BoardColumns struct {
	// Name the name of the column
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
}

func toBoardColumnsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BoardColumns:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *BoardColumns) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Name the name of the column
		"name": toBoardColumnsObject(o.Name, false),
	}
}

func (o *BoardColumns) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BoardColumns) FromMap(kv map[string]interface{}) {

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
	o.setDefaults(false)
}

// BoardType is the enumeration type for type
type BoardType int32

// toBoardTypePointer is the enumeration pointer type for type
func toBoardTypePointer(v int32) *BoardType {
	nv := BoardType(v)
	return &nv
}

// toBoardTypeEnum is the enumeration pointer wrapper for type
func toBoardTypeEnum(v *BoardType) string {
	if v == nil {
		return toBoardTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *BoardType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = BoardType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "SCRUM":
			*v = BoardType(0)
		case "KANBAN":
			*v = BoardType(1)
		case "OTHER":
			*v = BoardType(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *BoardType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "SCRUM":
		*v = 0
	case "KANBAN":
		*v = 1
	case "OTHER":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v BoardType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("SCRUM")
	case 1:
		return json.Marshal("KANBAN")
	case 2:
		return json.Marshal("OTHER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v BoardType) String() string {
	switch int32(v) {
	case 0:
		return "SCRUM"
	case 1:
		return "KANBAN"
	case 2:
		return "OTHER"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *BoardType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case BoardType:
		*v = val
	case int32:
		*v = BoardType(int32(val))
	case int:
		*v = BoardType(int32(val))
	case string:
		switch val {
		case "SCRUM":
			*v = BoardType(0)
		case "KANBAN":
			*v = BoardType(1)
		case "OTHER":
			*v = BoardType(2)
		}
	}
	return nil
}

const (
	// BoardTypeScrum is the enumeration value for scrum
	BoardTypeScrum BoardType = 0
	// BoardTypeKanban is the enumeration value for kanban
	BoardTypeKanban BoardType = 1
	// BoardTypeOther is the enumeration value for other
	BoardTypeOther BoardType = 2
)

// BoardUpdatedDate represents the object structure for updated_date
type BoardUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toBoardUpdatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *BoardUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *BoardUpdatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toBoardUpdatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toBoardUpdatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toBoardUpdatedDateObject(o.Rfc3339, false),
	}
}

func (o *BoardUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *BoardUpdatedDate) FromMap(kv map[string]interface{}) {

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

// Board details about an agile board
type Board struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// BacklogIssueIds the issue ids that are part of the board but not active on the board
	BacklogIssueIds []string `json:"backlog_issue_ids" codec:"backlog_issue_ids" bson:"backlog_issue_ids" yaml:"backlog_issue_ids" faker:"-"`
	// Columns the columns for the board in the order that they are displayed
	Columns []BoardColumns `json:"columns" codec:"columns" bson:"columns" yaml:"columns" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Name the name of the board
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Type the type of board
	Type BoardType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedDate the timestamp that the board was updated
	UpdatedDate BoardUpdatedDate `json:"updated_date" codec:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Board)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Board)(nil)

func toBoardObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Board:
		return v.ToMap()

	case []BoardColumns:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case BoardType:
		return v.String()

	case BoardUpdatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Board
func (o *Board) String() string {
	return fmt.Sprintf("work.Board<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Board) GetTopicName() datamodel.TopicNameType {
	return BoardTopic
}

// GetStreamName returns the name of the stream
func (o *Board) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Board) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Board) GetModelName() datamodel.ModelNameType {
	return BoardModelName
}

// NewBoardID provides a template for generating an ID field for Board
func NewBoardID(customerID string, refID string, refType string) string {
	return hash.Values(customerID, refID, refType)
}

func (o *Board) setDefaults(frommap bool) {
	if o.BacklogIssueIds == nil {
		o.BacklogIssueIds = make([]string, 0)
	}
	if o.Columns == nil {
		o.Columns = make([]BoardColumns, 0)
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
func (o *Board) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Board) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Board) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Board")
}

// GetRefID returns the RefID for the object
func (o *Board) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Board) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Board) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Board) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Board) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Board) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = BoardModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Board) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Board) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Board
func (o *Board) Clone() datamodel.Model {
	c := new(Board)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Board) Anon() datamodel.Model {
	c := new(Board)
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
func (o *Board) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Board) UnmarshalJSON(data []byte) error {
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
func (o *Board) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Board) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Board objects are equal
func (o *Board) IsEqual(other *Board) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Board) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toBoardObject(o.Active, false),
		"backlog_issue_ids":       toBoardObject(o.BacklogIssueIds, false),
		"columns":                 toBoardObject(o.Columns, false),
		"customer_id":             toBoardObject(o.CustomerID, false),
		"deleted":                 toBoardObject(o.Deleted, false),
		"id":                      toBoardObject(o.ID, false),
		"integration_instance_id": toBoardObject(o.IntegrationInstanceID, true),
		"name":                    toBoardObject(o.Name, false),
		"ref_id":                  toBoardObject(o.RefID, false),
		"ref_type":                toBoardObject(o.RefType, false),

		"type":         o.Type.String(),
		"updated_date": toBoardObject(o.UpdatedDate, false),
		"updated_ts":   toBoardObject(o.UpdatedAt, false),
		"hashcode":     toBoardObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Board) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["backlog_issue_ids"]; ok {
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
								panic("unsupported type for backlog_issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for backlog_issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for backlog_issue_ids field")
				}
			}
			o.BacklogIssueIds = na
		}
	}
	if o.BacklogIssueIds == nil {
		o.BacklogIssueIds = make([]string, 0)
	}

	if o == nil {

		o.Columns = make([]BoardColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]BoardColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*BoardColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(BoardColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm BoardColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av BoardColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(BoardColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm BoardColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := BoardColumns{}
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
							var fm BoardColumns
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
	if val, ok := kv["type"].(BoardType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {

			ev := em["work.type"].(string)
			switch ev {
			case "scrum", "SCRUM":
				o.Type = 0
			case "kanban", "KANBAN":
				o.Type = 1
			case "other", "OTHER":
				o.Type = 2
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "scrum", "SCRUM":
				o.Type = 0
			case "kanban", "KANBAN":
				o.Type = 1
			case "other", "OTHER":
				o.Type = 2
			}
		}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(BoardUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*BoardUpdatedDate); ok {
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Board) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.BacklogIssueIds)
	args = append(args, o.Columns)
	args = append(args, o.CustomerID)
	args = append(args, o.Deleted)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Type)
	args = append(args, o.UpdatedDate)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Board) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Board) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// BoardPartial is a partial struct for upsert mutations for Board
type BoardPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// BacklogIssueIds the issue ids that are part of the board but not active on the board
	BacklogIssueIds []string `json:"backlog_issue_ids,omitempty"`
	// Columns the columns for the board in the order that they are displayed
	Columns []BoardColumns `json:"columns,omitempty"`
	// Deleted indicates that this board was deleted by the user and shouldn't be shown in the app
	Deleted *bool `json:"deleted,omitempty"`
	// Name the name of the board
	Name *string `json:"name,omitempty"`
	// Type the type of board
	Type *BoardType `json:"type,omitempty"`
	// UpdatedDate the timestamp that the board was updated
	UpdatedDate *BoardUpdatedDate `json:"updated_date,omitempty"`
}

var _ datamodel.PartialModel = (*BoardPartial)(nil)

// GetModelName returns the name of the model
func (o *BoardPartial) GetModelName() datamodel.ModelNameType {
	return BoardModelName
}

// ToMap returns the object as a map
func (o *BoardPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":            toBoardObject(o.Active, true),
		"backlog_issue_ids": toBoardObject(o.BacklogIssueIds, true),
		"columns":           toBoardObject(o.Columns, true),
		"deleted":           toBoardObject(o.Deleted, true),
		"name":              toBoardObject(o.Name, true),

		"type":         toBoardTypeEnum(o.Type),
		"updated_date": toBoardObject(o.UpdatedDate, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {

			if k == "backlog_issue_ids" {
				if arr, ok := v.([]string); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}

			if k == "columns" {
				if arr, ok := v.([]BoardColumns); ok {
					if len(arr) == 0 {
						delete(kv, k)
					}
				}
			}
			if k == "updated_date" {
				if dt, ok := v.(*BoardUpdatedDate); ok {
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
func (o *BoardPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *BoardPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *BoardPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *BoardPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *BoardPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["backlog_issue_ids"]; ok {
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
								panic("unsupported type for backlog_issue_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for backlog_issue_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for backlog_issue_ids field")
				}
			}
			o.BacklogIssueIds = na
		}
	}
	if o.BacklogIssueIds == nil {
		o.BacklogIssueIds = make([]string, 0)
	}

	if o == nil {

		o.Columns = make([]BoardColumns, 0)

	}
	if val, ok := kv["columns"]; ok {
		if sv, ok := val.([]BoardColumns); ok {
			o.Columns = sv
		} else if sp, ok := val.([]*BoardColumns); ok {
			o.Columns = o.Columns[:0]
			for _, e := range sp {
				o.Columns = append(o.Columns, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(BoardColumns); ok {
					o.Columns = append(o.Columns, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm BoardColumns
					fm.FromMap(av)
					o.Columns = append(o.Columns, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av BoardColumns
					av.FromMap(bkv)
					o.Columns = append(o.Columns, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(BoardColumns); ok {
					o.Columns = append(o.Columns, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm BoardColumns
					fm.FromMap(r)
					o.Columns = append(o.Columns, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := BoardColumns{}
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
							var fm BoardColumns
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
	if val, ok := kv["type"].(*BoardType); ok {
		o.Type = val
	} else if val, ok := kv["type"].(BoardType); ok {
		o.Type = &val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = toBoardTypePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["BoardType"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "scrum", "SCRUM":
						o.Type = toBoardTypePointer(0)
					case "kanban", "KANBAN":
						o.Type = toBoardTypePointer(1)
					case "other", "OTHER":
						o.Type = toBoardTypePointer(2)
					}
				}
			}
		}
	}

	if o.UpdatedDate == nil {
		o.UpdatedDate = &BoardUpdatedDate{}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(BoardUpdatedDate); ok {
			// struct
			o.UpdatedDate = &sv
		} else if sp, ok := val.(*BoardUpdatedDate); ok {
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

	o.setDefaults(false)
}
