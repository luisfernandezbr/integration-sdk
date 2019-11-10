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
	// CodequalityResponseTopic is the default topic name
	CodequalityResponseTopic datamodel.TopicNameType = "agent_CodequalityResponse_topic"

	// CodequalityResponseTable is the default table name
	CodequalityResponseTable datamodel.ModelNameType = "agent_codequalityresponse"

	// CodequalityResponseModelName is the model name
	CodequalityResponseModelName datamodel.ModelNameType = "agent.CodequalityResponse"
)

const (
	// CodequalityResponseArchitectureColumn is the architecture column name
	CodequalityResponseArchitectureColumn = "Architecture"
	// CodequalityResponseCustomerIDColumn is the customer_id column name
	CodequalityResponseCustomerIDColumn = "CustomerID"
	// CodequalityResponseDataColumn is the data column name
	CodequalityResponseDataColumn = "Data"
	// CodequalityResponseDistroColumn is the distro column name
	CodequalityResponseDistroColumn = "Distro"
	// CodequalityResponseErrorColumn is the error column name
	CodequalityResponseErrorColumn = "Error"
	// CodequalityResponseEventDateColumn is the event_date column name
	CodequalityResponseEventDateColumn = "EventDate"
	// CodequalityResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	CodequalityResponseEventDateColumnEpochColumn = "EventDate.Epoch"
	// CodequalityResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	CodequalityResponseEventDateColumnOffsetColumn = "EventDate.Offset"
	// CodequalityResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	CodequalityResponseEventDateColumnRfc3339Column = "EventDate.Rfc3339"
	// CodequalityResponseFreeSpaceColumn is the free_space column name
	CodequalityResponseFreeSpaceColumn = "FreeSpace"
	// CodequalityResponseGoVersionColumn is the go_version column name
	CodequalityResponseGoVersionColumn = "GoVersion"
	// CodequalityResponseHostnameColumn is the hostname column name
	CodequalityResponseHostnameColumn = "Hostname"
	// CodequalityResponseIDColumn is the id column name
	CodequalityResponseIDColumn = "ID"
	// CodequalityResponseIntegrationIDColumn is the integration_id column name
	CodequalityResponseIntegrationIDColumn = "IntegrationID"
	// CodequalityResponseLastExportDateColumn is the last_export_date column name
	CodequalityResponseLastExportDateColumn = "LastExportDate"
	// CodequalityResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnEpochColumn = "LastExportDate.Epoch"
	// CodequalityResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnOffsetColumn = "LastExportDate.Offset"
	// CodequalityResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnRfc3339Column = "LastExportDate.Rfc3339"
	// CodequalityResponseMemoryColumn is the memory column name
	CodequalityResponseMemoryColumn = "Memory"
	// CodequalityResponseMessageColumn is the message column name
	CodequalityResponseMessageColumn = "Message"
	// CodequalityResponseNumCPUColumn is the num_cpu column name
	CodequalityResponseNumCPUColumn = "NumCPU"
	// CodequalityResponseOSColumn is the os column name
	CodequalityResponseOSColumn = "OS"
	// CodequalityResponseProjectsColumn is the projects column name
	CodequalityResponseProjectsColumn = "Projects"
	// CodequalityResponseProjectsColumnCustomerIDColumn is the customer_id column property of the Projects name
	CodequalityResponseProjectsColumnCustomerIDColumn = "Projects.CustomerID"
	// CodequalityResponseProjectsColumnIDColumn is the id column property of the Projects name
	CodequalityResponseProjectsColumnIDColumn = "Projects.ID"
	// CodequalityResponseProjectsColumnIdentifierColumn is the identifier column property of the Projects name
	CodequalityResponseProjectsColumnIdentifierColumn = "Projects.Identifier"
	// CodequalityResponseProjectsColumnNameColumn is the name column property of the Projects name
	CodequalityResponseProjectsColumnNameColumn = "Projects.Name"
	// CodequalityResponseProjectsColumnRefIDColumn is the ref_id column property of the Projects name
	CodequalityResponseProjectsColumnRefIDColumn = "Projects.RefID"
	// CodequalityResponseProjectsColumnRefTypeColumn is the ref_type column property of the Projects name
	CodequalityResponseProjectsColumnRefTypeColumn = "Projects.RefType"
	// CodequalityResponseRefIDColumn is the ref_id column name
	CodequalityResponseRefIDColumn = "RefID"
	// CodequalityResponseRefTypeColumn is the ref_type column name
	CodequalityResponseRefTypeColumn = "RefType"
	// CodequalityResponseRequestIDColumn is the request_id column name
	CodequalityResponseRequestIDColumn = "RequestID"
	// CodequalityResponseSuccessColumn is the success column name
	CodequalityResponseSuccessColumn = "Success"
	// CodequalityResponseSystemIDColumn is the system_id column name
	CodequalityResponseSystemIDColumn = "SystemID"
	// CodequalityResponseTypeColumn is the type column name
	CodequalityResponseTypeColumn = "Type"
	// CodequalityResponseUpdatedAtColumn is the updated_ts column name
	CodequalityResponseUpdatedAtColumn = "UpdatedAt"
	// CodequalityResponseUptimeColumn is the uptime column name
	CodequalityResponseUptimeColumn = "Uptime"
	// CodequalityResponseUUIDColumn is the uuid column name
	CodequalityResponseUUIDColumn = "UUID"
	// CodequalityResponseVersionColumn is the version column name
	CodequalityResponseVersionColumn = "Version"
)

// CodequalityResponseEventDate represents the object structure for event_date
type CodequalityResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityResponseEventDate) FromMap(kv map[string]interface{}) {

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

// CodequalityResponseLastExportDate represents the object structure for last_export_date
type CodequalityResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCodequalityResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *CodequalityResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// CodequalityResponseProjects represents the object structure for projects
type CodequalityResponseProjects struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"id" yaml:"id" faker:"-"`
	// Identifier the common identifier of the project
	Identifier string `json:"identifier" codec:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// Name the name of the project
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
}

func toCodequalityResponseProjectsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityResponseProjects:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CodequalityResponseProjects) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// CustomerID the customer id for the model instance
		"customer_id": toCodequalityResponseProjectsObject(o.CustomerID, false),
		// ID the primary key for the model instance
		"id": toCodequalityResponseProjectsObject(o.ID, false),
		// Identifier the common identifier of the project
		"identifier": toCodequalityResponseProjectsObject(o.Identifier, false),
		// Name the name of the project
		"name": toCodequalityResponseProjectsObject(o.Name, false),
		// RefID the source system id for the model instance
		"ref_id": toCodequalityResponseProjectsObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toCodequalityResponseProjectsObject(o.RefType, false),
	}
}

func (o *CodequalityResponseProjects) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityResponseProjects) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["identifier"].(string); ok {
		o.Identifier = val
	} else {
		if val, ok := kv["identifier"]; ok {
			if val == nil {
				o.Identifier = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Identifier = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// CodequalityResponseType is the enumeration type for type
type CodequalityResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *CodequalityResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = CodequalityResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = CodequalityResponseType(0)
		case "PING":
			*v = CodequalityResponseType(1)
		case "CRASH":
			*v = CodequalityResponseType(2)
		case "LOG":
			*v = CodequalityResponseType(3)
		case "INTEGRATION":
			*v = CodequalityResponseType(4)
		case "EXPORT":
			*v = CodequalityResponseType(5)
		case "PROJECT":
			*v = CodequalityResponseType(6)
		case "REPO":
			*v = CodequalityResponseType(7)
		case "USER":
			*v = CodequalityResponseType(8)
		case "UNINSTALL":
			*v = CodequalityResponseType(9)
		case "UPGRADE":
			*v = CodequalityResponseType(10)
		case "START":
			*v = CodequalityResponseType(11)
		case "STOP":
			*v = CodequalityResponseType(12)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v CodequalityResponseType) UnmarshalJSON(buf []byte) error {
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
func (v CodequalityResponseType) MarshalJSON() ([]byte, error) {
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
func (v CodequalityResponseType) String() string {
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
	CodequalityResponseTypeEnroll CodequalityResponseType = 0
	// TypePing is the enumeration value for ping
	CodequalityResponseTypePing CodequalityResponseType = 1
	// TypeCrash is the enumeration value for crash
	CodequalityResponseTypeCrash CodequalityResponseType = 2
	// TypeLog is the enumeration value for log
	CodequalityResponseTypeLog CodequalityResponseType = 3
	// TypeIntegration is the enumeration value for integration
	CodequalityResponseTypeIntegration CodequalityResponseType = 4
	// TypeExport is the enumeration value for export
	CodequalityResponseTypeExport CodequalityResponseType = 5
	// TypeProject is the enumeration value for project
	CodequalityResponseTypeProject CodequalityResponseType = 6
	// TypeRepo is the enumeration value for repo
	CodequalityResponseTypeRepo CodequalityResponseType = 7
	// TypeUser is the enumeration value for user
	CodequalityResponseTypeUser CodequalityResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	CodequalityResponseTypeUninstall CodequalityResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	CodequalityResponseTypeUpgrade CodequalityResponseType = 10
	// TypeStart is the enumeration value for start
	CodequalityResponseTypeStart CodequalityResponseType = 11
	// TypeStop is the enumeration value for stop
	CodequalityResponseTypeStop CodequalityResponseType = 12
)

// CodequalityResponse an agent response to an action request adding codequality entities
type CodequalityResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" codec:"data,omitempty" bson:"data" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" codec:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" codec:"error,omitempty" bson:"error" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate CodequalityResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" codec:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate CodequalityResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" codec:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// Projects the codequality projects exported
	Projects []CodequalityResponseProjects `json:"projects" codec:"projects" bson:"projects" yaml:"projects" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" codec:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type CodequalityResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
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
var _ datamodel.Model = (*CodequalityResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*CodequalityResponse)(nil)

func toCodequalityResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CodequalityResponse:
		return v.ToMap()

	case CodequalityResponseEventDate:
		return v.ToMap()

	case CodequalityResponseLastExportDate:
		return v.ToMap()

	case []CodequalityResponseProjects:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case CodequalityResponseType:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of CodequalityResponse
func (o *CodequalityResponse) String() string {
	return fmt.Sprintf("agent.CodequalityResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *CodequalityResponse) GetTopicName() datamodel.TopicNameType {
	return CodequalityResponseTopic
}

// GetStreamName returns the name of the stream
func (o *CodequalityResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *CodequalityResponse) GetTableName() string {
	return CodequalityResponseTable.String()
}

// GetModelName returns the name of the model
func (o *CodequalityResponse) GetModelName() datamodel.ModelNameType {
	return CodequalityResponseModelName
}

// NewCodequalityResponseID provides a template for generating an ID field for CodequalityResponse
func NewCodequalityResponseID(customerID string, refType string, refID string) string {
	return hash.Values("CodequalityResponse", customerID, refType, refID)
}

func (o *CodequalityResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Projects == nil {
		o.Projects = make([]CodequalityResponseProjects, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("CodequalityResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *CodequalityResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *CodequalityResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *CodequalityResponse) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for CodequalityResponse")
}

// GetRefID returns the RefID for the object
func (o *CodequalityResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *CodequalityResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *CodequalityResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *CodequalityResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *CodequalityResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *CodequalityResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CodequalityResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *CodequalityResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *CodequalityResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of CodequalityResponse
func (o *CodequalityResponse) Clone() datamodel.Model {
	c := new(CodequalityResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *CodequalityResponse) Anon() datamodel.Model {
	c := new(CodequalityResponse)
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
func (o *CodequalityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *CodequalityResponse) UnmarshalJSON(data []byte) error {
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
func (o *CodequalityResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CodequalityResponse objects are equal
func (o *CodequalityResponse) IsEqual(other *CodequalityResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toCodequalityResponseObject(o.Architecture, false),
		"customer_id":      toCodequalityResponseObject(o.CustomerID, false),
		"data":             toCodequalityResponseObject(o.Data, true),
		"distro":           toCodequalityResponseObject(o.Distro, false),
		"error":            toCodequalityResponseObject(o.Error, true),
		"event_date":       toCodequalityResponseObject(o.EventDate, false),
		"free_space":       toCodequalityResponseObject(o.FreeSpace, false),
		"go_version":       toCodequalityResponseObject(o.GoVersion, false),
		"hostname":         toCodequalityResponseObject(o.Hostname, false),
		"id":               toCodequalityResponseObject(o.ID, false),
		"integration_id":   toCodequalityResponseObject(o.IntegrationID, false),
		"last_export_date": toCodequalityResponseObject(o.LastExportDate, false),
		"memory":           toCodequalityResponseObject(o.Memory, false),
		"message":          toCodequalityResponseObject(o.Message, false),
		"num_cpu":          toCodequalityResponseObject(o.NumCPU, false),
		"os":               toCodequalityResponseObject(o.OS, false),
		"projects":         toCodequalityResponseObject(o.Projects, false),
		"ref_id":           toCodequalityResponseObject(o.RefID, false),
		"ref_type":         toCodequalityResponseObject(o.RefType, false),
		"request_id":       toCodequalityResponseObject(o.RequestID, false),
		"success":          toCodequalityResponseObject(o.Success, false),
		"system_id":        toCodequalityResponseObject(o.SystemID, false),

		"type":       o.Type.String(),
		"updated_ts": toCodequalityResponseObject(o.UpdatedAt, false),
		"uptime":     toCodequalityResponseObject(o.Uptime, false),
		"uuid":       toCodequalityResponseObject(o.UUID, false),
		"version":    toCodequalityResponseObject(o.Version, false),
		"hashcode":   toCodequalityResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *CodequalityResponse) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(CodequalityResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*CodequalityResponseEventDate); ok {
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

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(CodequalityResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*CodequalityResponseLastExportDate); ok {
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

	if o == nil {

		o.Projects = make([]CodequalityResponseProjects, 0)

	}
	if val, ok := kv["projects"]; ok {
		if sv, ok := val.([]CodequalityResponseProjects); ok {
			o.Projects = sv
		} else if sp, ok := val.([]*CodequalityResponseProjects); ok {
			o.Projects = o.Projects[:0]
			for _, e := range sp {
				o.Projects = append(o.Projects, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(CodequalityResponseProjects); ok {
					o.Projects = append(o.Projects, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm CodequalityResponseProjects
					fm.FromMap(av)
					o.Projects = append(o.Projects, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av CodequalityResponseProjects
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for projects field entry: " + reflect.TypeOf(ae).String())
					}
					o.Projects = append(o.Projects, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(CodequalityResponseProjects); ok {
					o.Projects = append(o.Projects, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm CodequalityResponseProjects
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := CodequalityResponseProjects{}
					fm.FromMap(r)
					o.Projects = append(o.Projects, fm)
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
							var fm CodequalityResponseProjects
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Projects = append(o.Projects, fm)
						}
					}
				}
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

	if val, ok := kv["type"].(CodequalityResponseType); ok {
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
func (o *CodequalityResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.Projects)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *CodequalityResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
