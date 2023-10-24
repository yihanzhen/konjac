package grammar

import (
	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/adjective"
	"github.com/yihanzhen/konjac/pkg/lexical/auxverb"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
	"github.com/yihanzhen/konjac/pkg/types"
)

var MakePolite = &Rule{
	Name: "Make Polite",
	items: []lexical.Appendable{
		verb.NewVerbConjugation(types.Conjuntive),
		auxverb.VerbPoliteMaker,
	},
}

var AssertEasy = &Rule{
	Name: "Assert Easy",
	items: []lexical.Appendable{
		verb.NewVerbConjugation(types.Conjuntive),
		adjective.EasyAssertor,
	},
}

var AssertHard = &Rule{
	Name: "Assert Hard",
	items: []lexical.Appendable{
		verb.NewVerbConjugation(types.Conjuntive),
		adjective.HardAssertor,
	},
}
