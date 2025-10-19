// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureMultiCamSession] class.
var aVCaptureMultiCamSessionClass = _AVCaptureMultiCamSessionClass{objc.GetClass("AVCaptureMultiCamSession")}

type _AVCaptureMultiCamSessionClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureMultiCamSession] class.
type IAVCaptureMultiCamSession interface {
	IAVCaptureSession
}

// A capture session that supports simultaneous capture from multiple inputs of the same media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession

type AVCaptureMultiCamSession struct {
	AVCaptureSession
}

// AVCaptureMultiCamSessionFrom constructs a [AVCaptureMultiCamSession] from an unsafe.Pointer.
//
// A capture session that supports simultaneous capture from multiple inputs of the same media type.
func AVCaptureMultiCamSessionFrom(ptr unsafe.Pointer) AVCaptureMultiCamSession {
	return AVCaptureMultiCamSession{
		AVCaptureSession: AVCaptureSessionFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureMultiCamSessionClass) Alloc() AVCaptureMultiCamSession {
	rv := objc.Send[AVCaptureMultiCamSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureMultiCamSessionClass) New() AVCaptureMultiCamSession {
	rv := objc.Send[AVCaptureMultiCamSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureMultiCamSession) Init() AVCaptureMultiCamSession {
	rv := objc.Send[AVCaptureMultiCamSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureMultiCamSession) Autorelease() AVCaptureMultiCamSession {
	rv := objc.Send[AVCaptureMultiCamSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureMultiCamSession creates a new AVCaptureMultiCamSession instance.
func NewAVCaptureMultiCamSession() AVCaptureMultiCamSession {
	return aVCaptureMultiCamSessionClass.New()
}




