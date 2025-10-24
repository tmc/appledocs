// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKClinicalType */


/* debug [class_header]: Header for HKClinicalType */
// The class instance for the [HKClinicalType] class.
var (
	HKClinicalTypeClass     _HKClinicalTypeClass
	HKClinicalTypeClassOnce sync.Once
)

func getHKClinicalTypeClass() _HKClinicalTypeClass {
	HKClinicalTypeClassOnce.Do(func() {
		HKClinicalTypeClass = _HKClinicalTypeClass{objc.GetClass("HKClinicalType")}
	})
	return HKClinicalTypeClass
}

type _HKClinicalTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKClinicalType */
// An interface definition for the [HKClinicalType] class.
type IHKClinicalType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKClinicalType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKClinicalType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKClinicalType */
// Alloc allocates a new instance without initialization.
func (hc _HKClinicalTypeClass) Alloc() HKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKClinicalTypeClass) New() HKClinicalType {
	rv := objc.Send[HKClinicalType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKClinicalType) Init() HKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKClinicalType) Autorelease() HKClinicalType {
	rv := objc.Send[HKClinicalType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKClinicalType creates a new HKClinicalType instance.
func NewHKClinicalType() HKClinicalType {
	return getHKClinicalTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKClinicalType */
// A type that identifies samples that contain clinical record data.


// A type that identifies samples that contain clinical record data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalType
type HKClinicalType struct {
	HKSampleType
}

// HKClinicalTypeFrom constructs a [HKClinicalType] from an unsafe.Pointer.
//
// A type that identifies samples that contain clinical record data.
func HKClinicalTypeFrom(ptr unsafe.Pointer) HKClinicalType {
	return HKClinicalType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKClinicalType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKClinicalType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKClinicalType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKClinicalType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKClinicalType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKClinicalType */



