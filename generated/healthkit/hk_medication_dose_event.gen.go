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
	// properties:
	DoseQuantity() float64 /* primitive/slice/pointer. */
	SetDoseQuantity(value float64 /* primitive/slice/pointer. */)
	LogStatus() unsafe.Pointer
	SetLogStatus(value unsafe.Pointer)
	MedicationConceptIdentifier() IHKHealthConceptIdentifier
	SetMedicationConceptIdentifier(value IHKHealthConceptIdentifier)
	MedicationDoseEventType() IHKMedicationDoseEventType
	SetMedicationDoseEventType(value IHKMedicationDoseEventType)
	ScheduleType() unsafe.Pointer
	SetScheduleType(value unsafe.Pointer)
	ScheduledDate() foundation.objc.IObject /* cross-framework: Date */
	SetScheduledDate(value foundation.objc.IObject /* cross-framework: Date */)
	ScheduledDoseQuantity() float64 /* primitive/slice/pointer. */
	SetScheduledDoseQuantity(value float64 /* primitive/slice/pointer. */)
	Unit() IHKUnit
	SetUnit(value IHKUnit)
	// methods:
}



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



// The quantity of the medication taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/dosequantity-4cb5m
func (h_ HKMedicationDoseEvent) DoseQuantity() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](h_.ID, objc.Sel("doseQuantity"))
	return rv
}


// The quantity of the medication taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/dosequantity-4cb5m
func (h_ HKMedicationDoseEvent) SetDoseQuantity(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setDoseQuantity:"), value)
}


// The log status the system assigns to this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/logstatus-swift.property
func (h_ HKMedicationDoseEvent) LogStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("logStatus"))
	return rv
}


// The log status the system assigns to this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/logstatus-swift.property
func (h_ HKMedicationDoseEvent) SetLogStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setLogStatus:"), value)
}


// The identifier of the medication concept the system associates with this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/medicationconceptidentifier
func (h_ HKMedicationDoseEvent) MedicationConceptIdentifier() IHKHealthConceptIdentifier {
	rv := objc.Send[HKHealthConceptIdentifier](h_.ID, objc.Sel("medicationConceptIdentifier"))
	return rv
}


// The identifier of the medication concept the system associates with this dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/medicationconceptidentifier
func (h_ HKMedicationDoseEvent) SetMedicationConceptIdentifier(value IHKHealthConceptIdentifier) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMedicationConceptIdentifier:"), value)
}


// The data type that identified the samples that store medication dose event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/medicationdoseeventtype
func (h_ HKMedicationDoseEvent) MedicationDoseEventType() IHKMedicationDoseEventType {
	rv := objc.Send[HKMedicationDoseEventType](h_.ID, objc.Sel("medicationDoseEventType"))
	return rv
}


// The data type that identified the samples that store medication dose event data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/medicationdoseeventtype
func (h_ HKMedicationDoseEvent) SetMedicationDoseEventType(value IHKMedicationDoseEventType) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setMedicationDoseEventType:"), value)
}


// The scheduling context for this logged dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduletype-swift.property
func (h_ HKMedicationDoseEvent) ScheduleType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("scheduleType"))
	return rv
}


// The scheduling context for this logged dose event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduletype-swift.property
func (h_ HKMedicationDoseEvent) SetScheduleType(value unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setScheduleType:"), value)
}


// The date and time the person takes the medication, if scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduleddate
func (h_ HKMedicationDoseEvent) ScheduledDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](h_.ID, objc.Sel("scheduledDate"))
	return rv
}


// The date and time the person takes the medication, if scheduled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduleddate
func (h_ HKMedicationDoseEvent) SetScheduledDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setScheduledDate:"), value)
}


// The quantity of the medication scheduled to be taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduleddosequantity-477ge
func (h_ HKMedicationDoseEvent) ScheduledDoseQuantity() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](h_.ID, objc.Sel("scheduledDoseQuantity"))
	return rv
}


// The quantity of the medication scheduled to be taken.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/scheduleddosequantity-477ge
func (h_ HKMedicationDoseEvent) SetScheduledDoseQuantity(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setScheduledDoseQuantity:"), value)
}


// The unit that the system associates with the medication when the person logs the dose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/unit
func (h_ HKMedicationDoseEvent) Unit() IHKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unit"))
	return rv
}


// The unit that the system associates with the medication when the person logs the dose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkmedicationdoseevent/unit
func (h_ HKMedicationDoseEvent) SetUnit(value IHKUnit) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setUnit:"), value)
}



