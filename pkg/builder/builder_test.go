package builder

import (
	"testing"

	"github.com/yihanzhen/konjac/pkg/grammar"
	"github.com/yihanzhen/konjac/pkg/vocabulary"
)

func TestCanonical(t *testing.T) {
	b := NewBuilder()
	if err := b.AddVerb("食べる", vocabulary.AddVerbOption{}); err != nil {
		t.Fatal(err)
	}
	b.Steps([]*AppendStep{
		NewAppendStep(b.Get("食べる"), grammar.MakePolite),
	})
	wr, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("build result: %v", wr)
}
