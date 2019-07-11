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
	// ProjectRequestTopic is the default topic name
	ProjectRequestTopic datamodel.TopicNameType = "agent_ProjectRequest_topic"

	// ProjectRequestStream is the default stream name
	ProjectRequestStream datamodel.TopicNameType = "agent_ProjectRequest_stream"

	// ProjectRequestTable is the default table name
	ProjectRequestTable datamodel.TopicNameType = "agent_ProjectRequest"

	// ProjectRequestModelName is the model name
	ProjectRequestModelName datamodel.ModelNameType = "agent.ProjectRequest"
)

const (
	// ProjectRequestCustomerIDColumn is the customer_id column name
	ProjectRequestCustomerIDColumn = "customer_id"
	// ProjectRequestDateColumn is the date column name
	ProjectRequestDateColumn = "date"
	// ProjectRequestDateColumnEpochColumn is the epoch column property of the Date name
	ProjectRequestDateColumnEpochColumn = "date->epoch"
	// ProjectRequestDateColumnOffsetColumn is the offset column property of the Date name
	ProjectRequestDateColumnOffsetColumn = "date->offset"
	// ProjectRequestDateColumnRfc3339Column is the rfc3339 column property of the Date name
	ProjectRequestDateColumnRfc3339Column = "date->rfc3339"
	// ProjectRequestIDColumn is the id column name
	ProjectRequestIDColumn = "id"
	// ProjectRequestIntegrationColumn is the integration column name
	ProjectRequestIntegrationColumn = "integration"
	// ProjectRequestIntegrationColumnActiveColumn is the active column property of the Integration name
	ProjectRequestIntegrationColumnActiveColumn = "integration->active"
	// ProjectRequestIntegrationColumnAuthorizationColumn is the authorization column property of the Integration name
	ProjectRequestIntegrationColumnAuthorizationColumn = "integration->authorization"
	// ProjectRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	ProjectRequestIntegrationColumnErroredColumn = "integration->errored"
	// ProjectRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	ProjectRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// ProjectRequestIntegrationColumnNameColumn is the name column property of the Integration name
	ProjectRequestIntegrationColumnNameColumn = "integration->name"
	// ProjectRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	ProjectRequestIntegrationColumnProgressColumn = "integration->progress"
	// ProjectRequestIntegrationColumnValidatedColumn is the validated column property of the Integration name
	ProjectRequestIntegrationColumnValidatedColumn = "integration->validated"
	// ProjectRequestIntegrationColumnValidatedAtColumn is the validated_ts column property of the Integration name
	ProjectRequestIntegrationColumnValidatedAtColumn = "integration->validated_ts"
	// ProjectRequestIntegrationColumnValidationMessageColumn is the validation_message column property of the Integration name
	ProjectRequestIntegrationColumnValidationMessageColumn = "integration->validation_message"
	// ProjectRequestLocationColumn is the location column name
	ProjectRequestLocationColumn = "location"
	// ProjectRequestRefIDColumn is the ref_id column name
	ProjectRequestRefIDColumn = "ref_id"
	// ProjectRequestRefTypeColumn is the ref_type column name
	ProjectRequestRefTypeColumn = "ref_type"
	// ProjectRequestUUIDColumn is the uuid column name
	ProjectRequestUUIDColumn = "uuid"
)

// ProjectRequestAuthorization represents the object structure for authorization
type ProjectRequestAuthorization struct {
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

func (o *ProjectRequestAuthorization) ToMap() map[string]interface{} {
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

// ProjectRequestDate represents the object structure for date
type ProjectRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *ProjectRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// ProjectRequestIntegration represents the object structure for integration
type ProjectRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization ProjectRequestAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress ProjectRequestProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedAt Timestamp when validated
	ValidatedAt *int64 `json:"validated_ts" bson:"validated_ts" yaml:"validated_ts" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func (o *ProjectRequestIntegration) ToMap() map[string]interface{} {
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
		// ValidatedAt Timestamp when validated
		"validated_ts": o.ValidatedAt,
		// ValidationMessage The validation message from the agent
		"validation_message": o.ValidationMessage,
	}
}

// Location is the enumeration type for location
type ProjectRequestLocation int32

// String returns the string value for Location
func (v ProjectRequestLocation) String() string {
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
	ProjectRequestLocationPrivate ProjectRequestLocation = 0
	// LocationCloud is the enumeration value for cloud
	ProjectRequestLocationCloud ProjectRequestLocation = 1
)

// ProjectRequestProgress represents the object structure for progress
type ProjectRequestProgress struct {
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
}

func (o *ProjectRequestProgress) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Completed The total amount processed thus far
		"completed": o.Completed,
		// Message Any relevant messaging during processing
		"message": o.Message,
		// Total The total amount to be processed
		"total": o.Total,
	}
}

// ProjectRequest an agent action to request adding new projects
type ProjectRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Date the date when the request was made
	Date ProjectRequestDate `json:"date" bson:"date" yaml:"date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration ProjectRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location ProjectRequestLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
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
var _ datamodel.Model = (*ProjectRequest)(nil)

func toProjectRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toProjectRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toProjectRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toProjectRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toProjectRequestObjectNil(isavro, isoptional)
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
	case *ProjectRequest:
		return v.ToMap()
	case ProjectRequest:
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
			arr = append(arr, toProjectRequestObject(av, isavro, false, ""))
		}
		return arr

	case ProjectRequestAuthorization:
		vv := o.(ProjectRequestAuthorization)
		return vv.ToMap()
	case *ProjectRequestAuthorization:
		return (*o.(*ProjectRequestAuthorization)).ToMap()
	case []ProjectRequestAuthorization:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ProjectRequestAuthorization) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ProjectRequestAuthorization:
		arr := make([]interface{}, 0)
		vv := o.(*[]ProjectRequestAuthorization)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ProjectRequestDate:
		vv := o.(ProjectRequestDate)
		return vv.ToMap()
	case *ProjectRequestDate:
		return (*o.(*ProjectRequestDate)).ToMap()
	case []ProjectRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ProjectRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ProjectRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]ProjectRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ProjectRequestIntegration:
		vv := o.(ProjectRequestIntegration)
		return vv.ToMap()
	case *ProjectRequestIntegration:
		return (*o.(*ProjectRequestIntegration)).ToMap()
	case []ProjectRequestIntegration:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ProjectRequestIntegration) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ProjectRequestIntegration:
		arr := make([]interface{}, 0)
		vv := o.(*[]ProjectRequestIntegration)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case ProjectRequestLocation:
		if !isavro {
			return (o.(ProjectRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(ProjectRequestLocation)).String(),
		}
	case *ProjectRequestLocation:
		if !isavro {
			return (o.(*ProjectRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(*ProjectRequestLocation)).String(),
		}
	case ProjectRequestProgress:
		vv := o.(ProjectRequestProgress)
		return vv.ToMap()
	case *ProjectRequestProgress:
		return (*o.(*ProjectRequestProgress)).ToMap()
	case []ProjectRequestProgress:
		arr := make([]interface{}, 0)
		for _, i := range o.([]ProjectRequestProgress) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]ProjectRequestProgress:
		arr := make([]interface{}, 0)
		vv := o.(*[]ProjectRequestProgress)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of ProjectRequest
func (o *ProjectRequest) String() string {
	return fmt.Sprintf("agent.ProjectRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectRequest) GetTopicName() datamodel.TopicNameType {
	return ProjectRequestTopic
}

// GetModelName returns the name of the model
func (o *ProjectRequest) GetModelName() datamodel.ModelNameType {
	return ProjectRequestModelName
}

func (o *ProjectRequest) setDefaults() {
	o.Integration.Authorization.AccessToken = &emptyString
	o.Integration.Authorization.APIToken = &emptyString
	o.Integration.Authorization.Authorization = &emptyString
	o.Integration.Authorization.Password = &emptyString
	o.Integration.Authorization.RefreshToken = &emptyString
	o.Integration.Authorization.URL = &emptyString
	o.Integration.Authorization.Username = &emptyString
	o.Integration.Errored = &emptyBool
	o.Integration.Exclusions = []string{}
	o.Integration.Validated = &emptyBool
	o.Integration.ValidatedAt = &emptyInt
	o.Integration.ValidationMessage = &emptyString

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ProjectRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectRequest) GetTimestamp() time.Time {
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
	case ProjectRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for ProjectRequest")
}

// GetRefID returns the RefID for the object
func (o *ProjectRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Timestamp:         "date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *ProjectRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *ProjectRequest) GetCustomerID() string {
	return o.CustomerID
}

// Clone returns an exact copy of ProjectRequest
func (o *ProjectRequest) Clone() datamodel.Model {
	c := new(ProjectRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectRequest) Anon() datamodel.Model {
	c := new(ProjectRequest)
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
func (o *ProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecProjectRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *ProjectRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecProjectRequest == nil {
		c, err := GetProjectRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecProjectRequest = c
	}
	return cachedCodecProjectRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *ProjectRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *ProjectRequest) FromAvroBinary(value []byte) error {
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
func (o *ProjectRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ProjectRequest objects are equal
func (o *ProjectRequest) IsEqual(other *ProjectRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"customer_id": toProjectRequestObject(o.CustomerID, isavro, false, "string"),
		"date":        toProjectRequestObject(o.Date, isavro, false, "date"),
		"id":          toProjectRequestObject(o.ID, isavro, false, "string"),
		"integration": toProjectRequestObject(o.Integration, isavro, false, "integration"),
		"location":    toProjectRequestObject(o.Location, isavro, false, "location"),
		"ref_id":      toProjectRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":    toProjectRequestObject(o.RefType, isavro, false, "string"),
		"uuid":        toProjectRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":    toProjectRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequest) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["date"].(ProjectRequestDate); ok {
		o.Date = val
	} else {
		val := kv["date"]
		if val == nil {
			o.Date = ProjectRequestDate{}
		} else {
			o.Date = ProjectRequestDate{}
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
	if val, ok := kv["integration"].(ProjectRequestIntegration); ok {
		o.Integration = val
	} else {
		val := kv["integration"]
		if val == nil {
			o.Integration = ProjectRequestIntegration{}
		} else {
			o.Integration = ProjectRequestIntegration{}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.Integration)

		}
	}
	if val, ok := kv["location"].(ProjectRequestLocation); ok {
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
func (o *ProjectRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.Date)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.Location)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetProjectRequestAvroSchemaSpec creates the avro schema specification for ProjectRequest
func GetProjectRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "ProjectRequest",
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
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}, "doc": "the date when the request was made", "type": "record", "name": "date"},
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration",
				"type": map[string]interface{}{"type": "record", "name": "integration", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "If true, the integration is still active"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "integration.authorization", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "access_token", "doc": "Access token"}, map[string]interface{}{"type": "string", "name": "api_token", "doc": "API Token for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "authorization", "doc": "the agents encrypted authorization"}, map[string]interface{}{"doc": "Password for instance, if relevant", "type": "string", "name": "password"}, map[string]interface{}{"type": "string", "name": "refresh_token", "doc": "Refresh token"}, map[string]interface{}{"name": "url", "doc": "URL of instance if relevant", "type": "string"}, map[string]interface{}{"type": "string", "name": "username", "doc": "Username for instance, if relevant"}}, "doc": "Authorization information"}, "name": "authorization", "doc": "Authorization information"}, map[string]interface{}{"type": "boolean", "name": "errored", "doc": "If authorization failed by the agent"}, map[string]interface{}{"type": map[string]interface{}{"type": "array", "name": "exclusions", "items": "string"}, "name": "exclusions", "doc": "The exclusion list for this integration"}, map[string]interface{}{"type": "string", "name": "name", "doc": "The user friendly name of the integration"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "integration.progress", "fields": []interface{}{map[string]interface{}{"type": "long", "name": "completed", "doc": "The total amount processed thus far"}, map[string]interface{}{"type": "string", "name": "message", "doc": "Any relevant messaging during processing"}, map[string]interface{}{"doc": "The total amount to be processed", "type": "long", "name": "total"}}, "doc": "Agent processing progress"}, "name": "progress", "doc": "Agent processing progress"}, map[string]interface{}{"type": "boolean", "name": "validated", "doc": "If the validation has been run against this instance"}, map[string]interface{}{"type": "long", "name": "validated_ts", "doc": "Timestamp when validated"}, map[string]interface{}{"type": "string", "name": "validation_message", "doc": "The validation message from the agent"}}, "doc": "the integration details to use"},
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
				"name": "uuid",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetProjectRequestAvroSchema creates the avro schema for ProjectRequest
func GetProjectRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetProjectRequestAvroSchemaSpec())
}

// TransformProjectRequestFunc is a function for transforming ProjectRequest during processing
type TransformProjectRequestFunc func(input *ProjectRequest) (*ProjectRequest, error)

// NewProjectRequestPipe creates a pipe for processing ProjectRequest items
func NewProjectRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformProjectRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewProjectRequestInputStream(input, errors)
	var stream chan ProjectRequest
	if len(transforms) > 0 {
		stream = make(chan ProjectRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewProjectRequestOutputStream(output, stream, errors)
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

// NewProjectRequestInputStreamDir creates a channel for reading ProjectRequest as JSON newlines from a directory of files
func NewProjectRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformProjectRequestFunc) (chan ProjectRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/project_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan ProjectRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for project_request")
		ch := make(chan ProjectRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewProjectRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan ProjectRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewProjectRequestInputStreamFile creates an channel for reading ProjectRequest as JSON newlines from filename
func NewProjectRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformProjectRequestFunc) (chan ProjectRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan ProjectRequest)
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
			ch := make(chan ProjectRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewProjectRequestInputStream(f, errors, transforms...)
}

// NewProjectRequestInputStream creates an channel for reading ProjectRequest as JSON newlines from stream
func NewProjectRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformProjectRequestFunc) (chan ProjectRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan ProjectRequest, 1000)
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
			var item ProjectRequest
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

// NewProjectRequestOutputStreamDir will output json newlines from channel and save in dir
func NewProjectRequestOutputStreamDir(dir string, ch chan ProjectRequest, errors chan<- error, transforms ...TransformProjectRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/project_request\\.json(\\.gz)?$")
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
	return NewProjectRequestOutputStream(gz, ch, errors, transforms...)
}

// NewProjectRequestOutputStream will output json newlines from channel to the stream
func NewProjectRequestOutputStream(stream io.WriteCloser, ch chan ProjectRequest, errors chan<- error, transforms ...TransformProjectRequestFunc) <-chan bool {
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

// ProjectRequestSendEvent is an event detail for sending data
type ProjectRequestSendEvent struct {
	ProjectRequest *ProjectRequest
	headers        map[string]string
	time           time.Time
	key            string
}

var _ datamodel.ModelSendEvent = (*ProjectRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *ProjectRequestSendEvent) Key() string {
	if e.key == "" {
		return e.ProjectRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *ProjectRequestSendEvent) Object() datamodel.Model {
	return e.ProjectRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *ProjectRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *ProjectRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// ProjectRequestSendEventOpts is a function handler for setting opts
type ProjectRequestSendEventOpts func(o *ProjectRequestSendEvent)

// WithProjectRequestSendEventKey sets the key value to a value different than the object ID
func WithProjectRequestSendEventKey(key string) ProjectRequestSendEventOpts {
	return func(o *ProjectRequestSendEvent) {
		o.key = key
	}
}

// WithProjectRequestSendEventTimestamp sets the timestamp value
func WithProjectRequestSendEventTimestamp(tv time.Time) ProjectRequestSendEventOpts {
	return func(o *ProjectRequestSendEvent) {
		o.time = tv
	}
}

// WithProjectRequestSendEventHeader sets the timestamp value
func WithProjectRequestSendEventHeader(key, value string) ProjectRequestSendEventOpts {
	return func(o *ProjectRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewProjectRequestSendEvent returns a new ProjectRequestSendEvent instance
func NewProjectRequestSendEvent(o *ProjectRequest, opts ...ProjectRequestSendEventOpts) *ProjectRequestSendEvent {
	res := &ProjectRequestSendEvent{
		ProjectRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewProjectRequestProducer will stream data from the channel
func NewProjectRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*ProjectRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.ProjectRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewProjectRequestConsumer will stream data from the topic into the provided channel
func NewProjectRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object ProjectRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.ProjectRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.ProjectRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.ProjectRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &ProjectRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object ProjectRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &ProjectRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// ProjectRequestReceiveEvent is an event detail for receiving data
type ProjectRequestReceiveEvent struct {
	ProjectRequest *ProjectRequest
	message        eventing.Message
	eof            bool
}

var _ datamodel.ModelReceiveEvent = (*ProjectRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *ProjectRequestReceiveEvent) Object() datamodel.Model {
	return e.ProjectRequest
}

// Message returns the underlying message data for the event
func (e *ProjectRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *ProjectRequestReceiveEvent) EOF() bool {
	return e.eof
}

// ProjectRequestProducer implements the datamodel.ModelEventProducer
type ProjectRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*ProjectRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *ProjectRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *ProjectRequestProducer) Close() error {
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
func (o *ProjectRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *ProjectRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ProjectRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewProjectRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewProjectRequestProducerChannel returns a channel which can be used for producing Model events
func NewProjectRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewProjectRequestProducerChannelSize(producer, 0, errors)
}

// NewProjectRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewProjectRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &ProjectRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewProjectRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// ProjectRequestConsumer implements the datamodel.ModelEventConsumer
type ProjectRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*ProjectRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *ProjectRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *ProjectRequestConsumer) Close() error {
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
func (o *ProjectRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ProjectRequestConsumer{
		ch:       ch,
		callback: NewProjectRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewProjectRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewProjectRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &ProjectRequestConsumer{
		ch:       ch,
		callback: NewProjectRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
