// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKUserAnnotatedMedication */


/* debug [class_header]: Header for HKUserAnnotatedMedication */
// The class instance for the [HKUserAnnotatedMedication] class.
var (
	HKUserAnnotatedMedicationClass     _HKUserAnnotatedMedicationClass
	HKUserAnnotatedMedicationClassOnce sync.Once
)

func getHKUserAnnotatedMedicationClass() _HKUserAnnotatedMedicationClass {
	HKUserAnnotatedMedicationClassOnce.Do(func() {
		HKUserAnnotatedMedicationClass = _HKUserAnnotatedMedicationClass{objc.GetClass("HKUserAnnotatedMedication")}
	})
	return HKUserAnnotatedMedicationClass
}

type _HKUserAnnotatedMedicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKUserAnnotatedMedication */
// An interface definition for the [HKUserAnnotatedMedication] class.
type IHKUserAnnotatedMedication interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKUserAnnotatedMedication */
	// properties:
	HasSchedule() bool
	IsArchived() bool
	Medication() IHKMedicationConcept
	Nickname() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKUserAnnotatedMedication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKUserAnnotatedMedication */
// Alloc allocates a new instance without initialization.
func (hc _HKUserAnnotatedMedicationClass) Alloc() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKUserAnnotatedMedicationClass) New() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUserAnnotatedMedication) Init() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUserAnnotatedMedication) Autorelease() HKUserAnnotatedMedication {
	rv := objc.Send[HKUserAnnotatedMedication](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUserAnnotatedMedication creates a new HKUserAnnotatedMedication instance.
func NewHKUserAnnotatedMedication() HKUserAnnotatedMedication {
	return getHKUserAnnotatedMedicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKUserAnnotatedMedication */
// A reference to the tracked medication and the details a person can customize.
//
// The details are relevant to the medication tracking experience.


// A reference to the tracked medication and the details a person can customize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication
type HKUserAnnotatedMedication struct {
	objectivec.Object
}

// HKUserAnnotatedMedicationFrom constructs a [HKUserAnnotatedMedication] from an unsafe.Pointer.
//
// A reference to the tracked medication and the details a person can customize.
func HKUserAnnotatedMedicationFrom(ptr unsafe.Pointer) HKUserAnnotatedMedication {
	return HKUserAnnotatedMedication{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKUserAnnotatedMedication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKUserAnnotatedMedication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKUserAnnotatedMedication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKUserAnnotatedMedication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKUserAnnotatedMedication */

// A Boolean value that indicates whether a medication has a schedule set up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/hasSchedule
func (h_ HKUserAnnotatedMedication) HasSchedule() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("hasSchedule"))
	return rv
}/* debug [instance_properties/getter]: hasSchedule */


// A Boolean value that indicates whether a medication is archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/isArchived
func (h_ HKUserAnnotatedMedication) IsArchived() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isArchived"))
	return rv
}/* debug [instance_properties/getter]: isArchived */


// A reference to the specific medication a person is tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/medication
func (h_ HKUserAnnotatedMedication) Medication() IHKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("medication"))
	return rv
}/* debug [instance_properties/getter]: medication */


// The nickname that a person added to a medication during the entry experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUserAnnotatedMedication/nickname
func (h_ HKUserAnnotatedMedication) Nickname() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("nickname"))
	return rv
}/* debug [instance_properties/getter]: nickname */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKUserAnnotatedMedication */



