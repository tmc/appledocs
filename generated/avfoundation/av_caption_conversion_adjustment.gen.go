// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionConversionAdjustment */


/* debug [class_header]: Header for AVCaptionConversionAdjustment */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionConversionAdjustment */
// An interface definition for the [CaptionConversionAdjustment] class.
type ICaptionConversionAdjustment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionConversionAdjustment */
	// properties:
	AdjustmentType() CaptionConversionAdjustmentType /* typedef */
	Adjustment() IAVCaptionConversionAdjustment
	SetAdjustment(value IAVCaptionConversionAdjustment)
	RangeOfCaptions() corefoundation.Range
	SetRangeOfCaptions(value corefoundation.Range)
	WarningType() objectivec.IObject
	SetWarningType(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionConversionAdjustment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionConversionAdjustment */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionConversionAdjustment */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionConversionAdjustment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionConversionAdjustment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionConversionAdjustment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionConversionAdjustment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionConversionAdjustment */

// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment/adjustmentType-swift.property
func (c_ CaptionConversionAdjustment) AdjustmentType() CaptionConversionAdjustmentType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("adjustmentType"))
	return rv
}/* debug [instance_properties/getter]: adjustmentType */


// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/adjustment
func (c_ CaptionConversionAdjustment) Adjustment() IAVCaptionConversionAdjustment {
	rv := objc.Send[CaptionConversionAdjustment](c_.ID, objc.Sel("adjustment"))
	return rv
}/* debug [instance_properties/getter]: adjustment */


// A correction the converter makes when it converts a caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/adjustment
func (c_ CaptionConversionAdjustment) SetAdjustment(value IAVCaptionConversionAdjustment) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdjustment:"), value)
}/* debug [instance_properties/setter]: adjustment */


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/rangeofcaptions
func (c_ CaptionConversionAdjustment) RangeOfCaptions() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](c_.ID, objc.Sel("rangeOfCaptions"))
	return rv
}/* debug [instance_properties/getter]: rangeOfCaptions */


// The range of the captions for which the system issued a warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/rangeofcaptions
func (c_ CaptionConversionAdjustment) SetRangeOfCaptions(value corefoundation.Range) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRangeOfCaptions:"), value)
}/* debug [instance_properties/setter]: rangeOfCaptions */


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/warningtype-swift.property
func (c_ CaptionConversionAdjustment) WarningType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("warningType"))
	return rv
}/* debug [instance_properties/getter]: warningType */


// A type that indicates the nature of the validation warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionwarning/warningtype-swift.property
func (c_ CaptionConversionAdjustment) SetWarningType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWarningType:"), value)
}/* debug [instance_properties/setter]: warningType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionConversionAdjustment */



