// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MutableTimedMetadataGroup] class.
type IMutableTimedMetadataGroup interface {
	ITimedMetadataGroup
	

	// properties:
	Items() []MetadataItem
	SetItems(value []MetadataItem)
	TimeRange() objectivec.IObject
	SetTimeRange(value objectivec.IObject)


	

	// methods:


}





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

























// An array of metadata items in the timed metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/items
func (m_ MutableTimedMetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}


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
}


// The time range of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/timeRange
func (m_ MutableTimedMetadataGroup) TimeRange() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("timeRange"))
	return rv
}


// The time range of the timed metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableTimedMetadataGroup/timeRange
func (m_ MutableTimedMetadataGroup) SetTimeRange(value objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimeRange:"), value)
}








