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
	"github.com/pinpt/go-common/slice"
	pstrings "github.com/pinpt/go-common/strings"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// RepoResponseTopic is the default topic name
	RepoResponseTopic datamodel.TopicNameType = "agent_RepoResponse_topic"

	// RepoResponseStream is the default stream name
	RepoResponseStream datamodel.TopicNameType = "agent_RepoResponse_stream"

	// RepoResponseTable is the default table name
	RepoResponseTable datamodel.TopicNameType = "agent_RepoResponse"

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
	// RepoResponseRequestIDColumn is the request_id column name
	RepoResponseRequestIDColumn = "request_id"
	// RepoResponseSuccessColumn is the success column name
	RepoResponseSuccessColumn = "success"
	// RepoResponseTypeColumn is the type column name
	RepoResponseTypeColumn = "type"
	// RepoResponseUUIDColumn is the uuid column name
	RepoResponseUUIDColumn = "uuid"
	// RepoResponseVersionColumn is the version column name
	RepoResponseVersionColumn = "version"
)

// 0 architecture
// architecture {"description":"the architecture of the agent machine","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"architecture","relation":false,"subtype":"","type":"string"}

// 1 created_date
// created_date {"description":"the timestamp of the latest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":true,"is_object":false,"name":"created_date","relation":false,"subtype":"","type":"common.Date"}

// customer_id {"description":"the customer id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"customer_id","relation":true,"subtype":"","type":"string"}

// epoch {"description":"the date in epoch format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"epoch","relation":false,"subtype":"","type":"int"}

// id {"description":"the primary key for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"id","relation":false,"subtype":"","type":"string"}

// offset {"description":"the timezone offset from GMT","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"offset","relation":false,"subtype":"","type":"int"}

// ref_id {"description":"the source system id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_id","relation":false,"subtype":"","type":"string"}

// ref_type {"description":"the source system identifier for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_type","relation":false,"subtype":"","type":"string"}

// rfc3339 {"description":"the date in RFC3339 format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"rfc3339","relation":false,"subtype":"","type":"string"}

// RepoResponseLastCommitCreatedDate represents the object structure for created_date
type RepoResponseLastCommitCreatedDate struct {
	// CustomerID the customer id for the model instance
	CustomerID string `json:"customer_id" bson:"customer_id" yaml:"customer_id" faker:"-"`
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// ID the primary key for the model instance
	ID string `json:"id" bson:"_id" yaml:"id" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// RefID the source system id for the model instance
	RefID string `json:"ref_id" bson:"ref_id" yaml:"ref_id" faker:"-"`
	// RefType the source system identifier for the model instance
	RefType string `json:"ref_type" bson:"ref_type" yaml:"ref_type" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *RepoResponseLastCommitCreatedDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// CustomerID the customer id for the model instance
		"customer_id": o.CustomerID,
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// ID the primary key for the model instance
		"id": o.ID,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// RefID the source system id for the model instance
		"ref_id": o.RefID,
		// RefType the source system identifier for the model instance
		"ref_type": o.RefType,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// 2 customer_id
// customer_id {"description":"the customer id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"customer_id","relation":true,"subtype":"","type":"string"}

// 3 data
// data {"description":"extra data that is specific about this event","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"data","relation":false,"subtype":"","type":"string"}

// 4 distro
// distro {"description":"the agent os distribution","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"distro","relation":false,"subtype":"","type":"string"}

// 5 error
// error {"description":"an error message related to this event","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"error","relation":false,"subtype":"","type":"string"}

// 6 event_date
// event_date {"description":"the date of the event","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"event_date","relation":false,"subtype":"","type":"object"}

// epoch {"description":"the date in epoch format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"epoch","relation":false,"subtype":"","type":"int"}

// offset {"description":"the timezone offset from GMT","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"offset","relation":false,"subtype":"","type":"int"}

// rfc3339 {"description":"the date in RFC3339 format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"rfc3339","relation":false,"subtype":"","type":"string"}

// RepoResponseEventDate represents the object structure for event_date
type RepoResponseEventDate struct {
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
}

func (o *RepoResponseEventDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
	}
}

// 7 free_space
// free_space {"description":"the amount of free space in bytes for the agent machine","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"free_space","relation":false,"subtype":"","type":"long"}

// 8 go_version
// go_version {"description":"the go version that the agent build was built with","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"go_version","relation":false,"subtype":"","type":"string"}

// 9 hostname
// hostname {"description":"the agent hostname","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"hostname","relation":false,"subtype":"","type":"string"}

// 10 id
// id {"description":"the primary key for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"id","relation":false,"subtype":"","type":"string"}

// 11 integration_id
// integration_id {"description":"the integration id","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"integration_id","relation":false,"subtype":"","type":"string"}

// 12 memory
// memory {"description":"the amount of memory in bytes for the agent machine","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"memory","relation":false,"subtype":"","type":"long"}

// 13 message
// message {"description":"a message related to this event","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"message","relation":false,"subtype":"","type":"string"}

// 14 num_cpu
// num_cpu {"description":"the number of CPU the agent is running","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"num_cpu","relation":false,"subtype":"","type":"int"}

// 15 os
// os {"description":"the agent operating system","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"os","relation":false,"subtype":"","type":"string"}

// 16 ref_id
// ref_id {"description":"the source system id for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_id","relation":false,"subtype":"","type":"string"}

// 17 ref_type
// ref_type {"description":"the source system identifier for the model instance","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"ref_type","relation":false,"subtype":"","type":"string"}

// 18 repos
// repos {"description":"the repos exported","is_array":true,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"repos","relation":false,"subtype":"admin.Repo","type":"object"}

// active {"description":"the status of the repo determined by an Admin","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"active","relation":false,"subtype":"","type":"boolean"}

// created_date {"description":"the creation date","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"created_date","relation":false,"subtype":"","type":"object"}

// rfc3339 {"description":"the date in RFC3339 format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"rfc3339","relation":false,"subtype":"","type":"string"}

// epoch {"description":"the date in epoch format","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"epoch","relation":false,"subtype":"","type":"int"}

// offset {"description":"the timezone offset from GMT","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"offset","relation":false,"subtype":"","type":"int"}

// RepoResponseReposCreatedDate represents the object structure for created_date
type RepoResponseReposCreatedDate struct {
	// Rfc3339 the date in RFC3339 format
	Rfc3339 string `json:"rfc3339" bson:"rfc3339" yaml:"rfc3339" faker:"-"`
	// Epoch the date in epoch format
	Epoch int64 `json:"epoch" bson:"epoch" yaml:"epoch" faker:"-"`
	// Offset the timezone offset from GMT
	Offset int64 `json:"offset" bson:"offset" yaml:"offset" faker:"-"`
}

func (o *RepoResponseReposCreatedDate) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Rfc3339 the date in RFC3339 format
		"rfc3339": o.Rfc3339,
		// Epoch the date in epoch format
		"epoch": o.Epoch,
		// Offset the timezone offset from GMT
		"offset": o.Offset,
	}
}

// description {"description":"the description of the repository","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"description","relation":false,"subtype":"","type":"string"}

// language {"description":"the programming language defined for the repository","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"language","relation":false,"subtype":"","type":"string"}

// last_commit {"description":"the most recent commit to the repo","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"last_commit","relation":false,"subtype":"","type":"object"}

// commit_id {"description":"the id of the latest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"commit_id","relation":false,"subtype":"","type":"string"}

// url {"description":"the url of the lastest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"url","relation":false,"subtype":"","type":"string"}

// message {"description":"the commit message of the latest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"message","relation":false,"subtype":"","type":"string"}

// created_date {"description":"the timestamp of the latest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":true,"is_object":false,"name":"created_date","relation":false,"subtype":"","type":"common.Date"}

// author {"description":"the author of the latest commit","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":true,"name":"author","relation":false,"subtype":"","type":"object"}

// name {"description":"the author name","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"name","relation":false,"subtype":"","type":"string"}

// email {"description":"the email of the author","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"email","relation":false,"subtype":"","type":"string"}

// avatar_url {"description":"the avatar_url for the author","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"avatar_url","relation":false,"subtype":"","type":"string"}

// RepoResponseLastCommitAuthor represents the object structure for author
type RepoResponseLastCommitAuthor struct {
	// Name the author name
	Name string `json:"name" bson:"name" yaml:"name" faker:"person"`
	// Email the email of the author
	Email string `json:"email" bson:"email" yaml:"email" faker:"email"`
	// AvatarURL the avatar_url for the author
	AvatarURL string `json:"avatar_url" bson:"avatar_url" yaml:"avatar_url" faker:"avatar"`
}

func (o *RepoResponseLastCommitAuthor) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Name the author name
		"name": o.Name,
		// Email the email of the author
		"email": o.Email,
		// AvatarURL the avatar_url for the author
		"avatar_url": o.AvatarURL,
	}
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
	CreatedDate RepoResponseLastCommitCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
	// Author the author of the latest commit
	Author RepoResponseLastCommitAuthor `json:"author" bson:"author" yaml:"author" faker:"-"`
}

func (o *RepoResponseReposLastCommit) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// CommitID the id of the latest commit
		"commit_id": o.CommitID,
		// URL the url of the lastest commit
		"url": o.URL,
		// Message the commit message of the latest commit
		"message": o.Message,
		// CreatedDate the timestamp of the latest commit
		"created_date": o.CreatedDate,
		// Author the author of the latest commit
		"author": o.Author,
	}
}

// name {"description":"the name of the repository","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"name","relation":false,"subtype":"","type":"string"}

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
}

func (o *RepoResponseRepos) ToMap() map[string]interface{} {
	return map[string]interface{}{
		// Active the status of the repo determined by an Admin
		"active": o.Active,
		// CreatedDate the creation date
		"created_date": o.CreatedDate,
		// Description the description of the repository
		"description": o.Description,
		// Language the programming language defined for the repository
		"language": o.Language,
		// LastCommit the most recent commit to the repo
		"last_commit": o.LastCommit,
		// Name the name of the repository
		"name": o.Name,
	}
}

// 19 request_id
// request_id {"description":"the request id that this response is correlated to","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"request_id","relation":false,"subtype":"","type":"string"}

// 20 success
// success {"description":"if the response was successful","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"success","relation":false,"subtype":"","type":"boolean"}

// 21 type
// type {"description":"the type of event","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"type","relation":false,"subtype":"","type":"enum"}

// RepoResponseType is the enumeration type for type
type RepoResponseType int32

// String returns the string value for Type
func (v RepoResponseType) String() string {
	switch int32(v) {
	case 0:
		return "enroll"
	case 1:
		return "ping"
	case 2:
		return "crash"
	case 3:
		return "integration"
	case 4:
		return "export"
	case 5:
		return "project"
	case 6:
		return "repo"
	case 7:
		return "user"
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
	// TypeIntegration is the enumeration value for integration
	RepoResponseTypeIntegration RepoResponseType = 3
	// TypeExport is the enumeration value for export
	RepoResponseTypeExport RepoResponseType = 4
	// TypeProject is the enumeration value for project
	RepoResponseTypeProject RepoResponseType = 5
	// TypeRepo is the enumeration value for repo
	RepoResponseTypeRepo RepoResponseType = 6
	// TypeUser is the enumeration value for user
	RepoResponseTypeUser RepoResponseType = 7
)

// 22 uuid
// uuid {"description":"the agent unique identifier","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"uuid","relation":false,"subtype":"","type":"string"}

// 23 version
// version {"description":"the agent version","is_array":false,"is_hidden":false,"is_map":false,"is_nested":false,"is_object":false,"name":"version","relation":false,"subtype":"","type":"string"}

// RepoResponse an agent response to an action request adding repo(s)
type RepoResponse struct {
	// Architecture the architecture of the agent machine
	Architecture string `json:"architecture" bson:"architecture" yaml:"architecture" faker:"-"`
	// LastCommitCreatedDate the timestamp of the latest commit
	LastCommitCreatedDate RepoResponseLastCommitCreatedDate `json:"created_date" bson:"created_date" yaml:"created_date" faker:"-"`
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
	// Type the type of event
	Type RepoResponseType `json:"type" bson:"type" yaml:"type" faker:"-"`
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
	if o == nil {
		return toRepoResponseObjectNil(isavro, isoptional)
	}
	switch v := o.(type) {
	case nil:
		return toRepoResponseObjectNil(isavro, isoptional)
	case string, int, int8, int16, int32, int64, float32, float64, bool:
		if isavro && isoptional {
			return goavro.Union(avrotype, v)
		}
		return v
	case *string:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int8:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int16:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int32:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *int64:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float32:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *float64:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case *bool:
		if isavro && isoptional {
			if v == nil {
				return toRepoResponseObjectNil(isavro, isoptional)
			}
			pv := *v
			return goavro.Union(avrotype, pv)
		}
		return v
	case map[string]interface{}:
		return o
	case *map[string]interface{}:
		return v
	case map[string]string:
		return v
	case *map[string]string:
		return *v
	case *RepoResponse:
		return v.ToMap()
	case RepoResponse:
		return v.ToMap()
	case []string, []int64, []float64, []bool:
		return o
	case *[]string:
		return (*(o.(*[]string)))
	case *[]int64:
		return (*(o.(*[]int64)))
	case *[]float64:
		return (*(o.(*[]float64)))
	case *[]bool:
		return (*(o.(*[]bool)))
	case []interface{}:
		a := o.([]interface{})
		arr := make([]interface{}, 0)
		for _, av := range a {
			arr = append(arr, toRepoResponseObject(av, isavro, false, ""))
		}
		return arr

	case RepoResponseEventDate:
		vv := o.(RepoResponseEventDate)
		return vv.ToMap()
	case *RepoResponseEventDate:
		return map[string]interface{}{
			"agent.event_date": (*o.(*RepoResponseEventDate)).ToMap(),
		}
	case []RepoResponseEventDate:
		arr := make([]interface{}, 0)
		for _, i := range o.([]RepoResponseEventDate) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]RepoResponseEventDate:
		arr := make([]interface{}, 0)
		vv := o.(*[]RepoResponseEventDate)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case RepoResponseRepos:
		vv := o.(RepoResponseRepos)
		return vv.ToMap()
	case *RepoResponseRepos:
		return map[string]interface{}{
			"agent.repos": (*o.(*RepoResponseRepos)).ToMap(),
		}
	case []RepoResponseRepos:
		arr := make([]interface{}, 0)
		for _, i := range o.([]RepoResponseRepos) {
			arr = append(arr, i.ToMap())
		}
		return arr
	case *[]RepoResponseRepos:
		arr := make([]interface{}, 0)
		vv := o.(*[]RepoResponseRepos)
		for _, i := range *vv {
			arr = append(arr, i.ToMap())
		}
		return arr
	case RepoResponseType:
		if !isavro {
			return (o.(RepoResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(RepoResponseType)).String(),
		}
	case *RepoResponseType:
		if !isavro {
			return (o.(*RepoResponseType)).String()
		}
		return map[string]string{
			"agent.type": (o.(*RepoResponseType)).String(),
		}
	}
	panic("couldn't figure out the object type: " + reflect.TypeOf(o).String())
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

func (o *RepoResponse) setDefaults() {
	if o.Data == nil {
		o.Data = &emptyString
	}
	if o.Error == nil {
		o.Error = &emptyString
	}
	if o.Repos == nil {
		o.Repos = []RepoResponseRepos{}
	}

	o.GetID()
	o.GetRefID()
	o.Hash()
}

// GetID returns the ID for the object
func (o *RepoResponse) GetID() string {
	if o.ID == "" {
		// we will attempt to generate a consistent, unique ID from a hash
		o.ID = hash.Values("RepoResponse", o.CustomerID, o.RefType, o.GetRefID())
	}
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
	var dt interface{} = o.EventDate
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
	case RepoResponseEventDate:
		return datetime.DateFromEpoch(v.Epoch)
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
	return &datamodel.ModelTopicConfig{
		Key:               "uuid",
		Timestamp:         "event_date",
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
	return nil
}

var cachedCodecRepoResponse *goavro.Codec

// GetAvroCodec returns the avro codec for this model
func (o *RepoResponse) GetAvroCodec() *goavro.Codec {
	if cachedCodecRepoResponse == nil {
		c, err := GetRepoResponseAvroSchema()
		if err != nil {
			panic(err)
		}
		cachedCodecRepoResponse = c
	}
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
	o.setDefaults()
	return map[string]interface{}{
		"architecture":   toRepoResponseObject(o.Architecture, isavro, false, "string"),
		"customer_id":    toRepoResponseObject(o.CustomerID, isavro, false, "string"),
		"data":           toRepoResponseObject(o.Data, isavro, true, "string"),
		"distro":         toRepoResponseObject(o.Distro, isavro, false, "string"),
		"error":          toRepoResponseObject(o.Error, isavro, true, "string"),
		"event_date":     toRepoResponseObject(o.EventDate, isavro, false, "event_date"),
		"free_space":     toRepoResponseObject(o.FreeSpace, isavro, false, "long"),
		"go_version":     toRepoResponseObject(o.GoVersion, isavro, false, "string"),
		"hostname":       toRepoResponseObject(o.Hostname, isavro, false, "string"),
		"id":             toRepoResponseObject(o.ID, isavro, false, "string"),
		"integration_id": toRepoResponseObject(o.IntegrationID, isavro, false, "string"),
		"memory":         toRepoResponseObject(o.Memory, isavro, false, "long"),
		"message":        toRepoResponseObject(o.Message, isavro, false, "string"),
		"num_cpu":        toRepoResponseObject(o.NumCPU, isavro, false, "long"),
		"os":             toRepoResponseObject(o.OS, isavro, false, "string"),
		"ref_id":         toRepoResponseObject(o.RefID, isavro, false, "string"),
		"ref_type":       toRepoResponseObject(o.RefType, isavro, false, "string"),
		"repos":          toRepoResponseObject(o.Repos, isavro, false, "repos"),
		"request_id":     toRepoResponseObject(o.RequestID, isavro, false, "string"),
		"success":        toRepoResponseObject(o.Success, isavro, false, "boolean"),
		"type":           toRepoResponseObject(o.Type, isavro, false, "type"),
		"uuid":           toRepoResponseObject(o.UUID, isavro, false, "string"),
		"version":        toRepoResponseObject(o.Version, isavro, false, "string"),
		"hashcode":       toRepoResponseObject(o.Hashcode, isavro, false, "string"),
	}
}

// FromMap attempts to load data into object from a map
func (o *RepoResponse) FromMap(kv map[string]interface{}) {
	// if coming from db
	if id, ok := kv["_id"]; ok && id != "" {
		kv["id"] = id
	}
	if val, ok := kv["architecture"].(string); ok {
		o.Architecture = val
	} else {
		val := kv["architecture"]
		if val == nil {
			o.Architecture = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Architecture = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["customer_id"].(string); ok {
		o.CustomerID = val
	} else {
		val := kv["customer_id"]
		if val == nil {
			o.CustomerID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.CustomerID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["data"].(*string); ok {
		o.Data = val
	} else if val, ok := kv["data"].(string); ok {
		o.Data = &val
	} else {
		val := kv["data"]
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
	if val, ok := kv["distro"].(string); ok {
		o.Distro = val
	} else {
		val := kv["distro"]
		if val == nil {
			o.Distro = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Distro = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["error"].(*string); ok {
		o.Error = val
	} else if val, ok := kv["error"].(string); ok {
		o.Error = &val
	} else {
		val := kv["error"]
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
	if val, ok := kv["event_date"].(RepoResponseEventDate); ok {
		o.EventDate = val
	} else {
		val := kv["event_date"]
		if val == nil {
			o.EventDate = RepoResponseEventDate{}
		} else {
			o.EventDate = RepoResponseEventDate{}
			if m, ok := val.(map[interface{}]interface{}); ok {
				si := make(map[string]interface{})
				for k, v := range m {
					if key, ok := k.(string); ok {
						si[key] = v
					}
				}
				val = si
			}
			b, _ := json.Marshal(val)
			json.Unmarshal(b, &o.EventDate)

		}
	}
	if val, ok := kv["free_space"].(int64); ok {
		o.FreeSpace = val
	} else {
		val := kv["free_space"]
		if val == nil {
			o.FreeSpace = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.FreeSpace = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["go_version"].(string); ok {
		o.GoVersion = val
	} else {
		val := kv["go_version"]
		if val == nil {
			o.GoVersion = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.GoVersion = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["hostname"].(string); ok {
		o.Hostname = val
	} else {
		val := kv["hostname"]
		if val == nil {
			o.Hostname = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Hostname = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["id"].(string); ok {
		o.ID = val
	} else {
		val := kv["id"]
		if val == nil {
			o.ID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.ID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["integration_id"].(string); ok {
		o.IntegrationID = val
	} else {
		val := kv["integration_id"]
		if val == nil {
			o.IntegrationID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.IntegrationID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["memory"].(int64); ok {
		o.Memory = val
	} else {
		val := kv["memory"]
		if val == nil {
			o.Memory = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.Memory = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["message"].(string); ok {
		o.Message = val
	} else {
		val := kv["message"]
		if val == nil {
			o.Message = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Message = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["num_cpu"].(int64); ok {
		o.NumCPU = val
	} else {
		val := kv["num_cpu"]
		if val == nil {
			o.NumCPU = number.ToInt64Any(nil)
		} else {
			if tv, ok := val.(time.Time); ok {
				val = datetime.TimeToEpoch(tv)
			}
			o.NumCPU = number.ToInt64Any(val)
		}
	}
	if val, ok := kv["os"].(string); ok {
		o.OS = val
	} else {
		val := kv["os"]
		if val == nil {
			o.OS = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.OS = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_id"].(string); ok {
		o.RefID = val
	} else {
		val := kv["ref_id"]
		if val == nil {
			o.RefID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["ref_type"].(string); ok {
		o.RefType = val
	} else {
		val := kv["ref_type"]
		if val == nil {
			o.RefType = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RefType = fmt.Sprintf("%v", val)
		}
	}
	if val := kv["repos"]; val != nil {
		na := make([]RepoResponseRepos, 0)
		if a, ok := val.([]RepoResponseRepos); ok {
			na = append(na, a...)
		} else {
			if a, ok := val.([]interface{}); ok {
				for _, ae := range a {
					if av, ok := ae.(RepoResponseRepos); ok {
						na = append(na, av)
					} else {
						if badMap, ok := ae.(map[interface{}]interface{}); ok {
							ae = slice.ConvertToStringToInterface(badMap)
						}
						b, _ := json.Marshal(ae)
						var av RepoResponseRepos
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for repos field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else if a, ok := val.(primitive.A); ok {
				for _, ae := range a {
					if av, ok := ae.(RepoResponseRepos); ok {
						na = append(na, av)
					} else {
						b, _ := json.Marshal(ae)
						var av RepoResponseRepos
						if err := json.Unmarshal(b, &av); err != nil {
							panic("unsupported type for repos field entry: " + reflect.TypeOf(ae).String())
						}
						na = append(na, av)
					}
				}
			} else {
				fmt.Println(reflect.TypeOf(val).String())
				panic("unsupported type for repos field")
			}
		}
		o.Repos = na
	} else {
		o.Repos = []RepoResponseRepos{}
	}
	if o.Repos == nil {
		o.Repos = make([]RepoResponseRepos, 0)
	}
	if val, ok := kv["request_id"].(string); ok {
		o.RequestID = val
	} else {
		val := kv["request_id"]
		if val == nil {
			o.RequestID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.RequestID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["success"].(bool); ok {
		o.Success = val
	} else {
		val := kv["success"]
		if val == nil {
			o.Success = number.ToBoolAny(nil)
		} else {
			o.Success = number.ToBoolAny(val)
		}
	}
	if val, ok := kv["type"].(RepoResponseType); ok {
		o.Type = val
	} else {
		if em, ok := kv["type"].(map[string]interface{}); ok {
			ev := em["agent.type"].(string)
			switch ev {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "repo":
				o.Type = 6
			case "user":
				o.Type = 7
			}
		}
		if em, ok := kv["type"].(string); ok {
			switch em {
			case "enroll":
				o.Type = 0
			case "ping":
				o.Type = 1
			case "crash":
				o.Type = 2
			case "integration":
				o.Type = 3
			case "export":
				o.Type = 4
			case "project":
				o.Type = 5
			case "repo":
				o.Type = 6
			case "user":
				o.Type = 7
			}
		}
	}
	if val, ok := kv["uuid"].(string); ok {
		o.UUID = val
	} else {
		val := kv["uuid"]
		if val == nil {
			o.UUID = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.UUID = fmt.Sprintf("%v", val)
		}
	}
	if val, ok := kv["version"].(string); ok {
		o.Version = val
	} else {
		val := kv["version"]
		if val == nil {
			o.Version = ""
		} else {
			if m, ok := val.(map[string]interface{}); ok {
				val = pjson.Stringify(m)
			}
			o.Version = fmt.Sprintf("%v", val)
		}
	}
	o.setDefaults()
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
	args = append(args, o.Memory)
	args = append(args, o.Message)
	args = append(args, o.NumCPU)
	args = append(args, o.OS)
	args = append(args, o.RefID)
	args = append(args, o.RefType)
	args = append(args, o.Repos)
	args = append(args, o.RequestID)
	args = append(args, o.Success)
	args = append(args, o.Type)
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
				"type": map[string]interface{}{"fields": []interface{}{map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the date of the event", "type": "record", "name": "event_date"},
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
				"type": map[string]interface{}{"type": "array", "name": "repos", "items": map[string]interface{}{"type": "record", "name": "repos", "fields": []interface{}{map[string]interface{}{"type": "boolean", "name": "active", "doc": "the status of the repo determined by an Admin"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "repos.created_date", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}, map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"name": "offset", "doc": "the timezone offset from GMT", "type": "long"}}, "doc": "the creation date"}, "name": "created_date", "doc": "the creation date"}, map[string]interface{}{"type": "string", "name": "description", "doc": "the description of the repository"}, map[string]interface{}{"name": "language", "doc": "the programming language defined for the repository", "type": "string"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "repos.last_commit", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "commit_id", "doc": "the id of the latest commit"}, map[string]interface{}{"type": "string", "name": "url", "doc": "the url of the lastest commit"}, map[string]interface{}{"type": "string", "name": "message", "doc": "the commit message of the latest commit"}, map[string]interface{}{"type": map[string]interface{}{"type": "record", "name": "last_commit.created_date", "fields": []interface{}{map[string]interface{}{"type": "string", "name": "customer_id", "doc": "the customer id for the model instance"}, map[string]interface{}{"type": "long", "name": "epoch", "doc": "the date in epoch format"}, map[string]interface{}{"type": "string", "name": "id", "doc": "the primary key for the model instance"}, map[string]interface{}{"type": "long", "name": "offset", "doc": "the timezone offset from GMT"}, map[string]interface{}{"type": "string", "name": "ref_id", "doc": "the source system id for the model instance"}, map[string]interface{}{"type": "string", "name": "ref_type", "doc": "the source system identifier for the model instance"}, map[string]interface{}{"type": "string", "name": "rfc3339", "doc": "the date in RFC3339 format"}}, "doc": "the timestamp of the latest commit"}, "name": "created_date", "doc": "the timestamp of the latest commit"}, map[string]interface{}{"name": "author", "doc": "the author of the latest commit", "type": map[string]interface{}{"type": "record", "name": "last_commit.author", "fields": []interface{}{map[string]interface{}{"name": "name", "doc": "the author name", "type": "string"}, map[string]interface{}{"type": "string", "name": "email", "doc": "the email of the author"}, map[string]interface{}{"name": "avatar_url", "doc": "the avatar_url for the author", "type": "string"}}, "doc": "the author of the latest commit"}}}, "doc": "the most recent commit to the repo"}, "name": "last_commit", "doc": "the most recent commit to the repo"}, map[string]interface{}{"type": "string", "name": "name", "doc": "the name of the repository"}}, "doc": "the repos exported"}},
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
				"name": "type",
				"type": []interface{}{
					map[string]interface{}{
						"type":    "enum",
						"name":    "type",
						"symbols": []interface{}{"enroll", "ping", "crash", "integration", "export", "project", "repo", "user"},
					},
				},
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
					if tv.IsZero() {
						tv = time.Now() // if its still zero, use the ingest time
					}
					msg := eventing.Message{
						Encoding:  eventing.AvroEncoding,
						Key:       item.Key(),
						Value:     binary,
						Codec:     codec,
						Headers:   headers,
						Timestamp: tv,
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
