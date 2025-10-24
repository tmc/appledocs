// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RectangleObservation] class.
var (
	RectangleObservationClass     _RectangleObservationClass
	RectangleObservationClassOnce sync.Once
)

func getRectangleObservationClass() _RectangleObservationClass {
	RectangleObservationClassOnce.Do(func() {
		RectangleObservationClass = _RectangleObservationClass{objc.GetClass("VNRectangleObservation")}
	})
	return RectangleObservationClass
}

type _RectangleObservationClass struct {
	class objc.Class
}

// An interface definition for the [RectangleObservation] class.
type IRectangleObservation interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Vision classes.


// A parent class referenced by other Vision classes. [Full Topic]
type RectangleObservation struct {
	objectivec.Object
}

// RectangleObservationFrom constructs a [RectangleObservation] from an unsafe.Pointer.
//
// A parent class referenced by other Vision classes.
func RectangleObservationFrom(ptr unsafe.Pointer) RectangleObservation {
	return RectangleObservation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RectangleObservationClass) Alloc() RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RectangleObservationClass) New() RectangleObservation {
	rv := objc.Send[RectangleObservation](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RectangleObservation) Init() RectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RectangleObservation) Autorelease() RectangleObservation {
	rv := objc.Send[RectangleObservation](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRectangleObservation creates a new RectangleObservation instance.
func NewRectangleObservation() RectangleObservation {
	return getRectangleObservationClass().New()
}




