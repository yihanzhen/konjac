package types

import "fmt"

type auxVerbType struct {
	Name string
}

type AuxVerbType interface {
	enumType() auxVerbType
	fmt.Stringer
}

func (a auxVerbType) enumType() auxVerbType {
	return a
}

func (a auxVerbType) String() string {
	return a.Name
}

var (
	// Assertor is だ.
	Assertor AuxVerbType = auxVerbType{Name: "assertor"}
	// PoliteAssertor is です.
	PoliteAssertor AuxVerbType = auxVerbType{Name: "polite assertor"}
	// AdjPoliteMaker is です, but used to make an adjective polite.
	AdjPoliteMaker AuxVerbType = auxVerbType{Name: "polite maker"}
	// AdjPoliteMaker is です, but used to make an adjective polite.
	VerbPoliteMaker AuxVerbType = auxVerbType{Name: "polite maker"}
	// Negator is ない.
	Negator AuxVerbType = auxVerbType{Name: "negator"}
	// PassiveMaker is られる.
	PassiveMaker AuxVerbType = auxVerbType{Name: "passive maker"}
	// CausativeMaker is させる.
	CausativeMaker AuxVerbType = auxVerbType{Name: "causative maker"}
	// PastMaker is た.
	PastMaker AuxVerbType = auxVerbType{Name: "past maker"}
	// Pauser is で or て.
	Pauser AuxVerbType = auxVerbType{Name: "pauser"}
	// StateMaker is で or て.
	StateMaker AuxVerbType = auxVerbType{Name: "state maker"}
	// ConditionalMaker is　ば.
	ConditionalMaker AuxVerbType = auxVerbType{Name: "conditional maker"}
	// PastConditionalMaker is たら.
	PastConditionalMaker AuxVerbType = auxVerbType{Name: "past conditional maker"}
	//　VolitionalMaker is よう.
	VolitionalMaker AuxVerbType = auxVerbType{Name: "volitional maker"}
	// PotentionalMaker is られる.
	PotentialMaker AuxVerbType = auxVerbType{Name: "potential maker"}
	// WishMaker is たい.
	WishMaker AuxVerbType = auxVerbType{Name: "wish maker"}
)
