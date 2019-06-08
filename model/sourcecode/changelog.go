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
	"github.com/pinpt/go-common/event"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

// ChangelogTopic is the default topic name
const ChangelogTopic datamodel.TopicNameType = "sourcecode_Changelog_topic"

// ChangelogStream is the default stream name
const ChangelogStream datamodel.TopicNameType = "sourcecode_Changelog_stream"

// ChangelogTable is the default table name
const ChangelogTable datamodel.TopicNameType = "sourcecode_Changelog"

// ChangelogModelName is the model name
const ChangelogModelName datamodel.ModelNameType = "sourcecode.Changelog"

// Changelog the change log for each commit in a repo
type Changelog struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" bson:"filename" yaml:"filename" faker:"-"`
	// Language the detected language
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// Loc the count of lines in the file
	Loc int64 `json:"loc" bson:"loc" yaml:"loc" faker:"-"`
	// Sloc the count of source lines in the file based on language rules
	Sloc int64 `json:"sloc" bson:"sloc" yaml:"sloc" faker:"-"`
	// Blanks the count of blank lines in the file
	Blanks int64 `json:"blanks" bson:"blanks" yaml:"blanks" faker:"-"`
	// Comments the count of comment lines in the file based on language rules
	Comments int64 `json:"comments" bson:"comments" yaml:"comments" faker:"-"`
	// Complexity the cyclomatic complexity for the change
	Complexity int64 `json:"complexity" bson:"complexity" yaml:"complexity" faker:"-"`
	// DateAt the date of the change
	DateAt int64 `json:"date_ts" bson:"date_ts" yaml:"date_ts" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Ordinal the order of the commit in the commit stream
	Ordinal int64 `json:"ordinal" bson:"ordinal" yaml:"ordinal" faker:"-"`
	// Sha the commit SHA
	Sha string `json:"sha" bson:"sha" yaml:"sha" faker:"-"`
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
	return fmt.Sprintf("sourcecode.Changelog<%s>", o.ID)
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
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Changelog) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Changelog) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Changelog")
}

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Changelog) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Changelog) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "sourcecode_changelog",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Changelog) IsEvented() bool {
	return true
}

// GetTopicConfig returns the topic config object
func (o *Changelog) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "repo_id",
		Timestamp:         "date_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
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
		"id":            o.GetID(),
		"ref_id":        o.GetRefID(),
		"ref_type":      o.RefType,
		"customer_id":   o.CustomerID,
		"hashcode":      o.Hash(),
		"repo_id":       toChangelogObject(o.RepoID, isavro, false, "string"),
		"filename":      toChangelogObject(o.Filename, isavro, false, "string"),
		"language":      toChangelogObject(o.Language, isavro, false, "string"),
		"loc":           toChangelogObject(o.Loc, isavro, false, "long"),
		"sloc":          toChangelogObject(o.Sloc, isavro, false, "long"),
		"blanks":        toChangelogObject(o.Blanks, isavro, false, "long"),
		"comments":      toChangelogObject(o.Comments, isavro, false, "long"),
		"complexity":    toChangelogObject(o.Complexity, isavro, false, "long"),
		"date_ts":       toChangelogObject(o.DateAt, isavro, false, "long"),
		"author_ref_id": toChangelogObject(o.AuthorRefID, isavro, false, "string"),
		"ordinal":       toChangelogObject(o.Ordinal, isavro, false, "long"),
		"sha":           toChangelogObject(o.Sha, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		val := kv["author_ref_id"]
		if val == nil {
			o.AuthorRefID = ""
		} else {
			o.AuthorRefID = fmt.Sprintf("%v", val)
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
	args = append(args, o.RepoID)
	args = append(args, o.Filename)
	args = append(args, o.Language)
	args = append(args, o.Loc)
	args = append(args, o.Sloc)
	args = append(args, o.Blanks)
	args = append(args, o.Comments)
	args = append(args, o.Complexity)
	args = append(args, o.DateAt)
	args = append(args, o.AuthorRefID)
	args = append(args, o.Ordinal)
	args = append(args, o.Sha)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func GetChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Changelog",
		"connect.name": "sourcecode.Changelog",
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
				"name": "repo_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "filename",
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
				"name": "sloc",
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
				"name": "complexity",
				"type": "long",
			},
			map[string]interface{}{
				"name": "date_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "author_ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
			},
			map[string]interface{}{
				"name": "sha",
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
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/changelog\\.json(\\.gz)?$"))
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
	fp := filepath.Join(dir, "/sourcecode/changelog\\.json(\\.gz)?$")
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
				tv := item.Timestamp()
				if tv.IsZero() {
					tv = object.GetTimestamp() // if not provided in the message, use the objects value
				}
				msg := event.Message{
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: item.Timestamp(),
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.Changelog but received on of type %v", reflect.TypeOf(item.Object()))
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
				return fmt.Errorf("error unmarshaling json data into sourcecode.Changelog: %s", err)
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
