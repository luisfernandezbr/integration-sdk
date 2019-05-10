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
	"regexp"
	"strings"

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// BranchDefaultTopic is the default topic name
const BranchDefaultTopic = "sourcecode_Branch_topic"

// BranchDefaultStream is the default stream name
const BranchDefaultStream = "sourcecode_Branch_stream"

// BranchDefaultTable is the default table name
const BranchDefaultTable = "sourcecode_Branch"

// Branch git repo's branch
type Branch struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Name                string   `json:"name" yaml:"name"`
	IsDefault           bool     `json:"is_default" yaml:"is_default"`
	IsMerged            bool     `json:"is_merged" yaml:"is_merged"`
	MergeCommit         bool     `json:"merge_commit" yaml:"merge_commit"`
	BranchedFromCommits []string `json:"branched_from_commits" yaml:"branched_from_commits"`
	Commits             []string `json:"commits" yaml:"commits"`
	BehindDefaultCount  int64    `json:"behind_default_count" yaml:"behind_default_count"`
	AheadDefaultCount   int64    `json:"ahead_default_count" yaml:"ahead_default_count"`
	RepoID              string   `json:"repo_id" yaml:"repo_id"`
}

// String returns a string representation of Branch
func (o *Branch) String() string {
	return fmt.Sprintf("sourcecode.v1.Branch<%s>", o.ID)
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

// Stringify returns the object in JSON format as a string
func (o *Branch) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Branch objects are equal
func (o *Branch) IsEqual(other *Branch) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Branch) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                    o.GetID(),
		"ref_id":                o.GetRefID(),
		"ref_type":              o.RefType,
		"customer_id":           o.CustomerID,
		"hashcode":              o.Hash(),
		"name":                  o.Name,
		"is_default":            o.IsDefault,
		"is_merged":             o.IsMerged,
		"merge_commit":          o.MergeCommit,
		"branched_from_commits": o.BranchedFromCommits,
		"commits":               o.Commits,
		"behind_default_count":  o.BehindDefaultCount,
		"ahead_default_count":   o.AheadDefaultCount,
		"repo_id":               o.RepoID,
	}
}

// FromMap attempts to load data into object from a map
func (o *Branch) FromMap(kv map[string]interface{}) {
	if val, ok := kv["id"].(string); ok {
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
	if val, ok := kv["is_default"].(bool); ok {
		o.IsDefault = val
	} else {
		val := kv["is_default"]
		if val == nil {
			o.IsDefault = number.ToBoolAny(nil)
		} else {
			o.IsDefault = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["is_merged"].(bool); ok {
		o.IsMerged = val
	} else {
		val := kv["is_merged"]
		if val == nil {
			o.IsMerged = number.ToBoolAny(nil)
		} else {
			o.IsMerged = number.ToBoolAny(val)
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
		if a, ok := val.([]interface{}); ok {
			var arr []string
			for _, b := range a {
				arr = append(arr, b.(string))
			}
			o.BranchedFromCommits = arr
		} else {
			o.BranchedFromCommits = []string{}
		}
	} else {
		o.BranchedFromCommits = []string{}
	}
	if val := kv["commits"]; val != nil {
		if a, ok := val.([]interface{}); ok {
			var arr []string
			for _, b := range a {
				arr = append(arr, b.(string))
			}
			o.Commits = arr
		} else {
			o.Commits = []string{}
		}
	} else {
		o.Commits = []string{}
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
	args = append(args, o.Name)
	args = append(args, o.IsDefault)
	args = append(args, o.IsMerged)
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
		"namespace":    "sourcecode.v1",
		"name":         "Branch",
		"connect.name": "sourcecode.v1.Branch",
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
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "is_default",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "is_merged",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "merge_commit",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "branched_from_commits",
				"type": map[string]string{"type": "array", "items": "string"},
			},
			map[string]interface{}{
				"name": "commits",
				"type": map[string]string{"type": "array", "items": "string"},
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

// CreateBranchKQLStreamSQL creates KQL Stream SQL for Branch
func CreateBranchKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", BranchDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", BranchDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// CreateBranchKQLTableSQL creates KQL Table SQL for Branch
func CreateBranchKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", BranchDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", BranchDefaultTopic))
	builder.WriteString(");")
	return builder.String()
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
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/branch\\.json(\\.gz)?$"))
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
	fp := filepath.Join(dir, "/sourcecode/v1/branch\\.json(\\.gz)?$")
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
func CreateBranchProducer(producer util.Producer, ch chan Branch, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateBranchAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item.String(), err)
			}
		}
	}()
	return done
}

// CreateBranchConsumer will stream data from the default topic into the provided channel
func CreateBranchConsumer(factory util.ConsumerFactory, topic string, ch chan Branch, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateBranchConsumerForTopic(factory, BranchDefaultTopic, ch, errors)
}

// CreateBranchConsumerForTopic will stream data from the topic into the provided channel
func CreateBranchConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan Branch, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
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
		consumer, err := factory.CreateConsumer(topic, callback)
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
