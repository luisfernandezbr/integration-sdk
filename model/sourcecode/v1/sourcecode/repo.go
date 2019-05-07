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
	"github.com/pinpt/integration-sdk/util"
)

// RepoDefaultTopic is the default topic name
const RepoDefaultTopic = "sourcecode.Repo.topic"

// RepoDefaultStream is the default stream name
const RepoDefaultStream = "sourcecode.Repo.topicstream"

// RepoDefaultTable is the default table name
const RepoDefaultTable = "sourcecode.Repo"

// Repo the repo holds source code
type Repo struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

// String returns a string representation of Repo
func (o *Repo) String() string {
	return fmt.Sprintf("sourcecode.v1.Repo<%s>", o.ID)
}

func (o *Repo) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Repo) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Repo", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Repo) GetRefID() string {
	return o.RefID
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
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *Repo) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Repo objects are equal
func (o *Repo) IsEqual(other *Repo) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Repo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        o.Name,
		"url":         o.URL,
	}
}

// FromMap attempts to load data into object from a map
func (o *Repo) FromMap(kv map[string]interface{}) {
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
		o.Name = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		val := kv["url"]
		o.URL = fmt.Sprintf("%v", val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Repo) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.Name)
		args = append(args, o.URL)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateRepoAvroSchemaSpec creates the avro schema specification for Repo
func CreateRepoAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode.v1",
		"name":         "Repo",
		"connect.name": "sourcecode.v1.Repo",
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
				"name": "url",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateRepoAvroSchema creates the avro schema for Repo
func CreateRepoAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateRepoAvroSchemaSpec())
}

// CreateRepoKQLStreamSQL creates KQL Stream SQL for Repo
func CreateRepoKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", RepoDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", RepoDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// CreateRepoKQLTableSQL creates KQL Table SQL for Repo
func CreateRepoKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", RepoDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", RepoDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// TransformRepoFunc is a function for transforming Repo during processing
type TransformRepoFunc func(input *Repo) (*Repo, error)

// CreateRepoPipe creates a pipe for processing Repo items
func CreateRepoPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateRepoInputStream(input, errors)
	var stream chan Repo
	if len(transforms) > 0 {
		stream = make(chan Repo, 1000)
	} else {
		stream = inch
	}
	outdone := CreateRepoOutputStream(output, stream, errors)
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

// CreateRepoInputStreamDir creates a channel for reading Repo as JSON newlines from a directory of files
func CreateRepoInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/repo\\.json(\\.gz)?$"))
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
		return CreateRepoInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Repo)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateRepoInputStreamFile creates an channel for reading Repo as JSON newlines from filename
func CreateRepoInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
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
	return CreateRepoInputStream(f, errors, transforms...)
}

// CreateRepoInputStream creates an channel for reading Repo as JSON newlines from stream
func CreateRepoInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoFunc) (chan Repo, <-chan bool) {
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

// CreateRepoOutputStreamDir will output json newlines from channel and save in dir
func CreateRepoOutputStreamDir(dir string, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/v1/repo\\.json(\\.gz)?$")
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
	return CreateRepoOutputStream(gz, ch, errors, transforms...)
}

// CreateRepoOutputStream will output json newlines from channel to the stream
func CreateRepoOutputStream(stream io.WriteCloser, ch chan Repo, errors chan<- error, transforms ...TransformRepoFunc) <-chan bool {
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

// CreateRepoProducer will stream data from the channel
func CreateRepoProducer(producer *util.KafkaProducer, ch chan Repo, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateRepoAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item, err)
			}
		}
	}()
	return done
}

// CreateRepoConsumer will stream data from the default topic into the provided channel
func CreateRepoConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Repo, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateRepoConsumerForTopic(kafkaServers, schemaRegistryServers, RepoDefaultTopic, consumerGroupID, ch, errors)
}

// CreateRepoConsumerForTopic will stream data from the topic into the provided channel
func CreateRepoConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Repo, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object Repo
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into Repo: %s", err)
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
