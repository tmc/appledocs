// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasTrackingID] class.
var hasTrackingIDClass = _hasTrackingIDClass{objc.GetClass("hasTrackingID")}

type _hasTrackingIDClass struct {
	class objc.Class
}

// An interface definition for the [hasTrackingID] class.
type IhasTrackingID interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-c.ivar

type hasTrackingID struct {
	objectivec.Object
}

// hasTrackingIDFrom constructs a [hasTrackingID] from an unsafe.Pointer.
func hasTrackingIDFrom(ptr unsafe.Pointer) hasTrackingID {
	return hasTrackingID{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (hc _hasTrackingIDClass) Alloc() hasTrackingID {
	rv := objc.Send[hasTrackingID](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (hc _hasTrackingIDClass) New() hasTrackingID {
	rv := objc.Send[hasTrackingID](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasTrackingID) Init() hasTrackingID {
	rv := objc.Send[hasTrackingID](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasTrackingID) Autorelease() hasTrackingID {
	rv := objc.Send[hasTrackingID](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasTrackingID creates a new hasTrackingID instance.
func NewhasTrackingID() hasTrackingID {
	return hasTrackingIDClass.New()
}




