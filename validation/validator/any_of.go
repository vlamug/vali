package validator

import (
	"fmt"
	"strings"

	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func AnyOf(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if len(rule.AnyOf) == 0 {
		return true, ""
	}

	if !node.IsFound() {
		return true, ""
	}

	value, ok := util.ExtractStringValueFromYaml(node)
	if !ok {
		return true, ""
	}

	exists := false
	for _, option := range rule.AnyOf {
		if rule.CaseInsens != nil && *rule.CaseInsens {
			option = strings.ToLower(option)
			value = strings.ToLower(option)
		}

		if option == value {
			exists = true
			break
		}
	}
	if !exists {
		return false, fmt.Sprintf("value '%s' of field '%s' is not any of : %v", value, rule.Field, rule.AnyOf)
	}

	return true, ""
}
