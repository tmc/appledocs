// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartPhotoPlaybackIntent] class.
var (
	INStartPhotoPlaybackIntentClass     _INStartPhotoPlaybackIntentClass
	INStartPhotoPlaybackIntentClassOnce sync.Once
)

func getINStartPhotoPlaybackIntentClass() _INStartPhotoPlaybackIntentClass {
	INStartPhotoPlaybackIntentClassOnce.Do(func() {
		INStartPhotoPlaybackIntentClass = _INStartPhotoPlaybackIntentClass{objc.GetClass("INStartPhotoPlaybackIntent")}
	})
	return INStartPhotoPlaybackIntentClass
}

type _INStartPhotoPlaybackIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartPhotoPlaybackIntent] class.
type IINStartPhotoPlaybackIntent interface {
	IINIntent
}

// A request to search for photos and initiate a slideshow with the results.
//
// The system creates an object when the user asks to start a slideshow of a set of photos. This intent object contains the parameters to use when searching for the photos, including the possible name of a photo album, the people in the photos, or the location of the photos. Use this intent object to perform the search and initiate the slideshow in your app. When performing the search, use only the parameters provided and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. After a successful search, Siri launches your app so that it can begin the slideshow. For a list of other intents in the photos domain, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartPhotoPlaybackIntent
type INStartPhotoPlaybackIntent struct {
	INIntent
}

// INStartPhotoPlaybackIntentFrom constructs a [INStartPhotoPlaybackIntent] from an unsafe.Pointer.
//
// A request to search for photos and initiate a slideshow with the results.
func INStartPhotoPlaybackIntentFrom(ptr unsafe.Pointer) INStartPhotoPlaybackIntent {
	return INStartPhotoPlaybackIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartPhotoPlaybackIntentClass) Alloc() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartPhotoPlaybackIntentClass) New() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartPhotoPlaybackIntent) Init() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartPhotoPlaybackIntent) Autorelease() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartPhotoPlaybackIntent creates a new INStartPhotoPlaybackIntent instance.
func NewINStartPhotoPlaybackIntent() INStartPhotoPlaybackIntent {
	return getINStartPhotoPlaybackIntentClass().New()
}




