// DO NOT EDIT -- generated code

// Package auth - the system which contains customer and user authentication and authorization
package auth

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
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
)

const (
	// ACLGrantTopic is the default topic name
	ACLGrantTopic datamodel.TopicNameType = "auth_ACLGrant_topic"

	// ACLGrantStream is the default stream name
	ACLGrantStream datamodel.TopicNameType = "auth_ACLGrant_stream"

	// ACLGrantTable is the default table name
	ACLGrantTable datamodel.TopicNameType = "auth_ACLGrant"

	// ACLGrantModelName is the model name
	ACLGrantModelName datamodel.ModelNameType = "auth.ACLGrant"
)

const (
	// ACLGrantIDColumn is the id column name
	ACLGrantIDColumn = "id"
	// ACLGrantRefIDColumn is the ref_id column name
	ACLGrantRefIDColumn = "ref_id"
	// ACLGrantRefTypeColumn is the ref_type column name
	ACLGrantRefTypeColumn = "ref_type"
	// ACLGrantCustomerIDColumn is the customer_id column name
	ACLGrantCustomerIDColumn = "customer_id"
	// ACLGrantCreatedAtColumn is the created_ts column name
	ACLGrantCreatedAtColumn = "created_ts"
	// ACLGrantUpdatedAtColumn is the updated_ts column name
	ACLGrantUpdatedAtColumn = "updated_ts"
	// ACLGrantResourceIDColumn is the resource_id column name
	ACLGrantResourceIDColumn = "resource_id"
	// ACLGrantRoleIDColumn is the role_id column name
	ACLGrantRoleIDColumn = "role_id"
	// ACLGrantPermissionColumn is the permission column name
	ACLGrantPermissionColumn = "permission"
)

// Permission is the enumeration type for permission
type Permission int32

// String returns the string value for Permission
func (v Permission) String() string {
	switch int32(v) {
	case 0:
		return "admin"
	case 1:
		return "read"
	case 2:
		return "readwrite"
	}
	return "unset"
}

const (
	// PermissionAdmin is the enumeration value for admin
	PermissionAdmin Permission = 0
	// PermissionRead is the enumeration value for read
	PermissionRead Permission = 1
	// PermissionReadwrite is the enumeration value for readwrite
	PermissionReadwrite Permission = 2
)

// ACLGrant a ACL grant controls permisson for a specific role to a resource
type ACLGrant struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`

	// custom types

	// ResourceID the resource id for the grant
	ResourceID string `json:"resource_id" bson:"resource_id" yaml:"resource_id" faker:"-"`
	// RoleID the role id for the grant
	RoleID string `json:"role_id" bson:"role_id" yaml:"role_id" faker:"-"`
	// Permission the permission that the grant provides
	Permission Permission `json:"permission" bson:"permission" yaml:"permission" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ACLGrant)(nil)

func toACLGrantObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toACLGrantObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toACLGrantObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toACLGrantObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toACLGrantObjectNil(isavro, isoptional)
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
	case *ACLGrant:
		return v.ToMap()
	case ACLGrant:
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
			arr = append(arr, toACLGrantObject(av, isavro, false, ""))
		}
		return arr
	case Permission:
		if !isavro {
			return (o.(Permission)).String()
		}
		return map[string]string{
			"auth.permission": (o.(Permission)).String(),
		}
	case *Permission:
		if !isavro {
			return (o.(*Permission)).String()
		}
		return map[string]string{
			"auth.permission": (o.(*Permission)).String(),
		}
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of ACLGrant
func (o *ACLGrant) String() string {
	return fmt.Sprintf("auth.ACLGrant<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ACLGrant) GetTopicName() datamodel.TopicNameType {
	return ACLGrantTopic
}

// GetModelName returns the name of the model
func (o *ACLGrant) GetModelName() datamodel.ModelNameType {
	return ACLGrantModelName
}

func (o *ACLGrant) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ACLGrant) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.ResourceID, o.RoleID, o.Permission)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ACLGrant) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ACLGrant) GetTimestamp() time.Time {
	var dt interface{} = o.CreatedAt
	switch v := dt.(type) {
	case int64:
		return datetime.DateFromEpoch(v).UTC()
	case string:
		tv, err := datetime.ISODateToTime(v)
		if err != nil {
			panic(err)
		}
		return tv.UTC()
	case time.Time:
		return v.UTC()
	}
	panic("not sure how to handle the date time format for ACLGrant")
}

// GetRefID returns the RefID for the object
func (o *ACLGrant) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ACLGrant) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ACLGrant) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "auth_aclgrant",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ACLGrant) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ACLGrant) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ACLGrantModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ACLGrant) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "created_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetCustomerID will return the customer_id
func (o *ACLGrant) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of ACLGrant
func (o *ACLGrant) Clone() datamodel.Model {
	c := new(ACLGrant)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ACLGrant) Anon() datamodel.Model {
	c := new(ACLGrant)
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
func (o *ACLGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ACLGrant) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecACLGrant *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ACLGrant) GetAvroCodec() *goavro.Codec {
	if cachedCodecACLGrant == nil {
		c, err := GetACLGrantAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecACLGrant = c
	}
	return cachedCodecACLGrant
}

// ToAvroBinary returns the data as Avro binary data
func (o *ACLGrant) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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

// Stringify returns the object in JSON format as a string
func (o *ACLGrant) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ACLGrant objects are equal
func (o *ACLGrant) IsEqual(other *ACLGrant) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ACLGrant) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	return map[string]interface{}{
		"id":          o.GetID(),
		"ref_id":      o.GetRefID(),
		"ref_type":    o.RefType,
		"customer_id": o.CustomerID,
		"hashcode":    o.Hash(),
		"created_ts":  o.CreatedAt,
		"updated_ts":  o.UpdatedAt,
		"resource_id": toACLGrantObject(o.ResourceID, isavro, false, "string"),
		"role_id":     toACLGrantObject(o.RoleID, isavro, false, "string"),
		"permission":  toACLGrantObject(o.Permission, isavro, false, "permission"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ACLGrant) FromMap(kv map[string]interface{}) {
	// make sure that these have values if empty
	o.setDefaults()
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
	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else if val, ok := kv["created_ts"].(time.Time); ok {
		o.CreatedAt = datetime.TimeToEpoch(val)
	}
	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else if val, ok := kv["updated_ts"].(time.Time); ok {
		o.UpdatedAt = datetime.TimeToEpoch(val)
	}
	if val, ok := kv["resource_id"].(string); ok {
		o.ResourceID = val
	} else {
		val := kv["resource_id"]
		if val == nil {
			o.ResourceID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ResourceID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["role_id"].(string); ok {
		o.RoleID = val
	} else {
		val := kv["role_id"]
		if val == nil {
			o.RoleID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RoleID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["permission"].(Permission); ok {
		o.Permission = val
	} else {
		if em, ok := kv["permission"].(map[string]interface{}); ok {
			ev := em["auth.permission"].(string)
			switch ev {
			case "admin":
				o.Permission = 0
			case "read":
				o.Permission = 1
			case "readwrite":
				o.Permission = 2
			}
		}
		if em, ok := kv["permission"].(string); ok {
			switch em {
			case "admin":
				o.Permission = 0
			case "read":
				o.Permission = 1
			case "readwrite":
				o.Permission = 2
			}
		}
	}
}

// Hash will return a hashcode for the object
func (o *ACLGrant) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.ResourceID)
	args = append(args, o.RoleID)
	args = append(args, o.Permission)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateACLGrant creates a new ACLGrant in the database
func CreateACLGrant(ctx context.Context, db datamodel.Storage, o *ACLGrant) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteACLGrant deletes a ACLGrant in the database
func DeleteACLGrant(ctx context.Context, db datamodel.Storage, o *ACLGrant) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateACLGrant updates a ACLGrant in the database
func UpdateACLGrant(ctx context.Context, db datamodel.Storage, o *ACLGrant) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindACLGrant returns a ACLGrant from the database
func FindACLGrant(ctx context.Context, db datamodel.Storage, id string) (*ACLGrant, error) {
	kv, err := db.FindOne(ctx, ACLGrantModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*ACLGrant), nil
}

// FindACLGrants returns all ACLGrant from the database matching keys
func FindACLGrants(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*ACLGrant, error) {
	res, err := db.Find(ctx, ACLGrantModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*ACLGrant, 0)
		for _, m := range res {
			arr = append(arr, m.(*ACLGrant))
		}
		return arr, nil
	}
	return nil, nil
}

// GetACLGrantAvroSchemaSpec creates the avro schema specification for ACLGrant
func GetACLGrantAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "auth",
		"name":      "ACLGrant",
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
				"name": "created_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "resource_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "role_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "permission",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "permission",
						"symbols": []interface{}{"admin", "read", "readwrite"},
					},
				},
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetACLGrantAvroSchema creates the avro schema for ACLGrant
func GetACLGrantAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetACLGrantAvroSchemaSpec())
}

// TransformACLGrantFunc is a function for transforming ACLGrant during processing
type TransformACLGrantFunc func(input *ACLGrant) (*ACLGrant, error)

// NewACLGrantPipe creates a pipe for processing ACLGrant items
func NewACLGrantPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformACLGrantFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewACLGrantInputStream(input, errors)
	var stream chan ACLGrant
	if len(transforms) > 0 {
		stream = make(chan ACLGrant, 1000)
	} else {
		stream = inch
	}
	outdone := NewACLGrantOutputStream(output, stream, errors)
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

// NewACLGrantInputStreamDir creates a channel for reading ACLGrant as JSON newlines from a directory of files
func NewACLGrantInputStreamDir(dir string, errors chan<- error, transforms ...TransformACLGrantFunc) (chan ACLGrant, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/auth/acl_grant\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ACLGrant)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for acl_grant")
		ch := make(chan ACLGrant)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewACLGrantInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ACLGrant)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewACLGrantInputStreamFile creates an channel for reading ACLGrant as JSON newlines from filename
func NewACLGrantInputStreamFile(filename string, errors chan<- error, transforms ...TransformACLGrantFunc) (chan ACLGrant, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ACLGrant)
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
			ch := make(chan ACLGrant)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewACLGrantInputStream(f, errors, transforms...)
}

// NewACLGrantInputStream creates an channel for reading ACLGrant as JSON newlines from stream
func NewACLGrantInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformACLGrantFunc) (chan ACLGrant, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ACLGrant, 1000)
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
			var item ACLGrant
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

// NewACLGrantOutputStreamDir will output json newlines from channel and save in dir
func NewACLGrantOutputStreamDir(dir string, ch chan ACLGrant, errors chan<- error, transforms ...TransformACLGrantFunc) <-chan bool {
	fp := filepath.Join(dir, "/auth/acl_grant\\.json(\\.gz)?$")
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
	return NewACLGrantOutputStream(gz, ch, errors, transforms...)
}

// NewACLGrantOutputStream will output json newlines from channel to the stream
func NewACLGrantOutputStream(stream io.WriteCloser, ch chan ACLGrant, errors chan<- error, transforms ...TransformACLGrantFunc) <-chan bool {
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

// ACLGrantSendEvent is an event detail for sending data
type ACLGrantSendEvent struct {
	ACLGrant *ACLGrant
	headers  map[string]string
	time     time.Time
	key      string
}

var _ datamodel.ModelSendEvent = (*ACLGrantSendEvent)(nil)

// Key is the key to use for the message
func (e *ACLGrantSendEvent) Key() string {
	if e.key == "" {
		return e.ACLGrant.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ACLGrantSendEvent) Object() datamodel.Model {
	return e.ACLGrant
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ACLGrantSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ACLGrantSendEvent) Timestamp() time.Time {
	return e.time
}

// ACLGrantSendEventOpts is a function handler for setting opts
type ACLGrantSendEventOpts func(o *ACLGrantSendEvent)

// WithACLGrantSendEventKey sets the key value to a value different than the object ID
func WithACLGrantSendEventKey(key string) ACLGrantSendEventOpts {
	return func(o *ACLGrantSendEvent) {
		o.key = key
	}
}

// WithACLGrantSendEventTimestamp sets the timestamp value
func WithACLGrantSendEventTimestamp(tv time.Time) ACLGrantSendEventOpts {
	return func(o *ACLGrantSendEvent) {
		o.time = tv
	}
}

// WithACLGrantSendEventHeader sets the timestamp value
func WithACLGrantSendEventHeader(key, value string) ACLGrantSendEventOpts {
	return func(o *ACLGrantSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewACLGrantSendEvent returns a new ACLGrantSendEvent instance
func NewACLGrantSendEvent(o *ACLGrant, opts ...ACLGrantSendEventOpts) *ACLGrantSendEvent {
	res := &ACLGrantSendEvent{
		ACLGrant: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewACLGrantProducer will stream data from the channel
func NewACLGrantProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*ACLGrant); ok {
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
				errors <- fmt.Errorf("invalid event received. expected an object of type auth.ACLGrant but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewACLGrantConsumer will stream data from the topic into the provided channel
func NewACLGrantConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object ACLGrant
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into auth.ACLGrant: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ACLGrantReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object ACLGrant
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ACLGrantReceiveEvent{nil, msg, true}
		},
	})
}

// ACLGrantReceiveEvent is an event detail for receiving data
type ACLGrantReceiveEvent struct {
	ACLGrant *ACLGrant
	message  eventing.Message
	eof      bool
}

var _ datamodel.ModelReceiveEvent = (*ACLGrantReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ACLGrantReceiveEvent) Object() datamodel.Model {
	return e.ACLGrant
}

// Message returns the underlying message data for the event
func (e *ACLGrantReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ACLGrantReceiveEvent) EOF() bool {
	return e.eof
}

// ACLGrantProducer implements the datamodel.ModelEventProducer
type ACLGrantProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*ACLGrantProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ACLGrantProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ACLGrantProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *ACLGrant) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ACLGrantProducer{
		ch:   ch,
		done: NewACLGrantProducer(producer, ch, errors),
	}
}

// NewACLGrantProducerChannel returns a channel which can be used for producing Model events
func NewACLGrantProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &ACLGrantProducer{
		ch:   ch,
		done: NewACLGrantProducer(producer, ch, errors),
	}
}

// ACLGrantConsumer implements the datamodel.ModelEventConsumer
type ACLGrantConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*ACLGrantConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ACLGrantConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ACLGrantConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *ACLGrant) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewACLGrantConsumer(consumer, ch, errors)
	return &ACLGrantConsumer{
		ch: ch,
	}
}

// NewACLGrantConsumerChannel returns a consumer channel which can be used to consume Model events
func NewACLGrantConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewACLGrantConsumer(consumer, ch, errors)
	return &ACLGrantConsumer{
		ch: ch,
	}
}
