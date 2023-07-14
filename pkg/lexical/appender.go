package lexical

import (
	"strings"

	"github.com/yihanzhen/konjac/pkg/types"
)

type appendable interface {
	Append(*Appender) (AppendOption, error)
	Write([]string) ([]string, error)
	Syntax([]string) error
}

type Appender struct {
	items          []appendable
	lexime         types.LeximeType
	contentLexime  types.LeximeType
	functionLexime types.LeximeType
	conjucation    types.ConjugationType

	verbStat *VerbStat
}

type AppendOption struct {
	Append         appendable
	SetLexime      types.LeximeType
	SetConjugation types.ConjugationType
	SetVerbStat    *VerbStat
}

type VerbStat struct {
	ConjugationRule types.VerbConjugationRule
}

type WriteResult struct {
	Sentence, Syntax string
}

func (a *Appender) ContentLexime() types.LeximeType {
	return a.contentLexime
}

func (a *Appender) FunctionLexime() types.LeximeType {
	return a.functionLexime
}

func (a *Appender) Lexime() types.LeximeType {
	return a.lexime
}

func (a *Appender) VerbStat() *VerbStat {
	return a.verbStat
}

func (a *Appender) Append(o *AppendOption) {
	a.items = append(a.items, o.Append)
	if o.SetLexime != nil {
		a.lexime = o.SetLexime
		a.conjucation = nil
	}
	if o.SetConjugation != nil {
		a.conjucation = o.SetConjugation
	}
	if o.SetVerbStat != nil {
		a.verbStat = o.SetVerbStat
	}
}

func (a *Appender) Write() (WriteResult, error) {
	var sentence, syntax []string
	var err error
	for _, item := range a.items {
		sentence, err = item.Write(sentence)
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
