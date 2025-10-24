// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaLibrary */


/* debug [class_header]: Header for MPMediaLibrary */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaLibrary */
// An interface definition for the [MediaLibrary] class.
type IMediaLibrary interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaLibrary */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaLibrary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaLibrary */
// Alloc allocates a new instance without initialization.
func (mc _MediaLibraryClass) Alloc() MediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaLibrary */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaLibrary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaLibrary */

// Returns whether the app can access the user’s media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/authorizationStatus()
func (mc _MediaLibraryClass) AuthorizationStatus() MediaLibraryAuthorizationStatus {
	rv := objc.Send[MediaLibraryAuthorizationStatus](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AuthorizationStatus) */


// Returns an instance of the default media library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/default()
func (mc _MediaLibraryClass) DefaultMediaLibrary() IMediaLibrary {
	rv := objc.Send[MediaLibrary](objc.ID(mc.class), objc.Sel("defaultMediaLibrary"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultMediaLibrary) */


// Displays a user interface so that the user can authorize whether your app may view the media library’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaLibrary/requestAuthorization(_:)
func (mc _MediaLibraryClass) RequestAuthorization(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("requestAuthorization:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestAuthorization) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaLibrary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaLibrary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaLibrary */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaLibrary */


