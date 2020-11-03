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
	// RepoErrorTopic is the default topic name
	RepoErrorTopic datamodel.TopicNameType = "sourcecode_RepoError"

	// RepoErrorTable is the default table name
	RepoErrorTable datamodel.ModelNameType = "sourcecode_repoerror"

	// RepoErrorModelName is the model name
	RepoErrorModelName datamodel.ModelNameType = "sourcecode.RepoError"
)

const (
	// RepoErrorModelCreatedAtColumn is the column json value created_ts
	RepoErrorModelCreatedAtColumn = "created_ts"
	// RepoErrorModelCustomerIDColumn is the column json value customer_id
	RepoErrorModelCustomerIDColumn = "customer_id"
	// RepoErrorModelErrorMessageColumn is the column json value error_message
	RepoErrorModelErrorMessageColumn = "error_message"
	// RepoErrorModelErroredColumn is the column json value errored
	RepoErrorModelErroredColumn = "errored"
	// RepoErrorModelIDColumn is the column json value id
	RepoErrorModelIDColumn = "id"
	// RepoErrorModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	RepoErrorModelIntegrationInstanceIDColumn = "integration_instance_id"
	// RepoErrorModelRefIDColumn is the column json value ref_id
	RepoErrorModelRefIDColumn = "ref_id"
	// RepoErrorModelRefTypeColumn is the column json value ref_type
	RepoErrorModelRefTypeColumn = "ref_type"
	// RepoErrorModelRepoIDColumn is the column json value repo_id
	RepoErrorModelRepoIDColumn = "repo_id"
	// RepoErrorModelUpdatedAtColumn is the column json value updated_ts
	RepoErrorModelUpdatedAtColumn = "updated_ts"
)

// RepoError the repo error contains details about any errors found during processing
type RepoError struct {
	// CreatedAt the date the record was created in Epoch time
	CreatedAt int64 `json:"created_ts" codec:"created_ts" bson:"created_ts" yaml:"created_ts" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty" codec:"error_message,omitempty" bson:"error_message" yaml:"error_message,omitempty" faker:"-"`
	// Errored if the repo is in an errored state
	Errored bool `json:"errored" codec:"errored" bson:"errored" yaml:"errored" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the repo id
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedAt the date the record was updated in Epoch time
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoError)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*RepoError)(nil)

func toRepoErrorObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *RepoError:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of RepoError
func (o *RepoError) String() string {
	return fmt.Sprintf("sourcecode.RepoError<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoError) GetTopicName() datamodel.TopicNameType {
	return RepoErrorTopic
}

// GetStreamName returns the name of the stream
func (o *RepoError) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *RepoError) GetTableName() string {
	return RepoErrorTable.String()
}

// GetModelName returns the name of the model
func (o *RepoError) GetModelName() datamodel.ModelNameType {
	return RepoErrorModelName
}

// NewRepoErrorID provides a template for generating an ID field for RepoError
func NewRepoErrorID(customerID string, RepoID string) string {
	return hash.Values(customerID, RepoID)
}

func (o *RepoError) setDefaults(frommap bool) {

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
func (o *RepoError) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoError) GetTopicKey() string {
	var i interface{} = o.ID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoError) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for RepoError")
}

// GetRefID returns the RefID for the object
func (o *RepoError) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoError) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *RepoError) IsMutable() bool {
	return true
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoError) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoError) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoError) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoErrorModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoError) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *RepoError) GetCustomerID() string {
	return o.CustomerID
}

// SetCustomerID will return the customer_id
func (o *RepoError) SetCustomerID(id string) {
	o.CustomerID = id
}

// GetRefType will return the ref_type
func (o *RepoError) GetRefType() string {
	return o.RefType
}

// SetRefType will return the ref_type
func (o *RepoError) SetRefType(t string) {
	o.RefType = t
}

// Clone returns an exact copy of RepoError
func (o *RepoError) Clone() datamodel.Model {
	c := new(RepoError)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RepoError) Anon() datamodel.Model {
	c := new(RepoError)
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
func (o *RepoError) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoError) UnmarshalJSON(data []byte) error {
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
func (o *RepoError) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *RepoError) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two RepoError objects are equal
func (o *RepoError) IsEqual(other *RepoError) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoError) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"created_ts":              toRepoErrorObject(o.CreatedAt, false),
		"customer_id":             toRepoErrorObject(o.CustomerID, false),
		"error_message":           toRepoErrorObject(o.ErrorMessage, true),
		"errored":                 toRepoErrorObject(o.Errored, false),
		"id":                      toRepoErrorObject(o.ID, false),
		"integration_instance_id": toRepoErrorObject(o.IntegrationInstanceID, true),
		"ref_id":                  toRepoErrorObject(o.RefID, false),
		"ref_type":                toRepoErrorObject(o.RefType, false),
		"repo_id":                 toRepoErrorObject(o.RepoID, false),
		"updated_ts":              toRepoErrorObject(o.UpdatedAt, false),
		"hashcode":                toRepoErrorObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoError) FromMap(kv map[string]interface{}) {

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
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *RepoError) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.CreatedAt)
	args = append(args, o.CustomerID)
	args = append(args, o.ErrorMessage)
	args = append(args, o.Errored)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedAt)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// SetIntegrationInstanceID will set the integration instance ID
func (o *RepoError) SetIntegrationInstanceID(id string) {
	if id == "" {
		o.IntegrationInstanceID = nil
	} else {
		o.IntegrationInstanceID = &id
	}
}

// GetIntegrationInstanceID will return the integration instance ID
func (o *RepoError) GetIntegrationInstanceID() *string {
	return o.IntegrationInstanceID
}

// GetHydrationQuery returns a query for all fields, and one level deep of relations.
// This query requires "id" to be in the query variables.
func (o *RepoError) GetHydrationQuery() string {
	return `query GoRepoErrorQuery($id: ID, $nocache: Boolean) {
	sourcecode {
		RepoError(_id: $id, nocache: $nocache) {
			created_ts
			customer_id
			error_message
			errored
			_id
			integration_instance_id
			ref_id
			ref_type
			repo_id
			updated_ts
		}
	}
}
`
}

func getRepoErrorQueryFields() string {
	var sb strings.Builder

	// scalar
	sb.WriteString("\t\t\tcreated_ts\n")
	// scalar
	sb.WriteString("\t\t\tcustomer_id\n")
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
	return sb.String()
}

// RepoErrorPageInfo is a grapqhl PageInfo
type RepoErrorPageInfo struct {
	StartCursor     string `json:"startCursor,omitempty"`
	EndCursor       string `json:"endCursor,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
}

// RepoErrorCacheInfo is the grapqhl CacheInfo
type RepoErrorCacheInfo struct {
	Cached bool    `json:"cached,omitempty"`
	ID     *string `json:"id,omitempty"`
	ETag   *string `json:"etag,omitempty"`
}

// RepoErrorConnection is a grapqhl connection
type RepoErrorConnection struct {
	Edges      []*RepoErrorEdge   `json:"edges,omitempty"`
	PageInfo   RepoErrorPageInfo  `json:"pageInfo,omitempty"`
	CacheInfo  RepoErrorCacheInfo `json:"cacheInfo,omitempty"`
	TotalCount *int64             `json:"totalCount,omitempty"`
}

// RepoErrorEdge is a grapqhl edge
type RepoErrorEdge struct {
	Node *RepoError `json:"node,omitempty"`
}

// QueryManyRepoErrorNode is a grapqhl query many node
type QueryManyRepoErrorNode struct {
	Object *RepoErrorConnection `json:"RepoErrors,omitempty"`
}

// QueryManyRepoErrorData is a grapqhl query many data node
type QueryManyRepoErrorData struct {
	Data *QueryManyRepoErrorNode `json:"sourcecode,omitempty"`
}

// QueryOneRepoErrorNode is a grapqhl query one node
type QueryOneRepoErrorNode struct {
	Object *RepoError `json:"RepoError,omitempty"`
}

// QueryOneRepoErrorData is a grapqhl query one data node
type QueryOneRepoErrorData struct {
	Data *QueryOneRepoErrorNode `json:"sourcecode,omitempty"`
}

// RepoErrorQuery is query struct
type RepoErrorQuery struct {
	Filters []string      `json:"filters,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
}

// RepoErrorOrder is the order direction
type RepoErrorOrder *string

var (
	// RepoErrorOrderDesc is descending
	RepoErrorOrderDesc RepoErrorOrder = pstrings.Pointer("DESC")
	// RepoErrorOrderAsc is ascending
	RepoErrorOrderAsc RepoErrorOrder = pstrings.Pointer("ASC")
)

// RepoErrorQueryInput is query input struct
type RepoErrorQueryInput struct {
	First   *int64          `json:"first,omitempty"`
	Last    *int64          `json:"last,omitempty"`
	Before  *string         `json:"before,omitempty"`
	After   *string         `json:"after,omitempty"`
	Query   *RepoErrorQuery `json:"query,omitempty"`
	OrderBy *string         `json:"orderBy,omitempty"`
	Order   RepoErrorOrder  `json:"order,omitempty"`
	NoCache bool            `json:"nocache,omitempty"`
}

// NewRepoErrorQuery is a convenience for building a *RepoErrorQuery
func NewRepoErrorQuery(params ...interface{}) *RepoErrorQueryInput {
	if len(params)%2 != 0 {
		panic("incorrect number of arguments passed")
	}
	q := &RepoErrorQuery{
		Filters: make([]string, 0),
		Params:  make([]interface{}, 0),
	}
	for i := 0; i < len(params); i += 2 {
		q.Filters = append(q.Filters, params[i].(string))
		q.Params = append(q.Params, params[i+1])
	}
	return &RepoErrorQueryInput{
		Query: q,
	}
}

// FindRepoError will query an RepoError by id
func FindRepoError(client graphql.Client, id string) (*RepoError, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("query GoRepoErrorQuery($id: ID) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tRepoError(_id: $id) {\n")
	sb.WriteString(getRepoErrorQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneRepoErrorData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindRepoErrorWithoutCache will query an RepoError by id avoiding the cache
func FindRepoErrorWithoutCache(client graphql.Client, id string) (*RepoError, error) {
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["nocache"] = true
	var sb strings.Builder
	sb.WriteString("query GoRepoErrorQuery($id: ID, $nocache: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tRepoError(_id: $id, nocache: $nocache) {\n")
	sb.WriteString(getRepoErrorQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryOneRepoErrorData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	if res.Data != nil {
		return res.Data.Object, nil
	}
	return nil, nil
}

// FindRepoErrors will query for any RepoErrors matching the query
func FindRepoErrors(client graphql.Client, input *RepoErrorQueryInput) (*RepoErrorConnection, error) {
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
	sb.WriteString("query GoRepoErrorQueryMany($first: Int, $last: Int, $before: Cursor, $after: Cursor, $query: QueryInput, $order: SortOrderEnum, $orderBy: SourcecodeRepoErrorColumnEnum, $nocache: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tRepoErrors(first: $first last: $last before: $before after: $after query: $query orderBy: $orderBy order: $order nocache: $nocache) {\n")
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
	sb.WriteString(getRepoErrorQueryFields())
	sb.WriteString("\t\t\t\t}\n")
	sb.WriteString("\t\t\t}\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res QueryManyRepoErrorData
	if err := client.Query(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// FindRepoErrorsPaginatedCallback is a callback function for handling each page
type FindRepoErrorsPaginatedCallback func(conn *RepoErrorConnection) (bool, error)

// FindRepoErrorsPaginated will query for any RepoErrors matching the query and return each page callback
func FindRepoErrorsPaginated(client graphql.Client, query *RepoErrorQuery, pageSize int64, callback FindRepoErrorsPaginatedCallback) error {
	input := &RepoErrorQueryInput{
		First: &pageSize,
		Query: query,
	}
	for {
		res, err := FindRepoErrors(client, input)
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

// FindRepoErrorsPaginatedWithoutCache will query for any RepoErrors matching the query and return each page callback
func FindRepoErrorsPaginatedWithoutCache(client graphql.Client, query *RepoErrorQuery, pageSize int64, callback FindRepoErrorsPaginatedCallback) error {
	input := &RepoErrorQueryInput{
		First:   &pageSize,
		Query:   query,
		NoCache: true,
	}
	for {
		res, err := FindRepoErrors(client, input)
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

// UpdateRepoErrorNode is a grapqhl update node
type UpdateRepoErrorNode struct {
	Object *RepoError `json:"updateRepoError,omitempty"`
}

// UpdateRepoErrorData is a grapqhl update data node
type UpdateRepoErrorData struct {
	Data *UpdateRepoErrorNode `json:"sourcecode,omitempty"`
}

// ExecRepoErrorUpdateMutation returns a graphql update mutation result for RepoError
func ExecRepoErrorUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) (*RepoError, error) {
	if !upsert && id == "" {
		return nil, fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoRepoErrorUpdateMutation($id: String, $input: UpdateSourcecodeRepoErrorInput, $upsert: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tupdateRepoError(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString(getRepoErrorQueryFields())
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return nil, err
	}
	return res.Data.Object, nil
}

// ExecRepoErrorSilentUpdateMutation returns a graphql update mutation result for RepoError
func ExecRepoErrorSilentUpdateMutation(client graphql.Client, id string, input graphql.Variables, upsert bool) error {
	if !upsert && id == "" {
		return fmt.Errorf("id is required with upsert false")
	}
	variables := make(graphql.Variables)
	variables["id"] = id
	variables["upsert"] = upsert
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoRepoErrorUpdateMutation($id: String, $input: UpdateSourcecodeRepoErrorInput, $upsert: Boolean) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tupdateRepoError(_id: $id, input: $input, upsert: $upsert) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// ExecRepoErrorDeleteMutation executes a graphql delete mutation for RepoError
func ExecRepoErrorDeleteMutation(client graphql.Client, id string) error {
	variables := make(graphql.Variables)
	variables["id"] = id
	var sb strings.Builder
	sb.WriteString("mutation GoRepoErrorDeleteMutation($id: String!) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tdeleteRepoError(_id: $id)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res interface{}
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

func CreateRepoError(client graphql.Client, model RepoError) error {
	variables := make(graphql.Variables)
	input := model.ToMap()
	delete(input, "hashcode")
	delete(input, "customer_id")
	delete(input, "created_ts")
	delete(input, "id")
	input["_id"] = model.GetID()
	variables["input"] = input
	var sb strings.Builder
	sb.WriteString("mutation GoCreateRepoError($input: CreateSourcecodeRepoErrorInput!) {\n")
	sb.WriteString("\tsourcecode {\n")
	sb.WriteString("\t\tcreateRepoError(input: $input) {\n")
	sb.WriteString("\t\t\t_id")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
	var res UpdateRepoErrorData
	if err := client.Mutate(sb.String(), variables, &res); err != nil {
		return err
	}
	return nil
}

// RepoErrorPartial is a partial struct for upsert mutations for RepoError
type RepoErrorPartial struct {
	// ErrorMessage the error message
	ErrorMessage *string `json:"error_message,omitempty"`
	// Errored if the repo is in an errored state
	Errored *bool `json:"errored,omitempty"`
	// RepoID the repo id
	RepoID *string `json:"repo_id,omitempty"`
}

var _ datamodel.PartialModel = (*RepoErrorPartial)(nil)

// GetModelName returns the name of the model
func (o *RepoErrorPartial) GetModelName() datamodel.ModelNameType {
	return RepoErrorModelName
}

// ToMap returns the object as a map
func (o *RepoErrorPartial) ToMap() map[string]interface{} {
	kv := map[string]interface{}{
		"error_message": toRepoErrorObject(o.ErrorMessage, true),
		"errored":       toRepoErrorObject(o.Errored, true),
		"repo_id":       toRepoErrorObject(o.RepoID, true),
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
func (o *RepoErrorPartial) Stringify() string {
	return pjson.Stringify(o.ToMap())
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoErrorPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoErrorPartial) UnmarshalJSON(data []byte) error {
	kv := make(map[string]interface{})
	if err := json.Unmarshal(data, &kv); err != nil {
		return err
	}
	o.FromMap(kv)
	return nil
}

func (o *RepoErrorPartial) setDefaults(frommap bool) {
}

// FromMap attempts to load data into object from a map
func (o *RepoErrorPartial) FromMap(kv map[string]interface{}) {
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
	o.setDefaults(false)
}
