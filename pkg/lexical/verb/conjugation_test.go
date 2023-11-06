package verb

import (
	errorsutil "errors"
	"fmt"
	"testing"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestConjugate(t *testing.T) {
	cases := []struct {
		typ            types.ConjugationType
		rule           types.VerbConjugationRule
		verb           string
		wantConjugated string
		wantError      error
	}{
		{
			typ:            types.Conjuntive,
			rule:           types.GroupOne,
			verb:           "のむ",
			wantConjugated: "のみ",
		},
		{
			typ:            types.Conjuntive,
			rule:           types.GroupTwo,
			verb:           "たべる",
			wantConjugated: "たべ",
		},
		{
			typ:       types.Conjuntive,
			rule:      types.KaIrregular,
			verb:      "くる",
			wantError: errors.UnimplementedError,
		},
		{
			typ:       types.Conjuntive,
			rule:      types.SaIrregular,
			verb:      "する",
			wantError: errors.UnimplementedError,
		},
		{
			typ:            types.Irrealis,
			rule:           types.GroupOne,
			verb:           "のむ",
			wantConjugated: "のま",
		},
		{
			typ:            types.Irrealis,
			rule:           types.GroupTwo,
			verb:           "たべる",
			wantConjugated: "たべ",
		},
		{
			typ:       types.Irrealis,
			rule:      types.KaIrregular,
			verb:      "くる",
			wantError: errors.UnimplementedError,
		},
		{
			typ:       types.Irrealis,
			rule:      types.SaIrregular,
			verb:      "する",
			wantError: errors.UnimplementedError,
		},
		{
			typ:            types.Attributive,
			rule:           types.GroupOne,
			verb:           "のむ",
			wantConjugated: "のむ",
		},
		{
			typ:            types.Attributive,
			rule:           types.GroupTwo,
			verb:           "たべる",
			wantConjugated: "たべる",
		},
		{
			typ:       types.Attributive,
			rule:      types.KaIrregular,
			verb:      "くる",
			wantError: errors.UnimplementedError,
		},
		{
			typ:       types.Attributive,
			rule:      types.SaIrregular,
			verb:      "する",
			wantError: errors.UnimplementedError,
		},
		{
			typ:       types.Conjuntive,
			rule:      types.GroupOne,
			verb:      "not hiragana",
			wantError: errors.IllegalArgError,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s-%s-%s", tc.typ, tc.rule, tc.verb), func(t *testing.T) {
			gotConjugated, gotErr := verbConjugationWritable{
				typ:  tc.typ,
				rule: tc.rule,
			}.conjugate(tc.verb)
			if !errorsutil.Is(gotErr, tc.wantError) {
				t.Errorf("conjugate has unexpected error: got %v, want %v", gotErr, tc.wantError)
			}
			if gotErr != nil {
				if gotConjugated != tc.wantConjugated {
					t.Errorf("conjugate has unexpected result: got %v, want %v", gotConjugated, tc.wantConjugated)
				}
			}
		})
	}
}
