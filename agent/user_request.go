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
	// UserRequestTopic is the default topic name
	UserRequestTopic datamodel.TopicNameType = "agent_UserRequest_topic"

	// UserRequestStream is the default stream name
	UserRequestStream datamodel.TopicNameType = "agent_UserRequest_stream"

	// UserRequestTable is the default table name
	UserRequestTable datamodel.TopicNameType = "agent_UserRequest"

	// UserRequestModelName is the model name
	UserRequestModelName datamodel.ModelNameType = "agent.UserRequest"
)

const (
	// UserRequestCustomerIDColumn is the customer_id column name
	UserRequestCustomerIDColumn = "customer_id"
	// UserRequestIDColumn is the id column name
	UserRequestIDColumn = "id"
	// UserRequestIntegrationColumn is the integration column name
	UserRequestIntegrationColumn = "integration"
	// UserRequestIntegrationColumnActiveColumn is the active column property of the Integration name
	UserRequestIntegrationColumnActiveColumn = "integration->active"
	// UserRequestIntegrationColumnAuthorizationColumn is the authorization column property of the Integration name
	UserRequestIntegrationColumnAuthorizationColumn = "integration->authorization"
	// UserRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	UserRequestIntegrationColumnErroredColumn = "integration->errored"
	// UserRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	UserRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// UserRequestIntegrationColumnNameColumn is the name column property of the Integration name
	UserRequestIntegrationColumnNameColumn = "integration->name"
	// UserRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	UserRequestIntegrationColumnProgressColumn = "integration->progress"
	// UserRequestIntegrationColumnValidatedColumn is the validated column property of the Integration name
	UserRequestIntegrationColumnValidatedColumn = "integration->validated"
	// UserRequestIntegrationColumnValidatedDateColumn is the validated_date column property of the Integration name
	UserRequestIntegrationColumnValidatedDateColumn = "integration->validated_date"
	// UserRequestIntegrationColumnValidationMessageColumn is the validation_message column property of the Integration name
	UserRequestIntegrationColumnValidationMessageColumn = "integration->validation_message"
	// UserRequestLocationColumn is the location column name
	UserRequestLocationColumn = "location"
	// UserRequestRefIDColumn is the ref_id column name
	UserRequestRefIDColumn = "ref_id"
	// UserRequestRefTypeColumn is the ref_type column name
	UserRequestRefTypeColumn = "ref_type"
	// UserRequestRequestDateColumn is the request_date column name
	UserRequestRequestDateColumn = "request_date"
	// UserRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	UserRequestRequestDateColumnEpochColumn = "request_date->epoch"
	// UserRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	UserRequestRequestDateColumnOffsetColumn = "request_date->offset"
	// UserRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	UserRequestRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// UserRequestUUIDColumn is the uuid column name
	UserRequestUUIDColumn = "uuid"
)

// UserRequestAuthorization represents the object structure for authorization
type UserRequestAuthorization struct {
	// AccessToken Access token
	AccessToken *string `json:"access_token" bson:"access_token" yaml:"access_token" faker:"-"`
	// APIToken API Token for instance, if relevant
	APIToken *string `json:"api_token" bson:"api_token" yaml:"api_token" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Password Password for instance, if relevant
	Password *string `json:"password" bson:"password" yaml:"password" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken *string `json:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url" bson:"url" yaml:"url" faker:"-"`
	// Username Username for instance, if relevant
	Username *string `json:"username" bson:"username" yaml:"username" faker:"-"`
}

func (o *UserRequestAuthorization) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": o.AccessToken,
		// APIToken API Token for instance, if relevant
		"api_token": o.APIToken,
		// Authorization the agents encrypted authorization
		"authorization": o.Authorization,
		// Password Password for instance, if relevant
		"password": o.Password,
		// RefreshToken Refresh token
		"refresh_token": o.RefreshToken,
		// URL URL of instance if relevant
		"url": o.URL,
		// Username Username for instance, if relevant
		"username": o.Username,
	}
}

// UserRequestIntegration represents the object structure for integration
type UserRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization UserRequestAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress UserRequestProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate UserRequestValidatedDate `json:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func (o *UserRequestIntegration) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": o.Active,
		// Authorization Authorization information
		"authorization": o.Authorization,
		// Errored If authorization failed by the agent
		"errored": o.Errored,
		// Exclusions The exclusion list for this integration
		"exclusions": o.Exclusions,
		// Name The user friendly name of the integration
		"name": o.Name,
		// Progress Agent processing progress
		"progress": o.Progress,
		// Validated If the validation has been run against this instance
		"validated": o.Validated,
		// ValidatedDate Date when validated
		"validated_date": o.ValidatedDate,
		// ValidationMessage The validation message from the agent
		"validation_message": o.ValidationMessage,
	}
}

// UserRequestLocation is the enumeration type for location
type UserRequestLocation int32

// String returns the string value for Location
func (v UserRequestLocation) String() string {
	switch int32(v) {
	case 0:
		return "private"
	case 1:
		return "cloud"
	}
	return "unset"
}

const (
	// LocationPrivate is the enumeration value for private
	UserRequestLocationPrivate UserRequestLocation = 0
	// LocationCloud is the enumeration value for cloud
	UserRequestLocationCloud UserRequestLocation = 1
)

// UserRequestProgress represents the object structure for progress
type UserRequestProgress struct {
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
}

func (o *UserRequestProgress) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Completed The total amount processed thus far
		"completed": o.Completed,
		// Message Any relevant messaging during processing
		"message": o.Message,
		// Total The total amount to be processed
		"total": o.Total,
	}
}

// UserRequestRequestDate represents the object structure for request_date
type UserRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *UserRequestRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// UserRequestValidatedDate represents the object structure for validated_date
type UserRequestValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *UserRequestValidatedDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// UserRequest an agent action to request adding new users
type UserRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration UserRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location UserRequestLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate UserRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UserRequest)(nil)

func toUserRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toUserRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toUserRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toUserRequestObjectNil(isavro, isoptional)
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
	case *UserRequest:
		return v.ToMap()
	case UserRequest:
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
			arr = append(arr, toUserRequestObject(av, isavro, false, ""))
		}
		return arr

	case UserRequestAuthorization:
		vv := o.(UserRequestAuthorization)
		return vv.ToMap()
	case *UserRequestAuthorization:
		return map[string]interface{}{
			"agent.authorization": (*o.(*UserRequestAuthorization)).ToMap(),
		}
	case []UserRequestAuthorization:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserRequestAuthorization) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserRequestAuthorization:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserRequestAuthorization)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserRequestIntegration:
		vv := o.(UserRequestIntegration)
		return vv.ToMap()
	case *UserRequestIntegration:
		return map[string]interface{}{
			"agent.integration": (*o.(*UserRequestIntegration)).ToMap(),
		}
	case []UserRequestIntegration:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserRequestIntegration) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserRequestIntegration:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserRequestIntegration)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserRequestLocation:
		if !isavro {
			return (o.(UserRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(UserRequestLocation)).String(),
		}
	case *UserRequestLocation:
		if !isavro {
			return (o.(*UserRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(*UserRequestLocation)).String(),
		}
	case UserRequestProgress:
		vv := o.(UserRequestProgress)
		return vv.ToMap()
	case *UserRequestProgress:
		return map[string]interface{}{
			"agent.progress": (*o.(*UserRequestProgress)).ToMap(),
		}
	case []UserRequestProgress:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserRequestProgress) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserRequestProgress:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserRequestProgress)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserRequestRequestDate:
		vv := o.(UserRequestRequestDate)
		return vv.ToMap()
	case *UserRequestRequestDate:
		return map[string]interface{}{
			"agent.request_date": (*o.(*UserRequestRequestDate)).ToMap(),
		}
	case []UserRequestRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserRequestRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserRequestRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserRequestRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserRequestValidatedDate:
		vv := o.(UserRequestValidatedDate)
		return vv.ToMap()
	case *UserRequestValidatedDate:
		return map[string]interface{}{
			"agent.validated_date": (*o.(*UserRequestValidatedDate)).ToMap(),
		}
	case []UserRequestValidatedDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserRequestValidatedDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserRequestValidatedDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserRequestValidatedDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of UserRequest
func (o *UserRequest) String() string {
	return fmt.Sprintf("agent.UserRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UserRequest) GetTopicName() datamodel.TopicNameType {
	return UserRequestTopic
}

// GetModelName returns the name of the model
func (o *UserRequest) GetModelName() datamodel.ModelNameType {
	return UserRequestModelName
}

func (o *UserRequest) setDefaults() {
	if o.Integration.Authorization.AccessToken == nil {
		o.Integration.Authorization.AccessToken = &emptyString
	}
	if o.Integration.Authorization.APIToken == nil {
		o.Integration.Authorization.APIToken = &emptyString
	}
	if o.Integration.Authorization.Authorization == nil {
		o.Integration.Authorization.Authorization = &emptyString
	}
	if o.Integration.Authorization.Password == nil {
		o.Integration.Authorization.Password = &emptyString
	}
	if o.Integration.Authorization.RefreshToken == nil {
		o.Integration.Authorization.RefreshToken = &emptyString
	}
	if o.Integration.Authorization.URL == nil {
		o.Integration.Authorization.URL = &emptyString
	}
	if o.Integration.Authorization.Username == nil {
		o.Integration.Authorization.Username = &emptyString
	}
	if o.Integration.Errored == nil {
		o.Integration.Errored = &emptyBool
	}
	if o.Integration.Exclusions == nil {
		o.Integration.Exclusions = []string{}
	}
	if o.Integration.Validated == nil {
		o.Integration.Validated = &emptyBool
	}
	if o.Integration.ValidationMessage == nil {
		o.Integration.ValidationMessage = &emptyString
	}

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UserRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UserRequest) GetTimestamp() time.Time {
	var dt interface{} = o.RequestDate
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
	case UserRequestRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for UserRequest")
}

// GetRefID returns the RefID for the object
func (o *UserRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UserRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UserRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UserRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *UserRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UserRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "request_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *UserRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *UserRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of UserRequest
func (o *UserRequest) Clone() datamodel.Model {
	c := new(UserRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UserRequest) Anon() datamodel.Model {
	c := new(UserRequest)
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
func (o *UserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecUserRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *UserRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecUserRequest == nil {
		c, err := GetUserRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUserRequest = c
	}
	return cachedCodecUserRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *UserRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *UserRequest) FromAvroBinary(value []byte) error {
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
func (o *UserRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserRequest objects are equal
func (o *UserRequest) IsEqual(other *UserRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"customer_id":  toUserRequestObject(o.CustomerID, isavro, false, "string"),
		"id":           toUserRequestObject(o.ID, isavro, false, "string"),
		"integration":  toUserRequestObject(o.Integration, isavro, false, "integration"),
		"location":     toUserRequestObject(o.Location, isavro, false, "location"),
		"ref_id":       toUserRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toUserRequestObject(o.RefType, isavro, false, "string"),
		"request_date": toUserRequestObject(o.RequestDate, isavro, false, "request_date"),
		"uuid":         toUserRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toUserRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequest) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["integration"].(UserRequestIntegration); ok {
		o.Integration = val
	} else {
		val := kv["integration"]
		if val == nil {
			o.Integration = UserRequestIntegration{}
		} else {
			o.Integration = UserRequestIntegration{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Integration)

		}
	}
	if val, ok := kv["location"].(UserRequestLocation); ok {
		o.Location = val
	} else {
		if em, ok := kv["location"].(map[string]interface{}); ok {
			ev := em["agent.location"].(string)
			switch ev {
			case "private":
				o.Location = 0
			case "cloud":
				o.Location = 1
			}
		}
		if em, ok := kv["location"].(string); ok {
			switch em {
			case "private":
				o.Location = 0
			case "cloud":
				o.Location = 1
			}
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
	if val, ok := kv["request_date"].(UserRequestRequestDate); ok {
		o.RequestDate = val
	} else {
		val := kv["request_date"]
		if val == nil {
			o.RequestDate = UserRequestRequestDate{}
		} else {
			o.RequestDate = UserRequestRequestDate{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.RequestDate)

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
func (o *UserRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.Location)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUserRequestAvroSchemaSpec creates the avro schema specification for UserRequest
func GetUserRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "UserRequest",
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
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration",
				"type": map[string]interface{}{"type": "record", "name": "integration", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "If true, the integration is still active"}, map[string]interface{}{"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "string", "name": "access_token", "doc": "Access token"}, map[string]interface{}{"type": "string", "name": "api_token", "doc": "API Token for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "authorization", "doc": "the agents encrypted authorization"}, map[string]interface{}{"type": "string", "name": "password", "doc": "Password for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "refresh_token", "doc": "Refresh token"}, map[string]interface{}{"type": "string", "name": "url", "doc": "URL of instance if relevant"}, map[string]interface{}{"name": "username", "doc": "Username for instance, if relevant", "type": "string"}}, "doc": "Authorization information", "type": "record", "name": "integration.authorization"}, "name": "authorization", "doc": "Authorization information"}, map[string]interface{}{"type": "boolean", "name": "errored", "doc": "If authorization failed by the agent"}, map[string]interface{}{"type": map[string]interface{}{"type": "array", "name": "exclusions", "items": "string"}, "name": "exclusions", "doc": "The exclusion list for this integration"}, map[string]interface{}{"type": "string", "name": "name", "doc": "The user friendly name of the integration"}, map[string]interface{}{"doc": "Agent processing progress", "type": map[string]interface{}{"type": "record", "name": "integration.progress", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "completed", "doc": "The total amount processed thus far"}, map[string]interface{}{"type": "string", "name": "message", "doc": "Any relevant messaging during processing"}, map[string]interface{}{"type": "long", "name": "total", "doc": "The total amount to be processed"}}, "doc": "Agent processing progress"}, "name": "progress"}, map[string]interface{}{"type": "boolean", "name": "validated", "doc": "If the validation has been run against this instance"}, map[string]interface{}{"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "Date when validated", "type": "record", "name": "integration.validated_date"}, "name": "validated_date", "doc": "Date when validated"}, map[string]interface{}{"doc": "The validation message from the agent", "type": "string", "name": "validation_message"}}, "doc": "the integration details to use"},
			},
			map[string]interface{}{
				"name": "location",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "location",
						"symbols": []interface{}{"private", "cloud"},
					},
				},
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
				"name": "request_date",
				"type": map[string]interface{}{"type": "record", "name": "request_date", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date when the request was made"},
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *UserRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetUserRequestAvroSchema creates the avro schema for UserRequest
func GetUserRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserRequestAvroSchemaSpec())
}

// TransformUserRequestFunc is a function for transforming UserRequest during processing
type TransformUserRequestFunc func(input *UserRequest) (*UserRequest, error)

// NewUserRequestPipe creates a pipe for processing UserRequest items
func NewUserRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUserRequestInputStream(input, errors)
	var stream chan UserRequest
	if len(transforms) > 0 {
		stream = make(chan UserRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewUserRequestOutputStream(output, stream, errors)
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

// NewUserRequestInputStreamDir creates a channel for reading UserRequest as JSON newlines from a directory of files
func NewUserRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserRequestFunc) (chan UserRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/user_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan UserRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user_request")
		ch := make(chan UserRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewUserRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan UserRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUserRequestInputStreamFile creates an channel for reading UserRequest as JSON newlines from filename
func NewUserRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserRequestFunc) (chan UserRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan UserRequest)
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
			ch := make(chan UserRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewUserRequestInputStream(f, errors, transforms...)
}

// NewUserRequestInputStream creates an channel for reading UserRequest as JSON newlines from stream
func NewUserRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserRequestFunc) (chan UserRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan UserRequest, 1000)
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
			var item UserRequest
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

// NewUserRequestOutputStreamDir will output json newlines from channel and save in dir
func NewUserRequestOutputStreamDir(dir string, ch chan UserRequest, errors chan<- error, transforms ...TransformUserRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/user_request\\.json(\\.gz)?$")
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
	return NewUserRequestOutputStream(gz, ch, errors, transforms...)
}

// NewUserRequestOutputStream will output json newlines from channel to the stream
func NewUserRequestOutputStream(stream io.WriteCloser, ch chan UserRequest, errors chan<- error, transforms ...TransformUserRequestFunc) <-chan bool {
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

// UserRequestSendEvent is an event detail for sending data
type UserRequestSendEvent struct {
	UserRequest *UserRequest
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*UserRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *UserRequestSendEvent) Key() string {
	if e.key == "" {
		return e.UserRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserRequestSendEvent) Object() datamodel.Model {
	return e.UserRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// UserRequestSendEventOpts is a function handler for setting opts
type UserRequestSendEventOpts func(o *UserRequestSendEvent)

// WithUserRequestSendEventKey sets the key value to a value different than the object ID
func WithUserRequestSendEventKey(key string) UserRequestSendEventOpts {
	return func(o *UserRequestSendEvent) {
		o.key = key
	}
}

// WithUserRequestSendEventTimestamp sets the timestamp value
func WithUserRequestSendEventTimestamp(tv time.Time) UserRequestSendEventOpts {
	return func(o *UserRequestSendEvent) {
		o.time = tv
	}
}

// WithUserRequestSendEventHeader sets the timestamp value
func WithUserRequestSendEventHeader(key, value string) UserRequestSendEventOpts {
	return func(o *UserRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserRequestSendEvent returns a new UserRequestSendEvent instance
func NewUserRequestSendEvent(o *UserRequest, opts ...UserRequestSendEventOpts) *UserRequestSendEvent {
	res := &UserRequestSendEvent{
		UserRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserRequestProducer will stream data from the channel
func NewUserRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*UserRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.UserRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUserRequestConsumer will stream data from the topic into the provided channel
func NewUserRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object UserRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.UserRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.UserRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.UserRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &UserRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object UserRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UserRequestReceiveEvent is an event detail for receiving data
type UserRequestReceiveEvent struct {
	UserRequest *UserRequest
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*UserRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserRequestReceiveEvent) Object() datamodel.Model {
	return e.UserRequest
}

// Message returns the underlying message data for the event
func (e *UserRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserRequestReceiveEvent) EOF() bool {
	return e.eof
}

// UserRequestProducer implements the datamodel.ModelEventProducer
type UserRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UserRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserRequestProducer) Close() error {
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
func (o *UserRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *UserRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUserRequestProducerChannel returns a channel which can be used for producing Model events
func NewUserRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUserRequestProducerChannelSize(producer, 0, errors)
}

// NewUserRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewUserRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// UserRequestConsumer implements the datamodel.ModelEventConsumer
type UserRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UserRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserRequestConsumer) Close() error {
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
func (o *UserRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserRequestConsumer{
		ch:       ch,
		callback: NewUserRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUserRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserRequestConsumer{
		ch:       ch,
		callback: NewUserRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
