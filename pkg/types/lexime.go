package types

import "fmt"

type leximeType struct {
	name string
}

// LeximeType is the type that defines a lexime.
type LeximeType interface {
	enumType() leximeType
	IsContent() bool
	fmt.Stringer
}

// String implements fmt.Stringer.
func (l leximeType) String() string {
	return l.name
}

func (l leximeType) enumType() leximeType {
	return l
}

// Whether the leximeType is a content lexime. Every lexime type
// except particle is a content lexime.
func (l leximeType) IsContent() bool {
	return !(l == Particle)
}

var (
	Verb      LeximeType = leximeType{name: "verb"}
	Noun      LeximeType = leximeType{name: "noun"}
	Adverb    leximeType = leximeType{name: "adverb"}
	AdjNoun   LeximeType = leximeType{name: "adjectival noun"}
	Adjective LeximeType = leximeType{name: "adjective"}
	AuxVerb   LeximeType = leximeType{name: "auxiliary verb"}
	Particle  LeximeType = leximeType{name: "particle"}
)
