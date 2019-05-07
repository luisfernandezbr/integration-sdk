package jira

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jhaynie/go-gator/orm"

	"github.com/pinpt/go-common/flatten"
	"github.com/pinpt/integration-sdk/model/work/v1/work"
	"github.com/pinpt/integration-sdk/util"
)

var issueMap = map[string][]string{
	"identifier":     []string{"$.id"},
	"key":            []string{"$.key"},
	"fields_summary": []string{"$.fields_summary"},
	"issue_type":     []string{"$.fields_issuetype_name"},
	"project_id":     []string{"$.fields_project_id"},
	"labels":         []string{"$.fields_labels"},
	"parent_id":      []string{"$.fields_parent_id"},
	"status":         []string{"$.fields_status_name"},
	"resolution":     []string{"$.fields_resolution_name"},
	"duedate":        []string{"$.fields_duedate"},
	// "custom_fields":       []string{"$.customfields"},
	"created_at":          []string{"$.fields_created"},
	"ref_creator_user_id": []string{"$.fields_creator_accountid"},
	"issuetype_id":        []string{"$.fields_issuetype_id"},
	"combo":               []string{"concat($.fields_priority_name, $.fields_issuetype_id)"},
	"url":                 []string{"$.url"},
	"ref_reporter_user_id": []string{"$.fields_reporter_accountid"},
	"ref_assignee_user_id": []string{"$.fields_assignee_accountid"},
	"comp":                 []string{"$.fields_components"},
	"active":               []string{"$.fields_creator_active"},
	"username":             []string{"$.fields_creator_name"},
	"displayname":          []string{"$.fields_creator_displayname"},
	"email":                []string{"$.fields_creator_emailaddress"},
	"avatar_url":           []string{"$.fields_creator_avatarurls_48x48"},
	"priority_id":          []string{"$.fields_priority_id"},
}

func readerFromFile(filename string) (io.ReadCloser, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s. %v", filename, err)
	}
	var r io.ReadCloser = f
	if filepath.Ext(filename) == ".gz" {
		r, err = gzip.NewReader(f)
		if err != nil {
			f.Close()
			return nil, fmt.Errorf("error opening gzip %s. %v", filename, err)
		}
	}
	return r, nil
}

// LoadIssues will convert a legacy commit hashmap to a new sourcecode.Commit
func LoadIssues(fn string) ([]work.Issue, error) {
	reader, err := readerFromFile(fn)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var allissues []work.Issue
	if err := orm.Deserialize(reader, func(line json.RawMessage) error {
		var raw map[string]interface{}
		json.Unmarshal(line, &raw)
		raw, err = flatten.Flatten(raw)
		if err != nil {
			return err
		}
		var issue work.Issue
		if err := util.CreateObject(&issue, raw, issueMap); err != nil {
			b, _ := json.MarshalIndent(line, "", "  ")
			fmt.Println(string(b))
			return err
		}
		allissues = append(allissues, issue)
		return nil
	}); err != nil {
		return nil, err
	}
	return allissues, nil
}
