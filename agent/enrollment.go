// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"encoding/json"
	"fmt"
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
)

const (

	// EnrollmentTable is the default table name
	EnrollmentTable datamodel.ModelNameType = "agent_enrollment"

	// EnrollmentModelName is the model name
	EnrollmentModelName datamodel.ModelNameType = "agent.Enrollment"
)

const (
	// EnrollmentModelAgentVersionColumn is the column json value agent_version
	EnrollmentModelAgentVersionColumn = "agent_version"
	// EnrollmentModelArchitectureColumn is the column json value architecture
	EnrollmentModelArchitectureColumn = "architecture"
	// EnrollmentModelCreatedDateColumn is the column json value created_date
	EnrollmentModelCreatedDateColumn = "created_date"
	// EnrollmentModelCreatedDateEpochColumn is the column json value epoch
	EnrollmentModelCreatedDateEpochColumn = "epoch"
	// EnrollmentModelCreatedDateOffsetColumn is the column json value offset
	EnrollmentModelCreatedDateOffsetColumn = "offset"
	// EnrollmentModelCreatedDateRfc3339Column is the column json value rfc3339
	EnrollmentModelCreatedDateRfc3339Column = "rfc3339"
	// EnrollmentModelCreatedAtColumn is the column json value created_ts
	EnrollmentModelCreatedAtColumn = "created_ts"
	// EnrollmentModelCustomerIDColumn is the column json value customer_id
	EnrollmentModelCustomerIDColumn = "customer_id"
	// EnrollmentModelGoVersionColumn is the column json value go_version
	EnrollmentModelGoVersionColumn = "go_version"
	// EnrollmentModelHostnameColumn is the column json value hostname
	EnrollmentModelHostnameColumn = "hostname"
	// EnrollmentModelIDColumn is the column json value id
	EnrollmentModelIDColumn = "id"
	// EnrollmentModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	EnrollmentModelIntegrationInstanceIDColumn = "integration_instance_id"
	// EnrollmentModelLastPingDateColumn is the column json value last_ping_date
	EnrollmentModelLastPingDateColumn = "last_ping_date"
	// EnrollmentModelLastPingDateEpochColumn is the column json value epoch
	EnrollmentModelLastPingDateEpochColumn = "epoch"
	// EnrollmentModelLastPingDateOffsetColumn is the column json value offset
	EnrollmentModelLastPingDateOffsetColumn = "offset"
	// EnrollmentModelLastPingDateRfc3339Column is the column json value rfc3339
	EnrollmentModelLastPingDateRfc3339Column = "rfc3339"
	// EnrollmentModelLastShutdownDateColumn is the column json value last_shutdown_date
	EnrollmentModelLastShutdownDateColumn = "last_shutdown_date"
	// EnrollmentModelLastShutdownDateEpochColumn is the column json value epoch
	EnrollmentModelLastShutdownDateEpochColumn = "epoch"
	// EnrollmentModelLastShutdownDateOffsetColumn is the column json value offset
	EnrollmentModelLastShutdownDateOffsetColumn = "offset"
	// EnrollmentModelLastShutdownDateRfc3339Column is the column json value rfc3339
	EnrollmentModelLastShutdownDateRfc3339Column = "rfc3339"
	// EnrollmentModelLastStartupDateColumn is the column json value last_startup_date
	EnrollmentModelLastStartupDateColumn = "last_startup_date"
	// EnrollmentModelLastStartupDateEpochColumn is the column json value epoch
	EnrollmentModelLastStartupDateEpochColumn = "epoch"
	// EnrollmentModelLastStartupDateOffsetColumn is the column json value offset
	EnrollmentModelLastStartupDateOffsetColumn = "offset"
	// EnrollmentModelLastStartupDateRfc3339Column is the column json value rfc3339
	EnrollmentModelLastStartupDateRfc3339Column = "rfc3339"
	// EnrollmentModelNumCPUColumn is the column json value num_cpu
	EnrollmentModelNumCPUColumn = "num_cpu"
	// EnrollmentModelOSColumn is the column json value os
	EnrollmentModelOSColumn = "os"
	// EnrollmentModelRefIDColumn is the column json value ref_id
	EnrollmentModelRefIDColumn = "ref_id"
	// EnrollmentModelRefTypeColumn is the column json value ref_type
	EnrollmentModelRefTypeColumn = "ref_type"
	// EnrollmentModelRunningColumn is the column json value running
	EnrollmentModelRunningColumn = "running"
	// EnrollmentModelSystemIDColumn is the column json value system_id
	EnrollmentModelSystemIDColumn = "system_id"
	// EnrollmentModelUpdatedAtColumn is the column json value updated_ts
	EnrollmentModelUpdatedAtColumn = "updated_ts"
	// EnrollmentModelUserIDColumn is the column json value user_id
	EnrollmentModelUserIDColumn = "user_id"
)

// EnrollmentCreatedDate represents the object structure for created_date
type EnrollmentCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollmentCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollmentCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollmentCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollmentCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollmentCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollmentCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollmentCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollmentCreatedDate) FromMap(kv map[string]interface{}) {

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

// EnrollmentLastPingDate represents the object structure for last_ping_date
type EnrollmentLastPingDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollmentLastPingDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollmentLastPingDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollmentLastPingDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollmentLastPingDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollmentLastPingDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollmentLastPingDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollmentLastPingDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollmentLastPingDate) FromMap(kv map[string]interface{}) {

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

// EnrollmentLastShutdownDate represents the object structure for last_shutdown_date
type EnrollmentLastShutdownDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollmentLastShutdownDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollmentLastShutdownDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollmentLastShutdownDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollmentLastShutdownDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollmentLastShutdownDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollmentLastShutdownDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollmentLastShutdownDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollmentLastShutdownDate) FromMap(kv map[string]interface{}) {

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

// EnrollmentLastStartupDate represents the object structure for last_startup_date
type EnrollmentLastStartupDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEnrollmentLastStartupDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EnrollmentLastStartupDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EnrollmentLastStartupDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEnrollmentLastStartupDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEnrollmentLastStartupDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEnrollmentLastStartupDateObject(o.Rfc3339, false),
	}
}

func (o *EnrollmentLastStartupDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EnrollmentLastStartupDate) FromMap(kv map[string]interface{}) {

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

// Enrollment The integration instance for a customer
type Enrollment struct {
	// AgentVersion the version of the agent
	AgentVersion string `json:"agent_version" codec:"agent_version" bson:"agent_version" yaml:"agent_version" faker:"-"`
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" codec:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CreatedDate when the agent was created
	CreatedDate EnrollmentCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" codec:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" codec:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// LastPingDate last time the agent pinged
	LastPingDate EnrollmentLastPingDate `json:"last_ping_date" codec:"last_ping_date" bson:"last_ping_date" yaml:"last_ping_date" faker:"-"`
	// LastShutdownDate last time the agent shut down
	LastShutdownDate EnrollmentLastShutdownDate `json:"last_shutdown_date" codec:"last_shutdown_date" bson:"last_shutdown_date" yaml:"last_shutdown_date" faker:"-"`
	// LastStartupDate last time the agent started up
	LastStartupDate EnrollmentLastStartupDate `json:"last_startup_date" codec:"last_startup_date" bson:"last_startup_date" yaml:"last_startup_date" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" codec:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" codec:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Running if the agent is alive
	Running bool `json:"running" codec:"running" bson:"running" yaml:"running" faker:"-"`
	// SystemID the system id
	SystemID string `json:"system_id" codec:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UserID the user that created this enrollment
	UserID string `json:"user_id" codec:"user_id" bson:"user_id" yaml:"user_id" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Enrollment)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Enrollment)(nil)

func toEnrollmentObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Enrollment:
		return v.ToMap()

	case EnrollmentCreatedDate:
		return v.ToMap()

	case EnrollmentLastPingDate:
		return v.ToMap()

	case EnrollmentLastShutdownDate:
		return v.ToMap()

	case EnrollmentLastStartupDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Enrollment
func (o *Enrollment) String() string {
	return fmt.Sprintf("agent.Enrollment<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Enrollment) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *Enrollment) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Enrollment) GetTableName() string {
	return EnrollmentTable.String()
}

// GetModelName returns the name of the model
func (o *Enrollment) GetModelName() datamodel.ModelNameType {
	return EnrollmentModelName
}

// NewEnrollmentID provides a template for generating an ID field for Enrollment
func NewEnrollmentID(customerID string, SystemID string) string {
	return hash.Values(customerID, SystemID)
}

func (o *Enrollment) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.SystemID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Enrollment) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Enrollment) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Enrollment) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *Enrollment) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Enrollment) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Enrollment) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Enrollment) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Enrollment) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *Enrollment) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EnrollmentModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Enrollment) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *Enrollment) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Enrollment
func (o *Enrollment) Clone() datamodel.Model {
	c := new(Enrollment)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Enrollment) Anon() datamodel.Model {
	c := new(Enrollment)
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
func (o *Enrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Enrollment) UnmarshalJSON(data []byte) error {
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
func (o *Enrollment) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Enrollment) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Enrollment objects are equal
func (o *Enrollment) IsEqual(other *Enrollment) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Enrollment) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"agent_version":           toEnrollmentObject(o.AgentVersion, false),
		"architecture":            toEnrollmentObject(o.Architecture, false),
		"created_date":            toEnrollmentObject(o.CreatedDate, false),
		"created_ts":              toEnrollmentObject(o.CreatedAt, false),
		"customer_id":             toEnrollmentObject(o.CustomerID, false),
		"go_version":              toEnrollmentObject(o.GoVersion, false),
		"hostname":                toEnrollmentObject(o.Hostname, false),
		"id":                      toEnrollmentObject(o.ID, false),
		"integration_instance_id": toEnrollmentObject(o.IntegrationInstanceID, true),
		"last_ping_date":          toEnrollmentObject(o.LastPingDate, false),
		"last_shutdown_date":      toEnrollmentObject(o.LastShutdownDate, false),
		"last_startup_date":       toEnrollmentObject(o.LastStartupDate, false),
		"num_cpu":                 toEnrollmentObject(o.NumCPU, false),
		"os":                      toEnrollmentObject(o.OS, false),
		"ref_id":                  toEnrollmentObject(o.RefID, false),
		"ref_type":                toEnrollmentObject(o.RefType, false),
		"running":                 toEnrollmentObject(o.Running, false),
		"system_id":               toEnrollmentObject(o.SystemID, false),
		"updated_ts":              toEnrollmentObject(o.UpdatedAt, false),
		"user_id":                 toEnrollmentObject(o.UserID, false),
		"hashcode":                toEnrollmentObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Enrollment) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["agent_version"].(string); ok {
		o.AgentVersion = val
	} else {
		if val, ok := kv["agent_version"]; ok {
			if val == nil {
				o.AgentVersion = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.AgentVersion = fmt.Sprintf("%v", val)
			}
		}
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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(EnrollmentCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*EnrollmentCreatedDate); ok {
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

	if val, ok := kv["last_ping_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastPingDate.FromMap(kv)
		} else if sv, ok := val.(EnrollmentLastPingDate); ok {
			// struct
			o.LastPingDate = sv
		} else if sp, ok := val.(*EnrollmentLastPingDate); ok {
			// struct pointer
			o.LastPingDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastPingDate.Epoch = dt.Epoch
			o.LastPingDate.Rfc3339 = dt.Rfc3339
			o.LastPingDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastPingDate.Epoch = dt.Epoch
			o.LastPingDate.Rfc3339 = dt.Rfc3339
			o.LastPingDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastPingDate.Epoch = dt.Epoch
				o.LastPingDate.Rfc3339 = dt.Rfc3339
				o.LastPingDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastPingDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_shutdown_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastShutdownDate.FromMap(kv)
		} else if sv, ok := val.(EnrollmentLastShutdownDate); ok {
			// struct
			o.LastShutdownDate = sv
		} else if sp, ok := val.(*EnrollmentLastShutdownDate); ok {
			// struct pointer
			o.LastShutdownDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastShutdownDate.Epoch = dt.Epoch
			o.LastShutdownDate.Rfc3339 = dt.Rfc3339
			o.LastShutdownDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastShutdownDate.Epoch = dt.Epoch
			o.LastShutdownDate.Rfc3339 = dt.Rfc3339
			o.LastShutdownDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastShutdownDate.Epoch = dt.Epoch
				o.LastShutdownDate.Rfc3339 = dt.Rfc3339
				o.LastShutdownDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastShutdownDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["last_startup_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastStartupDate.FromMap(kv)
		} else if sv, ok := val.(EnrollmentLastStartupDate); ok {
			// struct
			o.LastStartupDate = sv
		} else if sp, ok := val.(*EnrollmentLastStartupDate); ok {
			// struct pointer
			o.LastStartupDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastStartupDate.Epoch = dt.Epoch
			o.LastStartupDate.Rfc3339 = dt.Rfc3339
			o.LastStartupDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastStartupDate.Epoch = dt.Epoch
			o.LastStartupDate.Rfc3339 = dt.Rfc3339
			o.LastStartupDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastStartupDate.Epoch = dt.Epoch
				o.LastStartupDate.Rfc3339 = dt.Rfc3339
				o.LastStartupDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastStartupDate.FromMap(map[string]interface{}{})
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
	if val, ok := kv["running"].(bool); ok {
		o.Running = val
	} else {
		if val, ok := kv["running"]; ok {
			if val == nil {
				o.Running = false
			} else {
				o.Running = number.ToBoolAny(val)
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
	if val, ok := kv["user_id"].(string); ok {
		o.UserID = val
	} else {
		if val, ok := kv["user_id"]; ok {
			if val == nil {
				o.UserID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.UserID = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Enrollment) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AgentVersion)
	args = append(args, o.Architecture)
	args = append(args, o.CreatedDate)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.LastPingDate)
	args = append(args, o.LastShutdownDate)
	args = append(args, o.LastStartupDate)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Running)
	args = append(args, o.SystemID)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UserID)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getEnrollmentQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tagent_version\n")
	// scalar
	sb.WriteString("\t\t\tarchitecture\n")
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
	// scalar
	sb.WriteString("\t\t\tgo_version\n")
	// scalar
	sb.WriteString("\t\t\thostname\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// object with fields
	sb.WriteString("\t\t\tlast_ping_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_shutdown_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// object with fields
	sb.WriteString("\t\t\tlast_startup_date {\n")

	// scalar
	sb.WriteString("\t\t\tepoch\n")
	// scalar
	sb.WriteString("\t\t\toffset\n")
	// scalar
	sb.WriteString("\t\t\trfc3339\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tnum_cpu\n")
	// scalar
	sb.WriteString("\t\t\tos\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\trunning\n")
	// scalar
	sb.WriteString("\t\t\tsystem_id\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// scalar
	sb.WriteString("\t\t\tuser_id\n")
	return sb.String()
}

// EnrollmentPageInfo is a grapqhl PageInfo
type EnrollmentPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// EnrollmentConnection is a grapqhl connection
type EnrollmentConnection struct {
	Edges      []*EnrollmentEdge  `json:"edges,omitempty"`
	PageInfo   EnrollmentPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64             `json:"totalCount,omitempty"`
}

// EnrollmentEdge is a grapqhl edge
type EnrollmentEdge struct {
	Node *Enrollment `json:"node,omitempty"`
}

// QueryManyEnrollmentNode is a grapqhl query many node
type QueryManyEnrollmentNode struct {
	Object *EnrollmentConnection `json:"Enrollments,omitempty"`
}

// QueryManyEnrollmentData is a grapqhl query many data node
type QueryManyEnrollmentData struct {
	Data *QueryManyEnrollmentNode `json:"agent,omitempty"`
}

// QueryOneEnrollmentNode is a grapqhl query one node
type QueryOneEnrollmentNode struct {
	Object *Enrollment `json:"Enrollment,omitempty"`
}

// QueryOneEnrollmentData is a grapqhl query one data node
type QueryOneEnrollmentData struct {
	Data *QueryOneEnrollmentNode `json:"agent,omitempty"`
}

// EnrollmentQuery is query struct
type EnrollmentQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// EnrollmentOrder is the order direction
type EnrollmentOrder *string

var (
	// EnrollmentOrderDesc is descending
	EnrollmentOrderDesc EnrollmentOrder = pstrings.Pointer("DESC")
	// EnrollmentOrderAsc is ascending
	EnrollmentOrderAsc EnrollmentOrder = pstrings.Pointer("ASC")
)

// EnrollmentQueryInput is query input struct
type EnrollmentQueryInput struct {
	First   *int64           `json:"first,omitempty"`
	Last    *int64           `json:"last,omitempty"`
	Before  *string          `json:"before,omitempty"`
	After   *string          `json:"after,omitempty"`
	Query   *EnrollmentQuery `json:"query,omitempty"`
	OrderBy *string          `json:"orderBy,omitempty"`
	Order   EnrollmentOrder  `json:"order,omitempty"`
}

// NewEnrollmentQuery is a convenience for building a *EnrollmentQuery
func NewEnrollmentQuery(params ...interface{}) *EnrollmentQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &EnrollmentQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &EnrollmentQueryInput{
		Query: q,
	}
}

// FindEnrollment will query an Enrollment by id
func FindEnrollment(client graphql.Client, id string) (*Enrollment, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query EnrollmentQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tEnrollment(_id: $id) {\n")
	sb.WriteString(getEnrollmentQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneEnrollmentData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindEnrollments will query for any Enrollments matching the query
func FindEnrollments(client graphql.Client, input *EnrollmentQueryInput) (*EnrollmentConnection, error) {
	variables := make(graphql.Variables)
	if input != nil {
		variables["first"] = input.First
		variables["last"] = input.Last
		variables["before"] = input.Before
		variables["after"] = input.After
		variables["query"] = input.Query
		if input.OrderBy != nil {
			variables["orderBy"] = *input.OrderBy
		}
		if input.Order != nil {
			variables["order"] = *input.Order
		}
	}
	var sb strings.Builder
	sb.WriteString("query EnrollmentQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentEnrollmentColumnEnum) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tEnrollments(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getEnrollmentQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyEnrollmentData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindEnrollmentsPaginatedCallback is a callback function for handling each page
type FindEnrollmentsPaginatedCallback func(conn *EnrollmentConnection) (bool, error)

// FindEnrollmentsPaginated will query for any Enrollments matching the query and return each page callback
func FindEnrollmentsPaginated(client graphql.Client, query *EnrollmentQuery, pageSize int64, callback FindEnrollmentsPaginatedCallback) error {
	input := &EnrollmentQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindEnrollments(client, input)
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

// UpdateEnrollmentNode is a grapqhl update node
type UpdateEnrollmentNode struct {
	Object *Enrollment `json:"updateEnrollment,omitempty"`
}

// UpdateEnrollmentData is a grapqhl update data node
type UpdateEnrollmentData struct {
	Data *UpdateEnrollmentNode `json:"agent,omitempty"`
}

// ExecEnrollmentUpdateMutation returns a graphql update mutation result for Enrollment
func ExecEnrollmentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*Enrollment, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation EnrollmentUpdateMutation($id: String, $input: UpdateAgentEnrollmentInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateEnrollment(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getEnrollmentQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateEnrollmentData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecEnrollmentSilentUpdateMutation returns a graphql update mutation result for Enrollment
func ExecEnrollmentSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation EnrollmentUpdateMutation($id: String, $input: UpdateAgentEnrollmentInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateEnrollment(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateEnrollmentData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateEnrollment(client graphql.Client, model Enrollment) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CreateEnrollment($input: CreateAgentEnrollmentInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateEnrollment(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateEnrollmentData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}
