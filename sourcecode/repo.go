// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// RepoTopic is the default topic name
	RepoTopic datamodel.TopicNameType = "sourcecode_Repo_topic"

	// RepoStream is the default stream name
	RepoStream datamodel.TopicNameType = "sourcecode_Repo_stream"

	// RepoTable is the default table name
	RepoTable datamodel.TopicNameType = "sourcecode_repo"

	// RepoModelName is the model name
	RepoModelName datamodel.ModelNameType = "sourcecode.Repo"
)

const (
	// RepoActiveColumn is the active column name
	RepoActiveColumn = "active"
	// RepoCustomerIDColumn is the customer_id column name
	RepoCustomerIDColumn = "customer_id"
	// RepoDefaultBranchColumn is the default_branch column name
	RepoDefaultBranchColumn = "default_branch"
	// RepoDescriptionColumn is the description column name
	RepoDescriptionColumn = "description"
	// RepoIDColumn is the id column name
	RepoIDColumn = "id"
	// RepoLanguageColumn is the language column name
	RepoLanguageColumn = "language"
	// RepoNameColumn is the name column name
	RepoNameColumn = "name"
	// RepoRefIDColumn is the ref_id column name
	RepoRefIDColumn = "ref_id"
	// RepoRefTypeColumn is the ref_type column name
	RepoRefTypeColumn = "ref_type"
	// RepoUpdatedAtColumn is the updated_ts column name
	RepoUpdatedAtColumn = "updated_ts"
	// RepoURLColumn is the url column name
	RepoURLColumn = "url"
)

// Repo the repo holds source code
type Repo struct {
	// Active the status of the repo
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DefaultBranch the repo default branch
	DefaultBranch string `json:"default_branch" bson:"default_branch" yaml:"default_branch" faker:"-"`
	// Description brief explanation of the repo
	Description string `json:"description" bson:"description" yaml:"description" faker:"sentence"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Language the programming language the source code is primarily written in
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// Name the name of the repo
	Name string `json:"name" bson:"name" yaml:"name" faker:"repo"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the repo home page
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Repo)(nil)

func toRepoObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *Repo:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of Repo
func (o *Repo) String() string {
	return fmt.Sprintf("sourcecode.Repo<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Repo) GetTopicName() datamodel.TopicNameType {
	return RepoTopic
}

// GetModelName returns the name of the model
func (o *Repo) GetModelName() datamodel.ModelNameType {
	return RepoModelName
}

// NewRepoID provides a template for generating an ID field for Repo
func NewRepoID(customerID string, refType string, refID string) string {
	return hash.Values("Repo", customerID, refType, refID)
}

func (o *Repo) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Repo", o.CustomerID, o.RefType, o.GetRefID())
	}

	{
		v := "master"

		o.DefaultBranch = v

	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Repo) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Repo) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Repo) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedAt
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
	panic("not sure how to handle the date time format for Repo")
}

// GetRefID returns the RefID for the object
func (o *Repo) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Repo) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Repo) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Repo) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Repo) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Repo) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *Repo) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *Repo) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Repo
func (o *Repo) Clone() datamodel.Model {
	c := new(Repo)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Repo) Anon() datamodel.Model {
	c := new(Repo)
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *Repo) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *Repo) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Repo) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Repo) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

var cachedCodecRepo *goavro.Codec
var cachedCodecRepoLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *Repo) GetAvroCodec() *goavro.Codec {
	cachedCodecRepoLock.Lock()
	if cachedCodecRepo == nil {
		c, err := GetRepoAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecRepo = c
	}
	cachedCodecRepoLock.Unlock()
	return cachedCodecRepo
}

// ToAvroBinary returns the data as Avro binary data
func (o *Repo) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *Repo) FromAvroBinary(value []byte) error {
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
func (o *Repo) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Repo objects are equal
func (o *Repo) IsEqual(other *Repo) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Repo) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"active":         toRepoObject(o.Active, isavro, false, "boolean"),
		"customer_id":    toRepoObject(o.CustomerID, isavro, false, "string"),
		"default_branch": toRepoObject(o.DefaultBranch, isavro, false, "string"),
		"description":    toRepoObject(o.Description, isavro, false, "string"),
		"id":             toRepoObject(o.ID, isavro, false, "string"),
		"language":       toRepoObject(o.Language, isavro, false, "string"),
		"name":           toRepoObject(o.Name, isavro, false, "string"),
		"ref_id":         toRepoObject(o.RefID, isavro, false, "string"),
		"ref_type":       toRepoObject(o.RefType, isavro, false, "string"),
		"updated_ts":     toRepoObject(o.UpdatedAt, isavro, false, "long"),
		"url":            toRepoObject(o.URL, isavro, false, "string"),
		"hashcode":       toRepoObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Repo) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
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

	if val, ok := kv["default_branch"].(string); ok {
		o.DefaultBranch = val
	} else {
		if val, ok := kv["default_branch"]; ok {
			if val == nil {
				o.DefaultBranch = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.DefaultBranch = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Description = fmt.Sprintf("%v", val)
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

	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Language = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
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

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Repo) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CustomerID)
	args = append(args, o.DefaultBranch)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Language)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetRepoAvroSchemaSpec creates the avro schema specification for Repo
func GetRepoAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "Repo",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "default_branch",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
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
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Repo) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetRepoAvroSchema creates the avro schema for Repo
func GetRepoAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetRepoAvroSchemaSpec())
}

// TransformRepoFunc is a function for transforming Repo during processing
type TransformRepoFunc func(input *Repo) (*Repo, error)

// NewRepoPipe creates a pipe for processing Repo items
func NewRepoPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewRepoInputStream(input, errors)
	var stream chan Repo
	if len(transforms) > 0 {
		stream = make(chan Repo, 1000)
	} else {
		stream = inch
	}
	outdone := NewRepoOutputStream(output, stream, errors)
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

// NewRepoInputStreamDir creates a channel for reading Repo as JSON newlines from a directory of files
func NewRepoInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/repo\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for repo")
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewRepoInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewRepoInputStreamFile creates an channel for reading Repo as JSON newlines from filename
func NewRepoInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Repo)
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
			ch := make(chan Repo)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewRepoInputStream(f, errors, transforms...)
}

// NewRepoInputStream creates an channel for reading Repo as JSON newlines from stream
func NewRepoInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Repo, 1000)
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
			var item Repo
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

// NewRepoOutputStreamDir will output json newlines from channel and save in dir
func NewRepoOutputStreamDir(dir string, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/repo\\.json(\\.gz)?$")
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
	return NewRepoOutputStream(gz, ch, errors, transforms...)
}

// NewRepoOutputStream will output json newlines from channel to the stream
func NewRepoOutputStream(stream io.WriteCloser, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
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

// RepoSendEvent is an event detail for sending data
type RepoSendEvent struct {
	Repo    *Repo
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*RepoSendEvent)(nil)

// Key is the key to use for the message
func (e *RepoSendEvent) Key() string {
	if e.key == "" {
		return e.Repo.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *RepoSendEvent) Object() datamodel.Model {
	return e.Repo
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *RepoSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *RepoSendEvent) Timestamp() time.Time {
	return e.time
}

// RepoSendEventOpts is a function handler for setting opts
type RepoSendEventOpts func(o *RepoSendEvent)

// WithRepoSendEventKey sets the key value to a value different than the object ID
func WithRepoSendEventKey(key string) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		o.key = key
	}
}

// WithRepoSendEventTimestamp sets the timestamp value
func WithRepoSendEventTimestamp(tv time.Time) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		o.time = tv
	}
}

// WithRepoSendEventHeader sets the timestamp value
func WithRepoSendEventHeader(key, value string) RepoSendEventOpts {
	return func(o *RepoSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewRepoSendEvent returns a new RepoSendEvent instance
func NewRepoSendEvent(o *Repo, opts ...RepoSendEventOpts) *RepoSendEvent {
	res := &RepoSendEvent{
		Repo: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewRepoProducer will stream data from the channel
func NewRepoProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
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
				if object, ok := item.Object().(*Repo); ok {
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
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Repo but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewRepoConsumer will stream data from the topic into the provided channel
func NewRepoConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object Repo
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.Repo: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into sourcecode.Repo: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.Repo")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &RepoReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object Repo
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &RepoReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// RepoReceiveEvent is an event detail for receiving data
type RepoReceiveEvent struct {
	Repo    *Repo
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*RepoReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *RepoReceiveEvent) Object() datamodel.Model {
	return e.Repo
}

// Message returns the underlying message data for the event
func (e *RepoReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *RepoReceiveEvent) EOF() bool {
	return e.eof
}

// RepoProducer implements the datamodel.ModelEventProducer
type RepoProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*RepoProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *RepoProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *RepoProducer) Close() error {
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
func (o *Repo) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *Repo) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoProducer(newctx, producer, ch, errors, empty),
	}
}

// NewRepoProducerChannel returns a channel which can be used for producing Model events
func NewRepoProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewRepoProducerChannelSize(producer, 0, errors)
}

// NewRepoProducerChannelSize returns a channel which can be used for producing Model events
func NewRepoProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoProducer(newctx, producer, ch, errors, empty),
	}
}

// RepoConsumer implements the datamodel.ModelEventConsumer
type RepoConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*RepoConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *RepoConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *RepoConsumer) Close() error {
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
func (o *Repo) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoConsumer{
		ch:       ch,
		callback: NewRepoConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewRepoConsumerChannel returns a consumer channel which can be used to consume Model events
func NewRepoConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoConsumer{
		ch:       ch,
		callback: NewRepoConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
