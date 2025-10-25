// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVPlayerItemTrack */


/* debug [class_header]: Header for AVPlayerItemTrack */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PlayerItemTrack */
// An interface definition for the [PlayerItemTrack] class.
type IPlayerItemTrack interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PlayerItemTrack */
	// properties:
	AssetTrack() IAVAssetTrack
	CurrentVideoFrameRate() float32
	Enabled() bool
	SetEnabled(value bool)
	VideoFieldMode() objc.IObject /* cross-framework: NSString */
	SetVideoFieldMode(value objc.IObject /* cross-framework: NSString */)
	IsEnabled() bool
	SetIsEnabled(value bool)
	AVPlayerItemTrackVideoFieldModeDeinterlaceFields() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PlayerItemTrack */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PlayerItemTrack */
// Alloc allocates a new instance without initialization.
func (pc _PlayerItemTrackClass) Alloc() PlayerItemTrack {
	rv := objc.Send[PlayerItemTrack](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PlayerItemTrack */
// An object that represents the presentation state of an asset track during playback.


// An object that represents the presentation state of an asset track during playback.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PlayerItemTrack *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PlayerItemTrack */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PlayerItemTrack */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PlayerItemTrack */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PlayerItemTrack */

// An asset track that provides the media for the player item track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/assetTrack
func (p_ PlayerItemTrack) AssetTrack() IAVAssetTrack {
	rv := objc.Send[AssetTrack](p_.ID, objc.Sel("assetTrack"))
	return rv
}/* debug [instance_properties/getter]: assetTrack */


// The current frame rate of the video track as it plays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/currentVideoFrameRate
func (p_ PlayerItemTrack) CurrentVideoFrameRate() float32 {
	rv := objc.Send[float32](p_.ID, objc.Sel("currentVideoFrameRate"))
	return rv
}/* debug [instance_properties/getter]: currentVideoFrameRate */


// A Boolean value that indicates whether the player item presents the track’s media during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/isEnabled
func (p_ PlayerItemTrack) Enabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the player item presents the track’s media during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/isEnabled
func (p_ PlayerItemTrack) SetEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// A mode that specifies the handling of video frames that contain multiple fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/videoFieldMode
func (p_ PlayerItemTrack) VideoFieldMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("videoFieldMode"))
	return rv
}/* debug [instance_properties/getter]: videoFieldMode */


// A mode that specifies the handling of video frames that contain multiple fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/videoFieldMode
func (p_ PlayerItemTrack) SetVideoFieldMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVideoFieldMode:"), value)
}/* debug [instance_properties/setter]: videoFieldMode */


// A Boolean value that indicates whether the player item presents the track’s media during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/isenabled
func (p_ PlayerItemTrack) IsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the player item presents the track’s media during playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/isenabled
func (p_ PlayerItemTrack) SetIsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// A video field mode that requests deinterlacing of video fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrackvideofieldmodedeinterlacefields
func (p_ PlayerItemTrack) AVPlayerItemTrackVideoFieldModeDeinterlaceFields() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("AVPlayerItemTrackVideoFieldModeDeinterlaceFields"))
	return rv
}/* debug [instance_properties/getter]: AVPlayerItemTrackVideoFieldModeDeinterlaceFields */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVPlayerItemTrack */



