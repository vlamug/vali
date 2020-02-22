package util

import (
	"fmt"
	"strconv"

	"github.com/smallfish/simpleyaml"
)

func ExtractStringValueFromYaml(yaml *simpleyaml.Yaml) (string, bool) {
	if value, err := yaml.String(); err == nil {
		return value, true
	} else if value, err := yaml.Int(); err == nil {
		return strconv.Itoa(value), true
	} else if value, err := yaml.Bool(); err == nil {
		strValue := "false"
		if value {
			strValue = "true"
		}

		return strValue, true
	} else if value, err := yaml.Float(); err == nil {
		return fmt.Sprintf("%g", value), true
	}

	return "", false
}

func IsNumberYaml(yaml *simpleyaml.Yaml) bool {
	if _, err := yaml.Int(); err == nil {
		return true
	} else if _, err := yaml.Float(); err == nil {
		return true
	}

	return false
}
