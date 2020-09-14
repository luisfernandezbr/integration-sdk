// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

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
)

const (
	// RepoWebhookTopic is the default topic name
	RepoWebhookTopic datamodel.TopicNameType = "sourcecode_RepoWebhook"

	// RepoWebhookTable is the default table name
	RepoWebhookTable datamodel.ModelNameType = "sourcecode_repowebhook"

	// RepoWebhookModelName is the model name
	RepoWebhookModelName datamodel.ModelNameType = "sourcecode.RepoWebhook"
)

const (
	// RepoWebhookModelCreatedAtColumn is the column json value created_ts
	RepoWebhookModelCreatedAtColumn = "created_ts"
	// RepoWebhookModelCustomerIDColumn is the column json value customer_id
	RepoWebhookModelCustomerIDColumn = "customer_id"
	// RepoWebhookModelEnabledColumn is the column json value enabled
	RepoWebhookModelEnabledColumn = "enabled"
	// RepoWebhookModelErrorMessageColumn is the column json value error_message
	RepoWebhookModelErrorMessageColumn = "error_message"
	// RepoWebhookModelErroredColumn is the column json value errored
	RepoWebhookModelErroredColumn = "errored"
	// RepoWebhookModelIDColumn is the column json value id
	RepoWebhookModelIDColumn = "id"
	// RepoWebhookModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	RepoWebhookModelIntegrationInstanceIDColumn = "integration_instance_id"
	// RepoWebhookModelRefIDColumn is the column json value ref_id
	RepoWebhookModelRefIDColumn = "ref_id"
	// RepoWebhookModelRefTypeColumn is the column json value ref_type
	RepoWebhookModelRefTypeColumn = "ref_type"
	// RepoWebhookModelRepoIDColumn is the column json value repo_id
	RepoWebhookModelRepoIDColumn = "repo_id"
	// RepoWebhookModelUpdatedAtColumn is the column json value updated_ts
	RepoWebhookModelUpdatedAtColumn = "updated_ts"
	// RepoWebhookModelURLColumn is the column json value url
	RepoWebhookModelURLColumn = "url"
)

// RepoWebhook the repo webhook contains details about the webhook if one is installed at the repo level
type RepoWebhook struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Enabled if webhooks are enabled for this repo
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the webhook has an error
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the id of the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoWebhook)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*RepoWebhook)(nil)

func toRepoWebhookObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoWebhook:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of RepoWebhook
func (o *RepoWebhook) String() string {
	return fmt.Sprintf("sourcecode.RepoWebhook<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoWebhook) GetTopicName() datamodel.TopicNameType {
	return RepoWebhookTopic
}

// GetStreamName returns the name of the stream
func (o *RepoWebhook) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *RepoWebhook) GetTableName() string {
	return RepoWebhookTable.String()
}

// GetModelName returns the name of the model
func (o *RepoWebhook) GetModelName() datamodel.ModelNameType {
	return RepoWebhookModelName
}

// NewRepoWebhookID provides a template for generating an ID field for RepoWebhook
func NewRepoWebhookID(customerID string, RepoID string) string {
	return hash.Values(customerID, RepoID)
}

func (o *RepoWebhook) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.RepoID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoWebhook) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoWebhook) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoWebhook) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for RepoWebhook")
}

// GetRefID returns the RefID for the object
func (o *RepoWebhook) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoWebhook) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *RepoWebhook) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoWebhook) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoWebhook) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoWebhook) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoWebhookModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoWebhook) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *RepoWebhook) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *RepoWebhook) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *RepoWebhook) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *RepoWebhook) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of RepoWebhook
func (o *RepoWebhook) Clone() datamodel.Model {
	c := new(RepoWebhook)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RepoWebhook) Anon() datamodel.Model {
	c := new(RepoWebhook)
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
func (o *RepoWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoWebhook) UnmarshalJSON(data []byte) error {
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
func (o *RepoWebhook) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *RepoWebhook) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two RepoWebhook objects are equal
func (o *RepoWebhook) IsEqual(other *RepoWebhook) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoWebhook) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toRepoWebhookObject(o.CreatedAt, false),
		"customer_id":             toRepoWebhookObject(o.CustomerID, false),
		"enabled":                 toRepoWebhookObject(o.Enabled, false),
		"error_message":           toRepoWebhookObject(o.ErrorMessage, true),
		"errored":                 toRepoWebhookObject(o.Errored, false),
		"id":                      toRepoWebhookObject(o.ID, false),
		"integration_instance_id": toRepoWebhookObject(o.IntegrationInstanceID, true),
		"ref_id":                  toRepoWebhookObject(o.RefID, false),
		"ref_type":                toRepoWebhookObject(o.RefType, false),
		"repo_id":                 toRepoWebhookObject(o.RepoID, false),
		"updated_ts":              toRepoWebhookObject(o.UpdatedAt, false),
		"url":                     toRepoWebhookObject(o.URL, true),
		"hashcode":                toRepoWebhookObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoWebhook) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.RepoID = fmt.Sprintf("%v", val)
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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *RepoWebhook) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Enabled)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *RepoWebhook) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *RepoWebhook) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *RepoWebhook) GetHydrationQuery() string {
	return `query GoRepoWebhookQuery($id: ID) {
	sourcecode {
		RepoWebhook(_id: $id) {
			created_ts
			customer_id
			enabled
			error_message
			errored
			_id
			integration_instance_id
			ref_id
			ref_type
			repo_id
			updated_ts
			url
		}
	}
}
`
}

func getRepoWebhookQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
	// scalar
	sb.WriteString("\t\t\tenabled\n")
	// scalar
	sb.WriteString("\t\t\terror_message\n")
	// scalar
	sb.WriteString("\t\t\terrored\n")
	// id
	sb.WriteString("\t\t\t_id\n")
	// scalar
	sb.WriteString("\t\t\tintegration_instance_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\trepo_id\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// scalar
	sb.WriteString("\t\t\turl\n")
	return sb.String()
}

// RepoWebhookPageInfo is a grapqhl PageInfo
type RepoWebhookPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// RepoWebhookConnection is a grapqhl connection
type RepoWebhookConnection struct {
	Edges      []*RepoWebhookEdge  `json:"edges,omitempty"`
	PageInfo   RepoWebhookPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64              `json:"totalCount,omitempty"`
}

// RepoWebhookEdge is a grapqhl edge
type RepoWebhookEdge struct {
	Node *RepoWebhook `json:"node,omitempty"`
}

// QueryManyRepoWebhookNode is a grapqhl query many node
type QueryManyRepoWebhookNode struct {
	Object *RepoWebhookConnection `json:"RepoWebhooks,omitempty"`
}

// QueryManyRepoWebhookData is a grapqhl query many data node
type QueryManyRepoWebhookData struct {
	Data *QueryManyRepoWebhookNode `json:"sourcecode,omitempty"`
}

// QueryOneRepoWebhookNode is a grapqhl query one node
type QueryOneRepoWebhookNode struct {
	Object *RepoWebhook `json:"RepoWebhook,omitempty"`
}

// QueryOneRepoWebhookData is a grapqhl query one data node
type QueryOneRepoWebhookData struct {
	Data *QueryOneRepoWebhookNode `json:"sourcecode,omitempty"`
}

// RepoWebhookQuery is query struct
type RepoWebhookQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// RepoWebhookOrder is the order direction
type RepoWebhookOrder *string

var (
	// RepoWebhookOrderDesc is descending
	RepoWebhookOrderDesc RepoWebhookOrder = pstrings.Pointer("DESC")
	// RepoWebhookOrderAsc is ascending
	RepoWebhookOrderAsc RepoWebhookOrder = pstrings.Pointer("ASC")
)

// RepoWebhookQueryInput is query input struct
type RepoWebhookQueryInput struct {
	First   *int64            `json:"first,omitempty"`
	Last    *int64            `json:"last,omitempty"`
	Before  *string           `json:"before,omitempty"`
	After   *string           `json:"after,omitempty"`
	Query   *RepoWebhookQuery `json:"query,omitempty"`
	OrderBy *string           `json:"orderBy,omitempty"`
	Order   RepoWebhookOrder  `json:"order,omitempty"`
}

// NewRepoWebhookQuery is a convenience for building a *RepoWebhookQuery
func NewRepoWebhookQuery(params ...interface{}) *RepoWebhookQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &RepoWebhookQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &RepoWebhookQueryInput{
		Query: q,
	}
}

// FindRepoWebhook will query an RepoWebhook by id
func FindRepoWebhook(client graphql.Client, id string) (*RepoWebhook, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoRepoWebhookQuery($id: ID) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tRepoWebhook(_id: $id) {\n")
	sb.WriteString(getRepoWebhookQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneRepoWebhookData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindRepoWebhooks will query for any RepoWebhooks matching the query
func FindRepoWebhooks(client graphql.Client, input *RepoWebhookQueryInput) (*RepoWebhookConnection, error) {
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
	sb.WriteString("query GoRepoWebhookQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: SourcecodeRepoWebhookColumnEnum) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tRepoWebhooks(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getRepoWebhookQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyRepoWebhookData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindRepoWebhooksPaginatedCallback is a callback function for handling each page
type FindRepoWebhooksPaginatedCallback func(conn *RepoWebhookConnection) (bool, error)

// FindRepoWebhooksPaginated will query for any RepoWebhooks matching the query and return each page callback
func FindRepoWebhooksPaginated(client graphql.Client, query *RepoWebhookQuery, pageSize int64, callback FindRepoWebhooksPaginatedCallback) error {
	input := &RepoWebhookQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindRepoWebhooks(client, input)
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

// UpdateRepoWebhookNode is a grapqhl update node
type UpdateRepoWebhookNode struct {
	Object *RepoWebhook `json:"updateRepoWebhook,omitempty"`
}

// UpdateRepoWebhookData is a grapqhl update data node
type UpdateRepoWebhookData struct {
	Data *UpdateRepoWebhookNode `json:"sourcecode,omitempty"`
}

// ExecRepoWebhookUpdateMutation returns a graphql update mutation result for RepoWebhook
func ExecRepoWebhookUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*RepoWebhook, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoRepoWebhookUpdateMutation($id: String, $input: UpdateSourcecodeRepoWebhookInput, $upsert: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tupdateRepoWebhook(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getRepoWebhookQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecRepoWebhookSilentUpdateMutation returns a graphql update mutation result for RepoWebhook
func ExecRepoWebhookSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoRepoWebhookUpdateMutation($id: String, $input: UpdateSourcecodeRepoWebhookInput, $upsert: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tupdateRepoWebhook(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecRepoWebhookDeleteMutation executes a graphql delete mutation for RepoWebhook
func ExecRepoWebhookDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoRepoWebhookDeleteMutation($id: String!) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tdeleteRepoWebhook(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateRepoWebhook(client graphql.Client, model RepoWebhook) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateRepoWebhook($input: CreateSourcecodeRepoWebhookInput!) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tcreateRepoWebhook(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// RepoWebhookPartial is a partial struct for upsert mutations for RepoWebhook
type RepoWebhookPartial struct {
	// Enabled if webhooks are enabled for this repo
	Enabled *bool `json:"enabled,omitempty"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored if the webhook has an error
	Errored *bool `json:"errored,omitempty"`
	// RepoID the id of the repo
	RepoID *string `json:"repo_id,omitempty"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*RepoWebhookPartial)(nil)

// GetModelName returns the name of the model
func (o *RepoWebhookPartial) GetModelName() datamodel.ModelNameType {
	return RepoWebhookModelName
}

// ToMap returns the object as a map
func (o *RepoWebhookPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"enabled":       toRepoWebhookObject(o.Enabled, true),
		"error_message": toRepoWebhookObject(o.ErrorMessage, true),
		"errored":       toRepoWebhookObject(o.Errored, true),
		"repo_id":       toRepoWebhookObject(o.RepoID, true),
		"url":           toRepoWebhookObject(o.URL, true),
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
func (o *RepoWebhookPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoWebhookPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoWebhookPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *RepoWebhookPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *RepoWebhookPartial) FromMap(kv map[string]interface{}) {
	if val, ok := kv["enabled"].(*bool); ok {
		o.Enabled = val
	} else if val, ok := kv["enabled"].(bool); ok {
		o.Enabled = &val
	} else {
		if val, ok := kv["enabled"]; ok {
			if val == nil {
				o.Enabled = nil
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["bool"]
				}
				o.Enabled = number.BoolPointer(number.ToBoolAny(val))
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
	if val, ok := kv["repo_id"].(*string); ok {
		o.RepoID = val
	} else if val, ok := kv["repo_id"].(string); ok {
		o.RepoID = &val
	} else {
		if val, ok := kv["repo_id"]; ok {
			if val == nil {
				o.RepoID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.RepoID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
	o.setDefaults(false)
}
