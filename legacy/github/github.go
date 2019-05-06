package github

import (
	"github.com/pinpt/integration-sdk/model/sourcecode/v1/sourcecode"
	"github.com/pinpt/integration-sdk/util"
)

var commitMap = map[string][]string{
	"customer_id":   []string{"$.customer_id"},
	"ref_id":        []string{"$.id", "$.sha"},
	"ref_type":      []string{"github"},
	"sha":           []string{"$.sha", "$.id"},
	"message":       []string{"$.commit.message"},
	"url":           []string{"$.html_url"},
	"created_ts":    []string{"epoch($.commit.author.date)"},
	"author_ref_id": []string{"string($.author.id)", "hash($.customer_id, $.commit.author.email)"},
	"branch":        []string{"$.branch"},
	"additions":     []string{"$.additions"},
	"deletions":     []string{"$.deletions"},
	"files_changed": []string{"len($.files)"},
	"ordinal":       []string{"$.ordinal"},
}

var commitFileMap = map[string][]string{
	"customer_id":     []string{"$.customer_id"},
	"ref_id":          []string{"hash($.customer_id, $.id, $.filename)"},
	"ref_type":        []string{"github"},
	"commit_id":       []string{"$.id"},
	"filename":        []string{"$.filename"},
	"language":        []string{"$.language"},
	"additions":       []string{"$.additions"},
	"deletions":       []string{"$.deletions"},
	"blanks":          []string{"$.blanks"},
	"comments":        []string{"$.comments"},
	"status":          []string{"$.status"},
	"binary":          []string{"$.binary"},
	"loc":             []string{"$.loc"},
	"sloc":            []string{"$.sloc"},
	"ordinal":         []string{"$.ordinal"},
	"complexity":      []string{"$.complexity"},
	"excluded":        []string{"$.excluded"},
	"excluded_reason": []string{"$.excluded_reason"},
}

// ConvertCommit will convert a legacy commit hashmap to a new sourcecode.Commit
func ConvertCommit(kv map[string]interface{}) (*sourcecode.Commit, error) {
	var commit sourcecode.Commit
	if err := util.CreateObject(&commit, kv, commitMap); err != nil {
		return nil, err
	}
	return &commit, nil
}

// ConvertCommitFiles will convert a legacy commit detail files array to a new sourcecode.CommitFile
func ConvertCommitFiles(kv map[string]interface{}) ([]*sourcecode.CommitFile, error) {
	commitfiles := make([]*sourcecode.CommitFile, 0)
	if files, ok := kv["files"].([]interface{}); ok {
		for _, filekv := range files {
			newkv := util.CloneMap(kv)
			for k, v := range filekv.(map[string]interface{}) {
				newkv[k] = v
			}
			var commitfile sourcecode.CommitFile
			if err := util.CreateObject(&commitfile, newkv, commitFileMap); err != nil {
				return nil, err
			}
			commitfiles = append(commitfiles, &commitfile)
		}
	}

	return commitfiles, nil
}

// LoadCommitAndCommitDetailMap will load the appropriate files and merge them into a single
// map for commit and commit detail
func LoadCommitAndCommitDetailMap(dir string, customerid string) (map[string]map[string]interface{}, error) {
	return util.JoinTable(dir, "commit", "$.id", "commit_detail", "$.id", map[string]interface{}{"customer_id": customerid})
}
