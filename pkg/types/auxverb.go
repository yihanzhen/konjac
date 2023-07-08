package types

import "fmt"

type auxVerbType struct {
	name string
}

type AuxVerbType interface {
	enumType() auxVerbType
	fmt.Stringer
}

func (a auxVerbType) enumType() auxVerbType {
	return a
}

func (a auxVerbType) String() string {
	return a.name
}

var (
	// Assertor is だ.
	Assertor AuxVerbType = auxVerbType{name: "assertor"}
	// PoliteAssertor is です.
	PoliteAssertor AuxVerbType = auxVerbType{name: "polite assertor"}
	// PoliteMaker is です, but used to make an adjective polite.
	PoliteMaker AuxVerbType = auxVerbType{name: "polite maker"}
	// Negator is ない.
	Negator AuxVerbType = auxVerbType{name: "negator"}
	// PassiveMaker is られる.
	PassiveMaker AuxVerbType = auxVerbType{name: "passive maker"}
	// CausativeMaker is させる.
	CausativeMaker AuxVerbType = auxVerbType{name: "causative maker"}
	// PastMaker is た.
	PastMaker AuxVerbType = auxVerbType{name: "past maker"}
	// Pauser is で or て.
	Pauser AuxVerbType = auxVerbType{name: "pauser"}
	// StateMaker is で or て.
	StateMaker AuxVerbType = auxVerbType{name: "state maker"}
	// ConditionalMaker is　ば.
	ConditionalMaker AuxVerbType = auxVerbType{name: "conditional maker"}
	// PastConditionalMaker is たら.
	PastConditionalMaker AuxVerbType = auxVerbType{name: "past conditional maker"}
	//　VolitionalMaker is よう.
	VolitionalMaker AuxVerbType = auxVerbType{name: "volitional maker"}
	// PotentionalMaker is られる.
	PotentialMaker AuxVerbType = auxVerbType{name: "potential maker"}
	// WishMaker is たい.
	WishMaker AuxVerbType = auxVerbType{name: "wish maker"}
)
