// (C) 2022 Mike Spreitzer.  All rights reserved.

package boolean

import (
	"strconv"
	"testing"
)

type TestVars struct{}

var _ Variables = TestVars{}

func (TestVars) GetVariable(vbl string) (string, bool) {
	if len(vbl) != 1 {
		return "", false
	}
	return strconv.FormatUint(uint64(vbl[0]), 16), true
}

func TestDNF(t *testing.T) {
	if !Eq("A", "41").Evaluate(TestVars{}) {
		t.Error("Eq(A, 41) evald false")
	}
	if Eq("A", "42").Evaluate(TestVars{}) {
		t.Error("Eq(A, 42) evald true")
	}
	if !NEq("B", "41").Evaluate(TestVars{}) {
		t.Error("NEq(B, 41)) evald false")
	}
	if NEq("B", "42").Evaluate(TestVars{}) {
		t.Error("NEq(B, 42)) evald true")
	}
	if !(AndAtoms{}).Evaluate(TestVars{}) {
		t.Error("Empty conjunction evald false")
	}
	if !(AndAtoms{Eq("A", "41")}).Evaluate(TestVars{}) {
		t.Error("And(Eq(A,41)) evald false")
	}
	if (AndAtoms{Eq("A", "42")}).Evaluate(TestVars{}) {
		t.Error("And(Eq(A,42)) evald true")
	}
	if (AndAtoms{Eq("A", "41"), NEq("A", "41")}).Evaluate(TestVars{}) {
		t.Error("And(Eq(A,41),NEq(A,41)) evald true")
	}
	if !(AndAtoms{Eq("A", "41"), NEq("B", "41")}).Evaluate(TestVars{}) {
		t.Error("And(Eq(A,41),NEq(B,41)) evald false")
	}
	if (OrAndsAtoms{}).Evaluate(TestVars{}) {
		t.Error("Empty OrAndsAtoms evald true")
	}
	if (OrAndsAtoms{AndAtoms{Eq("A", "42")}}).Evaluate(TestVars{}) {
		t.Error("Or(And(Eq(A,42))) evald true")
	}
	if !(OrAndsAtoms{AndAtoms{Eq("A", "42")}, AndAtoms{Eq("A", "41"), NEq("B", "41")}}).Evaluate(TestVars{}) {
		t.Error("Or(And(Eq(A,42)),And(Eq(A,41),NEq(B,41))) evald false")
	}
}
