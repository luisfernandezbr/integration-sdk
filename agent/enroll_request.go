// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// EnrollRequestTopic is the default topic name
	EnrollRequestTopic datamodel.TopicNameType = "agent_EnrollRequest_topic"

	// EnrollRequestStream is the default stream name
	EnrollRequestStream datamodel.TopicNameType = "agent_EnrollRequest_stream"

	// EnrollRequestTable is the default table name
	EnrollRequestTable datamodel.TopicNameType = "agent_enrollrequest"

	// EnrollRequestModelName is the model name
	EnrollRequestModelName datamodel.ModelNameType = "agent.EnrollRequest"
)

const (
	// EnrollRequestCodeColumn is the code column name
	EnrollRequestCodeColumn = "code"
	// EnrollRequestIDColumn is the id column name
	EnrollRequestIDColumn = "id"
	// EnrollRequestRequestDateColumn is the request_date column name
	EnrollRequestRequestDateColumn = "request_date"
	// EnrollRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	EnrollRequestRequestDateColumnEpochColumn = "request_date->epoch"
	// EnrollRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	EnrollRequestRequestDateColumnOffsetColumn = "request_date->offset"
	// EnrollRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	EnrollRequestRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// EnrollRequestUpdatedAtColumn is the updated_ts column name
	EnrollRequestUpdatedAtColumn = "updated_ts"
	// EnrollRequestUUIDColumn is the uuid column name
	EnrollRequestUUIDColumn = "uuid"
)

// EnrollRequestRequestDate represents the object structure for request_date
type EnrollRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *EnrollRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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

// EnrollRequest an agent request to enroll a new agent machine
type EnrollRequest struct {
	// Code The agent enrollment code
	Code string `json:"code" codec:"code" bson:"code" yaml:"code" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate EnrollRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*EnrollRequest)(nil)

func toEnrollRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollRequest:
		return v.ToMap()

	case EnrollRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of EnrollRequest
func (o *EnrollRequest) String() string {
	return fmt.Sprintf("agent.EnrollRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *EnrollRequest) GetTopicName() datamodel.TopicNameType {
	return EnrollRequestTopic
}

// GetModelName returns the name of the model
func (o *EnrollRequest) GetModelName() datamodel.ModelNameType {
	return EnrollRequestModelName
}

// GetStreamName returns the name of the stream
func (o *EnrollRequest) GetStreamName() string {
	return EnrollRequestStream.String()
}

// GetTableName returns the name of the table
func (o *EnrollRequest) GetTableName() string {
	return EnrollRequestTable.String()
}

// NewEnrollRequestID provides a template for generating an ID field for EnrollRequest
func NewEnrollRequestID(Code string, UUID string) string {
	return hash.Values(Code, UUID)
}

func (o *EnrollRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.Code, o.UUID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// GetID returns the ID for the object
func (o *EnrollRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *EnrollRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *EnrollRequest) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for EnrollRequest")
}

// IsMaterialized returns true if the model is materialized
func (o *EnrollRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *EnrollRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *EnrollRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *EnrollRequest) SetEventHeaders(kv map[string]string) {
	kv["model"] = EnrollRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *EnrollRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "uuid",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// Clone returns an exact copy of EnrollRequest
func (o *EnrollRequest) Clone() datamodel.Model {
	c := new(EnrollRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *EnrollRequest) Anon() datamodel.Model {
	c := new(EnrollRequest)
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
func (o *EnrollRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *EnrollRequest) UnmarshalJSON(data []byte) error {
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
func (o *EnrollRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two EnrollRequest objects are equal
func (o *EnrollRequest) IsEqual(other *EnrollRequest) bool {
	return o.GetID() == other.GetID()
}

// ToMap returns the object as a map
func (o *EnrollRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"code":         toEnrollRequestObject(o.Code, false),
		"id":           toEnrollRequestObject(o.ID, false),
		"request_date": toEnrollRequestObject(o.RequestDate, false),
		"updated_ts":   toEnrollRequestObject(o.UpdatedAt, false),
		"uuid":         toEnrollRequestObject(o.UUID, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["code"].(string); ok {
		o.Code = val
	} else {
		if val, ok := kv["code"]; ok {
			if val == nil {
				o.Code = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Code = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(EnrollRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*EnrollRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
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

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *EnrollRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: true,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}
