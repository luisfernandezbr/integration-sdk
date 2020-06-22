// DO NOT EDIT -- generated code

// Package work - the system which contains project work
package work

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
	"github.com/pinpt/go-common/v10/slice"
	pstrings "github.com/pinpt/go-common/v10/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ConfigTopic is the default topic name
	ConfigTopic datamodel.TopicNameType = "work_Config"

	// ConfigTable is the default table name
	ConfigTable datamodel.ModelNameType = "work_config"

	// ConfigModelName is the model name
	ConfigModelName datamodel.ModelNameType = "work.Config"
)

const (
	// ConfigModelCreatedAtColumn is the column json value created_ts
	ConfigModelCreatedAtColumn = "created_ts"
	// ConfigModelCustomerIDColumn is the column json value customer_id
	ConfigModelCustomerIDColumn = "customer_id"
	// ConfigModelIDColumn is the column json value id
	ConfigModelIDColumn = "id"
	// ConfigModelIntegrationIDColumn is the column json value integration_id
	ConfigModelIntegrationIDColumn = "integration_id"
	// ConfigModelRefIDColumn is the column json value ref_id
	ConfigModelRefIDColumn = "ref_id"
	// ConfigModelRefTypeColumn is the column json value ref_type
	ConfigModelRefTypeColumn = "ref_type"
	// ConfigModelStatusesColumn is the column json value statuses
	ConfigModelStatusesColumn = "statuses"
	// ConfigModelStatusesClosedStatusColumn is the column json value closed_status
	ConfigModelStatusesClosedStatusColumn = "closed_status"
	// ConfigModelStatusesInProgressStatusColumn is the column json value in_progress_status
	ConfigModelStatusesInProgressStatusColumn = "in_progress_status"
	// ConfigModelStatusesOpenStatusColumn is the column json value open_status
	ConfigModelStatusesOpenStatusColumn = "open_status"
	// ConfigModelUpdatedAtColumn is the column json value updated_ts
	ConfigModelUpdatedAtColumn = "updated_ts"
)

// ConfigStatuses represents the object structure for statuses
type ConfigStatuses struct {
	// ClosedStatus The closed state names
	ClosedStatus []string `json:"closed_status" codec:"closed_status" bson:"closed_status" yaml:"closed_status" faker:"-"`
	// InProgressStatus The in progress state names
	InProgressStatus []string `json:"in_progress_status" codec:"in_progress_status" bson:"in_progress_status" yaml:"in_progress_status" faker:"-"`
	// OpenStatus The open status
	OpenStatus []string `json:"open_status" codec:"open_status" bson:"open_status" yaml:"open_status" faker:"-"`
}

func toConfigStatusesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ConfigStatuses:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *ConfigStatuses) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// ClosedStatus The closed state names
		"closed_status": toConfigStatusesObject(o.ClosedStatus, false),
		// InProgressStatus The in progress state names
		"in_progress_status": toConfigStatusesObject(o.InProgressStatus, false),
		// OpenStatus The open status
		"open_status": toConfigStatusesObject(o.OpenStatus, false),
	}
}

func (o *ConfigStatuses) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *ConfigStatuses) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
	o.setDefaults(false)
}

// Config Configuration for work management systems
type Config struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The ID reference to the integration instance
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Statuses The mapping of statuses for this integration
	Statuses ConfigStatuses `json:"statuses" codec:"statuses" bson:"statuses" yaml:"statuses" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Config)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Config)(nil)

func toConfigObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Config:
		return v.ToMap()

	case ConfigStatuses:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Config
func (o *Config) String() string {
	return fmt.Sprintf("work.Config<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Config) GetTopicName() datamodel.TopicNameType {
	return ConfigTopic
}

// GetStreamName returns the name of the stream
func (o *Config) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Config) GetTableName() string {
	return ConfigTable.String()
}

// GetModelName returns the name of the model
func (o *Config) GetModelName() datamodel.ModelNameType {
	return ConfigModelName
}

// NewConfigID provides a template for generating an ID field for Config
func NewConfigID(customerID string, refType string, IntegrationID string) string {
	return hash.Values(customerID, refType, IntegrationID)
}

func (o *Config) setDefaults(frommap bool) {
	if o.Statuses.ClosedStatus == nil {
		o.Statuses.ClosedStatus = make([]string, 0)
	}
	if o.Statuses.InProgressStatus == nil {
		o.Statuses.InProgressStatus = make([]string, 0)
	}
	if o.Statuses.OpenStatus == nil {
		o.Statuses.OpenStatus = make([]string, 0)
	}

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.RefType, o.IntegrationID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Config) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Config) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Config) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Config")
}

// GetRefID returns the RefID for the object
func (o *Config) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Config) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Config) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Config) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Config) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Config) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ConfigModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Config) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Config) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Config
func (o *Config) Clone() datamodel.Model {
	c := new(Config)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Config) Anon() datamodel.Model {
	c := new(Config)
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
func (o *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Config) UnmarshalJSON(data []byte) error {
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
func (o *Config) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *Config) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two Config objects are equal
func (o *Config) IsEqual(other *Config) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Config) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":     toConfigObject(o.CreatedAt, false),
		"customer_id":    toConfigObject(o.CustomerID, false),
		"id":             toConfigObject(o.ID, false),
		"integration_id": toConfigObject(o.IntegrationID, false),
		"ref_id":         toConfigObject(o.RefID, false),
		"ref_type":       toConfigObject(o.RefType, false),
		"statuses":       toConfigObject(o.Statuses, false),
		"updated_ts":     toConfigObject(o.UpdatedAt, false),
		"hashcode":       toConfigObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Config) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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
		} else if sv, ok := val.(ConfigStatuses); ok {
			// struct
			o.Statuses = sv
		} else if sp, ok := val.(*ConfigStatuses); ok {
			// struct pointer
			o.Statuses = *sp
		}
	} else {
		o.Statuses.FromMap(map[string]interface{}{})
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
func (o *Config) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Statuses)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

func getConfigQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// object with fields
	sb.WriteString("\t\t\tstatuses {\n")

	// scalar
	sb.WriteString("\t\t\tclosed_status\n")
	// scalar
	sb.WriteString("\t\t\tin_progress_status\n")
	// scalar
	sb.WriteString("\t\t\topen_status\n")
	sb.WriteString("\t\t\t}\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	return sb.String()
}

// ConfigPageInfo is a grapqhl PageInfo
type ConfigPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// ConfigConnection is a grapqhl connection
type ConfigConnection struct {
	Edges      []*ConfigEdge  `json:"edges,omitempty"`
	PageInfo   ConfigPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64         `json:"totalCount,omitempty"`
}

// ConfigEdge is a grapqhl edge
type ConfigEdge struct {
	Node *Config `json:"node,omitempty"`
}

// QueryManyConfigNode is a grapqhl query many node
type QueryManyConfigNode struct {
	Object *ConfigConnection `json:"Configs,omitempty"`
}

// QueryManyConfigData is a grapqhl query many data node
type QueryManyConfigData struct {
	Data *QueryManyConfigNode `json:"work,omitempty"`
}

// QueryOneConfigNode is a grapqhl query one node
type QueryOneConfigNode struct {
	Object *Config `json:"Config,omitempty"`
}

// QueryOneConfigData is a grapqhl query one data node
type QueryOneConfigData struct {
	Data *QueryOneConfigNode `json:"work,omitempty"`
}

// ConfigQuery is query struct
type ConfigQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// ConfigOrder is the order direction
type ConfigOrder *string

var (
	// ConfigOrderDesc is descending
	ConfigOrderDesc ConfigOrder = pstrings.Pointer("DESC")
	// ConfigOrderAsc is ascending
	ConfigOrderAsc ConfigOrder = pstrings.Pointer("ASC")
)

// ConfigQueryInput is query input struct
type ConfigQueryInput struct {
	First   *int64       `json:"first,omitempty"`
	Last    *int64       `json:"last,omitempty"`
	Before  *string      `json:"before,omitempty"`
	After   *string      `json:"after,omitempty"`
	Query   *ConfigQuery `json:"query,omitempty"`
	OrderBy *string      `json:"orderBy,omitempty"`
	Order   ConfigOrder  `json:"order,omitempty"`
}

// NewConfigQuery is a convenience for building a *ConfigQuery
func NewConfigQuery(params ...interface{}) *ConfigQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &ConfigQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &ConfigQueryInput{
		Query: q,
	}
}

// FindConfig will query an Config by id
func FindConfig(client graphql.Client, id string) (*Config, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query ConfigQuery($id: ID) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tConfig(_id: $id) {\n")
	sb.WriteString(getConfigQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneConfigData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindConfigs will query for any Configs matching the query
func FindConfigs(client graphql.Client, input *ConfigQueryInput) (*ConfigConnection, error) {
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
	sb.WriteString("query ConfigQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: WorkConfigColumnEnum) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tConfigs(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getConfigQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyConfigData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindConfigsPaginatedCallback is a callback function for handling each page
type FindConfigsPaginatedCallback func(conn *ConfigConnection) (bool, error)

// FindConfigsPaginated will query for any Configs matching the query and return each page callback
func FindConfigsPaginated(client graphql.Client, query *ConfigQuery, pageSize int64, callback FindConfigsPaginatedCallback) error {
	input := &ConfigQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindConfigs(client, input)
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

// UpdateConfigNode is a grapqhl update node
type UpdateConfigNode struct {
	Object *Config `json:"updateConfig,omitempty"`
}

// UpdateConfigData is a grapqhl update data node
type UpdateConfigData struct {
	Data *UpdateConfigNode `json:"work,omitempty"`
}

// ExecConfigUpdateMutation returns a graphql update mutation result for Config
func ExecConfigUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*Config, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation ConfigUpdateMutation($id: String, $input: UpdateWorkConfigInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateConfig(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getConfigQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateConfigData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecConfigSilentUpdateMutation returns a graphql update mutation result for Config
func ExecConfigSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation ConfigUpdateMutation($id: String, $input: UpdateWorkConfigInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateConfig(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateConfigData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateConfig(client graphql.Client, model Config) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation CreateConfig($input: CreateWorkConfigInput!) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tcreateConfig(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateConfigData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}
