package validator

import (
	"fmt"

	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func IsMap(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.IsMap == nil || !*rule.IsMap {
		return true, ""
	}

	if !node.IsFound() {
		return true, ""
	}

	if node.IsMap() {
		return true, ""
	}

	return false, fmt.Sprintf("field '%s' is not a map", rule.Field)
}
