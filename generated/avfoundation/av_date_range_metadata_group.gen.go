// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDateRangeMetadataGroup */


/* debug [class_header]: Header for AVDateRangeMetadataGroup */
// The class instance for the [DateRangeMetadataGroup] class.
var (
	DateRangeMetadataGroupClass     _DateRangeMetadataGroupClass
	DateRangeMetadataGroupClassOnce sync.Once
)

func getDateRangeMetadataGroupClass() _DateRangeMetadataGroupClass {
	DateRangeMetadataGroupClassOnce.Do(func() {
		DateRangeMetadataGroupClass = _DateRangeMetadataGroupClass{objc.GetClass("AVDateRangeMetadataGroup")}
	})
	return DateRangeMetadataGroupClass
}

type _DateRangeMetadataGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DateRangeMetadataGroup */
// An interface definition for the [DateRangeMetadataGroup] class.
type IDateRangeMetadataGroup interface {
	IMetadataGroup
	
/* debug [class_interface_properties]: Properties for DateRangeMetadataGroup */
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	Items() []MetadataItem
	StartDate() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DateRangeMetadataGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DateRangeMetadataGroup */
// Alloc allocates a new instance without initialization.
func (dc _DateRangeMetadataGroupClass) Alloc() DateRangeMetadataGroup {
	rv := objc.Send[DateRangeMetadataGroup](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DateRangeMetadataGroupClass) New() DateRangeMetadataGroup {
	rv := objc.Send[DateRangeMetadataGroup](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateRangeMetadataGroup) Init() DateRangeMetadataGroup {
	rv := objc.Send[DateRangeMetadataGroup](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateRangeMetadataGroup) Autorelease() DateRangeMetadataGroup {
	rv := objc.Send[DateRangeMetadataGroup](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateRangeMetadataGroup creates a new DateRangeMetadataGroup instance.
func NewDateRangeMetadataGroup() DateRangeMetadataGroup {
	return getDateRangeMetadataGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DateRangeMetadataGroup */
// A collection of metadata items that are valid for use within a specific date range.


// A collection of metadata items that are valid for use within a specific date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup
type DateRangeMetadataGroup struct {
	MetadataGroup
}

// DateRangeMetadataGroupFrom constructs a [DateRangeMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use within a specific date range.
func DateRangeMetadataGroupFrom(ptr unsafe.Pointer) DateRangeMetadataGroup {
	return DateRangeMetadataGroup{
		MetadataGroup: MetadataGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DateRangeMetadataGroup */

// Initializes an instance of with a collection of metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/init(items:start:end:)
func NewDateRangeMetadataGroupWithItemsStartDateEndDate(items []MetadataItem, startDate objc.IObject /* cross-framework: NSDate */, endDate objc.IObject /* cross-framework: NSDate */) DateRangeMetadataGroup {
	instance := getDateRangeMetadataGroupClass().Alloc()
	rv := objc.Send[DateRangeMetadataGroup](instance.ID, objc.Sel("initWithItems:startDate:endDate:"), items, startDate, endDate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDateRangeMetadataGroupWithItemsStartDateEndDate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DateRangeMetadataGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DateRangeMetadataGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DateRangeMetadataGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DateRangeMetadataGroup */

// The end date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/endDate
func (d_ DateRangeMetadataGroup) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// An array of associated metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/items
func (d_ DateRangeMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](d_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// The start date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup/startDate
func (d_ DateRangeMetadataGroup) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDateRangeMetadataGroup */


