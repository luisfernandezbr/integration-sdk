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

	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	number "github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// PullRequestDefaultTopic is the default topic name
const PullRequestDefaultTopic = "sourcecode_PullRequest_topic"

// PullRequestDefaultStream is the default stream name
const PullRequestDefaultStream = "sourcecode_PullRequest_stream"

// PullRequestDefaultTable is the default table name
const PullRequestDefaultTable = "sourcecode_PullRequest"

// PullRequest the pull request for a given repo
type PullRequest struct {
	// built in types

	ID         string `json:"pull_request_id" yaml:"pull_request_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`
	// custom types

	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id"`
	// Title the title of the pull request
	Title string `json:"title" yaml:"title"`
	// Description the description of the pull request
	Description string `json:"description" yaml:"description"`
	// URL the url to the pull request home page
	URL string `json:"url" yaml:"url"`
	// CreatedAt the timestamp in UTC that the pull request was created
	CreatedAt int64 `json:"created_ts" yaml:"created_ts"`
	// MergedAt the timestamp in UTC that the pull request was merged
	MergedAt int64 `json:"merged_ts" yaml:"merged_ts"`
	// ClosedAt the timestamp in UTC that the pull request was closed
	ClosedAt int64 `json:"closed_ts" yaml:"closed_ts"`
	// UpdatedAt the timestamp in UTC that the pull request was closed
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts"`
	// Status the status of the pull request
	Status string `json:"status" yaml:"status"`
	// UserRefID the user ref_id in the source system
	UserRefID string `json:"user_ref_id" yaml:"user_ref_id"`
}

func toPullRequestObject(o interface{}, isavro bool) interface{} {
	if o == nil {
		return nil
	}
	switch v := o.(type) {
	case nil:
		return nil
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		return v
	case *string, *int, *int8, *int16, *int32, *int64, *float32, *float64, *bool:
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *PullRequest:
		return v.ToMap()
	case PullRequest:
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
			arr = append(arr, toPullRequestObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of PullRequest
func (o *PullRequest) String() string {
	return fmt.Sprintf("sourcecode.PullRequest<%s>", o.ID)
}

func (o *PullRequest) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("PullRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *PullRequest) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *PullRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecPullRequest *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *PullRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecPullRequest == nil {
		c, err := CreatePullRequestAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecPullRequest = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecPullRequest.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecPullRequest.BinaryFromNative(nil, native)
	return buf, cachedCodecPullRequest, err
}

// Stringify returns the object in JSON format as a string
func (o *PullRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequest objects are equal
func (o *PullRequest) IsEqual(other *PullRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"pull_request_id": o.GetID(),
		"ref_id":          o.GetRefID(),
		"ref_type":        o.RefType,
		"customer_id":     o.CustomerID,
		"hashcode":        o.Hash(),
		"repo_id":         toPullRequestObject(o.RepoID, isavro),
		"title":           toPullRequestObject(o.Title, isavro),
		"description":     toPullRequestObject(o.Description, isavro),
		"url":             toPullRequestObject(o.URL, isavro),
		"created_ts":      toPullRequestObject(o.CreatedAt, isavro),
		"merged_ts":       toPullRequestObject(o.MergedAt, isavro),
		"closed_ts":       toPullRequestObject(o.ClosedAt, isavro),
		"updated_ts":      toPullRequestObject(o.UpdatedAt, isavro),
		"status":          toPullRequestObject(o.Status, isavro),
		"user_ref_id":     toPullRequestObject(o.UserRefID, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequest) FromMap(kv map[string]interface{}) {
	if val, ok := kv["pull_request_id"].(string); ok {
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
	if val, ok := kv["title"].(string); ok {
		o.Title = val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = ""
		} else {
			o.Title = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		if val == nil {
			o.Description = ""
		} else {
			o.Description = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		if val == nil {
			o.URL = ""
		} else {
			o.URL = fmt.Sprintf("%v", val)
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
	if val, ok := kv["merged_ts"].(int64); ok {
		o.MergedAt = val
	} else {
		val := kv["merged_ts"]
		if val == nil {
			o.MergedAt = number.ToInt64Any(nil)
		} else {
			o.MergedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["closed_ts"].(int64); ok {
		o.ClosedAt = val
	} else {
		val := kv["closed_ts"]
		if val == nil {
			o.ClosedAt = number.ToInt64Any(nil)
		} else {
			o.ClosedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		val := kv["updated_ts"]
		if val == nil {
			o.UpdatedAt = number.ToInt64Any(nil)
		} else {
			o.UpdatedAt = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		if val == nil {
			o.Status = ""
		} else {
			o.Status = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["user_ref_id"].(string); ok {
		o.UserRefID = val
	} else {
		val := kv["user_ref_id"]
		if val == nil {
			o.UserRefID = ""
		} else {
			o.UserRefID = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *PullRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.RepoID)
	args = append(args, o.Title)
	args = append(args, o.Description)
	args = append(args, o.URL)
	args = append(args, o.CreatedAt)
	args = append(args, o.MergedAt)
	args = append(args, o.ClosedAt)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Status)
	args = append(args, o.UserRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreatePullRequestAvroSchemaSpec creates the avro schema specification for PullRequest
func CreatePullRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "PullRequest",
		"connect.name": "sourcecode.PullRequest",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "pull_request_id",
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
				"name": "title",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "url",
				"type": "string",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "merged_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "closed_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "user_ref_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreatePullRequestAvroSchema creates the avro schema for PullRequest
func CreatePullRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreatePullRequestAvroSchemaSpec())
}

// TransformPullRequestFunc is a function for transforming PullRequest during processing
type TransformPullRequestFunc func(input *PullRequest) (*PullRequest, error)

// CreatePullRequestPipe creates a pipe for processing PullRequest items
func CreatePullRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformPullRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreatePullRequestInputStream(input, errors)
	var stream chan PullRequest
	if len(transforms) > 0 {
		stream = make(chan PullRequest, 1000)
	} else {
		stream = inch
	}
	outdone := CreatePullRequestOutputStream(output, stream, errors)
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

// CreatePullRequestInputStreamDir creates a channel for reading PullRequest as JSON newlines from a directory of files
func CreatePullRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/pull_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for pull_request")
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreatePullRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan PullRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreatePullRequestInputStreamFile creates an channel for reading PullRequest as JSON newlines from filename
func CreatePullRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan PullRequest)
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
			ch := make(chan PullRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreatePullRequestInputStream(f, errors, transforms...)
}

// CreatePullRequestInputStream creates an channel for reading PullRequest as JSON newlines from stream
func CreatePullRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformPullRequestFunc) (chan PullRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan PullRequest, 1000)
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
			var item PullRequest
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

// CreatePullRequestOutputStreamDir will output json newlines from channel and save in dir
func CreatePullRequestOutputStreamDir(dir string, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/pull_request\\.json(\\.gz)?$")
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
	return CreatePullRequestOutputStream(gz, ch, errors, transforms...)
}

// CreatePullRequestOutputStream will output json newlines from channel to the stream
func CreatePullRequestOutputStream(stream io.WriteCloser, ch chan PullRequest, errors chan<- error, transforms ...TransformPullRequestFunc) <-chan bool {
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

// CreatePullRequestProducer will stream data from the channel
func CreatePullRequestProducer(producer util.Producer, ch chan PullRequest, errors chan<- error) <-chan bool {
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

// CreatePullRequestConsumer will stream data from the default topic into the provided channel
func CreatePullRequestConsumer(factory util.ConsumerFactory, topic string, ch chan PullRequest, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreatePullRequestConsumerForTopic(factory, PullRequestDefaultTopic, ch, errors)
}

// CreatePullRequestConsumerForTopic will stream data from the topic into the provided channel
func CreatePullRequestConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan PullRequest, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object PullRequest
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into PullRequest: %s", err)
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
