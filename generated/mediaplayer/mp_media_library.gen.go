// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaLibrary] class.
var (
	MediaLibraryClass     _MediaLibraryClass
	MediaLibraryClassOnce sync.Once
)

func getMediaLibraryClass() _MediaLibraryClass {
	MediaLibraryClassOnce.Do(func() {
		MediaLibraryClass = _MediaLibraryClass{objc.GetClass("MPMediaLibrary")}
	})
	return MediaLibraryClass
}

type _MediaLibraryClass struct {
	class objc.Class
}

// An interface definition for the [MediaLibrary] class.
type IMediaLibrary interface {
	objectivec.IObject
	// properties:
	// methods:
}

// An object that represents the state of synced media items on a device.
//
// A user may sync their device, changing the contents on the device, while your app is running. You can use the notification provided by this class to ensure that your app’s cache of the user’s library is up-to-date. To retrieve media items from the media library, build a custom query as described in and .


// An object that represents the state of synced media items on a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary
type MediaLibrary struct {
	objectivec.Object
}

// MediaLibraryFrom constructs a [MediaLibrary] from an unsafe.Pointer.
//
// An object that represents the state of synced media items on a device.
func MediaLibraryFrom(ptr unsafe.Pointer) MediaLibrary {
	return MediaLibrary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryClass) Alloc() MediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaLibraryClass) New() MediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaLibrary) Init() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaLibrary) Autorelease() MediaLibrary {
	rv := objc.Send[MediaLibrary](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaLibrary creates a new MediaLibrary instance.
func NewMediaLibrary() MediaLibrary {
	return getMediaLibraryClass().New()
}



// Returns whether the app can access the user’s media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/authorizationStatus()
func (mc _MediaLibraryClass) AuthorizationStatus() MediaLibraryAuthorizationStatus {
	rv := objc.Send[MediaLibraryAuthorizationStatus](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}


// Returns an instance of the default media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/default()
func (mc _MediaLibraryClass) DefaultMediaLibrary() IMediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("defaultMediaLibrary"))
	return rv
}


// Displays a user interface so that the user can authorize whether your app may view the media library’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/requestAuthorization(_:)
func (mc _MediaLibraryClass) RequestAuthorization(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("requestAuthorization:"), completionHandler)
}


