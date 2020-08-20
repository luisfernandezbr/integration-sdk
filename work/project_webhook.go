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
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (
	// ProjectWebhookTopic is the default topic name
	ProjectWebhookTopic datamodel.TopicNameType = "work_ProjectWebhook"

	// ProjectWebhookTable is the default table name
	ProjectWebhookTable datamodel.ModelNameType = "work_projectwebhook"

	// ProjectWebhookModelName is the model name
	ProjectWebhookModelName datamodel.ModelNameType = "work.ProjectWebhook"
)

const (
	// ProjectWebhookModelCreatedAtColumn is the column json value created_ts
	ProjectWebhookModelCreatedAtColumn = "created_ts"
	// ProjectWebhookModelCustomerIDColumn is the column json value customer_id
	ProjectWebhookModelCustomerIDColumn = "customer_id"
	// ProjectWebhookModelEnabledColumn is the column json value enabled
	ProjectWebhookModelEnabledColumn = "enabled"
	// ProjectWebhookModelErrorMessageColumn is the column json value error_message
	ProjectWebhookModelErrorMessageColumn = "error_message"
	// ProjectWebhookModelErroredColumn is the column json value errored
	ProjectWebhookModelErroredColumn = "errored"
	// ProjectWebhookModelIDColumn is the column json value id
	ProjectWebhookModelIDColumn = "id"
	// ProjectWebhookModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	ProjectWebhookModelIntegrationInstanceIDColumn = "integration_instance_id"
	// ProjectWebhookModelProjectIDColumn is the column json value project_id
	ProjectWebhookModelProjectIDColumn = "project_id"
	// ProjectWebhookModelRefIDColumn is the column json value ref_id
	ProjectWebhookModelRefIDColumn = "ref_id"
	// ProjectWebhookModelRefTypeColumn is the column json value ref_type
	ProjectWebhookModelRefTypeColumn = "ref_type"
	// ProjectWebhookModelUpdatedAtColumn is the column json value updated_ts
	ProjectWebhookModelUpdatedAtColumn = "updated_ts"
	// ProjectWebhookModelURLColumn is the column json value url
	ProjectWebhookModelURLColumn = "url"
)

// ProjectWebhook the project webhook contains details about the webhook if one is installed at the project level
type ProjectWebhook struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Enabled if webhooks are enabled for this project
	Enabled bool `json:"enabled" codec:"enabled" bson:"enabled" yaml:"enabled" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the webhook has an error
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// ProjectID the project id
	ProjectID string `json:"project_id" codec:"project_id" bson:"project_id" yaml:"project_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty" codec:"url,omitempty" bson:"url" yaml:"url,omitempty" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*ProjectWebhook)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*ProjectWebhook)(nil)

func toProjectWebhookObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *ProjectWebhook:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of ProjectWebhook
func (o *ProjectWebhook) String() string {
	return fmt.Sprintf("work.ProjectWebhook<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *ProjectWebhook) GetTopicName() datamodel.TopicNameType {
	return ProjectWebhookTopic
}

// GetStreamName returns the name of the stream
func (o *ProjectWebhook) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *ProjectWebhook) GetTableName() string {
	return ProjectWebhookTable.String()
}

// GetModelName returns the name of the model
func (o *ProjectWebhook) GetModelName() datamodel.ModelNameType {
	return ProjectWebhookModelName
}

// NewProjectWebhookID provides a template for generating an ID field for ProjectWebhook
func NewProjectWebhookID(customerID string, ProjectID string) string {
	return hash.Values(customerID, ProjectID)
}

func (o *ProjectWebhook) setDefaults(frommap bool) {

	if o.ID == "" {
		// set the id from the spec provided in the model
		o.ID = hash.Values(o.CustomerID, o.ProjectID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *ProjectWebhook) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *ProjectWebhook) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *ProjectWebhook) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for ProjectWebhook")
}

// GetRefID returns the RefID for the object
func (o *ProjectWebhook) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *ProjectWebhook) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *ProjectWebhook) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *ProjectWebhook) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *ProjectWebhook) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *ProjectWebhook) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = ProjectWebhookModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *ProjectWebhook) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *ProjectWebhook) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of ProjectWebhook
func (o *ProjectWebhook) Clone() datamodel.Model {
	c := new(ProjectWebhook)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *ProjectWebhook) Anon() datamodel.Model {
	c := new(ProjectWebhook)
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
func (o *ProjectWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectWebhook) UnmarshalJSON(data []byte) error {
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
func (o *ProjectWebhook) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *ProjectWebhook) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two ProjectWebhook objects are equal
func (o *ProjectWebhook) IsEqual(other *ProjectWebhook) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *ProjectWebhook) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toProjectWebhookObject(o.CreatedAt, false),
		"customer_id":             toProjectWebhookObject(o.CustomerID, false),
		"enabled":                 toProjectWebhookObject(o.Enabled, false),
		"error_message":           toProjectWebhookObject(o.ErrorMessage, true),
		"errored":                 toProjectWebhookObject(o.Errored, false),
		"id":                      toProjectWebhookObject(o.ID, false),
		"integration_instance_id": toProjectWebhookObject(o.IntegrationInstanceID, true),
		"project_id":              toProjectWebhookObject(o.ProjectID, false),
		"ref_id":                  toProjectWebhookObject(o.RefID, false),
		"ref_type":                toProjectWebhookObject(o.RefType, false),
		"updated_ts":              toProjectWebhookObject(o.UpdatedAt, false),
		"url":                     toProjectWebhookObject(o.URL, true),
		"hashcode":                toProjectWebhookObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *ProjectWebhook) FromMap(kv map[string]interface{}) {

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
	if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.ProjectID = fmt.Sprintf("%v", val)
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
func (o *ProjectWebhook) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.Enabled)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.ProjectID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *ProjectWebhook) GetHydrationQuery() string {
	return `query GoProjectWebhookQuery($id: ID) {
	work {
		ProjectWebhook(_id: $id) {
			created_ts
			customer_id
			enabled
			error_message
			errored
			_id
			integration_instance_id
			project_id
			ref_id
			ref_type
			updated_ts
			url
		}
	}
}
`
}

func getProjectWebhookQueryFields() string {
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
	sb.WriteString("\t\t\tproject_id\n")
	// scalar
	sb.WriteString("\t\t\tref_id\n")
	// scalar
	sb.WriteString("\t\t\tref_type\n")
	// scalar
	sb.WriteString("\t\t\tupdated_ts\n")
	// scalar
	sb.WriteString("\t\t\turl\n")
	return sb.String()
}

// ProjectWebhookPageInfo is a grapqhl PageInfo
type ProjectWebhookPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// ProjectWebhookConnection is a grapqhl connection
type ProjectWebhookConnection struct {
	Edges      []*ProjectWebhookEdge  `json:"edges,omitempty"`
	PageInfo   ProjectWebhookPageInfo `json:"pageInfo,omitempty"`
	TotalCount *int64                 `json:"totalCount,omitempty"`
}

// ProjectWebhookEdge is a grapqhl edge
type ProjectWebhookEdge struct {
	Node *ProjectWebhook `json:"node,omitempty"`
}

// QueryManyProjectWebhookNode is a grapqhl query many node
type QueryManyProjectWebhookNode struct {
	Object *ProjectWebhookConnection `json:"ProjectWebhooks,omitempty"`
}

// QueryManyProjectWebhookData is a grapqhl query many data node
type QueryManyProjectWebhookData struct {
	Data *QueryManyProjectWebhookNode `json:"work,omitempty"`
}

// QueryOneProjectWebhookNode is a grapqhl query one node
type QueryOneProjectWebhookNode struct {
	Object *ProjectWebhook `json:"ProjectWebhook,omitempty"`
}

// QueryOneProjectWebhookData is a grapqhl query one data node
type QueryOneProjectWebhookData struct {
	Data *QueryOneProjectWebhookNode `json:"work,omitempty"`
}

// ProjectWebhookQuery is query struct
type ProjectWebhookQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// ProjectWebhookOrder is the order direction
type ProjectWebhookOrder *string

var (
	// ProjectWebhookOrderDesc is descending
	ProjectWebhookOrderDesc ProjectWebhookOrder = pstrings.Pointer("DESC")
	// ProjectWebhookOrderAsc is ascending
	ProjectWebhookOrderAsc ProjectWebhookOrder = pstrings.Pointer("ASC")
)

// ProjectWebhookQueryInput is query input struct
type ProjectWebhookQueryInput struct {
	First   *int64               `json:"first,omitempty"`
	Last    *int64               `json:"last,omitempty"`
	Before  *string              `json:"before,omitempty"`
	After   *string              `json:"after,omitempty"`
	Query   *ProjectWebhookQuery `json:"query,omitempty"`
	OrderBy *string              `json:"orderBy,omitempty"`
	Order   ProjectWebhookOrder  `json:"order,omitempty"`
}

// NewProjectWebhookQuery is a convenience for building a *ProjectWebhookQuery
func NewProjectWebhookQuery(params ...interface{}) *ProjectWebhookQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &ProjectWebhookQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &ProjectWebhookQueryInput{
		Query: q,
	}
}

// FindProjectWebhook will query an ProjectWebhook by id
func FindProjectWebhook(client graphql.Client, id string) (*ProjectWebhook, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoProjectWebhookQuery($id: ID) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tProjectWebhook(_id: $id) {\n")
	sb.WriteString(getProjectWebhookQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneProjectWebhookData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindProjectWebhooks will query for any ProjectWebhooks matching the query
func FindProjectWebhooks(client graphql.Client, input *ProjectWebhookQueryInput) (*ProjectWebhookConnection, error) {
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
	sb.WriteString("query GoProjectWebhookQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: WorkProjectWebhookColumnEnum) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tProjectWebhooks(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order) {\n")
	sb.WriteString("\t\t\tpageInfo {\n")
	sb.WriteString("\t\t\t\tstartCursor\n")
	sb.WriteString("\t\t\t\tendCursor\n")
	sb.WriteString("\t\t\t\thasNextPage\n")
	sb.WriteString("\t\t\t\thasPreviousPage\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t\ttotalCount\n")
	sb.WriteString("\t\t\tedges {\n")
	sb.WriteString("\t\t\t\tnode {\n")
	sb.WriteString(getProjectWebhookQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyProjectWebhookData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindProjectWebhooksPaginatedCallback is a callback function for handling each page
type FindProjectWebhooksPaginatedCallback func(conn *ProjectWebhookConnection) (bool, error)

// FindProjectWebhooksPaginated will query for any ProjectWebhooks matching the query and return each page callback
func FindProjectWebhooksPaginated(client graphql.Client, query *ProjectWebhookQuery, pageSize int64, callback FindProjectWebhooksPaginatedCallback) error {
	input := &ProjectWebhookQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindProjectWebhooks(client, input)
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

// UpdateProjectWebhookNode is a grapqhl update node
type UpdateProjectWebhookNode struct {
	Object *ProjectWebhook `json:"updateProjectWebhook,omitempty"`
}

// UpdateProjectWebhookData is a grapqhl update data node
type UpdateProjectWebhookData struct {
	Data *UpdateProjectWebhookNode `json:"work,omitempty"`
}

// ExecProjectWebhookUpdateMutation returns a graphql update mutation result for ProjectWebhook
func ExecProjectWebhookUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*ProjectWebhook, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoProjectWebhookUpdateMutation($id: String, $input: UpdateWorkProjectWebhookInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateProjectWebhook(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getProjectWebhookQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecProjectWebhookSilentUpdateMutation returns a graphql update mutation result for ProjectWebhook
func ExecProjectWebhookSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoProjectWebhookUpdateMutation($id: String, $input: UpdateWorkProjectWebhookInput, $upsert: Boolean) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tupdateProjectWebhook(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecProjectWebhookDeleteMutation executes a graphql delete mutation for ProjectWebhook
func ExecProjectWebhookDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoProjectWebhookDeleteMutation($id: String!) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tdeleteProjectWebhook(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateProjectWebhook(client graphql.Client, model ProjectWebhook) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateProjectWebhook($input: CreateWorkProjectWebhookInput!) {\n")
	sb.WriteString("\twork {\n")
	sb.WriteString("\t\tcreateProjectWebhook(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateProjectWebhookData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ProjectWebhookPartial is a partial struct for upsert mutations for ProjectWebhook
type ProjectWebhookPartial struct {
	// Enabled if webhooks are enabled for this project
	Enabled *bool `json:"enabled,omitempty"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored if the webhook has an error
	Errored *bool `json:"errored,omitempty"`
	// ProjectID the project id
	ProjectID *string `json:"project_id,omitempty"`
	// URL the url the webhook for the webhook
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*ProjectWebhookPartial)(nil)

// GetModelName returns the name of the model
func (o *ProjectWebhookPartial) GetModelName() datamodel.ModelNameType {
	return ProjectWebhookModelName
}

// ToMap returns the object as a map
func (o *ProjectWebhookPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"enabled":       toProjectWebhookObject(o.Enabled, true),
		"error_message": toProjectWebhookObject(o.ErrorMessage, true),
		"errored":       toProjectWebhookObject(o.Errored, true),
		"project_id":    toProjectWebhookObject(o.ProjectID, true),
		"url":           toProjectWebhookObject(o.URL, true),
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
func (o *ProjectWebhookPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *ProjectWebhookPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *ProjectWebhookPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *ProjectWebhookPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *ProjectWebhookPartial) FromMap(kv map[string]interface{}) {
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
	if val, ok := kv["project_id"].(*string); ok {
		o.ProjectID = val
	} else if val, ok := kv["project_id"].(string); ok {
		o.ProjectID = &val
	} else {
		if val, ok := kv["project_id"]; ok {
			if val == nil {
				o.ProjectID = pstrings.Pointer("")
			} else {
				// if coming in as map, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.ProjectID = pstrings.Pointer(fmt.Sprintf("%v", val))
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
