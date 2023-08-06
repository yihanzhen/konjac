package vocabulary

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
)

type Vocabulary struct {
	words map[string]lexical.Appendable
}

func NewVocabulary() *Vocabulary {
	return &Vocabulary{
		words: map[string]lexical.Appendable{},
	}
}

func (v *Vocabulary) AddVerb(writing string, opt AddVerbOption) error {
	vb, err := verb.NewVerb(writing, opt.NewVerbOptions...)
	if err != nil {
		return fmt.Errorf("Vocabulary.NewVerb: %w", err)
	}
	v.words[writing] = vb
	for _, w := range opt.AlternativeWrites {
		v.words[w] = vb
	}
	return nil
}

type AddVerbOption struct {
	NewVerbOptions    []verb.NewVerbOption
	AlternativeWrites []string
}

func (v *Vocabulary) Get(write string) lexical.Appendable {
	w, exists := v.words[write]
	if !exists {
		return unknownWord{input: write}
	}
	return w
}

type unknownWord struct {
	input string
}

// Write implements lexical.appendable.
func (u unknownWord) Write(sentence []string) ([]string, error) {
	return sentence, fmt.Errorf("unknownWord.Write: should never be called: %w", errors.IllegalStateError)
}

// Syntax implements lexical.appendable.
func (u unknownWord) Syntax(syntax []string) ([]string, error) {
	return syntax, fmt.Errorf("unknownWord.Syntax: should never be called: %w", errors.IllegalStateError)
}

// Append implements lexical.appendable.
func (u unknownWord) Append(appender *lexical.AppenderState) (*lexical.AppenderMutation, error) {
	return nil, fmt.Errorf("%w: got %v", errors.UnknownWordError, u.input)
}
