package types

import "fmt"

type leximeType struct {
	Name string
}

// LeximeType is the type that defines a lexime.
type LeximeType interface {
	enumType() leximeType
	IsContent() bool
	fmt.Stringer
}

// String implements fmt.Stringer.
func (l leximeType) String() string {
	return l.Name
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
	Verb      LeximeType = leximeType{Name: "verb"}
	Noun      LeximeType = leximeType{Name: "noun"}
	Adverb    leximeType = leximeType{Name: "adverb"}
	AdjNoun   LeximeType = leximeType{Name: "adjectival noun"}
	Adjective LeximeType = leximeType{Name: "adjective"}
	AuxVerb   LeximeType = leximeType{Name: "auxiliary verb"}
	Particle  LeximeType = leximeType{Name: "particle"}
)
