// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKGlassesLensSpecification */


/* debug [class_header]: Header for HKGlassesLensSpecification */
// The class instance for the [HKGlassesLensSpecification] class.
var (
	HKGlassesLensSpecificationClass     _HKGlassesLensSpecificationClass
	HKGlassesLensSpecificationClassOnce sync.Once
)

func getHKGlassesLensSpecificationClass() _HKGlassesLensSpecificationClass {
	HKGlassesLensSpecificationClassOnce.Do(func() {
		HKGlassesLensSpecificationClass = _HKGlassesLensSpecificationClass{objc.GetClass("HKGlassesLensSpecification")}
	})
	return HKGlassesLensSpecificationClass
}

type _HKGlassesLensSpecificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKGlassesLensSpecification */
// An interface definition for the [HKGlassesLensSpecification] class.
type IHKGlassesLensSpecification interface {
	IHKLensSpecification
	
/* debug [class_interface_properties]: Properties for HKGlassesLensSpecification */
	// properties:
	FarPupillaryDistance() IHKQuantity
	NearPupillaryDistance() IHKQuantity
	Prism() IHKVisionPrism
	VertexDistance() IHKQuantity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKGlassesLensSpecification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKGlassesLensSpecification */
// Alloc allocates a new instance without initialization.
func (hc _HKGlassesLensSpecificationClass) Alloc() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKGlassesLensSpecificationClass) New() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKGlassesLensSpecification) Init() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKGlassesLensSpecification) Autorelease() HKGlassesLensSpecification {
	rv := objc.Send[HKGlassesLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKGlassesLensSpecification creates a new HKGlassesLensSpecification instance.
func NewHKGlassesLensSpecification() HKGlassesLensSpecification {
	return getHKGlassesLensSpecificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKGlassesLensSpecification */
// An object that contains the glasses prescription data for one eye.
//
// To create a sample that stores a glasses prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add a correction for eye alignment, create an object. To add information about the distance between the eye and the back of the lens, or the pupil and the center of the nose, create , , and values. All of these use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.


// An object that contains the glasses prescription data for one eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification
type HKGlassesLensSpecification struct {
	HKLensSpecification
}

// HKGlassesLensSpecificationFrom constructs a [HKGlassesLensSpecification] from an unsafe.Pointer.
//
// An object that contains the glasses prescription data for one eye.
func HKGlassesLensSpecificationFrom(ptr unsafe.Pointer) HKGlassesLensSpecification {
	return HKGlassesLensSpecification{
		HKLensSpecification: HKLensSpecificationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKGlassesLensSpecification */

// Creates a new glasses lens specification, containing the prescription data for one eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification/init(sphere:cylinder:axis:addPower:vertexDistance:prism:farPupillaryDistance:nearPupillaryDistance:)
func NewHKGlassesLensSpecificationWithSphereCylinderAxisAddPowerVertexDistancePrismFarPupillaryDistanceNearPupillaryDistance(sphere IHKQuantity, cylinder IHKQuantity, axis IHKQuantity, addPower IHKQuantity, vertexDistance IHKQuantity, prism IHKVisionPrism, farPupillaryDistance IHKQuantity, nearPupillaryDistance IHKQuantity) HKGlassesLensSpecification {
	instance := getHKGlassesLensSpecificationClass().Alloc()
	rv := objc.Send[HKGlassesLensSpecification](instance.ID, objc.Sel("initWithSphere:cylinder:axis:addPower:vertexDistance:prism:farPupillaryDistance:nearPupillaryDistance:"), sphere, cylinder, axis, addPower, vertexDistance, prism, farPupillaryDistance, nearPupillaryDistance)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKGlassesLensSpecificationWithSphereCylinderAxisAddPowerVertexDistancePrismFarPupillaryDistanceNearPupillaryDistance */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKGlassesLensSpecification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKGlassesLensSpecification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKGlassesLensSpecification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKGlassesLensSpecification */

// The distance between the pupil and the center of the nose when looking at an object far away, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification/farPupillaryDistance
func (h_ HKGlassesLensSpecification) FarPupillaryDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("farPupillaryDistance"))
	return rv
}/* debug [instance_properties/getter]: farPupillaryDistance */


// The distance between the pupil and the center of the nose when looking at a nearby object, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification/nearPupillaryDistance
func (h_ HKGlassesLensSpecification) NearPupillaryDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("nearPupillaryDistance"))
	return rv
}/* debug [instance_properties/getter]: nearPupillaryDistance */


// An object that contains information about the eye alignment correction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification/prism
func (h_ HKGlassesLensSpecification) Prism() IHKVisionPrism {
	rv := objc.Send[HKVisionPrism](h_.ID, objc.Sel("prism"))
	return rv
}/* debug [instance_properties/getter]: prism */


// The distance between the back of the lens and the eye, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKGlassesLensSpecification/vertexDistance
func (h_ HKGlassesLensSpecification) VertexDistance() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("vertexDistance"))
	return rv
}/* debug [instance_properties/getter]: vertexDistance */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKGlassesLensSpecification */


