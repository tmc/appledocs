// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [EKEvent] class.
type IEKEvent interface {
	IEKCalendarItem
	CompareStartDateWithEvent(other unsafe.Pointer) unsafe.Pointer
	Refresh() bool
}

// A class that represents an event in a calendar.
//
// Use the method to create a new event. Use the properties in the class to get and modify certain information about an event. Other properties, such as the event’s title and calendar, are inherited from the parent class .
//
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

// Alloc allocates a new instance without initialization.
func (ec _EKEventClass) Alloc() EKEvent {
	rv := objc.Send[EKEvent](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates and returns a new event belonging to a specified event store.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/init(eventStore:)
func NewEKEventWithEventStore(eventStore unsafe.Pointer) EKEvent {
	rv := objc.Send[EKEvent](objc.ID(getEKEventClass().class), objc.Sel("eventWithEventStore:"), eventStore)
	return rv
}


// Creates and returns a new event belonging to a specified event store.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/init(eventStore:)
func (ec _EKEventClass) EventWithEventStore(eventStore unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("eventWithEventStore:"), eventStore)
	return rv
}

// Compares the start date of the receiving event with the start date of another event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/compareStartDate(with:)
func (e_ EKEvent) CompareStartDateWithEvent(other unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("compareStartDateWithEvent:"), other)
	return rv
}

// Updates the event’s data with the current information in the Calendar database.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/refresh()
func (e_ EKEvent) Refresh() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("refresh"))
	return rv
}

// The availability setting for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/availability
func (e_ EKEvent) Availability() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("availability"))
	return rv
}


// SetAvailability sets the value of the availability property.
// The availability setting for the event.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/availability
func (e_ EKEvent) SetAvailability(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAvailability:"), value)
}

// The contact identifier of the person for this birthday event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayContactIdentifier
func (e_ EKEvent) BirthdayContactIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("birthdayContactIdentifier"))
	return rv
}

// The Address Book framework record identifier of the person for this birthday event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayPersonID
func (e_ EKEvent) BirthdayPersonID() int {
	rv := objc.Send[int](e_.ID, objc.Sel("birthdayPersonID"))
	return rv
}

// The Address Book framework record identifier of the person for this birthday event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/birthdayPersonUniqueID
func (e_ EKEvent) BirthdayPersonUniqueID() string {
	rv := objc.Send[string](e_.ID, objc.Sel("birthdayPersonUniqueID"))
	return rv
}

// The end date for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/endDate
func (e_ EKEvent) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The end date for the event.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/endDate
func (e_ EKEvent) SetEndDate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEndDate:"), value)
}

// A unique identifier for the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/eventIdentifier
func (e_ EKEvent) EventIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("eventIdentifier"))
	return rv
}

// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isAllDay
func (e_ EKEvent) AllDay() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("allDay"))
	return rv
}


// SetAllDay sets the value of the allDay property.
// A Boolean value that indicates whether the event is an all-day event.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isAllDay
func (e_ EKEvent) SetAllDay(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setAllDay:"), value)
}

// A Boolean value that indicates whether an event is a detached instance of a repeating event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/isDetached
func (e_ EKEvent) IsDetached() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDetached"))
	return rv
}

// The original occurrence date of an event if it is part of a recurring series.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/occurrenceDate
func (e_ EKEvent) OccurrenceDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("occurrenceDate"))
	return rv
}

// The organizer associated with the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/organizer
func (e_ EKEvent) Organizer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("organizer"))
	return rv
}

// The start date of the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/startDate
func (e_ EKEvent) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The start date of the event.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/startDate
func (e_ EKEvent) SetStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStartDate:"), value)
}

// The status of the event.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/status
func (e_ EKEvent) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("status"))
	return rv
}

// The event’s location with a potential geocoordinate.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/structuredLocation
func (e_ EKEvent) StructuredLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("structuredLocation"))
	return rv
}


// SetStructuredLocation sets the value of the structuredLocation property.
// The event’s location with a potential geocoordinate.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKEvent/structuredLocation
func (e_ EKEvent) SetStructuredLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setStructuredLocation:"), value)
}


