package grammar

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/testutil"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestVerbRules(t *testing.T) {
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
			name:     "やすい",
			appender: testutil.AppenderWithVerb(t, "食べる"),
			rule:     AssertEasy,
			wantAppendState: lexical.AppenderState{
				Lexime:        types.Adjective,
				ContentLexime: types.Adjective,
			},
			wantSentence: "食べやすい",
			wantSyntax:   `group two verb "食べる";Rule(Assert Easy){;conjugate as conjunctive;ajdective "やすい";}`,
		},
		{
			name:     "にくい",
			appender: testutil.AppenderWithVerb(t, "食べる"),
			rule:     AssertHard,
			wantAppendState: lexical.AppenderState{
				Lexime:        types.Adjective,
				ContentLexime: types.Adjective,
			},
			wantSentence: "食べにくい",
			wantSyntax:   `group two verb "食べる";Rule(Assert Hard){;conjugate as conjunctive;ajdective "にくい";}`,
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
