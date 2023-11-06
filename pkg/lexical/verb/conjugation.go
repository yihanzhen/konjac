package verb

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/kana"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

type VerbConjugation struct {
	typ types.ConjugationType
}

type verbConjugationWritable struct {
	typ  types.ConjugationType
	rule types.VerbConjugationRule
}

func NewVerbConjugation(conjugationType types.ConjugationType) VerbConjugation {
	return VerbConjugation{typ: conjugationType}
}

// Write implements lexical.writable.
func (v verbConjugationWritable) Write(sentence []string) ([]string, error) {
	if len(sentence) == 0 {
		return nil, fmt.Errorf("VerbConjugation.Write: sentence is empty %w", errors.IllegalStateError)
	}
	if v.rule == nil {
		return nil, fmt.Errorf("VerbConjugation.Write: conjugationRule unset: %w", errors.IllegalStateError)
	}
	conjugated, err := v.conjugate(sentence[len(sentence)-1])
	if err != nil {
		return nil, fmt.Errorf("VerbConjugation.Write: %w, details: %v", errors.IllegalStateError, err)
	}
	return append(sentence[0:len(sentence)-1], conjugated), nil
}

// Syntax implements lexical.writable.
func (v verbConjugationWritable) Syntax(syntax []string) ([]string, error) {
	return append(syntax, fmt.Sprintf("conjugate as %s", v.typ)), nil
}

// Append implements lexical.appendable.
func (v VerbConjugation) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	if state.Lexime != types.Verb {
		return nil, fmt.Errorf("VerbConjugation.Append: not appended to a verb: %w", errors.IllegalStateError)
	}
	if state.VerbState == nil {
		return nil, fmt.Errorf("VerbConjugation.Append: VerbState does not exist: %w", errors.IllegalStateError)
	}
	return &lexical.AppenderMutation{
		Append: verbConjugationWritable{
			typ:  v.typ,
			rule: state.VerbState.ConjugationRule,
		},
		SetConjugation: types.ConjugationType(v.typ),
	}, nil
}

func groupOneConjugate(verb string, conjugationType types.ConjugationType) (string, error) {
	switch conjugationType {
	case types.Irrealis:
		return kana.LastRuneToCol(verb, 0)
	case types.Conjuntive:
		return kana.LastRuneToCol(verb, 1)
	case types.Attributive:
		return verb, nil
	default:
		return "", errors.UnimplementedError
	}
}

func groupTwoConjugate(verb string, conjugationType types.ConjugationType) (string, error) {
	switch conjugationType {
	case types.Irrealis:
		return kana.TrimLastRune(verb)
	case types.Conjuntive:
		return kana.TrimLastRune(verb)
	case types.Attributive:
		return verb, nil
	default:
		return "", errors.UnimplementedError
	}
}

func (v verbConjugationWritable) conjugate(verb string) (string, error) {
	var err error
	switch v.rule {
	case types.GroupOne:
		verb, err = groupOneConjugate(verb, v.typ)

	case types.GroupTwo:
		verb, err = groupTwoConjugate(verb, v.typ)
	default:
		return "", errors.UnimplementedError

	}
	if err != nil {
		return "", fmt.Errorf("conjugationType: %s, verbConjugationRule: %s: %w", v.typ, v.rule, err)
	}
	return verb, nil
}
