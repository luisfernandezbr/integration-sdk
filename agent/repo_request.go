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

	// RepoRequestTable is the default table name
	RepoRequestTable datamodel.ModelNameType = "agent_reporequest"

	// RepoRequestModelName is the model name
	RepoRequestModelName datamodel.ModelNameType = "agent.RepoRequest"
)

const (
	// RepoRequestModelCustomerIDColumn is the column json value customer_id
	RepoRequestModelCustomerIDColumn = "customer_id"
	// RepoRequestModelIDColumn is the column json value id
	RepoRequestModelIDColumn = "id"
	// RepoRequestModelIntegrationColumn is the column json value integration
	RepoRequestModelIntegrationColumn = "integration"
	// RepoRequestModelIntegrationActiveColumn is the column json value active
	RepoRequestModelIntegrationActiveColumn = "active"
	// RepoRequestModelIntegrationAuthorizationColumn is the column json value authorization
	RepoRequestModelIntegrationAuthorizationColumn = "authorization"
	// RepoRequestModelIntegrationAuthorizationAccessTokenColumn is the column json value access_token
	RepoRequestModelIntegrationAuthorizationAccessTokenColumn = "access_token"
	// RepoRequestModelIntegrationAuthorizationRefreshTokenColumn is the column json value refresh_token
	RepoRequestModelIntegrationAuthorizationRefreshTokenColumn = "refresh_token"
	// RepoRequestModelIntegrationAuthorizationURLColumn is the column json value url
	RepoRequestModelIntegrationAuthorizationURLColumn = "url"
	// RepoRequestModelIntegrationAuthorizationUsernameColumn is the column json value username
	RepoRequestModelIntegrationAuthorizationUsernameColumn = "username"
	// RepoRequestModelIntegrationAuthorizationPasswordColumn is the column json value password
	RepoRequestModelIntegrationAuthorizationPasswordColumn = "password"
	// RepoRequestModelIntegrationAuthorizationAPITokenColumn is the column json value api_token
	RepoRequestModelIntegrationAuthorizationAPITokenColumn = "api_token"
	// RepoRequestModelIntegrationAuthorizationCollectionNameColumn is the column json value collection_name
	RepoRequestModelIntegrationAuthorizationCollectionNameColumn = "collection_name"
	// RepoRequestModelIntegrationAuthorizationAPIKeyColumn is the column json value api_key
	RepoRequestModelIntegrationAuthorizationAPIKeyColumn = "api_key"
	// RepoRequestModelIntegrationAuthorizationAuthorizationColumn is the column json value authorization
	RepoRequestModelIntegrationAuthorizationAuthorizationColumn = "authorization"
	// RepoRequestModelIntegrationAuthorizationHostnameColumn is the column json value hostname
	RepoRequestModelIntegrationAuthorizationHostnameColumn = "hostname"
	// RepoRequestModelIntegrationAuthorizationAPIVersionColumn is the column json value api_version
	RepoRequestModelIntegrationAuthorizationAPIVersionColumn = "api_version"
	// RepoRequestModelIntegrationAuthorizationOrganizationColumn is the column json value organization
	RepoRequestModelIntegrationAuthorizationOrganizationColumn = "organization"
	// RepoRequestModelIntegrationConfigColumn is the column json value config
	RepoRequestModelIntegrationConfigColumn = "config"
	// RepoRequestModelIntegrationCreatedByProfileIDColumn is the column json value created_by_profile_id
	RepoRequestModelIntegrationCreatedByProfileIDColumn = "created_by_profile_id"
	// RepoRequestModelIntegrationCreatedByUserIDColumn is the column json value created_by_user_id
	RepoRequestModelIntegrationCreatedByUserIDColumn = "created_by_user_id"
	// RepoRequestModelIntegrationCustomerIDColumn is the column json value customer_id
	RepoRequestModelIntegrationCustomerIDColumn = "customer_id"
	// RepoRequestModelIntegrationEntityErrorsColumn is the column json value entity_errors
	RepoRequestModelIntegrationEntityErrorsColumn = "entity_errors"
	// RepoRequestModelIntegrationEntityErrorsIDColumn is the column json value id
	RepoRequestModelIntegrationEntityErrorsIDColumn = "id"
	// RepoRequestModelIntegrationEntityErrorsRefIDColumn is the column json value ref_id
	RepoRequestModelIntegrationEntityErrorsRefIDColumn = "ref_id"
	// RepoRequestModelIntegrationEntityErrorsErrorColumn is the column json value error
	RepoRequestModelIntegrationEntityErrorsErrorColumn = "error"
	// RepoRequestModelIntegrationErrorMessageColumn is the column json value error_message
	RepoRequestModelIntegrationErrorMessageColumn = "error_message"
	// RepoRequestModelIntegrationErroredColumn is the column json value errored
	RepoRequestModelIntegrationErroredColumn = "errored"
	// RepoRequestModelIntegrationExclusionsColumn is the column json value exclusions
	RepoRequestModelIntegrationExclusionsColumn = "exclusions"
	// RepoRequestModelIntegrationExportableColumn is the column json value exportable
	RepoRequestModelIntegrationExportableColumn = "exportable"
	// RepoRequestModelIntegrationIDColumn is the column json value id
	RepoRequestModelIntegrationIDColumn = "id"
	// RepoRequestModelIntegrationInclusionsColumn is the column json value inclusions
	RepoRequestModelIntegrationInclusionsColumn = "inclusions"
	// RepoRequestModelIntegrationIntegrationInstanceIDColumn is the column json value integration_instance_id
	RepoRequestModelIntegrationIntegrationInstanceIDColumn = "integration_instance_id"
	// RepoRequestModelIntegrationIntervalColumn is the column json value interval
	RepoRequestModelIntegrationIntervalColumn = "interval"
	// RepoRequestModelIntegrationLastExportCompletedDateColumn is the column json value last_export_completed_date
	RepoRequestModelIntegrationLastExportCompletedDateColumn = "last_export_completed_date"
	// RepoRequestModelIntegrationLastExportCompletedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationLastExportCompletedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationLastExportCompletedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationLastExportCompletedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationLastExportCompletedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationLastExportCompletedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationLastExportRequestedDateColumn is the column json value last_export_requested_date
	RepoRequestModelIntegrationLastExportRequestedDateColumn = "last_export_requested_date"
	// RepoRequestModelIntegrationLastExportRequestedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationLastExportRequestedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationLastExportRequestedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationLastExportRequestedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationLastExportRequestedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationLastExportRequestedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationLastProcessingCompletedDateColumn is the column json value last_processing_completed_date
	RepoRequestModelIntegrationLastProcessingCompletedDateColumn = "last_processing_completed_date"
	// RepoRequestModelIntegrationLastProcessingCompletedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationLastProcessingCompletedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationLastProcessingCompletedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationLastProcessingCompletedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationLastProcessingCompletedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationLastProcessingCompletedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationLastProcessingStartedDateColumn is the column json value last_processing_started_date
	RepoRequestModelIntegrationLastProcessingStartedDateColumn = "last_processing_started_date"
	// RepoRequestModelIntegrationLastProcessingStartedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationLastProcessingStartedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationLastProcessingStartedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationLastProcessingStartedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationLastProcessingStartedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationLastProcessingStartedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationLocationColumn is the column json value location
	RepoRequestModelIntegrationLocationColumn = "location"
	// RepoRequestModelIntegrationNameColumn is the column json value name
	RepoRequestModelIntegrationNameColumn = "name"
	// RepoRequestModelIntegrationOnboardCompletedDateColumn is the column json value onboard_completed_date
	RepoRequestModelIntegrationOnboardCompletedDateColumn = "onboard_completed_date"
	// RepoRequestModelIntegrationOnboardCompletedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationOnboardCompletedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationOnboardCompletedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationOnboardCompletedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationOnboardCompletedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationOnboardCompletedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationOnboardRequestedDateColumn is the column json value onboard_requested_date
	RepoRequestModelIntegrationOnboardRequestedDateColumn = "onboard_requested_date"
	// RepoRequestModelIntegrationOnboardRequestedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationOnboardRequestedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationOnboardRequestedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationOnboardRequestedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationOnboardRequestedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationOnboardRequestedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationOnboardingColumn is the column json value onboarding
	RepoRequestModelIntegrationOnboardingColumn = "onboarding"
	// RepoRequestModelIntegrationOrganizationColumn is the column json value organization
	RepoRequestModelIntegrationOrganizationColumn = "organization"
	// RepoRequestModelIntegrationPausedColumn is the column json value paused
	RepoRequestModelIntegrationPausedColumn = "paused"
	// RepoRequestModelIntegrationProcessedColumn is the column json value processed
	RepoRequestModelIntegrationProcessedColumn = "processed"
	// RepoRequestModelIntegrationRefIDColumn is the column json value ref_id
	RepoRequestModelIntegrationRefIDColumn = "ref_id"
	// RepoRequestModelIntegrationRefTypeColumn is the column json value ref_type
	RepoRequestModelIntegrationRefTypeColumn = "ref_type"
	// RepoRequestModelIntegrationServerVersionColumn is the column json value server_version
	RepoRequestModelIntegrationServerVersionColumn = "server_version"
	// RepoRequestModelIntegrationStateColumn is the column json value state
	RepoRequestModelIntegrationStateColumn = "state"
	// RepoRequestModelIntegrationSystemTypeColumn is the column json value system_type
	RepoRequestModelIntegrationSystemTypeColumn = "system_type"
	// RepoRequestModelIntegrationTeamIDColumn is the column json value team_id
	RepoRequestModelIntegrationTeamIDColumn = "team_id"
	// RepoRequestModelIntegrationThrottledColumn is the column json value throttled
	RepoRequestModelIntegrationThrottledColumn = "throttled"
	// RepoRequestModelIntegrationThrottledUntilColumn is the column json value throttled_until
	RepoRequestModelIntegrationThrottledUntilColumn = "throttled_until"
	// RepoRequestModelIntegrationThrottledUntilEpochColumn is the column json value epoch
	RepoRequestModelIntegrationThrottledUntilEpochColumn = "epoch"
	// RepoRequestModelIntegrationThrottledUntilOffsetColumn is the column json value offset
	RepoRequestModelIntegrationThrottledUntilOffsetColumn = "offset"
	// RepoRequestModelIntegrationThrottledUntilRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationThrottledUntilRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationValidatedColumn is the column json value validated
	RepoRequestModelIntegrationValidatedColumn = "validated"
	// RepoRequestModelIntegrationValidatedDateColumn is the column json value validated_date
	RepoRequestModelIntegrationValidatedDateColumn = "validated_date"
	// RepoRequestModelIntegrationValidatedDateEpochColumn is the column json value epoch
	RepoRequestModelIntegrationValidatedDateEpochColumn = "epoch"
	// RepoRequestModelIntegrationValidatedDateOffsetColumn is the column json value offset
	RepoRequestModelIntegrationValidatedDateOffsetColumn = "offset"
	// RepoRequestModelIntegrationValidatedDateRfc3339Column is the column json value rfc3339
	RepoRequestModelIntegrationValidatedDateRfc3339Column = "rfc3339"
	// RepoRequestModelIntegrationValidationMessageColumn is the column json value validation_message
	RepoRequestModelIntegrationValidationMessageColumn = "validation_message"
	// RepoRequestModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	RepoRequestModelIntegrationInstanceIDColumn = "integration_instance_id"
	// RepoRequestModelRefIDColumn is the column json value ref_id
	RepoRequestModelRefIDColumn = "ref_id"
	// RepoRequestModelRefTypeColumn is the column json value ref_type
	RepoRequestModelRefTypeColumn = "ref_type"
	// RepoRequestModelRequestDateColumn is the column json value request_date
	RepoRequestModelRequestDateColumn = "request_date"
	// RepoRequestModelRequestDateEpochColumn is the column json value epoch
	RepoRequestModelRequestDateEpochColumn = "epoch"
	// RepoRequestModelRequestDateOffsetColumn is the column json value offset
	RepoRequestModelRequestDateOffsetColumn = "offset"
	// RepoRequestModelRequestDateRfc3339Column is the column json value rfc3339
	RepoRequestModelRequestDateRfc3339Column = "rfc3339"
	// RepoRequestModelUUIDColumn is the column json value uuid
	RepoRequestModelUUIDColumn = "uuid"
)

// RepoRequestIntegrationAuthorization represents the object structure for authorization
type RepoRequestIntegrationAuthorization struct {
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

func toRepoRequestIntegrationAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toRepoRequestIntegrationAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toRepoRequestIntegrationAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toRepoRequestIntegrationAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toRepoRequestIntegrationAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toRepoRequestIntegrationAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toRepoRequestIntegrationAuthorizationObject(o.APIToken, true),
		// CollectionName Collection name for instance, if relevant
		"collection_name": toRepoRequestIntegrationAuthorizationObject(o.CollectionName, true),
		// APIKey API Key for instance, if relevant
		"api_key": toRepoRequestIntegrationAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toRepoRequestIntegrationAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toRepoRequestIntegrationAuthorizationObject(o.Hostname, true),
		// APIVersion the api version of the integration
		"api_version": toRepoRequestIntegrationAuthorizationObject(o.APIVersion, true),
		// Organization Organization for instance, if relevant
		"organization": toRepoRequestIntegrationAuthorizationObject(o.Organization, true),
	}
}

func (o *RepoRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationEntityErrors represents the object structure for entity_errors
type RepoRequestIntegrationEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
}

func toRepoRequestIntegrationEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toRepoRequestIntegrationEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toRepoRequestIntegrationEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export
		"error": toRepoRequestIntegrationEntityErrorsObject(o.Error, false),
	}
}

func (o *RepoRequestIntegrationEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationEntityErrors) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationLastExportCompletedDate represents the object structure for last_export_completed_date
type RepoRequestIntegrationLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationLastExportRequestedDate represents the object structure for last_export_requested_date
type RepoRequestIntegrationLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationLastProcessingCompletedDate represents the object structure for last_processing_completed_date
type RepoRequestIntegrationLastProcessingCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationLastProcessingCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationLastProcessingCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationLastProcessingCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationLastProcessingCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationLastProcessingCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationLastProcessingCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationLastProcessingCompletedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationLastProcessingStartedDate represents the object structure for last_processing_started_date
type RepoRequestIntegrationLastProcessingStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationLastProcessingStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationLastProcessingStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationLastProcessingStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationLastProcessingStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationLastProcessingStartedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationLastProcessingStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationLastProcessingStartedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationLocation is the enumeration type for location
type RepoRequestIntegrationLocation int32

// toRepoRequestIntegrationLocationPointer is the enumeration pointer type for location
func toRepoRequestIntegrationLocationPointer(v int32) *RepoRequestIntegrationLocation {
	nv := RepoRequestIntegrationLocation(v)
	return &nv
}

// toRepoRequestIntegrationLocationEnum is the enumeration pointer wrapper for location
func toRepoRequestIntegrationLocationEnum(v *RepoRequestIntegrationLocation) string {
	if v == nil {
		return toRepoRequestIntegrationLocationPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *RepoRequestIntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoRequestIntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = RepoRequestIntegrationLocation(0)
		case "CLOUD":
			*v = RepoRequestIntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *RepoRequestIntegrationLocation) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "PRIVATE":
		*v = 0
	case "CLOUD":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoRequestIntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationLocation
func (v RepoRequestIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *RepoRequestIntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoRequestIntegrationLocation:
		*v = val
	case int32:
		*v = RepoRequestIntegrationLocation(int32(val))
	case int:
		*v = RepoRequestIntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = RepoRequestIntegrationLocation(0)
		case "CLOUD":
			*v = RepoRequestIntegrationLocation(1)
		}
	}
	return nil
}

const (
	// RepoRequestIntegrationLocationPrivate is the enumeration value for private
	RepoRequestIntegrationLocationPrivate RepoRequestIntegrationLocation = 0
	// RepoRequestIntegrationLocationCloud is the enumeration value for cloud
	RepoRequestIntegrationLocationCloud RepoRequestIntegrationLocation = 1
)

// RepoRequestIntegrationOnboardCompletedDate represents the object structure for onboard_completed_date
type RepoRequestIntegrationOnboardCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationOnboardCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationOnboardCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationOnboardCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationOnboardCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationOnboardCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationOnboardCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationOnboardCompletedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationOnboardRequestedDate represents the object structure for onboard_requested_date
type RepoRequestIntegrationOnboardRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationOnboardRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationOnboardRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationOnboardRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationOnboardRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationOnboardRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationOnboardRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationOnboardRequestedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationState is the enumeration type for state
type RepoRequestIntegrationState int32

// toRepoRequestIntegrationStatePointer is the enumeration pointer type for state
func toRepoRequestIntegrationStatePointer(v int32) *RepoRequestIntegrationState {
	nv := RepoRequestIntegrationState(v)
	return &nv
}

// toRepoRequestIntegrationStateEnum is the enumeration pointer wrapper for state
func toRepoRequestIntegrationStateEnum(v *RepoRequestIntegrationState) string {
	if v == nil {
		return toRepoRequestIntegrationStatePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *RepoRequestIntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoRequestIntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = RepoRequestIntegrationState(0)
		case "EXPORTING":
			*v = RepoRequestIntegrationState(1)
		case "PROCESSING":
			*v = RepoRequestIntegrationState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *RepoRequestIntegrationState) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "IDLE":
		*v = 0
	case "EXPORTING":
		*v = 1
	case "PROCESSING":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoRequestIntegrationState) MarshalJSON() ([]byte, error) {
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
func (v RepoRequestIntegrationState) String() string {
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
func (v *RepoRequestIntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoRequestIntegrationState:
		*v = val
	case int32:
		*v = RepoRequestIntegrationState(int32(val))
	case int:
		*v = RepoRequestIntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = RepoRequestIntegrationState(0)
		case "EXPORTING":
			*v = RepoRequestIntegrationState(1)
		case "PROCESSING":
			*v = RepoRequestIntegrationState(2)
		}
	}
	return nil
}

const (
	// RepoRequestIntegrationStateIdle is the enumeration value for idle
	RepoRequestIntegrationStateIdle RepoRequestIntegrationState = 0
	// RepoRequestIntegrationStateExporting is the enumeration value for exporting
	RepoRequestIntegrationStateExporting RepoRequestIntegrationState = 1
	// RepoRequestIntegrationStateProcessing is the enumeration value for processing
	RepoRequestIntegrationStateProcessing RepoRequestIntegrationState = 2
)

// RepoRequestIntegrationSystemType is the enumeration type for system_type
type RepoRequestIntegrationSystemType int32

// toRepoRequestIntegrationSystemTypePointer is the enumeration pointer type for system_type
func toRepoRequestIntegrationSystemTypePointer(v int32) *RepoRequestIntegrationSystemType {
	nv := RepoRequestIntegrationSystemType(v)
	return &nv
}

// toRepoRequestIntegrationSystemTypeEnum is the enumeration pointer wrapper for system_type
func toRepoRequestIntegrationSystemTypeEnum(v *RepoRequestIntegrationSystemType) string {
	if v == nil {
		return toRepoRequestIntegrationSystemTypePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *RepoRequestIntegrationSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = RepoRequestIntegrationSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = RepoRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = RepoRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = RepoRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = RepoRequestIntegrationSystemType(3)
		case "USER":
			*v = RepoRequestIntegrationSystemType(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *RepoRequestIntegrationSystemType) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "WORK":
		*v = 0
	case "SOURCECODE":
		*v = 1
	case "CODEQUALITY":
		*v = 2
	case "CALENDAR":
		*v = 3
	case "USER":
		*v = 4
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v RepoRequestIntegrationSystemType) MarshalJSON() ([]byte, error) {
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
func (v RepoRequestIntegrationSystemType) String() string {
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
func (v *RepoRequestIntegrationSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case RepoRequestIntegrationSystemType:
		*v = val
	case int32:
		*v = RepoRequestIntegrationSystemType(int32(val))
	case int:
		*v = RepoRequestIntegrationSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = RepoRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = RepoRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = RepoRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = RepoRequestIntegrationSystemType(3)
		case "USER":
			*v = RepoRequestIntegrationSystemType(4)
		}
	}
	return nil
}

const (
	// RepoRequestIntegrationSystemTypeWork is the enumeration value for work
	RepoRequestIntegrationSystemTypeWork RepoRequestIntegrationSystemType = 0
	// RepoRequestIntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	RepoRequestIntegrationSystemTypeSourcecode RepoRequestIntegrationSystemType = 1
	// RepoRequestIntegrationSystemTypeCodequality is the enumeration value for codequality
	RepoRequestIntegrationSystemTypeCodequality RepoRequestIntegrationSystemType = 2
	// RepoRequestIntegrationSystemTypeCalendar is the enumeration value for calendar
	RepoRequestIntegrationSystemTypeCalendar RepoRequestIntegrationSystemType = 3
	// RepoRequestIntegrationSystemTypeUser is the enumeration value for user
	RepoRequestIntegrationSystemTypeUser RepoRequestIntegrationSystemType = 4
)

// RepoRequestIntegrationThrottledUntil represents the object structure for throttled_until
type RepoRequestIntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegrationValidatedDate represents the object structure for validated_date
type RepoRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestIntegrationValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestIntegrationValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestIntegrationValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestIntegrationValidatedDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

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

// RepoRequestIntegration represents the object structure for integration
type RepoRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization RepoRequestIntegrationAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []RepoRequestIntegrationEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
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
	LastExportCompletedDate RepoRequestIntegrationLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made (set by the backend)
	LastExportRequestedDate RepoRequestIntegrationLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingCompletedDate when the processing completes (set by the pipeline)
	LastProcessingCompletedDate RepoRequestIntegrationLastProcessingCompletedDate `json:"last_processing_completed_date" codec:"last_processing_completed_date" bson:"last_processing_completed_date" yaml:"last_processing_completed_date" faker:"-"`
	// LastProcessingStartedDate when the processing starts (set by the pipeline)
	LastProcessingStartedDate RepoRequestIntegrationLastProcessingStartedDate `json:"last_processing_started_date" codec:"last_processing_started_date" bson:"last_processing_started_date" yaml:"last_processing_started_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location RepoRequestIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OnboardCompletedDate when the last request for integration metadata was finished
	OnboardCompletedDate RepoRequestIntegrationOnboardCompletedDate `json:"onboard_completed_date" codec:"onboard_completed_date" bson:"onboard_completed_date" yaml:"onboard_completed_date" faker:"-"`
	// OnboardRequestedDate when the last request for integration metadata was made
	OnboardRequestedDate RepoRequestIntegrationOnboardRequestedDate `json:"onboard_requested_date" codec:"onboard_requested_date" bson:"onboard_requested_date" yaml:"onboard_requested_date" faker:"-"`
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
	State RepoRequestIntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType RepoRequestIntegrationSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
	TeamID *string `json:"team_id,omitempty" codec:"team_id,omitempty" bson:"team_id" yaml:"team_id,omitempty" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *RepoRequestIntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate RepoRequestIntegrationValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toRepoRequestIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestIntegration:
		return v.ToMap()

	case RepoRequestIntegrationAuthorization:
		return v.ToMap()

	case []RepoRequestIntegrationEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case RepoRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	case RepoRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	case RepoRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	case RepoRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	case RepoRequestIntegrationLocation:
		return v.String()

	case RepoRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	case RepoRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	case RepoRequestIntegrationState:
		return v.String()

	case RepoRequestIntegrationSystemType:
		return v.String()

	case *RepoRequestIntegrationThrottledUntil:
		return v.ToMap()

	case RepoRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toRepoRequestIntegrationObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toRepoRequestIntegrationObject(o.Authorization, false),
		// Config the integration configuration controlled by the integration itself
		"config": toRepoRequestIntegrationObject(o.Config, true),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toRepoRequestIntegrationObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toRepoRequestIntegrationObject(o.CreatedByUserID, true),
		// CustomerID the customer id for the model instance
		"customer_id": toRepoRequestIntegrationObject(o.CustomerID, false),
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toRepoRequestIntegrationObject(o.EntityErrors, false),
		// ErrorMessage The error message from an export run
		"error_message": toRepoRequestIntegrationObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent or any other error
		"errored": toRepoRequestIntegrationObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toRepoRequestIntegrationObject(o.Exclusions, false),
		// Exportable a flag to indicate if the integration is ready for export
		"exportable": toRepoRequestIntegrationObject(o.Exportable, false),
		// ID the primary key for the model instance
		"id": toRepoRequestIntegrationObject(o.ID, false),
		// Inclusions The inclusion list for this integration
		"inclusions": toRepoRequestIntegrationObject(o.Inclusions, false),
		// IntegrationInstanceID the integration instance id
		"integration_instance_id": toRepoRequestIntegrationObject(o.IntegrationInstanceID, true),
		// Interval the interval in milliseconds for how often an export job is scheduled
		"interval": toRepoRequestIntegrationObject(o.Interval, false),
		// LastExportCompletedDate when the export response was received (set by the backend)
		"last_export_completed_date": toRepoRequestIntegrationObject(o.LastExportCompletedDate, false),
		// LastExportRequestedDate when the export request was made (set by the backend)
		"last_export_requested_date": toRepoRequestIntegrationObject(o.LastExportRequestedDate, false),
		// LastProcessingCompletedDate when the processing completes (set by the pipeline)
		"last_processing_completed_date": toRepoRequestIntegrationObject(o.LastProcessingCompletedDate, false),
		// LastProcessingStartedDate when the processing starts (set by the pipeline)
		"last_processing_started_date": toRepoRequestIntegrationObject(o.LastProcessingStartedDate, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toRepoRequestIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toRepoRequestIntegrationObject(o.Name, false),
		// OnboardCompletedDate when the last request for integration metadata was finished
		"onboard_completed_date": toRepoRequestIntegrationObject(o.OnboardCompletedDate, false),
		// OnboardRequestedDate when the last request for integration metadata was made
		"onboard_requested_date": toRepoRequestIntegrationObject(o.OnboardRequestedDate, false),
		// Onboarding true if the agent is fetching metadata for the integration
		"onboarding": toRepoRequestIntegrationObject(o.Onboarding, false),
		// Organization The origanization authorized. Used for azure integrations
		"organization": toRepoRequestIntegrationObject(o.Organization, true),
		// Paused true if the agent is paused and should not start new scheduled jobs
		"paused": toRepoRequestIntegrationObject(o.Paused, false),
		// Processed If the integration has been processed at least once
		"processed": toRepoRequestIntegrationObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toRepoRequestIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toRepoRequestIntegrationObject(o.RefType, false),
		// ServerVersion the server version for this integration
		"server_version": toRepoRequestIntegrationObject(o.ServerVersion, true),
		// State the current state of the integration
		"state": toRepoRequestIntegrationObject(o.State, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toRepoRequestIntegrationObject(o.SystemType, false),
		// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
		"team_id": toRepoRequestIntegrationObject(o.TeamID, true),
		// Throttled Set to true when integration is throttled.
		"throttled": toRepoRequestIntegrationObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toRepoRequestIntegrationObject(o.ThrottledUntil, true),
		// Validated If the validation has been run against this instance
		"validated": toRepoRequestIntegrationObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toRepoRequestIntegrationObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toRepoRequestIntegrationObject(o.ValidationMessage, true),
	}
}

func (o *RepoRequestIntegration) setDefaults(frommap bool) {

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
func (o *RepoRequestIntegration) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(RepoRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*RepoRequestIntegrationAuthorization); ok {
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

		o.EntityErrors = make([]RepoRequestIntegrationEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]RepoRequestIntegrationEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*RepoRequestIntegrationEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(RepoRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm RepoRequestIntegrationEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av RepoRequestIntegrationEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(RepoRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm RepoRequestIntegrationEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := RepoRequestIntegrationEntityErrors{}
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
							var fm RepoRequestIntegrationEntityErrors
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
		} else if sv, ok := val.(RepoRequestIntegrationLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationLastExportCompletedDate); ok {
			// struct pointer
			o.LastExportCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportCompletedDate.Epoch = dt.Epoch
			o.LastExportCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastExportCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
		} else if sv, ok := val.(RepoRequestIntegrationLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationLastExportRequestedDate); ok {
			// struct pointer
			o.LastExportRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportRequestedDate.Epoch = dt.Epoch
			o.LastExportRequestedDate.Rfc3339 = dt.Rfc3339
			o.LastExportRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
		} else if sv, ok := val.(RepoRequestIntegrationLastProcessingCompletedDate); ok {
			// struct
			o.LastProcessingCompletedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationLastProcessingCompletedDate); ok {
			// struct pointer
			o.LastProcessingCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingCompletedDate.Epoch = dt.Epoch
			o.LastProcessingCompletedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
		} else if sv, ok := val.(RepoRequestIntegrationLastProcessingStartedDate); ok {
			// struct
			o.LastProcessingStartedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationLastProcessingStartedDate); ok {
			// struct pointer
			o.LastProcessingStartedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingStartedDate.Epoch = dt.Epoch
			o.LastProcessingStartedDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingStartedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["location"].(RepoRequestIntegrationLocation); ok {
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
		} else if sv, ok := val.(RepoRequestIntegrationOnboardCompletedDate); ok {
			// struct
			o.OnboardCompletedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationOnboardCompletedDate); ok {
			// struct pointer
			o.OnboardCompletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.OnboardCompletedDate.Epoch = dt.Epoch
			o.OnboardCompletedDate.Rfc3339 = dt.Rfc3339
			o.OnboardCompletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
		} else if sv, ok := val.(RepoRequestIntegrationOnboardRequestedDate); ok {
			// struct
			o.OnboardRequestedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationOnboardRequestedDate); ok {
			// struct pointer
			o.OnboardRequestedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.OnboardRequestedDate.Epoch = dt.Epoch
			o.OnboardRequestedDate.Rfc3339 = dt.Rfc3339
			o.OnboardRequestedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
	if val, ok := kv["state"].(RepoRequestIntegrationState); ok {
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
	if val, ok := kv["system_type"].(RepoRequestIntegrationSystemType); ok {
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
		o.ThrottledUntil = &RepoRequestIntegrationThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(RepoRequestIntegrationThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*RepoRequestIntegrationThrottledUntil); ok {
			// struct pointer
			o.ThrottledUntil = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ThrottledUntil.Epoch = dt.Epoch
			o.ThrottledUntil.Rfc3339 = dt.Rfc3339
			o.ThrottledUntil.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
		} else if sv, ok := val.(RepoRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*RepoRequestIntegrationValidatedDate); ok {
			// struct pointer
			o.ValidatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.ValidatedDate.Epoch = dt.Epoch
			o.ValidatedDate.Rfc3339 = dt.Rfc3339
			o.ValidatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

// RepoRequestRequestDate represents the object structure for request_date
type RepoRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *RepoRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toRepoRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *RepoRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// RepoRequest an agent action to request adding new repos
type RepoRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration RepoRequestIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate RepoRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*RepoRequest)(nil)

func toRepoRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoRequest:
		return v.ToMap()

	case RepoRequestIntegration:
		return v.ToMap()

	case RepoRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of RepoRequest
func (o *RepoRequest) String() string {
	return fmt.Sprintf("agent.RepoRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *RepoRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *RepoRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *RepoRequest) GetModelName() datamodel.ModelNameType {
	return RepoRequestModelName
}

// NewRepoRequestID provides a template for generating an ID field for RepoRequest
func NewRepoRequestID(customerID string, refType string, refID string) string {
	return hash.Values("RepoRequest", customerID, refType, refID)
}

func (o *RepoRequest) setDefaults(frommap bool) {
	if o.Integration.EntityErrors == nil {
		o.Integration.EntityErrors = make([]RepoRequestIntegrationEntityErrors, 0)
	}
	if o.Integration.Exclusions == nil {
		o.Integration.Exclusions = make([]string, 0)
	}
	if o.Integration.Inclusions == nil {
		o.Integration.Inclusions = make([]string, 0)
	}
	if o.Integration.ThrottledUntil == nil {
		o.Integration.ThrottledUntil = &RepoRequestIntegrationThrottledUntil{}
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *RepoRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *RepoRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoRequest) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoRequest) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoRequestModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *RepoRequest) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *RepoRequest) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *RepoRequest) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *RepoRequest) SetRefType(t string) {
	o.RefType = t
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *RepoRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *RepoRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two RepoRequest objects are equal
func (o *RepoRequest) IsEqual(other *RepoRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toRepoRequestObject(o.CustomerID, false),
		"id":                      toRepoRequestObject(o.ID, false),
		"integration":             toRepoRequestObject(o.Integration, false),
		"integration_instance_id": toRepoRequestObject(o.IntegrationInstanceID, true),
		"ref_id":                  toRepoRequestObject(o.RefID, false),
		"ref_type":                toRepoRequestObject(o.RefType, false),
		"request_date":            toRepoRequestObject(o.RequestDate, false),
		"uuid":                    toRepoRequestObject(o.UUID, false),
		"hashcode":                toRepoRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(RepoRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*RepoRequestIntegration); ok {
			// struct pointer
			o.Integration = *sp
		}
	} else {
		o.Integration.FromMap(map[string]interface{}{})
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
		} else if sv, ok := val.(RepoRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*RepoRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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
func (o *RepoRequest) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestDate)
	args = append(args, o.UUID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *RepoRequest) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *RepoRequest) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// RepoRequestPartial is a partial struct for upsert mutations for RepoRequest
type RepoRequestPartial struct {
	// Integration the integration details to use
	Integration *RepoRequestIntegration `json:"integration,omitempty"`
	// RequestDate the date when the request was made
	RequestDate *RepoRequestRequestDate `json:"request_date,omitempty"`
	// UUID the agent unique identifier
	UUID *string `json:"uuid,omitempty"`
}

var _ datamodel.PartialModel = (*RepoRequestPartial)(nil)

// GetModelName returns the name of the model
func (o *RepoRequestPartial) GetModelName() datamodel.ModelNameType {
	return RepoRequestModelName
}

// ToMap returns the object as a map
func (o *RepoRequestPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"integration":  toRepoRequestObject(o.Integration, true),
		"request_date": toRepoRequestObject(o.RequestDate, true),
		"uuid":         toRepoRequestObject(o.UUID, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "request_date" {
				if dt, ok := v.(*RepoRequestRequestDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *RepoRequestPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoRequestPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoRequestPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *RepoRequestPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *RepoRequestPartial) FromMap(kv map[string]interface{}) {

	if o.Integration == nil {
		o.Integration = &RepoRequestIntegration{}
	}

	if val, ok := kv["integration"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Integration.FromMap(kv)
		} else if sv, ok := val.(RepoRequestIntegration); ok {
			// struct
			o.Integration = &sv
		} else if sp, ok := val.(*RepoRequestIntegration); ok {
			// struct pointer
			o.Integration = sp
		}
	} else {
		o.Integration.FromMap(map[string]interface{}{})
	}

	if o.RequestDate == nil {
		o.RequestDate = &RepoRequestRequestDate{}
	}

	if val, ok := kv["request_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.RequestDate.FromMap(kv)
		} else if sv, ok := val.(RepoRequestRequestDate); ok {
			// struct
			o.RequestDate = &sv
		} else if sp, ok := val.(*RepoRequestRequestDate); ok {
			// struct pointer
			o.RequestDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.RequestDate.Epoch = dt.Epoch
			o.RequestDate.Rfc3339 = dt.Rfc3339
			o.RequestDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["uuid"].(*string); ok {
		o.UUID = val
	} else if val, ok := kv["uuid"].(string); ok {
		o.UUID = &val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UUID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}
