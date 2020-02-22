package validator

import (
	"github.com/vlamug/vali/validation/data"

	"github.com/smallfish/simpleyaml"
)

type ValidateFunc func(node *simpleyaml.Yaml, rule data.Rule) (bool, string)
