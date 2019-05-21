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

	pjson "github.com/pinpt/go-common/json"
	number "github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
)

// UserDefaultTopic is the default topic name
const UserDefaultTopic = "customer_User_topic"

// UserDefaultStream is the default stream name
const UserDefaultStream = "customer_User_stream"

// UserDefaultTable is the default table name
const UserDefaultTable = "customer_User"

// User the customer's user record
type User struct {
	// built in types

	ID         string `json:"user_id" yaml:"user_id"`
	RefID      string `json:"ref_id" yaml:"ref_id"`
	RefType    string `json:"ref_type" yaml:"ref_type"`
	CustomerID string `json:"customer_id" yaml:"customer_id"`
	Hashcode   string `json:"hashcode" yaml:"hashcode"`
	// custom types

	// Name name of the user
	Name string `json:"name" yaml:"name"`
	// Email the email of the user
	Email string `json:"email" yaml:"email"`
	// Title the title of the user
	Title *string `json:"title" yaml:"title"`
	// Location the location of the user
	Location *string `json:"location" yaml:"location"`
	// AvatarURL the user avatar url
	AvatarURL *string `json:"avatar_url" yaml:"avatar_url"`
	// ManagerID the manager user id
	ManagerID *string `json:"manager_id" yaml:"manager_id"`
	// Active if true, the user is active and able to login
	Active bool `json:"active" yaml:"active"`
	// Trackable if true, the user is trackable in the pinpoint system
	Trackable bool `json:"trackable" yaml:"trackable"`
	// CreatedAt when the user was created in epoch timestamp
	CreatedAt int64 `json:"created_ts" yaml:"created_ts"`
	// UpdatedAt when the user record was updated in epoch timestamp
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts"`
	// DeletedAt when the user record was deleted in epoch timestamp
	DeletedAt *int64 `json:"deleted_ts" yaml:"deleted_ts"`
	// HiredAt when the user was hired in epoch timestamp
	HiredAt *int64 `json:"hired_ts" yaml:"hired_ts"`
	// TerminatedAt when the user was terminated in epoch timestamp
	TerminatedAt *int64 `json:"terminated_ts" yaml:"terminated_ts"`
	// CostCenterID the id of the cost center
	CostCenterID *string `json:"cost_center_id" yaml:"cost_center_id"`
	// TeamID the team id that the user is part of
	TeamID *string `json:"team_id" yaml:"team_id"`
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
	return fmt.Sprintf("customer.User<%s>", o.ID)
}

func (o *User) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *User) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.Email)
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
		"user_id":        o.GetID(),
		"ref_id":         o.GetRefID(),
		"ref_type":       o.RefType,
		"customer_id":    o.CustomerID,
		"hashcode":       o.Hash(),
		"name":           toUserObject(o.Name, isavro),
		"email":          toUserObject(o.Email, isavro),
		"title":          toUserObject(o.Title, isavro),
		"location":       toUserObject(o.Location, isavro),
		"avatar_url":     toUserObject(o.AvatarURL, isavro),
		"manager_id":     toUserObject(o.ManagerID, isavro),
		"active":         toUserObject(o.Active, isavro),
		"trackable":      toUserObject(o.Trackable, isavro),
		"created_ts":     toUserObject(o.CreatedAt, isavro),
		"updated_ts":     toUserObject(o.UpdatedAt, isavro),
		"deleted_ts":     toUserObject(o.DeletedAt, isavro),
		"hired_ts":       toUserObject(o.HiredAt, isavro),
		"terminated_ts":  toUserObject(o.TerminatedAt, isavro),
		"cost_center_id": toUserObject(o.CostCenterID, isavro),
		"team_id":        toUserObject(o.TeamID, isavro),
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
	if val, ok := kv["email"].(string); ok {
		o.Email = val
	} else {
		val := kv["email"]
		if val == nil {
			o.Email = ""
		} else {
			o.Email = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["title"].(*string); ok {
		o.Title = val
	} else if val, ok := kv["title"].(string); ok {
		o.Title = &val
	} else {
		val := kv["title"]
		if val == nil {
			o.Title = pstrings.Pointer("")
		} else {
			o.Title = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["location"].(*string); ok {
		o.Location = val
	} else if val, ok := kv["location"].(string); ok {
		o.Location = &val
	} else {
		val := kv["location"]
		if val == nil {
			o.Location = pstrings.Pointer("")
		} else {
			o.Location = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["manager_id"].(*string); ok {
		o.ManagerID = val
	} else if val, ok := kv["manager_id"].(string); ok {
		o.ManagerID = &val
	} else {
		val := kv["manager_id"]
		if val == nil {
			o.ManagerID = pstrings.Pointer("")
		} else {
			o.ManagerID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		val := kv["active"]
		if val == nil {
			o.Active = number.ToBoolAny(nil)
		} else {
			o.Active = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["trackable"].(bool); ok {
		o.Trackable = val
	} else {
		val := kv["trackable"]
		if val == nil {
			o.Trackable = number.ToBoolAny(nil)
		} else {
			o.Trackable = number.ToBoolAny(val)
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
	if val, ok := kv["deleted_ts"].(*int64); ok {
		o.DeletedAt = val
	} else if val, ok := kv["deleted_ts"].(int64); ok {
		o.DeletedAt = &val
	} else {
		val := kv["deleted_ts"]
		if val == nil {
			o.DeletedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			o.DeletedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["hired_ts"].(*int64); ok {
		o.HiredAt = val
	} else if val, ok := kv["hired_ts"].(int64); ok {
		o.HiredAt = &val
	} else {
		val := kv["hired_ts"]
		if val == nil {
			o.HiredAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			o.HiredAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["terminated_ts"].(*int64); ok {
		o.TerminatedAt = val
	} else if val, ok := kv["terminated_ts"].(int64); ok {
		o.TerminatedAt = &val
	} else {
		val := kv["terminated_ts"]
		if val == nil {
			o.TerminatedAt = number.Int64Pointer(number.ToInt64Any(nil))
		} else {
			o.TerminatedAt = number.Int64Pointer(number.ToInt64Any(val))
		}
	}
	if val, ok := kv["cost_center_id"].(*string); ok {
		o.CostCenterID = val
	} else if val, ok := kv["cost_center_id"].(string); ok {
		o.CostCenterID = &val
	} else {
		val := kv["cost_center_id"]
		if val == nil {
			o.CostCenterID = pstrings.Pointer("")
		} else {
			o.CostCenterID = pstrings.Pointer(fmt.Sprintf("%v", val))
		}
	}
	if val, ok := kv["team_id"].(*string); ok {
		o.TeamID = val
	} else if val, ok := kv["team_id"].(string); ok {
		o.TeamID = &val
	} else {
		val := kv["team_id"]
		if val == nil {
			o.TeamID = pstrings.Pointer("")
		} else {
			o.TeamID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	args = append(args, o.Email)
	args = append(args, o.Title)
	args = append(args, o.Location)
	args = append(args, o.AvatarURL)
	args = append(args, o.ManagerID)
	args = append(args, o.Active)
	args = append(args, o.Trackable)
	args = append(args, o.CreatedAt)
	args = append(args, o.UpdatedAt)
	args = append(args, o.DeletedAt)
	args = append(args, o.HiredAt)
	args = append(args, o.TerminatedAt)
	args = append(args, o.CostCenterID)
	args = append(args, o.TeamID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateUser creates a new User in the database
func CreateUser(ctx context.Context, db Db, o *User) error {
	return db.Create(ctx, o.ToMap())
}

// DeleteUser deletes a User in the database
func DeleteUser(ctx context.Context, db Db, o *User) error {
	return db.Delete(ctx, o.GetID())
}

// UpdateUser updates a User in the database
func UpdateUser(ctx context.Context, db Db, o *User) error {
	return db.Update(ctx, o.GetID(), o.ToMap())
}

// FindUser returns a User from the database
func FindUser(ctx context.Context, db Db, id string) (*User, error) {
	kv, err := db.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	var result User
	result.FromMap(kv)
	return &result, nil
}

// FindUsers returns all User from the database matching keys
func FindUsers(ctx context.Context, db Db, kv map[string]interface{}) ([]*User, error) {
	res, err := db.Find(ctx, kv)
	if err != nil {
		return nil, err
	}
	arr := make([]*User, 0)
	for _, kv := range res {
		var result User
		result.FromMap(kv)
		arr = append(arr, &result)
	}
	return arr, nil
}

// CreateUserAvroSchemaSpec creates the avro schema specification for User
func CreateUserAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "User",
		"connect.name": "customer.User",
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
				"name": "email",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "title",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "location",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "avatar_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "manager_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "active",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "trackable",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name":    "deleted_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "hired_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "terminated_ts",
				"type":    []interface{}{"null", "long"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "cost_center_id",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name":    "team_id",
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
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/user\\.json(\\.gz)?$"))
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
	fp := filepath.Join(dir, "/customer/user\\.json(\\.gz)?$")
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
