// DO NOT EDIT -- generated code

// Package reprocess - the pipeline models
package reprocess

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
	"strings"
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
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// MaterializedLogTopic is the default topic name
	MaterializedLogTopic datamodel.TopicNameType = "pipeline_reprocess_MaterializedLog_topic"

	// MaterializedLogStream is the default stream name
	MaterializedLogStream datamodel.TopicNameType = "pipeline_reprocess_MaterializedLog_stream"

	// MaterializedLogTable is the default table name
	MaterializedLogTable datamodel.TopicNameType = "pipeline_reprocess_MaterializedLog"

	// MaterializedLogModelName is the model name
	MaterializedLogModelName datamodel.ModelNameType = "pipeline.reprocess.MaterializedLog"
)

const (
	// MaterializedLogCustomerIDColumn is the customer_id column name
	MaterializedLogCustomerIDColumn = "customer_id"
	// MaterializedLogIDColumn is the id column name
	MaterializedLogIDColumn = "id"
	// MaterializedLogModelColumn is the model column name
	MaterializedLogModelColumn = "model"
	// MaterializedLogModelIdsColumn is the model_ids column name
	MaterializedLogModelIdsColumn = "model_ids"
	// MaterializedLogRecordCountColumn is the record_count column name
	MaterializedLogRecordCountColumn = "record_count"
	// MaterializedLogRefIDColumn is the ref_id column name
	MaterializedLogRefIDColumn = "ref_id"
	// MaterializedLogRefTypeColumn is the ref_type column name
	MaterializedLogRefTypeColumn = "ref_type"
	// MaterializedLogSavedAtColumn is the saved_ts column name
	MaterializedLogSavedAtColumn = "saved_ts"
)

// MaterializedLog
type MaterializedLog struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Model the name of the model
	Model string `json:"model" bson:"model" yaml:"model" faker:"-"`
	// ModelIds ids of the changed records
	ModelIds []string `json:"model_ids" bson:"model_ids" yaml:"model_ids" faker:"-"`
	// RecordCount count of records in batched save
	RecordCount int64 `json:"record_count" bson:"record_count" yaml:"record_count" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SavedAt timestamp of data persisted
	SavedAt int64 `json:"saved_ts" bson:"saved_ts" yaml:"saved_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*MaterializedLog)(nil)

func toMaterializedLogObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toMaterializedLogObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *MaterializedLog:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of MaterializedLog
func (o *MaterializedLog) String() string {
	return fmt.Sprintf("pipeline.reprocess.MaterializedLog<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *MaterializedLog) GetTopicName() datamodel.TopicNameType {
	return MaterializedLogTopic
}

// GetModelName returns the name of the model
func (o *MaterializedLog) GetModelName() datamodel.ModelNameType {
	return MaterializedLogModelName
}

func (o *MaterializedLog) setDefaults(frommap bool) {
	if o.ModelIds == nil {
		o.ModelIds = make([]string, 0)
	}

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *MaterializedLog) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("MaterializedLog", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *MaterializedLog) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *MaterializedLog) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *MaterializedLog) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *MaterializedLog) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *MaterializedLog) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *MaterializedLog) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *MaterializedLog) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = MaterializedLogModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *MaterializedLog) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *MaterializedLog) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *MaterializedLog) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of MaterializedLog
func (o *MaterializedLog) Clone() datamodel.Model {
	c := new(MaterializedLog)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *MaterializedLog) Anon() datamodel.Model {
	c := new(MaterializedLog)
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
func (o *MaterializedLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *MaterializedLog) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecMaterializedLog *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *MaterializedLog) GetAvroCodec() *goavro.Codec {
	if cachedCodecMaterializedLog == nil {
		c, err := GetMaterializedLogAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecMaterializedLog = c
	}
	return cachedCodecMaterializedLog
}

// ToAvroBinary returns the data as Avro binary data
func (o *MaterializedLog) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *MaterializedLog) FromAvroBinary(value []byte) error {
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
func (o *MaterializedLog) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two MaterializedLog objects are equal
func (o *MaterializedLog) IsEqual(other *MaterializedLog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *MaterializedLog) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.ModelIds == nil {
			o.ModelIds = make([]string, 0)
		}
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"customer_id":  toMaterializedLogObject(o.CustomerID, isavro, false, "string"),
		"id":           toMaterializedLogObject(o.ID, isavro, false, "string"),
		"model":        toMaterializedLogObject(o.Model, isavro, false, "string"),
		"model_ids":    toMaterializedLogObject(o.ModelIds, isavro, false, "model_ids"),
		"record_count": toMaterializedLogObject(o.RecordCount, isavro, false, "long"),
		"ref_id":       toMaterializedLogObject(o.RefID, isavro, false, "string"),
		"ref_type":     toMaterializedLogObject(o.RefType, isavro, false, "string"),
		"saved_ts":     toMaterializedLogObject(o.SavedAt, isavro, false, "long"),
		"hashcode":     toMaterializedLogObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *MaterializedLog) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["model"].(string); ok {
		o.Model = val
	} else {
		if val, ok := kv["model"]; ok {
			if val == nil {
				o.Model = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Model = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["model_ids"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for model_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for model_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for model_ids field")
				}
			}
			o.ModelIds = na
		}
	}
	if o.ModelIds == nil {
		o.ModelIds = make([]string, 0)
	}

	if val, ok := kv["record_count"].(int64); ok {
		o.RecordCount = val
	} else {
		if val, ok := kv["record_count"]; ok {
			if val == nil {
				o.RecordCount = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.RecordCount = number.ToInt64Any(val)
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

	if val, ok := kv["saved_ts"].(int64); ok {
		o.SavedAt = val
	} else {
		if val, ok := kv["saved_ts"]; ok {
			if val == nil {
				o.SavedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.SavedAt = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *MaterializedLog) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Model)
	args = append(args, o.ModelIds)
	args = append(args, o.RecordCount)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.SavedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetMaterializedLogAvroSchemaSpec creates the avro schema specification for MaterializedLog
func GetMaterializedLogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.reprocess",
		"name":      "MaterializedLog",
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
				"name": "model",
				"type": "string",
			},
			map[string]interface{}{
				"name": "model_ids",
				"type": map[string]interface{}{"type": "array", "name": "model_ids", "items": "string"},
			},
			map[string]interface{}{
				"name": "record_count",
				"type": "long",
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
				"name": "saved_ts",
				"type": "long",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *MaterializedLog) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetMaterializedLogAvroSchema creates the avro schema for MaterializedLog
func GetMaterializedLogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetMaterializedLogAvroSchemaSpec())
}

// TransformMaterializedLogFunc is a function for transforming MaterializedLog during processing
type TransformMaterializedLogFunc func(input *MaterializedLog) (*MaterializedLog, error)

// NewMaterializedLogPipe creates a pipe for processing MaterializedLog items
func NewMaterializedLogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformMaterializedLogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewMaterializedLogInputStream(input, errors)
	var stream chan MaterializedLog
	if len(transforms) > 0 {
		stream = make(chan MaterializedLog, 1000)
	} else {
		stream = inch
	}
	outdone := NewMaterializedLogOutputStream(output, stream, errors)
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

// NewMaterializedLogInputStreamDir creates a channel for reading MaterializedLog as JSON newlines from a directory of files
func NewMaterializedLogInputStreamDir(dir string, errors chan<- error, transforms ...TransformMaterializedLogFunc) (chan MaterializedLog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.reprocess/materialized_log\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan MaterializedLog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for materialized_log")
		ch := make(chan MaterializedLog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewMaterializedLogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan MaterializedLog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewMaterializedLogInputStreamFile creates an channel for reading MaterializedLog as JSON newlines from filename
func NewMaterializedLogInputStreamFile(filename string, errors chan<- error, transforms ...TransformMaterializedLogFunc) (chan MaterializedLog, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan MaterializedLog)
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
			ch := make(chan MaterializedLog)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewMaterializedLogInputStream(f, errors, transforms...)
}

// NewMaterializedLogInputStream creates an channel for reading MaterializedLog as JSON newlines from stream
func NewMaterializedLogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformMaterializedLogFunc) (chan MaterializedLog, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan MaterializedLog, 1000)
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
			var item MaterializedLog
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

// NewMaterializedLogOutputStreamDir will output json newlines from channel and save in dir
func NewMaterializedLogOutputStreamDir(dir string, ch chan MaterializedLog, errors chan<- error, transforms ...TransformMaterializedLogFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline.reprocess/materialized_log\\.json(\\.gz)?$")
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
	return NewMaterializedLogOutputStream(gz, ch, errors, transforms...)
}

// NewMaterializedLogOutputStream will output json newlines from channel to the stream
func NewMaterializedLogOutputStream(stream io.WriteCloser, ch chan MaterializedLog, errors chan<- error, transforms ...TransformMaterializedLogFunc) <-chan bool {
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

// MaterializedLogSendEvent is an event detail for sending data
type MaterializedLogSendEvent struct {
	MaterializedLog *MaterializedLog
	headers         map[string]string
	time            time.Time
	key             string
}

var _ datamodel.ModelSendEvent = (*MaterializedLogSendEvent)(nil)

// Key is the key to use for the message
func (e *MaterializedLogSendEvent) Key() string {
	if e.key == "" {
		return e.MaterializedLog.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *MaterializedLogSendEvent) Object() datamodel.Model {
	return e.MaterializedLog
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *MaterializedLogSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *MaterializedLogSendEvent) Timestamp() time.Time {
	return e.time
}

// MaterializedLogSendEventOpts is a function handler for setting opts
type MaterializedLogSendEventOpts func(o *MaterializedLogSendEvent)

// WithMaterializedLogSendEventKey sets the key value to a value different than the object ID
func WithMaterializedLogSendEventKey(key string) MaterializedLogSendEventOpts {
	return func(o *MaterializedLogSendEvent) {
		o.key = key
	}
}

// WithMaterializedLogSendEventTimestamp sets the timestamp value
func WithMaterializedLogSendEventTimestamp(tv time.Time) MaterializedLogSendEventOpts {
	return func(o *MaterializedLogSendEvent) {
		o.time = tv
	}
}

// WithMaterializedLogSendEventHeader sets the timestamp value
func WithMaterializedLogSendEventHeader(key, value string) MaterializedLogSendEventOpts {
	return func(o *MaterializedLogSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewMaterializedLogSendEvent returns a new MaterializedLogSendEvent instance
func NewMaterializedLogSendEvent(o *MaterializedLog, opts ...MaterializedLogSendEventOpts) *MaterializedLogSendEvent {
	res := &MaterializedLogSendEvent{
		MaterializedLog: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewMaterializedLogProducer will stream data from the channel
func NewMaterializedLogProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*MaterializedLog); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.reprocess.MaterializedLog but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewMaterializedLogConsumer will stream data from the topic into the provided channel
func NewMaterializedLogConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object MaterializedLog
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.reprocess.MaterializedLog: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into pipeline.reprocess.MaterializedLog: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.reprocess.MaterializedLog")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &MaterializedLogReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object MaterializedLog
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &MaterializedLogReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// MaterializedLogReceiveEvent is an event detail for receiving data
type MaterializedLogReceiveEvent struct {
	MaterializedLog *MaterializedLog
	message         eventing.Message
	eof             bool
}

var _ datamodel.ModelReceiveEvent = (*MaterializedLogReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *MaterializedLogReceiveEvent) Object() datamodel.Model {
	return e.MaterializedLog
}

// Message returns the underlying message data for the event
func (e *MaterializedLogReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *MaterializedLogReceiveEvent) EOF() bool {
	return e.eof
}

// MaterializedLogProducer implements the datamodel.ModelEventProducer
type MaterializedLogProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*MaterializedLogProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *MaterializedLogProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *MaterializedLogProducer) Close() error {
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
func (o *MaterializedLog) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *MaterializedLog) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &MaterializedLogProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewMaterializedLogProducer(newctx, producer, ch, errors, empty),
	}
}

// NewMaterializedLogProducerChannel returns a channel which can be used for producing Model events
func NewMaterializedLogProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewMaterializedLogProducerChannelSize(producer, 0, errors)
}

// NewMaterializedLogProducerChannelSize returns a channel which can be used for producing Model events
func NewMaterializedLogProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &MaterializedLogProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewMaterializedLogProducer(newctx, producer, ch, errors, empty),
	}
}

// MaterializedLogConsumer implements the datamodel.ModelEventConsumer
type MaterializedLogConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*MaterializedLogConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *MaterializedLogConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *MaterializedLogConsumer) Close() error {
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
func (o *MaterializedLog) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &MaterializedLogConsumer{
		ch:       ch,
		callback: NewMaterializedLogConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewMaterializedLogConsumerChannel returns a consumer channel which can be used to consume Model events
func NewMaterializedLogConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &MaterializedLogConsumer{
		ch:       ch,
		callback: NewMaterializedLogConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
