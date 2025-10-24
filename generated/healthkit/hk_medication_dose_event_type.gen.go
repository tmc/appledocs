// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class HKMedicationDoseEventType */


/* debug [class_header]: Header for HKMedicationDoseEventType */
// The class instance for the [HKMedicationDoseEventType] class.
var (
	HKMedicationDoseEventTypeClass     _HKMedicationDoseEventTypeClass
	HKMedicationDoseEventTypeClassOnce sync.Once
)

func getHKMedicationDoseEventTypeClass() _HKMedicationDoseEventTypeClass {
	HKMedicationDoseEventTypeClassOnce.Do(func() {
		HKMedicationDoseEventTypeClass = _HKMedicationDoseEventTypeClass{objc.GetClass("HKMedicationDoseEventType")}
	})
	return HKMedicationDoseEventTypeClass
}

type _HKMedicationDoseEventTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKMedicationDoseEventType */
// An interface definition for the [HKMedicationDoseEventType] class.
type IHKMedicationDoseEventType interface {
	IHKSampleType
	
/* debug [class_interface_properties]: Properties for HKMedicationDoseEventType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKMedicationDoseEventType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKMedicationDoseEventType */
// Alloc allocates a new instance without initialization.
func (hc _HKMedicationDoseEventTypeClass) Alloc() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKMedicationDoseEventTypeClass) New() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKMedicationDoseEventType) Init() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKMedicationDoseEventType) Autorelease() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKMedicationDoseEventType creates a new HKMedicationDoseEventType instance.
func NewHKMedicationDoseEventType() HKMedicationDoseEventType {
	return getHKMedicationDoseEventTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKMedicationDoseEventType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEventType
type HKMedicationDoseEventType struct {
	HKSampleType
}

// HKMedicationDoseEventTypeFrom constructs a [HKMedicationDoseEventType] from an unsafe.Pointer.
func HKMedicationDoseEventTypeFrom(ptr unsafe.Pointer) HKMedicationDoseEventType {
	return HKMedicationDoseEventType{
		HKSampleType: HKSampleTypeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKMedicationDoseEventType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKMedicationDoseEventType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKMedicationDoseEventType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKMedicationDoseEventType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKMedicationDoseEventType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKMedicationDoseEventType */



