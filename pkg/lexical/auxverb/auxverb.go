package auxverb

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
	"github.com/yihanzhen/konjac/pkg/util"
)

type AuxVerb struct {
	checkState          func(*lexical.AppenderState) error
	allowedConjugations map[types.ConjugationType]bool
	auxVerbType         types.AuxVerbType
	Writing             string
}

var VerbPoliteMaker = &AuxVerb{
	checkState: func(s *lexical.AppenderState) error {
		if s.Lexime != types.Verb {
			return fmt.Errorf("got %v, want verb: %w", s.Lexime, errors.IllegalAppenderLeximeError)
		}
		if s.Conjugation != types.Conjuntive {
			return fmt.Errorf("got %v, want conjunctive: %w", s.Conjugation, errors.IllegalAppenderConjugationError)
		}
		return nil
	},
	allowedConjugations: util.EmptySetOf[types.ConjugationType](),
	auxVerbType:         types.VerbPoliteMaker,
	Writing:             "ます",
}

// Write implements lexical.appendable.
func (v *AuxVerb) Write(sentence []string) ([]string, error) {
	sentence = append(sentence, v.Writing)
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (v *AuxVerb) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, fmt.Sprintf("auxverb %q %q", types.VerbPoliteMaker, v.Writing))
	return syntax, nil
}

// Append implements lexical.appendable.
func (v *AuxVerb) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	if err := v.checkState(state); err != nil {
		return nil, fmt.Errorf("AuxVerb.Append: %w", err)
	}
	return &lexical.AppenderMutation{
		Append:    v,
		SetLexime: types.AuxVerb,
		SetAuxVerbState: &lexical.AuxVerbState{
			AuxVerbType: v.auxVerbType,
		},
	}, nil
}
