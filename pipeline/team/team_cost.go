// DO NOT EDIT -- generated code

// Package team - team cost pipeline
package team

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
	"sync"
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
	// TeamCostTopic is the default topic name
	TeamCostTopic datamodel.TopicNameType = "pipeline_team_TeamCost_topic"

	// TeamCostStream is the default stream name
	TeamCostStream datamodel.TopicNameType = "pipeline_team_TeamCost_stream"

	// TeamCostTable is the default table name
	TeamCostTable datamodel.TopicNameType = "pipeline_team_TeamCost"

	// TeamCostModelName is the model name
	TeamCostModelName datamodel.ModelNameType = "pipeline.team.TeamCost"
)

const (
	// TeamCostCostColumn is the cost column name
	TeamCostCostColumn = "cost"
	// TeamCostCountColumn is the count column name
	TeamCostCountColumn = "count"
	// TeamCostCustomerIDColumn is the customer_id column name
	TeamCostCustomerIDColumn = "customer_id"
	// TeamCostIDColumn is the id column name
	TeamCostIDColumn = "id"
	// TeamCostRefIDColumn is the ref_id column name
	TeamCostRefIDColumn = "ref_id"
	// TeamCostRefTypeColumn is the ref_type column name
	TeamCostRefTypeColumn = "ref_type"
	// TeamCostTeamIDColumn is the team_id column name
	TeamCostTeamIDColumn = "team_id"
	// TeamCostUpdatedDateColumn is the updated_date column name
	TeamCostUpdatedDateColumn = "updated_date"
	// TeamCostUpdatedDateColumnEpochColumn is the epoch column property of the UpdatedDate name
	TeamCostUpdatedDateColumnEpochColumn = "updated_date->epoch"
	// TeamCostUpdatedDateColumnOffsetColumn is the offset column property of the UpdatedDate name
	TeamCostUpdatedDateColumnOffsetColumn = "updated_date->offset"
	// TeamCostUpdatedDateColumnRfc3339Column is the rfc3339 column property of the UpdatedDate name
	TeamCostUpdatedDateColumnRfc3339Column = "updated_date->rfc3339"
)

// TeamCostUpdatedDate represents the object structure for updated_date
type TeamCostUpdatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toTeamCostUpdatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toTeamCostUpdatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *TeamCostUpdatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *TeamCostUpdatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toTeamCostUpdatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toTeamCostUpdatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toTeamCostUpdatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *TeamCostUpdatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *TeamCostUpdatedDate) FromMap(kv map[string]interface{}) {

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Epoch = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["offset"].(int64); ok {
		o.Offset = val
	} else {
		if val, ok := kv["offset"]; ok {
			if val == nil {
				o.Offset = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Offset = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["rfc3339"].(string); ok {
		o.Rfc3339 = val
	} else {
		if val, ok := kv["rfc3339"]; ok {
			if val == nil {
				o.Rfc3339 = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// TeamCost team cost
type TeamCost struct {
	// Cost the cost of the team
	Cost float64 `json:"cost" bson:"cost" yaml:"cost" faker:"-"`
	// Count the team count
	Count int64 `json:"count" bson:"count" yaml:"count" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// TeamID ref_id of the team
	TeamID string `json:"team_id" bson:"team_id" yaml:"team_id" faker:"-"`
	// UpdatedDate date object
	UpdatedDate TeamCostUpdatedDate `json:"updated_date" bson:"updated_date" yaml:"updated_date" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*TeamCost)(nil)

func toTeamCostObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toTeamCostObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *TeamCost:
		return v.ToMap(isavro)

	case TeamCostUpdatedDate:
		return v.ToMap(isavro)
	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of TeamCost
func (o *TeamCost) String() string {
	return fmt.Sprintf("pipeline.team.TeamCost<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *TeamCost) GetTopicName() datamodel.TopicNameType {
	return TeamCostTopic
}

// GetModelName returns the name of the model
func (o *TeamCost) GetModelName() datamodel.ModelNameType {
	return TeamCostModelName
}

func (o *TeamCost) setDefaults(frommap bool) {

	o.GetID()

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *TeamCost) GetID() string {
	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.TeamID)
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *TeamCost) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *TeamCost) GetTimestamp() time.Time {
	var dt interface{} = o.UpdatedDate
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
	case TeamCostUpdatedDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for TeamCost")
}

// GetRefID returns the RefID for the object
func (o *TeamCost) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *TeamCost) IsMaterialized() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *TeamCost) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	idletime, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	return &datamodel.ModelMaterializeConfig{
		KeyName:   "id",
		TableName: "pipeline_team_teamcost",
		BatchSize: 5000,
		IdleTime:  idletime,
	}
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *TeamCost) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *TeamCost) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = TeamCostModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *TeamCost) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "id",
		Timestamp:         "updated_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *TeamCost) GetStateKey() string {
	key := "id"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *TeamCost) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of TeamCost
func (o *TeamCost) Clone() datamodel.Model {
	c := new(TeamCost)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *TeamCost) Anon() datamodel.Model {
	c := new(TeamCost)
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
func (o *TeamCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *TeamCost) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecTeamCost *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *TeamCost) GetAvroCodec() *goavro.Codec {
	if cachedCodecTeamCost == nil {
		c, err := GetTeamCostAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecTeamCost = c
	}
	return cachedCodecTeamCost
}

// ToAvroBinary returns the data as Avro binary data
func (o *TeamCost) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *TeamCost) FromAvroBinary(value []byte) error {
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
func (o *TeamCost) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two TeamCost objects are equal
func (o *TeamCost) IsEqual(other *TeamCost) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *TeamCost) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(true)
	return map[string]interface{}{
		"cost":         toTeamCostObject(o.Cost, isavro, false, "float"),
		"count":        toTeamCostObject(o.Count, isavro, false, "long"),
		"customer_id":  toTeamCostObject(o.CustomerID, isavro, false, "string"),
		"id":           toTeamCostObject(o.ID, isavro, false, "string"),
		"ref_id":       toTeamCostObject(o.RefID, isavro, false, "string"),
		"ref_type":     toTeamCostObject(o.RefType, isavro, false, "string"),
		"team_id":      toTeamCostObject(o.TeamID, isavro, false, "string"),
		"updated_date": toTeamCostObject(o.UpdatedDate, isavro, false, "updated_date"),
		"hashcode":     toTeamCostObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *TeamCost) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["cost"].(float64); ok {
		o.Cost = val
	} else {
		if val, ok := kv["cost"]; ok {
			if val == nil {
				o.Cost = number.ToFloat64Any(nil)
			} else {
				o.Cost = number.ToFloat64Any(val)
			}
		}
	}

	if val, ok := kv["count"].(int64); ok {
		o.Count = val
	} else {
		if val, ok := kv["count"]; ok {
			if val == nil {
				o.Count = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Count = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		if val, ok := kv["customer_id"]; ok {
			if val == nil {
				o.CustomerID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["team_id"].(string); ok {
		o.TeamID = val
	} else {
		if val, ok := kv["team_id"]; ok {
			if val == nil {
				o.TeamID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.TeamID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["updated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.UpdatedDate.FromMap(kv)
		} else if sv, ok := val.(TeamCostUpdatedDate); ok {
			// struct
			o.UpdatedDate = sv
		} else if sp, ok := val.(*TeamCostUpdatedDate); ok {
			// struct pointer
			o.UpdatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.UpdatedDate.Epoch = dt.Epoch
			o.UpdatedDate.Rfc3339 = dt.Rfc3339
			o.UpdatedDate.Offset = dt.Offset
		}
	} else {
		o.UpdatedDate.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *TeamCost) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Cost)
	args = append(args, o.Count)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.TeamID)
	args = append(args, o.UpdatedDate)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetTeamCostAvroSchemaSpec creates the avro schema specification for TeamCost
func GetTeamCostAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "pipeline.team",
		"name":      "TeamCost",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "cost",
				"type": "float",
			},
			map[string]interface{}{
				"name": "count",
				"type": "long",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
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
				"name": "team_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "updated_date",
				"type": map[string]interface{}{"doc": "date object", "type": "record", "name": "updated_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"doc": "the date in RFC3339 format", "type": "string", "name": "rfc3339"}}},
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *TeamCost) GetEventAPIConfig() datamodel.EventAPIConfig {
	return datamodel.EventAPIConfig{
		Publish: datamodel.EventAPIPublish{
			Public: false,
		},
		Subscribe: datamodel.EventAPISubscribe{
			Public: false,
			Key:    "",
		},
	}
}

// GetTeamCostAvroSchema creates the avro schema for TeamCost
func GetTeamCostAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetTeamCostAvroSchemaSpec())
}

// TransformTeamCostFunc is a function for transforming TeamCost during processing
type TransformTeamCostFunc func(input *TeamCost) (*TeamCost, error)

// NewTeamCostPipe creates a pipe for processing TeamCost items
func NewTeamCostPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformTeamCostFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewTeamCostInputStream(input, errors)
	var stream chan TeamCost
	if len(transforms) > 0 {
		stream = make(chan TeamCost, 1000)
	} else {
		stream = inch
	}
	outdone := NewTeamCostOutputStream(output, stream, errors)
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

// NewTeamCostInputStreamDir creates a channel for reading TeamCost as JSON newlines from a directory of files
func NewTeamCostInputStreamDir(dir string, errors chan<- error, transforms ...TransformTeamCostFunc) (chan TeamCost, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/pipeline.team/team_cost\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan TeamCost)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for team_cost")
		ch := make(chan TeamCost)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewTeamCostInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan TeamCost)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewTeamCostInputStreamFile creates an channel for reading TeamCost as JSON newlines from filename
func NewTeamCostInputStreamFile(filename string, errors chan<- error, transforms ...TransformTeamCostFunc) (chan TeamCost, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan TeamCost)
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
			ch := make(chan TeamCost)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewTeamCostInputStream(f, errors, transforms...)
}

// NewTeamCostInputStream creates an channel for reading TeamCost as JSON newlines from stream
func NewTeamCostInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformTeamCostFunc) (chan TeamCost, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan TeamCost, 1000)
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
			var item TeamCost
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

// NewTeamCostOutputStreamDir will output json newlines from channel and save in dir
func NewTeamCostOutputStreamDir(dir string, ch chan TeamCost, errors chan<- error, transforms ...TransformTeamCostFunc) <-chan bool {
	fp := filepath.Join(dir, "/pipeline.team/team_cost\\.json(\\.gz)?$")
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
	return NewTeamCostOutputStream(gz, ch, errors, transforms...)
}

// NewTeamCostOutputStream will output json newlines from channel to the stream
func NewTeamCostOutputStream(stream io.WriteCloser, ch chan TeamCost, errors chan<- error, transforms ...TransformTeamCostFunc) <-chan bool {
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

// TeamCostSendEvent is an event detail for sending data
type TeamCostSendEvent struct {
	TeamCost *TeamCost
	headers  map[string]string
	time     time.Time
	key      string
}

var _ datamodel.ModelSendEvent = (*TeamCostSendEvent)(nil)

// Key is the key to use for the message
func (e *TeamCostSendEvent) Key() string {
	if e.key == "" {
		return e.TeamCost.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *TeamCostSendEvent) Object() datamodel.Model {
	return e.TeamCost
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *TeamCostSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *TeamCostSendEvent) Timestamp() time.Time {
	return e.time
}

// TeamCostSendEventOpts is a function handler for setting opts
type TeamCostSendEventOpts func(o *TeamCostSendEvent)

// WithTeamCostSendEventKey sets the key value to a value different than the object ID
func WithTeamCostSendEventKey(key string) TeamCostSendEventOpts {
	return func(o *TeamCostSendEvent) {
		o.key = key
	}
}

// WithTeamCostSendEventTimestamp sets the timestamp value
func WithTeamCostSendEventTimestamp(tv time.Time) TeamCostSendEventOpts {
	return func(o *TeamCostSendEvent) {
		o.time = tv
	}
}

// WithTeamCostSendEventHeader sets the timestamp value
func WithTeamCostSendEventHeader(key, value string) TeamCostSendEventOpts {
	return func(o *TeamCostSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewTeamCostSendEvent returns a new TeamCostSendEvent instance
func NewTeamCostSendEvent(o *TeamCost, opts ...TeamCostSendEventOpts) *TeamCostSendEvent {
	res := &TeamCostSendEvent{
		TeamCost: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewTeamCostProducer will stream data from the channel
func NewTeamCostProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*TeamCost); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type pipeline.team.TeamCost but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewTeamCostConsumer will stream data from the topic into the provided channel
func NewTeamCostConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object TeamCost
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into pipeline.team.TeamCost: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into pipeline.team.TeamCost: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for pipeline.team.TeamCost")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &TeamCostReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object TeamCost
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &TeamCostReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// TeamCostReceiveEvent is an event detail for receiving data
type TeamCostReceiveEvent struct {
	TeamCost *TeamCost
	message  eventing.Message
	eof      bool
}

var _ datamodel.ModelReceiveEvent = (*TeamCostReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *TeamCostReceiveEvent) Object() datamodel.Model {
	return e.TeamCost
}

// Message returns the underlying message data for the event
func (e *TeamCostReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *TeamCostReceiveEvent) EOF() bool {
	return e.eof
}

// TeamCostProducer implements the datamodel.ModelEventProducer
type TeamCostProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*TeamCostProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *TeamCostProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *TeamCostProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *TeamCost) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *TeamCost) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &TeamCostProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewTeamCostProducer(newctx, producer, ch, errors, empty),
	}
}

// NewTeamCostProducerChannel returns a channel which can be used for producing Model events
func NewTeamCostProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewTeamCostProducerChannelSize(producer, 0, errors)
}

// NewTeamCostProducerChannelSize returns a channel which can be used for producing Model events
func NewTeamCostProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &TeamCostProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewTeamCostProducer(newctx, producer, ch, errors, empty),
	}
}

// TeamCostConsumer implements the datamodel.ModelEventConsumer
type TeamCostConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*TeamCostConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *TeamCostConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *TeamCostConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *TeamCost) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &TeamCostConsumer{
		ch:       ch,
		callback: NewTeamCostConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewTeamCostConsumerChannel returns a consumer channel which can be used to consume Model events
func NewTeamCostConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &TeamCostConsumer{
		ch:       ch,
		callback: NewTeamCostConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
