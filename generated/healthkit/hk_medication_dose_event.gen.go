// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class HKMedicationDoseEvent */


/* debug [class_header]: Header for HKMedicationDoseEvent */
// The class instance for the [HKMedicationDoseEvent] class.
var (
	HKMedicationDoseEventClass     _HKMedicationDoseEventClass
	HKMedicationDoseEventClassOnce sync.Once
)

func getHKMedicationDoseEventClass() _HKMedicationDoseEventClass {
	HKMedicationDoseEventClassOnce.Do(func() {
		HKMedicationDoseEventClass = _HKMedicationDoseEventClass{objc.GetClass("HKMedicationDoseEvent")}
	})
	return HKMedicationDoseEventClass
}

type _HKMedicationDoseEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKMedicationDoseEvent */
// An interface definition for the [HKMedicationDoseEvent] class.
type IHKMedicationDoseEvent interface {
	IHKSample
	
/* debug [class_interface_properties]: Properties for HKMedicationDoseEvent */
	// properties:
	DoseQuantity() objc.IObject /* cross-framework: NSNumber */
	LogStatus() HKMedicationDoseEventLogStatus
	MedicationConceptIdentifier() IHKHealthConceptIdentifier
	MedicationDoseEventType() IHKMedicationDoseEventType
	ScheduledDate() objc.IObject /* cross-framework: NSDate */
	ScheduledDoseQuantity() objc.IObject /* cross-framework: NSNumber */
	ScheduleType() HKMedicationDoseEventScheduleType
	Unit() IHKUnit
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKMedicationDoseEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKMedicationDoseEvent */
// Alloc allocates a new instance without initialization.
func (hc _HKMedicationDoseEventClass) Alloc() HKMedicationDoseEvent {
	rv := objc.Send[HKMedicationDoseEvent](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HKMedicationDoseEventClass) New() HKMedicationDoseEvent {
	rv := objc.Send[HKMedicationDoseEvent](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKMedicationDoseEvent) Init() HKMedicationDoseEvent {
	rv := objc.Send[HKMedicationDoseEvent](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKMedicationDoseEvent) Autorelease() HKMedicationDoseEvent {
	rv := objc.Send[HKMedicationDoseEvent](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKMedicationDoseEvent creates a new HKMedicationDoseEvent instance.
func NewHKMedicationDoseEvent() HKMedicationDoseEvent {
	return getHKMedicationDoseEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKMedicationDoseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent
type HKMedicationDoseEvent struct {
	HKSample
}

// HKMedicationDoseEventFrom constructs a [HKMedicationDoseEvent] from an unsafe.Pointer.
func HKMedicationDoseEventFrom(ptr unsafe.Pointer) HKMedicationDoseEvent {
	return HKMedicationDoseEvent{
		HKSample: HKSampleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKMedicationDoseEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKMedicationDoseEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKMedicationDoseEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKMedicationDoseEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKMedicationDoseEvent */

// The dose quantity the person reports as taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/doseQuantity-52hxr
func (h_ HKMedicationDoseEvent) DoseQuantity() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](h_.ID, objc.Sel("doseQuantity"))
	return rv
}/* debug [instance_properties/getter]: doseQuantity */


// The log status the system assigns to this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/logStatus-swift.property
func (h_ HKMedicationDoseEvent) LogStatus() HKMedicationDoseEventLogStatus {
	rv := objc.Send[HKMedicationDoseEventLogStatus](h_.ID, objc.Sel("logStatus"))
	return rv
}/* debug [instance_properties/getter]: logStatus */


// The identifier of the medication concept the system associates with this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/medicationConceptIdentifier
func (h_ HKMedicationDoseEvent) MedicationConceptIdentifier() IHKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("medicationConceptIdentifier"))
	return rv
}/* debug [instance_properties/getter]: medicationConceptIdentifier */


// The data type that identified the samples that store medication dose event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/medicationDoseEventType
func (h_ HKMedicationDoseEvent) MedicationDoseEventType() IHKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("medicationDoseEventType"))
	return rv
}/* debug [instance_properties/getter]: medicationDoseEventType */


// The date and time the person takes the medication, if scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduledDate
func (h_ HKMedicationDoseEvent) ScheduledDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("scheduledDate"))
	return rv
}/* debug [instance_properties/getter]: scheduledDate */


// The dose quantity a person is expected to take based on their medication schedule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduledDoseQuantity-7ffhr
func (h_ HKMedicationDoseEvent) ScheduledDoseQuantity() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](h_.ID, objc.Sel("scheduledDoseQuantity"))
	return rv
}/* debug [instance_properties/getter]: scheduledDoseQuantity */


// The scheduling context for this logged dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduleType-swift.property
func (h_ HKMedicationDoseEvent) ScheduleType() HKMedicationDoseEventScheduleType {
	rv := objc.Send[HKMedicationDoseEventScheduleType](h_.ID, objc.Sel("scheduleType"))
	return rv
}/* debug [instance_properties/getter]: scheduleType */


// The unit that the system associates with the medication when the person logs the dose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/unit
func (h_ HKMedicationDoseEvent) Unit() IHKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unit"))
	return rv
}/* debug [instance_properties/getter]: unit */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKMedicationDoseEvent */



