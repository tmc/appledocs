// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchCallHistoryIntent] class.
var (
	INSearchCallHistoryIntentClass     _INSearchCallHistoryIntentClass
	INSearchCallHistoryIntentClassOnce sync.Once
)

func getINSearchCallHistoryIntentClass() _INSearchCallHistoryIntentClass {
	INSearchCallHistoryIntentClassOnce.Do(func() {
		INSearchCallHistoryIntentClass = _INSearchCallHistoryIntentClass{objc.GetClass("INSearchCallHistoryIntent")}
	})
	return INSearchCallHistoryIntentClass
}

type _INSearchCallHistoryIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchCallHistoryIntent] class.
type IINSearchCallHistoryIntent interface {
	IINIntent
}

// A request to list the calls matching the specified criteria.
//
// SiriKit creates objects when the user asks to see previous calls from their call history. This intent object contains the values for you to match when searching the user’s call history. Users can search for calls involving a specific person, calls that occurred on specific dates, or calls that are of a specific type such as missed calls. When performing the search, use only the parameters provided and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler confirms the request and creates an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchCallHistoryIntent
type INSearchCallHistoryIntent struct {
	INIntent
}

// INSearchCallHistoryIntentFrom constructs a [INSearchCallHistoryIntent] from an unsafe.Pointer.
//
// A request to list the calls matching the specified criteria.
func INSearchCallHistoryIntentFrom(ptr unsafe.Pointer) INSearchCallHistoryIntent {
	return INSearchCallHistoryIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchCallHistoryIntentClass) Alloc() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchCallHistoryIntentClass) New() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchCallHistoryIntent) Init() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchCallHistoryIntent) Autorelease() INSearchCallHistoryIntent {
	rv := objc.Send[INSearchCallHistoryIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchCallHistoryIntent creates a new INSearchCallHistoryIntent instance.
func NewINSearchCallHistoryIntent() INSearchCallHistoryIntent {
	return getINSearchCallHistoryIntentClass().New()
}




