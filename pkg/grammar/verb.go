package grammar

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/auxverb"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
	"github.com/yihanzhen/konjac/pkg/types"
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

var MakePolite = &Rule{
	Name: "Make Polite",
	items: []lexical.Appendable{
		verb.NewVerbConjugation(types.Conjuntive),
		auxverb.VerbPoliteMaker,
	},
}
