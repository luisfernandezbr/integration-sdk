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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CommitTopic is the default topic name
	CommitTopic datamodel.TopicNameType = "sourcecode_Commit_topic"

	// CommitModelName is the model name
	CommitModelName datamodel.ModelNameType = "sourcecode.Commit"
)

const (
	// CommitAdditionsColumn is the additions column name
	CommitAdditionsColumn = "Additions"
	// CommitAuthorRefIDColumn is the author_ref_id column name
	CommitAuthorRefIDColumn = "AuthorRefID"
	// CommitBlanksColumn is the blanks column name
	CommitBlanksColumn = "Blanks"
	// CommitBranchIdsColumn is the branch_ids column name
	CommitBranchIdsColumn = "BranchIds"
	// CommitCommentsColumn is the comments column name
	CommitCommentsColumn = "Comments"
	// CommitCommitterRefIDColumn is the committer_ref_id column name
	CommitCommitterRefIDColumn = "CommitterRefID"
	// CommitComplexityColumn is the complexity column name
	CommitComplexityColumn = "Complexity"
	// CommitCreatedDateColumn is the created_date column name
	CommitCreatedDateColumn = "CreatedDate"
	// CommitCreatedDateColumnEpochColumn is the epoch column property of the CreatedDate name
	CommitCreatedDateColumnEpochColumn = "CreatedDate.Epoch"
	// CommitCreatedDateColumnOffsetColumn is the offset column property of the CreatedDate name
	CommitCreatedDateColumnOffsetColumn = "CreatedDate.Offset"
	// CommitCreatedDateColumnRfc3339Column is the rfc3339 column property of the CreatedDate name
	CommitCreatedDateColumnRfc3339Column = "CreatedDate.Rfc3339"
	// CommitCustomerIDColumn is the customer_id column name
	CommitCustomerIDColumn = "CustomerID"
	// CommitDeletionsColumn is the deletions column name
	CommitDeletionsColumn = "Deletions"
	// CommitExcludedColumn is the excluded column name
	CommitExcludedColumn = "Excluded"
	// CommitFilesColumn is the files column name
	CommitFilesColumn = "Files"
	// CommitFilesColumnAdditionsColumn is the additions column property of the Files name
	CommitFilesColumnAdditionsColumn = "Files.Additions"
	// CommitFilesColumnBinaryColumn is the binary column property of the Files name
	CommitFilesColumnBinaryColumn = "Files.Binary"
	// CommitFilesColumnBlanksColumn is the blanks column property of the Files name
	CommitFilesColumnBlanksColumn = "Files.Blanks"
	// CommitFilesColumnCommentsColumn is the comments column property of the Files name
	CommitFilesColumnCommentsColumn = "Files.Comments"
	// CommitFilesColumnCommitIDColumn is the commit_id column property of the Files name
	CommitFilesColumnCommitIDColumn = "Files.CommitID"
	// CommitFilesColumnComplexityColumn is the complexity column property of the Files name
	CommitFilesColumnComplexityColumn = "Files.Complexity"
	// CommitFilesColumnCreatedDateColumn is the created_date column property of the Files name
	CommitFilesColumnCreatedDateColumn = "Files.CreatedDate"
	// CommitFilesColumnDeletionsColumn is the deletions column property of the Files name
	CommitFilesColumnDeletionsColumn = "Files.Deletions"
	// CommitFilesColumnExcludedColumn is the excluded column property of the Files name
	CommitFilesColumnExcludedColumn = "Files.Excluded"
	// CommitFilesColumnExcludedReasonColumn is the excluded_reason column property of the Files name
	CommitFilesColumnExcludedReasonColumn = "Files.ExcludedReason"
	// CommitFilesColumnFilenameColumn is the filename column property of the Files name
	CommitFilesColumnFilenameColumn = "Files.Filename"
	// CommitFilesColumnLanguageColumn is the language column property of the Files name
	CommitFilesColumnLanguageColumn = "Files.Language"
	// CommitFilesColumnLicenseColumn is the license column property of the Files name
	CommitFilesColumnLicenseColumn = "Files.License"
	// CommitFilesColumnLicenseConfidenceColumn is the license_confidence column property of the Files name
	CommitFilesColumnLicenseConfidenceColumn = "Files.LicenseConfidence"
	// CommitFilesColumnLocColumn is the loc column property of the Files name
	CommitFilesColumnLocColumn = "Files.Loc"
	// CommitFilesColumnOrdinalColumn is the ordinal column property of the Files name
	CommitFilesColumnOrdinalColumn = "Files.Ordinal"
	// CommitFilesColumnRenamedColumn is the renamed column property of the Files name
	CommitFilesColumnRenamedColumn = "Files.Renamed"
	// CommitFilesColumnRenamedFromColumn is the renamed_from column property of the Files name
	CommitFilesColumnRenamedFromColumn = "Files.RenamedFrom"
	// CommitFilesColumnRenamedToColumn is the renamed_to column property of the Files name
	CommitFilesColumnRenamedToColumn = "Files.RenamedTo"
	// CommitFilesColumnRepoIDColumn is the repo_id column property of the Files name
	CommitFilesColumnRepoIDColumn = "Files.RepoID"
	// CommitFilesColumnSizeColumn is the size column property of the Files name
	CommitFilesColumnSizeColumn = "Files.Size"
	// CommitFilesColumnSlocColumn is the sloc column property of the Files name
	CommitFilesColumnSlocColumn = "Files.Sloc"
	// CommitFilesColumnStatusColumn is the status column property of the Files name
	CommitFilesColumnStatusColumn = "Files.Status"
	// CommitFilesChangedColumn is the files_changed column name
	CommitFilesChangedColumn = "FilesChanged"
	// CommitGpgSignedColumn is the gpg_signed column name
	CommitGpgSignedColumn = "GpgSigned"
	// CommitIDColumn is the id column name
	CommitIDColumn = "ID"
	// CommitLocColumn is the loc column name
	CommitLocColumn = "Loc"
	// CommitMessageColumn is the message column name
	CommitMessageColumn = "Message"
	// CommitOrdinalColumn is the ordinal column name
	CommitOrdinalColumn = "Ordinal"
	// CommitRefIDColumn is the ref_id column name
	CommitRefIDColumn = "RefID"
	// CommitRefTypeColumn is the ref_type column name
	CommitRefTypeColumn = "RefType"
	// CommitRepoIDColumn is the repo_id column name
	CommitRepoIDColumn = "RepoID"
	// CommitShaColumn is the sha column name
	CommitShaColumn = "Sha"
	// CommitSizeColumn is the size column name
	CommitSizeColumn = "Size"
	// CommitSlocColumn is the sloc column name
	CommitSlocColumn = "Sloc"
	// CommitUpdatedAtColumn is the updated_ts column name
	CommitUpdatedAtColumn = "UpdatedAt"
	// CommitURLColumn is the url column name
	CommitURLColumn = "URL"
)

// CommitCreatedDate represents the object structure for created_date
type CommitCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCommitCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CommitCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CommitCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCommitCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCommitCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCommitCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *CommitCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitCreatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// CommitFilesCreatedDate represents the object structure for created_date
type CommitFilesCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toCommitFilesCreatedDateObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CommitFilesCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CommitFilesCreatedDate) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toCommitFilesCreatedDateObject(o.Epoch, false),
		// Offset the timezone offset from GMT
		"offset": toCommitFilesCreatedDateObject(o.Offset, false),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toCommitFilesCreatedDateObject(o.Rfc3339, false),
	}
}

func (o *CommitFilesCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitFilesCreatedDate) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["epoch"].(int64); ok {
		o.Epoch = val
	} else {
		if val, ok := kv["epoch"]; ok {
			if val == nil {
				o.Epoch = number.ToInt64Any(nil)
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
				o.Offset = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Rfc3339 = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// CommitFiles represents the object structure for files
type CommitFiles struct {
	// Additions the number of additions for the commit file
	Additions int64 `json:"additions" yaml:"additions" faker:"-"`
	// Binary indicates if the file was detected to be a binary file
	Binary bool `json:"binary" yaml:"binary" faker:"-"`
	// Blanks the number of blank lines in the file
	Blanks int64 `json:"blanks" yaml:"blanks" faker:"-"`
	// Comments the number of comment lines in the file
	Comments int64 `json:"comments" yaml:"comments" faker:"-"`
	// CommitID the unique id for the commit
	CommitID string `json:"commit_id" yaml:"commit_id" faker:"-"`
	// Complexity the complexity value for the file change
	Complexity int64 `json:"complexity" yaml:"complexity" faker:"-"`
	// CreatedDate the timestamp in UTC that the commit was created
	CreatedDate CommitFilesCreatedDate `json:"created_date" yaml:"created_date" faker:"-"`
	// Deletions the number of deletions for the commit file
	Deletions int64 `json:"deletions" yaml:"deletions" faker:"-"`
	// Excluded if the file was excluded from processing
	Excluded bool `json:"excluded" yaml:"excluded" faker:"-"`
	// ExcludedReason if the file was excluded, the reason
	ExcludedReason string `json:"excluded_reason" yaml:"excluded_reason" faker:"-"`
	// Filename the filename
	Filename string `json:"filename" yaml:"filename" faker:"-"`
	// Language the language that was detected for the file
	Language string `json:"language" yaml:"language" faker:"-"`
	// License the license which was detected for the file
	License string `json:"license" yaml:"license" faker:"-"`
	// LicenseConfidence the license confidence from the detection engine
	LicenseConfidence float64 `json:"license_confidence" yaml:"license_confidence" faker:"-"`
	// Loc the number of lines in the file
	Loc int64 `json:"loc" yaml:"loc" faker:"-"`
	// Ordinal the order value for the file in the change set
	Ordinal int64 `json:"ordinal" yaml:"ordinal" faker:"-"`
	// Renamed if the file was renamed
	Renamed bool `json:"renamed" yaml:"renamed" faker:"-"`
	// RenamedFrom the original file name
	RenamedFrom string `json:"renamed_from" yaml:"renamed_from" faker:"-"`
	// RenamedTo the final file name
	RenamedTo string `json:"renamed_to" yaml:"renamed_to" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
	// Size the size of the file
	Size int64 `json:"size" yaml:"size" faker:"-"`
	// Sloc the number of source lines in the file
	Sloc int64 `json:"sloc" yaml:"sloc" faker:"-"`
	// Status the status of the change
	Status string `json:"status" yaml:"status" faker:"-"`
}

func toCommitFilesObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *CommitFiles:
		return v.ToMap()

	case CommitFilesCreatedDate:
		return v.ToMap()

	default:
		return o
	}
}

func (o *CommitFiles) ToMap() map[string]interface{} {
	o.setDefaults(true)
	return map[string]interface{}{
		// Additions the number of additions for the commit file
		"additions": toCommitFilesObject(o.Additions, false),
		// Binary indicates if the file was detected to be a binary file
		"binary": toCommitFilesObject(o.Binary, false),
		// Blanks the number of blank lines in the file
		"blanks": toCommitFilesObject(o.Blanks, false),
		// Comments the number of comment lines in the file
		"comments": toCommitFilesObject(o.Comments, false),
		// CommitID the unique id for the commit
		"commit_id": toCommitFilesObject(o.CommitID, false),
		// Complexity the complexity value for the file change
		"complexity": toCommitFilesObject(o.Complexity, false),
		// CreatedDate the timestamp in UTC that the commit was created
		"created_date": toCommitFilesObject(o.CreatedDate, false),
		// Deletions the number of deletions for the commit file
		"deletions": toCommitFilesObject(o.Deletions, false),
		// Excluded if the file was excluded from processing
		"excluded": toCommitFilesObject(o.Excluded, false),
		// ExcludedReason if the file was excluded, the reason
		"excluded_reason": toCommitFilesObject(o.ExcludedReason, false),
		// Filename the filename
		"filename": toCommitFilesObject(o.Filename, false),
		// Language the language that was detected for the file
		"language": toCommitFilesObject(o.Language, false),
		// License the license which was detected for the file
		"license": toCommitFilesObject(o.License, false),
		// LicenseConfidence the license confidence from the detection engine
		"license_confidence": toCommitFilesObject(o.LicenseConfidence, false),
		// Loc the number of lines in the file
		"loc": toCommitFilesObject(o.Loc, false),
		// Ordinal the order value for the file in the change set
		"ordinal": toCommitFilesObject(o.Ordinal, false),
		// Renamed if the file was renamed
		"renamed": toCommitFilesObject(o.Renamed, false),
		// RenamedFrom the original file name
		"renamed_from": toCommitFilesObject(o.RenamedFrom, false),
		// RenamedTo the final file name
		"renamed_to": toCommitFilesObject(o.RenamedTo, false),
		// RepoID the unique id for the repo
		"repo_id": toCommitFilesObject(o.RepoID, false),
		// Size the size of the file
		"size": toCommitFilesObject(o.Size, false),
		// Sloc the number of source lines in the file
		"sloc": toCommitFilesObject(o.Sloc, false),
		// Status the status of the change
		"status": toCommitFilesObject(o.Status, false),
	}
}

func (o *CommitFiles) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *CommitFiles) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["additions"].(int64); ok {
		o.Additions = val
	} else {
		if val, ok := kv["additions"]; ok {
			if val == nil {
				o.Additions = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Additions = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["binary"].(bool); ok {
		o.Binary = val
	} else {
		if val, ok := kv["binary"]; ok {
			if val == nil {
				o.Binary = number.ToBoolAny(nil)
			} else {
				o.Binary = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		if val, ok := kv["blanks"]; ok {
			if val == nil {
				o.Blanks = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Blanks = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		if val, ok := kv["comments"]; ok {
			if val == nil {
				o.Comments = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Comments = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["commit_id"].(string); ok {
		o.CommitID = val
	} else {
		if val, ok := kv["commit_id"]; ok {
			if val == nil {
				o.CommitID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CommitID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		if val, ok := kv["complexity"]; ok {
			if val == nil {
				o.Complexity = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Complexity = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(CommitFilesCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*CommitFilesCreatedDate); ok {
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

	if val, ok := kv["deletions"].(int64); ok {
		o.Deletions = val
	} else {
		if val, ok := kv["deletions"]; ok {
			if val == nil {
				o.Deletions = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Deletions = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		if val, ok := kv["excluded"]; ok {
			if val == nil {
				o.Excluded = number.ToBoolAny(nil)
			} else {
				o.Excluded = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["excluded_reason"].(string); ok {
		o.ExcludedReason = val
	} else {
		if val, ok := kv["excluded_reason"]; ok {
			if val == nil {
				o.ExcludedReason = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.ExcludedReason = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["filename"].(string); ok {
		o.Filename = val
	} else {
		if val, ok := kv["filename"]; ok {
			if val == nil {
				o.Filename = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Filename = fmt.Sprintf("%v", val)
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

	if val, ok := kv["license"].(string); ok {
		o.License = val
	} else {
		if val, ok := kv["license"]; ok {
			if val == nil {
				o.License = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.License = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["license_confidence"].(float64); ok {
		o.LicenseConfidence = val
	} else {
		if val, ok := kv["license_confidence"]; ok {
			if val == nil {
				o.LicenseConfidence = number.ToFloat64Any(nil)
			} else {
				o.LicenseConfidence = number.ToFloat64Any(val)
			}
		}
	}

	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		if val, ok := kv["loc"]; ok {
			if val == nil {
				o.Loc = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Loc = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		if val, ok := kv["ordinal"]; ok {
			if val == nil {
				o.Ordinal = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Ordinal = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["renamed"].(bool); ok {
		o.Renamed = val
	} else {
		if val, ok := kv["renamed"]; ok {
			if val == nil {
				o.Renamed = number.ToBoolAny(nil)
			} else {
				o.Renamed = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["renamed_from"].(string); ok {
		o.RenamedFrom = val
	} else {
		if val, ok := kv["renamed_from"]; ok {
			if val == nil {
				o.RenamedFrom = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RenamedFrom = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["renamed_to"].(string); ok {
		o.RenamedTo = val
	} else {
		if val, ok := kv["renamed_to"]; ok {
			if val == nil {
				o.RenamedTo = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RenamedTo = fmt.Sprintf("%v", val)
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

	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		if val, ok := kv["size"]; ok {
			if val == nil {
				o.Size = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Size = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		if val, ok := kv["sloc"]; ok {
			if val == nil {
				o.Sloc = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Sloc = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["status"].(string); ok {
		o.Status = val
	} else {
		if val, ok := kv["status"]; ok {
			if val == nil {
				o.Status = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Status = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Commit the commit is a specific change in a repo
type Commit struct {
	// Additions the number of additions for the commit
	Additions int64 `json:"additions" yaml:"additions" faker:"-"`
	// AuthorRefID the author ref_id in the source system
	AuthorRefID string `json:"author_ref_id" yaml:"author_ref_id" faker:"-"`
	// Blanks the number of blank lines in the commit
	Blanks int64 `json:"blanks" yaml:"blanks" faker:"-"`
	// BranchIds the branches that the commit was a part of
	BranchIds []string `json:"branch_ids" yaml:"branch_ids" faker:"-"`
	// Comments the number of comment lines in the commit
	Comments int64 `json:"comments" yaml:"comments" faker:"-"`
	// CommitterRefID the committer ref_id in the source system
	CommitterRefID string `json:"committer_ref_id" yaml:"committer_ref_id" faker:"-"`
	// Complexity the complexity value for the change
	Complexity int64 `json:"complexity" yaml:"complexity" faker:"-"`
	// CreatedDate date when the commit was created
	CreatedDate CommitCreatedDate `json:"created_date" yaml:"created_date" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" yaml:"customer_id" faker:"-"`
	// Deletions the number of deletions for the commit
	Deletions int64 `json:"deletions" yaml:"deletions" faker:"-"`
	// Excluded if the commit was excluded
	Excluded bool `json:"excluded" yaml:"excluded" faker:"-"`
	// Files the files touched by this commit
	Files []CommitFiles `json:"files" yaml:"files" faker:"-"`
	// FilesChanged the number of files changed for the commit
	FilesChanged int64 `json:"files_changed" yaml:"files_changed" faker:"-"`
	// GpgSigned if the commit was signed
	GpgSigned bool `json:"gpg_signed" yaml:"gpg_signed" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Loc the number of lines in the commit
	Loc int64 `json:"loc" yaml:"loc" faker:"-"`
	// Message the commit message
	Message string `json:"message" yaml:"message" faker:"commit_message"`
	// Ordinal the order of the commit in the commit stream
	Ordinal int64 `json:"ordinal" yaml:"ordinal" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" yaml:"ref_type" faker:"-"`
	// RepoID the unique id for the repo
	RepoID string `json:"repo_id" yaml:"repo_id" faker:"-"`
	// Sha the unique sha for the commit
	Sha string `json:"sha" yaml:"sha" faker:"sha"`
	// Size the size of all files in the commit
	Size int64 `json:"size" yaml:"size" faker:"-"`
	// Sloc the number of source lines in the commit
	Sloc int64 `json:"sloc" yaml:"sloc" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" yaml:"updated_ts" faker:"-"`
	// URL the url to the commit detail
	URL string `json:"url" yaml:"url" faker:"url"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*Commit)(nil)

// ensure that this type implements the streamed data model interface
var _ datamodel.StreamedModel = (*Commit)(nil)

func toCommitObject(o interface{}, isoptional bool) interface{} {
	switch v := o.(type) {
	case *Commit:
		return v.ToMap()

	case CommitCreatedDate:
		return v.ToMap()

	case []CommitFiles:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap())
		}
		return arr

	default:
		return o
	}
}

// String returns a string representation of Commit
func (o *Commit) String() string {
	return fmt.Sprintf("sourcecode.Commit<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *Commit) GetTopicName() datamodel.TopicNameType {
	return CommitTopic
}

// GetStreamName returns the name of the stream
func (o *Commit) GetStreamName() string {
	return ""
}

// GetTableName returns the name of the table
func (o *Commit) GetTableName() string {
	return ""
}

// GetModelName returns the name of the model
func (o *Commit) GetModelName() datamodel.ModelNameType {
	return CommitModelName
}

// NewCommitID provides a template for generating an ID field for Commit
func NewCommitID(customerID string, refID string, refType string, RepoID string) string {
	return hash.Values(customerID, refID, refType, RepoID)
}

func (o *Commit) setDefaults(frommap bool) {
	if o.BranchIds == nil {
		o.BranchIds = make([]string, 0)
	}
	if o.Files == nil {
		o.Files = make([]CommitFiles, 0)
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
func (o *Commit) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *Commit) GetTopicKey() string {
	var i interface{} = o.RepoID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *Commit) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for Commit")
}

// GetRefID returns the RefID for the object
func (o *Commit) GetRefID() string {
	o.RefID = o.Sha
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *Commit) IsMaterialized() bool {
	return false
}

// IsMutable returns true if the model is mutable
func (o *Commit) IsMutable() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *Commit) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *Commit) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *Commit) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = CommitModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *Commit) GetTopicConfig() *datamodel.ModelTopicConfig {
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
func (o *Commit) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of Commit
func (o *Commit) Clone() datamodel.Model {
	c := new(Commit)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *Commit) Anon() datamodel.Model {
	c := new(Commit)
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
func (o *Commit) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *Commit) UnmarshalJSON(data []byte) error {
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
func (o *Commit) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two Commit objects are equal
func (o *Commit) IsEqual(other *Commit) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *Commit) ToMap() map[string]interface{} {
	o.setDefaults(false)
	return map[string]interface{}{
		"additions":        toCommitObject(o.Additions, false),
		"author_ref_id":    toCommitObject(o.AuthorRefID, false),
		"blanks":           toCommitObject(o.Blanks, false),
		"branch_ids":       toCommitObject(o.BranchIds, false),
		"comments":         toCommitObject(o.Comments, false),
		"committer_ref_id": toCommitObject(o.CommitterRefID, false),
		"complexity":       toCommitObject(o.Complexity, false),
		"created_date":     toCommitObject(o.CreatedDate, false),
		"customer_id":      toCommitObject(o.CustomerID, false),
		"deletions":        toCommitObject(o.Deletions, false),
		"excluded":         toCommitObject(o.Excluded, false),
		"files":            toCommitObject(o.Files, false),
		"files_changed":    toCommitObject(o.FilesChanged, false),
		"gpg_signed":       toCommitObject(o.GpgSigned, false),
		"id":               toCommitObject(o.ID, false),
		"loc":              toCommitObject(o.Loc, false),
		"message":          toCommitObject(o.Message, false),
		"ordinal":          toCommitObject(o.Ordinal, false),
		"ref_id":           toCommitObject(o.RefID, false),
		"ref_type":         toCommitObject(o.RefType, false),
		"repo_id":          toCommitObject(o.RepoID, false),
		"sha":              toCommitObject(o.Sha, false),
		"size":             toCommitObject(o.Size, false),
		"sloc":             toCommitObject(o.Sloc, false),
		"updated_ts":       toCommitObject(o.UpdatedAt, false),
		"url":              toCommitObject(o.URL, false),
		"hashcode":         toCommitObject(o.Hashcode, false),
	}
}

// FromMap attempts to load data into object from a map
func (o *Commit) FromMap(kv map[string]interface{}) {

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
				o.Additions = number.ToInt64Any(nil)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.AuthorRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["blanks"].(int64); ok {
		o.Blanks = val
	} else {
		if val, ok := kv["blanks"]; ok {
			if val == nil {
				o.Blanks = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Blanks = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["branch_ids"]; ok {
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
								panic("unsupported type for branch_ids field entry: " + reflect.TypeOf(ae).String())
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
								panic("unsupported type for branch_ids field entry: " + reflect.TypeOf(ae).String())
							}
							na = append(na, av)
						}
					}
				} else {
					fmt.Println(reflect.TypeOf(val).String())
					panic("unsupported type for branch_ids field")
				}
			}
			o.BranchIds = na
		}
	}
	if o.BranchIds == nil {
		o.BranchIds = make([]string, 0)
	}

	if val, ok := kv["comments"].(int64); ok {
		o.Comments = val
	} else {
		if val, ok := kv["comments"]; ok {
			if val == nil {
				o.Comments = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Comments = number.ToInt64Any(val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.CommitterRefID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["complexity"].(int64); ok {
		o.Complexity = val
	} else {
		if val, ok := kv["complexity"]; ok {
			if val == nil {
				o.Complexity = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Complexity = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(CommitCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*CommitCreatedDate); ok {
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
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
				o.Deletions = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Deletions = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["excluded"].(bool); ok {
		o.Excluded = val
	} else {
		if val, ok := kv["excluded"]; ok {
			if val == nil {
				o.Excluded = number.ToBoolAny(nil)
			} else {
				o.Excluded = number.ToBoolAny(val)
			}
		}
	}

	if o == nil {

		o.Files = make([]CommitFiles, 0)

	}
	if val, ok := kv["files"]; ok {
		if sv, ok := val.([]CommitFiles); ok {
			o.Files = sv
		} else if sp, ok := val.([]*CommitFiles); ok {
			o.Files = o.Files[:0]
			for _, e := range sp {
				o.Files = append(o.Files, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(CommitFiles); ok {
					o.Files = append(o.Files, av)
				} else if av, ok := ae.(primitive.M); ok {
					var fm CommitFiles
					fm.FromMap(av)
					o.Files = append(o.Files, fm)
				} else {
					b, _ := json.Marshal(ae)
					var av CommitFiles
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for files field entry: " + reflect.TypeOf(ae).String())
					}
					o.Files = append(o.Files, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(CommitFiles); ok {
					o.Files = append(o.Files, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm CommitFiles
					fm.FromMap(r)
					o.Files = append(o.Files, fm)
				} else if r, ok := item.(primitive.M); ok {
					fm := CommitFiles{}
					fm.FromMap(r)
					o.Files = append(o.Files, fm)
				}
			}
		} else {
			arr := reflect.ValueOf(val)
			if arr.Kind() == reflect.Slice {
				for i := 0; i < arr.Len(); i++ {
					item := arr.Index(i)
					if item.CanAddr() {
						v := item.Addr().MethodByName("ToMap")
						if !v.IsNil() {
							m := v.Call([]reflect.Value{})
							var fm CommitFiles
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Files = append(o.Files, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["files_changed"].(int64); ok {
		o.FilesChanged = val
	} else {
		if val, ok := kv["files_changed"]; ok {
			if val == nil {
				o.FilesChanged = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FilesChanged = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["gpg_signed"].(bool); ok {
		o.GpgSigned = val
	} else {
		if val, ok := kv["gpg_signed"]; ok {
			if val == nil {
				o.GpgSigned = number.ToBoolAny(nil)
			} else {
				o.GpgSigned = number.ToBoolAny(val)
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

	if val, ok := kv["loc"].(int64); ok {
		o.Loc = val
	} else {
		if val, ok := kv["loc"]; ok {
			if val == nil {
				o.Loc = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Loc = number.ToInt64Any(val)
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
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Message = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["ordinal"].(int64); ok {
		o.Ordinal = val
	} else {
		if val, ok := kv["ordinal"]; ok {
			if val == nil {
				o.Ordinal = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Ordinal = number.ToInt64Any(val)
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

	if val, ok := kv["sha"].(string); ok {
		o.Sha = val
	} else {
		if val, ok := kv["sha"]; ok {
			if val == nil {
				o.Sha = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Sha = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["size"].(int64); ok {
		o.Size = val
	} else {
		if val, ok := kv["size"]; ok {
			if val == nil {
				o.Size = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Size = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["sloc"].(int64); ok {
		o.Sloc = val
	} else {
		if val, ok := kv["sloc"]; ok {
			if val == nil {
				o.Sloc = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Sloc = number.ToInt64Any(val)
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
func (o *Commit) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Additions)
	args = append(args, o.AuthorRefID)
	args = append(args, o.Blanks)
	args = append(args, o.BranchIds)
	args = append(args, o.Comments)
	args = append(args, o.CommitterRefID)
	args = append(args, o.Complexity)
	args = append(args, o.CreatedDate)
	args = append(args, o.CustomerID)
	args = append(args, o.Deletions)
	args = append(args, o.Excluded)
	args = append(args, o.Files)
	args = append(args, o.FilesChanged)
	args = append(args, o.GpgSigned)
	args = append(args, o.ID)
	args = append(args, o.Loc)
	args = append(args, o.Message)
	args = append(args, o.Ordinal)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.RepoID)
	args = append(args, o.Sha)
	args = append(args, o.Size)
	args = append(args, o.Sloc)
	args = append(args, o.UpdatedAt)
	args = append(args, o.URL)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *Commit) GetEventAPIConfig() datamodel.EventAPIConfig {
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
