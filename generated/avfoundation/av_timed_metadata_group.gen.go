// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVTimedMetadataGroup */


/* debug [class_header]: Header for AVTimedMetadataGroup */
// The class instance for the [TimedMetadataGroup] class.
var (
	TimedMetadataGroupClass     _TimedMetadataGroupClass
	TimedMetadataGroupClassOnce sync.Once
)

func getTimedMetadataGroupClass() _TimedMetadataGroupClass {
	TimedMetadataGroupClassOnce.Do(func() {
		TimedMetadataGroupClass = _TimedMetadataGroupClass{objc.GetClass("AVTimedMetadataGroup")}
	})
	return TimedMetadataGroupClass
}

type _TimedMetadataGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TimedMetadataGroup */
// An interface definition for the [TimedMetadataGroup] class.
type ITimedMetadataGroup interface {
	IMetadataGroup
	
/* debug [class_interface_properties]: Properties for TimedMetadataGroup */
	// properties:
	Items() []MetadataItem
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TimedMetadataGroup */
	// methods:
	CopyFormatDescription() MetadataFormatDescriptionRef /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TimedMetadataGroup */
// Alloc allocates a new instance without initialization.
func (tc _TimedMetadataGroupClass) Alloc() TimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TimedMetadataGroupClass) New() TimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TimedMetadataGroup) Init() TimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TimedMetadataGroup) Autorelease() TimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimedMetadataGroup creates a new TimedMetadataGroup instance.
func NewTimedMetadataGroup() TimedMetadataGroup {
	return getTimedMetadataGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TimedMetadataGroup */
// A collection of metadata items that are valid for use during a specific time range.
//
// For example, are used to represent chapters, optionally containing metadata items for chapter titles and chapter images.


// A collection of metadata items that are valid for use during a specific time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup
type TimedMetadataGroup struct {
	MetadataGroup
}

// TimedMetadataGroupFrom constructs a [TimedMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use during a specific time range.
func TimedMetadataGroupFrom(ptr unsafe.Pointer) TimedMetadataGroup {
	return TimedMetadataGroup{
		MetadataGroup: MetadataGroupFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TimedMetadataGroup */

// Creates a timed metadata group initialized with the given metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(items:timeRange:)
func NewTimedMetadataGroupWithItemsTimeRange(items []MetadataItem, timeRange TimeRange /* not a class type */) TimedMetadataGroup {
	instance := getTimedMetadataGroupClass().Alloc()
	rv := objc.Send[TimedMetadataGroup](instance.ID, objc.Sel("initWithItems:timeRange:"), items, timeRange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimedMetadataGroupWithItemsTimeRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(sampleBuffer:)-bjuo
func NewTimedMetadataGroupWithSampleBuffer(sampleBuffer SampleBufferRef /* not a class type */) TimedMetadataGroup {
	instance := getTimedMetadataGroupClass().Alloc()
	rv := objc.Send[TimedMetadataGroup](instance.ID, objc.Sel("initWithSampleBuffer:"), sampleBuffer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTimedMetadataGroupWithSampleBuffer */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TimedMetadataGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TimedMetadataGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TimedMetadataGroup */

// Creates a format description based on the receiver’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/copyFormatDescription()
func (t_ TimedMetadataGroup) CopyFormatDescription() MetadataFormatDescriptionRef /* not a class type */ {
	rv := objc.Send[MetadataFormatDescriptionRef](t_.ID, objc.Sel("copyFormatDescription"))
	return rv
}/* debug [instance_methods/method]: CopyFormatDescription */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TimedMetadataGroup */

// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/items
func (t_ TimedMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](t_.ID, objc.Sel("items"))
	return rv
}/* debug [instance_properties/getter]: items */


// The time range for the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/timeRange
func (t_ TimedMetadataGroup) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](t_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVTimedMetadataGroup */


