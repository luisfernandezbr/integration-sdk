// DO NOT EDIT -- generated code

// Package agent - the agent communicates with pinpoint cloud to send integration data
package agent

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"
	"time"

	"github.com/bxcodec/faker"
	"github.com/linkedin/goavro"
	"github.com/pinpt/go-common/datamodel"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/eventing"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	pjson "github.com/pinpt/go-common/json"
	"github.com/pinpt/go-common/number"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// RepoResponseTopic is the default topic name
	RepoResponseTopic datamodel.TopicNameType = "agent_RepoResponse_topic"

	// RepoResponseStream is the default stream name
	RepoResponseStream datamodel.TopicNameType = "agent_RepoResponse_stream"

	// RepoResponseTable is the default table name
	RepoResponseTable datamodel.TopicNameType = "agent_reporesponse"

	// RepoResponseModelName is the model name
	RepoResponseModelName datamodel.ModelNameType = "agent.RepoResponse"
)

const (
	// RepoResponseArchitectureColumn is the architecture column name
	RepoResponseArchitectureColumn = "architecture"
	// RepoResponseCustomerIDColumn is the customer_id column name
	RepoResponseCustomerIDColumn = "customer_id"
	// RepoResponseDataColumn is the data column name
	RepoResponseDataColumn = "data"
	// RepoResponseDistroColumn is the distro column name
	RepoResponseDistroColumn = "distro"
	// RepoResponseErrorColumn is the error column name
	RepoResponseErrorColumn = "error"
	// RepoResponseEventDateColumn is the event_date column name
	RepoResponseEventDateColumn = "event_date"
	// RepoResponseEventDateColumnEpochColumn is the epoch column property of the EventDate name
	RepoResponseEventDateColumnEpochColumn = "event_date->epoch"
	// RepoResponseEventDateColumnOffsetColumn is the offset column property of the EventDate name
	RepoResponseEventDateColumnOffsetColumn = "event_date->offset"
	// RepoResponseEventDateColumnRfc3339Column is the rfc3339 column property of the EventDate name
	RepoResponseEventDateColumnRfc3339Column = "event_date->rfc3339"
	// RepoResponseFreeSpaceColumn is the free_space column name
	RepoResponseFreeSpaceColumn = "free_space"
	// RepoResponseGoVersionColumn is the go_version column name
	RepoResponseGoVersionColumn = "go_version"
	// RepoResponseHostnameColumn is the hostname column name
	RepoResponseHostnameColumn = "hostname"
	// RepoResponseIDColumn is the id column name
	RepoResponseIDColumn = "id"
	// RepoResponseIntegrationIDColumn is the integration_id column name
	RepoResponseIntegrationIDColumn = "integration_id"
	// RepoResponseLastExportDateColumn is the last_export_date column name
	RepoResponseLastExportDateColumn = "last_export_date"
	// RepoResponseLastExportDateColumnEpochColumn is the epoch column property of the LastExportDate name
	RepoResponseLastExportDateColumnEpochColumn = "last_export_date->epoch"
	// RepoResponseLastExportDateColumnOffsetColumn is the offset column property of the LastExportDate name
	RepoResponseLastExportDateColumnOffsetColumn = "last_export_date->offset"
	// RepoResponseLastExportDateColumnRfc3339Column is the rfc3339 column property of the LastExportDate name
	RepoResponseLastExportDateColumnRfc3339Column = "last_export_date->rfc3339"
	// RepoResponseMemoryColumn is the memory column name
	RepoResponseMemoryColumn = "memory"
	// RepoResponseMessageColumn is the message column name
	RepoResponseMessageColumn = "message"
	// RepoResponseNumCPUColumn is the num_cpu column name
	RepoResponseNumCPUColumn = "num_cpu"
	// RepoResponseOSColumn is the os column name
	RepoResponseOSColumn = "os"
	// RepoResponseRefIDColumn is the ref_id column name
	RepoResponseRefIDColumn = "ref_id"
	// RepoResponseRefTypeColumn is the ref_type column name
	RepoResponseRefTypeColumn = "ref_type"
	// RepoResponseReposColumn is the repos column name
	RepoResponseReposColumn = "repos"
	// RepoResponseReposColumnActiveColumn is the active column property of the Repos name
	RepoResponseReposColumnActiveColumn = "repos->active"
	// RepoResponseReposColumnCreatedDateColumn is the created_date column property of the Repos name
	RepoResponseReposColumnCreatedDateColumn = "repos->created_date"
	// RepoResponseReposColumnDescriptionColumn is the description column property of the Repos name
	RepoResponseReposColumnDescriptionColumn = "repos->description"
	// RepoResponseReposColumnLanguageColumn is the language column property of the Repos name
	RepoResponseReposColumnLanguageColumn = "repos->language"
	// RepoResponseReposColumnLastCommitColumn is the last_commit column property of the Repos name
	RepoResponseReposColumnLastCommitColumn = "repos->last_commit"
	// RepoResponseReposColumnNameColumn is the name column property of the Repos name
	RepoResponseReposColumnNameColumn = "repos->name"
	// RepoResponseReposColumnRefIDColumn is the ref_id column property of the Repos name
	RepoResponseReposColumnRefIDColumn = "repos->ref_id"
	// RepoResponseReposColumnRefTypeColumn is the ref_type column property of the Repos name
	RepoResponseReposColumnRefTypeColumn = "repos->ref_type"
	// RepoResponseRequestIDColumn is the request_id column name
	RepoResponseRequestIDColumn = "request_id"
	// RepoResponseSuccessColumn is the success column name
	RepoResponseSuccessColumn = "success"
	// RepoResponseSystemIDColumn is the system_id column name
	RepoResponseSystemIDColumn = "system_id"
	// RepoResponseTypeColumn is the type column name
	RepoResponseTypeColumn = "type"
	// RepoResponseUpdatedAtColumn is the updated_ts column name
	RepoResponseUpdatedAtColumn = "updated_ts"
	// RepoResponseUptimeColumn is the uptime column name
	RepoResponseUptimeColumn = "uptime"
	// RepoResponseUUIDColumn is the uuid column name
	RepoResponseUUIDColumn = "uuid"
	// RepoResponseVersionColumn is the version column name
	RepoResponseVersionColumn = "version"
)

// RepoResponseEventDate represents the object structure for event_date
type RepoResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseEventDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseEventDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseEventDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseEventDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseEventDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseEventDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseEventDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *RepoResponseEventDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseEventDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseLastExportDate represents the object structure for last_export_date
type RepoResponseLastExportDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseLastExportDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseLastExportDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseLastExportDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseLastExportDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseLastExportDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseLastExportDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseLastExportDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *RepoResponseLastExportDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseLastExportDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseReposCreatedDate represents the object structure for created_date
type RepoResponseReposCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseReposCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseReposCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseReposCreatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseReposCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseReposCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseReposCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseReposCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *RepoResponseReposCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseReposCreatedDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseReposLastCommitCreatedDate represents the object structure for created_date
type RepoResponseReposLastCommitCreatedDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func toRepoResponseReposLastCommitCreatedDateObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseReposLastCommitCreatedDateObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseReposLastCommitCreatedDate:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseReposLastCommitCreatedDate) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": toRepoResponseReposLastCommitCreatedDateObject(o.Epoch, isavro, false, "long"),
		// Offset the timezone offset from GMT
		"offset": toRepoResponseReposLastCommitCreatedDateObject(o.Offset, isavro, false, "long"),
		// Rfc3339 the date in RFC3339 format
		"rfc3339": toRepoResponseReposLastCommitCreatedDateObject(o.Rfc3339, isavro, false, "string"),
	}
}

func (o *RepoResponseReposLastCommitCreatedDate) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseReposLastCommitCreatedDate) FromMap(kv map[string]interface{}) {

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

// RepoResponseReposLastCommitAuthor represents the object structure for author
type RepoResponseReposLastCommitAuthor struct {
	// Name the author name
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// Email the email of the author
	Email string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// AvatarURL the avatar_url for the author
	AvatarURL string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
}

func toRepoResponseReposLastCommitAuthorObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseReposLastCommitAuthorObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseReposLastCommitAuthor:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseReposLastCommitAuthor) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Name the author name
		"name": toRepoResponseReposLastCommitAuthorObject(o.Name, isavro, false, "string"),
		// Email the email of the author
		"email": toRepoResponseReposLastCommitAuthorObject(o.Email, isavro, false, "string"),
		// AvatarURL the avatar_url for the author
		"avatar_url": toRepoResponseReposLastCommitAuthorObject(o.AvatarURL, isavro, false, "string"),
	}
}

func (o *RepoResponseReposLastCommitAuthor) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseReposLastCommitAuthor) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["email"].(string); ok {
		o.Email = val
	} else {
		if val, ok := kv["email"]; ok {
			if val == nil {
				o.Email = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Email = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["avatar_url"].(string); ok {
		o.AvatarURL = val
	} else {
		if val, ok := kv["avatar_url"]; ok {
			if val == nil {
				o.AvatarURL = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.AvatarURL = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// RepoResponseReposLastCommit represents the object structure for last_commit
type RepoResponseReposLastCommit struct {
	// CommitID the id of the latest commit
	CommitID string `json:"commit_id" bson:"commit_id" yaml:"commit_id" faker:"sha"`
	// URL the url of the lastest commit
	URL string `json:"url" bson:"url" yaml:"url" faker:"url"`
	// Message the commit message of the latest commit
	Message string `json:"message" bson:"message" yaml:"message" faker:"commit_message"`
	// CreatedDate the timestamp of the latest commit
	CreatedDate RepoResponseReposLastCommitCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// Author the author of the latest commit
	Author RepoResponseReposLastCommitAuthor `json:"author" bson:"author" yaml:"author" faker:"-"`
}

func toRepoResponseReposLastCommitObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseReposLastCommitObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseReposLastCommit:
		return v.ToMap(isavro)

	case RepoResponseReposLastCommitCreatedDate:
		return v.ToMap(isavro)

	case RepoResponseReposLastCommitAuthor:
		return v.ToMap(isavro)
	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseReposLastCommit) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// CommitID the id of the latest commit
		"commit_id": toRepoResponseReposLastCommitObject(o.CommitID, isavro, false, "string"),
		// URL the url of the lastest commit
		"url": toRepoResponseReposLastCommitObject(o.URL, isavro, false, "string"),
		// Message the commit message of the latest commit
		"message": toRepoResponseReposLastCommitObject(o.Message, isavro, false, "string"),
		// CreatedDate the timestamp of the latest commit
		"created_date": toRepoResponseReposLastCommitObject(o.CreatedDate, isavro, false, "created_date"),
		// Author the author of the latest commit
		"author": toRepoResponseReposLastCommitObject(o.Author, isavro, false, "author"),
	}
}

func (o *RepoResponseReposLastCommit) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseReposLastCommit) FromMap(kv map[string]interface{}) {

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseReposLastCommitCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*RepoResponseReposLastCommitCreatedDate); ok {
			// struct pointer
			o.CreatedDate = *sp
		}
	} else {
		o.CreatedDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["author"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.Author.FromMap(kv)
		} else if sv, ok := val.(RepoResponseReposLastCommitAuthor); ok {
			// struct
			o.Author = sv
		} else if sp, ok := val.(*RepoResponseReposLastCommitAuthor); ok {
			// struct pointer
			o.Author = *sp
		}
	} else {
		o.Author.FromMap(map[string]interface{}{})
	}

	o.setDefaults(false)
}

// RepoResponseRepos represents the object structure for repos
type RepoResponseRepos struct {
	// Active the status of the repo determined by an Admin
	Active bool `json:"active" bson:"active" yaml:"active" faker:"-"`
	// CreatedDate the creation date
	CreatedDate RepoResponseReposCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// Description the description of the repository
	Description string `json:"description" bson:"description" yaml:"description" faker:"sentence"`
	// Language the programming language defined for the repository
	Language string `json:"language" bson:"language" yaml:"language" faker:"-"`
	// LastCommit the most recent commit to the repo
	LastCommit RepoResponseReposLastCommit `json:"last_commit" bson:"last_commit" yaml:"last_commit" faker:"-"`
	// Name the name of the repository
	Name string `json:"name" bson:"name" yaml:"name" faker:"repo"`
	// RefID the repo ID
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the record type
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
}

func toRepoResponseReposObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseReposObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponseRepos:
		return v.ToMap(isavro)

	case RepoResponseReposCreatedDate:
		return v.ToMap(isavro)

	case RepoResponseReposLastCommit:
		return v.ToMap(isavro)

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

func (o *RepoResponseRepos) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	o.setDefaults(true)
	return map[string]interface{}{
		// Active the status of the repo determined by an Admin
		"active": toRepoResponseReposObject(o.Active, isavro, false, "boolean"),
		// CreatedDate the creation date
		"created_date": toRepoResponseReposObject(o.CreatedDate, isavro, false, "created_date"),
		// Description the description of the repository
		"description": toRepoResponseReposObject(o.Description, isavro, false, "string"),
		// Language the programming language defined for the repository
		"language": toRepoResponseReposObject(o.Language, isavro, false, "string"),
		// LastCommit the most recent commit to the repo
		"last_commit": toRepoResponseReposObject(o.LastCommit, isavro, false, "last_commit"),
		// Name the name of the repository
		"name": toRepoResponseReposObject(o.Name, isavro, false, "string"),
		// RefID the repo ID
		"ref_id": toRepoResponseReposObject(o.RefID, isavro, false, "string"),
		// RefType the record type
		"ref_type": toRepoResponseReposObject(o.RefType, isavro, false, "string"),
	}
}

func (o *RepoResponseRepos) setDefaults(frommap bool) {

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponseRepos) FromMap(kv map[string]interface{}) {

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

	if val, ok := kv["created_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.CreatedDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseReposCreatedDate); ok {
			// struct
			o.CreatedDate = sv
		} else if sp, ok := val.(*RepoResponseReposCreatedDate); ok {
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

	if val, ok := kv["last_commit"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastCommit.FromMap(kv)
		} else if sv, ok := val.(RepoResponseReposLastCommit); ok {
			// struct
			o.LastCommit = sv
		} else if sp, ok := val.(*RepoResponseReposLastCommit); ok {
			// struct pointer
			o.LastCommit = *sp
		}
	} else {
		o.LastCommit.FromMap(map[string]interface{}{})
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
	o.setDefaults(false)
}

// RepoResponseType is the enumeration type for type
type RepoResponseType int32

// String returns the string value for Type
func (v RepoResponseType) String() string {
	switch int32(v) {
	case 0:
		return "ENROLL"
	case 1:
		return "PING"
	case 2:
		return "CRASH"
	case 3:
		return "LOG"
	case 4:
		return "INTEGRATION"
	case 5:
		return "EXPORT"
	case 6:
		return "PROJECT"
	case 7:
		return "REPO"
	case 8:
		return "USER"
	case 9:
		return "UNINSTALL"
	case 10:
		return "UPGRADE"
	case 11:
		return "START"
	case 12:
		return "STOP"
	}
	return "unset"
}

const (
	// TypeEnroll is the enumeration value for enroll
	RepoResponseTypeEnroll RepoResponseType = 0
	// TypePing is the enumeration value for ping
	RepoResponseTypePing RepoResponseType = 1
	// TypeCrash is the enumeration value for crash
	RepoResponseTypeCrash RepoResponseType = 2
	// TypeLog is the enumeration value for log
	RepoResponseTypeLog RepoResponseType = 3
	// TypeIntegration is the enumeration value for integration
	RepoResponseTypeIntegration RepoResponseType = 4
	// TypeExport is the enumeration value for export
	RepoResponseTypeExport RepoResponseType = 5
	// TypeProject is the enumeration value for project
	RepoResponseTypeProject RepoResponseType = 6
	// TypeRepo is the enumeration value for repo
	RepoResponseTypeRepo RepoResponseType = 7
	// TypeUser is the enumeration value for user
	RepoResponseTypeUser RepoResponseType = 8
	// TypeUninstall is the enumeration value for uninstall
	RepoResponseTypeUninstall RepoResponseType = 9
	// TypeUpgrade is the enumeration value for upgrade
	RepoResponseTypeUpgrade RepoResponseType = 10
	// TypeStart is the enumeration value for start
	RepoResponseTypeStart RepoResponseType = 11
	// TypeStop is the enumeration value for stop
	RepoResponseTypeStop RepoResponseType = 12
)

// RepoResponse an agent response to an action request adding repo(s)
type RepoResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Data extra data that is specific about this event
	Data *string `json:"data" bson:"data" yaml:"data" faker:"-"`
	// Distro the agent os distribution
	Distro string `json:"distro" bson:"distro" yaml:"distro" faker:"-"`
	// Error an error message related to this event
	Error *string `json:"error" bson:"error" yaml:"error" faker:"-"`
	// EventDate the date of the event
	EventDate RepoResponseEventDate `json:"event_date" bson:"event_date" yaml:"event_date" faker:"-"`
	// FreeSpace the amount of free space in bytes for the agent machine
	FreeSpace int64 `json:"free_space" bson:"free_space" yaml:"free_space" faker:"-"`
	// GoVersion the go version that the agent build was built with
	GoVersion string `json:"go_version" bson:"go_version" yaml:"go_version" faker:"-"`
	// Hostname the agent hostname
	Hostname string `json:"hostname" bson:"hostname" yaml:"hostname" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// IntegrationID the integration id
	IntegrationID string `json:"integration_id" bson:"integration_id" yaml:"integration_id" faker:"-"`
	// LastExportDate the last export date
	LastExportDate RepoResponseLastExportDate `json:"last_export_date" bson:"last_export_date" yaml:"last_export_date" faker:"-"`
	// Memory the amount of memory in bytes for the agent machine
	Memory int64 `json:"memory" bson:"memory" yaml:"memory" faker:"-"`
	// Message a message related to this event
	Message string `json:"message" bson:"message" yaml:"message" faker:"-"`
	// NumCPU the number of CPU the agent is running
	NumCPU int64 `json:"num_cpu" bson:"num_cpu" yaml:"num_cpu" faker:"-"`
	// OS the agent operating system
	OS string `json:"os" bson:"os" yaml:"os" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Repos the repos exported
	Repos []RepoResponseRepos `json:"repos" bson:"repos" yaml:"repos" faker:"-"`
	// RequestID the request id that this response is correlated to
	RequestID string `json:"request_id" bson:"request_id" yaml:"request_id" faker:"-"`
	// Success if the response was successful
	Success bool `json:"success" bson:"success" yaml:"success" faker:"-"`
	// SystemID system unique device ID
	SystemID string `json:"system_id" bson:"system_id" yaml:"system_id" faker:"-"`
	// Type the type of event
	Type RepoResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
	// UpdatedAt the timestamp that the model was last updated fo real
	UpdatedAt int64 `json:"updated_ts" bson:"updated_ts" yaml:"updated_ts" faker:"-"`
	// Uptime the uptime in milliseconds since the agent started
	Uptime int64 `json:"uptime" bson:"uptime" yaml:"uptime" faker:"-"`
	// UUID the agent unique identifier
	UUID string `json:"uuid" bson:"uuid" yaml:"uuid" faker:"-"`
	// Version the agent version
	Version string `json:"version" bson:"version" yaml:"version" faker:"-"`
	// Hashcode stores the hash of the value of this object whereby two objects with the same hashcode are functionality equal
	Hashcode string `json:"hashcode" bson:"hashcode" yaml:"hashcode" faker:"-"`
}

// ensure that this type implements the data model interface
var _ datamodel.Model = (*RepoResponse)(nil)

func toRepoResponseObjectNil(isavro bool, isoptional bool) interface{} {
	if isavro && isoptional {
		return goavro.Union("null", nil)
	}
	return nil
}

func toRepoResponseObject(o interface{}, isavro bool, isoptional bool, avrotype string) interface{} {
	if res, ok := datamodel.ToGolangObject(o, isavro, isoptional, avrotype); ok {
		return res
	}
	switch v := o.(type) {
	case *RepoResponse:
		return v.ToMap(isavro)

	case RepoResponseEventDate:
		return v.ToMap(isavro)

	case RepoResponseLastExportDate:
		return v.ToMap(isavro)

	case []RepoResponseRepos:
		arr := make([]interface{}, 0)
		for _, i := range v {
			arr = append(arr, i.ToMap(isavro))
		}
		return arr

	case RepoResponseType:
		return v.String()

	default:
		panic("couldn't figure out the object type: " + reflect.TypeOf(v).String())
	}
}

// String returns a string representation of RepoResponse
func (o *RepoResponse) String() string {
	return fmt.Sprintf("agent.RepoResponse<%s>", o.ID)
}

// GetTopicName returns the name of the topic if evented
func (o *RepoResponse) GetTopicName() datamodel.TopicNameType {
	return RepoResponseTopic
}

// GetModelName returns the name of the model
func (o *RepoResponse) GetModelName() datamodel.ModelNameType {
	return RepoResponseModelName
}

func (o *RepoResponse) setDefaults(frommap bool) {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Repos == nil {
		o.Repos = make([]RepoResponseRepos, 0)
	}

	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoResponse", o.CustomerID, o.RefType, o.GetRefID())
	}

	if frommap {
		o.FromMap(map[string]interface{}{})
	}
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoResponse) GetID() string {
	return o.ID
}

// GetTopicKey returns the topic message key when sending this model as a ModelSendEvent
func (o *RepoResponse) GetTopicKey() string {
	var i interface{} = o.UUID
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

// GetTimestamp returns the timestamp for the model or now if not provided
func (o *RepoResponse) GetTimestamp() time.Time {
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
	panic("not sure how to handle the date time format for RepoResponse")
}

// GetRefID returns the RefID for the object
func (o *RepoResponse) GetRefID() string {
	return o.RefID
}

// IsMaterialized returns true if the model is materialized
func (o *RepoResponse) IsMaterialized() bool {
	return false
}

// GetModelMaterializeConfig returns the materialization config if materialized or nil if not
func (o *RepoResponse) GetModelMaterializeConfig() *datamodel.ModelMaterializeConfig {
	return nil
}

// IsEvented returns true if the model supports eventing and implements ModelEventProvider
func (o *RepoResponse) IsEvented() bool {
	return true
}

// SetEventHeaders will set any event headers for the object instance
func (o *RepoResponse) SetEventHeaders(kv map[string]string) {
	kv["customer_id"] = o.CustomerID
	kv["model"] = RepoResponseModelName.String()
}

// GetTopicConfig returns the topic config object
func (o *RepoResponse) GetTopicConfig() *datamodel.ModelTopicConfig {
	retention, err := time.ParseDuration("168h0m0s")
	if err != nil {
		panic("Invalid topic retention duration provided: 168h0m0s. " + err.Error())
	}

	ttl, err := time.ParseDuration("0s")
	if err != nil {
		ttl = 0
	}
	if ttl == 0 && retention != 0 {
		ttl = retention // they should be the same if not set
	}
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "updated_ts",
		NumPartitions:     8,
		ReplicationFactor: 3,
		Retention:         retention,
		MaxSize:           5242880,
		TTL:               ttl,
	}
}

// GetStateKey returns a key for use in state store
func (o *RepoResponse) GetStateKey() string {
	key := "uuid"
	return fmt.Sprintf("%s_%s", key, o.GetID())
}

// GetCustomerID will return the customer_id
func (o *RepoResponse) GetCustomerID() string {

	return o.CustomerID

}

// Clone returns an exact copy of RepoResponse
func (o *RepoResponse) Clone() datamodel.Model {
	c := new(RepoResponse)
	c.FromMap(o.ToMap())
	return c
}

// Anon returns the data structure as anonymous data
func (o *RepoResponse) Anon() datamodel.Model {
	c := new(RepoResponse)
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

// MarshalBinary returns the bytes for marshaling to binary
func (o *RepoResponse) MarshalBinary() ([]byte, error) {
	return o.MarshalJSON()
}

func (o *RepoResponse) UnmarshalBinary(data []byte) error {
	return o.UnmarshalJSON(data)
}

// MarshalJSON returns the bytes for marshaling to json
func (o *RepoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ToMap())
}

// UnmarshalJSON will unmarshal the json buffer into the object
func (o *RepoResponse) UnmarshalJSON(data []byte) error {
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

var cachedCodecRepoResponse *goavro.Codec
var cachedCodecRepoResponseLock sync.Mutex

// GetAvroCodec returns the avro codec for this model
func (o *RepoResponse) GetAvroCodec() *goavro.Codec {
	cachedCodecRepoResponseLock.Lock()
	if cachedCodecRepoResponse == nil {
		c, err := GetRepoResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecRepoResponse = c
	}
	cachedCodecRepoResponseLock.Unlock()
	return cachedCodecRepoResponse
}

// ToAvroBinary returns the data as Avro binary data
func (o *RepoResponse) ToAvroBinary() ([]byte, *goavro.Codec, error) {
	kv := o.ToMap(true)
	jbuf, _ := json.Marshal(kv)
	codec := o.GetAvroCodec()
	native, _, err := codec.NativeFromTextual(jbuf)
	if err != nil {
		return nil, nil, err
	}
	// Convert native Go form to binary Avro data
	buf, err := codec.BinaryFromNative(nil, native)
	return buf, codec, err
}

// FromAvroBinary will convert from Avro binary data into data in this object
func (o *RepoResponse) FromAvroBinary(value []byte) error {
	var nullHeader = []byte{byte(0)}
	// if this still has the schema encoded in the header, move past it to the avro payload
	if bytes.HasPrefix(value, nullHeader) {
		value = value[5:]
	}
	kv, _, err := o.GetAvroCodec().NativeFromBinary(value)
	if err != nil {
		return err
	}
	o.FromMap(kv.(map[string]interface{}))
	return nil
}

// Stringify returns the object in JSON format as a string
func (o *RepoResponse) Stringify() string {
	o.Hash()
	return pjson.Stringify(o)
}

// IsEqual returns true if the two RepoResponse objects are equal
func (o *RepoResponse) IsEqual(other *RepoResponse) bool {
	return o.Hash() == other.Hash()
}

// ToMap returns the object as a map
func (o *RepoResponse) ToMap(avro ...bool) map[string]interface{} {
	var isavro bool
	if len(avro) > 0 && avro[0] {
		isavro = true
	}
	if isavro {
		if o.Repos == nil {
			o.Repos = make([]RepoResponseRepos, 0)
		}
	}
	o.setDefaults(false)
	return map[string]interface{}{
		"architecture":     toRepoResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":      toRepoResponseObject(o.CustomerID, isavro, false, "string"),
		"data":             toRepoResponseObject(o.Data, isavro, true, "string"),
		"distro":           toRepoResponseObject(o.Distro, isavro, false, "string"),
		"error":            toRepoResponseObject(o.Error, isavro, true, "string"),
		"event_date":       toRepoResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":       toRepoResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":       toRepoResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":         toRepoResponseObject(o.Hostname, isavro, false, "string"),
		"id":               toRepoResponseObject(o.ID, isavro, false, "string"),
		"integration_id":   toRepoResponseObject(o.IntegrationID, isavro, false, "string"),
		"last_export_date": toRepoResponseObject(o.LastExportDate, isavro, false, "last_export_date"),
		"memory":           toRepoResponseObject(o.Memory, isavro, false, "long"),
		"message":          toRepoResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":          toRepoResponseObject(o.NumCPU, isavro, false, "long"),
		"os":               toRepoResponseObject(o.OS, isavro, false, "string"),
		"ref_id":           toRepoResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":         toRepoResponseObject(o.RefType, isavro, false, "string"),
		"repos":            toRepoResponseObject(o.Repos, isavro, false, "repos"),
		"request_id":       toRepoResponseObject(o.RequestID, isavro, false, "string"),
		"success":          toRepoResponseObject(o.Success, isavro, false, "boolean"),
		"system_id":        toRepoResponseObject(o.SystemID, isavro, false, "string"),
		"type":             toRepoResponseObject(o.Type, isavro, false, "type"),
		"updated_ts":       toRepoResponseObject(o.UpdatedAt, isavro, false, "long"),
		"uptime":           toRepoResponseObject(o.Uptime, isavro, false, "long"),
		"uuid":             toRepoResponseObject(o.UUID, isavro, false, "string"),
		"version":          toRepoResponseObject(o.Version, isavro, false, "string"),
		"hashcode":         toRepoResponseObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponse) FromMap(kv map[string]interface{}) {

	o.ID = ""

	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}

	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		if val, ok := kv["architecture"]; ok {
			if val == nil {
				o.Architecture = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Architecture = fmt.Sprintf("%v", val)
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

	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		if val, ok := kv["data"]; ok {
			if val == nil {
				o.Data = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Data = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		if val, ok := kv["distro"]; ok {
			if val == nil {
				o.Distro = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Distro = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		if val, ok := kv["error"]; ok {
			if val == nil {
				o.Error = pstrings.Pointer("")
			} else {
				// if coming in as avro union, convert it back
				if kv, ok := val.(map[string]interface{}); ok {
					val = kv["string"]
				}
				o.Error = pstrings.Pointer(fmt.Sprintf("%v", val))
			}
		}
	}

	if val, ok := kv["event_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.EventDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseEventDate); ok {
			// struct
			o.EventDate = sv
		} else if sp, ok := val.(*RepoResponseEventDate); ok {
			// struct pointer
			o.EventDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.EventDate.Epoch = dt.Epoch
			o.EventDate.Rfc3339 = dt.Rfc3339
			o.EventDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.EventDate.Epoch = dt.Epoch
				o.EventDate.Rfc3339 = dt.Rfc3339
				o.EventDate.Offset = dt.Offset
			}
		}
	} else {
		o.EventDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		if val, ok := kv["free_space"]; ok {
			if val == nil {
				o.FreeSpace = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.FreeSpace = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		if val, ok := kv["go_version"]; ok {
			if val == nil {
				o.GoVersion = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.GoVersion = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		if val, ok := kv["hostname"]; ok {
			if val == nil {
				o.Hostname = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Hostname = fmt.Sprintf("%v", val)
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

	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		if val, ok := kv["integration_id"]; ok {
			if val == nil {
				o.IntegrationID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.IntegrationID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["last_export_date"]; ok {
		if kv, ok := val.(map[string]interface{}); ok {
			o.LastExportDate.FromMap(kv)
		} else if sv, ok := val.(RepoResponseLastExportDate); ok {
			// struct
			o.LastExportDate = sv
		} else if sp, ok := val.(*RepoResponseLastExportDate); ok {
			// struct pointer
			o.LastExportDate = *sp
		} else if dt, ok := val.(*datetime.Date); ok && dt != nil {
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if tv, ok := val.(time.Time); ok && !tv.IsZero() {
			dt, err := datetime.NewDateWithTime(tv)
			if err != nil {
				panic(err)
			}
			o.LastExportDate.Epoch = dt.Epoch
			o.LastExportDate.Rfc3339 = dt.Rfc3339
			o.LastExportDate.Offset = dt.Offset
		} else if s, ok := val.(string); ok && s != "" {
			dt, err := datetime.NewDate(s)
			if err == nil {
				o.LastExportDate.Epoch = dt.Epoch
				o.LastExportDate.Rfc3339 = dt.Rfc3339
				o.LastExportDate.Offset = dt.Offset
			}
		}
	} else {
		o.LastExportDate.FromMap(map[string]interface{}{})
	}

	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		if val, ok := kv["memory"]; ok {
			if val == nil {
				o.Memory = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Memory = number.ToInt64Any(val)
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

	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		if val, ok := kv["num_cpu"]; ok {
			if val == nil {
				o.NumCPU = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.NumCPU = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		if val, ok := kv["os"]; ok {
			if val == nil {
				o.OS = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.OS = fmt.Sprintf("%v", val)
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

	if o == nil {

		o.Repos = make([]RepoResponseRepos, 0)

	}
	if val, ok := kv["repos"]; ok {
		if sv, ok := val.([]RepoResponseRepos); ok {
			o.Repos = sv
		} else if sp, ok := val.([]*RepoResponseRepos); ok {
			o.Repos = o.Repos[:0]
			for _, e := range sp {
				o.Repos = append(o.Repos, *e)
			}
		} else if a, ok := val.(primitive.A); ok {
			for _, ae := range a {
				if av, ok := ae.(RepoResponseRepos); ok {
					o.Repos = append(o.Repos, av)
				} else {
					b, _ := json.Marshal(ae)
					var av RepoResponseRepos
					if err := json.Unmarshal(b, &av); err != nil {
						panic("unsupported type for repos field entry: " + reflect.TypeOf(ae).String())
					}
					o.Repos = append(o.Repos, av)
				}
			}
		} else if arr, ok := val.([]interface{}); ok {
			for _, item := range arr {
				if r, ok := item.(RepoResponseRepos); ok {
					o.Repos = append(o.Repos, r)
				} else if r, ok := item.(map[string]interface{}); ok {
					var fm RepoResponseRepos
					fm.FromMap(r)
					o.Repos = append(o.Repos, fm)
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
							var fm RepoResponseRepos
							fm.FromMap(m[0].Interface().(map[string]interface{}))
							o.Repos = append(o.Repos, fm)
						}
					}
				}
			}
		}
	}

	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		if val, ok := kv["request_id"]; ok {
			if val == nil {
				o.RequestID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.RequestID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		if val, ok := kv["success"]; ok {
			if val == nil {
				o.Success = number.ToBoolAny(nil)
			} else {
				o.Success = number.ToBoolAny(val)
			}
		}
	}

	if val, ok := kv["system_id"].(string); ok {
		o.SystemID = val
	} else {
		if val, ok := kv["system_id"]; ok {
			if val == nil {
				o.SystemID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.SystemID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["type"].(RepoResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll", "ENROLL":
				o.Type = 0
			case "ping", "PING":
				o.Type = 1
			case "crash", "CRASH":
				o.Type = 2
			case "log", "LOG":
				o.Type = 3
			case "integration", "INTEGRATION":
				o.Type = 4
			case "export", "EXPORT":
				o.Type = 5
			case "project", "PROJECT":
				o.Type = 6
			case "repo", "REPO":
				o.Type = 7
			case "user", "USER":
				o.Type = 8
			case "uninstall", "UNINSTALL":
				o.Type = 9
			case "upgrade", "UPGRADE":
				o.Type = 10
			case "start", "START":
				o.Type = 11
			case "stop", "STOP":
				o.Type = 12
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

	if val, ok := kv["uptime"].(int64); ok {
		o.Uptime = val
	} else {
		if val, ok := kv["uptime"]; ok {
			if val == nil {
				o.Uptime = number.ToInt64Any(nil)
			} else {
				if tv, ok := val.(time.Time); ok {
					val = datetime.TimeToEpoch(tv)
				}
				o.Uptime = number.ToInt64Any(val)
			}
		}
	}

	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		if val, ok := kv["uuid"]; ok {
			if val == nil {
				o.UUID = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.UUID = fmt.Sprintf("%v", val)
			}
		}
	}

	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		if val, ok := kv["version"]; ok {
			if val == nil {
				o.Version = ""
			} else {
				if m, ok := val.(map[string]interface{}); ok {
					val = pjson.Stringify(m)
				}
				o.Version = fmt.Sprintf("%v", val)
			}
		}
	}
	o.setDefaults(false)
}

// Hash will return a hashcode for the object
func (o *RepoResponse) Hash() string {
	args := make([]interface{}, 0)
	args = append(args, o.Architecture)
	args = append(args, o.CustomerID)
	args = append(args, o.Data)
	args = append(args, o.Distro)
	args = append(args, o.Error)
	args = append(args, o.EventDate)
	args = append(args, o.FreeSpace)
	args = append(args, o.GoVersion)
	args = append(args, o.Hostname)
	args = append(args, o.ID)
	args = append(args, o.IntegrationID)
	args = append(args, o.LastExportDate)
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Repos)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.SystemID)
	args = append(args, o.Type)
	args = append(args, o.UpdatedAt)
	args = append(args, o.Uptime)
	args = append(args, o.UUID)
	args = append(args, o.Version)
	o.Hashcode = hash.Values(args...)
	return o.Hashcode
}

// GetRepoResponseAvroSchemaSpec creates the avro schema specification for RepoResponse
func GetRepoResponseAvroSchemaSpec() string {
	spec := map[string]interface{}{
		"type":      "record",
		"namespace": "agent",
		"name":      "RepoResponse",
		"fields": []map[string]interface{}{
			map[string]interface{}{
				"name": "hashcode",
				"type": "string",
			},
			map[string]interface{}{
				"name": "architecture",
				"type": "string",
			},
			map[string]interface{}{
				"name": "customer_id",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "data",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "distro",
				"type": "string",
			},
			map[string]interface{}{
				"name":    "error",
				"type":    []interface{}{"null", "string"},
				"default": nil,
			},
			map[string]interface{}{
				"name": "event_date",
				"type": map[string]interface{}{"doc": "the date of the event", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "event_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "free_space",
				"type": "long",
			},
			map[string]interface{}{
				"name": "go_version",
				"type": "string",
			},
			map[string]interface{}{
				"name": "hostname",
				"type": "string",
			},
			map[string]interface{}{
				"name": "id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "integration_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "last_export_date",
				"type": map[string]interface{}{"doc": "the last export date", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "last_export_date", "type": "record"},
			},
			map[string]interface{}{
				"name": "memory",
				"type": "long",
			},
			map[string]interface{}{
				"name": "message",
				"type": "string",
			},
			map[string]interface{}{
				"name": "num_cpu",
				"type": "long",
			},
			map[string]interface{}{
				"name": "os",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "ref_type",
				"type": "string",
			},
			map[string]interface{}{
				"name": "repos",
				"type": map[string]interface{}{"items": map[string]interface{}{"doc": "the repos exported", "fields": []interface{}{map[string]interface{}{"doc": "the status of the repo determined by an Admin", "name": "active", "type": "boolean"}, map[string]interface{}{"doc": "the creation date", "name": "created_date", "type": map[string]interface{}{"doc": "the creation date", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "repos.created_date", "type": "record"}}, map[string]interface{}{"doc": "the description of the repository", "name": "description", "type": "string"}, map[string]interface{}{"doc": "the programming language defined for the repository", "name": "language", "type": "string"}, map[string]interface{}{"doc": "the most recent commit to the repo", "name": "last_commit", "type": map[string]interface{}{"doc": "the most recent commit to the repo", "fields": []interface{}{map[string]interface{}{"doc": "the id of the latest commit", "name": "commit_id", "type": "string"}, map[string]interface{}{"doc": "the url of the lastest commit", "name": "url", "type": "string"}, map[string]interface{}{"doc": "the commit message of the latest commit", "name": "message", "type": "string"}, map[string]interface{}{"doc": "the timestamp of the latest commit", "name": "created_date", "type": map[string]interface{}{"doc": "the timestamp of the latest commit", "fields": []interface{}{map[string]interface{}{"doc": "the date in epoch format", "name": "epoch", "type": "long"}, map[string]interface{}{"doc": "the timezone offset from GMT", "name": "offset", "type": "long"}, map[string]interface{}{"doc": "the date in RFC3339 format", "name": "rfc3339", "type": "string"}}, "name": "last_commit.created_date", "type": "record"}}, map[string]interface{}{"doc": "the author of the latest commit", "name": "author", "type": map[string]interface{}{"doc": "the author of the latest commit", "fields": []interface{}{map[string]interface{}{"doc": "the author name", "name": "name", "type": "string"}, map[string]interface{}{"doc": "the email of the author", "name": "email", "type": "string"}, map[string]interface{}{"doc": "the avatar_url for the author", "name": "avatar_url", "type": "string"}}, "name": "last_commit.author", "type": "record"}}}, "name": "repos.last_commit", "type": "record"}}, map[string]interface{}{"doc": "the name of the repository", "name": "name", "type": "string"}, map[string]interface{}{"doc": "the repo ID", "name": "ref_id", "type": "string"}, map[string]interface{}{"doc": "the record type", "name": "ref_type", "type": "string"}}, "name": "repos", "type": "record"}, "name": "repos", "type": "array"},
			},
			map[string]interface{}{
				"name": "request_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "success",
				"type": "boolean",
			},
			map[string]interface{}{
				"name": "system_id",
				"type": "string",
			},
			map[string]interface{}{
				"name": "type",
				"type": map[string]interface{}{
					"type":    "enum",
					"name":    "type",
					"symbols": []interface{}{"ENROLL", "PING", "CRASH", "LOG", "INTEGRATION", "EXPORT", "PROJECT", "REPO", "USER", "UNINSTALL", "UPGRADE", "START", "STOP"},
				},
			},
			map[string]interface{}{
				"name": "updated_ts",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uptime",
				"type": "long",
			},
			map[string]interface{}{
				"name": "uuid",
				"type": "string",
			},
			map[string]interface{}{
				"name": "version",
				"type": "string",
			},
		},
	}
	return pjson.Stringify(spec, true)
}

// GetEventAPIConfig returns the EventAPIConfig
func (o *RepoResponse) GetEventAPIConfig() datamodel.EventAPIConfig {
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

// GetRepoResponseAvroSchema creates the avro schema for RepoResponse
func GetRepoResponseAvroSchema() (*goavro.Codec, error) {
	return goavro.NewCodec(GetRepoResponseAvroSchemaSpec())
}

// TransformRepoResponseFunc is a function for transforming RepoResponse during processing
type TransformRepoResponseFunc func(input *RepoResponse) (*RepoResponse, error)

// NewRepoResponsePipe creates a pipe for processing RepoResponse items
func NewRepoResponsePipe(input io.ReadCloser, output io.WriteCloser, errors chan error, transforms ...TransformRepoResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	inch, indone := NewRepoResponseInputStream(input, errors)
	var stream chan RepoResponse
	if len(transforms) > 0 {
		stream = make(chan RepoResponse, 1000)
	} else {
		stream = inch
	}
	outdone := NewRepoResponseOutputStream(output, stream, errors)
	go func() {
		if len(transforms) > 0 {
			var stop bool
			for item := range inch {
				input := &item
				for _, transform := range transforms {
					out, err := transform(input)
					if err != nil {
						stop = true
						errors <- err
						break
					}
					if out == nil {
						input = nil
						break
					} else {
						input = out
					}
				}
				if stop {
					break
				}
				if input != nil {
					stream <- *input
				}
			}
			close(stream)
		}
		<-indone
		<-outdone
		done <- true
	}()
	return done
}

// NewRepoResponseInputStreamDir creates a channel for reading RepoResponse as JSON newlines from a directory of files
func NewRepoResponseInputStreamDir(dir string, errors chan<- error, transforms ...TransformRepoResponseFunc) (chan RepoResponse, <-chan bool) {
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("/agent/repo_response\\.json(\\.gz)?$"))
	if err != nil {
		errors <- err
		ch := make(chan RepoResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	l := len(files)
	if l > 1 {
		errors <- fmt.Errorf("too many files matched our finder regular expression for repo_response")
		ch := make(chan RepoResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	} else if l == 1 {
		return NewRepoResponseInputStreamFile(files[0], errors, transforms...)
	} else {
		ch := make(chan RepoResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
}

// NewRepoResponseInputStreamFile creates an channel for reading RepoResponse as JSON newlines from filename
func NewRepoResponseInputStreamFile(filename string, errors chan<- error, transforms ...TransformRepoResponseFunc) (chan RepoResponse, <-chan bool) {
	of, err := os.Open(filename)
	if err != nil {
		errors <- err
		ch := make(chan RepoResponse)
		close(ch)
		done := make(chan bool, 1)
		done <- true
		return ch, done
	}
	var f io.ReadCloser = of
	if filepath.Ext(filename) == ".gz" {
		gz, err := gzip.NewReader(f)
		if err != nil {
			of.Close()
			errors <- err
			ch := make(chan RepoResponse)
			close(ch)
			done := make(chan bool, 1)
			done <- true
			return ch, done
		}
		f = gz
	}
	return NewRepoResponseInputStream(f, errors, transforms...)
}

// NewRepoResponseInputStream creates an channel for reading RepoResponse as JSON newlines from stream
func NewRepoResponseInputStream(stream io.ReadCloser, errors chan<- error, transforms ...TransformRepoResponseFunc) (chan RepoResponse, <-chan bool) {
	done := make(chan bool, 1)
	ch := make(chan RepoResponse, 1000)
	go func() {
		defer func() { stream.Close(); close(ch); done <- true }()
		r := bufio.NewReader(stream)
		for {
			buf, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				errors <- err
				return
			}
			var item RepoResponse
			if err := json.Unmarshal(buf, &item); err != nil {
				errors <- err
				return
			}
			in := &item
			var skip bool
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				ch <- *in
			}
		}
	}()
	return ch, done
}

// NewRepoResponseOutputStreamDir will output json newlines from channel and save in dir
func NewRepoResponseOutputStreamDir(dir string, ch chan RepoResponse, errors chan<- error, transforms ...TransformRepoResponseFunc) <-chan bool {
	fp := filepath.Join(dir, "/agent/repo_response\\.json(\\.gz)?$")
	os.MkdirAll(filepath.Dir(fp), 0777)
	of, err := os.Create(fp)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	gz, err := gzip.NewWriterLevel(of, gzip.BestCompression)
	if err != nil {
		errors <- err
		done := make(chan bool, 1)
		done <- true
		return done
	}
	return NewRepoResponseOutputStream(gz, ch, errors, transforms...)
}

// NewRepoResponseOutputStream will output json newlines from channel to the stream
func NewRepoResponseOutputStream(stream io.WriteCloser, ch chan RepoResponse, errors chan<- error, transforms ...TransformRepoResponseFunc) <-chan bool {
	done := make(chan bool, 1)
	go func() {
		defer func() {
			if gz, ok := stream.(*gzip.Writer); ok {
				gz.Flush()
				gz.Close()
			}
			stream.Close()
			done <- true
		}()
		for item := range ch {
			in := &item
			var skip bool
			var err error
			for _, transform := range transforms {
				in, err = transform(in)
				if err != nil {
					errors <- err
					return
				}
				if in == nil {
					skip = true
					break
				}
			}
			if !skip {
				buf, err := json.Marshal(in)
				if err != nil {
					errors <- err
					return
				}
				stream.Write(buf)
				stream.Write([]byte{'\n'})
			}
		}
	}()
	return done
}

// RepoResponseSendEvent is an event detail for sending data
type RepoResponseSendEvent struct {
	RepoResponse *RepoResponse
	headers      map[string]string
	time         time.Time
	key          string
}

var _ datamodel.ModelSendEvent = (*RepoResponseSendEvent)(nil)

// Key is the key to use for the message
func (e *RepoResponseSendEvent) Key() string {
	if e.key == "" {
		return e.RepoResponse.GetID()
	}
	return e.key
}

// Object returns an instance of the Model that will be send
func (e *RepoResponseSendEvent) Object() datamodel.Model {
	return e.RepoResponse
}

// Headers returns any headers for the event. can be nil to not send any additional headers
func (e *RepoResponseSendEvent) Headers() map[string]string {
	return e.headers
}

// Timestamp returns the event timestamp. If empty, will default to time.Now()
func (e *RepoResponseSendEvent) Timestamp() time.Time {
	return e.time
}

// RepoResponseSendEventOpts is a function handler for setting opts
type RepoResponseSendEventOpts func(o *RepoResponseSendEvent)

// WithRepoResponseSendEventKey sets the key value to a value different than the object ID
func WithRepoResponseSendEventKey(key string) RepoResponseSendEventOpts {
	return func(o *RepoResponseSendEvent) {
		o.key = key
	}
}

// WithRepoResponseSendEventTimestamp sets the timestamp value
func WithRepoResponseSendEventTimestamp(tv time.Time) RepoResponseSendEventOpts {
	return func(o *RepoResponseSendEvent) {
		o.time = tv
	}
}

// WithRepoResponseSendEventHeader sets the timestamp value
func WithRepoResponseSendEventHeader(key, value string) RepoResponseSendEventOpts {
	return func(o *RepoResponseSendEvent) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}

// NewRepoResponseSendEvent returns a new RepoResponseSendEvent instance
func NewRepoResponseSendEvent(o *RepoResponse, opts ...RepoResponseSendEventOpts) *RepoResponseSendEvent {
	res := &RepoResponseSendEvent{
		RepoResponse: o,
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(res)
		}
	}
	return res
}

// NewRepoResponseProducer will stream data from the channel
func NewRepoResponseProducer(ctx context.Context, producer eventing.Producer, ch <-chan datamodel.ModelSendEvent, errors chan<- error, empty chan<- bool) <-chan bool {
	done := make(chan bool, 1)
	emptyTime := time.Unix(0, 0)
	go func() {
		defer func() { done <- true }()
		for {
			select {
			case <-ctx.Done():
				return
			case item := <-ch:
				if item == nil {
					empty <- true
					return
				}
				if object, ok := item.Object().(*RepoResponse); ok {
					binary, codec, err := object.ToAvroBinary()
					if err != nil {
						errors <- fmt.Errorf("error encoding %s to avro binary data. %v", object.String(), err)
						return
					}
					headers := map[string]string{}
					object.SetEventHeaders(headers)
					for k, v := range item.Headers() {
						headers[k] = v
					}
					tv := item.Timestamp()
					if tv.IsZero() {
						tv = object.GetTimestamp() // if not provided in the message, use the objects value
					}
					if tv.IsZero() || tv.Equal(emptyTime) {
						tv = time.Now() // if its still zero, use the ingest time
					}
					// add generated message headers
					headers["message-id"] = pstrings.NewUUIDV4()
					headers["message-ts"] = fmt.Sprintf("%v", datetime.EpochNow())
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
						Partition: -1, // select any partition based on partitioner strategy in kafka
						Topic:     object.GetTopicName().String(),
					}
					if err := producer.Send(ctx, msg); err != nil {
						errors <- fmt.Errorf("error sending %s. %v", object.String(), err)
					}
				} else {
					errors <- fmt.Errorf("invalid event received. expected an object of type agent.RepoResponse but received on of type %v", reflect.TypeOf(item.Object()))
				}
			}
		}
	}()
	return done
}

// NewRepoResponseConsumer will stream data from the topic into the provided channel
func NewRepoResponseConsumer(consumer eventing.Consumer, ch chan<- datamodel.ModelReceiveEvent, errors chan<- error) *eventing.ConsumerCallbackAdapter {
	adapter := &eventing.ConsumerCallbackAdapter{
		OnDataReceived: func(msg eventing.Message) error {
			var object RepoResponse
			switch msg.Encoding {
			case eventing.JSONEncoding:
				if err := json.Unmarshal(msg.Value, &object); err != nil {
					return fmt.Errorf("error unmarshaling json data into agent.RepoResponse: %s", err)
				}
			case eventing.AvroEncoding:
				if err := object.FromAvroBinary(msg.Value); err != nil {
					return fmt.Errorf("error unmarshaling avro data into agent.RepoResponse: %s", err)
				}
			default:
				return fmt.Errorf("unsure of the encoding since it was not set for agent.RepoResponse")
			}

			// ignore messages that have exceeded the TTL
			cfg := object.GetTopicConfig()
			if cfg != nil && cfg.TTL != 0 && msg.Timestamp.UTC().Add(cfg.TTL).Sub(time.Now().UTC()) < 0 {
				// if disable auto and we're skipping, we need to commit the message
				if !msg.IsAutoCommit() {
					msg.Commit()
				}
				return nil
			}
			msg.Codec = object.GetAvroCodec() // match the codec

			ch <- &RepoResponseReceiveEvent{&object, msg, false}
			return nil
		},
		OnErrorReceived: func(err error) {
			errors <- err
		},
		OnEOF: func(topic string, partition int32, offset int64) {
			var object RepoResponse
			var msg eventing.Message
			msg.Topic = topic
			msg.Partition = partition
			msg.Codec = object.GetAvroCodec() // match the codec
			ch <- &RepoResponseReceiveEvent{nil, msg, true}
		},
	}
	consumer.Consume(adapter)
	return adapter
}

// RepoResponseReceiveEvent is an event detail for receiving data
type RepoResponseReceiveEvent struct {
	RepoResponse *RepoResponse
	message      eventing.Message
	eof          bool
}

var _ datamodel.ModelReceiveEvent = (*RepoResponseReceiveEvent)(nil)

// Object returns an instance of the Model that was received
func (e *RepoResponseReceiveEvent) Object() datamodel.Model {
	return e.RepoResponse
}

// Message returns the underlying message data for the event
func (e *RepoResponseReceiveEvent) Message() eventing.Message {
	return e.message
}

// EOF returns true if an EOF event was received. in this case, the Object and Message will return nil
func (e *RepoResponseReceiveEvent) EOF() bool {
	return e.eof
}

// RepoResponseProducer implements the datamodel.ModelEventProducer
type RepoResponseProducer struct {
	ch       chan datamodel.ModelSendEvent
	done     <-chan bool
	producer eventing.Producer
	closed   bool
	mu       sync.Mutex
	ctx      context.Context
	cancel   context.CancelFunc
	empty    chan bool
}

var _ datamodel.ModelEventProducer = (*RepoResponseProducer)(nil)

// Channel returns the producer channel to produce new events
func (p *RepoResponseProducer) Channel() chan<- datamodel.ModelSendEvent {
	return p.ch
}

// Close is called to shutdown the producer
func (p *RepoResponseProducer) Close() error {
	p.mu.Lock()
	closed := p.closed
	p.closed = true
	p.mu.Unlock()
	if !closed {
		close(p.ch)
		<-p.empty
		p.cancel()
		<-p.done
	}
	return nil
}

// NewProducerChannel returns a channel which can be used for producing Model events
func (o *RepoResponse) NewProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return o.NewProducerChannelSize(producer, 0, errors)
}

// NewProducerChannelSize returns a channel which can be used for producing Model events
func (o *RepoResponse) NewProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// NewRepoResponseProducerChannel returns a channel which can be used for producing Model events
func NewRepoResponseProducerChannel(producer eventing.Producer, errors chan<- error) datamodel.ModelEventProducer {
	return NewRepoResponseProducerChannelSize(producer, 0, errors)
}

// NewRepoResponseProducerChannelSize returns a channel which can be used for producing Model events
func NewRepoResponseProducerChannelSize(producer eventing.Producer, size int, errors chan<- error) datamodel.ModelEventProducer {
	ch := make(chan datamodel.ModelSendEvent, size)
	empty := make(chan bool, 1)
	newctx, cancel := context.WithCancel(context.Background())
	return &RepoResponseProducer{
		ch:       ch,
		ctx:      newctx,
		cancel:   cancel,
		producer: producer,
		empty:    empty,
		done:     NewRepoResponseProducer(newctx, producer, ch, errors, empty),
	}
}

// RepoResponseConsumer implements the datamodel.ModelEventConsumer
type RepoResponseConsumer struct {
	ch       chan datamodel.ModelReceiveEvent
	consumer eventing.Consumer
	callback *eventing.ConsumerCallbackAdapter
	closed   bool
	mu       sync.Mutex
}

var _ datamodel.ModelEventConsumer = (*RepoResponseConsumer)(nil)

// Channel returns the consumer channel to consume new events
func (c *RepoResponseConsumer) Channel() <-chan datamodel.ModelReceiveEvent {
	return c.ch
}

// Close is called to shutdown the producer
func (c *RepoResponseConsumer) Close() error {
	c.mu.Lock()
	closed := c.closed
	c.closed = true
	c.mu.Unlock()
	var err error
	if !closed {
		c.callback.Close()
		err = c.consumer.Close()
	}
	return err
}

// NewConsumerChannel returns a consumer channel which can be used to consume Model events
func (o *RepoResponse) NewConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoResponseConsumer{
		ch:       ch,
		callback: NewRepoResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}

// NewRepoResponseConsumerChannel returns a consumer channel which can be used to consume Model events
func NewRepoResponseConsumerChannel(consumer eventing.Consumer, errors chan<- error) datamodel.ModelEventConsumer {
	ch := make(chan datamodel.ModelReceiveEvent)
	return &RepoResponseConsumer{
		ch:       ch,
		callback: NewRepoResponseConsumer(consumer, ch, errors),
		consumer: consumer,
	}
}
