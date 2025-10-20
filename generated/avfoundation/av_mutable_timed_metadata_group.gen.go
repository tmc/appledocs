// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A mutable collection of metadata items that are valid for use during a specific time range.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MutableTimedMetadataGroupClass) Alloc() MutableTimedMetadataGroup {
	rv := objc.Send[MutableTimedMetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




