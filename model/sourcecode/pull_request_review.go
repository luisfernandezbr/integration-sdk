// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// PullRequestReviewTopic is the default topic name
	PullRequestReviewTopic datamodel.TopicNameType = "sourcecode_PullRequestReview_topic"

	// PullRequestReviewStream is the default stream name
	PullRequestReviewStream datamodel.TopicNameType = "sourcecode_PullRequestReview_stream"

	// PullRequestReviewTable is the default table name
	PullRequestReviewTable datamodel.TopicNameType = "sourcecode_PullRequestReview"

	// PullRequestReviewModelName is the model name
	PullRequestReviewModelName datamodel.ModelNameType = "sourcecode.PullRequestReview"
)

const (
	// PullRequestReviewIDColumn is the id column name
	PullRequestReviewIDColumn = "id"
	// PullRequestReviewRefIDColumn is the ref_id column name
	PullRequestReviewRefIDColumn = "ref_id"
	// PullRequestReviewRefTypeColumn is the ref_type column name
	PullRequestReviewRefTypeColumn = "ref_type"
	// PullRequestReviewCustomerIDColumn is the customer_id column name
	PullRequestReviewCustomerIDColumn = "customer_id"
	// PullRequestReviewPullRequestIDColumn is the pull_request_id column name
	PullRequestReviewPullRequestIDColumn = "pull_request_id"
	// PullRequestReviewRepoIDColumn is the repo_id column name
	PullRequestReviewRepoIDColumn = "repo_id"
	// PullRequestReviewStateColumn is the state column name
	PullRequestReviewStateColumn = "state"
	// PullRequestReviewCreatedAtColumn is the created_ts column name
	PullRequestReviewCreatedAtColumn = "created_ts"
	// PullRequestReviewUserRefIDColumn is the user_ref_id column name
	PullRequestReviewUserRefIDColumn = "user_ref_id"
)

// PullRequestReview the review for a given pull request
type PullRequestReview struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	// custom types

	// PullRequestID the pull request this review is associated with
	PullRequestID string `json:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// State the state of the review
	State string `json:"state" bson:"state" yaml:"state" faker:"commit_message"`
	// CreatedAt the timestamp in UTC that the review was created
	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" bson:"user_ref_id" yaml:"user_ref_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestReview)(nil)

func toPullRequestReviewObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toPullRequestReviewObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toPullRequestReviewObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toPullRequestReviewObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toPullRequestReviewObjectNil(isavro, isoptional)
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
	case *PullRequestReview:
		return v.ToMap()
	case PullRequestReview:
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
			arr = append(arr, toPullRequestReviewObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of PullRequestReview
func (o *PullRequestReview) String() string {
	return fmt.Sprintf("sourcecode.PullRequestReview<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestReview) GetTopicName() datamodel.TopicNameType {
	return PullRequestReviewTopic
}

// GetModelName returns the name of the model
func (o *PullRequestReview) GetModelName() datamodel.ModelNameType {
	return PullRequestReviewModelName
}

func (o *PullRequestReview) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequestReview) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequestReview", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestReview) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestReview) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedAt
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
	panic("not sure how to handle the date time format for PullRequestReview")
}

// GetRefID returns the RefID for the object
func (o *PullRequestReview) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestReview) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestReview) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "sourcecode_pullrequestreview",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestReview) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestReview) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestReviewModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestReview) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "created_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *PullRequestReview) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of PullRequestReview
func (o *PullRequestReview) Clone() datamodel.Model {
	c := new(PullRequestReview)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestReview) Anon() datamodel.Model {
	c := new(PullRequestReview)
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
func (o *PullRequestReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestReview) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecPullRequestReview *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *PullRequestReview) GetAvroCodec() *goavro.Codec {
	if cachedCodecPullRequestReview == nil {
		c, err := GetPullRequestReviewAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecPullRequestReview = c
	}
	return cachedCodecPullRequestReview
}

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequestReview) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *PullRequestReview) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequestReview objects are equal
func (o *PullRequestReview) IsEqual(other *PullRequestReview) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestReview) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":              o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"pull_request_id": toPullRequestReviewObject(o.PullRequestID, isavro, false, "string"),
		"repo_id":         toPullRequestReviewObject(o.RepoID, isavro, false, "string"),
		"state":           toPullRequestReviewObject(o.State, isavro, false, "string"),
		"created_ts":      toPullRequestReviewObject(o.CreatedAt, isavro, false, "long"),
		"user_ref_id":     toPullRequestReviewObject(o.UserRefID, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestReview) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	o.setDefaults()
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
	if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = val
	} else {
		val := kv["pull_request_id"]
		if val == nil {
			o.PullRequestID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.PullRequestID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		if val == nil {
			o.RepoID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["state"].(string); ok {
		o.State = val
	} else {
		val := kv["state"]
		if val == nil {
			o.State = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.State = fmt.Sprintf("%v", val)
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
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		val := kv["user_ref_id"]
		if val == nil {
			o.UserRefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UserRefID = fmt.Sprintf("%v", val)
		}
	}
}

// Hash will return a hashcode for the object
func (o *PullRequestReview) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.PullRequestID)
	args = append(args, o.RepoID)
	args = append(args, o.State)
	args = append(args, o.CreatedAt)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetPullRequestReviewAvroSchemaSpec creates the avro schema specification for PullRequestReview
func GetPullRequestReviewAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "PullRequestReview",
		"connect.name": "sourcecode.PullRequestReview",
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
				"name": "pull_request_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "state",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "user_ref_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetPullRequestReviewAvroSchema creates the avro schema for PullRequestReview
func GetPullRequestReviewAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetPullRequestReviewAvroSchemaSpec())
}

// TransformPullRequestReviewFunc is a function for transforming PullRequestReview during processing
type TransformPullRequestReviewFunc func(input *PullRequestReview) (*PullRequestReview, error)

// NewPullRequestReviewPipe creates a pipe for processing PullRequestReview items
func NewPullRequestReviewPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestReviewFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewPullRequestReviewInputStream(input, errors)
	var stream chan PullRequestReview
	if len(transforms) > 0 {
		stream = make(chan PullRequestReview, 1000)
	} else {
		stream = inch
	}
	outdone := NewPullRequestReviewOutputStream(output, stream, errors)
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

// NewPullRequestReviewInputStreamDir creates a channel for reading PullRequestReview as JSON newlines from a directory of files
func NewPullRequestReviewInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestReviewFunc) (chan PullRequestReview, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/pull_request_review\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequestReview)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request_review")
		ch := make(chan PullRequestReview)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewPullRequestReviewInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequestReview)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewPullRequestReviewInputStreamFile creates an channel for reading PullRequestReview as JSON newlines from filename
func NewPullRequestReviewInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestReviewFunc) (chan PullRequestReview, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequestReview)
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
			ch := make(chan PullRequestReview)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewPullRequestReviewInputStream(f, errors, transforms...)
}

// NewPullRequestReviewInputStream creates an channel for reading PullRequestReview as JSON newlines from stream
func NewPullRequestReviewInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestReviewFunc) (chan PullRequestReview, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequestReview, 1000)
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
			var item PullRequestReview
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

// NewPullRequestReviewOutputStreamDir will output json newlines from channel and save in dir
func NewPullRequestReviewOutputStreamDir(dir string, ch chan PullRequestReview, errors chan<- error, transforms ...TransformPullRequestReviewFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/pull_request_review\\.json(\\.gz)?$")
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
	return NewPullRequestReviewOutputStream(gz, ch, errors, transforms...)
}

// NewPullRequestReviewOutputStream will output json newlines from channel to the stream
func NewPullRequestReviewOutputStream(stream io.WriteCloser, ch chan PullRequestReview, errors chan<- error, transforms ...TransformPullRequestReviewFunc) <-chan bool {
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

// PullRequestReviewSendEvent is an event detail for sending data
type PullRequestReviewSendEvent struct {
	PullRequestReview *PullRequestReview
	headers           map[string]string
	time              time.Time
	key               string
}

var _ datamodel.ModelSendEvent = (*PullRequestReviewSendEvent)(nil)

// Key is the key to use for the message
func (e *PullRequestReviewSendEvent) Key() string {
	if e.key == "" {
		return e.PullRequestReview.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *PullRequestReviewSendEvent) Object() datamodel.Model {
	return e.PullRequestReview
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *PullRequestReviewSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *PullRequestReviewSendEvent) Timestamp() time.Time {
	return e.time
}

// PullRequestReviewSendEventOpts is a function handler for setting opts
type PullRequestReviewSendEventOpts func(o *PullRequestReviewSendEvent)

// WithPullRequestReviewSendEventKey sets the key value to a value different than the object ID
func WithPullRequestReviewSendEventKey(key string) PullRequestReviewSendEventOpts {
	return func(o *PullRequestReviewSendEvent) {
		o.key = key
	}
}

// WithPullRequestReviewSendEventTimestamp sets the timestamp value
func WithPullRequestReviewSendEventTimestamp(tv time.Time) PullRequestReviewSendEventOpts {
	return func(o *PullRequestReviewSendEvent) {
		o.time = tv
	}
}

// WithPullRequestReviewSendEventHeader sets the timestamp value
func WithPullRequestReviewSendEventHeader(key, value string) PullRequestReviewSendEventOpts {
	return func(o *PullRequestReviewSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewPullRequestReviewSendEvent returns a new PullRequestReviewSendEvent instance
func NewPullRequestReviewSendEvent(o *PullRequestReview, opts ...PullRequestReviewSendEventOpts) *PullRequestReviewSendEvent {
	res := &PullRequestReviewSendEvent{
		PullRequestReview: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewPullRequestReviewProducer will stream data from the channel
func NewPullRequestReviewProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*PullRequestReview); ok {
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
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.PullRequestReview but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewPullRequestReviewConsumer will stream data from the topic into the provided channel
func NewPullRequestReviewConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object PullRequestReview
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.PullRequestReview: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestReviewReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object PullRequestReview
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &PullRequestReviewReceiveEvent{nil, msg, true}
		},
	})
}

// PullRequestReviewReceiveEvent is an event detail for receiving data
type PullRequestReviewReceiveEvent struct {
	PullRequestReview *PullRequestReview
	message           eventing.Message
	eof               bool
}

var _ datamodel.ModelReceiveEvent = (*PullRequestReviewReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *PullRequestReviewReceiveEvent) Object() datamodel.Model {
	return e.PullRequestReview
}

// Message returns the underlying message data for the event
func (e *PullRequestReviewReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *PullRequestReviewReceiveEvent) EOF() bool {
	return e.eof
}

// PullRequestReviewProducer implements the datamodel.ModelEventProducer
type PullRequestReviewProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*PullRequestReviewProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *PullRequestReviewProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *PullRequestReviewProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *PullRequestReview) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &PullRequestReviewProducer{
		ch:   ch,
		done: NewPullRequestReviewProducer(producer, ch, errors),
	}
}

// NewPullRequestReviewProducerChannel returns a channel which can be used for producing Model events
func NewPullRequestReviewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &PullRequestReviewProducer{
		ch:   ch,
		done: NewPullRequestReviewProducer(producer, ch, errors),
	}
}

// PullRequestReviewConsumer implements the datamodel.ModelEventConsumer
type PullRequestReviewConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*PullRequestReviewConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *PullRequestReviewConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *PullRequestReviewConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *PullRequestReview) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewPullRequestReviewConsumer(consumer, ch, errors)
	return &PullRequestReviewConsumer{
		ch: ch,
	}
}

// NewPullRequestReviewConsumerChannel returns a consumer channel which can be used to consume Model events
func NewPullRequestReviewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewPullRequestReviewConsumer(consumer, ch, errors)
	return &PullRequestReviewConsumer{
		ch: ch,
	}
}
