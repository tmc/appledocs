// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasTrackingFrameCount] class.
var (
	HasTrackingFrameCountClass     _hasTrackingFrameCountClass
	HasTrackingFrameCountClassOnce sync.Once
)

func gethasTrackingFrameCountClass() _hasTrackingFrameCountClass {
	HasTrackingFrameCountClassOnce.Do(func() {
		HasTrackingFrameCountClass = _hasTrackingFrameCountClass{objc.GetClass("hasTrackingFrameCount")}
	})
	return HasTrackingFrameCountClass
}

type _hasTrackingFrameCountClass struct {
	class objc.Class
}

// An interface definition for the [hasTrackingFrameCount] class.
type IhasTrackingFrameCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-c.ivar
type hasTrackingFrameCount struct {
	objectivec.Object
}

// hasTrackingFrameCountFrom constructs a [hasTrackingFrameCount] from an unsafe.Pointer.
func hasTrackingFrameCountFrom(ptr unsafe.Pointer) hasTrackingFrameCount {
	return hasTrackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _hasTrackingFrameCountClass) Alloc() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _hasTrackingFrameCountClass) New() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasTrackingFrameCount) Init() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasTrackingFrameCount) Autorelease() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasTrackingFrameCount creates a new hasTrackingFrameCount instance.
func NewhasTrackingFrameCount() hasTrackingFrameCount {
	return gethasTrackingFrameCountClass().New()
}




