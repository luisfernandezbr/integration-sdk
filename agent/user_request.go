// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// UserRequestTopic is the default topic name
	UserRequestTopic datamodel.TopicNameType = "agent_UserRequest_topic"

	// UserRequestStream is the default stream name
	UserRequestStream datamodel.TopicNameType = "agent_UserRequest_stream"

	// UserRequestTable is the default table name
	UserRequestTable datamodel.TopicNameType = "agent_userrequest"

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
	// UserRequestIntegrationColumnCustomerIDColumn is the customer_id column property of the Integration name
	UserRequestIntegrationColumnCustomerIDColumn = "integration->customer_id"
	// UserRequestIntegrationColumnErroredColumn is the errored column property of the Integration name
	UserRequestIntegrationColumnErroredColumn = "integration->errored"
	// UserRequestIntegrationColumnExclusionsColumn is the exclusions column property of the Integration name
	UserRequestIntegrationColumnExclusionsColumn = "integration->exclusions"
	// UserRequestIntegrationColumnIDColumn is the id column property of the Integration name
	UserRequestIntegrationColumnIDColumn = "integration->id"
	// UserRequestIntegrationColumnLocationColumn is the location column property of the Integration name
	UserRequestIntegrationColumnLocationColumn = "integration->location"
	// UserRequestIntegrationColumnNameColumn is the name column property of the Integration name
	UserRequestIntegrationColumnNameColumn = "integration->name"
	// UserRequestIntegrationColumnProgressColumn is the progress column property of the Integration name
	UserRequestIntegrationColumnProgressColumn = "integration->progress"
	// UserRequestIntegrationColumnRefIDColumn is the ref_id column property of the Integration name
	UserRequestIntegrationColumnRefIDColumn = "integration->ref_id"
	// UserRequestIntegrationColumnRefTypeColumn is the ref_type column property of the Integration name
	UserRequestIntegrationColumnRefTypeColumn = "integration->ref_type"
	// UserRequestIntegrationColumnSystemTypeColumn is the system_type column property of the Integration name
	UserRequestIntegrationColumnSystemTypeColumn = "integration->system_type"
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
	// UserRequestSystemTypeColumn is the system_type column name
	UserRequestSystemTypeColumn = "system_type"
	// UserRequestUpdatedAtColumn is the updated_ts column name
	UserRequestUpdatedAtColumn = "updated_ts"
	// UserRequestUUIDColumn is the uuid column name
	UserRequestUUIDColumn = "uuid"
)

// UserRequestIntegrationAuthorization represents the object structure for authorization
type UserRequestIntegrationAuthorization struct {
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
	// Hostname Hostname for instance, if relevant
	Hostname *string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
}

func toUserRequestIntegrationAuthorizationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestIntegrationAuthorizationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *UserRequestIntegrationAuthorization) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toUserRequestIntegrationAuthorizationObject(o.AccessToken, isavro, true, "string"),
		// RefreshToken Refresh token
		"refresh_token": toUserRequestIntegrationAuthorizationObject(o.RefreshToken, isavro, true, "string"),
		// URL URL of instance if relevant
		"url": toUserRequestIntegrationAuthorizationObject(o.URL, isavro, true, "string"),
		// Username Username for instance, if relevant
		"username": toUserRequestIntegrationAuthorizationObject(o.Username, isavro, true, "string"),
		// Password Password for instance, if relevant
		"password": toUserRequestIntegrationAuthorizationObject(o.Password, isavro, true, "string"),
		// APIToken API Token for instance, if relevant
		"api_token": toUserRequestIntegrationAuthorizationObject(o.APIToken, isavro, true, "string"),
		// Collection Collection for instance, if relevant
		"collection": toUserRequestIntegrationAuthorizationObject(o.Collection, isavro, true, "string"),
		// APIKey API Key for instance, if relevant
		"api_key": toUserRequestIntegrationAuthorizationObject(o.APIKey, isavro, true, "string"),
		// Authorization the agents encrypted authorization
		"authorization": toUserRequestIntegrationAuthorizationObject(o.Authorization, isavro, true, "string"),
		// Hostname Hostname for instance, if relevant
		"hostname": toUserRequestIntegrationAuthorizationObject(o.Hostname, isavro, true, "string"),
	}
}

func (o *UserRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["hostname"].(*string); ok {
		o.Hostname = val
	} else if val, ok := kv["hostname"].(string); ok {
		o.Hostname = &val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Hostname = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// UserRequestIntegrationLocation is the enumeration type for location
type UserRequestIntegrationLocation int32

// String returns the string value for IntegrationLocation
func (v UserRequestIntegrationLocation) String() string {
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
	UserRequestIntegrationLocationPrivate UserRequestIntegrationLocation = 0
	// IntegrationLocationCloud is the enumeration value for cloud
	UserRequestIntegrationLocationCloud UserRequestIntegrationLocation = 1
)

// UserRequestIntegrationProgress represents the object structure for progress
type UserRequestIntegrationProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func toUserRequestIntegrationProgressObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestIntegrationProgressObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequestIntegrationProgress:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *UserRequestIntegrationProgress) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": toUserRequestIntegrationProgressObject(o.Message, isavro, false, "string"),
		// Total The total amount to be processed
		"total": toUserRequestIntegrationProgressObject(o.Total, isavro, false, "long"),
		// Completed The total amount processed thus far
		"completed": toUserRequestIntegrationProgressObject(o.Completed, isavro, false, "long"),
	}
}

func (o *UserRequestIntegrationProgress) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequestIntegrationProgress) FromMap(kv map[string]interface{}) {

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

// UserRequestIntegrationSystemType is the enumeration type for system_type
type UserRequestIntegrationSystemType int32

// String returns the string value for IntegrationSystemType
func (v UserRequestIntegrationSystemType) String() string {
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
	UserRequestIntegrationSystemTypeWork UserRequestIntegrationSystemType = 0
	// IntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	UserRequestIntegrationSystemTypeSourcecode UserRequestIntegrationSystemType = 1
	// IntegrationSystemTypeCodequality is the enumeration value for codequality
	UserRequestIntegrationSystemTypeCodequality UserRequestIntegrationSystemType = 2
)

// UserRequestIntegrationValidatedDate represents the object structure for validated_date
type UserRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserRequestIntegrationValidatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestIntegrationValidatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *UserRequestIntegrationValidatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserRequestIntegrationValidatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUserRequestIntegrationValidatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserRequestIntegrationValidatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *UserRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

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

// UserRequestIntegration represents the object structure for integration
type UserRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization UserRequestIntegrationAuthorization `json:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location UserRequestIntegrationLocation `json:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress UserRequestIntegrationProgress `json:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType UserRequestIntegrationSystemType `json:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated" bson:"validated" yaml:"validated" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate UserRequestIntegrationValidatedDate `json:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message" bson:"validation_message" yaml:"validation_message" faker:"-"`
}

func toUserRequestIntegrationObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestIntegrationObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequestIntegration:
		return v.ToMap(isavro)

	case UserRequestIntegrationAuthorization:
		return v.ToMap(isavro)

	case UserRequestIntegrationLocation:
		return v.String()

	case UserRequestIntegrationProgress:
		return v.ToMap(isavro)

	case UserRequestIntegrationSystemType:
		return v.String()

	case UserRequestIntegrationValidatedDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *UserRequestIntegration) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toUserRequestIntegrationObject(o.Active, isavro, false, "boolean"),
		// Authorization Authorization information
		"authorization": toUserRequestIntegrationObject(o.Authorization, isavro, false, "authorization"),
		// CustomerID the customer id for the model instance
		"customer_id": toUserRequestIntegrationObject(o.CustomerID, isavro, false, "string"),
		// Errored If authorization failed by the agent
		"errored": toUserRequestIntegrationObject(o.Errored, isavro, true, "boolean"),
		// Exclusions The exclusion list for this integration
		"exclusions": toUserRequestIntegrationObject(o.Exclusions, isavro, false, "exclusions"),
		// ID the primary key for the model instance
		"id": toUserRequestIntegrationObject(o.ID, isavro, false, "string"),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toUserRequestIntegrationObject(o.Location, isavro, false, "location"),
		// Name The user friendly name of the integration
		"name": toUserRequestIntegrationObject(o.Name, isavro, false, "string"),
		// Progress Agent processing progress
		"progress": toUserRequestIntegrationObject(o.Progress, isavro, false, "progress"),
		// RefID the source system id for the model instance
		"ref_id": toUserRequestIntegrationObject(o.RefID, isavro, false, "string"),
		// RefType the source system identifier for the model instance
		"ref_type": toUserRequestIntegrationObject(o.RefType, isavro, false, "string"),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toUserRequestIntegrationObject(o.SystemType, isavro, false, "system_type"),
		// Validated If the validation has been run against this instance
		"validated": toUserRequestIntegrationObject(o.Validated, isavro, true, "boolean"),
		// ValidatedDate Date when validated
		"validated_date": toUserRequestIntegrationObject(o.ValidatedDate, isavro, false, "validated_date"),
		// ValidationMessage The validation message from the agent
		"validation_message": toUserRequestIntegrationObject(o.ValidationMessage, isavro, true, "string"),
	}
}

func (o *UserRequestIntegration) setDefaults(frommap bool) {

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
func (o *UserRequestIntegration) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(UserRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*UserRequestIntegrationAuthorization); ok {
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

	if val, ok := kv["location"].(UserRequestIntegrationLocation); ok {
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
		} else if sv, ok := val.(UserRequestIntegrationProgress); ok {
			// struct
			o.Progress = sv
		} else if sp, ok := val.(*UserRequestIntegrationProgress); ok {
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

	if val, ok := kv["system_type"].(UserRequestIntegrationSystemType); ok {
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
		} else if sv, ok := val.(UserRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*UserRequestIntegrationValidatedDate); ok {
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

// UserRequestRequestDate represents the object structure for request_date
type UserRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toUserRequestRequestDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserRequestRequestDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequestRequestDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *UserRequestRequestDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toUserRequestRequestDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toUserRequestRequestDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toUserRequestRequestDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *UserRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// UserRequest an agent action to request adding new users
type UserRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration UserRequestIntegration `json:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate UserRequestRequestDate `json:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *UserRequest:
		return v.ToMap(isavro)

	case UserRequestIntegration:
		return v.ToMap(isavro)

	case UserRequestRequestDate:
		return v.ToMap(isavro)

	default:
		return o
	}
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

// GetStreamName returns the name of the stream
func (o *UserRequest) GetStreamName() string {
	return UserRequestStream.String()
}

// GetTableName returns the name of the table
func (o *UserRequest) GetTableName() string {
	return UserRequestTable.String()
}

// NewUserRequestID provides a template for generating an ID field for UserRequest
func NewUserRequestID(customerID string, refType string, refID string) string {
	return hash.Values("UserRequest", customerID, refType, refID)
}

func (o *UserRequest) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserRequest) GetID() string {
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *UserRequest) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *UserRequest) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

var cachedCodecUserRequest *goavro.Codec
var cachedCodecUserRequestLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *UserRequest) GetAvroCodec() *goavro.Codec {
	cachedCodecUserRequestLock.Lock()
	if cachedCodecUserRequest == nil {
		c, err := GetUserRequestAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUserRequest = c
	}
	cachedCodecUserRequestLock.Unlock()
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
	o.Hash()
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
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toUserRequestObject(o.CustomerID, isavro, false, "string"),
		"id":           toUserRequestObject(o.ID, isavro, false, "string"),
		"integration":  toUserRequestObject(o.Integration, isavro, false, "integration"),
		"ref_id":       toUserRequestObject(o.RefID, isavro, false, "string"),
		"ref_type":     toUserRequestObject(o.RefType, isavro, false, "string"),
		"request_date": toUserRequestObject(o.RequestDate, isavro, false, "request_date"),
		"updated_ts":   toUserRequestObject(o.UpdatedAt, isavro, false, "long"),
		"uuid":         toUserRequestObject(o.UUID, isavro, false, "string"),
		"hashcode":     toUserRequestObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(UserRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*UserRequestIntegration); ok {
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
		} else if sv, ok := val.(UserRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*UserRequestRequestDate); ok {
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
func (o *UserRequest) Hash() string {
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
				"type": map[string]interface{}{"doc": "the integration details to use", "fields": []interface{}{map[string]interface{}{"doc": "If true, the integration is still active", "name": "active", "type": "boolean"}, map[string]interface{}{"doc": "Authorization information", "name": "authorization", "type": map[string]interface{}{"doc": "Authorization information", "fields": []interface{}{map[string]interface{}{"default": nil, "doc": "Access token", "name": "access_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Refresh token", "name": "refresh_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "URL of instance if relevant", "name": "url", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Username for instance, if relevant", "name": "username", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Password for instance, if relevant", "name": "password", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Token for instance, if relevant", "name": "api_token", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Collection for instance, if relevant", "name": "collection", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "API Key for instance, if relevant", "name": "api_key", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "the agents encrypted authorization", "name": "authorization", "type": []interface{}{"null", "string"}}, map[string]interface{}{"default": nil, "doc": "Hostname for instance, if relevant", "name": "hostname", "type": []interface{}{"null", "string"}}}, "name": "integration.authorization", "type": "record"}}, map[string]interface{}{"doc": "the customer id for the model instance", "name": "customer_id", "type": "string"}, map[string]interface{}{"default": nil, "doc": "If authorization failed by the agent", "name": "errored", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "The exclusion list for this integration", "name": "exclusions", "type": map[string]interface{}{"items": "string", "name": "exclusions", "type": "array"}}, map[string]interface{}{"doc": "the primary key for the model instance", "name": "id", "type": "string"}, map[string]interface{}{"doc": "The location of this integration (on-premise / private or cloud)", "name": "location", "type": map[string]interface{}{"default": "PRIVATE", "doc": "The location of this integration (on-premise / private or cloud)", "name": "integration.location", "symbols": []string{"PRIVATE", "CLOUD"}, "type": "enum"}}, map[string]interface{}{"doc": "The user friendly name of the integration", "name": "name", "type": "string"}, map[string]interface{}{"doc": "Agent processing progress", "name": "progress", "type": map[string]interface{}{"doc": "Agent processing progress", "fields": []interface{}{map[string]interface{}{"doc": "Any relevant messaging during processing", "name": "message", "type": "string"}, map[string]interface{}{"doc": "The total amount to be processed", "name": "total", "type": "long"}, map[string]interface{}{"doc": "The total amount processed thus far", "name": "completed", "type": "long"}}, "name": "integration.progress", "type": "record"}}, map[string]interface{}{"doc": "the source system id for the model instance", "name": "ref_id", "type": "string"}, map[string]interface{}{"doc": "the source system identifier for the model instance", "name": "ref_type", "type": "string"}, map[string]interface{}{"doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "system_type", "type": map[string]interface{}{"default": "WORK", "doc": "The system type of the integration (sourcecode / work (jira) / codequality / etc.)", "name": "integration.system_type", "symbols": []string{"WORK", "SOURCECODE", "CODEQUALITY"}, "type": "enum"}}, map[string]interface{}{"default": nil, "doc": "If the validation has been run against this instance", "name": "validated", "type": []interface{}{"null", "boolean"}}, map[string]interface{}{"doc": "Date when validated", "name": "validated_date", "type": map[string]interface{}{"doc": "Date when validated", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "integration.validated_date", "type": "record"}}, map[string]interface{}{"default": nil, "doc": "The validation message from the agent", "name": "validation_message", "type": []interface{}{"null", "string"}}}, "name": "integration", "type": "record"},
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
	emptyTime := time.Unix(0, 0)
	var numPartitions int
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
					if numPartitions == 0 {
						numPartitions = object.GetTopicConfig().NumPartitions
					}
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
					// determine the partition selection by using the partition key
					// and taking the modulo over the number of partitions for the topic
					partition := hash.Modulo(item.Key(), numPartitions)
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       object.GetID(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: int32(partition),
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
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
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
