// (C) 2022 Mike Spreitzer.  All rights reserved.

package simple

import (
	"strconv"
	"testing"

	"github.com/MikeSpreitzer/dynamatch"
	"github.com/MikeSpreitzer/dynamatch/boolean"
)

type TestVars struct{}

var _ boolean.Variables = TestVars{}

func (TestVars) GetVariable(vbl string) (string, bool) {
	if len(vbl) != 1 {
		return "", false
	}
	return strconv.FormatUint(uint64(vbl[0]), 16), true
}

type TestSub struct {
	cond boolean.OrAndsAtoms
}

var _ dynamatch.Subscription = &TestSub{}

func (sub *TestSub) GetCondition() boolean.OrAndsAtoms {
	return sub.cond
}

func TestMatcher(t *testing.T) {
	mr := New()
	sub1 := &TestSub{boolean.OrAndsAtoms{boolean.AndAtoms{boolean.Eq("A", "41")}}}
	mr.AddSubscription(sub1)
	expectMatches(t, mr, sub1)
	sub2 := &TestSub{boolean.OrAndsAtoms{boolean.AndAtoms{boolean.Eq("A", "42")}}}
	mr.AddSubscription(sub2)
	expectMatches(t, mr, sub1)
	sub3 := &TestSub{boolean.OrAndsAtoms{boolean.AndAtoms{boolean.Eq("B", "42")}}}
	mr.AddSubscription(sub3)
	expectMatches(t, mr, sub3, sub1)
	mr.AddSubscription(sub1)
	expectMatches(t, mr, sub1, sub3)
	mr.RemoveSubscription(sub1)
	expectMatches(t, mr, sub3)
}

func expectMatches(t *testing.T, mr dynamatch.Matcher, subslice ...dynamatch.Subscription) {
	actual := listMatches(mr)
	expected := subs(subslice...)
	if !subsEqual(actual, expected) {
		t.Errorf("Got %#+v instead of %#+v", actual, expected)
	}
}

func subs(subslice ...dynamatch.Subscription) Subscriptions {
	ans := Subscriptions{}
	for _, sub := range subslice {
		ans[sub] = sub.GetCondition()
	}
	return ans
}

func listMatches(mr dynamatch.Matcher) Subscriptions {
	subs := Subscriptions{}
	note := func(sub dynamatch.Subscription) error {
		subs[sub] = sub.GetCondition()
		return nil
	}
	mr.EnumerateMatchingSubscriptions(TestVars{}, note)
	return subs
}

func subsEqual(x, y Subscriptions) bool {
	if len(x) != len(y) {
		return false
	}
	for xi := range x {
		if _, ok := y[xi]; !ok {
			return false
		}
	}
	return true
}
