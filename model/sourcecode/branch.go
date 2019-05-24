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
	"strings"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	number "github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/datamodel"
)

// BranchTopic is the default topic name
const BranchTopic datamodel.TopicNameType = "sourcecode_Branch_topic"

// BranchStream is the default stream name
const BranchStream datamodel.TopicNameType = "sourcecode_Branch_stream"

// BranchTable is the default table name
const BranchTable datamodel.TopicNameType = "sourcecode_Branch"

// BranchModelName is the model name
const BranchModelName datamodel.ModelNameType = "sourcecode.Branch"

// Branch git branches
type Branch struct {
	// built in types

	ID         string `json:"branch_id" yaml:"branch_id" faker:"-"`
	RefID      string `json:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" yaml:"hashcode" faker:"-"`
	// custom types

	// Name name of the branch
	Name string `json:"name" yaml:"name" faker:"-"`
	// Default wether is the default branch or not
	Default bool `json:"default" yaml:"default" faker:"-"`
	// Merged wether it has been merged
	Merged bool `json:"merged" yaml:"merged" faker:"-"`
	// MergeCommit commit from the merge
	MergeCommit bool `json:"merge_commit" yaml:"merge_commit" faker:"-"`
	// BranchedFromCommits branched from commits
	BranchedFromCommits []string `json:"branched_from_commits" yaml:"branched_from_commits" faker:"-"`
	// Commits list of commits on this branch
	Commits []string `json:"commits" yaml:"commits" faker:"-"`
	// BehindDefaultCount behind default count
	BehindDefaultCount int64 `json:"behind_default_count" yaml:"behind_default_count" faker:"-"`
	// AheadDefaultCount ahead default count
	AheadDefaultCount int64 `json:"ahead_default_count" yaml:"ahead_default_count" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Branch)(nil)

func toBranchObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toBranchObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toBranchObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toBranchObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toBranchObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Branch:
		return v.ToMap()
	case Branch:
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
			arr = append(arr, toBranchObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Branch
func (o *Branch) String() string {
	return fmt.Sprintf("sourcecode.Branch<%s>", o.ID)
}

func (o *Branch) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Branch) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Branch", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Branch) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Branch) IsMaterialized() bool {
	return true
}

// MaterializedName returns the name of the materialized table
func (o *Branch) MaterializedName() string {
	return "sourcecode_branch"
}

// Clone returns an exact copy of Branch
func (o *Branch) Clone() *Branch {
	c := new(Branch)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Branch) Anon() *Branch {
	c := new(Branch)
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
func (o *Branch) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Branch) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecBranch *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *Branch) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecBranch == nil {
		c, err := CreateBranchAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecBranch = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecBranch.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecBranch.BinaryFromNative(nil, native)
	return buf, cachedCodecBranch, err
}

// Stringify returns the object in JSON format as a string
func (o *Branch) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Branch objects are equal
func (o *Branch) IsEqual(other *Branch) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Branch) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.BranchedFromCommits == nil {
			o.BranchedFromCommits = make([]string, 0)
		}
		if o.Commits == nil {
			o.Commits = make([]string, 0)
		}
	}
	return map[string]interface{}{
		"branch_id":             o.GetID(),
		"ref_id":                o.GetRefID(),
		"ref_type":              o.RefType,
		"customer_id":           o.CustomerID,
		"hashcode":              o.Hash(),
		"name":                  toBranchObject(o.Name, isavro, false, "string"),
		"default":               toBranchObject(o.Default, isavro, false, "boolean"),
		"merged":                toBranchObject(o.Merged, isavro, false, "boolean"),
		"merge_commit":          toBranchObject(o.MergeCommit, isavro, false, "boolean"),
		"branched_from_commits": toBranchObject(o.BranchedFromCommits, isavro, false, "string"),
		"commits":               toBranchObject(o.Commits, isavro, false, "string"),
		"behind_default_count":  toBranchObject(o.BehindDefaultCount, isavro, false, "long"),
		"ahead_default_count":   toBranchObject(o.AheadDefaultCount, isavro, false, "long"),
		"repo_id":               toBranchObject(o.RepoID, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Branch) FromMap(kv map[string]interface{}) {
	if val, ok := kv["branch_id"].(string); ok {
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
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			o.Name = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["default"].(bool); ok {
		o.Default = val
	} else {
		val := kv["default"]
		if val == nil {
			o.Default = number.ToBoolAny(nil)
		} else {
			o.Default = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["merged"].(bool); ok {
		o.Merged = val
	} else {
		val := kv["merged"]
		if val == nil {
			o.Merged = number.ToBoolAny(nil)
		} else {
			o.Merged = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["merge_commit"].(bool); ok {
		o.MergeCommit = val
	} else {
		val := kv["merge_commit"]
		if val == nil {
			o.MergeCommit = number.ToBoolAny(nil)
		} else {
			o.MergeCommit = number.ToBoolAny(val)
		}
	}
	if val := kv["branched_from_commits"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						panic("unsupported type for branched_from_commits field entry: " + reflect.TypeOf(ae).String())
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for branched_from_commits field")
			}
		}
		o.BranchedFromCommits = na
	} else {
		o.BranchedFromCommits = []string{}
	}
	if o.BranchedFromCommits == nil {
		o.BranchedFromCommits = make([]string, 0)
	}
	if val := kv["commits"]; val != nil {
		na := make([]string, 0)
		if a, ok := val.([]string); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(string); ok {
						na = append(na, av)
					} else {
						panic("unsupported type for commits field entry: " + reflect.TypeOf(ae).String())
					}
				}
			} else if s, ok := val.(string); ok {
				for _, sv := range strings.Split(s, ",") {
					na = append(na, strings.TrimSpace(sv))
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for commits field")
			}
		}
		o.Commits = na
	} else {
		o.Commits = []string{}
	}
	if o.Commits == nil {
		o.Commits = make([]string, 0)
	}
	if val, ok := kv["behind_default_count"].(int64); ok {
		o.BehindDefaultCount = val
	} else {
		val := kv["behind_default_count"]
		if val == nil {
			o.BehindDefaultCount = number.ToInt64Any(nil)
		} else {
			o.BehindDefaultCount = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["ahead_default_count"].(int64); ok {
		o.AheadDefaultCount = val
	} else {
		val := kv["ahead_default_count"]
		if val == nil {
			o.AheadDefaultCount = number.ToInt64Any(nil)
		} else {
			o.AheadDefaultCount = number.ToInt64Any(val)
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
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Branch) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Default)
	args = append(args, o.Merged)
	args = append(args, o.MergeCommit)
	args = append(args, o.BranchedFromCommits)
	args = append(args, o.Commits)
	args = append(args, o.BehindDefaultCount)
	args = append(args, o.AheadDefaultCount)
	args = append(args, o.RepoID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateBranchAvroSchemaSpec creates the avro schema specification for Branch
func CreateBranchAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "Branch",
		"connect.name": "sourcecode.Branch",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "branch_id",
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
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "default",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "merged",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "merge_commit",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "branched_from_commits",
				"type": map[string]interface{}{
					"type":  "array",
					"items": "string",
				},
			},
			map[string]interface{}{
				"name": "commits",
				"type": map[string]interface{}{
					"type":  "array",
					"items": "string",
				},
			},
			map[string]interface{}{
				"name": "behind_default_count",
				"type": "long",
			},
			map[string]interface{}{
				"name": "ahead_default_count",
				"type": "long",
			},
			map[string]interface{}{
				"name": "repo_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateBranchAvroSchema creates the avro schema for Branch
func CreateBranchAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateBranchAvroSchemaSpec())
}

// TransformBranchFunc is a function for transforming Branch during processing
type TransformBranchFunc func(input *Branch) (*Branch, error)

// CreateBranchPipe creates a pipe for processing Branch items
func CreateBranchPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformBranchFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateBranchInputStream(input, errors)
	var stream chan Branch
	if len(transforms) > 0 {
		stream = make(chan Branch, 1000)
	} else {
		stream = inch
	}
	outdone := CreateBranchOutputStream(output, stream, errors)
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

// CreateBranchInputStreamDir creates a channel for reading Branch as JSON newlines from a directory of files
func CreateBranchInputStreamDir(dir string, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/branch\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for branch")
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateBranchInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Branch)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateBranchInputStreamFile creates an channel for reading Branch as JSON newlines from filename
func CreateBranchInputStreamFile(filename string, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Branch)
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
			ch := make(chan Branch)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateBranchInputStream(f, errors, transforms...)
}

// CreateBranchInputStream creates an channel for reading Branch as JSON newlines from stream
func CreateBranchInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformBranchFunc) (chan Branch, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Branch, 1000)
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
			var item Branch
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

// CreateBranchOutputStreamDir will output json newlines from channel and save in dir
func CreateBranchOutputStreamDir(dir string, ch chan Branch, errors chan<- error, transforms ...TransformBranchFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/branch\\.json(\\.gz)?$")
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
	return CreateBranchOutputStream(gz, ch, errors, transforms...)
}

// CreateBranchOutputStream will output json newlines from channel to the stream
func CreateBranchOutputStream(stream io.WriteCloser, ch chan Branch, errors chan<- error, transforms ...TransformBranchFunc) <-chan bool {
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

// CreateBranchProducer will stream data from the channel
func CreateBranchProducer(producer datamodel.Producer, ch chan Branch, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			binary, codec, err := item.ToAvroBinary()
			if err != nil {
				errors <- fmt.Errorf("error encoding %s to avro binary data. %v", item.String(), err)
				return
			}
			if err := producer.Send(ctx, codec, []byte(item.ID), binary); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateBranchConsumer will stream data from the default topic into the provided channel
func CreateBranchConsumer(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Branch, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateBranchConsumerForTopic(factory, BranchTopic, ch, errors)
}

// CreateBranchConsumerForTopic will stream data from the topic into the provided channel
func CreateBranchConsumerForTopic(factory datamodel.ConsumerFactory, topic datamodel.TopicNameType, ch chan Branch, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := datamodel.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Branch
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Branch: %s", err)
				}
				ch <- object
				return nil
			},
			OnErrorReceived: func(err error) {
				errors <- err
			},
		}
		consumer, err := factory.CreateConsumer(string(topic), callback)
		if err != nil {
			errors <- err
			return
		}
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}