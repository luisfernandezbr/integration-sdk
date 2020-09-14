// DO NOT EDIT -- generated code

// Package calendar - public calendar data models
package calendar

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
)

const (
	// EventTopic is the default topic name
	EventTopic datamodel.TopicNameType = "calendar_Event"

	// EventTable is the default table name
	EventTable datamodel.ModelNameType = "calendar_event"

	// EventModelName is the model name
	EventModelName datamodel.ModelNameType = "calendar.Event"
)

const (
	// EventModelActiveColumn is the column json value active
	EventModelActiveColumn = "active"
	// EventModelAttendeeRefIDColumn is the column json value attendee_ref_id
	EventModelAttendeeRefIDColumn = "attendee_ref_id"
	// EventModelBusyColumn is the column json value busy
	EventModelBusyColumn = "busy"
	// EventModelCalendarIDColumn is the column json value calendar_id
	EventModelCalendarIDColumn = "calendar_id"
	// EventModelCustomerIDColumn is the column json value customer_id
	EventModelCustomerIDColumn = "customer_id"
	// EventModelDescriptionColumn is the column json value description
	EventModelDescriptionColumn = "description"
	// EventModelEndDateColumn is the column json value end_date
	EventModelEndDateColumn = "end_date"
	// EventModelEndDateEpochColumn is the column json value epoch
	EventModelEndDateEpochColumn = "epoch"
	// EventModelEndDateOffsetColumn is the column json value offset
	EventModelEndDateOffsetColumn = "offset"
	// EventModelEndDateRfc3339Column is the column json value rfc3339
	EventModelEndDateRfc3339Column = "rfc3339"
	// EventModelIDColumn is the column json value id
	EventModelIDColumn = "id"
	// EventModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	EventModelIntegrationInstanceIDColumn = "integration_instance_id"
	// EventModelLocationColumn is the column json value location
	EventModelLocationColumn = "location"
	// EventModelLocationDetailsColumn is the column json value details
	EventModelLocationDetailsColumn = "details"
	// EventModelLocationNameColumn is the column json value name
	EventModelLocationNameColumn = "name"
	// EventModelLocationURLColumn is the column json value url
	EventModelLocationURLColumn = "url"
	// EventModelNameColumn is the column json value name
	EventModelNameColumn = "name"
	// EventModelOwnerRefIDColumn is the column json value owner_ref_id
	EventModelOwnerRefIDColumn = "owner_ref_id"
	// EventModelRefIDColumn is the column json value ref_id
	EventModelRefIDColumn = "ref_id"
	// EventModelRefTypeColumn is the column json value ref_type
	EventModelRefTypeColumn = "ref_type"
	// EventModelStartDateColumn is the column json value start_date
	EventModelStartDateColumn = "start_date"
	// EventModelStartDateEpochColumn is the column json value epoch
	EventModelStartDateEpochColumn = "epoch"
	// EventModelStartDateOffsetColumn is the column json value offset
	EventModelStartDateOffsetColumn = "offset"
	// EventModelStartDateRfc3339Column is the column json value rfc3339
	EventModelStartDateRfc3339Column = "rfc3339"
	// EventModelStatusColumn is the column json value status
	EventModelStatusColumn = "status"
	// EventModelUpdatedAtColumn is the column json value updated_ts
	EventModelUpdatedAtColumn = "updated_ts"
)

// EventEndDate represents the object structure for end_date
type EventEndDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEventEndDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventEndDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventEndDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEventEndDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEventEndDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEventEndDateObject(o.Rfc3339, false),
	}
}

func (o *EventEndDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventEndDate) FromMap(kv map[string]interface{}) {

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

// EventLocation represents the object structure for location
type EventLocation struct {
	// Details raw data from the location property the calendar
	Details string `json:"details" codec:"details" bson:"details" yaml:"details" faker:"-"`
	// Name the name of the place
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// URL link to the meeting if any
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"-"`
}

func toEventLocationObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventLocation:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventLocation) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Details raw data from the location property the calendar
		"details": toEventLocationObject(o.Details, false),
		// Name the name of the place
		"name": toEventLocationObject(o.Name, false),
		// URL link to the meeting if any
		"url": toEventLocationObject(o.URL, false),
	}
}

func (o *EventLocation) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventLocation) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["details"].(string); ok {
		o.Details = val
	} else {
		if val, ok := kv["details"]; ok {
			if val == nil {
				o.Details = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Details = fmt.Sprintf("%v", val)
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
	if val, ok := kv["url"].(string); ok {
		o.URL = val
	} else {
		if val, ok := kv["url"]; ok {
			if val == nil {
				o.URL = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// EventStartDate represents the object structure for start_date
type EventStartDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toEventStartDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *EventStartDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *EventStartDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toEventStartDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toEventStartDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toEventStartDateObject(o.Rfc3339, false),
	}
}

func (o *EventStartDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *EventStartDate) FromMap(kv map[string]interface{}) {

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

// EventStatus is the enumeration type for status
type EventStatus int32

// toEventStatusPointer is the enumeration pointer type for status
func toEventStatusPointer(v int32) *EventStatus {
	nv := EventStatus(v)
	return &nv
}

// toEventStatusEnum is the enumeration pointer wrapper for status
func toEventStatusEnum(v *EventStatus) string {
	if v == nil {
		return toEventStatusPointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *EventStatus) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = EventStatus(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "CONFIRMED":
			*v = EventStatus(0)
		case "TENTATIVE":
			*v = EventStatus(1)
		case "CANCELLED":
			*v = EventStatus(2)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *EventStatus) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "CONFIRMED":
		*v = 0
	case "TENTATIVE":
		*v = 1
	case "CANCELLED":
		*v = 2
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v EventStatus) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("CONFIRMED")
	case 1:
		return json.Marshal("TENTATIVE")
	case 2:
		return json.Marshal("CANCELLED")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for Status
func (v EventStatus) String() string {
	switch int32(v) {
	case 0:
		return "CONFIRMED"
	case 1:
		return "TENTATIVE"
	case 2:
		return "CANCELLED"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *EventStatus) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case EventStatus:
		*v = val
	case int32:
		*v = EventStatus(int32(val))
	case int:
		*v = EventStatus(int32(val))
	case string:
		switch val {
		case "CONFIRMED":
			*v = EventStatus(0)
		case "TENTATIVE":
			*v = EventStatus(1)
		case "CANCELLED":
			*v = EventStatus(2)
		}
	}
	return nil
}

const (
	// EventStatusConfirmed is the enumeration value for confirmed
	EventStatusConfirmed EventStatus = 0
	// EventStatusTentative is the enumeration value for tentative
	EventStatusTentative EventStatus = 1
	// EventStatusCancelled is the enumeration value for cancelled
	EventStatusCancelled EventStatus = 2
)

// Event event on a calendar
type Event struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// AttendeeRefID ref_id of the user from whom's calendar this event was exported
	AttendeeRefID string `json:"attendee_ref_id" codec:"attendee_ref_id" bson:"attendee_ref_id" yaml:"attendee_ref_id" faker:"-"`
	// Busy true if the user is marked as busy in this event
	Busy bool `json:"busy" codec:"busy" bson:"busy" yaml:"busy" faker:"-"`
	// CalendarID unique project id this event belongs to (deprecated)
	CalendarID string `json:"calendar_id" codec:"calendar_id" bson:"calendar_id" yaml:"calendar_id" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Description the description of the event
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"-"`
	// EndDate the event end date
	EndDate EventEndDate `json:"end_date" codec:"end_date" bson:"end_date" yaml:"end_date" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Location location of the event, could be a place or a link to a meeting
	Location EventLocation `json:"location" codec:"location" bson:"location" yaml:"location" faker:"-"`
	// Name the name of the event
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// OwnerRefID owner ref_id of the event
	OwnerRefID string `json:"owner_ref_id" codec:"owner_ref_id" bson:"owner_ref_id" yaml:"owner_ref_id" faker:"-"`
	// RefID the event ref ID
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// StartDate the event start date
	StartDate EventStartDate `json:"start_date" codec:"start_date" bson:"start_date" yaml:"start_date" faker:"-"`
	// Status the status of the user in this event
	Status EventStatus `json:"status" codec:"status" bson:"status" yaml:"status" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Event)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Event)(nil)

func toEventObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Event:
		return v.ToMap()

	case EventEndDate:
		return v.ToMap()

	case EventLocation:
		return v.ToMap()

	case EventStartDate:
		return v.ToMap()

	case EventStatus:
		return v.String()

	default:
		return o
	}
}

// String returns a string representation of Event
func (o *Event) String() string {
	return fmt.Sprintf("calendar.Event<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Event) GetTopicName() datamodel.TopicNameType {
	return EventTopic
}

// GetStreamName returns the name of the stream
func (o *Event) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Event) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Event) GetModelName() datamodel.ModelNameType {
	return EventModelName
}

// NewEventID provides a template for generating an ID field for Event
func NewEventID(customerID string, refType string, refID string, OwnerRefID string) string {
	return hash.Values(customerID, refType, refID, OwnerRefID)
}

func (o *Event) setDefaults(frommap bool) {

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefType, o.RefID, o.OwnerRefID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Event) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Event) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Event) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Event")
}

// GetRefID returns the RefID for the object
func (o *Event) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Event) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Event) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Event) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Event) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Event) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = EventModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Event) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "id",
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
func (o *Event) GetCustomerID() string {

	return o.CustomerID

}

// SetCustomerID will return the customer_id
func (o *Event) SetCustomerID(id string) {

	o.CustomerID = id

}

// GetRefType will return the ref_type
func (o *Event) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *Event) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of Event
func (o *Event) Clone() datamodel.Model {
	c := new(Event)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Event) Anon() datamodel.Model {
	c := new(Event)
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
func (o *Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Event) UnmarshalJSON(data []byte) error {
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
func (o *Event) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Event) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Event objects are equal
func (o *Event) IsEqual(other *Event) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Event) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":                  toEventObject(o.Active, false),
		"attendee_ref_id":         toEventObject(o.AttendeeRefID, false),
		"busy":                    toEventObject(o.Busy, false),
		"calendar_id":             toEventObject(o.CalendarID, false),
		"customer_id":             toEventObject(o.CustomerID, false),
		"description":             toEventObject(o.Description, false),
		"end_date":                toEventObject(o.EndDate, false),
		"id":                      toEventObject(o.ID, false),
		"integration_instance_id": toEventObject(o.IntegrationInstanceID, true),
		"location":                toEventObject(o.Location, false),
		"name":                    toEventObject(o.Name, false),
		"owner_ref_id":            toEventObject(o.OwnerRefID, false),
		"ref_id":                  toEventObject(o.RefID, false),
		"ref_type":                toEventObject(o.RefType, false),
		"start_date":              toEventObject(o.StartDate, false),

		"status":     o.Status.String(),
		"updated_ts": toEventObject(o.UpdatedAt, false),
		"hashcode":   toEventObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Event) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["attendee_ref_id"].(string); ok {
		o.AttendeeRefID = val
	} else {
		if val, ok := kv["attendee_ref_id"]; ok {
			if val == nil {
				o.AttendeeRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.AttendeeRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["busy"].(bool); ok {
		o.Busy = val
	} else {
		if val, ok := kv["busy"]; ok {
			if val == nil {
				o.Busy = false
			} else {
				o.Busy = number.ToBoolAny(val)
			}
		}
	}
	if val, ok := kv["calendar_id"].(string); ok {
		o.CalendarID = val
	} else {
		if val, ok := kv["calendar_id"]; ok {
			if val == nil {
				o.CalendarID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CalendarID = fmt.Sprintf("%v", val)
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
	if val, ok := kv["description"].(string); ok {
		o.Description = val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Description = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(EventEndDate); ok {
			// struct
			o.EndDate = sv
		} else if sp, ok := val.(*EventEndDate); ok {
			// struct pointer
			o.EndDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["location"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Location.FromMap(kv)
		} else if sv, ok := val.(EventLocation); ok {
			// struct
			o.Location = sv
		} else if sp, ok := val.(*EventLocation); ok {
			// struct pointer
			o.Location = *sp
		}
	} else {
		o.Location.FromMap(map[string]interface{}{})
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
	if val, ok := kv["owner_ref_id"].(string); ok {
		o.OwnerRefID = val
	} else {
		if val, ok := kv["owner_ref_id"]; ok {
			if val == nil {
				o.OwnerRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.OwnerRefID = fmt.Sprintf("%v", val)
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

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(EventStartDate); ok {
			// struct
			o.StartDate = sv
		} else if sp, ok := val.(*EventStartDate); ok {
			// struct pointer
			o.StartDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["status"].(EventStatus); ok {
		o.Status = val
	} else {
		if em, ok := kv["status"].(map[string]interface{}); ok {

			ev := em["calendar.status"].(string)
			switch ev {
			case "confirmed", "CONFIRMED":
				o.Status = 0
			case "tentative", "TENTATIVE":
				o.Status = 1
			case "cancelled", "CANCELLED":
				o.Status = 2
			}
		}
		if em, ok := kv["status"].(string); ok {
			switch em {
			case "confirmed", "CONFIRMED":
				o.Status = 0
			case "tentative", "TENTATIVE":
				o.Status = 1
			case "cancelled", "CANCELLED":
				o.Status = 2
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Event) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.AttendeeRefID)
	args = append(args, o.Busy)
	args = append(args, o.CalendarID)
	args = append(args, o.CustomerID)
	args = append(args, o.Description)
	args = append(args, o.EndDate)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Location)
	args = append(args, o.Name)
	args = append(args, o.OwnerRefID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.StartDate)
	args = append(args, o.Status)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *Event) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *Event) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// EventPartial is a partial struct for upsert mutations for Event
type EventPartial struct {
	// Active indicates that this model is displayed in a source system, false if the model is deleted
	Active *bool `json:"active,omitempty"`
	// AttendeeRefID ref_id of the user from whom's calendar this event was exported
	AttendeeRefID *string `json:"attendee_ref_id,omitempty"`
	// Busy true if the user is marked as busy in this event
	Busy *bool `json:"busy,omitempty"`
	// CalendarID unique project id this event belongs to (deprecated)
	CalendarID *string `json:"calendar_id,omitempty"`
	// Description the description of the event
	Description *string `json:"description,omitempty"`
	// EndDate the event end date
	EndDate *EventEndDate `json:"end_date,omitempty"`
	// Location location of the event, could be a place or a link to a meeting
	Location *EventLocation `json:"location,omitempty"`
	// Name the name of the event
	Name *string `json:"name,omitempty"`
	// OwnerRefID owner ref_id of the event
	OwnerRefID *string `json:"owner_ref_id,omitempty"`
	// RefID the event ref ID
	RefID *string `json:"ref_id,omitempty"`
	// RefType the record type
	RefType *string `json:"ref_type,omitempty"`
	// StartDate the event start date
	StartDate *EventStartDate `json:"start_date,omitempty"`
	// Status the status of the user in this event
	Status *EventStatus `json:"status,omitempty"`
}

var _ datamodel.PartialModel = (*EventPartial)(nil)

// GetModelName returns the name of the model
func (o *EventPartial) GetModelName() datamodel.ModelNameType {
	return EventModelName
}

// ToMap returns the object as a map
func (o *EventPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"active":          toEventObject(o.Active, true),
		"attendee_ref_id": toEventObject(o.AttendeeRefID, true),
		"busy":            toEventObject(o.Busy, true),
		"calendar_id":     toEventObject(o.CalendarID, true),
		"description":     toEventObject(o.Description, true),
		"end_date":        toEventObject(o.EndDate, true),
		"location":        toEventObject(o.Location, true),
		"name":            toEventObject(o.Name, true),
		"owner_ref_id":    toEventObject(o.OwnerRefID, true),
		"ref_id":          toEventObject(o.RefID, true),
		"ref_type":        toEventObject(o.RefType, true),
		"start_date":      toEventObject(o.StartDate, true),

		"status": toEventStatusEnum(o.Status),
	}
	for k, v := range kv {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(kv, k)
		} else {
			if k == "end_date" {
				if dt, ok := v.(*EventEndDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
			if k == "start_date" {
				if dt, ok := v.(*EventStartDate); ok {
					if dt.Epoch == 0 && dt.Offset == 0 && dt.Rfc3339 == "" {
						delete(kv, k)
					}
				}
			}
		}
	}
	return kv
}

// Stringify returns the object in JSON format as a string
func (o *EventPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *EventPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *EventPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *EventPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *EventPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["active"].(*bool); ok {
		o.Active = val
	} else if val, ok := kv["active"].(bool); ok {
		o.Active = &val
	} else {
		if val, ok := kv["active"]; ok {
			if val == nil {
				o.Active = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Active = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["attendee_ref_id"].(*string); ok {
		o.AttendeeRefID = val
	} else if val, ok := kv["attendee_ref_id"].(string); ok {
		o.AttendeeRefID = &val
	} else {
		if val, ok := kv["attendee_ref_id"]; ok {
			if val == nil {
				o.AttendeeRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.AttendeeRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["busy"].(*bool); ok {
		o.Busy = val
	} else if val, ok := kv["busy"].(bool); ok {
		o.Busy = &val
	} else {
		if val, ok := kv["busy"]; ok {
			if val == nil {
				o.Busy = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Busy = number.BoolPointer(number.ToBoolAny(val))
			}
		}
	}
	if val, ok := kv["calendar_id"].(*string); ok {
		o.CalendarID = val
	} else if val, ok := kv["calendar_id"].(string); ok {
		o.CalendarID = &val
	} else {
		if val, ok := kv["calendar_id"]; ok {
			if val == nil {
				o.CalendarID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.CalendarID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["description"].(*string); ok {
		o.Description = val
	} else if val, ok := kv["description"].(string); ok {
		o.Description = &val
	} else {
		if val, ok := kv["description"]; ok {
			if val == nil {
				o.Description = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Description = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.EndDate == nil {
		o.EndDate = &EventEndDate{}
	}

	if val, ok := kv["end_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EndDate.FromMap(kv)
		} else if sv, ok := val.(EventEndDate); ok {
			// struct
			o.EndDate = &sv
		} else if sp, ok := val.(*EventEndDate); ok {
			// struct pointer
			o.EndDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EndDate.Epoch = dt.Epoch
			o.EndDate.Rfc3339 = dt.Rfc3339
			o.EndDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if o.Location == nil {
		o.Location = &EventLocation{}
	}

	if val, ok := kv["location"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Location.FromMap(kv)
		} else if sv, ok := val.(EventLocation); ok {
			// struct
			o.Location = &sv
		} else if sp, ok := val.(*EventLocation); ok {
			// struct pointer
			o.Location = sp
		}
	} else {
		o.Location.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["name"].(*string); ok {
		o.Name = val
	} else if val, ok := kv["name"].(string); ok {
		o.Name = &val
	} else {
		if val, ok := kv["name"]; ok {
			if val == nil {
				o.Name = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Name = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["owner_ref_id"].(*string); ok {
		o.OwnerRefID = val
	} else if val, ok := kv["owner_ref_id"].(string); ok {
		o.OwnerRefID = &val
	} else {
		if val, ok := kv["owner_ref_id"]; ok {
			if val == nil {
				o.OwnerRefID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.OwnerRefID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	if val, ok := kv["ref_type"].(*string); ok {
		o.RefType = val
	} else if val, ok := kv["ref_type"].(string); ok {
		o.RefType = &val
	} else {
		if val, ok := kv["ref_type"]; ok {
			if val == nil {
				o.RefType = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RefType = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if o.StartDate == nil {
		o.StartDate = &EventStartDate{}
	}

	if val, ok := kv["start_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.StartDate.FromMap(kv)
		} else if sv, ok := val.(EventStartDate); ok {
			// struct
			o.StartDate = &sv
		} else if sp, ok := val.(*EventStartDate); ok {
			// struct pointer
			o.StartDate = sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.StartDate.Epoch = dt.Epoch
			o.StartDate.Rfc3339 = dt.Rfc3339
			o.StartDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt := datetime.NewDateWithTime(tv)
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

	if val, ok := kv["status"].(*EventStatus); ok {
		o.Status = val
	} else if val, ok := kv["status"].(EventStatus); ok {
		o.Status = &val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = toEventStatusPointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["EventStatus"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "confirmed", "CONFIRMED":
						o.Status = toEventStatusPointer(0)
					case "tentative", "TENTATIVE":
						o.Status = toEventStatusPointer(1)
					case "cancelled", "CANCELLED":
						o.Status = toEventStatusPointer(2)
					}
				}
			}
		}
	}
	o.setDefaults(false)
}
