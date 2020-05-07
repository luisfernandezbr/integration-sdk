// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// ExportResponseTable is the default table name
	ExportResponseTable datamodel.ModelNameType = "agent_exportresponse"

	// ExportResponseModelName is the model name
	ExportResponseModelName datamodel.ModelNameType = "agent.ExportResponse"
)

const (
	// ExportResponseModelArchitectureColumn is the column json value architecture
	ExportResponseModelArchitectureColumn = "architecture"
	// ExportResponseModelCustomerIDColumn is the column json value customer_id
	ExportResponseModelCustomerIDColumn = "customer_id"
	// ExportResponseModelDataColumn is the column json value data
	ExportResponseModelDataColumn = "data"
	// ExportResponseModelDistroColumn is the column json value distro
	ExportResponseModelDistroColumn = "distro"
	// ExportResponseModelEndDateColumn is the column json value end_date
	ExportResponseModelEndDateColumn = "end_date"
	// ExportResponseModelEndDateEpochColumn is the column json value epoch
	ExportResponseModelEndDateEpochColumn = "epoch"
	// ExportResponseModelEndDateOffsetColumn is the column json value offset
	ExportResponseModelEndDateOffsetColumn = "offset"
	// ExportResponseModelEndDateRfc3339Column is the column json value rfc3339
	ExportResponseModelEndDateRfc3339Column = "rfc3339"
	// ExportResponseModelErrorColumn is the column json value error
	ExportResponseModelErrorColumn = "error"
	// ExportResponseModelEventDateColumn is the column json value event_date
	ExportResponseModelEventDateColumn = "event_date"
	// ExportResponseModelEventDateEpochColumn is the column json value epoch
	ExportResponseModelEventDateEpochColumn = "epoch"
	// ExportResponseModelEventDateOffsetColumn is the column json value offset
	ExportResponseModelEventDateOffsetColumn = "offset"
	// ExportResponseModelEventDateRfc3339Column is the column json value rfc3339
	ExportResponseModelEventDateRfc3339Column = "rfc3339"
	// ExportResponseModelFreeSpaceColumn is the column json value free_space
	ExportResponseModelFreeSpaceColumn = "free_space"
	// ExportResponseModelGoVersionColumn is the column json value go_version
	ExportResponseModelGoVersionColumn = "go_version"
	// ExportResponseModelHostnameColumn is the column json value hostname
	ExportResponseModelHostnameColumn = "hostname"
	// ExportResponseModelIDColumn is the column json value id
	ExportResponseModelIDColumn = "id"
	// ExportResponseModelIntegrationsColumn is the column json value integrations
	ExportResponseModelIntegrationsColumn = "integrations"
	// ExportResponseModelIntegrationsEntityErrorsColumn is the column json value entity_errors
	ExportResponseModelIntegrationsEntityErrorsColumn = "entity_errors"
	// ExportResponseModelIntegrationsEntityErrorsIDColumn is the column json value id
	ExportResponseModelIntegrationsEntityErrorsIDColumn = "id"
	// ExportResponseModelIntegrationsEntityErrorsRefIDColumn is the column json value ref_id
	ExportResponseModelIntegrationsEntityErrorsRefIDColumn = "ref_id"
	// ExportResponseModelIntegrationsEntityErrorsErrorColumn is the column json value error
	ExportResponseModelIntegrationsEntityErrorsErrorColumn = "error"
	// ExportResponseModelIntegrationsEntityErrorsWebhookErrorColumn is the column json value webhook_error
	ExportResponseModelIntegrationsEntityErrorsWebhookErrorColumn = "webhook_error"
	// ExportResponseModelIntegrationsErrorColumn is the column json value error
	ExportResponseModelIntegrationsErrorColumn = "error"
	// ExportResponseModelIntegrationsExportTypeColumn is the column json value export_type
	ExportResponseModelIntegrationsExportTypeColumn = "export_type"
	// ExportResponseModelIntegrationsIntegrationIDColumn is the column json value integration_id
	ExportResponseModelIntegrationsIntegrationIDColumn = "integration_id"
	// ExportResponseModelIntegrationsNameColumn is the column json value name
	ExportResponseModelIntegrationsNameColumn = "name"
	// ExportResponseModelIntegrationsSystemTypeColumn is the column json value system_type
	ExportResponseModelIntegrationsSystemTypeColumn = "system_type"
	// ExportResponseModelJobIDColumn is the column json value job_id
	ExportResponseModelJobIDColumn = "job_id"
	// ExportResponseModelLastExportDateColumn is the column json value last_export_date
	ExportResponseModelLastExportDateColumn = "last_export_date"
	// ExportResponseModelLastExportDateEpochColumn is the column json value epoch
	ExportResponseModelLastExportDateEpochColumn = "epoch"
	// ExportResponseModelLastExportDateOffsetColumn is the column json value offset
	ExportResponseModelLastExportDateOffsetColumn = "offset"
	// ExportResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	ExportResponseModelLastExportDateRfc3339Column = "rfc3339"
	// ExportResponseModelMemoryColumn is the column json value memory
	ExportResponseModelMemoryColumn = "memory"
	// ExportResponseModelMessageColumn is the column json value message
	ExportResponseModelMessageColumn = "message"
	// ExportResponseModelNumCPUColumn is the column json value num_cpu
	ExportResponseModelNumCPUColumn = "num_cpu"
	// ExportResponseModelOSColumn is the column json value os
	ExportResponseModelOSColumn = "os"
	// ExportResponseModelRefIDColumn is the column json value ref_id
	ExportResponseModelRefIDColumn = "ref_id"
	// ExportResponseModelRefTypeColumn is the column json value ref_type
	ExportResponseModelRefTypeColumn = "ref_type"
	// ExportResponseModelRequestIDColumn is the column json value request_id
	ExportResponseModelRequestIDColumn = "request_id"
	// ExportResponseModelSizeColumn is the column json value size
	ExportResponseModelSizeColumn = "size"
	// ExportResponseModelStartDateColumn is the column json value start_date
	ExportResponseModelStartDateColumn = "start_date"
	// ExportResponseModelStartDateEpochColumn is the column json value epoch
	ExportResponseModelStartDateEpochColumn = "epoch"
	// ExportResponseModelStartDateOffsetColumn is the column json value offset
	ExportResponseModelStartDateOffsetColumn = "offset"
	// ExportResponseModelStartDateRfc3339Column is the column json value rfc3339
	ExportResponseModelStartDateRfc3339Column = "rfc3339"
	// ExportResponseModelStateColumn is the column json value state
	ExportResponseModelStateColumn = "state"
	// ExportResponseModelSuccessColumn is the column json value success
	ExportResponseModelSuccessColumn = "success"
	// ExportResponseModelSystemIDColumn is the column json value system_id
	ExportResponseModelSystemIDColumn = "system_id"
	// ExportResponseModelTypeColumn is the column json value type
	ExportResponseModelTypeColumn = "type"
	// ExportResponseModelUploadPartCountColumn is the column json value upload_part_count
	ExportResponseModelUploadPartCountColumn = "upload_part_count"
	// ExportResponseModelUploadURLColumn is the column json value upload_url
	ExportResponseModelUploadURLColumn = "upload_url"
	// ExportResponseModelUptimeColumn is the column json value uptime
	ExportResponseModelUptimeColumn = "uptime"
	// ExportResponseModelUUIDColumn is the column json value uuid
	ExportResponseModelUUIDColumn = "uuid"
	// ExportResponseModelVersionColumn is the column json value version
	ExportResponseModelVersionColumn = "version"
)

// ExportResponseEndDate represents the object structure for end_date
type ExportResponseEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportResponseEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseEndDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportResponseEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportResponseEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportResponseEndDateObject(o.Rfc3339, false),
	}
}

func (o *ExportResponseEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseEndDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseEventDate represents the object structure for event_date
type ExportResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *ExportResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseEventDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseIntegrationsEntityErrors represents the object structure for entity_errors
type ExportResponseIntegrationsEntityErrors struct {
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// Error error message if set integration failed when trying to export this entity
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// WebhookError error message if set integration failed when installing webhook for this entity
	WebhookError string `json:"webhook_error" codec:"webhook_error" bson:"webhook_error" yaml:"webhook_error" faker:"-"`
}

func toExportResponseIntegrationsEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseIntegrationsEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseIntegrationsEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ID entity id
		"id": toExportResponseIntegrationsEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toExportResponseIntegrationsEntityErrorsObject(o.RefID, false),
		// Error error message if set integration failed when trying to export this entity
		"error": toExportResponseIntegrationsEntityErrorsObject(o.Error, false),
		// WebhookError error message if set integration failed when installing webhook for this entity
		"webhook_error": toExportResponseIntegrationsEntityErrorsObject(o.WebhookError, false),
	}
}

func (o *ExportResponseIntegrationsEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseIntegrationsEntityErrors) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseIntegrationsExportType is the enumeration type for export_type
type ExportResponseIntegrationsExportType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportResponseIntegrationsExportType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportResponseIntegrationsExportType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "HISTORICAL":
			*v = ExportResponseIntegrationsExportType(0)
		case "INCREMENTAL":
			*v = ExportResponseIntegrationsExportType(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportResponseIntegrationsExportType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "HISTORICAL":
		v = 0
	case "INCREMENTAL":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportResponseIntegrationsExportType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("HISTORICAL")
	case 1:
		return json.Marshal("INCREMENTAL")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationsExportType
func (v ExportResponseIntegrationsExportType) String() string {
	switch int32(v) {
	case 0:
		return "HISTORICAL"
	case 1:
		return "INCREMENTAL"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportResponseIntegrationsExportType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportResponseIntegrationsExportType:
		*v = val
	case int32:
		*v = ExportResponseIntegrationsExportType(int32(val))
	case int:
		*v = ExportResponseIntegrationsExportType(int32(val))
	case string:
		switch val {
		case "HISTORICAL":
			*v = ExportResponseIntegrationsExportType(0)
		case "INCREMENTAL":
			*v = ExportResponseIntegrationsExportType(1)
		}
	}
	return nil
}

const (
	// ExportResponseIntegrationsExportTypeHistorical is the enumeration value for historical
	ExportResponseIntegrationsExportTypeHistorical ExportResponseIntegrationsExportType = 0
	// ExportResponseIntegrationsExportTypeIncremental is the enumeration value for incremental
	ExportResponseIntegrationsExportTypeIncremental ExportResponseIntegrationsExportType = 1
)

// ExportResponseIntegrationsSystemType is the enumeration type for system_type
type ExportResponseIntegrationsSystemType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportResponseIntegrationsSystemType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportResponseIntegrationsSystemType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "WORK":
			*v = ExportResponseIntegrationsSystemType(0)
		case "SOURCECODE":
			*v = ExportResponseIntegrationsSystemType(1)
		case "CODEQUALITY":
			*v = ExportResponseIntegrationsSystemType(2)
		case "USER":
			*v = ExportResponseIntegrationsSystemType(3)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportResponseIntegrationsSystemType) UnmarshalJSON(buf []byte) error {
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

// MarshalJSON marshals the enum value
func (v ExportResponseIntegrationsSystemType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("WORK")
	case 1:
		return json.Marshal("SOURCECODE")
	case 2:
		return json.Marshal("CODEQUALITY")
	case 3:
		return json.Marshal("USER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationsSystemType
func (v ExportResponseIntegrationsSystemType) String() string {
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

// FromInterface for decoding from an interface
func (v *ExportResponseIntegrationsSystemType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportResponseIntegrationsSystemType:
		*v = val
	case int32:
		*v = ExportResponseIntegrationsSystemType(int32(val))
	case int:
		*v = ExportResponseIntegrationsSystemType(int32(val))
	case string:
		switch val {
		case "WORK":
			*v = ExportResponseIntegrationsSystemType(0)
		case "SOURCECODE":
			*v = ExportResponseIntegrationsSystemType(1)
		case "CODEQUALITY":
			*v = ExportResponseIntegrationsSystemType(2)
		case "USER":
			*v = ExportResponseIntegrationsSystemType(3)
		}
	}
	return nil
}

const (
	// ExportResponseIntegrationsSystemTypeWork is the enumeration value for work
	ExportResponseIntegrationsSystemTypeWork ExportResponseIntegrationsSystemType = 0
	// ExportResponseIntegrationsSystemTypeSourcecode is the enumeration value for sourcecode
	ExportResponseIntegrationsSystemTypeSourcecode ExportResponseIntegrationsSystemType = 1
	// ExportResponseIntegrationsSystemTypeCodequality is the enumeration value for codequality
	ExportResponseIntegrationsSystemTypeCodequality ExportResponseIntegrationsSystemType = 2
	// ExportResponseIntegrationsSystemTypeUser is the enumeration value for user
	ExportResponseIntegrationsSystemTypeUser ExportResponseIntegrationsSystemType = 3
)

// ExportResponseIntegrations represents the object structure for integrations
type ExportResponseIntegrations struct {
	// EntityErrors export status and error per entity in the integration
	EntityErrors []ExportResponseIntegrationsEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// Error if set integration failed with this error
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// ExportType the integration export type
	ExportType ExportResponseIntegrationsExportType `json:"export_type" codec:"export_type" bson:"export_type" yaml:"export_type" faker:"-"`
	// IntegrationID the id of the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// Name the friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
	SystemType ExportResponseIntegrationsSystemType `json:"system_type" codec:"system_type" bson:"system_type" yaml:"system_type" faker:"-"`
}

func toExportResponseIntegrationsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseIntegrations:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseIntegrations) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// EntityErrors export status and error per entity in the integration
		"entity_errors": toExportResponseIntegrationsObject(o.EntityErrors, false),
		// Error if set integration failed with this error
		"error": toExportResponseIntegrationsObject(o.Error, false),
		// ExportType the integration export type
		"export_type": toExportResponseIntegrationsObject(o.ExportType, false),
		// IntegrationID the id of the integration
		"integration_id": toExportResponseIntegrationsObject(o.IntegrationID, false),
		// Name the friendly name of the integration
		"name": toExportResponseIntegrationsObject(o.Name, false),
		// SystemType The system type of the integration (sourcecode / work (jira) / codequality / etc.)
		"system_type": toExportResponseIntegrationsObject(o.SystemType, false),
	}
}

func (o *ExportResponseIntegrations) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseIntegrations) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseLastExportDate represents the object structure for last_export_date
type ExportResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *ExportResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseLastExportDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseStartDate represents the object structure for start_date
type ExportResponseStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportResponseStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponseStartDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportResponseStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportResponseStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportResponseStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportResponseStartDateObject(o.Rfc3339, false),
	}
}

func (o *ExportResponseStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponseStartDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	o.setDefaults(false)
}

// ExportResponseState is the enumeration type for state
type ExportResponseState int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportResponseState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportResponseState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "STARTING":
			*v = ExportResponseState(0)
		case "PROGRESS":
			*v = ExportResponseState(1)
		case "COMPLETED":
			*v = ExportResponseState(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportResponseState) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "STARTING":
		v = 0
	case "PROGRESS":
		v = 1
	case "COMPLETED":
		v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportResponseState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("STARTING")
	case 1:
		return json.Marshal("PROGRESS")
	case 2:
		return json.Marshal("COMPLETED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v ExportResponseState) String() string {
	switch int32(v) {
	case 0:
		return "STARTING"
	case 1:
		return "PROGRESS"
	case 2:
		return "COMPLETED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportResponseState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportResponseState:
		*v = val
	case int32:
		*v = ExportResponseState(int32(val))
	case int:
		*v = ExportResponseState(int32(val))
	case string:
		switch val {
		case "STARTING":
			*v = ExportResponseState(0)
		case "PROGRESS":
			*v = ExportResponseState(1)
		case "COMPLETED":
			*v = ExportResponseState(2)
		}
	}
	return nil
}

const (
	// ExportResponseStateStarting is the enumeration value for starting
	ExportResponseStateStarting ExportResponseState = 0
	// ExportResponseStateProgress is the enumeration value for progress
	ExportResponseStateProgress ExportResponseState = 1
	// ExportResponseStateCompleted is the enumeration value for completed
	ExportResponseStateCompleted ExportResponseState = 2
)

// ExportResponseType is the enumeration type for type
type ExportResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *ExportResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = ExportResponseType(0)
		case "PING":
			*v = ExportResponseType(1)
		case "CRASH":
			*v = ExportResponseType(2)
		case "LOG":
			*v = ExportResponseType(3)
		case "INTEGRATION":
			*v = ExportResponseType(4)
		case "EXPORT":
			*v = ExportResponseType(5)
		case "PROJECT":
			*v = ExportResponseType(6)
		case "REPO":
			*v = ExportResponseType(7)
		case "USER":
			*v = ExportResponseType(8)
		case "CALENDAR":
			*v = ExportResponseType(9)
		case "UNINSTALL":
			*v = ExportResponseType(10)
		case "UPGRADE":
			*v = ExportResponseType(11)
		case "START":
			*v = ExportResponseType(12)
		case "STOP":
			*v = ExportResponseType(13)
		case "PAUSE":
			*v = ExportResponseType(14)
		case "RESUME":
			*v = ExportResponseType(15)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v ExportResponseType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENROLL":
		v = 0
	case "PING":
		v = 1
	case "CRASH":
		v = 2
	case "LOG":
		v = 3
	case "INTEGRATION":
		v = 4
	case "EXPORT":
		v = 5
	case "PROJECT":
		v = 6
	case "REPO":
		v = 7
	case "USER":
		v = 8
	case "CALENDAR":
		v = 9
	case "UNINSTALL":
		v = 10
	case "UPGRADE":
		v = 11
	case "START":
		v = 12
	case "STOP":
		v = 13
	case "PAUSE":
		v = 14
	case "RESUME":
		v = 15
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportResponseType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENROLL")
	case 1:
		return json.Marshal("PING")
	case 2:
		return json.Marshal("CRASH")
	case 3:
		return json.Marshal("LOG")
	case 4:
		return json.Marshal("INTEGRATION")
	case 5:
		return json.Marshal("EXPORT")
	case 6:
		return json.Marshal("PROJECT")
	case 7:
		return json.Marshal("REPO")
	case 8:
		return json.Marshal("USER")
	case 9:
		return json.Marshal("CALENDAR")
	case 10:
		return json.Marshal("UNINSTALL")
	case 11:
		return json.Marshal("UPGRADE")
	case 12:
		return json.Marshal("START")
	case 13:
		return json.Marshal("STOP")
	case 14:
		return json.Marshal("PAUSE")
	case 15:
		return json.Marshal("RESUME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Type
func (v ExportResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "CALENDAR"
	case 10:
		return "UNINSTALL"
	case 11:
		return "UPGRADE"
	case 12:
		return "START"
	case 13:
		return "STOP"
	case 14:
		return "PAUSE"
	case 15:
		return "RESUME"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportResponseType) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportResponseType:
		*v = val
	case int32:
		*v = ExportResponseType(int32(val))
	case int:
		*v = ExportResponseType(int32(val))
	case string:
		switch val {
		case "ENROLL":
			*v = ExportResponseType(0)
		case "PING":
			*v = ExportResponseType(1)
		case "CRASH":
			*v = ExportResponseType(2)
		case "LOG":
			*v = ExportResponseType(3)
		case "INTEGRATION":
			*v = ExportResponseType(4)
		case "EXPORT":
			*v = ExportResponseType(5)
		case "PROJECT":
			*v = ExportResponseType(6)
		case "REPO":
			*v = ExportResponseType(7)
		case "USER":
			*v = ExportResponseType(8)
		case "CALENDAR":
			*v = ExportResponseType(9)
		case "UNINSTALL":
			*v = ExportResponseType(10)
		case "UPGRADE":
			*v = ExportResponseType(11)
		case "START":
			*v = ExportResponseType(12)
		case "STOP":
			*v = ExportResponseType(13)
		case "PAUSE":
			*v = ExportResponseType(14)
		case "RESUME":
			*v = ExportResponseType(15)
		}
	}
	return nil
}

const (
	// ExportResponseTypeEnroll is the enumeration value for enroll
	ExportResponseTypeEnroll ExportResponseType = 0
	// ExportResponseTypePing is the enumeration value for ping
	ExportResponseTypePing ExportResponseType = 1
	// ExportResponseTypeCrash is the enumeration value for crash
	ExportResponseTypeCrash ExportResponseType = 2
	// ExportResponseTypeLog is the enumeration value for log
	ExportResponseTypeLog ExportResponseType = 3
	// ExportResponseTypeIntegration is the enumeration value for integration
	ExportResponseTypeIntegration ExportResponseType = 4
	// ExportResponseTypeExport is the enumeration value for export
	ExportResponseTypeExport ExportResponseType = 5
	// ExportResponseTypeProject is the enumeration value for project
	ExportResponseTypeProject ExportResponseType = 6
	// ExportResponseTypeRepo is the enumeration value for repo
	ExportResponseTypeRepo ExportResponseType = 7
	// ExportResponseTypeUser is the enumeration value for user
	ExportResponseTypeUser ExportResponseType = 8
	// ExportResponseTypeCalendar is the enumeration value for calendar
	ExportResponseTypeCalendar ExportResponseType = 9
	// ExportResponseTypeUninstall is the enumeration value for uninstall
	ExportResponseTypeUninstall ExportResponseType = 10
	// ExportResponseTypeUpgrade is the enumeration value for upgrade
	ExportResponseTypeUpgrade ExportResponseType = 11
	// ExportResponseTypeStart is the enumeration value for start
	ExportResponseTypeStart ExportResponseType = 12
	// ExportResponseTypeStop is the enumeration value for stop
	ExportResponseTypeStop ExportResponseType = 13
	// ExportResponseTypePause is the enumeration value for pause
	ExportResponseTypePause ExportResponseType = 14
	// ExportResponseTypeResume is the enumeration value for resume
	ExportResponseTypeResume ExportResponseType = 15
)

// ExportResponse an agent response to an action request for export
type ExportResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// EndDate the export end date
	EndDate ExportResponseEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate ExportResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integrations the integrations that were exported
	Integrations []ExportResponseIntegrations `json:"integrations" codec:"integrations" bson:"integrations" yaml:"integrations" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate ExportResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Size the size of the upload in bytes
	Size int64 `json:"size" codec:"size" bson:"size" yaml:"size" faker:"-"`
	// StartDate the export start date
	StartDate ExportResponseStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// State the state of the response
	State ExportResponseState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type ExportResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UploadPartCount the number of parts that the upload sent to the agent upload server
	UploadPartCount int64 `json:"upload_part_count" codec:"upload_part_count" bson:"upload_part_count" yaml:"upload_part_count" faker:"-"`
	// UploadURL the upload URL where the job was placed. will be NULL if not uploaded
	UploadURL *string `json:"upload_url,omitempty" codec:"upload_url,omitempty" bson:"upload_url" yaml:"upload_url,omitempty" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ExportResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ExportResponse)(nil)

func toExportResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportResponse:
		return v.ToMap()

	case ExportResponseEndDate:
		return v.ToMap()

	case ExportResponseEventDate:
		return v.ToMap()

	case []ExportResponseIntegrations:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case ExportResponseLastExportDate:
		return v.ToMap()

	case ExportResponseStartDate:
		return v.ToMap()

	case ExportResponseState:
		return v.String()

	case ExportResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of ExportResponse
func (o *ExportResponse) String() string {
	return fmt.Sprintf("agent.ExportResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ExportResponse) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *ExportResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *ExportResponse) GetModelName() datamodel.ModelNameType {
	return ExportResponseModelName
}

// NewExportResponseID provides a template for generating an ID field for ExportResponse
func NewExportResponseID(customerID string, refType string, refID string) string {
	return hash.Values("ExportResponse", customerID, refType, refID)
}

func (o *ExportResponse) setDefaults(frommap bool) {
	if o.Integrations == nil {
		o.Integrations = make([]ExportResponseIntegrations, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("ExportResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ExportResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ExportResponse) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *ExportResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ExportResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ExportResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ExportResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ExportResponse) IsEvented() bool {
	return false
}

// GetTopicConfig returns the topic config object
func (o *ExportResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *ExportResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ExportResponse
func (o *ExportResponse) Clone() datamodel.Model {
	c := new(ExportResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ExportResponse) Anon() datamodel.Model {
	c := new(ExportResponse)
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
func (o *ExportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportResponse) UnmarshalJSON(data []byte) error {
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
func (o *ExportResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ExportResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ExportResponse objects are equal
func (o *ExportResponse) IsEqual(other *ExportResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ExportResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toExportResponseObject(o.Architecture, false),
		"customer_id":      toExportResponseObject(o.CustomerID, false),
		"data":             toExportResponseObject(o.Data, true),
		"distro":           toExportResponseObject(o.Distro, false),
		"end_date":         toExportResponseObject(o.EndDate, false),
		"error":            toExportResponseObject(o.Error, true),
		"event_date":       toExportResponseObject(o.EventDate, false),
		"free_space":       toExportResponseObject(o.FreeSpace, false),
		"go_version":       toExportResponseObject(o.GoVersion, false),
		"hostname":         toExportResponseObject(o.Hostname, false),
		"id":               toExportResponseObject(o.ID, false),
		"integrations":     toExportResponseObject(o.Integrations, false),
		"job_id":           toExportResponseObject(o.JobID, false),
		"last_export_date": toExportResponseObject(o.LastExportDate, false),
		"memory":           toExportResponseObject(o.Memory, false),
		"message":          toExportResponseObject(o.Message, false),
		"num_cpu":          toExportResponseObject(o.NumCPU, false),
		"os":               toExportResponseObject(o.OS, false),
		"ref_id":           toExportResponseObject(o.RefID, false),
		"ref_type":         toExportResponseObject(o.RefType, false),
		"request_id":       toExportResponseObject(o.RequestID, false),
		"size":             toExportResponseObject(o.Size, false),
		"start_date":       toExportResponseObject(o.StartDate, false),

		"state":     o.State.String(),
		"success":   toExportResponseObject(o.Success, false),
		"system_id": toExportResponseObject(o.SystemID, false),

		"type":              o.Type.String(),
		"upload_part_count": toExportResponseObject(o.UploadPartCount, false),
		"upload_url":        toExportResponseObject(o.UploadURL, true),
		"uptime":            toExportResponseObject(o.Uptime, false),
		"uuid":              toExportResponseObject(o.UUID, false),
		"version":           toExportResponseObject(o.Version, false),
		"hashcode":          toExportResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Architecture = fmt.Sprintf("%v", val)
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

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Distro = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(ExportResponseEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*ExportResponseEndDate); ok {
			// struct pointer
			o.EndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EndDate.Epoch = dt.Epoch
				o.EndDate.Rfc3339 = dt.Rfc3339
				o.EndDate.Offset = dt.Offset
			}
		}
	} else {
		o.EndDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(ExportResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*ExportResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Hostname = fmt.Sprintf("%v", val)
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

	if o == nil {

		o.Integrations = make([]ExportResponseIntegrations, 0)

	}
	if val, ok := kv["integrations"]; ok {
		if sv, ok := val.([]ExportResponseIntegrations); ok {
			o.Integrations = sv
		} else if sp, ok := val.([]*ExportResponseIntegrations); ok {
			o.Integrations = o.Integrations[:0]
			for _, e := range sp {
				o.Integrations = append(o.Integrations, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ExportResponseIntegrations); ok {
					o.Integrations = append(o.Integrations, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ExportResponseIntegrations
					fm.FromMap(av)
					o.Integrations = append(o.Integrations, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ExportResponseIntegrations
					av.FromMap(bkv)
					o.Integrations = append(o.Integrations, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ExportResponseIntegrations); ok {
					o.Integrations = append(o.Integrations, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ExportResponseIntegrations
					fm.FromMap(r)
					o.Integrations = append(o.Integrations, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ExportResponseIntegrations{}
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
							var fm ExportResponseIntegrations
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

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(ExportResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*ExportResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.OS = fmt.Sprintf("%v", val)
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

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RequestID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		if val, ok := kv["size"]; ok {
			if val == nil {
				o.Size = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Size = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(ExportResponseStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*ExportResponseStartDate); ok {
			// struct pointer
			o.StartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.StartDate.Epoch = dt.Epoch
				o.StartDate.Rfc3339 = dt.Rfc3339
				o.StartDate.Offset = dt.Offset
			}
		}
	} else {
		o.StartDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["state"].(ExportResponseState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "starting", "STARTING":
				o.State = 0
			case "progress", "PROGRESS":
				o.State = 1
			case "completed", "COMPLETED":
				o.State = 2
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "starting", "STARTING":
				o.State = 0
			case "progress", "PROGRESS":
				o.State = 1
			case "completed", "COMPLETED":
				o.State = 2
			}
		}
	}

	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = false
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(ExportResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {

			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "calendar", "CALENDAR":
				o.Type = 9
			case "uninstall", "UNINSTALL":
				o.Type = 10
			case "upgrade", "UPGRADE":
				o.Type = 11
			case "start", "START":
				o.Type = 12
			case "stop", "STOP":
				o.Type = 13
			case "pause", "PAUSE":
				o.Type = 14
			case "resume", "RESUME":
				o.Type = 15
			}
		}
	}

	if val, ok := kv["upload_part_count"].(int64); ok {
		o.UploadPartCount = val
	} else {
		if val, ok := kv["upload_part_count"]; ok {
			if val == nil {
				o.UploadPartCount = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UploadPartCount = number.ToInt64Any(val)
			}
		}
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

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
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

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *ExportResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.EndDate)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.Integrations)
	args = append(args, o.JobID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Size)
	args = append(args, o.StartDate)
	args = append(args, o.State)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.UploadPartCount)
	args = append(args, o.UploadURL)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *ExportResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
