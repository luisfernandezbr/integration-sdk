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

	// WorkStatusResponseTable is the default table name
	WorkStatusResponseTable datamodel.ModelNameType = "agent_workstatusresponse"

	// WorkStatusResponseModelName is the model name
	WorkStatusResponseModelName datamodel.ModelNameType = "agent.WorkStatusResponse"
)

const (
	// WorkStatusResponseModelArchitectureColumn is the column json value architecture
	WorkStatusResponseModelArchitectureColumn = "architecture"
	// WorkStatusResponseModelCustomerIDColumn is the column json value customer_id
	WorkStatusResponseModelCustomerIDColumn = "customer_id"
	// WorkStatusResponseModelDataColumn is the column json value data
	WorkStatusResponseModelDataColumn = "data"
	// WorkStatusResponseModelDistroColumn is the column json value distro
	WorkStatusResponseModelDistroColumn = "distro"
	// WorkStatusResponseModelErrorColumn is the column json value error
	WorkStatusResponseModelErrorColumn = "error"
	// WorkStatusResponseModelEventDateColumn is the column json value event_date
	WorkStatusResponseModelEventDateColumn = "event_date"
	// WorkStatusResponseModelEventDateEpochColumn is the column json value epoch
	WorkStatusResponseModelEventDateEpochColumn = "epoch"
	// WorkStatusResponseModelEventDateOffsetColumn is the column json value offset
	WorkStatusResponseModelEventDateOffsetColumn = "offset"
	// WorkStatusResponseModelEventDateRfc3339Column is the column json value rfc3339
	WorkStatusResponseModelEventDateRfc3339Column = "rfc3339"
	// WorkStatusResponseModelFreeSpaceColumn is the column json value free_space
	WorkStatusResponseModelFreeSpaceColumn = "free_space"
	// WorkStatusResponseModelGoVersionColumn is the column json value go_version
	WorkStatusResponseModelGoVersionColumn = "go_version"
	// WorkStatusResponseModelHostnameColumn is the column json value hostname
	WorkStatusResponseModelHostnameColumn = "hostname"
	// WorkStatusResponseModelIDColumn is the column json value id
	WorkStatusResponseModelIDColumn = "id"
	// WorkStatusResponseModelIntegrationIDColumn is the column json value integration_id
	WorkStatusResponseModelIntegrationIDColumn = "integration_id"
	// WorkStatusResponseModelLastExportDateColumn is the column json value last_export_date
	WorkStatusResponseModelLastExportDateColumn = "last_export_date"
	// WorkStatusResponseModelLastExportDateEpochColumn is the column json value epoch
	WorkStatusResponseModelLastExportDateEpochColumn = "epoch"
	// WorkStatusResponseModelLastExportDateOffsetColumn is the column json value offset
	WorkStatusResponseModelLastExportDateOffsetColumn = "offset"
	// WorkStatusResponseModelLastExportDateRfc3339Column is the column json value rfc3339
	WorkStatusResponseModelLastExportDateRfc3339Column = "rfc3339"
	// WorkStatusResponseModelMemoryColumn is the column json value memory
	WorkStatusResponseModelMemoryColumn = "memory"
	// WorkStatusResponseModelMessageColumn is the column json value message
	WorkStatusResponseModelMessageColumn = "message"
	// WorkStatusResponseModelNumCPUColumn is the column json value num_cpu
	WorkStatusResponseModelNumCPUColumn = "num_cpu"
	// WorkStatusResponseModelOSColumn is the column json value os
	WorkStatusResponseModelOSColumn = "os"
	// WorkStatusResponseModelRefIDColumn is the column json value ref_id
	WorkStatusResponseModelRefIDColumn = "ref_id"
	// WorkStatusResponseModelRefTypeColumn is the column json value ref_type
	WorkStatusResponseModelRefTypeColumn = "ref_type"
	// WorkStatusResponseModelRequestIDColumn is the column json value request_id
	WorkStatusResponseModelRequestIDColumn = "request_id"
	// WorkStatusResponseModelSuccessColumn is the column json value success
	WorkStatusResponseModelSuccessColumn = "success"
	// WorkStatusResponseModelSystemIDColumn is the column json value system_id
	WorkStatusResponseModelSystemIDColumn = "system_id"
	// WorkStatusResponseModelTypeColumn is the column json value type
	WorkStatusResponseModelTypeColumn = "type"
	// WorkStatusResponseModelUptimeColumn is the column json value uptime
	WorkStatusResponseModelUptimeColumn = "uptime"
	// WorkStatusResponseModelUUIDColumn is the column json value uuid
	WorkStatusResponseModelUUIDColumn = "uuid"
	// WorkStatusResponseModelVersionColumn is the column json value version
	WorkStatusResponseModelVersionColumn = "version"
	// WorkStatusResponseModelWorkConfigColumn is the column json value work_config
	WorkStatusResponseModelWorkConfigColumn = "work_config"
	// WorkStatusResponseModelWorkConfigCustomerIDColumn is the column json value customer_id
	WorkStatusResponseModelWorkConfigCustomerIDColumn = "customer_id"
	// WorkStatusResponseModelWorkConfigIDColumn is the column json value id
	WorkStatusResponseModelWorkConfigIDColumn = "id"
	// WorkStatusResponseModelWorkConfigIntegrationIDColumn is the column json value integration_id
	WorkStatusResponseModelWorkConfigIntegrationIDColumn = "integration_id"
	// WorkStatusResponseModelWorkConfigRefIDColumn is the column json value ref_id
	WorkStatusResponseModelWorkConfigRefIDColumn = "ref_id"
	// WorkStatusResponseModelWorkConfigRefTypeColumn is the column json value ref_type
	WorkStatusResponseModelWorkConfigRefTypeColumn = "ref_type"
	// WorkStatusResponseModelWorkConfigStatusesColumn is the column json value statuses
	WorkStatusResponseModelWorkConfigStatusesColumn = "statuses"
	// WorkStatusResponseModelWorkConfigStatusesOpenStatusColumn is the column json value open_status
	WorkStatusResponseModelWorkConfigStatusesOpenStatusColumn = "open_status"
	// WorkStatusResponseModelWorkConfigStatusesInProgressStatusColumn is the column json value in_progress_status
	WorkStatusResponseModelWorkConfigStatusesInProgressStatusColumn = "in_progress_status"
	// WorkStatusResponseModelWorkConfigStatusesInReviewStatusColumn is the column json value in_review_status
	WorkStatusResponseModelWorkConfigStatusesInReviewStatusColumn = "in_review_status"
	// WorkStatusResponseModelWorkConfigStatusesReleasedStatusColumn is the column json value released_status
	WorkStatusResponseModelWorkConfigStatusesReleasedStatusColumn = "released_status"
	// WorkStatusResponseModelWorkConfigStatusesReopenedStatusColumn is the column json value reopened_status
	WorkStatusResponseModelWorkConfigStatusesReopenedStatusColumn = "reopened_status"
	// WorkStatusResponseModelWorkConfigStatusesClosedStatusColumn is the column json value closed_status
	WorkStatusResponseModelWorkConfigStatusesClosedStatusColumn = "closed_status"
	// WorkStatusResponseModelWorkConfigTopLevelIssueColumn is the column json value top_level_issue
	WorkStatusResponseModelWorkConfigTopLevelIssueColumn = "top_level_issue"
	// WorkStatusResponseModelWorkConfigTopLevelIssueNameColumn is the column json value name
	WorkStatusResponseModelWorkConfigTopLevelIssueNameColumn = "name"
	// WorkStatusResponseModelWorkConfigTopLevelIssueTypeColumn is the column json value type
	WorkStatusResponseModelWorkConfigTopLevelIssueTypeColumn = "type"
)

// WorkStatusResponseEventDate represents the object structure for event_date
type WorkStatusResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
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

// WorkStatusResponseLastExportDate represents the object structure for last_export_date
type WorkStatusResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
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

// WorkStatusResponseType is the enumeration type for type
type WorkStatusResponseType int32

// UnmarshalBSONValue for unmarshaling value
func (v *WorkStatusResponseType) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = WorkStatusResponseType(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "ENROLL":
			*v = WorkStatusResponseType(0)
		case "PING":
			*v = WorkStatusResponseType(1)
		case "CRASH":
			*v = WorkStatusResponseType(2)
		case "LOG":
			*v = WorkStatusResponseType(3)
		case "INTEGRATION":
			*v = WorkStatusResponseType(4)
		case "EXPORT":
			*v = WorkStatusResponseType(5)
		case "PROJECT":
			*v = WorkStatusResponseType(6)
		case "REPO":
			*v = WorkStatusResponseType(7)
		case "USER":
			*v = WorkStatusResponseType(8)
		case "CALENDAR":
			*v = WorkStatusResponseType(9)
		case "UNINSTALL":
			*v = WorkStatusResponseType(10)
		case "UPGRADE":
			*v = WorkStatusResponseType(11)
		case "START":
			*v = WorkStatusResponseType(12)
		case "STOP":
			*v = WorkStatusResponseType(13)
		case "PAUSE":
			*v = WorkStatusResponseType(14)
		case "RESUME":
			*v = WorkStatusResponseType(15)
		}
	}
	return nil
}

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

const (
	// WorkStatusResponseTypeEnroll is the enumeration value for enroll
	WorkStatusResponseTypeEnroll WorkStatusResponseType = 0
	// WorkStatusResponseTypePing is the enumeration value for ping
	WorkStatusResponseTypePing WorkStatusResponseType = 1
	// WorkStatusResponseTypeCrash is the enumeration value for crash
	WorkStatusResponseTypeCrash WorkStatusResponseType = 2
	// WorkStatusResponseTypeLog is the enumeration value for log
	WorkStatusResponseTypeLog WorkStatusResponseType = 3
	// WorkStatusResponseTypeIntegration is the enumeration value for integration
	WorkStatusResponseTypeIntegration WorkStatusResponseType = 4
	// WorkStatusResponseTypeExport is the enumeration value for export
	WorkStatusResponseTypeExport WorkStatusResponseType = 5
	// WorkStatusResponseTypeProject is the enumeration value for project
	WorkStatusResponseTypeProject WorkStatusResponseType = 6
	// WorkStatusResponseTypeRepo is the enumeration value for repo
	WorkStatusResponseTypeRepo WorkStatusResponseType = 7
	// WorkStatusResponseTypeUser is the enumeration value for user
	WorkStatusResponseTypeUser WorkStatusResponseType = 8
	// WorkStatusResponseTypeCalendar is the enumeration value for calendar
	WorkStatusResponseTypeCalendar WorkStatusResponseType = 9
	// WorkStatusResponseTypeUninstall is the enumeration value for uninstall
	WorkStatusResponseTypeUninstall WorkStatusResponseType = 10
	// WorkStatusResponseTypeUpgrade is the enumeration value for upgrade
	WorkStatusResponseTypeUpgrade WorkStatusResponseType = 11
	// WorkStatusResponseTypeStart is the enumeration value for start
	WorkStatusResponseTypeStart WorkStatusResponseType = 12
	// WorkStatusResponseTypeStop is the enumeration value for stop
	WorkStatusResponseTypeStop WorkStatusResponseType = 13
	// WorkStatusResponseTypePause is the enumeration value for pause
	WorkStatusResponseTypePause WorkStatusResponseType = 14
	// WorkStatusResponseTypeResume is the enumeration value for resume
	WorkStatusResponseTypeResume WorkStatusResponseType = 15
)

// WorkStatusResponseWorkConfigStatuses represents the object structure for statuses
type WorkStatusResponseWorkConfigStatuses struct {
	// OpenStatus The open status
	OpenStatus []string `json:"open_status" codec:"open_status" bson:"open_status" yaml:"open_status" faker:"-"`
	// InProgressStatus The in progress state names
	InProgressStatus []string `json:"in_progress_status" codec:"in_progress_status" bson:"in_progress_status" yaml:"in_progress_status" faker:"-"`
	// InReviewStatus The in review state names
	InReviewStatus []string `json:"in_review_status" codec:"in_review_status" bson:"in_review_status" yaml:"in_review_status" faker:"-"`
	// ReleasedStatus The released state names
	ReleasedStatus []string `json:"released_status" codec:"released_status" bson:"released_status" yaml:"released_status" faker:"-"`
	// ReopenedStatus The reopened state names
	ReopenedStatus []string `json:"reopened_status" codec:"reopened_status" bson:"reopened_status" yaml:"reopened_status" faker:"-"`
	// ClosedStatus The closed state names
	ClosedStatus []string `json:"closed_status" codec:"closed_status" bson:"closed_status" yaml:"closed_status" faker:"-"`
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
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Type type of the top level issue
	Type string `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
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

	if val, ok := kv["type"].(string); ok {
		o.Type = val
	} else {
		if val, ok := kv["type"]; ok {
			if val == nil {
				o.Type = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Type = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// WorkStatusResponseWorkConfig represents the object structure for work_config
type WorkStatusResponseWorkConfig struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"id" yaml:"id" faker:"-"`
	// IntegrationID The ID reference to the integration instance
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Statuses The mapping of statuses for this integration
	Statuses WorkStatusResponseWorkConfigStatuses `json:"statuses" codec:"statuses" bson:"statuses" yaml:"statuses" faker:"-"`
	// TopLevelIssue details about the top level issue
	TopLevelIssue WorkStatusResponseWorkConfigTopLevelIssue `json:"top_level_issue" codec:"top_level_issue" bson:"top_level_issue" yaml:"top_level_issue" faker:"-"`
}

func toWorkStatusResponseWorkConfigObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *WorkStatusResponseWorkConfig:
		return v.ToMap()

	case WorkStatusResponseWorkConfigStatuses:
		return v.ToMap()

	case WorkStatusResponseWorkConfigTopLevelIssue:
		return v.ToMap()
	default:
		return o
	}
}

func (o *WorkStatusResponseWorkConfig) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// CustomerID the customer id for the model instance
		"customer_id": toWorkStatusResponseWorkConfigObject(o.CustomerID, false),
		// ID the primary key for the model instance
		"id": toWorkStatusResponseWorkConfigObject(o.ID, false),
		// IntegrationID The ID reference to the integration instance
		"integration_id": toWorkStatusResponseWorkConfigObject(o.IntegrationID, false),
		// RefID the source system id for the model instance
		"ref_id": toWorkStatusResponseWorkConfigObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toWorkStatusResponseWorkConfigObject(o.RefType, false),
		// Statuses The mapping of statuses for this integration
		"statuses": toWorkStatusResponseWorkConfigObject(o.Statuses, false),
		// TopLevelIssue details about the top level issue
		"top_level_issue": toWorkStatusResponseWorkConfigObject(o.TopLevelIssue, false),
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

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
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

	o.setDefaults(false)
}

// WorkStatusResponse an agent response to an action request for fetching issue statuses
type WorkStatusResponse struct {
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
	EventDate WorkStatusResponseEventDate `json:"event_date" codec:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	LastExportDate WorkStatusResponseLastExportDate `json:"last_export_date" codec:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
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
	// Success if the response was successful
	Success bool `json:"success" codec:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type WorkStatusResponseType `json:"type" codec:"type" bson:"type" yaml:"type" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" codec:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" codec:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" codec:"version" bson:"version" yaml:"version" faker:"-"`
	// WorkConfig The work config structure to use
	WorkConfig WorkStatusResponseWorkConfig `json:"work_config" codec:"work_config" bson:"work_config" yaml:"work_config" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
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
	return ""
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
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *WorkStatusResponse) GetTimestamp() time.Time {
	return time.Now().UTC()
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
	return false
}

// GetTopicConfig returns the topic config object
func (o *WorkStatusResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
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

// StringifyPretty returns the object in JSON format as a string prettified
func (o *WorkStatusResponse) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
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

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
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
