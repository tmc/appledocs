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
	// properties:
	MediaDestination() unsafe.Pointer
	SetMediaDestination(value unsafe.Pointer)
	MediaItems() INMediaItem
	SetMediaItems(value INMediaItem)
	MediaSearch() INMediaSearch
	SetMediaSearch(value INMediaSearch)
	// methods:
}

// A request to add a media item.
//
// Siri creates an object when the user asks to add a media item. The intents object contains the media to add. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object that contains the media to add.


// A request to add a media item.
//
// [Full Topic]
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



// The location for the media to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediadestination
func (i_ INAddMediaIntent) MediaDestination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaDestination"))
	return rv
}


// The location for the media to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediadestination
func (i_ INAddMediaIntent) SetMediaDestination(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaDestination:"), value)
}


// The media content to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediaitems
func (i_ INAddMediaIntent) MediaItems() INMediaItem {
	rv := objc.Send[INMediaItem](i_.ID, objc.Sel("mediaItems"))
	return rv
}


// The media content to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediaitems
func (i_ INAddMediaIntent) SetMediaItems(value INMediaItem) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaItems:"), value)
}


// The location to search for the media item to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediasearch
func (i_ INAddMediaIntent) MediaSearch() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("mediaSearch"))
	return rv
}


// The location to search for the media item to add.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inaddmediaintent/mediasearch
func (i_ INAddMediaIntent) SetMediaSearch(value INMediaSearch) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaSearch:"), value)
}



