// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptureSystemExposureBiasSlider] class.
type ICaptureSystemExposureBiasSlider interface {
	ICaptureControl
	

	// properties:
	SystemRecommendedExposureBiasRange() float32
	SetSystemRecommendedExposureBiasRange(value float32)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)


	

	// methods:


}





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






// Creates a slider to control the exposure bias of the specified capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:)
func NewCaptureSystemExposureBiasSliderWithDevice(device IAVCaptureDevice) CaptureSystemExposureBiasSlider {
	instance := getCaptureSystemExposureBiasSliderClass().Alloc()
	rv := objc.Send[CaptureSystemExposureBiasSlider](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Creates a slider to control the exposure bias of the specified capture device with an action to respond to exposure bias changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemExposureBiasSlider/init(device:action:)
func NewCaptureSystemExposureBiasSliderWithDeviceAction(device IAVCaptureDevice, action func(float32)) CaptureSystemExposureBiasSlider {
	instance := getCaptureSystemExposureBiasSliderClass().Alloc()
	rv := objc.Send[CaptureSystemExposureBiasSlider](instance.ID, objc.Sel("initWithDevice:action:"), device, action)
	rv.Autorelease()
	return rv
}






















// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureSystemExposureBiasSlider) SystemRecommendedExposureBiasRange() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return rv
}


// The system’s recommended exposure bias range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedexposurebiasrange
func (c_ CaptureSystemExposureBiasSlider) SetSystemRecommendedExposureBiasRange(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemRecommendedExposureBiasRange:"), value)
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemExposureBiasSlider) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemExposureBiasSlider) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}







