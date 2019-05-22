// DO NOT EDIT -- generated code

// Package codequality - the system which contains code quality
package codequality

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
)

// MetricDefaultTopic is the default topic name
const MetricDefaultTopic = "codequality_Metric_topic"

// MetricDefaultStream is the default stream name
const MetricDefaultStream = "codequality_Metric_stream"

// MetricDefaultTable is the default table name
const MetricDefaultTable = "codequality_Metric"

// Metric individual metric details
type Metric struct {
	// built in types

	ID         string `json:"metric_id" yaml:"metric_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`
	// custom types

	// DateAt the date when the metric was created
	DateAt int64 `json:"date_ts" yaml:"date_ts"`
	// ProjectID the the project id
	ProjectID string `json:"project_id" yaml:"project_id"`
	// Name the metric name
	Name string `json:"name" yaml:"name"`
	// Value the value of the metric
	Value string `json:"value" yaml:"value"`
}

func toMetricObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toMetricObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toMetricObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toMetricObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toSprintObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case *Metric:
		return v.ToMap()
	case Metric:
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
			arr = append(arr, toMetricObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of Metric
func (o *Metric) String() string {
	return fmt.Sprintf("codequality.Metric<%s>", o.ID)
}

func (o *Metric) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Metric) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Metric", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Metric) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Metric) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Metric) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecMetric *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *Metric) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecMetric == nil {
		c, err := CreateMetricAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecMetric = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecMetric.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecMetric.BinaryFromNative(nil, native)
	return buf, cachedCodecMetric, err
}

// Stringify returns the object in JSON format as a string
func (o *Metric) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Metric objects are equal
func (o *Metric) IsEqual(other *Metric) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Metric) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"metric_id":   o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"date_ts":     toMetricObject(o.DateAt, isavro, false, "long"),
		"project_id":  toMetricObject(o.ProjectID, isavro, false, "string"),
		"name":        toMetricObject(o.Name, isavro, false, "string"),
		"value":       toMetricObject(o.Value, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *Metric) FromMap(kv map[string]interface{}) {
	if val, ok := kv["metric_id"].(string); ok {
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
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		val := kv["project_id"]
		if val == nil {
			o.ProjectID = ""
		} else {
			o.ProjectID = fmt.Sprintf("%v", val)
		}
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
	if val, ok := kv["value"].(string); ok {
		o.Value = val
	} else {
		val := kv["value"]
		if val == nil {
			o.Value = ""
		} else {
			o.Value = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *Metric) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.DateAt)
	args = append(args, o.ProjectID)
	args = append(args, o.Name)
	args = append(args, o.Value)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateMetricAvroSchemaSpec creates the avro schema specification for Metric
func CreateMetricAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "codequality",
		"name":         "Metric",
		"connect.name": "codequality.Metric",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "metric_id",
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
				"name": "project_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "value",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateMetricAvroSchema creates the avro schema for Metric
func CreateMetricAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateMetricAvroSchemaSpec())
}

// TransformMetricFunc is a function for transforming Metric during processing
type TransformMetricFunc func(input *Metric) (*Metric, error)

// CreateMetricPipe creates a pipe for processing Metric items
func CreateMetricPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformMetricFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateMetricInputStream(input, errors)
	var stream chan Metric
	if len(transforms) > 0 {
		stream = make(chan Metric, 1000)
	} else {
		stream = inch
	}
	outdone := CreateMetricOutputStream(output, stream, errors)
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

// CreateMetricInputStreamDir creates a channel for reading Metric as JSON newlines from a directory of files
func CreateMetricInputStreamDir(dir string, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/codequality/metric\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for metric")
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateMetricInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan Metric)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateMetricInputStreamFile creates an channel for reading Metric as JSON newlines from filename
func CreateMetricInputStreamFile(filename string, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan Metric)
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
			ch := make(chan Metric)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateMetricInputStream(f, errors, transforms...)
}

// CreateMetricInputStream creates an channel for reading Metric as JSON newlines from stream
func CreateMetricInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformMetricFunc) (chan Metric, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan Metric, 1000)
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
			var item Metric
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

// CreateMetricOutputStreamDir will output json newlines from channel and save in dir
func CreateMetricOutputStreamDir(dir string, ch chan Metric, errors chan<- error, transforms ...TransformMetricFunc) <-chan bool {
	fp := filepath.Join(dir, "/codequality/metric\\.json(\\.gz)?$")
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
	return CreateMetricOutputStream(gz, ch, errors, transforms...)
}

// CreateMetricOutputStream will output json newlines from channel to the stream
func CreateMetricOutputStream(stream io.WriteCloser, ch chan Metric, errors chan<- error, transforms ...TransformMetricFunc) <-chan bool {
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

// CreateMetricProducer will stream data from the channel
func CreateMetricProducer(producer Producer, ch chan Metric, errors chan<- error) <-chan bool {
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

// CreateMetricConsumer will stream data from the default topic into the provided channel
func CreateMetricConsumer(factory ConsumerFactory, topic string, ch chan Metric, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateMetricConsumerForTopic(factory, MetricDefaultTopic, ch, errors)
}

// CreateMetricConsumerForTopic will stream data from the topic into the provided channel
func CreateMetricConsumerForTopic(factory ConsumerFactory, topic string, ch chan Metric, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object Metric
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into Metric: %s", err)
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
