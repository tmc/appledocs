// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [HKMedicationDoseEvent] class.
type IHKMedicationDoseEvent interface {
	IHKSample
}

//
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

// Alloc allocates a new instance without initialization.
func (hc _HKMedicationDoseEventClass) Alloc() HKMedicationDoseEvent {
	rv := objc.Send[HKMedicationDoseEvent](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The dose quantity the person reports as taken.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/doseQuantity-52hxr
func (h_ HKMedicationDoseEvent) DoseQuantity() foundation.Number {
	rv := objc.Send[foundation.Number](h_.ID, objc.Sel("doseQuantity"))
	return rv
}

// The log status the system assigns to this dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/logStatus-swift.property
func (h_ HKMedicationDoseEvent) LogStatus() HKMedicationDoseEventLogStatus {
	rv := objc.Send[HKMedicationDoseEventLogStatus](h_.ID, objc.Sel("logStatus"))
	return rv
}

// The identifier of the medication concept the system associates with this dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/medicationConceptIdentifier
func (h_ HKMedicationDoseEvent) MedicationConceptIdentifier() HKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("medicationConceptIdentifier"))
	return rv
}

// The data type that identified the samples that store medication dose event data.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/medicationDoseEventType
func (h_ HKMedicationDoseEvent) MedicationDoseEventType() HKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("medicationDoseEventType"))
	return rv
}

// The scheduling context for this logged dose event.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduleType-swift.property
func (h_ HKMedicationDoseEvent) ScheduleType() HKMedicationDoseEventScheduleType {
	rv := objc.Send[HKMedicationDoseEventScheduleType](h_.ID, objc.Sel("scheduleType"))
	return rv
}

// The date and time the person takes the medication, if scheduled.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduledDate
func (h_ HKMedicationDoseEvent) ScheduledDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("scheduledDate"))
	return rv
}

// The dose quantity a person is expected to take based on their medication schedule.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/scheduledDoseQuantity-7ffhr
func (h_ HKMedicationDoseEvent) ScheduledDoseQuantity() foundation.Number {
	rv := objc.Send[foundation.Number](h_.ID, objc.Sel("scheduledDoseQuantity"))
	return rv
}

// The unit that the system associates with the medication when the person logs the dose.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMedicationDoseEvent/unit
func (h_ HKMedicationDoseEvent) Unit() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unit"))
	return rv
}



