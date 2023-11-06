package noun

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

// Noun represents a noun.
type Noun struct {
	Writing string
}

// NewNoun creates a new adjective.
func NewNoun(writing string) Noun {
	return Noun{Writing: writing}
}

// Write implements lexical.appendable.
func (n Noun) Write(sentence []string) ([]string, error) {
	sentence = append(sentence, n.Writing)
	return sentence, nil
}

// Syntax implements lexical.appendable.
func (n Noun) Syntax(syntax []string) ([]string, error) {
	syntax = append(syntax, fmt.Sprintf("noun %q", n.Writing))
	return syntax, nil
}

// Append implements lexical.appendable.
func (n Noun) Append(appender *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return &lexical.AppenderMutation{
		Append:    n,
		SetLexime: types.Noun,
	}, nil
}
