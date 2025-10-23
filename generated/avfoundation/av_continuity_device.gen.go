// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ContinuityDevice] class.
var (
	ContinuityDeviceClass     _ContinuityDeviceClass
	ContinuityDeviceClassOnce sync.Once
)

func getContinuityDeviceClass() _ContinuityDeviceClass {
	ContinuityDeviceClassOnce.Do(func() {
		ContinuityDeviceClass = _ContinuityDeviceClass{objc.GetClass("AVContinuityDevice")}
	})
	return ContinuityDeviceClass
}

type _ContinuityDeviceClass struct {
	class objc.Class
}

// An interface definition for the [ContinuityDevice] class.
type IContinuityDevice interface {
	objectivec.IObject
	// properties:
	VideoDevices() []CaptureDevice /* primitive/slice/pointer */
	AudioSessionInputs() AudioSessionPortDescription /* not a class type */
	SetAudioSessionInputs(value AudioSessionPortDescription /* not a class type */)
	ConnectionID() foundation.UUID /* not a class type */
	SetConnectionID(value foundation.UUID /* not a class type */)
	IsConnected() bool /* primitive/slice/pointer */
	SetIsConnected(value bool /* primitive/slice/pointer */)
	// methods:
}

// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
//
// Each continuity device instance represents another iOS device that’s nearby. Your app can access the other device’s cameras and microphones with its and properties, respectively.


// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice
type ContinuityDevice struct {
	objectivec.Object
}

// ContinuityDeviceFrom constructs a [ContinuityDevice] from an unsafe.Pointer.
//
// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
func ContinuityDeviceFrom(ptr unsafe.Pointer) ContinuityDevice {
	return ContinuityDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContinuityDeviceClass) Alloc() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContinuityDeviceClass) New() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContinuityDevice) Init() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContinuityDevice) Autorelease() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContinuityDevice creates a new ContinuityDevice instance.
func NewContinuityDevice() ContinuityDevice {
	return getContinuityDeviceClass().New()
}



// An array of the continuity device’s video-capture devices available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice/videoDevices
func (c_ ContinuityDevice) VideoDevices() []CaptureDevice /* primitive/slice/pointer */ {
	rv := objc.Send[[]CaptureDevice](c_.ID, objc.Sel("videoDevices"))
	return rv
}


// An array of the continuity device’s audio session port descriptions that’s available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/audiosessioninputs
func (c_ ContinuityDevice) AudioSessionInputs() AudioSessionPortDescription /* not a class type */ {
	rv := objc.Send[AudioSessionPortDescription](c_.ID, objc.Sel("audioSessionInputs"))
	return rv
}


// An array of the continuity device’s audio session port descriptions that’s available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/audiosessioninputs
func (c_ ContinuityDevice) SetAudioSessionInputs(value AudioSessionPortDescription /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSessionInputs:"), value)
}


// A universally unique value that identifies a specific continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/connectionid
func (c_ ContinuityDevice) ConnectionID() foundation.UUID /* not a class type */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("connectionID"))
	return rv
}


// A universally unique value that identifies a specific continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/connectionid
func (c_ ContinuityDevice) SetConnectionID(value foundation.UUID /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConnectionID:"), value)
}


// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) IsConnected() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) SetIsConnected(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}



