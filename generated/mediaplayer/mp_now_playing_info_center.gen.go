// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NowPlayingInfoCenter] class.
var (
	NowPlayingInfoCenterClass     _NowPlayingInfoCenterClass
	NowPlayingInfoCenterClassOnce sync.Once
)

func getNowPlayingInfoCenterClass() _NowPlayingInfoCenterClass {
	NowPlayingInfoCenterClassOnce.Do(func() {
		NowPlayingInfoCenterClass = _NowPlayingInfoCenterClass{objc.GetClass("MPNowPlayingInfoCenter")}
	})
	return NowPlayingInfoCenterClass
}

type _NowPlayingInfoCenterClass struct {
	class objc.Class
}

// An interface definition for the [NowPlayingInfoCenter] class.
type INowPlayingInfoCenter interface {
	objectivec.IObject
}

// An object for setting the Now Playing information for media that your app plays.
//
// If your app also provides Now Playing information containing information about the current track, use this object to update that information at appropriate times. This object contains a dictionary describing the playing item. The system displays Now Playing information on the device’s Lock Screen and in the media controls in Control Center. If the user directs playback of your media to Apple TV using AirPlay, the Now Playing information appears on the television screen. If the user connects a device to an iPod accessory, such as in a car, the accessory may display Now Playing information. The information you can specify includes all of the Now Playing metadata properties (see the Accessing Now Playing metadata properties topic group below), and the following subset of properties: You don’t have direct control over what information the system displays, or its formatting. You set the values in the dictionary and the system or the connected accessory handles displaying the information in a consistent manner for all apps. You can ensure that your app interacts well with other apps providing Now Playing information by following the best practices in the sample code project.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter
type NowPlayingInfoCenter struct {
	objectivec.Object
}

// NowPlayingInfoCenterFrom constructs a [NowPlayingInfoCenter] from an unsafe.Pointer.
//
// An object for setting the Now Playing information for media that your app plays.
func NowPlayingInfoCenterFrom(ptr unsafe.Pointer) NowPlayingInfoCenter {
	return NowPlayingInfoCenter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NowPlayingInfoCenterClass) Alloc() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NowPlayingInfoCenterClass) New() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NowPlayingInfoCenter) Init() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NowPlayingInfoCenter) Autorelease() NowPlayingInfoCenter {
	rv := objc.Send[NowPlayingInfoCenter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNowPlayingInfoCenter creates a new NowPlayingInfoCenter instance.
func NewNowPlayingInfoCenter() NowPlayingInfoCenter {
	return getNowPlayingInfoCenterClass().New()
}


// Returns the singleton Now Playing info center.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/default()
func (nc _NowPlayingInfoCenterClass) DefaultCenter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("defaultCenter"))
	return rv
}

// Keys related to animated artwork that are supported by the current platform.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/supportedAnimatedArtworkKeys
func (nc _NowPlayingInfoCenterClass) SupportedAnimatedArtworkKeys() []string {
	rv := objc.Send[[]string](objc.ID(nc.class), objc.Sel("supportedAnimatedArtworkKeys"))
	return rv
}
// The current Now Playing information for the default Now Playing info center.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/nowPlayingInfo
func (n_ NowPlayingInfoCenter) NowPlayingInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("nowPlayingInfo"))
	return rv
}


// SetNowPlayingInfo sets the value of the nowPlayingInfo property.
// The current Now Playing information for the default Now Playing info center.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/nowPlayingInfo
func (n_ NowPlayingInfoCenter) SetNowPlayingInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNowPlayingInfo:"), value)
}

// The current playback state of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/playbackState
func (n_ NowPlayingInfoCenter) PlaybackState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("playbackState"))
	return rv
}


// SetPlaybackState sets the value of the playbackState property.
// The current playback state of the app.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/playbackState
func (n_ NowPlayingInfoCenter) SetPlaybackState(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPlaybackState:"), value)
}

// Keys related to animated artwork that are supported by the current platform.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPNowPlayingInfoCenter/supportedAnimatedArtworkKeys
func (n_ NowPlayingInfoCenter) SupportedAnimatedArtworkKeys() []string {
	rv := objc.Send[[]string](n_.ID, objc.Sel("supportedAnimatedArtworkKeys"))
	return rv
}



