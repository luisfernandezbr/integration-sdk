// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// ExportTable is the default table name
	ExportTable datamodel.ModelNameType = "agent_export"

	// ExportModelName is the model name
	ExportModelName datamodel.ModelNameType = "agent.Export"
)

const (
	// ExportModelCustomerIDColumn is the column json value customer_id
	ExportModelCustomerIDColumn = "customer_id"
	// ExportModelIDColumn is the column json value id
	ExportModelIDColumn = "id"
	// ExportModelIntegrationColumn is the column json value integration
	ExportModelIntegrationColumn = "integration"
	// ExportModelIntegrationActiveColumn is the column json value active
	ExportModelIntegrationActiveColumn = "active"
	// ExportModelIntegrationAutoConfigureColumn is the column json value auto_configure
	ExportModelIntegrationAutoConfigureColumn = "auto_configure"
	// ExportModelIntegrationConfigColumn is the column json value config
	ExportModelIntegrationConfigColumn = "config"
	// ExportModelIntegrationCreatedByProfileIDColumn is the column json value created_by_profile_id
	ExportModelIntegrationCreatedByProfileIDColumn = "created_by_profile_id"
	// ExportModelIntegrationCreatedByUserIDColumn is the column json value created_by_user_id
	ExportModelIntegrationCreatedByUserIDColumn = "created_by_user_id"
	// ExportModelIntegrationCreatedDateColumn is the column json value created_date
	ExportModelIntegrationCreatedDateColumn = "created_date"
	// ExportModelIntegrationCreatedDateEpochColumn is the column json value epoch
	ExportModelIntegrationCreatedDateEpochColumn = "epoch"
	// ExportModelIntegrationCreatedDateOffsetColumn is the column json value offset
	ExportModelIntegrationCreatedDateOffsetColumn = "offset"
	// ExportModelIntegrationCreatedDateRfc3339Column is the column json value rfc3339
	ExportModelIntegrationCreatedDateRfc3339Column = "rfc3339"
	// ExportModelIntegrationCustomerIDColumn is the column json value customer_id
	ExportModelIntegrationCustomerIDColumn = "customer_id"
	// ExportModelIntegrationDeletedColumn is the column json value deleted
	ExportModelIntegrationDeletedColumn = "deleted"
	// ExportModelIntegrationDeletedByProfileIDColumn is the column json value deleted_by_profile_id
	ExportModelIntegrationDeletedByProfileIDColumn = "deleted_by_profile_id"
	// ExportModelIntegrationDeletedByUserIDColumn is the column json value deleted_by_user_id
	ExportModelIntegrationDeletedByUserIDColumn = "deleted_by_user_id"
	// ExportModelIntegrationDeletedDateColumn is the column json value deleted_date
	ExportModelIntegrationDeletedDateColumn = "deleted_date"
	// ExportModelIntegrationDeletedDateEpochColumn is the column json value epoch
	ExportModelIntegrationDeletedDateEpochColumn = "epoch"
	// ExportModelIntegrationDeletedDateOffsetColumn is the column json value offset
	ExportModelIntegrationDeletedDateOffsetColumn = "offset"
	// ExportModelIntegrationDeletedDateRfc3339Column is the column json value rfc3339
	ExportModelIntegrationDeletedDateRfc3339Column = "rfc3339"
	// ExportModelIntegrationEnrollmentIDColumn is the column json value enrollment_id
	ExportModelIntegrationEnrollmentIDColumn = "enrollment_id"
	// ExportModelIntegrationErrorMessageColumn is the column json value error_message
	ExportModelIntegrationErrorMessageColumn = "error_message"
	// ExportModelIntegrationErroredColumn is the column json value errored
	ExportModelIntegrationErroredColumn = "errored"
	// ExportModelIntegrationExportAcknowledgedColumn is the column json value export_acknowledged
	ExportModelIntegrationExportAcknowledgedColumn = "export_acknowledged"
	// ExportModelIntegrationIDColumn is the column json value id
	ExportModelIntegrationIDColumn = "id"
	// ExportModelIntegrationIntegrationIDColumn is the column json value integration_id
	ExportModelIntegrationIntegrationIDColumn = "integration_id"
	// ExportModelIntegrationIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportModelIntegrationIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportModelIntegrationIntervalColumn is the column json value interval
	ExportModelIntegrationIntervalColumn = "interval"
	// ExportModelIntegrationJobIDColumn is the column json value job_id
	ExportModelIntegrationJobIDColumn = "job_id"
	// ExportModelIntegrationLocationColumn is the column json value location
	ExportModelIntegrationLocationColumn = "location"
	// ExportModelIntegrationNameColumn is the column json value name
	ExportModelIntegrationNameColumn = "name"
	// ExportModelIntegrationPausedColumn is the column json value paused
	ExportModelIntegrationPausedColumn = "paused"
	// ExportModelIntegrationPrivateKeyColumn is the column json value private_key
	ExportModelIntegrationPrivateKeyColumn = "private_key"
	// ExportModelIntegrationProcessedColumn is the column json value processed
	ExportModelIntegrationProcessedColumn = "processed"
	// ExportModelIntegrationRefIDColumn is the column json value ref_id
	ExportModelIntegrationRefIDColumn = "ref_id"
	// ExportModelIntegrationRefTypeColumn is the column json value ref_type
	ExportModelIntegrationRefTypeColumn = "ref_type"
	// ExportModelIntegrationRequiresHistoricalColumn is the column json value requires_historical
	ExportModelIntegrationRequiresHistoricalColumn = "requires_historical"
	// ExportModelIntegrationSetupColumn is the column json value setup
	ExportModelIntegrationSetupColumn = "setup"
	// ExportModelIntegrationStateColumn is the column json value state
	ExportModelIntegrationStateColumn = "state"
	// ExportModelIntegrationThrottledColumn is the column json value throttled
	ExportModelIntegrationThrottledColumn = "throttled"
	// ExportModelIntegrationThrottledUntilColumn is the column json value throttled_until
	ExportModelIntegrationThrottledUntilColumn = "throttled_until"
	// ExportModelIntegrationThrottledUntilEpochColumn is the column json value epoch
	ExportModelIntegrationThrottledUntilEpochColumn = "epoch"
	// ExportModelIntegrationThrottledUntilOffsetColumn is the column json value offset
	ExportModelIntegrationThrottledUntilOffsetColumn = "offset"
	// ExportModelIntegrationThrottledUntilRfc3339Column is the column json value rfc3339
	ExportModelIntegrationThrottledUntilRfc3339Column = "rfc3339"
	// ExportModelIntegrationUpgradeDateAtColumn is the column json value upgrade_date_ts
	ExportModelIntegrationUpgradeDateAtColumn = "upgrade_date_ts"
	// ExportModelIntegrationUpgradeExpiresDateAtColumn is the column json value upgrade_expires_date_ts
	ExportModelIntegrationUpgradeExpiresDateAtColumn = "upgrade_expires_date_ts"
	// ExportModelIntegrationUpgradeMessageColumn is the column json value upgrade_message
	ExportModelIntegrationUpgradeMessageColumn = "upgrade_message"
	// ExportModelIntegrationUpgradeRequiredColumn is the column json value upgrade_required
	ExportModelIntegrationUpgradeRequiredColumn = "upgrade_required"
	// ExportModelIntegrationWebhooksColumn is the column json value webhooks
	ExportModelIntegrationWebhooksColumn = "webhooks"
	// ExportModelIntegrationWebhooksEnabledColumn is the column json value enabled
	ExportModelIntegrationWebhooksEnabledColumn = "enabled"
	// ExportModelIntegrationWebhooksURLColumn is the column json value url
	ExportModelIntegrationWebhooksURLColumn = "url"
	// ExportModelIntegrationWebhooksErroredColumn is the column json value errored
	ExportModelIntegrationWebhooksErroredColumn = "errored"
	// ExportModelIntegrationWebhooksErrorMessageColumn is the column json value error_message
	ExportModelIntegrationWebhooksErrorMessageColumn = "error_message"
	// ExportModelIntegrationWebhooksRefIDColumn is the column json value ref_id
	ExportModelIntegrationWebhooksRefIDColumn = "ref_id"
	// ExportModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ExportModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ExportModelJobIDColumn is the column json value job_id
	ExportModelJobIDColumn = "job_id"
	// ExportModelRefIDColumn is the column json value ref_id
	ExportModelRefIDColumn = "ref_id"
	// ExportModelRefTypeColumn is the column json value ref_type
	ExportModelRefTypeColumn = "ref_type"
	// ExportModelReprocessHistoricalColumn is the column json value reprocess_historical
	ExportModelReprocessHistoricalColumn = "reprocess_historical"
)

// ExportIntegrationCreatedDate represents the object structure for created_date
type ExportIntegrationCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportIntegrationCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportIntegrationCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportIntegrationCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportIntegrationCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportIntegrationCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportIntegrationCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportIntegrationCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportIntegrationCreatedDate) FromMap(kv map[string]interface{}) {

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

// ExportIntegrationDeletedDate represents the object structure for deleted_date
type ExportIntegrationDeletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportIntegrationDeletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportIntegrationDeletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportIntegrationDeletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportIntegrationDeletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportIntegrationDeletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportIntegrationDeletedDateObject(o.Rfc3339, false),
	}
}

func (o *ExportIntegrationDeletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportIntegrationDeletedDate) FromMap(kv map[string]interface{}) {

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

// ExportIntegrationLocation is the enumeration type for location
type ExportIntegrationLocation int32

// toExportIntegrationLocationPointer is the enumeration pointer type for location
func toExportIntegrationLocationPointer(v int32) *ExportIntegrationLocation {
	nv := ExportIntegrationLocation(v)
	return &nv
}

// toExportIntegrationLocationEnum is the enumeration pointer wrapper for location
func toExportIntegrationLocationEnum(v *ExportIntegrationLocation) string {
	if v == nil {
		return toExportIntegrationLocationPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *ExportIntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportIntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = ExportIntegrationLocation(0)
		case "CLOUD":
			*v = ExportIntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *ExportIntegrationLocation) UnmarshalJSON(buf []byte) error {
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
func (v ExportIntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationLocation
func (v ExportIntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportIntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportIntegrationLocation:
		*v = val
	case int32:
		*v = ExportIntegrationLocation(int32(val))
	case int:
		*v = ExportIntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = ExportIntegrationLocation(0)
		case "CLOUD":
			*v = ExportIntegrationLocation(1)
		}
	}
	return nil
}

const (
	// ExportIntegrationLocationPrivate is the enumeration value for private
	ExportIntegrationLocationPrivate ExportIntegrationLocation = 0
	// ExportIntegrationLocationCloud is the enumeration value for cloud
	ExportIntegrationLocationCloud ExportIntegrationLocation = 1
)

// ExportIntegrationSetup is the enumeration type for setup
type ExportIntegrationSetup int32

// toExportIntegrationSetupPointer is the enumeration pointer type for setup
func toExportIntegrationSetupPointer(v int32) *ExportIntegrationSetup {
	nv := ExportIntegrationSetup(v)
	return &nv
}

// toExportIntegrationSetupEnum is the enumeration pointer wrapper for setup
func toExportIntegrationSetupEnum(v *ExportIntegrationSetup) string {
	if v == nil {
		return toExportIntegrationSetupPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *ExportIntegrationSetup) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportIntegrationSetup(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "CONFIG":
			*v = ExportIntegrationSetup(0)
		case "READY":
			*v = ExportIntegrationSetup(1)
		case "RUNNING":
			*v = ExportIntegrationSetup(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *ExportIntegrationSetup) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "CONFIG":
		*v = 0
	case "READY":
		*v = 1
	case "RUNNING":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportIntegrationSetup) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("CONFIG")
	case 1:
		return json.Marshal("READY")
	case 2:
		return json.Marshal("RUNNING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationSetup
func (v ExportIntegrationSetup) String() string {
	switch int32(v) {
	case 0:
		return "CONFIG"
	case 1:
		return "READY"
	case 2:
		return "RUNNING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportIntegrationSetup) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportIntegrationSetup:
		*v = val
	case int32:
		*v = ExportIntegrationSetup(int32(val))
	case int:
		*v = ExportIntegrationSetup(int32(val))
	case string:
		switch val {
		case "CONFIG":
			*v = ExportIntegrationSetup(0)
		case "READY":
			*v = ExportIntegrationSetup(1)
		case "RUNNING":
			*v = ExportIntegrationSetup(2)
		}
	}
	return nil
}

const (
	// ExportIntegrationSetupConfig is the enumeration value for config
	ExportIntegrationSetupConfig ExportIntegrationSetup = 0
	// ExportIntegrationSetupReady is the enumeration value for ready
	ExportIntegrationSetupReady ExportIntegrationSetup = 1
	// ExportIntegrationSetupRunning is the enumeration value for running
	ExportIntegrationSetupRunning ExportIntegrationSetup = 2
)

// ExportIntegrationState is the enumeration type for state
type ExportIntegrationState int32

// toExportIntegrationStatePointer is the enumeration pointer type for state
func toExportIntegrationStatePointer(v int32) *ExportIntegrationState {
	nv := ExportIntegrationState(v)
	return &nv
}

// toExportIntegrationStateEnum is the enumeration pointer wrapper for state
func toExportIntegrationStateEnum(v *ExportIntegrationState) string {
	if v == nil {
		return toExportIntegrationStatePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *ExportIntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = ExportIntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = ExportIntegrationState(0)
		case "EXPORTING":
			*v = ExportIntegrationState(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *ExportIntegrationState) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "IDLE":
		*v = 0
	case "EXPORTING":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v ExportIntegrationState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("EXPORTING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for IntegrationState
func (v ExportIntegrationState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "EXPORTING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *ExportIntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case ExportIntegrationState:
		*v = val
	case int32:
		*v = ExportIntegrationState(int32(val))
	case int:
		*v = ExportIntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = ExportIntegrationState(0)
		case "EXPORTING":
			*v = ExportIntegrationState(1)
		}
	}
	return nil
}

const (
	// ExportIntegrationStateIdle is the enumeration value for idle
	ExportIntegrationStateIdle ExportIntegrationState = 0
	// ExportIntegrationStateExporting is the enumeration value for exporting
	ExportIntegrationStateExporting ExportIntegrationState = 1
)

// ExportIntegrationThrottledUntil represents the object structure for throttled_until
type ExportIntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toExportIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportIntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportIntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toExportIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toExportIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toExportIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *ExportIntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportIntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

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

// ExportIntegrationWebhooks represents the object structure for webhooks
type ExportIntegrationWebhooks struct {
	// Enabled if webhooks are enabled for this instance
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Errored if the webhook has an error
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// RefID an optional reference id related to the webhook and what its related to
	RefID *string `json:"ref_id,omitempty" codec:"ref_id,omitempty" bson:"ref_id" yaml:"ref_id,omitempty" faker:"-"`
}

func toExportIntegrationWebhooksObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportIntegrationWebhooks:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportIntegrationWebhooks) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Enabled if webhooks are enabled for this instance
		"enabled": toExportIntegrationWebhooksObject(o.Enabled, false),
		// URL the url the webhook for the webhook
		"url": toExportIntegrationWebhooksObject(o.URL, true),
		// Errored if the webhook has an error
		"errored": toExportIntegrationWebhooksObject(o.Errored, false),
		// ErrorMessage the error message
		"error_message": toExportIntegrationWebhooksObject(o.ErrorMessage, true),
		// RefID an optional reference id related to the webhook and what its related to
		"ref_id": toExportIntegrationWebhooksObject(o.RefID, true),
	}
}

func (o *ExportIntegrationWebhooks) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportIntegrationWebhooks) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = false
			} else {
				o.Enabled = number.ToBoolAny(val)
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
	if val, ok := kv["errored"].(bool); ok {
		o.Errored = val
	} else {
		if val, ok := kv["errored"]; ok {
			if val == nil {
				o.Errored = false
			} else {
				o.Errored = number.ToBoolAny(val)
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
	if val, ok := kv["ref_id"].(*string); ok {
		o.RefID = val
	} else if val, ok := kv["ref_id"].(string); ok {
		o.RefID = &val
	} else {
		if val, ok := kv["ref_id"]; ok {
			if val == nil {
				o.RefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	o.setDefaults(false)
}

// ExportIntegration represents the object structure for integration
type ExportIntegration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// AutoConfigure if this integration should be auto configured by the integration before activating
	AutoConfigure bool `json:"auto_configure" codec:"auto_configure" bson:"auto_configure" yaml:"auto_configure" faker:"-"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CreatedDate when the integration was created
	CreatedDate ExportIntegrationCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deleted If true, the integration has been deleted
	Deleted bool `json:"deleted" codec:"deleted" bson:"deleted" yaml:"deleted" faker:"-"`
	// DeletedByProfileID The id of the profile for the user that deleted the integration
	DeletedByProfileID *string `json:"deleted_by_profile_id,omitempty" codec:"deleted_by_profile_id,omitempty" bson:"deleted_by_profile_id" yaml:"deleted_by_profile_id,omitempty" faker:"-"`
	// DeletedByUserID The id of the user that deleted the integration
	DeletedByUserID *string `json:"deleted_by_user_id,omitempty" codec:"deleted_by_user_id,omitempty" bson:"deleted_by_user_id" yaml:"deleted_by_user_id,omitempty" faker:"-"`
	// DeletedDate when the integration was deleted
	DeletedDate ExportIntegrationDeletedDate `json:"deleted_date" codec:"deleted_date" bson:"deleted_date" yaml:"deleted_date" faker:"-"`
	// EnrollmentID if the integration is linked to a self-managed agent, it will have the enrollment_id set otherwise will be null
	EnrollmentID *string `json:"enrollment_id,omitempty" codec:"enrollment_id,omitempty" bson:"enrollment_id" yaml:"enrollment_id,omitempty" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// ExportAcknowledged Set to true an export has been received by the agent.
	ExportAcknowledged *bool `json:"export_acknowledged,omitempty" codec:"export_acknowledged,omitempty" bson:"export_acknowledged" yaml:"export_acknowledged,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The unique id for the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval int64 `json:"interval" codec:"interval" bson:"interval" yaml:"interval" faker:"-"`
	// JobID the current (if EXPORTING) or previous (if IDLE) job id which is useful for debugging
	JobID *string `json:"job_id,omitempty" codec:"job_id,omitempty" bson:"job_id" yaml:"job_id,omitempty" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location ExportIntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused bool `json:"paused" codec:"paused" bson:"paused" yaml:"paused" faker:"-"`
	// PrivateKey the private key for the instance if needed by an integration
	PrivateKey *string `json:"private_key,omitempty" codec:"private_key,omitempty" bson:"private_key" yaml:"private_key,omitempty" faker:"-"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequiresHistorical flag which can be set to trigger a historical on the next scheduler visit
	RequiresHistorical bool `json:"requires_historical" codec:"requires_historical" bson:"requires_historical" yaml:"requires_historical" faker:"-"`
	// Setup the setup state of the integration
	Setup ExportIntegrationSetup `json:"setup" codec:"setup" bson:"setup" yaml:"setup" faker:"-"`
	// State the current state of the integration
	State ExportIntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *ExportIntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// UpgradeDateAt If upgrade is required, the date when the upgrade was set
	UpgradeDateAt *int64 `json:"upgrade_date_ts,omitempty" codec:"upgrade_date_ts,omitempty" bson:"upgrade_date_ts" yaml:"upgrade_date_ts,omitempty" faker:"-"`
	// UpgradeExpiresDateAt If upgrade is required and there's a due date this will be set
	UpgradeExpiresDateAt *int64 `json:"upgrade_expires_date_ts,omitempty" codec:"upgrade_expires_date_ts,omitempty" bson:"upgrade_expires_date_ts" yaml:"upgrade_expires_date_ts,omitempty" faker:"-"`
	// UpgradeMessage If upgrade is required, the message to display to the user
	UpgradeMessage *string `json:"upgrade_message,omitempty" codec:"upgrade_message,omitempty" bson:"upgrade_message" yaml:"upgrade_message,omitempty" faker:"-"`
	// UpgradeRequired If true, the integration requires a manual upgrade
	UpgradeRequired bool `json:"upgrade_required" codec:"upgrade_required" bson:"upgrade_required" yaml:"upgrade_required" faker:"-"`
	// Webhooks for any webhooks installed at the integration instance
	Webhooks []ExportIntegrationWebhooks `json:"webhooks" codec:"webhooks" bson:"webhooks" yaml:"webhooks" faker:"-"`
}

func toExportIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ExportIntegration:
		return v.ToMap()

	case ExportIntegrationCreatedDate:
		return v.ToMap()

	case ExportIntegrationDeletedDate:
		return v.ToMap()

	case ExportIntegrationLocation:
		return v.String()

	case ExportIntegrationSetup:
		return v.String()

	case ExportIntegrationState:
		return v.String()

	case *ExportIntegrationThrottledUntil:
		return v.ToMap()

	case []ExportIntegrationWebhooks:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ExportIntegration) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Active If true, the integration is still active
		"active": toExportIntegrationObject(o.Active, false),
		// AutoConfigure if this integration should be auto configured by the integration before activating
		"auto_configure": toExportIntegrationObject(o.AutoConfigure, false),
		// Config the integration configuration controlled by the integration itself
		"config": toExportIntegrationObject(o.Config, true),
		// CreatedByProfileID The id of the profile for the user that created the integration
		"created_by_profile_id": toExportIntegrationObject(o.CreatedByProfileID, true),
		// CreatedByUserID The id of the user that created the integration
		"created_by_user_id": toExportIntegrationObject(o.CreatedByUserID, true),
		// CreatedDate when the integration was created
		"created_date": toExportIntegrationObject(o.CreatedDate, false),
		// CustomerID the customer id for the model instance
		"customer_id": toExportIntegrationObject(o.CustomerID, false),
		// Deleted If true, the integration has been deleted
		"deleted": toExportIntegrationObject(o.Deleted, false),
		// DeletedByProfileID The id of the profile for the user that deleted the integration
		"deleted_by_profile_id": toExportIntegrationObject(o.DeletedByProfileID, true),
		// DeletedByUserID The id of the user that deleted the integration
		"deleted_by_user_id": toExportIntegrationObject(o.DeletedByUserID, true),
		// DeletedDate when the integration was deleted
		"deleted_date": toExportIntegrationObject(o.DeletedDate, false),
		// EnrollmentID if the integration is linked to a self-managed agent, it will have the enrollment_id set otherwise will be null
		"enrollment_id": toExportIntegrationObject(o.EnrollmentID, true),
		// ErrorMessage The error message from an export run
		"error_message": toExportIntegrationObject(o.ErrorMessage, true),
		// Errored If authorization failed by the agent or any other error
		"errored": toExportIntegrationObject(o.Errored, true),
		// ExportAcknowledged Set to true an export has been received by the agent.
		"export_acknowledged": toExportIntegrationObject(o.ExportAcknowledged, true),
		// ID the primary key for the model instance
		"id": toExportIntegrationObject(o.ID, false),
		// IntegrationID The unique id for the integration
		"integration_id": toExportIntegrationObject(o.IntegrationID, false),
		// IntegrationInstanceID the integration instance id
		"integration_instance_id": toExportIntegrationObject(o.IntegrationInstanceID, true),
		// Interval the interval in milliseconds for how often an export job is scheduled
		"interval": toExportIntegrationObject(o.Interval, false),
		// JobID the current (if EXPORTING) or previous (if IDLE) job id which is useful for debugging
		"job_id": toExportIntegrationObject(o.JobID, true),
		// Location The location of this integration (on-premise / private or cloud)
		"location": toExportIntegrationObject(o.Location, false),
		// Name The user friendly name of the integration
		"name": toExportIntegrationObject(o.Name, false),
		// Paused true if the agent is paused and should not start new scheduled jobs
		"paused": toExportIntegrationObject(o.Paused, false),
		// PrivateKey the private key for the instance if needed by an integration
		"private_key": toExportIntegrationObject(o.PrivateKey, true),
		// Processed If the integration has been processed at least once
		"processed": toExportIntegrationObject(o.Processed, true),
		// RefID the source system id for the model instance
		"ref_id": toExportIntegrationObject(o.RefID, false),
		// RefType the source system identifier for the model instance
		"ref_type": toExportIntegrationObject(o.RefType, false),
		// RequiresHistorical flag which can be set to trigger a historical on the next scheduler visit
		"requires_historical": toExportIntegrationObject(o.RequiresHistorical, false),
		// Setup the setup state of the integration
		"setup": toExportIntegrationObject(o.Setup, false),
		// State the current state of the integration
		"state": toExportIntegrationObject(o.State, false),
		// Throttled Set to true when integration is throttled.
		"throttled": toExportIntegrationObject(o.Throttled, true),
		// ThrottledUntil After throttling integration, we set this field for estimated resume date.
		"throttled_until": toExportIntegrationObject(o.ThrottledUntil, true),
		// UpgradeDateAt If upgrade is required, the date when the upgrade was set
		"upgrade_date_ts": toExportIntegrationObject(o.UpgradeDateAt, true),
		// UpgradeExpiresDateAt If upgrade is required and there's a due date this will be set
		"upgrade_expires_date_ts": toExportIntegrationObject(o.UpgradeExpiresDateAt, true),
		// UpgradeMessage If upgrade is required, the message to display to the user
		"upgrade_message": toExportIntegrationObject(o.UpgradeMessage, true),
		// UpgradeRequired If true, the integration requires a manual upgrade
		"upgrade_required": toExportIntegrationObject(o.UpgradeRequired, false),
		// Webhooks for any webhooks installed at the integration instance
		"webhooks": toExportIntegrationObject(o.Webhooks, false),
	}
}

func (o *ExportIntegration) setDefaults(frommap bool) {

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	if o.ExportAcknowledged == nil {
		var v bool
		o.ExportAcknowledged = &v
	}

	if o.Processed == nil {
		var v bool
		o.Processed = &v
	}

	if o.Throttled == nil {
		var v bool
		o.Throttled = &v
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ExportIntegration) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["auto_configure"].(bool); ok {
		o.AutoConfigure = val
	} else {
		if val, ok := kv["auto_configure"]; ok {
			if val == nil {
				o.AutoConfigure = false
			} else {
				o.AutoConfigure = number.ToBoolAny(val)
			}
		}
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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(ExportIntegrationCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*ExportIntegrationCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.CreatedDate.Epoch = dt.Epoch
				o.CreatedDate.Rfc3339 = dt.Rfc3339
				o.CreatedDate.Offset = dt.Offset
			}
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
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
	if val, ok := kv["deleted"].(bool); ok {
		o.Deleted = val
	} else {
		if val, ok := kv["deleted"]; ok {
			if val == nil {
				o.Deleted = false
			} else {
				o.Deleted = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["deleted_by_profile_id"].(*string); ok {
		o.DeletedByProfileID = val
	} else if val, ok := kv["deleted_by_profile_id"].(string); ok {
		o.DeletedByProfileID = &val
	} else {
		if val, ok := kv["deleted_by_profile_id"]; ok {
			if val == nil {
				o.DeletedByProfileID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByProfileID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["deleted_by_user_id"].(*string); ok {
		o.DeletedByUserID = val
	} else if val, ok := kv["deleted_by_user_id"].(string); ok {
		o.DeletedByUserID = &val
	} else {
		if val, ok := kv["deleted_by_user_id"]; ok {
			if val == nil {
				o.DeletedByUserID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.DeletedByUserID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["deleted_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.DeletedDate.FromMap(kv)
		} else if sv, ok := val.(ExportIntegrationDeletedDate); ok {
			// struct
			o.DeletedDate = sv
		} else if sp, ok := val.(*ExportIntegrationDeletedDate); ok {
			// struct pointer
			o.DeletedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
			o.DeletedDate.Epoch = dt.Epoch
			o.DeletedDate.Rfc3339 = dt.Rfc3339
			o.DeletedDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.DeletedDate.Epoch = dt.Epoch
				o.DeletedDate.Rfc3339 = dt.Rfc3339
				o.DeletedDate.Offset = dt.Offset
			}
		}
	} else {
		o.DeletedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["enrollment_id"].(*string); ok {
		o.EnrollmentID = val
	} else if val, ok := kv["enrollment_id"].(string); ok {
		o.EnrollmentID = &val
	} else {
		if val, ok := kv["enrollment_id"]; ok {
			if val == nil {
				o.EnrollmentID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.EnrollmentID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["export_acknowledged"].(*bool); ok {
		o.ExportAcknowledged = val
	} else if val, ok := kv["export_acknowledged"].(bool); ok {
		o.ExportAcknowledged = &val
	} else {
		if val, ok := kv["export_acknowledged"]; ok {
			if val == nil {
				o.ExportAcknowledged = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ExportAcknowledged = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["location"].(ExportIntegrationLocation); ok {
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
	if val, ok := kv["private_key"].(*string); ok {
		o.PrivateKey = val
	} else if val, ok := kv["private_key"].(string); ok {
		o.PrivateKey = &val
	} else {
		if val, ok := kv["private_key"]; ok {
			if val == nil {
				o.PrivateKey = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.PrivateKey = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["requires_historical"].(bool); ok {
		o.RequiresHistorical = val
	} else {
		if val, ok := kv["requires_historical"]; ok {
			if val == nil {
				o.RequiresHistorical = false
			} else {
				o.RequiresHistorical = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["setup"].(ExportIntegrationSetup); ok {
		o.Setup = val
	} else {
		if em, ok := kv["setup"].(map[string]interface{}); ok {

			ev := em["agent.setup"].(string)
			switch ev {
			case "config", "CONFIG":
				o.Setup = 0
			case "ready", "READY":
				o.Setup = 1
			case "running", "RUNNING":
				o.Setup = 2
			}
		}
		if em, ok := kv["setup"].(string); ok {
			switch em {
			case "config", "CONFIG":
				o.Setup = 0
			case "ready", "READY":
				o.Setup = 1
			case "running", "RUNNING":
				o.Setup = 2
			}
		}
	}
	if val, ok := kv["state"].(ExportIntegrationState); ok {
		o.State = val
	} else {
		if em, ok := kv["state"].(map[string]interface{}); ok {

			ev := em["agent.state"].(string)
			switch ev {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
			}
		}
		if em, ok := kv["state"].(string); ok {
			switch em {
			case "idle", "IDLE":
				o.State = 0
			case "exporting", "EXPORTING":
				o.State = 1
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
		o.ThrottledUntil = &ExportIntegrationThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(ExportIntegrationThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*ExportIntegrationThrottledUntil); ok {
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

	if val, ok := kv["upgrade_date_ts"].(*int64); ok {
		o.UpgradeDateAt = val
	} else if val, ok := kv["upgrade_date_ts"].(int64); ok {
		o.UpgradeDateAt = &val
	} else {
		if val, ok := kv["upgrade_date_ts"]; ok {
			if val == nil {
				o.UpgradeDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_expires_date_ts"].(*int64); ok {
		o.UpgradeExpiresDateAt = val
	} else if val, ok := kv["upgrade_expires_date_ts"].(int64); ok {
		o.UpgradeExpiresDateAt = &val
	} else {
		if val, ok := kv["upgrade_expires_date_ts"]; ok {
			if val == nil {
				o.UpgradeExpiresDateAt = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["long"]
				}
				o.UpgradeExpiresDateAt = number.Int64Pointer(number.ToInt64Any(val))
			}
		}
	}
	if val, ok := kv["upgrade_message"].(*string); ok {
		o.UpgradeMessage = val
	} else if val, ok := kv["upgrade_message"].(string); ok {
		o.UpgradeMessage = &val
	} else {
		if val, ok := kv["upgrade_message"]; ok {
			if val == nil {
				o.UpgradeMessage = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.UpgradeMessage = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["upgrade_required"].(bool); ok {
		o.UpgradeRequired = val
	} else {
		if val, ok := kv["upgrade_required"]; ok {
			if val == nil {
				o.UpgradeRequired = false
			} else {
				o.UpgradeRequired = number.ToBoolAny(val)
			}
		}
	}

	if o == nil {

		o.Webhooks = make([]ExportIntegrationWebhooks, 0)

	}
	if val, ok := kv["webhooks"]; ok {
		if sv, ok := val.([]ExportIntegrationWebhooks); ok {
			o.Webhooks = sv
		} else if sp, ok := val.([]*ExportIntegrationWebhooks); ok {
			o.Webhooks = o.Webhooks[:0]
			for _, e := range sp {
				o.Webhooks = append(o.Webhooks, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(ExportIntegrationWebhooks); ok {
					o.Webhooks = append(o.Webhooks, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm ExportIntegrationWebhooks
					fm.FromMap(av)
					o.Webhooks = append(o.Webhooks, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av ExportIntegrationWebhooks
					av.FromMap(bkv)
					o.Webhooks = append(o.Webhooks, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(ExportIntegrationWebhooks); ok {
					o.Webhooks = append(o.Webhooks, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm ExportIntegrationWebhooks
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := ExportIntegrationWebhooks{}
					fm.FromMap(r)
					o.Webhooks = append(o.Webhooks, fm)
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
							var fm ExportIntegrationWebhooks
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Webhooks = append(o.Webhooks, fm)
						}
					}
				}
			}
		}
	}

	o.setDefaults(false)
}

// Export an agent action to request an export
type Export struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Integration The integrations that should be exported and their current configuration
	Integration ExportIntegration `json:"integration" codec:"integration" bson:"integration" yaml:"integration" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// JobID The job ID
	JobID string `json:"job_id" codec:"job_id" bson:"job_id" yaml:"job_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical bool `json:"reprocess_historical" codec:"reprocess_historical" bson:"reprocess_historical" yaml:"reprocess_historical" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Export)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Export)(nil)

func toExportObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Export:
		return v.ToMap()

	case ExportIntegration:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Export
func (o *Export) String() string {
	return fmt.Sprintf("agent.Export<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Export) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Export) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Export) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Export) GetModelName() datamodel.ModelNameType {
	return ExportModelName
}

// NewExportID provides a template for generating an ID field for Export
func NewExportID(customerID string, refType string, refID string) string {
	return hash.Values("Export", customerID, refType, refID)
}

func (o *Export) setDefaults(frommap bool) {
	if o.Integration.ThrottledUntil == nil {
		o.Integration.ThrottledUntil = &ExportIntegrationThrottledUntil{}
	}
	if o.Integration.Webhooks == nil {
		o.Integration.Webhooks = make([]ExportIntegrationWebhooks, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Export", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Export) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Export) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Export) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Export) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Export) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Export) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Export) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Export) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Export) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ExportModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Export) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Export) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Export
func (o *Export) Clone() datamodel.Model {
	c := new(Export)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Export) Anon() datamodel.Model {
	c := new(Export)
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
func (o *Export) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Export) UnmarshalJSON(data []byte) error {
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
func (o *Export) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Export) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Export objects are equal
func (o *Export) IsEqual(other *Export) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Export) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"customer_id":             toExportObject(o.CustomerID, false),
		"id":                      toExportObject(o.ID, false),
		"integration":             toExportObject(o.Integration, false),
		"integration_instance_id": toExportObject(o.IntegrationInstanceID, true),
		"job_id":                  toExportObject(o.JobID, false),
		"ref_id":                  toExportObject(o.RefID, false),
		"ref_type":                toExportObject(o.RefType, false),
		"reprocess_historical":    toExportObject(o.ReprocessHistorical, false),
		"hashcode":                toExportObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Export) FromMap(kv map[string]interface{}) {

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
		} else if sv, ok := val.(ExportIntegration); ok {
			// struct
			o.Integration = sv
		} else if sp, ok := val.(*ExportIntegration); ok {
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
	if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = false
			} else {
				o.ReprocessHistorical = number.ToBoolAny(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Export) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.Integration)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.JobID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.ReprocessHistorical)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// ExportPartial is a partial struct for upsert mutations for Export
type ExportPartial struct {
	// Integration The integrations that should be exported and their current configuration
	Integration *ExportIntegration `json:"integration,omitempty"`
	// JobID The job ID
	JobID *string `json:"job_id,omitempty"`
	// ReprocessHistorical This field is to differentiate between historical and incrementals
	ReprocessHistorical *bool `json:"reprocess_historical,omitempty"`
}

var _ datamodel.PartialModel = (*ExportPartial)(nil)

// GetModelName returns the name of the model
func (o *ExportPartial) GetModelName() datamodel.ModelNameType {
	return ExportModelName
}

// ToMap returns the object as a map
func (o *ExportPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"integration":          toExportObject(o.Integration, true),
		"job_id":               toExportObject(o.JobID, true),
		"reprocess_historical": toExportObject(o.ReprocessHistorical, true),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *ExportPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ExportPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ExportPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ExportPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ExportPartial) FromMap(kv map[string]interface{}) {

	if o.Integration == nil {
		o.Integration = &ExportIntegration{}
	}

	if val, ok := kv["integration"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Integration.FromMap(kv)
		} else if sv, ok := val.(ExportIntegration); ok {
			// struct
			o.Integration = &sv
		} else if sp, ok := val.(*ExportIntegration); ok {
			// struct pointer
			o.Integration = sp
		}
	} else {
		o.Integration.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["job_id"].(*string); ok {
		o.JobID = val
	} else if val, ok := kv["job_id"].(string); ok {
		o.JobID = &val
	} else {
		if val, ok := kv["job_id"]; ok {
			if val == nil {
				o.JobID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.JobID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["reprocess_historical"].(*bool); ok {
		o.ReprocessHistorical = val
	} else if val, ok := kv["reprocess_historical"].(bool); ok {
		o.ReprocessHistorical = &val
	} else {
		if val, ok := kv["reprocess_historical"]; ok {
			if val == nil {
				o.ReprocessHistorical = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.ReprocessHistorical = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	o.setDefaults(false)
}
