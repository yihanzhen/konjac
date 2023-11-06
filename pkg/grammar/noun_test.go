package grammar

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/testutil"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestNounRules(t *testing.T) {
	cases := []struct {
		name            string
		appender        *lexical.Appender
		rule            *Rule
		wantAppendState lexical.AppenderState
		wantAppendErr   error
		wantSentence    string
		wantSyntax      string
	}{
		{
			name:     "場合",
			appender: testutil.AppenderWithVerb(t, "食べる"),
			rule:     AttributingToOccasion,
			wantAppendState: lexical.AppenderState{
				Lexime:        types.Noun,
				ContentLexime: types.Noun,
			},
			wantSentence: "食べる場合",
			wantSyntax:   `group two verb "食べる";Rule(Attributing To Occasion){;conjugate as attributive;noun "場合";}`,
		},
		{
			name:     "悪い",
			appender: testutil.AppenderWithAdjective(t, "悪い"),
			rule:     AttributingToOccasion,
			wantAppendState: lexical.AppenderState{
				Lexime:        types.Noun,
				ContentLexime: types.Noun,
			},
			wantSentence: "悪い場合",
			wantSyntax:   `adjective "悪い";Rule(Attributing To Occasion){;conjugate as attributive;noun "場合";}`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := tc.appender.Append(tc.rule)
			if !errors.Is(gotErr, tc.wantAppendErr) {
				t.Errorf("append got unexpected error: got %v, want %v", gotErr, tc.wantAppendErr)
			}
			if gotErr == nil {
				if diff := cmp.Diff(tc.wantAppendState, tc.appender.GetState()); diff != "" {
					t.Errorf("Append got unexpected state: (-got,+want): %v", diff)
				}
			}
			got, err := tc.appender.WriteResult()
			if err != nil {
				t.Fatalf("WriteResult got unexpected error: %v", err)
			}
			if got.Sentence != tc.wantSentence {
				t.Errorf("WriteResult has unexpected sentence: got %v, want %v", got.Sentence, tc.wantSentence)
			}
			if got.Syntax != tc.wantSyntax {
				t.Errorf("WriteResult has unexpected syntax: got %v, want %v", got.Syntax, tc.wantSyntax)
			}
		})
	}
}
