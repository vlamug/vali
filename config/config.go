package config

import (
	"io/ioutil"

	"github.com/vlamug/vali/validation/data"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Rules []data.Rule `yaml:"rules"`
}

func MakeConfigFromFile(path string) (*Config, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(raw, cfg); err != nil {
		return nil, err
	}

	rules, err := compileRules(cfg.Rules)
	if err != nil {
		return nil, err
	}
	cfg.Rules = rules

	return cfg, nil
}

func compileRules(rules []data.Rule) ([]data.Rule, error) {
	compileRuleItems := func(rule data.Rule) (data.Rule, error) {
		if len(rule.Items) > 0 {
			items, err := compileRules(rule.Items)
			if err != nil {
				return rule, err
			}
			rule.Items = items
		}

		return rule, nil
	}

	var compiledRules []data.Rule
	for i := range rules {
		rule := rules[i]

		if err := rule.Compile(); err != nil {
			return nil, err
		}

		if len(rule.Fields) == 0 {
			rule, err := compileRuleItems(rule)
			if err != nil {
				return nil, err
			}

			compiledRules = append(compiledRules, rule)
		} else {
			for _, field := range rule.Fields {
				rule := rules[i]
				rule.Field = field
				rule.Fields = nil

				rule, err := compileRuleItems(rule)
				if err != nil {
					return nil, err
				}
				compiledRules = append(compiledRules, rule)
			}
		}

	}

	return compiledRules, nil
}
