// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// ChangelogTopic is the default topic name
	ChangelogTopic datamodel.TopicNameType = "work_Changelog_topic"

	// ChangelogStream is the default stream name
	ChangelogStream datamodel.TopicNameType = "work_Changelog_stream"

	// ChangelogTable is the default table name
	ChangelogTable datamodel.TopicNameType = "work_Changelog"

	// ChangelogModelName is the model name
	ChangelogModelName datamodel.ModelNameType = "work.Changelog"
)

const (
	// ChangelogCreatedColumn is the created column name
	ChangelogCreatedColumn = "created"
	// ChangelogCreatedColumnEpochColumn is the epoch column property of the Created name
	ChangelogCreatedColumnEpochColumn = "created->epoch"
	// ChangelogCreatedColumnOffsetColumn is the offset column property of the Created name
	ChangelogCreatedColumnOffsetColumn = "created->offset"
	// ChangelogCreatedColumnRfc3339Column is the rfc3339 column property of the Created name
	ChangelogCreatedColumnRfc3339Column = "created->rfc3339"
	// ChangelogCustomerIDColumn is the customer_id column name
	ChangelogCustomerIDColumn = "customer_id"
	// ChangelogFieldColumn is the field column name
	ChangelogFieldColumn = "field"
	// ChangelogFieldTypeColumn is the field_type column name
	ChangelogFieldTypeColumn = "field_type"
	// ChangelogFromColumn is the from column name
	ChangelogFromColumn = "from"
	// ChangelogFromStringColumn is the from_string column name
	ChangelogFromStringColumn = "from_string"
	// ChangelogIDColumn is the id column name
	ChangelogIDColumn = "id"
	// ChangelogIssueIDColumn is the issue_id column name
	ChangelogIssueIDColumn = "issue_id"
	// ChangelogOrdinalColumn is the ordinal column name
	ChangelogOrdinalColumn = "ordinal"
	// ChangelogRefIDColumn is the ref_id column name
	ChangelogRefIDColumn = "ref_id"
	// ChangelogRefTypeColumn is the ref_type column name
	ChangelogRefTypeColumn = "ref_type"
	// ChangelogToColumn is the to column name
	ChangelogToColumn = "to"
	// ChangelogToStringColumn is the to_string column name
	ChangelogToStringColumn = "to_string"
	// ChangelogUserIDColumn is the user_id column name
	ChangelogUserIDColumn = "user_id"
)

// ChangelogCreated represents the object structure for created
type ChangelogCreated struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *ChangelogCreated) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Changelog change log
type Changelog struct {
	// Created the date when this change was created
	Created ChangelogCreated `json:"created" bson:"created" yaml:"created" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Field name of the field that was changed
	Field string `json:"field" bson:"field" yaml:"field" faker:"-"`
	// FieldType type of the field that was changed
	FieldType string `json:"field_type" bson:"field_type" yaml:"field_type" faker:"-"`
	// From id of the change from
	From string `json:"from" bson:"from" yaml:"from" faker:"-"`
	// FromString name of the change from
	FromString string `json:"from_string" bson:"from_string" yaml:"from_string" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IssueID id of the issue
	IssueID string `json:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// Ordinal so we can order correctly in queries since dates could be equal
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// To id of the change to
	To string `json:"to" bson:"to" yaml:"to" faker:"-"`
	// ToString name of the change to
	ToString string `json:"to_string" bson:"to_string" yaml:"to_string" faker:"-"`
	// UserID id of the user of this change
	UserID string `json:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Changelog)(nil)

func toChangelogObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toChangelogObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toChangelogObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toChangelogObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *Changelog:
		return v.ToMap()
	case Changelog:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toChangelogObject(av, isavro, false, ""))
		}
		return arr

	case ChangelogCreated:
		vv := o.(ChangelogCreated)
		return vv.ToMap()
	case *ChangelogCreated:
		return map[string]interface{}{
			"work.created": (*o.(*ChangelogCreated)).ToMap(),
		}
	case []ChangelogCreated:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ChangelogCreated) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ChangelogCreated:
		arr := make([]interface{}, 0)
		vv := o.(*[]ChangelogCreated)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("work.Changelog<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Changelog) GetTopicName() datamodel.TopicNameType {
	return ChangelogTopic
}

// GetModelName returns the name of the model
func (o *Changelog) GetModelName() datamodel.ModelNameType {
	return ChangelogModelName
}

func (o *Changelog) setDefaults() {

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Changelog) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Changelog", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Changelog) GetTopicKey() string {
	var i interface{} = o.IssueID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Changelog) GetTimestamp() time.Time {
	var dt interface{} = o.Created
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
	case ChangelogCreated:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Changelog")
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Changelog) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Changelog) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Changelog) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Changelog) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ChangelogModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Changelog) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "issue_id",
		Timestamp:         "created",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Changelog) GetStateKey() string {
	key := "issue_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Changelog) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Changelog
func (o *Changelog) Clone() datamodel.Model {
	c := new(Changelog)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Changelog) Anon() datamodel.Model {
	c := new(Changelog)
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
func (o *Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Changelog) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecChangelog *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Changelog) GetAvroCodec() *goavro.Codec {
	if cachedCodecChangelog == nil {
		c, err := GetChangelogAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecChangelog = c
	}
	return cachedCodecChangelog
}

// ToAvroBinary returns the data as Avro binary data
func (o *Changelog) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *Changelog) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Changelog) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"created":     toChangelogObject(o.Created, isavro, false, "created"),
		"customer_id": toChangelogObject(o.CustomerID, isavro, false, "string"),
		"field":       toChangelogObject(o.Field, isavro, false, "string"),
		"field_type":  toChangelogObject(o.FieldType, isavro, false, "string"),
		"from":        toChangelogObject(o.From, isavro, false, "string"),
		"from_string": toChangelogObject(o.FromString, isavro, false, "string"),
		"id":          toChangelogObject(o.ID, isavro, false, "string"),
		"issue_id":    toChangelogObject(o.IssueID, isavro, false, "string"),
		"ordinal":     toChangelogObject(o.Ordinal, isavro, false, "long"),
		"ref_id":      toChangelogObject(o.RefID, isavro, false, "string"),
		"ref_type":    toChangelogObject(o.RefType, isavro, false, "string"),
		"to":          toChangelogObject(o.To, isavro, false, "string"),
		"to_string":   toChangelogObject(o.ToString, isavro, false, "string"),
		"user_id":     toChangelogObject(o.UserID, isavro, false, "string"),
		"hashcode":    toChangelogObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["created"].(ChangelogCreated); ok {
		o.Created = val
	} else {
		val := kv["created"]
		if val == nil {
			o.Created = ChangelogCreated{}
		} else {
			o.Created = ChangelogCreated{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Created)

		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field"].(string); ok {
		o.Field = val
	} else {
		val := kv["field"]
		if val == nil {
			o.Field = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Field = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field_type"].(string); ok {
		o.FieldType = val
	} else {
		val := kv["field_type"]
		if val == nil {
			o.FieldType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.FieldType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from"].(string); ok {
		o.From = val
	} else {
		val := kv["from"]
		if val == nil {
			o.From = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.From = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from_string"].(string); ok {
		o.FromString = val
	} else {
		val := kv["from_string"]
		if val == nil {
			o.FromString = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.FromString = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		val := kv["issue_id"]
		if val == nil {
			o.IssueID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.IssueID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		if val == nil {
			o.Ordinal = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Ordinal = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to"].(string); ok {
		o.To = val
	} else {
		val := kv["to"]
		if val == nil {
			o.To = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.To = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to_string"].(string); ok {
		o.ToString = val
	} else {
		val := kv["to_string"]
		if val == nil {
			o.ToString = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ToString = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		val := kv["user_id"]
		if val == nil {
			o.UserID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UserID = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Created)
	args = append(args, o.CustomerID)
	args = append(args, o.Field)
	args = append(args, o.FieldType)
	args = append(args, o.From)
	args = append(args, o.FromString)
	args = append(args, o.ID)
	args = append(args, o.IssueID)
	args = append(args, o.Ordinal)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.To)
	args = append(args, o.ToString)
	args = append(args, o.UserID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func GetChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "work",
		"name":      "Changelog",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created",
				"type": map[string]interface{}{"type": "record", "name": "created", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when this change was created"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from_string",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "issue_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to_string",
				"type": "string",
			},
			map[string]interface{}{
				"name": "user_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Changelog) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetChangelogAvroSchema creates the avro schema for Changelog
func GetChangelogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetChangelogAvroSchemaSpec())
}

// TransformChangelogFunc is a function for transforming Changelog during processing
type TransformChangelogFunc func(input *Changelog) (*Changelog, error)

// NewChangelogPipe creates a pipe for processing Changelog items
func NewChangelogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewChangelogInputStream(input, errors)
	var stream chan Changelog
	if len(transforms) > 0 {
		stream = make(chan Changelog, 1000)
	} else {
		stream = inch
	}
	outdone := NewChangelogOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewChangelogInputStreamDir creates a channel for reading Changelog as JSON newlines from a directory of files
func NewChangelogInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/changelog\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for changelog")
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewChangelogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewChangelogInputStreamFile creates an channel for reading Changelog as JSON newlines from filename
func NewChangelogInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan Changelog)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewChangelogInputStream(f, errors, transforms...)
}

// NewChangelogInputStream creates an channel for reading Changelog as JSON newlines from stream
func NewChangelogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Changelog, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item Changelog
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewChangelogOutputStreamDir will output json newlines from channel and save in dir
func NewChangelogOutputStreamDir(dir string, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/changelog\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewChangelogOutputStream(gz, ch, errors, transforms...)
}

// NewChangelogOutputStream will output json newlines from channel to the stream
func NewChangelogOutputStream(stream io.WriteCloser, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// ChangelogSendEvent is an event detail for sending data
type ChangelogSendEvent struct {
	Changelog *Changelog
	headers   map[string]string
	time      time.Time
	key       string
}

var _ datamodel.ModelSendEvent = (*ChangelogSendEvent)(nil)

// Key is the key to use for the message
func (e *ChangelogSendEvent) Key() string {
	if e.key == "" {
		return e.Changelog.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ChangelogSendEvent) Object() datamodel.Model {
	return e.Changelog
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ChangelogSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ChangelogSendEvent) Timestamp() time.Time {
	return e.time
}

// ChangelogSendEventOpts is a function handler for setting opts
type ChangelogSendEventOpts func(o *ChangelogSendEvent)

// WithChangelogSendEventKey sets the key value to a value different than the object ID
func WithChangelogSendEventKey(key string) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		o.key = key
	}
}

// WithChangelogSendEventTimestamp sets the timestamp value
func WithChangelogSendEventTimestamp(tv time.Time) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		o.time = tv
	}
}

// WithChangelogSendEventHeader sets the timestamp value
func WithChangelogSendEventHeader(key, value string) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewChangelogSendEvent returns a new ChangelogSendEvent instance
func NewChangelogSendEvent(o *Changelog, opts ...ChangelogSendEventOpts) *ChangelogSendEvent {
	res := &ChangelogSendEvent{
		Changelog: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewChangelogProducer will stream data from the channel
func NewChangelogProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*Changelog); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type work.Changelog but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewChangelogConsumer will stream data from the topic into the provided channel
func NewChangelogConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Changelog
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.Changelog: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into work.Changelog: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for work.Changelog")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &ChangelogReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Changelog
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ChangelogReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ChangelogReceiveEvent is an event detail for receiving data
type ChangelogReceiveEvent struct {
	Changelog *Changelog
	message   eventing.Message
	eof       bool
}

var _ datamodel.ModelReceiveEvent = (*ChangelogReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ChangelogReceiveEvent) Object() datamodel.Model {
	return e.Changelog
}

// Message returns the underlying message data for the event
func (e *ChangelogReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ChangelogReceiveEvent) EOF() bool {
	return e.eof
}

// ChangelogProducer implements the datamodel.ModelEventProducer
type ChangelogProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ChangelogProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ChangelogProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ChangelogProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Changelog) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Changelog) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ChangelogProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewChangelogProducer(newctx, producer, ch, errors, empty),
	}
}

// NewChangelogProducerChannel returns a channel which can be used for producing Model events
func NewChangelogProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewChangelogProducerChannelSize(producer, 0, errors)
}

// NewChangelogProducerChannelSize returns a channel which can be used for producing Model events
func NewChangelogProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ChangelogProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewChangelogProducer(newctx, producer, ch, errors, empty),
	}
}

// ChangelogConsumer implements the datamodel.ModelEventConsumer
type ChangelogConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ChangelogConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ChangelogConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ChangelogConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Changelog) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ChangelogConsumer{
		ch:       ch,
		callback: NewChangelogConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewChangelogConsumerChannel returns a consumer channel which can be used to consume Model events
func NewChangelogConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ChangelogConsumer{
		ch:       ch,
		callback: NewChangelogConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
