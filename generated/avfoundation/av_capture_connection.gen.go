// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CaptureConnection] class.
var (
	CaptureConnectionClass     _CaptureConnectionClass
	CaptureConnectionClassOnce sync.Once
)

func getCaptureConnectionClass() _CaptureConnectionClass {
	CaptureConnectionClassOnce.Do(func() {
		CaptureConnectionClass = _CaptureConnectionClass{objc.GetClass("AVCaptureConnection")}
	})
	return CaptureConnectionClass
}

type _CaptureConnectionClass struct {
	class objc.Class
}

// An interface definition for the [CaptureConnection] class.
type ICaptureConnection interface {
	objectivec.IObject
}

// An object that represents a connection from a capture input to a capture output.
//
// Capture inputs have one or more input ports (instances of ). Capture outputs can accept data from one or more sources (for example, an object accepts both video and audio data). You can add an instance to a session using the method only if the method returns . When using the or method, the session forms connections automatically between all compatible inputs and outputs. You only need to add connections manually when adding an input or output with no connections. You can also use connections to enable or disable the flow of data from a given input or to a given output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection
type CaptureConnection struct {
	objectivec.Object
}

// CaptureConnectionFrom constructs a [CaptureConnection] from an unsafe.Pointer.
//
// An object that represents a connection from a capture input to a capture output.
func CaptureConnectionFrom(ptr unsafe.Pointer) CaptureConnection {
	return CaptureConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureConnectionClass) Alloc() CaptureConnection {
	rv := objc.Send[CaptureConnection](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureConnectionClass) New() CaptureConnection {
	rv := objc.Send[CaptureConnection](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureConnection) Init() CaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureConnection) Autorelease() CaptureConnection {
	rv := objc.Send[CaptureConnection](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureConnection creates a new CaptureConnection instance.
func NewCaptureConnection() CaptureConnection {
	return getCaptureConnectionClass().New()
}


// The smallest time interval the connection can apply between consecutive video frames.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) VideoMinFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoMinFrameDuration"))
	return rv
}


// SetVideoMinFrameDuration sets the value of the videoMinFrameDuration property.
// The smallest time interval the connection can apply between consecutive video frames.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c_ CaptureConnection) SetVideoMinFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}
// An orientation that tells the connection how to rotate a video flowing through it.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) VideoOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoOrientation"))
	return rv
}


// SetVideoOrientation sets the value of the videoOrientation property.
// An orientation that tells the connection how to rotate a video flowing through it.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c_ CaptureConnection) SetVideoOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoOrientation:"), value)
}


