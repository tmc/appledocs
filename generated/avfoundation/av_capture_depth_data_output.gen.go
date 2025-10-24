// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CaptureDepthDataOutput] class.
type ICaptureDepthDataOutput interface {
	ICaptureOutput
	// properties:
	AlwaysDiscardsLateDepthData() bool
	SetAlwaysDiscardsLateDepthData(value bool)
	Delegate() CaptureDepthDataOutputDelegate /* not a class type */
	SetDelegate(value CaptureDepthDataOutputDelegate /* not a class type */)
	DelegateCallbackQueue() unsafe.Pointer
	SetDelegateCallbackQueue(value unsafe.Pointer)
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
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CaptureDepthDataOutputClass) Alloc() CaptureDepthDataOutput {
	rv := objc.Send[CaptureDepthDataOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that determines whether the capture output should discard any depth data that is not processed before the next depth data is captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/alwaysdiscardslatedepthdata
func (c_ CaptureDepthDataOutput) AlwaysDiscardsLateDepthData() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("alwaysDiscardsLateDepthData"))
	return rv
}


// A Boolean value that determines whether the capture output should discard any depth data that is not processed before the next depth data is captured.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/alwaysdiscardslatedepthdata
func (c_ CaptureDepthDataOutput) SetAlwaysDiscardsLateDepthData(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlwaysDiscardsLateDepthData:"), value)
}


// A delegate object that receives depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/delegate
func (c_ CaptureDepthDataOutput) Delegate() CaptureDepthDataOutputDelegate /* not a class type */ {
	rv := objc.Send[CaptureDepthDataOutputDelegate](c_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object that receives depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/delegate
func (c_ CaptureDepthDataOutput) SetDelegate(value CaptureDepthDataOutputDelegate /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// A dispatch queue for delivering depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/delegatecallbackqueue
func (c_ CaptureDepthDataOutput) DelegateCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}


// A dispatch queue for delivering depth data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/delegatecallbackqueue
func (c_ CaptureDepthDataOutput) SetDelegateCallbackQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegateCallbackQueue:"), value)
}


// A Boolean value that determines whether the depth data output should filter depth data to smooth out noise and fill invalid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/isfilteringenabled
func (c_ CaptureDepthDataOutput) IsFilteringEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFilteringEnabled"))
	return rv
}


// A Boolean value that determines whether the depth data output should filter depth data to smooth out noise and fill invalid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedepthdataoutput/isfilteringenabled
func (c_ CaptureDepthDataOutput) SetIsFilteringEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFilteringEnabled:"), value)
}


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDepthDataOutput) SupportedDepthDataFormats() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("supportedDepthDataFormats"))
	return rv
}


// The list of data formats compatible with this video format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/supporteddepthdataformats
func (c_ CaptureDepthDataOutput) SetSupportedDepthDataFormats(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportedDepthDataFormats:"), value)
}


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDepthDataOutput) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeDepthDataFormat"))
	return rv
}


// The currently active depth data format of the capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c_ CaptureDepthDataOutput) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveDepthDataFormat:"), value)
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDepthDataOutput) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[CaptureDeviceFormat](c_.ID, objc.Sel("activeFormat"))
	return rv
}


// The capture format in use by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c_ CaptureDepthDataOutput) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActiveFormat:"), value)
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CaptureDepthDataOutput) IsDepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}


// A Boolean value that determines whether the photo output captures depth data along with the photo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/isdepthdatadeliveryenabled
func (c_ CaptureDepthDataOutput) SetIsDepthDataDeliveryEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDepthDataDeliveryEnabled:"), value)
}



