// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureDeviceDiscoverySession] class.
var (
	aVCaptureDeviceDiscoverySessionClass     _AVCaptureDeviceDiscoverySessionClass
	aVCaptureDeviceDiscoverySessionClassOnce sync.Once
)

func getAVCaptureDeviceDiscoverySessionClass() _AVCaptureDeviceDiscoverySessionClass {
	aVCaptureDeviceDiscoverySessionClassOnce.Do(func() {
		aVCaptureDeviceDiscoverySessionClass = _AVCaptureDeviceDiscoverySessionClass{objc.GetClass("AVCaptureDeviceDiscoverySession")}
	})
	return aVCaptureDeviceDiscoverySessionClass
}

type _AVCaptureDeviceDiscoverySessionClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureDeviceDiscoverySession] class.
type IAVCaptureDeviceDiscoverySession interface {
	objectivec.IObject
}

// An object that finds capture devices that match specific search criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession
type AVCaptureDeviceDiscoverySession struct {
	objectivec.Object
}

// AVCaptureDeviceDiscoverySessionFrom constructs a [AVCaptureDeviceDiscoverySession] from an unsafe.Pointer.
//
// An object that finds capture devices that match specific search criteria.
func AVCaptureDeviceDiscoverySessionFrom(ptr unsafe.Pointer) AVCaptureDeviceDiscoverySession {
	return AVCaptureDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVCaptureDeviceDiscoverySessionClass) Alloc() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVCaptureDeviceDiscoverySessionClass) New() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureDeviceDiscoverySession) Init() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureDeviceDiscoverySession) Autorelease() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceDiscoverySession creates a new AVCaptureDeviceDiscoverySession instance.
func NewAVCaptureDeviceDiscoverySession() AVCaptureDeviceDiscoverySession {
	return getAVCaptureDeviceDiscoverySessionClass().New()
}




