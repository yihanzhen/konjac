package grammar

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/adjective"
	"github.com/yihanzhen/konjac/pkg/lexical/noun"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
	"github.com/yihanzhen/konjac/pkg/types"
)

var attributor = &ConditionalAppendable{
	cond: func(state *lexical.AppenderState) (lexical.Appendable, error) {
		if state.Lexime == types.Verb {
			return verb.NewVerbConjugation(types.Attributive), nil
		}
		if state.Lexime == types.Adjective {
			return adjective.NewConjugation(types.Attributive, types.AdjectiveConjuntiveModeUnset), nil
		}
		if state.Lexime == types.AdjNoun {
			return nil, fmt.Errorf("AdjNoun attributor: %w", errors.UnimplementedError)
		}
		if state.Lexime == types.Noun {
			return nil, fmt.Errorf("Noun attributor: %w", errors.UnimplementedError)
		}
		return nil, fmt.Errorf("attributor: state not attributable %v: %w", state, errors.IllegalStateError)
	},
}

var AttributingTo = &Rule{
	Name: "Attributing To",
	items: []lexical.Appendable{
		attributor,
	},
}

var AttributingToOccasion = &Rule{
	Name: "Attributing To Occasion",
	items: []lexical.Appendable{
		attributor,
		noun.Occasion,
	},
}
