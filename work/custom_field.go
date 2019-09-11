// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	// CustomFieldTopic is the default topic name
	CustomFieldTopic datamodel.TopicNameType = "work_CustomField_topic"

	// CustomFieldStream is the default stream name
	CustomFieldStream datamodel.TopicNameType = "work_CustomField_stream"

	// CustomFieldTable is the default table name
	CustomFieldTable datamodel.TopicNameType = "work_customfield"

	// CustomFieldModelName is the model name
	CustomFieldModelName datamodel.ModelNameType = "work.CustomField"
)

const (
	// CustomFieldCustomerIDColumn is the customer_id column name
	CustomFieldCustomerIDColumn = "customer_id"
	// CustomFieldIDColumn is the id column name
	CustomFieldIDColumn = "id"
	// CustomFieldKeyColumn is the key column name
	CustomFieldKeyColumn = "key"
	// CustomFieldNameColumn is the name column name
	CustomFieldNameColumn = "name"
	// CustomFieldRefIDColumn is the ref_id column name
	CustomFieldRefIDColumn = "ref_id"
	// CustomFieldRefTypeColumn is the ref_type column name
	CustomFieldRefTypeColumn = "ref_type"
	// CustomFieldUpdatedAtColumn is the updated_ts column name
	CustomFieldUpdatedAtColumn = "updated_ts"
)

// CustomField user defined fields
type CustomField struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Key key of the field
	Key string `json:"key" bson:"key" yaml:"key" faker:"-"`
	// Name the name of the field
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
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
var _ datamodel.Model = (*CustomField)(nil)

func toCustomFieldObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCustomFieldObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CustomField:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of CustomField
func (o *CustomField) String() string {
	return fmt.Sprintf("work.CustomField<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CustomField) GetTopicName() datamodel.TopicNameType {
	return CustomFieldTopic
}

// GetModelName returns the name of the model
func (o *CustomField) GetModelName() datamodel.ModelNameType {
	return CustomFieldModelName
}

// GetStreamName returns the name of the stream
func (o *CustomField) GetStreamName() string {
	return CustomFieldStream.String()
}

// GetTableName returns the name of the table
func (o *CustomField) GetTableName() string {
	return CustomFieldTable.String()
}

// NewCustomFieldID provides a template for generating an ID field for CustomField
func NewCustomFieldID(customerID string, refType string, refID string) string {
	return hash.Values("CustomField", customerID, refType, refID)
}

func (o *CustomField) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CustomField", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CustomField) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CustomField) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CustomField) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for CustomField")
}

// GetRefID returns the RefID for the object
func (o *CustomField) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CustomField) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CustomField) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CustomField) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CustomField) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CustomFieldModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CustomField) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "customer_id",
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
func (o *CustomField) GetStateKey() string {
	key := "customer_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *CustomField) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CustomField
func (o *CustomField) Clone() datamodel.Model {
	c := new(CustomField)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CustomField) Anon() datamodel.Model {
	c := new(CustomField)
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
func (o *CustomField) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *CustomField) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CustomField) UnmarshalJSON(data []byte) error {
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

var cachedCodecCustomField *goavro.Codec
var cachedCodecCustomFieldLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *CustomField) GetAvroCodec() *goavro.Codec {
	cachedCodecCustomFieldLock.Lock()
	if cachedCodecCustomField == nil {
		c, err := GetCustomFieldAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCustomField = c
	}
	cachedCodecCustomFieldLock.Unlock()
	return cachedCodecCustomField
}

// ToAvroBinary returns the data as Avro binary data
func (o *CustomField) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CustomField) FromAvroBinary(value []byte) error {
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
func (o *CustomField) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CustomField objects are equal
func (o *CustomField) IsEqual(other *CustomField) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CustomField) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id": toCustomFieldObject(o.CustomerID, isavro, false, "string"),
		"id":          toCustomFieldObject(o.ID, isavro, false, "string"),
		"key":         toCustomFieldObject(o.Key, isavro, false, "string"),
		"name":        toCustomFieldObject(o.Name, isavro, false, "string"),
		"ref_id":      toCustomFieldObject(o.RefID, isavro, false, "string"),
		"ref_type":    toCustomFieldObject(o.RefType, isavro, false, "string"),
		"updated_ts":  toCustomFieldObject(o.UpdatedAt, isavro, false, "long"),
		"hashcode":    toCustomFieldObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CustomField) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["key"].(string); ok {
		o.Key = val
	} else {
		if val, ok := kv["key"]; ok {
			if val == nil {
				o.Key = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Key = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
func (o *CustomField) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Key)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCustomFieldAvroSchemaSpec creates the avro schema specification for CustomField
func GetCustomFieldAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "work",
		"name":      "CustomField",
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
				"name": "key",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
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
func (o *CustomField) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetCustomFieldAvroSchema creates the avro schema for CustomField
func GetCustomFieldAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCustomFieldAvroSchemaSpec())
}

// CustomFieldSendEvent is an event detail for sending data
type CustomFieldSendEvent struct {
	CustomField *CustomField
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*CustomFieldSendEvent)(nil)

// Key is the key to use for the message
func (e *CustomFieldSendEvent) Key() string {
	if e.key == "" {
		return e.CustomField.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CustomFieldSendEvent) Object() datamodel.Model {
	return e.CustomField
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CustomFieldSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CustomFieldSendEvent) Timestamp() time.Time {
	return e.time
}

// CustomFieldSendEventOpts is a function handler for setting opts
type CustomFieldSendEventOpts func(o *CustomFieldSendEvent)

// WithCustomFieldSendEventKey sets the key value to a value different than the object ID
func WithCustomFieldSendEventKey(key string) CustomFieldSendEventOpts {
	return func(o *CustomFieldSendEvent) {
		o.key = key
	}
}

// WithCustomFieldSendEventTimestamp sets the timestamp value
func WithCustomFieldSendEventTimestamp(tv time.Time) CustomFieldSendEventOpts {
	return func(o *CustomFieldSendEvent) {
		o.time = tv
	}
}

// WithCustomFieldSendEventHeader sets the timestamp value
func WithCustomFieldSendEventHeader(key, value string) CustomFieldSendEventOpts {
	return func(o *CustomFieldSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCustomFieldSendEvent returns a new CustomFieldSendEvent instance
func NewCustomFieldSendEvent(o *CustomField, opts ...CustomFieldSendEventOpts) *CustomFieldSendEvent {
	res := &CustomFieldSendEvent{
		CustomField: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCustomFieldProducer will stream data from the channel
func NewCustomFieldProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*CustomField); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type work.CustomField but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewCustomFieldConsumer will stream data from the topic into the provided channel
func NewCustomFieldConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object CustomField
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.CustomField: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into work.CustomField: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for work.CustomField")
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

			ch <- &CustomFieldReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object CustomField
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CustomFieldReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// CustomFieldReceiveEvent is an event detail for receiving data
type CustomFieldReceiveEvent struct {
	CustomField *CustomField
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*CustomFieldReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CustomFieldReceiveEvent) Object() datamodel.Model {
	return e.CustomField
}

// Message returns the underlying message data for the event
func (e *CustomFieldReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CustomFieldReceiveEvent) EOF() bool {
	return e.eof
}

// CustomFieldProducer implements the datamodel.ModelEventProducer
type CustomFieldProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*CustomFieldProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CustomFieldProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CustomFieldProducer) Close() error {
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
func (o *CustomField) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *CustomField) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CustomFieldProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCustomFieldProducer(newctx, producer, ch, errors, empty),
	}
}

// NewCustomFieldProducerChannel returns a channel which can be used for producing Model events
func NewCustomFieldProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewCustomFieldProducerChannelSize(producer, 0, errors)
}

// NewCustomFieldProducerChannelSize returns a channel which can be used for producing Model events
func NewCustomFieldProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CustomFieldProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCustomFieldProducer(newctx, producer, ch, errors, empty),
	}
}

// CustomFieldConsumer implements the datamodel.ModelEventConsumer
type CustomFieldConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*CustomFieldConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CustomFieldConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CustomFieldConsumer) Close() error {
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
func (o *CustomField) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CustomFieldConsumer{
		ch:       ch,
		callback: NewCustomFieldConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewCustomFieldConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCustomFieldConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CustomFieldConsumer{
		ch:       ch,
		callback: NewCustomFieldConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
