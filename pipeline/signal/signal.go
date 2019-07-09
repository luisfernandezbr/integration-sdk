// DO NOT EDIT -- generated code

// Package signal - the pipeline models
package signal

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
	// SignalTopic is the default topic name
	SignalTopic datamodel.TopicNameType = "pipeline_signal_Signal_topic"

	// SignalStream is the default stream name
	SignalStream datamodel.TopicNameType = "pipeline_signal_Signal_stream"

	// SignalTable is the default table name
	SignalTable datamodel.TopicNameType = "pipeline_signal_Signal"

	// SignalModelName is the model name
	SignalModelName datamodel.ModelNameType = "pipeline.signal.Signal"
)

const (
	// SignalActiveColumn is the active column name
	SignalActiveColumn = "active"
	// SignalCustomerIDColumn is the customer_id column name
	SignalCustomerIDColumn = "customer_id"
	// SignalDateColumn is the date column name
	SignalDateColumn = "date"
	// SignalDateAtColumn is the date_ts column name
	SignalDateAtColumn = "date_ts"
	// SignalIDColumn is the id column name
	SignalIDColumn = "id"
	// SignalMetadataColumn is the metadata column name
	SignalMetadataColumn = "metadata"
	// SignalNameColumn is the name column name
	SignalNameColumn = "name"
	// SignalPercentileRankColumn is the percentile_rank column name
	SignalPercentileRankColumn = "percentile_rank"
	// SignalRefIDColumn is the ref_id column name
	SignalRefIDColumn = "ref_id"
	// SignalRefKeyColumn is the ref_key column name
	SignalRefKeyColumn = "ref_key"
	// SignalRefNameColumn is the ref_name column name
	SignalRefNameColumn = "ref_name"
	// SignalRefTypeColumn is the ref_type column name
	SignalRefTypeColumn = "ref_type"
	// SignalTimeUnitColumn is the time_unit column name
	SignalTimeUnitColumn = "time_unit"
	// SignalTrendColumn is the trend column name
	SignalTrendColumn = "trend"
	// SignalValueColumn is the value column name
	SignalValueColumn = "value"
)

// SignalDate represents the object structure for date
type SignalDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *SignalDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Signal signal table contains all the calculated data from the data analytics system
type Signal struct {
	// Active the active state for the signal
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Date date object
	Date SignalDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// DateAt the date when the signal was calculated
	DateAt int64 `json:"date_ts" bson:"date_ts" yaml:"date_ts" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Metadata the metadata for the signal
	Metadata string `json:"metadata" bson:"metadata" yaml:"metadata" faker:"-"`
	// Name the name for the signal
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// PercentileRank the percentile_rank for the signal
	PercentileRank float64 `json:"percentile_rank" bson:"percentile_rank" yaml:"percentile_rank" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefKey the ref_key for the signal
	RefKey string `json:"ref_key" bson:"ref_key" yaml:"ref_key" faker:"-"`
	// RefName the name of the entity for which the signal was calculated
	RefName string `json:"ref_name" bson:"ref_name" yaml:"ref_name" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TimeUnit the time unit for the signal
	TimeUnit int64 `json:"time_unit" bson:"time_unit" yaml:"time_unit" faker:"-"`
	// Trend the trend detail for the signal
	Trend string `json:"trend" bson:"trend" yaml:"trend" faker:"-"`
	// Value the value for the signal
	Value float64 `json:"value" bson:"value" yaml:"value" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Signal)(nil)

func toSignalObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toSignalObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toSignalObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toSignalObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toSignalObjectNil(isavro, isoptional)
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
	case *Signal:
		return v.ToMap()
	case Signal:
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
			arr = append(arr, toSignalObject(av, isavro, false, ""))
		}
		return arr

	case SignalDate:
		vv := o.(SignalDate)
		return vv.ToMap()
	case *SignalDate:
		return (*o.(*SignalDate)).ToMap()
	case []SignalDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]SignalDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]SignalDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]SignalDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Signal
func (o *Signal) String() string {
	return fmt.Sprintf("pipeline.signal.Signal<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Signal) GetTopicName() datamodel.TopicNameType {
	return SignalTopic
}

// GetModelName returns the name of the model
func (o *Signal) GetModelName() datamodel.ModelNameType {
	return SignalModelName
}

func (o *Signal) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Signal) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Signal", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Signal) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Signal) GetTimestamp() time.Time {
	var dt interface{} = o.DateAt
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
	panic("not sure how to handle the date time format for Signal")
}

// GetRefID returns the RefID for the object
func (o *Signal) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Signal) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Signal) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "pipeline_signal_signal",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Signal) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Signal) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = SignalModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Signal) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "date_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetStateKey returns a key for use in state store
func (o *Signal) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Signal) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Signal
func (o *Signal) Clone() datamodel.Model {
	c := new(Signal)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Signal) Anon() datamodel.Model {
	c := new(Signal)
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
func (o *Signal) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Signal) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecSignal *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Signal) GetAvroCodec() *goavro.Codec {
	if cachedCodecSignal == nil {
		c, err := GetSignalAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecSignal = c
	}
	return cachedCodecSignal
}

// ToAvroBinary returns the data as Avro binary data
func (o *Signal) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Signal) FromAvroBinary(value []byte) error {
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
func (o *Signal) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Signal objects are equal
func (o *Signal) IsEqual(other *Signal) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Signal) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"active":          toSignalObject(o.Active, isavro, false, "boolean"),
		"customer_id":     toSignalObject(o.CustomerID, isavro, false, "string"),
		"date":            toSignalObject(o.Date, isavro, false, "date"),
		"date_ts":         toSignalObject(o.DateAt, isavro, false, "long"),
		"id":              toSignalObject(o.ID, isavro, false, "string"),
		"metadata":        toSignalObject(o.Metadata, isavro, false, "string"),
		"name":            toSignalObject(o.Name, isavro, false, "string"),
		"percentile_rank": toSignalObject(o.PercentileRank, isavro, false, "double"),
		"ref_id":          toSignalObject(o.RefID, isavro, false, "string"),
		"ref_key":         toSignalObject(o.RefKey, isavro, false, "string"),
		"ref_name":        toSignalObject(o.RefName, isavro, false, "string"),
		"ref_type":        toSignalObject(o.RefType, isavro, false, "string"),
		"time_unit":       toSignalObject(o.TimeUnit, isavro, false, "long"),
		"trend":           toSignalObject(o.Trend, isavro, false, "string"),
		"value":           toSignalObject(o.Value, isavro, false, "double"),
		"hashcode":        toSignalObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Signal) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		val := kv["active"]
		if val == nil {
			o.Active = number.ToBoolAny(nil)
		} else {
			o.Active = number.ToBoolAny(val)
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
	if val, ok := kv["date"].(SignalDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = SignalDate{}
		} else {
			o.Date = SignalDate{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Date)

		}
	}
	if val, ok := kv["date_ts"].(int64); ok {
		o.DateAt = val
	} else {
		val := kv["date_ts"]
		if val == nil {
			o.DateAt = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.DateAt = number.ToInt64Any(val)
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
	if val, ok := kv["metadata"].(string); ok {
		o.Metadata = val
	} else {
		val := kv["metadata"]
		if val == nil {
			o.Metadata = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Metadata = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["percentile_rank"].(float64); ok {
		o.PercentileRank = val
	} else {
		val := kv["percentile_rank"]
		if val == nil {
			o.PercentileRank = number.ToFloat64Any(nil)
		} else {
			o.PercentileRank = number.ToFloat64Any(val)
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
	if val, ok := kv["ref_key"].(string); ok {
		o.RefKey = val
	} else {
		val := kv["ref_key"]
		if val == nil {
			o.RefKey = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefKey = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_name"].(string); ok {
		o.RefName = val
	} else {
		val := kv["ref_name"]
		if val == nil {
			o.RefName = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefName = fmt.Sprintf("%v", val)
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
	if val, ok := kv["time_unit"].(int64); ok {
		o.TimeUnit = val
	} else {
		val := kv["time_unit"]
		if val == nil {
			o.TimeUnit = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.TimeUnit = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["trend"].(string); ok {
		o.Trend = val
	} else {
		val := kv["trend"]
		if val == nil {
			o.Trend = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Trend = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["value"].(float64); ok {
		o.Value = val
	} else {
		val := kv["value"]
		if val == nil {
			o.Value = number.ToFloat64Any(nil)
		} else {
			o.Value = number.ToFloat64Any(val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Signal) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CustomerID)
	args = append(args, o.Date)
	args = append(args, o.DateAt)
	args = append(args, o.ID)
	args = append(args, o.Metadata)
	args = append(args, o.Name)
	args = append(args, o.PercentileRank)
	args = append(args, o.RefID)
	args = append(args, o.RefKey)
	args = append(args, o.RefName)
	args = append(args, o.RefType)
	args = append(args, o.TimeUnit)
	args = append(args, o.Trend)
	args = append(args, o.Value)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetSignalAvroSchemaSpec creates the avro schema specification for Signal
func GetSignalAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.signal",
		"name":      "Signal",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "date",
				"type": map[string]interface{}{"type": "record", "name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "date object"},
			},
			map[string]interface{}{
				"name": "date_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "metadata",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "percentile_rank",
				"type": "double",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_key",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "time_unit",
				"type": "long",
			},
			map[string]interface{}{
				"name": "trend",
				"type": "string",
			},
			map[string]interface{}{
				"name": "value",
				"type": "double",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetSignalAvroSchema creates the avro schema for Signal
func GetSignalAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetSignalAvroSchemaSpec())
}

// TransformSignalFunc is a function for transforming Signal during processing
type TransformSignalFunc func(input *Signal) (*Signal, error)

// NewSignalPipe creates a pipe for processing Signal items
func NewSignalPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformSignalFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewSignalInputStream(input, errors)
	var stream chan Signal
	if len(transforms) > 0 {
		stream = make(chan Signal, 1000)
	} else {
		stream = inch
	}
	outdone := NewSignalOutputStream(output, stream, errors)
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

// NewSignalInputStreamDir creates a channel for reading Signal as JSON newlines from a directory of files
func NewSignalInputStreamDir(dir string, errors chan<- error, transforms ...TransformSignalFunc) (chan Signal, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.signal/signal\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Signal)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for signal")
		ch := make(chan Signal)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewSignalInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Signal)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewSignalInputStreamFile creates an channel for reading Signal as JSON newlines from filename
func NewSignalInputStreamFile(filename string, errors chan<- error, transforms ...TransformSignalFunc) (chan Signal, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Signal)
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
			ch := make(chan Signal)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewSignalInputStream(f, errors, transforms...)
}

// NewSignalInputStream creates an channel for reading Signal as JSON newlines from stream
func NewSignalInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformSignalFunc) (chan Signal, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Signal, 1000)
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
			var item Signal
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

// NewSignalOutputStreamDir will output json newlines from channel and save in dir
func NewSignalOutputStreamDir(dir string, ch chan Signal, errors chan<- error, transforms ...TransformSignalFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline.signal/signal\\.json(\\.gz)?$")
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
	return NewSignalOutputStream(gz, ch, errors, transforms...)
}

// NewSignalOutputStream will output json newlines from channel to the stream
func NewSignalOutputStream(stream io.WriteCloser, ch chan Signal, errors chan<- error, transforms ...TransformSignalFunc) <-chan bool {
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

// SignalSendEvent is an event detail for sending data
type SignalSendEvent struct {
	Signal  *Signal
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*SignalSendEvent)(nil)

// Key is the key to use for the message
func (e *SignalSendEvent) Key() string {
	if e.key == "" {
		return e.Signal.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *SignalSendEvent) Object() datamodel.Model {
	return e.Signal
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *SignalSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *SignalSendEvent) Timestamp() time.Time {
	return e.time
}

// SignalSendEventOpts is a function handler for setting opts
type SignalSendEventOpts func(o *SignalSendEvent)

// WithSignalSendEventKey sets the key value to a value different than the object ID
func WithSignalSendEventKey(key string) SignalSendEventOpts {
	return func(o *SignalSendEvent) {
		o.key = key
	}
}

// WithSignalSendEventTimestamp sets the timestamp value
func WithSignalSendEventTimestamp(tv time.Time) SignalSendEventOpts {
	return func(o *SignalSendEvent) {
		o.time = tv
	}
}

// WithSignalSendEventHeader sets the timestamp value
func WithSignalSendEventHeader(key, value string) SignalSendEventOpts {
	return func(o *SignalSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewSignalSendEvent returns a new SignalSendEvent instance
func NewSignalSendEvent(o *Signal, opts ...SignalSendEventOpts) *SignalSendEvent {
	res := &SignalSendEvent{
		Signal: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewSignalProducer will stream data from the channel
func NewSignalProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Signal); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.signal.Signal but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewSignalConsumer will stream data from the topic into the provided channel
func NewSignalConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Signal
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.signal.Signal: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into pipeline.signal.Signal: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.signal.Signal")
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &SignalReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Signal
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &SignalReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// SignalReceiveEvent is an event detail for receiving data
type SignalReceiveEvent struct {
	Signal  *Signal
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*SignalReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *SignalReceiveEvent) Object() datamodel.Model {
	return e.Signal
}

// Message returns the underlying message data for the event
func (e *SignalReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *SignalReceiveEvent) EOF() bool {
	return e.eof
}

// SignalProducer implements the datamodel.ModelEventProducer
type SignalProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*SignalProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *SignalProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *SignalProducer) Close() error {
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
func (o *Signal) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Signal) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &SignalProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewSignalProducer(newctx, producer, ch, errors, empty),
	}
}

// NewSignalProducerChannel returns a channel which can be used for producing Model events
func NewSignalProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewSignalProducerChannelSize(producer, 0, errors)
}

// NewSignalProducerChannelSize returns a channel which can be used for producing Model events
func NewSignalProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &SignalProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewSignalProducer(newctx, producer, ch, errors, empty),
	}
}

// SignalConsumer implements the datamodel.ModelEventConsumer
type SignalConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*SignalConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *SignalConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *SignalConsumer) Close() error {
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
func (o *Signal) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &SignalConsumer{
		ch:       ch,
		callback: NewSignalConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewSignalConsumerChannel returns a consumer channel which can be used to consume Model events
func NewSignalConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &SignalConsumer{
		ch:       ch,
		callback: NewSignalConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
