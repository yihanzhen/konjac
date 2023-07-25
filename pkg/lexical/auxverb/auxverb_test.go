package auxverb

import (
	errorsutil "errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/types"
)

func TestAppend(t *testing.T) {
	cases := []struct {
		name               string
		auxVerb            *AuxVerb
		state              *lexical.AppenderState
		wantAppendMutation *lexical.AppenderMutation
		wantErr            error
	}{
		{
			name:    "ます success",
			auxVerb: VerbPoliteMaker,
			state: &lexical.AppenderState{
				Lexime:      types.Verb,
				Conjugation: types.Conjuntive,
			},
			wantAppendMutation: &lexical.AppenderMutation{
				Append:    VerbPoliteMaker,
				SetLexime: types.AuxVerb,
				SetAuxVerbState: &lexical.AuxVerbState{
					AuxVerbType: types.VerbPoliteMaker,
				},
			},
		},
		{
			name:    "ます not verb error",
			auxVerb: VerbPoliteMaker,
			state: &lexical.AppenderState{
				Lexime:      types.Adjective,
				Conjugation: types.Conjuntive,
			},
			wantErr: errors.IllegalAppenderLeximeError,
		},
		{
			name:    "ます not conjunctive error",
			auxVerb: VerbPoliteMaker,
			state: &lexical.AppenderState{
				Lexime:      types.Verb,
				Conjugation: types.Irrealis,
			},
			wantErr: errors.IllegalAppenderConjugationError,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotMut, err := tc.auxVerb.Append(tc.state)
			if !errorsutil.Is(err, tc.wantErr) {
				t.Errorf("got error %v, want error %v", err, tc.wantErr)
				return
			}
			if err == nil {
				if diff := cmp.Diff(tc.wantAppendMutation, gotMut, cmpopts.IgnoreUnexported(AuxVerb{})); diff != "" {
					t.Errorf("unexpected diff in AppenderMutation: (-want,+got):\n%s", diff)
				}
			}
		})
	}
}
