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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// WorkStatusRequestTable is the default table name
	WorkStatusRequestTable datamodel.ModelNameType = "agent_workstatusrequest"

	// WorkStatusRequestModelName is the model name
	WorkStatusRequestModelName datamodel.ModelNameType = "agent.WorkStatusRequest"
)

const (
	// WorkStatusRequestModelCustomerIDColumn is the column json value customer_id
	WorkStatusRequestModelCustomerIDColumn = "customer_id"
	// WorkStatusRequestModelIDColumn is the column json value id
	WorkStatusRequestModelIDColumn = "id"
	// WorkStatusRequestModelIntegrationColumn is the column json value integration
	WorkStatusRequestModelIntegrationColumn = "integration"
	// WorkStatusRequestModelIntegrationActiveColumn is the column json value active
	WorkStatusRequestModelIntegrationActiveColumn = "active"
	// WorkStatusRequestModelIntegrationAuthorizationColumn is the column json value authorization
	WorkStatusRequestModelIntegrationAuthorizationColumn = "authorization"
	// WorkStatusRequestModelIntegrationAuthorizationAccessTokenColumn is the column json value access_token
	WorkStatusRequestModelIntegrationAuthorizationAccessTokenColumn = "access_token"
	// WorkStatusRequestModelIntegrationAuthorizationRefreshTokenColumn is the column json value refresh_token
	WorkStatusRequestModelIntegrationAuthorizationRefreshTokenColumn = "refresh_token"
	// WorkStatusRequestModelIntegrationAuthorizationURLColumn is the column json value url
	WorkStatusRequestModelIntegrationAuthorizationURLColumn = "url"
	// WorkStatusRequestModelIntegrationAuthorizationUsernameColumn is the column json value username
	WorkStatusRequestModelIntegrationAuthorizationUsernameColumn = "username"
	// WorkStatusRequestModelIntegrationAuthorizationPasswordColumn is the column json value password
	WorkStatusRequestModelIntegrationAuthorizationPasswordColumn = "password"
	// WorkStatusRequestModelIntegrationAuthorizationAPITokenColumn is the column json value api_token
	WorkStatusRequestModelIntegrationAuthorizationAPITokenColumn = "api_token"
	// WorkStatusRequestModelIntegrationAuthorizationCollectionNameColumn is the column json value collection_name
	WorkStatusRequestModelIntegrationAuthorizationCollectionNameColumn = "collection_name"
	// WorkStatusRequestModelIntegrationAuthorizationAPIKeyColumn is the column json value api_key
	WorkStatusRequestModelIntegrationAuthorizationAPIKeyColumn = "api_key"
	// WorkStatusRequestModelIntegrationAuthorizationAuthorizationColumn is the column json value authorization
	WorkStatusRequestModelIntegrationAuthorizationAuthorizationColumn = "authorization"
	// WorkStatusRequestModelIntegrationAuthorizationHostnameColumn is the column json value hostname
	WorkStatusRequestModelIntegrationAuthorizationHostnameColumn = "hostname"
	// WorkStatusRequestModelIntegrationAuthorizationAPIVersionColumn is the column json value api_version
	WorkStatusRequestModelIntegrationAuthorizationAPIVersionColumn = "api_version"
	// WorkStatusRequestModelIntegrationAuthorizationOrganizationColumn is the column json value organization
	WorkStatusRequestModelIntegrationAuthorizationOrganizationColumn = "organization"
	// WorkStatusRequestModelIntegrationCreatedByProfileIDColumn is the column json value created_by_profile_id
	WorkStatusRequestModelIntegrationCreatedByProfileIDColumn = "created_by_profile_id"
	// WorkStatusRequestModelIntegrationCreatedByUserIDColumn is the column json value created_by_user_id
	WorkStatusRequestModelIntegrationCreatedByUserIDColumn = "created_by_user_id"
	// WorkStatusRequestModelIntegrationCustomerIDColumn is the column json value customer_id
	WorkStatusRequestModelIntegrationCustomerIDColumn = "customer_id"
	// WorkStatusRequestModelIntegrationEntityErrorsColumn is the column json value entity_errors
	WorkStatusRequestModelIntegrationEntityErrorsColumn = "entity_errors"
	// WorkStatusRequestModelIntegrationEntityErrorsIDColumn is the column json value id
	WorkStatusRequestModelIntegrationEntityErrorsIDColumn = "id"
	// WorkStatusRequestModelIntegrationEntityErrorsRefIDColumn is the column json value ref_id
	WorkStatusRequestModelIntegrationEntityErrorsRefIDColumn = "ref_id"
	// WorkStatusRequestModelIntegrationEntityErrorsErrorColumn is the column json value error
	WorkStatusRequestModelIntegrationEntityErrorsErrorColumn = "error"
	// WorkStatusRequestModelIntegrationErrorMessageColumn is the column json value error_message
	WorkStatusRequestModelIntegrationErrorMessageColumn = "error_message"
	// WorkStatusRequestModelIntegrationErroredColumn is the column json value errored
	WorkStatusRequestModelIntegrationErroredColumn = "errored"
	// WorkStatusRequestModelIntegrationExclusionsColumn is the column json value exclusions
	WorkStatusRequestModelIntegrationExclusionsColumn = "exclusions"
	// WorkStatusRequestModelIntegrationExportableColumn is the column json value exportable
	WorkStatusRequestModelIntegrationExportableColumn = "exportable"
	// WorkStatusRequestModelIntegrationIDColumn is the column json value id
	WorkStatusRequestModelIntegrationIDColumn = "id"
	// WorkStatusRequestModelIntegrationInclusionsColumn is the column json value inclusions
	WorkStatusRequestModelIntegrationInclusionsColumn = "inclusions"
	// WorkStatusRequestModelIntegrationLastExportCompletedDateColumn is the column json value last_export_completed_date
	WorkStatusRequestModelIntegrationLastExportCompletedDateColumn = "last_export_completed_date"
	// WorkStatusRequestModelIntegrationLastExportCompletedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationLastExportCompletedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationLastExportCompletedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationLastExportCompletedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationLastExportCompletedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationLastExportCompletedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationLastExportRequestedDateColumn is the column json value last_export_requested_date
	WorkStatusRequestModelIntegrationLastExportRequestedDateColumn = "last_export_requested_date"
	// WorkStatusRequestModelIntegrationLastExportRequestedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationLastExportRequestedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationLastExportRequestedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationLastExportRequestedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationLastExportRequestedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationLastExportRequestedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationLastProcessingCompletedDateColumn is the column json value last_processing_completed_date
	WorkStatusRequestModelIntegrationLastProcessingCompletedDateColumn = "last_processing_completed_date"
	// WorkStatusRequestModelIntegrationLastProcessingCompletedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationLastProcessingCompletedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationLastProcessingCompletedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationLastProcessingCompletedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationLastProcessingCompletedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationLastProcessingCompletedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationLastProcessingStartedDateColumn is the column json value last_processing_started_date
	WorkStatusRequestModelIntegrationLastProcessingStartedDateColumn = "last_processing_started_date"
	// WorkStatusRequestModelIntegrationLastProcessingStartedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationLastProcessingStartedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationLastProcessingStartedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationLastProcessingStartedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationLastProcessingStartedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationLastProcessingStartedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationLocationColumn is the column json value location
	WorkStatusRequestModelIntegrationLocationColumn = "location"
	// WorkStatusRequestModelIntegrationNameColumn is the column json value name
	WorkStatusRequestModelIntegrationNameColumn = "name"
	// WorkStatusRequestModelIntegrationOnboardCompletedDateColumn is the column json value onboard_completed_date
	WorkStatusRequestModelIntegrationOnboardCompletedDateColumn = "onboard_completed_date"
	// WorkStatusRequestModelIntegrationOnboardCompletedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationOnboardCompletedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationOnboardCompletedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationOnboardCompletedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationOnboardCompletedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationOnboardCompletedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationOnboardRequestedDateColumn is the column json value onboard_requested_date
	WorkStatusRequestModelIntegrationOnboardRequestedDateColumn = "onboard_requested_date"
	// WorkStatusRequestModelIntegrationOnboardRequestedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationOnboardRequestedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationOnboardRequestedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationOnboardRequestedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationOnboardRequestedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationOnboardRequestedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationOnboardingColumn is the column json value onboarding
	WorkStatusRequestModelIntegrationOnboardingColumn = "onboarding"
	// WorkStatusRequestModelIntegrationOrganizationColumn is the column json value organization
	WorkStatusRequestModelIntegrationOrganizationColumn = "organization"
	// WorkStatusRequestModelIntegrationProcessedColumn is the column json value processed
	WorkStatusRequestModelIntegrationProcessedColumn = "processed"
	// WorkStatusRequestModelIntegrationRefIDColumn is the column json value ref_id
	WorkStatusRequestModelIntegrationRefIDColumn = "ref_id"
	// WorkStatusRequestModelIntegrationRefTypeColumn is the column json value ref_type
	WorkStatusRequestModelIntegrationRefTypeColumn = "ref_type"
	// WorkStatusRequestModelIntegrationServerVersionColumn is the column json value server_version
	WorkStatusRequestModelIntegrationServerVersionColumn = "server_version"
	// WorkStatusRequestModelIntegrationStateColumn is the column json value state
	WorkStatusRequestModelIntegrationStateColumn = "state"
	// WorkStatusRequestModelIntegrationSystemTypeColumn is the column json value system_type
	WorkStatusRequestModelIntegrationSystemTypeColumn = "system_type"
	// WorkStatusRequestModelIntegrationTeamIDColumn is the column json value team_id
	WorkStatusRequestModelIntegrationTeamIDColumn = "team_id"
	// WorkStatusRequestModelIntegrationThrottledColumn is the column json value throttled
	WorkStatusRequestModelIntegrationThrottledColumn = "throttled"
	// WorkStatusRequestModelIntegrationThrottledUntilColumn is the column json value throttled_until
	WorkStatusRequestModelIntegrationThrottledUntilColumn = "throttled_until"
	// WorkStatusRequestModelIntegrationThrottledUntilEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationThrottledUntilEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationThrottledUntilOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationThrottledUntilOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationThrottledUntilRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationThrottledUntilRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationValidatedColumn is the column json value validated
	WorkStatusRequestModelIntegrationValidatedColumn = "validated"
	// WorkStatusRequestModelIntegrationValidatedDateColumn is the column json value validated_date
	WorkStatusRequestModelIntegrationValidatedDateColumn = "validated_date"
	// WorkStatusRequestModelIntegrationValidatedDateEpochColumn is the column json value epoch
	WorkStatusRequestModelIntegrationValidatedDateEpochColumn = "epoch"
	// WorkStatusRequestModelIntegrationValidatedDateOffsetColumn is the column json value offset
	WorkStatusRequestModelIntegrationValidatedDateOffsetColumn = "offset"
	// WorkStatusRequestModelIntegrationValidatedDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelIntegrationValidatedDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelIntegrationValidationMessageColumn is the column json value validation_message
	WorkStatusRequestModelIntegrationValidationMessageColumn = "validation_message"
	// WorkStatusRequestModelRefIDColumn is the column json value ref_id
	WorkStatusRequestModelRefIDColumn = "ref_id"
	// WorkStatusRequestModelRefTypeColumn is the column json value ref_type
	WorkStatusRequestModelRefTypeColumn = "ref_type"
	// WorkStatusRequestModelRequestDateColumn is the column json value request_date
	WorkStatusRequestModelRequestDateColumn = "request_date"
	// WorkStatusRequestModelRequestDateEpochColumn is the column json value epoch
	WorkStatusRequestModelRequestDateEpochColumn = "epoch"
	// WorkStatusRequestModelRequestDateOffsetColumn is the column json value offset
	WorkStatusRequestModelRequestDateOffsetColumn = "offset"
	// WorkStatusRequestModelRequestDateRfc3339Column is the column json value rfc3339
	WorkStatusRequestModelRequestDateRfc3339Column = "rfc3339"
	// WorkStatusRequestModelUUIDColumn is the column json value uuid
	WorkStatusRequestModelUUIDColumn = "uuid"
)

// WorkStatusRequestIntegrationAuthorization represents the object structure for authorization
type WorkStatusRequestIntegrationAuthorization struct {
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

func toWorkStatusRequestIntegrationAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toWorkStatusRequestIntegrationAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toWorkStatusRequestIntegrationAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toWorkStatusRequestIntegrationAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toWorkStatusRequestIntegrationAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toWorkStatusRequestIntegrationAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toWorkStatusRequestIntegrationAuthorizationObject(o.APIToken, true),
		// CollectionName Collection name for instance, if relevant
		"collection_name": toWorkStatusRequestIntegrationAuthorizationObject(o.CollectionName, true),
		// APIKey API Key for instance, if relevant
		"api_key": toWorkStatusRequestIntegrationAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toWorkStatusRequestIntegrationAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toWorkStatusRequestIntegrationAuthorizationObject(o.Hostname, true),
		// APIVersion the api version of the integration
		"api_version": toWorkStatusRequestIntegrationAuthorizationObject(o.APIVersion, true),
		// Organization Organization for instance, if relevant
		"organization": toWorkStatusRequestIntegrationAuthorizationObject(o.Organization, true),
	}
}

func (o *WorkStatusRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationEntityErrors represents the object structure for entity_errors
type WorkStatusRequestIntegrationEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
}

func toWorkStatusRequestIntegrationEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toWorkStatusRequestIntegrationEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toWorkStatusRequestIntegrationEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export
		"error": toWorkStatusRequestIntegrationEntityErrorsObject(o.Error, false),
	}
}

func (o *WorkStatusRequestIntegrationEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationEntityErrors) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationLastExportCompletedDate represents the object structure for last_export_completed_date
type WorkStatusRequestIntegrationLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationLastExportRequestedDate represents the object structure for last_export_requested_date
type WorkStatusRequestIntegrationLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationLastProcessingCompletedDate represents the object structure for last_processing_completed_date
type WorkStatusRequestIntegrationLastProcessingCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationLastProcessingCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationLastProcessingCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationLastProcessingCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationLastProcessingCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationLastProcessingCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationLastProcessingCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationLastProcessingCompletedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationLastProcessingStartedDate represents the object structure for last_processing_started_date
type WorkStatusRequestIntegrationLastProcessingStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationLastProcessingStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationLastProcessingStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationLastProcessingStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationLastProcessingStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationLastProcessingStartedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationLastProcessingStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationLastProcessingStartedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationLocation is the enumeration type for location
type WorkStatusRequestIntegrationLocation int32

// UnmarshalBSONValue for unmarshaling value
func (v *WorkStatusRequestIntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WorkStatusRequestIntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = WorkStatusRequestIntegrationLocation(0)
		case "CLOUD":
			*v = WorkStatusRequestIntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusRequestIntegrationLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusRequestIntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationLocation
func (v WorkStatusRequestIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *WorkStatusRequestIntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case int32:
		*v = WorkStatusRequestIntegrationLocation(int32(val))
	case int:
		*v = WorkStatusRequestIntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = WorkStatusRequestIntegrationLocation(0)
		case "CLOUD":
			*v = WorkStatusRequestIntegrationLocation(1)
		}
	}
	return nil
}

const (
	// WorkStatusRequestIntegrationLocationPrivate is the enumeration value for private
	WorkStatusRequestIntegrationLocationPrivate WorkStatusRequestIntegrationLocation = 0
	// WorkStatusRequestIntegrationLocationCloud is the enumeration value for cloud
	WorkStatusRequestIntegrationLocationCloud WorkStatusRequestIntegrationLocation = 1
)

// WorkStatusRequestIntegrationOnboardCompletedDate represents the object structure for onboard_completed_date
type WorkStatusRequestIntegrationOnboardCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationOnboardCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationOnboardCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationOnboardCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationOnboardCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationOnboardCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationOnboardCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationOnboardCompletedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationOnboardRequestedDate represents the object structure for onboard_requested_date
type WorkStatusRequestIntegrationOnboardRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationOnboardRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationOnboardRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationOnboardRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationOnboardRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationOnboardRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationOnboardRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationOnboardRequestedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationState is the enumeration type for state
type WorkStatusRequestIntegrationState int32

// UnmarshalBSONValue for unmarshaling value
func (v *WorkStatusRequestIntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WorkStatusRequestIntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = WorkStatusRequestIntegrationState(0)
		case "EXPORTING":
			*v = WorkStatusRequestIntegrationState(1)
		case "PROCESSING":
			*v = WorkStatusRequestIntegrationState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusRequestIntegrationState) UnmarshalJSON(buf []byte) error {
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
func (v WorkStatusRequestIntegrationState) MarshalJSON() ([]byte, error) {
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
func (v WorkStatusRequestIntegrationState) String() string {
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
func (v *WorkStatusRequestIntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case int32:
		*v = WorkStatusRequestIntegrationState(int32(val))
	case int:
		*v = WorkStatusRequestIntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = WorkStatusRequestIntegrationState(0)
		case "EXPORTING":
			*v = WorkStatusRequestIntegrationState(1)
		case "PROCESSING":
			*v = WorkStatusRequestIntegrationState(2)
		}
	}
	return nil
}

const (
	// WorkStatusRequestIntegrationStateIdle is the enumeration value for idle
	WorkStatusRequestIntegrationStateIdle WorkStatusRequestIntegrationState = 0
	// WorkStatusRequestIntegrationStateExporting is the enumeration value for exporting
	WorkStatusRequestIntegrationStateExporting WorkStatusRequestIntegrationState = 1
	// WorkStatusRequestIntegrationStateProcessing is the enumeration value for processing
	WorkStatusRequestIntegrationStateProcessing WorkStatusRequestIntegrationState = 2
)

// WorkStatusRequestIntegrationSystemType is the enumeration type for system_type
type WorkStatusRequestIntegrationSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *WorkStatusRequestIntegrationSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WorkStatusRequestIntegrationSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = WorkStatusRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = WorkStatusRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = WorkStatusRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = WorkStatusRequestIntegrationSystemType(3)
		case "USER":
			*v = WorkStatusRequestIntegrationSystemType(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusRequestIntegrationSystemType) UnmarshalJSON(buf []byte) error {
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
func (v WorkStatusRequestIntegrationSystemType) MarshalJSON() ([]byte, error) {
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
func (v WorkStatusRequestIntegrationSystemType) String() string {
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
func (v *WorkStatusRequestIntegrationSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case int32:
		*v = WorkStatusRequestIntegrationSystemType(int32(val))
	case int:
		*v = WorkStatusRequestIntegrationSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = WorkStatusRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = WorkStatusRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = WorkStatusRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = WorkStatusRequestIntegrationSystemType(3)
		case "USER":
			*v = WorkStatusRequestIntegrationSystemType(4)
		}
	}
	return nil
}

const (
	// WorkStatusRequestIntegrationSystemTypeWork is the enumeration value for work
	WorkStatusRequestIntegrationSystemTypeWork WorkStatusRequestIntegrationSystemType = 0
	// WorkStatusRequestIntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	WorkStatusRequestIntegrationSystemTypeSourcecode WorkStatusRequestIntegrationSystemType = 1
	// WorkStatusRequestIntegrationSystemTypeCodequality is the enumeration value for codequality
	WorkStatusRequestIntegrationSystemTypeCodequality WorkStatusRequestIntegrationSystemType = 2
	// WorkStatusRequestIntegrationSystemTypeCalendar is the enumeration value for calendar
	WorkStatusRequestIntegrationSystemTypeCalendar WorkStatusRequestIntegrationSystemType = 3
	// WorkStatusRequestIntegrationSystemTypeUser is the enumeration value for user
	WorkStatusRequestIntegrationSystemTypeUser WorkStatusRequestIntegrationSystemType = 4
)

// WorkStatusRequestIntegrationThrottledUntil represents the object structure for throttled_until
type WorkStatusRequestIntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegrationValidatedDate represents the object structure for validated_date
type WorkStatusRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestIntegrationValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestIntegrationValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestIntegrationValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestIntegrationValidatedDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequestIntegration represents the object structure for integration
type WorkStatusRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization WorkStatusRequestIntegrationAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []WorkStatusRequestIntegrationEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// Exclusions The exclusion list for this integration
	Exclusions []string `json:"exclusions" codec:"exclusions" bson:"exclusions" yaml:"exclusions" faker:"-"`
	// Exportable a flag to indicate if the integration is ready for export
	Exportable bool `json:"exportable" codec:"exportable" bson:"exportable" yaml:"exportable" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"id" yaml:"id" faker:"-"`
	// Inclusions The inclusion list for this integration
	Inclusions []string `json:"inclusions" codec:"inclusions" bson:"inclusions" yaml:"inclusions" faker:"-"`
	// LastExportCompletedDate when the export response was received (set by the backend)
	LastExportCompletedDate WorkStatusRequestIntegrationLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made (set by the backend)
	LastExportRequestedDate WorkStatusRequestIntegrationLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingCompletedDate when the processing completes (set by the pipeline)
	LastProcessingCompletedDate WorkStatusRequestIntegrationLastProcessingCompletedDate `json:"last_processing_completed_date" codec:"last_processing_completed_date" bson:"last_processing_completed_date" yaml:"last_processing_completed_date" faker:"-"`
	// LastProcessingStartedDate when the processing starts (set by the pipeline)
	LastProcessingStartedDate WorkStatusRequestIntegrationLastProcessingStartedDate `json:"last_processing_started_date" codec:"last_processing_started_date" bson:"last_processing_started_date" yaml:"last_processing_started_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location WorkStatusRequestIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OnboardCompletedDate when the last request for integration metadata was finished
	OnboardCompletedDate WorkStatusRequestIntegrationOnboardCompletedDate `json:"onboard_completed_date" codec:"onboard_completed_date" bson:"onboard_completed_date" yaml:"onboard_completed_date" faker:"-"`
	// OnboardRequestedDate when the last request for integration metadata was made
	OnboardRequestedDate WorkStatusRequestIntegrationOnboardRequestedDate `json:"onboard_requested_date" codec:"onboard_requested_date" bson:"onboard_requested_date" yaml:"onboard_requested_date" faker:"-"`
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
	State WorkStatusRequestIntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType WorkStatusRequestIntegrationSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
	TeamID *string `json:"team_id,omitempty" codec:"team_id,omitempty" bson:"team_id" yaml:"team_id,omitempty" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *WorkStatusRequestIntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate WorkStatusRequestIntegrationValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toWorkStatusRequestIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestIntegration:
		return v.ToMap()

	case WorkStatusRequestIntegrationAuthorization:
		return v.ToMap()

	case []WorkStatusRequestIntegrationEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case WorkStatusRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationLocation:
		return v.String()

	case WorkStatusRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	case WorkStatusRequestIntegrationState:
		return v.String()

	case WorkStatusRequestIntegrationSystemType:
		return v.String()

	case *WorkStatusRequestIntegrationThrottledUntil:
		return v.ToMap()

	case WorkStatusRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toWorkStatusRequestIntegrationObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toWorkStatusRequestIntegrationObject(o.Authorization, false),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toWorkStatusRequestIntegrationObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toWorkStatusRequestIntegrationObject(o.CreatedByUserID, true),
		// CustomerID the customer id for the model instance
		"customer_id": toWorkStatusRequestIntegrationObject(o.CustomerID, false),
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toWorkStatusRequestIntegrationObject(o.EntityErrors, false),
		// ErrorMessage The error message from an export run
		"error_message": toWorkStatusRequestIntegrationObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent
		"errored": toWorkStatusRequestIntegrationObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toWorkStatusRequestIntegrationObject(o.Exclusions, false),
		// Exportable a flag to indicate if the integration is ready for export
		"exportable": toWorkStatusRequestIntegrationObject(o.Exportable, false),
		// ID the primary key for the model instance
		"id": toWorkStatusRequestIntegrationObject(o.ID, false),
		// Inclusions The inclusion list for this integration
		"inclusions": toWorkStatusRequestIntegrationObject(o.Inclusions, false),
		// LastExportCompletedDate when the export response was received (set by the backend)
		"last_export_completed_date": toWorkStatusRequestIntegrationObject(o.LastExportCompletedDate, false),
		// LastExportRequestedDate when the export request was made (set by the backend)
		"last_export_requested_date": toWorkStatusRequestIntegrationObject(o.LastExportRequestedDate, false),
		// LastProcessingCompletedDate when the processing completes (set by the pipeline)
		"last_processing_completed_date": toWorkStatusRequestIntegrationObject(o.LastProcessingCompletedDate, false),
		// LastProcessingStartedDate when the processing starts (set by the pipeline)
		"last_processing_started_date": toWorkStatusRequestIntegrationObject(o.LastProcessingStartedDate, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toWorkStatusRequestIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toWorkStatusRequestIntegrationObject(o.Name, false),
		// OnboardCompletedDate when the last request for integration metadata was finished
		"onboard_completed_date": toWorkStatusRequestIntegrationObject(o.OnboardCompletedDate, false),
		// OnboardRequestedDate when the last request for integration metadata was made
		"onboard_requested_date": toWorkStatusRequestIntegrationObject(o.OnboardRequestedDate, false),
		// Onboarding true if the agent is fetching metadata for the integration
		"onboarding": toWorkStatusRequestIntegrationObject(o.Onboarding, false),
		// Organization The origanization authorized. Used for azure integrations
		"organization": toWorkStatusRequestIntegrationObject(o.Organization, true),
		// Processed If the integration has been processed at least once
		"processed": toWorkStatusRequestIntegrationObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toWorkStatusRequestIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toWorkStatusRequestIntegrationObject(o.RefType, false),
		// ServerVersion the server version for this integration
		"server_version": toWorkStatusRequestIntegrationObject(o.ServerVersion, true),
		// State the current state of the integration
		"state": toWorkStatusRequestIntegrationObject(o.State, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toWorkStatusRequestIntegrationObject(o.SystemType, false),
		// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
		"team_id": toWorkStatusRequestIntegrationObject(o.TeamID, true),
		// Throttled Set to true when integration is throttled.
		"throttled": toWorkStatusRequestIntegrationObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toWorkStatusRequestIntegrationObject(o.ThrottledUntil, true),
		// Validated If the validation has been run against this instance
		"validated": toWorkStatusRequestIntegrationObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toWorkStatusRequestIntegrationObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toWorkStatusRequestIntegrationObject(o.ValidationMessage, true),
	}
}

func (o *WorkStatusRequestIntegration) setDefaults(frommap bool) {

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
func (o *WorkStatusRequestIntegration) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(WorkStatusRequestIntegrationAuthorization); ok {
			// struct
			o.Authorization = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationAuthorization); ok {
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

		o.EntityErrors = make([]WorkStatusRequestIntegrationEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]WorkStatusRequestIntegrationEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*WorkStatusRequestIntegrationEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(WorkStatusRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm WorkStatusRequestIntegrationEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av WorkStatusRequestIntegrationEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(WorkStatusRequestIntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm WorkStatusRequestIntegrationEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := WorkStatusRequestIntegrationEntityErrors{}
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
							var fm WorkStatusRequestIntegrationEntityErrors
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationLastExportCompletedDate); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationLastExportRequestedDate); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationLastProcessingCompletedDate); ok {
			// struct
			o.LastProcessingCompletedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationLastProcessingCompletedDate); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationLastProcessingStartedDate); ok {
			// struct
			o.LastProcessingStartedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationLastProcessingStartedDate); ok {
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

	if val, ok := kv["location"].(WorkStatusRequestIntegrationLocation); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationOnboardCompletedDate); ok {
			// struct
			o.OnboardCompletedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationOnboardCompletedDate); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationOnboardRequestedDate); ok {
			// struct
			o.OnboardRequestedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationOnboardRequestedDate); ok {
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

	if val, ok := kv["state"].(WorkStatusRequestIntegrationState); ok {
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

	if val, ok := kv["system_type"].(WorkStatusRequestIntegrationSystemType); ok {
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
		o.ThrottledUntil = &WorkStatusRequestIntegrationThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(WorkStatusRequestIntegrationThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationThrottledUntil); ok {
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
		} else if sv, ok := val.(WorkStatusRequestIntegrationValidatedDate); ok {
			// struct
			o.ValidatedDate = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegrationValidatedDate); ok {
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

// WorkStatusRequestRequestDate represents the object structure for request_date
type WorkStatusRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *WorkStatusRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequestRequestDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusRequest an agent action to request fetching issue statuses
type WorkStatusRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration WorkStatusRequestIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate WorkStatusRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WorkStatusRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WorkStatusRequest)(nil)

func toWorkStatusRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusRequest:
		return v.ToMap()

	case WorkStatusRequestIntegration:
		return v.ToMap()

	case WorkStatusRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of WorkStatusRequest
func (o *WorkStatusRequest) String() string {
	return fmt.Sprintf("agent.WorkStatusRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WorkStatusRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *WorkStatusRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WorkStatusRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WorkStatusRequest) GetModelName() datamodel.ModelNameType {
	return WorkStatusRequestModelName
}

// NewWorkStatusRequestID provides a template for generating an ID field for WorkStatusRequest
func NewWorkStatusRequestID(customerID string, refType string, refID string) string {
	return hash.Values("WorkStatusRequest", customerID, refType, refID)
}

func (o *WorkStatusRequest) setDefaults(frommap bool) {
	if o.Integration.EntityErrors == nil {
		o.Integration.EntityErrors = make([]WorkStatusRequestIntegrationEntityErrors, 0)
	}
	if o.Integration.ThrottledUntil == nil {
		o.Integration.ThrottledUntil = &WorkStatusRequestIntegrationThrottledUntil{}
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WorkStatusRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WorkStatusRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WorkStatusRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WorkStatusRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *WorkStatusRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WorkStatusRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WorkStatusRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WorkStatusRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WorkStatusRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *WorkStatusRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *WorkStatusRequest) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WorkStatusRequest
func (o *WorkStatusRequest) Clone() datamodel.Model {
	c := new(WorkStatusRequest)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WorkStatusRequest) Anon() datamodel.Model {
	c := new(WorkStatusRequest)
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
func (o *WorkStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WorkStatusRequest) UnmarshalJSON(data []byte) error {
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
func (o *WorkStatusRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WorkStatusRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two WorkStatusRequest objects are equal
func (o *WorkStatusRequest) IsEqual(other *WorkStatusRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WorkStatusRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toWorkStatusRequestObject(o.CustomerID, false),
		"id":           toWorkStatusRequestObject(o.ID, false),
		"integration":  toWorkStatusRequestObject(o.Integration, false),
		"ref_id":       toWorkStatusRequestObject(o.RefID, false),
		"ref_type":     toWorkStatusRequestObject(o.RefType, false),
		"request_date": toWorkStatusRequestObject(o.RequestDate, false),
		"uuid":         toWorkStatusRequestObject(o.UUID, false),
		"hashcode":     toWorkStatusRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(WorkStatusRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*WorkStatusRequestIntegration); ok {
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
		} else if sv, ok := val.(WorkStatusRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*WorkStatusRequestRequestDate); ok {
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
func (o *WorkStatusRequest) Hash() string {
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

// GetEventAPIConfig returns the EventAPIConfig
func (o *WorkStatusRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
