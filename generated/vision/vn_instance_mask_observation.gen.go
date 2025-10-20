// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	GenerateMaskForInstancesError(instances unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances unsafe.Pointer, requestHandler unsafe.Pointer, cropResult bool, error_ unsafe.Pointer) unsafe.Pointer
	GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances unsafe.Pointer, requestHandler unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
}

// An observation that contains an instance mask that labels instances in the mask.
//
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


// Creates a low-resolution mask from the instances you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMask(forInstances:)
func (i_ InstanceMaskObservation) GenerateMaskForInstancesError(instances unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("generateMaskForInstances:error:"), instances, error_)
	return rv
}

// Creates a high-resolution image where everything becomes transparent black, except for the instances you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMaskedImage(ofInstances:from:croppedToInstancesExtent:)
func (i_ InstanceMaskObservation) GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances unsafe.Pointer, requestHandler unsafe.Pointer, cropResult bool, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("generateMaskedImageOfInstances:fromRequestHandler:croppedToInstancesExtent:error:"), instances, requestHandler, cropResult, error_)
	return rv
}

// Creates a high-resolution mask where everything becomes transparent black, except for the instances you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateScaledMaskForImage(forInstances:from:)
func (i_ InstanceMaskObservation) GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances unsafe.Pointer, requestHandler unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("generateScaledMaskForImageForInstances:fromRequestHandler:error:"), instances, requestHandler, error_)
	return rv
}

// The collection that contains all instances, excluding the background.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/allInstances
func (i_ InstanceMaskObservation) AllInstances() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("allInstances"))
	return rv
}

// The resulting mask that represents all instances.
//
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/instanceMask
func (i_ InstanceMaskObservation) InstanceMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("instanceMask"))
	return rv
}



