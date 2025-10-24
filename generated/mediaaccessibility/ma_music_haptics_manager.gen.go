// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MAMusicHapticsManager */


/* debug [class_header]: Header for MAMusicHapticsManager */
// The class instance for the [MAMusicHapticsManager] class.
var (
	MAMusicHapticsManagerClass     _MAMusicHapticsManagerClass
	MAMusicHapticsManagerClassOnce sync.Once
)

func getMAMusicHapticsManagerClass() _MAMusicHapticsManagerClass {
	MAMusicHapticsManagerClassOnce.Do(func() {
		MAMusicHapticsManagerClass = _MAMusicHapticsManagerClass{objc.GetClass("MAMusicHapticsManager")}
	})
	return MAMusicHapticsManagerClass
}

type _MAMusicHapticsManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MAMusicHapticsManager */
// An interface definition for the [MAMusicHapticsManager] class.
type IMAMusicHapticsManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MAMusicHapticsManager */
	// properties:
	IsActive() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MAMusicHapticsManager */
	// methods:
	AddStatusObserver(statusHandler unsafe.Pointer) unsafe.Pointer
	CheckHapticTrackAvailabilityForMediaMatchingCodeCompletionHandler(internationalStandardRecordingCode objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	RemoveStatusObserver(registrationToken unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MAMusicHapticsManager */
// Alloc allocates a new instance without initialization.
func (mc _MAMusicHapticsManagerClass) Alloc() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MAMusicHapticsManagerClass) New() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MAMusicHapticsManager) Init() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MAMusicHapticsManager) Autorelease() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMAMusicHapticsManager creates a new MAMusicHapticsManager instance.
func NewMAMusicHapticsManager() MAMusicHapticsManager {
	return getMAMusicHapticsManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MAMusicHapticsManager */
// A class that reports information about the Music Haptics feature.
//
// Use the instance of to check information about the Music Haptics feature so you can respond accordingly in your app. For example, you can check whether Music Haptics is on, get a notification when it turns on or off, or check whether a haptic track is currently playing along with the Now Playing item.


// A class that reports information about the Music Haptics feature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager
type MAMusicHapticsManager struct {
	objectivec.Object
}

// MAMusicHapticsManagerFrom constructs a [MAMusicHapticsManager] from an unsafe.Pointer.
//
// A class that reports information about the Music Haptics feature.
func MAMusicHapticsManagerFrom(ptr unsafe.Pointer) MAMusicHapticsManager {
	return MAMusicHapticsManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MAMusicHapticsManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MAMusicHapticsManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MAMusicHapticsManager */

// The shared Music Haptics manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/shared
func (mc _MAMusicHapticsManagerClass) SharedManager() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](objc.ID(mc.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_properties_class/property]: sharedManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MAMusicHapticsManager */

// Adds an observer to monitor the status of haptic playback for the Now Playing song.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/addStatusObserver(_:)
func (m_ MAMusicHapticsManager) AddStatusObserver(statusHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addStatusObserver:"), statusHandler)
	return rv
}/* debug [instance_methods/method]: AddStatusObserver */


// Checks whether a haptic track is available for the song with the specified International Standard Recording Code (ISRC).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/checkHapticTrackAvailabilityForMedia(matchingCode:completionHandler:)
func (m_ MAMusicHapticsManager) CheckHapticTrackAvailabilityForMediaMatchingCodeCompletionHandler(internationalStandardRecordingCode objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("checkHapticTrackAvailabilityForMediaMatchingCode:completionHandler:"), internationalStandardRecordingCode, completionHandler)
}/* debug [instance_methods/method]: CheckHapticTrackAvailabilityForMediaMatchingCodeCompletionHandler */


// Removes the observer monitoring the status of haptic playback for the Now Playing song.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/removeStatusObserver(_:)
func (m_ MAMusicHapticsManager) RemoveStatusObserver(registrationToken unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeStatusObserver:"), registrationToken)
}/* debug [instance_methods/method]: RemoveStatusObserver */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MAMusicHapticsManager */

// A Boolean value that indicates whether the system setting for Music Haptics is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/isActive
func (m_ MAMusicHapticsManager) IsActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// The shared Music Haptics manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/shared
func (m_ MAMusicHapticsManager) SharedManager() IMAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](m_.ID, objc.Sel("sharedManager"))
	return rv
}/* debug [instance_properties/getter]: sharedManager */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MAMusicHapticsManager */






