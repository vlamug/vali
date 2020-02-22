package validator

import (
	"fmt"

	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func NotEmpty(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	// if validator is not specified
	if rule.NotEmpty == nil || !*rule.NotEmpty {
		return true, ""
	}

	// if node is not found
	if !node.IsFound() || node.IsMap() || node.IsArray() {
		return true, ""
	}

	// extract string value
	if value, ok := util.ExtractStringValueFromYaml(node); ok && value != "" {
		return true, ""
	}

	return false, fmt.Sprintf("field '%s' must not be empty", rule.Field)
}
