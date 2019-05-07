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

// ChangelogDefaultTopic is the default topic name
const ChangelogDefaultTopic = "sourcecode.Changelog.topic"

// ChangelogDefaultStream is the default stream name
const ChangelogDefaultStream = "sourcecode.Changelog.topicstream"

// ChangelogDefaultTable is the default table name
const ChangelogDefaultTable = "sourcecode.Changelog"

// Changelog the change log for each commit in a repo
type Changelog struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	RepoID      string `json:"repo_id" yaml:"repo_id"`
	Filename    string `json:"filename" yaml:"filename"`
	Language    string `json:"language" yaml:"language"`
	Loc         int64  `json:"loc" yaml:"loc"`
	Sloc        int64  `json:"sloc" yaml:"sloc"`
	Blanks      int64  `json:"blanks" yaml:"blanks"`
	Comments    int64  `json:"comments" yaml:"comments"`
	Complexity  int64  `json:"complexity" yaml:"complexity"`
	DateAt      int64  `json:"date_ts" yaml:"date_ts"`
	AuthorRefID string `json:"author_ref_id" yaml:"author_ref_id"`
	Ordinal     int64  `json:"ordinal" yaml:"ordinal"`
	Sha         string `json:"sha" yaml:"sha"`
}

// String returns a string representation of Changelog
func (o *Changelog) String() string {
	return fmt.Sprintf("sourcecode.v1.Changelog<%s>", o.ID)
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

// GetRefID returns the RefID for the object
func (o *Changelog) GetRefID() string {
	return o.RefID
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

// Stringify returns the object in JSON format as a string
func (o *Changelog) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Changelog objects are equal
func (o *Changelog) IsEqual(other *Changelog) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Changelog) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            o.GetID(),
		"ref_id":        o.GetRefID(),
		"ref_type":      o.RefType,
		"customer_id":   o.CustomerID,
		"hashcode":      o.Hash(),
		"repo_id":       o.RepoID,
		"filename":      o.Filename,
		"language":      o.Language,
		"loc":           o.Loc,
		"sloc":          o.Sloc,
		"blanks":        o.Blanks,
		"comments":      o.Comments,
		"complexity":    o.Complexity,
		"date_ts":       o.DateAt,
		"author_ref_id": o.AuthorRefID,
		"ordinal":       o.Ordinal,
		"sha":           o.Sha,
	}
}

// FromMap attempts to load data into object from a map
func (o *Changelog) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		val := kv["filename"]
		o.Filename = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		o.Language = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		val := kv["loc"]
		o.Loc = number.ToInt64Any(val)
	}
	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		val := kv["sloc"]
		o.Sloc = number.ToInt64Any(val)
	}
	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		val := kv["blanks"]
		o.Blanks = number.ToInt64Any(val)
	}
	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		val := kv["comments"]
		o.Comments = number.ToInt64Any(val)
	}
	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		val := kv["complexity"]
		o.Complexity = number.ToInt64Any(val)
	}
	if val, ok := kv["date_ts"].(int64); ok {
		o.DateAt = val
	} else {
		val := kv["date_ts"]
		o.DateAt = number.ToInt64Any(val)
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
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		val := kv["sha"]
		o.Sha = fmt.Sprintf("%v", val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Changelog) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
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
	}
	return o.Hashcode
}

// CreateChangelogAvroSchemaSpec creates the avro schema specification for Changelog
func CreateChangelogAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode.v1",
		"name":         "Changelog",
		"connect.name": "sourcecode.v1.Changelog",
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

// CreateChangelogAvroSchema creates the avro schema for Changelog
func CreateChangelogAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateChangelogAvroSchemaSpec())
}

// CreateChangelogKQLStreamSQL creates KQL Stream SQL for Changelog
func CreateChangelogKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", ChangelogDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", ChangelogDefaultTopic))
	builder.WriteString(", TIMESTAMP='date_ts'")
	builder.WriteString(");")
	return builder.String()
}

// CreateChangelogKQLTableSQL creates KQL Table SQL for Changelog
func CreateChangelogKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", ChangelogDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", ChangelogDefaultTopic))
	builder.WriteString(", TIMESTAMP='date_ts'")
	builder.WriteString(");")
	return builder.String()
}

// TransformChangelogFunc is a function for transforming Changelog during processing
type TransformChangelogFunc func(input *Changelog) (*Changelog, error)

// CreateChangelogPipe creates a pipe for processing Changelog items
func CreateChangelogPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformChangelogFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateChangelogInputStream(input, errors)
	var stream chan Changelog
	if len(transforms) > 0 {
		stream = make(chan Changelog, 1000)
	} else {
		stream = inch
	}
	outdone := CreateChangelogOutputStream(output, stream, errors)
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

// CreateChangelogInputStreamDir creates a channel for reading Changelog as JSON newlines from a directory of files
func CreateChangelogInputStreamDir(dir string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/changelog\\.json(\\.gz)?$"))
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
		return CreateChangelogInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Changelog)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateChangelogInputStreamFile creates an channel for reading Changelog as JSON newlines from filename
func CreateChangelogInputStreamFile(filename string, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
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
	return CreateChangelogInputStream(f, errors, transforms...)
}

// CreateChangelogInputStream creates an channel for reading Changelog as JSON newlines from stream
func CreateChangelogInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformChangelogFunc) (chan Changelog, <-chan bool) {
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

// CreateChangelogOutputStreamDir will output json newlines from channel and save in dir
func CreateChangelogOutputStreamDir(dir string, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/v1/changelog\\.json(\\.gz)?$")
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
	return CreateChangelogOutputStream(gz, ch, errors, transforms...)
}

// CreateChangelogOutputStream will output json newlines from channel to the stream
func CreateChangelogOutputStream(stream io.WriteCloser, ch chan Changelog, errors chan<- error, transforms ...TransformChangelogFunc) <-chan bool {
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

// CreateChangelogProducer will stream data from the channel
func CreateChangelogProducer(producer *util.KafkaProducer, ch chan Changelog, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateChangelogAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- fmt.Errorf("error sending %s. %v", item, err)
			}
		}
	}()
	return done
}

// CreateChangelogConsumer will stream data from the default topic into the provided channel
func CreateChangelogConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateChangelogConsumerForTopic(kafkaServers, schemaRegistryServers, ChangelogDefaultTopic, consumerGroupID, ch, errors)
}

// CreateChangelogConsumerForTopic will stream data from the topic into the provided channel
func CreateChangelogConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan Changelog, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object Changelog
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into Changelog: %s", err)
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
