// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INListCarsIntent] class.
var (
	INListCarsIntentClass     _INListCarsIntentClass
	INListCarsIntentClassOnce sync.Once
)

func getINListCarsIntentClass() _INListCarsIntentClass {
	INListCarsIntentClassOnce.Do(func() {
		INListCarsIntentClass = _INListCarsIntentClass{objc.GetClass("INListCarsIntent")}
	})
	return INListCarsIntentClass
}

type _INListCarsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INListCarsIntent] class.
type IINListCarsIntent interface {
	IINIntent
}

// An intent for retrieving a list of the user’s electric vehicles.
//
// Maps creates instances of when it needs to display a list of the user’s electric vehicles, usually before or during route planning. To handle this intent, you create an object that conforms to the protocol. The object handles, and optionally confirms, the request by providing an instance of . The response contains an array of objects, one for each of the user’s electric vehicles. Maps requires that each vehicle responds to requests so that it can use the information the intent provides—current charge, battery capacity, distance remaining, and so forth—to more accurately plan the route.


// An intent for retrieving a list of the user’s electric vehicles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INListCarsIntent
type INListCarsIntent struct {
	INIntent
}

// INListCarsIntentFrom constructs a [INListCarsIntent] from an unsafe.Pointer.
//
// An intent for retrieving a list of the user’s electric vehicles.
func INListCarsIntentFrom(ptr unsafe.Pointer) INListCarsIntent {
	return INListCarsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INListCarsIntentClass) Alloc() INListCarsIntent {
	rv := objc.Send[INListCarsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INListCarsIntentClass) New() INListCarsIntent {
	rv := objc.Send[INListCarsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INListCarsIntent) Init() INListCarsIntent {
	rv := objc.Send[INListCarsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INListCarsIntent) Autorelease() INListCarsIntent {
	rv := objc.Send[INListCarsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINListCarsIntent creates a new INListCarsIntent instance.
func NewINListCarsIntent() INListCarsIntent {
	return getINListCarsIntentClass().New()
}




