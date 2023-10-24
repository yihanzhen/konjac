package types

import "fmt"

type conjugationType struct {
	Name string
}

// ConjugationType is the type that defines a conjugation.
type ConjugationType interface {
	enumType() conjugationType
	fmt.Stringer
}

// String implements fmt.Stringer.
func (c conjugationType) String() string {
	return c.Name
}

func (c conjugationType) enumType() conjugationType {
	return c
}

var (
	Irrealis    ConjugationType = conjugationType{Name: "irrealis"}
	Conjuntive  ConjugationType = conjugationType{Name: "conjunctive"}
	Attributive ConjugationType = conjugationType{Name: "attributive"}
	Terminal    ConjugationType = conjugationType{Name: "terminal"}
	Volitional  ConjugationType = conjugationType{Name: "volitional"}
	Imperative  ConjugationType = conjugationType{Name: "imperative"}
	Conditional ConjugationType = conjugationType{Name: "conditional"}
)

type verbConjugationRule struct {
	Name string
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
	return v.Name
}

var (
	ConjugationRuleUnset VerbConjugationRule = verbConjugationRule{Name: "conjugation rule unset"}
	GroupOne             VerbConjugationRule = verbConjugationRule{Name: "group one"}
	GroupTwo             VerbConjugationRule = verbConjugationRule{Name: "group two"}
	KaIrregular          VerbConjugationRule = verbConjugationRule{Name: "ka irregular"}
	SaIrregular          VerbConjugationRule = verbConjugationRule{Name: "sa irregular"}
)

type AdjectiveConjuntiveMode interface {
	enumType() AdjectiveConjuntiveMode
	fmt.Stringer
}

type adjectiveConjuntiveMode struct {
	Name string
}

func (a adjectiveConjuntiveMode) enumType() AdjectiveConjuntiveMode {
	return a
}

func (a adjectiveConjuntiveMode) String() string {
	return a.Name
}

var (
	AdjectiveConjuntiveModeUnset AdjectiveConjuntiveMode = adjectiveConjuntiveMode{Name: "adjectival conjunctive unset"}
	Adjectival                   AdjectiveConjuntiveMode = adjectiveConjuntiveMode{Name: "adjectival conjunctive mode"}
	Adverbial                    AdjectiveConjuntiveMode = adjectiveConjuntiveMode{Name: "adverbial conjunctive mode"}
	Completed                    AdjectiveConjuntiveMode = adjectiveConjuntiveMode{Name: "completed conjunctive mode"}
)
