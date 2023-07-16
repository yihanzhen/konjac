package verb

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yihanzhen/konjac/pkg/lexical"
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

func TestAppend(t *testing.T) {
	cases := []struct {
		name         string
		verb         Verb
		state        *lexical.AppenderState
		wantMutation *lexical.AppenderMutation
		wantError    error
	}{
		{
			name: "success",
			verb: Verb{
				Writing:         "のむ",
				ConjugationRule: types.GroupOne,
			},
			wantMutation: &lexical.AppenderMutation{
				Append: Verb{
					Writing:         "のむ",
					ConjugationRule: types.GroupOne,
				},
				SetLexime: types.Verb,
				SetVerbState: &lexical.VerbState{
					ConjugationRule: types.GroupOne,
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotMut, gotErr := tc.verb.Append(tc.state)
			if !errors.Is(gotErr, tc.wantError) {
				t.Errorf("Verb.Append has unexpected error: want %v, got %v", tc.wantError, gotErr)
			}
			if gotErr == nil {
				if diff := cmp.Diff(gotMut, tc.wantMutation); diff != "" {
					t.Errorf("Verb.Append has unexpected result: (-want,+got):\n%s", diff)
				}
			}
		})
	}
}

func TestWrite(t *testing.T) {
	cases := []struct {
		name      string
		verb      Verb
		stc       []string
		wantStc   []string
		wantError error
	}{
		{
			name: "success",
			verb: Verb{
				Writing:         "のむ",
				ConjugationRule: types.GroupOne,
			},
			wantStc: []string{"のむ"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotStc, gotErr := tc.verb.Write(tc.stc)
			if !errors.Is(gotErr, tc.wantError) {
				t.Errorf("Verb.Write has unexpected error: want %v, got %v", tc.wantError, gotErr)
			}
			if gotErr == nil {
				if diff := cmp.Diff(gotStc, tc.wantStc); diff != "" {
					t.Errorf("Verb.Write has unexpected result: (-want,+got):\n%s", diff)
				}
			}
		})
	}
}

func TestSyntax(t *testing.T) {
	cases := []struct {
		name      string
		verb      Verb
		stx       []string
		wantStx   []string
		wantError error
	}{
		{
			name: "success",
			verb: Verb{
				Writing:         "のむ",
				ConjugationRule: types.GroupOne,
			},
			wantStx: []string{`group one verb "のむ"`},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotStc, gotErr := tc.verb.Syntax(tc.stx)
			if !errors.Is(gotErr, tc.wantError) {
				t.Errorf("Verb.Write has unexpected error: want %v, got %v", tc.wantError, gotErr)
			}
			if gotErr == nil {
				if diff := cmp.Diff(gotStc, tc.wantStx); diff != "" {
					t.Errorf("Verb.Write has unexpected result: (-want,+got):\n%s", diff)
				}
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
