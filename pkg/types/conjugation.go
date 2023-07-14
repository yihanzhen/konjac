package types

import "fmt"

type conjugationType struct {
	name string
}

// ConjugationType is the type that defines a conjugation.
type ConjugationType interface {
	enumType() conjugationType
	fmt.Stringer
}

// String implements fmt.Stringer.
func (c conjugationType) String() string {
	return c.name
}

func (c conjugationType) enumType() conjugationType {
	return c
}

var (
	Irrealis    ConjugationType = conjugationType{name: "irrealis"}
	Conjuntive  ConjugationType = conjugationType{name: "conjunctive"}
	Attributive ConjugationType = conjugationType{name: "attributive"}
	Terminal    ConjugationType = conjugationType{name: "terminal"}
	Volitional  ConjugationType = conjugationType{name: "volitional"}
	Imperative  ConjugationType = conjugationType{name: "imperative"}
	Conditional ConjugationType = conjugationType{name: "conditional"}
)

type verbConjugationRule struct {
	name string
}

// VerbConjugationRule is the conjugation rule of a verb.
type VerbConjugationRule interface {
	enumType() verbConjugationRule
	fmt.Stringer
}

func (v verbConjugationRule) enumType() verbConjugationRule {
	return v
}

func (v verbConjugationRule) String() string {
	return v.name
}

var (
	ConjugationRuleUnset VerbConjugationRule = verbConjugationRule{name: "conjugation rule unset"}
	GroupOne             VerbConjugationRule = verbConjugationRule{name: "group one"}
	GroupTwo             VerbConjugationRule = verbConjugationRule{name: "group two"}
	KaIrregular          VerbConjugationRule = verbConjugationRule{name: "ka irregular"}
	SaIrregular          VerbConjugationRule = verbConjugationRule{name: "sa irregular"}
)
