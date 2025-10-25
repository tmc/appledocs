// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSystemExposureBiasSlider */


/* debug [class_header]: Header for AVCaptureSystemExposureBiasSlider */
// The class instance for the [CaptureSystemExposureBiasSlider] class.
var (
	CaptureSystemExposureBiasSliderClass     _CaptureSystemExposureBiasSliderClass
	CaptureSystemExposureBiasSliderClassOnce sync.Once
)

func getCaptureSystemExposureBiasSliderClass() _CaptureSystemExposureBiasSliderClass {
	CaptureSystemExposureBiasSliderClassOnce.Do(func() {
		CaptureSystemExposureBiasSliderClass = _CaptureSystemExposureBiasSliderClass{objc.GetClass("AVCaptureSystemExposureBiasSlider")}
	})
	return CaptureSystemExposureBiasSliderClass
}

type _CaptureSystemExposureBiasSliderClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSystemExposureBiasSlider */
// An interface definition for the [CaptureSystemExposureBiasSlider] class.
type ICaptureSystemExposureBiasSlider interface {
	ICaptureControl
	
/* debug [class_interface_properties]: Properties for CaptureSystemExposureBiasSlider */
	// properties:
	SystemRecommendedExposureBiasRange() float32
	SetSystemRecommendedExposureBiasRange(value float32)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSystemExposureBiasSlider */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSystemExposureBiasSlider */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSystemExposureBiasSliderClass) Alloc() CaptureSystemExposureBiasSlider {
	rv := objc.Send[CaptureSystemExposureBiasSlider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSystemExposureBiasSliderClass) New() CaptureSystemExposureBiasSlider {
	rv := objc.Send[CaptureSystemExposureBiasSlider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSystemExposureBiasSlider) Init() CaptureSystemExposureBiasSlider {
	rv := objc.Send[CaptureSystemExposureBiasSlider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSystemExposureBiasSlider) Autorelease() CaptureSystemExposureBiasSlider {
	rv := objc.Send[CaptureSystemExposureBiasSlider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSystemExposureBiasSlider creates a new CaptureSystemExposureBiasSlider instance.
func NewCaptureSystemExposureBiasSlider() CaptureSystemExposureBiasSlider {
	return getCaptureSystemExposureBiasSliderClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSystemExposureBiasSlider */
// A control that adjusts the exposure bias of a capture device within the system-recommended range.
//
// This control defines its range by querying the property of the device’s active format. If a device’s value changes, the slider updates its range with the new format’s system-recommended value. To use this control, add it to the capture session by calling the session’s method.


// A control that adjusts the exposure bias of a capture device within the system-recommended range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider
type CaptureSystemExposureBiasSlider struct {
	CaptureControl
}

// CaptureSystemExposureBiasSliderFrom constructs a [CaptureSystemExposureBiasSlider] from an unsafe.Pointer.
//
// A control that adjusts the exposure bias of a capture device within the system-recommended range.
func CaptureSystemExposureBiasSliderFrom(ptr unsafe.Pointer) CaptureSystemExposureBiasSlider {
	return CaptureSystemExposureBiasSlider{
		CaptureControl: CaptureControlFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSystemExposureBiasSlider */

// Creates a slider to control the exposure bias of the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:)
func NewCaptureSystemExposureBiasSliderWithDevice(device IAVCaptureDevice) CaptureSystemExposureBiasSlider {
	instance := getCaptureSystemExposureBiasSliderClass().Alloc()
	rv := objc.Send[CaptureSystemExposureBiasSlider](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureSystemExposureBiasSliderWithDevice */


// Creates a slider to control the exposure bias of the specified capture device with an action to respond to exposure bias changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:action:)
func NewCaptureSystemExposureBiasSliderWithDeviceAction(device IAVCaptureDevice, action func(float32)) CaptureSystemExposureBiasSlider {
	instance := getCaptureSystemExposureBiasSliderClass().Alloc()
	rv := objc.Send[CaptureSystemExposureBiasSlider](instance.ID, objc.Sel("initWithDevice:action:"), device, action)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptureSystemExposureBiasSliderWithDeviceAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSystemExposureBiasSlider */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSystemExposureBiasSlider */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSystemExposureBiasSlider */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSystemExposureBiasSlider */

// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureSystemExposureBiasSlider) SystemRecommendedExposureBiasRange() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return rv
}/* debug [instance_properties/getter]: systemRecommendedExposureBiasRange */


// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureSystemExposureBiasSlider) SetSystemRecommendedExposureBiasRange(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemRecommendedExposureBiasRange:"), value)
}/* debug [instance_properties/setter]: systemRecommendedExposureBiasRange */


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemExposureBiasSlider) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}/* debug [instance_properties/getter]: activeFormat */


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemExposureBiasSlider) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}/* debug [instance_properties/setter]: activeFormat */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSystemExposureBiasSlider */


