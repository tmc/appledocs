// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionConversionTimeRangeAdjustment */


/* debug [class_header]: Header for AVCaptionConversionTimeRangeAdjustment */
// The class instance for the [CaptionConversionTimeRangeAdjustment] class.
var (
	CaptionConversionTimeRangeAdjustmentClass     _CaptionConversionTimeRangeAdjustmentClass
	CaptionConversionTimeRangeAdjustmentClassOnce sync.Once
)

func getCaptionConversionTimeRangeAdjustmentClass() _CaptionConversionTimeRangeAdjustmentClass {
	CaptionConversionTimeRangeAdjustmentClassOnce.Do(func() {
		CaptionConversionTimeRangeAdjustmentClass = _CaptionConversionTimeRangeAdjustmentClass{objc.GetClass("AVCaptionConversionTimeRangeAdjustment")}
	})
	return CaptionConversionTimeRangeAdjustmentClass
}

type _CaptionConversionTimeRangeAdjustmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionConversionTimeRangeAdjustment */
// An interface definition for the [CaptionConversionTimeRangeAdjustment] class.
type ICaptionConversionTimeRangeAdjustment interface {
	ICaptionConversionAdjustment
	
/* debug [class_interface_properties]: Properties for CaptionConversionTimeRangeAdjustment */
	// properties:
	DurationOffset() objc.IObject /* cross-framework: Time */
	StartTimeOffset() objc.IObject /* cross-framework: Time */
	AdjustmentType() objectivec.IObject
	SetAdjustmentType(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionConversionTimeRangeAdjustment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionConversionTimeRangeAdjustment */
// Alloc allocates a new instance without initialization.
func (cc _CaptionConversionTimeRangeAdjustmentClass) Alloc() CaptionConversionTimeRangeAdjustment {
	rv := objc.Send[CaptionConversionTimeRangeAdjustment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionConversionTimeRangeAdjustmentClass) New() CaptionConversionTimeRangeAdjustment {
	rv := objc.Send[CaptionConversionTimeRangeAdjustment](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionConversionTimeRangeAdjustment) Init() CaptionConversionTimeRangeAdjustment {
	rv := objc.Send[CaptionConversionTimeRangeAdjustment](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionConversionTimeRangeAdjustment) Autorelease() CaptionConversionTimeRangeAdjustment {
	rv := objc.Send[CaptionConversionTimeRangeAdjustment](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionConversionTimeRangeAdjustment creates a new CaptionConversionTimeRangeAdjustment instance.
func NewCaptionConversionTimeRangeAdjustment() CaptionConversionTimeRangeAdjustment {
	return getCaptionConversionTimeRangeAdjustmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionConversionTimeRangeAdjustment */
// An object that describes an adjustment to the time range of one or more captions.


// An object that describes an adjustment to the time range of one or more captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment
type CaptionConversionTimeRangeAdjustment struct {
	CaptionConversionAdjustment
}

// CaptionConversionTimeRangeAdjustmentFrom constructs a [CaptionConversionTimeRangeAdjustment] from an unsafe.Pointer.
//
// An object that describes an adjustment to the time range of one or more captions.
func CaptionConversionTimeRangeAdjustmentFrom(ptr unsafe.Pointer) CaptionConversionTimeRangeAdjustment {
	return CaptionConversionTimeRangeAdjustment{
		CaptionConversionAdjustment: CaptionConversionAdjustmentFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionConversionTimeRangeAdjustment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionConversionTimeRangeAdjustment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionConversionTimeRangeAdjustment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionConversionTimeRangeAdjustment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionConversionTimeRangeAdjustment */

// The time value by which the system offsets the durations of captions to correct a problem.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/durationOffset
func (c_ CaptionConversionTimeRangeAdjustment) DurationOffset() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("durationOffset"))
	return rv
}/* debug [instance_properties/getter]: durationOffset */


// The time value by which the system offsets the start times of captions to correct a problem.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/startTimeOffset
func (c_ CaptionConversionTimeRangeAdjustment) StartTimeOffset() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("startTimeOffset"))
	return rv
}/* debug [instance_properties/getter]: startTimeOffset */


// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment/adjustmenttype-swift.property
func (c_ CaptionConversionTimeRangeAdjustment) AdjustmentType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("adjustmentType"))
	return rv
}/* debug [instance_properties/getter]: adjustmentType */


// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment/adjustmenttype-swift.property
func (c_ CaptionConversionTimeRangeAdjustment) SetAdjustmentType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdjustmentType:"), value)
}/* debug [instance_properties/setter]: adjustmentType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionConversionTimeRangeAdjustment */



