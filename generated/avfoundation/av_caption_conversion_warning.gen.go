// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionConversionWarning */


/* debug [class_header]: Header for AVCaptionConversionWarning */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionConversionWarning */
// An interface definition for the [CaptionConversionWarning] class.
type ICaptionConversionWarning interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionConversionWarning */
	// properties:
	Adjustment() IAVCaptionConversionAdjustment
	RangeOfCaptions() corefoundation.Range
	WarningType() CaptionConversionWarningType /* typedef */
	Warnings() IAVCaptionConversionWarning
	SetWarnings(value IAVCaptionConversionWarning)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionConversionWarning */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionConversionWarning */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionConversionWarning */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionConversionWarning *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionConversionWarning */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionConversionWarning */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionConversionWarning */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionConversionWarning */

// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/adjustment
func (c_ CaptionConversionWarning) Adjustment() IAVCaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("adjustment"))
	return rv
}/* debug [instance_properties/getter]: adjustment */


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/rangeOfCaptions
func (c_ CaptionConversionWarning) RangeOfCaptions() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](c_.ID, objc.Sel("rangeOfCaptions"))
	return rv
}/* debug [instance_properties/getter]: rangeOfCaptions */


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/warningType-swift.property
func (c_ CaptionConversionWarning) WarningType() CaptionConversionWarningType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("warningType"))
	return rv
}/* debug [instance_properties/getter]: warningType */


// The collection of warnings the validator encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/warnings
func (c_ CaptionConversionWarning) Warnings() IAVCaptionConversionWarning {
	rv := objc.Send[CaptionConversionWarning](c_.ID, objc.Sel("warnings"))
	return rv
}/* debug [instance_properties/getter]: warnings */


// The collection of warnings the validator encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionvalidator/warnings
func (c_ CaptionConversionWarning) SetWarnings(value IAVCaptionConversionWarning) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWarnings:"), value)
}/* debug [instance_properties/setter]: warnings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionConversionWarning */



