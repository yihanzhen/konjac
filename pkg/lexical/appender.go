package lexical

import (
	"strings"

	"github.com/yihanzhen/konjac/pkg/types"
)

type appendable interface {
	Append() AppendOption
	Write([]string) error
	Syntax([]string) error
}

type Appender struct {
	items          []appendable
	lexime         types.LeximeType
	contentLexime  types.LeximeType
	functionLexime types.LeximeType
	conjucation    types.ConjugationType
}

type AppendOption struct {
	Append         appendable
	SetLexime      types.LeximeType
	SetConjugation types.ConjugationType
}

type WriteResult struct {
	Sentence, Syntax string
}

func (a *Appender) ContentLexime() types.LeximeType {
	return nil
}

func (a *Appender) FunctionLexime() types.LeximeType {
	return nil
}

func (a *Appender) Append(o *AppendOption) {
	a.items = append(a.items, o.Append)
	a.lexime = o.SetLexime
	a.conjucation = o.SetConjugation
}

func (a *Appender) Write() (WriteResult, error) {
	var sentence, syntax []string
	var err error
	for _, item := range a.items {
		err = item.Write(sentence)
		if err != nil {
			return WriteResult{}, err
		}
		err = item.Syntax(sentence)
		if err != nil {
			return WriteResult{}, err
		}
	}
	return WriteResult{
		Sentence: strings.Join(sentence, ""),
		Syntax:   strings.Join(syntax, ""),
	}, nil
}
