// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureDepthDataOutput */


/* debug [class_header]: Header for AVCaptureDepthDataOutput */
// The class instance for the [CaptureDepthDataOutput] class.
var (
	CaptureDepthDataOutputClass     _CaptureDepthDataOutputClass
	CaptureDepthDataOutputClassOnce sync.Once
)

func getCaptureDepthDataOutputClass() _CaptureDepthDataOutputClass {
	CaptureDepthDataOutputClassOnce.Do(func() {
		CaptureDepthDataOutputClass = _CaptureDepthDataOutputClass{objc.GetClass("AVCaptureDepthDataOutput")}
	})
	return CaptureDepthDataOutputClass
}

type _CaptureDepthDataOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDepthDataOutput */
// An interface definition for the [CaptureDepthDataOutput] class.
type ICaptureDepthDataOutput interface {
	ICaptureOutput
	
/* debug [class_interface_properties]: Properties for CaptureDepthDataOutput */
	// properties:
	IsFilteringEnabled() bool
	SetIsFilteringEnabled(value bool)
	SupportedDepthDataFormats() IAVCaptureDeviceFormat
	SetSupportedDepthDataFormats(value IAVCaptureDeviceFormat)
	ActiveDepthDataFormat() IAVCaptureDeviceFormat
	SetActiveDepthDataFormat(value IAVCaptureDeviceFormat)
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	IsDepthDataDeliveryEnabled() bool
	SetIsDepthDataDeliveryEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDepthDataOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDepthDataOutput */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDepthDataOutputClass) Alloc() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDepthDataOutputClass) New() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDepthDataOutput) Init() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDepthDataOutput) Autorelease() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDepthDataOutput creates a new CaptureDepthDataOutput instance.
func NewCaptureDepthDataOutput() CaptureDepthDataOutput {
	return getCaptureDepthDataOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDepthDataOutput */
// A capture output that records scene depth information on compatible camera devices.
//
// This output type captures objects containing per-pixel depth or disparity information, following a streaming delivery model similar to that used by . Alternatively, you can capture depth data alongside photos using (see the property). This object always provides depth data in the format expressed by the source object’s property. If you wish to receive depth data in another format, choose a new value for that property from those listed in the array of the device’s object.


// A capture output that records scene depth information on compatible camera devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
type CaptureDepthDataOutput struct {
	CaptureOutput
}

// CaptureDepthDataOutputFrom constructs a [CaptureDepthDataOutput] from an unsafe.Pointer.
//
// A capture output that records scene depth information on compatible camera devices.
func CaptureDepthDataOutputFrom(ptr unsafe.Pointer) CaptureDepthDataOutput {
	return CaptureDepthDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDepthDataOutput */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDepthDataOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDepthDataOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDepthDataOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDepthDataOutput */

// A Boolean value that determines whether the depth data output should filter depth data to smooth out noise and fill invalid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/isfilteringenabled
func (c_ CaptureDepthDataOutput) IsFilteringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFilteringEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFilteringEnabled */


// A Boolean value that determines whether the depth data output should filter depth data to smooth out noise and fill invalid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/isfilteringenabled
func (c_ CaptureDepthDataOutput) SetIsFilteringEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFilteringEnabled:"), value)
}/* debug [instance_properties/setter]: isFilteringEnabled */


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDepthDataOutput) SupportedDepthDataFormats() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("supportedDepthDataFormats"))
	return rv
}/* debug [instance_properties/getter]: supportedDepthDataFormats */


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDepthDataOutput) SetSupportedDepthDataFormats(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedDepthDataFormats:"), value)
}/* debug [instance_properties/setter]: supportedDepthDataFormats */


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDepthDataOutput) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeDepthDataFormat"))
	return rv
}/* debug [instance_properties/getter]: activeDepthDataFormat */


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDepthDataOutput) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveDepthDataFormat:"), value)
}/* debug [instance_properties/setter]: activeDepthDataFormat */


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDepthDataOutput) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}/* debug [instance_properties/getter]: activeFormat */


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDepthDataOutput) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}/* debug [instance_properties/setter]: activeFormat */


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CaptureDepthDataOutput) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDepthDataDeliveryEnabled */


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CaptureDepthDataOutput) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}/* debug [instance_properties/setter]: isDepthDataDeliveryEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureDepthDataOutput */


