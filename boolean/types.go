// (C) 2022 Mike Spreitzer.  All rights reserved.

package boolean

// OrAndsAtoms is a boolean expression in Disjunctive Normal Form
// whose atoms are string equality comparisons between
// a variable and a literal.
type OrAndsAtoms []AndAtoms

// AndAtoms is the AND of some string equality comparisons
// between variable and literal.
type AndAtoms []Atom

// Atom is one string equality comparison between a variable and a literal
type Atom struct {
	// Variable identifies the variable being compared.
	Variable string

	// CompareTo is the literal being compared with.
	CompareTo string

	// Negate indicates whether the result of the comparison is to be negated
	Negate bool
}

type Variables interface {
	GetVariable(string) (string, bool)
}
