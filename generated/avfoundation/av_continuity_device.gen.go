// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
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
	AudioSessionInputs() objc.IObject /* cross-framework: AudioSessionPortDescription */
	SetAudioSessionInputs(value objc.IObject /* cross-framework: AudioSessionPortDescription */)
	ConnectionID() objc.IObject /* cross-framework: UUID */
	SetConnectionID(value objc.IObject /* cross-framework: UUID */)
	IsConnected() bool
	SetIsConnected(value bool)
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



// An array of the continuity device’s audio session port descriptions that’s available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/audiosessioninputs
func (c_ ContinuityDevice) AudioSessionInputs() objc.IObject /* cross-framework: AudioSessionPortDescription */ {
	rv := objc.Send[avfaudio.AudioSessionPortDescription](c_.ID, objc.Sel("audioSessionInputs"))
	return rv
}


// An array of the continuity device’s audio session port descriptions that’s available to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/audiosessioninputs
func (c_ ContinuityDevice) SetAudioSessionInputs(value objc.IObject /* cross-framework: AudioSessionPortDescription */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSessionInputs:"), value)
}


// A universally unique value that identifies a specific continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/connectionid
func (c_ ContinuityDevice) ConnectionID() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("connectionID"))
	return rv
}


// A universally unique value that identifies a specific continuity device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/connectionid
func (c_ ContinuityDevice) SetConnectionID(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConnectionID:"), value)
}


// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) IsConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}


// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}


