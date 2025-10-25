// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionGroup */


/* debug [class_header]: Header for AVCaptionGroup */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionGroup */
// An interface definition for the [CaptionGroup] class.
type ICaptionGroup interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionGroup */
	// properties:
	Captions() []Caption
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionGroup */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionGroup */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionGroup */

// Creates a caption group with captions and a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(captions:timeRange:)
func NewCaptionGroupWithCaptionsTimeRange(captions []Caption, timeRange TimeRange /* not a class type */) CaptionGroup {
	instance := getCaptionGroupClass().Alloc()
	rv := objc.Send[CaptionGroup](instance.ID, objc.Sel("initWithCaptions:timeRange:"), captions, timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionGroupWithCaptionsTimeRange */


// Creates a caption group with a time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/init(timeRange:)
func NewCaptionGroupWithTimeRange(timeRange TimeRange /* not a class type */) CaptionGroup {
	instance := getCaptionGroupClass().Alloc()
	rv := objc.Send[CaptionGroup](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionGroupWithTimeRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionGroup */

// The captions associated with the caption group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/captions
func (c_ CaptionGroup) Captions() []Caption {
	rv := objc.Send[[]Caption](c_.ID, objc.Sel("captions"))
	return rv
}/* debug [instance_properties/getter]: captions */


// The time range of the caption group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionGroup/timeRange
func (c_ CaptionGroup) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionGroup */


