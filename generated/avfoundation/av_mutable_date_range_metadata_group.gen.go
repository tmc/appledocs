// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MutableDateRangeMetadataGroup] class.
type IMutableDateRangeMetadataGroup interface {
	IDateRangeMetadataGroup
	

	// properties:
	EndDate() foundation.foundation.INSDate
	SetEndDate(value foundation.foundation.INSDate)
	Items() []MetadataItem
	SetItems(value []MetadataItem)
	StartDate() foundation.foundation.INSDate
	SetStartDate(value foundation.foundation.INSDate)


	

	// methods:


}





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

























// The end date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/endDate
func (m_ MutableDateRangeMetadataGroup) EndDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("endDate"))
	return rv
}


// The end date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/endDate
func (m_ MutableDateRangeMetadataGroup) SetEndDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndDate:"), value)
}


// An array of associated metadata items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/items
func (m_ MutableDateRangeMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}


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
}


// The start date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/startDate
func (m_ MutableDateRangeMetadataGroup) StartDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("startDate"))
	return rv
}


// The start date for the metadata date range group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableDateRangeMetadataGroup/startDate
func (m_ MutableDateRangeMetadataGroup) SetStartDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartDate:"), value)
}








