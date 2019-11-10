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
	// ExportResponseTopic is the default topic name
	ExportResponseTopic datamodel.TopicNameType = "agent_ExportResponse_topic"

	// ExportResponseTable is the default table name
	ExportResponseTable datamodel.ModelNameType = "agent_exportresponse"

	// ExportResponseModelName is the model name
	ExportResponseModelName datamodel.ModelNameType = "agent.ExportResponse"
)

const (
	// ExportResponseArchitectureColumn is the architecture column name
	ExportResponseArchitectureColumn = "Architecture"
	// ExportResponseCustomerIDColumn is the customer_id column name
	ExportResponseCustomerIDColumn = "CustomerID"
	// ExportResponseDataColumn is the data column name
	ExportResponseDataColumn = "Data"
	// ExportResponseDistroColumn is the distro column name
	ExportResponseDistroColumn = "Distro"
	// ExportResponseEndDateColumn is the end_date column name
	ExportResponseEndDateColumn = "EndDate"
	// ExportResponseEndDateColumnEpochColumn is the epoch column property of the EndDate name
	ExportResponseEndDateColumnEpochColumn = "EndDate.Epoch"
	// ExportResponseEndDateColumnOffsetColumn is the offset column property of the EndDate name
	ExportResponseEndDateColumnOffsetColumn = "EndDate.Offset"
	// ExportResponseEndDateColumnRfc3339Column is the rfc3339 column property of the EndDate name
	ExportResponseEndDateColumnRfc3339Column = "EndDate.Rfc3339"
	// ExportResponseErrorColumn is the error column name
	ExportResponseErrorColumn = "Error"
	// ExportResponseEventDateColumn is the event_date column name
	ExportResponseEventDateColumn = "EventDate"
	// ExportResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	ExportResponseEventDateColumnEpochColumn = "EventDate.Epoch"
	// ExportResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	ExportResponseEventDateColumnOffsetColumn = "EventDate.Offset"
	// ExportResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	ExportResponseEventDateColumnRfc3339Column = "EventDate.Rfc3339"
	// ExportResponseExportTypeColumn is the export_type column name
	ExportResponseExportTypeColumn = "ExportType"
	// ExportResponseFreeSpaceColumn is the free_space column name
	ExportResponseFreeSpaceColumn = "FreeSpace"
	// ExportResponseGoVersionColumn is the go_version column name
	ExportResponseGoVersionColumn = "GoVersion"
	// ExportResponseHostnameColumn is the hostname column name
	ExportResponseHostnameColumn = "Hostname"
	// ExportResponseIDColumn is the id column name
	ExportResponseIDColumn = "ID"
	// ExportResponseIntegrationsColumn is the integrations column name
	ExportResponseIntegrationsColumn = "Integrations"
	// ExportResponseIntegrationsColumnExportTypeColumn is the export_type column property of the Integrations name
	ExportResponseIntegrationsColumnExportTypeColumn = "Integrations.ExportType"
	// ExportResponseIntegrationsColumnIntegrationIDColumn is the integration_id column property of the Integrations name
	ExportResponseIntegrationsColumnIntegrationIDColumn = "Integrations.IntegrationID"
	// ExportResponseIntegrationsColumnNameColumn is the name column property of the Integrations name
	ExportResponseIntegrationsColumnNameColumn = "Integrations.Name"
	// ExportResponseIntegrationsColumnSystemTypeColumn is the system_type column property of the Integrations name
	ExportResponseIntegrationsColumnSystemTypeColumn = "Integrations.SystemType"
	// ExportResponseJobIDColumn is the job_id column name
	ExportResponseJobIDColumn = "JobID"
	// ExportResponseLastExportDateColumn is the last_export_date column name
	ExportResponseLastExportDateColumn = "LastExportDate"
	// ExportResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	ExportResponseLastExportDateColumnEpochColumn = "LastExportDate.Epoch"
	// ExportResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	ExportResponseLastExportDateColumnOffsetColumn = "LastExportDate.Offset"
	// ExportResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	ExportResponseLastExportDateColumnRfc3339Column = "LastExportDate.Rfc3339"
	// ExportResponseMemoryColumn is the memory column name
	ExportResponseMemoryColumn = "Memory"
	// ExportResponseMessageColumn is the message column name
	ExportResponseMessageColumn = "Message"
	// ExportResponseNumCPUColumn is the num_cpu column name
	ExportResponseNumCPUColumn = "NumCPU"
	// ExportResponseOSColumn is the os column name
	ExportResponseOSColumn = "OS"
	// ExportResponseRefIDColumn is the ref_id column name
	ExportResponseRefIDColumn = "RefID"
	// ExportResponseRefTypeColumn is the ref_type column name
	ExportResponseRefTypeColumn = "RefType"
	// ExportResponseRequestIDColumn is the request_id column name
	ExportResponseRequestIDColumn = "RequestID"
	// ExportResponseSizeColumn is the size column name
	ExportResponseSizeColumn = "Size"
	// ExportResponseStartDateColumn is the start_date column name
	ExportResponseStartDateColumn = "StartDate"
	// ExportResponseStartDateColumnEpochColumn is the epoch column property of the StartDate name
	ExportResponseStartDateColumnEpochColumn = "StartDate.Epoch"
	// ExportResponseStartDateColumnOffsetColumn is the offset column property of the StartDate name
	ExportResponseStartDateColumnOffsetColumn = "StartDate.Offset"
	// ExportResponseStartDateColumnRfc3339Column is the rfc3339 column property of the StartDate name
	ExportResponseStartDateColumnRfc3339Column = "StartDate.Rfc3339"
	// ExportResponseStateColumn is the state column name
	ExportResponseStateColumn = "State"
	// ExportResponseSuccessColumn is the success column name
	ExportResponseSuccessColumn = "Success"
	// ExportResponseSystemIDColumn is the system_id column name
	ExportResponseSystemIDColumn = "SystemID"
	// ExportResponseSystemTypeColumn is the system_type column name
	ExportResponseSystemTypeColumn = "SystemType"
	// ExportResponseTypeColumn is the type column name
	ExportResponseTypeColumn = "Type"
	// ExportResponseUpdatedAtColumn is the updated_ts column name
	ExportResponseUpdatedAtColumn = "UpdatedAt"
	// ExportResponseUploadPartCountColumn is the upload_part_count column name
	ExportResponseUploadPartCountColumn = "UploadPartCount"
	// ExportResponseUploadURLColumn is the upload_url column name
	ExportResponseUploadURLColumn = "UploadURL"
	// ExportResponseUptimeColumn is the uptime column name
	ExportResponseUptimeColumn = "Uptime"
	// ExportResponseUUIDColumn is the uuid column name
	ExportResponseUUIDColumn = "UUID"
	// ExportResponseVersionColumn is the version column name
	ExportResponseVersionColumn = "Version"
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

const (
	// IntegrationsExportTypeHistorical is the enumeration value for historical
	ExportResponseIntegrationsExportTypeHistorical ExportResponseIntegrationsExportType = 0
	// IntegrationsExportTypeIncremental is the enumeration value for incremental
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

const (
	// IntegrationsSystemTypeWork is the enumeration value for work
	ExportResponseIntegrationsSystemTypeWork ExportResponseIntegrationsSystemType = 0
	// IntegrationsSystemTypeSourcecode is the enumeration value for sourcecode
	ExportResponseIntegrationsSystemTypeSourcecode ExportResponseIntegrationsSystemType = 1
	// IntegrationsSystemTypeCodequality is the enumeration value for codequality
	ExportResponseIntegrationsSystemTypeCodequality ExportResponseIntegrationsSystemType = 2
	// IntegrationsSystemTypeUser is the enumeration value for user
	ExportResponseIntegrationsSystemTypeUser ExportResponseIntegrationsSystemType = 3
)

// ExportResponseIntegrations represents the object structure for integrations
type ExportResponseIntegrations struct {
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

	case ExportResponseIntegrationsExportType:
		return v.String()

	case ExportResponseIntegrationsSystemType:
		return v.String()
	default:
		return o
	}
}

func (o *ExportResponseIntegrations) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
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

	if val, ok := kv["export_type"].(ExportResponseIntegrationsExportType); ok {
		o.ExportType = val
	} else {
		if em, ok := kv["export_type"].(map[string]interface{}); ok {
			ev := em["agent.export_type"].(string)
			switch ev {
			case "historical", "HISTORICAL":
				o.ExportType = 0
			case "incremental", "INCREMENTAL":
				o.ExportType = 1
			}
		}
		if em, ok := kv["export_type"].(string); ok {
			switch em {
			case "historical", "HISTORICAL":
				o.ExportType = 0
			case "incremental", "INCREMENTAL":
				o.ExportType = 1
			}
		}
	}

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["system_type"].(ExportResponseIntegrationsSystemType); ok {
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

const (
	// StateStarting is the enumeration value for starting
	ExportResponseStateStarting ExportResponseState = 0
	// StateProgress is the enumeration value for progress
	ExportResponseStateProgress ExportResponseState = 1
	// StateCompleted is the enumeration value for completed
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
		case "UNINSTALL":
			*v = ExportResponseType(9)
		case "UPGRADE":
			*v = ExportResponseType(10)
		case "START":
			*v = ExportResponseType(11)
		case "STOP":
			*v = ExportResponseType(12)
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
	case "UNINSTALL":
		v = 9
	case "UPGRADE":
		v = 10
	case "START":
		v = 11
	case "STOP":
		v = 12
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
		return json.Marshal("UNINSTALL")
	case 10:
		return json.Marshal("UPGRADE")
	case 11:
		return json.Marshal("START")
	case 12:
		return json.Marshal("STOP")
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
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	ExportResponseTypeEnroll ExportResponseType = 0
	// TypePing is the enumeration value for ping
	ExportResponseTypePing ExportResponseType = 1
	// TypeCrash is the enumeration value for crash
	ExportResponseTypeCrash ExportResponseType = 2
	// TypeLog is the enumeration value for log
	ExportResponseTypeLog ExportResponseType = 3
	// TypeIntegration is the enumeration value for integration
	ExportResponseTypeIntegration ExportResponseType = 4
	// TypeExport is the enumeration value for export
	ExportResponseTypeExport ExportResponseType = 5
	// TypeProject is the enumeration value for project
	ExportResponseTypeProject ExportResponseType = 6
	// TypeRepo is the enumeration value for repo
	ExportResponseTypeRepo ExportResponseType = 7
	// TypeUser is the enumeration value for user
	ExportResponseTypeUser ExportResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	ExportResponseTypeUninstall ExportResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	ExportResponseTypeUpgrade ExportResponseType = 10
	// TypeStart is the enumeration value for start
	ExportResponseTypeStart ExportResponseType = 11
	// TypeStop is the enumeration value for stop
	ExportResponseTypeStop ExportResponseType = 12
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
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
	return ExportResponseTopic
}

// GetStreamName returns the name of the stream
func (o *ExportResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ExportResponse) GetTableName() string {
	return ExportResponseTable.String()
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
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Integrations == nil {
		o.Integrations = make([]ExportResponseIntegrations, 0)
	}
	if o.UploadURL == nil {
		o.UploadURL = &emptyString
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
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ExportResponse) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for ExportResponse")
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
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ExportResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ExportResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		"updated_ts":        toExportResponseObject(o.UpdatedAt, false),
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EventDate.Epoch = dt.Epoch
				o.EventDate.Rfc3339 = dt.Rfc3339
				o.EventDate.Offset = dt.Offset
			}
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
					var av ExportResponseIntegrations
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for integrations field entry: " + reflect.TypeOf(ae).String())
					}
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportDate.Epoch = dt.Epoch
				o.LastExportDate.Rfc3339 = dt.Rfc3339
				o.LastExportDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.NumCPU = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Size = number.ToInt64Any(nil)
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
				o.Success = number.ToBoolAny(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
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
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			}
		}
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

	if val, ok := kv["upload_part_count"].(int64); ok {
		o.UploadPartCount = val
	} else {
		if val, ok := kv["upload_part_count"]; ok {
			if val == nil {
				o.UploadPartCount = number.ToInt64Any(nil)
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
				o.Uptime = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
	args = append(args, o.UpdatedAt)
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
