// DO NOT EDIT -- generated code

// Package web - the web model
package web

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
	// HookTopic is the default topic name
	HookTopic datamodel.TopicNameType = "web_Hook_topic"

	// HookStream is the default stream name
	HookStream datamodel.TopicNameType = "web_Hook_stream"

	// HookTable is the default table name
	HookTable datamodel.TopicNameType = "web_Hook"

	// HookModelName is the model name
	HookModelName datamodel.ModelNameType = "web.Hook"
)

const (
	// HookDataColumn is the data column name
	HookDataColumn = "data"
	// HookDateColumn is the date column name
	HookDateColumn = "date"
	// HookDateColumnEpochColumn is the epoch column property of the Date name
	HookDateColumnEpochColumn = "date->epoch"
	// HookDateColumnOffsetColumn is the offset column property of the Date name
	HookDateColumnOffsetColumn = "date->offset"
	// HookDateColumnRfc3339Column is the rfc3339 column property of the Date name
	HookDateColumnRfc3339Column = "date->rfc3339"
	// HookHeadersColumn is the headers column name
	HookHeadersColumn = "headers"
	// HookIDColumn is the id column name
	HookIDColumn = "id"
	// HookSystemColumn is the system column name
	HookSystemColumn = "system"
	// HookTokenColumn is the token column name
	HookTokenColumn = "token"
)

// HookDate represents the object structure for date
type HookDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *HookDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Hook hook is a webhook event which is received from an external source
type Hook struct {
	// Data the webhook data payload base64 encoded
	Data string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Date the date when the hook was received
	Date HookDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// Headers the headers of the incoming webhook
	Headers map[string]string `json:"headers" bson:"headers" yaml:"headers" faker:"-"`
	// ID the primary key for this model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// System the name of the system sending the hook
	System string `json:"system" bson:"system" yaml:"system" faker:"-"`
	// Token the token part of the url
	Token string `json:"token" bson:"token" yaml:"token" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Hook)(nil)

func toHookObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toHookObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toHookObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toHookObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toHookObjectNil(isavro, isoptional)
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
	case *Hook:
		return v.ToMap()
	case Hook:
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
			arr = append(arr, toHookObject(av, isavro, false, ""))
		}
		return arr

	case HookDate:
		vv := o.(HookDate)
		return vv.ToMap()
	case *HookDate:
		return map[string]interface{}{
			"web.date": (*o.(*HookDate)).ToMap(),
		}
	case []HookDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]HookDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]HookDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]HookDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Hook
func (o *Hook) String() string {
	return fmt.Sprintf("web.Hook<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Hook) GetTopicName() datamodel.TopicNameType {
	return HookTopic
}

// GetModelName returns the name of the model
func (o *Hook) GetModelName() datamodel.ModelNameType {
	return HookModelName
}

func (o *Hook) setDefaults() {

	o.GetID()
}

// GetID returns the ID for the object
func (o *Hook) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.System, o.Token, o.Date.Epoch)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Hook) GetTopicKey() string {
	var i interface{} = o.System
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Hook) GetTimestamp() time.Time {
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
	case HookDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Hook")
}

// IsMaterialized returns true if the model is materialized
func (o *Hook) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Hook) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Hook) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Hook) SetEventHeaders(kv map[string]string) {
	kv["model"] = HookModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Hook) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "system",
		Timestamp:         "date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Hook) GetStateKey() string {
	key := "system"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// Clone returns an exact copy of Hook
func (o *Hook) Clone() datamodel.Model {
	c := new(Hook)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Hook) Anon() datamodel.Model {
	c := new(Hook)
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
func (o *Hook) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Hook) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecHook *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Hook) GetAvroCodec() *goavro.Codec {
	if cachedCodecHook == nil {
		c, err := GetHookAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecHook = c
	}
	return cachedCodecHook
}

// ToAvroBinary returns the data as Avro binary data
func (o *Hook) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Hook) FromAvroBinary(value []byte) error {
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
func (o *Hook) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Hook objects are equal
func (o *Hook) IsEqual(other *Hook) bool {
	return o.GetID() == other.GetID()
}

// ToMap returns the object as a map
func (o *Hook) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Headers == nil {
			o.Headers = make(map[string]string)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"data":    toHookObject(o.Data, isavro, false, "string"),
		"date":    toHookObject(o.Date, isavro, false, "date"),
		"headers": toHookObject(o.Headers, isavro, false, "string"),
		"id":      toHookObject(o.ID, isavro, false, "string"),
		"system":  toHookObject(o.System, isavro, false, "string"),
		"token":   toHookObject(o.Token, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Hook) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["data"].(string); ok {
		o.Data = val
	} else {
		val := kv["data"]
		if val == nil {
			o.Data = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Data = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["date"].(HookDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = HookDate{}
		} else {
			o.Date = HookDate{}
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
	if val := kv["headers"]; val != nil {
		kv := make(map[string]string)
		if m, ok := val.(map[string]string); ok {
			kv = m
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				for k, v := range m {
					if mv, ok := v.(string); ok {
						kv[k] = mv
					} else {
						kv[k] = fmt.Sprintf("%v", v)
					}
				}
			} else {
				panic("unsupported type for headers field entry: " + reflect.TypeOf(val).String())
			}
		}
		o.Headers = kv
	} else {
		o.Headers = map[string]string{}
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
	if val, ok := kv["system"].(string); ok {
		o.System = val
	} else {
		val := kv["system"]
		if val == nil {
			o.System = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.System = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["token"].(string); ok {
		o.Token = val
	} else {
		val := kv["token"]
		if val == nil {
			o.Token = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Token = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// GetHookAvroSchemaSpec creates the avro schema specification for Hook
func GetHookAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "web",
		"name":      "Hook",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "data",
				"type": "string",
			},
			map[string]interface{}{
				"name": "date",
				"type": map[string]interface{}{"name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the hook was received", "type": "record"},
			},
			map[string]interface{}{
				"name": "headers",
				"type": map[string]interface{}{
					"type":   "map",
					"values": "string",
				},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "system",
				"type": "string",
			},
			map[string]interface{}{
				"name": "token",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Hook) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetHookAvroSchema creates the avro schema for Hook
func GetHookAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetHookAvroSchemaSpec())
}

// TransformHookFunc is a function for transforming Hook during processing
type TransformHookFunc func(input *Hook) (*Hook, error)

// NewHookPipe creates a pipe for processing Hook items
func NewHookPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformHookFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewHookInputStream(input, errors)
	var stream chan Hook
	if len(transforms) > 0 {
		stream = make(chan Hook, 1000)
	} else {
		stream = inch
	}
	outdone := NewHookOutputStream(output, stream, errors)
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

// NewHookInputStreamDir creates a channel for reading Hook as JSON newlines from a directory of files
func NewHookInputStreamDir(dir string, errors chan<- error, transforms ...TransformHookFunc) (chan Hook, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/web/hook\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Hook)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for hook")
		ch := make(chan Hook)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewHookInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Hook)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewHookInputStreamFile creates an channel for reading Hook as JSON newlines from filename
func NewHookInputStreamFile(filename string, errors chan<- error, transforms ...TransformHookFunc) (chan Hook, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Hook)
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
			ch := make(chan Hook)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewHookInputStream(f, errors, transforms...)
}

// NewHookInputStream creates an channel for reading Hook as JSON newlines from stream
func NewHookInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformHookFunc) (chan Hook, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Hook, 1000)
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
			var item Hook
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

// NewHookOutputStreamDir will output json newlines from channel and save in dir
func NewHookOutputStreamDir(dir string, ch chan Hook, errors chan<- error, transforms ...TransformHookFunc) <-chan bool {
	fp := filepath.Join(dir, "/web/hook\\.json(\\.gz)?$")
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
	return NewHookOutputStream(gz, ch, errors, transforms...)
}

// NewHookOutputStream will output json newlines from channel to the stream
func NewHookOutputStream(stream io.WriteCloser, ch chan Hook, errors chan<- error, transforms ...TransformHookFunc) <-chan bool {
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

// HookSendEvent is an event detail for sending data
type HookSendEvent struct {
	Hook    *Hook
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*HookSendEvent)(nil)

// Key is the key to use for the message
func (e *HookSendEvent) Key() string {
	if e.key == "" {
		return e.Hook.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *HookSendEvent) Object() datamodel.Model {
	return e.Hook
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *HookSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *HookSendEvent) Timestamp() time.Time {
	return e.time
}

// HookSendEventOpts is a function handler for setting opts
type HookSendEventOpts func(o *HookSendEvent)

// WithHookSendEventKey sets the key value to a value different than the object ID
func WithHookSendEventKey(key string) HookSendEventOpts {
	return func(o *HookSendEvent) {
		o.key = key
	}
}

// WithHookSendEventTimestamp sets the timestamp value
func WithHookSendEventTimestamp(tv time.Time) HookSendEventOpts {
	return func(o *HookSendEvent) {
		o.time = tv
	}
}

// WithHookSendEventHeader sets the timestamp value
func WithHookSendEventHeader(key, value string) HookSendEventOpts {
	return func(o *HookSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewHookSendEvent returns a new HookSendEvent instance
func NewHookSendEvent(o *Hook, opts ...HookSendEventOpts) *HookSendEvent {
	res := &HookSendEvent{
		Hook: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewHookProducer will stream data from the channel
func NewHookProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Hook); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type web.Hook but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewHookConsumer will stream data from the topic into the provided channel
func NewHookConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Hook
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into web.Hook: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into web.Hook: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for web.Hook")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &HookReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Hook
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &HookReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// HookReceiveEvent is an event detail for receiving data
type HookReceiveEvent struct {
	Hook    *Hook
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*HookReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *HookReceiveEvent) Object() datamodel.Model {
	return e.Hook
}

// Message returns the underlying message data for the event
func (e *HookReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *HookReceiveEvent) EOF() bool {
	return e.eof
}

// HookProducer implements the datamodel.ModelEventProducer
type HookProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*HookProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *HookProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *HookProducer) Close() error {
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
func (o *Hook) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Hook) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &HookProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewHookProducer(newctx, producer, ch, errors, empty),
	}
}

// NewHookProducerChannel returns a channel which can be used for producing Model events
func NewHookProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewHookProducerChannelSize(producer, 0, errors)
}

// NewHookProducerChannelSize returns a channel which can be used for producing Model events
func NewHookProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &HookProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewHookProducer(newctx, producer, ch, errors, empty),
	}
}

// HookConsumer implements the datamodel.ModelEventConsumer
type HookConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*HookConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *HookConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *HookConsumer) Close() error {
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
func (o *Hook) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &HookConsumer{
		ch:       ch,
		callback: NewHookConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewHookConsumerChannel returns a consumer channel which can be used to consume Model events
func NewHookConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &HookConsumer{
		ch:       ch,
		callback: NewHookConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
