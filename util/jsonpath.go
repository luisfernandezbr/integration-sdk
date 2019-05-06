package util

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/oliveagle/jsonpath"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
)

// Action is the signature for handling registered actions
type Action func(args ...interface{}) (interface{}, error)

var (
	actionRe = regexp.MustCompile("^(\\w+)\\((.*)\\)$")
	actions  = make(map[string]Action)
)

// RegisterAction will register a new action
func RegisterAction(name string, action Action) {
	actions[name] = action
}

func invokeAction(val string, o interface{}) (interface{}, error) {
	tok := actionRe.FindStringSubmatch(val)
	if tok != nil && len(tok) > 1 {
		actionName := tok[1]
		action := actions[actionName]
		if action == nil {
			return nil, fmt.Errorf("invalid action: %s", actionName)
		}
		args := []interface{}{}
		if len(tok) > 2 {
			for _, arg := range strings.Split(tok[2], ",") {
				arg = strings.TrimSpace(arg)
				if arg[0:1] != "@" && arg[0:1] != "$" {
					args = append(args, arg)
				} else {
					val, err := jsonpath.JsonPathLookup(o, arg)
					if err != nil {
						return nil, fmt.Errorf("error fetching json path: %v. %v", arg, err)
					}
					args = append(args, val)
				}
			}
		}
		return action(args...)
	}
	if val[0:1] != "@" && val[0:1] != "$" {
		return val, nil
	}
	return jsonpath.JsonPathLookup(o, val)
}

func init() {
	RegisterAction("epoch", func(args ...interface{}) (interface{}, error) {
		return datetime.ISODateToEpoch(args[0].(string))
	})
	RegisterAction("hash", func(args ...interface{}) (interface{}, error) {
		return hash.Values(args...), nil
	})
	RegisterAction("string", func(args ...interface{}) (interface{}, error) {
		return fmt.Sprintf("%v", args[0]), nil
	})
	RegisterAction("len", func(args ...interface{}) (interface{}, error) {
		if val, ok := args[0].([]interface{}); ok {
			return len(val), nil
		}
		if val, ok := args[0].(map[string]interface{}); ok {
			return len(val), nil
		}
		return 0, nil
	})
}
