// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [InstanceMaskObservation] class.
var (
	InstanceMaskObservationClass     _InstanceMaskObservationClass
	InstanceMaskObservationClassOnce sync.Once
)

func getInstanceMaskObservationClass() _InstanceMaskObservationClass {
	InstanceMaskObservationClassOnce.Do(func() {
		InstanceMaskObservationClass = _InstanceMaskObservationClass{objc.GetClass("VNInstanceMaskObservation")}
	})
	return InstanceMaskObservationClass
}

type _InstanceMaskObservationClass struct {
	class objc.Class
}

// An interface definition for the [InstanceMaskObservation] class.
type IInstanceMaskObservation interface {
	IObservation
	// properties:
	VNGenerateForegroundInstanceMaskRequestRevision1() int
	AllInstances() objc.IObject /* cross-framework: IndexSet */
	SetAllInstances(value objc.IObject /* cross-framework: IndexSet */)
	InstanceMask() PixelBuffer /* not a class type */
	SetInstanceMask(value PixelBuffer /* not a class type */)
	// methods:
}

// An observation that contains an instance mask that labels instances in the mask.


// An observation that contains an instance mask that labels instances in the mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation
type InstanceMaskObservation struct {
	Observation
}

// InstanceMaskObservationFrom constructs a [InstanceMaskObservation] from an unsafe.Pointer.
//
// An observation that contains an instance mask that labels instances in the mask.
func InstanceMaskObservationFrom(ptr unsafe.Pointer) InstanceMaskObservation {
	return InstanceMaskObservation{
		Observation: ObservationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InstanceMaskObservationClass) Alloc() InstanceMaskObservation {
	rv := objc.Send[InstanceMaskObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InstanceMaskObservationClass) New() InstanceMaskObservation {
	rv := objc.Send[InstanceMaskObservation](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InstanceMaskObservation) Init() InstanceMaskObservation {
	rv := objc.Send[InstanceMaskObservation](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InstanceMaskObservation) Autorelease() InstanceMaskObservation {
	rv := objc.Send[InstanceMaskObservation](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInstanceMaskObservation creates a new InstanceMaskObservation instance.
func NewInstanceMaskObservation() InstanceMaskObservation {
	return getInstanceMaskObservationClass().New()
}



// A constant for specifying the first revision of the foreground instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (i_ InstanceMaskObservation) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](i_.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}


// The collection that contains all instances, excluding the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vninstancemaskobservation/allinstances
func (i_ InstanceMaskObservation) AllInstances() objc.IObject /* cross-framework: IndexSet */ {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("allInstances"))
	return rv
}


// The collection that contains all instances, excluding the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vninstancemaskobservation/allinstances
func (i_ InstanceMaskObservation) SetAllInstances(value objc.IObject /* cross-framework: IndexSet */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAllInstances:"), value)
}


// The resulting mask that represents all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vninstancemaskobservation/instancemask
func (i_ InstanceMaskObservation) InstanceMask() PixelBuffer /* not a class type */ {
	rv := objc.Send[PixelBuffer](i_.ID, objc.Sel("instanceMask"))
	return rv
}


// The resulting mask that represents all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vninstancemaskobservation/instancemask
func (i_ InstanceMaskObservation) SetInstanceMask(value PixelBuffer /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setInstanceMask:"), value)
}



