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
	// ExportTriggerTopic is the default topic name
	ExportTriggerTopic datamodel.TopicNameType = "agent_ExportTrigger_topic"

	// ExportTriggerStream is the default stream name
	ExportTriggerStream datamodel.TopicNameType = "agent_ExportTrigger_stream"

	// ExportTriggerTable is the default table name
	ExportTriggerTable datamodel.TopicNameType = "agent_ExportTrigger"

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
	// ExportTriggerRequestDateColumn is the request_date column name
	ExportTriggerRequestDateColumn = "request_date"
	// ExportTriggerRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	ExportTriggerRequestDateColumnEpochColumn = "request_date->epoch"
	// ExportTriggerRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	ExportTriggerRequestDateColumnOffsetColumn = "request_date->offset"
	// ExportTriggerRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	ExportTriggerRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// ExportTriggerUUIDColumn is the uuid column name
	ExportTriggerUUIDColumn = "uuid"
)

// ExportTriggerRequestDate represents the object structure for request_date
type ExportTriggerRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *ExportTriggerRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

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
	// RequestDate the date when the request was made
	RequestDate ExportTriggerRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
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

	if res := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); res != nil {
		return res
	}
	switch v := o.(type) {
	case *ExportTrigger:
		return v.ToMap()
	case ExportTrigger:
		return v.ToMap()

	case ExportTriggerRequestDate:
		vv := o.(ExportTriggerRequestDate)
		return vv.ToMap()
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
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

func (o *ExportTrigger) setDefaults() {

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportTrigger) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportTrigger", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportTrigger) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportTrigger) GetTimestamp() time.Time {
	var dt interface{} = o.RequestDate
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
	case ExportTriggerRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
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
		Timestamp:         "request_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *ExportTrigger) GetStateKey() string {
	key := "uuid"
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
	return nil
}

var cachedCodecExportTrigger *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ExportTrigger) GetAvroCodec() *goavro.Codec {
	if cachedCodecExportTrigger == nil {
		c, err := GetExportTriggerAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecExportTrigger = c
	}
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
	o.setDefaults()
	return map[string]interface{}{
		"customer_id":  toExportTriggerObject(o.CustomerID, isavro, false, "string"),
		"id":           toExportTriggerObject(o.ID, isavro, false, "string"),
		"ref_id":       toExportTriggerObject(o.RefID, isavro, false, "string"),
		"ref_type":     toExportTriggerObject(o.RefType, isavro, false, "string"),
		"request_date": toExportTriggerObject(o.RequestDate, isavro, false, "request_date"),
		"uuid":         toExportTriggerObject(o.UUID, isavro, false, "string"),
		"hashcode":     toExportTriggerObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportTrigger) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["request_date"].(ExportTriggerRequestDate); ok {
		o.RequestDate = val
	} else {
		val := kv["request_date"]
		if val == nil {
			o.RequestDate = ExportTriggerRequestDate{}
		} else {
			o.RequestDate = ExportTriggerRequestDate{}
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
			json.Unmarshal(b, &o.RequestDate)

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

// Hash will return a hashcode for the object
func (o *ExportTrigger) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
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
				"name": "request_date",
				"type": map[string]interface{}{"doc": "the date when the request was made", "type": "record", "name": "request_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}},
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

// TransformExportTriggerFunc is a function for transforming ExportTrigger during processing
type TransformExportTriggerFunc func(input *ExportTrigger) (*ExportTrigger, error)

// NewExportTriggerPipe creates a pipe for processing ExportTrigger items
func NewExportTriggerPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformExportTriggerFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewExportTriggerInputStream(input, errors)
	var stream chan ExportTrigger
	if len(transforms) > 0 {
		stream = make(chan ExportTrigger, 1000)
	} else {
		stream = inch
	}
	outdone := NewExportTriggerOutputStream(output, stream, errors)
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

// NewExportTriggerInputStreamDir creates a channel for reading ExportTrigger as JSON newlines from a directory of files
func NewExportTriggerInputStreamDir(dir string, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/export_trigger\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for export_trigger")
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewExportTriggerInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ExportTrigger)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewExportTriggerInputStreamFile creates an channel for reading ExportTrigger as JSON newlines from filename
func NewExportTriggerInputStreamFile(filename string, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ExportTrigger)
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
			ch := make(chan ExportTrigger)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewExportTriggerInputStream(f, errors, transforms...)
}

// NewExportTriggerInputStream creates an channel for reading ExportTrigger as JSON newlines from stream
func NewExportTriggerInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformExportTriggerFunc) (chan ExportTrigger, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ExportTrigger, 1000)
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
			var item ExportTrigger
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

// NewExportTriggerOutputStreamDir will output json newlines from channel and save in dir
func NewExportTriggerOutputStreamDir(dir string, ch chan ExportTrigger, errors chan<- error, transforms ...TransformExportTriggerFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/export_trigger\\.json(\\.gz)?$")
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
	return NewExportTriggerOutputStream(gz, ch, errors, transforms...)
}

// NewExportTriggerOutputStream will output json newlines from channel to the stream
func NewExportTriggerOutputStream(stream io.WriteCloser, ch chan ExportTrigger, errors chan<- error, transforms ...TransformExportTriggerFunc) <-chan bool {
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
