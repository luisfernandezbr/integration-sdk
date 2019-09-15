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
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "work_User_topic"

	// UserStream is the default stream name
	UserStream datamodel.TopicNameType = "work_User_stream"

	// UserTable is the default table name
	UserTable datamodel.TopicNameType = "work_user"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "work.User"
)

const (
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "avatar_url"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "customer_id"
	// UserEmailColumn is the email column name
	UserEmailColumn = "email"
	// UserIDColumn is the id column name
	UserIDColumn = "id"
	// UserNameColumn is the name column name
	UserNameColumn = "name"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "ref_id"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "ref_type"
	// UserUpdatedAtColumn is the updated_ts column name
	UserUpdatedAtColumn = "updated_ts"
	// UserUsernameColumn is the username column name
	UserUsernameColumn = "username"
)

// User the work user
type User struct {
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-" badgerhold:"index"`
	// Email the email for the user
	Email *string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Name the name of the user
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Username the username of the user
	Username string `json:"username" bson:"username" yaml:"username" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*User)(nil)

func toUserObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *User:
		return v.ToMap(isavro)

	default:
		return o
	}
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("work.User<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *User) GetTopicName() datamodel.TopicNameType {
	return UserTopic
}

// GetModelName returns the name of the model
func (o *User) GetModelName() datamodel.ModelNameType {
	return UserModelName
}

// GetStreamName returns the name of the stream
func (o *User) GetStreamName() string {
	return UserStream.String()
}

// GetTableName returns the name of the table
func (o *User) GetTableName() string {
	return UserTable.String()
}

// NewUserID provides a template for generating an ID field for User
func NewUserID(customerID string, refType string, refID string) string {
	return hash.Values("User", customerID, refType, refID)
}

func (o *User) setDefaults(frommap bool) {
	if o.AvatarURL == nil {
		o.AvatarURL = &emptyString
	}
	if o.Email == nil {
		o.Email = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("User", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *User) GetTopicKey() string {
	var i interface{} = o.CustomerID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *User) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for User")
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *User) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *User) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *User) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *User) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *User) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *User) GetStateKey() string {
	key := "customer_id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *User) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of User
func (o *User) Clone() datamodel.Model {
	c := new(User)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *User) Anon() datamodel.Model {
	c := new(User)
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
func (o *User) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *User) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *User) UnmarshalJSON(data []byte) error {
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

var cachedCodecUser *goavro.Codec
var cachedCodecUserLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *User) GetAvroCodec() *goavro.Codec {
	cachedCodecUserLock.Lock()
	if cachedCodecUser == nil {
		c, err := GetUserAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUser = c
	}
	cachedCodecUserLock.Unlock()
	return cachedCodecUser
}

// ToAvroBinary returns the data as Avro binary data
func (o *User) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *User) FromAvroBinary(value []byte) error {
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
func (o *User) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two User objects are equal
func (o *User) IsEqual(other *User) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *User) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"avatar_url":  toUserObject(o.AvatarURL, isavro, true, "string"),
		"customer_id": toUserObject(o.CustomerID, isavro, false, "string"),
		"email":       toUserObject(o.Email, isavro, true, "string"),
		"id":          toUserObject(o.ID, isavro, false, "string"),
		"name":        toUserObject(o.Name, isavro, false, "string"),
		"ref_id":      toUserObject(o.RefID, isavro, false, "string"),
		"ref_type":    toUserObject(o.RefType, isavro, false, "string"),
		"updated_ts":  toUserObject(o.UpdatedAt, isavro, false, "long"),
		"username":    toUserObject(o.Username, isavro, false, "string"),
		"hashcode":    toUserObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["avatar_url"].(*string); ok {
		o.AvatarURL = val
	} else if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = &val
	} else {
		if val, ok := kv["avatar_url"]; ok {
			if val == nil {
				o.AvatarURL = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["email"].(*string); ok {
		o.Email = val
	} else if val, ok := kv["email"].(string); ok {
		o.Email = &val
	} else {
		if val, ok := kv["email"]; ok {
			if val == nil {
				o.Email = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Email = pstrings.Pointer(fmt.Sprintf("%v", val))
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

	if val, ok := kv["username"].(string); ok {
		o.Username = val
	} else {
		if val, ok := kv["username"]; ok {
			if val == nil {
				o.Username = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Username = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AvatarURL)
	args = append(args, o.CustomerID)
	args = append(args, o.Email)
	args = append(args, o.ID)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Username)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUserAvroSchemaSpec creates the avro schema specification for User
func GetUserAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "work",
		"name":      "User",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "email",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "id",
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
			map[string]interface{}{
				"name": "username",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *User) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetUserAvroSchema creates the avro schema for User
func GetUserAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserAvroSchemaSpec())
}

// UserSendEvent is an event detail for sending data
type UserSendEvent struct {
	User    *User
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*UserSendEvent)(nil)

// Key is the key to use for the message
func (e *UserSendEvent) Key() string {
	if e.key == "" {
		return e.User.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserSendEvent) Object() datamodel.Model {
	return e.User
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserSendEvent) Timestamp() time.Time {
	return e.time
}

// UserSendEventOpts is a function handler for setting opts
type UserSendEventOpts func(o *UserSendEvent)

// WithUserSendEventKey sets the key value to a value different than the object ID
func WithUserSendEventKey(key string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.key = key
	}
}

// WithUserSendEventTimestamp sets the timestamp value
func WithUserSendEventTimestamp(tv time.Time) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.time = tv
	}
}

// WithUserSendEventHeader sets the timestamp value
func WithUserSendEventHeader(key, value string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserSendEvent returns a new UserSendEvent instance
func NewUserSendEvent(o *User, opts ...UserSendEventOpts) *UserSendEvent {
	res := &UserSendEvent{
		User: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserProducer will stream data from the channel
func NewUserProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*User); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type work.User but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUserConsumer will stream data from the topic into the provided channel
func NewUserConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object User
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into work.User: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into work.User: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for work.User")
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

			ch <- &UserReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object User
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UserReceiveEvent is an event detail for receiving data
type UserReceiveEvent struct {
	User    *User
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*UserReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserReceiveEvent) Object() datamodel.Model {
	return e.User
}

// Message returns the underlying message data for the event
func (e *UserReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserReceiveEvent) EOF() bool {
	return e.eof
}

// UserProducer implements the datamodel.ModelEventProducer
type UserProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UserProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserProducer) Close() error {
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
func (o *User) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *User) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUserProducerChannel returns a channel which can be used for producing Model events
func NewUserProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUserProducerChannelSize(producer, 0, errors)
}

// NewUserProducerChannelSize returns a channel which can be used for producing Model events
func NewUserProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserProducer(newctx, producer, ch, errors, empty),
	}
}

// UserConsumer implements the datamodel.ModelEventConsumer
type UserConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UserConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserConsumer) Close() error {
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
func (o *User) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserConsumer{
		ch:       ch,
		callback: NewUserConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUserConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserConsumer{
		ch:       ch,
		callback: NewUserConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
