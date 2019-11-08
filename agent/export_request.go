// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ExportRequestTopic is the default topic name
	ExportRequestTopic datamodel.TopicNameType = "agent_ExportRequest_topic"

	// ExportRequestModelName is the model name
	ExportRequestModelName datamodel.ModelNameType = "agent.ExportRequest"
)

const (
	// ExportRequestCustomerIDColumn is the customer_id column name
	ExportRequestCustomerIDColumn = "CustomerID"
	// ExportRequestIDColumn is the id column name
	ExportRequestIDColumn = "ID"
	// ExportRequestIntegrationsColumn is the integrations column name
	ExportRequestIntegrationsColumn = "Integrations"
	// ExportRequestIntegrationsColumnActiveColumn is the active column property of the Integrations name
	ExportRequestIntegrationsColumnActiveColumn = "Integrations.Active"
	// ExportRequestIntegrationsColumnAuthorizationColumn is the authorization column property of the Integrations name
	ExportRequestIntegrationsColumnAuthorizationColumn = "Integrations.Authorization"
	// ExportRequestIntegrationsColumnCustomerIDColumn is the customer_id column property of the Integrations name
	ExportRequestIntegrationsColumnCustomerIDColumn = "Integrations.CustomerID"
	// ExportRequestIntegrationsColumnErroredColumn is the errored column property of the Integrations name
	ExportRequestIntegrationsColumnErroredColumn = "Integrations.Errored"
	// ExportRequestIntegrationsColumnExclusionsColumn is the exclusions column property of the Integrations name
	ExportRequestIntegrationsColumnExclusionsColumn = "Integrations.Exclusions"
	// ExportRequestIntegrationsColumnIDColumn is the id column property of the Integrations name
	ExportRequestIntegrationsColumnIDColumn = "Integrations.ID"
	// ExportRequestIntegrationsColumnLocationColumn is the location column property of the Integrations name
	ExportRequestIntegrationsColumnLocationColumn = "Integrations.Location"
	// ExportRequestIntegrationsColumnNameColumn is the name column property of the Integrations name
	ExportRequestIntegrationsColumnNameColumn = "Integrations.Name"
	// ExportRequestIntegrationsColumnProgressColumn is the progress column property of the Integrations name
	ExportRequestIntegrationsColumnProgressColumn = "Integrations.Progress"
	// ExportRequestIntegrationsColumnRefIDColumn is the ref_id column property of the Integrations name
	ExportRequestIntegrationsColumnRefIDColumn = "Integrations.RefID"
	// ExportRequestIntegrationsColumnRefTypeColumn is the ref_type column property of the Integrations name
	ExportRequestIntegrationsColumnRefTypeColumn = "Integrations.RefType"
	// ExportRequestIntegrationsColumnSystemTypeColumn is the system_type column property of the Integrations name
	ExportRequestIntegrationsColumnSystemTypeColumn = "Integrations.SystemType"
	// ExportRequestIntegrationsColumnValidatedColumn is the validated column property of the Integrations name
	ExportRequestIntegrationsColumnValidatedColumn = "Integrations.Validated"
	// ExportRequestIntegrationsColumnValidatedDateColumn is the validated_date column property of the Integrations name
	ExportRequestIntegrationsColumnValidatedDateColumn = "Integrations.ValidatedDate"
	// ExportRequestIntegrationsColumnValidationMessageColumn is the validation_message column property of the Integrations name
	ExportRequestIntegrationsColumnValidationMessageColumn = "Integrations.ValidationMessage"
	// ExportRequestJobIDColumn is the job_id column name
	ExportRequestJobIDColumn = "JobID"
	// ExportRequestLocationColumn is the location column name
	ExportRequestLocationColumn = "Location"
	// ExportRequestRefIDColumn is the ref_id column name
	ExportRequestRefIDColumn = "RefID"
	// ExportRequestRefTypeColumn is the ref_type column name
	ExportRequestRefTypeColumn = "RefType"
	// ExportRequestReprocessHistoricalColumn is the reprocess_historical column name
	ExportRequestReprocessHistoricalColumn = "ReprocessHistorical"
	// ExportRequestRequestDateColumn is the request_date column name
	ExportRequestRequestDateColumn = "RequestDate"
	// ExportRequestRequestDateColumnEpochColumn is the epoch column property of the RequestDate name
	ExportRequestRequestDateColumnEpochColumn = "RequestDate.Epoch"
	// ExportRequestRequestDateColumnOffsetColumn is the offset column property of the RequestDate name
	ExportRequestRequestDateColumnOffsetColumn = "RequestDate.Offset"
	// ExportRequestRequestDateColumnRfc3339Column is the rfc3339 column property of the RequestDate name
	ExportRequestRequestDateColumnRfc3339Column = "RequestDate.Rfc3339"
	// ExportRequestSystemTypeColumn is the system_type column name
	ExportRequestSystemTypeColumn = "SystemType"
	// ExportRequestUpdatedAtColumn is the updated_ts column name
	ExportRequestUpdatedAtColumn = "UpdatedAt"
	// ExportRequestUploadHeadersColumn is the upload_headers column name
	ExportRequestUploadHeadersColumn = "UploadHeaders"
	// ExportRequestUploadURLColumn is the upload_url column name
	ExportRequestUploadURLColumn = "UploadURL"
	// ExportRequestUUIDColumn is the uuid column name
	ExportRequestUUIDColumn = "UUID"
)

// ExportRequestIntegrationsAuthorization represents the object structure for authorization
type ExportRequestIntegrationsAuthorization struct {
	// AccessToken Access token
	AccessToken *string `json:"access_token,omitempty" codec:"access_token,omitempty" bson:"access_token" yaml:"access_token,omitempty" faker:"-"`
	// RefreshToken Refresh token
	RefreshToken *string `json:"refresh_token,omitempty" codec:"refresh_token,omitempty" bson:"refresh_token" yaml:"refresh_token,omitempty" faker:"-"`
	// URL URL of instance if relevant
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Username Username for instance, if relevant
	Username *string `json:"username,omitempty" codec:"username,omitempty" bson:"username" yaml:"username,omitempty" faker:"-"`
	// Password Password for instance, if relevant
	Password *string `json:"password,omitempty" codec:"password,omitempty" bson:"password" yaml:"password,omitempty" faker:"-"`
	// APIToken API Token for instance, if relevant
	APIToken *string `json:"api_token,omitempty" codec:"api_token,omitempty" bson:"api_token" yaml:"api_token,omitempty" faker:"-"`
	// Collection Collection for instance, if relevant
	Collection *string `json:"collection,omitempty" codec:"collection,omitempty" bson:"collection" yaml:"collection,omitempty" faker:"-"`
	// APIKey API Key for instance, if relevant
	APIKey *string `json:"api_key,omitempty" codec:"api_key,omitempty" bson:"api_key" yaml:"api_key,omitempty" faker:"-"`
	// Authorization the agents encrypted authorization
	Authorization *string `json:"authorization,omitempty" codec:"authorization,omitempty" bson:"authorization" yaml:"authorization,omitempty" faker:"-"`
	// Hostname Hostname for instance, if relevant
	Hostname *string `json:"hostname,omitempty" codec:"hostname,omitempty" bson:"hostname" yaml:"hostname,omitempty" faker:"-"`
	// APIVersion the api version of the integration
	APIVersion *string `json:"api_version,omitempty" codec:"api_version,omitempty" bson:"api_version" yaml:"api_version,omitempty" faker:"-"`
	// Organization Organization for instance, if relevant
	Organization *string `json:"organization,omitempty" codec:"organization,omitempty" bson:"organization" yaml:"organization,omitempty" faker:"-"`
}

func toExportRequestIntegrationsAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ExportRequestIntegrationsAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toExportRequestIntegrationsAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toExportRequestIntegrationsAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toExportRequestIntegrationsAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toExportRequestIntegrationsAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toExportRequestIntegrationsAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toExportRequestIntegrationsAuthorizationObject(o.APIToken, true),
		// Collection Collection for instance, if relevant
		"collection": toExportRequestIntegrationsAuthorizationObject(o.Collection, true),
		// APIKey API Key for instance, if relevant
		"api_key": toExportRequestIntegrationsAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toExportRequestIntegrationsAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toExportRequestIntegrationsAuthorizationObject(o.Hostname, true),
		// APIVersion the api version of the integration
		"api_version": toExportRequestIntegrationsAuthorizationObject(o.APIVersion, true),
		// Organization Organization for instance, if relevant
		"organization": toExportRequestIntegrationsAuthorizationObject(o.Organization, true),
	}
}

func (o *ExportRequestIntegrationsAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsAuthorization) FromMap(kv map[string]interface{}) {

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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
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
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Hostname = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["api_version"].(*string); ok {
		o.APIVersion = val
	} else if val, ok := kv["api_version"].(string); ok {
		o.APIVersion = &val
	} else {
		if val, ok := kv["api_version"]; ok {
			if val == nil {
				o.APIVersion = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.APIVersion = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["organization"].(*string); ok {
		o.Organization = val
	} else if val, ok := kv["organization"].(string); ok {
		o.Organization = &val
	} else {
		if val, ok := kv["organization"]; ok {
			if val == nil {
				o.Organization = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Organization = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLocation is the enumeration type for location
type ExportRequestIntegrationsLocation int32

// UnmarshalBSON unmarshals the enum value
func (v ExportRequestIntegrationsLocation) UnmarshalBSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalBSON marshals the enum value
func (v ExportRequestIntegrationsLocation) MarshalBSON() ([]byte, error) {
	switch v {
	case 0:
		return []byte("PRIVATE"), nil
	case 1:
		return []byte("CLOUD"), nil
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationsLocation
func (v ExportRequestIntegrationsLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

const (
	// IntegrationsLocationPrivate is the enumeration value for private
	ExportRequestIntegrationsLocationPrivate ExportRequestIntegrationsLocation = 0
	// IntegrationsLocationCloud is the enumeration value for cloud
	ExportRequestIntegrationsLocationCloud ExportRequestIntegrationsLocation = 1
)

// ExportRequestIntegrationsProgress represents the object structure for progress
type ExportRequestIntegrationsProgress struct {
	// Message Any relevant messaging during processing
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// Total The total amount to be processed
	Total int64 `json:"total" codec:"total" bson:"total" yaml:"total" faker:"-"`
	// Completed The total amount processed thus far
	Completed int64 `json:"completed" codec:"completed" bson:"completed" yaml:"completed" faker:"-"`
}

func toExportRequestIntegrationsProgressObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsProgress:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ExportRequestIntegrationsProgress) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Message Any relevant messaging during processing
		"message": toExportRequestIntegrationsProgressObject(o.Message, false),
		// Total The total amount to be processed
		"total": toExportRequestIntegrationsProgressObject(o.Total, false),
		// Completed The total amount processed thus far
		"completed": toExportRequestIntegrationsProgressObject(o.Completed, false),
	}
}

func (o *ExportRequestIntegrationsProgress) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsProgress) FromMap(kv map[string]interface{}) {

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

// ExportRequestIntegrationsSystemType is the enumeration type for system_type
type ExportRequestIntegrationsSystemType int32

// UnmarshalBSON unmarshals the enum value
func (v ExportRequestIntegrationsSystemType) UnmarshalBSON(buf []byte) error {
	switch string(buf) {
	case "WORK":
		v = 0
	case "SOURCECODE":
		v = 1
	case "CODEQUALITY":
		v = 2
	case "USER":
		v = 3
	}
	return nil
}

// MarshalBSON marshals the enum value
func (v ExportRequestIntegrationsSystemType) MarshalBSON() ([]byte, error) {
	switch v {
	case 0:
		return []byte("WORK"), nil
	case 1:
		return []byte("SOURCECODE"), nil
	case 2:
		return []byte("CODEQUALITY"), nil
	case 3:
		return []byte("USER"), nil
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationsSystemType
func (v ExportRequestIntegrationsSystemType) String() string {
	switch int32(v) {
	case 0:
		return "WORK"
	case 1:
		return "SOURCECODE"
	case 2:
		return "CODEQUALITY"
	case 3:
		return "USER"
	}
	return "unset"
}

const (
	// IntegrationsSystemTypeWork is the enumeration value for work
	ExportRequestIntegrationsSystemTypeWork ExportRequestIntegrationsSystemType = 0
	// IntegrationsSystemTypeSourcecode is the enumeration value for sourcecode
	ExportRequestIntegrationsSystemTypeSourcecode ExportRequestIntegrationsSystemType = 1
	// IntegrationsSystemTypeCodequality is the enumeration value for codequality
	ExportRequestIntegrationsSystemTypeCodequality ExportRequestIntegrationsSystemType = 2
	// IntegrationsSystemTypeUser is the enumeration value for user
	ExportRequestIntegrationsSystemTypeUser ExportRequestIntegrationsSystemType = 3
)

// ExportRequestIntegrationsValidatedDate represents the object structure for validated_date
type ExportRequestIntegrationsValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ExportRequestIntegrationsValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsValidatedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsValidatedDate) FromMap(kv map[string]interface{}) {

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

// ExportRequestIntegrations represents the object structure for integrations
type ExportRequestIntegrations struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization ExportRequestIntegrationsAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" codec:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location ExportRequestIntegrationsLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Progress Agent processing progress
	Progress ExportRequestIntegrationsProgress `json:"progress" codec:"progress" bson:"progress" yaml:"progress" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType ExportRequestIntegrationsSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate ExportRequestIntegrationsValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toExportRequestIntegrationsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrations:
		return v.ToMap()

	case ExportRequestIntegrationsAuthorization:
		return v.ToMap()

	case ExportRequestIntegrationsLocation:
		return v.String()

	case ExportRequestIntegrationsProgress:
		return v.ToMap()

	case ExportRequestIntegrationsSystemType:
		return v.String()

	case ExportRequestIntegrationsValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ExportRequestIntegrations) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toExportRequestIntegrationsObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toExportRequestIntegrationsObject(o.Authorization, false),
		// CustomerID the customer id for the model instance
		"customer_id": toExportRequestIntegrationsObject(o.CustomerID, false),
		// Errored If authorization failed by the agent
		"errored": toExportRequestIntegrationsObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toExportRequestIntegrationsObject(o.Exclusions, false),
		// ID the primary key for the model instance
		"id": toExportRequestIntegrationsObject(o.ID, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toExportRequestIntegrationsObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toExportRequestIntegrationsObject(o.Name, false),
		// Progress Agent processing progress
		"progress": toExportRequestIntegrationsObject(o.Progress, false),
		// RefID the source system id for the model instance
		"ref_id": toExportRequestIntegrationsObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toExportRequestIntegrationsObject(o.RefType, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toExportRequestIntegrationsObject(o.SystemType, false),
		// Validated If the validation has been run against this instance
		"validated": toExportRequestIntegrationsObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toExportRequestIntegrationsObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toExportRequestIntegrationsObject(o.ValidationMessage, true),
	}
}

func (o *ExportRequestIntegrations) setDefaults(frommap bool) {

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
func (o *ExportRequestIntegrations) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(ExportRequestIntegrationsAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsAuthorization); ok {
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
				// if coming in as map, convert it back
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

	if val, ok := kv["location"].(ExportRequestIntegrationsLocation); ok {
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
		} else if sv, ok := val.(ExportRequestIntegrationsProgress); ok {
			// struct
			o.Progress = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsProgress); ok {
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

	if val, ok := kv["system_type"].(ExportRequestIntegrationsSystemType); ok {
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
			case "user", "USER":
				o.SystemType = 3
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
			case "user", "USER":
				o.SystemType = 3
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
				// if coming in as map, convert it back
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
		} else if sv, ok := val.(ExportRequestIntegrationsValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsValidatedDate); ok {
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
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ValidationMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestRequestDate represents the object structure for request_date
type ExportRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *ExportRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// ExportRequest an agent action to request an export
type ExportRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integrations The integrations that should be exported and their current configuration
	Integrations []ExportRequestIntegrations `json:"integrations" codec:"integrations" bson:"integrations" yaml:"integrations" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical bool `json:"reprocess_historical" codec:"reprocess_historical" bson:"reprocess_historical" yaml:"reprocess_historical" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate ExportRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UploadHeaders Any upload headers to include in the HTTP upload
	UploadHeaders []string `json:"upload_headers" codec:"upload_headers" bson:"upload_headers" yaml:"upload_headers" faker:"-"`
	// UploadURL The one time upload URL to use for uploading a job to the Pinpoint cloud
	UploadURL *string `json:"upload_url,omitempty" codec:"upload_url,omitempty" bson:"upload_url" yaml:"upload_url,omitempty" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportRequest)(nil)

func toExportRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequest:
		return v.ToMap()

	case []ExportRequestIntegrations:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case ExportRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ExportRequest
func (o *ExportRequest) String() string {
	return fmt.Sprintf("agent.ExportRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportRequest) GetTopicName() datamodel.TopicNameType {
	return ExportRequestTopic
}

// GetStreamName returns the name of the stream
func (o *ExportRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ExportRequest) GetModelName() datamodel.ModelNameType {
	return ExportRequestModelName
}

// NewExportRequestID provides a template for generating an ID field for ExportRequest
func NewExportRequestID(customerID string, refType string, refID string) string {
	return hash.Values("ExportRequest", customerID, refType, refID)
}

func (o *ExportRequest) setDefaults(frommap bool) {
	if o.Integrations == nil {
		o.Integrations = make([]ExportRequestIntegrations, 0)
	}
	if o.UploadHeaders == nil {
		o.UploadHeaders = make([]string, 0)
	}
	if o.UploadURL == nil {
		o.UploadURL = &emptyString
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportRequest) GetID() string {
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

// IsMutable returns true if the model is mutable
func (o *ExportRequest) IsMutable() bool {
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *ExportRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two ExportRequest objects are equal
func (o *ExportRequest) IsEqual(other *ExportRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":          toExportRequestObject(o.CustomerID, false),
		"id":                   toExportRequestObject(o.ID, false),
		"integrations":         toExportRequestObject(o.Integrations, false),
		"job_id":               toExportRequestObject(o.JobID, false),
		"ref_id":               toExportRequestObject(o.RefID, false),
		"ref_type":             toExportRequestObject(o.RefType, false),
		"reprocess_historical": toExportRequestObject(o.ReprocessHistorical, false),
		"request_date":         toExportRequestObject(o.RequestDate, false),
		"updated_ts":           toExportRequestObject(o.UpdatedAt, false),
		"upload_headers":       toExportRequestObject(o.UploadHeaders, false),
		"upload_url":           toExportRequestObject(o.UploadURL, true),
		"uuid":                 toExportRequestObject(o.UUID, false),
		"hashcode":             toExportRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequest) FromMap(kv map[string]interface{}) {

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

	if o == nil {

		o.Integrations = make([]ExportRequestIntegrations, 0)

	}
	if val, ok := kv["integrations"]; ok {
		if sv, ok := val.([]ExportRequestIntegrations); ok {
			o.Integrations = sv
		} else if sp, ok := val.([]*ExportRequestIntegrations); ok {
			o.Integrations = o.Integrations[:0]
			for _, e := range sp {
				o.Integrations = append(o.Integrations, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ExportRequestIntegrations); ok {
					o.Integrations = append(o.Integrations, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ExportRequestIntegrations
					fm.FromMap(av)
					o.Integrations = append(o.Integrations, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av ExportRequestIntegrations
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for integrations field entry: " + reflect.TypeOf(ae).String())
					}
					o.Integrations = append(o.Integrations, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ExportRequestIntegrations); ok {
					o.Integrations = append(o.Integrations, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ExportRequestIntegrations
					fm.FromMap(r)
					o.Integrations = append(o.Integrations, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ExportRequestIntegrations{}
					fm.FromMap(r)
					o.Integrations = append(o.Integrations, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm ExportRequestIntegrations
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Integrations = append(o.Integrations, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["job_id"].(string); ok {
		o.JobID = val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.JobID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = number.ToBoolAny(nil)
			} else {
				o.ReprocessHistorical = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*ExportRequestRequestDate); ok {
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

	if val, ok := kv["upload_headers"]; ok {
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
								panic("unsupported type for upload_headers field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for upload_headers field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for upload_headers field")
				}
			}
			o.UploadHeaders = na
		}
	}
	if o.UploadHeaders == nil {
		o.UploadHeaders = make([]string, 0)
	}

	if val, ok := kv["upload_url"].(*string); ok {
		o.UploadURL = val
	} else if val, ok := kv["upload_url"].(string); ok {
		o.UploadURL = &val
	} else {
		if val, ok := kv["upload_url"]; ok {
			if val == nil {
				o.UploadURL = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UploadURL = pstrings.Pointer(fmt.Sprintf("%v", val))
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
func (o *ExportRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integrations)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	args = append(args, o.RequestDate)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UploadHeaders)
	args = append(args, o.UploadURL)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *ExportRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
