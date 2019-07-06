// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

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
)

const (
	// IntegrationRequestTopic is the default topic name
	IntegrationRequestTopic datamodel.TopicNameType = "agent_IntegrationRequest_topic"

	// IntegrationRequestStream is the default stream name
	IntegrationRequestStream datamodel.TopicNameType = "agent_IntegrationRequest_stream"

	// IntegrationRequestTable is the default table name
	IntegrationRequestTable datamodel.TopicNameType = "agent_IntegrationRequest"

	// IntegrationRequestModelName is the model name
	IntegrationRequestModelName datamodel.ModelNameType = "agent.IntegrationRequest"
)

const (
	// IntegrationRequestCustomerIDColumn is the customer_id column name
	IntegrationRequestCustomerIDColumn = "customer_id"
	// IntegrationRequestDateColumn is the date column name
	IntegrationRequestDateColumn = "date"
	// IntegrationRequestIDColumn is the id column name
	IntegrationRequestIDColumn = "id"
	// IntegrationRequestIntegrationColumn is the integration column name
	IntegrationRequestIntegrationColumn = "integration"
	// IntegrationRequestRefIDColumn is the ref_id column name
	IntegrationRequestRefIDColumn = "ref_id"
	// IntegrationRequestRefTypeColumn is the ref_type column name
	IntegrationRequestRefTypeColumn = "ref_type"
	// IntegrationRequestUUIDColumn is the uuid column name
	IntegrationRequestUUIDColumn = "uuid"
)

// IntegrationRequestAuthorization represents the object structure for authorization
type IntegrationRequestAuthorization struct {
	// AccessToken Access token
	AccessToken *string `json:"access_token" bson:"access_token" yaml:"access_token" faker:"-"`
	// APIToken API Token for instance, if relevant
	APIToken *string `json:"api_token" bson:"api_token" yaml:"api_token" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Errored If authorization failed by the agent
	Errored bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Password Password for instance, if relevant
	Password *string `json:"password" bson:"password" yaml:"password" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken *string `json:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url" bson:"url" yaml:"url" faker:"-"`
	// Username Username for instance, if relevant
	Username *string `json:"username" bson:"username" yaml:"username" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedAt Timestamp when validated
	ValidatedAt int64 `json:"validated_ts" bson:"validated_ts" yaml:"validated_ts" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func (o *IntegrationRequestAuthorization) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": o.AccessToken,
		// APIToken API Token for instance, if relevant
		"api_token": o.APIToken,
		// Authorization the agents encrypted authorization
		"authorization": o.Authorization,
		// Errored If authorization failed by the agent
		"errored": o.Errored,
		// Password Password for instance, if relevant
		"password": o.Password,
		// RefreshToken Refresh token
		"refresh_token": o.RefreshToken,
		// URL URL of instance if relevant
		"url": o.URL,
		// Username Username for instance, if relevant
		"username": o.Username,
		// Validated If the validation has been run against this instance
		"validated": o.Validated,
		// ValidatedAt Timestamp when validated
		"validated_ts": o.ValidatedAt,
		// ValidationMessage The validation message from the agent
		"validation_message": o.ValidationMessage,
	}
}

// IntegrationRequestDate represents the object structure for date
type IntegrationRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *IntegrationRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// IntegrationRequestIntegration represents the object structure for integration
type IntegrationRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization IntegrationRequestAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location string `json:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress IntegrationRequestProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
}

func (o *IntegrationRequestIntegration) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": o.Active,
		// Authorization Authorization information
		"authorization": o.Authorization,
		// Exclusions The exclusion list for this integration
		"exclusions": o.Exclusions,
		// Location The location of this integration (on-premise / private or cloud)
		"location": o.Location,
		// Name The user friendly name of the integration
		"name": o.Name,
		// Progress Agent processing progress
		"progress": o.Progress,
	}
}

// IntegrationRequestProgress represents the object structure for progress
type IntegrationRequestProgress struct {
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
}

func (o *IntegrationRequestProgress) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Completed The total amount processed thus far
		"completed": o.Completed,
		// Message Any relevant messaging during processing
		"message": o.Message,
		// Total The total amount to be processed
		"total": o.Total,
	}
}

// IntegrationRequest an agent action to request adding an integration
type IntegrationRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Date the date when the request was made
	Date IntegrationRequestDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to add
	Integration IntegrationRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationRequest)(nil)

func toIntegrationRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toIntegrationRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toIntegrationRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toIntegrationRequestObjectNil(isavro, isoptional)
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
	case *IntegrationRequest:
		return v.ToMap()
	case IntegrationRequest:
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
			arr = append(arr, toIntegrationRequestObject(av, isavro, false, ""))
		}
		return arr

	case IntegrationRequestAuthorization:
		vv := o.(IntegrationRequestAuthorization)
		return vv.ToMap()
	case *IntegrationRequestAuthorization:
		return (*o.(*IntegrationRequestAuthorization)).ToMap()
	case []IntegrationRequestAuthorization:
		arr := make([]interface{}, 0)
		for _, i := range o.([]IntegrationRequestAuthorization) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]IntegrationRequestAuthorization:
		arr := make([]interface{}, 0)
		vv := o.(*[]IntegrationRequestAuthorization)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case IntegrationRequestDate:
		vv := o.(IntegrationRequestDate)
		return vv.ToMap()
	case *IntegrationRequestDate:
		return (*o.(*IntegrationRequestDate)).ToMap()
	case []IntegrationRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]IntegrationRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]IntegrationRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]IntegrationRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case IntegrationRequestIntegration:
		vv := o.(IntegrationRequestIntegration)
		return vv.ToMap()
	case *IntegrationRequestIntegration:
		return (*o.(*IntegrationRequestIntegration)).ToMap()
	case []IntegrationRequestIntegration:
		arr := make([]interface{}, 0)
		for _, i := range o.([]IntegrationRequestIntegration) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]IntegrationRequestIntegration:
		arr := make([]interface{}, 0)
		vv := o.(*[]IntegrationRequestIntegration)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case IntegrationRequestProgress:
		vv := o.(IntegrationRequestProgress)
		return vv.ToMap()
	case *IntegrationRequestProgress:
		return (*o.(*IntegrationRequestProgress)).ToMap()
	case []IntegrationRequestProgress:
		arr := make([]interface{}, 0)
		for _, i := range o.([]IntegrationRequestProgress) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]IntegrationRequestProgress:
		arr := make([]interface{}, 0)
		vv := o.(*[]IntegrationRequestProgress)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of IntegrationRequest
func (o *IntegrationRequest) String() string {
	return fmt.Sprintf("agent.IntegrationRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationRequest) GetTopicName() datamodel.TopicNameType {
	return IntegrationRequestTopic
}

// GetModelName returns the name of the model
func (o *IntegrationRequest) GetModelName() datamodel.ModelNameType {
	return IntegrationRequestModelName
}

func (o *IntegrationRequest) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationRequest) GetTimestamp() time.Time {
	var dt interface{} = o.Date
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
	case IntegrationRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for IntegrationRequest")
}

// GetRefID returns the RefID for the object
func (o *IntegrationRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	duration, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         duration,
		MaxSize:           5242880,
	}
}

// GetStateKey returns a key for use in state store
func (o *IntegrationRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *IntegrationRequest) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of IntegrationRequest
func (o *IntegrationRequest) Clone() datamodel.Model {
	c := new(IntegrationRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationRequest) Anon() datamodel.Model {
	c := new(IntegrationRequest)
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
func (o *IntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecIntegrationRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *IntegrationRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecIntegrationRequest == nil {
		c, err := GetIntegrationRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecIntegrationRequest = c
	}
	return cachedCodecIntegrationRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *IntegrationRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *IntegrationRequest) FromAvroBinary(value []byte) error {
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
func (o *IntegrationRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two IntegrationRequest objects are equal
func (o *IntegrationRequest) IsEqual(other *IntegrationRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"customer_id": toIntegrationRequestObject(o.CustomerID, isavro, false, "string"),
		"date":        toIntegrationRequestObject(o.Date, isavro, false, "date"),
		"id":          toIntegrationRequestObject(o.ID, isavro, false, "string"),
		"integration": toIntegrationRequestObject(o.Integration, isavro, false, "integration"),
		"ref_id":      toIntegrationRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":    toIntegrationRequestObject(o.RefType, isavro, false, "string"),
		"uuid":        toIntegrationRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":    toIntegrationRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequest) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["date"].(IntegrationRequestDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = IntegrationRequestDate{}
		} else {
			o.Date = IntegrationRequestDate{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Date)

		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["integration"].(IntegrationRequestIntegration); ok {
		o.Integration = val
	} else {
		val := kv["integration"]
		if val == nil {
			o.Integration = IntegrationRequestIntegration{}
		} else {
			o.Integration = IntegrationRequestIntegration{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Integration)

		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *IntegrationRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Date)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetIntegrationRequestAvroSchemaSpec creates the avro schema specification for IntegrationRequest
func GetIntegrationRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "IntegrationRequest",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "date",
				"type": map[string]interface{}{"type": "record", "name": "date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the request was made"},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration",
				"type": map[string]interface{}{"type": "record", "name": "integration", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "If true, the integration is still active"}, map[string]interface{}{"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "string", "name": "access_token", "doc": "Access token"}, map[string]interface{}{"type": "string", "name": "api_token", "doc": "API Token for instance, if relevant"}, map[string]interface{}{"doc": "the agents encrypted authorization", "type": "string", "name": "authorization"}, map[string]interface{}{"name": "errored", "doc": "If authorization failed by the agent", "type": "boolean"}, map[string]interface{}{"type": "string", "name": "password", "doc": "Password for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "refresh_token", "doc": "Refresh token"}, map[string]interface{}{"type": "string", "name": "url", "doc": "URL of instance if relevant"}, map[string]interface{}{"type": "string", "name": "username", "doc": "Username for instance, if relevant"}, map[string]interface{}{"type": "boolean", "name": "validated", "doc": "If the validation has been run against this instance"}, map[string]interface{}{"doc": "Timestamp when validated", "type": "long", "name": "validated_ts"}, map[string]interface{}{"type": "string", "name": "validation_message", "doc": "The validation message from the agent"}}, "doc": "Authorization information", "type": "record", "name": "integration.authorization"}, "name": "authorization", "doc": "Authorization information"}, map[string]interface{}{"doc": "The exclusion list for this integration", "type": map[string]interface{}{"items": "string", "type": "array", "name": "exclusions"}, "name": "exclusions"}, map[string]interface{}{"type": "string", "name": "location", "doc": "The location of this integration (on-premise / private or cloud)"}, map[string]interface{}{"type": "string", "name": "name", "doc": "The user friendly name of the integration"}, map[string]interface{}{"name": "progress", "doc": "Agent processing progress", "type": map[string]interface{}{"type": "record", "name": "integration.progress", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "completed", "doc": "The total amount processed thus far"}, map[string]interface{}{"type": "string", "name": "message", "doc": "Any relevant messaging during processing"}, map[string]interface{}{"type": "long", "name": "total", "doc": "The total amount to be processed"}}, "doc": "Agent processing progress"}}}, "doc": "the integration details to add"},
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
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetIntegrationRequestAvroSchema creates the avro schema for IntegrationRequest
func GetIntegrationRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetIntegrationRequestAvroSchemaSpec())
}

// TransformIntegrationRequestFunc is a function for transforming IntegrationRequest during processing
type TransformIntegrationRequestFunc func(input *IntegrationRequest) (*IntegrationRequest, error)

// NewIntegrationRequestPipe creates a pipe for processing IntegrationRequest items
func NewIntegrationRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformIntegrationRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewIntegrationRequestInputStream(input, errors)
	var stream chan IntegrationRequest
	if len(transforms) > 0 {
		stream = make(chan IntegrationRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewIntegrationRequestOutputStream(output, stream, errors)
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

// NewIntegrationRequestInputStreamDir creates a channel for reading IntegrationRequest as JSON newlines from a directory of files
func NewIntegrationRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformIntegrationRequestFunc) (chan IntegrationRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/integration_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan IntegrationRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for integration_request")
		ch := make(chan IntegrationRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewIntegrationRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan IntegrationRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewIntegrationRequestInputStreamFile creates an channel for reading IntegrationRequest as JSON newlines from filename
func NewIntegrationRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformIntegrationRequestFunc) (chan IntegrationRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan IntegrationRequest)
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
			ch := make(chan IntegrationRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewIntegrationRequestInputStream(f, errors, transforms...)
}

// NewIntegrationRequestInputStream creates an channel for reading IntegrationRequest as JSON newlines from stream
func NewIntegrationRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformIntegrationRequestFunc) (chan IntegrationRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan IntegrationRequest, 1000)
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
			var item IntegrationRequest
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

// NewIntegrationRequestOutputStreamDir will output json newlines from channel and save in dir
func NewIntegrationRequestOutputStreamDir(dir string, ch chan IntegrationRequest, errors chan<- error, transforms ...TransformIntegrationRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/integration_request\\.json(\\.gz)?$")
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
	return NewIntegrationRequestOutputStream(gz, ch, errors, transforms...)
}

// NewIntegrationRequestOutputStream will output json newlines from channel to the stream
func NewIntegrationRequestOutputStream(stream io.WriteCloser, ch chan IntegrationRequest, errors chan<- error, transforms ...TransformIntegrationRequestFunc) <-chan bool {
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

// IntegrationRequestSendEvent is an event detail for sending data
type IntegrationRequestSendEvent struct {
	IntegrationRequest *IntegrationRequest
	headers            map[string]string
	time               time.Time
	key                string
}

var _ datamodel.ModelSendEvent = (*IntegrationRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *IntegrationRequestSendEvent) Key() string {
	if e.key == "" {
		return e.IntegrationRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *IntegrationRequestSendEvent) Object() datamodel.Model {
	return e.IntegrationRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *IntegrationRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *IntegrationRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// IntegrationRequestSendEventOpts is a function handler for setting opts
type IntegrationRequestSendEventOpts func(o *IntegrationRequestSendEvent)

// WithIntegrationRequestSendEventKey sets the key value to a value different than the object ID
func WithIntegrationRequestSendEventKey(key string) IntegrationRequestSendEventOpts {
	return func(o *IntegrationRequestSendEvent) {
		o.key = key
	}
}

// WithIntegrationRequestSendEventTimestamp sets the timestamp value
func WithIntegrationRequestSendEventTimestamp(tv time.Time) IntegrationRequestSendEventOpts {
	return func(o *IntegrationRequestSendEvent) {
		o.time = tv
	}
}

// WithIntegrationRequestSendEventHeader sets the timestamp value
func WithIntegrationRequestSendEventHeader(key, value string) IntegrationRequestSendEventOpts {
	return func(o *IntegrationRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewIntegrationRequestSendEvent returns a new IntegrationRequestSendEvent instance
func NewIntegrationRequestSendEvent(o *IntegrationRequest, opts ...IntegrationRequestSendEventOpts) *IntegrationRequestSendEvent {
	res := &IntegrationRequestSendEvent{
		IntegrationRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewIntegrationRequestProducer will stream data from the channel
func NewIntegrationRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*IntegrationRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.IntegrationRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewIntegrationRequestConsumer will stream data from the topic into the provided channel
func NewIntegrationRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object IntegrationRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.IntegrationRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into agent.IntegrationRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.IntegrationRequest")
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &IntegrationRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object IntegrationRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &IntegrationRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// IntegrationRequestReceiveEvent is an event detail for receiving data
type IntegrationRequestReceiveEvent struct {
	IntegrationRequest *IntegrationRequest
	message            eventing.Message
	eof                bool
}

var _ datamodel.ModelReceiveEvent = (*IntegrationRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *IntegrationRequestReceiveEvent) Object() datamodel.Model {
	return e.IntegrationRequest
}

// Message returns the underlying message data for the event
func (e *IntegrationRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *IntegrationRequestReceiveEvent) EOF() bool {
	return e.eof
}

// IntegrationRequestProducer implements the datamodel.ModelEventProducer
type IntegrationRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*IntegrationRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *IntegrationRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *IntegrationRequestProducer) Close() error {
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
func (o *IntegrationRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *IntegrationRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IntegrationRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIntegrationRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewIntegrationRequestProducerChannel returns a channel which can be used for producing Model events
func NewIntegrationRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewIntegrationRequestProducerChannelSize(producer, 0, errors)
}

// NewIntegrationRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewIntegrationRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &IntegrationRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewIntegrationRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// IntegrationRequestConsumer implements the datamodel.ModelEventConsumer
type IntegrationRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*IntegrationRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *IntegrationRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *IntegrationRequestConsumer) Close() error {
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
func (o *IntegrationRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IntegrationRequestConsumer{
		ch:       ch,
		callback: NewIntegrationRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewIntegrationRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewIntegrationRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &IntegrationRequestConsumer{
		ch:       ch,
		callback: NewIntegrationRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
