// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSmartFramingMonitor */


/* debug [class_header]: Header for AVCaptureSmartFramingMonitor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSmartFramingMonitor */
// An interface definition for the [CaptureSmartFramingMonitor] class.
type ICaptureSmartFramingMonitor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSmartFramingMonitor */
	// properties:
	SmartFramingMonitor() IAVCaptureSmartFramingMonitor
	SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor)
	VideoZoomFactor() float64
	SetVideoZoomFactor(value float64)
	IsMonitoring() bool
	SetIsMonitoring(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSmartFramingMonitor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSmartFramingMonitor */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSmartFramingMonitorClass) Alloc() CaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSmartFramingMonitor */
// An object associated with a capture device that monitors the scene and suggests an optimal framing.
//
// A smart framing monitor observes its associated device for objects of interest entering and exiting the camera’s field of view and recommends an optimal framing for good photographic composition. This framing recommendation consists of an aspect ratio and zoom factor. You may respond to the device’s framing recommendation by calling and setting on the associated device in whatever order best matches your animation between old and new framings.


// An object associated with a capture device that monitors the scene and suggests an optimal framing.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSmartFramingMonitor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSmartFramingMonitor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSmartFramingMonitor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSmartFramingMonitor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSmartFramingMonitor */

// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureSmartFramingMonitor) SmartFramingMonitor() IAVCaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}/* debug [instance_properties/getter]: smartFramingMonitor */


// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureSmartFramingMonitor) SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSmartFramingMonitor:"), value)
}/* debug [instance_properties/setter]: smartFramingMonitor */


// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c_ CaptureSmartFramingMonitor) VideoZoomFactor() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("videoZoomFactor"))
	return rv
}/* debug [instance_properties/getter]: videoZoomFactor */


// A value that controls the cropping and enlargement of images captured by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c_ CaptureSmartFramingMonitor) SetVideoZoomFactor(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoZoomFactor:"), value)
}/* debug [instance_properties/setter]: videoZoomFactor */


// Yes when the receiver is actively monitoring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesmartframingmonitor/ismonitoring
func (c_ CaptureSmartFramingMonitor) IsMonitoring() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMonitoring"))
	return rv
}/* debug [instance_properties/getter]: isMonitoring */


// Yes when the receiver is actively monitoring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesmartframingmonitor/ismonitoring
func (c_ CaptureSmartFramingMonitor) SetIsMonitoring(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMonitoring:"), value)
}/* debug [instance_properties/setter]: isMonitoring */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSmartFramingMonitor */


