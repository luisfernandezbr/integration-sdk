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

	// CodequalityRequestTable is the default table name
	CodequalityRequestTable datamodel.ModelNameType = "agent_codequalityrequest"

	// CodequalityRequestModelName is the model name
	CodequalityRequestModelName datamodel.ModelNameType = "agent.CodequalityRequest"
)

const (
	// CodequalityRequestModelCustomerIDColumn is the column json value customer_id
	CodequalityRequestModelCustomerIDColumn = "customer_id"
	// CodequalityRequestModelIDColumn is the column json value id
	CodequalityRequestModelIDColumn = "id"
	// CodequalityRequestModelIntegrationColumn is the column json value integration
	CodequalityRequestModelIntegrationColumn = "integration"
	// CodequalityRequestModelIntegrationActiveColumn is the column json value active
	CodequalityRequestModelIntegrationActiveColumn = "active"
	// CodequalityRequestModelIntegrationAuthorizationColumn is the column json value authorization
	CodequalityRequestModelIntegrationAuthorizationColumn = "authorization"
	// CodequalityRequestModelIntegrationAuthorizationAccessTokenColumn is the column json value access_token
	CodequalityRequestModelIntegrationAuthorizationAccessTokenColumn = "access_token"
	// CodequalityRequestModelIntegrationAuthorizationRefreshTokenColumn is the column json value refresh_token
	CodequalityRequestModelIntegrationAuthorizationRefreshTokenColumn = "refresh_token"
	// CodequalityRequestModelIntegrationAuthorizationURLColumn is the column json value url
	CodequalityRequestModelIntegrationAuthorizationURLColumn = "url"
	// CodequalityRequestModelIntegrationAuthorizationUsernameColumn is the column json value username
	CodequalityRequestModelIntegrationAuthorizationUsernameColumn = "username"
	// CodequalityRequestModelIntegrationAuthorizationPasswordColumn is the column json value password
	CodequalityRequestModelIntegrationAuthorizationPasswordColumn = "password"
	// CodequalityRequestModelIntegrationAuthorizationAPITokenColumn is the column json value api_token
	CodequalityRequestModelIntegrationAuthorizationAPITokenColumn = "api_token"
	// CodequalityRequestModelIntegrationAuthorizationCollectionNameColumn is the column json value collection_name
	CodequalityRequestModelIntegrationAuthorizationCollectionNameColumn = "collection_name"
	// CodequalityRequestModelIntegrationAuthorizationAPIKeyColumn is the column json value api_key
	CodequalityRequestModelIntegrationAuthorizationAPIKeyColumn = "api_key"
	// CodequalityRequestModelIntegrationAuthorizationAuthorizationColumn is the column json value authorization
	CodequalityRequestModelIntegrationAuthorizationAuthorizationColumn = "authorization"
	// CodequalityRequestModelIntegrationAuthorizationHostnameColumn is the column json value hostname
	CodequalityRequestModelIntegrationAuthorizationHostnameColumn = "hostname"
	// CodequalityRequestModelIntegrationAuthorizationAPIVersionColumn is the column json value api_version
	CodequalityRequestModelIntegrationAuthorizationAPIVersionColumn = "api_version"
	// CodequalityRequestModelIntegrationAuthorizationOrganizationColumn is the column json value organization
	CodequalityRequestModelIntegrationAuthorizationOrganizationColumn = "organization"
	// CodequalityRequestModelIntegrationCreatedByProfileIDColumn is the column json value created_by_profile_id
	CodequalityRequestModelIntegrationCreatedByProfileIDColumn = "created_by_profile_id"
	// CodequalityRequestModelIntegrationCreatedByUserIDColumn is the column json value created_by_user_id
	CodequalityRequestModelIntegrationCreatedByUserIDColumn = "created_by_user_id"
	// CodequalityRequestModelIntegrationCustomerIDColumn is the column json value customer_id
	CodequalityRequestModelIntegrationCustomerIDColumn = "customer_id"
	// CodequalityRequestModelIntegrationEntityErrorsColumn is the column json value entity_errors
	CodequalityRequestModelIntegrationEntityErrorsColumn = "entity_errors"
	// CodequalityRequestModelIntegrationEntityErrorsIDColumn is the column json value id
	CodequalityRequestModelIntegrationEntityErrorsIDColumn = "id"
	// CodequalityRequestModelIntegrationEntityErrorsRefIDColumn is the column json value ref_id
	CodequalityRequestModelIntegrationEntityErrorsRefIDColumn = "ref_id"
	// CodequalityRequestModelIntegrationEntityErrorsErrorColumn is the column json value error
	CodequalityRequestModelIntegrationEntityErrorsErrorColumn = "error"
	// CodequalityRequestModelIntegrationErrorMessageColumn is the column json value error_message
	CodequalityRequestModelIntegrationErrorMessageColumn = "error_message"
	// CodequalityRequestModelIntegrationErroredColumn is the column json value errored
	CodequalityRequestModelIntegrationErroredColumn = "errored"
	// CodequalityRequestModelIntegrationExclusionsColumn is the column json value exclusions
	CodequalityRequestModelIntegrationExclusionsColumn = "exclusions"
	// CodequalityRequestModelIntegrationExportableColumn is the column json value exportable
	CodequalityRequestModelIntegrationExportableColumn = "exportable"
	// CodequalityRequestModelIntegrationIDColumn is the column json value id
	CodequalityRequestModelIntegrationIDColumn = "id"
	// CodequalityRequestModelIntegrationInclusionsColumn is the column json value inclusions
	CodequalityRequestModelIntegrationInclusionsColumn = "inclusions"
	// CodequalityRequestModelIntegrationLastExportCompletedDateColumn is the column json value last_export_completed_date
	CodequalityRequestModelIntegrationLastExportCompletedDateColumn = "last_export_completed_date"
	// CodequalityRequestModelIntegrationLastExportCompletedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationLastExportCompletedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationLastExportCompletedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationLastExportCompletedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationLastExportCompletedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationLastExportCompletedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationLastExportRequestedDateColumn is the column json value last_export_requested_date
	CodequalityRequestModelIntegrationLastExportRequestedDateColumn = "last_export_requested_date"
	// CodequalityRequestModelIntegrationLastExportRequestedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationLastExportRequestedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationLastExportRequestedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationLastExportRequestedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationLastExportRequestedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationLastExportRequestedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationLastProcessingCompletedDateColumn is the column json value last_processing_completed_date
	CodequalityRequestModelIntegrationLastProcessingCompletedDateColumn = "last_processing_completed_date"
	// CodequalityRequestModelIntegrationLastProcessingCompletedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationLastProcessingCompletedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationLastProcessingCompletedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationLastProcessingCompletedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationLastProcessingCompletedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationLastProcessingCompletedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationLastProcessingStartedDateColumn is the column json value last_processing_started_date
	CodequalityRequestModelIntegrationLastProcessingStartedDateColumn = "last_processing_started_date"
	// CodequalityRequestModelIntegrationLastProcessingStartedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationLastProcessingStartedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationLastProcessingStartedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationLastProcessingStartedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationLastProcessingStartedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationLastProcessingStartedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationLocationColumn is the column json value location
	CodequalityRequestModelIntegrationLocationColumn = "location"
	// CodequalityRequestModelIntegrationNameColumn is the column json value name
	CodequalityRequestModelIntegrationNameColumn = "name"
	// CodequalityRequestModelIntegrationOnboardCompletedDateColumn is the column json value onboard_completed_date
	CodequalityRequestModelIntegrationOnboardCompletedDateColumn = "onboard_completed_date"
	// CodequalityRequestModelIntegrationOnboardCompletedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationOnboardCompletedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationOnboardCompletedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationOnboardCompletedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationOnboardCompletedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationOnboardCompletedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationOnboardRequestedDateColumn is the column json value onboard_requested_date
	CodequalityRequestModelIntegrationOnboardRequestedDateColumn = "onboard_requested_date"
	// CodequalityRequestModelIntegrationOnboardRequestedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationOnboardRequestedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationOnboardRequestedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationOnboardRequestedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationOnboardRequestedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationOnboardRequestedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationOnboardingColumn is the column json value onboarding
	CodequalityRequestModelIntegrationOnboardingColumn = "onboarding"
	// CodequalityRequestModelIntegrationOrganizationColumn is the column json value organization
	CodequalityRequestModelIntegrationOrganizationColumn = "organization"
	// CodequalityRequestModelIntegrationProcessedColumn is the column json value processed
	CodequalityRequestModelIntegrationProcessedColumn = "processed"
	// CodequalityRequestModelIntegrationRefIDColumn is the column json value ref_id
	CodequalityRequestModelIntegrationRefIDColumn = "ref_id"
	// CodequalityRequestModelIntegrationRefTypeColumn is the column json value ref_type
	CodequalityRequestModelIntegrationRefTypeColumn = "ref_type"
	// CodequalityRequestModelIntegrationServerVersionColumn is the column json value server_version
	CodequalityRequestModelIntegrationServerVersionColumn = "server_version"
	// CodequalityRequestModelIntegrationStateColumn is the column json value state
	CodequalityRequestModelIntegrationStateColumn = "state"
	// CodequalityRequestModelIntegrationSystemTypeColumn is the column json value system_type
	CodequalityRequestModelIntegrationSystemTypeColumn = "system_type"
	// CodequalityRequestModelIntegrationTeamIDColumn is the column json value team_id
	CodequalityRequestModelIntegrationTeamIDColumn = "team_id"
	// CodequalityRequestModelIntegrationThrottledColumn is the column json value throttled
	CodequalityRequestModelIntegrationThrottledColumn = "throttled"
	// CodequalityRequestModelIntegrationThrottledUntilColumn is the column json value throttled_until
	CodequalityRequestModelIntegrationThrottledUntilColumn = "throttled_until"
	// CodequalityRequestModelIntegrationThrottledUntilEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationThrottledUntilEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationThrottledUntilOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationThrottledUntilOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationThrottledUntilRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationThrottledUntilRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationValidatedColumn is the column json value validated
	CodequalityRequestModelIntegrationValidatedColumn = "validated"
	// CodequalityRequestModelIntegrationValidatedDateColumn is the column json value validated_date
	CodequalityRequestModelIntegrationValidatedDateColumn = "validated_date"
	// CodequalityRequestModelIntegrationValidatedDateEpochColumn is the column json value epoch
	CodequalityRequestModelIntegrationValidatedDateEpochColumn = "epoch"
	// CodequalityRequestModelIntegrationValidatedDateOffsetColumn is the column json value offset
	CodequalityRequestModelIntegrationValidatedDateOffsetColumn = "offset"
	// CodequalityRequestModelIntegrationValidatedDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelIntegrationValidatedDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelIntegrationValidationMessageColumn is the column json value validation_message
	CodequalityRequestModelIntegrationValidationMessageColumn = "validation_message"
	// CodequalityRequestModelRefIDColumn is the column json value ref_id
	CodequalityRequestModelRefIDColumn = "ref_id"
	// CodequalityRequestModelRefTypeColumn is the column json value ref_type
	CodequalityRequestModelRefTypeColumn = "ref_type"
	// CodequalityRequestModelRequestDateColumn is the column json value request_date
	CodequalityRequestModelRequestDateColumn = "request_date"
	// CodequalityRequestModelRequestDateEpochColumn is the column json value epoch
	CodequalityRequestModelRequestDateEpochColumn = "epoch"
	// CodequalityRequestModelRequestDateOffsetColumn is the column json value offset
	CodequalityRequestModelRequestDateOffsetColumn = "offset"
	// CodequalityRequestModelRequestDateRfc3339Column is the column json value rfc3339
	CodequalityRequestModelRequestDateRfc3339Column = "rfc3339"
	// CodequalityRequestModelUUIDColumn is the column json value uuid
	CodequalityRequestModelUUIDColumn = "uuid"
)

// CodequalityRequestIntegrationAuthorization represents the object structure for authorization
type CodequalityRequestIntegrationAuthorization struct {
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

func toCodequalityRequestIntegrationAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toCodequalityRequestIntegrationAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toCodequalityRequestIntegrationAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toCodequalityRequestIntegrationAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toCodequalityRequestIntegrationAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toCodequalityRequestIntegrationAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toCodequalityRequestIntegrationAuthorizationObject(o.APIToken, true),
		// CollectionName Collection name for instance, if relevant
		"collection_name": toCodequalityRequestIntegrationAuthorizationObject(o.CollectionName, true),
		// APIKey API Key for instance, if relevant
		"api_key": toCodequalityRequestIntegrationAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toCodequalityRequestIntegrationAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toCodequalityRequestIntegrationAuthorizationObject(o.Hostname, true),
		// APIVersion the api version of the integration
		"api_version": toCodequalityRequestIntegrationAuthorizationObject(o.APIVersion, true),
		// Organization Organization for instance, if relevant
		"organization": toCodequalityRequestIntegrationAuthorizationObject(o.Organization, true),
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

// CodequalityRequestIntegrationEntityErrors represents the object structure for entity_errors
type CodequalityRequestIntegrationEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
}

func toCodequalityRequestIntegrationEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toCodequalityRequestIntegrationEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toCodequalityRequestIntegrationEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export
		"error": toCodequalityRequestIntegrationEntityErrorsObject(o.Error, false),
	}
}

func (o *CodequalityRequestIntegrationEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationEntityErrors) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLastExportCompletedDate represents the object structure for last_export_completed_date
type CodequalityRequestIntegrationLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLastExportRequestedDate represents the object structure for last_export_requested_date
type CodequalityRequestIntegrationLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLastProcessingCompletedDate represents the object structure for last_processing_completed_date
type CodequalityRequestIntegrationLastProcessingCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationLastProcessingCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationLastProcessingCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationLastProcessingCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationLastProcessingCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationLastProcessingCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationLastProcessingCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationLastProcessingCompletedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLastProcessingStartedDate represents the object structure for last_processing_started_date
type CodequalityRequestIntegrationLastProcessingStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationLastProcessingStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationLastProcessingStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationLastProcessingStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationLastProcessingStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationLastProcessingStartedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationLastProcessingStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationLastProcessingStartedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationLocation is the enumeration type for location
type CodequalityRequestIntegrationLocation int32

// UnmarshalBSONValue for unmarshaling value
func (v *CodequalityRequestIntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CodequalityRequestIntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = CodequalityRequestIntegrationLocation(0)
		case "CLOUD":
			*v = CodequalityRequestIntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CodequalityRequestIntegrationLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v CodequalityRequestIntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

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

// FromInterface for decoding from an interface
func (v *CodequalityRequestIntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CodequalityRequestIntegrationLocation:
		*v = val
	case int32:
		*v = CodequalityRequestIntegrationLocation(int32(val))
	case int:
		*v = CodequalityRequestIntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = CodequalityRequestIntegrationLocation(0)
		case "CLOUD":
			*v = CodequalityRequestIntegrationLocation(1)
		}
	}
	return nil
}

const (
	// CodequalityRequestIntegrationLocationPrivate is the enumeration value for private
	CodequalityRequestIntegrationLocationPrivate CodequalityRequestIntegrationLocation = 0
	// CodequalityRequestIntegrationLocationCloud is the enumeration value for cloud
	CodequalityRequestIntegrationLocationCloud CodequalityRequestIntegrationLocation = 1
)

// CodequalityRequestIntegrationOnboardCompletedDate represents the object structure for onboard_completed_date
type CodequalityRequestIntegrationOnboardCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationOnboardCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationOnboardCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationOnboardCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationOnboardCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationOnboardCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationOnboardCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationOnboardCompletedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationOnboardRequestedDate represents the object structure for onboard_requested_date
type CodequalityRequestIntegrationOnboardRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationOnboardRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationOnboardRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationOnboardRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationOnboardRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationOnboardRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationOnboardRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationOnboardRequestedDate) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationState is the enumeration type for state
type CodequalityRequestIntegrationState int32

// UnmarshalBSONValue for unmarshaling value
func (v *CodequalityRequestIntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CodequalityRequestIntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = CodequalityRequestIntegrationState(0)
		case "EXPORTING":
			*v = CodequalityRequestIntegrationState(1)
		case "PROCESSING":
			*v = CodequalityRequestIntegrationState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CodequalityRequestIntegrationState) UnmarshalJSON(buf []byte) error {
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
func (v CodequalityRequestIntegrationState) MarshalJSON() ([]byte, error) {
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

// String returns the string value for IntegrationState
func (v CodequalityRequestIntegrationState) String() string {
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
func (v *CodequalityRequestIntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CodequalityRequestIntegrationState:
		*v = val
	case int32:
		*v = CodequalityRequestIntegrationState(int32(val))
	case int:
		*v = CodequalityRequestIntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = CodequalityRequestIntegrationState(0)
		case "EXPORTING":
			*v = CodequalityRequestIntegrationState(1)
		case "PROCESSING":
			*v = CodequalityRequestIntegrationState(2)
		}
	}
	return nil
}

const (
	// CodequalityRequestIntegrationStateIdle is the enumeration value for idle
	CodequalityRequestIntegrationStateIdle CodequalityRequestIntegrationState = 0
	// CodequalityRequestIntegrationStateExporting is the enumeration value for exporting
	CodequalityRequestIntegrationStateExporting CodequalityRequestIntegrationState = 1
	// CodequalityRequestIntegrationStateProcessing is the enumeration value for processing
	CodequalityRequestIntegrationStateProcessing CodequalityRequestIntegrationState = 2
)

// CodequalityRequestIntegrationSystemType is the enumeration type for system_type
type CodequalityRequestIntegrationSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *CodequalityRequestIntegrationSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CodequalityRequestIntegrationSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = CodequalityRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = CodequalityRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = CodequalityRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = CodequalityRequestIntegrationSystemType(3)
		case "USER":
			*v = CodequalityRequestIntegrationSystemType(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CodequalityRequestIntegrationSystemType) UnmarshalJSON(buf []byte) error {
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
func (v CodequalityRequestIntegrationSystemType) MarshalJSON() ([]byte, error) {
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

// String returns the string value for IntegrationSystemType
func (v CodequalityRequestIntegrationSystemType) String() string {
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
func (v *CodequalityRequestIntegrationSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case CodequalityRequestIntegrationSystemType:
		*v = val
	case int32:
		*v = CodequalityRequestIntegrationSystemType(int32(val))
	case int:
		*v = CodequalityRequestIntegrationSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = CodequalityRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = CodequalityRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = CodequalityRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = CodequalityRequestIntegrationSystemType(3)
		case "USER":
			*v = CodequalityRequestIntegrationSystemType(4)
		}
	}
	return nil
}

const (
	// CodequalityRequestIntegrationSystemTypeWork is the enumeration value for work
	CodequalityRequestIntegrationSystemTypeWork CodequalityRequestIntegrationSystemType = 0
	// CodequalityRequestIntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	CodequalityRequestIntegrationSystemTypeSourcecode CodequalityRequestIntegrationSystemType = 1
	// CodequalityRequestIntegrationSystemTypeCodequality is the enumeration value for codequality
	CodequalityRequestIntegrationSystemTypeCodequality CodequalityRequestIntegrationSystemType = 2
	// CodequalityRequestIntegrationSystemTypeCalendar is the enumeration value for calendar
	CodequalityRequestIntegrationSystemTypeCalendar CodequalityRequestIntegrationSystemType = 3
	// CodequalityRequestIntegrationSystemTypeUser is the enumeration value for user
	CodequalityRequestIntegrationSystemTypeUser CodequalityRequestIntegrationSystemType = 4
)

// CodequalityRequestIntegrationThrottledUntil represents the object structure for throttled_until
type CodequalityRequestIntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *CodequalityRequestIntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityRequestIntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

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

// CodequalityRequestIntegrationValidatedDate represents the object structure for validated_date
type CodequalityRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestIntegrationValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestIntegrationValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestIntegrationValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestIntegrationValidatedDateObject(o.Rfc3339, false),
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

// CodequalityRequestIntegration represents the object structure for integration
type CodequalityRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization CodequalityRequestIntegrationAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []CodequalityRequestIntegrationEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" codec:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Exportable a flag to indicate if the integration is ready for export
	Exportable bool `json:"exportable" codec:"exportable" bson:"exportable" yaml:"exportable" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Inclusions The inclusion list for this integration
	Inclusions []string `json:"inclusions" codec:"inclusions" bson:"inclusions" yaml:"inclusions" faker:"-"`
	// LastExportCompletedDate when the export response was received (set by the backend)
	LastExportCompletedDate CodequalityRequestIntegrationLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made (set by the backend)
	LastExportRequestedDate CodequalityRequestIntegrationLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingCompletedDate when the processing completes (set by the pipeline)
	LastProcessingCompletedDate CodequalityRequestIntegrationLastProcessingCompletedDate `json:"last_processing_completed_date" codec:"last_processing_completed_date" bson:"last_processing_completed_date" yaml:"last_processing_completed_date" faker:"-"`
	// LastProcessingStartedDate when the processing starts (set by the pipeline)
	LastProcessingStartedDate CodequalityRequestIntegrationLastProcessingStartedDate `json:"last_processing_started_date" codec:"last_processing_started_date" bson:"last_processing_started_date" yaml:"last_processing_started_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location CodequalityRequestIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OnboardCompletedDate when the last request for integration metadata was finished
	OnboardCompletedDate CodequalityRequestIntegrationOnboardCompletedDate `json:"onboard_completed_date" codec:"onboard_completed_date" bson:"onboard_completed_date" yaml:"onboard_completed_date" faker:"-"`
	// OnboardRequestedDate when the last request for integration metadata was made
	OnboardRequestedDate CodequalityRequestIntegrationOnboardRequestedDate `json:"onboard_requested_date" codec:"onboard_requested_date" bson:"onboard_requested_date" yaml:"onboard_requested_date" faker:"-"`
	// Onboarding true if the agent is fetching metadata for the integration
	Onboarding bool `json:"onboarding" codec:"onboarding" bson:"onboarding" yaml:"onboarding" faker:"-"`
	// Organization The origanization authorized. Used for azure integrations
	Organization *string `json:"organization,omitempty" codec:"organization,omitempty" bson:"organization" yaml:"organization,omitempty" faker:"-"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ServerVersion the server version for this integration
	ServerVersion *string `json:"server_version,omitempty" codec:"server_version,omitempty" bson:"server_version" yaml:"server_version,omitempty" faker:"-"`
	// State the current state of the integration
	State CodequalityRequestIntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType CodequalityRequestIntegrationSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
	TeamID *string `json:"team_id,omitempty" codec:"team_id,omitempty" bson:"team_id" yaml:"team_id,omitempty" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *CodequalityRequestIntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate CodequalityRequestIntegrationValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toCodequalityRequestIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestIntegration:
		return v.ToMap()

	case CodequalityRequestIntegrationAuthorization:
		return v.ToMap()

	case []CodequalityRequestIntegrationEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case CodequalityRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationLocation:
		return v.String()

	case CodequalityRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	case CodequalityRequestIntegrationState:
		return v.String()

	case CodequalityRequestIntegrationSystemType:
		return v.String()

	case *CodequalityRequestIntegrationThrottledUntil:
		return v.ToMap()

	case CodequalityRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toCodequalityRequestIntegrationObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toCodequalityRequestIntegrationObject(o.Authorization, false),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toCodequalityRequestIntegrationObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toCodequalityRequestIntegrationObject(o.CreatedByUserID, true),
		// CustomerID the customer id for the model instance
		"customer_id": toCodequalityRequestIntegrationObject(o.CustomerID, false),
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toCodequalityRequestIntegrationObject(o.EntityErrors, false),
		// ErrorMessage The error message from an export run
		"error_message": toCodequalityRequestIntegrationObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent
		"errored": toCodequalityRequestIntegrationObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toCodequalityRequestIntegrationObject(o.Exclusions, false),
		// Exportable a flag to indicate if the integration is ready for export
		"exportable": toCodequalityRequestIntegrationObject(o.Exportable, false),
		// ID the primary key for the model instance
		"id": toCodequalityRequestIntegrationObject(o.ID, false),
		// Inclusions The inclusion list for this integration
		"inclusions": toCodequalityRequestIntegrationObject(o.Inclusions, false),
		// LastExportCompletedDate when the export response was received (set by the backend)
		"last_export_completed_date": toCodequalityRequestIntegrationObject(o.LastExportCompletedDate, false),
		// LastExportRequestedDate when the export request was made (set by the backend)
		"last_export_requested_date": toCodequalityRequestIntegrationObject(o.LastExportRequestedDate, false),
		// LastProcessingCompletedDate when the processing completes (set by the pipeline)
		"last_processing_completed_date": toCodequalityRequestIntegrationObject(o.LastProcessingCompletedDate, false),
		// LastProcessingStartedDate when the processing starts (set by the pipeline)
		"last_processing_started_date": toCodequalityRequestIntegrationObject(o.LastProcessingStartedDate, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toCodequalityRequestIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toCodequalityRequestIntegrationObject(o.Name, false),
		// OnboardCompletedDate when the last request for integration metadata was finished
		"onboard_completed_date": toCodequalityRequestIntegrationObject(o.OnboardCompletedDate, false),
		// OnboardRequestedDate when the last request for integration metadata was made
		"onboard_requested_date": toCodequalityRequestIntegrationObject(o.OnboardRequestedDate, false),
		// Onboarding true if the agent is fetching metadata for the integration
		"onboarding": toCodequalityRequestIntegrationObject(o.Onboarding, false),
		// Organization The origanization authorized. Used for azure integrations
		"organization": toCodequalityRequestIntegrationObject(o.Organization, true),
		// Processed If the integration has been processed at least once
		"processed": toCodequalityRequestIntegrationObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toCodequalityRequestIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toCodequalityRequestIntegrationObject(o.RefType, false),
		// ServerVersion the server version for this integration
		"server_version": toCodequalityRequestIntegrationObject(o.ServerVersion, true),
		// State the current state of the integration
		"state": toCodequalityRequestIntegrationObject(o.State, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toCodequalityRequestIntegrationObject(o.SystemType, false),
		// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
		"team_id": toCodequalityRequestIntegrationObject(o.TeamID, true),
		// Throttled Set to true when integration is throttled.
		"throttled": toCodequalityRequestIntegrationObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toCodequalityRequestIntegrationObject(o.ThrottledUntil, true),
		// Validated If the validation has been run against this instance
		"validated": toCodequalityRequestIntegrationObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toCodequalityRequestIntegrationObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toCodequalityRequestIntegrationObject(o.ValidationMessage, true),
	}
}

func (o *CodequalityRequestIntegration) setDefaults(frommap bool) {

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
				o.Active = false
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

		o.EntityErrors = make([]CodequalityRequestIntegrationEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]CodequalityRequestIntegrationEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*CodequalityRequestIntegrationEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(CodequalityRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm CodequalityRequestIntegrationEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av CodequalityRequestIntegrationEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(CodequalityRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm CodequalityRequestIntegrationEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := CodequalityRequestIntegrationEntityErrors{}
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
							var fm CodequalityRequestIntegrationEntityErrors
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

	if val, ok := kv["last_export_completed_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportCompletedDate.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegrationLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationLastExportCompletedDate); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationLastExportRequestedDate); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationLastProcessingCompletedDate); ok {
			// struct
			o.LastProcessingCompletedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationLastProcessingCompletedDate); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationLastProcessingStartedDate); ok {
			// struct
			o.LastProcessingStartedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationLastProcessingStartedDate); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationOnboardCompletedDate); ok {
			// struct
			o.OnboardCompletedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationOnboardCompletedDate); ok {
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
		} else if sv, ok := val.(CodequalityRequestIntegrationOnboardRequestedDate); ok {
			// struct
			o.OnboardRequestedDate = sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationOnboardRequestedDate); ok {
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
	if val, ok := kv["state"].(CodequalityRequestIntegrationState); ok {
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
		o.ThrottledUntil = &CodequalityRequestIntegrationThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestIntegrationThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*CodequalityRequestIntegrationThrottledUntil); ok {
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

// CodequalityRequestRequestDate represents the object structure for request_date
type CodequalityRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *CodequalityRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityRequestRequestDateObject(o.Rfc3339, false),
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

// CodequalityRequest an agent action to request adding new codequality entities
type CodequalityRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration CodequalityRequestIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate CodequalityRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CodequalityRequest)(nil)

func toCodequalityRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityRequest:
		return v.ToMap()

	case CodequalityRequestIntegration:
		return v.ToMap()

	case CodequalityRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of CodequalityRequest
func (o *CodequalityRequest) String() string {
	return fmt.Sprintf("agent.CodequalityRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CodequalityRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *CodequalityRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CodequalityRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *CodequalityRequest) GetModelName() datamodel.ModelNameType {
	return CodequalityRequestModelName
}

// NewCodequalityRequestID provides a template for generating an ID field for CodequalityRequest
func NewCodequalityRequestID(customerID string, refType string, refID string) string {
	return hash.Values("CodequalityRequest", customerID, refType, refID)
}

func (o *CodequalityRequest) setDefaults(frommap bool) {
	if o.Integration.EntityErrors == nil {
		o.Integration.EntityErrors = make([]CodequalityRequestIntegrationEntityErrors, 0)
	}
	if o.Integration.Exclusions == nil {
		o.Integration.Exclusions = make([]string, 0)
	}
	if o.Integration.Inclusions == nil {
		o.Integration.Inclusions = make([]string, 0)
	}
	if o.Integration.ThrottledUntil == nil {
		o.Integration.ThrottledUntil = &CodequalityRequestIntegrationThrottledUntil{}
	}

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
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CodequalityRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *CodequalityRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CodequalityRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CodequalityRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CodequalityRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CodequalityRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *CodequalityRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CodequalityRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CodequalityRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
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

// Stringify returns the object in JSON format as a string
func (o *CodequalityRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *CodequalityRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two CodequalityRequest objects are equal
func (o *CodequalityRequest) IsEqual(other *CodequalityRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toCodequalityRequestObject(o.CustomerID, false),
		"id":           toCodequalityRequestObject(o.ID, false),
		"integration":  toCodequalityRequestObject(o.Integration, false),
		"ref_id":       toCodequalityRequestObject(o.RefID, false),
		"ref_type":     toCodequalityRequestObject(o.RefType, false),
		"request_date": toCodequalityRequestObject(o.RequestDate, false),
		"uuid":         toCodequalityRequestObject(o.UUID, false),
		"hashcode":     toCodequalityRequestObject(o.Hashcode, false),
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

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(CodequalityRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*CodequalityRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		}
	} else {
		o.RequestDate.FromMap(map[string]interface{}{})
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
func (o *CodequalityRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
