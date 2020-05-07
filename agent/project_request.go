// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

const (

	// ProjectRequestTable is the default table name
	ProjectRequestTable datamodel.ModelNameType = "agent_projectrequest"

	// ProjectRequestModelName is the model name
	ProjectRequestModelName datamodel.ModelNameType = "agent.ProjectRequest"
)

const (
	// ProjectRequestModelCustomerIDColumn is the column json value customer_id
	ProjectRequestModelCustomerIDColumn = "customer_id"
	// ProjectRequestModelIDColumn is the column json value id
	ProjectRequestModelIDColumn = "id"
	// ProjectRequestModelIntegrationColumn is the column json value integration
	ProjectRequestModelIntegrationColumn = "integration"
	// ProjectRequestModelIntegrationActiveColumn is the column json value active
	ProjectRequestModelIntegrationActiveColumn = "active"
	// ProjectRequestModelIntegrationAuthorizationColumn is the column json value authorization
	ProjectRequestModelIntegrationAuthorizationColumn = "authorization"
	// ProjectRequestModelIntegrationAuthorizationAccessTokenColumn is the column json value access_token
	ProjectRequestModelIntegrationAuthorizationAccessTokenColumn = "access_token"
	// ProjectRequestModelIntegrationAuthorizationRefreshTokenColumn is the column json value refresh_token
	ProjectRequestModelIntegrationAuthorizationRefreshTokenColumn = "refresh_token"
	// ProjectRequestModelIntegrationAuthorizationURLColumn is the column json value url
	ProjectRequestModelIntegrationAuthorizationURLColumn = "url"
	// ProjectRequestModelIntegrationAuthorizationUsernameColumn is the column json value username
	ProjectRequestModelIntegrationAuthorizationUsernameColumn = "username"
	// ProjectRequestModelIntegrationAuthorizationPasswordColumn is the column json value password
	ProjectRequestModelIntegrationAuthorizationPasswordColumn = "password"
	// ProjectRequestModelIntegrationAuthorizationAPITokenColumn is the column json value api_token
	ProjectRequestModelIntegrationAuthorizationAPITokenColumn = "api_token"
	// ProjectRequestModelIntegrationAuthorizationCollectionNameColumn is the column json value collection_name
	ProjectRequestModelIntegrationAuthorizationCollectionNameColumn = "collection_name"
	// ProjectRequestModelIntegrationAuthorizationAPIKeyColumn is the column json value api_key
	ProjectRequestModelIntegrationAuthorizationAPIKeyColumn = "api_key"
	// ProjectRequestModelIntegrationAuthorizationAuthorizationColumn is the column json value authorization
	ProjectRequestModelIntegrationAuthorizationAuthorizationColumn = "authorization"
	// ProjectRequestModelIntegrationAuthorizationHostnameColumn is the column json value hostname
	ProjectRequestModelIntegrationAuthorizationHostnameColumn = "hostname"
	// ProjectRequestModelIntegrationAuthorizationAPIVersionColumn is the column json value api_version
	ProjectRequestModelIntegrationAuthorizationAPIVersionColumn = "api_version"
	// ProjectRequestModelIntegrationAuthorizationOrganizationColumn is the column json value organization
	ProjectRequestModelIntegrationAuthorizationOrganizationColumn = "organization"
	// ProjectRequestModelIntegrationCreatedByProfileIDColumn is the column json value created_by_profile_id
	ProjectRequestModelIntegrationCreatedByProfileIDColumn = "created_by_profile_id"
	// ProjectRequestModelIntegrationCreatedByUserIDColumn is the column json value created_by_user_id
	ProjectRequestModelIntegrationCreatedByUserIDColumn = "created_by_user_id"
	// ProjectRequestModelIntegrationCustomerIDColumn is the column json value customer_id
	ProjectRequestModelIntegrationCustomerIDColumn = "customer_id"
	// ProjectRequestModelIntegrationEntityErrorsColumn is the column json value entity_errors
	ProjectRequestModelIntegrationEntityErrorsColumn = "entity_errors"
	// ProjectRequestModelIntegrationEntityErrorsIDColumn is the column json value id
	ProjectRequestModelIntegrationEntityErrorsIDColumn = "id"
	// ProjectRequestModelIntegrationEntityErrorsRefIDColumn is the column json value ref_id
	ProjectRequestModelIntegrationEntityErrorsRefIDColumn = "ref_id"
	// ProjectRequestModelIntegrationEntityErrorsErrorColumn is the column json value error
	ProjectRequestModelIntegrationEntityErrorsErrorColumn = "error"
	// ProjectRequestModelIntegrationErrorMessageColumn is the column json value error_message
	ProjectRequestModelIntegrationErrorMessageColumn = "error_message"
	// ProjectRequestModelIntegrationErroredColumn is the column json value errored
	ProjectRequestModelIntegrationErroredColumn = "errored"
	// ProjectRequestModelIntegrationExclusionsColumn is the column json value exclusions
	ProjectRequestModelIntegrationExclusionsColumn = "exclusions"
	// ProjectRequestModelIntegrationExportableColumn is the column json value exportable
	ProjectRequestModelIntegrationExportableColumn = "exportable"
	// ProjectRequestModelIntegrationIDColumn is the column json value id
	ProjectRequestModelIntegrationIDColumn = "id"
	// ProjectRequestModelIntegrationInclusionsColumn is the column json value inclusions
	ProjectRequestModelIntegrationInclusionsColumn = "inclusions"
	// ProjectRequestModelIntegrationLastExportCompletedDateColumn is the column json value last_export_completed_date
	ProjectRequestModelIntegrationLastExportCompletedDateColumn = "last_export_completed_date"
	// ProjectRequestModelIntegrationLastExportCompletedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationLastExportCompletedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationLastExportCompletedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationLastExportCompletedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationLastExportCompletedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationLastExportCompletedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationLastExportRequestedDateColumn is the column json value last_export_requested_date
	ProjectRequestModelIntegrationLastExportRequestedDateColumn = "last_export_requested_date"
	// ProjectRequestModelIntegrationLastExportRequestedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationLastExportRequestedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationLastExportRequestedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationLastExportRequestedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationLastExportRequestedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationLastExportRequestedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationLastProcessingCompletedDateColumn is the column json value last_processing_completed_date
	ProjectRequestModelIntegrationLastProcessingCompletedDateColumn = "last_processing_completed_date"
	// ProjectRequestModelIntegrationLastProcessingCompletedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationLastProcessingCompletedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationLastProcessingCompletedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationLastProcessingCompletedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationLastProcessingCompletedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationLastProcessingCompletedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationLastProcessingStartedDateColumn is the column json value last_processing_started_date
	ProjectRequestModelIntegrationLastProcessingStartedDateColumn = "last_processing_started_date"
	// ProjectRequestModelIntegrationLastProcessingStartedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationLastProcessingStartedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationLastProcessingStartedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationLastProcessingStartedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationLastProcessingStartedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationLastProcessingStartedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationLocationColumn is the column json value location
	ProjectRequestModelIntegrationLocationColumn = "location"
	// ProjectRequestModelIntegrationNameColumn is the column json value name
	ProjectRequestModelIntegrationNameColumn = "name"
	// ProjectRequestModelIntegrationOnboardCompletedDateColumn is the column json value onboard_completed_date
	ProjectRequestModelIntegrationOnboardCompletedDateColumn = "onboard_completed_date"
	// ProjectRequestModelIntegrationOnboardCompletedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationOnboardCompletedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationOnboardCompletedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationOnboardCompletedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationOnboardCompletedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationOnboardCompletedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationOnboardRequestedDateColumn is the column json value onboard_requested_date
	ProjectRequestModelIntegrationOnboardRequestedDateColumn = "onboard_requested_date"
	// ProjectRequestModelIntegrationOnboardRequestedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationOnboardRequestedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationOnboardRequestedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationOnboardRequestedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationOnboardRequestedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationOnboardRequestedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationOnboardingColumn is the column json value onboarding
	ProjectRequestModelIntegrationOnboardingColumn = "onboarding"
	// ProjectRequestModelIntegrationOrganizationColumn is the column json value organization
	ProjectRequestModelIntegrationOrganizationColumn = "organization"
	// ProjectRequestModelIntegrationProcessedColumn is the column json value processed
	ProjectRequestModelIntegrationProcessedColumn = "processed"
	// ProjectRequestModelIntegrationRefIDColumn is the column json value ref_id
	ProjectRequestModelIntegrationRefIDColumn = "ref_id"
	// ProjectRequestModelIntegrationRefTypeColumn is the column json value ref_type
	ProjectRequestModelIntegrationRefTypeColumn = "ref_type"
	// ProjectRequestModelIntegrationServerVersionColumn is the column json value server_version
	ProjectRequestModelIntegrationServerVersionColumn = "server_version"
	// ProjectRequestModelIntegrationStateColumn is the column json value state
	ProjectRequestModelIntegrationStateColumn = "state"
	// ProjectRequestModelIntegrationSystemTypeColumn is the column json value system_type
	ProjectRequestModelIntegrationSystemTypeColumn = "system_type"
	// ProjectRequestModelIntegrationTeamIDColumn is the column json value team_id
	ProjectRequestModelIntegrationTeamIDColumn = "team_id"
	// ProjectRequestModelIntegrationThrottledColumn is the column json value throttled
	ProjectRequestModelIntegrationThrottledColumn = "throttled"
	// ProjectRequestModelIntegrationThrottledUntilColumn is the column json value throttled_until
	ProjectRequestModelIntegrationThrottledUntilColumn = "throttled_until"
	// ProjectRequestModelIntegrationThrottledUntilEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationThrottledUntilEpochColumn = "epoch"
	// ProjectRequestModelIntegrationThrottledUntilOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationThrottledUntilOffsetColumn = "offset"
	// ProjectRequestModelIntegrationThrottledUntilRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationThrottledUntilRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationValidatedColumn is the column json value validated
	ProjectRequestModelIntegrationValidatedColumn = "validated"
	// ProjectRequestModelIntegrationValidatedDateColumn is the column json value validated_date
	ProjectRequestModelIntegrationValidatedDateColumn = "validated_date"
	// ProjectRequestModelIntegrationValidatedDateEpochColumn is the column json value epoch
	ProjectRequestModelIntegrationValidatedDateEpochColumn = "epoch"
	// ProjectRequestModelIntegrationValidatedDateOffsetColumn is the column json value offset
	ProjectRequestModelIntegrationValidatedDateOffsetColumn = "offset"
	// ProjectRequestModelIntegrationValidatedDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelIntegrationValidatedDateRfc3339Column = "rfc3339"
	// ProjectRequestModelIntegrationValidationMessageColumn is the column json value validation_message
	ProjectRequestModelIntegrationValidationMessageColumn = "validation_message"
	// ProjectRequestModelRefIDColumn is the column json value ref_id
	ProjectRequestModelRefIDColumn = "ref_id"
	// ProjectRequestModelRefTypeColumn is the column json value ref_type
	ProjectRequestModelRefTypeColumn = "ref_type"
	// ProjectRequestModelRequestDateColumn is the column json value request_date
	ProjectRequestModelRequestDateColumn = "request_date"
	// ProjectRequestModelRequestDateEpochColumn is the column json value epoch
	ProjectRequestModelRequestDateEpochColumn = "epoch"
	// ProjectRequestModelRequestDateOffsetColumn is the column json value offset
	ProjectRequestModelRequestDateOffsetColumn = "offset"
	// ProjectRequestModelRequestDateRfc3339Column is the column json value rfc3339
	ProjectRequestModelRequestDateRfc3339Column = "rfc3339"
	// ProjectRequestModelUUIDColumn is the column json value uuid
	ProjectRequestModelUUIDColumn = "uuid"
)

// ProjectRequestIntegrationAuthorization represents the object structure for authorization
type ProjectRequestIntegrationAuthorization struct {
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

func toProjectRequestIntegrationAuthorizationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationAuthorization:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationAuthorization) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AccessToken Access token
		"access_token": toProjectRequestIntegrationAuthorizationObject(o.AccessToken, true),
		// RefreshToken Refresh token
		"refresh_token": toProjectRequestIntegrationAuthorizationObject(o.RefreshToken, true),
		// URL URL of instance if relevant
		"url": toProjectRequestIntegrationAuthorizationObject(o.URL, true),
		// Username Username for instance, if relevant
		"username": toProjectRequestIntegrationAuthorizationObject(o.Username, true),
		// Password Password for instance, if relevant
		"password": toProjectRequestIntegrationAuthorizationObject(o.Password, true),
		// APIToken API Token for instance, if relevant
		"api_token": toProjectRequestIntegrationAuthorizationObject(o.APIToken, true),
		// CollectionName Collection name for instance, if relevant
		"collection_name": toProjectRequestIntegrationAuthorizationObject(o.CollectionName, true),
		// APIKey API Key for instance, if relevant
		"api_key": toProjectRequestIntegrationAuthorizationObject(o.APIKey, true),
		// Authorization the agents encrypted authorization
		"authorization": toProjectRequestIntegrationAuthorizationObject(o.Authorization, true),
		// Hostname Hostname for instance, if relevant
		"hostname": toProjectRequestIntegrationAuthorizationObject(o.Hostname, true),
		// APIVersion the api version of the integration
		"api_version": toProjectRequestIntegrationAuthorizationObject(o.APIVersion, true),
		// Organization Organization for instance, if relevant
		"organization": toProjectRequestIntegrationAuthorizationObject(o.Organization, true),
	}
}

func (o *ProjectRequestIntegrationAuthorization) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationAuthorization) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationEntityErrors represents the object structure for entity_errors
type ProjectRequestIntegrationEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
}

func toProjectRequestIntegrationEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toProjectRequestIntegrationEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toProjectRequestIntegrationEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export
		"error": toProjectRequestIntegrationEntityErrorsObject(o.Error, false),
	}
}

func (o *ProjectRequestIntegrationEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationEntityErrors) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationLastExportCompletedDate represents the object structure for last_export_completed_date
type ProjectRequestIntegrationLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationLastExportCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationLastExportRequestedDate represents the object structure for last_export_requested_date
type ProjectRequestIntegrationLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationLastExportRequestedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationLastProcessingCompletedDate represents the object structure for last_processing_completed_date
type ProjectRequestIntegrationLastProcessingCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationLastProcessingCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationLastProcessingCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationLastProcessingCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationLastProcessingCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationLastProcessingCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationLastProcessingCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationLastProcessingCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationLastProcessingCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationLastProcessingStartedDate represents the object structure for last_processing_started_date
type ProjectRequestIntegrationLastProcessingStartedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationLastProcessingStartedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationLastProcessingStartedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationLastProcessingStartedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationLastProcessingStartedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationLastProcessingStartedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationLastProcessingStartedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationLastProcessingStartedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationLastProcessingStartedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationLocation is the enumeration type for location
type ProjectRequestIntegrationLocation int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectRequestIntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectRequestIntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = ProjectRequestIntegrationLocation(0)
		case "CLOUD":
			*v = ProjectRequestIntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectRequestIntegrationLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ProjectRequestIntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationLocation
func (v ProjectRequestIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ProjectRequestIntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectRequestIntegrationLocation:
		*v = val
	case int32:
		*v = ProjectRequestIntegrationLocation(int32(val))
	case int:
		*v = ProjectRequestIntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = ProjectRequestIntegrationLocation(0)
		case "CLOUD":
			*v = ProjectRequestIntegrationLocation(1)
		}
	}
	return nil
}

const (
	// ProjectRequestIntegrationLocationPrivate is the enumeration value for private
	ProjectRequestIntegrationLocationPrivate ProjectRequestIntegrationLocation = 0
	// ProjectRequestIntegrationLocationCloud is the enumeration value for cloud
	ProjectRequestIntegrationLocationCloud ProjectRequestIntegrationLocation = 1
)

// ProjectRequestIntegrationOnboardCompletedDate represents the object structure for onboard_completed_date
type ProjectRequestIntegrationOnboardCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationOnboardCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationOnboardCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationOnboardCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationOnboardCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationOnboardCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationOnboardCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationOnboardCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationOnboardCompletedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationOnboardRequestedDate represents the object structure for onboard_requested_date
type ProjectRequestIntegrationOnboardRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationOnboardRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationOnboardRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationOnboardRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationOnboardRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationOnboardRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationOnboardRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationOnboardRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationOnboardRequestedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationState is the enumeration type for state
type ProjectRequestIntegrationState int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectRequestIntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectRequestIntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = ProjectRequestIntegrationState(0)
		case "EXPORTING":
			*v = ProjectRequestIntegrationState(1)
		case "PROCESSING":
			*v = ProjectRequestIntegrationState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectRequestIntegrationState) UnmarshalJSON(buf []byte) error {
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
func (v ProjectRequestIntegrationState) MarshalJSON() ([]byte, error) {
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
func (v ProjectRequestIntegrationState) String() string {
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
func (v *ProjectRequestIntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectRequestIntegrationState:
		*v = val
	case int32:
		*v = ProjectRequestIntegrationState(int32(val))
	case int:
		*v = ProjectRequestIntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = ProjectRequestIntegrationState(0)
		case "EXPORTING":
			*v = ProjectRequestIntegrationState(1)
		case "PROCESSING":
			*v = ProjectRequestIntegrationState(2)
		}
	}
	return nil
}

const (
	// ProjectRequestIntegrationStateIdle is the enumeration value for idle
	ProjectRequestIntegrationStateIdle ProjectRequestIntegrationState = 0
	// ProjectRequestIntegrationStateExporting is the enumeration value for exporting
	ProjectRequestIntegrationStateExporting ProjectRequestIntegrationState = 1
	// ProjectRequestIntegrationStateProcessing is the enumeration value for processing
	ProjectRequestIntegrationStateProcessing ProjectRequestIntegrationState = 2
)

// ProjectRequestIntegrationSystemType is the enumeration type for system_type
type ProjectRequestIntegrationSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ProjectRequestIntegrationSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ProjectRequestIntegrationSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = ProjectRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = ProjectRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = ProjectRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = ProjectRequestIntegrationSystemType(3)
		case "USER":
			*v = ProjectRequestIntegrationSystemType(4)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ProjectRequestIntegrationSystemType) UnmarshalJSON(buf []byte) error {
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
func (v ProjectRequestIntegrationSystemType) MarshalJSON() ([]byte, error) {
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
func (v ProjectRequestIntegrationSystemType) String() string {
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
func (v *ProjectRequestIntegrationSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ProjectRequestIntegrationSystemType:
		*v = val
	case int32:
		*v = ProjectRequestIntegrationSystemType(int32(val))
	case int:
		*v = ProjectRequestIntegrationSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = ProjectRequestIntegrationSystemType(0)
		case "SOURCECODE":
			*v = ProjectRequestIntegrationSystemType(1)
		case "CODEQUALITY":
			*v = ProjectRequestIntegrationSystemType(2)
		case "CALENDAR":
			*v = ProjectRequestIntegrationSystemType(3)
		case "USER":
			*v = ProjectRequestIntegrationSystemType(4)
		}
	}
	return nil
}

const (
	// ProjectRequestIntegrationSystemTypeWork is the enumeration value for work
	ProjectRequestIntegrationSystemTypeWork ProjectRequestIntegrationSystemType = 0
	// ProjectRequestIntegrationSystemTypeSourcecode is the enumeration value for sourcecode
	ProjectRequestIntegrationSystemTypeSourcecode ProjectRequestIntegrationSystemType = 1
	// ProjectRequestIntegrationSystemTypeCodequality is the enumeration value for codequality
	ProjectRequestIntegrationSystemTypeCodequality ProjectRequestIntegrationSystemType = 2
	// ProjectRequestIntegrationSystemTypeCalendar is the enumeration value for calendar
	ProjectRequestIntegrationSystemTypeCalendar ProjectRequestIntegrationSystemType = 3
	// ProjectRequestIntegrationSystemTypeUser is the enumeration value for user
	ProjectRequestIntegrationSystemTypeUser ProjectRequestIntegrationSystemType = 4
)

// ProjectRequestIntegrationThrottledUntil represents the object structure for throttled_until
type ProjectRequestIntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegrationValidatedDate represents the object structure for validated_date
type ProjectRequestIntegrationValidatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestIntegrationValidatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegrationValidatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegrationValidatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestIntegrationValidatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestIntegrationValidatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestIntegrationValidatedDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestIntegrationValidatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestIntegrationValidatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestIntegration represents the object structure for integration
type ProjectRequestIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Authorization Authorization information
	Authorization ProjectRequestIntegrationAuthorization `json:"authorization" codec:"authorization" bson:"authorization" yaml:"authorization" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []ProjectRequestIntegrationEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
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
	LastExportCompletedDate ProjectRequestIntegrationLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made (set by the backend)
	LastExportRequestedDate ProjectRequestIntegrationLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingCompletedDate when the processing completes (set by the pipeline)
	LastProcessingCompletedDate ProjectRequestIntegrationLastProcessingCompletedDate `json:"last_processing_completed_date" codec:"last_processing_completed_date" bson:"last_processing_completed_date" yaml:"last_processing_completed_date" faker:"-"`
	// LastProcessingStartedDate when the processing starts (set by the pipeline)
	LastProcessingStartedDate ProjectRequestIntegrationLastProcessingStartedDate `json:"last_processing_started_date" codec:"last_processing_started_date" bson:"last_processing_started_date" yaml:"last_processing_started_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location ProjectRequestIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OnboardCompletedDate when the last request for integration metadata was finished
	OnboardCompletedDate ProjectRequestIntegrationOnboardCompletedDate `json:"onboard_completed_date" codec:"onboard_completed_date" bson:"onboard_completed_date" yaml:"onboard_completed_date" faker:"-"`
	// OnboardRequestedDate when the last request for integration metadata was made
	OnboardRequestedDate ProjectRequestIntegrationOnboardRequestedDate `json:"onboard_requested_date" codec:"onboard_requested_date" bson:"onboard_requested_date" yaml:"onboard_requested_date" faker:"-"`
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
	State ProjectRequestIntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType ProjectRequestIntegrationSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
	// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
	TeamID *string `json:"team_id,omitempty" codec:"team_id,omitempty" bson:"team_id" yaml:"team_id,omitempty" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *ProjectRequestIntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// Validated If the validation has been run against this instance
	Validated *bool `json:"validated,omitempty" codec:"validated,omitempty" bson:"validated" yaml:"validated,omitempty" faker:"-"`
	// ValidatedDate Date when validated
	ValidatedDate ProjectRequestIntegrationValidatedDate `json:"validated_date" codec:"validated_date" bson:"validated_date" yaml:"validated_date" faker:"-"`
	// ValidationMessage The validation message from the agent
	ValidationMessage *string `json:"validation_message,omitempty" codec:"validation_message,omitempty" bson:"validation_message" yaml:"validation_message,omitempty" faker:"-"`
}

func toProjectRequestIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestIntegration:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toProjectRequestIntegrationObject(o.Active, false),
		// Authorization Authorization information
		"authorization": toProjectRequestIntegrationObject(o.Authorization, false),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toProjectRequestIntegrationObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toProjectRequestIntegrationObject(o.CreatedByUserID, true),
		// CustomerID the customer id for the model instance
		"customer_id": toProjectRequestIntegrationObject(o.CustomerID, false),
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toProjectRequestIntegrationObject(o.EntityErrors, false),
		// ErrorMessage The error message from an export run
		"error_message": toProjectRequestIntegrationObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent
		"errored": toProjectRequestIntegrationObject(o.Errored, true),
		// Exclusions The exclusion list for this integration
		"exclusions": toProjectRequestIntegrationObject(o.Exclusions, false),
		// Exportable a flag to indicate if the integration is ready for export
		"exportable": toProjectRequestIntegrationObject(o.Exportable, false),
		// ID the primary key for the model instance
		"id": toProjectRequestIntegrationObject(o.ID, false),
		// Inclusions The inclusion list for this integration
		"inclusions": toProjectRequestIntegrationObject(o.Inclusions, false),
		// LastExportCompletedDate when the export response was received (set by the backend)
		"last_export_completed_date": toProjectRequestIntegrationObject(o.LastExportCompletedDate, false),
		// LastExportRequestedDate when the export request was made (set by the backend)
		"last_export_requested_date": toProjectRequestIntegrationObject(o.LastExportRequestedDate, false),
		// LastProcessingCompletedDate when the processing completes (set by the pipeline)
		"last_processing_completed_date": toProjectRequestIntegrationObject(o.LastProcessingCompletedDate, false),
		// LastProcessingStartedDate when the processing starts (set by the pipeline)
		"last_processing_started_date": toProjectRequestIntegrationObject(o.LastProcessingStartedDate, false),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toProjectRequestIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toProjectRequestIntegrationObject(o.Name, false),
		// OnboardCompletedDate when the last request for integration metadata was finished
		"onboard_completed_date": toProjectRequestIntegrationObject(o.OnboardCompletedDate, false),
		// OnboardRequestedDate when the last request for integration metadata was made
		"onboard_requested_date": toProjectRequestIntegrationObject(o.OnboardRequestedDate, false),
		// Onboarding true if the agent is fetching metadata for the integration
		"onboarding": toProjectRequestIntegrationObject(o.Onboarding, false),
		// Organization The origanization authorized. Used for azure integrations
		"organization": toProjectRequestIntegrationObject(o.Organization, true),
		// Processed If the integration has been processed at least once
		"processed": toProjectRequestIntegrationObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toProjectRequestIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toProjectRequestIntegrationObject(o.RefType, false),
		// ServerVersion the server version for this integration
		"server_version": toProjectRequestIntegrationObject(o.ServerVersion, true),
		// State the current state of the integration
		"state": toProjectRequestIntegrationObject(o.State, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toProjectRequestIntegrationObject(o.SystemType, false),
		// TeamID The optional team_id for this integration. If set the integration is scoped to a specific team, otherwise global.
		"team_id": toProjectRequestIntegrationObject(o.TeamID, true),
		// Throttled Set to true when integration is throttled.
		"throttled": toProjectRequestIntegrationObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toProjectRequestIntegrationObject(o.ThrottledUntil, true),
		// Validated If the validation has been run against this instance
		"validated": toProjectRequestIntegrationObject(o.Validated, true),
		// ValidatedDate Date when validated
		"validated_date": toProjectRequestIntegrationObject(o.ValidatedDate, false),
		// ValidationMessage The validation message from the agent
		"validation_message": toProjectRequestIntegrationObject(o.ValidationMessage, true),
	}
}

func (o *ProjectRequestIntegration) setDefaults(frommap bool) {

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
func (o *ProjectRequestIntegration) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequestRequestDate represents the object structure for request_date
type ProjectRequestRequestDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toProjectRequestRequestDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ProjectRequestRequestDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toProjectRequestRequestDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toProjectRequestRequestDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toProjectRequestRequestDateObject(o.Rfc3339, false),
	}
}

func (o *ProjectRequestRequestDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequestRequestDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ProjectRequest an agent action to request adding new projects
type ProjectRequest struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration the integration details to use
	Integration ProjectRequestIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestDate the date when the request was made
	RequestDate ProjectRequestRequestDate `json:"request_date" codec:"request_date" bson:"request_date" yaml:"request_date" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ProjectRequest)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectRequest)(nil)

func toProjectRequestObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectRequest:
		return v.ToMap()

	case ProjectRequestIntegration:
		return v.ToMap()

	case ProjectRequestRequestDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ProjectRequest
func (o *ProjectRequest) String() string {
	return fmt.Sprintf("agent.ProjectRequest<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectRequest) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ProjectRequest) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectRequest) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ProjectRequest) GetModelName() datamodel.ModelNameType {
	return ProjectRequestModelName
}

// NewProjectRequestID provides a template for generating an ID field for ProjectRequest
func NewProjectRequestID(customerID string, refType string, refID string) string {
	return hash.Values("ProjectRequest", customerID, refType, refID)
}

func (o *ProjectRequest) setDefaults(frommap bool) {
	if o.Integration.EntityErrors == nil {
		o.Integration.EntityErrors = make([]ProjectRequestIntegrationEntityErrors, 0)
	}
	if o.Integration.Exclusions == nil {
		o.Integration.Exclusions = make([]string, 0)
	}
	if o.Integration.Inclusions == nil {
		o.Integration.Inclusions = make([]string, 0)
	}
	if o.Integration.ThrottledUntil == nil {
		o.Integration.ThrottledUntil = &ProjectRequestIntegrationThrottledUntil{}
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ProjectRequest", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectRequest) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectRequest) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectRequest) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ProjectRequest) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectRequest) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectRequest) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectRequest) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectRequest) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *ProjectRequest) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
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
	if idstr, ok := kv["id"].(string); ok {
		o.ID = idstr
	}
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *ProjectRequest) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ProjectRequest) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ProjectRequest objects are equal
func (o *ProjectRequest) IsEqual(other *ProjectRequest) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectRequest) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":  toProjectRequestObject(o.CustomerID, false),
		"id":           toProjectRequestObject(o.ID, false),
		"integration":  toProjectRequestObject(o.Integration, false),
		"ref_id":       toProjectRequestObject(o.RefID, false),
		"ref_type":     toProjectRequestObject(o.RefType, false),
		"request_date": toProjectRequestObject(o.RequestDate, false),
		"uuid":         toProjectRequestObject(o.UUID, false),
		"hashcode":     toProjectRequestObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectRequest) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(ProjectRequestIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*ProjectRequestIntegration); ok {
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
		} else if sv, ok := val.(ProjectRequestRequestDate); ok {
			// struct
			o.RequestDate = sv
		} else if sp, ok := val.(*ProjectRequestRequestDate); ok {
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
func (o *ProjectRequest) Hash() string {
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
func (o *ProjectRequest) GetEventAPIConfig() datamodel.EventAPIConfig {
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
