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
	// RepoRequestTopic is the default topic name
	RepoRequestTopic datamodel.TopicNameType = "agent_RepoRequest_topic"

	// RepoRequestStream is the default stream name
	RepoRequestStream datamodel.TopicNameType = "agent_RepoRequest_stream"

	// RepoRequestTable is the default table name
	RepoRequestTable datamodel.TopicNameType = "agent_RepoRequest"

	// RepoRequestModelName is the model name
	RepoRequestModelName datamodel.ModelNameType = "agent.RepoRequest"
)

const (
	// RepoRequestCustomerIDColumn is the customer_id column name
	RepoRequestCustomerIDColumn = "customer_id"
	// RepoRequestIDColumn is the id column name
	RepoRequestIDColumn = "id"
	// RepoRequestIntegrationColumn is the integration column name
	RepoRequestIntegrationColumn = "integration"
	// RepoRequestIntegrationColumnActiveColumn is the active column property of the Integration name
	RepoRequestIntegrationColumnActiveColumn = "integration->active"
	// RepoRequestIntegrationColumnAuthorizationColumn is the authorization column property of the Integration name
	RepoRequestIntegrationColumnAuthorizationColumn = "integration->authorization"
	// RepoRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	RepoRequestIntegrationColumnErroredColumn = "integration->errored"
	// RepoRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	RepoRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// RepoRequestIntegrationColumnNameColumn is the name column property of the Integration name
	RepoRequestIntegrationColumnNameColumn = "integration->name"
	// RepoRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	RepoRequestIntegrationColumnProgressColumn = "integration->progress"
	// RepoRequestIntegrationColumnValidatedColumn is the validated column property of the Integration name
	RepoRequestIntegrationColumnValidatedColumn = "integration->validated"
	// RepoRequestIntegrationColumnValidatedDateColumn is the validated_date column property of the Integration name
	RepoRequestIntegrationColumnValidatedDateColumn = "integration->validated_date"
	// RepoRequestIntegrationColumnValidationMessageColumn is the validation_message column property of the Integration name
	RepoRequestIntegrationColumnValidationMessageColumn = "integration->validation_message"
	// RepoRequestLocationColumn is the location column name
	RepoRequestLocationColumn = "location"
	// RepoRequestRefIDColumn is the ref_id column name
	RepoRequestRefIDColumn = "ref_id"
	// RepoRequestRefTypeColumn is the ref_type column name
	RepoRequestRefTypeColumn = "ref_type"
	// RepoRequestRequestDateColumn is the request_date column name
	RepoRequestRequestDateColumn = "request_date"
	// RepoRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	RepoRequestRequestDateColumnEpochColumn = "request_date->epoch"
	// RepoRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	RepoRequestRequestDateColumnOffsetColumn = "request_date->offset"
	// RepoRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	RepoRequestRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// RepoRequestUUIDColumn is the uuid column name
	RepoRequestUUIDColumn = "uuid"
)

// 0 customer_id
// customer_id {"description":"the customer id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"customer_id","relation":true,"subtype":"","type":"string"}

// 1 id
// id {"description":"the primary key for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"id","relation":false,"subtype":"","type":"string"}

// 2 integration
// integration {"description":"the integration details to use","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"integration","relation":false,"subtype":"object","type":"object"}

// active {"description":"If true, the integration is still active","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"active","relation":false,"subtype":"","type":"boolean"}

// authorization {"description":"Authorization information","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"authorization","relation":false,"subtype":"","type":"object"}

// access_token {"description":"Access token","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"access_token","relation":false,"subtype":"","type":"string"}

// refresh_token {"description":"Refresh token","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"refresh_token","relation":false,"subtype":"","type":"string"}

// url {"description":"URL of instance if relevant","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"url","relation":false,"subtype":"","type":"string"}

// username {"description":"Username for instance, if relevant","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"username","relation":false,"subtype":"","type":"string"}

// password {"description":"Password for instance, if relevant","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"password","relation":false,"subtype":"","type":"string"}

// api_token {"description":"API Token for instance, if relevant","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"api_token","relation":false,"subtype":"","type":"string"}

// authorization {"description":"the agents encrypted authorization","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"authorization","relation":false,"subtype":"","type":"string"}

// RepoRequestIntegrationAuthorization represents the object structure for authorization
type RepoRequestIntegrationAuthorization struct {
	// AccessToken Access token
	AccessToken *string `json:"access_token" bson:"access_token" yaml:"access_token" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken *string `json:"refresh_token" bson:"refresh_token" yaml:"refresh_token" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url" bson:"url" yaml:"url" faker:"-"`
	// Username Username for instance, if relevant
	Username *string `json:"username" bson:"username" yaml:"username" faker:"-"`
	// Password Password for instance, if relevant
	Password *string `json:"password" bson:"password" yaml:"password" faker:"-"`
	// APIToken API Token for instance, if relevant
	APIToken *string `json:"api_token" bson:"api_token" yaml:"api_token" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
}

func (o *RepoRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": o.AccessToken,
		// RefreshToken Refresh token
		"refresh_token": o.RefreshToken,
		// URL URL of instance if relevant
		"url": o.URL,
		// Username Username for instance, if relevant
		"username": o.Username,
		// Password Password for instance, if relevant
		"password": o.Password,
		// APIToken API Token for instance, if relevant
		"api_token": o.APIToken,
		// Authorization the agents encrypted authorization
		"authorization": o.Authorization,
	}
}

// errored {"description":"If authorization failed by the agent","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"errored","relation":false,"subtype":"","type":"boolean"}

// exclusions {"description":"The exclusion list for this integration","is_array":true,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"exclusions","relation":false,"subtype":"","type":"string"}

// name {"description":"The user friendly name of the integration","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"name","relation":false,"subtype":"","type":"string"}

// progress {"description":"Agent processing progress","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"progress","relation":false,"subtype":"","type":"object"}

// message {"description":"Any relevant messaging during processing","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"message","relation":false,"subtype":"","type":"string"}

// total {"description":"The total amount to be processed","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"total","relation":false,"subtype":"","type":"int"}

// completed {"description":"The total amount processed thus far","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"completed","relation":false,"subtype":"","type":"int"}

// RepoRequestIntegrationProgress represents the object structure for progress
type RepoRequestIntegrationProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func (o *RepoRequestIntegrationProgress) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": o.Message,
		// Total The total amount to be processed
		"total": o.Total,
		// Completed The total amount processed thus far
		"completed": o.Completed,
	}
}

// validated {"description":"If the validation has been run against this instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"validated","relation":false,"subtype":"","type":"boolean"}

// validated_date {"description":"Date when validated","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"validated_date","relation":false,"subtype":"","type":"object"}

// rfc3339 {"description":"the date in RFC3339 format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"rfc3339","relation":false,"subtype":"","type":"string"}

// epoch {"description":"the date in epoch format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"epoch","relation":false,"subtype":"","type":"int"}

// offset {"description":"the timezone offset from GMT","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"offset","relation":false,"subtype":"","type":"int"}

// RepoRequestIntegrationValidatedDate represents the object structure for validated_date
type RepoRequestIntegrationValidatedDate struct {
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
}

func (o *RepoRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
	}
}

// validation_message {"description":"The validation message from the agent","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"validation_message","relation":false,"subtype":"","type":"string"}

// RepoRequestIntegration represents the object structure for integration
type RepoRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization RepoRequestIntegrationAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress RepoRequestIntegrationProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate RepoRequestIntegrationValidatedDate `json:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func (o *RepoRequestIntegration) ToMap() map[string]interface{} {
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

// 3 location
// location {"description":"The location of this integration (on-premise / private or cloud)","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"location","relation":false,"subtype":"","type":"enum"}

// RepoRequestLocation is the enumeration type for location
type RepoRequestLocation int32

// String returns the string value for Location
func (v RepoRequestLocation) String() string {
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
	RepoRequestLocationPrivate RepoRequestLocation = 0
	// LocationCloud is the enumeration value for cloud
	RepoRequestLocationCloud RepoRequestLocation = 1
)

// 4 ref_id
// ref_id {"description":"the source system id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_id","relation":false,"subtype":"","type":"string"}

// 5 ref_type
// ref_type {"description":"the source system identifier for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_type","relation":false,"subtype":"","type":"string"}

// 6 request_date
// request_date {"description":"the date when the request was made","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"request_date","relation":false,"subtype":"","type":"object"}

// epoch {"description":"the date in epoch format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"epoch","relation":false,"subtype":"","type":"int"}

// offset {"description":"the timezone offset from GMT","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"offset","relation":false,"subtype":"","type":"int"}

// rfc3339 {"description":"the date in RFC3339 format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"rfc3339","relation":false,"subtype":"","type":"string"}

// RepoRequestRequestDate represents the object structure for request_date
type RepoRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *RepoRequestRequestDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// 7 uuid
// uuid {"description":"the agent unique identifier","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"uuid","relation":false,"subtype":"","type":"string"}

// RepoRequest an agent action to request adding new repos
type RepoRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration RepoRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location RepoRequestLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate RepoRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoRequest)(nil)

func toRepoRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toRepoRequestObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toRepoRequestObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toRepoRequestObjectNil(isavro, isoptional)
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
	case *RepoRequest:
		return v.ToMap()
	case RepoRequest:
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
			arr = append(arr, toRepoRequestObject(av, isavro, false, ""))
		}
		return arr

	case RepoRequestIntegration:
		vv := o.(RepoRequestIntegration)
		return vv.ToMap()
	case *RepoRequestIntegration:
		return map[string]interface{}{
			"agent.integration": (*o.(*RepoRequestIntegration)).ToMap(),
		}
	case []RepoRequestIntegration:
		arr := make([]interface{}, 0)
		for _, i := range o.([]RepoRequestIntegration) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]RepoRequestIntegration:
		arr := make([]interface{}, 0)
		vv := o.(*[]RepoRequestIntegration)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case RepoRequestLocation:
		if !isavro {
			return (o.(RepoRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(RepoRequestLocation)).String(),
		}
	case *RepoRequestLocation:
		if !isavro {
			return (o.(*RepoRequestLocation)).String()
		}
		return map[string]string{
			"agent.location": (o.(*RepoRequestLocation)).String(),
		}
	case RepoRequestRequestDate:
		vv := o.(RepoRequestRequestDate)
		return vv.ToMap()
	case *RepoRequestRequestDate:
		return map[string]interface{}{
			"agent.request_date": (*o.(*RepoRequestRequestDate)).ToMap(),
		}
	case []RepoRequestRequestDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]RepoRequestRequestDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]RepoRequestRequestDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]RepoRequestRequestDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of RepoRequest
func (o *RepoRequest) String() string {
	return fmt.Sprintf("agent.RepoRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoRequest) GetTopicName() datamodel.TopicNameType {
	return RepoRequestTopic
}

// GetModelName returns the name of the model
func (o *RepoRequest) GetModelName() datamodel.ModelNameType {
	return RepoRequestModelName
}

func (o *RepoRequest) setDefaults() {
	if o.Integration.Authorization.AccessToken == nil {
		o.Integration.Authorization.AccessToken = &emptyString
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
	if o.Integration.Authorization.Password == nil {
		o.Integration.Authorization.Password = &emptyString
	}
	if o.Integration.Authorization.APIToken == nil {
		o.Integration.Authorization.APIToken = &emptyString
	}
	if o.Integration.Authorization.Authorization == nil {
		o.Integration.Authorization.Authorization = &emptyString
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
func (o *RepoRequest) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoRequest", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoRequest) GetTimestamp() time.Time {
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
	case RepoRequestRequestDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for RepoRequest")
}

// GetRefID returns the RefID for the object
func (o *RepoRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *RepoRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *RepoRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of RepoRequest
func (o *RepoRequest) Clone() datamodel.Model {
	c := new(RepoRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RepoRequest) Anon() datamodel.Model {
	c := new(RepoRequest)
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
func (o *RepoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoRequest) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecRepoRequest *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *RepoRequest) GetAvroCodec() *goavro.Codec {
	if cachedCodecRepoRequest == nil {
		c, err := GetRepoRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecRepoRequest = c
	}
	return cachedCodecRepoRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *RepoRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *RepoRequest) FromAvroBinary(value []byte) error {
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
func (o *RepoRequest) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two RepoRequest objects are equal
func (o *RepoRequest) IsEqual(other *RepoRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults()
	return map[string]interface{}{
		"customer_id":  toRepoRequestObject(o.CustomerID, isavro, false, "string"),
		"id":           toRepoRequestObject(o.ID, isavro, false, "string"),
		"integration":  toRepoRequestObject(o.Integration, isavro, false, "integration"),
		"location":     toRepoRequestObject(o.Location, isavro, false, "location"),
		"ref_id":       toRepoRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toRepoRequestObject(o.RefType, isavro, false, "string"),
		"request_date": toRepoRequestObject(o.RequestDate, isavro, false, "request_date"),
		"uuid":         toRepoRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toRepoRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequest) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["integration"].(RepoRequestIntegration); ok {
		o.Integration = val
	} else {
		val := kv["integration"]
		if val == nil {
			o.Integration = RepoRequestIntegration{}
		} else {
			o.Integration = RepoRequestIntegration{}
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
	if val, ok := kv["location"].(RepoRequestLocation); ok {
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
	if val, ok := kv["request_date"].(RepoRequestRequestDate); ok {
		o.RequestDate = val
	} else {
		val := kv["request_date"]
		if val == nil {
			o.RequestDate = RepoRequestRequestDate{}
		} else {
			o.RequestDate = RepoRequestRequestDate{}
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
func (o *RepoRequest) Hash() string {
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

// GetRepoRequestAvroSchemaSpec creates the avro schema specification for RepoRequest
func GetRepoRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "RepoRequest",
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
				"type": map[string]interface{}{"name": "integration", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "If true, the integration is still active"}, map[string]interface{}{"name": "authorization", "doc": "Authorization information", "type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"doc": "Access token", "type": "string", "name": "access_token"}, map[string]interface{}{"type": "string", "name": "refresh_token", "doc": "Refresh token"}, map[string]interface{}{"type": "string", "name": "url", "doc": "URL of instance if relevant"}, map[string]interface{}{"type": "string", "name": "username", "doc": "Username for instance, if relevant"}, map[string]interface{}{"type": "string", "name": "password", "doc": "Password for instance, if relevant"}, map[string]interface{}{"doc": "API Token for instance, if relevant", "type": "string", "name": "api_token"}, map[string]interface{}{"type": "string", "name": "authorization", "doc": "the agents encrypted authorization"}}, "doc": "Authorization information", "type": "record", "name": "integration.authorization"}}, map[string]interface{}{"type": "boolean", "name": "errored", "doc": "If authorization failed by the agent"}, map[string]interface{}{"type": map[string]interface{}{"items": "string", "type": "array", "name": "exclusions"}, "name": "exclusions", "doc": "The exclusion list for this integration"}, map[string]interface{}{"type": "string", "name": "name", "doc": "The user friendly name of the integration"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "integration.progress", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "message", "doc": "Any relevant messaging during processing"}, map[string]interface{}{"type": "long", "name": "total", "doc": "The total amount to be processed"}, map[string]interface{}{"name": "completed", "doc": "The total amount processed thus far", "type": "long"}}, "doc": "Agent processing progress"}, "name": "progress", "doc": "Agent processing progress"}, map[string]interface{}{"doc": "If the validation has been run against this instance", "type": "boolean", "name": "validated"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "integration.validated_date", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}, map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"name": "offset", "doc": "the timezone offset from GMT", "type": "long"}}, "doc": "Date when validated"}, "name": "validated_date", "doc": "Date when validated"}, map[string]interface{}{"doc": "The validation message from the agent", "type": "string", "name": "validation_message"}}, "doc": "the integration details to use", "type": "record"},
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
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "type": "long", "name": "epoch"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"name": "rfc3339", "doc": "the date in RFC3339 format", "type": "string"}}, "doc": "the date when the request was made", "type": "record", "name": "request_date"},
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
func (o *RepoRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetRepoRequestAvroSchema creates the avro schema for RepoRequest
func GetRepoRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetRepoRequestAvroSchemaSpec())
}

// TransformRepoRequestFunc is a function for transforming RepoRequest during processing
type TransformRepoRequestFunc func(input *RepoRequest) (*RepoRequest, error)

// NewRepoRequestPipe creates a pipe for processing RepoRequest items
func NewRepoRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewRepoRequestInputStream(input, errors)
	var stream chan RepoRequest
	if len(transforms) > 0 {
		stream = make(chan RepoRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewRepoRequestOutputStream(output, stream, errors)
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

// NewRepoRequestInputStreamDir creates a channel for reading RepoRequest as JSON newlines from a directory of files
func NewRepoRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoRequestFunc) (chan RepoRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/repo_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan RepoRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for repo_request")
		ch := make(chan RepoRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewRepoRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan RepoRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewRepoRequestInputStreamFile creates an channel for reading RepoRequest as JSON newlines from filename
func NewRepoRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoRequestFunc) (chan RepoRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan RepoRequest)
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
			ch := make(chan RepoRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewRepoRequestInputStream(f, errors, transforms...)
}

// NewRepoRequestInputStream creates an channel for reading RepoRequest as JSON newlines from stream
func NewRepoRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoRequestFunc) (chan RepoRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan RepoRequest, 1000)
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
			var item RepoRequest
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

// NewRepoRequestOutputStreamDir will output json newlines from channel and save in dir
func NewRepoRequestOutputStreamDir(dir string, ch chan RepoRequest, errors chan<- error, transforms ...TransformRepoRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/repo_request\\.json(\\.gz)?$")
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
	return NewRepoRequestOutputStream(gz, ch, errors, transforms...)
}

// NewRepoRequestOutputStream will output json newlines from channel to the stream
func NewRepoRequestOutputStream(stream io.WriteCloser, ch chan RepoRequest, errors chan<- error, transforms ...TransformRepoRequestFunc) <-chan bool {
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

// RepoRequestSendEvent is an event detail for sending data
type RepoRequestSendEvent struct {
	RepoRequest *RepoRequest
	headers     map[string]string
	time        time.Time
	key         string
}

var _ datamodel.ModelSendEvent = (*RepoRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *RepoRequestSendEvent) Key() string {
	if e.key == "" {
		return e.RepoRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *RepoRequestSendEvent) Object() datamodel.Model {
	return e.RepoRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *RepoRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *RepoRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// RepoRequestSendEventOpts is a function handler for setting opts
type RepoRequestSendEventOpts func(o *RepoRequestSendEvent)

// WithRepoRequestSendEventKey sets the key value to a value different than the object ID
func WithRepoRequestSendEventKey(key string) RepoRequestSendEventOpts {
	return func(o *RepoRequestSendEvent) {
		o.key = key
	}
}

// WithRepoRequestSendEventTimestamp sets the timestamp value
func WithRepoRequestSendEventTimestamp(tv time.Time) RepoRequestSendEventOpts {
	return func(o *RepoRequestSendEvent) {
		o.time = tv
	}
}

// WithRepoRequestSendEventHeader sets the timestamp value
func WithRepoRequestSendEventHeader(key, value string) RepoRequestSendEventOpts {
	return func(o *RepoRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewRepoRequestSendEvent returns a new RepoRequestSendEvent instance
func NewRepoRequestSendEvent(o *RepoRequest, opts ...RepoRequestSendEventOpts) *RepoRequestSendEvent {
	res := &RepoRequestSendEvent{
		RepoRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewRepoRequestProducer will stream data from the channel
func NewRepoRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*RepoRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.RepoRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewRepoRequestConsumer will stream data from the topic into the provided channel
func NewRepoRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object RepoRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.RepoRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.RepoRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.RepoRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &RepoRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object RepoRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &RepoRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// RepoRequestReceiveEvent is an event detail for receiving data
type RepoRequestReceiveEvent struct {
	RepoRequest *RepoRequest
	message     eventing.Message
	eof         bool
}

var _ datamodel.ModelReceiveEvent = (*RepoRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *RepoRequestReceiveEvent) Object() datamodel.Model {
	return e.RepoRequest
}

// Message returns the underlying message data for the event
func (e *RepoRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *RepoRequestReceiveEvent) EOF() bool {
	return e.eof
}

// RepoRequestProducer implements the datamodel.ModelEventProducer
type RepoRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*RepoRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *RepoRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *RepoRequestProducer) Close() error {
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
func (o *RepoRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *RepoRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewRepoRequestProducerChannel returns a channel which can be used for producing Model events
func NewRepoRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewRepoRequestProducerChannelSize(producer, 0, errors)
}

// NewRepoRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewRepoRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// RepoRequestConsumer implements the datamodel.ModelEventConsumer
type RepoRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*RepoRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *RepoRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *RepoRequestConsumer) Close() error {
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
func (o *RepoRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoRequestConsumer{
		ch:       ch,
		callback: NewRepoRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewRepoRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewRepoRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoRequestConsumer{
		ch:       ch,
		callback: NewRepoRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
