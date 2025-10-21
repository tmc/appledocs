// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSetCarLockStatusIntent] class.
var (
	INSetCarLockStatusIntentClass     _INSetCarLockStatusIntentClass
	INSetCarLockStatusIntentClassOnce sync.Once
)

func getINSetCarLockStatusIntentClass() _INSetCarLockStatusIntentClass {
	INSetCarLockStatusIntentClassOnce.Do(func() {
		INSetCarLockStatusIntentClass = _INSetCarLockStatusIntentClass{objc.GetClass("INSetCarLockStatusIntent")}
	})
	return INSetCarLockStatusIntentClass
}

type _INSetCarLockStatusIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetCarLockStatusIntent] class.
type IINSetCarLockStatusIntent interface {
	IINIntent
}

// A request to lock or unlock the user’s car.
//
// When the user asks to lock or unlock the car, Siri creates an object. This intent object can contain the name of the user’s car and the requested lock status. Use this object to lock or unlock the car. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetCarLockStatusIntent
type INSetCarLockStatusIntent struct {
	INIntent
}

// INSetCarLockStatusIntentFrom constructs a [INSetCarLockStatusIntent] from an unsafe.Pointer.
//
// A request to lock or unlock the user’s car.
func INSetCarLockStatusIntentFrom(ptr unsafe.Pointer) INSetCarLockStatusIntent {
	return INSetCarLockStatusIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetCarLockStatusIntentClass) Alloc() INSetCarLockStatusIntent {
	rv := objc.Send[INSetCarLockStatusIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetCarLockStatusIntentClass) New() INSetCarLockStatusIntent {
	rv := objc.Send[INSetCarLockStatusIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetCarLockStatusIntent) Init() INSetCarLockStatusIntent {
	rv := objc.Send[INSetCarLockStatusIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetCarLockStatusIntent) Autorelease() INSetCarLockStatusIntent {
	rv := objc.Send[INSetCarLockStatusIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetCarLockStatusIntent creates a new INSetCarLockStatusIntent instance.
func NewINSetCarLockStatusIntent() INSetCarLockStatusIntent {
	return getINSetCarLockStatusIntentClass().New()
}




