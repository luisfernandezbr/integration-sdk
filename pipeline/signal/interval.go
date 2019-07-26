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
	// IntervalTopic is the default topic name
	IntervalTopic datamodel.TopicNameType = "pipeline_signal_Interval_topic"

	// IntervalStream is the default stream name
	IntervalStream datamodel.TopicNameType = "pipeline_signal_Interval_stream"

	// IntervalTable is the default table name
	IntervalTable datamodel.TopicNameType = "pipeline_signal_Interval"

	// IntervalModelName is the model name
	IntervalModelName datamodel.ModelNameType = "pipeline.signal.Interval"
)

const (
	// IntervalEndDateColumn is the end_date column name
	IntervalEndDateColumn = "end_date"
	// IntervalEndDateColumnEpochColumn is the epoch column property of the EndDate name
	IntervalEndDateColumnEpochColumn = "end_date->epoch"
	// IntervalEndDateColumnOffsetColumn is the offset column property of the EndDate name
	IntervalEndDateColumnOffsetColumn = "end_date->offset"
	// IntervalEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	IntervalEndDateColumnRfc3339Column = "end_date->rfc3339"
	// IntervalIDColumn is the id column name
	IntervalIDColumn = "id"
	// IntervalPriorEndDateColumn is the prior_end_date column name
	IntervalPriorEndDateColumn = "prior_end_date"
	// IntervalPriorEndDateColumnEpochColumn is the epoch column property of the PriorEndDate name
	IntervalPriorEndDateColumnEpochColumn = "prior_end_date->epoch"
	// IntervalPriorEndDateColumnOffsetColumn is the offset column property of the PriorEndDate name
	IntervalPriorEndDateColumnOffsetColumn = "prior_end_date->offset"
	// IntervalPriorEndDateColumnRfc3339Column is the rfc3339 column property of the PriorEndDate name
	IntervalPriorEndDateColumnRfc3339Column = "prior_end_date->rfc3339"
	// IntervalPriorStartDateColumn is the prior_start_date column name
	IntervalPriorStartDateColumn = "prior_start_date"
	// IntervalPriorStartDateColumnEpochColumn is the epoch column property of the PriorStartDate name
	IntervalPriorStartDateColumnEpochColumn = "prior_start_date->epoch"
	// IntervalPriorStartDateColumnOffsetColumn is the offset column property of the PriorStartDate name
	IntervalPriorStartDateColumnOffsetColumn = "prior_start_date->offset"
	// IntervalPriorStartDateColumnRfc3339Column is the rfc3339 column property of the PriorStartDate name
	IntervalPriorStartDateColumnRfc3339Column = "prior_start_date->rfc3339"
	// IntervalStartDateColumn is the start_date column name
	IntervalStartDateColumn = "start_date"
	// IntervalStartDateColumnEpochColumn is the epoch column property of the StartDate name
	IntervalStartDateColumnEpochColumn = "start_date->epoch"
	// IntervalStartDateColumnOffsetColumn is the offset column property of the StartDate name
	IntervalStartDateColumnOffsetColumn = "start_date->offset"
	// IntervalStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	IntervalStartDateColumnRfc3339Column = "start_date->rfc3339"
	// IntervalTimeUnitColumn is the time_unit column name
	IntervalTimeUnitColumn = "time_unit"
)

// IntervalEndDate represents the object structure for end_date
type IntervalEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntervalEndDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntervalEndDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntervalEndDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntervalEndDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntervalEndDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntervalEndDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntervalEndDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntervalEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntervalEndDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntervalPriorEndDate represents the object structure for prior_end_date
type IntervalPriorEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntervalPriorEndDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntervalPriorEndDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntervalPriorEndDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntervalPriorEndDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntervalPriorEndDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntervalPriorEndDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntervalPriorEndDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntervalPriorEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntervalPriorEndDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntervalPriorStartDate represents the object structure for prior_start_date
type IntervalPriorStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntervalPriorStartDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntervalPriorStartDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntervalPriorStartDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntervalPriorStartDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntervalPriorStartDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntervalPriorStartDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntervalPriorStartDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntervalPriorStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntervalPriorStartDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// IntervalStartDate represents the object structure for start_date
type IntervalStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntervalStartDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntervalStartDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntervalStartDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntervalStartDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntervalStartDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntervalStartDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntervalStartDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntervalStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntervalStartDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Interval time intervals for signal data
type Interval struct {
	// EndDate the date at the end of the time interval - the current signal date
	EndDate IntervalEndDate `json:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// PriorEndDate the date at the end of the prior interval - the prior signal date
	PriorEndDate IntervalPriorEndDate `json:"prior_end_date" bson:"prior_end_date" yaml:"prior_end_date" faker:"-"`
	// PriorStartDate the date at the start of the prior time interval
	PriorStartDate IntervalPriorStartDate `json:"prior_start_date" bson:"prior_start_date" yaml:"prior_start_date" faker:"-"`
	// StartDate the date at the start of the time interval
	StartDate IntervalStartDate `json:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// TimeUnit the time unit for the time interval
	TimeUnit int64 `json:"time_unit" bson:"time_unit" yaml:"time_unit" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Interval)(nil)

func toIntervalObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntervalObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Interval:
		return v.ToMap(isavro)

	case IntervalEndDate:
		return v.ToMap(isavro)

	case IntervalPriorEndDate:
		return v.ToMap(isavro)

	case IntervalPriorStartDate:
		return v.ToMap(isavro)

	case IntervalStartDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Interval
func (o *Interval) String() string {
	return fmt.Sprintf("pipeline.signal.Interval<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Interval) GetTopicName() datamodel.TopicNameType {
	return IntervalTopic
}

// GetModelName returns the name of the model
func (o *Interval) GetModelName() datamodel.ModelNameType {
	return IntervalModelName
}

func (o *Interval) setDefaults(frommap bool) {

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// GetID returns the ID for the object
func (o *Interval) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.TimeUnit)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Interval) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Interval) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// IsMaterialized returns true if the model is materialized
func (o *Interval) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Interval) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "pipeline_signal_interval",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Interval) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Interval) SetEventHeaders(kv map[string]string) {
	kv["model"] = IntervalModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Interval) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Interval) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// Clone returns an exact copy of Interval
func (o *Interval) Clone() datamodel.Model {
	c := new(Interval)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Interval) Anon() datamodel.Model {
	c := new(Interval)
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
func (o *Interval) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Interval) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecInterval *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Interval) GetAvroCodec() *goavro.Codec {
	if cachedCodecInterval == nil {
		c, err := GetIntervalAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecInterval = c
	}
	return cachedCodecInterval
}

// ToAvroBinary returns the data as Avro binary data
func (o *Interval) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Interval) FromAvroBinary(value []byte) error {
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
func (o *Interval) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Interval objects are equal
func (o *Interval) IsEqual(other *Interval) bool {
	return o.GetID() == other.GetID()
}

// ToMap returns the object as a map
func (o *Interval) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"end_date":         toIntervalObject(o.EndDate, isavro, false, "end_date"),
		"id":               toIntervalObject(o.ID, isavro, false, "string"),
		"prior_end_date":   toIntervalObject(o.PriorEndDate, isavro, false, "prior_end_date"),
		"prior_start_date": toIntervalObject(o.PriorStartDate, isavro, false, "prior_start_date"),
		"start_date":       toIntervalObject(o.StartDate, isavro, false, "start_date"),
		"time_unit":        toIntervalObject(o.TimeUnit, isavro, false, "long"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Interval) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(IntervalEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*IntervalEndDate); ok {
			// struct pointer
			o.EndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		}
	} else {
		o.EndDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["prior_end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.PriorEndDate.FromMap(kv)
		} else if sv, ok := val.(IntervalPriorEndDate); ok {
			// struct
			o.PriorEndDate = sv
		} else if sp, ok := val.(*IntervalPriorEndDate); ok {
			// struct pointer
			o.PriorEndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.PriorEndDate.Epoch = dt.Epoch
			o.PriorEndDate.Rfc3339 = dt.Rfc3339
			o.PriorEndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.PriorEndDate.Epoch = dt.Epoch
			o.PriorEndDate.Rfc3339 = dt.Rfc3339
			o.PriorEndDate.Offset = dt.Offset
		}
	} else {
		o.PriorEndDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["prior_start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.PriorStartDate.FromMap(kv)
		} else if sv, ok := val.(IntervalPriorStartDate); ok {
			// struct
			o.PriorStartDate = sv
		} else if sp, ok := val.(*IntervalPriorStartDate); ok {
			// struct pointer
			o.PriorStartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.PriorStartDate.Epoch = dt.Epoch
			o.PriorStartDate.Rfc3339 = dt.Rfc3339
			o.PriorStartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.PriorStartDate.Epoch = dt.Epoch
			o.PriorStartDate.Rfc3339 = dt.Rfc3339
			o.PriorStartDate.Offset = dt.Offset
		}
	} else {
		o.PriorStartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(IntervalStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*IntervalStartDate); ok {
			// struct pointer
			o.StartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		}
	} else {
		o.StartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["time_unit"].(int64); ok {
		o.TimeUnit = val
	} else {
		if val, ok := kv["time_unit"]; ok {
			if val == nil {
				o.TimeUnit = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.TimeUnit = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// GetIntervalAvroSchemaSpec creates the avro schema specification for Interval
func GetIntervalAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.signal",
		"name":      "Interval",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "end_date",
				"type": map[string]interface{}{"type": "record", "name": "end_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"doc": "the timezone offset from GMT", "type": "long", "name": "offset"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date at the end of the time interval - the current signal date"},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "prior_end_date",
				"type": map[string]interface{}{"type": "record", "name": "prior_end_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date at the end of the prior interval - the prior signal date"},
			},
			map[string]interface{}{
				"name": "prior_start_date",
				"type": map[string]interface{}{"type": "record", "name": "prior_start_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"doc": "the timezone offset from GMT", "type": "long", "name": "offset"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}, "doc": "the date at the start of the prior time interval"},
			},
			map[string]interface{}{
				"name": "start_date",
				"type": map[string]interface{}{"type": "record", "name": "start_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"doc": "the timezone offset from GMT", "type": "long", "name": "offset"}, map[string]interface{}{"doc": "the date in RFC3339 format", "type": "string", "name": "rfc3339"}}, "doc": "the date at the start of the time interval"},
			},
			map[string]interface{}{
				"name": "time_unit",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Interval) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetIntervalAvroSchema creates the avro schema for Interval
func GetIntervalAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetIntervalAvroSchemaSpec())
}

// TransformIntervalFunc is a function for transforming Interval during processing
type TransformIntervalFunc func(input *Interval) (*Interval, error)

// NewIntervalPipe creates a pipe for processing Interval items
func NewIntervalPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformIntervalFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewIntervalInputStream(input, errors)
	var stream chan Interval
	if len(transforms) > 0 {
		stream = make(chan Interval, 1000)
	} else {
		stream = inch
	}
	outdone := NewIntervalOutputStream(output, stream, errors)
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

// NewIntervalInputStreamDir creates a channel for reading Interval as JSON newlines from a directory of files
func NewIntervalInputStreamDir(dir string, errors chan<- error, transforms ...TransformIntervalFunc) (chan Interval, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.signal/interval\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Interval)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for interval")
		ch := make(chan Interval)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewIntervalInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Interval)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewIntervalInputStreamFile creates an channel for reading Interval as JSON newlines from filename
func NewIntervalInputStreamFile(filename string, errors chan<- error, transforms ...TransformIntervalFunc) (chan Interval, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Interval)
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
			ch := make(chan Interval)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewIntervalInputStream(f, errors, transforms...)
}

// NewIntervalInputStream creates an channel for reading Interval as JSON newlines from stream
func NewIntervalInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformIntervalFunc) (chan Interval, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Interval, 1000)
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
			var item Interval
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

// NewIntervalOutputStreamDir will output json newlines from channel and save in dir
func NewIntervalOutputStreamDir(dir string, ch chan Interval, errors chan<- error, transforms ...TransformIntervalFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline.signal/interval\\.json(\\.gz)?$")
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
	return NewIntervalOutputStream(gz, ch, errors, transforms...)
}

// NewIntervalOutputStream will output json newlines from channel to the stream
func NewIntervalOutputStream(stream io.WriteCloser, ch chan Interval, errors chan<- error, transforms ...TransformIntervalFunc) <-chan bool {
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

// IntervalSendEvent is an event detail for sending data
type IntervalSendEvent struct {
	Interval *Interval
	headers  map[string]string
	time     time.Time
	key      string
}

var _ datamodel.ModelSendEvent = (*IntervalSendEvent)(nil)

// Key is the key to use for the message
func (e *IntervalSendEvent) Key() string {
	if e.key == "" {
		return e.Interval.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *IntervalSendEvent) Object() datamodel.Model {
	return e.Interval
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *IntervalSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *IntervalSendEvent) Timestamp() time.Time {
	return e.time
}

// IntervalSendEventOpts is a function handler for setting opts
type IntervalSendEventOpts func(o *IntervalSendEvent)

// WithIntervalSendEventKey sets the key value to a value different than the object ID
func WithIntervalSendEventKey(key string) IntervalSendEventOpts {
	return func(o *IntervalSendEvent) {
		o.key = key
	}
}

// WithIntervalSendEventTimestamp sets the timestamp value
func WithIntervalSendEventTimestamp(tv time.Time) IntervalSendEventOpts {
	return func(o *IntervalSendEvent) {
		o.time = tv
	}
}

// WithIntervalSendEventHeader sets the timestamp value
func WithIntervalSendEventHeader(key, value string) IntervalSendEventOpts {
	return func(o *IntervalSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewIntervalSendEvent returns a new IntervalSendEvent instance
func NewIntervalSendEvent(o *Interval, opts ...IntervalSendEventOpts) *IntervalSendEvent {
	res := &IntervalSendEvent{
		Interval: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewIntervalProducer will stream data from the channel
func NewIntervalProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Interval); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.signal.Interval but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewIntervalConsumer will stream data from the topic into the provided channel
func NewIntervalConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Interval
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.signal.Interval: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into pipeline.signal.Interval: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.signal.Interval")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &IntervalReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Interval
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &IntervalReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// IntervalReceiveEvent is an event detail for receiving data
type IntervalReceiveEvent struct {
	Interval *Interval
	message  eventing.Message
	eof      bool
}

var _ datamodel.ModelReceiveEvent = (*IntervalReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *IntervalReceiveEvent) Object() datamodel.Model {
	return e.Interval
}

// Message returns the underlying message data for the event
func (e *IntervalReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *IntervalReceiveEvent) EOF() bool {
	return e.eof
}

// IntervalProducer implements the datamodel.ModelEventProducer
type IntervalProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*IntervalProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *IntervalProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *IntervalProducer) Close() error {
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
func (o *Interval) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Interval) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IntervalProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIntervalProducer(newctx, producer, ch, errors, empty),
	}
}

// NewIntervalProducerChannel returns a channel which can be used for producing Model events
func NewIntervalProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewIntervalProducerChannelSize(producer, 0, errors)
}

// NewIntervalProducerChannelSize returns a channel which can be used for producing Model events
func NewIntervalProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IntervalProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIntervalProducer(newctx, producer, ch, errors, empty),
	}
}

// IntervalConsumer implements the datamodel.ModelEventConsumer
type IntervalConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*IntervalConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *IntervalConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *IntervalConsumer) Close() error {
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
func (o *Interval) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IntervalConsumer{
		ch:       ch,
		callback: NewIntervalConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewIntervalConsumerChannel returns a consumer channel which can be used to consume Model events
func NewIntervalConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IntervalConsumer{
		ch:       ch,
		callback: NewIntervalConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
