package data

import (
	"regexp"
)

type Rule struct {
	Field  string   `yaml:"field"`
	Fields []string `yaml:"fields"`
	Items  Items    `yaml:"items"`

	Match      *string `yaml:"match"`
	MatchReRaw string  `yaml:"match_re"`
	MatchRe    *regexp.Regexp
	AnyOf      []string `yaml:"anyOf"`
	Required   *bool    `yaml:"required"`
	NotEmpty   *bool    `yaml:"notEmpty"`
	IsNumber   *bool    `yaml:"isNumber"`
	Absent     *bool    `yaml:"absent"`
	IsMap      *bool    `yaml:"isMap"`
	IsArray    *bool    `yaml:"isArray"`
	CaseInsens *bool    `yaml:"caseInsens"`
}

type Items []Rule

func (r *Rule) Compile() error {
	if r.MatchReRaw != "" {
		r.MatchRe = regexp.MustCompile(r.MatchReRaw)
	}

	return nil
}
