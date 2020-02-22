package validator

import (
	"fmt"

	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func Absent(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.Absent == nil || !*rule.Absent {
		return true, ""
	}

	if node.IsFound() {
		return false, fmt.Sprintf("field '%s' must not be here", rule.Field)
	}

	return true, ""
}
