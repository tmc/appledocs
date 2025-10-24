// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKContactsLensSpecification */


/* debug [class_header]: Header for HKContactsLensSpecification */
// The class instance for the [HKContactsLensSpecification] class.
var (
	HKContactsLensSpecificationClass     _HKContactsLensSpecificationClass
	HKContactsLensSpecificationClassOnce sync.Once
)

func getHKContactsLensSpecificationClass() _HKContactsLensSpecificationClass {
	HKContactsLensSpecificationClassOnce.Do(func() {
		HKContactsLensSpecificationClass = _HKContactsLensSpecificationClass{objc.GetClass("HKContactsLensSpecification")}
	})
	return HKContactsLensSpecificationClass
}

type _HKContactsLensSpecificationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKContactsLensSpecification */
// An interface definition for the [HKContactsLensSpecification] class.
type IHKContactsLensSpecification interface {
	IHKLensSpecification
	
/* debug [class_interface_properties]: Properties for HKContactsLensSpecification */
	// properties:
	BaseCurve() IHKQuantity
	Diameter() IHKQuantity
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKContactsLensSpecification */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKContactsLensSpecification */
// Alloc allocates a new instance without initialization.
func (hc _HKContactsLensSpecificationClass) Alloc() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKContactsLensSpecificationClass) New() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKContactsLensSpecification) Init() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKContactsLensSpecification) Autorelease() HKContactsLensSpecification {
	rv := objc.Send[HKContactsLensSpecification](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKContactsLensSpecification creates a new HKContactsLensSpecification instance.
func NewHKContactsLensSpecification() HKContactsLensSpecification {
	return getHKContactsLensSpecificationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKContactsLensSpecification */
// An object that contains the contacts prescription data for one eye.
//
// To create a sample that stores a contacts prescription, start by defining a specification for each eye. Each lens specification object requires a parameter. This measures the lens’s strength for correcting either nearsightedness or farsightedness (measured in units). Next, create values for any of the prescription’s optional parameters. For example, if the prescription corrects for astigmatism, create the and values. The value uses units, while the uses . To add a multifocal correction for reading, create an value using units. To add fitting information for the contact lens, create and values. Both of these values use millimeters. Then you can create the lens specification. After you create your lens specifications, you can create an sample. Then save the sample to the HealthKit store. Finally, add an image or PDF of the prescription to the sample as an attachment.


// An object that contains the contacts prescription data for one eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsLensSpecification
type HKContactsLensSpecification struct {
	HKLensSpecification
}

// HKContactsLensSpecificationFrom constructs a [HKContactsLensSpecification] from an unsafe.Pointer.
//
// An object that contains the contacts prescription data for one eye.
func HKContactsLensSpecificationFrom(ptr unsafe.Pointer) HKContactsLensSpecification {
	return HKContactsLensSpecification{
		HKLensSpecification: HKLensSpecificationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKContactsLensSpecification */

// Creates a new contact lens specification, containing the prescription data for one eye.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsLensSpecification/init(sphere:cylinder:axis:addPower:baseCurve:diameter:)
func NewHKContactsLensSpecificationWithSphereCylinderAxisAddPowerBaseCurveDiameter(sphere IHKQuantity, cylinder IHKQuantity, axis IHKQuantity, addPower IHKQuantity, baseCurve IHKQuantity, diameter IHKQuantity) HKContactsLensSpecification {
	instance := getHKContactsLensSpecificationClass().Alloc()
	rv := objc.Send[HKContactsLensSpecification](instance.ID, objc.Sel("initWithSphere:cylinder:axis:addPower:baseCurve:diameter:"), sphere, cylinder, axis, addPower, baseCurve, diameter)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewHKContactsLensSpecificationWithSphereCylinderAxisAddPowerBaseCurveDiameter */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKContactsLensSpecification */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKContactsLensSpecification */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKContactsLensSpecification */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKContactsLensSpecification */

// Part of the contact’s fit, it measures the curve of the back side of the contact, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsLensSpecification/baseCurve
func (h_ HKContactsLensSpecification) BaseCurve() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("baseCurve"))
	return rv
}/* debug [instance_properties/getter]: baseCurve */


// Part of the contact’s fit, it measures the diameter of the lens, measured in mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKContactsLensSpecification/diameter
func (h_ HKContactsLensSpecification) Diameter() IHKQuantity {
	rv := objc.Send[HKQuantity](h_.ID, objc.Sel("diameter"))
	return rv
}/* debug [instance_properties/getter]: diameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKContactsLensSpecification */


