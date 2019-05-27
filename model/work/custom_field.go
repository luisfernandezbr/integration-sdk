// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
)

// CustomFieldTopic is the default topic name
const CustomFieldTopic datamodel.TopicNameType = "work_CustomField_topic"

// CustomFieldStream is the default stream name
const CustomFieldStream datamodel.TopicNameType = "work_CustomField_stream"

// CustomFieldTable is the default table name
const CustomFieldTable datamodel.TopicNameType = "work_CustomField"

// CustomFieldModelName is the model name
const CustomFieldModelName datamodel.ModelNameType = "work.CustomField"

// CustomField user defined fields
type CustomField struct {
	// built in types

	ID         string `json:"custom_field_id" yaml:"custom_field_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name the name of the field
	Name string `json:"name" yaml:"name" faker:"-"`
	// Key key of the field
	Key string `json:"key" yaml:"key" faker:"-"`
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
	if o == nil {
		return toCustomFieldObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCustomFieldObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCustomFieldObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *CustomField:
		return v.ToMap()
	case CustomField:
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
			arr = append(arr, toCustomFieldObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
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

func (o *CustomField) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CustomField) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CustomField", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
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

// GetTopicConfig returns the topic config object
func (o *CustomField) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		NumPartitions:     4,
		ReplicationFactor: 1,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// Clone returns an exact copy of CustomField
func (o *CustomField) Clone() *CustomField {
	c := new(CustomField)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CustomField) Anon() *CustomField {
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
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCustomField *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *CustomField) GetAvroCodec() *goavro.Codec {
	if cachedCodecCustomField == nil {
		c, err := GetCustomFieldAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCustomField = c
	}
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

// Stringify returns the object in JSON format as a string
func (o *CustomField) Stringify() string {
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
	return map[string]interface{}{
		"custom_field_id": o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"name":            toCustomFieldObject(o.Name, isavro, false, "string"),
		"key":             toCustomFieldObject(o.Key, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CustomField) FromMap(kv map[string]interface{}) {
	if val, ok := kv["custom_field_id"].(string); ok {
		o.ID = val
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	}
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["key"].(string); ok {
		o.Key = val
	} else {
		val := kv["key"]
		if val == nil {
			o.Key = ""
		} else {
			o.Key = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CustomField) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Key)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCustomFieldAvroSchemaSpec creates the avro schema specification for CustomField
func GetCustomFieldAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "CustomField",
		"connect.name": "work.CustomField",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "custom_field_id",
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "key",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetCustomFieldAvroSchema creates the avro schema for CustomField
func GetCustomFieldAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCustomFieldAvroSchemaSpec())
}

// TransformCustomFieldFunc is a function for transforming CustomField during processing
type TransformCustomFieldFunc func(input *CustomField) (*CustomField, error)

// NewCustomFieldPipe creates a pipe for processing CustomField items
func NewCustomFieldPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCustomFieldFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCustomFieldInputStream(input, errors)
	var stream chan CustomField
	if len(transforms) > 0 {
		stream = make(chan CustomField, 1000)
	} else {
		stream = inch
	}
	outdone := NewCustomFieldOutputStream(output, stream, errors)
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

// NewCustomFieldInputStreamDir creates a channel for reading CustomField as JSON newlines from a directory of files
func NewCustomFieldInputStreamDir(dir string, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/custom_field\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for custom_field")
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCustomFieldInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CustomField)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCustomFieldInputStreamFile creates an channel for reading CustomField as JSON newlines from filename
func NewCustomFieldInputStreamFile(filename string, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CustomField)
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
			ch := make(chan CustomField)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCustomFieldInputStream(f, errors, transforms...)
}

// NewCustomFieldInputStream creates an channel for reading CustomField as JSON newlines from stream
func NewCustomFieldInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCustomFieldFunc) (chan CustomField, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CustomField, 1000)
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
			var item CustomField
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

// NewCustomFieldOutputStreamDir will output json newlines from channel and save in dir
func NewCustomFieldOutputStreamDir(dir string, ch chan CustomField, errors chan<- error, transforms ...TransformCustomFieldFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/custom_field\\.json(\\.gz)?$")
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
	return NewCustomFieldOutputStream(gz, ch, errors, transforms...)
}

// NewCustomFieldOutputStream will output json newlines from channel to the stream
func NewCustomFieldOutputStream(stream io.WriteCloser, ch chan CustomField, errors chan<- error, transforms ...TransformCustomFieldFunc) <-chan bool {
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
func NewCustomFieldProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*CustomField); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
				}
				for k, v := range item.Headers() {
					headers[k] = v
				}
				msg := event.Message{
					Key:     item.Key(),
					Value:   binary,
					Codec:   codec,
					Headers: headers,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type work.CustomField but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewCustomFieldConsumer will stream data from the topic into the provided channel
func NewCustomFieldConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object CustomField
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into work.CustomField: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CustomFieldReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// CustomFieldReceiveEvent is an event detail for receiving data
type CustomFieldReceiveEvent struct {
	CustomField *CustomField
	message     event.Message
}

var _ datamodel.ModelReceiveEvent = (*CustomFieldReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CustomFieldReceiveEvent) Object() datamodel.Model {
	return e.CustomField
}

// Message returns the underlying message data for the event
func (e *CustomFieldReceiveEvent) Message() event.Message {
	return e.message
}

// CustomFieldProducer implements the datamodel.ModelEventProducer
type CustomFieldProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*CustomFieldProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CustomFieldProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CustomFieldProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *CustomField) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CustomFieldProducer{
		ch:   ch,
		done: NewCustomFieldProducer(producer, ch, errors),
	}
}

// NewCustomFieldProducerChannel returns a channel which can be used for producing Model events
func NewCustomFieldProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CustomFieldProducer{
		ch:   ch,
		done: NewCustomFieldProducer(producer, ch, errors),
	}
}

// CustomFieldConsumer implements the datamodel.ModelEventConsumer
type CustomFieldConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*CustomFieldConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CustomFieldConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CustomFieldConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *CustomField) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCustomFieldConsumer(consumer, ch, errors)
	return &CustomFieldConsumer{
		ch: ch,
	}
}

// NewCustomFieldConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCustomFieldConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCustomFieldConsumer(consumer, ch, errors)
	return &CustomFieldConsumer{
		ch: ch,
	}
}
