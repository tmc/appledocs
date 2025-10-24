// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionGroup] class.
var (
	CaptionGroupClass     _CaptionGroupClass
	CaptionGroupClassOnce sync.Once
)

func getCaptionGroupClass() _CaptionGroupClass {
	CaptionGroupClassOnce.Do(func() {
		CaptionGroupClass = _CaptionGroupClass{objc.GetClass("AVCaptionGroup")}
	})
	return CaptionGroupClass
}

type _CaptionGroupClass struct {
	class objc.Class
}





// An interface definition for the [CaptionGroup] class.
type ICaptionGroup interface {
	objectivec.IObject
	

	// properties:
	Captions() []Caption
	TimeRange() TimeRange /* not a class type */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionGroupClass) Alloc() CaptionGroup {
	rv := objc.Send[CaptionGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionGroupClass) New() CaptionGroup {
	rv := objc.Send[CaptionGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionGroup) Init() CaptionGroup {
	rv := objc.Send[CaptionGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionGroup) Autorelease() CaptionGroup {
	rv := objc.Send[CaptionGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionGroup creates a new CaptionGroup instance.
func NewCaptionGroup() CaptionGroup {
	return getCaptionGroupClass().New()
}





// An object that represents zero or more captions that intersect in time.


// An object that represents zero or more captions that intersect in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup
type CaptionGroup struct {
	objectivec.Object
}

// CaptionGroupFrom constructs a [CaptionGroup] from an unsafe.Pointer.
//
// An object that represents zero or more captions that intersect in time.
func CaptionGroupFrom(ptr unsafe.Pointer) CaptionGroup {
	return CaptionGroup{objectivec.Object{objc.ID(ptr)}}
}






// Creates a caption group with captions and a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(captions:timeRange:)
func NewCaptionGroupWithCaptionsTimeRange(captions []Caption, timeRange TimeRange /* not a class type */) CaptionGroup {
	instance := getCaptionGroupClass().Alloc()
	rv := objc.Send[CaptionGroup](instance.ID, objc.Sel("initWithCaptions:timeRange:"), captions, timeRange)
	rv.Autorelease()
	return rv
}


// Creates a caption group with a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(timeRange:)
func NewCaptionGroupWithTimeRange(timeRange TimeRange /* not a class type */) CaptionGroup {
	instance := getCaptionGroupClass().Alloc()
	rv := objc.Send[CaptionGroup](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}






















// The captions associated with the caption group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/captions
func (c_ CaptionGroup) Captions() []Caption {
	rv := objc.Send[[]Caption](c_.ID, objc.Sel("captions"))
	return rv
}


// The time range of the caption group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/timeRange
func (c_ CaptionGroup) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}







