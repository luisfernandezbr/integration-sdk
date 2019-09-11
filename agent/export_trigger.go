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
	// ExportTriggerTopic is the default topic name
	ExportTriggerTopic datamodel.TopicNameType = "agent_ExportTrigger_topic"

	// ExportTriggerStream is the default stream name
	ExportTriggerStream datamodel.TopicNameType = "agent_ExportTrigger_stream"

	// ExportTriggerTable is the default table name
	ExportTriggerTable datamodel.TopicNameType = "agent_exporttrigger"

	// ExportTriggerModelName is the model name
	ExportTriggerModelName datamodel.ModelNameType = "agent.ExportTrigger"
)

const (
	// ExportTriggerCustomerIDColumn is the customer_id column name
	ExportTriggerCustomerIDColumn = "customer_id"
	// ExportTriggerIDColumn is the id column name
	ExportTriggerIDColumn = "id"
	// ExportTriggerRefIDColumn is the ref_id column name
	ExportTriggerRefIDColumn = "ref_id"
	// ExportTriggerRefTypeColumn is the ref_type column name
	ExportTriggerRefTypeColumn = "ref_type"
	// ExportTriggerReprocessHistoricalColumn is the reprocess_historical column name
	ExportTriggerReprocessHistoricalColumn = "reprocess_historical"
	// ExportTriggerUpdatedAtColumn is the updated_ts column name
	ExportTriggerUpdatedAtColumn = "updated_ts"
	// ExportTriggerUUIDColumn is the uuid column name
	ExportTriggerUUIDColumn = "uuid"
)

// ExportTrigger used to trigger an agent.ExportRequest
type ExportTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical bool `json:"reprocess_historical" bson:"reprocess_historical" yaml:"reprocess_historical" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID This UUID of the agent to trigger
	UUID *string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportTrigger)(nil)

func toExportTriggerObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toExportTriggerObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *ExportTrigger:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of ExportTrigger
func (o *ExportTrigger) String() string {
	return fmt.Sprintf("agent.ExportTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportTrigger) GetTopicName() datamodel.TopicNameType {
	return ExportTriggerTopic
}

// GetModelName returns the name of the model
func (o *ExportTrigger) GetModelName() datamodel.ModelNameType {
	return ExportTriggerModelName
}

// GetStreamName returns the name of the stream
func (o *ExportTrigger) GetStreamName() string {
	return ExportTriggerStream.String()
}

// GetTableName returns the name of the table
func (o *ExportTrigger) GetTableName() string {
	return ExportTriggerTable.String()
}

// NewExportTriggerID provides a template for generating an ID field for ExportTrigger
func NewExportTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("ExportTrigger", customerID, refType, refID)
}

func (o *ExportTrigger) setDefaults(frommap bool) {
	if o.UUID == nil {
		o.UUID = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportTrigger) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportTrigger) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for ExportTrigger")
}

// GetRefID returns the RefID for the object
func (o *ExportTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportTrigger) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportTrigger) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("24h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 24h0m0s. " + err.Error())
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
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *ExportTrigger) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *ExportTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ExportTrigger
func (o *ExportTrigger) Clone() datamodel.Model {
	c := new(ExportTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportTrigger) Anon() datamodel.Model {
	c := new(ExportTrigger)
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
func (o *ExportTrigger) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *ExportTrigger) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportTrigger) UnmarshalJSON(data []byte) error {
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

var cachedCodecExportTrigger *goavro.Codec
var cachedCodecExportTriggerLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *ExportTrigger) GetAvroCodec() *goavro.Codec {
	cachedCodecExportTriggerLock.Lock()
	if cachedCodecExportTrigger == nil {
		c, err := GetExportTriggerAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecExportTrigger = c
	}
	cachedCodecExportTriggerLock.Unlock()
	return cachedCodecExportTrigger
}

// ToAvroBinary returns the data as Avro binary data
func (o *ExportTrigger) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *ExportTrigger) FromAvroBinary(value []byte) error {
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
func (o *ExportTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ExportTrigger objects are equal
func (o *ExportTrigger) IsEqual(other *ExportTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportTrigger) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":          toExportTriggerObject(o.CustomerID, isavro, false, "string"),
		"id":                   toExportTriggerObject(o.ID, isavro, false, "string"),
		"ref_id":               toExportTriggerObject(o.RefID, isavro, false, "string"),
		"ref_type":             toExportTriggerObject(o.RefType, isavro, false, "string"),
		"reprocess_historical": toExportTriggerObject(o.ReprocessHistorical, isavro, false, "boolean"),
		"updated_ts":           toExportTriggerObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":                 toExportTriggerObject(o.UUID, isavro, true, "string"),
		"hashcode":             toExportTriggerObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportTrigger) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = number.ToBoolAny(nil)
			} else {
				o.ReprocessHistorical = number.ToBoolAny(val)
			}
		}
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

	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ExportTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetExportTriggerAvroSchemaSpec creates the avro schema specification for ExportTrigger
func GetExportTriggerAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "ExportTrigger",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
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
				"name": "reprocess_historical",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "uuid",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *ExportTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetExportTriggerAvroSchema creates the avro schema for ExportTrigger
func GetExportTriggerAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetExportTriggerAvroSchemaSpec())
}

// ExportTriggerSendEvent is an event detail for sending data
type ExportTriggerSendEvent struct {
	ExportTrigger *ExportTrigger
	headers       map[string]string
	time          time.Time
	key           string
}

var _ datamodel.ModelSendEvent = (*ExportTriggerSendEvent)(nil)

// Key is the key to use for the message
func (e *ExportTriggerSendEvent) Key() string {
	if e.key == "" {
		return e.ExportTrigger.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ExportTriggerSendEvent) Object() datamodel.Model {
	return e.ExportTrigger
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ExportTriggerSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ExportTriggerSendEvent) Timestamp() time.Time {
	return e.time
}

// ExportTriggerSendEventOpts is a function handler for setting opts
type ExportTriggerSendEventOpts func(o *ExportTriggerSendEvent)

// WithExportTriggerSendEventKey sets the key value to a value different than the object ID
func WithExportTriggerSendEventKey(key string) ExportTriggerSendEventOpts {
	return func(o *ExportTriggerSendEvent) {
		o.key = key
	}
}

// WithExportTriggerSendEventTimestamp sets the timestamp value
func WithExportTriggerSendEventTimestamp(tv time.Time) ExportTriggerSendEventOpts {
	return func(o *ExportTriggerSendEvent) {
		o.time = tv
	}
}

// WithExportTriggerSendEventHeader sets the timestamp value
func WithExportTriggerSendEventHeader(key, value string) ExportTriggerSendEventOpts {
	return func(o *ExportTriggerSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewExportTriggerSendEvent returns a new ExportTriggerSendEvent instance
func NewExportTriggerSendEvent(o *ExportTrigger, opts ...ExportTriggerSendEventOpts) *ExportTriggerSendEvent {
	res := &ExportTriggerSendEvent{
		ExportTrigger: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewExportTriggerProducer will stream data from the channel
func NewExportTriggerProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*ExportTrigger); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.ExportTrigger but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewExportTriggerConsumer will stream data from the topic into the provided channel
func NewExportTriggerConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object ExportTrigger
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.ExportTrigger: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.ExportTrigger: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.ExportTrigger")
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

			ch <- &ExportTriggerReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object ExportTrigger
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ExportTriggerReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ExportTriggerReceiveEvent is an event detail for receiving data
type ExportTriggerReceiveEvent struct {
	ExportTrigger *ExportTrigger
	message       eventing.Message
	eof           bool
}

var _ datamodel.ModelReceiveEvent = (*ExportTriggerReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ExportTriggerReceiveEvent) Object() datamodel.Model {
	return e.ExportTrigger
}

// Message returns the underlying message data for the event
func (e *ExportTriggerReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ExportTriggerReceiveEvent) EOF() bool {
	return e.eof
}

// ExportTriggerProducer implements the datamodel.ModelEventProducer
type ExportTriggerProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ExportTriggerProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ExportTriggerProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ExportTriggerProducer) Close() error {
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
func (o *ExportTrigger) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *ExportTrigger) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ExportTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewExportTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// NewExportTriggerProducerChannel returns a channel which can be used for producing Model events
func NewExportTriggerProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewExportTriggerProducerChannelSize(producer, 0, errors)
}

// NewExportTriggerProducerChannelSize returns a channel which can be used for producing Model events
func NewExportTriggerProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ExportTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewExportTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// ExportTriggerConsumer implements the datamodel.ModelEventConsumer
type ExportTriggerConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ExportTriggerConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ExportTriggerConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ExportTriggerConsumer) Close() error {
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
func (o *ExportTrigger) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ExportTriggerConsumer{
		ch:       ch,
		callback: NewExportTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewExportTriggerConsumerChannel returns a consumer channel which can be used to consume Model events
func NewExportTriggerConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ExportTriggerConsumer{
		ch:       ch,
		callback: NewExportTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
