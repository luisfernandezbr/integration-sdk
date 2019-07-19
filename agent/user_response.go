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
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// UserResponseTopic is the default topic name
	UserResponseTopic datamodel.TopicNameType = "agent_UserResponse_topic"

	// UserResponseStream is the default stream name
	UserResponseStream datamodel.TopicNameType = "agent_UserResponse_stream"

	// UserResponseTable is the default table name
	UserResponseTable datamodel.TopicNameType = "agent_UserResponse"

	// UserResponseModelName is the model name
	UserResponseModelName datamodel.ModelNameType = "agent.UserResponse"
)

const (
	// UserResponseArchitectureColumn is the architecture column name
	UserResponseArchitectureColumn = "architecture"
	// UserResponseCustomerIDColumn is the customer_id column name
	UserResponseCustomerIDColumn = "customer_id"
	// UserResponseDataColumn is the data column name
	UserResponseDataColumn = "data"
	// UserResponseDistroColumn is the distro column name
	UserResponseDistroColumn = "distro"
	// UserResponseErrorColumn is the error column name
	UserResponseErrorColumn = "error"
	// UserResponseEventDateColumn is the event_date column name
	UserResponseEventDateColumn = "event_date"
	// UserResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	UserResponseEventDateColumnEpochColumn = "event_date->epoch"
	// UserResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	UserResponseEventDateColumnOffsetColumn = "event_date->offset"
	// UserResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	UserResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// UserResponseFreeSpaceColumn is the free_space column name
	UserResponseFreeSpaceColumn = "free_space"
	// UserResponseGoVersionColumn is the go_version column name
	UserResponseGoVersionColumn = "go_version"
	// UserResponseHostnameColumn is the hostname column name
	UserResponseHostnameColumn = "hostname"
	// UserResponseIDColumn is the id column name
	UserResponseIDColumn = "id"
	// UserResponseIntegrationIDColumn is the integration_id column name
	UserResponseIntegrationIDColumn = "integration_id"
	// UserResponseMemoryColumn is the memory column name
	UserResponseMemoryColumn = "memory"
	// UserResponseMessageColumn is the message column name
	UserResponseMessageColumn = "message"
	// UserResponseNumCPUColumn is the num_cpu column name
	UserResponseNumCPUColumn = "num_cpu"
	// UserResponseOSColumn is the os column name
	UserResponseOSColumn = "os"
	// UserResponseRefIDColumn is the ref_id column name
	UserResponseRefIDColumn = "ref_id"
	// UserResponseRefTypeColumn is the ref_type column name
	UserResponseRefTypeColumn = "ref_type"
	// UserResponseRequestIDColumn is the request_id column name
	UserResponseRequestIDColumn = "request_id"
	// UserResponseSuccessColumn is the success column name
	UserResponseSuccessColumn = "success"
	// UserResponseTypeColumn is the type column name
	UserResponseTypeColumn = "type"
	// UserResponseUsersColumn is the users column name
	UserResponseUsersColumn = "users"
	// UserResponseUsersColumnActiveColumn is the active column property of the Users name
	UserResponseUsersColumnActiveColumn = "users->active"
	// UserResponseUsersColumnAvatarURLColumn is the avatar_url column property of the Users name
	UserResponseUsersColumnAvatarURLColumn = "users->avatar_url"
	// UserResponseUsersColumnEmailsColumn is the emails column property of the Users name
	UserResponseUsersColumnEmailsColumn = "users->emails"
	// UserResponseUsersColumnGroupsColumn is the groups column property of the Users name
	UserResponseUsersColumnGroupsColumn = "users->groups"
	// UserResponseUsersColumnNameColumn is the name column property of the Users name
	UserResponseUsersColumnNameColumn = "users->name"
	// UserResponseUsersColumnUsernameColumn is the username column property of the Users name
	UserResponseUsersColumnUsernameColumn = "users->username"
	// UserResponseUUIDColumn is the uuid column name
	UserResponseUUIDColumn = "uuid"
	// UserResponseVersionColumn is the version column name
	UserResponseVersionColumn = "version"
)

// UserResponseEventDate represents the object structure for event_date
type UserResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *UserResponseEventDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// UserResponseGroups represents the object structure for groups
type UserResponseGroups struct {
	// GroupID Group id
	GroupID string `json:"group_id" bson:"group_id" yaml:"group_id" faker:"-"`
	// Name Group name
	Name string `json:"name" bson:"name" yaml:"name" faker:"-"`
}

func (o *UserResponseGroups) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// GroupID Group id
		"group_id": o.GroupID,
		// Name Group name
		"name": o.Name,
	}
}

// UserResponseType is the enumeration type for type
type UserResponseType int32

// String returns the string value for Type
func (v UserResponseType) String() string {
	switch int32(v) {
	case 0:
		return "enroll"
	case 1:
		return "ping"
	case 2:
		return "crash"
	case 3:
		return "integration"
	case 4:
		return "export"
	case 5:
		return "project"
	case 6:
		return "user"
	case 7:
		return "repo"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	UserResponseTypeEnroll UserResponseType = 0
	// TypePing is the enumeration value for ping
	UserResponseTypePing UserResponseType = 1
	// TypeCrash is the enumeration value for crash
	UserResponseTypeCrash UserResponseType = 2
	// TypeIntegration is the enumeration value for integration
	UserResponseTypeIntegration UserResponseType = 3
	// TypeExport is the enumeration value for export
	UserResponseTypeExport UserResponseType = 4
	// TypeProject is the enumeration value for project
	UserResponseTypeProject UserResponseType = 5
	// TypeUser is the enumeration value for user
	UserResponseTypeUser UserResponseType = 6
	// TypeRepo is the enumeration value for repo
	UserResponseTypeRepo UserResponseType = 7
)

// UserResponseUsers represents the object structure for users
type UserResponseUsers struct {
	// Active if user is active
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// AvatarURL the url to users avatar
	AvatarURL *string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
	// Emails the email for the user
	Emails []string `json:"emails" bson:"emails" yaml:"emails" faker:"-"`
	// Groups Group names
	Groups []UserResponseGroups `json:"groups" bson:"groups" yaml:"groups" faker:"-"`
	// Name the name of the user
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// Username the username of the user
	Username string `json:"username" bson:"username" yaml:"username" faker:"username"`
}

func (o *UserResponseUsers) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Active if user is active
		"active": o.Active,
		// AvatarURL the url to users avatar
		"avatar_url": o.AvatarURL,
		// Emails the email for the user
		"emails": o.Emails,
		// Groups Group names
		"groups": o.Groups,
		// Name the name of the user
		"name": o.Name,
		// Username the username of the user
		"username": o.Username,
	}
}

// UserResponse an agent response to an action request adding user(s)
type UserResponse struct {
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
	EventDate UserResponseEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
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
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" bson:"success" yaml:"success" faker:"-"`
	// Type the type of event
	Type UserResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// Users the exported users
	Users []UserResponseUsers `json:"users" bson:"users" yaml:"users" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*UserResponse)(nil)

func toUserResponseObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toUserResponseObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if o == nil {
		return toUserResponseObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toUserResponseObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toUserResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *UserResponse:
		return v.ToMap()
	case UserResponse:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toUserResponseObject(av, isavro, false, ""))
		}
		return arr

	case UserResponseEventDate:
		vv := o.(UserResponseEventDate)
		return vv.ToMap()
	case *UserResponseEventDate:
		return map[string]interface{}{
			"agent.event_date": (*o.(*UserResponseEventDate)).ToMap(),
		}
	case []UserResponseEventDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserResponseEventDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserResponseEventDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserResponseEventDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserResponseGroups:
		vv := o.(UserResponseGroups)
		return vv.ToMap()
	case *UserResponseGroups:
		return map[string]interface{}{
			"agent.groups": (*o.(*UserResponseGroups)).ToMap(),
		}
	case []UserResponseGroups:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserResponseGroups) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserResponseGroups:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserResponseGroups)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case UserResponseType:
		if !isavro {
			return (o.(UserResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(UserResponseType)).String(),
		}
	case *UserResponseType:
		if !isavro {
			return (o.(*UserResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(*UserResponseType)).String(),
		}
	case UserResponseUsers:
		vv := o.(UserResponseUsers)
		return vv.ToMap()
	case *UserResponseUsers:
		return map[string]interface{}{
			"agent.users": (*o.(*UserResponseUsers)).ToMap(),
		}
	case []UserResponseUsers:
		arr := make([]interface{}, 0)
		for _, i := range o.([]UserResponseUsers) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]UserResponseUsers:
		arr := make([]interface{}, 0)
		vv := o.(*[]UserResponseUsers)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
}

// String returns a string representation of UserResponse
func (o *UserResponse) String() string {
	return fmt.Sprintf("agent.UserResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *UserResponse) GetTopicName() datamodel.TopicNameType {
	return UserResponseTopic
}

// GetModelName returns the name of the model
func (o *UserResponse) GetModelName() datamodel.ModelNameType {
	return UserResponseModelName
}

func (o *UserResponse) setDefaults() {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Users == nil {
		o.Users = []UserResponseUsers{}
	}

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *UserResponse) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("UserResponse", o.CustomerID, o.RefType, o.GetRefID())
	}
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *UserResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *UserResponse) GetTimestamp() time.Time {
	var dt interface{} = o.EventDate
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
	case UserResponseEventDate:
		return datetime.DateFromEpoch(v.Epoch)
	}
	panic("not sure how to handle the date time format for UserResponse")
}

// GetRefID returns the RefID for the object
func (o *UserResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *UserResponse) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *UserResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *UserResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *UserResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = UserResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *UserResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "event_date",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *UserResponse) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *UserResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of UserResponse
func (o *UserResponse) Clone() datamodel.Model {
	c := new(UserResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *UserResponse) Anon() datamodel.Model {
	c := new(UserResponse)
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
func (o *UserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *UserResponse) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

var cachedCodecUserResponse *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *UserResponse) GetAvroCodec() *goavro.Codec {
	if cachedCodecUserResponse == nil {
		c, err := GetUserResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecUserResponse = c
	}
	return cachedCodecUserResponse
}

// ToAvroBinary returns the data as Avro binary data
func (o *UserResponse) ToAvroBinary() ([]byte, *goavro.Codec, error) {
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
func (o *UserResponse) FromAvroBinary(value []byte) error {
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
func (o *UserResponse) Stringify() string {
	return pjson.Stringify(o)
}

// IsEqual returns true if the two UserResponse objects are equal
func (o *UserResponse) IsEqual(other *UserResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *UserResponse) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Users == nil {
			o.Users = make([]UserResponseUsers, 0)
		}
	}
	o.setDefaults()
	return map[string]interface{}{
		"architecture":   toUserResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":    toUserResponseObject(o.CustomerID, isavro, false, "string"),
		"data":           toUserResponseObject(o.Data, isavro, true, "string"),
		"distro":         toUserResponseObject(o.Distro, isavro, false, "string"),
		"error":          toUserResponseObject(o.Error, isavro, true, "string"),
		"event_date":     toUserResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":     toUserResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":     toUserResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":       toUserResponseObject(o.Hostname, isavro, false, "string"),
		"id":             toUserResponseObject(o.ID, isavro, false, "string"),
		"integration_id": toUserResponseObject(o.IntegrationID, isavro, false, "string"),
		"memory":         toUserResponseObject(o.Memory, isavro, false, "long"),
		"message":        toUserResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":        toUserResponseObject(o.NumCPU, isavro, false, "long"),
		"os":             toUserResponseObject(o.OS, isavro, false, "string"),
		"ref_id":         toUserResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":       toUserResponseObject(o.RefType, isavro, false, "string"),
		"request_id":     toUserResponseObject(o.RequestID, isavro, false, "string"),
		"success":        toUserResponseObject(o.Success, isavro, false, "boolean"),
		"type":           toUserResponseObject(o.Type, isavro, false, "type"),
		"users":          toUserResponseObject(o.Users, isavro, false, "users"),
		"uuid":           toUserResponseObject(o.UUID, isavro, false, "string"),
		"version":        toUserResponseObject(o.Version, isavro, false, "string"),
		"hashcode":       toUserResponseObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *UserResponse) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		val := kv["architecture"]
		if val == nil {
			o.Architecture = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Architecture = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		val := kv["data"]
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
	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		val := kv["distro"]
		if val == nil {
			o.Distro = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Distro = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		val := kv["error"]
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
	if val, ok := kv["event_date"].(UserResponseEventDate); ok {
		o.EventDate = val
	} else {
		val := kv["event_date"]
		if val == nil {
			o.EventDate = UserResponseEventDate{}
		} else {
			o.EventDate = UserResponseEventDate{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.EventDate)

		}
	}
	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		val := kv["free_space"]
		if val == nil {
			o.FreeSpace = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.FreeSpace = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		val := kv["go_version"]
		if val == nil {
			o.GoVersion = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.GoVersion = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		val := kv["hostname"]
		if val == nil {
			o.Hostname = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Hostname = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		val := kv["integration_id"]
		if val == nil {
			o.IntegrationID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.IntegrationID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		val := kv["memory"]
		if val == nil {
			o.Memory = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Memory = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		if val == nil {
			o.Message = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Message = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		val := kv["num_cpu"]
		if val == nil {
			o.NumCPU = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.NumCPU = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		val := kv["os"]
		if val == nil {
			o.OS = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.OS = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		val := kv["request_id"]
		if val == nil {
			o.RequestID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RequestID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		val := kv["success"]
		if val == nil {
			o.Success = number.ToBoolAny(nil)
		} else {
			o.Success = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["type"].(UserResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "user":
				o.Type = 6
			case "repo":
				o.Type = 7
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "user":
				o.Type = 6
			case "repo":
				o.Type = 7
			}
		}
	}
	if val := kv["users"]; val != nil {
		na := make([]UserResponseUsers, 0)
		if a, ok := val.([]UserResponseUsers); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(UserResponseUsers); ok {
						na = append(na, av)
					} else {
						if badMap, ok := ae.(map[interface{}]interface{}); ok {
							ae = slice.ConvertToStringToInterface(badMap)
						}
						b, _ := json.Marshal(ae)
						var av UserResponseUsers
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for users field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if a, ok := val.(primitive.A); ok {
				for _, ae := range a {
					if av, ok := ae.(UserResponseUsers); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av UserResponseUsers
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for users field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for users field")
			}
		}
		o.Users = na
	} else {
		o.Users = []UserResponseUsers{}
	}
	if o.Users == nil {
		o.Users = make([]UserResponseUsers, 0)
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		val := kv["version"]
		if val == nil {
			o.Version = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Version = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
}

// Hash will return a hashcode for the object
func (o *UserResponse) Hash() string {
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
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.Type)
	args = append(args, o.Users)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetUserResponseAvroSchemaSpec creates the avro schema specification for UserResponse
func GetUserResponseAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "UserResponse",
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
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "type": "long", "name": "epoch"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the event", "type": "record", "name": "event_date"},
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
				"name": "type",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "type",
						"symbols": []interface{}{"enroll", "ping", "crash", "integration", "export", "project", "user", "repo"},
					},
				},
			},
			map[string]interface{}{
				"name": "users",
				"type": map[string]interface{}{"type": "array", "name": "users", "items": map[string]interface{}{"name": "users", "fields": []interface{}{map[string]interface{}{"doc": "if user is active", "type": "boolean", "name": "active"}, map[string]interface{}{"type": "string", "name": "avatar_url", "doc": "the url to users avatar"}, map[string]interface{}{"type": map[string]interface{}{"type": "array", "name": "emails", "items": "string"}, "name": "emails", "doc": "the email for the user"}, map[string]interface{}{"type": map[string]interface{}{"type": "array", "name": "groups", "items": map[string]interface{}{"type": "record", "name": "groups", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "group_id", "doc": "Group id"}, map[string]interface{}{"type": "string", "name": "name", "doc": "Group name"}}, "doc": "Group names"}}, "name": "groups", "doc": "Group names"}, map[string]interface{}{"doc": "the name of the user", "type": "string", "name": "name"}, map[string]interface{}{"type": "string", "name": "username", "doc": "the username of the user"}}, "doc": "the exported users", "type": "record"}},
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
func (o *UserResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetUserResponseAvroSchema creates the avro schema for UserResponse
func GetUserResponseAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetUserResponseAvroSchemaSpec())
}

// TransformUserResponseFunc is a function for transforming UserResponse during processing
type TransformUserResponseFunc func(input *UserResponse) (*UserResponse, error)

// NewUserResponsePipe creates a pipe for processing UserResponse items
func NewUserResponsePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformUserResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewUserResponseInputStream(input, errors)
	var stream chan UserResponse
	if len(transforms) > 0 {
		stream = make(chan UserResponse, 1000)
	} else {
		stream = inch
	}
	outdone := NewUserResponseOutputStream(output, stream, errors)
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

// NewUserResponseInputStreamDir creates a channel for reading UserResponse as JSON newlines from a directory of files
func NewUserResponseInputStreamDir(dir string, errors chan<- error, transforms ...TransformUserResponseFunc) (chan UserResponse, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/user_response\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan UserResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for user_response")
		ch := make(chan UserResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewUserResponseInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan UserResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewUserResponseInputStreamFile creates an channel for reading UserResponse as JSON newlines from filename
func NewUserResponseInputStreamFile(filename string, errors chan<- error, transforms ...TransformUserResponseFunc) (chan UserResponse, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan UserResponse)
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
			ch := make(chan UserResponse)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewUserResponseInputStream(f, errors, transforms...)
}

// NewUserResponseInputStream creates an channel for reading UserResponse as JSON newlines from stream
func NewUserResponseInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformUserResponseFunc) (chan UserResponse, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan UserResponse, 1000)
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
			var item UserResponse
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

// NewUserResponseOutputStreamDir will output json newlines from channel and save in dir
func NewUserResponseOutputStreamDir(dir string, ch chan UserResponse, errors chan<- error, transforms ...TransformUserResponseFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/user_response\\.json(\\.gz)?$")
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
	return NewUserResponseOutputStream(gz, ch, errors, transforms...)
}

// NewUserResponseOutputStream will output json newlines from channel to the stream
func NewUserResponseOutputStream(stream io.WriteCloser, ch chan UserResponse, errors chan<- error, transforms ...TransformUserResponseFunc) <-chan bool {
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

// UserResponseSendEvent is an event detail for sending data
type UserResponseSendEvent struct {
	UserResponse *UserResponse
	headers      map[string]string
	time         time.Time
	key          string
}

var _ datamodel.ModelSendEvent = (*UserResponseSendEvent)(nil)

// Key is the key to use for the message
func (e *UserResponseSendEvent) Key() string {
	if e.key == "" {
		return e.UserResponse.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *UserResponseSendEvent) Object() datamodel.Model {
	return e.UserResponse
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *UserResponseSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *UserResponseSendEvent) Timestamp() time.Time {
	return e.time
}

// UserResponseSendEventOpts is a function handler for setting opts
type UserResponseSendEventOpts func(o *UserResponseSendEvent)

// WithUserResponseSendEventKey sets the key value to a value different than the object ID
func WithUserResponseSendEventKey(key string) UserResponseSendEventOpts {
	return func(o *UserResponseSendEvent) {
		o.key = key
	}
}

// WithUserResponseSendEventTimestamp sets the timestamp value
func WithUserResponseSendEventTimestamp(tv time.Time) UserResponseSendEventOpts {
	return func(o *UserResponseSendEvent) {
		o.time = tv
	}
}

// WithUserResponseSendEventHeader sets the timestamp value
func WithUserResponseSendEventHeader(key, value string) UserResponseSendEventOpts {
	return func(o *UserResponseSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewUserResponseSendEvent returns a new UserResponseSendEvent instance
func NewUserResponseSendEvent(o *UserResponse, opts ...UserResponseSendEventOpts) *UserResponseSendEvent {
	res := &UserResponseSendEvent{
		UserResponse: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewUserResponseProducer will stream data from the channel
func NewUserResponseProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
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
				if object, ok := item.Object().(*UserResponse); ok {
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
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.UserResponse but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewUserResponseConsumer will stream data from the topic into the provided channel
func NewUserResponseConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object UserResponse
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.UserResponse: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.UserResponse: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.UserResponse")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &UserResponseReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object UserResponse
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &UserResponseReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// UserResponseReceiveEvent is an event detail for receiving data
type UserResponseReceiveEvent struct {
	UserResponse *UserResponse
	message      eventing.Message
	eof          bool
}

var _ datamodel.ModelReceiveEvent = (*UserResponseReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *UserResponseReceiveEvent) Object() datamodel.Model {
	return e.UserResponse
}

// Message returns the underlying message data for the event
func (e *UserResponseReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *UserResponseReceiveEvent) EOF() bool {
	return e.eof
}

// UserResponseProducer implements the datamodel.ModelEventProducer
type UserResponseProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*UserResponseProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *UserResponseProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *UserResponseProducer) Close() error {
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
func (o *UserResponse) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *UserResponse) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// NewUserResponseProducerChannel returns a channel which can be used for producing Model events
func NewUserResponseProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewUserResponseProducerChannelSize(producer, 0, errors)
}

// NewUserResponseProducerChannelSize returns a channel which can be used for producing Model events
func NewUserResponseProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &UserResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewUserResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// UserResponseConsumer implements the datamodel.ModelEventConsumer
type UserResponseConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*UserResponseConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *UserResponseConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *UserResponseConsumer) Close() error {
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
func (o *UserResponse) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserResponseConsumer{
		ch:       ch,
		callback: NewUserResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewUserResponseConsumerChannel returns a consumer channel which can be used to consume Model events
func NewUserResponseConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &UserResponseConsumer{
		ch:       ch,
		callback: NewUserResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}