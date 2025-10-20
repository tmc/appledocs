// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerItemTrack] class.
var (
	PlayerItemTrackClass     _PlayerItemTrackClass
	PlayerItemTrackClassOnce sync.Once
)

func getPlayerItemTrackClass() _PlayerItemTrackClass {
	PlayerItemTrackClassOnce.Do(func() {
		PlayerItemTrackClass = _PlayerItemTrackClass{objc.GetClass("AVPlayerItemTrack")}
	})
	return PlayerItemTrackClass
}

type _PlayerItemTrackClass struct {
	class objc.Class
}

// An interface definition for the [PlayerItemTrack] class.
type IPlayerItemTrack interface {
	objectivec.IObject
}

// An object that represents the presentation state of an asset track during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack
type PlayerItemTrack struct {
	objectivec.Object
}

// PlayerItemTrackFrom constructs a [PlayerItemTrack] from an unsafe.Pointer.
//
// An object that represents the presentation state of an asset track during playback.
func PlayerItemTrackFrom(ptr unsafe.Pointer) PlayerItemTrack {
	return PlayerItemTrack{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerItemTrackClass) Alloc() PlayerItemTrack {
	rv := objc.Send[PlayerItemTrack](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerItemTrackClass) New() PlayerItemTrack {
	rv := objc.Send[PlayerItemTrack](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerItemTrack) Init() PlayerItemTrack {
	rv := objc.Send[PlayerItemTrack](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerItemTrack) Autorelease() PlayerItemTrack {
	rv := objc.Send[PlayerItemTrack](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerItemTrack creates a new PlayerItemTrack instance.
func NewPlayerItemTrack() PlayerItemTrack {
	return getPlayerItemTrackClass().New()
}


// The current frame rate of the video track as it plays.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/currentVideoFrameRate
func (p_ PlayerItemTrack) CurrentVideoFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("currentVideoFrameRate"))
	return rv
}

// A Boolean value that indicates whether the player item presents the track’s media during playback.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/isEnabled
func (p_ PlayerItemTrack) Enabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("enabled"))
	return rv
}


// SetEnabled sets the value of the enabled property.
// A Boolean value that indicates whether the player item presents the track’s media during playback.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/isEnabled
func (p_ PlayerItemTrack) SetEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnabled:"), value)
}
// A mode that specifies the handling of video frames that contain multiple fields.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/videoFieldMode
func (p_ PlayerItemTrack) VideoFieldMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("videoFieldMode"))
	return rv
}


// SetVideoFieldMode sets the value of the videoFieldMode property.
// A mode that specifies the handling of video frames that contain multiple fields.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/videoFieldMode
func (p_ PlayerItemTrack) SetVideoFieldMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFieldMode:"), value)
}


