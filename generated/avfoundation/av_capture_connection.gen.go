// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureConnection] class.
var (
	aVCaptureConnectionClass     _AVCaptureConnectionClass
	aVCaptureConnectionClassOnce sync.Once
)

func getAVCaptureConnectionClass() _AVCaptureConnectionClass {
	aVCaptureConnectionClassOnce.Do(func() {
		aVCaptureConnectionClass = _AVCaptureConnectionClass{objc.GetClass("AVCaptureConnection")}
	})
	return aVCaptureConnectionClass
}

type _AVCaptureConnectionClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureConnection] class.
type IAVCaptureConnection interface {
	objectivec.IObject
}

// An object that represents a connection from a capture input to a capture output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection
type AVCaptureConnection struct {
	objectivec.Object
}

// AVCaptureConnectionFrom constructs a [AVCaptureConnection] from an unsafe.Pointer.
//
// An object that represents a connection from a capture input to a capture output.
func AVCaptureConnectionFrom(ptr unsafe.Pointer) AVCaptureConnection {
	return AVCaptureConnection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureConnectionClass) Alloc() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureConnectionClass) New() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureConnection) Init() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureConnection) Autorelease() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureConnection creates a new AVCaptureConnection instance.
func NewAVCaptureConnection() AVCaptureConnection {
	return getAVCaptureConnectionClass().New()
}




