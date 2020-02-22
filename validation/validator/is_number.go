package validator

import (
	"fmt"
	"strconv"

	"github.com/vlamug/vali/util"
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func IsNumber(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	// if validator is not specified
	if rule.IsNumber == nil || !*rule.IsNumber {
		return true, ""
	}

	// if node is not found
	if !node.IsFound() {
		return true, ""
	}

	// check first that it is number
	if util.IsNumberYaml(node) {
		return true, ""
	}

	// try then extract number from string
	value, ok := util.ExtractStringValueFromYaml(node)
	if ok {
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return false, fmt.Sprintf("field '%s' with value '%s' is a not number", rule.Field, value)
		} else {
			return true, ""
		}
	}

	return false, fmt.Sprintf("field '%s' is not a number", rule.Field)
}
