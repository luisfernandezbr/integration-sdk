// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

// MetricTopic is the default topic name
const MetricTopic datamodel.TopicNameType = "codequality_Metric_topic"

// MetricStream is the default stream name
const MetricStream datamodel.TopicNameType = "codequality_Metric_stream"

// MetricTable is the default table name
const MetricTable datamodel.TopicNameType = "codequality_Metric"

// MetricModelName is the model name
const MetricModelName datamodel.ModelNameType = "codequality.Metric"

// Metric individual metric details
type Metric struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// DateAt the date when the metric was created
	DateAt int64 `json:"date_ts" bson:"date_ts" yaml:"date_ts" faker:"-"`
	// ProjectID the the project id
	ProjectID string `json:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// Name the metric name
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Value the value of the metric
	Value string `json:"value" bson:"value" yaml:"value" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Metric)(nil)

func toMetricObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toMetricObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toMetricObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toMetricObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toMetricObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Metric:
		return v.ToMap()
	case Metric:
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
			arr = append(arr, toMetricObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Metric
func (o *Metric) String() string {
	return fmt.Sprintf("codequality.Metric<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Metric) GetTopicName() datamodel.TopicNameType {
	return MetricTopic
}

// GetModelName returns the name of the model
func (o *Metric) GetModelName() datamodel.ModelNameType {
	return MetricModelName
}

func (o *Metric) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Metric) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Metric", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Metric) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Metric) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Metric")
}

// GetRefID returns the RefID for the object
func (o *Metric) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Metric) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Metric) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Metric) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Metric) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MetricModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Metric) GetTopicConfig() *datamodel.ModelTopicConfig {
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

// GetCustomerID will return the customer_id
func (o *Metric) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of Metric
func (o *Metric) Clone() datamodel.Model {
	c := new(Metric)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Metric) Anon() datamodel.Model {
	c := new(Metric)
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
func (o *Metric) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Metric) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecMetric *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Metric) GetAvroCodec() *goavro.Codec {
	if cachedCodecMetric == nil {
		c, err := GetMetricAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecMetric = c
	}
	return cachedCodecMetric
}

// ToAvroBinary returns the data as Avro binary data
func (o *Metric) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Metric) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Metric objects are equal
func (o *Metric) IsEqual(other *Metric) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Metric) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"date_ts":     toMetricObject(o.DateAt, isavro, false, "long"),
		"project_id":  toMetricObject(o.ProjectID, isavro, false, "string"),
		"name":        toMetricObject(o.Name, isavro, false, "string"),
		"value":       toMetricObject(o.Value, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Metric) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else if val, ok := kv["_id"].(string); ok {
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
	if val, ok := kv["date_ts"].(int64); ok {
		o.DateAt = val
	} else {
		val := kv["date_ts"]
		if val == nil {
			o.DateAt = number.ToInt64Any(nil)
		} else {
			o.DateAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		val := kv["project_id"]
		if val == nil {
			o.ProjectID = ""
		} else {
			o.ProjectID = fmt.Sprintf("%v", val)
		}
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
	if val, ok := kv["value"].(string); ok {
		o.Value = val
	} else {
		val := kv["value"]
		if val == nil {
			o.Value = ""
		} else {
			o.Value = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Metric) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.DateAt)
	args = append(args, o.ProjectID)
	args = append(args, o.Name)
	args = append(args, o.Value)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetMetricAvroSchemaSpec creates the avro schema specification for Metric
func GetMetricAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "codequality",
		"name":         "Metric",
		"connect.name": "codequality.Metric",
		"fields": []map[string]interface{}{
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
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "date_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "project_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "value",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetMetricAvroSchema creates the avro schema for Metric
func GetMetricAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetMetricAvroSchemaSpec())
}

// TransformMetricFunc is a function for transforming Metric during processing
type TransformMetricFunc func(input *Metric) (*Metric, error)

// NewMetricPipe creates a pipe for processing Metric items
func NewMetricPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformMetricFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewMetricInputStream(input, errors)
	var stream chan Metric
	if len(transforms) > 0 {
		stream = make(chan Metric, 1000)
	} else {
		stream = inch
	}
	outdone := NewMetricOutputStream(output, stream, errors)
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

// NewMetricInputStreamDir creates a channel for reading Metric as JSON newlines from a directory of files
func NewMetricInputStreamDir(dir string, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/codequality/metric\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for metric")
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewMetricInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewMetricInputStreamFile creates an channel for reading Metric as JSON newlines from filename
func NewMetricInputStreamFile(filename string, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Metric)
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
			ch := make(chan Metric)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewMetricInputStream(f, errors, transforms...)
}

// NewMetricInputStream creates an channel for reading Metric as JSON newlines from stream
func NewMetricInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Metric, 1000)
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
			var item Metric
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

// NewMetricOutputStreamDir will output json newlines from channel and save in dir
func NewMetricOutputStreamDir(dir string, ch chan Metric, errors chan<- error, transforms ...TransformMetricFunc) <-chan bool {
	fp := filepath.Join(dir, "/codequality/metric\\.json(\\.gz)?$")
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
	return NewMetricOutputStream(gz, ch, errors, transforms...)
}

// NewMetricOutputStream will output json newlines from channel to the stream
func NewMetricOutputStream(stream io.WriteCloser, ch chan Metric, errors chan<- error, transforms ...TransformMetricFunc) <-chan bool {
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

// MetricSendEvent is an event detail for sending data
type MetricSendEvent struct {
	Metric  *Metric
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*MetricSendEvent)(nil)

// Key is the key to use for the message
func (e *MetricSendEvent) Key() string {
	if e.key == "" {
		return e.Metric.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *MetricSendEvent) Object() datamodel.Model {
	return e.Metric
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *MetricSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *MetricSendEvent) Timestamp() time.Time {
	return e.time
}

// MetricSendEventOpts is a function handler for setting opts
type MetricSendEventOpts func(o *MetricSendEvent)

// WithMetricSendEventKey sets the key value to a value different than the object ID
func WithMetricSendEventKey(key string) MetricSendEventOpts {
	return func(o *MetricSendEvent) {
		o.key = key
	}
}

// WithMetricSendEventTimestamp sets the timestamp value
func WithMetricSendEventTimestamp(tv time.Time) MetricSendEventOpts {
	return func(o *MetricSendEvent) {
		o.time = tv
	}
}

// WithMetricSendEventHeader sets the timestamp value
func WithMetricSendEventHeader(key, value string) MetricSendEventOpts {
	return func(o *MetricSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewMetricSendEvent returns a new MetricSendEvent instance
func NewMetricSendEvent(o *Metric, opts ...MetricSendEventOpts) *MetricSendEvent {
	res := &MetricSendEvent{
		Metric: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewMetricProducer will stream data from the channel
func NewMetricProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Metric); ok {
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
				msg := event.Message{
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type codequality.Metric but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewMetricConsumer will stream data from the topic into the provided channel
func NewMetricConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Metric
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into codequality.Metric: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &MetricReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// MetricReceiveEvent is an event detail for receiving data
type MetricReceiveEvent struct {
	Metric  *Metric
	message event.Message
}

var _ datamodel.ModelReceiveEvent = (*MetricReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *MetricReceiveEvent) Object() datamodel.Model {
	return e.Metric
}

// Message returns the underlying message data for the event
func (e *MetricReceiveEvent) Message() event.Message {
	return e.message
}

// MetricProducer implements the datamodel.ModelEventProducer
type MetricProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*MetricProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *MetricProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *MetricProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Metric) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &MetricProducer{
		ch:   ch,
		done: NewMetricProducer(producer, ch, errors),
	}
}

// NewMetricProducerChannel returns a channel which can be used for producing Model events
func NewMetricProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &MetricProducer{
		ch:   ch,
		done: NewMetricProducer(producer, ch, errors),
	}
}

// MetricConsumer implements the datamodel.ModelEventConsumer
type MetricConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*MetricConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *MetricConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *MetricConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Metric) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewMetricConsumer(consumer, ch, errors)
	return &MetricConsumer{
		ch: ch,
	}
}

// NewMetricConsumerChannel returns a consumer channel which can be used to consume Model events
func NewMetricConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewMetricConsumer(consumer, ch, errors)
	return &MetricConsumer{
		ch: ch,
	}
}
