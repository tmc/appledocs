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
}

// An object that represents a stream of data that a capture input provides.
//
// Instances of have one or more input ports, one for each data stream they can produce. For example, an object presenting one video data stream has one port.
//
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




