// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureFraming */


/* debug [class_header]: Header for AVCaptureFraming */
// The class instance for the [CaptureFraming] class.
var (
	CaptureFramingClass     _CaptureFramingClass
	CaptureFramingClassOnce sync.Once
)

func getCaptureFramingClass() _CaptureFramingClass {
	CaptureFramingClassOnce.Do(func() {
		CaptureFramingClass = _CaptureFramingClass{objc.GetClass("AVCaptureFraming")}
	})
	return CaptureFramingClass
}

type _CaptureFramingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureFraming */
// An interface definition for the [CaptureFraming] class.
type ICaptureFraming interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureFraming */
	// properties:
	SmartFramingMonitor() IAVCaptureSmartFramingMonitor
	SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureFraming */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureFraming */
// Alloc allocates a new instance without initialization.
func (cc _CaptureFramingClass) Alloc() CaptureFraming {
	rv := objc.Send[CaptureFraming](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureFramingClass) New() CaptureFraming {
	rv := objc.Send[CaptureFraming](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureFraming) Init() CaptureFraming {
	rv := objc.Send[CaptureFraming](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureFraming) Autorelease() CaptureFraming {
	rv := objc.Send[CaptureFraming](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureFraming creates a new CaptureFraming instance.
func NewCaptureFraming() CaptureFraming {
	return getCaptureFramingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureFraming */
// A framing, consisting of an aspect ratio and a zoom factor.
//
// An provides framing recommendations using this object.


// A framing, consisting of an aspect ratio and a zoom factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFraming
type CaptureFraming struct {
	objectivec.Object
}

// CaptureFramingFrom constructs a [CaptureFraming] from an unsafe.Pointer.
//
// A framing, consisting of an aspect ratio and a zoom factor.
func CaptureFramingFrom(ptr unsafe.Pointer) CaptureFraming {
	return CaptureFraming{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureFraming *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureFraming */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureFraming */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureFraming */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureFraming */

// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureFraming) SmartFramingMonitor() IAVCaptureSmartFramingMonitor {
	rv := objc.Send[CaptureSmartFramingMonitor](c_.ID, objc.Sel("smartFramingMonitor"))
	return rv
}/* debug [instance_properties/getter]: smartFramingMonitor */


// A monitor owned by the device that recommends an optimal framing based on the content in the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/smartframingmonitor
func (c_ CaptureFraming) SetSmartFramingMonitor(value IAVCaptureSmartFramingMonitor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSmartFramingMonitor:"), value)
}/* debug [instance_properties/setter]: smartFramingMonitor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureFraming */


