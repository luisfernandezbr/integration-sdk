// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	"github.com/pinpt/go-common/slice"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// BranchModelName is the model name
	BranchModelName datamodel.ModelNameType = "sourcecode.Branch"
)

// Branch git branches
type Branch struct {
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
	// FirstBranchedFromCommitID the first commit id from which the branch was created
	FirstBranchedFromCommitID string `json:"first_branched_from_commit_id" codec:"first_branched_from_commit_id" bson:"first_branched_from_commit_id" yaml:"first_branched_from_commit_id" faker:"-"`
	// FirstBranchedFromCommitSha the first commit sha from which the branch was created
	FirstBranchedFromCommitSha string `json:"first_branched_from_commit_sha" codec:"first_branched_from_commit_sha" bson:"first_branched_from_commit_sha" yaml:"first_branched_from_commit_sha" faker:"-"`
	// FirstCommitID the first commit id on the branch
	FirstCommitID string `json:"first_commit_id" codec:"first_commit_id" bson:"first_commit_id" yaml:"first_commit_id" faker:"-"`
	// FirstCommitSha the first commit sha on the branch
	FirstCommitSha string `json:"first_commit_sha" codec:"first_commit_sha" bson:"first_commit_sha" yaml:"first_commit_sha" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// MergeCommitID commit id in which the branch was merged
	MergeCommitID string `json:"merge_commit_id" codec:"merge_commit_id" bson:"merge_commit_id" yaml:"merge_commit_id" faker:"-"`
	// MergeCommitSha commit sha in which the branch was merged
	MergeCommitSha string `json:"merge_commit_sha" codec:"merge_commit_sha" bson:"merge_commit_sha" yaml:"merge_commit_sha" faker:"-"`
	// Merged true if the branch has been merged into the default branch
	Merged bool `json:"merged" codec:"merged" bson:"merged" yaml:"merged" faker:"-"`
	// Name name of the branch
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"-"`
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
var _ datamodel.Model = (*Branch)(nil)

func toBranchObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Branch:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Branch
func (o *Branch) String() string {
	return fmt.Sprintf("sourcecode.Branch<%s>", o.ID)
}

// GetModelName returns the name of the model
func (o *Branch) GetModelName() datamodel.ModelNameType {
	return BranchModelName
}

// NewBranchID provides a template for generating an ID field for Branch
func NewBranchID(refType string, RepoID string, customerID string, Name string, FirstCommitID string) string {
	return hash.Values(refType, RepoID, customerID, Name, FirstCommitID)
}

func (o *Branch) setDefaults(frommap bool) {
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

	if o.ID == "" {
		o.ID = hash.Values(o.RefType, o.RepoID, o.CustomerID, o.Name, o.FirstCommitID)
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Branch) GetID() string {
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Branch) GetRefID() string {
	return o.RefID
}

// GetCustomerID will return the customer_id
func (o *Branch) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Branch
func (o *Branch) Clone() datamodel.Model {
	c := new(Branch)
	c.FromMap(o.ToMap())
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Branch) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Branch) UnmarshalJSON(data []byte) error {
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
func (o *Branch) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Branch objects are equal
func (o *Branch) IsEqual(other *Branch) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Branch) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"ahead_default_count":            toBranchObject(o.AheadDefaultCount, false),
		"behind_default_count":           toBranchObject(o.BehindDefaultCount, false),
		"branched_from_commit_ids":       toBranchObject(o.BranchedFromCommitIds, false),
		"branched_from_commit_shas":      toBranchObject(o.BranchedFromCommitShas, false),
		"commit_ids":                     toBranchObject(o.CommitIds, false),
		"commit_shas":                    toBranchObject(o.CommitShas, false),
		"customer_id":                    toBranchObject(o.CustomerID, false),
		"default":                        toBranchObject(o.Default, false),
		"first_branched_from_commit_id":  toBranchObject(o.FirstBranchedFromCommitID, false),
		"first_branched_from_commit_sha": toBranchObject(o.FirstBranchedFromCommitSha, false),
		"first_commit_id":                toBranchObject(o.FirstCommitID, false),
		"first_commit_sha":               toBranchObject(o.FirstCommitSha, false),
		"id":                             toBranchObject(o.ID, false),
		"merge_commit_id":                toBranchObject(o.MergeCommitID, false),
		"merge_commit_sha":               toBranchObject(o.MergeCommitSha, false),
		"merged":                         toBranchObject(o.Merged, false),
		"name":                           toBranchObject(o.Name, false),
		"ref_id":                         toBranchObject(o.RefID, false),
		"ref_type":                       toBranchObject(o.RefType, false),
		"repo_id":                        toBranchObject(o.RepoID, false),
		"updated_ts":                     toBranchObject(o.UpdatedAt, false),
		"url":                            toBranchObject(o.URL, false),
		"hashcode":                       toBranchObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Branch) FromMap(kv map[string]interface{}) {

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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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

	if val, ok := kv["first_branched_from_commit_id"].(string); ok {
		o.FirstBranchedFromCommitID = val
	} else {
		if val, ok := kv["first_branched_from_commit_id"]; ok {
			if val == nil {
				o.FirstBranchedFromCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstBranchedFromCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_branched_from_commit_sha"].(string); ok {
		o.FirstBranchedFromCommitSha = val
	} else {
		if val, ok := kv["first_branched_from_commit_sha"]; ok {
			if val == nil {
				o.FirstBranchedFromCommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstBranchedFromCommitSha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_commit_id"].(string); ok {
		o.FirstCommitID = val
	} else {
		if val, ok := kv["first_commit_id"]; ok {
			if val == nil {
				o.FirstCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstCommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["first_commit_sha"].(string); ok {
		o.FirstCommitSha = val
	} else {
		if val, ok := kv["first_commit_sha"]; ok {
			if val == nil {
				o.FirstCommitSha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.FirstCommitSha = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["merge_commit_id"].(string); ok {
		o.MergeCommitID = val
	} else {
		if val, ok := kv["merge_commit_id"]; ok {
			if val == nil {
				o.MergeCommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Name = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.URL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *Branch) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.AheadDefaultCount)
	args = append(args, o.BehindDefaultCount)
	args = append(args, o.BranchedFromCommitIds)
	args = append(args, o.BranchedFromCommitShas)
	args = append(args, o.CommitIds)
	args = append(args, o.CommitShas)
	args = append(args, o.CustomerID)
	args = append(args, o.Default)
	args = append(args, o.FirstBranchedFromCommitID)
	args = append(args, o.FirstBranchedFromCommitSha)
	args = append(args, o.FirstCommitID)
	args = append(args, o.FirstCommitSha)
	args = append(args, o.ID)
	args = append(args, o.MergeCommitID)
	args = append(args, o.MergeCommitSha)
	args = append(args, o.Merged)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
