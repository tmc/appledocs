// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionConversionAdjustment] class.
var (
	CaptionConversionAdjustmentClass     _CaptionConversionAdjustmentClass
	CaptionConversionAdjustmentClassOnce sync.Once
)

func getCaptionConversionAdjustmentClass() _CaptionConversionAdjustmentClass {
	CaptionConversionAdjustmentClassOnce.Do(func() {
		CaptionConversionAdjustmentClass = _CaptionConversionAdjustmentClass{objc.GetClass("AVCaptionConversionAdjustment")}
	})
	return CaptionConversionAdjustmentClass
}

type _CaptionConversionAdjustmentClass struct {
	class objc.Class
}





// An interface definition for the [CaptionConversionAdjustment] class.
type ICaptionConversionAdjustment interface {
	objectivec.IObject
	

	// properties:
	AdjustmentType() CaptionConversionAdjustmentType
	Adjustment() IAVCaptionConversionAdjustment
	SetAdjustment(value IAVCaptionConversionAdjustment)
	RangeOfCaptions() foundation.Range
	SetRangeOfCaptions(value foundation.Range)
	WarningType() objectivec.IObject
	SetWarningType(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionConversionAdjustmentClass) Alloc() CaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionConversionAdjustmentClass) New() CaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionConversionAdjustment) Init() CaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionConversionAdjustment) Autorelease() CaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionConversionAdjustment creates a new CaptionConversionAdjustment instance.
func NewCaptionConversionAdjustment() CaptionConversionAdjustment {
	return getCaptionConversionAdjustmentClass().New()
}





// An object that describes an adjustment to correct a problem found during validation of a caption conversion.


// An object that describes an adjustment to correct a problem found during validation of a caption conversion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment
type CaptionConversionAdjustment struct {
	objectivec.Object
}

// CaptionConversionAdjustmentFrom constructs a [CaptionConversionAdjustment] from an unsafe.Pointer.
//
// An object that describes an adjustment to correct a problem found during validation of a caption conversion.
func CaptionConversionAdjustmentFrom(ptr unsafe.Pointer) CaptionConversionAdjustment {
	return CaptionConversionAdjustment{objectivec.Object{objc.ID(ptr)}}
}

























// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment/adjustmentType-swift.property
func (c_ CaptionConversionAdjustment) AdjustmentType() CaptionConversionAdjustmentType {
	rv := objc.Send[CaptionConversionAdjustmentType](c_.ID, objc.Sel("adjustmentType"))
	return rv
}


// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/adjustment
func (c_ CaptionConversionAdjustment) Adjustment() IAVCaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("adjustment"))
	return rv
}


// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/adjustment
func (c_ CaptionConversionAdjustment) SetAdjustment(value IAVCaptionConversionAdjustment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdjustment:"), value)
}


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/rangeofcaptions
func (c_ CaptionConversionAdjustment) RangeOfCaptions() foundation.Range {
	rv := objc.Send[foundation.Range](c_.ID, objc.Sel("rangeOfCaptions"))
	return rv
}


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/rangeofcaptions
func (c_ CaptionConversionAdjustment) SetRangeOfCaptions(value foundation.Range) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRangeOfCaptions:"), value)
}


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/warningtype-swift.property
func (c_ CaptionConversionAdjustment) WarningType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("warningType"))
	return rv
}


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/warningtype-swift.property
func (c_ CaptionConversionAdjustment) SetWarningType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWarningType:"), value)
}








