// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureAudioChannel */


/* debug [class_header]: Header for AVCaptureAudioChannel */
// The class instance for the [CaptureAudioChannel] class.
var (
	CaptureAudioChannelClass     _CaptureAudioChannelClass
	CaptureAudioChannelClassOnce sync.Once
)

func getCaptureAudioChannelClass() _CaptureAudioChannelClass {
	CaptureAudioChannelClassOnce.Do(func() {
		CaptureAudioChannelClass = _CaptureAudioChannelClass{objc.GetClass("AVCaptureAudioChannel")}
	})
	return CaptureAudioChannelClass
}

type _CaptureAudioChannelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureAudioChannel */
// An interface definition for the [CaptureAudioChannel] class.
type ICaptureAudioChannel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureAudioChannel */
	// properties:
	AveragePowerLevel() float32
	Enabled() bool
	SetEnabled(value bool)
	PeakHoldLevel() float32
	Volume() float32
	SetVolume(value float32)
	IsEnabled() bool
	SetIsEnabled(value bool)
	Connections() IAVCaptureConnection
	SetConnections(value IAVCaptureConnection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureAudioChannel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureAudioChannel */
// Alloc allocates a new instance without initialization.
func (cc _CaptureAudioChannelClass) Alloc() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureAudioChannelClass) New() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureAudioChannel) Init() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureAudioChannel) Autorelease() CaptureAudioChannel {
	rv := objc.Send[CaptureAudioChannel](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureAudioChannel creates a new CaptureAudioChannel instance.
func NewCaptureAudioChannel() CaptureAudioChannel {
	return getCaptureAudioChannelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureAudioChannel */
// An object that monitors average and peak power levels for an audio channel in a capture connection.
//
// You don’t create instances of this class directly. Instead, an object that connects an audio input to an audio output provides an array of objects, one for each channel of audio available. You can poll for audio levels by iterating through these audio channel objects.


// An object that monitors average and peak power levels for an audio channel in a capture connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel
type CaptureAudioChannel struct {
	objectivec.Object
}

// CaptureAudioChannelFrom constructs a [CaptureAudioChannel] from an unsafe.Pointer.
//
// An object that monitors average and peak power levels for an audio channel in a capture connection.
func CaptureAudioChannelFrom(ptr unsafe.Pointer) CaptureAudioChannel {
	return CaptureAudioChannel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureAudioChannel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureAudioChannel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureAudioChannel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureAudioChannel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureAudioChannel */

// The instantaneous average power level in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/averagePowerLevel
func (c_ CaptureAudioChannel) AveragePowerLevel() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("averagePowerLevel"))
	return rv
}/* debug [instance_properties/getter]: averagePowerLevel */


// A Boolean value that indicates whether the channel is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/isEnabled
func (c_ CaptureAudioChannel) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the channel is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/isEnabled
func (c_ CaptureAudioChannel) SetEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The peak hold power level in decibels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/peakHoldLevel
func (c_ CaptureAudioChannel) PeakHoldLevel() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("peakHoldLevel"))
	return rv
}/* debug [instance_properties/getter]: peakHoldLevel */


// The current volume (gain) of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/volume
func (c_ CaptureAudioChannel) Volume() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("volume"))
	return rv
}/* debug [instance_properties/getter]: volume */


// The current volume (gain) of the channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioChannel/volume
func (c_ CaptureAudioChannel) SetVolume(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVolume:"), value)
}/* debug [instance_properties/setter]: volume */


// A Boolean value that indicates whether the channel is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/isenabled
func (c_ CaptureAudioChannel) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the channel is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiochannel/isenabled
func (c_ CaptureAudioChannel) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */


// The connections between inputs and outputs that a capture session contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/connections
func (c_ CaptureAudioChannel) Connections() IAVCaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("connections"))
	return rv
}/* debug [instance_properties/getter]: connections */


// The connections between inputs and outputs that a capture session contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/connections
func (c_ CaptureAudioChannel) SetConnections(value IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConnections:"), value)
}/* debug [instance_properties/setter]: connections */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureAudioChannel */



