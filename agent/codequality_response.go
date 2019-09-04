// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CodequalityResponseTopic is the default topic name
	CodequalityResponseTopic datamodel.TopicNameType = "agent_CodequalityResponse_topic"

	// CodequalityResponseStream is the default stream name
	CodequalityResponseStream datamodel.TopicNameType = "agent_CodequalityResponse_stream"

	// CodequalityResponseTable is the default table name
	CodequalityResponseTable datamodel.TopicNameType = "agent_codequalityresponse"

	// CodequalityResponseModelName is the model name
	CodequalityResponseModelName datamodel.ModelNameType = "agent.CodequalityResponse"
)

const (
	// CodequalityResponseArchitectureColumn is the architecture column name
	CodequalityResponseArchitectureColumn = "architecture"
	// CodequalityResponseCustomerIDColumn is the customer_id column name
	CodequalityResponseCustomerIDColumn = "customer_id"
	// CodequalityResponseDataColumn is the data column name
	CodequalityResponseDataColumn = "data"
	// CodequalityResponseDistroColumn is the distro column name
	CodequalityResponseDistroColumn = "distro"
	// CodequalityResponseErrorColumn is the error column name
	CodequalityResponseErrorColumn = "error"
	// CodequalityResponseEventDateColumn is the event_date column name
	CodequalityResponseEventDateColumn = "event_date"
	// CodequalityResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	CodequalityResponseEventDateColumnEpochColumn = "event_date->epoch"
	// CodequalityResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	CodequalityResponseEventDateColumnOffsetColumn = "event_date->offset"
	// CodequalityResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	CodequalityResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// CodequalityResponseFreeSpaceColumn is the free_space column name
	CodequalityResponseFreeSpaceColumn = "free_space"
	// CodequalityResponseGoVersionColumn is the go_version column name
	CodequalityResponseGoVersionColumn = "go_version"
	// CodequalityResponseHostnameColumn is the hostname column name
	CodequalityResponseHostnameColumn = "hostname"
	// CodequalityResponseIDColumn is the id column name
	CodequalityResponseIDColumn = "id"
	// CodequalityResponseIntegrationIDColumn is the integration_id column name
	CodequalityResponseIntegrationIDColumn = "integration_id"
	// CodequalityResponseLastExportDateColumn is the last_export_date column name
	CodequalityResponseLastExportDateColumn = "last_export_date"
	// CodequalityResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnEpochColumn = "last_export_date->epoch"
	// CodequalityResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnOffsetColumn = "last_export_date->offset"
	// CodequalityResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	CodequalityResponseLastExportDateColumnRfc3339Column = "last_export_date->rfc3339"
	// CodequalityResponseMemoryColumn is the memory column name
	CodequalityResponseMemoryColumn = "memory"
	// CodequalityResponseMessageColumn is the message column name
	CodequalityResponseMessageColumn = "message"
	// CodequalityResponseNumCPUColumn is the num_cpu column name
	CodequalityResponseNumCPUColumn = "num_cpu"
	// CodequalityResponseOSColumn is the os column name
	CodequalityResponseOSColumn = "os"
	// CodequalityResponseProjectsColumn is the projects column name
	CodequalityResponseProjectsColumn = "projects"
	// CodequalityResponseProjectsColumnCustomerIDColumn is the customer_id column property of the Projects name
	CodequalityResponseProjectsColumnCustomerIDColumn = "projects->customer_id"
	// CodequalityResponseProjectsColumnIDColumn is the id column property of the Projects name
	CodequalityResponseProjectsColumnIDColumn = "projects->id"
	// CodequalityResponseProjectsColumnIdentifierColumn is the identifier column property of the Projects name
	CodequalityResponseProjectsColumnIdentifierColumn = "projects->identifier"
	// CodequalityResponseProjectsColumnNameColumn is the name column property of the Projects name
	CodequalityResponseProjectsColumnNameColumn = "projects->name"
	// CodequalityResponseProjectsColumnRefIDColumn is the ref_id column property of the Projects name
	CodequalityResponseProjectsColumnRefIDColumn = "projects->ref_id"
	// CodequalityResponseProjectsColumnRefTypeColumn is the ref_type column property of the Projects name
	CodequalityResponseProjectsColumnRefTypeColumn = "projects->ref_type"
	// CodequalityResponseRefIDColumn is the ref_id column name
	CodequalityResponseRefIDColumn = "ref_id"
	// CodequalityResponseRefTypeColumn is the ref_type column name
	CodequalityResponseRefTypeColumn = "ref_type"
	// CodequalityResponseRequestIDColumn is the request_id column name
	CodequalityResponseRequestIDColumn = "request_id"
	// CodequalityResponseSuccessColumn is the success column name
	CodequalityResponseSuccessColumn = "success"
	// CodequalityResponseSystemIDColumn is the system_id column name
	CodequalityResponseSystemIDColumn = "system_id"
	// CodequalityResponseTypeColumn is the type column name
	CodequalityResponseTypeColumn = "type"
	// CodequalityResponseUpdatedAtColumn is the updated_ts column name
	CodequalityResponseUpdatedAtColumn = "updated_ts"
	// CodequalityResponseUptimeColumn is the uptime column name
	CodequalityResponseUptimeColumn = "uptime"
	// CodequalityResponseUUIDColumn is the uuid column name
	CodequalityResponseUUIDColumn = "uuid"
	// CodequalityResponseVersionColumn is the version column name
	CodequalityResponseVersionColumn = "version"
)

// CodequalityResponseEventDate represents the object structure for event_date
type CodequalityResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityResponseEventDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityResponseEventDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityResponseEventDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *CodequalityResponseEventDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityResponseEventDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toCodequalityResponseEventDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityResponseEventDateObject(o.Rfc3339, isavro, false, "string"),
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
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCodequalityResponseLastExportDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityResponseLastExportDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityResponseLastExportDate:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *CodequalityResponseLastExportDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCodequalityResponseLastExportDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toCodequalityResponseLastExportDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCodequalityResponseLastExportDateObject(o.Rfc3339, isavro, false, "string"),
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
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Identifier the common identifier of the project
	Identifier string `json:"identifier" bson:"identifier" yaml:"identifier" faker:"-"`
	// Name the name of the project
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
}

func toCodequalityResponseProjectsObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityResponseProjectsObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityResponseProjects:
		return v.ToMap(isavro)

	default:
		return o
	}
}

func (o *CodequalityResponseProjects) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// CustomerID the customer id for the model instance
		"customer_id": toCodequalityResponseProjectsObject(o.CustomerID, isavro, false, "string"),
		// ID the primary key for the model instance
		"id": toCodequalityResponseProjectsObject(o.ID, isavro, false, "string"),
		// Identifier the common identifier of the project
		"identifier": toCodequalityResponseProjectsObject(o.Identifier, isavro, false, "string"),
		// Name the name of the project
		"name": toCodequalityResponseProjectsObject(o.Name, isavro, false, "string"),
		// RefID the source system id for the model instance
		"ref_id": toCodequalityResponseProjectsObject(o.RefID, isavro, false, "string"),
		// RefType the source system identifier for the model instance
		"ref_type": toCodequalityResponseProjectsObject(o.RefType, isavro, false, "string"),
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
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
	// EventDate the date of the event
	EventDate CodequalityResponseEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate CodequalityResponseLastExportDate `json:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" bson:"os" yaml:"os" faker:"-"`
	// Projects the codequality projects exported
	Projects []CodequalityResponseProjects `json:"projects" bson:"projects" yaml:"projects" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type CodequalityResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*CodequalityResponse)(nil)

func toCodequalityResponseObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toCodequalityResponseObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *CodequalityResponse:
		return v.ToMap(isavro)

	case CodequalityResponseEventDate:
		return v.ToMap(isavro)

	case CodequalityResponseLastExportDate:
		return v.ToMap(isavro)

	case []CodequalityResponseProjects:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap(isavro))
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

// GetStateKey returns a key for use in state store
func (o *CodequalityResponse) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *CodequalityResponse) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *CodequalityResponse) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
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

var cachedCodecCodequalityResponse *goavro.Codec
var cachedCodecCodequalityResponseLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *CodequalityResponse) GetAvroCodec() *goavro.Codec {
	cachedCodecCodequalityResponseLock.Lock()
	if cachedCodecCodequalityResponse == nil {
		c, err := GetCodequalityResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecCodequalityResponse = c
	}
	cachedCodecCodequalityResponseLock.Unlock()
	return cachedCodecCodequalityResponse
}

// ToAvroBinary returns the data as Avro binary data
func (o *CodequalityResponse) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *CodequalityResponse) FromAvroBinary(value []byte) error {
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
func (o *CodequalityResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two CodequalityResponse objects are equal
func (o *CodequalityResponse) IsEqual(other *CodequalityResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *CodequalityResponse) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Projects == nil {
			o.Projects = make([]CodequalityResponseProjects, 0)
		}
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toCodequalityResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":      toCodequalityResponseObject(o.CustomerID, isavro, false, "string"),
		"data":             toCodequalityResponseObject(o.Data, isavro, true, "string"),
		"distro":           toCodequalityResponseObject(o.Distro, isavro, false, "string"),
		"error":            toCodequalityResponseObject(o.Error, isavro, true, "string"),
		"event_date":       toCodequalityResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":       toCodequalityResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":       toCodequalityResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":         toCodequalityResponseObject(o.Hostname, isavro, false, "string"),
		"id":               toCodequalityResponseObject(o.ID, isavro, false, "string"),
		"integration_id":   toCodequalityResponseObject(o.IntegrationID, isavro, false, "string"),
		"last_export_date": toCodequalityResponseObject(o.LastExportDate, isavro, false, "last_export_date"),
		"memory":           toCodequalityResponseObject(o.Memory, isavro, false, "long"),
		"message":          toCodequalityResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":          toCodequalityResponseObject(o.NumCPU, isavro, false, "long"),
		"os":               toCodequalityResponseObject(o.OS, isavro, false, "string"),
		"projects":         toCodequalityResponseObject(o.Projects, isavro, false, "projects"),
		"ref_id":           toCodequalityResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":         toCodequalityResponseObject(o.RefType, isavro, false, "string"),
		"request_id":       toCodequalityResponseObject(o.RequestID, isavro, false, "string"),
		"success":          toCodequalityResponseObject(o.Success, isavro, false, "boolean"),
		"system_id":        toCodequalityResponseObject(o.SystemID, isavro, false, "string"),
		"type":             toCodequalityResponseObject(o.Type, isavro, false, "type"),
		"updated_ts":       toCodequalityResponseObject(o.UpdatedAt, isavro, false, "long"),
		"uptime":           toCodequalityResponseObject(o.Uptime, isavro, false, "long"),
		"uuid":             toCodequalityResponseObject(o.UUID, isavro, false, "string"),
		"version":          toCodequalityResponseObject(o.Version, isavro, false, "string"),
		"hashcode":         toCodequalityResponseObject(o.Hashcode, isavro, false, "string"),
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
				// if coming in as avro union, convert it back
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
				// if coming in as avro union, convert it back
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

// GetCodequalityResponseAvroSchemaSpec creates the avro schema specification for CodequalityResponse
func GetCodequalityResponseAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "CodequalityResponse",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "architecture",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "data",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "distro",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "error",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "event_date",
				"type": map[string]interface{}{"doc": "the date of the event", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "event_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "free_space",
				"type": "long",
			},
			map[string]interface{}{
				"name": "go_version",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hostname",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "last_export_date",
				"type": map[string]interface{}{"doc": "the last export date", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "last_export_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "memory",
				"type": "long",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "num_cpu",
				"type": "long",
			},
			map[string]interface{}{
				"name": "os",
				"type": "string",
			},
			map[string]interface{}{
				"name": "projects",
				"type": map[string]interface{}{"items": map[string]interface{}{"doc": "the codequality projects exported", "fields": []interface{}{map[string]interface{}{"doc": "the customer id for the model instance", "name": "customer_id", "type": "string"}, map[string]interface{}{"doc": "the primary key for the model instance", "name": "id", "type": "string"}, map[string]interface{}{"doc": "the common identifier of the project", "name": "identifier", "type": "string"}, map[string]interface{}{"doc": "the name of the project", "name": "name", "type": "string"}, map[string]interface{}{"doc": "the source system id for the model instance", "name": "ref_id", "type": "string"}, map[string]interface{}{"doc": "the source system identifier for the model instance", "name": "ref_type", "type": "string"}}, "name": "projects", "type": "record"}, "name": "projects", "type": "array"},
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
				"name": "request_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "success",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "system_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "type",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "type",
					"symbols": []interface{}{"ENROLL", "PING", "CRASH", "LOG", "INTEGRATION", "EXPORT", "PROJECT", "REPO", "USER", "UNINSTALL", "UPGRADE", "START", "STOP"},
				},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uptime",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
			map[string]interface{}{
				"name": "version",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
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

// GetCodequalityResponseAvroSchema creates the avro schema for CodequalityResponse
func GetCodequalityResponseAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetCodequalityResponseAvroSchemaSpec())
}

// TransformCodequalityResponseFunc is a function for transforming CodequalityResponse during processing
type TransformCodequalityResponseFunc func(input *CodequalityResponse) (*CodequalityResponse, error)

// NewCodequalityResponsePipe creates a pipe for processing CodequalityResponse items
func NewCodequalityResponsePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformCodequalityResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewCodequalityResponseInputStream(input, errors)
	var stream chan CodequalityResponse
	if len(transforms) > 0 {
		stream = make(chan CodequalityResponse, 1000)
	} else {
		stream = inch
	}
	outdone := NewCodequalityResponseOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewCodequalityResponseInputStreamDir creates a channel for reading CodequalityResponse as JSON newlines from a directory of files
func NewCodequalityResponseInputStreamDir(dir string, errors chan<- error, transforms ...TransformCodequalityResponseFunc) (chan CodequalityResponse, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/codequality_response\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan CodequalityResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for codequality_response")
		ch := make(chan CodequalityResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewCodequalityResponseInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan CodequalityResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewCodequalityResponseInputStreamFile creates an channel for reading CodequalityResponse as JSON newlines from filename
func NewCodequalityResponseInputStreamFile(filename string, errors chan<- error, transforms ...TransformCodequalityResponseFunc) (chan CodequalityResponse, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan CodequalityResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan CodequalityResponse)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewCodequalityResponseInputStream(f, errors, transforms...)
}

// NewCodequalityResponseInputStream creates an channel for reading CodequalityResponse as JSON newlines from stream
func NewCodequalityResponseInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformCodequalityResponseFunc) (chan CodequalityResponse, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan CodequalityResponse, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item CodequalityResponse
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewCodequalityResponseOutputStreamDir will output json newlines from channel and save in dir
func NewCodequalityResponseOutputStreamDir(dir string, ch chan CodequalityResponse, errors chan<- error, transforms ...TransformCodequalityResponseFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/codequality_response\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewCodequalityResponseOutputStream(gz, ch, errors, transforms...)
}

// NewCodequalityResponseOutputStream will output json newlines from channel to the stream
func NewCodequalityResponseOutputStream(stream io.WriteCloser, ch chan CodequalityResponse, errors chan<- error, transforms ...TransformCodequalityResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// CodequalityResponseSendEvent is an event detail for sending data
type CodequalityResponseSendEvent struct {
	CodequalityResponse *CodequalityResponse
	headers             map[string]string
	time                time.Time
	key                 string
}

var _ datamodel.ModelSendEvent = (*CodequalityResponseSendEvent)(nil)

// Key is the key to use for the message
func (e *CodequalityResponseSendEvent) Key() string {
	if e.key == "" {
		return e.CodequalityResponse.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *CodequalityResponseSendEvent) Object() datamodel.Model {
	return e.CodequalityResponse
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *CodequalityResponseSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *CodequalityResponseSendEvent) Timestamp() time.Time {
	return e.time
}

// CodequalityResponseSendEventOpts is a function handler for setting opts
type CodequalityResponseSendEventOpts func(o *CodequalityResponseSendEvent)

// WithCodequalityResponseSendEventKey sets the key value to a value different than the object ID
func WithCodequalityResponseSendEventKey(key string) CodequalityResponseSendEventOpts {
	return func(o *CodequalityResponseSendEvent) {
		o.key = key
	}
}

// WithCodequalityResponseSendEventTimestamp sets the timestamp value
func WithCodequalityResponseSendEventTimestamp(tv time.Time) CodequalityResponseSendEventOpts {
	return func(o *CodequalityResponseSendEvent) {
		o.time = tv
	}
}

// WithCodequalityResponseSendEventHeader sets the timestamp value
func WithCodequalityResponseSendEventHeader(key, value string) CodequalityResponseSendEventOpts {
	return func(o *CodequalityResponseSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewCodequalityResponseSendEvent returns a new CodequalityResponseSendEvent instance
func NewCodequalityResponseSendEvent(o *CodequalityResponse, opts ...CodequalityResponseSendEventOpts) *CodequalityResponseSendEvent {
	res := &CodequalityResponseSendEvent{
		CodequalityResponse: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewCodequalityResponseProducer will stream data from the channel
func NewCodequalityResponseProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
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
				if object, ok := item.Object().(*CodequalityResponse); ok {
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
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.CodequalityResponse but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewCodequalityResponseConsumer will stream data from the topic into the provided channel
func NewCodequalityResponseConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object CodequalityResponse
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.CodequalityResponse: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.CodequalityResponse: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.CodequalityResponse")
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

			ch <- &CodequalityResponseReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object CodequalityResponse
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &CodequalityResponseReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// CodequalityResponseReceiveEvent is an event detail for receiving data
type CodequalityResponseReceiveEvent struct {
	CodequalityResponse *CodequalityResponse
	message             eventing.Message
	eof                 bool
}

var _ datamodel.ModelReceiveEvent = (*CodequalityResponseReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *CodequalityResponseReceiveEvent) Object() datamodel.Model {
	return e.CodequalityResponse
}

// Message returns the underlying message data for the event
func (e *CodequalityResponseReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *CodequalityResponseReceiveEvent) EOF() bool {
	return e.eof
}

// CodequalityResponseProducer implements the datamodel.ModelEventProducer
type CodequalityResponseProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*CodequalityResponseProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *CodequalityResponseProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *CodequalityResponseProducer) Close() error {
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
func (o *CodequalityResponse) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *CodequalityResponse) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// NewCodequalityResponseProducerChannel returns a channel which can be used for producing Model events
func NewCodequalityResponseProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewCodequalityResponseProducerChannelSize(producer, 0, errors)
}

// NewCodequalityResponseProducerChannelSize returns a channel which can be used for producing Model events
func NewCodequalityResponseProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &CodequalityResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewCodequalityResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// CodequalityResponseConsumer implements the datamodel.ModelEventConsumer
type CodequalityResponseConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*CodequalityResponseConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *CodequalityResponseConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *CodequalityResponseConsumer) Close() error {
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
func (o *CodequalityResponse) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityResponseConsumer{
		ch:       ch,
		callback: NewCodequalityResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewCodequalityResponseConsumerChannel returns a consumer channel which can be used to consume Model events
func NewCodequalityResponseConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &CodequalityResponseConsumer{
		ch:       ch,
		callback: NewCodequalityResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
