// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKPrescriptionType */


/* debug [class_header]: Header for HKPrescriptionType */
// The class instance for the [HKPrescriptionType] class.
var (
	HKPrescriptionTypeClass     _HKPrescriptionTypeClass
	HKPrescriptionTypeClassOnce sync.Once
)

func getHKPrescriptionTypeClass() _HKPrescriptionTypeClass {
	HKPrescriptionTypeClassOnce.Do(func() {
		HKPrescriptionTypeClass = _HKPrescriptionTypeClass{objc.GetClass("HKPrescriptionType")}
	})
	return HKPrescriptionTypeClass
}

type _HKPrescriptionTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKPrescriptionType */
// An interface definition for the [HKPrescriptionType] class.
type IHKPrescriptionType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKPrescriptionType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKPrescriptionType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKPrescriptionType */
// Alloc allocates a new instance without initialization.
func (hc _HKPrescriptionTypeClass) Alloc() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKPrescriptionTypeClass) New() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKPrescriptionType) Init() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKPrescriptionType) Autorelease() HKPrescriptionType {
	rv := objc.Send[HKPrescriptionType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKPrescriptionType creates a new HKPrescriptionType instance.
func NewHKPrescriptionType() HKPrescriptionType {
	return getHKPrescriptionTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKPrescriptionType */
// A type that identifies samples that store a prescription.
//
// The class is a concrete subclass of the class. To create a vision prescription type instances, use the convenience method. Use this data type to request permission to save vision prescriptions to the HealthKit store.


// A type that identifies samples that store a prescription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKPrescriptionType
type HKPrescriptionType struct {
	HKSampleType
}

// HKPrescriptionTypeFrom constructs a [HKPrescriptionType] from an unsafe.Pointer.
//
// A type that identifies samples that store a prescription.
func HKPrescriptionTypeFrom(ptr unsafe.Pointer) HKPrescriptionType {
	return HKPrescriptionType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKPrescriptionType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKPrescriptionType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKPrescriptionType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKPrescriptionType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKPrescriptionType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKPrescriptionType */



