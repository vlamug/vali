package validator

import (
	"fmt"

	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func MatchRe(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.MatchRe == nil {
		return true, ""
	}

	if !node.IsFound() {
		return true, ""
	}

	if node.IsMap() {
		return false, fmt.Sprintf("field '%s' does not match regular expression '%s', it's a map", rule.Field, rule.MatchRe.String())
	}

	if node.IsArray() {
		return false, fmt.Sprintf("field '%s' does not match regular expression '%s', it's an array", rule.Field, rule.MatchRe.String())
	}

	value, ok := util.ExtractStringValueFromYaml(node)
	if !ok {
		return true, ""
	}

	if !rule.MatchRe.MatchString(value) {
		return false, fmt.Sprintf("field '%s' with value '%s' does not match regular expression '%s'", rule.Field, value, rule.MatchRe.String())
	}

	return true, ""
}
