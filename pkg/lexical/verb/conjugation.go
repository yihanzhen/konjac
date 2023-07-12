package verb

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/kana"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

type VerbConjugation struct {
	conjugationType types.ConjugationType
	conjugationRule types.VerbConjugationRule
}

func NewVerbConjugation(conjugationType types.ConjugationType) *VerbConjugation {
	return &VerbConjugation{
		conjugationType: conjugationType,
	}
}

// Write implements lexical.appendable.
func (v *VerbConjugation) Write(sentence []string) ([]string, error) {
	if len(sentence) == 0 {
		return nil, fmt.Errorf("VerbConjugation.Write: sentence is empty %w", errors.IllegalStateError)
	}
	if v.conjugationRule == nil {
		return nil, fmt.Errorf("VerbConjugation.Write: conjugationRule unset: %w", errors.IllegalStateError)
	}
	conjugated, err := v.conjugate(sentence[len(sentence)-1], v.conjugationRule)
	if err != nil {
		return nil, fmt.Errorf("VerbConjugation.Write: %w, details: %v", errors.IllegalStateError, err)
	}
	return append(sentence[0:len(sentence)-1], conjugated), nil

}

// Append implements lexical.appendable.
func (v *VerbConjugation) Append(appender lexical.Appender) (lexical.AppendOption, error) {
	if appender.Lexime() != types.Verb {
		return lexical.AppendOption{}, fmt.Errorf("VerbConjugation.Append: not appended to a verb: %w", errors.IllegalStateError)
	}
	v.conjugationRule = appender.VerbStat().ConjugationRule
	return lexical.AppendOption{
		SetConjugation: v.conjugationType,
	}, nil
}

func groupOneConjugate(verb string, conjugationType types.ConjugationType) (string, error) {
	switch conjugationType {
	case types.Conjuntive:
		verb, err := kana.LastRuneToCol(verb, 1)
		if err != nil {
			return "", err
		}
		return verb, nil
	default:
		return "", errors.UnimplementedError
	}
}

func groupTwoConjugate(verb string, conjugationType types.ConjugationType) (string, error) {
	return kana.TrimLastRune(verb)
}

func (v *VerbConjugation) conjugate(verb string, conjugationRule types.VerbConjugationRule) (string, error) {
	var err error
	switch conjugationRule {
	case types.GroupOne:
		verb, err = groupOneConjugate(verb, v.conjugationType)

	case types.GroupTwo:
		verb, err = groupTwoConjugate(verb, v.conjugationType)
	default:
		return "", errors.UnimplementedError

	}
	if err != nil {
		return "", fmt.Errorf("conjugationType: %s, verbConjugationRule: %s: %w", v.conjugationType, conjugationRule, err)
	}
	return verb, nil
}
