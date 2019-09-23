// DO NOT EDIT -- generated code

// Package sourcecode - the system which contains source code
package sourcecode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
)

const (
	// RepoModelName is the model name
	RepoModelName datamodel.ModelNameType = "sourcecode.Repo"
)

// Repo the repo holds source code
type Repo struct {
	// Active the status of the repo
	Active bool `json:"active" codec:"active" bson:"active" yaml:"active" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" codec:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// DefaultBranch the repo default branch
	DefaultBranch string `json:"default_branch" codec:"default_branch" bson:"default_branch" yaml:"default_branch" faker:"-"`
	// Description brief explanation of the repo
	Description string `json:"description" codec:"description" bson:"description" yaml:"description" faker:"sentence"`
	// ID the primary key for the model instance
	ID string `json:"id" codec:"id" bson:"_id" yaml:"id" faker:"-"`
	// Language the programming language the source code is primarily written in
	Language string `json:"language" codec:"language" bson:"language" yaml:"language" faker:"-"`
	// Name the name of the repo
	Name string `json:"name" codec:"name" bson:"name" yaml:"name" faker:"repo"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" codec:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" codec:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" codec:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the repo home page
	URL string `json:"url" codec:"url" bson:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" codec:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Repo)(nil)

func toRepoObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Repo:
		return v.ToMap()

	default:
		return o
	}
}

// String returns a string representation of Repo
func (o *Repo) String() string {
	return fmt.Sprintf("sourcecode.Repo<%s>", o.ID)
}

// GetModelName returns the name of the model
func (o *Repo) GetModelName() datamodel.ModelNameType {
	return RepoModelName
}

// NewRepoID provides a template for generating an ID field for Repo
func NewRepoID(customerID string, refType string, refID string) string {
	return hash.Values("Repo", customerID, refType, refID)
}

func (o *Repo) setDefaults(frommap bool) {

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("Repo", o.CustomerID, o.RefType, o.GetRefID())
	}

	{
		v := "master"

		o.DefaultBranch = v

	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *Repo) GetID() string {
	return o.ID
}

// GetRefID returns the RefID for the object
func (o *Repo) GetRefID() string {
	return o.RefID
}

// GetCustomerID will return the customer_id
func (o *Repo) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Repo
func (o *Repo) Clone() datamodel.Model {
	c := new(Repo)
	c.FromMap(o.ToMap())
	return c
}

// MarshalJSON returns the bytes for marshaling to json
func (o *Repo) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Repo) UnmarshalJSON(data []byte) error {
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
func (o *Repo) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Repo objects are equal
func (o *Repo) IsEqual(other *Repo) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Repo) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"active":         toRepoObject(o.Active, false),
		"customer_id":    toRepoObject(o.CustomerID, false),
		"default_branch": toRepoObject(o.DefaultBranch, false),
		"description":    toRepoObject(o.Description, false),
		"id":             toRepoObject(o.ID, false),
		"language":       toRepoObject(o.Language, false),
		"name":           toRepoObject(o.Name, false),
		"ref_id":         toRepoObject(o.RefID, false),
		"ref_type":       toRepoObject(o.RefType, false),
		"updated_ts":     toRepoObject(o.UpdatedAt, false),
		"url":            toRepoObject(o.URL, false),
		"hashcode":       toRepoObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Repo) FromMap(kv map[string]interface{}) {

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
				o.Active = number.ToBoolAny(nil)
			} else {
				o.Active = number.ToBoolAny(val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CustomerID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["default_branch"].(string); ok {
		o.DefaultBranch = val
	} else {
		if val, ok := kv["default_branch"]; ok {
			if val == nil {
				o.DefaultBranch = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.DefaultBranch = fmt.Sprintf("%v", val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Description = fmt.Sprintf("%v", val)
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

	if val, ok := kv["language"].(string); ok {
		o.Language = val
	} else {
		if val, ok := kv["language"]; ok {
			if val == nil {
				o.Language = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Language = fmt.Sprintf("%v", val)
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
func (o *Repo) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Active)
	args = append(args, o.CustomerID)
	args = append(args, o.DefaultBranch)
	args = append(args, o.Description)
	args = append(args, o.ID)
	args = append(args, o.Language)
	args = append(args, o.Name)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}
