// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INHangUpCallIntent] class.
var (
	INHangUpCallIntentClass     _INHangUpCallIntentClass
	INHangUpCallIntentClassOnce sync.Once
)

func getINHangUpCallIntentClass() _INHangUpCallIntentClass {
	INHangUpCallIntentClassOnce.Do(func() {
		INHangUpCallIntentClass = _INHangUpCallIntentClass{objc.GetClass("INHangUpCallIntent")}
	})
	return INHangUpCallIntentClass
}

type _INHangUpCallIntentClass struct {
	class objc.Class
}

// An interface definition for the [INHangUpCallIntent] class.
type IINHangUpCallIntent interface {
	IINIntent
}

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INHangUpCallIntent
type INHangUpCallIntent struct {
	INIntent
}

// INHangUpCallIntentFrom constructs a [INHangUpCallIntent] from an unsafe.Pointer.
func INHangUpCallIntentFrom(ptr unsafe.Pointer) INHangUpCallIntent {
	return INHangUpCallIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INHangUpCallIntentClass) Alloc() INHangUpCallIntent {
	rv := objc.Send[INHangUpCallIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INHangUpCallIntentClass) New() INHangUpCallIntent {
	rv := objc.Send[INHangUpCallIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INHangUpCallIntent) Init() INHangUpCallIntent {
	rv := objc.Send[INHangUpCallIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INHangUpCallIntent) Autorelease() INHangUpCallIntent {
	rv := objc.Send[INHangUpCallIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINHangUpCallIntent creates a new INHangUpCallIntent instance.
func NewINHangUpCallIntent() INHangUpCallIntent {
	return getINHangUpCallIntentClass().New()
}




