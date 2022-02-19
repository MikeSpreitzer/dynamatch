// (C) 2022 Mike Spreitzer.  All rights reserved.

package simple

import (
	"sync"
	"sync/atomic"

	"github.com/MikeSpreitzer/dynamatch"
	"github.com/MikeSpreitzer/dynamatch/boolean"
)

type Implementation struct {
	current atomic.Value
	lock    sync.Mutex
}

type Subscriptions map[dynamatch.Subscription]boolean.OrAndsAtoms

func New() dynamatch.Matcher {
	subs := make(Subscriptions)
	impl := &Implementation{}
	impl.current.Store(subs)
	return impl
}

func (impl *Implementation) AddSubscription(sub dynamatch.Subscription) {
	impl.lock.Lock()
	defer impl.lock.Unlock()
	oldSubs := impl.current.Load().(Subscriptions)
	newSubs := copySubs(oldSubs)
	newSubs[sub] = sub.GetCondition()
	impl.current.Store(newSubs)
}

func (impl *Implementation) RemoveSubscription(sub dynamatch.Subscription) {
	impl.lock.Lock()
	defer impl.lock.Unlock()
	oldSubs := impl.current.Load().(Subscriptions)
	newSubs := copySubs(oldSubs)
	delete(newSubs, sub)
	impl.current.Store(newSubs)
}

func copySubs(oldSubs Subscriptions) Subscriptions {
	newSubs := make(Subscriptions, len(oldSubs))
	for key, val := range oldSubs {
		newSubs[key] = val
	}
	return newSubs
}

func (impl *Implementation) EnumerateMatchingSubscriptions(vars boolean.Variables, consume func(dynamatch.Subscription) error) error {
	subs := impl.current.Load().(Subscriptions)
	return subs.EnumMatches(vars, consume)
}

func (subs Subscriptions) EnumMatches(vars boolean.Variables, consume func(dynamatch.Subscription) error) error {
	for sub, dnf := range subs {
		if dnf.Evaluate(vars) {
			if err := consume(sub); err != nil {
				return err
			}
		}
	}
	return nil
}
