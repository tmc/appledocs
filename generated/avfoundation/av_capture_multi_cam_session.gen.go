// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureMultiCamSession] class.
var (
	CaptureMultiCamSessionClass     _CaptureMultiCamSessionClass
	CaptureMultiCamSessionClassOnce sync.Once
)

func getCaptureMultiCamSessionClass() _CaptureMultiCamSessionClass {
	CaptureMultiCamSessionClassOnce.Do(func() {
		CaptureMultiCamSessionClass = _CaptureMultiCamSessionClass{objc.GetClass("AVCaptureMultiCamSession")}
	})
	return CaptureMultiCamSessionClass
}

type _CaptureMultiCamSessionClass struct {
	class objc.Class
}

// An interface definition for the [CaptureMultiCamSession] class.
type ICaptureMultiCamSession interface {
	ICaptureSession
	// properties:
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	HardwareCost() float32 /* primitive/slice/pointer. */
	SetHardwareCost(value float32 /* primitive/slice/pointer. */)
	SystemPressureCost() float32 /* primitive/slice/pointer. */
	SetSystemPressureCost(value float32 /* primitive/slice/pointer. */)
	// methods:
}

// A capture session that supports simultaneous capture from multiple inputs of the same media type.
//
// The session preset for a multicamera session is always . Set each capture device’s value to the desired quality of service. You can dynamically enable and disable this session’s individual camera inputs without interrupting capture preview. To stop an individual camera, disable all of its connections or connected ports. The camera then stops streaming data to save power and bandwidth. Other inputs that are streaming data through the session are unaffected.


// A capture session that supports simultaneous capture from multiple inputs of the same media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
type CaptureMultiCamSession struct {
	CaptureSession
}

// CaptureMultiCamSessionFrom constructs a [CaptureMultiCamSession] from an unsafe.Pointer.
//
// A capture session that supports simultaneous capture from multiple inputs of the same media type.
func CaptureMultiCamSessionFrom(ptr unsafe.Pointer) CaptureMultiCamSession {
	return CaptureMultiCamSession{
		CaptureSession: CaptureSessionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureMultiCamSessionClass) Alloc() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureMultiCamSessionClass) New() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMultiCamSession) Init() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMultiCamSession) Autorelease() CaptureMultiCamSession {
	rv := objc.Send[CaptureMultiCamSession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMultiCamSession creates a new CaptureMultiCamSession instance.
func NewCaptureMultiCamSession() CaptureMultiCamSession {
	return getCaptureMultiCamSessionClass().New()
}



// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureMultiCamSession) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureMultiCamSession) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}


// A value that indicates the percentage of the session’s available hardware budget currently in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemulticamsession/hardwarecost
func (c_ CaptureMultiCamSession) HardwareCost() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("hardwareCost"))
	return rv
}


// A value that indicates the percentage of the session’s available hardware budget currently in use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemulticamsession/hardwarecost
func (c_ CaptureMultiCamSession) SetHardwareCost(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHardwareCost:"), value)
}


// A value that indicates the system pressure cost of the current session configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemulticamsession/systempressurecost
func (c_ CaptureMultiCamSession) SystemPressureCost() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](c_.ID, objc.Sel("systemPressureCost"))
	return rv
}


// A value that indicates the system pressure cost of the current session configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemulticamsession/systempressurecost
func (c_ CaptureMultiCamSession) SetSystemPressureCost(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemPressureCost:"), value)
}



