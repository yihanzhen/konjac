package builder

import (
	"testing"

	"github.com/yihanzhen/konjac/pkg/lexical/auxverb"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestCanonical(t *testing.T) {
	b := NewBuilder()
	v, err := verb.NewVerb("食べる")
	if err != nil {
		t.Fatal(err)
	}
	b.Steps([]*AppendStep{
		NewAppendStep(v, verb.NewVerbConjugation(types.Conjuntive), auxverb.VerbPoliteMaker),
	})
	wr, err := b.Build()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("build result: %v", wr)
}
