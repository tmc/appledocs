// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MAMusicHapticsManager] class.
type IMAMusicHapticsManager interface {
	objectivec.IObject
	AddStatusObserver(statusHandler unsafe.Pointer) objc.ID
	CheckHapticTrackAvailabilityForMediaMatchingCodeCompletionHandler(internationalStandardRecordingCode appkit.string, completionHandler unsafe.Pointer)
	RemoveStatusObserver(registrationToken objectivec.IObject)
}

// A class that reports information about the Music Haptics feature.
//
// Use the instance of to check information about the Music Haptics feature so you can respond accordingly in your app. For example, you can check whether Music Haptics is on, get a notification when it turns on or off, or check whether a haptic track is currently playing along with the Now Playing item.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MAMusicHapticsManagerClass) Alloc() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The shared Music Haptics manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/shared
func (mc _MAMusicHapticsManagerClass) SharedManager() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](objc.ID(mc.class), objc.Sel("sharedManager"))
	return rv
}
// Adds an observer to monitor the status of haptic playback for the Now Playing song.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/addStatusObserver(_:)
func (m_ MAMusicHapticsManager) AddStatusObserver(statusHandler unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("addStatusObserver:"), statusHandler)
	return rv
}

// Checks whether a haptic track is available for the song with the specified International Standard Recording Code (ISRC).
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/checkHapticTrackAvailabilityForMedia(matchingCode:completionHandler:)
func (m_ MAMusicHapticsManager) CheckHapticTrackAvailabilityForMediaMatchingCodeCompletionHandler(internationalStandardRecordingCode appkit.string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("checkHapticTrackAvailabilityForMediaMatchingCode:completionHandler:"), internationalStandardRecordingCode, completionHandler)
}

// Removes the observer monitoring the status of haptic playback for the Now Playing song.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/removeStatusObserver(_:)
func (m_ MAMusicHapticsManager) RemoveStatusObserver(registrationToken objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeStatusObserver:"), registrationToken)
}

// A Boolean value that indicates whether the system setting for Music Haptics is on.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/isActive
func (m_ MAMusicHapticsManager) IsActive() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isActive"))
	return rv
}

// The shared Music Haptics manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAMusicHapticsManager/shared
func (m_ MAMusicHapticsManager) SharedManager() MAMusicHapticsManager {
	rv := objc.Send[MAMusicHapticsManager](m_.ID, objc.Sel("sharedManager"))
	return rv
}




