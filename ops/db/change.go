// DO NOT EDIT -- generated code

// Package db - the ops models
package db

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
	// ChangeTopic is the default topic name
	ChangeTopic datamodel.TopicNameType = "ops_db_Change_topic"

	// ChangeStream is the default stream name
	ChangeStream datamodel.TopicNameType = "ops_db_Change_stream"

	// ChangeTable is the default table name
	ChangeTable datamodel.TopicNameType = "ops_db_Change"

	// ChangeModelName is the model name
	ChangeModelName datamodel.ModelNameType = "ops.db.Change"
)

const (
	// ChangeActionColumn is the action column name
	ChangeActionColumn = "action"
	// ChangeChangeDateColumn is the change_date column name
	ChangeChangeDateColumn = "change_date"
	// ChangeChangeDateColumnEpochColumn is the epoch column property of the ChangeDate name
	ChangeChangeDateColumnEpochColumn = "change_date->epoch"
	// ChangeChangeDateColumnOffsetColumn is the offset column property of the ChangeDate name
	ChangeChangeDateColumnOffsetColumn = "change_date->offset"
	// ChangeChangeDateColumnRfc3339Column is the rfc3339 column property of the ChangeDate name
	ChangeChangeDateColumnRfc3339Column = "change_date->rfc3339"
	// ChangeCustomerIDColumn is the customer_id column name
	ChangeCustomerIDColumn = "customer_id"
	// ChangeDataColumn is the data column name
	ChangeDataColumn = "data"
	// ChangeIDColumn is the id column name
	ChangeIDColumn = "id"
	// ChangeModelColumn is the model column name
	ChangeModelColumn = "model"
	// ChangeRefIDColumn is the ref_id column name
	ChangeRefIDColumn = "ref_id"
	// ChangeRefTypeColumn is the ref_type column name
	ChangeRefTypeColumn = "ref_type"
)

// ChangeAction is the enumeration type for action
type ChangeAction int32

// String returns the string value for Action
func (v ChangeAction) String() string {
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
	ChangeActionCreate ChangeAction = 0
	// ActionUpdate is the enumeration value for update
	ChangeActionUpdate ChangeAction = 1
	// ActionDelete is the enumeration value for delete
	ChangeActionDelete ChangeAction = 2
)

// ChangeChangeDate represents the object structure for change_date
type ChangeChangeDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *ChangeChangeDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// Change db change will contain all the changes to a specific data model from the DB
type Change struct {
	// Action the action that was taken
	Action ChangeAction `json:"action" bson:"action" yaml:"action" faker:"-"`
	// ChangeDate the date when the change was made
	ChangeDate ChangeChangeDate `json:"change_date" bson:"change_date" yaml:"change_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data the data payload of the change
	Data string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Model the name of the model
	Model string `json:"model" bson:"model" yaml:"model" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Change)(nil)

func toChangeObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toChangeObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {

	if res := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); res != nil {
		return res
	}
	switch v := o.(type) {
	case *Change:
		return v.ToMap()
	case Change:
		return v.ToMap()

	case ChangeAction:
		if !isavro {
			return (o.(ChangeAction)).String()
		}
		return (o.(ChangeAction)).String()

	case ChangeChangeDate:
		vv := o.(ChangeChangeDate)
		return vv.ToMap()
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Change
func (o *Change) String() string {
	return fmt.Sprintf("ops.db.Change<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Change) GetTopicName() datamodel.TopicNameType {
	return ChangeTopic
}

// GetModelName returns the name of the model
func (o *Change) GetModelName() datamodel.ModelNameType {
	return ChangeModelName
}

func (o *Change) setDefaults() {

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Change) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.Model, o.Action, o.ChangeDate.Epoch)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Change) GetTopicKey() string {
	var i interface{} = o.Model
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Change) GetTimestamp() time.Time {
	var dt interface{} = o.ChangeDate
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
	case ChangeChangeDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for Change")
}

// GetRefID returns the RefID for the object
func (o *Change) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Change) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Change) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Change) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Change) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ChangeModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Change) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "model",
		Timestamp:         "change_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Change) GetStateKey() string {
	key := "model"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Change) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Change
func (o *Change) Clone() datamodel.Model {
	c := new(Change)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Change) Anon() datamodel.Model {
	c := new(Change)
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
func (o *Change) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Change) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecChange *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Change) GetAvroCodec() *goavro.Codec {
	if cachedCodecChange == nil {
		c, err := GetChangeAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecChange = c
	}
	return cachedCodecChange
}

// ToAvroBinary returns the data as Avro binary data
func (o *Change) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Change) FromAvroBinary(value []byte) error {
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
func (o *Change) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Change objects are equal
func (o *Change) IsEqual(other *Change) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Change) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"action":      toChangeObject(o.Action, isavro, false, "action"),
		"change_date": toChangeObject(o.ChangeDate, isavro, false, "change_date"),
		"customer_id": toChangeObject(o.CustomerID, isavro, false, "string"),
		"data":        toChangeObject(o.Data, isavro, false, "string"),
		"id":          toChangeObject(o.ID, isavro, false, "string"),
		"model":       toChangeObject(o.Model, isavro, false, "string"),
		"ref_id":      toChangeObject(o.RefID, isavro, false, "string"),
		"ref_type":    toChangeObject(o.RefType, isavro, false, "string"),
		"hashcode":    toChangeObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Change) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["action"].(ChangeAction); ok {
		o.Action = val
	} else {
		if em, ok := kv["action"].(map[string]interface{}); ok {
			ev := em["ops.db.action"].(string)
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
	if val, ok := kv["change_date"].(ChangeChangeDate); ok {
		o.ChangeDate = val
	} else {
		val := kv["change_date"]
		if val == nil {
			o.ChangeDate = ChangeChangeDate{}
		} else {
			o.ChangeDate = ChangeChangeDate{}
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
			json.Unmarshal(b, &o.ChangeDate)

		}
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
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Change) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Action)
	args = append(args, o.ChangeDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.ID)
	args = append(args, o.Model)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetChangeAvroSchemaSpec creates the avro schema specification for Change
func GetChangeAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "ops.db",
		"name":      "Change",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "action",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "action",
					"symbols": []interface{}{"create", "update", "delete"},
				},
			},
			map[string]interface{}{
				"name": "change_date",
				"type": map[string]interface{}{"type": "record", "name": "change_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the change was made"},
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "data",
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
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Change) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetChangeAvroSchema creates the avro schema for Change
func GetChangeAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetChangeAvroSchemaSpec())
}

// TransformChangeFunc is a function for transforming Change during processing
type TransformChangeFunc func(input *Change) (*Change, error)

// NewChangePipe creates a pipe for processing Change items
func NewChangePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangeFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewChangeInputStream(input, errors)
	var stream chan Change
	if len(transforms) > 0 {
		stream = make(chan Change, 1000)
	} else {
		stream = inch
	}
	outdone := NewChangeOutputStream(output, stream, errors)
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

// NewChangeInputStreamDir creates a channel for reading Change as JSON newlines from a directory of files
func NewChangeInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangeFunc) (chan Change, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/ops.db/change\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Change)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for change")
		ch := make(chan Change)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewChangeInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Change)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewChangeInputStreamFile creates an channel for reading Change as JSON newlines from filename
func NewChangeInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangeFunc) (chan Change, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Change)
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
			ch := make(chan Change)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewChangeInputStream(f, errors, transforms...)
}

// NewChangeInputStream creates an channel for reading Change as JSON newlines from stream
func NewChangeInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangeFunc) (chan Change, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Change, 1000)
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
			var item Change
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

// NewChangeOutputStreamDir will output json newlines from channel and save in dir
func NewChangeOutputStreamDir(dir string, ch chan Change, errors chan<- error, transforms ...TransformChangeFunc) <-chan bool {
	fp := filepath.Join(dir, "/ops.db/change\\.json(\\.gz)?$")
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
	return NewChangeOutputStream(gz, ch, errors, transforms...)
}

// NewChangeOutputStream will output json newlines from channel to the stream
func NewChangeOutputStream(stream io.WriteCloser, ch chan Change, errors chan<- error, transforms ...TransformChangeFunc) <-chan bool {
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

// ChangeSendEvent is an event detail for sending data
type ChangeSendEvent struct {
	Change  *Change
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*ChangeSendEvent)(nil)

// Key is the key to use for the message
func (e *ChangeSendEvent) Key() string {
	if e.key == "" {
		return e.Change.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ChangeSendEvent) Object() datamodel.Model {
	return e.Change
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ChangeSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ChangeSendEvent) Timestamp() time.Time {
	return e.time
}

// ChangeSendEventOpts is a function handler for setting opts
type ChangeSendEventOpts func(o *ChangeSendEvent)

// WithChangeSendEventKey sets the key value to a value different than the object ID
func WithChangeSendEventKey(key string) ChangeSendEventOpts {
	return func(o *ChangeSendEvent) {
		o.key = key
	}
}

// WithChangeSendEventTimestamp sets the timestamp value
func WithChangeSendEventTimestamp(tv time.Time) ChangeSendEventOpts {
	return func(o *ChangeSendEvent) {
		o.time = tv
	}
}

// WithChangeSendEventHeader sets the timestamp value
func WithChangeSendEventHeader(key, value string) ChangeSendEventOpts {
	return func(o *ChangeSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewChangeSendEvent returns a new ChangeSendEvent instance
func NewChangeSendEvent(o *Change, opts ...ChangeSendEventOpts) *ChangeSendEvent {
	res := &ChangeSendEvent{
		Change: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewChangeProducer will stream data from the channel
func NewChangeProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*Change); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type ops.db.Change but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewChangeConsumer will stream data from the topic into the provided channel
func NewChangeConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Change
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into ops.db.Change: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into ops.db.Change: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for ops.db.Change")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &ChangeReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Change
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ChangeReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ChangeReceiveEvent is an event detail for receiving data
type ChangeReceiveEvent struct {
	Change  *Change
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*ChangeReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ChangeReceiveEvent) Object() datamodel.Model {
	return e.Change
}

// Message returns the underlying message data for the event
func (e *ChangeReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ChangeReceiveEvent) EOF() bool {
	return e.eof
}

// ChangeProducer implements the datamodel.ModelEventProducer
type ChangeProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ChangeProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ChangeProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ChangeProducer) Close() error {
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
func (o *Change) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Change) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ChangeProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewChangeProducer(newctx, producer, ch, errors, empty),
	}
}

// NewChangeProducerChannel returns a channel which can be used for producing Model events
func NewChangeProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewChangeProducerChannelSize(producer, 0, errors)
}

// NewChangeProducerChannelSize returns a channel which can be used for producing Model events
func NewChangeProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ChangeProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewChangeProducer(newctx, producer, ch, errors, empty),
	}
}

// ChangeConsumer implements the datamodel.ModelEventConsumer
type ChangeConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ChangeConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ChangeConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ChangeConsumer) Close() error {
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
func (o *Change) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ChangeConsumer{
		ch:       ch,
		callback: NewChangeConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewChangeConsumerChannel returns a consumer channel which can be used to consume Model events
func NewChangeConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ChangeConsumer{
		ch:       ch,
		callback: NewChangeConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
