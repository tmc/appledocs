// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPPlayableContentManager */


/* debug [class_header]: Header for MPPlayableContentManager */
// The class instance for the [PlayableContentManager] class.
var (
	PlayableContentManagerClass     _PlayableContentManagerClass
	PlayableContentManagerClassOnce sync.Once
)

func getPlayableContentManagerClass() _PlayableContentManagerClass {
	PlayableContentManagerClassOnce.Do(func() {
		PlayableContentManagerClass = _PlayableContentManagerClass{objc.GetClass("MPPlayableContentManager")}
	})
	return PlayableContentManagerClass
}

type _PlayableContentManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayableContentManager */
// An interface definition for the [PlayableContentManager] class.
type IPlayableContentManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayableContentManager */
	// properties:
	ImageCropRect() corefoundation.CGRect
	SetImageCropRect(value corefoundation.CGRect)
	ShowsRouteButton() bool
	SetShowsRouteButton(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayableContentManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayableContentManager */
// Alloc allocates a new instance without initialization.
func (pc _PlayableContentManagerClass) Alloc() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayableContentManagerClass) New() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayableContentManager) Init() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayableContentManager) Autorelease() PlayableContentManager {
	rv := objc.Send[PlayableContentManager](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayableContentManager creates a new PlayableContentManager instance.
func NewPlayableContentManager() PlayableContentManager {
	return getPlayableContentManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayableContentManager */
// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
//
// The app provides data to the content manager so that the media player can browse the content provided. A delegate provides the media player the ability to perform actions that manage the app’s playback queue. You don’t create a new content manager directly, instead you grab the shared content manager using the method. After getting the shared content manager, your next step depends on the features your app supports: To provide content navigation and suggested content for CarPlay, immediately set both the and properties. After setting these properties, use the and methods to load the information from the data source. To provide suggested content when the user connects headphones, a Bluetooth stereo, or another output device, set only the property. After you set a delegate, iOS automatically calls methods in the protocol allowing you to suggest appropriate content.


// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager
type PlayableContentManager struct {
	objectivec.Object
}

// PlayableContentManagerFrom constructs a [PlayableContentManager] from an unsafe.Pointer.
//
// A shared content manager for controlling interactions between your media app and system-provided or external media player interfaces.
func PlayableContentManagerFrom(ptr unsafe.Pointer) PlayableContentManager {
	return PlayableContentManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayableContentManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayableContentManager */

// Returns the current content manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManager/shared()
func (pc _PlayableContentManagerClass) SharedContentManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("sharedContentManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedContentManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayableContentManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayableContentManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayableContentManager */

// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) ImageCropRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p_.ID, objc.Sel("imageCropRect"))
	return rv
}/* debug [instance_properties/getter]: imageCropRect */


// The bounds, in points, of the content area for the full size image associated with the media item artwork.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/imagecroprect
func (p_ PlayableContentManager) SetImageCropRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageCropRect:"), value)
}/* debug [instance_properties/setter]: imageCropRect */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) ShowsRouteButton() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsRouteButton"))
	return rv
}/* debug [instance_properties/getter]: showsRouteButton */


// A Boolean value that indicates whether the route button is visible in the volume view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpvolumeview/showsroutebutton
func (p_ PlayableContentManager) SetShowsRouteButton(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsRouteButton:"), value)
}/* debug [instance_properties/setter]: showsRouteButton */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPPlayableContentManager */


