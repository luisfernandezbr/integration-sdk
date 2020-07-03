// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker"
	"github.com/pinpt/go-common/v10/datamodel"
	"github.com/pinpt/go-common/v10/datetime"
	"github.com/pinpt/go-common/v10/hash"
	pjson "github.com/pinpt/go-common/v10/json"
	"github.com/pinpt/go-common/v10/number"
	pstrings "github.com/pinpt/go-common/v10/strings"
)

const (

	// PullRequestCommitTable is the default table name
	PullRequestCommitTable datamodel.ModelNameType = "sourcecode_pullrequestcommit"

	// PullRequestCommitModelName is the model name
	PullRequestCommitModelName datamodel.ModelNameType = "sourcecode.PullRequestCommit"
)

const (
	// PullRequestCommitModelAdditionsColumn is the column json value additions
	PullRequestCommitModelAdditionsColumn = "additions"
	// PullRequestCommitModelAuthorRefIDColumn is the column json value author_ref_id
	PullRequestCommitModelAuthorRefIDColumn = "author_ref_id"
	// PullRequestCommitModelBranchIDColumn is the column json value branch_id
	PullRequestCommitModelBranchIDColumn = "branch_id"
	// PullRequestCommitModelCommitterRefIDColumn is the column json value committer_ref_id
	PullRequestCommitModelCommitterRefIDColumn = "committer_ref_id"
	// PullRequestCommitModelCreatedDateColumn is the column json value created_date
	PullRequestCommitModelCreatedDateColumn = "created_date"
	// PullRequestCommitModelCreatedDateEpochColumn is the column json value epoch
	PullRequestCommitModelCreatedDateEpochColumn = "epoch"
	// PullRequestCommitModelCreatedDateOffsetColumn is the column json value offset
	PullRequestCommitModelCreatedDateOffsetColumn = "offset"
	// PullRequestCommitModelCreatedDateRfc3339Column is the column json value rfc3339
	PullRequestCommitModelCreatedDateRfc3339Column = "rfc3339"
	// PullRequestCommitModelCustomerIDColumn is the column json value customer_id
	PullRequestCommitModelCustomerIDColumn = "customer_id"
	// PullRequestCommitModelDeletionsColumn is the column json value deletions
	PullRequestCommitModelDeletionsColumn = "deletions"
	// PullRequestCommitModelIDColumn is the column json value id
	PullRequestCommitModelIDColumn = "id"
	// PullRequestCommitModelIntegrationInstanceIDColumn is the column json value integration_instance_id
	PullRequestCommitModelIntegrationInstanceIDColumn = "integration_instance_id"
	// PullRequestCommitModelMessageColumn is the column json value message
	PullRequestCommitModelMessageColumn = "message"
	// PullRequestCommitModelPullRequestIDColumn is the column json value pull_request_id
	PullRequestCommitModelPullRequestIDColumn = "pull_request_id"
	// PullRequestCommitModelRefIDColumn is the column json value ref_id
	PullRequestCommitModelRefIDColumn = "ref_id"
	// PullRequestCommitModelRefTypeColumn is the column json value ref_type
	PullRequestCommitModelRefTypeColumn = "ref_type"
	// PullRequestCommitModelRepoIDColumn is the column json value repo_id
	PullRequestCommitModelRepoIDColumn = "repo_id"
	// PullRequestCommitModelShaColumn is the column json value sha
	PullRequestCommitModelShaColumn = "sha"
	// PullRequestCommitModelURLColumn is the column json value url
	PullRequestCommitModelURLColumn = "url"
)

// PullRequestCommitCreatedDate represents the object structure for created_date
type PullRequestCommitCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" codec:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" codec:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" codec:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toPullRequestCommitCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommitCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// ToMap returns the object as a map
func (o *PullRequestCommitCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toPullRequestCommitCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toPullRequestCommitCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toPullRequestCommitCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *PullRequestCommitCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommitCreatedDate) FromMap(kv map[string]interface{}) {

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

// PullRequestCommit the pull request commit is a specific change in a repo that was extracted from a pull request
type PullRequestCommit struct {
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" codec:"additions" bson:"additions" yaml:"additions" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" codec:"author_ref_id" bson:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// BranchID the branch that the commit was a part of
	BranchID string `json:"branch_id" codec:"branch_id" bson:"branch_id" yaml:"branch_id" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" codec:"committer_ref_id" bson:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// CreatedDate date when the commit was created
	CreatedDate PullRequestCommitCreatedDate `json:"created_date" codec:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" codec:"deletions" bson:"deletions" yaml:"deletions" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationInstanceID the integration instance id
	IntegrationInstanceID *string `json:"integration_instance_id,omitempty" codec:"integration_instance_id,omitempty" bson:"integration_instance_id" yaml:"integration_instance_id,omitempty" faker:"-"`
	// Message the commit message
	Message string `json:"message" codec:"message" bson:"message" yaml:"message" faker:"commit_message"`
	// PullRequestID the pull request this commit was taken from
	PullRequestID string `json:"pull_request_id" codec:"pull_request_id" bson:"pull_request_id" yaml:"pull_request_id" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" codec:"repo_id" bson:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" codec:"sha" bson:"sha" yaml:"sha" faker:"sha"`
	// URL the url to the commit detail
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*PullRequestCommit)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*PullRequestCommit)(nil)

func toPullRequestCommitObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *PullRequestCommit:
		return v.ToMap()

	case PullRequestCommitCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of PullRequestCommit
func (o *PullRequestCommit) String() string {
	return fmt.Sprintf("sourcecode.PullRequestCommit<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *PullRequestCommit) GetTopicName() datamodel.TopicNameType {
	return ""
}

// GetStreamName returns the name of the stream
func (o *PullRequestCommit) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *PullRequestCommit) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *PullRequestCommit) GetModelName() datamodel.ModelNameType {
	return PullRequestCommitModelName
}

// NewPullRequestCommitID provides a template for generating an ID field for PullRequestCommit
func NewPullRequestCommitID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *PullRequestCommit) setDefaults(frommap bool) {

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
func (o *PullRequestCommit) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *PullRequestCommit) GetTopicKey() string {
	return ""
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *PullRequestCommit) GetTimestamp() time.Time {
	return time.Now().UTC()
}

// GetRefID returns the RefID for the object
func (o *PullRequestCommit) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *PullRequestCommit) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *PullRequestCommit) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *PullRequestCommit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *PullRequestCommit) IsEvented() bool {
	return false
}

// SetEventHeaders will set any event headers for the object instance
func (o *PullRequestCommit) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = PullRequestCommitModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *PullRequestCommit) GetTopicConfig() *datamodel.ModelTopicConfig {
	return nil
}

// GetCustomerID will return the customer_id
func (o *PullRequestCommit) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of PullRequestCommit
func (o *PullRequestCommit) Clone() datamodel.Model {
	c := new(PullRequestCommit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *PullRequestCommit) Anon() datamodel.Model {
	c := new(PullRequestCommit)
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
func (o *PullRequestCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *PullRequestCommit) UnmarshalJSON(data []byte) error {
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
func (o *PullRequestCommit) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// StringifyPretty returns the object in JSON format as a string prettified
func (o *PullRequestCommit) StringifyPretty() string {
	o.Hash()
	return pjson.Stringify(o, true)
}

// IsEqual returns true if the two PullRequestCommit objects are equal
func (o *PullRequestCommit) IsEqual(other *PullRequestCommit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *PullRequestCommit) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"additions":               toPullRequestCommitObject(o.Additions, false),
		"author_ref_id":           toPullRequestCommitObject(o.AuthorRefID, false),
		"branch_id":               toPullRequestCommitObject(o.BranchID, false),
		"committer_ref_id":        toPullRequestCommitObject(o.CommitterRefID, false),
		"created_date":            toPullRequestCommitObject(o.CreatedDate, false),
		"customer_id":             toPullRequestCommitObject(o.CustomerID, false),
		"deletions":               toPullRequestCommitObject(o.Deletions, false),
		"id":                      toPullRequestCommitObject(o.ID, false),
		"integration_instance_id": toPullRequestCommitObject(o.IntegrationInstanceID, true),
		"message":                 toPullRequestCommitObject(o.Message, false),
		"pull_request_id":         toPullRequestCommitObject(o.PullRequestID, false),
		"ref_id":                  toPullRequestCommitObject(o.RefID, false),
		"ref_type":                toPullRequestCommitObject(o.RefType, false),
		"repo_id":                 toPullRequestCommitObject(o.RepoID, false),
		"sha":                     toPullRequestCommitObject(o.Sha, false),
		"url":                     toPullRequestCommitObject(o.URL, false),
		"hashcode":                toPullRequestCommitObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *PullRequestCommit) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		if val, ok := kv["additions"]; ok {
			if val == nil {
				o.Additions = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Additions = number.ToInt64Any(val)
			}
		}
	}
	if val, ok := kv["author_ref_id"].(string); ok {
		o.AuthorRefID = val
	} else {
		if val, ok := kv["author_ref_id"]; ok {
			if val == nil {
				o.AuthorRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.AuthorRefID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["branch_id"].(string); ok {
		o.BranchID = val
	} else {
		if val, ok := kv["branch_id"]; ok {
			if val == nil {
				o.BranchID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.BranchID = fmt.Sprintf("%v", val)
			}
		}
	}
	if val, ok := kv["committer_ref_id"].(string); ok {
		o.CommitterRefID = val
	} else {
		if val, ok := kv["committer_ref_id"]; ok {
			if val == nil {
				o.CommitterRefID = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.CommitterRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(PullRequestCommitCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*PullRequestCommitCreatedDate); ok {
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
	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		if val, ok := kv["deletions"]; ok {
			if val == nil {
				o.Deletions = 0
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Deletions = number.ToInt64Any(val)
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
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		if val, ok := kv["message"]; ok {
			if val == nil {
				o.Message = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Message = fmt.Sprintf("%v", val)
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
	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = ""
			} else {
				v := pstrings.Value(val)
				if v != "" {
					if m, ok := val.(map[string]interface{}); ok && m != nil {
						val = pjson.Stringify(m)
					}
				} else {
					val = v
				}
				o.Sha = fmt.Sprintf("%v", val)
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
func (o *PullRequestCommit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Additions)
	args = append(args, o.AuthorRefID)
	args = append(args, o.BranchID)
	args = append(args, o.CommitterRefID)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Deletions)
	args = append(args, o.ID)
	args = append(args, o.IntegrationInstanceID)
	args = append(args, o.Message)
	args = append(args, o.PullRequestID)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// PullRequestCommitPartial is a partial struct for upsert mutations for PullRequestCommit
type PullRequestCommitPartial struct {
	// Additions the number of additions for the commit
	Additions *int64 `json:"additions,omitempty"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID *string `json:"author_ref_id,omitempty"`
	// BranchID the branch that the commit was a part of
	BranchID *string `json:"branch_id,omitempty"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID *string `json:"committer_ref_id,omitempty"`
	// CreatedDate date when the commit was created
	CreatedDate *PullRequestCommitCreatedDate `json:"created_date,omitempty"`
	// Deletions the number of deletions for the commit
	Deletions *int64 `json:"deletions,omitempty"`
	// Message the commit message
	Message *string `json:"message,omitempty"`
	// PullRequestID the pull request this commit was taken from
	PullRequestID *string `json:"pull_request_id,omitempty"`
	// RepoID the unique id for the repo
	RepoID *string `json:"repo_id,omitempty"`
	// Sha the unique sha for the commit
	Sha *string `json:"sha,omitempty"`
	// URL the url to the commit detail
	URL *string `json:"url,omitempty"`
}

var _ datamodel.PartialModel = (*PullRequestCommitPartial)(nil)

// GetModelName returns the name of the model
func (o *PullRequestCommitPartial) GetModelName() datamodel.ModelNameType {
	return PullRequestCommitModelName
}

// ToMap returns the object as a map
func (o *PullRequestCommitPartial) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"additions":        toPullRequestCommitObject(o.Additions, true),
		"author_ref_id":    toPullRequestCommitObject(o.AuthorRefID, true),
		"branch_id":        toPullRequestCommitObject(o.BranchID, true),
		"committer_ref_id": toPullRequestCommitObject(o.CommitterRefID, true),
		"created_date":     toPullRequestCommitObject(o.CreatedDate, true),
		"deletions":        toPullRequestCommitObject(o.Deletions, true),
		"message":          toPullRequestCommitObject(o.Message, true),
		"pull_request_id":  toPullRequestCommitObject(o.PullRequestID, true),
		"repo_id":          toPullRequestCommitObject(o.RepoID, true),
		"sha":              toPullRequestCommitObject(o.Sha, true),
		"url":              toPullRequestCommitObject(o.URL, true),
	}
}

// Stringify returns the object in JSON format as a string
func (o *PullRequestCommitPartial) Stringify() string {
	return pjson.Stringify(o)
}
