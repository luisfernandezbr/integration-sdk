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
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

// CommitActivityTopic is the default topic name
const CommitActivityTopic datamodel.TopicNameType = "sourcecode_CommitActivity_topic"

// CommitActivityStream is the default stream name
const CommitActivityStream datamodel.TopicNameType = "sourcecode_CommitActivity_stream"

// CommitActivityTable is the default table name
const CommitActivityTable datamodel.TopicNameType = "sourcecode_CommitActivity"

// CommitActivityModelName is the model name
const CommitActivityModelName datamodel.ModelNameType = "sourcecode.CommitActivity"

// CommitActivity the file change for a specific author
type CommitActivity struct {
	// built in types

	ID         string `json:"commit_activity_id" yaml:"commit_activity_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// DateAt the timestamp in UTC that the commit of this file was created
	DateAt int64 `json:"date_ts" yaml:"date_ts" faker:"-"`
	// Blanks the number of blank lines in the file
	Blanks int64 `json:"blanks" yaml:"blanks" faker:"-"`
	// Comments the number of comment lines in the file
	Comments int64 `json:"comments" yaml:"comments" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" yaml:"filename" faker:"-"`
	// Sha the sha for this commit
	Sha string `json:"sha" yaml:"sha" faker:"-"`
	// Language the language that was detected for the file
	Language string `json:"language" yaml:"language" faker:"-"`
	// Loc the number of lines in the file
	Loc int64 `json:"loc" yaml:"loc" faker:"-"`
	// Ordinal the order value for the file in the change set
	Ordinal int64 `json:"ordinal" yaml:"ordinal" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
	// Sloc the number of source lines in the file
	Sloc int64 `json:"sloc" yaml:"sloc" faker:"-"`
	// UserID the user who modified this file
	UserID string `json:"user_id" yaml:"user_id" faker:"-"`
	// Complexity the complexity value for the file change
	Complexity int64 `json:"complexity" yaml:"complexity" faker:"-"`
	// ComplexityWeighted the weight of the complexity
	ComplexityWeighted float64 `json:"complexity_weighted" yaml:"complexity_weighted" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CommitActivity)(nil)

func toCommitActivityObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCommitActivityObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toCommitActivityObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCommitActivityObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCommitActivityObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *CommitActivity:
		return v.ToMap()
	case CommitActivity:
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
			arr = append(arr, toCommitActivityObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of CommitActivity
func (o *CommitActivity) String() string {
	return fmt.Sprintf("sourcecode.CommitActivity<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CommitActivity) GetTopicName() datamodel.TopicNameType {
	return CommitActivityTopic
}

// GetModelName returns the name of the model
func (o *CommitActivity) GetModelName() datamodel.ModelNameType {
	return CommitActivityModelName
}

func (o *CommitActivity) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CommitActivity) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CommitActivity", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CommitActivity) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CommitActivity) IsMaterialized() bool {
	return true
}

// MaterializedName returns the name of the materialized table
func (o *CommitActivity) MaterializedName() string {
	return "sourcecode_commitactivity"
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CommitActivity) IsEvented() bool {
	return false

}

// Clone returns an exact copy of CommitActivity
func (o *CommitActivity) Clone() *CommitActivity {
	c := new(CommitActivity)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CommitActivity) Anon() *CommitActivity {
	c := new(CommitActivity)
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
func (o *CommitActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CommitActivity) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecCommitActivity *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *CommitActivity) GetAvroCodec() *goavro.Codec {
	if cachedCodecCommitActivity == nil {
		c, err := GetCommitActivityAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCommitActivity = c
	}
	return cachedCodecCommitActivity
}

// ToAvroBinary returns the data as Avro binary data
func (o *CommitActivity) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CommitActivity) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CommitActivity objects are equal
func (o *CommitActivity) IsEqual(other *CommitActivity) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CommitActivity) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"commit_activity_id":  o.GetID(),
		"ref_id":              o.GetRefID(),
		"ref_type":            o.RefType,
		"customer_id":         o.CustomerID,
		"hashcode":            o.Hash(),
		"date_ts":             toCommitActivityObject(o.DateAt, isavro, false, "long"),
		"blanks":              toCommitActivityObject(o.Blanks, isavro, false, "long"),
		"comments":            toCommitActivityObject(o.Comments, isavro, false, "long"),
		"filename":            toCommitActivityObject(o.Filename, isavro, false, "string"),
		"sha":                 toCommitActivityObject(o.Sha, isavro, false, "string"),
		"language":            toCommitActivityObject(o.Language, isavro, false, "string"),
		"loc":                 toCommitActivityObject(o.Loc, isavro, false, "long"),
		"ordinal":             toCommitActivityObject(o.Ordinal, isavro, false, "long"),
		"repo_id":             toCommitActivityObject(o.RepoID, isavro, false, "string"),
		"sloc":                toCommitActivityObject(o.Sloc, isavro, false, "long"),
		"user_id":             toCommitActivityObject(o.UserID, isavro, false, "string"),
		"complexity":          toCommitActivityObject(o.Complexity, isavro, false, "long"),
		"complexity_weighted": toCommitActivityObject(o.ComplexityWeighted, isavro, false, "float"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitActivity) FromMap(kv map[string]interface{}) {
	if val, ok := kv["commit_activity_id"].(string); ok {
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
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		val := kv["blanks"]
		if val == nil {
			o.Blanks = number.ToInt64Any(nil)
		} else {
			o.Blanks = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		val := kv["comments"]
		if val == nil {
			o.Comments = number.ToInt64Any(nil)
		} else {
			o.Comments = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		val := kv["filename"]
		if val == nil {
			o.Filename = ""
		} else {
			o.Filename = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		val := kv["sha"]
		if val == nil {
			o.Sha = ""
		} else {
			o.Sha = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		if val == nil {
			o.Language = ""
		} else {
			o.Language = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		val := kv["loc"]
		if val == nil {
			o.Loc = number.ToInt64Any(nil)
		} else {
			o.Loc = number.ToInt64Any(val)
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		if val == nil {
			o.RepoID = ""
		} else {
			o.RepoID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		val := kv["sloc"]
		if val == nil {
			o.Sloc = number.ToInt64Any(nil)
		} else {
			o.Sloc = number.ToInt64Any(val)
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
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		val := kv["complexity"]
		if val == nil {
			o.Complexity = number.ToInt64Any(nil)
		} else {
			o.Complexity = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["complexity_weighted"].(float64); ok {
		o.ComplexityWeighted = val
	} else {
		val := kv["complexity_weighted"]
		if val == nil {
			o.ComplexityWeighted = number.ToFloat64Any(nil)
		} else {
			o.ComplexityWeighted = number.ToFloat64Any(val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CommitActivity) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.DateAt)
	args = append(args, o.Blanks)
	args = append(args, o.Comments)
	args = append(args, o.Filename)
	args = append(args, o.Sha)
	args = append(args, o.Language)
	args = append(args, o.Loc)
	args = append(args, o.Ordinal)
	args = append(args, o.RepoID)
	args = append(args, o.Sloc)
	args = append(args, o.UserID)
	args = append(args, o.Complexity)
	args = append(args, o.ComplexityWeighted)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetCommitActivityAvroSchemaSpec creates the avro schema specification for CommitActivity
func GetCommitActivityAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "CommitActivity",
		"connect.name": "sourcecode.CommitActivity",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "commit_activity_id",
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
				"name": "blanks",
				"type": "long",
			},
			map[string]interface{}{
				"name": "comments",
				"type": "long",
			},
			map[string]interface{}{
				"name": "filename",
				"type": "string",
			},
			map[string]interface{}{
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name": "loc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "sloc",
				"type": "long",
			},
			map[string]interface{}{
				"name": "user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "complexity_weighted",
				"type": "float",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetCommitActivityAvroSchema creates the avro schema for CommitActivity
func GetCommitActivityAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCommitActivityAvroSchemaSpec())
}

// TransformCommitActivityFunc is a function for transforming CommitActivity during processing
type TransformCommitActivityFunc func(input *CommitActivity) (*CommitActivity, error)

// NewCommitActivityPipe creates a pipe for processing CommitActivity items
func NewCommitActivityPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitActivityFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCommitActivityInputStream(input, errors)
	var stream chan CommitActivity
	if len(transforms) > 0 {
		stream = make(chan CommitActivity, 1000)
	} else {
		stream = inch
	}
	outdone := NewCommitActivityOutputStream(output, stream, errors)
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

// NewCommitActivityInputStreamDir creates a channel for reading CommitActivity as JSON newlines from a directory of files
func NewCommitActivityInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitActivityFunc) (chan CommitActivity, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/commit_activity\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CommitActivity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit_activity")
		ch := make(chan CommitActivity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCommitActivityInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CommitActivity)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCommitActivityInputStreamFile creates an channel for reading CommitActivity as JSON newlines from filename
func NewCommitActivityInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitActivityFunc) (chan CommitActivity, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CommitActivity)
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
			ch := make(chan CommitActivity)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCommitActivityInputStream(f, errors, transforms...)
}

// NewCommitActivityInputStream creates an channel for reading CommitActivity as JSON newlines from stream
func NewCommitActivityInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitActivityFunc) (chan CommitActivity, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CommitActivity, 1000)
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
			var item CommitActivity
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

// NewCommitActivityOutputStreamDir will output json newlines from channel and save in dir
func NewCommitActivityOutputStreamDir(dir string, ch chan CommitActivity, errors chan<- error, transforms ...TransformCommitActivityFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/commit_activity\\.json(\\.gz)?$")
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
	return NewCommitActivityOutputStream(gz, ch, errors, transforms...)
}

// NewCommitActivityOutputStream will output json newlines from channel to the stream
func NewCommitActivityOutputStream(stream io.WriteCloser, ch chan CommitActivity, errors chan<- error, transforms ...TransformCommitActivityFunc) <-chan bool {
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

// CommitActivitySendEvent is an event detail for sending data
type CommitActivitySendEvent struct {
	CommitActivity *CommitActivity
	headers        map[string]string
	time           time.Time
	key            string
}

var _ datamodel.ModelSendEvent = (*CommitActivitySendEvent)(nil)

// Key is the key to use for the message
func (e *CommitActivitySendEvent) Key() string {
	if e.key == "" {
		return e.CommitActivity.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CommitActivitySendEvent) Object() datamodel.Model {
	return e.CommitActivity
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CommitActivitySendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CommitActivitySendEvent) Timestamp() time.Time {
	return e.time
}

// CommitActivitySendEventOpts is a function handler for setting opts
type CommitActivitySendEventOpts func(o *CommitActivitySendEvent)

// WithCommitActivitySendEventKey sets the key value to a value different than the object ID
func WithCommitActivitySendEventKey(key string) CommitActivitySendEventOpts {
	return func(o *CommitActivitySendEvent) {
		o.key = key
	}
}

// WithCommitActivitySendEventTimestamp sets the timestamp value
func WithCommitActivitySendEventTimestamp(tv time.Time) CommitActivitySendEventOpts {
	return func(o *CommitActivitySendEvent) {
		o.time = tv
	}
}

// WithCommitActivitySendEventHeader sets the timestamp value
func WithCommitActivitySendEventHeader(key, value string) CommitActivitySendEventOpts {
	return func(o *CommitActivitySendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCommitActivitySendEvent returns a new CommitActivitySendEvent instance
func NewCommitActivitySendEvent(o *CommitActivity, opts ...CommitActivitySendEventOpts) *CommitActivitySendEvent {
	res := &CommitActivitySendEvent{
		CommitActivity: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCommitActivityProducer will stream data from the channel
func NewCommitActivityProducer(producer event.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*CommitActivity); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{
					"customer_id": object.CustomerID,
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
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.CommitActivity but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewCommitActivityConsumer will stream data from the topic into the provided channel
func NewCommitActivityConsumer(consumer event.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(event.ConsumerCallback{
		OnDataReceived: func(msg event.Message) error {
			var object CommitActivity
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into sourcecode.CommitActivity: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CommitActivityReceiveEvent{&object, msg}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
	})
}

// CommitActivityReceiveEvent is an event detail for receiving data
type CommitActivityReceiveEvent struct {
	CommitActivity *CommitActivity
	message        event.Message
}

var _ datamodel.ModelReceiveEvent = (*CommitActivityReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CommitActivityReceiveEvent) Object() datamodel.Model {
	return e.CommitActivity
}

// Message returns the underlying message data for the event
func (e *CommitActivityReceiveEvent) Message() event.Message {
	return e.message
}

// CommitActivityProducer implements the datamodel.ModelEventProducer
type CommitActivityProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*CommitActivityProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CommitActivityProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CommitActivityProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *CommitActivity) NewProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitActivityProducer{
		ch:   ch,
		done: NewCommitActivityProducer(producer, ch, errors),
	}
}

// NewCommitActivityProducerChannel returns a channel which can be used for producing Model events
func NewCommitActivityProducerChannel(producer event.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CommitActivityProducer{
		ch:   ch,
		done: NewCommitActivityProducer(producer, ch, errors),
	}
}

// CommitActivityConsumer implements the datamodel.ModelEventConsumer
type CommitActivityConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*CommitActivityConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CommitActivityConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CommitActivityConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *CommitActivity) NewConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitActivityConsumer(consumer, ch, errors)
	return &CommitActivityConsumer{
		ch: ch,
	}
}

// NewCommitActivityConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCommitActivityConsumerChannel(consumer event.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCommitActivityConsumer(consumer, ch, errors)
	return &CommitActivityConsumer{
		ch: ch,
	}
}
