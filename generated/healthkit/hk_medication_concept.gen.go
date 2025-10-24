// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKMedicationConcept */


/* debug [class_header]: Header for HKMedicationConcept */
// The class instance for the [HKMedicationConcept] class.
var (
	HKMedicationConceptClass     _HKMedicationConceptClass
	HKMedicationConceptClassOnce sync.Once
)

func getHKMedicationConceptClass() _HKMedicationConceptClass {
	HKMedicationConceptClassOnce.Do(func() {
		HKMedicationConceptClass = _HKMedicationConceptClass{objc.GetClass("HKMedicationConcept")}
	})
	return HKMedicationConceptClass
}

type _HKMedicationConceptClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKMedicationConcept */
// An interface definition for the [HKMedicationConcept] class.
type IHKMedicationConcept interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKMedicationConcept */
	// properties:
	DisplayText() objc.IObject /* cross-framework: NSString */
	GeneralForm() HKMedicationGeneralForm /* typedef */
	Identifier() IHKHealthConceptIdentifier
	RelatedCodings() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKMedicationConcept */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKMedicationConcept */
// Alloc allocates a new instance without initialization.
func (hc _HKMedicationConceptClass) Alloc() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKMedicationConceptClass) New() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKMedicationConcept) Init() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKMedicationConcept) Autorelease() HKMedicationConcept {
	rv := objc.Send[HKMedicationConcept](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKMedicationConcept creates a new HKMedicationConcept instance.
func NewHKMedicationConcept() HKMedicationConcept {
	return getHKMedicationConceptClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKMedicationConcept */
// An object that describes a specific medication concept.
//
// A medication concept represents the idea of a medication, like ibuprofen or insulin. It can have clinical significance, or can be created by the person using your app.


// An object that describes a specific medication concept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept
type HKMedicationConcept struct {
	objectivec.Object
}

// HKMedicationConceptFrom constructs a [HKMedicationConcept] from an unsafe.Pointer.
//
// An object that describes a specific medication concept.
func HKMedicationConceptFrom(ptr unsafe.Pointer) HKMedicationConcept {
	return HKMedicationConcept{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKMedicationConcept *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKMedicationConcept */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKMedicationConcept */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKMedicationConcept */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKMedicationConcept */

// The display name for this medication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/displayText
func (h_ HKMedicationConcept) DisplayText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("displayText"))
	return rv
}/* debug [instance_properties/getter]: displayText */


// The general form the medication is manufactured in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/generalForm
func (h_ HKMedicationConcept) GeneralForm() HKMedicationGeneralForm /* typedef */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("generalForm"))
	return rv
}/* debug [instance_properties/getter]: generalForm */


// The unique identifier for the specific medication concept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/identifier
func (h_ HKMedicationConcept) Identifier() IHKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The set of related clinical codings for the medication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationConcept/relatedCodings
func (h_ HKMedicationConcept) RelatedCodings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("relatedCodings"))
	return rv
}/* debug [instance_properties/getter]: relatedCodings */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKMedicationConcept */



