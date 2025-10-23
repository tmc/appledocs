// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TimedMetadataGroup] class.
type ITimedMetadataGroup interface {
	objectivec.IObject
	TimeRange() unsafe.Pointer
	Items() AVMetadataItem
	SetItems(value IAVMetadataItem)
}

// A collection of metadata items that are valid for use during a specific time range.
//
// For example, are used to represent chapters, optionally containing metadata items for chapter titles and chapter images.


// A collection of metadata items that are valid for use during a specific time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup
type TimedMetadataGroup struct {
	objectivec.Object
}

// TimedMetadataGroupFrom constructs a [TimedMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use during a specific time range.
func TimedMetadataGroupFrom(ptr unsafe.Pointer) TimedMetadataGroup {
	return TimedMetadataGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TimedMetadataGroupClass) Alloc() TimedMetadataGroup {
	rv := objc.Send[TimedMetadataGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a timed metadata group initialized with the given metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/init(items:timeRange:)
func NewTimedMetadataGroupWithItemsTimeRange(items []MetadataItem, timeRange unsafe.Pointer) TimedMetadataGroup {
	instance := getTimedMetadataGroupClass().Alloc()
	rv := objc.Send[TimedMetadataGroup](instance.ID, objc.Sel("initWithItems:timeRange:"), items, timeRange)
	rv.Autorelease()
	return rv
}



// The time range for the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup/timeRange
func (t_ TimedMetadataGroup) TimeRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("timeRange"))
	return rv
}


// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/items
func (t_ TimedMetadataGroup) Items() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](t_.ID, objc.Sel("items"))
	return rv
}


// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/items
func (t_ TimedMetadataGroup) SetItems(value IAVMetadataItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItems:"), value)
}


