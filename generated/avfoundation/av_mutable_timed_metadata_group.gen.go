// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMutableTimedMetadataGroup */


/* debug [class_header]: Header for AVMutableTimedMetadataGroup */
// The class instance for the [MutableTimedMetadataGroup] class.
var (
	MutableTimedMetadataGroupClass     _MutableTimedMetadataGroupClass
	MutableTimedMetadataGroupClassOnce sync.Once
)

func getMutableTimedMetadataGroupClass() _MutableTimedMetadataGroupClass {
	MutableTimedMetadataGroupClassOnce.Do(func() {
		MutableTimedMetadataGroupClass = _MutableTimedMetadataGroupClass{objc.GetClass("AVMutableTimedMetadataGroup")}
	})
	return MutableTimedMetadataGroupClass
}

type _MutableTimedMetadataGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableTimedMetadataGroup */
// An interface definition for the [MutableTimedMetadataGroup] class.
type IMutableTimedMetadataGroup interface {
	ITimedMetadataGroup
	
/* debug [class_interface_properties]: Properties for MutableTimedMetadataGroup */
	// properties:
	Items() []MetadataItem
	SetItems(value []MetadataItem)
	TimeRange() TimeRange /* not a class type */
	SetTimeRange(value TimeRange /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableTimedMetadataGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableTimedMetadataGroup */
// Alloc allocates a new instance without initialization.
func (mc _MutableTimedMetadataGroupClass) Alloc() MutableTimedMetadataGroup {
	rv := objc.Send[MutableTimedMetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableTimedMetadataGroupClass) New() MutableTimedMetadataGroup {
	rv := objc.Send[MutableTimedMetadataGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableTimedMetadataGroup) Init() MutableTimedMetadataGroup {
	rv := objc.Send[MutableTimedMetadataGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableTimedMetadataGroup) Autorelease() MutableTimedMetadataGroup {
	rv := objc.Send[MutableTimedMetadataGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableTimedMetadataGroup creates a new MutableTimedMetadataGroup instance.
func NewMutableTimedMetadataGroup() MutableTimedMetadataGroup {
	return getMutableTimedMetadataGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableTimedMetadataGroup */
// A mutable collection of metadata items that are valid for use during a specific time range.


// A mutable collection of metadata items that are valid for use during a specific time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup
type MutableTimedMetadataGroup struct {
	TimedMetadataGroup
}

// MutableTimedMetadataGroupFrom constructs a [MutableTimedMetadataGroup] from an unsafe.Pointer.
//
// A mutable collection of metadata items that are valid for use during a specific time range.
func MutableTimedMetadataGroupFrom(ptr unsafe.Pointer) MutableTimedMetadataGroup {
	return MutableTimedMetadataGroup{
		TimedMetadataGroup: TimedMetadataGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableTimedMetadataGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableTimedMetadataGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableTimedMetadataGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableTimedMetadataGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableTimedMetadataGroup */

// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/items
func (m_ MutableTimedMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/items
func (m_ MutableTimedMetadataGroup) SetItems(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setItems:"), nsArray)
}/* debug [instance_properties/setter]: items */


// The time range of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/timeRange
func (m_ MutableTimedMetadataGroup) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](m_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// The time range of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/timeRange
func (m_ MutableTimedMetadataGroup) SetTimeRange(value TimeRange /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}/* debug [instance_properties/setter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableTimedMetadataGroup */



