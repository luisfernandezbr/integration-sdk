package github

import (
	"regexp"

	"github.com/pinpt/go-common/fileutil"

	"github.com/pinpt/integration-sdk/model/sourcecode/v1/sourcecode"
	"github.com/pinpt/integration-sdk/util"
)

var commitMap = map[string][]string{
	"id":            []string{"hash($.customer_id, $.sha, $.branch)"},
	"customer_id":   []string{"$.customer_id"},
	"ref_id":        []string{"$.sha"},
	"ref_type":      []string{"github"},
	"repo_id":       []string{"$.repo_id"},
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
	"id":              []string{"hash($.customer_id, $.sha, $.branch, $.filename)"},
	"customer_id":     []string{"$.customer_id"},
	"repo_id":         []string{"$.repo_id"},
	"ref_id":          []string{"$.sha"},
	"ref_type":        []string{"github"},
	"commit_id":       []string{"hash($.customer_id, $.sha, $.branch)"},
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

var changelogMap = map[string][]string{
	"id":            []string{"hash($.customer_id, $.sha, $.filename, $.user_id, $.date)"},
	"customer_id":   []string{"$.customer_id"},
	"repo_id":       []string{"$.repo_id"},
	"ref_id":        []string{"$.sha"},
	"ref_type":      []string{"github"},
	"sha":           []string{"$.sha"},
	"filename":      []string{"$.filename"},
	"language":      []string{"$.language"},
	"blanks":        []string{"$.blanks"},
	"comments":      []string{"$.comments"},
	"loc":           []string{"$.loc"},
	"sloc":          []string{"$.sloc"},
	"ordinal":       []string{"$.ordinal"},
	"complexity":    []string{"$.complexity"},
	"author_ref_id": []string{"hash($.customer_id, $.user_id)"},
	"date_ts":       []string{"$.date"},
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

// ConvertChangelog will convert a legacy commit activity into a changelog
func ConvertChangelog(kv map[string]interface{}) (*sourcecode.Changelog, error) {
	var changelog sourcecode.Changelog
	if err := util.CreateObject(&changelog, kv, changelogMap); err != nil {
		return nil, err
	}
	return &changelog, nil
}

// LoadCommitAndCommitDetailMap will load the appropriate files and merge them into a single
// map for commit and commit detail
func LoadCommitAndCommitDetailMap(dir string, repoid string, customerid string) (map[string]map[string]interface{}, error) {
	return util.JoinTable(dir, "commit", "$.id", "commit_detail", "$.id", map[string]interface{}{"customer_id": customerid, "repo_id": repoid})
}

// LoadChangelog will load the commit activity table and return
func LoadChangelog(dir string, repoid string, customerid string) (map[string]map[string]interface{}, error) {
	res := make(map[string]map[string]interface{})
	files, err := fileutil.FindFiles(dir, regexp.MustCompile("commit_activity\\.json(\\.gz)?$"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if err := util.StreamToMap(file, "$.id", res, true); err != nil {
			return nil, err
		}
	}
	for _, kv := range res {
		kv["customer_id"] = customerid
		kv["repo_id"] = repoid
	}
	return res, nil
}
