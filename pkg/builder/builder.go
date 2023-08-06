package builder

import (
	"fmt"
	"sync"

	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/vocabulary"
)

type Builder struct {
	*vocabulary.Vocabulary
	appender    *lexical.Appender
	appendSteps []*AppendStep
	stepOnce    sync.Once
}

func NewBuilder() *Builder {
	return &Builder{
		Vocabulary: vocabulary.NewVocabulary(),
		appender:   lexical.NewAppender(),
	}
}

type AppendStep struct {
	items []lexical.Appendable
}

func NewAppendStep(items ...lexical.Appendable) *AppendStep {
	return &AppendStep{
		items: items,
	}
}

func (b *Builder) Steps(appendSteps []*AppendStep) {
	b.stepOnce.Do(func() {
		b.appendSteps = appendSteps
	})
}

func (b *Builder) Build() (lexical.WriteResult, error) {
	for _, step := range b.appendSteps {
		for _, item := range step.items {
			if err := b.appender.Append(item); err != nil {
				return lexical.WriteResult{}, fmt.Errorf("Builder.Build: %w", err)
			}
		}
	}

	wr, err := b.appender.WriteResult()
	if err != nil {
		return lexical.WriteResult{}, fmt.Errorf("Builder.Build: %w", err)
	}
	return wr, nil
}
