// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CaptionConversionTimeRangeAdjustment] class.
type ICaptionConversionTimeRangeAdjustment interface {
	ICaptionConversionAdjustment
	

	// properties:
	DurationOffset() objectivec.IObject
	StartTimeOffset() objectivec.IObject
	AdjustmentType() objectivec.IObject
	SetAdjustmentType(value objectivec.IObject)


	

	// methods:


}





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

























// The time value by which the system offsets the durations of captions to correct a problem.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/durationOffset
func (c_ CaptionConversionTimeRangeAdjustment) DurationOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("durationOffset"))
	return rv
}


// The time value by which the system offsets the start times of captions to correct a problem.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/startTimeOffset
func (c_ CaptionConversionTimeRangeAdjustment) StartTimeOffset() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("startTimeOffset"))
	return rv
}


// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment/adjustmenttype-swift.property
func (c_ CaptionConversionTimeRangeAdjustment) AdjustmentType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("adjustmentType"))
	return rv
}


// The type of caption conversion adjustment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptionconversionadjustment/adjustmenttype-swift.property
func (c_ CaptionConversionTimeRangeAdjustment) SetAdjustmentType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdjustmentType:"), value)
}








