package lexical

import (
	"fmt"
	"strings"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/types"
)

type Appendable interface {
	Append(*AppenderState) (*AppenderMutation, error)
}

func AdditionalAppendables(items ...Appendable) func(*Appender) error {
	return func(a *Appender) error {
		for _, item := range items {
			if err := a.Append(item); err != nil {
				return err
			}
		}
		return nil
	}
}

type writable interface {
	Write([]string) ([]string, error)
	Syntax([]string) ([]string, error)
}

// AppenderState is the current state of elements in the Appender.
type AppenderState struct {
	Lexime         types.LeximeType
	ContentLexime  types.LeximeType
	FunctionLexime types.LeximeType
	Conjugation    types.ConjugationType
	VerbState      *VerbState
}

// Appender is used to append words to construct a sentence.
type Appender struct {
	items []writable
	state *AppenderState
}

func NewAppender() *Appender {
	return &Appender{
		state: &AppenderState{},
	}
}

func (a *Appender) GetState() AppenderState {
	return *a.state
}

// AppenderMutation is the mutation to the Appender after appending a word.
type AppenderMutation struct {
	Append           writable
	SetAppenderState *AppenderState
	SetLexime        types.LeximeType
	SetConjugation   types.ConjugationType
	SetVerbState     *VerbState
	SetAuxVerbState  *AuxVerbState
	CustomMutates    []func(*Appender) error
}

// VerbState is the extra information of the verb if current head of the Appender is a verb.
type VerbState struct {
	ConjugationRule types.VerbConjugationRule
}

// AuxVerbState is the extra information of the auxverb if the current head of the Appender is an auxverb.
type AuxVerbState struct {
	types.AuxVerbType
}

// WriteResult is the result of Write.
type WriteResult struct {
	Sentence, Syntax string
}

// Append appends a word to Appender.
func (a *Appender) Append(item Appendable) error {
	if a.state == nil {
		a.state = &AppenderState{}
	}
	m, err := item.Append(a.state)
	if err != nil {
		return fmt.Errorf("Appnder.Append: %w", err)
	}
	a.items = append(a.items, m.Append)
	if m.SetAppenderState != nil {
		a.state = m.SetAppenderState
	}
	if m.SetLexime != nil {
		a.state.Lexime = m.SetLexime
		a.state.Conjugation = nil
		a.state.VerbState = nil
		if m.SetLexime.IsContent() {
			a.state.ContentLexime = m.SetLexime
		} else {
			a.state.FunctionLexime = m.SetLexime
		}
	}
	if m.SetConjugation != nil {
		a.state.Conjugation = m.SetConjugation
	}
	if m.SetVerbState != nil {
		if m.SetLexime != types.Verb {
			return fmt.Errorf("Appender.Append: cannot SetVerbState when SetLexime != Verb: %w", errors.IllegalArgError)
		}
		a.state.VerbState = m.SetVerbState
	}
	for _, cm := range m.CustomMutates {
		if err := cm(a); err != nil {
			return fmt.Errorf("Appender.Append: cannot apply CustomMutate: %w", err)
		}
	}
	return nil
}

// Write writes the items in Appender to a sentence.
func (a *Appender) WriteResult() (WriteResult, error) {
	var sentence, syntax []string
	var err error
	for _, item := range a.items {
		sentence, err = item.Write(sentence)
		if err != nil {
			return WriteResult{}, fmt.Errorf("Appender.Write: %w", err)
		}
		syntax, err = item.Syntax(syntax)
		if err != nil {
			return WriteResult{}, fmt.Errorf("Appender.Write: %w", err)
		}
	}
	return WriteResult{
		Sentence: strings.Join(sentence, ""),
		Syntax:   strings.Join(syntax, ";"),
	}, nil
}
