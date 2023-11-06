package grammar

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/lexical"
)

type Rule struct {
	Name  string
	items []lexical.Appendable
}

func (r *Rule) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return &lexical.AppenderMutation{
		Append:        r,
		CustomMutates: []func(*lexical.Appender) error{lexical.AdditionalAppendables(append(r.items, RuleEnder)...)},
	}, nil
}

// Write implements lexical.appendable.
func (r *Rule) Write(sentence []string) ([]string, error) {
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (r *Rule) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, fmt.Sprintf("Rule(%v){", r.Name))
	return syntax, nil
}

type ConditionalAppendable struct {
	cond func(*lexical.AppenderState) (lexical.Appendable, error)
}

func (c *ConditionalAppendable) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	a, err := c.cond(state)
	if err != nil {
		return nil, fmt.Errorf("ConditionalAppendable.Append: %w", err)
	}
	return a.Append(state)
}

type ruleEnder struct{}

var RuleEnder = ruleEnder{}

func (r ruleEnder) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return &lexical.AppenderMutation{
		Append: r,
	}, nil
}

// Write implements lexical.appendable.
func (r ruleEnder) Write(sentence []string) ([]string, error) {
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (r ruleEnder) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, "}")
	return syntax, nil
}
