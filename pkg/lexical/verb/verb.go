package verb

import (
	"fmt"
	"strings"

	"github.com/yihanzhen/konjac/pkg/kana"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

// Verb represents a verb.
type Verb struct {
	Writing         string
	ConjugationRule types.VerbConjugationRule
}

// Write implements lexical.appendable.
func (v Verb) Write(sentence []string) ([]string, error) {
	sentence = append(sentence, v.Writing)
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (v Verb) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, fmt.Sprintf("%s verb %q", v.ConjugationRule, v.Writing))
	return syntax, nil
}

// Append implements lexical.appendable.
func (v Verb) Append(appender *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return &lexical.AppenderMutation{
		Append:    v,
		SetLexime: types.Verb,
		SetVerbState: &lexical.VerbState{
			ConjugationRule: v.ConjugationRule,
		},
	}, nil
}

// NewVerb creates a new verb.
func NewVerb(write string, opts ...NewVerbOption) (Verb, error) {
	if l := len(opts); l > 1 {
		return Verb{}, fmt.Errorf("lexical.NewVerb: too many NewVerbOptions. Got %v, want 0 or 1", l)
	}
	opt := InferConjugationRule
	if len(opts) > 0 {
		opt = opts[0]
	}
	v := Verb{Writing: write}
	if err := opt(&v); err != nil {
		return Verb{}, fmt.Errorf("lexical.NewVerb: %w", err)
	}
	return v, nil
}

// NewVerbOption is option to create a new Verb.
type NewVerbOption func(*Verb) error

// InferConjugationRule infers the conjugationRule from the v.write.
// This is the default option if no option is provided.
func InferConjugationRule(v *Verb) error {
	conjRule, err := inferConjugationRule(v.Writing)
	if err != nil {
		return fmt.Errorf("lexical.InferConjugationRule: %w", err)
	}
	v.ConjugationRule = conjRule
	return nil
}

// ForceGroupOne sets the conjugationRule to GroupOne.
func ForceGroupOne(v *Verb) error {
	v.ConjugationRule = types.GroupOne
	return nil
}

// ConjugationReference provides a write for inferring the conjugationRule
// of the verb. This is useful when the second to last character of the verb
// is written in kanjis, for example, 見る.
// It errors when v.write is not fully composed of hiraganas.
func ConjugationReference(ref string) func(*Verb) error {
	return func(v *Verb) error {
		cr, err := inferConjugationRule(ref)
		if err != nil {
			return fmt.Errorf("lexical.ConjugationReference: %w", err)
		}
		v.ConjugationRule = cr
		return nil
	}
}

func inferConjugationRule(write string) (types.VerbConjugationRule, error) {
	writeRunes := []rune(write)
	if suf := writeRunes[len(writeRunes)-1]; !kana.IsCol(suf, 2) {
		return types.ConjugationRuleUnset, fmt.Errorf("got suffix %c, want ending in the third column %q", suf, write)
	}
	if write == "くる" || write == "来る" {
		return types.KaIrregular, nil
	}
	if strings.HasSuffix(write, "する") {
		return types.SaIrregular, nil
	}
	if !strings.HasSuffix(write, "る") {
		return types.GroupOne, nil
	}

	secondToLastRune := writeRunes[len(writeRunes)-2]
	if kana.IsCol(secondToLastRune, 0) || kana.IsCol(secondToLastRune, 2) || kana.IsCol(secondToLastRune, 4) {
		return types.GroupOne, nil
	}
	if kana.IsCol(secondToLastRune, 1) || kana.IsCol(secondToLastRune, 3) {
		return types.GroupTwo, nil
	}
	return types.ConjugationRuleUnset, fmt.Errorf("got second to last rune %v, want hiragana", secondToLastRune)
}
