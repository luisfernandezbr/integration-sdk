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
	"github.com/pinpt/go-common/v10/graphql"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (

	// IntegrationTable is the default table name
	IntegrationTable datamodel.ModelNameType = "agent_integration"

	// IntegrationModelName is the model name
	IntegrationModelName datamodel.ModelNameType = "agent.Integration"
)

const (
	// IntegrationModelActiveColumn is the column json value active
	IntegrationModelActiveColumn = "active"
	// IntegrationModelConfigColumn is the column json value config
	IntegrationModelConfigColumn = "config"
	// IntegrationModelCreatedByProfileIDColumn is the column json value created_by_profile_id
	IntegrationModelCreatedByProfileIDColumn = "created_by_profile_id"
	// IntegrationModelCreatedByUserIDColumn is the column json value created_by_user_id
	IntegrationModelCreatedByUserIDColumn = "created_by_user_id"
	// IntegrationModelCreatedDateColumn is the column json value created_date
	IntegrationModelCreatedDateColumn = "created_date"
	// IntegrationModelCreatedDateEpochColumn is the column json value epoch
	IntegrationModelCreatedDateEpochColumn = "epoch"
	// IntegrationModelCreatedDateOffsetColumn is the column json value offset
	IntegrationModelCreatedDateOffsetColumn = "offset"
	// IntegrationModelCreatedDateRfc3339Column is the column json value rfc3339
	IntegrationModelCreatedDateRfc3339Column = "rfc3339"
	// IntegrationModelCreatedAtColumn is the column json value created_ts
	IntegrationModelCreatedAtColumn = "created_ts"
	// IntegrationModelCustomerIDColumn is the column json value customer_id
	IntegrationModelCustomerIDColumn = "customer_id"
	// IntegrationModelEntityErrorsColumn is the column json value entity_errors
	IntegrationModelEntityErrorsColumn = "entity_errors"
	// IntegrationModelEntityErrorsErrorColumn is the column json value error
	IntegrationModelEntityErrorsErrorColumn = "error"
	// IntegrationModelEntityErrorsIDColumn is the column json value id
	IntegrationModelEntityErrorsIDColumn = "id"
	// IntegrationModelEntityErrorsRefIDColumn is the column json value ref_id
	IntegrationModelEntityErrorsRefIDColumn = "ref_id"
	// IntegrationModelErrorMessageColumn is the column json value error_message
	IntegrationModelErrorMessageColumn = "error_message"
	// IntegrationModelErroredColumn is the column json value errored
	IntegrationModelErroredColumn = "errored"
	// IntegrationModelIDColumn is the column json value id
	IntegrationModelIDColumn = "id"
	// IntegrationModelIntervalColumn is the column json value interval
	IntegrationModelIntervalColumn = "interval"
	// IntegrationModelLastExportCompletedDateColumn is the column json value last_export_completed_date
	IntegrationModelLastExportCompletedDateColumn = "last_export_completed_date"
	// IntegrationModelLastExportCompletedDateEpochColumn is the column json value epoch
	IntegrationModelLastExportCompletedDateEpochColumn = "epoch"
	// IntegrationModelLastExportCompletedDateOffsetColumn is the column json value offset
	IntegrationModelLastExportCompletedDateOffsetColumn = "offset"
	// IntegrationModelLastExportCompletedDateRfc3339Column is the column json value rfc3339
	IntegrationModelLastExportCompletedDateRfc3339Column = "rfc3339"
	// IntegrationModelLastExportRequestedDateColumn is the column json value last_export_requested_date
	IntegrationModelLastExportRequestedDateColumn = "last_export_requested_date"
	// IntegrationModelLastExportRequestedDateEpochColumn is the column json value epoch
	IntegrationModelLastExportRequestedDateEpochColumn = "epoch"
	// IntegrationModelLastExportRequestedDateOffsetColumn is the column json value offset
	IntegrationModelLastExportRequestedDateOffsetColumn = "offset"
	// IntegrationModelLastExportRequestedDateRfc3339Column is the column json value rfc3339
	IntegrationModelLastExportRequestedDateRfc3339Column = "rfc3339"
	// IntegrationModelLastProcessingDateColumn is the column json value last_processing_date
	IntegrationModelLastProcessingDateColumn = "last_processing_date"
	// IntegrationModelLastProcessingDateEpochColumn is the column json value epoch
	IntegrationModelLastProcessingDateEpochColumn = "epoch"
	// IntegrationModelLastProcessingDateOffsetColumn is the column json value offset
	IntegrationModelLastProcessingDateOffsetColumn = "offset"
	// IntegrationModelLastProcessingDateRfc3339Column is the column json value rfc3339
	IntegrationModelLastProcessingDateRfc3339Column = "rfc3339"
	// IntegrationModelLocationColumn is the column json value location
	IntegrationModelLocationColumn = "location"
	// IntegrationModelNameColumn is the column json value name
	IntegrationModelNameColumn = "name"
	// IntegrationModelPausedColumn is the column json value paused
	IntegrationModelPausedColumn = "paused"
	// IntegrationModelProcessedColumn is the column json value processed
	IntegrationModelProcessedColumn = "processed"
	// IntegrationModelRefIDColumn is the column json value ref_id
	IntegrationModelRefIDColumn = "ref_id"
	// IntegrationModelRefTypeColumn is the column json value ref_type
	IntegrationModelRefTypeColumn = "ref_type"
	// IntegrationModelStateColumn is the column json value state
	IntegrationModelStateColumn = "state"
	// IntegrationModelThrottledColumn is the column json value throttled
	IntegrationModelThrottledColumn = "throttled"
	// IntegrationModelThrottledUntilColumn is the column json value throttled_until
	IntegrationModelThrottledUntilColumn = "throttled_until"
	// IntegrationModelThrottledUntilEpochColumn is the column json value epoch
	IntegrationModelThrottledUntilEpochColumn = "epoch"
	// IntegrationModelThrottledUntilOffsetColumn is the column json value offset
	IntegrationModelThrottledUntilOffsetColumn = "offset"
	// IntegrationModelThrottledUntilRfc3339Column is the column json value rfc3339
	IntegrationModelThrottledUntilRfc3339Column = "rfc3339"
	// IntegrationModelUpdatedAtColumn is the column json value updated_ts
	IntegrationModelUpdatedAtColumn = "updated_ts"
)

// IntegrationCreatedDate represents the object structure for created_date
type IntegrationCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationCreatedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationEntityErrors represents the object structure for entity_errors
type IntegrationEntityErrors struct {
	// Error error message if set integration failed when trying to export
	Error string `json:"error" codec:"error" bson:"error" yaml:"error" faker:"-"`
	// ID entity id
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// RefID entity ref_id
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
}

func toIntegrationEntityErrorsObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationEntityErrors:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationEntityErrors) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Error error message if set integration failed when trying to export
		"error": toIntegrationEntityErrorsObject(o.Error, false),
		// ID entity id
		"id": toIntegrationEntityErrorsObject(o.ID, false),
		// RefID entity ref_id
		"ref_id": toIntegrationEntityErrorsObject(o.RefID, false),
	}
}

func (o *IntegrationEntityErrors) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationEntityErrors) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	o.setDefaults(false)
}

// IntegrationLastExportCompletedDate represents the object structure for last_export_completed_date
type IntegrationLastExportCompletedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationLastExportCompletedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationLastExportCompletedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationLastExportCompletedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationLastExportCompletedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationLastExportCompletedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationLastExportCompletedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationLastExportCompletedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationLastExportCompletedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationLastExportRequestedDate represents the object structure for last_export_requested_date
type IntegrationLastExportRequestedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationLastExportRequestedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationLastExportRequestedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationLastExportRequestedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationLastExportRequestedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationLastExportRequestedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationLastExportRequestedDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationLastExportRequestedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationLastExportRequestedDate) FromMap(kv map[string]interface{}) {

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

// IntegrationLastProcessingDate represents the object structure for last_processing_date
type IntegrationLastProcessingDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationLastProcessingDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationLastProcessingDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationLastProcessingDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationLastProcessingDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationLastProcessingDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationLastProcessingDateObject(o.Rfc3339, false),
	}
}

func (o *IntegrationLastProcessingDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationLastProcessingDate) FromMap(kv map[string]interface{}) {

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

// IntegrationLocation is the enumeration type for location
type IntegrationLocation int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationLocation) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationLocation(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "PRIVATE":
			*v = IntegrationLocation(0)
		case "CLOUD":
			*v = IntegrationLocation(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationLocation) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "PRIVATE":
		v = 0
	case "CLOUD":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationLocation) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("PRIVATE")
	case 1:
		return json.Marshal("CLOUD")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Location
func (v IntegrationLocation) String() string {
	switch int32(v) {
	case 0:
		return "PRIVATE"
	case 1:
		return "CLOUD"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationLocation) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationLocation:
		*v = val
	case int32:
		*v = IntegrationLocation(int32(val))
	case int:
		*v = IntegrationLocation(int32(val))
	case string:
		switch val {
		case "PRIVATE":
			*v = IntegrationLocation(0)
		case "CLOUD":
			*v = IntegrationLocation(1)
		}
	}
	return nil
}

const (
	// IntegrationLocationPrivate is the enumeration value for private
	IntegrationLocationPrivate IntegrationLocation = 0
	// IntegrationLocationCloud is the enumeration value for cloud
	IntegrationLocationCloud IntegrationLocation = 1
)

// IntegrationState is the enumeration type for state
type IntegrationState int32

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationState) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationState(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "IDLE":
			*v = IntegrationState(0)
		case "EXPORTING":
			*v = IntegrationState(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v IntegrationState) UnmarshalJSON(buf []byte) error {
	switch string(buf) {
	case "IDLE":
		v = 0
	case "EXPORTING":
		v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationState) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("IDLE")
	case 1:
		return json.Marshal("EXPORTING")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for State
func (v IntegrationState) String() string {
	switch int32(v) {
	case 0:
		return "IDLE"
	case 1:
		return "EXPORTING"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationState) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationState:
		*v = val
	case int32:
		*v = IntegrationState(int32(val))
	case int:
		*v = IntegrationState(int32(val))
	case string:
		switch val {
		case "IDLE":
			*v = IntegrationState(0)
		case "EXPORTING":
			*v = IntegrationState(1)
		}
	}
	return nil
}

const (
	// IntegrationStateIdle is the enumeration value for idle
	IntegrationStateIdle IntegrationState = 0
	// IntegrationStateExporting is the enumeration value for exporting
	IntegrationStateExporting IntegrationState = 1
)

// IntegrationThrottledUntil represents the object structure for throttled_until
type IntegrationThrottledUntil struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toIntegrationThrottledUntilObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *IntegrationThrottledUntil) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toIntegrationThrottledUntilObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toIntegrationThrottledUntilObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toIntegrationThrottledUntilObject(o.Rfc3339, false),
	}
}

func (o *IntegrationThrottledUntil) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationThrottledUntil) FromMap(kv map[string]interface{}) {

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

// Integration The integration instance for a customer
type Integration struct {
	// Active If true, the integration is still active
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// Config the integration configuration controlled by the integration itself
	Config *string `json:"config,omitempty" codec:"config,omitempty" bson:"config" yaml:"config,omitempty" faker:"-"`
	// CreatedByProfileID The id of the profile for the user that created the integration
	CreatedByProfileID *string `json:"created_by_profile_id,omitempty" codec:"created_by_profile_id,omitempty" bson:"created_by_profile_id" yaml:"created_by_profile_id,omitempty" faker:"-"`
	// CreatedByUserID The id of the user that created the integration
	CreatedByUserID *string `json:"created_by_user_id,omitempty" codec:"created_by_user_id,omitempty" bson:"created_by_user_id" yaml:"created_by_user_id,omitempty" faker:"-"`
	// CreatedDate when the integration was created
	CreatedDate IntegrationCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// EntityErrors export status and error per entity in the integration
	EntityErrors []IntegrationEntityErrors `json:"entity_errors" codec:"entity_errors" bson:"entity_errors" yaml:"entity_errors" faker:"-"`
	// ErrorMessage The error message from an export run
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored If authorization failed by the agent or any other error
	Errored *bool `json:"errored,omitempty" codec:"errored,omitempty" bson:"errored" yaml:"errored,omitempty" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Interval the interval in milliseconds for how often an export job is scheduled
	Interval int64 `json:"interval" codec:"interval" bson:"interval" yaml:"interval" faker:"-"`
	// LastExportCompletedDate when the export response was received
	LastExportCompletedDate IntegrationLastExportCompletedDate `json:"last_export_completed_date" codec:"last_export_completed_date" bson:"last_export_completed_date" yaml:"last_export_completed_date" faker:"-"`
	// LastExportRequestedDate when the export request was made
	LastExportRequestedDate IntegrationLastExportRequestedDate `json:"last_export_requested_date" codec:"last_export_requested_date" bson:"last_export_requested_date" yaml:"last_export_requested_date" faker:"-"`
	// LastProcessingDate when the data was last processed
	LastProcessingDate IntegrationLastProcessingDate `json:"last_processing_date" codec:"last_processing_date" bson:"last_processing_date" yaml:"last_processing_date" faker:"-"`
	// Location The location of this integration (on-premise / private or cloud)
	Location IntegrationLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name The user friendly name of the integration
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// Paused true if the agent is paused and should not start new scheduled jobs
	Paused bool `json:"paused" codec:"paused" bson:"paused" yaml:"paused" faker:"-"`
	// Processed If the integration has been processed at least once
	Processed *bool `json:"processed,omitempty" codec:"processed,omitempty" bson:"processed" yaml:"processed,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// State the current state of the integration
	State IntegrationState `json:"state" codec:"state" bson:"state" yaml:"state" faker:"-"`
	// Throttled Set to true when integration is throttled.
	Throttled *bool `json:"throttled,omitempty" codec:"throttled,omitempty" bson:"throttled" yaml:"throttled,omitempty" faker:"-"`
	// ThrottledUntil After throttling integration, we set this field for estimated resume date.
	ThrottledUntil *IntegrationThrottledUntil `json:"throttled_until,omitempty" codec:"throttled_until,omitempty" bson:"throttled_until" yaml:"throttled_until,omitempty" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Integration)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Integration)(nil)

func toIntegrationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Integration:
		return v.ToMap()

	case IntegrationCreatedDate:
		return v.ToMap()

	case []IntegrationEntityErrors:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	case IntegrationLastExportCompletedDate:
		return v.ToMap()

	case IntegrationLastExportRequestedDate:
		return v.ToMap()

	case IntegrationLastProcessingDate:
		return v.ToMap()

	case IntegrationLocation:
		return v.String()

	case IntegrationState:
		return v.String()

	case *IntegrationThrottledUntil:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Integration
func (o *Integration) String() string {
	return fmt.Sprintf("agent.Integration<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Integration) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Integration) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Integration) GetTableName() string {
	return IntegrationTable.String()
}

// GetModelName returns the name of the model
func (o *Integration) GetModelName() datamodel.ModelNameType {
	return IntegrationModelName
}

// NewIntegrationID provides a template for generating an ID field for Integration
func NewIntegrationID(customerID string) string {
	return hash.Values(customerID, datetime.EpochNow(), randomString(64))
}

func (o *Integration) setDefaults(frommap bool) {
	if o.EntityErrors == nil {
		o.EntityErrors = make([]IntegrationEntityErrors, 0)
	}
	if o.ThrottledUntil == nil {
		o.ThrottledUntil = &IntegrationThrottledUntil{}
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, datetime.EpochNow(), randomString(64))
	}

	if o.Errored == nil {
		var v bool
		o.Errored = &v
	}

	{

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
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Integration) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Integration) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Integration) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Integration) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Integration) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Integration) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Integration) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Integration) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Integration) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Integration) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Integration) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Integration
func (o *Integration) Clone() datamodel.Model {
	c := new(Integration)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Integration) Anon() datamodel.Model {
	c := new(Integration)
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
func (o *Integration) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Integration) UnmarshalJSON(data []byte) error {
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
func (o *Integration) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Integration) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Integration objects are equal
func (o *Integration) IsEqual(other *Integration) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Integration) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                     toIntegrationObject(o.Active, false),
		"config":                     toIntegrationObject(o.Config, true),
		"created_by_profile_id":      toIntegrationObject(o.CreatedByProfileID, true),
		"created_by_user_id":         toIntegrationObject(o.CreatedByUserID, true),
		"created_date":               toIntegrationObject(o.CreatedDate, false),
		"created_ts":                 toIntegrationObject(o.CreatedAt, false),
		"customer_id":                toIntegrationObject(o.CustomerID, false),
		"entity_errors":              toIntegrationObject(o.EntityErrors, false),
		"error_message":              toIntegrationObject(o.ErrorMessage, true),
		"errored":                    toIntegrationObject(o.Errored, true),
		"id":                         toIntegrationObject(o.ID, false),
		"interval":                   toIntegrationObject(o.Interval, false),
		"last_export_completed_date": toIntegrationObject(o.LastExportCompletedDate, false),
		"last_export_requested_date": toIntegrationObject(o.LastExportRequestedDate, false),
		"last_processing_date":       toIntegrationObject(o.LastProcessingDate, false),

		"location":  o.Location.String(),
		"name":      toIntegrationObject(o.Name, false),
		"paused":    toIntegrationObject(o.Paused, false),
		"processed": toIntegrationObject(o.Processed, true),
		"ref_id":    toIntegrationObject(o.RefID, false),
		"ref_type":  toIntegrationObject(o.RefType, false),

		"state":           o.State.String(),
		"throttled":       toIntegrationObject(o.Throttled, true),
		"throttled_until": toIntegrationObject(o.ThrottledUntil, true),
		"updated_ts":      toIntegrationObject(o.UpdatedAt, false),
		"hashcode":        toIntegrationObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Integration) FromMap(kv map[string]interface{}) {

	o.ID = ""

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
		} else if sv, ok := val.(IntegrationCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*IntegrationCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.CreatedDate.Epoch = dt.Epoch
			o.CreatedDate.Rfc3339 = dt.Rfc3339
			o.CreatedDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
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

	if val, ok := kv["created_ts"].(int64); ok {
		o.CreatedAt = val
	} else {
		if val, ok := kv["created_ts"]; ok {
			if val == nil {
				o.CreatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.CreatedAt = number.ToInt64Any(val)
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

		o.EntityErrors = make([]IntegrationEntityErrors, 0)

	}
	if val, ok := kv["entity_errors"]; ok {
		if sv, ok := val.([]IntegrationEntityErrors); ok {
			o.EntityErrors = sv
		} else if sp, ok := val.([]*IntegrationEntityErrors); ok {
			o.EntityErrors = o.EntityErrors[:0]
			for _, e := range sp {
				o.EntityErrors = append(o.EntityErrors, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(IntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm IntegrationEntityErrors
					fm.FromMap(av)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else {
					b, _ := json.Marshal(ae)
					bkv := make(map[string]interface{})
					json.Unmarshal(b, &bkv)
					var av IntegrationEntityErrors
					av.FromMap(bkv)
					o.EntityErrors = append(o.EntityErrors, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(IntegrationEntityErrors); ok {
					o.EntityErrors = append(o.EntityErrors, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm IntegrationEntityErrors
					fm.FromMap(r)
					o.EntityErrors = append(o.EntityErrors, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := IntegrationEntityErrors{}
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
							var fm IntegrationEntityErrors
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
		} else if sv, ok := val.(IntegrationLastExportCompletedDate); ok {
			// struct
			o.LastExportCompletedDate = sv
		} else if sp, ok := val.(*IntegrationLastExportCompletedDate); ok {
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
		} else if sv, ok := val.(IntegrationLastExportRequestedDate); ok {
			// struct
			o.LastExportRequestedDate = sv
		} else if sp, ok := val.(*IntegrationLastExportRequestedDate); ok {
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

	if val, ok := kv["last_processing_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastProcessingDate.FromMap(kv)
		} else if sv, ok := val.(IntegrationLastProcessingDate); ok {
			// struct
			o.LastProcessingDate = sv
		} else if sp, ok := val.(*IntegrationLastProcessingDate); ok {
			// struct pointer
			o.LastProcessingDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastProcessingDate.Epoch = dt.Epoch
			o.LastProcessingDate.Rfc3339 = dt.Rfc3339
			o.LastProcessingDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastProcessingDate.Epoch = dt.Epoch
				o.LastProcessingDate.Rfc3339 = dt.Rfc3339
				o.LastProcessingDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastProcessingDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["location"].(IntegrationLocation); ok {
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
	if val, ok := kv["state"].(IntegrationState); ok {
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
		o.ThrottledUntil = &IntegrationThrottledUntil{}
	}

	if val, ok := kv["throttled_until"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.ThrottledUntil.FromMap(kv)
		} else if sv, ok := val.(IntegrationThrottledUntil); ok {
			// struct
			o.ThrottledUntil = &sv
		} else if sp, ok := val.(*IntegrationThrottledUntil); ok {
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

	if val, ok := kv["updated_ts"].(int64); ok {
		o.UpdatedAt = val
	} else {
		if val, ok := kv["updated_ts"]; ok {
			if val == nil {
				o.UpdatedAt = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Integration) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.Config)
	args = append(args, o.CreatedByProfileID)
	args = append(args, o.CreatedByUserID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.EntityErrors)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ID)
	args = append(args, o.Interval)
	args = append(args, o.LastExportCompletedDate)
	args = append(args, o.LastExportRequestedDate)
	args = append(args, o.LastProcessingDate)
	args = append(args, o.Location)
	args = append(args, o.Name)
	args = append(args, o.Paused)
	args = append(args, o.Processed)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.State)
	args = append(args, o.Throttled)
	args = append(args, o.ThrottledUntil)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getIntegrationQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tactive\n")
	// scalar
	sb.WriteString("\t\t\tconfig\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_profile_id\n")
	// scalar
	sb.WriteString("\t\t\tcreated_by_user_id\n")
	// object with fields
	sb.WriteString("\t\t\tcreated_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// object with fields
	sb.WriteString("\t\t\tentity_errors {\n")

	// scalar
	sb.WriteString("\t\t\terror\n")
	// scalar
	sb.WriteString("\t\t\tid\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tinterval\n")
	// object with fields
	sb.WriteString("\t\t\tlast_export_completed_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_export_requested_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_processing_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tlocation\n")
	// scalar
	sb.WriteString("\t\t\tname\n")
	// scalar
	sb.WriteString("\t\t\tpaused\n")
	// scalar
	sb.WriteString("\t\t\tprocessed\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tstate\n")
	// scalar
	sb.WriteString("\t\t\tthrottled\n")
	// object with fields
	sb.WriteString("\t\t\tthrottled_until {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// IntegrationPageInfo is a grapqhl PageInfo
type IntegrationPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// IntegrationConnection is a grapqhl connection
type IntegrationConnection struct {
	Edges      []*IntegrationEdge  `json:"edges,omitempty"`
	PageInfo   IntegrationPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64              `json:"totalCount,omitempty"`
}

// IntegrationEdge is a grapqhl edge
type IntegrationEdge struct {
	Node *Integration `json:"node,omitempty"`
}

// QueryManyIntegrationNode is a grapqhl query many node
type QueryManyIntegrationNode struct {
	Object *IntegrationConnection `json:"Integrations,omitempty"`
}

// QueryManyIntegrationData is a grapqhl query many data node
type QueryManyIntegrationData struct {
	Data *QueryManyIntegrationNode `json:"agent,omitempty"`
}

// QueryOneIntegrationNode is a grapqhl query one node
type QueryOneIntegrationNode struct {
	Object *Integration `json:"Integration,omitempty"`
}

// QueryOneIntegrationData is a grapqhl query one data node
type QueryOneIntegrationData struct {
	Data *QueryOneIntegrationNode `json:"agent,omitempty"`
}

// IntegrationQuery is query struct
type IntegrationQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// IntegrationQueryInput is query input struct
type IntegrationQueryInput struct {
	First  *int64            `json:"first,omitempty"`
	Last   *int64            `json:"last,omitempty"`
	Before *string           `json:"before,omitempty"`
	After  *string           `json:"after,omitempty"`
	Query  *IntegrationQuery `json:"query,omitempty"`
}

// NewIntegrationQuery is a convenience for building a *IntegrationQuery
func NewIntegrationQuery(params ...interface{}) *IntegrationQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &IntegrationQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &IntegrationQueryInput{
		Query: q,
	}
}

// FindIntegration will query an Integration by id
func FindIntegration(client graphql.Client, id string) (*Integration, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query IntegrationQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegration(_id: $id) {\n")
	sb.WriteString(getIntegrationQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrations will query for any Integrations matching the query
func FindIntegrations(client graphql.Client, input *IntegrationQueryInput) (*IntegrationConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
	}
	var sb strings.Builder
	sb.WriteString("query IntegrationQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrations(first: $first last: $last before: $before after: $after query: $query) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getIntegrationQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyIntegrationData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindIntegrationsPaginatedCallback is a callback function for handling each page
type FindIntegrationsPaginatedCallback func(conn *IntegrationConnection) (bool, error)

// FindIntegrationsPaginated will query for any Integrations matching the query and return each page callback
func FindIntegrationsPaginated(client graphql.Client, query *IntegrationQuery, pageSize int64, callback FindIntegrationsPaginatedCallback) error {
	input := &IntegrationQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindIntegrations(client, input)
		if err != nil {
			return err
		}
		if res == nil {
			break
		}
		ok, err := callback(res)
		if err != nil {
			return err
		}
		if !ok || !res.PageInfo.HasNextPage {
			break
		}
		input.After = &res.PageInfo.EndCursor
	}
	return nil
}

// UpdateIntegrationNode is a grapqhl update node
type UpdateIntegrationNode struct {
	Object *Integration `json:"updateIntegration,omitempty"`
}

// UpdateIntegrationData is a grapqhl update data node
type UpdateIntegrationData struct {
	Data *UpdateIntegrationNode `json:"agent,omitempty"`
}

// ExecIntegrationUpdateMutation returns a graphql update mutation result for Integration
func ExecIntegrationUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*Integration, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation IntegrationUpdateMutation($id: String, $input: UpdateAgentIntegrationInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegration(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getIntegrationQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}
