// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKUserAnnotatedMedicationType */


/* debug [class_header]: Header for HKUserAnnotatedMedicationType */
// The class instance for the [HKUserAnnotatedMedicationType] class.
var (
	HKUserAnnotatedMedicationTypeClass     _HKUserAnnotatedMedicationTypeClass
	HKUserAnnotatedMedicationTypeClassOnce sync.Once
)

func getHKUserAnnotatedMedicationTypeClass() _HKUserAnnotatedMedicationTypeClass {
	HKUserAnnotatedMedicationTypeClassOnce.Do(func() {
		HKUserAnnotatedMedicationTypeClass = _HKUserAnnotatedMedicationTypeClass{objc.GetClass("HKUserAnnotatedMedicationType")}
	})
	return HKUserAnnotatedMedicationTypeClass
}

type _HKUserAnnotatedMedicationTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKUserAnnotatedMedicationType */
// An interface definition for the [HKUserAnnotatedMedicationType] class.
type IHKUserAnnotatedMedicationType interface {
	IHKObjectType
	
/* debug [class_interface_properties]: Properties for HKUserAnnotatedMedicationType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKUserAnnotatedMedicationType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKUserAnnotatedMedicationType */
// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationTypeClass) Alloc() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKUserAnnotatedMedicationTypeClass) New() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUserAnnotatedMedicationType) Init() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUserAnnotatedMedicationType) Autorelease() HKUserAnnotatedMedicationType {
	rv := objc.Send[HKUserAnnotatedMedicationType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUserAnnotatedMedicationType creates a new HKUserAnnotatedMedicationType instance.
func NewHKUserAnnotatedMedicationType() HKUserAnnotatedMedicationType {
	return getHKUserAnnotatedMedicationTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKUserAnnotatedMedicationType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedicationType
type HKUserAnnotatedMedicationType struct {
	HKObjectType
}

// HKUserAnnotatedMedicationTypeFrom constructs a [HKUserAnnotatedMedicationType] from an unsafe.Pointer.
func HKUserAnnotatedMedicationTypeFrom(ptr unsafe.Pointer) HKUserAnnotatedMedicationType {
	return HKUserAnnotatedMedicationType{
		HKObjectType: HKObjectTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKUserAnnotatedMedicationType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKUserAnnotatedMedicationType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKUserAnnotatedMedicationType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKUserAnnotatedMedicationType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKUserAnnotatedMedicationType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKUserAnnotatedMedicationType */



