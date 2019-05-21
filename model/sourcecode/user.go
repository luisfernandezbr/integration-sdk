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

	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
)

// UserDefaultTopic is the default topic name
const UserDefaultTopic = "sourcecode_User_topic"

// UserDefaultStream is the default stream name
const UserDefaultStream = "sourcecode_User_stream"

// UserDefaultTable is the default table name
const UserDefaultTable = "sourcecode_User"

// User the source code user
type User struct {
	// built in types

	ID         string `json:"user_id" yaml:"user_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`
	// custom types

	// Name the name of the user
	Name string `json:"name" yaml:"name"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url" yaml:"avatar_url"`
	// Email the email for the user
	Email *string `json:"email" yaml:"email"`
}

func toUserObject(o interface{}, isavro bool) interface{} {
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
	case *User:
		return v.ToMap()
	case User:
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
			arr = append(arr, toUserObject(av, isavro))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("sourcecode.User<%s>", o.ID)
}

func (o *User) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("User", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// MarshalJSON returns the bytes for marshaling to json
func (o *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *User) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	// make sure that these have values if empty
	o.setDefaults()
	return nil
}

var cachedCodecUser *goavro.Codec

// ToAvroBinary returns the data as Avro binary data
func (o *User) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	if cachedCodecUser == nil {
		c, err := CreateUserAvroSchema()
		if err != nil {
			return nil, nil, err
		}
		cachedCodecUser = c
	}
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	native, _, err := cachedCodecUser.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := cachedCodecUser.BinaryFromNative(nil, native)
	return buf, cachedCodecUser, err
}

// Stringify returns the object in JSON format as a string
func (o *User) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two User objects are equal
func (o *User) IsEqual(other *User) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *User) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"user_id":     o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"name":        toUserObject(o.Name, isavro),
		"avatar_url":  toUserObject(o.AvatarURL, isavro),
		"email":       toUserObject(o.Email, isavro),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {
	if val, ok := kv["user_id"].(string); ok {
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
	if val, ok := kv["avatar_url"].(*string); ok {
		o.AvatarURL = val
	} else if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = &val
	} else {
		val := kv["avatar_url"]
		if val == nil {
			o.AvatarURL = pstrings.Pointer("")
		} else {
			o.AvatarURL = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["email"].(*string); ok {
		o.Email = val
	} else if val, ok := kv["email"].(string); ok {
		o.Email = &val
	} else {
		val := kv["email"]
		if val == nil {
			o.Email = pstrings.Pointer("")
		} else {
			o.Email = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	// make sure that these have values if empty
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *User) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.AvatarURL)
	args = append(args, o.Email)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUserAvroSchemaSpec creates the avro schema specification for User
func CreateUserAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "sourcecode",
		"name":         "User",
		"connect.name": "sourcecode.User",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "user_id",
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
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "email",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// CreateUserAvroSchema creates the avro schema for User
func CreateUserAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(CreateUserAvroSchemaSpec())
}

// TransformUserFunc is a function for transforming User during processing
type TransformUserFunc func(input *User) (*User, error)

// CreateUserPipe creates a pipe for processing User items
func CreateUserPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := CreateUserInputStream(input, errors)
	var stream chan User
	if len(transforms) > 0 {
		stream = make(chan User, 1000)
	} else {
		stream = inch
	}
	outdone := CreateUserOutputStream(output, stream, errors)
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

// CreateUserInputStreamDir creates a channel for reading User as JSON newlines from a directory of files
func CreateUserInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/sourcecode/user\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user")
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return CreateUserInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// CreateUserInputStreamFile creates an channel for reading User as JSON newlines from filename
func CreateUserInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan User)
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
			ch := make(chan User)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return CreateUserInputStream(f, errors, transforms...)
}

// CreateUserInputStream creates an channel for reading User as JSON newlines from stream
func CreateUserInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan User, 1000)
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
			var item User
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

// CreateUserOutputStreamDir will output json newlines from channel and save in dir
func CreateUserOutputStreamDir(dir string, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
	fp := filepath.Join(dir, "/sourcecode/user\\.json(\\.gz)?$")
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
	return CreateUserOutputStream(gz, ch, errors, transforms...)
}

// CreateUserOutputStream will output json newlines from channel to the stream
func CreateUserOutputStream(stream io.WriteCloser, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
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

// CreateUserProducer will stream data from the channel
func CreateUserProducer(producer util.Producer, ch chan User, errors chan<- error) <-chan bool {
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

// CreateUserConsumer will stream data from the default topic into the provided channel
func CreateUserConsumer(factory util.ConsumerFactory, topic string, ch chan User, errors chan<- error) (<-chan bool, chan<- bool) {
	return CreateUserConsumerForTopic(factory, UserDefaultTopic, ch, errors)
}

// CreateUserConsumerForTopic will stream data from the topic into the provided channel
func CreateUserConsumerForTopic(factory util.ConsumerFactory, topic string, ch chan User, errors chan<- error) (<-chan bool, chan<- bool) {
	done := make(chan bool, 1)
	closed := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		callback := util.ConsumerCallback{
			OnDataReceived: func(key []byte, value []byte) error {
				var object User
				if err := json.Unmarshal(value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into User: %s", err)
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
