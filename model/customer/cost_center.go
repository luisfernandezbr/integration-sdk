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
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// CostCenterTopic is the default topic name
	CostCenterTopic datamodel.TopicNameType = "customer_CostCenter_topic"

	// CostCenterStream is the default stream name
	CostCenterStream datamodel.TopicNameType = "customer_CostCenter_stream"

	// CostCenterTable is the default table name
	CostCenterTable datamodel.TopicNameType = "customer_CostCenter"

	// CostCenterModelName is the model name
	CostCenterModelName datamodel.ModelNameType = "customer.CostCenter"
)

const (
	// CostCenterIDColumn is the id column name
	CostCenterIDColumn = "id"
	// CostCenterRefIDColumn is the ref_id column name
	CostCenterRefIDColumn = "ref_id"
	// CostCenterRefTypeColumn is the ref_type column name
	CostCenterRefTypeColumn = "ref_type"
	// CostCenterCustomerIDColumn is the customer_id column name
	CostCenterCustomerIDColumn = "customer_id"
	// CostCenterCreatedAtColumn is the created_ts column name
	CostCenterCreatedAtColumn = "created_ts"
	// CostCenterUpdatedAtColumn is the updated_ts column name
	CostCenterUpdatedAtColumn = "updated_ts"
	// CostCenterNameColumn is the name column name
	CostCenterNameColumn = "name"
	// CostCenterDescriptionColumn is the description column name
	CostCenterDescriptionColumn = "description"
	// CostCenterCostColumn is the cost column name
	CostCenterCostColumn = "cost"
)

// CostCenter a cost center represents information about users and their cost
type CostCenter struct {
	// built in types

	ID         string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	RefID      string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	RefType    string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	Hashcode   string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`

	CreatedAt int64 `json:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`

	// custom types

	// Name the name of the cost center
	Name string `json:"name" bson:"name" yaml:"name" faker:"costcenter"`
	// Description the description for the cost center
	Description string `json:"description" bson:"description" yaml:"description" faker:"-"`
	// Cost the cost value of the cost center
	Cost float64 `json:"cost" bson:"cost" yaml:"cost" faker:"salary"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CostCenter)(nil)

func toCostCenterObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCostCenterObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toCostCenterObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toCostCenterObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toCostCenterObjectNil(isavro, isoptional)
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
	case *CostCenter:
		return v.ToMap()
	case CostCenter:
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
			arr = append(arr, toCostCenterObject(av, isavro, false, ""))
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of CostCenter
func (o *CostCenter) String() string {
	return fmt.Sprintf("customer.CostCenter<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CostCenter) GetTopicName() datamodel.TopicNameType {
	return CostCenterTopic
}

// GetModelName returns the name of the model
func (o *CostCenter) GetModelName() datamodel.ModelNameType {
	return CostCenterModelName
}

func (o *CostCenter) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CostCenter) GetID() string {
	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, randomString(64))
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CostCenter) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CostCenter) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CostCenter) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CostCenter) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CostCenter) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "customer_costcenter",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CostCenter) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CostCenter) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CostCenterModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CostCenter) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *CostCenter) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of CostCenter
func (o *CostCenter) Clone() datamodel.Model {
	c := new(CostCenter)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CostCenter) Anon() datamodel.Model {
	c := new(CostCenter)
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
func (o *CostCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CostCenter) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	// make sure that these have values if empty
	o.setDefaults()
	o.FromMap(kv)
	return nil
}

var cachedCodecCostCenter *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *CostCenter) GetAvroCodec() *goavro.Codec {
	if cachedCodecCostCenter == nil {
		c, err := GetCostCenterAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCostCenter = c
	}
	return cachedCodecCostCenter
}

// ToAvroBinary returns the data as Avro binary data
func (o *CostCenter) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CostCenter) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CostCenter objects are equal
func (o *CostCenter) IsEqual(other *CostCenter) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CostCenter) ToMap(avro ...bool) map[string]interface{} {
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
		"name":        toCostCenterObject(o.Name, isavro, false, "string"),
		"description": toCostCenterObject(o.Description, isavro, false, "string"),
		"cost":        toCostCenterObject(o.Cost, isavro, false, "float"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CostCenter) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		val := kv["description"]
		if val == nil {
			o.Description = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Description = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["cost"].(float64); ok {
		o.Cost = val
	} else {
		val := kv["cost"]
		if val == nil {
			o.Cost = number.ToFloat64Any(nil)
		} else {
			o.Cost = number.ToFloat64Any(val)
		}
	}
}

// Hash will return a hashcode for the object
func (o *CostCenter) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.Name)
	args = append(args, o.Description)
	args = append(args, o.Cost)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// CreateCostCenter creates a new CostCenter in the database
func CreateCostCenter(ctx context.Context, db datamodel.Storage, o *CostCenter) error {
	o.setDefaults()
	return db.Create(ctx, o)
}

// DeleteCostCenter deletes a CostCenter in the database
func DeleteCostCenter(ctx context.Context, db datamodel.Storage, o *CostCenter) error {
	o.setDefaults()
	return db.Delete(ctx, o)
}

// UpdateCostCenter updates a CostCenter in the database
func UpdateCostCenter(ctx context.Context, db datamodel.Storage, o *CostCenter) error {
	o.setDefaults()
	return db.Update(ctx, o)
}

// FindCostCenter returns a CostCenter from the database
func FindCostCenter(ctx context.Context, db datamodel.Storage, id string) (*CostCenter, error) {
	kv, err := db.FindOne(ctx, CostCenterModelName, id)
	if err != nil {
		return nil, err
	}
	if kv == nil {
		return nil, nil
	}
	return kv.(*CostCenter), nil
}

// FindCostCenters returns all CostCenter from the database matching keys
func FindCostCenters(ctx context.Context, db datamodel.Storage, kv map[string]interface{}) ([]*CostCenter, error) {
	res, err := db.Find(ctx, CostCenterModelName, kv)
	if err != nil {
		return nil, err
	}
	if res != nil {
		arr := make([]*CostCenter, 0)
		for _, m := range res {
			arr = append(arr, m.(*CostCenter))
		}
		return arr, nil
	}
	return nil, nil
}

// GetCostCenterAvroSchemaSpec creates the avro schema specification for CostCenter
func GetCostCenterAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":         "record",
		"namespace":    "customer",
		"name":         "CostCenter",
		"connect.name": "customer.CostCenter",
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
				"name": "name",
				"type": "string",
			},
			map[string]interface{}{
				"name": "description",
				"type": "string",
			},
			map[string]interface{}{
				"name": "cost",
				"type": "float",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetCostCenterAvroSchema creates the avro schema for CostCenter
func GetCostCenterAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCostCenterAvroSchemaSpec())
}

// TransformCostCenterFunc is a function for transforming CostCenter during processing
type TransformCostCenterFunc func(input *CostCenter) (*CostCenter, error)

// NewCostCenterPipe creates a pipe for processing CostCenter items
func NewCostCenterPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCostCenterFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCostCenterInputStream(input, errors)
	var stream chan CostCenter
	if len(transforms) > 0 {
		stream = make(chan CostCenter, 1000)
	} else {
		stream = inch
	}
	outdone := NewCostCenterOutputStream(output, stream, errors)
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

// NewCostCenterInputStreamDir creates a channel for reading CostCenter as JSON newlines from a directory of files
func NewCostCenterInputStreamDir(dir string, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/customer/cost_center\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for cost_center")
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCostCenterInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CostCenter)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCostCenterInputStreamFile creates an channel for reading CostCenter as JSON newlines from filename
func NewCostCenterInputStreamFile(filename string, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CostCenter)
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
			ch := make(chan CostCenter)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCostCenterInputStream(f, errors, transforms...)
}

// NewCostCenterInputStream creates an channel for reading CostCenter as JSON newlines from stream
func NewCostCenterInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCostCenterFunc) (chan CostCenter, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CostCenter, 1000)
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
			var item CostCenter
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

// NewCostCenterOutputStreamDir will output json newlines from channel and save in dir
func NewCostCenterOutputStreamDir(dir string, ch chan CostCenter, errors chan<- error, transforms ...TransformCostCenterFunc) <-chan bool {
	fp := filepath.Join(dir, "/customer/cost_center\\.json(\\.gz)?$")
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
	return NewCostCenterOutputStream(gz, ch, errors, transforms...)
}

// NewCostCenterOutputStream will output json newlines from channel to the stream
func NewCostCenterOutputStream(stream io.WriteCloser, ch chan CostCenter, errors chan<- error, transforms ...TransformCostCenterFunc) <-chan bool {
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

// CostCenterSendEvent is an event detail for sending data
type CostCenterSendEvent struct {
	CostCenter *CostCenter
	headers    map[string]string
	time       time.Time
	key        string
}

var _ datamodel.ModelSendEvent = (*CostCenterSendEvent)(nil)

// Key is the key to use for the message
func (e *CostCenterSendEvent) Key() string {
	if e.key == "" {
		return e.CostCenter.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CostCenterSendEvent) Object() datamodel.Model {
	return e.CostCenter
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CostCenterSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CostCenterSendEvent) Timestamp() time.Time {
	return e.time
}

// CostCenterSendEventOpts is a function handler for setting opts
type CostCenterSendEventOpts func(o *CostCenterSendEvent)

// WithCostCenterSendEventKey sets the key value to a value different than the object ID
func WithCostCenterSendEventKey(key string) CostCenterSendEventOpts {
	return func(o *CostCenterSendEvent) {
		o.key = key
	}
}

// WithCostCenterSendEventTimestamp sets the timestamp value
func WithCostCenterSendEventTimestamp(tv time.Time) CostCenterSendEventOpts {
	return func(o *CostCenterSendEvent) {
		o.time = tv
	}
}

// WithCostCenterSendEventHeader sets the timestamp value
func WithCostCenterSendEventHeader(key, value string) CostCenterSendEventOpts {
	return func(o *CostCenterSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCostCenterSendEvent returns a new CostCenterSendEvent instance
func NewCostCenterSendEvent(o *CostCenter, opts ...CostCenterSendEventOpts) *CostCenterSendEvent {
	res := &CostCenterSendEvent{
		CostCenter: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCostCenterProducer will stream data from the channel
func NewCostCenterProducer(producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		ctx := context.Background()
		for item := range ch {
			if object, ok := item.Object().(*CostCenter); ok {
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
				errors <- fmt.Errorf("invalid event received. expected an object of type customer.CostCenter but received on of type %v", reflect.TypeOf(item.Object()))
			}
		}
	}()
	return done
}

// NewCostCenterConsumer will stream data from the topic into the provided channel
func NewCostCenterConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) {
	consumer.Consume(&eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object CostCenter
			if err := json.Unmarshal(msg.Value, &object); err != nil {
				return fmt.Errorf("error unmarshaling json data into customer.CostCenter: %s", err)
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CostCenterReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object CostCenter
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CostCenterReceiveEvent{nil, msg, true}
		},
	})
}

// CostCenterReceiveEvent is an event detail for receiving data
type CostCenterReceiveEvent struct {
	CostCenter *CostCenter
	message    eventing.Message
	eof        bool
}

var _ datamodel.ModelReceiveEvent = (*CostCenterReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CostCenterReceiveEvent) Object() datamodel.Model {
	return e.CostCenter
}

// Message returns the underlying message data for the event
func (e *CostCenterReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CostCenterReceiveEvent) EOF() bool {
	return e.eof
}

// CostCenterProducer implements the datamodel.ModelEventProducer
type CostCenterProducer struct {
	ch   chan datamodel.ModelSendEvent
	done <-chan bool
}

var _ datamodel.ModelEventProducer = (*CostCenterProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CostCenterProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CostCenterProducer) Close() error {
	close(p.ch)
	<-p.done
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *CostCenter) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CostCenterProducer{
		ch:   ch,
		done: NewCostCenterProducer(producer, ch, errors),
	}
}

// NewCostCenterProducerChannel returns a channel which can be used for producing Model events
func NewCostCenterProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	return &CostCenterProducer{
		ch:   ch,
		done: NewCostCenterProducer(producer, ch, errors),
	}
}

// CostCenterConsumer implements the datamodel.ModelEventConsumer
type CostCenterConsumer struct {
	ch chan datamodel.ModelReceiveEvent
}

var _ datamodel.ModelEventConsumer = (*CostCenterConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CostCenterConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CostCenterConsumer) Close() error {
	close(c.ch)
	return nil
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *CostCenter) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCostCenterConsumer(consumer, ch, errors)
	return &CostCenterConsumer{
		ch: ch,
	}
}

// NewCostCenterConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCostCenterConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	NewCostCenterConsumer(consumer, ch, errors)
	return &CostCenterConsumer{
		ch: ch,
	}
}
