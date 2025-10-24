// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionConversionWarning] class.
var (
	CaptionConversionWarningClass     _CaptionConversionWarningClass
	CaptionConversionWarningClassOnce sync.Once
)

func getCaptionConversionWarningClass() _CaptionConversionWarningClass {
	CaptionConversionWarningClassOnce.Do(func() {
		CaptionConversionWarningClass = _CaptionConversionWarningClass{objc.GetClass("AVCaptionConversionWarning")}
	})
	return CaptionConversionWarningClass
}

type _CaptionConversionWarningClass struct {
	class objc.Class
}





// An interface definition for the [CaptionConversionWarning] class.
type ICaptionConversionWarning interface {
	objectivec.IObject
	

	// properties:
	Adjustment() IAVCaptionConversionAdjustment
	RangeOfCaptions() corefoundation.Range
	WarningType() CaptionConversionWarningType /* typedef */
	Warnings() IAVCaptionConversionWarning
	SetWarnings(value IAVCaptionConversionWarning)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionConversionWarningClass) Alloc() CaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionConversionWarningClass) New() CaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionConversionWarning) Init() CaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionConversionWarning) Autorelease() CaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionConversionWarning creates a new CaptionConversionWarning instance.
func NewCaptionConversionWarning() CaptionConversionWarning {
	return getCaptionConversionWarningClass().New()
}





// An object that represents a conversion warning produced by a validator.


// An object that represents a conversion warning produced by a validator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning
type CaptionConversionWarning struct {
	objectivec.Object
}

// CaptionConversionWarningFrom constructs a [CaptionConversionWarning] from an unsafe.Pointer.
//
// An object that represents a conversion warning produced by a validator.
func CaptionConversionWarningFrom(ptr unsafe.Pointer) CaptionConversionWarning {
	return CaptionConversionWarning{objectivec.Object{objc.ID(ptr)}}
}

























// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/adjustment
func (c_ CaptionConversionWarning) Adjustment() IAVCaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("adjustment"))
	return rv
}


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/rangeOfCaptions
func (c_ CaptionConversionWarning) RangeOfCaptions() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](c_.ID, objc.Sel("rangeOfCaptions"))
	return rv
}


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/warningType-swift.property
func (c_ CaptionConversionWarning) WarningType() CaptionConversionWarningType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("warningType"))
	return rv
}


// The collection of warnings the validator encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/warnings
func (c_ CaptionConversionWarning) Warnings() IAVCaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](c_.ID, objc.Sel("warnings"))
	return rv
}


// The collection of warnings the validator encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/warnings
func (c_ CaptionConversionWarning) SetWarnings(value IAVCaptionConversionWarning) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWarnings:"), value)
}








