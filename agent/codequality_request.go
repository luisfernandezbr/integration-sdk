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
	// CodequalityRequestTopic is the default topic name
	CodequalityRequestTopic datamodel.TopicNameType = "agent_CodequalityRequest_topic"

	// CodequalityRequestStream is the default stream name
	CodequalityRequestStream datamodel.TopicNameType = "agent_CodequalityRequest_stream"

	// CodequalityRequestTable is the default table name
	CodequalityRequestTable datamodel.TopicNameType = "agent_codequalityrequest"

	// CodequalityRequestModelName is the model name
	CodequalityRequestModelName datamodel.ModelNameType = "agent.CodequalityRequest"
)

const (
	// CodequalityRequestCustomerIDColumn is the customer_id column name
	CodequalityRequestCustomerIDColumn = "customer_id"
	// CodequalityRequestIDColumn is the id column name
	CodequalityRequestIDColumn = "id"
	// CodequalityRequestIntegrationColumn is the integration column name
	CodequalityRequestIntegrationColumn = "integration"
	// CodequalityRequestIntegrationColumnActiveColumn is the active column property of the Integration name
	CodequalityRequestIntegrationColumnActiveColumn = "integration->active"
	// CodequalityRequestIntegrationColumnAuthorizationColumn is the authorization column property of the Integration name
	CodequalityRequestIntegrationColumnAuthorizationColumn = "integration->authorization"
	// CodequalityRequestIntegrationColumnCustomerIDColumn is the customer_id column property of the Integration name
	CodequalityRequestIntegrationColumnCustomerIDColumn = "integration->customer_id"
	// CodequalityRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	CodequalityRequestIntegrationColumnErroredColumn = "integration->errored"
	// CodequalityRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	CodequalityRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// CodequalityRequestIntegrationColumnIDColumn is the id column property of the Integration name
	CodequalityRequestIntegrationColumnIDColumn = "integration->id"
	// CodequalityRequestIntegrationColumnLocationColumn is the location column property of the Integration name
	CodequalityRequestIntegrationColumnLocationColumn = "integration->location"
	// CodequalityRequestIntegrationColumnNameColumn is the name column property of the Integration name
	CodequalityRequestIntegrationColumnNameColumn = "integration->name"
	// CodequalityRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	CodequalityRequestIntegrationColumnProgressColumn = "integration->progress"
	// CodequalityRequestIntegrationColumnRefIDColumn is the ref_id column property of the Integration name
	CodequalityRequestIntegrationColumnRefIDColumn = "integration->ref_id"
	// CodequalityRequestIntegrationColumnRefTypeColumn is the ref_type column property of the Integration name
	CodequalityRequestIntegrationColumnRefTypeColumn = "integration->ref_type"
	// CodequalityRequestIntegrationColumnSystemTypeColumn is the system_type column property of the Integration name
	CodequalityRequestIntegrationColumnSystemTypeColumn = "integration->system_type"
	// CodequalityRequestIntegrationColumnValidatedColumn is the validated column property of the Integration name
	CodequalityRequestIntegrationColumnValidatedColumn = "integration->validated"
	// CodequalityRequestIntegrationColumnValidatedDateColumn is the validated_date column property of the Integration name
	CodequalityRequestIntegrationColumnValidatedDateColumn = "integration->validated_date"
	// CodequalityRequestIntegrationColumnValidationMessageColumn is the validation_message column property of the Integration name
	CodequalityRequestIntegrationColumnValidationMessageColumn = "integration->validation_message"
	// CodequalityRequestLocationColumn is the location column name
	CodequalityRequestLocationColumn = "location"
	// CodequalityRequestRefIDColumn is the ref_id column name
	CodequalityRequestRefIDColumn = "ref_id"
	// CodequalityRequestRefTypeColumn is the ref_type column name
	CodequalityRequestRefTypeColumn = "ref_type"
	// CodequalityRequestRequestDateColumn is the request_date column name
	CodequalityRequestRequestDateColumn = "request_date"
	// CodequalityRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	CodequalityRequestRequestDateColumnEpochColumn = "request_date->epoch"
	// CodequalityRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	CodequalityRequestRequestDateColumnOffsetColumn = "request_date->offset"
	// CodequalityRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	CodequalityRequestRequestDateColumnRfc3339Column = "request_date->rfc3339"
	// CodequalityRequestSystemTypeColumn is the system_type column name
	CodequalityRequestSystemTypeColumn = "system_type"
	// CodequalityRequestUpdatedAtColumn is the updated_ts column name
	CodequalityRequestUpdatedAtColumn = "updated_ts"
	// CodequalityRequestUUIDColumn is the uuid column name
	CodequalityRequestUUIDColumn = "uuid"
)

// CodequalityRequestIntegrationAuthorization represents the object structure for authorization
type CodequalityRequestIntegrationAuthorization struct {
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

func toCodequalityRequestIntegrationAuthorizationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestIntegrationAuthorizationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *CodequalityRequestIntegrationAuthorization) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toCodequalityRequestIntegrationAuthorizationObject(o.AccessToken, isavro, true, "string"),
		// RefreshToken Refresh token
		"refresh_token": toCodequalityRequestIntegrationAuthorizationObject(o.RefreshToken, isavro, true, "string"),
		// URL URL of instance if relevant
		"url": toCodequalityRequestIntegrationAuthorizationObject(o.URL, isavro, true, "string"),
		// Username Username for instance, if relevant
		"username": toCodequalityRequestIntegrationAuthorizationObject(o.Username, isavro, true, "string"),
		// Password Password for instance, if relevant
		"password": toCodequalityRequestIntegrationAuthorizationObject(o.Password, isavro, true, "string"),
		// APIToken API Token for instance, if relevant
		"api_token": toCodequalityRequestIntegrationAuthorizationObject(o.APIToken, isavro, true, "string"),
		// Collection Collection for instance, if relevant
		"collection": toCodequalityRequestIntegrationAuthorizationObject(o.Collection, isavro, true, "string"),
		// APIKey API Key for instance, if relevant
		"api_key": toCodequalityRequestIntegrationAuthorizationObject(o.APIKey, isavro, true, "string"),
		// Authorization the agents encrypted authorization
		"authorization": toCodequalityRequestIntegrationAuthorizationObject(o.Authorization, isavro, true, "string"),
	}
}

func (o *CodequalityRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLocation is the enumeration type for location
type CodequalityRequestIntegrationLocation int32

// String returns the string value for IntegrationLocation
func (v CodequalityRequestIntegrationLocation) String() string {
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
	CodequalityRequestIntegrationLocationPrivate CodequalityRequestIntegrationLocation = 0
	// IntegrationLocationCloud is the enumeration value for cloud
	CodequalityRequestIntegrationLocationCloud CodequalityRequestIntegrationLocation = 1
)

// CodequalityRequestIntegrationProgress represents the object structure for progress
type CodequalityRequestIntegrationProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func toCodequalityRequestIntegrationProgressObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestIntegrationProgressObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequestIntegrationProgress:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *CodequalityRequestIntegrationProgress) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": toCodequalityRequestIntegrationProgressObject(o.Message, isavro, false, "string"),
		// Total The total amount to be processed
		"total": toCodequalityRequestIntegrationProgressObject(o.Total, isavro, false, "long"),
		// Completed The total amount processed thus far
		"completed": toCodequalityRequestIntegrationProgressObject(o.Completed, isavro, false, "long"),
	}
}

func (o *CodequalityRequestIntegrationProgress) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationProgress) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationSystemType is the enumeration type for system_type
type CodequalityRequestIntegrationSystemType int32

// String returns the string value for IntegrationSystemType
func (v CodequalityRequestIntegrationSystemType) String() string {
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
	CodequalityRequestIntegrationSystemTypeWork CodequalityRequestIntegrationSystemType = 0
	// IntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	CodequalityRequestIntegrationSystemTypeSourcecode CodequalityRequestIntegrationSystemType = 1
	// IntegrationSystemTypeCodequality is the enumeration value for codequality
	CodequalityRequestIntegrationSystemTypeCodequality CodequalityRequestIntegrationSystemType = 2
)

// CodequalityRequestIntegrationValidatedDate represents the object structure for validated_date
type CodequalityRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationValidatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestIntegrationValidatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *CodequalityRequestIntegrationValidatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationValidatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationValidatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationValidatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *CodequalityRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegration represents the object structure for integration
type CodequalityRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization CodequalityRequestIntegrationAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location CodequalityRequestIntegrationLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress CodequalityRequestIntegrationProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType CodequalityRequestIntegrationSystemType `json:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate CodequalityRequestIntegrationValidatedDate `json:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func toCodequalityRequestIntegrationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestIntegrationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequestIntegration:
		return v.ToMap(isavro)

	case CodequalityRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	case CodequalityRequestIntegrationLocation:
		return v.String()

	case CodequalityRequestIntegrationProgress:
		return v.ToMap(isavro)

	case CodequalityRequestIntegrationSystemType:
		return v.String()

	case CodequalityRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *CodequalityRequestIntegration) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toCodequalityRequestIntegrationObject(o.Active, isavro, false, "boolean"),
		// Authorization Authorization information
		"authorization": toCodequalityRequestIntegrationObject(o.Authorization, isavro, false, "authorization"),
		// CustomerID the customer id for the model instance
		"customer_id": toCodequalityRequestIntegrationObject(o.CustomerID, isavro, false, "string"),
		// Errored If authorization failed by the agent
		"errored": toCodequalityRequestIntegrationObject(o.Errored, isavro, true, "boolean"),
		// Exclusions The exclusion list for this integration
		"exclusions": toCodequalityRequestIntegrationObject(o.Exclusions, isavro, false, "exclusions"),
		// ID the primary key for the model instance
		"id": toCodequalityRequestIntegrationObject(o.ID, isavro, false, "string"),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toCodequalityRequestIntegrationObject(o.Location, isavro, false, "location"),
		// Name The user friendly name of the integration
		"name": toCodequalityRequestIntegrationObject(o.Name, isavro, false, "string"),
		// Progress Agent processing progress
		"progress": toCodequalityRequestIntegrationObject(o.Progress, isavro, false, "progress"),
		// RefID the source system id for the model instance
		"ref_id": toCodequalityRequestIntegrationObject(o.RefID, isavro, false, "string"),
		// RefType the source system identifier for the model instance
		"ref_type": toCodequalityRequestIntegrationObject(o.RefType, isavro, false, "string"),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toCodequalityRequestIntegrationObject(o.SystemType, isavro, false, "system_type"),
		// Validated If the validation has been run against this instance
		"validated": toCodequalityRequestIntegrationObject(o.Validated, isavro, true, "boolean"),
		// ValidatedDate Date when validated
		"validated_date": toCodequalityRequestIntegrationObject(o.ValidatedDate, isavro, false, "validated_date"),
		// ValidationMessage The validation message from the agent
		"validation_message": toCodequalityRequestIntegrationObject(o.ValidationMessage, isavro, true, "string"),
	}
}

func (o *CodequalityRequestIntegration) setDefaults(frommap bool) {

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
func (o *CodequalityRequestIntegration) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(CodequalityRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationAuthorization); ok {
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

	if val, ok := kv["location"].(CodequalityRequestIntegrationLocation); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationProgress); ok {
			// struct
			o.Progress = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationProgress); ok {
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

	if val, ok := kv["system_type"].(CodequalityRequestIntegrationSystemType); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationValidatedDate); ok {
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

// CodequalityRequestRequestDate represents the object structure for request_date
type CodequalityRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestRequestDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestRequestDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequestRequestDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *CodequalityRequestRequestDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestRequestDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestRequestDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestRequestDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *CodequalityRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequest an agent action to request adding new codequality entities
type CodequalityRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration CodequalityRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate CodequalityRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityRequest)(nil)

func toCodequalityRequestObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityRequestObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityRequest:
		return v.ToMap(isavro)

	case CodequalityRequestIntegration:
		return v.ToMap(isavro)

	case CodequalityRequestRequestDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of CodequalityRequest
func (o *CodequalityRequest) String() string {
	return fmt.Sprintf("agent.CodequalityRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CodequalityRequest) GetTopicName() datamodel.TopicNameType {
	return CodequalityRequestTopic
}

// GetModelName returns the name of the model
func (o *CodequalityRequest) GetModelName() datamodel.ModelNameType {
	return CodequalityRequestModelName
}

func (o *CodequalityRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CodequalityRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CodequalityRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CodequalityRequest) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CodequalityRequest) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for CodequalityRequest")
}

// GetRefID returns the RefID for the object
func (o *CodequalityRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CodequalityRequest) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CodequalityRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CodequalityRequest) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CodequalityRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CodequalityRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CodequalityRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
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
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *CodequalityRequest) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *CodequalityRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CodequalityRequest
func (o *CodequalityRequest) Clone() datamodel.Model {
	c := new(CodequalityRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CodequalityRequest) Anon() datamodel.Model {
	c := new(CodequalityRequest)
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
func (o *CodequalityRequest) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *CodequalityRequest) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *CodequalityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CodequalityRequest) UnmarshalJSON(data []byte) error {
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

var cachedCodecCodequalityRequest *goavro.Codec
var cachedCodecCodequalityRequestLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *CodequalityRequest) GetAvroCodec() *goavro.Codec {
	cachedCodecCodequalityRequestLock.Lock()
	if cachedCodecCodequalityRequest == nil {
		c, err := GetCodequalityRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCodequalityRequest = c
	}
	cachedCodecCodequalityRequestLock.Unlock()
	return cachedCodecCodequalityRequest
}

// ToAvroBinary returns the data as Avro binary data
func (o *CodequalityRequest) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CodequalityRequest) FromAvroBinary(value []byte) error {
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
func (o *CodequalityRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CodequalityRequest objects are equal
func (o *CodequalityRequest) IsEqual(other *CodequalityRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityRequest) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toCodequalityRequestObject(o.CustomerID, isavro, false, "string"),
		"id":           toCodequalityRequestObject(o.ID, isavro, false, "string"),
		"integration":  toCodequalityRequestObject(o.Integration, isavro, false, "integration"),
		"ref_id":       toCodequalityRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toCodequalityRequestObject(o.RefType, isavro, false, "string"),
		"request_date": toCodequalityRequestObject(o.RequestDate, isavro, false, "request_date"),
		"updated_ts":   toCodequalityRequestObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":         toCodequalityRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toCodequalityRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(CodequalityRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*CodequalityRequestIntegration); ok {
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
		} else if sv, ok := val.(CodequalityRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*CodequalityRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.RequestDate.Epoch = dt.Epoch
				o.RequestDate.Rfc3339 = dt.Rfc3339
				o.RequestDate.Offset = dt.Offset
			}
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
func (o *CodequalityRequest) Hash() string {
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

// GetCodequalityRequestAvroSchemaSpec creates the avro schema specification for CodequalityRequest
func GetCodequalityRequestAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "CodequalityRequest",
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
				"type": map[string]interface{}{"doc": "the integration details to use", "fields": []interface{}{map[string]interface{}{"doc": "If true, the integration is still active", "name": "active", "type": "boolean"}, map[string]interface{}{"doc": "Authorization information", "name": "authorization", "type": map[string]interface{}{"doc": "Authorization information", "fields": []interface{}{map[string]interface{}{"default": nil, "doc": "Access token", "name": "access_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Refresh token", "name": "refresh_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "URL of instance if relevant", "name": "url", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Username for instance, if relevant", "name": "username", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Password for instance, if relevant", "name": "password", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Token for instance, if relevant", "name": "api_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Collection for instance, if relevant", "name": "collection", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Key for instance, if relevant", "name": "api_key", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "the agents encrypted authorization", "name": "authorization", "type": []interface{}{"null", "string"}}}, "name": "integration.authorization", "type": "record"}}, map[string]interface{}{"doc": "the customer id for the model instance", "name": "customer_id", "type": "string"}, map[string]interface{}{"default": nil, "doc": "If authorization failed by the agent", "name": "errored", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "The exclusion list for this integration", "name": "exclusions", "type": map[string]interface{}{"items": "string", "name": "exclusions", "type": "array"}}, map[string]interface{}{"doc": "the primary key for the model instance", "name": "id", "type": "string"}, map[string]interface{}{"doc": "The location of this integration (on-premise / private or cloud)", "name": "location", "type": map[string]interface{}{"default": "PRIVATE", "doc": "The location of this integration (on-premise / private or cloud)", "name": "integration.location", "symbols": []string{"PRIVATE", "CLOUD"}, "type": "enum"}}, map[string]interface{}{"doc": "The user friendly name of the integration", "name": "name", "type": "string"}, map[string]interface{}{"doc": "Agent processing progress", "name": "progress", "type": map[string]interface{}{"doc": "Agent processing progress", "fields": []interface{}{map[string]interface{}{"doc": "Any relevant messaging during processing", "name": "message", "type": "string"}, map[string]interface{}{"doc": "The total amount to be processed", "name": "total", "type": "long"}, map[string]interface{}{"doc": "The total amount processed thus far", "name": "completed", "type": "long"}}, "name": "integration.progress", "type": "record"}}, map[string]interface{}{"doc": "the source system id for the model instance", "name": "ref_id", "type": "string"}, map[string]interface{}{"doc": "the source system identifier for the model instance", "name": "ref_type", "type": "string"}, map[string]interface{}{"doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "system_type", "type": map[string]interface{}{"default": "WORK", "doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "integration.system_type", "symbols": []string{"WORK", "SOURCECODE", "CODEQUALITY"}, "type": "enum"}}, map[string]interface{}{"default": nil, "doc": "If the validation has been run against this instance", "name": "validated", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "Date when validated", "name": "validated_date", "type": map[string]interface{}{"doc": "Date when validated", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "integration.validated_date", "type": "record"}}, map[string]interface{}{"default": nil, "doc": "The validation message from the agent", "name": "validation_message", "type": []interface{}{"null", "string"}}}, "name": "integration", "type": "record"},
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
func (o *CodequalityRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetCodequalityRequestAvroSchema creates the avro schema for CodequalityRequest
func GetCodequalityRequestAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCodequalityRequestAvroSchemaSpec())
}

// TransformCodequalityRequestFunc is a function for transforming CodequalityRequest during processing
type TransformCodequalityRequestFunc func(input *CodequalityRequest) (*CodequalityRequest, error)

// NewCodequalityRequestPipe creates a pipe for processing CodequalityRequest items
func NewCodequalityRequestPipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCodequalityRequestFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCodequalityRequestInputStream(input, errors)
	var stream chan CodequalityRequest
	if len(transforms) > 0 {
		stream = make(chan CodequalityRequest, 1000)
	} else {
		stream = inch
	}
	outdone := NewCodequalityRequestOutputStream(output, stream, errors)
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

// NewCodequalityRequestInputStreamDir creates a channel for reading CodequalityRequest as JSON newlines from a directory of files
func NewCodequalityRequestInputStreamDir(dir string, errors chan<- error, transforms ...TransformCodequalityRequestFunc) (chan CodequalityRequest, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/codequality_request\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CodequalityRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for codequality_request")
		ch := make(chan CodequalityRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCodequalityRequestInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CodequalityRequest)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCodequalityRequestInputStreamFile creates an channel for reading CodequalityRequest as JSON newlines from filename
func NewCodequalityRequestInputStreamFile(filename string, errors chan<- error, transforms ...TransformCodequalityRequestFunc) (chan CodequalityRequest, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CodequalityRequest)
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
			ch := make(chan CodequalityRequest)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCodequalityRequestInputStream(f, errors, transforms...)
}

// NewCodequalityRequestInputStream creates an channel for reading CodequalityRequest as JSON newlines from stream
func NewCodequalityRequestInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCodequalityRequestFunc) (chan CodequalityRequest, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CodequalityRequest, 1000)
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
			var item CodequalityRequest
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

// NewCodequalityRequestOutputStreamDir will output json newlines from channel and save in dir
func NewCodequalityRequestOutputStreamDir(dir string, ch chan CodequalityRequest, errors chan<- error, transforms ...TransformCodequalityRequestFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/codequality_request\\.json(\\.gz)?$")
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
	return NewCodequalityRequestOutputStream(gz, ch, errors, transforms...)
}

// NewCodequalityRequestOutputStream will output json newlines from channel to the stream
func NewCodequalityRequestOutputStream(stream io.WriteCloser, ch chan CodequalityRequest, errors chan<- error, transforms ...TransformCodequalityRequestFunc) <-chan bool {
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

// CodequalityRequestSendEvent is an event detail for sending data
type CodequalityRequestSendEvent struct {
	CodequalityRequest *CodequalityRequest
	headers            map[string]string
	time               time.Time
	key                string
}

var _ datamodel.ModelSendEvent = (*CodequalityRequestSendEvent)(nil)

// Key is the key to use for the message
func (e *CodequalityRequestSendEvent) Key() string {
	if e.key == "" {
		return e.CodequalityRequest.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CodequalityRequestSendEvent) Object() datamodel.Model {
	return e.CodequalityRequest
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CodequalityRequestSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CodequalityRequestSendEvent) Timestamp() time.Time {
	return e.time
}

// CodequalityRequestSendEventOpts is a function handler for setting opts
type CodequalityRequestSendEventOpts func(o *CodequalityRequestSendEvent)

// WithCodequalityRequestSendEventKey sets the key value to a value different than the object ID
func WithCodequalityRequestSendEventKey(key string) CodequalityRequestSendEventOpts {
	return func(o *CodequalityRequestSendEvent) {
		o.key = key
	}
}

// WithCodequalityRequestSendEventTimestamp sets the timestamp value
func WithCodequalityRequestSendEventTimestamp(tv time.Time) CodequalityRequestSendEventOpts {
	return func(o *CodequalityRequestSendEvent) {
		o.time = tv
	}
}

// WithCodequalityRequestSendEventHeader sets the timestamp value
func WithCodequalityRequestSendEventHeader(key, value string) CodequalityRequestSendEventOpts {
	return func(o *CodequalityRequestSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCodequalityRequestSendEvent returns a new CodequalityRequestSendEvent instance
func NewCodequalityRequestSendEvent(o *CodequalityRequest, opts ...CodequalityRequestSendEventOpts) *CodequalityRequestSendEvent {
	res := &CodequalityRequestSendEvent{
		CodequalityRequest: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCodequalityRequestProducer will stream data from the channel
func NewCodequalityRequestProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
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
				if object, ok := item.Object().(*CodequalityRequest); ok {
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
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.CodequalityRequest but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewCodequalityRequestConsumer will stream data from the topic into the provided channel
func NewCodequalityRequestConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object CodequalityRequest
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.CodequalityRequest: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.CodequalityRequest: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.CodequalityRequest")
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

			ch <- &CodequalityRequestReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object CodequalityRequest
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CodequalityRequestReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// CodequalityRequestReceiveEvent is an event detail for receiving data
type CodequalityRequestReceiveEvent struct {
	CodequalityRequest *CodequalityRequest
	message            eventing.Message
	eof                bool
}

var _ datamodel.ModelReceiveEvent = (*CodequalityRequestReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CodequalityRequestReceiveEvent) Object() datamodel.Model {
	return e.CodequalityRequest
}

// Message returns the underlying message data for the event
func (e *CodequalityRequestReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CodequalityRequestReceiveEvent) EOF() bool {
	return e.eof
}

// CodequalityRequestProducer implements the datamodel.ModelEventProducer
type CodequalityRequestProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*CodequalityRequestProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CodequalityRequestProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CodequalityRequestProducer) Close() error {
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
func (o *CodequalityRequest) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *CodequalityRequest) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// NewCodequalityRequestProducerChannel returns a channel which can be used for producing Model events
func NewCodequalityRequestProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewCodequalityRequestProducerChannelSize(producer, 0, errors)
}

// NewCodequalityRequestProducerChannelSize returns a channel which can be used for producing Model events
func NewCodequalityRequestProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityRequestProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityRequestProducer(newctx, producer, ch, errors, empty),
	}
}

// CodequalityRequestConsumer implements the datamodel.ModelEventConsumer
type CodequalityRequestConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*CodequalityRequestConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CodequalityRequestConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CodequalityRequestConsumer) Close() error {
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
func (o *CodequalityRequest) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityRequestConsumer{
		ch:       ch,
		callback: NewCodequalityRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewCodequalityRequestConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCodequalityRequestConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityRequestConsumer{
		ch:       ch,
		callback: NewCodequalityRequestConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
