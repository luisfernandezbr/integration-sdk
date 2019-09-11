// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollRequestRequestDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toEnrollRequestRequestDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *EnrollRequestRequestDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *EnrollRequestRequestDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollRequestRequestDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toEnrollRequestRequestDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollRequestRequestDateObject(o.Rfc3339, isavro, false, "string"),
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
	Code string `json:"code" bson:"code" yaml:"code" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate EnrollRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*EnrollRequest)(nil)

func toEnrollRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toEnrollRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *EnrollRequest:
		return v.ToMap(isavro)

	case EnrollRequestRequestDate:
		return v.ToMap(isavro)

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

// GetStateKey returns a key for use in state store
func (o *EnrollRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *EnrollRequest) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *EnrollRequest) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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

var cachedCodecEnrollRequest *goavro.Codec
var cachedCodecEnrollRequestLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *EnrollRequest) GetAvroCodec() *goavro.Codec {
	cachedCodecEnrollRequestLock.Lock()
	if cachedCodecEnrollRequest == nil {
		c, err := GetEnrollRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecEnrollRequest = c
	}
	cachedCodecEnrollRequestLock.Unlock()
	return cachedCodecEnrollRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *EnrollRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *EnrollRequest) FromAvroBinary(value []byte) error {
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
func (o *EnrollRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two EnrollRequest objects are equal
func (o *EnrollRequest) IsEqual(other *EnrollRequest) bool {
	return o.GetID() == other.GetID()
}

// ToMap returns the object as a map
func (o *EnrollRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"code":         toEnrollRequestObject(o.Code, isavro, false, "string"),
		"id":           toEnrollRequestObject(o.ID, isavro, false, "string"),
		"request_date": toEnrollRequestObject(o.RequestDate, isavro, false, "request_date"),
		"updated_ts":   toEnrollRequestObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":         toEnrollRequestObject(o.UUID, isavro, false, "string"),
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

// GetEnrollRequestAvroSchemaSpec creates the avro schema specification for EnrollRequest
func GetEnrollRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "EnrollRequest",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "code",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "request_date",
				"type": map[string]interface{}{"doc": "the date when the request was made", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "request_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
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

// GetEnrollRequestAvroSchema creates the avro schema for EnrollRequest
func GetEnrollRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetEnrollRequestAvroSchemaSpec())
}

// EnrollRequestSendEvent is an event detail for sending data
type EnrollRequestSendEvent struct {
	EnrollRequest *EnrollRequest
	headers       map[string]string
	time          time.Time
	key           string
}

var _ datamodel.ModelSendEvent = (*EnrollRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *EnrollRequestSendEvent) Key() string {
	if e.key == "" {
		return e.EnrollRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *EnrollRequestSendEvent) Object() datamodel.Model {
	return e.EnrollRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *EnrollRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *EnrollRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// EnrollRequestSendEventOpts is a function handler for setting opts
type EnrollRequestSendEventOpts func(o *EnrollRequestSendEvent)

// WithEnrollRequestSendEventKey sets the key value to a value different than the object ID
func WithEnrollRequestSendEventKey(key string) EnrollRequestSendEventOpts {
	return func(o *EnrollRequestSendEvent) {
		o.key = key
	}
}

// WithEnrollRequestSendEventTimestamp sets the timestamp value
func WithEnrollRequestSendEventTimestamp(tv time.Time) EnrollRequestSendEventOpts {
	return func(o *EnrollRequestSendEvent) {
		o.time = tv
	}
}

// WithEnrollRequestSendEventHeader sets the timestamp value
func WithEnrollRequestSendEventHeader(key, value string) EnrollRequestSendEventOpts {
	return func(o *EnrollRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewEnrollRequestSendEvent returns a new EnrollRequestSendEvent instance
func NewEnrollRequestSendEvent(o *EnrollRequest, opts ...EnrollRequestSendEventOpts) *EnrollRequestSendEvent {
	res := &EnrollRequestSendEvent{
		EnrollRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewEnrollRequestProducer will stream data from the channel
func NewEnrollRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
	var numPartitions int
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
				if object, ok := item.Object().(*EnrollRequest); ok {
					if numPartitions == 0 {
						numPartitions = object.GetTopicConfig().NumPartitions
					}
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
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					// determine the partition selection by using the partition key
					// and taking the modulo over the number of partitions for the topic
					partition := hash.Modulo(item.Key(), numPartitions)
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.(ModelWithID).GetID(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: int32(partition),
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.EnrollRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewEnrollRequestConsumer will stream data from the topic into the provided channel
func NewEnrollRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object EnrollRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.EnrollRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.EnrollRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.EnrollRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &EnrollRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object EnrollRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &EnrollRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// EnrollRequestReceiveEvent is an event detail for receiving data
type EnrollRequestReceiveEvent struct {
	EnrollRequest *EnrollRequest
	message       eventing.Message
	eof           bool
}

var _ datamodel.ModelReceiveEvent = (*EnrollRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *EnrollRequestReceiveEvent) Object() datamodel.Model {
	return e.EnrollRequest
}

// Message returns the underlying message data for the event
func (e *EnrollRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *EnrollRequestReceiveEvent) EOF() bool {
	return e.eof
}

// EnrollRequestProducer implements the datamodel.ModelEventProducer
type EnrollRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*EnrollRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *EnrollRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *EnrollRequestProducer) Close() error {
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
func (o *EnrollRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *EnrollRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnrollRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnrollRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewEnrollRequestProducerChannel returns a channel which can be used for producing Model events
func NewEnrollRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewEnrollRequestProducerChannelSize(producer, 0, errors)
}

// NewEnrollRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewEnrollRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &EnrollRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewEnrollRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// EnrollRequestConsumer implements the datamodel.ModelEventConsumer
type EnrollRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*EnrollRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *EnrollRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *EnrollRequestConsumer) Close() error {
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
func (o *EnrollRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnrollRequestConsumer{
		ch:       ch,
		callback: NewEnrollRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewEnrollRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewEnrollRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &EnrollRequestConsumer{
		ch:       ch,
		callback: NewEnrollRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
