// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVMetadataGroup] class.
var (
	aVMetadataGroupClass     _AVMetadataGroupClass
	aVMetadataGroupClassOnce sync.Once
)

func getAVMetadataGroupClass() _AVMetadataGroupClass {
	aVMetadataGroupClassOnce.Do(func() {
		aVMetadataGroupClass = _AVMetadataGroupClass{objc.GetClass("AVMetadataGroup")}
	})
	return aVMetadataGroupClass
}

type _AVMetadataGroupClass struct {
	class objc.Class
}

// An interface definition for the [AVMetadataGroup] class.
type IAVMetadataGroup interface {
	objectivec.IObject
}

// A parent class referenced by other AVFoundation classes. [Full Topic]
type AVMetadataGroup struct {
	objectivec.Object
}

// AVMetadataGroupFrom constructs a [AVMetadataGroup] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func AVMetadataGroupFrom(ptr unsafe.Pointer) AVMetadataGroup {
	return AVMetadataGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVMetadataGroupClass) Alloc() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVMetadataGroupClass) New() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVMetadataGroup) Init() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVMetadataGroup) Autorelease() AVMetadataGroup {
	rv := objc.Send[AVMetadataGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetadataGroup creates a new AVMetadataGroup instance.
func NewAVMetadataGroup() AVMetadataGroup {
	return getAVMetadataGroupClass().New()
}




