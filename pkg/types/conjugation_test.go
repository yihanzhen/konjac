package types

import "testing"

func TestConjugate(t *testing.T) {
	cases := []struct {
		conj       ConjugationType
		wantString string
	}{
		{
			conj:       Irrealis,
			wantString: "irrealis",
		},
		{
			conj:       Conjuntive,
			wantString: "conjunctive",
		},
		{
			conj:       Attributive,
			wantString: "attributive",
		},
		{
			conj:       Terminal,
			wantString: "terminal",
		},
		{
			conj:       Volitional,
			wantString: "volitional",
		},
		{
			conj:       Imperative,
			wantString: "imperative",
		},
		{
			conj:       Conditional,
			wantString: "conditional",
		},
	}
	for _, tc := range cases {
		t.Run(tc.wantString, func(t *testing.T) {
			if gotString := tc.conj.String(); gotString != tc.wantString {
				t.Errorf("String() has unexpected output: got %s, want %s", gotString, tc.wantString)
			}
		})
	}
}
