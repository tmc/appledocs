// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureInputPort] class.
var (
	CaptureInputPortClass     _CaptureInputPortClass
	CaptureInputPortClassOnce sync.Once
)

func getCaptureInputPortClass() _CaptureInputPortClass {
	CaptureInputPortClassOnce.Do(func() {
		CaptureInputPortClass = _CaptureInputPortClass{objc.GetClass("AVCaptureInputPort")}
	})
	return CaptureInputPortClass
}

type _CaptureInputPortClass struct {
	class objc.Class
}

// An interface definition for the [CaptureInputPort] class.
type ICaptureInputPort interface {
	objectivec.IObject
	Clock() unsafe.Pointer
	SetClock(value unsafe.Pointer)
	FormatDescription() unsafe.Pointer
	SetFormatDescription(value unsafe.Pointer)
	Input() AVCaptureInput
	SetInput(value IAVCaptureInput)
	IsEnabled() bool
	SetIsEnabled(value bool)
	MediaType() MediaType
	SetMediaType(value MediaType)
	SourceDevicePosition() unsafe.Pointer
	SetSourceDevicePosition(value unsafe.Pointer)
	SourceDeviceType() unsafe.Pointer
	SetSourceDeviceType(value unsafe.Pointer)
	Ports() AVCaptureInputPort
	SetPorts(value IAVCaptureInputPort)
}

// An object that represents a stream of data that a capture input provides.
//
// Instances of have one or more input ports, one for each data stream they can produce. For example, an object presenting one video data stream has one port.


// An object that represents a stream of data that a capture input provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port

type CaptureInputPort struct {
	objectivec.Object
}

// CaptureInputPortFrom constructs a [CaptureInputPort] from an unsafe.Pointer.
//
// An object that represents a stream of data that a capture input provides.
func CaptureInputPortFrom(ptr unsafe.Pointer) CaptureInputPort {
	return CaptureInputPort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureInputPortClass) Alloc() CaptureInputPort {
	rv := objc.Send[CaptureInputPort](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureInputPortClass) New() CaptureInputPort {
	rv := objc.Send[CaptureInputPort](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureInputPort) Init() CaptureInputPort {
	rv := objc.Send[CaptureInputPort](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureInputPort) Autorelease() CaptureInputPort {
	rv := objc.Send[CaptureInputPort](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureInputPort creates a new CaptureInputPort instance.
func NewCaptureInputPort() CaptureInputPort {
	return getCaptureInputPortClass().New()
}



// An object that represents the capture device’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/clock

func (c_ CaptureInputPort) Clock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("clock"))
	return rv
}


// An object that represents the capture device’s clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/clock

func (c_ CaptureInputPort) SetClock(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setClock:"), value)
}


// A description of the port format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/formatdescription

func (c_ CaptureInputPort) FormatDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("formatDescription"))
	return rv
}


// A description of the port format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/formatdescription

func (c_ CaptureInputPort) SetFormatDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFormatDescription:"), value)
}


// The input object that owns the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/input

func (c_ CaptureInputPort) Input() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](c_.ID, objc.Sel("input"))
	return rv
}


// The input object that owns the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/input

func (c_ CaptureInputPort) SetInput(value IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInput:"), value)
}


// A Boolean value that indicates whether the port is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/isenabled

func (c_ CaptureInputPort) IsEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value that indicates whether the port is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/isenabled

func (c_ CaptureInputPort) SetIsEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsEnabled:"), value)
}


// The media type of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/mediatype

func (c_ CaptureInputPort) MediaType() MediaType {
	rv := objc.Send[MediaType](c_.ID, objc.Sel("mediaType"))
	return rv
}


// The media type of the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/mediatype

func (c_ CaptureInputPort) SetMediaType(value MediaType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaType:"), value)
}


// The position of the source device providing input through this port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/sourcedeviceposition

func (c_ CaptureInputPort) SourceDevicePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sourceDevicePosition"))
	return rv
}


// The position of the source device providing input through this port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/sourcedeviceposition

func (c_ CaptureInputPort) SetSourceDevicePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceDevicePosition:"), value)
}


// The device type of the source camera that provides data to the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/sourcedevicetype

func (c_ CaptureInputPort) SourceDeviceType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sourceDeviceType"))
	return rv
}


// The device type of the source camera that provides data to the port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/port/sourcedevicetype

func (c_ CaptureInputPort) SetSourceDeviceType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSourceDeviceType:"), value)
}


// The ports available on a capture input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/ports

func (c_ CaptureInputPort) Ports() AVCaptureInputPort {
	rv := objc.Send[AVCaptureInputPort](c_.ID, objc.Sel("ports"))
	return rv
}


// The ports available on a capture input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureinput/ports

func (c_ CaptureInputPort) SetPorts(value IAVCaptureInputPort) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPorts:"), value)
}



