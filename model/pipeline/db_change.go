// DO NOT EDIT -- generated code

// Package pipeline - the pipeline models
package pipeline

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

const (
	// DbChangeTopic is the default topic name
	DbChangeTopic datamodel.TopicNameType = "pipeline_DbChange_topic"

	// DbChangeStream is the default stream name
	DbChangeStream datamodel.TopicNameType = "pipeline_DbChange_stream"

	// DbChangeTable is the default table name
	DbChangeTable datamodel.TopicNameType = "pipeline_DbChange"
	// DbChangeModelName is the model name
	DbChangeModelName datamodel.ModelNameType = "pipeline.DbChange"
)

// Action is the enumeration type for action
type Action int32

// String returns the string value for Action
func (v Action) String() string {
	switch int32(v) {
	case 0:
		return "create"
	case 1:
		return "update"
	case 2:
		return "delete"
	}
	return "unset"
}

const (
	// ActionCreate is the enumeration value for create
	ActionCreate Action = 0
	// ActionUpdate is the enumeration value for update
	ActionUpdate Action = 1
	// ActionDelete is the enumeration value for delete
	ActionDelete Action = 2
)

const (
	// DbChangeDateAtColumn is the date_ts column name
	DbChangeDateAtColumn = "date_ts"
	// DbChangeModelColumn is the model column name
	DbChangeModelColumn = "model"
	// DbChangeActionColumn is the action column name
	DbChangeActionColumn = "action"
	// DbChangeDataColumn is the data column name
	DbChangeDataColumn = "data"
)

// DbChange db change will contain all the changes to a specific data model from the DB
type DbChange struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// DateAt the date when the change was made
	DateAt int64 `json:"date_ts" bson:"date_ts" yaml:"date_ts" faker:"-"`
	// Model the name of the model
	Model string `json:"model" bson:"model" yaml:"model" faker:"-"`
	// Action the action that was taken
	Action Action `json:"action" bson:"action" yaml:"action" faker:"-"`
	// Data the data payload of the change
	Data string `json:"data" bson:"data" yaml:"data" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*DbChange)(nil)

func toDbChangeObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toDbChangeObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toDbChangeObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toDbChangeObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toDbChangeObjectNil(isavro, isoptional)
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
	case *DbChange:
		return v.ToMap()
	case DbChange:
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
			arr = append(arr, toDbChangeObject(av, isavro, false, ""))
		}
		return arr
	case Action:
		if !isavro {
			return (o.(Action)).String()
		}
		return map[string]string{
			"pipeline.action": (o.(Action)).String(),
		}
	case *Action:
		if !isavro {
			return (o.(*Action)).String()
		}
		return map[string]string{
			"pipeline.action": (o.(*Action)).String(),
		}
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of DbChange
func (o *DbChange) String() string {
	return fmt.Sprintf("pipeline.DbChange<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *DbChange) GetTopicName() datamodel.TopicNameType {
	return DbChangeTopic
}

// GetModelName returns the name of the model
func (o *DbChange) GetModelName() datamodel.ModelNameType {
	return DbChangeModelName
}

func (o *DbChange) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *DbChange) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.Model, o.Action, o.DateAt)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *DbChange) GetTopicKey() string {
	var i interface{} = o.Model
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *DbChange) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for DbChange")
}

// GetRefID returns the RefID for the object
func (o *DbChange) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *DbChange) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *DbChange) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *DbChange) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *DbChange) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = DbChangeModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *DbChange) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "model",
		Timestamp:         "date_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *DbChange) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of DbChange
func (o *DbChange) Clone() datamodel.Model {
	c := new(DbChange)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *DbChange) Anon() datamodel.Model {
	c := new(DbChange)
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
func (o *DbChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *DbChange) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecDbChange *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *DbChange) GetAvroCodec() *goavro.Codec {
	if cachedCodecDbChange == nil {
		c, err := GetDbChangeAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecDbChange = c
	}
	return cachedCodecDbChange
}

// ToAvroBinary returns the data as Avro binary data
func (o *DbChange) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *DbChange) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two DbChange objects are equal
func (o *DbChange) IsEqual(other *DbChange) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *DbChange) ToMap(avro ...bool) map[string]interface{} {
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
		"date_ts":     toDbChangeObject(o.DateAt, isavro, false, "long"),
		"model":       toDbChangeObject(o.Model, isavro, false, "string"),
		"action":      toDbChangeObject(o.Action, isavro, false, "action"),
		"data":        toDbChangeObject(o.Data, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *DbChange) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["model"].(string); ok {
		o.Model = val
	} else {
		val := kv["model"]
		if val == nil {
			o.Model = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Model = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["action"].(Action); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {
			ev := em["pipeline.action"].(string)
			switch ev {
			case "create":
				o.Action = 0
			case "update":
				o.Action = 1
			case "delete":
				o.Action = 2
			}
		}
		if em, ok := kv["action"].(string); ok {
			switch em {
			case "create":
				o.Action = 0
			case "update":
				o.Action = 1
			case "delete":
				o.Action = 2
			}
		}
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
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *DbChange) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.DateAt)
	args = append(args, o.Model)
	args = append(args, o.Action)
	args = append(args, o.Data)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetDbChangeAvroSchemaSpec creates the avro schema specification for DbChange
func GetDbChangeAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "pipeline",
		"name":         "DbChange",
		"connect.name": "pipeline.DbChange",
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
				"name": "model",
				"type": "string",
			},
			map[string]interface{}{
				"name": "action",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "action",
						"symbols": []interface{}{"create", "update", "delete"},
					},
				},
			},
			map[string]interface{}{
				"name": "data",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetDbChangeAvroSchema creates the avro schema for DbChange
func GetDbChangeAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetDbChangeAvroSchemaSpec())
}

// TransformDbChangeFunc is a function for transforming DbChange during processing
type TransformDbChangeFunc func(input *DbChange) (*DbChange, error)

// NewDbChangePipe creates a pipe for processing DbChange items
func NewDbChangePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformDbChangeFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewDbChangeInputStream(input, errors)
	var stream chan DbChange
	if len(transforms) > 0 {
		stream = make(chan DbChange, 1000)
	} else {
		stream = inch
	}
	outdone := NewDbChangeOutputStream(output, stream, errors)
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

// NewDbChangeInputStreamDir creates a channel for reading DbChange as JSON newlines from a directory of files
func NewDbChangeInputStreamDir(dir string, errors chan<- error, transforms ...TransformDbChangeFunc) (chan DbChange, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline/db_change\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan DbChange)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for db_change")
		ch := make(chan DbChange)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewDbChangeInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan DbChange)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewDbChangeInputStreamFile creates an channel for reading DbChange as JSON newlines from filename
func NewDbChangeInputStreamFile(filename string, errors chan<- error, transforms ...TransformDbChangeFunc) (chan DbChange, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan DbChange)
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
			ch := make(chan DbChange)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewDbChangeInputStream(f, errors, transforms...)
}

// NewDbChangeInputStream creates an channel for reading DbChange as JSON newlines from stream
func NewDbChangeInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformDbChangeFunc) (chan DbChange, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan DbChange, 1000)
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
			var item DbChange
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

// NewDbChangeOutputStreamDir will output json newlines from channel and save in dir
func NewDbChangeOutputStreamDir(dir string, ch chan DbChange, errors chan<- error, transforms ...TransformDbChangeFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline/db_change\\.json(\\.gz)?$")
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
	return NewDbChangeOutputStream(gz, ch, errors, transforms...)
}

// NewDbChangeOutputStream will output json newlines from channel to the stream
func NewDbChangeOutputStream(stream io.WriteCloser, ch chan DbChange, errors chan<- error, transforms ...TransformDbChangeFunc) <-chan bool {
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

// DbChangeSendEvent is an event detail for sending data
type DbChangeSendEvent struct {
	DbChange *DbChange
	headers  map[string]string
	time     time.Time
	key      string
}

var _ datamodel.ModelSendEvent = (*DbChangeSendEvent)(nil)

// Key is the key to use for the message
func (e *DbChangeSendEvent) Key() string {
	if e.key == "" {
		return e.DbChange.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *DbChangeSendEvent) Object() datamodel.Model {
	return e.DbChange
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *DbChangeSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *DbChangeSendEvent) Timestamp() time.Time {
	return e.time
}

// DbChangeSendEventOpts is a function handler for setting opts
type DbChangeSendEventOpts func(o *DbChangeSendEvent)

// WithDbChangeSendEventKey sets the key value to a value different than the object ID
func WithDbChangeSendEventKey(key string) DbChangeSendEventOpts {
	return func(o *DbChangeSendEvent) {
		o.key = key
	}
}

// WithDbChangeSendEventTimestamp sets the timestamp value
func WithDbChangeSendEventTimestamp(tv time.Time) DbChangeSendEventOpts {
	return func(o *DbChangeSendEvent) {
		o.time = tv
	}
}

// WithDbChangeSendEventHeader sets the timestamp value
func WithDbChangeSendEventHeader(key, value string) DbChangeSendEventOpts {
	return func(o *DbChangeSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewDbChangeSendEvent returns a new DbChangeSendEvent instance
func NewDbChangeSendEvent(o *DbChange, opts ...DbChangeSendEventOpts) *DbChangeSendEvent {
	res := &DbChangeSendEvent{
		DbChange: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewDbChangeProducer will stream data from the channel
func NewDbChangeProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*DbChange); ok {
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
				errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.DbChange but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewDbChangeConsumer will stream data from the topic into the provided channel
func NewDbChangeConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object DbChange
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into pipeline.DbChange: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &DbChangeReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// DbChangeReceiveEvent is an event detail for receiving data
type DbChangeReceiveEvent struct {
	DbChange *DbChange
	message  event.Message
}

var _ datamodel.ModelReceiveEvent = (*DbChangeReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *DbChangeReceiveEvent) Object() datamodel.Model {
	return e.DbChange
}

// Message returns the underlying message data for the event
func (e *DbChangeReceiveEvent) Message() event.Message {
	return e.message
}

// DbChangeProducer implements the datamodel.ModelEventProducer
type DbChangeProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*DbChangeProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *DbChangeProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *DbChangeProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *DbChange) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &DbChangeProducer{
		ch:   ch,
		done: NewDbChangeProducer(producer, ch, errors),
	}
}

// NewDbChangeProducerChannel returns a channel which can be used for producing Model events
func NewDbChangeProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &DbChangeProducer{
		ch:   ch,
		done: NewDbChangeProducer(producer, ch, errors),
	}
}

// DbChangeConsumer implements the datamodel.ModelEventConsumer
type DbChangeConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*DbChangeConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *DbChangeConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *DbChangeConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *DbChange) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewDbChangeConsumer(consumer, ch, errors)
	return &DbChangeConsumer{
		ch: ch,
	}
}

// NewDbChangeConsumerChannel returns a consumer channel which can be used to consume Model events
func NewDbChangeConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewDbChangeConsumer(consumer, ch, errors)
	return &DbChangeConsumer{
		ch: ch,
	}
}
