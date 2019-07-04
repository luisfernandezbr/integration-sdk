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
	pstrings "github.com/pinpt/go-common/strings"
)

const (
	// ExportRequestTopic is the default topic name
	ExportRequestTopic datamodel.TopicNameType = "agent_ExportRequest_topic"

	// ExportRequestStream is the default stream name
	ExportRequestStream datamodel.TopicNameType = "agent_ExportRequest_stream"

	// ExportRequestTable is the default table name
	ExportRequestTable datamodel.TopicNameType = "agent_ExportRequest"

	// ExportRequestModelName is the model name
	ExportRequestModelName datamodel.ModelNameType = "agent.ExportRequest"
)

const (
	// ExportRequestCustomerIDColumn is the customer_id column name
	ExportRequestCustomerIDColumn = "customer_id"
	// ExportRequestDateColumn is the date column name
	ExportRequestDateColumn = "date"
	// ExportRequestIDColumn is the id column name
	ExportRequestIDColumn = "id"
	// ExportRequestIntegrationsColumn is the integrations column name
	ExportRequestIntegrationsColumn = "integrations"
	// ExportRequestRefIDColumn is the ref_id column name
	ExportRequestRefIDColumn = "ref_id"
	// ExportRequestRefTypeColumn is the ref_type column name
	ExportRequestRefTypeColumn = "ref_type"
	// ExportRequestUploadURLColumn is the upload_url column name
	ExportRequestUploadURLColumn = "upload_url"
	// ExportRequestUUIDColumn is the uuid column name
	ExportRequestUUIDColumn = "uuid"
)

// ExportRequestAuthorization represents the object structure for authorization
type ExportRequestAuthorization struct {
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

func (o *ExportRequestAuthorization) ToMap() map[string]interface{} {
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

// ExportRequestDate represents the object structure for date
type ExportRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *ExportRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// ExportRequestIntegrations represents the object structure for integrations
type ExportRequestIntegrations struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization ExportRequestAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location string `json:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress ExportRequestProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
}

func (o *ExportRequestIntegrations) ToMap() map[string]interface{} {
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

// ExportRequestProgress represents the object structure for progress
type ExportRequestProgress struct {
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
}

func (o *ExportRequestProgress) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Completed The total amount processed thus far
		"completed": o.Completed,
		// Message Any relevant messaging during processing
		"message": o.Message,
		// Total The total amount to be processed
		"total": o.Total,
	}
}

// ExportRequest an agent action to request an export
type ExportRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Date the date when the request was made
	Date ExportRequestDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integrations The integrations that should be exported and their current configuration
	Integrations []ExportRequestIntegrations `json:"integrations" bson:"integrations" yaml:"integrations" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UploadURL The one time upload URL to use for uploading a job to the Pinpoint cloud
	UploadURL *string `json:"upload_url" bson:"upload_url" yaml:"upload_url" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportRequest)(nil)

func toExportRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toExportRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toExportRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toExportRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toExportRequestObjectNil(isavro, isoptional)
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
	case *ExportRequest:
		return v.ToMap()
	case ExportRequest:
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
			arr = append(arr, toExportRequestObject(av, isavro, false, ""))
		}
		return arr

	case ExportRequestAuthorization:
		vv := o.(ExportRequestAuthorization)
		return vv.ToMap()
	case *ExportRequestAuthorization:
		return (*o.(*ExportRequestAuthorization)).ToMap()
	case []ExportRequestAuthorization:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ExportRequestAuthorization) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ExportRequestAuthorization:
		arr := make([]interface{}, 0)
		vv := o.(*[]ExportRequestAuthorization)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ExportRequestDate:
		vv := o.(ExportRequestDate)
		return vv.ToMap()
	case *ExportRequestDate:
		return (*o.(*ExportRequestDate)).ToMap()
	case []ExportRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ExportRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ExportRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]ExportRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ExportRequestIntegrations:
		vv := o.(ExportRequestIntegrations)
		return vv.ToMap()
	case *ExportRequestIntegrations:
		return (*o.(*ExportRequestIntegrations)).ToMap()
	case []ExportRequestIntegrations:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ExportRequestIntegrations) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ExportRequestIntegrations:
		arr := make([]interface{}, 0)
		vv := o.(*[]ExportRequestIntegrations)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ExportRequestProgress:
		vv := o.(ExportRequestProgress)
		return vv.ToMap()
	case *ExportRequestProgress:
		return (*o.(*ExportRequestProgress)).ToMap()
	case []ExportRequestProgress:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ExportRequestProgress) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ExportRequestProgress:
		arr := make([]interface{}, 0)
		vv := o.(*[]ExportRequestProgress)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of ExportRequest
func (o *ExportRequest) String() string {
	return fmt.Sprintf("agent.ExportRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportRequest) GetTopicName() datamodel.TopicNameType {
	return ExportRequestTopic
}

// GetModelName returns the name of the model
func (o *ExportRequest) GetModelName() datamodel.ModelNameType {
	return ExportRequestModelName
}

func (o *ExportRequest) setDefaults() {
	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportRequest) GetTimestamp() time.Time {
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
	case ExportRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for ExportRequest")
}

// GetRefID returns the RefID for the object
func (o *ExportRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *ExportRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *ExportRequest) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of ExportRequest
func (o *ExportRequest) Clone() datamodel.Model {
	c := new(ExportRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportRequest) Anon() datamodel.Model {
	c := new(ExportRequest)
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
func (o *ExportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecExportRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ExportRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecExportRequest == nil {
		c, err := GetExportRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecExportRequest = c
	}
	return cachedCodecExportRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *ExportRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *ExportRequest) FromAvroBinary(value []byte) error {
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
func (o *ExportRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ExportRequest objects are equal
func (o *ExportRequest) IsEqual(other *ExportRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Integrations == nil {
			o.Integrations = make([]ExportRequestIntegrations, 0)
		}
	}
	return map[string]interface{}{
		"customer_id":  toExportRequestObject(o.CustomerID, isavro, false, "string"),
		"date":         toExportRequestObject(o.Date, isavro, false, "date"),
		"id":           toExportRequestObject(o.ID, isavro, false, "string"),
		"integrations": toExportRequestObject(o.Integrations, isavro, false, "integrations"),
		"ref_id":       toExportRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toExportRequestObject(o.RefType, isavro, false, "string"),
		"upload_url":   toExportRequestObject(o.UploadURL, isavro, true, "string"),
		"uuid":         toExportRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toExportRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequest) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["date"].(ExportRequestDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = ExportRequestDate{}
		} else {
			o.Date = ExportRequestDate{}
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
	if val := kv["integrations"]; val != nil {
		na := make([]ExportRequestIntegrations, 0)
		if a, ok := val.([]ExportRequestIntegrations); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(ExportRequestIntegrations); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av ExportRequestIntegrations
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for integrations field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for integrations field")
			}
		}
		o.Integrations = na
	} else {
		o.Integrations = []ExportRequestIntegrations{}
	}
	if o.Integrations == nil {
		o.Integrations = make([]ExportRequestIntegrations, 0)
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
	if val, ok := kv["upload_url"].(*string); ok {
		o.UploadURL = val
	} else if val, ok := kv["upload_url"].(string); ok {
		o.UploadURL = &val
	} else {
		val := kv["upload_url"]
		if val == nil {
			o.UploadURL = pstrings.Pointer("")
		} else {
			// if coming in as avro union, convert it back
			if kv, ok := val.(map[string]interface{}); ok {
				val = kv["string"]
			}
			o.UploadURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
func (o *ExportRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.GetID())
	args = append(args, o.GetRefID())
	args = append(args, o.RefType)
	args = append(args, o.CustomerID)
	args = append(args, o.CustomerID)
	args = append(args, o.Date)
	args = append(args, o.ID)
	args = append(args, o.Integrations)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UploadURL)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetExportRequestAvroSchemaSpec creates the avro schema specification for ExportRequest
func GetExportRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "ExportRequest",
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
				"name": "integrations",
				"type": map[string]interface{}{"items": map[string]interface{}{"type": "record", "name": "integrations", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "If true, the integration is still active"}, map[string]interface{}{"doc": "Authorization information", "type": map[string]interface{}{"doc": "Authorization information", "type": "record", "name": "integrations.authorization", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "access_token", "doc": "Access token"}, map[string]interface{}{"name": "api_token", "doc": "API Token for instance, if relevant", "type": "string"}, map[string]interface{}{"type": "string", "name": "authorization", "doc": "the agents encrypted authorization"}, map[string]interface{}{"type": "boolean", "name": "errored", "doc": "If authorization failed by the agent"}, map[string]interface{}{"type": "string", "name": "password", "doc": "Password for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "refresh_token", "doc": "Refresh token"}, map[string]interface{}{"doc": "URL of instance if relevant", "type": "string", "name": "url"}, map[string]interface{}{"name": "username", "doc": "Username for instance, if relevant", "type": "string"}, map[string]interface{}{"type": "boolean", "name": "validated", "doc": "If the validation has been run against this instance"}, map[string]interface{}{"name": "validated_ts", "doc": "Timestamp when validated", "type": "long"}, map[string]interface{}{"type": "string", "name": "validation_message", "doc": "The validation message from the agent"}}}, "name": "authorization"}, map[string]interface{}{"type": map[string]interface{}{"items": "string", "type": "array", "name": "exclusions"}, "name": "exclusions", "doc": "The exclusion list for this integration"}, map[string]interface{}{"type": "string", "name": "location", "doc": "The location of this integration (on-premise / private or cloud)"}, map[string]interface{}{"type": "string", "name": "name", "doc": "The user friendly name of the integration"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "integrations.progress", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "completed", "doc": "The total amount processed thus far"}, map[string]interface{}{"type": "string", "name": "message", "doc": "Any relevant messaging during processing"}, map[string]interface{}{"type": "long", "name": "total", "doc": "The total amount to be processed"}}, "doc": "Agent processing progress"}, "name": "progress", "doc": "Agent processing progress"}}, "doc": "The integrations that should be exported and their current configuration"}, "type": "array", "name": "integrations"},
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
				"name":    "upload_url",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetExportRequestAvroSchema creates the avro schema for ExportRequest
func GetExportRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetExportRequestAvroSchemaSpec())
}

// TransformExportRequestFunc is a function for transforming ExportRequest during processing
type TransformExportRequestFunc func(input *ExportRequest) (*ExportRequest, error)

// NewExportRequestPipe creates a pipe for processing ExportRequest items
func NewExportRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformExportRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewExportRequestInputStream(input, errors)
	var stream chan ExportRequest
	if len(transforms) > 0 {
		stream = make(chan ExportRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewExportRequestOutputStream(output, stream, errors)
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

// NewExportRequestInputStreamDir creates a channel for reading ExportRequest as JSON newlines from a directory of files
func NewExportRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformExportRequestFunc) (chan ExportRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/export_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ExportRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for export_request")
		ch := make(chan ExportRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewExportRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ExportRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewExportRequestInputStreamFile creates an channel for reading ExportRequest as JSON newlines from filename
func NewExportRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformExportRequestFunc) (chan ExportRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ExportRequest)
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
			ch := make(chan ExportRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewExportRequestInputStream(f, errors, transforms...)
}

// NewExportRequestInputStream creates an channel for reading ExportRequest as JSON newlines from stream
func NewExportRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformExportRequestFunc) (chan ExportRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ExportRequest, 1000)
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
			var item ExportRequest
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

// NewExportRequestOutputStreamDir will output json newlines from channel and save in dir
func NewExportRequestOutputStreamDir(dir string, ch chan ExportRequest, errors chan<- error, transforms ...TransformExportRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/export_request\\.json(\\.gz)?$")
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
	return NewExportRequestOutputStream(gz, ch, errors, transforms...)
}

// NewExportRequestOutputStream will output json newlines from channel to the stream
func NewExportRequestOutputStream(stream io.WriteCloser, ch chan ExportRequest, errors chan<- error, transforms ...TransformExportRequestFunc) <-chan bool {
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

// ExportRequestSendEvent is an event detail for sending data
type ExportRequestSendEvent struct {
	ExportRequest *ExportRequest
	headers       map[string]string
	time          time.Time
	key           string
}

var _ datamodel.ModelSendEvent = (*ExportRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *ExportRequestSendEvent) Key() string {
	if e.key == "" {
		return e.ExportRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ExportRequestSendEvent) Object() datamodel.Model {
	return e.ExportRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ExportRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ExportRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// ExportRequestSendEventOpts is a function handler for setting opts
type ExportRequestSendEventOpts func(o *ExportRequestSendEvent)

// WithExportRequestSendEventKey sets the key value to a value different than the object ID
func WithExportRequestSendEventKey(key string) ExportRequestSendEventOpts {
	return func(o *ExportRequestSendEvent) {
		o.key = key
	}
}

// WithExportRequestSendEventTimestamp sets the timestamp value
func WithExportRequestSendEventTimestamp(tv time.Time) ExportRequestSendEventOpts {
	return func(o *ExportRequestSendEvent) {
		o.time = tv
	}
}

// WithExportRequestSendEventHeader sets the timestamp value
func WithExportRequestSendEventHeader(key, value string) ExportRequestSendEventOpts {
	return func(o *ExportRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewExportRequestSendEvent returns a new ExportRequestSendEvent instance
func NewExportRequestSendEvent(o *ExportRequest, opts ...ExportRequestSendEventOpts) *ExportRequestSendEvent {
	res := &ExportRequestSendEvent{
		ExportRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewExportRequestProducer will stream data from the channel
func NewExportRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*ExportRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.ExportRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewExportRequestConsumer will stream data from the topic into the provided channel
func NewExportRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object ExportRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.ExportRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avri data into agent.ExportRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.ExportRequest")
			}
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ExportRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object ExportRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ExportRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ExportRequestReceiveEvent is an event detail for receiving data
type ExportRequestReceiveEvent struct {
	ExportRequest *ExportRequest
	message       eventing.Message
	eof           bool
}

var _ datamodel.ModelReceiveEvent = (*ExportRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ExportRequestReceiveEvent) Object() datamodel.Model {
	return e.ExportRequest
}

// Message returns the underlying message data for the event
func (e *ExportRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ExportRequestReceiveEvent) EOF() bool {
	return e.eof
}

// ExportRequestProducer implements the datamodel.ModelEventProducer
type ExportRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ExportRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ExportRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ExportRequestProducer) Close() error {
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
func (o *ExportRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ExportRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewExportRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewExportRequestProducerChannel returns a channel which can be used for producing Model events
func NewExportRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ExportRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewExportRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// ExportRequestConsumer implements the datamodel.ModelEventConsumer
type ExportRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ExportRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ExportRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ExportRequestConsumer) Close() error {
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
func (o *ExportRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ExportRequestConsumer{
		ch:       ch,
		callback: NewExportRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewExportRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewExportRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ExportRequestConsumer{
		ch:       ch,
		callback: NewExportRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
