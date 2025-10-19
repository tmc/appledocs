// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trackingFrameCount] class.
var trackingFrameCountClass = _trackingFrameCountClass{objc.GetClass("trackingFrameCount")}

type _trackingFrameCountClass struct {
	class objc.Class
}

// An interface definition for the [trackingFrameCount] class.
type ItrackingFrameCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-c.ivar

type trackingFrameCount struct {
	objectivec.Object
}

// trackingFrameCountFrom constructs a [trackingFrameCount] from an unsafe.Pointer.
func trackingFrameCountFrom(ptr unsafe.Pointer) trackingFrameCount {
	return trackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (tc _trackingFrameCountClass) Alloc() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (tc _trackingFrameCountClass) New() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trackingFrameCount) Init() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trackingFrameCount) Autorelease() trackingFrameCount {
	rv := objc.Send[trackingFrameCount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrackingFrameCount creates a new trackingFrameCount instance.
func NewtrackingFrameCount() trackingFrameCount {
	return trackingFrameCountClass.New()
}




