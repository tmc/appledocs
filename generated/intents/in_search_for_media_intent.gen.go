// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForMediaIntent] class.
var (
	INSearchForMediaIntentClass     _INSearchForMediaIntentClass
	INSearchForMediaIntentClassOnce sync.Once
)

func getINSearchForMediaIntentClass() _INSearchForMediaIntentClass {
	INSearchForMediaIntentClassOnce.Do(func() {
		INSearchForMediaIntentClass = _INSearchForMediaIntentClass{objc.GetClass("INSearchForMediaIntent")}
	})
	return INSearchForMediaIntentClass
}

type _INSearchForMediaIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForMediaIntent] class.
type IINSearchForMediaIntent interface {
	IINIntent
}

// A request to search for a media item.
//
// The system creates an object when the user asks Siri to search for a media item, or searches for a musician or band in Spotlight. To handle this intent, provide a handler that conforms to the protocol. Your handler should confirm the request and create an object with the media to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMediaIntent
type INSearchForMediaIntent struct {
	INIntent
}

// INSearchForMediaIntentFrom constructs a [INSearchForMediaIntent] from an unsafe.Pointer.
//
// A request to search for a media item.
func INSearchForMediaIntentFrom(ptr unsafe.Pointer) INSearchForMediaIntent {
	return INSearchForMediaIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForMediaIntentClass) Alloc() INSearchForMediaIntent {
	rv := objc.Send[INSearchForMediaIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForMediaIntentClass) New() INSearchForMediaIntent {
	rv := objc.Send[INSearchForMediaIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForMediaIntent) Init() INSearchForMediaIntent {
	rv := objc.Send[INSearchForMediaIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForMediaIntent) Autorelease() INSearchForMediaIntent {
	rv := objc.Send[INSearchForMediaIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForMediaIntent creates a new INSearchForMediaIntent instance.
func NewINSearchForMediaIntent() INSearchForMediaIntent {
	return getINSearchForMediaIntentClass().New()
}




