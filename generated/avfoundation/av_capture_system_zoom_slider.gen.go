// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureSystemZoomSlider] class.
var (
	CaptureSystemZoomSliderClass     _CaptureSystemZoomSliderClass
	CaptureSystemZoomSliderClassOnce sync.Once
)

func getCaptureSystemZoomSliderClass() _CaptureSystemZoomSliderClass {
	CaptureSystemZoomSliderClassOnce.Do(func() {
		CaptureSystemZoomSliderClass = _CaptureSystemZoomSliderClass{objc.GetClass("AVCaptureSystemZoomSlider")}
	})
	return CaptureSystemZoomSliderClass
}

type _CaptureSystemZoomSliderClass struct {
	class objc.Class
}





// An interface definition for the [CaptureSystemZoomSlider] class.
type ICaptureSystemZoomSlider interface {
	ICaptureControl
	

	// properties:
	SystemRecommendedVideoZoomRange() float64
	SetSystemRecommendedVideoZoomRange(value float64)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureSystemZoomSliderClass) Alloc() CaptureSystemZoomSlider {
	rv := objc.Send[CaptureSystemZoomSlider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureSystemZoomSliderClass) New() CaptureSystemZoomSlider {
	rv := objc.Send[CaptureSystemZoomSlider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSystemZoomSlider) Init() CaptureSystemZoomSlider {
	rv := objc.Send[CaptureSystemZoomSlider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSystemZoomSlider) Autorelease() CaptureSystemZoomSlider {
	rv := objc.Send[CaptureSystemZoomSlider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSystemZoomSlider creates a new CaptureSystemZoomSlider instance.
func NewCaptureSystemZoomSlider() CaptureSystemZoomSlider {
	return getCaptureSystemZoomSliderClass().New()
}





// A control that adjusts the video zoom factor of a capture device within the system-recommended range.
//
// The system sets the slider’s range to the value of the property of the device’s active format. If a device’s value changes, the slider updates its range to the new format’s recommendation. To use this control, add it to the capture session by calling the session’s method.


// A control that adjusts the video zoom factor of a capture device within the system-recommended range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider
type CaptureSystemZoomSlider struct {
	CaptureControl
}

// CaptureSystemZoomSliderFrom constructs a [CaptureSystemZoomSlider] from an unsafe.Pointer.
//
// A control that adjusts the video zoom factor of a capture device within the system-recommended range.
func CaptureSystemZoomSliderFrom(ptr unsafe.Pointer) CaptureSystemZoomSlider {
	return CaptureSystemZoomSlider{
		CaptureControl: CaptureControlFrom(ptr),
	}
}






// Creates a slider to control the video zoom factor of a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider/init(device:)
func NewCaptureSystemZoomSliderWithDevice(device IAVCaptureDevice) CaptureSystemZoomSlider {
	instance := getCaptureSystemZoomSliderClass().Alloc()
	rv := objc.Send[CaptureSystemZoomSlider](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Creates a slider to control the zoom level of the specified capture device with an action to respond to zoom changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSystemZoomSlider/init(device:action:)
func NewCaptureSystemZoomSliderWithDeviceAction(device IAVCaptureDevice, action func(float64)) CaptureSystemZoomSlider {
	instance := getCaptureSystemZoomSliderClass().Alloc()
	rv := objc.Send[CaptureSystemZoomSlider](instance.ID, objc.Sel("initWithDevice:action:"), device, action)
	rv.Autorelease()
	return rv
}






















// The system’s recommended zoom range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedvideozoomrange
func (c_ CaptureSystemZoomSlider) SystemRecommendedVideoZoomRange() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("systemRecommendedVideoZoomRange"))
	return rv
}


// The system’s recommended zoom range for this device format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/systemrecommendedvideozoomrange
func (c_ CaptureSystemZoomSlider) SetSystemRecommendedVideoZoomRange(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSystemRecommendedVideoZoomRange:"), value)
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemZoomSlider) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureSystemZoomSlider) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}







