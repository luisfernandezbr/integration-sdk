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
	// UserTriggerTopic is the default topic name
	UserTriggerTopic datamodel.TopicNameType = "agent_UserTrigger_topic"

	// UserTriggerStream is the default stream name
	UserTriggerStream datamodel.TopicNameType = "agent_UserTrigger_stream"

	// UserTriggerTable is the default table name
	UserTriggerTable datamodel.TopicNameType = "agent_usertrigger"

	// UserTriggerModelName is the model name
	UserTriggerModelName datamodel.ModelNameType = "agent.UserTrigger"
)

const (
	// UserTriggerCustomerIDColumn is the customer_id column name
	UserTriggerCustomerIDColumn = "customer_id"
	// UserTriggerIDColumn is the id column name
	UserTriggerIDColumn = "id"
	// UserTriggerIntegrationIDColumn is the integration_id column name
	UserTriggerIntegrationIDColumn = "integration_id"
	// UserTriggerRefIDColumn is the ref_id column name
	UserTriggerRefIDColumn = "ref_id"
	// UserTriggerRefTypeColumn is the ref_type column name
	UserTriggerRefTypeColumn = "ref_type"
	// UserTriggerUpdatedAtColumn is the updated_ts column name
	UserTriggerUpdatedAtColumn = "updated_ts"
)

// UserTrigger used to trigger an agent.UserRequest
type UserTrigger struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UserTrigger)(nil)

func toUserTriggerObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserTriggerObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserTrigger:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of UserTrigger
func (o *UserTrigger) String() string {
	return fmt.Sprintf("agent.UserTrigger<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UserTrigger) GetTopicName() datamodel.TopicNameType {
	return UserTriggerTopic
}

// GetModelName returns the name of the model
func (o *UserTrigger) GetModelName() datamodel.ModelNameType {
	return UserTriggerModelName
}

// GetStreamName returns the name of the stream
func (o *UserTrigger) GetStreamName() string {
	return UserTriggerStream.String()
}

// GetTableName returns the name of the table
func (o *UserTrigger) GetTableName() string {
	return UserTriggerTable.String()
}

// NewUserTriggerID provides a template for generating an ID field for UserTrigger
func NewUserTriggerID(customerID string, refType string, refID string) string {
	return hash.Values("UserTrigger", customerID, refType, refID)
}

func (o *UserTrigger) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserTrigger) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UserTrigger) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UserTrigger) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for UserTrigger")
}

// GetRefID returns the RefID for the object
func (o *UserTrigger) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UserTrigger) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UserTrigger) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UserTrigger) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *UserTrigger) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserTriggerModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UserTrigger) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *UserTrigger) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *UserTrigger) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of UserTrigger
func (o *UserTrigger) Clone() datamodel.Model {
	c := new(UserTrigger)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UserTrigger) Anon() datamodel.Model {
	c := new(UserTrigger)
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
func (o *UserTrigger) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *UserTrigger) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UserTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserTrigger) UnmarshalJSON(data []byte) error {
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

var cachedCodecUserTrigger *goavro.Codec
var cachedCodecUserTriggerLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *UserTrigger) GetAvroCodec() *goavro.Codec {
	cachedCodecUserTriggerLock.Lock()
	if cachedCodecUserTrigger == nil {
		c, err := GetUserTriggerAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUserTrigger = c
	}
	cachedCodecUserTriggerLock.Unlock()
	return cachedCodecUserTrigger
}

// ToAvroBinary returns the data as Avro binary data
func (o *UserTrigger) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *UserTrigger) FromAvroBinary(value []byte) error {
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
func (o *UserTrigger) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserTrigger objects are equal
func (o *UserTrigger) IsEqual(other *UserTrigger) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserTrigger) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":    toUserTriggerObject(o.CustomerID, isavro, false, "string"),
		"id":             toUserTriggerObject(o.ID, isavro, false, "string"),
		"integration_id": toUserTriggerObject(o.IntegrationID, isavro, false, "string"),
		"ref_id":         toUserTriggerObject(o.RefID, isavro, false, "string"),
		"ref_type":       toUserTriggerObject(o.RefType, isavro, false, "string"),
		"updated_ts":     toUserTriggerObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":       toUserTriggerObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserTrigger) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *UserTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUserTriggerAvroSchemaSpec creates the avro schema specification for UserTrigger
func GetUserTriggerAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "UserTrigger",
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
				"name": "integration_id",
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
				"name": "updated_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *UserTrigger) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetUserTriggerAvroSchema creates the avro schema for UserTrigger
func GetUserTriggerAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserTriggerAvroSchemaSpec())
}

// UserTriggerSendEvent is an event detail for sending data
type UserTriggerSendEvent struct {
	UserTrigger *UserTrigger
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*UserTriggerSendEvent)(nil)

// Key is the key to use for the message
func (e *UserTriggerSendEvent) Key() string {
	if e.key == "" {
		return e.UserTrigger.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserTriggerSendEvent) Object() datamodel.Model {
	return e.UserTrigger
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserTriggerSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserTriggerSendEvent) Timestamp() time.Time {
	return e.time
}

// UserTriggerSendEventOpts is a function handler for setting opts
type UserTriggerSendEventOpts func(o *UserTriggerSendEvent)

// WithUserTriggerSendEventKey sets the key value to a value different than the object ID
func WithUserTriggerSendEventKey(key string) UserTriggerSendEventOpts {
	return func(o *UserTriggerSendEvent) {
		o.key = key
	}
}

// WithUserTriggerSendEventTimestamp sets the timestamp value
func WithUserTriggerSendEventTimestamp(tv time.Time) UserTriggerSendEventOpts {
	return func(o *UserTriggerSendEvent) {
		o.time = tv
	}
}

// WithUserTriggerSendEventHeader sets the timestamp value
func WithUserTriggerSendEventHeader(key, value string) UserTriggerSendEventOpts {
	return func(o *UserTriggerSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserTriggerSendEvent returns a new UserTriggerSendEvent instance
func NewUserTriggerSendEvent(o *UserTrigger, opts ...UserTriggerSendEventOpts) *UserTriggerSendEvent {
	res := &UserTriggerSendEvent{
		UserTrigger: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserTriggerProducer will stream data from the channel
func NewUserTriggerProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*UserTrigger); ok {
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
						Key:       object.GetID(),
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.UserTrigger but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUserTriggerConsumer will stream data from the topic into the provided channel
func NewUserTriggerConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object UserTrigger
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.UserTrigger: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.UserTrigger: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.UserTrigger")
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

			ch <- &UserTriggerReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object UserTrigger
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserTriggerReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UserTriggerReceiveEvent is an event detail for receiving data
type UserTriggerReceiveEvent struct {
	UserTrigger *UserTrigger
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*UserTriggerReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserTriggerReceiveEvent) Object() datamodel.Model {
	return e.UserTrigger
}

// Message returns the underlying message data for the event
func (e *UserTriggerReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserTriggerReceiveEvent) EOF() bool {
	return e.eof
}

// UserTriggerProducer implements the datamodel.ModelEventProducer
type UserTriggerProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UserTriggerProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserTriggerProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserTriggerProducer) Close() error {
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
func (o *UserTrigger) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *UserTrigger) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUserTriggerProducerChannel returns a channel which can be used for producing Model events
func NewUserTriggerProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUserTriggerProducerChannelSize(producer, 0, errors)
}

// NewUserTriggerProducerChannelSize returns a channel which can be used for producing Model events
func NewUserTriggerProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserTriggerProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserTriggerProducer(newctx, producer, ch, errors, empty),
	}
}

// UserTriggerConsumer implements the datamodel.ModelEventConsumer
type UserTriggerConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UserTriggerConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserTriggerConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserTriggerConsumer) Close() error {
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
func (o *UserTrigger) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserTriggerConsumer{
		ch:       ch,
		callback: NewUserTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUserTriggerConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserTriggerConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserTriggerConsumer{
		ch:       ch,
		callback: NewUserTriggerConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
