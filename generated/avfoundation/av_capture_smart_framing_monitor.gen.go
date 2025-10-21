// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureSmartFramingMonitor] class.
var (
	CaptureSmartFramingMonitorClass     _CaptureSmartFramingMonitorClass
	CaptureSmartFramingMonitorClassOnce sync.Once
)

func getCaptureSmartFramingMonitorClass() _CaptureSmartFramingMonitorClass {
	CaptureSmartFramingMonitorClassOnce.Do(func() {
		CaptureSmartFramingMonitorClass = _CaptureSmartFramingMonitorClass{objc.GetClass("AVCaptureSmartFramingMonitor")}
	})
	return CaptureSmartFramingMonitorClass
}

type _CaptureSmartFramingMonitorClass struct {
	class objc.Class
}

// An interface definition for the [CaptureSmartFramingMonitor] class.
type ICaptureSmartFramingMonitor interface {
	objectivec.IObject
}

// An object associated with a capture device that monitors the scene and suggests an optimal framing.
//
// A smart framing monitor observes its associated device for objects of interest entering and exiting the camera’s field of view and recommends an optimal framing for good photographic composition. This framing recommendation consists of an aspect ratio and zoom factor. You may respond to the device’s framing recommendation by calling and setting on the associated device in whatever order best matches your animation between old and new framings.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor
type CaptureSmartFramingMonitor struct {
	objectivec.Object
}

// CaptureSmartFramingMonitorFrom constructs a [CaptureSmartFramingMonitor] from an unsafe.Pointer.
//
// An object associated with a capture device that monitors the scene and suggests an optimal framing.
func CaptureSmartFramingMonitorFrom(ptr unsafe.Pointer) CaptureSmartFramingMonitor {
	return CaptureSmartFramingMonitor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureSmartFramingMonitorClass) Alloc() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureSmartFramingMonitorClass) New() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSmartFramingMonitor) Init() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSmartFramingMonitor) Autorelease() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSmartFramingMonitor creates a new CaptureSmartFramingMonitor instance.
func NewCaptureSmartFramingMonitor() CaptureSmartFramingMonitor {
	return getCaptureSmartFramingMonitorClass().New()
}


// The latest recommended framing from the monitor.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSmartFramingMonitor/recommendedFraming
func (c_ CaptureSmartFramingMonitor) RecommendedFraming() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recommendedFraming"))
	return rv
}



