// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForPhotosIntent] class.
var (
	INSearchForPhotosIntentClass     _INSearchForPhotosIntentClass
	INSearchForPhotosIntentClassOnce sync.Once
)

func getINSearchForPhotosIntentClass() _INSearchForPhotosIntentClass {
	INSearchForPhotosIntentClassOnce.Do(func() {
		INSearchForPhotosIntentClass = _INSearchForPhotosIntentClass{objc.GetClass("INSearchForPhotosIntent")}
	})
	return INSearchForPhotosIntentClass
}

type _INSearchForPhotosIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForPhotosIntent] class.
type IINSearchForPhotosIntent interface {
	IINIntent
}

// A request for the list of photos that match the specified criteria.
//
// The system creates an object when the user asks to search for photos in an app. The intent object contains the parameters to use during the search, including the possible name of a photo album, the people in the photos, or the location of the photos. Use this intent object to validate the search parameters and to begin the search process. When performing the search, use only the provided parameters and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results. For a list of other intents in the photos domain, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForPhotosIntent
type INSearchForPhotosIntent struct {
	INIntent
}

// INSearchForPhotosIntentFrom constructs a [INSearchForPhotosIntent] from an unsafe.Pointer.
//
// A request for the list of photos that match the specified criteria.
func INSearchForPhotosIntentFrom(ptr unsafe.Pointer) INSearchForPhotosIntent {
	return INSearchForPhotosIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForPhotosIntentClass) Alloc() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForPhotosIntentClass) New() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForPhotosIntent) Init() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForPhotosIntent) Autorelease() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForPhotosIntent creates a new INSearchForPhotosIntent instance.
func NewINSearchForPhotosIntent() INSearchForPhotosIntent {
	return getINSearchForPhotosIntentClass().New()
}




