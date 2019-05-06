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
	pstrings "github.com/pinpt/go-common/strings"
	"github.com/pinpt/integration-sdk/util"
)

// CommitFileDefaultTopic is the default topic name
const CommitFileDefaultTopic = "sourcecode.CommitFile.topic"

// CommitFileDefaultStream is the default stream name
const CommitFileDefaultStream = "sourcecode.CommitFile.topicstream"

// CommitFileDefaultTable is the default table name
const CommitFileDefaultTable = "sourcecode.CommitFile"

// CommitFile the file change for a specific commit
type CommitFile struct {
	// built in types

	ID         string `json:"id" yaml:"id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	CommitID          string  `json:"commit_id" yaml:"commit_id"`
	Filename          string  `json:"filename" yaml:"filename"`
	Additions         int64   `json:"additions" yaml:"additions"`
	Deletions         int64   `json:"deletions" yaml:"deletions"`
	Status            string  `json:"status" yaml:"status"`
	Binary            bool    `json:"binary" yaml:"binary"`
	Language          string  `json:"language" yaml:"language"`
	Excluded          bool    `json:"excluded" yaml:"excluded"`
	ExcludedReason    string  `json:"excluded_reason" yaml:"excluded_reason"`
	Ordinal           int64   `json:"ordinal" yaml:"ordinal"`
	Loc               int64   `json:"loc" yaml:"loc"`
	Sloc              int64   `json:"sloc" yaml:"sloc"`
	Blanks            int64   `json:"blanks" yaml:"blanks"`
	Comments          int64   `json:"comments" yaml:"comments"`
	Complexity        int64   `json:"complexity" yaml:"complexity"`
	License           *string `json:"license" yaml:"license"`
	LicenseConfidence float64 `json:"license_confidence" yaml:"license_confidence"`
}

// String returns a string representation of CommitFile
func (o *CommitFile) String() string {
	return fmt.Sprintf("sourcecode.v1.CommitFile<%s>", o.ID)
}

func (o *CommitFile) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CommitFile) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CommitFile", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *CommitFile) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CommitFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CommitFile) UnmarshalJSON(data []byte) error {
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
func (o *CommitFile) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CommitFile objects are equal
func (o *CommitFile) IsEqual(other *CommitFile) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CommitFile) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                 o.GetID(),
		"ref_id":             o.GetRefID(),
		"ref_type":           o.RefType,
		"customer_id":        o.CustomerID,
		"hashcode":           o.Hash(),
		"commit_id":          o.CommitID,
		"filename":           o.Filename,
		"additions":          o.Additions,
		"deletions":          o.Deletions,
		"status":             o.Status,
		"binary":             o.Binary,
		"language":           o.Language,
		"excluded":           o.Excluded,
		"excluded_reason":    o.ExcludedReason,
		"ordinal":            o.Ordinal,
		"loc":                o.Loc,
		"sloc":               o.Sloc,
		"blanks":             o.Blanks,
		"comments":           o.Comments,
		"complexity":         o.Complexity,
		"license":            o.License,
		"license_confidence": o.LicenseConfidence,
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitFile) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		val := kv["commit_id"]
		o.CommitID = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		val := kv["filename"]
		o.Filename = fmt.Sprintf("%v", val)
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
	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		val := kv["status"]
		o.Status = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["binary"].(bool); ok {
		o.Binary = val
	} else {
		val := kv["binary"]
		o.Binary = number.ToBoolAny(val)
	}
	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		val := kv["language"]
		o.Language = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		val := kv["excluded"]
		o.Excluded = number.ToBoolAny(val)
	}
	if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = val
	} else {
		val := kv["excluded_reason"]
		o.ExcludedReason = fmt.Sprintf("%v", val)
	}
	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		val := kv["ordinal"]
		o.Ordinal = number.ToInt64Any(val)
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
	if val, ok := kv["license"].(*string); ok {
		o.License = val
	} else if val, ok := kv["license"].(string); ok {
		o.License = &val
	} else {
		val := kv["license"]
		o.License = pstrings.Pointer(fmt.Sprintf("%v", val))
	}
	if val, ok := kv["license_confidence"].(float64); ok {
		o.LicenseConfidence = val
	} else {
		val := kv["license_confidence"]
		o.LicenseConfidence = number.ToFloat64Any(val)
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *CommitFile) Hash() string {
	if o.Hashcode == "" {
		args := make([]interface{}, 0)
		args = append(args, o.GetID())
		args = append(args, o.GetRefID())
		args = append(args, o.RefType)
		args = append(args, o.CommitID)
		args = append(args, o.Filename)
		args = append(args, o.Additions)
		args = append(args, o.Deletions)
		args = append(args, o.Status)
		args = append(args, o.Binary)
		args = append(args, o.Language)
		args = append(args, o.Excluded)
		args = append(args, o.ExcludedReason)
		args = append(args, o.Ordinal)
		args = append(args, o.Loc)
		args = append(args, o.Sloc)
		args = append(args, o.Blanks)
		args = append(args, o.Comments)
		args = append(args, o.Complexity)
		args = append(args, o.License)
		args = append(args, o.LicenseConfidence)
		o.Hashcode = hash.Values(args...)
	}
	return o.Hashcode
}

// CreateCommitFileAvroSchemaSpec creates the avro schema specification for CommitFile
func CreateCommitFileAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode.v1",
		"name":         "CommitFile",
		"connect.name": "sourcecode.v1.CommitFile",
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
				"name": "commit_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "filename",
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
				"name": "status",
				"type": "string",
			},
			map[string]interface{}{
				"name": "binary",
				"type": "bool",
			},
			map[string]interface{}{
				"name": "language",
				"type": "string",
			},
			map[string]interface{}{
				"name": "excluded",
				"type": "bool",
			},
			map[string]interface{}{
				"name": "excluded_reason",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ordinal",
				"type": "long",
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
				"name": "license",
				"type": []string{"null", "string"},
			},
			map[string]interface{}{
				"name": "license_confidence",
				"type": "float",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateCommitFileAvroSchema creates the avro schema for CommitFile
func CreateCommitFileAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateCommitFileAvroSchemaSpec())
}

// CreateCommitFileKQLStreamSQL creates KQL Stream SQL for CommitFile
func CreateCommitFileKQLStreamSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE STREAM %s ", CommitFileDefaultStream))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", CommitFileDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// CreateCommitFileKQLTableSQL creates KQL Table SQL for CommitFile
func CreateCommitFileKQLTableSQL() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("CREATE TABLE %s ", CommitFileDefaultTable))
	builder.WriteString(fmt.Sprintf("WITH (KAFKA_TOPIC='%s', VALUE_FORMAT='AVRO', KEY='id'", CommitFileDefaultTopic))
	builder.WriteString(");")
	return builder.String()
}

// TransformCommitFileFunc is a function for transforming CommitFile during processing
type TransformCommitFileFunc func(input *CommitFile) (*CommitFile, error)

// CreateCommitFilePipe creates a pipe for processing CommitFile items
func CreateCommitFilePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCommitFileFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateCommitFileInputStream(input, errors)
	var stream chan CommitFile
	if len(transforms) > 0 {
		stream = make(chan CommitFile, 1000)
	} else {
		stream = inch
	}
	outdone := CreateCommitFileOutputStream(output, stream, errors)
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

// CreateCommitFileInputStreamDir creates a channel for reading CommitFile as JSON newlines from a directory of files
func CreateCommitFileInputStreamDir(dir string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/v1/commit_file\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for commit_file")
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateCommitFileInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CommitFile)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateCommitFileInputStreamFile creates an channel for reading CommitFile as JSON newlines from filename
func CreateCommitFileInputStreamFile(filename string, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CommitFile)
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
			ch := make(chan CommitFile)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateCommitFileInputStream(f, errors, transforms...)
}

// CreateCommitFileInputStream creates an channel for reading CommitFile as JSON newlines from stream
func CreateCommitFileInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCommitFileFunc) (chan CommitFile, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CommitFile, 1000)
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
			var item CommitFile
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

// CreateCommitFileOutputStreamDir will output json newlines from channel and save in dir
func CreateCommitFileOutputStreamDir(dir string, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/v1/commit_file\\.json(\\.gz)?$")
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
	return CreateCommitFileOutputStream(gz, ch, errors, transforms...)
}

// CreateCommitFileOutputStream will output json newlines from channel to the stream
func CreateCommitFileOutputStream(stream io.WriteCloser, ch chan CommitFile, errors chan<- error, transforms ...TransformCommitFileFunc) <-chan bool {
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

// CreateCommitFileProducer will stream data from the channel
func CreateCommitFileProducer(producer *util.KafkaProducer, ch chan CommitFile, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		schemaspec := CreateCommitFileAvroSchemaSpec()
		ctx := context.Background()
		for item := range ch {
			if err := producer.Send(ctx, schemaspec, []byte(item.ID), []byte(item.Stringify())); err != nil {
				errors <- err
			}
		}
	}()
	return done
}

// CreateCommitFileConsumer will stream data from the default topic into the provided channel
func CreateCommitFileConsumer(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan CommitFile, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateCommitFileConsumerForTopic(kafkaServers, schemaRegistryServers, CommitFileDefaultTopic, consumerGroupID, ch, errors)
}

// CreateCommitFileConsumerForTopic will stream data from the topic into the provided channel
func CreateCommitFileConsumerForTopic(kafkaServers []string, schemaRegistryServers []string, topic string, consumerGroupID string, ch chan CommitFile, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		consumerCallbacks := kafka.ConsumerCallbacks{
			OnDataReceived: func(msg kafka.Message) {
				var object CommitFile
				if err := json.Unmarshal([]byte(msg.Value), &object); err != nil {
					errors <- fmt.Errorf("error unmarshaling json data into CommitFile: %s", err)
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
