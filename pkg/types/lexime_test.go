package types

import "testing"

func TestLexime(t *testing.T) {
	cases := []struct {
		lex        LeximeType
		wantString string
	}{
		{
			lex:        Verb,
			wantString: "verb",
		},
		{
			lex:        Noun,
			wantString: "noun",
		},
		{
			lex:        Adjective,
			wantString: "adjective",
		},
		{
			lex:        AdjNoun,
			wantString: "adjectival noun",
		},
		{
			lex:        Adverb,
			wantString: "adverb",
		},
		{
			lex:        Particle,
			wantString: "particle",
		},
		{
			lex:        AuxVerb,
			wantString: "auxiliary verb",
		},
	}
	for _, tc := range cases {
		t.Run(tc.wantString, func(t *testing.T) {
			if gotString := tc.lex.String(); gotString != tc.wantString {
				t.Errorf("String() has unexpected output: got %s, want %s", gotString, tc.wantString)
			}
		})
	}
}
