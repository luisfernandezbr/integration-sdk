// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// UserTopic is the default topic name
	UserTopic datamodel.TopicNameType = "sourcecode_User_topic"

	// UserStream is the default stream name
	UserStream datamodel.TopicNameType = "sourcecode_User_stream"

	// UserTable is the default table name
	UserTable datamodel.TopicNameType = "sourcecode_User"

	// UserModelName is the model name
	UserModelName datamodel.ModelNameType = "sourcecode.User"
)

const (
	// UserIDColumn is the id column name
	UserIDColumn = "id"
	// UserRefIDColumn is the ref_id column name
	UserRefIDColumn = "ref_id"
	// UserRefTypeColumn is the ref_type column name
	UserRefTypeColumn = "ref_type"
	// UserCustomerIDColumn is the customer_id column name
	UserCustomerIDColumn = "customer_id"
	// UserNameColumn is the name column name
	UserNameColumn = "name"
	// UserAvatarURLColumn is the avatar_url column name
	UserAvatarURLColumn = "avatar_url"
	// UserEmailColumn is the email column name
	UserEmailColumn = "email"
	// UserUsernameColumn is the username column name
	UserUsernameColumn = "username"
	// UserAssociatedRefIDColumn is the associated_ref_id column name
	UserAssociatedRefIDColumn = "associated_ref_id"
)

// User the source code user
type User struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	// custom types

	// Name the name of the user
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
	// Email the email for the user
	Email *string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// Username username of the user
	Username *string `json:"username" bson:"username" yaml:"username" faker:"username"`
	// AssociatedRefID the ref id associated for this user in another system
	AssociatedRefID *string `json:"associated_ref_id" bson:"associated_ref_id" yaml:"associated_ref_id" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*User)(nil)

func toUserObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toUserObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toUserObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toUserObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
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
			arr = append(arr, toUserObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of User
func (o *User) String() string {
	return fmt.Sprintf("sourcecode.User<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *User) GetTopicName() datamodel.TopicNameType {
	return UserTopic
}

// GetModelName returns the name of the model
func (o *User) GetModelName() datamodel.ModelNameType {
	return UserModelName
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

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *User) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *User) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *User) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *User) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *User) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "sourcecode_user",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *User) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *User) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *User) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *User) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of User
func (o *User) Clone() datamodel.Model {
	c := new(User)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *User) Anon() datamodel.Model {
	c := new(User)
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
func (o *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *User) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecUser *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *User) GetAvroCodec() *goavro.Codec {
	if cachedCodecUser == nil {
		c, err := GetUserAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUser = c
	}
	return cachedCodecUser
}

// ToAvroBinary returns the data as Avro binary data
func (o *User) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *User) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
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
		"id":                o.GetID(),
		"ref_id":            o.GetRefID(),
		"ref_type":          o.RefType,
		"customer_id":       o.CustomerID,
		"hashcode":          o.Hash(),
		"name":              toUserObject(o.Name, isavro, false, "string"),
		"avatar_url":        toUserObject(o.AvatarURL, isavro, true, "string"),
		"email":             toUserObject(o.Email, isavro, true, "string"),
		"username":          toUserObject(o.Username, isavro, true, "string"),
		"associated_ref_id": toUserObject(o.AssociatedRefID, isavro, true, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *User) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else if val, ok := kv["_id"].(string); ok {
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
	o.setDefaults()
	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		val := kv["name"]
		if val == nil {
			o.Name = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
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
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
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
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.Email = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["username"].(*string); ok {
		o.Username = val
	} else if val, ok := kv["username"].(string); ok {
		o.Username = &val
	} else {
		val := kv["username"]
		if val == nil {
			o.Username = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.Username = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["associated_ref_id"].(*string); ok {
		o.AssociatedRefID = val
	} else if val, ok := kv["associated_ref_id"].(string); ok {
		o.AssociatedRefID = &val
	} else {
		val := kv["associated_ref_id"]
		if val == nil {
			o.AssociatedRefID = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.AssociatedRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
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
	args = append(args, o.Username)
	args = append(args, o.AssociatedRefID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUserAvroSchemaSpec creates the avro schema specification for User
func GetUserAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "sourcecode",
		"name":      "User",
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
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "email",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "username",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "associated_ref_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetUserAvroSchema creates the avro schema for User
func GetUserAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserAvroSchemaSpec())
}

// TransformUserFunc is a function for transforming User during processing
type TransformUserFunc func(input *User) (*User, error)

// NewUserPipe creates a pipe for processing User items
func NewUserPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUserInputStream(input, errors)
	var stream chan User
	if len(transforms) > 0 {
		stream = make(chan User, 1000)
	} else {
		stream = inch
	}
	outdone := NewUserOutputStream(output, stream, errors)
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

// NewUserInputStreamDir creates a channel for reading User as JSON newlines from a directory of files
func NewUserInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
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
		return NewUserInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan User)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUserInputStreamFile creates an channel for reading User as JSON newlines from filename
func NewUserInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
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
	return NewUserInputStream(f, errors, transforms...)
}

// NewUserInputStream creates an channel for reading User as JSON newlines from stream
func NewUserInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserFunc) (chan User, <-chan bool) {
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

// NewUserOutputStreamDir will output json newlines from channel and save in dir
func NewUserOutputStreamDir(dir string, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
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
	return NewUserOutputStream(gz, ch, errors, transforms...)
}

// NewUserOutputStream will output json newlines from channel to the stream
func NewUserOutputStream(stream io.WriteCloser, ch chan User, errors chan<- error, transforms ...TransformUserFunc) <-chan bool {
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

// UserSendEvent is an event detail for sending data
type UserSendEvent struct {
	User    *User
	headers map[string]string
	time    time.Time
	key     string
}

var _ datamodel.ModelSendEvent = (*UserSendEvent)(nil)

// Key is the key to use for the message
func (e *UserSendEvent) Key() string {
	if e.key == "" {
		return e.User.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserSendEvent) Object() datamodel.Model {
	return e.User
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserSendEvent) Timestamp() time.Time {
	return e.time
}

// UserSendEventOpts is a function handler for setting opts
type UserSendEventOpts func(o *UserSendEvent)

// WithUserSendEventKey sets the key value to a value different than the object ID
func WithUserSendEventKey(key string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.key = key
	}
}

// WithUserSendEventTimestamp sets the timestamp value
func WithUserSendEventTimestamp(tv time.Time) UserSendEventOpts {
	return func(o *UserSendEvent) {
		o.time = tv
	}
}

// WithUserSendEventHeader sets the timestamp value
func WithUserSendEventHeader(key, value string) UserSendEventOpts {
	return func(o *UserSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserSendEvent returns a new UserSendEvent instance
func NewUserSendEvent(o *User, opts ...UserSendEventOpts) *UserSendEvent {
	res := &UserSendEvent{
		User: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserProducer will stream data from the channel
func NewUserProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*User); ok {
				binary, codec, err := object.ToAvroBinary()
				if err != nil {
					errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
					return
				}
				headers := map[string]string{}
				object.SetEventHeaders(headers)
				for k, v := range item.Headers() {
					headers[k] = v
				}
				tv := item.Timestamp()
				if tv.IsZero() {
					tv = object.GetTimestamp() // if not provided in the message, use the objects value
				}
				if tv.IsZero() {
					tv = time.Now() // if its still zero, use the ingest time
				}
				msg := eventing.Message{
					Encoding:  eventing.AvroEncoding,
					Key:       item.Key(),
					Value:     binary,
					Codec:     codec,
					Headers:   headers,
					Timestamp: tv,
					Topic:     object.GetTopicName().String(),
				}
				if err := producer.Send(ctx, msg); err != nil {
					errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
				}
			} else {
				errors <- fmt.Errorf("invalid event received. expected an object of type sourcecode.User but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewUserConsumer will stream data from the topic into the provided channel
func NewUserConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object User
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into sourcecode.User: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into sourcecode.User: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for sourcecode.User")
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object User
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserReceiveEvent{nil, msg, true}
		},
	})
}

// UserReceiveEvent is an event detail for receiving data
type UserReceiveEvent struct {
	User    *User
	message eventing.Message
	eof     bool
}

var _ datamodel.ModelReceiveEvent = (*UserReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserReceiveEvent) Object() datamodel.Model {
	return e.User
}

// Message returns the underlying message data for the event
func (e *UserReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserReceiveEvent) EOF() bool {
	return e.eof
}

// UserProducer implements the datamodel.ModelEventProducer
type UserProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*UserProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *User) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &UserProducer{
		ch:   ch,
		done: NewUserProducer(producer, ch, errors),
	}
}

// NewUserProducerChannel returns a channel which can be used for producing Model events
func NewUserProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &UserProducer{
		ch:   ch,
		done: NewUserProducer(producer, ch, errors),
	}
}

// UserConsumer implements the datamodel.ModelEventConsumer
type UserConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*UserConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *User) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewUserConsumer(consumer, ch, errors)
	return &UserConsumer{
		ch: ch,
	}
}

// NewUserConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewUserConsumer(consumer, ch, errors)
	return &UserConsumer{
		ch: ch,
	}
}
