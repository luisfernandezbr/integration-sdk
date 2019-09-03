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
	"strings"
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
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// IntegrationRequestTopic is the default topic name
	IntegrationRequestTopic datamodel.TopicNameType = "agent_IntegrationRequest_topic"

	// IntegrationRequestStream is the default stream name
	IntegrationRequestStream datamodel.TopicNameType = "agent_IntegrationRequest_stream"

	// IntegrationRequestTable is the default table name
	IntegrationRequestTable datamodel.TopicNameType = "agent_integrationrequest"

	// IntegrationRequestModelName is the model name
	IntegrationRequestModelName datamodel.ModelNameType = "agent.IntegrationRequest"
)

const (
	// IntegrationRequestCustomerIDColumn is the customer_id column name
	IntegrationRequestCustomerIDColumn = "customer_id"
	// IntegrationRequestIDColumn is the id column name
	IntegrationRequestIDColumn = "id"
	// IntegrationRequestIntegrationColumn is the integration column name
	IntegrationRequestIntegrationColumn = "integration"
	// IntegrationRequestIntegrationColumnActiveColumn is the active column property of the Integration name
	IntegrationRequestIntegrationColumnActiveColumn = "integration->active"
	// IntegrationRequestIntegrationColumnAuthorizationColumn is the authorization column property of the Integration name
	IntegrationRequestIntegrationColumnAuthorizationColumn = "integration->authorization"
	// IntegrationRequestIntegrationColumnCustomerIDColumn is the customer_id column property of the Integration name
	IntegrationRequestIntegrationColumnCustomerIDColumn = "integration->customer_id"
	// IntegrationRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	IntegrationRequestIntegrationColumnErroredColumn = "integration->errored"
	// IntegrationRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	IntegrationRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// IntegrationRequestIntegrationColumnIDColumn is the id column property of the Integration name
	IntegrationRequestIntegrationColumnIDColumn = "integration->id"
	// IntegrationRequestIntegrationColumnLocationColumn is the location column property of the Integration name
	IntegrationRequestIntegrationColumnLocationColumn = "integration->location"
	// IntegrationRequestIntegrationColumnNameColumn is the name column property of the Integration name
	IntegrationRequestIntegrationColumnNameColumn = "integration->name"
	// IntegrationRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	IntegrationRequestIntegrationColumnProgressColumn = "integration->progress"
	// IntegrationRequestIntegrationColumnRefIDColumn is the ref_id column property of the Integration name
	IntegrationRequestIntegrationColumnRefIDColumn = "integration->ref_id"
	// IntegrationRequestIntegrationColumnRefTypeColumn is the ref_type column property of the Integration name
	IntegrationRequestIntegrationColumnRefTypeColumn = "integration->ref_type"
	// IntegrationRequestIntegrationColumnSystemTypeColumn is the system_type column property of the Integration name
	IntegrationRequestIntegrationColumnSystemTypeColumn = "integration->system_type"
	// IntegrationRequestIntegrationColumnValidatedColumn is the validated column property of the Integration name
	IntegrationRequestIntegrationColumnValidatedColumn = "integration->validated"
	// IntegrationRequestIntegrationColumnValidatedDateColumn is the validated_date column property of the Integration name
	IntegrationRequestIntegrationColumnValidatedDateColumn = "integration->validated_date"
	// IntegrationRequestIntegrationColumnValidationMessageColumn is the validation_message column property of the Integration name
	IntegrationRequestIntegrationColumnValidationMessageColumn = "integration->validation_message"
	// IntegrationRequestLocationColumn is the location column name
	IntegrationRequestLocationColumn = "location"
	// IntegrationRequestRefIDColumn is the ref_id column name
	IntegrationRequestRefIDColumn = "ref_id"
	// IntegrationRequestRefTypeColumn is the ref_type column name
	IntegrationRequestRefTypeColumn = "ref_type"
	// IntegrationRequestRequestDateColumn is the request_date column name
	IntegrationRequestRequestDateColumn = "request_date"
	// IntegrationRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	IntegrationRequestRequestDateColumnEpochColumn = "request_date->epoch"
	// IntegrationRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	IntegrationRequestRequestDateColumnOffsetColumn = "request_date->offset"
	// IntegrationRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	IntegrationRequestRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// IntegrationRequestSystemTypeColumn is the system_type column name
	IntegrationRequestSystemTypeColumn = "system_type"
	// IntegrationRequestUpdatedAtColumn is the updated_ts column name
	IntegrationRequestUpdatedAtColumn = "updated_ts"
	// IntegrationRequestUUIDColumn is the uuid column name
	IntegrationRequestUUIDColumn = "uuid"
)

// IntegrationRequestIntegrationAuthorization represents the object structure for authorization
type IntegrationRequestIntegrationAuthorization struct {
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
	// Collection Collection for instance, if relevant
	Collection *string `json:"collection" bson:"collection" yaml:"collection" faker:"-"`
	// APIKey API Key for instance, if relevant
	APIKey *string `json:"api_key" bson:"api_key" yaml:"api_key" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
}

func toIntegrationRequestIntegrationAuthorizationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestIntegrationAuthorizationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntegrationRequestIntegrationAuthorization) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toIntegrationRequestIntegrationAuthorizationObject(o.AccessToken, isavro, true, "string"),
		// RefreshToken Refresh token
		"refresh_token": toIntegrationRequestIntegrationAuthorizationObject(o.RefreshToken, isavro, true, "string"),
		// URL URL of instance if relevant
		"url": toIntegrationRequestIntegrationAuthorizationObject(o.URL, isavro, true, "string"),
		// Username Username for instance, if relevant
		"username": toIntegrationRequestIntegrationAuthorizationObject(o.Username, isavro, true, "string"),
		// Password Password for instance, if relevant
		"password": toIntegrationRequestIntegrationAuthorizationObject(o.Password, isavro, true, "string"),
		// APIToken API Token for instance, if relevant
		"api_token": toIntegrationRequestIntegrationAuthorizationObject(o.APIToken, isavro, true, "string"),
		// Collection Collection for instance, if relevant
		"collection": toIntegrationRequestIntegrationAuthorizationObject(o.Collection, isavro, true, "string"),
		// APIKey API Key for instance, if relevant
		"api_key": toIntegrationRequestIntegrationAuthorizationObject(o.APIKey, isavro, true, "string"),
		// Authorization the agents encrypted authorization
		"authorization": toIntegrationRequestIntegrationAuthorizationObject(o.Authorization, isavro, true, "string"),
	}
}

func (o *IntegrationRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["access_token"].(*string); ok {
		o.AccessToken = val
	} else if val, ok := kv["access_token"].(string); ok {
		o.AccessToken = &val
	} else {
		if val, ok := kv["access_token"]; ok {
			if val == nil {
				o.AccessToken = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AccessToken = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["refresh_token"].(*string); ok {
		o.RefreshToken = val
	} else if val, ok := kv["refresh_token"].(string); ok {
		o.RefreshToken = &val
	} else {
		if val, ok := kv["refresh_token"]; ok {
			if val == nil {
				o.RefreshToken = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefreshToken = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["url"].(*string); ok {
		o.URL = val
	} else if val, ok := kv["url"].(string); ok {
		o.URL = &val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.URL = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["username"].(*string); ok {
		o.Username = val
	} else if val, ok := kv["username"].(string); ok {
		o.Username = &val
	} else {
		if val, ok := kv["username"]; ok {
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
	}

	if val, ok := kv["password"].(*string); ok {
		o.Password = val
	} else if val, ok := kv["password"].(string); ok {
		o.Password = &val
	} else {
		if val, ok := kv["password"]; ok {
			if val == nil {
				o.Password = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Password = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["api_token"].(*string); ok {
		o.APIToken = val
	} else if val, ok := kv["api_token"].(string); ok {
		o.APIToken = &val
	} else {
		if val, ok := kv["api_token"]; ok {
			if val == nil {
				o.APIToken = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.APIToken = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["collection"].(*string); ok {
		o.Collection = val
	} else if val, ok := kv["collection"].(string); ok {
		o.Collection = &val
	} else {
		if val, ok := kv["collection"]; ok {
			if val == nil {
				o.Collection = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Collection = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["api_key"].(*string); ok {
		o.APIKey = val
	} else if val, ok := kv["api_key"].(string); ok {
		o.APIKey = &val
	} else {
		if val, ok := kv["api_key"]; ok {
			if val == nil {
				o.APIKey = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.APIKey = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["authorization"].(*string); ok {
		o.Authorization = val
	} else if val, ok := kv["authorization"].(string); ok {
		o.Authorization = &val
	} else {
		if val, ok := kv["authorization"]; ok {
			if val == nil {
				o.Authorization = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Authorization = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationRequestIntegrationLocation is the enumeration type for location
type IntegrationRequestIntegrationLocation int32

// String returns the string value for IntegrationLocation
func (v IntegrationRequestIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

const (
	// IntegrationLocationPrivate is the enumeration value for private
	IntegrationRequestIntegrationLocationPrivate IntegrationRequestIntegrationLocation = 0
	// IntegrationLocationCloud is the enumeration value for cloud
	IntegrationRequestIntegrationLocationCloud IntegrationRequestIntegrationLocation = 1
)

// IntegrationRequestIntegrationProgress represents the object structure for progress
type IntegrationRequestIntegrationProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func toIntegrationRequestIntegrationProgressObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestIntegrationProgressObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequestIntegrationProgress:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntegrationRequestIntegrationProgress) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": toIntegrationRequestIntegrationProgressObject(o.Message, isavro, false, "string"),
		// Total The total amount to be processed
		"total": toIntegrationRequestIntegrationProgressObject(o.Total, isavro, false, "long"),
		// Completed The total amount processed thus far
		"completed": toIntegrationRequestIntegrationProgressObject(o.Completed, isavro, false, "long"),
	}
}

func (o *IntegrationRequestIntegrationProgress) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequestIntegrationProgress) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["total"].(int64); ok {
		o.Total = val
	} else {
		if val, ok := kv["total"]; ok {
			if val == nil {
				o.Total = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Total = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["completed"].(int64); ok {
		o.Completed = val
	} else {
		if val, ok := kv["completed"]; ok {
			if val == nil {
				o.Completed = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Completed = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationRequestIntegrationSystemType is the enumeration type for system_type
type IntegrationRequestIntegrationSystemType int32

// String returns the string value for IntegrationSystemType
func (v IntegrationRequestIntegrationSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	}
	return "unset"
}

const (
	// IntegrationSystemTypeWork is the enumeration value for work
	IntegrationRequestIntegrationSystemTypeWork IntegrationRequestIntegrationSystemType = 0
	// IntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	IntegrationRequestIntegrationSystemTypeSourcecode IntegrationRequestIntegrationSystemType = 1
	// IntegrationSystemTypeCodequality is the enumeration value for codequality
	IntegrationRequestIntegrationSystemTypeCodequality IntegrationRequestIntegrationSystemType = 2
)

// IntegrationRequestIntegrationValidatedDate represents the object structure for validated_date
type IntegrationRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationRequestIntegrationValidatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestIntegrationValidatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntegrationRequestIntegrationValidatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationRequestIntegrationValidatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntegrationRequestIntegrationValidatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationRequestIntegrationValidatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntegrationRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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

// IntegrationRequestIntegration represents the object structure for integration
type IntegrationRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization IntegrationRequestIntegrationAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location IntegrationRequestIntegrationLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress IntegrationRequestIntegrationProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType IntegrationRequestIntegrationSystemType `json:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate IntegrationRequestIntegrationValidatedDate `json:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func toIntegrationRequestIntegrationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestIntegrationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequestIntegration:
		return v.ToMap(isavro)

	case IntegrationRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	case IntegrationRequestIntegrationLocation:
		return v.String()

	case IntegrationRequestIntegrationProgress:
		return v.ToMap(isavro)

	case IntegrationRequestIntegrationSystemType:
		return v.String()

	case IntegrationRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntegrationRequestIntegration) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toIntegrationRequestIntegrationObject(o.Active, isavro, false, "boolean"),
		// Authorization Authorization information
		"authorization": toIntegrationRequestIntegrationObject(o.Authorization, isavro, false, "authorization"),
		// CustomerID the customer id for the model instance
		"customer_id": toIntegrationRequestIntegrationObject(o.CustomerID, isavro, false, "string"),
		// Errored If authorization failed by the agent
		"errored": toIntegrationRequestIntegrationObject(o.Errored, isavro, true, "boolean"),
		// Exclusions The exclusion list for this integration
		"exclusions": toIntegrationRequestIntegrationObject(o.Exclusions, isavro, false, "exclusions"),
		// ID the primary key for the model instance
		"id": toIntegrationRequestIntegrationObject(o.ID, isavro, false, "string"),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toIntegrationRequestIntegrationObject(o.Location, isavro, false, "location"),
		// Name The user friendly name of the integration
		"name": toIntegrationRequestIntegrationObject(o.Name, isavro, false, "string"),
		// Progress Agent processing progress
		"progress": toIntegrationRequestIntegrationObject(o.Progress, isavro, false, "progress"),
		// RefID the source system id for the model instance
		"ref_id": toIntegrationRequestIntegrationObject(o.RefID, isavro, false, "string"),
		// RefType the source system identifier for the model instance
		"ref_type": toIntegrationRequestIntegrationObject(o.RefType, isavro, false, "string"),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toIntegrationRequestIntegrationObject(o.SystemType, isavro, false, "system_type"),
		// Validated If the validation has been run against this instance
		"validated": toIntegrationRequestIntegrationObject(o.Validated, isavro, true, "boolean"),
		// ValidatedDate Date when validated
		"validated_date": toIntegrationRequestIntegrationObject(o.ValidatedDate, isavro, false, "validated_date"),
		// ValidationMessage The validation message from the agent
		"validation_message": toIntegrationRequestIntegrationObject(o.ValidationMessage, isavro, true, "string"),
	}
}

func (o *IntegrationRequestIntegration) setDefaults(frommap bool) {

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	if o.Validated == nil {
		var v bool
		o.Validated = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequestIntegration) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["active"].(bool); ok {
		o.Active = val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["authorization"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Authorization.FromMap(kv)
		} else if sv, ok := val.(IntegrationRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*IntegrationRequestIntegrationAuthorization); ok {
			// struct pointer
			o.Authorization = *sp
		}
	} else {
		o.Authorization.FromMap(map[string]interface{}{})
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

	if val, ok := kv["errored"].(*bool); ok {
		o.Errored = val
	} else if val, ok := kv["errored"].(bool); ok {
		o.Errored = &val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = number.BoolPointer(number.ToBoolAny(nil))
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Errored = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if val, ok := kv["exclusions"]; ok {
		if val != nil {
			na := make([]string, 0)
			if a, ok := val.([]string); ok {
				na = append(na, a...)
			} else {
				if a, ok := val.([]interface{}); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							if badMap, ok := ae.(map[interface{}]interface{}); ok {
								ae = slice.ConvertToStringToInterface(badMap)
							}
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for exclusions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else if s, ok := val.(string); ok {
					for _, sv := range strings.Split(s, ",") {
						na = append(na, strings.TrimSpace(sv))
					}
				} else if a, ok := val.(primitive.A); ok {
					for _, ae := range a {
						if av, ok := ae.(string); ok {
							na = append(na, av)
						} else {
							b, _ := json.Marshal(ae)
							var av string
							if err := json.Unmarshal(b, &av); err != nil {
								panic("unsupported type for exclusions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for exclusions field")
				}
			}
			o.Exclusions = na
		}
	}
	if o.Exclusions == nil {
		o.Exclusions = make([]string, 0)
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

	if val, ok := kv["location"].(IntegrationRequestIntegrationLocation); ok {
		o.Location = val
	} else {
		if em, ok := kv["location"].(map[string]interface{}); ok {
			ev := em["agent.location"].(string)
			switch ev {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
			}
		}
		if em, ok := kv["location"].(string); ok {
			switch em {
			case "private", "PRIVATE":
				o.Location = 0
			case "cloud", "CLOUD":
				o.Location = 1
			}
		}
	}

	if val, ok := kv["name"].(string); ok {
		o.Name = val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["progress"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Progress.FromMap(kv)
		} else if sv, ok := val.(IntegrationRequestIntegrationProgress); ok {
			// struct
			o.Progress = sv
		} else if sp, ok := val.(*IntegrationRequestIntegrationProgress); ok {
			// struct pointer
			o.Progress = *sp
		}
	} else {
		o.Progress.FromMap(map[string]interface{}{})
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

	if val, ok := kv["system_type"].(IntegrationRequestIntegrationSystemType); ok {
		o.SystemType = val
	} else {
		if em, ok := kv["system_type"].(map[string]interface{}); ok {
			ev := em["agent.system_type"].(string)
			switch ev {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			}
		}
		if em, ok := kv["system_type"].(string); ok {
			switch em {
			case "work", "WORK":
				o.SystemType = 0
			case "sourcecode", "SOURCECODE":
				o.SystemType = 1
			case "codequality", "CODEQUALITY":
				o.SystemType = 2
			}
		}
	}

	if val, ok := kv["validated"].(*bool); ok {
		o.Validated = val
	} else if val, ok := kv["validated"].(bool); ok {
		o.Validated = &val
	} else {
		if val, ok := kv["validated"]; ok {
			if val == nil {
				o.Validated = number.BoolPointer(number.ToBoolAny(nil))
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Validated = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if val, ok := kv["validated_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ValidatedDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*IntegrationRequestIntegrationValidatedDate); ok {
			// struct pointer
			o.ValidatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ValidatedDate.Epoch = dt.Epoch
			o.ValidatedDate.Rfc3339 = dt.Rfc3339
			o.ValidatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ValidatedDate.Epoch = dt.Epoch
			o.ValidatedDate.Rfc3339 = dt.Rfc3339
			o.ValidatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ValidatedDate.Epoch = dt.Epoch
				o.ValidatedDate.Rfc3339 = dt.Rfc3339
				o.ValidatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.ValidatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["validation_message"].(*string); ok {
		o.ValidationMessage = val
	} else if val, ok := kv["validation_message"].(string); ok {
		o.ValidationMessage = &val
	} else {
		if val, ok := kv["validation_message"]; ok {
			if val == nil {
				o.ValidationMessage = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ValidationMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// IntegrationRequestRequestDate represents the object structure for request_date
type IntegrationRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationRequestRequestDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toIntegrationRequestRequestDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequestRequestDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *IntegrationRequestRequestDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationRequestRequestDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toIntegrationRequestRequestDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationRequestRequestDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *IntegrationRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

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

// IntegrationRequest an agent action to request adding an integration
type IntegrationRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to add
	Integration IntegrationRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate IntegrationRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *IntegrationRequest:
		return v.ToMap(isavro)

	case IntegrationRequestIntegration:
		return v.ToMap(isavro)

	case IntegrationRequestRequestDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
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

// NewIntegrationRequestID provides a template for generating an ID field for IntegrationRequest
func NewIntegrationRequestID(customerID string, refType string, refID string) string {
	return hash.Values("IntegrationRequest", customerID, refType, refID)
}

func (o *IntegrationRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("IntegrationRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationRequest) GetID() string {
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
	var dt interface{} = o.UpdatedAt
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
	retention, err := time.ParseDuration("87360h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 87360h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		CleanupPolicy:     datamodel.CleanupPolicy("compact"),
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *IntegrationRequest) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *IntegrationRequest) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

var cachedCodecIntegrationRequest *goavro.Codec
var cachedCodecIntegrationRequestLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *IntegrationRequest) GetAvroCodec() *goavro.Codec {
	cachedCodecIntegrationRequestLock.Lock()
	if cachedCodecIntegrationRequest == nil {
		c, err := GetIntegrationRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecIntegrationRequest = c
	}
	cachedCodecIntegrationRequestLock.Unlock()
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
	o.Hash()
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
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toIntegrationRequestObject(o.CustomerID, isavro, false, "string"),
		"id":           toIntegrationRequestObject(o.ID, isavro, false, "string"),
		"integration":  toIntegrationRequestObject(o.Integration, isavro, false, "integration"),
		"ref_id":       toIntegrationRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toIntegrationRequestObject(o.RefType, isavro, false, "string"),
		"request_date": toIntegrationRequestObject(o.RequestDate, isavro, false, "request_date"),
		"updated_ts":   toIntegrationRequestObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":         toIntegrationRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toIntegrationRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationRequest) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["integration"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Integration.FromMap(kv)
		} else if sv, ok := val.(IntegrationRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*IntegrationRequestIntegration); ok {
			// struct pointer
			o.Integration = *sp
		}
	} else {
		o.Integration.FromMap(map[string]interface{}{})
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*IntegrationRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UpdatedAt)
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
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration",
				"type": map[string]interface{}{"doc": "the integration details to add", "fields": []interface{}{map[string]interface{}{"doc": "If true, the integration is still active", "name": "active", "type": "boolean"}, map[string]interface{}{"doc": "Authorization information", "name": "authorization", "type": map[string]interface{}{"doc": "Authorization information", "fields": []interface{}{map[string]interface{}{"default": nil, "doc": "Access token", "name": "access_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Refresh token", "name": "refresh_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "URL of instance if relevant", "name": "url", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Username for instance, if relevant", "name": "username", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Password for instance, if relevant", "name": "password", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Token for instance, if relevant", "name": "api_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Collection for instance, if relevant", "name": "collection", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Key for instance, if relevant", "name": "api_key", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "the agents encrypted authorization", "name": "authorization", "type": []interface{}{"null", "string"}}}, "name": "integration.authorization", "type": "record"}}, map[string]interface{}{"doc": "the customer id for the model instance", "name": "customer_id", "type": "string"}, map[string]interface{}{"default": nil, "doc": "If authorization failed by the agent", "name": "errored", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "The exclusion list for this integration", "name": "exclusions", "type": map[string]interface{}{"items": "string", "name": "exclusions", "type": "array"}}, map[string]interface{}{"doc": "the primary key for the model instance", "name": "id", "type": "string"}, map[string]interface{}{"doc": "The location of this integration (on-premise / private or cloud)", "name": "location", "type": map[string]interface{}{"default": "PRIVATE", "doc": "The location of this integration (on-premise / private or cloud)", "name": "integration.location", "symbols": []string{"PRIVATE", "CLOUD"}, "type": "enum"}}, map[string]interface{}{"doc": "The user friendly name of the integration", "name": "name", "type": "string"}, map[string]interface{}{"doc": "Agent processing progress", "name": "progress", "type": map[string]interface{}{"doc": "Agent processing progress", "fields": []interface{}{map[string]interface{}{"doc": "Any relevant messaging during processing", "name": "message", "type": "string"}, map[string]interface{}{"doc": "The total amount to be processed", "name": "total", "type": "long"}, map[string]interface{}{"doc": "The total amount processed thus far", "name": "completed", "type": "long"}}, "name": "integration.progress", "type": "record"}}, map[string]interface{}{"doc": "the source system id for the model instance", "name": "ref_id", "type": "string"}, map[string]interface{}{"doc": "the source system identifier for the model instance", "name": "ref_type", "type": "string"}, map[string]interface{}{"doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "system_type", "type": map[string]interface{}{"default": "WORK", "doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "integration.system_type", "symbols": []string{"WORK", "SOURCECODE", "CODEQUALITY"}, "type": "enum"}}, map[string]interface{}{"default": nil, "doc": "If the validation has been run against this instance", "name": "validated", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "Date when validated", "name": "validated_date", "type": map[string]interface{}{"doc": "Date when validated", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "integration.validated_date", "type": "record"}}, map[string]interface{}{"default": nil, "doc": "The validation message from the agent", "name": "validation_message", "type": []interface{}{"null", "string"}}}, "name": "integration", "type": "record"},
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
				"type": map[string]interface{}{"doc": "the date when the request was made", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "request_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
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
func (o *IntegrationRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
	emptyTime := time.Unix(0, 0)
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
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
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
					return fmt.Errorf("error unmarshaling avro data into agent.IntegrationRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.IntegrationRequest")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
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
