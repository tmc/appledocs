// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INActivateCarSignalIntent] class.
var (
	INActivateCarSignalIntentClass     _INActivateCarSignalIntentClass
	INActivateCarSignalIntentClassOnce sync.Once
)

func getINActivateCarSignalIntentClass() _INActivateCarSignalIntentClass {
	INActivateCarSignalIntentClassOnce.Do(func() {
		INActivateCarSignalIntentClass = _INActivateCarSignalIntentClass{objc.GetClass("INActivateCarSignalIntent")}
	})
	return INActivateCarSignalIntentClass
}

type _INActivateCarSignalIntentClass struct {
	class objc.Class
}

// An interface definition for the [INActivateCarSignalIntent] class.
type IINActivateCarSignalIntent interface {
	IINIntent
}

// A request to activate the signals on the user’s car.
//
// When the user asks for an audible or visual signal from the car, Siri creates an object. This intent object can contain the name of the user’s car and the signal options. Use this object to trigger audible and visual signals from the car. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INActivateCarSignalIntent
type INActivateCarSignalIntent struct {
	INIntent
}

// INActivateCarSignalIntentFrom constructs a [INActivateCarSignalIntent] from an unsafe.Pointer.
//
// A request to activate the signals on the user’s car.
func INActivateCarSignalIntentFrom(ptr unsafe.Pointer) INActivateCarSignalIntent {
	return INActivateCarSignalIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INActivateCarSignalIntentClass) Alloc() INActivateCarSignalIntent {
	rv := objc.Send[INActivateCarSignalIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INActivateCarSignalIntentClass) New() INActivateCarSignalIntent {
	rv := objc.Send[INActivateCarSignalIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INActivateCarSignalIntent) Init() INActivateCarSignalIntent {
	rv := objc.Send[INActivateCarSignalIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INActivateCarSignalIntent) Autorelease() INActivateCarSignalIntent {
	rv := objc.Send[INActivateCarSignalIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINActivateCarSignalIntent creates a new INActivateCarSignalIntent instance.
func NewINActivateCarSignalIntent() INActivateCarSignalIntent {
	return getINActivateCarSignalIntentClass().New()
}




