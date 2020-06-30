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
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// ExportRequestTable is the default table name
	ExportRequestTable datamodel.ModelNameType = "agent_exportrequest"

	// ExportRequestModelName is the model name
	ExportRequestModelName datamodel.ModelNameType = "agent.ExportRequest"
)

const (
	// ExportRequestModelCustomerIDColumn is the column json value customer_id
	ExportRequestModelCustomerIDColumn = "customer_id"
	// ExportRequestModelIDColumn is the column json value id
	ExportRequestModelIDColumn = "id"
	// ExportRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportRequestModelIntegrationsColumn is the column json value integrations
	ExportRequestModelIntegrationsColumn = "integrations"
	// ExportRequestModelIntegrationsActiveColumn is the column json value active
	ExportRequestModelIntegrationsActiveColumn = "active"
	// ExportRequestModelIntegrationsAuthorizationColumn is the column json value authorization
	ExportRequestModelIntegrationsAuthorizationColumn = "authorization"
	// ExportRequestModelIntegrationsAuthorizationAccessTokenColumn is the column json value access_token
	ExportRequestModelIntegrationsAuthorizationAccessTokenColumn = "access_token"
	// ExportRequestModelIntegrationsAuthorizationRefreshTokenColumn is the column json value refresh_token
	ExportRequestModelIntegrationsAuthorizationRefreshTokenColumn = "refresh_token"
	// ExportRequestModelIntegrationsAuthorizationURLColumn is the column json value url
	ExportRequestModelIntegrationsAuthorizationURLColumn = "url"
	// ExportRequestModelIntegrationsAuthorizationUsernameColumn is the column json value username
	ExportRequestModelIntegrationsAuthorizationUsernameColumn = "username"
	// ExportRequestModelIntegrationsAuthorizationPasswordColumn is the column json value password
	ExportRequestModelIntegrationsAuthorizationPasswordColumn = "password"
	// ExportRequestModelIntegrationsAuthorizationAPITokenColumn is the column json value api_token
	ExportRequestModelIntegrationsAuthorizationAPITokenColumn = "api_token"
	// ExportRequestModelIntegrationsAuthorizationCollectionNameColumn is the column json value collection_name
	ExportRequestModelIntegrationsAuthorizationCollectionNameColumn = "collection_name"
	// ExportRequestModelIntegrationsAuthorizationAPIKeyColumn is the column json value api_key
	ExportRequestModelIntegrationsAuthorizationAPIKeyColumn = "api_key"
	// ExportRequestModelIntegrationsAuthorizationAuthorizationColumn is the column json value authorization
	ExportRequestModelIntegrationsAuthorizationAuthorizationColumn = "authorization"
	// ExportRequestModelIntegrationsAuthorizationHostnameColumn is the column json value hostname
	ExportRequestModelIntegrationsAuthorizationHostnameColumn = "hostname"
	// ExportRequestModelIntegrationsAuthorizationAPIVersionColumn is the column json value api_version
	ExportRequestModelIntegrationsAuthorizationAPIVersionColumn = "api_version"
	// ExportRequestModelIntegrationsAuthorizationOrganizationColumn is the column json value organization
	ExportRequestModelIntegrationsAuthorizationOrganizationColumn = "organization"
	// ExportRequestModelIntegrationsConfigColumn is the column json value config
	ExportRequestModelIntegrationsConfigColumn = "config"
	// ExportRequestModelIntegrationsCreatedByProfileIDColumn is the column json value created_by_profile_id
	ExportRequestModelIntegrationsCreatedByProfileIDColumn = "created_by_profile_id"
	// ExportRequestModelIntegrationsCreatedByUserIDColumn is the column json value created_by_user_id
	ExportRequestModelIntegrationsCreatedByUserIDColumn = "created_by_user_id"
	// ExportRequestModelIntegrationsCustomerIDColumn is the column json value customer_id
	ExportRequestModelIntegrationsCustomerIDColumn = "customer_id"
	// ExportRequestModelIntegrationsEntityErrorsColumn is the column json value entity_errors
	ExportRequestModelIntegrationsEntityErrorsColumn = "entity_errors"
	// ExportRequestModelIntegrationsEntityErrorsIDColumn is the column json value id
	ExportRequestModelIntegrationsEntityErrorsIDColumn = "id"
	// ExportRequestModelIntegrationsEntityErrorsRefIDColumn is the column json value ref_id
	ExportRequestModelIntegrationsEntityErrorsRefIDColumn = "ref_id"
	// ExportRequestModelIntegrationsEntityErrorsErrorColumn is the column json value error
	ExportRequestModelIntegrationsEntityErrorsErrorColumn = "error"
	// ExportRequestModelIntegrationsErrorMessageColumn is the column json value error_message
	ExportRequestModelIntegrationsErrorMessageColumn = "error_message"
	// ExportRequestModelIntegrationsErroredColumn is the column json value errored
	ExportRequestModelIntegrationsErroredColumn = "errored"
	// ExportRequestModelIntegrationsExclusionsColumn is the column json value exclusions
	ExportRequestModelIntegrationsExclusionsColumn = "exclusions"
	// ExportRequestModelIntegrationsExportableColumn is the column json value exportable
	ExportRequestModelIntegrationsExportableColumn = "exportable"
	// ExportRequestModelIntegrationsIDColumn is the column json value id
	ExportRequestModelIntegrationsIDColumn = "id"
	// ExportRequestModelIntegrationsInclusionsColumn is the column json value inclusions
	ExportRequestModelIntegrationsInclusionsColumn = "inclusions"
	// ExportRequestModelIntegrationsIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportRequestModelIntegrationsIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportRequestModelIntegrationsIntervalColumn is the column json value interval
	ExportRequestModelIntegrationsIntervalColumn = "interval"
	// ExportRequestModelIntegrationsLastExportCompletedDateColumn is the column json value last_export_completed_date
	ExportRequestModelIntegrationsLastExportCompletedDateColumn = "last_export_completed_date"
	// ExportRequestModelIntegrationsLastExportCompletedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsLastExportCompletedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsLastExportCompletedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsLastExportCompletedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsLastExportCompletedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsLastExportCompletedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsLastExportRequestedDateColumn is the column json value last_export_requested_date
	ExportRequestModelIntegrationsLastExportRequestedDateColumn = "last_export_requested_date"
	// ExportRequestModelIntegrationsLastExportRequestedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsLastExportRequestedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsLastExportRequestedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsLastExportRequestedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsLastExportRequestedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsLastExportRequestedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsLastProcessingCompletedDateColumn is the column json value last_processing_completed_date
	ExportRequestModelIntegrationsLastProcessingCompletedDateColumn = "last_processing_completed_date"
	// ExportRequestModelIntegrationsLastProcessingCompletedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsLastProcessingCompletedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsLastProcessingCompletedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsLastProcessingCompletedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsLastProcessingCompletedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsLastProcessingCompletedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsLastProcessingStartedDateColumn is the column json value last_processing_started_date
	ExportRequestModelIntegrationsLastProcessingStartedDateColumn = "last_processing_started_date"
	// ExportRequestModelIntegrationsLastProcessingStartedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsLastProcessingStartedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsLastProcessingStartedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsLastProcessingStartedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsLastProcessingStartedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsLastProcessingStartedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsLocationColumn is the column json value location
	ExportRequestModelIntegrationsLocationColumn = "location"
	// ExportRequestModelIntegrationsNameColumn is the column json value name
	ExportRequestModelIntegrationsNameColumn = "name"
	// ExportRequestModelIntegrationsOnboardCompletedDateColumn is the column json value onboard_completed_date
	ExportRequestModelIntegrationsOnboardCompletedDateColumn = "onboard_completed_date"
	// ExportRequestModelIntegrationsOnboardCompletedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsOnboardCompletedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsOnboardCompletedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsOnboardCompletedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsOnboardCompletedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsOnboardCompletedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsOnboardRequestedDateColumn is the column json value onboard_requested_date
	ExportRequestModelIntegrationsOnboardRequestedDateColumn = "onboard_requested_date"
	// ExportRequestModelIntegrationsOnboardRequestedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsOnboardRequestedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsOnboardRequestedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsOnboardRequestedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsOnboardRequestedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsOnboardRequestedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsOnboardingColumn is the column json value onboarding
	ExportRequestModelIntegrationsOnboardingColumn = "onboarding"
	// ExportRequestModelIntegrationsOrganizationColumn is the column json value organization
	ExportRequestModelIntegrationsOrganizationColumn = "organization"
	// ExportRequestModelIntegrationsPausedColumn is the column json value paused
	ExportRequestModelIntegrationsPausedColumn = "paused"
	// ExportRequestModelIntegrationsProcessedColumn is the column json value processed
	ExportRequestModelIntegrationsProcessedColumn = "processed"
	// ExportRequestModelIntegrationsRefIDColumn is the column json value ref_id
	ExportRequestModelIntegrationsRefIDColumn = "ref_id"
	// ExportRequestModelIntegrationsRefTypeColumn is the column json value ref_type
	ExportRequestModelIntegrationsRefTypeColumn = "ref_type"
	// ExportRequestModelIntegrationsServerVersionColumn is the column json value server_version
	ExportRequestModelIntegrationsServerVersionColumn = "server_version"
	// ExportRequestModelIntegrationsStateColumn is the column json value state
	ExportRequestModelIntegrationsStateColumn = "state"
	// ExportRequestModelIntegrationsSystemTypeColumn is the column json value system_type
	ExportRequestModelIntegrationsSystemTypeColumn = "system_type"
	// ExportRequestModelIntegrationsTeamIDColumn is the column json value team_id
	ExportRequestModelIntegrationsTeamIDColumn = "team_id"
	// ExportRequestModelIntegrationsThrottledColumn is the column json value throttled
	ExportRequestModelIntegrationsThrottledColumn = "throttled"
	// ExportRequestModelIntegrationsThrottledUntilColumn is the column json value throttled_until
	ExportRequestModelIntegrationsThrottledUntilColumn = "throttled_until"
	// ExportRequestModelIntegrationsThrottledUntilEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsThrottledUntilEpochColumn = "epoch"
	// ExportRequestModelIntegrationsThrottledUntilOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsThrottledUntilOffsetColumn = "offset"
	// ExportRequestModelIntegrationsThrottledUntilRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsThrottledUntilRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsValidatedColumn is the column json value validated
	ExportRequestModelIntegrationsValidatedColumn = "validated"
	// ExportRequestModelIntegrationsValidatedDateColumn is the column json value validated_date
	ExportRequestModelIntegrationsValidatedDateColumn = "validated_date"
	// ExportRequestModelIntegrationsValidatedDateEpochColumn is the column json value epoch
	ExportRequestModelIntegrationsValidatedDateEpochColumn = "epoch"
	// ExportRequestModelIntegrationsValidatedDateOffsetColumn is the column json value offset
	ExportRequestModelIntegrationsValidatedDateOffsetColumn = "offset"
	// ExportRequestModelIntegrationsValidatedDateRfc3339Column is the column json value rfc3339
	ExportRequestModelIntegrationsValidatedDateRfc3339Column = "rfc3339"
	// ExportRequestModelIntegrationsValidationMessageColumn is the column json value validation_message
	ExportRequestModelIntegrationsValidationMessageColumn = "validation_message"
	// ExportRequestModelJobIDColumn is the column json value job_id
	ExportRequestModelJobIDColumn = "job_id"
	// ExportRequestModelRefIDColumn is the column json value ref_id
	ExportRequestModelRefIDColumn = "ref_id"
	// ExportRequestModelRefTypeColumn is the column json value ref_type
	ExportRequestModelRefTypeColumn = "ref_type"
	// ExportRequestModelReprocessHistoricalColumn is the column json value reprocess_historical
	ExportRequestModelReprocessHistoricalColumn = "reprocess_historical"
	// ExportRequestModelRequestDateColumn is the column json value request_date
	ExportRequestModelRequestDateColumn = "request_date"
	// ExportRequestModelRequestDateEpochColumn is the column json value epoch
	ExportRequestModelRequestDateEpochColumn = "epoch"
	// ExportRequestModelRequestDateOffsetColumn is the column json value offset
	ExportRequestModelRequestDateOffsetColumn = "offset"
	// ExportRequestModelRequestDateRfc3339Column is the column json value rfc3339
	ExportRequestModelRequestDateRfc3339Column = "rfc3339"
	// ExportRequestModelUploadURLColumn is the column json value upload_url
	ExportRequestModelUploadURLColumn = "upload_url"
	// ExportRequestModelUUIDColumn is the column json value uuid
	ExportRequestModelUUIDColumn = "uuid"
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
	// CollectionName Collection name for instance, if relevant
	CollectionName *string `json:"collection_name,omitempty" codec:"collection_name,omitempty" bson:"collection_name" yaml:"collection_name,omitempty" faker:"-"`
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

// ToMap returns the object as a map
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
		// CollectionName Collection name for instance, if relevant
		"collection_name": toExportRequestIntegrationsAuthorizationObject(o.CollectionName, true),
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
	if val, ok := kv["collection_name"].(*string); ok {
		o.CollectionName = val
	} else if val, ok := kv["collection_name"].(string); ok {
		o.CollectionName = &val
	} else {
		if val, ok := kv["collection_name"]; ok {
			if val == nil {
				o.CollectionName = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CollectionName = pstrings.Pointer(fmt.Sprintf("%v", val))
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

// ExportRequestIntegrationsEntityErrors represents the object structure for entity_errors
type ExportRequestIntegrationsEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
}

func toExportRequestIntegrationsEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toExportRequestIntegrationsEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toExportRequestIntegrationsEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export
		"error": toExportRequestIntegrationsEntityErrorsObject(o.Error, false),
	}
}

func (o *ExportRequestIntegrationsEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsEntityErrors) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		if val, ok := kv["id"]; ok {
			if val == nil {
				o.ID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["error"].(string); ok {
		o.Error = val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Error = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLastExportCompletedDate represents the object structure for last_export_completed_date
type ExportRequestIntegrationsLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsLastExportCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLastExportRequestedDate represents the object structure for last_export_requested_date
type ExportRequestIntegrationsLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsLastExportRequestedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLastProcessingCompletedDate represents the object structure for last_processing_completed_date
type ExportRequestIntegrationsLastProcessingCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsLastProcessingCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsLastProcessingCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsLastProcessingCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsLastProcessingCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsLastProcessingCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsLastProcessingCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsLastProcessingCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsLastProcessingCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLastProcessingStartedDate represents the object structure for last_processing_started_date
type ExportRequestIntegrationsLastProcessingStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsLastProcessingStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsLastProcessingStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsLastProcessingStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsLastProcessingStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsLastProcessingStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsLastProcessingStartedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsLastProcessingStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsLastProcessingStartedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsLocation is the enumeration type for location
type ExportRequestIntegrationsLocation int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportRequestIntegrationsLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportRequestIntegrationsLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = ExportRequestIntegrationsLocation(0)
		case "CLOUD":
			*v = ExportRequestIntegrationsLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportRequestIntegrationsLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportRequestIntegrationsLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
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

// FromInterface for decoding from an interface
func (v *ExportRequestIntegrationsLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportRequestIntegrationsLocation:
		*v = val
	case int32:
		*v = ExportRequestIntegrationsLocation(int32(val))
	case int:
		*v = ExportRequestIntegrationsLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = ExportRequestIntegrationsLocation(0)
		case "CLOUD":
			*v = ExportRequestIntegrationsLocation(1)
		}
	}
	return nil
}

const (
	// ExportRequestIntegrationsLocationPrivate is the enumeration value for private
	ExportRequestIntegrationsLocationPrivate ExportRequestIntegrationsLocation = 0
	// ExportRequestIntegrationsLocationCloud is the enumeration value for cloud
	ExportRequestIntegrationsLocationCloud ExportRequestIntegrationsLocation = 1
)

// ExportRequestIntegrationsOnboardCompletedDate represents the object structure for onboard_completed_date
type ExportRequestIntegrationsOnboardCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsOnboardCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsOnboardCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsOnboardCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsOnboardCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsOnboardCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsOnboardCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsOnboardCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsOnboardCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsOnboardRequestedDate represents the object structure for onboard_requested_date
type ExportRequestIntegrationsOnboardRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsOnboardRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsOnboardRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsOnboardRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsOnboardRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsOnboardRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsOnboardRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsOnboardRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsOnboardRequestedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// ExportRequestIntegrationsState is the enumeration type for state
type ExportRequestIntegrationsState int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportRequestIntegrationsState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportRequestIntegrationsState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = ExportRequestIntegrationsState(0)
		case "EXPORTING":
			*v = ExportRequestIntegrationsState(1)
		case "PROCESSING":
			*v = ExportRequestIntegrationsState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportRequestIntegrationsState) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "IDLE":
		v = 0
	case "EXPORTING":
		v = 1
	case "PROCESSING":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportRequestIntegrationsState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("EXPORTING")
	case 2:
		return json.Marshal("PROCESSING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationsState
func (v ExportRequestIntegrationsState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "EXPORTING"
	case 2:
		return "PROCESSING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportRequestIntegrationsState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportRequestIntegrationsState:
		*v = val
	case int32:
		*v = ExportRequestIntegrationsState(int32(val))
	case int:
		*v = ExportRequestIntegrationsState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = ExportRequestIntegrationsState(0)
		case "EXPORTING":
			*v = ExportRequestIntegrationsState(1)
		case "PROCESSING":
			*v = ExportRequestIntegrationsState(2)
		}
	}
	return nil
}

const (
	// ExportRequestIntegrationsStateIdle is the enumeration value for idle
	ExportRequestIntegrationsStateIdle ExportRequestIntegrationsState = 0
	// ExportRequestIntegrationsStateExporting is the enumeration value for exporting
	ExportRequestIntegrationsStateExporting ExportRequestIntegrationsState = 1
	// ExportRequestIntegrationsStateProcessing is the enumeration value for processing
	ExportRequestIntegrationsStateProcessing ExportRequestIntegrationsState = 2
)

// ExportRequestIntegrationsSystemType is the enumeration type for system_type
type ExportRequestIntegrationsSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportRequestIntegrationsSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportRequestIntegrationsSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = ExportRequestIntegrationsSystemType(0)
		case "SOURCECODE":
			*v = ExportRequestIntegrationsSystemType(1)
		case "CODEQUALITY":
			*v = ExportRequestIntegrationsSystemType(2)
		case "CALENDAR":
			*v = ExportRequestIntegrationsSystemType(3)
		case "USER":
			*v = ExportRequestIntegrationsSystemType(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportRequestIntegrationsSystemType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "WORK":
		v = 0
	case "SOURCECODE":
		v = 1
	case "CODEQUALITY":
		v = 2
	case "CALENDAR":
		v = 3
	case "USER":
		v = 4
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportRequestIntegrationsSystemType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("WORK")
	case 1:
		return json.Marshal("SOURCECODE")
	case 2:
		return json.Marshal("CODEQUALITY")
	case 3:
		return json.Marshal("CALENDAR")
	case 4:
		return json.Marshal("USER")
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
		return "CALENDAR"
	case 4:
		return "USER"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportRequestIntegrationsSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportRequestIntegrationsSystemType:
		*v = val
	case int32:
		*v = ExportRequestIntegrationsSystemType(int32(val))
	case int:
		*v = ExportRequestIntegrationsSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = ExportRequestIntegrationsSystemType(0)
		case "SOURCECODE":
			*v = ExportRequestIntegrationsSystemType(1)
		case "CODEQUALITY":
			*v = ExportRequestIntegrationsSystemType(2)
		case "CALENDAR":
			*v = ExportRequestIntegrationsSystemType(3)
		case "USER":
			*v = ExportRequestIntegrationsSystemType(4)
		}
	}
	return nil
}

const (
	// ExportRequestIntegrationsSystemTypeWork is the enumeration value for work
	ExportRequestIntegrationsSystemTypeWork ExportRequestIntegrationsSystemType = 0
	// ExportRequestIntegrationsSystemTypeSourcecode is the enumeration value for sourcecode
	ExportRequestIntegrationsSystemTypeSourcecode ExportRequestIntegrationsSystemType = 1
	// ExportRequestIntegrationsSystemTypeCodequality is the enumeration value for codequality
	ExportRequestIntegrationsSystemTypeCodequality ExportRequestIntegrationsSystemType = 2
	// ExportRequestIntegrationsSystemTypeCalendar is the enumeration value for calendar
	ExportRequestIntegrationsSystemTypeCalendar ExportRequestIntegrationsSystemType = 3
	// ExportRequestIntegrationsSystemTypeUser is the enumeration value for user
	ExportRequestIntegrationsSystemTypeUser ExportRequestIntegrationsSystemType = 4
)

// ExportRequestIntegrationsThrottledUntil represents the object structure for throttled_until
type ExportRequestIntegrationsThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportRequestIntegrationsThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportRequestIntegrationsThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrationsThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportRequestIntegrationsThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportRequestIntegrationsThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportRequestIntegrationsThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *ExportRequestIntegrationsThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportRequestIntegrationsThrottledUntil) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

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

// ToMap returns the object as a map
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
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []ExportRequestIntegrationsEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" codec:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Exportable a flag to indicate if the integration is ready for export
	Exportable bool `json:"exportable" codec:"exportable" bson:"exportable" yaml:"exportable" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Inclusions The inclusion list for this integration
	Inclusions []string `json:"inclusions" codec:"inclusions" bson:"inclusions" yaml:"inclusions" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval int64 `json:"interval" codec:"interval" bson:"interval" yaml:"interval" faker:"-"`
	// LastExportCompletedDate when the export response was received (set by the backend)
	LastExportCompletedDate ExportRequestIntegrationsLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made (set by the backend)
	LastExportRequestedDate ExportRequestIntegrationsLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingCompletedDate when the processing completes (set by the pipeline)
	LastProcessingCompletedDate ExportRequestIntegrationsLastProcessingCompletedDate `json:"last_processing_completed_date" codec:"last_processing_completed_date" bson:"last_processing_completed_date" yaml:"last_processing_completed_date" faker:"-"`
	// LastProcessingStartedDate when the processing starts (set by the pipeline)
	LastProcessingStartedDate ExportRequestIntegrationsLastProcessingStartedDate `json:"last_processing_started_date" codec:"last_processing_started_date" bson:"last_processing_started_date" yaml:"last_processing_started_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location ExportRequestIntegrationsLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OnboardCompletedDate when the last request for integration metadata was finished
	OnboardCompletedDate ExportRequestIntegrationsOnboardCompletedDate `json:"onboard_completed_date" codec:"onboard_completed_date" bson:"onboard_completed_date" yaml:"onboard_completed_date" faker:"-"`
	// OnboardRequestedDate when the last request for integration metadata was made
	OnboardRequestedDate ExportRequestIntegrationsOnboardRequestedDate `json:"onboard_requested_date" codec:"onboard_requested_date" bson:"onboard_requested_date" yaml:"onboard_requested_date" faker:"-"`
	// Onboarding true if the agent is fetching metadata for the integration
	Onboarding bool `json:"onboarding" codec:"onboarding" bson:"onboarding" yaml:"onboarding" faker:"-"`
	// Organization The origanization authorized. Used for azure integrations
	Organization *string `json:"organization,omitempty" codec:"organization,omitempty" bson:"organization" yaml:"organization,omitempty" faker:"-"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused bool `json:"paused" codec:"paused" bson:"paused" yaml:"paused" faker:"-"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ServerVersion the server version for this integration
	ServerVersion *string `json:"server_version,omitempty" codec:"server_version,omitempty" bson:"server_version" yaml:"server_version,omitempty" faker:"-"`
	// State the current state of the integration
	State ExportRequestIntegrationsState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType ExportRequestIntegrationsSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
	TeamID *string `json:"team_id,omitempty" codec:"team_id,omitempty" bson:"team_id" yaml:"team_id,omitempty" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *ExportRequestIntegrationsThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
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

	case []ExportRequestIntegrationsEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case ExportRequestIntegrationsLastExportCompletedDate:
		return v.ToMap()

	case ExportRequestIntegrationsLastExportRequestedDate:
		return v.ToMap()

	case ExportRequestIntegrationsLastProcessingCompletedDate:
		return v.ToMap()

	case ExportRequestIntegrationsLastProcessingStartedDate:
		return v.ToMap()

	case ExportRequestIntegrationsLocation:
		return v.String()

	case ExportRequestIntegrationsOnboardCompletedDate:
		return v.ToMap()

	case ExportRequestIntegrationsOnboardRequestedDate:
		return v.ToMap()

	case ExportRequestIntegrationsState:
		return v.String()

	case ExportRequestIntegrationsSystemType:
		return v.String()

	case *ExportRequestIntegrationsThrottledUntil:
		return v.ToMap()

	case ExportRequestIntegrationsValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportRequestIntegrations) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toExportRequestIntegrationsObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toExportRequestIntegrationsObject(o.Authorization, false),
		// Config the integration configuration controlled by the integration itself
		"config": toExportRequestIntegrationsObject(o.Config, true),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toExportRequestIntegrationsObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toExportRequestIntegrationsObject(o.CreatedByUserID, true),
		// CustomerID the customer id for the model instance
		"customer_id": toExportRequestIntegrationsObject(o.CustomerID, false),
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toExportRequestIntegrationsObject(o.EntityErrors, false),
		// ErrorMessage The error message from an export run
		"error_message": toExportRequestIntegrationsObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent or any other error
		"errored": toExportRequestIntegrationsObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toExportRequestIntegrationsObject(o.Exclusions, false),
		// Exportable a flag to indicate if the integration is ready for export
		"exportable": toExportRequestIntegrationsObject(o.Exportable, false),
		// ID the primary key for the model instance
		"id": toExportRequestIntegrationsObject(o.ID, false),
		// Inclusions The inclusion list for this integration
		"inclusions": toExportRequestIntegrationsObject(o.Inclusions, false),
		// IntegrationInstanceID the integration instance id
		"integration_instance_id": toExportRequestIntegrationsObject(o.IntegrationInstanceID, true),
		// Interval the interval in milliseconds for how often an export job is scheduled
		"interval": toExportRequestIntegrationsObject(o.Interval, false),
		// LastExportCompletedDate when the export response was received (set by the backend)
		"last_export_completed_date": toExportRequestIntegrationsObject(o.LastExportCompletedDate, false),
		// LastExportRequestedDate when the export request was made (set by the backend)
		"last_export_requested_date": toExportRequestIntegrationsObject(o.LastExportRequestedDate, false),
		// LastProcessingCompletedDate when the processing completes (set by the pipeline)
		"last_processing_completed_date": toExportRequestIntegrationsObject(o.LastProcessingCompletedDate, false),
		// LastProcessingStartedDate when the processing starts (set by the pipeline)
		"last_processing_started_date": toExportRequestIntegrationsObject(o.LastProcessingStartedDate, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toExportRequestIntegrationsObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toExportRequestIntegrationsObject(o.Name, false),
		// OnboardCompletedDate when the last request for integration metadata was finished
		"onboard_completed_date": toExportRequestIntegrationsObject(o.OnboardCompletedDate, false),
		// OnboardRequestedDate when the last request for integration metadata was made
		"onboard_requested_date": toExportRequestIntegrationsObject(o.OnboardRequestedDate, false),
		// Onboarding true if the agent is fetching metadata for the integration
		"onboarding": toExportRequestIntegrationsObject(o.Onboarding, false),
		// Organization The origanization authorized. Used for azure integrations
		"organization": toExportRequestIntegrationsObject(o.Organization, true),
		// Paused true if the agent is paused and should not start new scheduled jobs
		"paused": toExportRequestIntegrationsObject(o.Paused, false),
		// Processed If the integration has been processed at least once
		"processed": toExportRequestIntegrationsObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toExportRequestIntegrationsObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toExportRequestIntegrationsObject(o.RefType, false),
		// ServerVersion the server version for this integration
		"server_version": toExportRequestIntegrationsObject(o.ServerVersion, true),
		// State the current state of the integration
		"state": toExportRequestIntegrationsObject(o.State, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toExportRequestIntegrationsObject(o.SystemType, false),
		// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
		"team_id": toExportRequestIntegrationsObject(o.TeamID, true),
		// Throttled Set to true when integration is throttled.
		"throttled": toExportRequestIntegrationsObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toExportRequestIntegrationsObject(o.ThrottledUntil, true),
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

	if o.Processed == nil {
		var v bool
		o.Processed = &v
	}

	if o.Throttled == nil {
		var v bool
		o.Throttled = &v
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
				o.Active = false
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

	if val, ok := kv["config"].(*string); ok {
		o.Config = val
	} else if val, ok := kv["config"].(string); ok {
		o.Config = &val
	} else {
		if val, ok := kv["config"]; ok {
			if val == nil {
				o.Config = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Config = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_profile_id"].(*string); ok {
		o.CreatedByProfileID = val
	} else if val, ok := kv["created_by_profile_id"].(string); ok {
		o.CreatedByProfileID = &val
	} else {
		if val, ok := kv["created_by_profile_id"]; ok {
			if val == nil {
				o.CreatedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["created_by_user_id"].(*string); ok {
		o.CreatedByUserID = val
	} else if val, ok := kv["created_by_user_id"].(string); ok {
		o.CreatedByUserID = &val
	} else {
		if val, ok := kv["created_by_user_id"]; ok {
			if val == nil {
				o.CreatedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CreatedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if o == nil {

		o.EntityErrors = make([]ExportRequestIntegrationsEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]ExportRequestIntegrationsEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*ExportRequestIntegrationsEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ExportRequestIntegrationsEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ExportRequestIntegrationsEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ExportRequestIntegrationsEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ExportRequestIntegrationsEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ExportRequestIntegrationsEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ExportRequestIntegrationsEntityErrors{}
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
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
							var fm ExportRequestIntegrationsEntityErrors
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.EntityErrors = append(o.EntityErrors, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["error_message"].(*string); ok {
		o.ErrorMessage = val
	} else if val, ok := kv["error_message"].(string); ok {
		o.ErrorMessage = &val
	} else {
		if val, ok := kv["error_message"]; ok {
			if val == nil {
				o.ErrorMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ErrorMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
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
				o.Errored = nil
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
	if val, ok := kv["exportable"].(bool); ok {
		o.Exportable = val
	} else {
		if val, ok := kv["exportable"]; ok {
			if val == nil {
				o.Exportable = false
			} else {
				o.Exportable = number.ToBoolAny(val)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["inclusions"]; ok {
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
								panic("unsupported type for inclusions field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for inclusions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for inclusions field")
				}
			}
			o.Inclusions = na
		}
	}
	if o.Inclusions == nil {
		o.Inclusions = make([]string, 0)
	}
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["interval"].(int64); ok {
		o.Interval = val
	} else {
		if val, ok := kv["interval"]; ok {
			if val == nil {
				o.Interval = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Interval = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportCompletedDate.Epoch = dt.Epoch
				o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastExportCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportCompletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_export_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportRequestedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportRequestedDate.Epoch = dt.Epoch
				o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
				o.LastExportRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportRequestedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_processing_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingCompletedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsLastProcessingCompletedDate); ok {
			// struct
			o.LastProcessingCompletedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsLastProcessingCompletedDate); ok {
			// struct pointer
			o.LastProcessingCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingCompletedDate.Epoch = dt.Epoch
			o.LastProcessingCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastProcessingCompletedDate.Epoch = dt.Epoch
			o.LastProcessingCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastProcessingCompletedDate.Epoch = dt.Epoch
				o.LastProcessingCompletedDate.Rfc3339 = dt.Rfc3339
				o.LastProcessingCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastProcessingCompletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_processing_started_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingStartedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsLastProcessingStartedDate); ok {
			// struct
			o.LastProcessingStartedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsLastProcessingStartedDate); ok {
			// struct pointer
			o.LastProcessingStartedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingStartedDate.Epoch = dt.Epoch
			o.LastProcessingStartedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingStartedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastProcessingStartedDate.Epoch = dt.Epoch
			o.LastProcessingStartedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingStartedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastProcessingStartedDate.Epoch = dt.Epoch
				o.LastProcessingStartedDate.Rfc3339 = dt.Rfc3339
				o.LastProcessingStartedDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastProcessingStartedDate.FromMap(map[string]interface{}{})
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Name = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["onboard_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.OnboardCompletedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsOnboardCompletedDate); ok {
			// struct
			o.OnboardCompletedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsOnboardCompletedDate); ok {
			// struct pointer
			o.OnboardCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.OnboardCompletedDate.Epoch = dt.Epoch
			o.OnboardCompletedDate.Rfc3339 = dt.Rfc3339
			o.OnboardCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.OnboardCompletedDate.Epoch = dt.Epoch
			o.OnboardCompletedDate.Rfc3339 = dt.Rfc3339
			o.OnboardCompletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.OnboardCompletedDate.Epoch = dt.Epoch
				o.OnboardCompletedDate.Rfc3339 = dt.Rfc3339
				o.OnboardCompletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.OnboardCompletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["onboard_requested_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.OnboardRequestedDate.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsOnboardRequestedDate); ok {
			// struct
			o.OnboardRequestedDate = sv
		} else if sp, ok := val.(*ExportRequestIntegrationsOnboardRequestedDate); ok {
			// struct pointer
			o.OnboardRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.OnboardRequestedDate.Epoch = dt.Epoch
			o.OnboardRequestedDate.Rfc3339 = dt.Rfc3339
			o.OnboardRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.OnboardRequestedDate.Epoch = dt.Epoch
			o.OnboardRequestedDate.Rfc3339 = dt.Rfc3339
			o.OnboardRequestedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.OnboardRequestedDate.Epoch = dt.Epoch
				o.OnboardRequestedDate.Rfc3339 = dt.Rfc3339
				o.OnboardRequestedDate.Offset = dt.Offset
			}
		}
	} else {
		o.OnboardRequestedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["onboarding"].(bool); ok {
		o.Onboarding = val
	} else {
		if val, ok := kv["onboarding"]; ok {
			if val == nil {
				o.Onboarding = false
			} else {
				o.Onboarding = number.ToBoolAny(val)
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
	if val, ok := kv["paused"].(bool); ok {
		o.Paused = val
	} else {
		if val, ok := kv["paused"]; ok {
			if val == nil {
				o.Paused = false
			} else {
				o.Paused = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["processed"].(*bool); ok {
		o.Processed = val
	} else if val, ok := kv["processed"].(bool); ok {
		o.Processed = &val
	} else {
		if val, ok := kv["processed"]; ok {
			if val == nil {
				o.Processed = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Processed = number.BoolPointer(number.ToBoolAny(val))
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RefType = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["server_version"].(*string); ok {
		o.ServerVersion = val
	} else if val, ok := kv["server_version"].(string); ok {
		o.ServerVersion = &val
	} else {
		if val, ok := kv["server_version"]; ok {
			if val == nil {
				o.ServerVersion = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ServerVersion = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["state"].(ExportRequestIntegrationsState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			case "processing", "PROCESSING":
				o.State = 2
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			case "processing", "PROCESSING":
				o.State = 2
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
			case "calendar", "CALENDAR":
				o.SystemType = 3
			case "user", "USER":
				o.SystemType = 4
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
			case "calendar", "CALENDAR":
				o.SystemType = 3
			case "user", "USER":
				o.SystemType = 4
			}
		}
	}
	if val, ok := kv["team_id"].(*string); ok {
		o.TeamID = val
	} else if val, ok := kv["team_id"].(string); ok {
		o.TeamID = &val
	} else {
		if val, ok := kv["team_id"]; ok {
			if val == nil {
				o.TeamID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.TeamID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["throttled"].(*bool); ok {
		o.Throttled = val
	} else if val, ok := kv["throttled"].(bool); ok {
		o.Throttled = &val
	} else {
		if val, ok := kv["throttled"]; ok {
			if val == nil {
				o.Throttled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Throttled = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}

	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &ExportRequestIntegrationsThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(ExportRequestIntegrationsThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*ExportRequestIntegrationsThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.ThrottledUntil.Epoch = dt.Epoch
				o.ThrottledUntil.Rfc3339 = dt.Rfc3339
				o.ThrottledUntil.Offset = dt.Offset
			}
		}
	} else {
		o.ThrottledUntil.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["validated"].(*bool); ok {
		o.Validated = val
	} else if val, ok := kv["validated"].(bool); ok {
		o.Validated = &val
	} else {
		if val, ok := kv["validated"]; ok {
			if val == nil {
				o.Validated = nil
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

// ToMap returns the object as a map
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
				o.Epoch = 0
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
				o.Offset = 0
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
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
	return ""
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
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
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
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
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

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportRequest objects are equal
func (o *ExportRequest) IsEqual(other *ExportRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toExportRequestObject(o.CustomerID, false),
		"id":                      toExportRequestObject(o.ID, false),
		"integration_instance_id": toExportRequestObject(o.IntegrationInstanceID, true),
		"integrations":            toExportRequestObject(o.Integrations, false),
		"job_id":                  toExportRequestObject(o.JobID, false),
		"ref_id":                  toExportRequestObject(o.RefID, false),
		"ref_type":                toExportRequestObject(o.RefType, false),
		"reprocess_historical":    toExportRequestObject(o.ReprocessHistorical, false),
		"request_date":            toExportRequestObject(o.RequestDate, false),
		"upload_url":              toExportRequestObject(o.UploadURL, true),
		"uuid":                    toExportRequestObject(o.UUID, false),
		"hashcode":                toExportRequestObject(o.Hashcode, false),
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["integration_instance_id"].(*string); ok {
		o.IntegrationInstanceID = val
	} else if val, ok := kv["integration_instance_id"].(string); ok {
		o.IntegrationInstanceID = &val
	} else {
		if val, ok := kv["integration_instance_id"]; ok {
			if val == nil {
				o.IntegrationInstanceID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationInstanceID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ExportRequestIntegrations
					av.FromMap(bkv)
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
				o.ReprocessHistorical = false
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
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
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
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Integrations)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	args = append(args, o.RequestDate)
	args = append(args, o.UploadURL)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
