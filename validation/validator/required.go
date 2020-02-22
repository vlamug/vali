package validator

import (
	"fmt"

	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

func Required(node *simpleyaml.Yaml, rule data.Rule) (bool, string) {
	if rule.Required != nil && *rule.Required == true && !node.IsFound() {
		return false, fmt.Sprintf("field '%s' is required", rule.Field)
	}

	return true, ""
}
