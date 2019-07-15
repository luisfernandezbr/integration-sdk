// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

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
)

const (
	// EnrollRequestTopic is the default topic name
	EnrollRequestTopic datamodel.TopicNameType = "agent_EnrollRequest_topic"

	// EnrollRequestStream is the default stream name
	EnrollRequestStream datamodel.TopicNameType = "agent_EnrollRequest_stream"

	// EnrollRequestTable is the default table name
	EnrollRequestTable datamodel.TopicNameType = "agent_EnrollRequest"

	// EnrollRequestModelName is the model name
	EnrollRequestModelName datamodel.ModelNameType = "agent.EnrollRequest"
)

const (
	// EnrollRequestCodeColumn is the code column name
	EnrollRequestCodeColumn = "code"
	// EnrollRequestDateColumn is the date column name
	EnrollRequestDateColumn = "date"
	// EnrollRequestDateColumnEpochColumn is the epoch column property of the Date name
	EnrollRequestDateColumnEpochColumn = "date->epoch"
	// EnrollRequestDateColumnOffsetColumn is the offset column property of the Date name
	EnrollRequestDateColumnOffsetColumn = "date->offset"
	// EnrollRequestDateColumnRfc3339Column is the rfc3339 column property of the Date name
	EnrollRequestDateColumnRfc3339Column = "date->rfc3339"
	// EnrollRequestIDColumn is the id column name
	EnrollRequestIDColumn = "id"
	// EnrollRequestUUIDColumn is the uuid column name
	EnrollRequestUUIDColumn = "uuid"
)

// EnrollRequestDate represents the object structure for date
type EnrollRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *EnrollRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// EnrollRequest an agent request to enroll a new agent machine
type EnrollRequest struct {
	// Code The agent enrollment code
	Code string `json:"code" bson:"code" yaml:"code" faker:"-"`
	// Date the date when the request was made
	Date EnrollRequestDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
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
	if o == nil {
		return toEnrollRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toEnrollRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toEnrollRequestObjectNil(isavro, isoptional)
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
	case *EnrollRequest:
		return v.ToMap()
	case EnrollRequest:
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
			arr = append(arr, toEnrollRequestObject(av, isavro, false, ""))
		}
		return arr

	case EnrollRequestDate:
		vv := o.(EnrollRequestDate)
		return vv.ToMap()
	case *EnrollRequestDate:
		return (*o.(*EnrollRequestDate)).ToMap()
	case []EnrollRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]EnrollRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]EnrollRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]EnrollRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
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

func (o *EnrollRequest) setDefaults() {

	o.GetID()
}

// GetID returns the ID for the object
func (o *EnrollRequest) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.Code, o.UUID)
	}
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
	var dt interface{} = o.Date
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
	case EnrollRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
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
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "date",
		NumPartitions:     8,
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
	return nil
}

var cachedCodecEnrollRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *EnrollRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecEnrollRequest == nil {
		c, err := GetEnrollRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecEnrollRequest = c
	}
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
	o.setDefaults()
	return map[string]interface{}{
		"code": toEnrollRequestObject(o.Code, isavro, false, "string"),
		"date": toEnrollRequestObject(o.Date, isavro, false, "date"),
		"id":   toEnrollRequestObject(o.ID, isavro, false, "string"),
		"uuid": toEnrollRequestObject(o.UUID, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollRequest) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["code"].(string); ok {
		o.Code = val
	} else {
		val := kv["code"]
		if val == nil {
			o.Code = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Code = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["date"].(EnrollRequestDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = EnrollRequestDate{}
		} else {
			o.Date = EnrollRequestDate{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Date)

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
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
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
				"name": "date",
				"type": map[string]interface{}{"type": "record", "name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"name": "offset", "doc": "the timezone offset from GMT", "type": "long"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the request was made"},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
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

// TransformEnrollRequestFunc is a function for transforming EnrollRequest during processing
type TransformEnrollRequestFunc func(input *EnrollRequest) (*EnrollRequest, error)

// NewEnrollRequestPipe creates a pipe for processing EnrollRequest items
func NewEnrollRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformEnrollRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewEnrollRequestInputStream(input, errors)
	var stream chan EnrollRequest
	if len(transforms) > 0 {
		stream = make(chan EnrollRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewEnrollRequestOutputStream(output, stream, errors)
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

// NewEnrollRequestInputStreamDir creates a channel for reading EnrollRequest as JSON newlines from a directory of files
func NewEnrollRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformEnrollRequestFunc) (chan EnrollRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/enroll_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan EnrollRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for enroll_request")
		ch := make(chan EnrollRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewEnrollRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan EnrollRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewEnrollRequestInputStreamFile creates an channel for reading EnrollRequest as JSON newlines from filename
func NewEnrollRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformEnrollRequestFunc) (chan EnrollRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan EnrollRequest)
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
			ch := make(chan EnrollRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewEnrollRequestInputStream(f, errors, transforms...)
}

// NewEnrollRequestInputStream creates an channel for reading EnrollRequest as JSON newlines from stream
func NewEnrollRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformEnrollRequestFunc) (chan EnrollRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan EnrollRequest, 1000)
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
			var item EnrollRequest
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

// NewEnrollRequestOutputStreamDir will output json newlines from channel and save in dir
func NewEnrollRequestOutputStreamDir(dir string, ch chan EnrollRequest, errors chan<- error, transforms ...TransformEnrollRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/enroll_request\\.json(\\.gz)?$")
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
	return NewEnrollRequestOutputStream(gz, ch, errors, transforms...)
}

// NewEnrollRequestOutputStream will output json newlines from channel to the stream
func NewEnrollRequestOutputStream(stream io.WriteCloser, ch chan EnrollRequest, errors chan<- error, transforms ...TransformEnrollRequestFunc) <-chan bool {
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
