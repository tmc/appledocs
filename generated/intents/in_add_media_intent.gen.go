// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INAddMediaIntent] class.
var (
	INAddMediaIntentClass     _INAddMediaIntentClass
	INAddMediaIntentClassOnce sync.Once
)

func getINAddMediaIntentClass() _INAddMediaIntentClass {
	INAddMediaIntentClassOnce.Do(func() {
		INAddMediaIntentClass = _INAddMediaIntentClass{objc.GetClass("INAddMediaIntent")}
	})
	return INAddMediaIntentClass
}

type _INAddMediaIntentClass struct {
	class objc.Class
}

// An interface definition for the [INAddMediaIntent] class.
type IINAddMediaIntent interface {
	IINIntent
}

// A request to add a media item.
//
// Siri creates an object when the user asks to add a media item. The intents object contains the media to add. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that contains the media to add.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INAddMediaIntent
type INAddMediaIntent struct {
	INIntent
}

// INAddMediaIntentFrom constructs a [INAddMediaIntent] from an unsafe.Pointer.
//
// A request to add a media item.
func INAddMediaIntentFrom(ptr unsafe.Pointer) INAddMediaIntent {
	return INAddMediaIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INAddMediaIntentClass) Alloc() INAddMediaIntent {
	rv := objc.Send[INAddMediaIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INAddMediaIntentClass) New() INAddMediaIntent {
	rv := objc.Send[INAddMediaIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INAddMediaIntent) Init() INAddMediaIntent {
	rv := objc.Send[INAddMediaIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INAddMediaIntent) Autorelease() INAddMediaIntent {
	rv := objc.Send[INAddMediaIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINAddMediaIntent creates a new INAddMediaIntent instance.
func NewINAddMediaIntent() INAddMediaIntent {
	return getINAddMediaIntentClass().New()
}




