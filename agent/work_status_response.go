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
	// WorkStatusResponseTopic is the default topic name
	WorkStatusResponseTopic datamodel.TopicNameType = "agent_WorkStatusResponse_topic"

	// WorkStatusResponseModelName is the model name
	WorkStatusResponseModelName datamodel.ModelNameType = "agent.WorkStatusResponse"
)

const (
	// WorkStatusResponseArchitectureColumn is the architecture column name
	WorkStatusResponseArchitectureColumn = "Architecture"
	// WorkStatusResponseCustomerIDColumn is the customer_id column name
	WorkStatusResponseCustomerIDColumn = "CustomerID"
	// WorkStatusResponseDataColumn is the data column name
	WorkStatusResponseDataColumn = "Data"
	// WorkStatusResponseDistroColumn is the distro column name
	WorkStatusResponseDistroColumn = "Distro"
	// WorkStatusResponseErrorColumn is the error column name
	WorkStatusResponseErrorColumn = "Error"
	// WorkStatusResponseEventDateColumn is the event_date column name
	WorkStatusResponseEventDateColumn = "EventDate"
	// WorkStatusResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	WorkStatusResponseEventDateColumnEpochColumn = "EventDate.Epoch"
	// WorkStatusResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	WorkStatusResponseEventDateColumnOffsetColumn = "EventDate.Offset"
	// WorkStatusResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	WorkStatusResponseEventDateColumnRfc3339Column = "EventDate.Rfc3339"
	// WorkStatusResponseFieldColumn is the field column name
	WorkStatusResponseFieldColumn = "Field"
	// WorkStatusResponseFreeSpaceColumn is the free_space column name
	WorkStatusResponseFreeSpaceColumn = "FreeSpace"
	// WorkStatusResponseGoVersionColumn is the go_version column name
	WorkStatusResponseGoVersionColumn = "GoVersion"
	// WorkStatusResponseHostnameColumn is the hostname column name
	WorkStatusResponseHostnameColumn = "Hostname"
	// WorkStatusResponseIDColumn is the id column name
	WorkStatusResponseIDColumn = "ID"
	// WorkStatusResponseIntegrationIDColumn is the integration_id column name
	WorkStatusResponseIntegrationIDColumn = "IntegrationID"
	// WorkStatusResponseIssueTypeColumn is the issue_type column name
	WorkStatusResponseIssueTypeColumn = "IssueType"
	// WorkStatusResponseLastExportDateColumn is the last_export_date column name
	WorkStatusResponseLastExportDateColumn = "LastExportDate"
	// WorkStatusResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	WorkStatusResponseLastExportDateColumnEpochColumn = "LastExportDate.Epoch"
	// WorkStatusResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	WorkStatusResponseLastExportDateColumnOffsetColumn = "LastExportDate.Offset"
	// WorkStatusResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	WorkStatusResponseLastExportDateColumnRfc3339Column = "LastExportDate.Rfc3339"
	// WorkStatusResponseMemoryColumn is the memory column name
	WorkStatusResponseMemoryColumn = "Memory"
	// WorkStatusResponseMessageColumn is the message column name
	WorkStatusResponseMessageColumn = "Message"
	// WorkStatusResponseNumCPUColumn is the num_cpu column name
	WorkStatusResponseNumCPUColumn = "NumCPU"
	// WorkStatusResponseOperatorColumn is the operator column name
	WorkStatusResponseOperatorColumn = "Operator"
	// WorkStatusResponseOSColumn is the os column name
	WorkStatusResponseOSColumn = "OS"
	// WorkStatusResponseRefIDColumn is the ref_id column name
	WorkStatusResponseRefIDColumn = "RefID"
	// WorkStatusResponseRefTypeColumn is the ref_type column name
	WorkStatusResponseRefTypeColumn = "RefType"
	// WorkStatusResponseRequestIDColumn is the request_id column name
	WorkStatusResponseRequestIDColumn = "RequestID"
	// WorkStatusResponseSuccessColumn is the success column name
	WorkStatusResponseSuccessColumn = "Success"
	// WorkStatusResponseSystemIDColumn is the system_id column name
	WorkStatusResponseSystemIDColumn = "SystemID"
	// WorkStatusResponseTypeColumn is the type column name
	WorkStatusResponseTypeColumn = "Type"
	// WorkStatusResponseUpdatedAtColumn is the updated_ts column name
	WorkStatusResponseUpdatedAtColumn = "UpdatedAt"
	// WorkStatusResponseUptimeColumn is the uptime column name
	WorkStatusResponseUptimeColumn = "Uptime"
	// WorkStatusResponseUUIDColumn is the uuid column name
	WorkStatusResponseUUIDColumn = "UUID"
	// WorkStatusResponseVersionColumn is the version column name
	WorkStatusResponseVersionColumn = "Version"
	// WorkStatusResponseWorkConfigColumn is the work_config column name
	WorkStatusResponseWorkConfigColumn = "WorkConfig"
	// WorkStatusResponseWorkConfigColumnAllResolutionsColumn is the all_resolutions column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnAllResolutionsColumn = "WorkConfig.AllResolutions"
	// WorkStatusResponseWorkConfigColumnAllStatusesColumn is the all_statuses column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnAllStatusesColumn = "WorkConfig.AllStatuses"
	// WorkStatusResponseWorkConfigColumnCustomerIDColumn is the customer_id column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnCustomerIDColumn = "WorkConfig.CustomerID"
	// WorkStatusResponseWorkConfigColumnFieldColumn is the field column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnFieldColumn = "WorkConfig.Field"
	// WorkStatusResponseWorkConfigColumnIDColumn is the id column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnIDColumn = "WorkConfig.ID"
	// WorkStatusResponseWorkConfigColumnIntegrationIDColumn is the integration_id column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnIntegrationIDColumn = "WorkConfig.IntegrationID"
	// WorkStatusResponseWorkConfigColumnIssueTypeColumn is the issue_type column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnIssueTypeColumn = "WorkConfig.IssueType"
	// WorkStatusResponseWorkConfigColumnLabelsColumn is the labels column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnLabelsColumn = "WorkConfig.Labels"
	// WorkStatusResponseWorkConfigColumnOperatorColumn is the operator column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnOperatorColumn = "WorkConfig.Operator"
	// WorkStatusResponseWorkConfigColumnPrioritiesColumn is the priorities column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnPrioritiesColumn = "WorkConfig.Priorities"
	// WorkStatusResponseWorkConfigColumnRefIDColumn is the ref_id column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnRefIDColumn = "WorkConfig.RefID"
	// WorkStatusResponseWorkConfigColumnRefTypeColumn is the ref_type column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnRefTypeColumn = "WorkConfig.RefType"
	// WorkStatusResponseWorkConfigColumnResolutionsColumn is the resolutions column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnResolutionsColumn = "WorkConfig.Resolutions"
	// WorkStatusResponseWorkConfigColumnStatusesColumn is the statuses column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnStatusesColumn = "WorkConfig.Statuses"
	// WorkStatusResponseWorkConfigColumnTopLevelIssueColumn is the top_level_issue column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnTopLevelIssueColumn = "WorkConfig.TopLevelIssue"
	// WorkStatusResponseWorkConfigColumnTypeRulesColumn is the type_rules column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnTypeRulesColumn = "WorkConfig.TypeRules"
	// WorkStatusResponseWorkConfigColumnTypesColumn is the types column property of the WorkConfig name
	WorkStatusResponseWorkConfigColumnTypesColumn = "WorkConfig.Types"
)

// WorkStatusResponseEventDate represents the object structure for event_date
type WorkStatusResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusResponseEventDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseEventDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *WorkStatusResponseEventDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusResponseEventDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusResponseEventDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusResponseEventDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseEventDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusResponseWorkConfigField is the enumeration type for field
type WorkStatusResponseWorkConfigField int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigField) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIORITY":
		v = 0
	case "TYPE":
		v = 1
	case "LABEL":
		v = 2
	case "CUSTOM_FIELD_NAME":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigField) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIORITY")
	case 1:
		return json.Marshal("TYPE")
	case 2:
		return json.Marshal("LABEL")
	case 3:
		return json.Marshal("CUSTOM_FIELD_NAME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigField
func (v WorkStatusResponseWorkConfigField) String() string {
	switch int32(v) {
	case 0:
		return "PRIORITY"
	case 1:
		return "TYPE"
	case 2:
		return "LABEL"
	case 3:
		return "CUSTOM_FIELD_NAME"
	}
	return "unset"
}

const (
	// WorkConfigFieldPriority is the enumeration value for priority
	WorkStatusResponseWorkConfigFieldPriority WorkStatusResponseWorkConfigField = 0
	// WorkConfigFieldType is the enumeration value for type
	WorkStatusResponseWorkConfigFieldType WorkStatusResponseWorkConfigField = 1
	// WorkConfigFieldLabel is the enumeration value for label
	WorkStatusResponseWorkConfigFieldLabel WorkStatusResponseWorkConfigField = 2
	// WorkConfigFieldCustomFieldName is the enumeration value for custom_field_name
	WorkStatusResponseWorkConfigFieldCustomFieldName WorkStatusResponseWorkConfigField = 3
)

// WorkStatusResponseWorkConfigIssueType is the enumeration type for issue_type
type WorkStatusResponseWorkConfigIssueType int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigIssueType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENHANCEMENT":
		v = 0
	case "BUG":
		v = 1
	case "FEATURE":
		v = 2
	case "OTHER":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigIssueType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENHANCEMENT")
	case 1:
		return json.Marshal("BUG")
	case 2:
		return json.Marshal("FEATURE")
	case 3:
		return json.Marshal("OTHER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigIssueType
func (v WorkStatusResponseWorkConfigIssueType) String() string {
	switch int32(v) {
	case 0:
		return "ENHANCEMENT"
	case 1:
		return "BUG"
	case 2:
		return "FEATURE"
	case 3:
		return "OTHER"
	}
	return "unset"
}

const (
	// WorkConfigIssueTypeEnhancement is the enumeration value for enhancement
	WorkStatusResponseWorkConfigIssueTypeEnhancement WorkStatusResponseWorkConfigIssueType = 0
	// WorkConfigIssueTypeBug is the enumeration value for bug
	WorkStatusResponseWorkConfigIssueTypeBug WorkStatusResponseWorkConfigIssueType = 1
	// WorkConfigIssueTypeFeature is the enumeration value for feature
	WorkStatusResponseWorkConfigIssueTypeFeature WorkStatusResponseWorkConfigIssueType = 2
	// WorkConfigIssueTypeOther is the enumeration value for other
	WorkStatusResponseWorkConfigIssueTypeOther WorkStatusResponseWorkConfigIssueType = 3
)

// WorkStatusResponseLastExportDate represents the object structure for last_export_date
type WorkStatusResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toWorkStatusResponseLastExportDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseLastExportDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *WorkStatusResponseLastExportDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toWorkStatusResponseLastExportDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toWorkStatusResponseLastExportDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toWorkStatusResponseLastExportDateObject(o.Rfc3339, false),
	}
}

func (o *WorkStatusResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// WorkStatusResponseWorkConfigOperator is the enumeration type for operator
type WorkStatusResponseWorkConfigOperator int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigOperator) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "EQUALS":
		v = 0
	case "NOT_EQUALS":
		v = 1
	case "EXISTS":
		v = 2
	case "NOT_EXISTS":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigOperator) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("EQUALS")
	case 1:
		return json.Marshal("NOT_EQUALS")
	case 2:
		return json.Marshal("EXISTS")
	case 3:
		return json.Marshal("NOT_EXISTS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigOperator
func (v WorkStatusResponseWorkConfigOperator) String() string {
	switch int32(v) {
	case 0:
		return "EQUALS"
	case 1:
		return "NOT_EQUALS"
	case 2:
		return "EXISTS"
	case 3:
		return "NOT_EXISTS"
	}
	return "unset"
}

const (
	// WorkConfigOperatorEquals is the enumeration value for equals
	WorkStatusResponseWorkConfigOperatorEquals WorkStatusResponseWorkConfigOperator = 0
	// WorkConfigOperatorNotEquals is the enumeration value for not_equals
	WorkStatusResponseWorkConfigOperatorNotEquals WorkStatusResponseWorkConfigOperator = 1
	// WorkConfigOperatorExists is the enumeration value for exists
	WorkStatusResponseWorkConfigOperatorExists WorkStatusResponseWorkConfigOperator = 2
	// WorkConfigOperatorNotExists is the enumeration value for not_exists
	WorkStatusResponseWorkConfigOperatorNotExists WorkStatusResponseWorkConfigOperator = 3
)

// WorkStatusResponseType is the enumeration type for type
type WorkStatusResponseType int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseType) UnmarshalJSON(buf []byte) error {
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
func (v WorkStatusResponseType) MarshalJSON() ([]byte, error) {
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
func (v WorkStatusResponseType) String() string {
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
	WorkStatusResponseTypeEnroll WorkStatusResponseType = 0
	// TypePing is the enumeration value for ping
	WorkStatusResponseTypePing WorkStatusResponseType = 1
	// TypeCrash is the enumeration value for crash
	WorkStatusResponseTypeCrash WorkStatusResponseType = 2
	// TypeLog is the enumeration value for log
	WorkStatusResponseTypeLog WorkStatusResponseType = 3
	// TypeIntegration is the enumeration value for integration
	WorkStatusResponseTypeIntegration WorkStatusResponseType = 4
	// TypeExport is the enumeration value for export
	WorkStatusResponseTypeExport WorkStatusResponseType = 5
	// TypeProject is the enumeration value for project
	WorkStatusResponseTypeProject WorkStatusResponseType = 6
	// TypeRepo is the enumeration value for repo
	WorkStatusResponseTypeRepo WorkStatusResponseType = 7
	// TypeUser is the enumeration value for user
	WorkStatusResponseTypeUser WorkStatusResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	WorkStatusResponseTypeUninstall WorkStatusResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	WorkStatusResponseTypeUpgrade WorkStatusResponseType = 10
	// TypeStart is the enumeration value for start
	WorkStatusResponseTypeStart WorkStatusResponseType = 11
	// TypeStop is the enumeration value for stop
	WorkStatusResponseTypeStop WorkStatusResponseType = 12
)

// WorkStatusResponseWorkConfigResolutions represents the object structure for resolutions
type WorkStatusResponseWorkConfigResolutions struct {
	// WorkDone work completed resolution
	WorkDone []string `json:"work_done" yaml:"work_done" faker:"-"`
	// NoWorkDone resolutions that represent not completed issues
	NoWorkDone []string `json:"no_work_done" yaml:"no_work_done" faker:"-"`
}

func toWorkStatusResponseWorkConfigResolutionsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfigResolutions:
		return v.ToMap()

	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfigResolutions) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// WorkDone work completed resolution
		"work_done": toWorkStatusResponseWorkConfigResolutionsObject(o.WorkDone, false),
		// NoWorkDone resolutions that represent not completed issues
		"no_work_done": toWorkStatusResponseWorkConfigResolutionsObject(o.NoWorkDone, false),
	}
}

func (o *WorkStatusResponseWorkConfigResolutions) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfigResolutions) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["work_done"]; ok {
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
								panic("unsupported type for work_done field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for work_done field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for work_done field")
				}
			}
			o.WorkDone = na
		}
	}
	if o.WorkDone == nil {
		o.WorkDone = make([]string, 0)
	}

	if val, ok := kv["no_work_done"]; ok {
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
								panic("unsupported type for no_work_done field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for no_work_done field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for no_work_done field")
				}
			}
			o.NoWorkDone = na
		}
	}
	if o.NoWorkDone == nil {
		o.NoWorkDone = make([]string, 0)
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfigStatuses represents the object structure for statuses
type WorkStatusResponseWorkConfigStatuses struct {
	// OpenStatus The open status
	OpenStatus []string `json:"open_status" yaml:"open_status" faker:"-"`
	// InProgressStatus The in progress state names
	InProgressStatus []string `json:"in_progress_status" yaml:"in_progress_status" faker:"-"`
	// InReviewStatus The in review state names
	InReviewStatus []string `json:"in_review_status" yaml:"in_review_status" faker:"-"`
	// ReleasedStatus The released state names
	ReleasedStatus []string `json:"released_status" yaml:"released_status" faker:"-"`
	// ReopenedStatus The reopened state names
	ReopenedStatus []string `json:"reopened_status" yaml:"reopened_status" faker:"-"`
	// ClosedStatus The closed state names
	ClosedStatus []string `json:"closed_status" yaml:"closed_status" faker:"-"`
}

func toWorkStatusResponseWorkConfigStatusesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfigStatuses:
		return v.ToMap()

	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfigStatuses) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// OpenStatus The open status
		"open_status": toWorkStatusResponseWorkConfigStatusesObject(o.OpenStatus, false),
		// InProgressStatus The in progress state names
		"in_progress_status": toWorkStatusResponseWorkConfigStatusesObject(o.InProgressStatus, false),
		// InReviewStatus The in review state names
		"in_review_status": toWorkStatusResponseWorkConfigStatusesObject(o.InReviewStatus, false),
		// ReleasedStatus The released state names
		"released_status": toWorkStatusResponseWorkConfigStatusesObject(o.ReleasedStatus, false),
		// ReopenedStatus The reopened state names
		"reopened_status": toWorkStatusResponseWorkConfigStatusesObject(o.ReopenedStatus, false),
		// ClosedStatus The closed state names
		"closed_status": toWorkStatusResponseWorkConfigStatusesObject(o.ClosedStatus, false),
	}
}

func (o *WorkStatusResponseWorkConfigStatuses) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfigStatuses) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["open_status"]; ok {
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
								panic("unsupported type for open_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for open_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for open_status field")
				}
			}
			o.OpenStatus = na
		}
	}
	if o.OpenStatus == nil {
		o.OpenStatus = make([]string, 0)
	}

	if val, ok := kv["in_progress_status"]; ok {
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
								panic("unsupported type for in_progress_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for in_progress_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for in_progress_status field")
				}
			}
			o.InProgressStatus = na
		}
	}
	if o.InProgressStatus == nil {
		o.InProgressStatus = make([]string, 0)
	}

	if val, ok := kv["in_review_status"]; ok {
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
								panic("unsupported type for in_review_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for in_review_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for in_review_status field")
				}
			}
			o.InReviewStatus = na
		}
	}
	if o.InReviewStatus == nil {
		o.InReviewStatus = make([]string, 0)
	}

	if val, ok := kv["released_status"]; ok {
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
								panic("unsupported type for released_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for released_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for released_status field")
				}
			}
			o.ReleasedStatus = na
		}
	}
	if o.ReleasedStatus == nil {
		o.ReleasedStatus = make([]string, 0)
	}

	if val, ok := kv["reopened_status"]; ok {
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
								panic("unsupported type for reopened_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for reopened_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for reopened_status field")
				}
			}
			o.ReopenedStatus = na
		}
	}
	if o.ReopenedStatus == nil {
		o.ReopenedStatus = make([]string, 0)
	}

	if val, ok := kv["closed_status"]; ok {
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
								panic("unsupported type for closed_status field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for closed_status field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for closed_status field")
				}
			}
			o.ClosedStatus = na
		}
	}
	if o.ClosedStatus == nil {
		o.ClosedStatus = make([]string, 0)
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfigTopLevelIssue represents the object structure for top_level_issue
type WorkStatusResponseWorkConfigTopLevelIssue struct {
	// Name name of the top level issue
	Name string `json:"name" yaml:"name" faker:"-"`
	// Type type of the top level issue
	Type string `json:"type" yaml:"type" faker:"-"`
}

func toWorkStatusResponseWorkConfigTopLevelIssueObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfigTopLevelIssue:
		return v.ToMap()

	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfigTopLevelIssue) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Name name of the top level issue
		"name": toWorkStatusResponseWorkConfigTopLevelIssueObject(o.Name, false),
		// Type type of the top level issue
		"type": toWorkStatusResponseWorkConfigTopLevelIssueObject(o.Type, false),
	}
}

func (o *WorkStatusResponseWorkConfigTopLevelIssue) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfigTopLevelIssue) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["type"].(string); ok {
		o.Type = val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Type = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfigTypeRulesPredicatesField is the enumeration type for field
type WorkStatusResponseWorkConfigTypeRulesPredicatesField int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesField) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIORITY":
		v = 0
	case "TYPE":
		v = 1
	case "LABEL":
		v = 2
	case "CUSTOM_FIELD_NAME":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesField) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIORITY")
	case 1:
		return json.Marshal("TYPE")
	case 2:
		return json.Marshal("LABEL")
	case 3:
		return json.Marshal("CUSTOM_FIELD_NAME")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigTypeRulesPredicatesField
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesField) String() string {
	switch int32(v) {
	case 0:
		return "PRIORITY"
	case 1:
		return "TYPE"
	case 2:
		return "LABEL"
	case 3:
		return "CUSTOM_FIELD_NAME"
	}
	return "unset"
}

const (
	// WorkConfigTypeRulesPredicatesFieldPriority is the enumeration value for priority
	WorkStatusResponseWorkConfigTypeRulesPredicatesFieldPriority WorkStatusResponseWorkConfigTypeRulesPredicatesField = 0
	// WorkConfigTypeRulesPredicatesFieldType is the enumeration value for type
	WorkStatusResponseWorkConfigTypeRulesPredicatesFieldType WorkStatusResponseWorkConfigTypeRulesPredicatesField = 1
	// WorkConfigTypeRulesPredicatesFieldLabel is the enumeration value for label
	WorkStatusResponseWorkConfigTypeRulesPredicatesFieldLabel WorkStatusResponseWorkConfigTypeRulesPredicatesField = 2
	// WorkConfigTypeRulesPredicatesFieldCustomFieldName is the enumeration value for custom_field_name
	WorkStatusResponseWorkConfigTypeRulesPredicatesFieldCustomFieldName WorkStatusResponseWorkConfigTypeRulesPredicatesField = 3
)

// WorkStatusResponseWorkConfigTypeRulesPredicatesOperator is the enumeration type for operator
type WorkStatusResponseWorkConfigTypeRulesPredicatesOperator int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesOperator) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "EQUALS":
		v = 0
	case "NOT_EQUALS":
		v = 1
	case "EXISTS":
		v = 2
	case "NOT_EXISTS":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesOperator) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("EQUALS")
	case 1:
		return json.Marshal("NOT_EQUALS")
	case 2:
		return json.Marshal("EXISTS")
	case 3:
		return json.Marshal("NOT_EXISTS")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigTypeRulesPredicatesOperator
func (v WorkStatusResponseWorkConfigTypeRulesPredicatesOperator) String() string {
	switch int32(v) {
	case 0:
		return "EQUALS"
	case 1:
		return "NOT_EQUALS"
	case 2:
		return "EXISTS"
	case 3:
		return "NOT_EXISTS"
	}
	return "unset"
}

const (
	// WorkConfigTypeRulesPredicatesOperatorEquals is the enumeration value for equals
	WorkStatusResponseWorkConfigTypeRulesPredicatesOperatorEquals WorkStatusResponseWorkConfigTypeRulesPredicatesOperator = 0
	// WorkConfigTypeRulesPredicatesOperatorNotEquals is the enumeration value for not_equals
	WorkStatusResponseWorkConfigTypeRulesPredicatesOperatorNotEquals WorkStatusResponseWorkConfigTypeRulesPredicatesOperator = 1
	// WorkConfigTypeRulesPredicatesOperatorExists is the enumeration value for exists
	WorkStatusResponseWorkConfigTypeRulesPredicatesOperatorExists WorkStatusResponseWorkConfigTypeRulesPredicatesOperator = 2
	// WorkConfigTypeRulesPredicatesOperatorNotExists is the enumeration value for not_exists
	WorkStatusResponseWorkConfigTypeRulesPredicatesOperatorNotExists WorkStatusResponseWorkConfigTypeRulesPredicatesOperator = 3
)

// WorkStatusResponseWorkConfigTypeRulesPredicates represents the object structure for predicates
type WorkStatusResponseWorkConfigTypeRulesPredicates struct {
	// Field The field to operate on
	Field WorkStatusResponseWorkConfigTypeRulesPredicatesField `json:"field" yaml:"field" faker:"-"`
	// Operator The operation
	Operator WorkStatusResponseWorkConfigTypeRulesPredicatesOperator `json:"operator" yaml:"operator" faker:"-"`
	// Value The value of the field
	Value *string `json:"value,omitempty" yaml:"value,omitempty" faker:"-"`
}

func toWorkStatusResponseWorkConfigTypeRulesPredicatesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfigTypeRulesPredicates:
		return v.ToMap()

	case WorkStatusResponseWorkConfigTypeRulesPredicatesField:
		return v.String()

	case WorkStatusResponseWorkConfigTypeRulesPredicatesOperator:
		return v.String()

	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfigTypeRulesPredicates) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Field The field to operate on
		"field": toWorkStatusResponseWorkConfigTypeRulesPredicatesObject(o.Field, false),
		// Operator The operation
		"operator": toWorkStatusResponseWorkConfigTypeRulesPredicatesObject(o.Operator, false),
		// Value The value of the field
		"value": toWorkStatusResponseWorkConfigTypeRulesPredicatesObject(o.Value, true),
	}
}

func (o *WorkStatusResponseWorkConfigTypeRulesPredicates) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfigTypeRulesPredicates) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["field"].(WorkStatusResponseWorkConfigTypeRulesPredicatesField); ok {
		o.Field = val
	} else {
		if em, ok := kv["field"].(map[string]interface{}); ok {
			ev := em["agent.field"].(string)
			switch ev {
			case "priority", "PRIORITY":
				o.Field = 0
			case "type", "TYPE":
				o.Field = 1
			case "label", "LABEL":
				o.Field = 2
			case "custom_field_name", "CUSTOM_FIELD_NAME":
				o.Field = 3
			}
		}
		if em, ok := kv["field"].(string); ok {
			switch em {
			case "priority", "PRIORITY":
				o.Field = 0
			case "type", "TYPE":
				o.Field = 1
			case "label", "LABEL":
				o.Field = 2
			case "custom_field_name", "CUSTOM_FIELD_NAME":
				o.Field = 3
			}
		}
	}

	if val, ok := kv["operator"].(WorkStatusResponseWorkConfigTypeRulesPredicatesOperator); ok {
		o.Operator = val
	} else {
		if em, ok := kv["operator"].(map[string]interface{}); ok {
			ev := em["agent.operator"].(string)
			switch ev {
			case "equals", "EQUALS":
				o.Operator = 0
			case "not_equals", "NOT_EQUALS":
				o.Operator = 1
			case "exists", "EXISTS":
				o.Operator = 2
			case "not_exists", "NOT_EXISTS":
				o.Operator = 3
			}
		}
		if em, ok := kv["operator"].(string); ok {
			switch em {
			case "equals", "EQUALS":
				o.Operator = 0
			case "not_equals", "NOT_EQUALS":
				o.Operator = 1
			case "exists", "EXISTS":
				o.Operator = 2
			case "not_exists", "NOT_EXISTS":
				o.Operator = 3
			}
		}
	}

	if val, ok := kv["value"].(*string); ok {
		o.Value = val
	} else if val, ok := kv["value"].(string); ok {
		o.Value = &val
	} else {
		if val, ok := kv["value"]; ok {
			if val == nil {
				o.Value = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Value = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfigTypeRulesIssueType is the enumeration type for issue_type
type WorkStatusResponseWorkConfigTypeRulesIssueType int32

// UnmarshalJSON unmarshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesIssueType) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "ENHANCEMENT":
		v = 0
	case "BUG":
		v = 1
	case "FEATURE":
		v = 2
	case "OTHER":
		v = 3
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v WorkStatusResponseWorkConfigTypeRulesIssueType) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("ENHANCEMENT")
	case 1:
		return json.Marshal("BUG")
	case 2:
		return json.Marshal("FEATURE")
	case 3:
		return json.Marshal("OTHER")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for WorkConfigTypeRulesIssueType
func (v WorkStatusResponseWorkConfigTypeRulesIssueType) String() string {
	switch int32(v) {
	case 0:
		return "ENHANCEMENT"
	case 1:
		return "BUG"
	case 2:
		return "FEATURE"
	case 3:
		return "OTHER"
	}
	return "unset"
}

const (
	// WorkConfigTypeRulesIssueTypeEnhancement is the enumeration value for enhancement
	WorkStatusResponseWorkConfigTypeRulesIssueTypeEnhancement WorkStatusResponseWorkConfigTypeRulesIssueType = 0
	// WorkConfigTypeRulesIssueTypeBug is the enumeration value for bug
	WorkStatusResponseWorkConfigTypeRulesIssueTypeBug WorkStatusResponseWorkConfigTypeRulesIssueType = 1
	// WorkConfigTypeRulesIssueTypeFeature is the enumeration value for feature
	WorkStatusResponseWorkConfigTypeRulesIssueTypeFeature WorkStatusResponseWorkConfigTypeRulesIssueType = 2
	// WorkConfigTypeRulesIssueTypeOther is the enumeration value for other
	WorkStatusResponseWorkConfigTypeRulesIssueTypeOther WorkStatusResponseWorkConfigTypeRulesIssueType = 3
)

// WorkStatusResponseWorkConfigTypeRules represents the object structure for type_rules
type WorkStatusResponseWorkConfigTypeRules struct {
	// Predicates The predicate of the rule
	Predicates []WorkStatusResponseWorkConfigTypeRulesPredicates `json:"predicates" yaml:"predicates" faker:"-"`
	// IssueType The type to map the issue as if all predicates are true
	IssueType WorkStatusResponseWorkConfigTypeRulesIssueType `json:"issue_type" yaml:"issue_type" faker:"-"`
}

func toWorkStatusResponseWorkConfigTypeRulesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfigTypeRules:
		return v.ToMap()

	case []WorkStatusResponseWorkConfigTypeRulesPredicates:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case WorkStatusResponseWorkConfigTypeRulesIssueType:
		return v.String()
	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfigTypeRules) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Predicates The predicate of the rule
		"predicates": toWorkStatusResponseWorkConfigTypeRulesObject(o.Predicates, false),
		// IssueType The type to map the issue as if all predicates are true
		"issue_type": toWorkStatusResponseWorkConfigTypeRulesObject(o.IssueType, false),
	}
}

func (o *WorkStatusResponseWorkConfigTypeRules) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfigTypeRules) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if o == nil {

		o.Predicates = make([]WorkStatusResponseWorkConfigTypeRulesPredicates, 0)

	}
	if val, ok := kv["predicates"]; ok {
		if sv, ok := val.([]WorkStatusResponseWorkConfigTypeRulesPredicates); ok {
			o.Predicates = sv
		} else if sp, ok := val.([]*WorkStatusResponseWorkConfigTypeRulesPredicates); ok {
			o.Predicates = o.Predicates[:0]
			for _, e := range sp {
				o.Predicates = append(o.Predicates, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(WorkStatusResponseWorkConfigTypeRulesPredicates); ok {
					o.Predicates = append(o.Predicates, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm WorkStatusResponseWorkConfigTypeRulesPredicates
					fm.FromMap(av)
					o.Predicates = append(o.Predicates, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av WorkStatusResponseWorkConfigTypeRulesPredicates
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for predicates field entry: " + reflect.TypeOf(ae).String())
					}
					o.Predicates = append(o.Predicates, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(WorkStatusResponseWorkConfigTypeRulesPredicates); ok {
					o.Predicates = append(o.Predicates, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm WorkStatusResponseWorkConfigTypeRulesPredicates
					fm.FromMap(r)
					o.Predicates = append(o.Predicates, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := WorkStatusResponseWorkConfigTypeRulesPredicates{}
					fm.FromMap(r)
					o.Predicates = append(o.Predicates, fm)
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
							var fm WorkStatusResponseWorkConfigTypeRulesPredicates
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Predicates = append(o.Predicates, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["issue_type"].(WorkStatusResponseWorkConfigTypeRulesIssueType); ok {
		o.IssueType = val
	} else {
		if em, ok := kv["issue_type"].(map[string]interface{}); ok {
			ev := em["agent.issue_type"].(string)
			switch ev {
			case "enhancement", "ENHANCEMENT":
				o.IssueType = 0
			case "bug", "BUG":
				o.IssueType = 1
			case "feature", "FEATURE":
				o.IssueType = 2
			case "other", "OTHER":
				o.IssueType = 3
			}
		}
		if em, ok := kv["issue_type"].(string); ok {
			switch em {
			case "enhancement", "ENHANCEMENT":
				o.IssueType = 0
			case "bug", "BUG":
				o.IssueType = 1
			case "feature", "FEATURE":
				o.IssueType = 2
			case "other", "OTHER":
				o.IssueType = 3
			}
		}
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfig represents the object structure for work_config
type WorkStatusResponseWorkConfig struct {
	// AllResolutions all the values of issue resolutions on the integration regardless of mapping
	AllResolutions []string `json:"all_resolutions" yaml:"all_resolutions" faker:"-"`
	// AllStatuses all the values of issue statuses on the integration regardless of mapping
	AllStatuses []string `json:"all_statuses" yaml:"all_statuses" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// Field The field to operate on
	Field WorkStatusResponseWorkConfigField `json:"field" yaml:"field" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The ID reference to the integration instance
	IntegrationID string `json:"integration_id" yaml:"integration_id" faker:"-"`
	// IssueType The type to map the issue as if all predicates are true
	IssueType WorkStatusResponseWorkConfigIssueType `json:"issue_type" yaml:"issue_type" faker:"-"`
	// Labels all the values of labels on the integration
	Labels []string `json:"labels" yaml:"labels" faker:"-"`
	// Operator The operation
	Operator WorkStatusResponseWorkConfigOperator `json:"operator" yaml:"operator" faker:"-"`
	// Priorities all the values for priority on the integration
	Priorities []string `json:"priorities" yaml:"priorities" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// Resolutions The mapping of resolutions that represent completion for this integration
	Resolutions WorkStatusResponseWorkConfigResolutions `json:"resolutions" yaml:"resolutions" faker:"-"`
	// Statuses The mapping of statuses for this integration
	Statuses WorkStatusResponseWorkConfigStatuses `json:"statuses" yaml:"statuses" faker:"-"`
	// TopLevelIssue details about the top level issue
	TopLevelIssue WorkStatusResponseWorkConfigTopLevelIssue `json:"top_level_issue" yaml:"top_level_issue" faker:"-"`
	// TypeRules The rules for mapping issue types for this integration
	TypeRules []WorkStatusResponseWorkConfigTypeRules `json:"type_rules" yaml:"type_rules" faker:"-"`
	// Types all the values for type on the integration
	Types []string `json:"types" yaml:"types" faker:"-"`
}

func toWorkStatusResponseWorkConfigObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfig:
		return v.ToMap()

	case WorkStatusResponseWorkConfigField:
		return v.String()

	case WorkStatusResponseWorkConfigIssueType:
		return v.String()

	case WorkStatusResponseWorkConfigOperator:
		return v.String()

	case WorkStatusResponseWorkConfigResolutions:
		return v.ToMap()

	case WorkStatusResponseWorkConfigStatuses:
		return v.ToMap()

	case WorkStatusResponseWorkConfigTopLevelIssue:
		return v.ToMap()

	case []WorkStatusResponseWorkConfigTypeRules:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfig) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// AllResolutions all the values of issue resolutions on the integration regardless of mapping
		"all_resolutions": toWorkStatusResponseWorkConfigObject(o.AllResolutions, false),
		// AllStatuses all the values of issue statuses on the integration regardless of mapping
		"all_statuses": toWorkStatusResponseWorkConfigObject(o.AllStatuses, false),
		// CustomerID the customer id for the model instance
		"customer_id": toWorkStatusResponseWorkConfigObject(o.CustomerID, false),
		// Field The field to operate on
		"field": toWorkStatusResponseWorkConfigObject(o.Field, false),
		// ID the primary key for the model instance
		"id": toWorkStatusResponseWorkConfigObject(o.ID, false),
		// IntegrationID The ID reference to the integration instance
		"integration_id": toWorkStatusResponseWorkConfigObject(o.IntegrationID, false),
		// IssueType The type to map the issue as if all predicates are true
		"issue_type": toWorkStatusResponseWorkConfigObject(o.IssueType, false),
		// Labels all the values of labels on the integration
		"labels": toWorkStatusResponseWorkConfigObject(o.Labels, false),
		// Operator The operation
		"operator": toWorkStatusResponseWorkConfigObject(o.Operator, false),
		// Priorities all the values for priority on the integration
		"priorities": toWorkStatusResponseWorkConfigObject(o.Priorities, false),
		// RefID the source system id for the model instance
		"ref_id": toWorkStatusResponseWorkConfigObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toWorkStatusResponseWorkConfigObject(o.RefType, false),
		// Resolutions The mapping of resolutions that represent completion for this integration
		"resolutions": toWorkStatusResponseWorkConfigObject(o.Resolutions, false),
		// Statuses The mapping of statuses for this integration
		"statuses": toWorkStatusResponseWorkConfigObject(o.Statuses, false),
		// TopLevelIssue details about the top level issue
		"top_level_issue": toWorkStatusResponseWorkConfigObject(o.TopLevelIssue, false),
		// TypeRules The rules for mapping issue types for this integration
		"type_rules": toWorkStatusResponseWorkConfigObject(o.TypeRules, false),
		// Types all the values for type on the integration
		"types": toWorkStatusResponseWorkConfigObject(o.Types, false),
	}
}

func (o *WorkStatusResponseWorkConfig) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponseWorkConfig) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["all_resolutions"]; ok {
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
								panic("unsupported type for all_resolutions field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for all_resolutions field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for all_resolutions field")
				}
			}
			o.AllResolutions = na
		}
	}
	if o.AllResolutions == nil {
		o.AllResolutions = make([]string, 0)
	}

	if val, ok := kv["all_statuses"]; ok {
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
								panic("unsupported type for all_statuses field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for all_statuses field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for all_statuses field")
				}
			}
			o.AllStatuses = na
		}
	}
	if o.AllStatuses == nil {
		o.AllStatuses = make([]string, 0)
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

	if val, ok := kv["field"].(WorkStatusResponseWorkConfigField); ok {
		o.Field = val
	} else {
		if em, ok := kv["field"].(map[string]interface{}); ok {
			ev := em["agent.field"].(string)
			switch ev {
			case "priority", "PRIORITY":
				o.Field = 0
			case "type", "TYPE":
				o.Field = 1
			case "label", "LABEL":
				o.Field = 2
			case "custom_field_name", "CUSTOM_FIELD_NAME":
				o.Field = 3
			}
		}
		if em, ok := kv["field"].(string); ok {
			switch em {
			case "priority", "PRIORITY":
				o.Field = 0
			case "type", "TYPE":
				o.Field = 1
			case "label", "LABEL":
				o.Field = 2
			case "custom_field_name", "CUSTOM_FIELD_NAME":
				o.Field = 3
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

	if val, ok := kv["issue_type"].(WorkStatusResponseWorkConfigIssueType); ok {
		o.IssueType = val
	} else {
		if em, ok := kv["issue_type"].(map[string]interface{}); ok {
			ev := em["agent.issue_type"].(string)
			switch ev {
			case "enhancement", "ENHANCEMENT":
				o.IssueType = 0
			case "bug", "BUG":
				o.IssueType = 1
			case "feature", "FEATURE":
				o.IssueType = 2
			case "other", "OTHER":
				o.IssueType = 3
			}
		}
		if em, ok := kv["issue_type"].(string); ok {
			switch em {
			case "enhancement", "ENHANCEMENT":
				o.IssueType = 0
			case "bug", "BUG":
				o.IssueType = 1
			case "feature", "FEATURE":
				o.IssueType = 2
			case "other", "OTHER":
				o.IssueType = 3
			}
		}
	}

	if val, ok := kv["labels"]; ok {
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
								panic("unsupported type for labels field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for labels field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for labels field")
				}
			}
			o.Labels = na
		}
	}
	if o.Labels == nil {
		o.Labels = make([]string, 0)
	}

	if val, ok := kv["operator"].(WorkStatusResponseWorkConfigOperator); ok {
		o.Operator = val
	} else {
		if em, ok := kv["operator"].(map[string]interface{}); ok {
			ev := em["agent.operator"].(string)
			switch ev {
			case "equals", "EQUALS":
				o.Operator = 0
			case "not_equals", "NOT_EQUALS":
				o.Operator = 1
			case "exists", "EXISTS":
				o.Operator = 2
			case "not_exists", "NOT_EXISTS":
				o.Operator = 3
			}
		}
		if em, ok := kv["operator"].(string); ok {
			switch em {
			case "equals", "EQUALS":
				o.Operator = 0
			case "not_equals", "NOT_EQUALS":
				o.Operator = 1
			case "exists", "EXISTS":
				o.Operator = 2
			case "not_exists", "NOT_EXISTS":
				o.Operator = 3
			}
		}
	}

	if val, ok := kv["priorities"]; ok {
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
								panic("unsupported type for priorities field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for priorities field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for priorities field")
				}
			}
			o.Priorities = na
		}
	}
	if o.Priorities == nil {
		o.Priorities = make([]string, 0)
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

	if val, ok := kv["resolutions"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Resolutions.FromMap(kv)
		} else if sv, ok := val.(WorkStatusResponseWorkConfigResolutions); ok {
			// struct
			o.Resolutions = sv
		} else if sp, ok := val.(*WorkStatusResponseWorkConfigResolutions); ok {
			// struct pointer
			o.Resolutions = *sp
		}
	} else {
		o.Resolutions.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["statuses"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Statuses.FromMap(kv)
		} else if sv, ok := val.(WorkStatusResponseWorkConfigStatuses); ok {
			// struct
			o.Statuses = sv
		} else if sp, ok := val.(*WorkStatusResponseWorkConfigStatuses); ok {
			// struct pointer
			o.Statuses = *sp
		}
	} else {
		o.Statuses.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["top_level_issue"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.TopLevelIssue.FromMap(kv)
		} else if sv, ok := val.(WorkStatusResponseWorkConfigTopLevelIssue); ok {
			// struct
			o.TopLevelIssue = sv
		} else if sp, ok := val.(*WorkStatusResponseWorkConfigTopLevelIssue); ok {
			// struct pointer
			o.TopLevelIssue = *sp
		}
	} else {
		o.TopLevelIssue.FromMap(map[string]interface{}{})
	}

	if o == nil {

		o.TypeRules = make([]WorkStatusResponseWorkConfigTypeRules, 0)

	}
	if val, ok := kv["type_rules"]; ok {
		if sv, ok := val.([]WorkStatusResponseWorkConfigTypeRules); ok {
			o.TypeRules = sv
		} else if sp, ok := val.([]*WorkStatusResponseWorkConfigTypeRules); ok {
			o.TypeRules = o.TypeRules[:0]
			for _, e := range sp {
				o.TypeRules = append(o.TypeRules, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(WorkStatusResponseWorkConfigTypeRules); ok {
					o.TypeRules = append(o.TypeRules, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm WorkStatusResponseWorkConfigTypeRules
					fm.FromMap(av)
					o.TypeRules = append(o.TypeRules, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av WorkStatusResponseWorkConfigTypeRules
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for type_rules field entry: " + reflect.TypeOf(ae).String())
					}
					o.TypeRules = append(o.TypeRules, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(WorkStatusResponseWorkConfigTypeRules); ok {
					o.TypeRules = append(o.TypeRules, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm WorkStatusResponseWorkConfigTypeRules
					fm.FromMap(r)
					o.TypeRules = append(o.TypeRules, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := WorkStatusResponseWorkConfigTypeRules{}
					fm.FromMap(r)
					o.TypeRules = append(o.TypeRules, fm)
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
							var fm WorkStatusResponseWorkConfigTypeRules
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.TypeRules = append(o.TypeRules, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["types"]; ok {
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
								panic("unsupported type for types field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for types field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for types field")
				}
			}
			o.Types = na
		}
	}
	if o.Types == nil {
		o.Types = make([]string, 0)
	}
	o.setDefaults(false)
}

// WorkStatusResponse an agent response to an action request for fetching issue statuses
type WorkStatusResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data,omitempty" yaml:"data,omitempty" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error,omitempty" yaml:"error,omitempty" faker:"-"`
	// EventDate the date of the event
	EventDate WorkStatusResponseEventDate `json:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate WorkStatusResponseLastExportDate `json:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type WorkStatusResponseType `json:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" yaml:"version" faker:"-"`
	// WorkConfig The work config structure to use
	WorkConfig WorkStatusResponseWorkConfig `json:"work_config" yaml:"work_config" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*WorkStatusResponse)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*WorkStatusResponse)(nil)

func toWorkStatusResponseObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponse:
		return v.ToMap()

	case WorkStatusResponseEventDate:
		return v.ToMap()

	case WorkStatusResponseLastExportDate:
		return v.ToMap()

	case WorkStatusResponseType:
		return v.String()

	case WorkStatusResponseWorkConfig:
		return v.ToMap()
	default:
		return o
	}
}

// String returns a string representation of WorkStatusResponse
func (o *WorkStatusResponse) String() string {
	return fmt.Sprintf("agent.WorkStatusResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *WorkStatusResponse) GetTopicName() datamodel.TopicNameType {
	return WorkStatusResponseTopic
}

// GetStreamName returns the name of the stream
func (o *WorkStatusResponse) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *WorkStatusResponse) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *WorkStatusResponse) GetModelName() datamodel.ModelNameType {
	return WorkStatusResponseModelName
}

// NewWorkStatusResponseID provides a template for generating an ID field for WorkStatusResponse
func NewWorkStatusResponseID(customerID string, refType string, refID string) string {
	return hash.Values("WorkStatusResponse", customerID, refType, refID)
}

func (o *WorkStatusResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.WorkConfig.TypeRules == nil {
		o.WorkConfig.TypeRules = make([]WorkStatusResponseWorkConfigTypeRules, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("WorkStatusResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *WorkStatusResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *WorkStatusResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WorkStatusResponse) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for WorkStatusResponse")
}

// GetRefID returns the RefID for the object
func (o *WorkStatusResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *WorkStatusResponse) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *WorkStatusResponse) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *WorkStatusResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *WorkStatusResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *WorkStatusResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = WorkStatusResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *WorkStatusResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *WorkStatusResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of WorkStatusResponse
func (o *WorkStatusResponse) Clone() datamodel.Model {
	c := new(WorkStatusResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *WorkStatusResponse) Anon() datamodel.Model {
	c := new(WorkStatusResponse)
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
func (o *WorkStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *WorkStatusResponse) UnmarshalJSON(data []byte) error {
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
func (o *WorkStatusResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two WorkStatusResponse objects are equal
func (o *WorkStatusResponse) IsEqual(other *WorkStatusResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *WorkStatusResponse) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toWorkStatusResponseObject(o.Architecture, false),
		"customer_id":      toWorkStatusResponseObject(o.CustomerID, false),
		"data":             toWorkStatusResponseObject(o.Data, true),
		"distro":           toWorkStatusResponseObject(o.Distro, false),
		"error":            toWorkStatusResponseObject(o.Error, true),
		"event_date":       toWorkStatusResponseObject(o.EventDate, false),
		"free_space":       toWorkStatusResponseObject(o.FreeSpace, false),
		"go_version":       toWorkStatusResponseObject(o.GoVersion, false),
		"hostname":         toWorkStatusResponseObject(o.Hostname, false),
		"id":               toWorkStatusResponseObject(o.ID, false),
		"integration_id":   toWorkStatusResponseObject(o.IntegrationID, false),
		"last_export_date": toWorkStatusResponseObject(o.LastExportDate, false),
		"memory":           toWorkStatusResponseObject(o.Memory, false),
		"message":          toWorkStatusResponseObject(o.Message, false),
		"num_cpu":          toWorkStatusResponseObject(o.NumCPU, false),
		"os":               toWorkStatusResponseObject(o.OS, false),
		"ref_id":           toWorkStatusResponseObject(o.RefID, false),
		"ref_type":         toWorkStatusResponseObject(o.RefType, false),
		"request_id":       toWorkStatusResponseObject(o.RequestID, false),
		"success":          toWorkStatusResponseObject(o.Success, false),
		"system_id":        toWorkStatusResponseObject(o.SystemID, false),

		"type":        o.Type.String(),
		"updated_ts":  toWorkStatusResponseObject(o.UpdatedAt, false),
		"uptime":      toWorkStatusResponseObject(o.Uptime, false),
		"uuid":        toWorkStatusResponseObject(o.UUID, false),
		"version":     toWorkStatusResponseObject(o.Version, false),
		"work_config": toWorkStatusResponseObject(o.WorkConfig, false),
		"hashcode":    toWorkStatusResponseObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *WorkStatusResponse) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(WorkStatusResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*WorkStatusResponseEventDate); ok {
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
		} else if sv, ok := val.(WorkStatusResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*WorkStatusResponseLastExportDate); ok {
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

	if val, ok := kv["type"].(WorkStatusResponseType); ok {
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

	if val, ok := kv["work_config"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.WorkConfig.FromMap(kv)
		} else if sv, ok := val.(WorkStatusResponseWorkConfig); ok {
			// struct
			o.WorkConfig = sv
		} else if sp, ok := val.(*WorkStatusResponseWorkConfig); ok {
			// struct pointer
			o.WorkConfig = *sp
		}
	} else {
		o.WorkConfig.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *WorkStatusResponse) Hash() string {
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
	args = append(args, o.WorkConfig)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *WorkStatusResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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
