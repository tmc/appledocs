// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKEvent */


/* debug [class_header]: Header for EKEvent */
// The class instance for the [EKEvent] class.
var (
	EKEventClass     _EKEventClass
	EKEventClassOnce sync.Once
)

func getEKEventClass() _EKEventClass {
	EKEventClassOnce.Do(func() {
		EKEventClass = _EKEventClass{objc.GetClass("EKEvent")}
	})
	return EKEventClass
}

type _EKEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKEvent */
// An interface definition for the [EKEvent] class.
type IEKEvent interface {
	IEKCalendarItem
	
/* debug [class_interface_properties]: Properties for EKEvent */
	// properties:
	Availability() EKEventAvailability
	SetAvailability(value EKEventAvailability)
	BirthdayContactIdentifier() objc.IObject /* cross-framework: NSString */
	BirthdayPersonUniqueID() objc.IObject /* cross-framework: NSString */
	EndDate() objc.IObject /* cross-framework: NSDate */
	SetEndDate(value objc.IObject /* cross-framework: NSDate */)
	EventIdentifier() objc.IObject /* cross-framework: NSString */
	AllDay() bool
	SetAllDay(value bool)
	IsDetached() bool
	OccurrenceDate() objc.IObject /* cross-framework: NSDate */
	Organizer() IEKParticipant
	StartDate() objc.IObject /* cross-framework: NSDate */
	SetStartDate(value objc.IObject /* cross-framework: NSDate */)
	Status() EKEventStatus
	StructuredLocation() IEKStructuredLocation
	SetStructuredLocation(value IEKStructuredLocation)
	IsAllDay() bool
	SetIsAllDay(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKEvent */
	// methods:
	CompareStartDateWithEvent(other IEKEvent) ComparisonResult /* not a class type */
	Refresh() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKEvent */
// Alloc allocates a new instance without initialization.
func (ec _EKEventClass) Alloc() EKEvent {
	rv := objc.Send[EKEvent](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKEventClass) New() EKEvent {
	rv := objc.Send[EKEvent](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKEvent) Init() EKEvent {
	rv := objc.Send[EKEvent](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKEvent) Autorelease() EKEvent {
	rv := objc.Send[EKEvent](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKEvent creates a new EKEvent instance.
func NewEKEvent() EKEvent {
	return getEKEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKEvent */
// A class that represents an event in a calendar.
//
// Use the method to create a new event. Use the properties in the class to get and modify certain information about an event. Other properties, such as the event’s title and calendar, are inherited from the parent class .


// A class that represents an event in a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent
type EKEvent struct {
	EKCalendarItem
}

// EKEventFrom constructs a [EKEvent] from an unsafe.Pointer.
//
// A class that represents an event in a calendar.
func EKEventFrom(ptr unsafe.Pointer) EKEvent {
	return EKEvent{
		EKCalendarItem: EKCalendarItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKEvent */

// Creates and returns a new event belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/init(eventStore:)
func NewEKEventWithEventStore(eventStore IEKEventStore) EKEvent {
	rv := objc.Send[EKEvent](objc.ID(getEKEventClass().class), objc.Sel("eventWithEventStore:"), eventStore)
	return rv
}/* debug [class_init_methods/constructor]: NewEKEventWithEventStore */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKEvent */

// Creates and returns a new event belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/init(eventStore:)
func (ec _EKEventClass) EventWithEventStore(eventStore IEKEventStore) EKEvent {
	rv := objc.Send[EKEvent](objc.ID(ec.class), objc.Sel("eventWithEventStore:"), eventStore)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EventWithEventStore) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKEvent */

// Compares the start date of the receiving event with the start date of another event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/compareStartDate(with:)
func (e_ EKEvent) CompareStartDateWithEvent(other IEKEvent) ComparisonResult /* not a class type */ {
	rv := objc.Send[ComparisonResult](e_.ID, objc.Sel("compareStartDateWithEvent:"), other)
	return rv
}/* debug [instance_methods/method]: CompareStartDateWithEvent */


// Updates the event’s data with the current information in the Calendar database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/refresh()
func (e_ EKEvent) Refresh() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("refresh"))
	return rv
}/* debug [instance_methods/method]: Refresh */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKEvent */

// The availability setting for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/availability
func (e_ EKEvent) Availability() EKEventAvailability {
	rv := objc.Send[EKEventAvailability](e_.ID, objc.Sel("availability"))
	return rv
}/* debug [instance_properties/getter]: availability */


// The availability setting for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/availability
func (e_ EKEvent) SetAvailability(value EKEventAvailability) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAvailability:"), value)
}/* debug [instance_properties/setter]: availability */


// The contact identifier of the person for this birthday event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayContactIdentifier
func (e_ EKEvent) BirthdayContactIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("birthdayContactIdentifier"))
	return rv
}/* debug [instance_properties/getter]: birthdayContactIdentifier */


// The Address Book framework record identifier of the person for this birthday event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayPersonUniqueID
func (e_ EKEvent) BirthdayPersonUniqueID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("birthdayPersonUniqueID"))
	return rv
}/* debug [instance_properties/getter]: birthdayPersonUniqueID */


// The end date for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/endDate
func (e_ EKEvent) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The end date for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/endDate
func (e_ EKEvent) SetEndDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEndDate:"), value)
}/* debug [instance_properties/setter]: endDate */


// A unique identifier for the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/eventIdentifier
func (e_ EKEvent) EventIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("eventIdentifier"))
	return rv
}/* debug [instance_properties/getter]: eventIdentifier */


// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isAllDay
func (e_ EKEvent) AllDay() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("allDay"))
	return rv
}/* debug [instance_properties/getter]: allDay */


// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isAllDay
func (e_ EKEvent) SetAllDay(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAllDay:"), value)
}/* debug [instance_properties/setter]: allDay */


// A Boolean value that indicates whether an event is a detached instance of a repeating event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isDetached
func (e_ EKEvent) IsDetached() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDetached"))
	return rv
}/* debug [instance_properties/getter]: isDetached */


// The original occurrence date of an event if it is part of a recurring series.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/occurrenceDate
func (e_ EKEvent) OccurrenceDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("occurrenceDate"))
	return rv
}/* debug [instance_properties/getter]: occurrenceDate */


// The organizer associated with the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/organizer
func (e_ EKEvent) Organizer() IEKParticipant {
	rv := objc.Send[EKParticipant](e_.ID, objc.Sel("organizer"))
	return rv
}/* debug [instance_properties/getter]: organizer */


// The start date of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/startDate
func (e_ EKEvent) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The start date of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/startDate
func (e_ EKEvent) SetStartDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStartDate:"), value)
}/* debug [instance_properties/setter]: startDate */


// The status of the event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/status
func (e_ EKEvent) Status() EKEventStatus {
	rv := objc.Send[EKEventStatus](e_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// The event’s location with a potential geocoordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/structuredLocation
func (e_ EKEvent) StructuredLocation() IEKStructuredLocation {
	rv := objc.Send[EKStructuredLocation](e_.ID, objc.Sel("structuredLocation"))
	return rv
}/* debug [instance_properties/getter]: structuredLocation */


// The event’s location with a potential geocoordinate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/structuredLocation
func (e_ EKEvent) SetStructuredLocation(value IEKStructuredLocation) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}/* debug [instance_properties/setter]: structuredLocation */


// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekevent/isallday
func (e_ EKEvent) IsAllDay() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isAllDay"))
	return rv
}/* debug [instance_properties/getter]: isAllDay */


// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekevent/isallday
func (e_ EKEvent) SetIsAllDay(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsAllDay:"), value)
}/* debug [instance_properties/setter]: isAllDay */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKEvent */


