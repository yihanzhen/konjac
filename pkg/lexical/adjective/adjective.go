package adjective

import (
	"fmt"
	"strings"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

// Adjective represents an adjective.
type Adjective struct {
	Writing string
}

// NewAdjective creates a new adjective.
func NewAdjective(writing string) (Adjective, error) {
	if !strings.HasSuffix(writing, "い") {
		return Adjective{}, fmt.Errorf("adjective.NewAdjective: adjectives must end with い: %w", errors.IllegalArgError)
	}
	return Adjective{Writing: writing}, nil
}

// Write implements lexical.appendable.
func (a Adjective) Write(sentence []string) ([]string, error) {
	sentence = append(sentence, a.Writing)
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (a Adjective) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, fmt.Sprintf("adjective %q", a.Writing))
	return syntax, nil
}

// Append implements lexical.appendable.
func (a Adjective) Append(appender *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return &lexical.AppenderMutation{
		Append:    a,
		SetLexime: types.Adjective,
	}, nil
}
