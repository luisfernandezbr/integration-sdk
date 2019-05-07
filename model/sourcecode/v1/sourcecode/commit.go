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

	cluster "github.com/bsm/sarama-cluster"
	kafka "github.com/dangkaka/go-kafka-avro"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/integration-sdk/util"
)

// CommitDefaultTopic is the default topic name
const CommitDefaultTopic = "sourcecode.Commit.topic"

// CommitDefaultStream is the default stream name
const CommitDefaultStream = "sourcecode.Commit.topicstream"

// CommitDefaultTable is the default table name
const CommitDefaultTable = "sourcecode.Commit"

// Commit the commit is a specific change in a repo
type Commit struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	RepoID       string `json:"repo_id" yaml:"repo_id"`
	Sha          string `json:"sha" yaml:"sha"`
	Message      string `json:"message" yaml:"message"`
	URL          string `json:"url" yaml:"url"`
	CreatedAt    int64  `json:"created_ts" yaml:"created_ts"`
	Branch       string `json:"branch" yaml:"branch"`
	Additions    int64  `json:"additions" yaml:"additions"`
	Deletions    int64  `json:"deletions" yaml:"deletions"`
	FilesChanged int64  `json:"files_changed" yaml:"files_changed"`
	AuthorRefID  string `json:"author_ref_id" yaml:"author_ref_id"`
	Ordinal      int64  `json:"ordinal" yaml:"ordinal"`
}

// String returns a string representation of Commit
func (o *Commit) String() string {
	return fmt.Sprintf("sourcecode.v1.Commit<%s>", o.ID)
}

func (o *Commit) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Commit) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Commit", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Commit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Commit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Commit) UnmarshalJSON(data []byte) error {
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
func (o *Commit) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Commit objects are equal
func (o *Commit) IsEqual(other *Commit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Commit) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            o.GetID(),
		"ref_id":        o.GetRefID(),
		"ref_type":      o.RefType,
		"customer_id":   o.CustomerID,
		"hashcode":      o.Hash(),
		"repo_id":       o.RepoID,
		"sha":           o.Sha,
		"message":       o.Message,
		"url":           o.URL,
		"created_ts":    o.CreatedAt,
		"branch":        o.Branch,
		"additions":     o.Additions,
		"deletions":     o.Deletions,
		"files_changed": o.FilesChanged,
		"author_ref_id": o.AuthorRefID,
		"ordinal":       o.Ordinal,
	}
}

// FromMap attempts to load data into object from a map
func (o *Commit) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		val := kv["repo_id"]
		o.RepoID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		val := kv["sha"]
		o.Sha = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		o.Message = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		o.URL = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		val := kv["created_ts"]
		o.CreatedAt = number.ToInt64Any(val)
	}
	if val, ok := kv["branch"].(string); ok {
		o.Branch = val
	} else {
		val := kv["branch"]
		o.Branch = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		val := kv["additions"]
		o.Additions = number.ToInt64Any(val)
	}
	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		val := kv["deletions"]
		o.Deletions = number.ToInt64Any(val)
	}
	if val, ok := kv["files_changed"].(int64); ok {
		o.FilesChanged = val
	} else {
		val := kv["files_changed"]
		o.FilesChanged = number.ToInt64Any(val)
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		val := kv["author_ref_id"]
		o.AuthorRefID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		o.Ordinal = number.ToInt64Any(val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Commit) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.RepoID)
		args = append(args, o.Sha)
		args = append(args, o.Message)
		args = append(args, o.URL)
		args = append(args, o.CreatedAt)
		args = append(args, o.Branch)
		args = append(args, o.Additions)
		args = append(args, o.Deletions)
		args = append(args, o.FilesChanged)
		args = append(args, o.AuthorRefID)
		args = append(args, o.Ordinal)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateCommitAvroSchemaSpec creates the avro schema specification for Commit
func CreateCommitAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode.v1",
		"name":         "Commit",
		"connect.name": "sourcecode.v1.Commit",
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
				"name": "sha",
				"type": "string",
			},
			map[string]interface{}{
				"name": "message",
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
				"name": "branch",
				"type": "string",
			},
			map[string]interface{}{
				"name": "additions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "deletions",
				"type": "long",
			},
			map[string]interface{}{
				"name": "files_changed",
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
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateCommitAvroSchema creates the avro schema for Commit
func CreateCommitAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCommitAvroSchemaSpec())
}

// CreateCommitKQLStreamSQL creates KQL Stream SQL for Commit
func CreateCommitKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", CommitDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", CommitDefaultTopic))
	builder.WriteString(", TIMESTAMP='created_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateCommitKQLTableSQL creates KQL Table SQL for Commit
func CreateCommitKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", CommitDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", CommitDefaultTopic))
	builder.WriteString(", TIMESTAMP='created_ts'")
	builder.WriteString(");")
	return builder.String()
}

// TransformCommitFunc is a function for transforming Commit during processing
type TransformCommitFunc func(input *Commit) (*Commit, error)

// CreateCommitPipe creates a pipe for processing Commit items
func CreateCommitPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCommitInputStream(input, errors)
	var stream chan Commit
	if len(transforms) > 0 {
		stream = make(chan Commit, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCommitOutputStream(output, stream, errors)
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

// CreateCommitInputStreamDir creates a channel for reading Commit as JSON newlines from a directory of files
func CreateCommitInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/commit\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit")
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCommitInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Commit)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCommitInputStreamFile creates an channel for reading Commit as JSON newlines from filename
func CreateCommitInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Commit)
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
			ch := make(chan Commit)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCommitInputStream(f, errors, transforms...)
}

// CreateCommitInputStream creates an channel for reading Commit as JSON newlines from stream
func CreateCommitInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFunc) (chan Commit, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Commit, 1000)
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
			var item Commit
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

// CreateCommitOutputStreamDir will output json newlines from channel and save in dir
func CreateCommitOutputStreamDir(dir string, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/v1/commit\\.json(\\.gz)?$")
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
	return CreateCommitOutputStream(gz, ch, errors, transforms...)
}

// CreateCommitOutputStream will output json newlines from channel to the stream
func CreateCommitOutputStream(stream io.WriteCloser, ch chan Commit, errors chan<- error, transforms ...TransformCommitFunc) <-chan bool {
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

// CreateCommitProducer will stream data from the channel
func CreateCommitProducer(producer *util.KafkaProducer, ch chan Commit, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateCommitAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item, err)
			}
		}
	}()
	return done
}

// CreateCommitConsumer will stream data from the default topic into the provided channel
func CreateCommitConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Commit, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateCommitConsumerForTopic(kafkaServers, schemaRegistryServers, CommitDefaultTopic, consumerGroupID, ch, errors)
}

// CreateCommitConsumerForTopic will stream data from the topic into the provided channel
func CreateCommitConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Commit, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object Commit
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into Commit: %s", err)
					return
				}
				ch <- object
			},
			OnError: func(err error) {
				errors <- err
			},
			OnNotification: func(notification *cluster.Notification) {
			},
		}
		consumer, err := kafka.NewAvroConsumer(kafkaServers, schemaRegistryServers, topic, consumerGroupID, consumerCallbacks)
		if err != nil {
			errors <- err
			return
		}
		go consumer.Consume()
		select {
		case <-closed:
			consumer.Close()
		}
	}()
	return done, closed
}
