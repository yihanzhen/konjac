package adjective

import (
	"fmt"
	"strings"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

var irregulars = map[string]bool{
	"いい": true,
	"よい": true,
	"良い": true,
}

type AdjectiveConjugation struct {
	typ types.ConjugationType
	mod types.AdjectiveConjuntiveMode
}

// Write implements lexical.writable.
func (a AdjectiveConjugation) Write(sentence []string) ([]string, error) {
	if len(sentence) == 0 {
		return nil, fmt.Errorf("AdjectiveConjugation.Write: sentence is empty %w", errors.IllegalStateError)
	}
	conjugated, err := a.conjugate(sentence[len(sentence)-1])
	if err != nil {
		return nil, fmt.Errorf("AdjectiveConjugation.Write: %w, details: %v", errors.IllegalStateError, err)
	}
	return append(sentence[0:len(sentence)-1], conjugated), nil
}

// Syntax implements lexical.writable.
func (a AdjectiveConjugation) Syntax(syntax []string) ([]string, error) {
	return append(syntax, fmt.Sprintf("conjugate as %s", a.typ)), nil
}

// Append implements lexical.appendable.
func (a AdjectiveConjugation) Append(state *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	if state.Lexime != types.Adjective {
		return nil, fmt.Errorf("AdjectiveConjugation.Append: not appended to an adj: %w", errors.IllegalStateError)
	}
	return &lexical.AppenderMutation{
		Append: a,
	}, nil
}

func (a AdjectiveConjugation) conjugate(word string) (string, error) {
	wordRunes := []rune(word)
	if _, ok := irregulars[word]; ok {
		return "", errors.UnimplementedError
	}
	if !strings.HasSuffix(word, "い") {
		return "", fmt.Errorf("got %v, want ending in い: %w", word, errors.IllegalStateError)
	}

	if a.typ == types.Conjuntive && a.mod == types.Adjectival {
		return string(wordRunes[0 : len(wordRunes)-1]), nil
	}
	if a.typ == types.Conjuntive && a.mod == types.Adverbial {
		return string(wordRunes[0:len(wordRunes)-1]) + "く", nil
	}
	if a.typ == types.Conjuntive && a.mod == types.Completed {
		return string(wordRunes[0:len(wordRunes)-1]) + "かっ", nil
	}
	if a.typ == types.Conjuntive {
		return "", fmt.Errorf("got conjunctive but unknown mode %s: %w", a.mod, errors.IllegalArgError)
	}
	if a.typ == types.Irrealis {
		return string(wordRunes[0:len(wordRunes)-1]) + "く", nil
	}
	if a.typ == types.Conditional {
		return string(wordRunes[0:len(wordRunes)-1]) + "け", nil
	}
	return "", fmt.Errorf("got conjugation %s and mode %s: %w", a.typ, a.mod, errors.IllegalArgError)
}
