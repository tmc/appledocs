// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INUpdateMediaAffinityIntent] class.
var (
	INUpdateMediaAffinityIntentClass     _INUpdateMediaAffinityIntentClass
	INUpdateMediaAffinityIntentClassOnce sync.Once
)

func getINUpdateMediaAffinityIntentClass() _INUpdateMediaAffinityIntentClass {
	INUpdateMediaAffinityIntentClassOnce.Do(func() {
		INUpdateMediaAffinityIntentClass = _INUpdateMediaAffinityIntentClass{objc.GetClass("INUpdateMediaAffinityIntent")}
	})
	return INUpdateMediaAffinityIntentClass
}

type _INUpdateMediaAffinityIntentClass struct {
	class objc.Class
}

// An interface definition for the [INUpdateMediaAffinityIntent] class.
type IINUpdateMediaAffinityIntent interface {
	IINIntent
	MediaItems() []INMediaItem
	MediaSearch() INMediaSearch
	AffinityType() unsafe.Pointer
	SetAffinityType(value unsafe.Pointer)
}

// A request to update the user’s affinity for a media item.
//
// Siri creates an object when the user expresses a preference for or disinterest in a media item. The intent object contains the media to update. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the media to update.


// A request to update the user’s affinity for a media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpdateMediaAffinityIntent
type INUpdateMediaAffinityIntent struct {
	INIntent
}

// INUpdateMediaAffinityIntentFrom constructs a [INUpdateMediaAffinityIntent] from an unsafe.Pointer.
//
// A request to update the user’s affinity for a media item.
func INUpdateMediaAffinityIntentFrom(ptr unsafe.Pointer) INUpdateMediaAffinityIntent {
	return INUpdateMediaAffinityIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INUpdateMediaAffinityIntentClass) Alloc() INUpdateMediaAffinityIntent {
	rv := objc.Send[INUpdateMediaAffinityIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUpdateMediaAffinityIntentClass) New() INUpdateMediaAffinityIntent {
	rv := objc.Send[INUpdateMediaAffinityIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUpdateMediaAffinityIntent) Init() INUpdateMediaAffinityIntent {
	rv := objc.Send[INUpdateMediaAffinityIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUpdateMediaAffinityIntent) Autorelease() INUpdateMediaAffinityIntent {
	rv := objc.Send[INUpdateMediaAffinityIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUpdateMediaAffinityIntent creates a new INUpdateMediaAffinityIntent instance.
func NewINUpdateMediaAffinityIntent() INUpdateMediaAffinityIntent {
	return getINUpdateMediaAffinityIntentClass().New()
}



// The media items to update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpdateMediaAffinityIntent/mediaItems
func (i_ INUpdateMediaAffinityIntent) MediaItems() []INMediaItem {
	rv := objc.Send[[]INMediaItem](i_.ID, objc.Sel("mediaItems"))
	return rv
}


// The type of item to search for.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpdateMediaAffinityIntent/mediaSearch
func (i_ INUpdateMediaAffinityIntent) MediaSearch() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("mediaSearch"))
	return rv
}


// The user’s affinity for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inupdatemediaaffinityintent/affinitytype
func (i_ INUpdateMediaAffinityIntent) AffinityType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("affinityType"))
	return rv
}


// The user’s affinity for the media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inupdatemediaaffinityintent/affinitytype
func (i_ INUpdateMediaAffinityIntent) SetAffinityType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAffinityType:"), value)
}



