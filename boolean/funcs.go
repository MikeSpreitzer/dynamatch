// (C) 2022 Mike Spreitzer.  All rights reserved.

package boolean

// Evaluate is a convienience function that evaluates the given arguments.
// Their semantics is defined at their types, not by this function.
func (dnf OrAndsAtoms) Evaluate(vars Variables) bool {
	for _, conj := range dnf {
		if conj.Evaluate(vars) {
			return true
		}
	}
	return false
}

// Evaluate is a convienience function that evaluates the given arguments.
// Their semantics is defined at their types, not by this function.
func (conj AndAtoms) Evaluate(vars Variables) bool {
	for _, atom := range conj {
		if !atom.Evaluate(vars) {
			return false
		}
	}
	return true
}

// Evaluate is a convienience function that evaluates the given arguments.
// Their semantics is defined at their types, not by this function.
func (atom Atom) Evaluate(vars Variables) bool {
	varVal, ok := vars.GetVariable(atom.Variable)
	return ok && (atom.Negate == (varVal != atom.CompareTo))
}
