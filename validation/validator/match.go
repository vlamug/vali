package validator

import (
	"fmt"
	"strings"

	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func Match(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.Match == nil {
		return true, ""
	}

	if !node.IsFound() {
		return true, ""
	}

	if node.IsMap() {
		return false, fmt.Sprintf("field '%s' does not match '%s', it's a map", rule.Field, *rule.Match)
	}

	if node.IsArray() {
		return false, fmt.Sprintf("field '%s' does not match '%s', it's an array", rule.Field, *rule.Match)
	}

	value, ok := util.ExtractStringValueFromYaml(node)
	if !ok {
		return true, ""
	}

	match := *rule.Match
	if rule.CaseInsens != nil && *rule.CaseInsens {
		match = strings.ToLower(match)
		value = strings.ToLower(value)
	}

	if match != value {
		return false, fmt.Sprintf("invalid '%s' field value: got '%s', expected '%s'", rule.Field, value, *rule.Match)
	}

	return true, ""
}
