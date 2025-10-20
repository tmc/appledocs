// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [DateRangeMetadataGroup] class.
type IDateRangeMetadataGroup interface {
	objectivec.IObject
}

// A collection of metadata items that are valid for use within a specific date range.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVDateRangeMetadataGroup
type DateRangeMetadataGroup struct {
	objectivec.Object
}

// DateRangeMetadataGroupFrom constructs a [DateRangeMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use within a specific date range.
func DateRangeMetadataGroupFrom(ptr unsafe.Pointer) DateRangeMetadataGroup {
	return DateRangeMetadataGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DateRangeMetadataGroupClass) Alloc() DateRangeMetadataGroup {
	rv := objc.Send[DateRangeMetadataGroup](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




