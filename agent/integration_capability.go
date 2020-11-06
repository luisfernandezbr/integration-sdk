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
)

const (

	// IntegrationCapabilityTable is the default table name
	IntegrationCapabilityTable datamodel.ModelNameType = "agent_integrationcapability"

	// IntegrationCapabilityModelName is the model name
	IntegrationCapabilityModelName datamodel.ModelNameType = "agent.IntegrationCapability"
)

const (
	// IntegrationCapabilityModelCreatedAtColumn is the column json value created_ts
	IntegrationCapabilityModelCreatedAtColumn = "created_ts"
	// IntegrationCapabilityModelCustomerIDColumn is the column json value customer_id
	IntegrationCapabilityModelCustomerIDColumn = "customer_id"
	// IntegrationCapabilityModelIDColumn is the column json value id
	IntegrationCapabilityModelIDColumn = "id"
	// IntegrationCapabilityModelIntegrationIDColumn is the column json value integration_id
	IntegrationCapabilityModelIntegrationIDColumn = "integration_id"
	// IntegrationCapabilityModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	IntegrationCapabilityModelIntegrationInstanceIDColumn = "integration_instance_id"
	// IntegrationCapabilityModelRefIDColumn is the column json value ref_id
	IntegrationCapabilityModelRefIDColumn = "ref_id"
	// IntegrationCapabilityModelRefTypeColumn is the column json value ref_type
	IntegrationCapabilityModelRefTypeColumn = "ref_type"
	// IntegrationCapabilityModelUpdatedAtColumn is the column json value updated_ts
	IntegrationCapabilityModelUpdatedAtColumn = "updated_ts"
	// IntegrationCapabilityModelUserLinkingScopeColumn is the column json value user_linking_scope
	IntegrationCapabilityModelUserLinkingScopeColumn = "user_linking_scope"
)

// IntegrationCapabilityUserLinkingScope is the enumeration type for user_linking_scope
type IntegrationCapabilityUserLinkingScope int32

// toIntegrationCapabilityUserLinkingScopePointer is the enumeration pointer type for user_linking_scope
func toIntegrationCapabilityUserLinkingScopePointer(v int32) *IntegrationCapabilityUserLinkingScope {
	nv := IntegrationCapabilityUserLinkingScope(v)
	return &nv
}

// toIntegrationCapabilityUserLinkingScopeEnum is the enumeration pointer wrapper for user_linking_scope
func toIntegrationCapabilityUserLinkingScopeEnum(v *IntegrationCapabilityUserLinkingScope) string {
	if v == nil {
		return toIntegrationCapabilityUserLinkingScopePointer(0).String()
	}
	return v.String()
}

// UnmarshalBSONValue for unmarshaling value
func (v *IntegrationCapabilityUserLinkingScope) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.RawValue{Type: t, Value: data}
	switch t {
	case bsontype.Int32:
		*v = IntegrationCapabilityUserLinkingScope(val.Int32())
	case bsontype.String:
		switch val.StringValue() {
		case "INTEGRATION_INSTANCE":
			*v = IntegrationCapabilityUserLinkingScope(0)
		case "GLOBAL":
			*v = IntegrationCapabilityUserLinkingScope(1)
		}
	}
	return nil
}

// UnmarshalJSON unmarshals the enum value
func (v *IntegrationCapabilityUserLinkingScope) UnmarshalJSON(buf []byte) error {
	var val string
	if err := json.Unmarshal(buf, &val); err != nil {
		return err
	}
	switch val {
	case "INTEGRATION_INSTANCE":
		*v = 0
	case "GLOBAL":
		*v = 1
	}
	return nil
}

// MarshalJSON marshals the enum value
func (v IntegrationCapabilityUserLinkingScope) MarshalJSON() ([]byte, error) {
	switch v {
	case 0:
		return json.Marshal("INTEGRATION_INSTANCE")
	case 1:
		return json.Marshal("GLOBAL")
	}
	return nil, fmt.Errorf("unexpected enum value")
}

// String returns the string value for UserLinkingScope
func (v IntegrationCapabilityUserLinkingScope) String() string {
	switch int32(v) {
	case 0:
		return "INTEGRATION_INSTANCE"
	case 1:
		return "GLOBAL"
	}
	return "unset"
}

// FromInterface for decoding from an interface
func (v *IntegrationCapabilityUserLinkingScope) FromInterface(o interface{}) error {
	switch val := o.(type) {
	case IntegrationCapabilityUserLinkingScope:
		*v = val
	case int32:
		*v = IntegrationCapabilityUserLinkingScope(int32(val))
	case int:
		*v = IntegrationCapabilityUserLinkingScope(int32(val))
	case string:
		switch val {
		case "INTEGRATION_INSTANCE":
			*v = IntegrationCapabilityUserLinkingScope(0)
		case "GLOBAL":
			*v = IntegrationCapabilityUserLinkingScope(1)
		}
	}
	return nil
}

const (
	// IntegrationCapabilityUserLinkingScopeIntegrationInstance is the enumeration value for integration_instance
	IntegrationCapabilityUserLinkingScopeIntegrationInstance IntegrationCapabilityUserLinkingScope = 0
	// IntegrationCapabilityUserLinkingScopeGlobal is the enumeration value for global
	IntegrationCapabilityUserLinkingScopeGlobal IntegrationCapabilityUserLinkingScope = 1
)

// IntegrationCapability Data about an integration instance that is useful but too common to be sent in a dbchange.
type IntegrationCapability struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID The unique id for the integration
	IntegrationID string `json:"integration_id" codec:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// UserLinkingScope how linked users are scoped, (ie. globally, or by customer/integration instance)
	UserLinkingScope IntegrationCapabilityUserLinkingScope `json:"user_linking_scope" codec:"user_linking_scope" bson:"user_linking_scope" yaml:"user_linking_scope" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*IntegrationCapability)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*IntegrationCapability)(nil)

func toIntegrationCapabilityObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *IntegrationCapability:
		return v.ToMap()

	case IntegrationCapabilityUserLinkingScope:
		return v.String()
	default:
		return o
	}
}

// String returns a string representation of IntegrationCapability
func (o *IntegrationCapability) String() string {
	return fmt.Sprintf("agent.IntegrationCapability<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *IntegrationCapability) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *IntegrationCapability) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *IntegrationCapability) GetTableName() string {
	return IntegrationCapabilityTable.String()
}

// GetModelName returns the name of the model
func (o *IntegrationCapability) GetModelName() datamodel.ModelNameType {
	return IntegrationCapabilityModelName
}

// NewIntegrationCapabilityID provides a template for generating an ID field for IntegrationCapability
func NewIntegrationCapabilityID(IntegrationID string) string {
	return hash.Values(IntegrationID)
}

func (o *IntegrationCapability) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.IntegrationID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *IntegrationCapability) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *IntegrationCapability) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *IntegrationCapability) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *IntegrationCapability) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *IntegrationCapability) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *IntegrationCapability) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *IntegrationCapability) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *IntegrationCapability) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *IntegrationCapability) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = IntegrationCapabilityModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *IntegrationCapability) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *IntegrationCapability) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *IntegrationCapability) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *IntegrationCapability) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *IntegrationCapability) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of IntegrationCapability
func (o *IntegrationCapability) Clone() datamodel.Model {
	c := new(IntegrationCapability)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *IntegrationCapability) Anon() datamodel.Model {
	c := new(IntegrationCapability)
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
func (o *IntegrationCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationCapability) UnmarshalJSON(data []byte) error {
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
func (o *IntegrationCapability) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *IntegrationCapability) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two IntegrationCapability objects are equal
func (o *IntegrationCapability) IsEqual(other *IntegrationCapability) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *IntegrationCapability) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toIntegrationCapabilityObject(o.CreatedAt, false),
		"customer_id":             toIntegrationCapabilityObject(o.CustomerID, false),
		"id":                      toIntegrationCapabilityObject(o.ID, false),
		"integration_id":          toIntegrationCapabilityObject(o.IntegrationID, false),
		"integration_instance_id": toIntegrationCapabilityObject(o.IntegrationInstanceID, true),
		"ref_id":                  toIntegrationCapabilityObject(o.RefID, false),
		"ref_type":                toIntegrationCapabilityObject(o.RefType, false),
		"updated_ts":              toIntegrationCapabilityObject(o.UpdatedAt, false),

		"user_linking_scope": o.UserLinkingScope.String(),
		"hashcode":           toIntegrationCapabilityObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *IntegrationCapability) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["user_linking_scope"].(IntegrationCapabilityUserLinkingScope); ok {
		o.UserLinkingScope = val
	} else {
		if em, ok := kv["user_linking_scope"].(map[string]interface{}); ok {

			ev := em["agent.user_linking_scope"].(string)
			switch ev {
			case "integration_instance", "INTEGRATION_INSTANCE":
				o.UserLinkingScope = 0
			case "global", "GLOBAL":
				o.UserLinkingScope = 1
			}
		}
		if em, ok := kv["user_linking_scope"].(string); ok {
			switch em {
			case "integration_instance", "INTEGRATION_INSTANCE":
				o.UserLinkingScope = 0
			case "global", "GLOBAL":
				o.UserLinkingScope = 1
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *IntegrationCapability) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.UserLinkingScope)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *IntegrationCapability) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *IntegrationCapability) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *IntegrationCapability) GetHydrationQuery() string {
	return `query GoIntegrationCapabilityQuery($id: ID, $nocache: Boolean) {
	agent {
		IntegrationCapability(_id: $id, nocache: $nocache) {
			created_ts
			customer_id
			_id
			integration_id
			integration_instance_id
			ref_id
			ref_type
			updated_ts
			user_linking_scope
		}
	}
}
`
}

func getIntegrationCapabilityQueryFields() string {
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
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// scalar
	sb.WriteString("\t\t\tuser_linking_scope\n")
	return sb.String()
}

// IntegrationCapabilityPageInfo is a grapqhl PageInfo
type IntegrationCapabilityPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// IntegrationCapabilityCacheInfo is the grapqhl CacheInfo
type IntegrationCapabilityCacheInfo struct {
	Cached bool    `json:"cached,omitempty"`
	ID     *string `json:"id,omitempty"`
	ETag   *string `json:"etag,omitempty"`
}

// IntegrationCapabilityConnection is a grapqhl connection
type IntegrationCapabilityConnection struct {
	Edges      []*IntegrationCapabilityEdge   `json:"edges,omitempty"`
	PageInfo   IntegrationCapabilityPageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  IntegrationCapabilityCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64                         `json:"totalCount,omitempty"`
}

// IntegrationCapabilityEdge is a grapqhl edge
type IntegrationCapabilityEdge struct {
	Node *IntegrationCapability `json:"node,omitempty"`
}

// QueryManyIntegrationCapabilityNode is a grapqhl query many node
type QueryManyIntegrationCapabilityNode struct {
	Object *IntegrationCapabilityConnection `json:"IntegrationCapabilities,omitempty"`
}

// QueryManyIntegrationCapabilityData is a grapqhl query many data node
type QueryManyIntegrationCapabilityData struct {
	Data *QueryManyIntegrationCapabilityNode `json:"agent,omitempty"`
}

// QueryOneIntegrationCapabilityNode is a grapqhl query one node
type QueryOneIntegrationCapabilityNode struct {
	Object *IntegrationCapability `json:"IntegrationCapability,omitempty"`
}

// QueryOneIntegrationCapabilityData is a grapqhl query one data node
type QueryOneIntegrationCapabilityData struct {
	Data *QueryOneIntegrationCapabilityNode `json:"agent,omitempty"`
}

// IntegrationCapabilityQuery is query struct
type IntegrationCapabilityQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// IntegrationCapabilityOrder is the order direction
type IntegrationCapabilityOrder *string

var (
	// IntegrationCapabilityOrderDesc is descending
	IntegrationCapabilityOrderDesc IntegrationCapabilityOrder = pstrings.Pointer("DESC")
	// IntegrationCapabilityOrderAsc is ascending
	IntegrationCapabilityOrderAsc IntegrationCapabilityOrder = pstrings.Pointer("ASC")
)

// IntegrationCapabilityQueryInput is query input struct
type IntegrationCapabilityQueryInput struct {
	First   *int64                      `json:"first,omitempty"`
	Last    *int64                      `json:"last,omitempty"`
	Before  *string                     `json:"before,omitempty"`
	After   *string                     `json:"after,omitempty"`
	Query   *IntegrationCapabilityQuery `json:"query,omitempty"`
	OrderBy *string                     `json:"orderBy,omitempty"`
	Order   IntegrationCapabilityOrder  `json:"order,omitempty"`
	NoCache bool                        `json:"nocache,omitempty"`
}

// NewIntegrationCapabilityQuery is a convenience for building a *IntegrationCapabilityQuery
func NewIntegrationCapabilityQuery(params ...interface{}) *IntegrationCapabilityQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &IntegrationCapabilityQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &IntegrationCapabilityQueryInput{
		Query: q,
	}
}

// FindIntegrationCapability will query an IntegrationCapability by id
func FindIntegrationCapability(client graphql.Client, id string) (*IntegrationCapability, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoIntegrationCapabilityQuery($id: ID) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationCapability(_id: $id) {\n")
	sb.WriteString(getIntegrationCapabilityQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationCapabilityData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationCapabilityWithoutCache will query an IntegrationCapability by id avoiding the cache
func FindIntegrationCapabilityWithoutCache(client graphql.Client, id string) (*IntegrationCapability, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoIntegrationCapabilityQuery($id: ID, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationCapability(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getIntegrationCapabilityQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneIntegrationCapabilityData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindIntegrationCapabilities will query for any IntegrationCapabilities matching the query
func FindIntegrationCapabilities(client graphql.Client, input *IntegrationCapabilityQueryInput) (*IntegrationCapabilityConnection, error) {
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
		variables["nocache"] = input.NoCache
	}
	var sb strings.Builder
	sb.WriteString("query GoIntegrationCapabilityQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: AgentIntegrationCapabilityColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tIntegrationCapabilities(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\tcacheInfo {\n")
	sb.WriteString("\t\t\t\tcached\n")
	sb.WriteString("\t\t\t\tid\n")
	sb.WriteString("\t\t\t\tetag\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getIntegrationCapabilityQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyIntegrationCapabilityData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindIntegrationCapabilitiesPaginatedCallback is a callback function for handling each page
type FindIntegrationCapabilitiesPaginatedCallback func(conn *IntegrationCapabilityConnection) (bool, error)

// FindIntegrationCapabilitiesPaginated will query for any IntegrationCapabilities matching the query and return each page callback
func FindIntegrationCapabilitiesPaginated(client graphql.Client, query *IntegrationCapabilityQuery, pageSize int64, callback FindIntegrationCapabilitiesPaginatedCallback) error {
	input := &IntegrationCapabilityQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindIntegrationCapabilities(client, input)
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

// FindIntegrationCapabilitiesPaginatedWithoutCache will query for any IntegrationCapabilities matching the query and return each page callback
func FindIntegrationCapabilitiesPaginatedWithoutCache(client graphql.Client, query *IntegrationCapabilityQuery, pageSize int64, callback FindIntegrationCapabilitiesPaginatedCallback) error {
	input := &IntegrationCapabilityQueryInput{
		First:   &pageSize,
		Query:   query,
		NoCache: true,
	}
	for {
		res, err := FindIntegrationCapabilities(client, input)
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

// UpdateIntegrationCapabilityNode is a grapqhl update node
type UpdateIntegrationCapabilityNode struct {
	Object *IntegrationCapability `json:"updateIntegrationCapability,omitempty"`
}

// UpdateIntegrationCapabilityData is a grapqhl update data node
type UpdateIntegrationCapabilityData struct {
	Data *UpdateIntegrationCapabilityNode `json:"agent,omitempty"`
}

// ExecIntegrationCapabilityUpdateMutation returns a graphql update mutation result for IntegrationCapability
func ExecIntegrationCapabilityUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*IntegrationCapability, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationCapabilityUpdateMutation($id: String, $input: UpdateAgentIntegrationCapabilityInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationCapability(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getIntegrationCapabilityQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationCapabilityData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecIntegrationCapabilitySilentUpdateMutation returns a graphql update mutation result for IntegrationCapability
func ExecIntegrationCapabilitySilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationCapabilityUpdateMutation($id: String, $input: UpdateAgentIntegrationCapabilityInput, $upsert: Boolean) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tupdateIntegrationCapability(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationCapabilityData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecIntegrationCapabilityDeleteMutation executes a graphql delete mutation for IntegrationCapability
func ExecIntegrationCapabilityDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoIntegrationCapabilityDeleteMutation($id: String!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tdeleteIntegrationCapability(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateIntegrationCapability(client graphql.Client, model IntegrationCapability) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateIntegrationCapability($input: CreateAgentIntegrationCapabilityInput!) {\n")
	sb.WriteString("\tagent {\n")
	sb.WriteString("\t\tcreateIntegrationCapability(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateIntegrationCapabilityData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// IntegrationCapabilityPartial is a partial struct for upsert mutations for IntegrationCapability
type IntegrationCapabilityPartial struct {
	// IntegrationID The unique id for the integration
	IntegrationID *string `json:"integration_id,omitempty"`
	// UserLinkingScope how linked users are scoped, (ie. globally, or by customer/integration instance)
	UserLinkingScope *IntegrationCapabilityUserLinkingScope `json:"user_linking_scope,omitempty"`
}

var _ datamodel.PartialModel = (*IntegrationCapabilityPartial)(nil)

// GetModelName returns the name of the model
func (o *IntegrationCapabilityPartial) GetModelName() datamodel.ModelNameType {
	return IntegrationCapabilityModelName
}

// ToMap returns the object as a map
func (o *IntegrationCapabilityPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"integration_id": toIntegrationCapabilityObject(o.IntegrationID, true),

		"user_linking_scope": toIntegrationCapabilityUserLinkingScopeEnum(o.UserLinkingScope),
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
func (o *IntegrationCapabilityPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *IntegrationCapabilityPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *IntegrationCapabilityPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *IntegrationCapabilityPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *IntegrationCapabilityPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["integration_id"].(*string); ok {
		o.IntegrationID = val
	} else if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = &val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.IntegrationID = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}
	if val, ok := kv["user_linking_scope"].(*IntegrationCapabilityUserLinkingScope); ok {
		o.UserLinkingScope = val
	} else if val, ok := kv["user_linking_scope"].(IntegrationCapabilityUserLinkingScope); ok {
		o.UserLinkingScope = &val
	} else {
		if val, ok := kv["user_linking_scope"]; ok {
			if val == nil {
				o.UserLinkingScope = toIntegrationCapabilityUserLinkingScopePointer(0)
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["IntegrationCapabilityUserLinkingScope"]
				}
				// this is an enum pointer
				if em, ok := val.(string); ok {
					switch em {
					case "integration_instance", "INTEGRATION_INSTANCE":
						o.UserLinkingScope = toIntegrationCapabilityUserLinkingScopePointer(0)
					case "global", "GLOBAL":
						o.UserLinkingScope = toIntegrationCapabilityUserLinkingScopePointer(1)
					}
				}
			}
		}
	}
	o.setDefaults(false)
}
