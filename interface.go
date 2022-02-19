// (C) 2022 Mike Spreitzer.  All rights reserved.

package dynamatch

import (
	"github.com/MikeSpreitzer/dynamatch/boolean"
)

// Matcher is about matching items against subscriptions.
// An item has a string-to-string mapping defining constant
// values for some variables.
// A subscription defines matching items by a boolean
// combination of variable-to-literal equality comparisons.
// A Matcher is safe for concurrent access.
// All the calls to Add and Remove Subscriptions appear
// to happen atomically in some total order.
// When the set of subscriptions is changed concurrently
// with an enumeration of matches, any of the subscriptions
// being added or removed may or may not be enumerated.
type Matcher interface {
	// AddSubscription adds a Subscription if it is not already present
	AddSubscription(Subscription)

	// RemoveSubscription ensures a Subscription is not included
	RemoveSubscription(Subscription)

	// EnumerateMatchingSubscriptions calls the given function on
	// each matching Subscription, stopping as soon as that function
	// returns an error and returning the same error.
	// The return is nil if the given function never returns an error.
	EnumerateMatchingSubscriptions(boolean.Variables, func(Subscription) error) error
}

// Subscription is something that matches items
// based on string comparisons.
type Subscription interface {
	// GetCondition returns a disjunctive normal form representation of
	// the whether a given item matches the Subscription.
	GetCondition() boolean.OrAndsAtoms
}
