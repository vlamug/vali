package validation

import (
	"fmt"
	"strings"

	"github.com/vlamug/vali/validation/data"
	"github.com/vlamug/vali/validation/validator"

	"github.com/smallfish/simpleyaml"
)

type Runner struct {
	rules      []data.Rule
	validators []validator.ValidateFunc
	report     data.Report
}

func NewRunner(rules []data.Rule, validators []validator.ValidateFunc, report data.Report) *Runner {
	return &Runner{rules: rules, validators: validators, report: report}
}

func (r *Runner) Run(node *simpleyaml.Yaml) (data.Report, error) {
	for _, rule := range r.rules {
		for _, msg := range r.applyRule(node, rule) {
			r.report.Add(msg)
		}
	}

	return r.report, nil
}

// applyRule runs validation rule
func (r *Runner) applyRule(node *simpleyaml.Yaml, rule data.Rule) []string {
	// first take needed node by path
	subNode := extractNodeByPath(node, rule.Field)

	// if there is no any array by specified path
	if len(rule.Items) == 0 {
		return r.applyValidators(subNode, rule)
	}

	size, err := subNode.GetArraySize()
	if err != nil {
		return nil
	}

	var msgs []string
	for i := 0; i < size; i++ {
		// get over each node
		node := subNode.GetIndex(i)
		// make context from all fields of node to use it in error message
		context := makeContext(node)

		// apply sub rules
		for _, item := range rule.Items {
			// apply sub rules and collect message if rule fails
			for _, msg := range r.applyRule(node, item) {
				msgs = append(msgs, fmt.Sprintf("%s, context: %s", msg, context))
			}
		}
	}

	return msgs
}

// applyValidators applies validators one by one and return error message with first failed validation
func (r *Runner) applyValidators(node *simpleyaml.Yaml, rule data.Rule) []string {
	for _, v := range r.validators {
		if isValid, message := v(node, rule); !isValid {
			return []string{message}
		}
	}

	return nil
}

// makeContext takes all fields of node and make context from them
func makeContext(node *simpleyaml.Yaml) string {
	m, err := node.Map()
	if err != nil {
		return ""
	}

	var context []string
	for k, v := range m {
		context = append(context, fmt.Sprintf("%s:%v", k, v))
	}

	return fmt.Sprintf("{%s}", strings.Join(context, ","))
}

// extractNodeByPath extracts node by specified path
func extractNodeByPath(node *simpleyaml.Yaml, path string) *simpleyaml.Yaml {
	subNode := node
	for _, f := range strings.Split(path, ".") {
		subNode = subNode.Get(f)
	}

	return subNode
}
