// DO NOT EDIT -- generated code

// Package customer - customer specific data
package customer

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
	"github.com/pinpt/integration-sdk/util"
)

// UserLocationDefaultTopic is the default topic name
const UserLocationDefaultTopic = "customer_UserLocation_topic"

// UserLocationDefaultStream is the default stream name
const UserLocationDefaultStream = "customer_UserLocation_stream"

// UserLocationDefaultTable is the default table name
const UserLocationDefaultTable = "customer_UserLocation"

// UserLocation blah blah blah
type UserLocation struct {
	// built in types

	ID         string `json:"user_location_id" yaml:"user_location_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`

	// custom types

	// UserID _
	UserID string `json:"user_id" yaml:"user_id"`
	// Location _
	Location string `json:"location" yaml:"location"`
	// Region _
	Region string `json:"region" yaml:"region"`
	// LocationID _
	LocationID string `json:"location_id" yaml:"location_id"`
	// RegionID _
	RegionID string `json:"region_id" yaml:"region_id"`
}

func toUserLocationObject(o interface{}, isavro bool) interface{} {
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
	case *UserLocation:
		return v.ToMap()
	case UserLocation:
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
			arr = append(arr, toUserLocationObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of UserLocation
func (o *UserLocation) String() string {
	return fmt.Sprintf("customer.UserLocation<%s>", o.ID)
}

func (o *UserLocation) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserLocation) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserLocation", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *UserLocation) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *UserLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserLocation) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecUserLocation *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *UserLocation) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecUserLocation == nil {
		c, err := CreateUserLocationAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecUserLocation = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecUserLocation.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecUserLocation.BinaryFromNative(nil, native)
	return buf, cachedCodecUserLocation, err
}

// Stringify returns the object in JSON format as a string
func (o *UserLocation) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserLocation objects are equal
func (o *UserLocation) IsEqual(other *UserLocation) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserLocation) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"user_location_id": o.GetID(),
		"ref_id":           o.GetRefID(),
		"ref_type":         o.RefType,
		"customer_id":      o.CustomerID,
		"hashcode":         o.Hash(),
		"user_id":          toUserLocationObject(o.UserID, isavro),
		"location":         toUserLocationObject(o.Location, isavro),
		"region":           toUserLocationObject(o.Region, isavro),
		"location_id":      toUserLocationObject(o.LocationID, isavro),
		"region_id":        toUserLocationObject(o.RegionID, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserLocation) FromMap(kv map[string]interface{}) {
	if val, ok := kv["user_location_id"].(string); ok {
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
	if val, ok := kv["location"].(string); ok {
		o.Location = val
	} else {
		val := kv["location"]
		if val == nil {
			o.Location = ""
		} else {
			o.Location = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["region"].(string); ok {
		o.Region = val
	} else {
		val := kv["region"]
		if val == nil {
			o.Region = ""
		} else {
			o.Region = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["location_id"].(string); ok {
		o.LocationID = val
	} else {
		val := kv["location_id"]
		if val == nil {
			o.LocationID = ""
		} else {
			o.LocationID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["region_id"].(string); ok {
		o.RegionID = val
	} else {
		val := kv["region_id"]
		if val == nil {
			o.RegionID = ""
		} else {
			o.RegionID = fmt.Sprintf("%v", val)
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *UserLocation) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.UserID)
	args = append(args, o.Location)
	args = append(args, o.Region)
	args = append(args, o.LocationID)
	args = append(args, o.RegionID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUserLocationAvroSchemaSpec creates the avro schema specification for UserLocation
func CreateUserLocationAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "UserLocation",
		"connect.name": "customer.UserLocation",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "user_location_id",
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
				"name": "user_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "location",
				"type": "string",
			},
			map[string]interface{}{
				"name": "region",
				"type": "string",
			},
			map[string]interface{}{
				"name": "location_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "region_id",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateUserLocationAvroSchema creates the avro schema for UserLocation
func CreateUserLocationAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateUserLocationAvroSchemaSpec())
}

// TransformUserLocationFunc is a function for transforming UserLocation during processing
type TransformUserLocationFunc func(input *UserLocation) (*UserLocation, error)

// CreateUserLocationPipe creates a pipe for processing UserLocation items
func CreateUserLocationPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserLocationFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateUserLocationInputStream(input, errors)
	var stream chan UserLocation
	if len(transforms) > 0 {
		stream = make(chan UserLocation, 1000)
	} else {
		stream = inch
	}
	outdone := CreateUserLocationOutputStream(output, stream, errors)
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

// CreateUserLocationInputStreamDir creates a channel for reading UserLocation as JSON newlines from a directory of files
func CreateUserLocationInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserLocationFunc) (chan UserLocation, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/user_location\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan UserLocation)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user_location")
		ch := make(chan UserLocation)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateUserLocationInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan UserLocation)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateUserLocationInputStreamFile creates an channel for reading UserLocation as JSON newlines from filename
func CreateUserLocationInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserLocationFunc) (chan UserLocation, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan UserLocation)
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
			ch := make(chan UserLocation)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateUserLocationInputStream(f, errors, transforms...)
}

// CreateUserLocationInputStream creates an channel for reading UserLocation as JSON newlines from stream
func CreateUserLocationInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserLocationFunc) (chan UserLocation, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan UserLocation, 1000)
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
			var item UserLocation
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

// CreateUserLocationOutputStreamDir will output json newlines from channel and save in dir
func CreateUserLocationOutputStreamDir(dir string, ch chan UserLocation, errors chan<- error, transforms ...TransformUserLocationFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/user_location\\.json(\\.gz)?$")
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
	return CreateUserLocationOutputStream(gz, ch, errors, transforms...)
}

// CreateUserLocationOutputStream will output json newlines from channel to the stream
func CreateUserLocationOutputStream(stream io.WriteCloser, ch chan UserLocation, errors chan<- error, transforms ...TransformUserLocationFunc) <-chan bool {
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

// CreateUserLocationProducer will stream data from the channel
func CreateUserLocationProducer(producer util.Producer, ch chan UserLocation, errors chan<- error) <-chan bool {
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

// CreateUserLocationConsumer will stream data from the default topic into the provided channel
func CreateUserLocationConsumer(factory util.ConsumerFactory, topic string, ch chan UserLocation, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateUserLocationConsumerForTopic(factory, UserLocationDefaultTopic, ch, errors)
}

// CreateUserLocationConsumerForTopic will stream data from the topic into the provided channel
func CreateUserLocationConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan UserLocation, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object UserLocation
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into UserLocation: %s", err)
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
