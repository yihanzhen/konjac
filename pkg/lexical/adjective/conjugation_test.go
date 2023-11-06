package adjective

import (
	"errors"
	"fmt"
	"testing"

	"github.com/yihanzhen/konjac/pkg/types"
)

func TestConjugate(t *testing.T) {
	cases := []struct {
		typ            types.ConjugationType
		mod            types.AdjectiveConjuntiveMode
		adj            string
		wantError      error
		wantConjugated string
	}{
		{
			typ:            types.Irrealis,
			mod:            types.AdjectiveConjuntiveModeUnset,
			adj:            "安い",
			wantConjugated: "安く",
		},
		{
			typ:            types.Conjuntive,
			mod:            types.Adjectival,
			adj:            "安い",
			wantConjugated: "安",
		},
		{
			typ:            types.Conjuntive,
			mod:            types.Adverbial,
			adj:            "安い",
			wantConjugated: "安く",
		},
		{
			typ:            types.Conjuntive,
			mod:            types.Completed,
			adj:            "安い",
			wantConjugated: "安かっ",
		},
		{
			typ:            types.Conditional,
			adj:            "安い",
			wantConjugated: "安け",
		},
		{
			typ:            types.Attributive,
			adj:            "安い",
			wantConjugated: "安い",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", tc.typ, tc.mod, tc.adj), func(t *testing.T) {
			gotConjugated, gotErr := AdjectiveConjugation{
				typ: tc.typ,
				mod: tc.mod,
			}.conjugate(tc.adj)
			if !errors.Is(gotErr, tc.wantError) {
				t.Errorf("conjugate has unexpected error: got %v, want %v", gotErr, tc.wantError)
			}
			if gotErr == nil {
				if gotConjugated != tc.wantConjugated {
					t.Errorf("conjugate has unexpected result: got %v, want %v", gotConjugated, tc.wantConjugated)
				}
			}
		})
	}

}
