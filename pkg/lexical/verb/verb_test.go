package verb

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestNewVerb(t *testing.T) {
	cases := []struct {
		name      string
		write     string
		opt       NewVerbOption
		wantVerb  *Verb
		wantError bool
	}{
		{
			name:  "no opt inferred group one",
			write: "のむ",
			wantVerb: &Verb{
				Writing:         "のむ",
				ConjugationRule: types.GroupOne,
			},
		},
		{
			name:  "no opt inferred group one",
			write: "食べる",
			wantVerb: &Verb{
				Writing:         "食べる",
				ConjugationRule: types.GroupTwo,
			},
		},
		{
			name:  "no opt inferred ka irregular",
			write: "くる",
			wantVerb: &Verb{
				Writing:         "くる",
				ConjugationRule: types.KaIrregular,
			},
		},
		{
			name:  "no opt inferred sa irregular",
			write: "見学する",
			wantVerb: &Verb{
				Writing:         "見学する",
				ConjugationRule: types.SaIrregular,
			},
		},
		{
			name:  "infer inferred group one",
			write: "のむ",
			opt:   InferConjugationRule,
			wantVerb: &Verb{
				Writing:         "のむ",
				ConjugationRule: types.GroupOne,
			},
		},
		{
			name:  "force group one",
			write: "入る",
			opt:   ForceGroupOne,
			wantVerb: &Verb{
				Writing:         "入る",
				ConjugationRule: types.GroupOne,
			},
		},
		{
			name:  "conjugation reference",
			write: "見る",
			opt:   ConjugationReference("みる"),
			wantVerb: &Verb{
				Writing:         "見る",
				ConjugationRule: types.GroupTwo,
			},
		},
		{
			name:      "error unable to infer",
			write:     "見る",
			wantError: true,
		},
		{
			name:      "error reference not all hiragana",
			write:     "見る",
			opt:       ConjugationReference("見る"),
			wantError: true,
		},
		{
			name:      "error reference incorrect write",
			write:     "見ら",
			wantError: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := NewVerb(tc.write, nilToEmpty(tc.opt)...)
			if err != nil {
				if !tc.wantError {
					t.Errorf("got error %v, want success", err)
				}
				return
			}
			if tc.wantError {
				t.Errorf("got verb %v, want error", v)
			}
			if diff := cmp.Diff(tc.wantVerb, &v, cmp.Comparer(func(a, b types.VerbConjugationRule) bool {
				return a == b
			})); diff != "" {
				t.Errorf("unexpected verb: (-want,+got)\n%s", diff)
			}
		})
	}
}

func nilToEmpty(opt NewVerbOption) []NewVerbOption {
	if opt == nil {
		return []NewVerbOption{}
	}
	return []NewVerbOption{opt}
}
