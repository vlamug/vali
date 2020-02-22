package validator

import (
	"fmt"

	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func IsArray(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.IsArray == nil || !*rule.IsArray {
		return true, ""
	}

	if !node.IsFound() {
		return true, ""
	}

	if node.IsArray() {
		return true, ""
	}

	return false, fmt.Sprintf("field '%s' is not an array", rule.Field)
}
