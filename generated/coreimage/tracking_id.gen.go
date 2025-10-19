// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trackingID] class.
var (
	trackingIDClass     _trackingIDClass
	trackingIDClassOnce sync.Once
)

func gettrackingIDClass() _trackingIDClass {
	trackingIDClassOnce.Do(func() {
		trackingIDClass = _trackingIDClass{objc.GetClass("trackingID")}
	})
	return trackingIDClass
}

type _trackingIDClass struct {
	class objc.Class
}

// An interface definition for the [trackingID] class.
type ItrackingID interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-c.ivar
type trackingID struct {
	objectivec.Object
}

// trackingIDFrom constructs a [trackingID] from an unsafe.Pointer.
func trackingIDFrom(ptr unsafe.Pointer) trackingID {
	return trackingID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _trackingIDClass) Alloc() trackingID {
	rv := objc.Send[trackingID](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _trackingIDClass) New() trackingID {
	rv := objc.Send[trackingID](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trackingID) Init() trackingID {
	rv := objc.Send[trackingID](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trackingID) Autorelease() trackingID {
	rv := objc.Send[trackingID](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrackingID creates a new trackingID instance.
func NewtrackingID() trackingID {
	return gettrackingIDClass().New()
}




