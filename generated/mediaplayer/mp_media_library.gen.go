// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	AddItemWithProductIDCompletionHandler(productID string, completionHandler unsafe.Pointer)
	BeginGeneratingLibraryChangeNotifications()
	EndGeneratingLibraryChangeNotifications()
	GetPlaylistWithUUIDCreationMetadataCompletionHandler(uuid unsafe.Pointer, creationMetadata unsafe.Pointer, completionHandler unsafe.Pointer)
}

// An object that represents the state of synced media items on a device.
//
// A user may sync their device, changing the contents on the device, while your app is running. You can use the notification provided by this class to ensure that your app’s cache of the user’s library is up-to-date. To retrieve media items from the media library, build a custom query as described in and .
//
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
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/authorizationStatus()
func (mc _MediaLibraryClass) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}

// Returns an instance of the default media library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/default()
func (mc _MediaLibraryClass) DefaultMediaLibrary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("defaultMediaLibrary"))
	return rv
}

// Displays a user interface so that the user can authorize whether your app may view the media library’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/requestAuthorization(_:)
func (mc _MediaLibraryClass) RequestAuthorization(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("requestAuthorization:"), completionHandler)
}

// Adds the designated item to the user’s music library.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/addItem(withProductID:completionHandler:)
func (m_ MediaLibrary) AddItemWithProductIDCompletionHandler(productID string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addItemWithProductID:completionHandler:"), objc.String(productID), completionHandler)
}

// Asks the media library to turn on notifications for whenever the library changes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/beginGeneratingLibraryChangeNotifications()
func (m_ MediaLibrary) BeginGeneratingLibraryChangeNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("beginGeneratingLibraryChangeNotifications"))
}

// Asks the media library to turn off notifications for whenever the library changes.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/endGeneratingLibraryChangeNotifications()
func (m_ MediaLibrary) EndGeneratingLibraryChangeNotifications() {
	objc.Send[objc.ID](m_.ID, objc.Sel("endGeneratingLibraryChangeNotifications"))
}

// Retrieves an app maintained existing playlist or creates a new playlist when no playlist exists.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/getPlaylist(with:creationMetadata:completionHandler:)
func (m_ MediaLibrary) GetPlaylistWithUUIDCreationMetadataCompletionHandler(uuid unsafe.Pointer, creationMetadata unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("getPlaylistWithUUID:creationMetadata:completionHandler:"), uuid, creationMetadata, completionHandler)
}

// The calendar date on which the media library was last modified.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/lastModifiedDate
func (m_ MediaLibrary) LastModifiedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lastModifiedDate"))
	return rv
}



