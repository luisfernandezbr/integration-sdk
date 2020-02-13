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
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// PullRequestBranchTopic is the default topic name
	PullRequestBranchTopic datamodel.TopicNameType = "sourcecode_PullRequestBranch_topic"

	// PullRequestBranchTable is the default table name
	PullRequestBranchTable datamodel.ModelNameType = "sourcecode_pullrequestbranch"

	// PullRequestBranchModelName is the model name
	PullRequestBranchModelName datamodel.ModelNameType = "sourcecode.PullRequestBranch"
)

// PullRequestBranch the pull request branch is a branch in a repo that was extracted from a pull request instead of from git
type PullRequestBranch struct {
	// AheadDefaultCount the number of commits that this branch is ahead of the default branch
	AheadDefaultCount int64 `json:"ahead_default_count" codec:"ahead_default_count" bson:"ahead_default_count" yaml:"ahead_default_count" faker:"-"`
	// BehindDefaultCount the number of commits that this branch is behind from the default branch
	BehindDefaultCount int64 `json:"behind_default_count" codec:"behind_default_count" bson:"behind_default_count" yaml:"behind_default_count" faker:"-"`
	// BranchedFromCommitIds the commit ids from which the branch was created
	BranchedFromCommitIds []string `json:"branched_from_commit_ids" codec:"branched_from_commit_ids" bson:"branched_from_commit_ids" yaml:"branched_from_commit_ids" faker:"-"`
	// BranchedFromCommitShas the commit shas from which the branch was created
	BranchedFromCommitShas []string `json:"branched_from_commit_shas" codec:"branched_from_commit_shas" bson:"branched_from_commit_shas" yaml:"branched_from_commit_shas" faker:"-"`
	// CommitIds list of commit ids contained on this branch
	CommitIds []string `json:"commit_ids" codec:"commit_ids" bson:"commit_ids" yaml:"commit_ids" faker:"-"`
	// CommitShas list of commit shas contained on this branch
	CommitShas []string `json:"commit_shas" codec:"commit_shas" bson:"commit_shas" yaml:"commit_shas" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Default true if the branch is the default branch
	Default bool `json:"default" codec:"default" bson:"default" yaml:"default" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationIds the integration IDs for this model object
	IntegrationIds []string `json:"integration_ids" codec:"integration_ids" bson:"integration_ids" yaml:"integration_ids" faker:"-"`
	// MergeCommitID commit id in which the branch was merged
	MergeCommitID string `json:"merge_commit_id" codec:"merge_commit_id" bson:"merge_commit_id" yaml:"merge_commit_id" faker:"-"`
	// MergeCommitSha commit sha in which the branch was merged
	MergeCommitSha string `json:"merge_commit_sha" codec:"merge_commit_sha" bson:"merge_commit_sha" yaml:"merge_commit_sha" faker:"-"`
	// Merged true if the branch has been merged into the default branch
	Merged bool `json:"merged" codec:"merged" bson:"merged" yaml:"merged" faker:"-"`
	// Name name of the branch
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
	// PullRequestID pull request that provided sha for this branch
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the URL to the source system for this branch
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestBranch)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequestBranch)(nil)

func toPullRequestBranchObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestBranch:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequestBranch
func (o *PullRequestBranch) String() string {
	return fmt.Sprintf("sourcecode.PullRequestBranch<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestBranch) GetTopicName() datamodel.TopicNameType {
	return PullRequestBranchTopic
}

// GetStreamName returns the name of the stream
func (o *PullRequestBranch) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequestBranch) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestBranch) GetModelName() datamodel.ModelNameType {
	return PullRequestBranchModelName
}

// NewPullRequestBranchID provides a template for generating an ID field for PullRequestBranch
func NewPullRequestBranchID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestBranch) setDefaults(frommap bool) {
	if o.BranchedFromCommitIds == nil {
		o.BranchedFromCommitIds = make([]string, 0)
	}
	if o.BranchedFromCommitShas == nil {
		o.BranchedFromCommitShas = make([]string, 0)
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
	}
	if o.IntegrationIds == nil {
		o.IntegrationIds = make([]string, 0)
	}

	if o.ID == "" {
		o.ID = hash.Values(o.CustomerID, o.RefID, o.RefType, o.RepoID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *PullRequestBranch) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestBranch) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestBranch) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for PullRequestBranch")
}

// GetRefID returns the RefID for the object
func (o *PullRequestBranch) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestBranch) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequestBranch) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestBranch) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestBranch) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestBranch) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestBranchModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestBranch) GetTopicConfig() *datamodel.ModelTopicConfig {
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
		Key:               "repo_id",
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
func (o *PullRequestBranch) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestBranch
func (o *PullRequestBranch) Clone() datamodel.Model {
	c := new(PullRequestBranch)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestBranch) Anon() datamodel.Model {
	c := new(PullRequestBranch)
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
func (o *PullRequestBranch) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestBranch) UnmarshalJSON(data []byte) error {
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
func (o *PullRequestBranch) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two PullRequestBranch objects are equal
func (o *PullRequestBranch) IsEqual(other *PullRequestBranch) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestBranch) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"ahead_default_count":       toPullRequestBranchObject(o.AheadDefaultCount, false),
		"behind_default_count":      toPullRequestBranchObject(o.BehindDefaultCount, false),
		"branched_from_commit_ids":  toPullRequestBranchObject(o.BranchedFromCommitIds, false),
		"branched_from_commit_shas": toPullRequestBranchObject(o.BranchedFromCommitShas, false),
		"commit_ids":                toPullRequestBranchObject(o.CommitIds, false),
		"commit_shas":               toPullRequestBranchObject(o.CommitShas, false),
		"customer_id":               toPullRequestBranchObject(o.CustomerID, false),
		"default":                   toPullRequestBranchObject(o.Default, false),
		"id":                        toPullRequestBranchObject(o.ID, false),
		"integration_ids":           toPullRequestBranchObject(o.IntegrationIds, false),
		"merge_commit_id":           toPullRequestBranchObject(o.MergeCommitID, false),
		"merge_commit_sha":          toPullRequestBranchObject(o.MergeCommitSha, false),
		"merged":                    toPullRequestBranchObject(o.Merged, false),
		"name":                      toPullRequestBranchObject(o.Name, false),
		"pull_request_id":           toPullRequestBranchObject(o.PullRequestID, false),
		"ref_id":                    toPullRequestBranchObject(o.RefID, false),
		"ref_type":                  toPullRequestBranchObject(o.RefType, false),
		"repo_id":                   toPullRequestBranchObject(o.RepoID, false),
		"updated_ts":                toPullRequestBranchObject(o.UpdatedAt, false),
		"url":                       toPullRequestBranchObject(o.URL, false),
		"hashcode":                  toPullRequestBranchObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestBranch) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["ahead_default_count"].(int64); ok {
		o.AheadDefaultCount = val
	} else {
		if val, ok := kv["ahead_default_count"]; ok {
			if val == nil {
				o.AheadDefaultCount = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.AheadDefaultCount = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["behind_default_count"].(int64); ok {
		o.BehindDefaultCount = val
	} else {
		if val, ok := kv["behind_default_count"]; ok {
			if val == nil {
				o.BehindDefaultCount = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.BehindDefaultCount = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["branched_from_commit_ids"]; ok {
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
								panic("unsupported type for branched_from_commit_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for branched_from_commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for branched_from_commit_ids field")
				}
			}
			o.BranchedFromCommitIds = na
		}
	}
	if o.BranchedFromCommitIds == nil {
		o.BranchedFromCommitIds = make([]string, 0)
	}

	if val, ok := kv["branched_from_commit_shas"]; ok {
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
								panic("unsupported type for branched_from_commit_shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for branched_from_commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for branched_from_commit_shas field")
				}
			}
			o.BranchedFromCommitShas = na
		}
	}
	if o.BranchedFromCommitShas == nil {
		o.BranchedFromCommitShas = make([]string, 0)
	}

	if val, ok := kv["commit_ids"]; ok {
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
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for commit_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_ids field")
				}
			}
			o.CommitIds = na
		}
	}
	if o.CommitIds == nil {
		o.CommitIds = make([]string, 0)
	}

	if val, ok := kv["commit_shas"]; ok {
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
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for commit_shas field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for commit_shas field")
				}
			}
			o.CommitShas = na
		}
	}
	if o.CommitShas == nil {
		o.CommitShas = make([]string, 0)
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

	if val, ok := kv["default"].(bool); ok {
		o.Default = val
	} else {
		if val, ok := kv["default"]; ok {
			if val == nil {
				o.Default = number.ToBoolAny(nil)
			} else {
				o.Default = number.ToBoolAny(val)
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

	if val, ok := kv["integration_ids"]; ok {
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
								panic("unsupported type for integration_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for integration_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for integration_ids field")
				}
			}
			o.IntegrationIds = na
		}
	}
	if o.IntegrationIds == nil {
		o.IntegrationIds = make([]string, 0)
	}

	if val, ok := kv["merge_commit_id"].(string); ok {
		o.MergeCommitID = val
	} else {
		if val, ok := kv["merge_commit_id"]; ok {
			if val == nil {
				o.MergeCommitID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.MergeCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merge_commit_sha"].(string); ok {
		o.MergeCommitSha = val
	} else {
		if val, ok := kv["merge_commit_sha"]; ok {
			if val == nil {
				o.MergeCommitSha = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.MergeCommitSha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merged"].(bool); ok {
		o.Merged = val
	} else {
		if val, ok := kv["merged"]; ok {
			if val == nil {
				o.Merged = number.ToBoolAny(nil)
			} else {
				o.Merged = number.ToBoolAny(val)
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

	if val, ok := kv["pull_request_id"].(string); ok {
		o.PullRequestID = val
	} else {
		if val, ok := kv["pull_request_id"]; ok {
			if val == nil {
				o.PullRequestID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.PullRequestID = fmt.Sprintf("%v", val)
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
				o.UpdatedAt = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.UpdatedAt = number.ToInt64Any(val)
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

// Hash will return a hashcode for the object
func (o *PullRequestBranch) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AheadDefaultCount)
	args = append(args, o.BehindDefaultCount)
	args = append(args, o.BranchedFromCommitIds)
	args = append(args, o.BranchedFromCommitShas)
	args = append(args, o.CommitIds)
	args = append(args, o.CommitShas)
	args = append(args, o.CustomerID)
	args = append(args, o.Default)
	args = append(args, o.ID)
	args = append(args, o.IntegrationIds)
	args = append(args, o.MergeCommitID)
	args = append(args, o.MergeCommitSha)
	args = append(args, o.Merged)
	args = append(args, o.Name)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *PullRequestBranch) GetEventAPIConfig() datamodel.EventAPIConfig {
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
