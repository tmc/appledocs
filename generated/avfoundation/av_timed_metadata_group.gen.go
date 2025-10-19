// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVTimedMetadataGroup] class.
var aVTimedMetadataGroupClass = _AVTimedMetadataGroupClass{objc.GetClass("AVTimedMetadataGroup")}

type _AVTimedMetadataGroupClass struct {
	class objc.Class
}

// An interface definition for the [AVTimedMetadataGroup] class.
type IAVTimedMetadataGroup interface {
	IAVMetadataGroup
}

// A collection of metadata items that are valid for use during a specific time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup

type AVTimedMetadataGroup struct {
	AVMetadataGroup
}

// AVTimedMetadataGroupFrom constructs a [AVTimedMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use during a specific time range.
func AVTimedMetadataGroupFrom(ptr unsafe.Pointer) AVTimedMetadataGroup {
	return AVTimedMetadataGroup{
		AVMetadataGroup: AVMetadataGroupFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _AVTimedMetadataGroupClass) Alloc() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVTimedMetadataGroupClass) New() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVTimedMetadataGroup) Init() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVTimedMetadataGroup) Autorelease() AVTimedMetadataGroup {
	rv := objc.Send[AVTimedMetadataGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVTimedMetadataGroup creates a new AVTimedMetadataGroup instance.
func NewAVTimedMetadataGroup() AVTimedMetadataGroup {
	return aVTimedMetadataGroupClass.New()
}




