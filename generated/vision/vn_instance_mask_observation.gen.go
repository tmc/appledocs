// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNInstanceMaskObservation */


/* debug [class_header]: Header for VNInstanceMaskObservation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InstanceMaskObservation */
// An interface definition for the [InstanceMaskObservation] class.
type IInstanceMaskObservation interface {
	IObservation
	
/* debug [class_interface_properties]: Properties for InstanceMaskObservation */
	// properties:
	AllInstances() foundation.IndexSet
	InstanceMask() PixelBufferRef /* not a class type */
	VNGenerateForegroundInstanceMaskRequestRevision1() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InstanceMaskObservation */
	// methods:
	GenerateMaskForInstancesError(instances foundation.IndexSet, error_ objectivec.IObject) PixelBufferRef /* not a class type */
	GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances foundation.IndexSet, requestHandler IVNImageRequestHandler, cropResult bool, error_ objectivec.IObject) PixelBufferRef /* not a class type */
	GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances foundation.IndexSet, requestHandler IVNImageRequestHandler, error_ objectivec.IObject) PixelBufferRef /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InstanceMaskObservation */
// Alloc allocates a new instance without initialization.
func (ic _InstanceMaskObservationClass) Alloc() InstanceMaskObservation {
	rv := objc.Send[InstanceMaskObservation](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InstanceMaskObservation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InstanceMaskObservation *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InstanceMaskObservation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InstanceMaskObservation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InstanceMaskObservation */

// Creates a low-resolution mask from the instances you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMask(forInstances:)
func (i_ InstanceMaskObservation) GenerateMaskForInstancesError(instances foundation.IndexSet, error_ objectivec.IObject) PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](i_.ID, objc.Sel("generateMaskForInstances:error:"), instances, error_)
	return rv
}/* debug [instance_methods/method]: GenerateMaskForInstancesError */


// Creates a high-resolution image where everything becomes transparent black, except for the instances you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateMaskedImage(ofInstances:from:croppedToInstancesExtent:)
func (i_ InstanceMaskObservation) GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError(instances foundation.IndexSet, requestHandler IVNImageRequestHandler, cropResult bool, error_ objectivec.IObject) PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](i_.ID, objc.Sel("generateMaskedImageOfInstances:fromRequestHandler:croppedToInstancesExtent:error:"), instances, requestHandler, cropResult, error_)
	return rv
}/* debug [instance_methods/method]: GenerateMaskedImageOfInstancesFromRequestHandlerCroppedToInstancesExtentError */


// Creates a high-resolution mask where everything becomes transparent black, except for the instances you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/generateScaledMaskForImage(forInstances:from:)
func (i_ InstanceMaskObservation) GenerateScaledMaskForImageForInstancesFromRequestHandlerError(instances foundation.IndexSet, requestHandler IVNImageRequestHandler, error_ objectivec.IObject) PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](i_.ID, objc.Sel("generateScaledMaskForImageForInstances:fromRequestHandler:error:"), instances, requestHandler, error_)
	return rv
}/* debug [instance_methods/method]: GenerateScaledMaskForImageForInstancesFromRequestHandlerError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InstanceMaskObservation */

// The collection that contains all instances, excluding the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/allInstances
func (i_ InstanceMaskObservation) AllInstances() foundation.IndexSet {
	rv := objc.Send[foundation.IndexSet](i_.ID, objc.Sel("allInstances"))
	return rv
}/* debug [instance_properties/getter]: allInstances */


// The resulting mask that represents all instances.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNInstanceMaskObservation/instanceMask
func (i_ InstanceMaskObservation) InstanceMask() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](i_.ID, objc.Sel("instanceMask"))
	return rv
}/* debug [instance_properties/getter]: instanceMask */


// A constant for specifying the first revision of the foreground instance mask request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/vision/vngenerateforegroundinstancemaskrequestrevision1
func (i_ InstanceMaskObservation) VNGenerateForegroundInstanceMaskRequestRevision1() int {
	rv := objc.Send[int](i_.ID, objc.Sel("VNGenerateForegroundInstanceMaskRequestRevision1"))
	return rv
}/* debug [instance_properties/getter]: VNGenerateForegroundInstanceMaskRequestRevision1 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNInstanceMaskObservation */



