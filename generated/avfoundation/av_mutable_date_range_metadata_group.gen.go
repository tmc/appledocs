// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMutableDateRangeMetadataGroup */


/* debug [class_header]: Header for AVMutableDateRangeMetadataGroup */
// The class instance for the [MutableDateRangeMetadataGroup] class.
var (
	MutableDateRangeMetadataGroupClass     _MutableDateRangeMetadataGroupClass
	MutableDateRangeMetadataGroupClassOnce sync.Once
)

func getMutableDateRangeMetadataGroupClass() _MutableDateRangeMetadataGroupClass {
	MutableDateRangeMetadataGroupClassOnce.Do(func() {
		MutableDateRangeMetadataGroupClass = _MutableDateRangeMetadataGroupClass{objc.GetClass("AVMutableDateRangeMetadataGroup")}
	})
	return MutableDateRangeMetadataGroupClass
}

type _MutableDateRangeMetadataGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableDateRangeMetadataGroup */
// An interface definition for the [MutableDateRangeMetadataGroup] class.
type IMutableDateRangeMetadataGroup interface {
	IDateRangeMetadataGroup
	
/* debug [class_interface_properties]: Properties for MutableDateRangeMetadataGroup */
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	SetEndDate(value objc.IObject /* cross-framework: NSDate */)
	Items() []MetadataItem
	SetItems(value []MetadataItem)
	StartDate() objc.IObject /* cross-framework: NSDate */
	SetStartDate(value objc.IObject /* cross-framework: NSDate */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableDateRangeMetadataGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableDateRangeMetadataGroup */
// Alloc allocates a new instance without initialization.
func (mc _MutableDateRangeMetadataGroupClass) Alloc() MutableDateRangeMetadataGroup {
	rv := objc.Send[MutableDateRangeMetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableDateRangeMetadataGroupClass) New() MutableDateRangeMetadataGroup {
	rv := objc.Send[MutableDateRangeMetadataGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableDateRangeMetadataGroup) Init() MutableDateRangeMetadataGroup {
	rv := objc.Send[MutableDateRangeMetadataGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableDateRangeMetadataGroup) Autorelease() MutableDateRangeMetadataGroup {
	rv := objc.Send[MutableDateRangeMetadataGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableDateRangeMetadataGroup creates a new MutableDateRangeMetadataGroup instance.
func NewMutableDateRangeMetadataGroup() MutableDateRangeMetadataGroup {
	return getMutableDateRangeMetadataGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableDateRangeMetadataGroup */
// A mutable collection of metadata items that are valid for use within a specific range of dates.


// A mutable collection of metadata items that are valid for use within a specific range of dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup
type MutableDateRangeMetadataGroup struct {
	DateRangeMetadataGroup
}

// MutableDateRangeMetadataGroupFrom constructs a [MutableDateRangeMetadataGroup] from an unsafe.Pointer.
//
// A mutable collection of metadata items that are valid for use within a specific range of dates.
func MutableDateRangeMetadataGroupFrom(ptr unsafe.Pointer) MutableDateRangeMetadataGroup {
	return MutableDateRangeMetadataGroup{
		DateRangeMetadataGroup: DateRangeMetadataGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableDateRangeMetadataGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableDateRangeMetadataGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableDateRangeMetadataGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableDateRangeMetadataGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableDateRangeMetadataGroup */

// The end date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/endDate
func (m_ MutableDateRangeMetadataGroup) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The end date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/endDate
func (m_ MutableDateRangeMetadataGroup) SetEndDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndDate:"), value)
}/* debug [instance_properties/setter]: endDate */


// An array of associated metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/items
func (m_ MutableDateRangeMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// An array of associated metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/items
func (m_ MutableDateRangeMetadataGroup) SetItems(value []MetadataItem) {
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


// The start date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/startDate
func (m_ MutableDateRangeMetadataGroup) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The start date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/startDate
func (m_ MutableDateRangeMetadataGroup) SetStartDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartDate:"), value)
}/* debug [instance_properties/setter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableDateRangeMetadataGroup */



