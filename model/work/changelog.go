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
	"github.com/pinpt/go-common/number"
)

// ChangelogTopic is the default topic name
const ChangelogTopic datamodel.TopicNameType = "work_Changelog_topic"

// ChangelogStream is the default stream name
const ChangelogStream datamodel.TopicNameType = "work_Changelog_stream"

// ChangelogTable is the default table name
const ChangelogTable datamodel.TopicNameType = "work_Changelog"

// ChangelogModelName is the model name
const ChangelogModelName datamodel.ModelNameType = "work.Changelog"

// Changelog change log
type Changelog struct {
	// built in types

	ID         string `json:"changelog_id" bson:"changelog_id" yaml:"changelog_id" faker:"-"`
	MongoID    string `json:"_id" bson:"_id" yaml:"_id" faker:"-"` // generated and used internally, do not set
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// IssueID id of the issue
	IssueID string `json:"issue_id" bson:"issue_id" yaml:"issue_id" faker:"-"`
	// CreatedAt the timestamp in UTC when this change was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// Ordinal so we can order correctly in queries since dates could be equal
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// UserID id of the user of this change
	UserID string `json:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Field name of the field that was changed
	Field string `json:"field" bson:"field" yaml:"field" faker:"-"`
	// FieldType type of the field that was changed
	FieldType string `json:"field_type" bson:"field_type" yaml:"field_type" faker:"-"`
	// From id of the change from
	From string `json:"from" bson:"from" yaml:"from" faker:"-"`
	// FromString name of the change from
	FromString string `json:"from_string" bson:"from_string" yaml:"from_string" faker:"-"`
	// To id of the change to
	To string `json:"to" bson:"to" yaml:"to" faker:"-"`
	// ToString name of the change to
	ToString string `json:"to_string" bson:"to_string" yaml:"to_string" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Changelog)(nil)

func toChangelogObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toChangelogObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toChangelogObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toChangelogObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toChangelogObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Changelog:
		return v.ToMap()
	case Changelog:
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
			arr = append(arr, toChangelogObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("work.Changelog<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Changelog) GetTopicName() datamodel.TopicNameType {
	return ChangelogTopic
}

// GetModelName returns the name of the model
func (o *Changelog) GetModelName() datamodel.ModelNameType {
	return ChangelogModelName
}

func (o *Changelog) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Changelog) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Changelog", o.CustomerID, o.RefType, o.GetRefID())
	}
	if o.MongoID == "" {
		o.MongoID = o.ID
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Changelog) GetTopicKey() string {
	var i interface{} = o.IssueID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Changelog) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Changelog) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Changelog) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *Changelog) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "issue_id",
		Timestamp:         "created_ts",
		NumPartitions:     0,
		ReplicationFactor: 0,
		Retention:         duration,
		MaxSize:           0,
	}
}

// Clone returns an exact copy of Changelog
func (o *Changelog) Clone() datamodel.Model {
	c := new(Changelog)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Changelog) Anon() datamodel.Model {
	c := new(Changelog)
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
func (o *Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Changelog) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecChangelog *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *Changelog) GetAvroCodec() *goavro.Codec {
	if cachedCodecChangelog == nil {
		c, err := GetChangelogAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecChangelog = c
	}
	return cachedCodecChangelog
}

// ToAvroBinary returns the data as Avro binary data
func (o *Changelog) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Changelog) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"changelog_id": o.GetID(),
		"ref_id":       o.GetRefID(),
		"ref_type":     o.RefType,
		"customer_id":  o.CustomerID,
		"hashcode":     o.Hash(),
		"issue_id":     toChangelogObject(o.IssueID, isavro, false, "string"),
		"created_ts":   toChangelogObject(o.CreatedAt, isavro, false, "long"),
		"ordinal":      toChangelogObject(o.Ordinal, isavro, false, "long"),
		"user_id":      toChangelogObject(o.UserID, isavro, false, "string"),
		"field":        toChangelogObject(o.Field, isavro, false, "string"),
		"field_type":   toChangelogObject(o.FieldType, isavro, false, "string"),
		"from":         toChangelogObject(o.From, isavro, false, "string"),
		"from_string":  toChangelogObject(o.FromString, isavro, false, "string"),
		"to":           toChangelogObject(o.To, isavro, false, "string"),
		"to_string":    toChangelogObject(o.ToString, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
	if val, ok := kv["changelog_id"].(string); ok {
		o.ID = val
		o.MongoID = val
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
	if val, ok := kv["issue_id"].(string); ok {
		o.IssueID = val
	} else {
		val := kv["issue_id"]
		if val == nil {
			o.IssueID = ""
		} else {
			o.IssueID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		if val == nil {
			o.CreatedAt = number.ToInt64Any(nil)
		} else {
			o.CreatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		if val == nil {
			o.Ordinal = number.ToInt64Any(nil)
		} else {
			o.Ordinal = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		val := kv["user_id"]
		if val == nil {
			o.UserID = ""
		} else {
			o.UserID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field"].(string); ok {
		o.Field = val
	} else {
		val := kv["field"]
		if val == nil {
			o.Field = ""
		} else {
			o.Field = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["field_type"].(string); ok {
		o.FieldType = val
	} else {
		val := kv["field_type"]
		if val == nil {
			o.FieldType = ""
		} else {
			o.FieldType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from"].(string); ok {
		o.From = val
	} else {
		val := kv["from"]
		if val == nil {
			o.From = ""
		} else {
			o.From = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["from_string"].(string); ok {
		o.FromString = val
	} else {
		val := kv["from_string"]
		if val == nil {
			o.FromString = ""
		} else {
			o.FromString = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to"].(string); ok {
		o.To = val
	} else {
		val := kv["to"]
		if val == nil {
			o.To = ""
		} else {
			o.To = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["to_string"].(string); ok {
		o.ToString = val
	} else {
		val := kv["to_string"]
		if val == nil {
			o.ToString = ""
		} else {
			o.ToString = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.IssueID)
	args = append(args, o.CreatedAt)
	args = append(args, o.Ordinal)
	args = append(args, o.UserID)
	args = append(args, o.Field)
	args = append(args, o.FieldType)
	args = append(args, o.From)
	args = append(args, o.FromString)
	args = append(args, o.To)
	args = append(args, o.ToString)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func GetChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "work",
		"name":         "Changelog",
		"connect.name": "work.Changelog",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "changelog_id",
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
				"name": "issue_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field",
				"type": "string",
			},
			map[string]interface{}{
				"name": "field_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from",
				"type": "string",
			},
			map[string]interface{}{
				"name": "from_string",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to",
				"type": "string",
			},
			map[string]interface{}{
				"name": "to_string",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetChangelogAvroSchema creates the avro schema for Changelog
func GetChangelogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetChangelogAvroSchemaSpec())
}

// TransformChangelogFunc is a function for transforming Changelog during processing
type TransformChangelogFunc func(input *Changelog) (*Changelog, error)

// NewChangelogPipe creates a pipe for processing Changelog items
func NewChangelogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewChangelogInputStream(input, errors)
	var stream chan Changelog
	if len(transforms) > 0 {
		stream = make(chan Changelog, 1000)
	} else {
		stream = inch
	}
	outdone := NewChangelogOutputStream(output, stream, errors)
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

// NewChangelogInputStreamDir creates a channel for reading Changelog as JSON newlines from a directory of files
func NewChangelogInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/work/changelog\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for changelog")
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewChangelogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewChangelogInputStreamFile creates an channel for reading Changelog as JSON newlines from filename
func NewChangelogInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Changelog)
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
			ch := make(chan Changelog)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewChangelogInputStream(f, errors, transforms...)
}

// NewChangelogInputStream creates an channel for reading Changelog as JSON newlines from stream
func NewChangelogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Changelog, 1000)
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
			var item Changelog
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

// NewChangelogOutputStreamDir will output json newlines from channel and save in dir
func NewChangelogOutputStreamDir(dir string, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	fp := filepath.Join(dir, "/work/changelog\\.json(\\.gz)?$")
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
	return NewChangelogOutputStream(gz, ch, errors, transforms...)
}

// NewChangelogOutputStream will output json newlines from channel to the stream
func NewChangelogOutputStream(stream io.WriteCloser, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
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

// ChangelogSendEvent is an event detail for sending data
type ChangelogSendEvent struct {
	Changelog *Changelog
	headers   map[string]string
	time      time.Time
	key       string
}

var _ datamodel.ModelSendEvent = (*ChangelogSendEvent)(nil)

// Key is the key to use for the message
func (e *ChangelogSendEvent) Key() string {
	if e.key == "" {
		return e.Changelog.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ChangelogSendEvent) Object() datamodel.Model {
	return e.Changelog
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ChangelogSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ChangelogSendEvent) Timestamp() time.Time {
	return e.time
}

// ChangelogSendEventOpts is a function handler for setting opts
type ChangelogSendEventOpts func(o *ChangelogSendEvent)

// WithChangelogSendEventKey sets the key value to a value different than the object ID
func WithChangelogSendEventKey(key string) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		o.key = key
	}
}

// WithChangelogSendEventTimestamp sets the timestamp value
func WithChangelogSendEventTimestamp(tv time.Time) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		o.time = tv
	}
}

// WithChangelogSendEventHeader sets the timestamp value
func WithChangelogSendEventHeader(key, value string) ChangelogSendEventOpts {
	return func(o *ChangelogSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewChangelogSendEvent returns a new ChangelogSendEvent instance
func NewChangelogSendEvent(o *Changelog, opts ...ChangelogSendEventOpts) *ChangelogSendEvent {
	res := &ChangelogSendEvent{
		Changelog: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewChangelogProducer will stream data from the channel
func NewChangelogProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*Changelog); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
					"model":       ChangelogModelName.String(),
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
				errors <- fmt.Errorf("invalid event received. expected an object of type work.Changelog but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewChangelogConsumer will stream data from the topic into the provided channel
func NewChangelogConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object Changelog
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into work.Changelog: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ChangelogReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// ChangelogReceiveEvent is an event detail for receiving data
type ChangelogReceiveEvent struct {
	Changelog *Changelog
	message   event.Message
}

var _ datamodel.ModelReceiveEvent = (*ChangelogReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ChangelogReceiveEvent) Object() datamodel.Model {
	return e.Changelog
}

// Message returns the underlying message data for the event
func (e *ChangelogReceiveEvent) Message() event.Message {
	return e.message
}

// ChangelogProducer implements the datamodel.ModelEventProducer
type ChangelogProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*ChangelogProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ChangelogProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ChangelogProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *Changelog) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ChangelogProducer{
		ch:   ch,
		done: NewChangelogProducer(producer, ch, errors),
	}
}

// NewChangelogProducerChannel returns a channel which can be used for producing Model events
func NewChangelogProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ChangelogProducer{
		ch:   ch,
		done: NewChangelogProducer(producer, ch, errors),
	}
}

// ChangelogConsumer implements the datamodel.ModelEventConsumer
type ChangelogConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*ChangelogConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ChangelogConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ChangelogConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *Changelog) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewChangelogConsumer(consumer, ch, errors)
	return &ChangelogConsumer{
		ch: ch,
	}
}

// NewChangelogConsumerChannel returns a consumer channel which can be used to consume Model events
func NewChangelogConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewChangelogConsumer(consumer, ch, errors)
	return &ChangelogConsumer{
		ch: ch,
	}
}
