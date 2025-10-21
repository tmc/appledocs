// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EKCalendarItem] class.
var (
	EKCalendarItemClass     _EKCalendarItemClass
	EKCalendarItemClassOnce sync.Once
)

func getEKCalendarItemClass() _EKCalendarItemClass {
	EKCalendarItemClassOnce.Do(func() {
		EKCalendarItemClass = _EKCalendarItemClass{objc.GetClass("EKCalendarItem")}
	})
	return EKCalendarItemClass
}

type _EKCalendarItemClass struct {
	class objc.Class
}

// An interface definition for the [EKCalendarItem] class.
type IEKCalendarItem interface {
	IEKObject
	AddAlarm(alarm unsafe.Pointer)
	AddRecurrenceRule(rule unsafe.Pointer)
	RemoveAlarm(alarm unsafe.Pointer)
	RemoveRecurrenceRule(rule unsafe.Pointer)
}

// An abstract superclass for calendar events and reminders.
//
// The is a an abstract superclass for calendar events and reminders. This class provides common properties and methods for accessing properties of calendar items such as the ability to set the calendar, title, and location, as well as support for attaching notes, displaying attendees, setting multiple alarms, and specifying recurrence rules.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem
type EKCalendarItem struct {
	EKObject
}

// EKCalendarItemFrom constructs a [EKCalendarItem] from an unsafe.Pointer.
//
// An abstract superclass for calendar events and reminders.
func EKCalendarItemFrom(ptr unsafe.Pointer) EKCalendarItem {
	return EKCalendarItem{
		EKObject: EKObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EKCalendarItemClass) Alloc() EKCalendarItem {
	rv := objc.Send[EKCalendarItem](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKCalendarItemClass) New() EKCalendarItem {
	rv := objc.Send[EKCalendarItem](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKCalendarItem) Init() EKCalendarItem {
	rv := objc.Send[EKCalendarItem](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKCalendarItem) Autorelease() EKCalendarItem {
	rv := objc.Send[EKCalendarItem](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKCalendarItem creates a new EKCalendarItem instance.
func NewEKCalendarItem() EKCalendarItem {
	return getEKCalendarItemClass().New()
}


// Adds an alarm to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addAlarm(_:)
func (e_ EKCalendarItem) AddAlarm(alarm unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addAlarm:"), alarm)
}

// Adds a recurrence rule to the recurrence rule array.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addRecurrenceRule(_:)
func (e_ EKCalendarItem) AddRecurrenceRule(rule unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addRecurrenceRule:"), rule)
}

// Removes an alarm from the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeAlarm(_:)
func (e_ EKCalendarItem) RemoveAlarm(alarm unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeAlarm:"), alarm)
}

// Removes a recurrence rule from the recurrence rule array.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeRecurrenceRule(_:)
func (e_ EKCalendarItem) RemoveRecurrenceRule(rule unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeRecurrenceRule:"), rule)
}

// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) Alarms() []EKAlarm {
	rv := objc.Send[[]EKAlarm](e_.ID, objc.Sel("alarms"))
	return rv
}


// SetAlarms sets the value of the alarms property.
// The alarms associated with the calendar item, as an array of objects.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) SetAlarms(value []EKAlarm) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setAlarms:"), nsArray)
}
// The attendees associated with the calendar item, as an array of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/attendees
func (e_ EKCalendarItem) Attendees() []EKParticipant {
	rv := objc.Send[[]EKParticipant](e_.ID, objc.Sel("attendees"))
	return rv
}

// The calendar for the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) Calendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("calendar"))
	return rv
}


// SetCalendar sets the value of the calendar property.
// The calendar for the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) SetCalendar(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCalendar:"), value)
}
// The calendar item’s external identifier as provided by the calendar server.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemExternalIdentifier
func (e_ EKCalendarItem) CalendarItemExternalIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarItemExternalIdentifier"))
	return rv
}

// The calendar item’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemIdentifier
func (e_ EKCalendarItem) CalendarItemIdentifier() string {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarItemIdentifier"))
	return rv
}

// The date that this calendar item was created.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/creationDate
func (e_ EKCalendarItem) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("creationDate"))
	return rv
}

// A Boolean value that indicates whether the calendar item has alarms.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAlarms
func (e_ EKCalendarItem) HasAlarms() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAlarms"))
	return rv
}

// A Boolean value that indicates whether the calendar item has attendees.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAttendees
func (e_ EKCalendarItem) HasAttendees() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAttendees"))
	return rv
}

// A Boolean value that indicates whether the calendar item has notes.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasNotes
func (e_ EKCalendarItem) HasNotes() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasNotes"))
	return rv
}

// A Boolean value that indicates whether the calendar item has recurrence rules.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasRecurrenceRules
func (e_ EKCalendarItem) HasRecurrenceRules() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasRecurrenceRules"))
	return rv
}

// The date that the calendar item was last modified.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/lastModifiedDate
func (e_ EKCalendarItem) LastModifiedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("lastModifiedDate"))
	return rv
}

// The location associated with the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) Location() string {
	rv := objc.Send[string](e_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The location associated with the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) SetLocation(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocation:"), objc.String(value))
}
// The notes associated with the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) Notes() string {
	rv := objc.Send[string](e_.ID, objc.Sel("notes"))
	return rv
}


// SetNotes sets the value of the notes property.
// The notes associated with the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) SetNotes(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNotes:"), objc.String(value))
}
// The recurrence rules for the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) RecurrenceRules() []EKRecurrenceRule {
	rv := objc.Send[[]EKRecurrenceRule](e_.ID, objc.Sel("recurrenceRules"))
	return rv
}


// SetRecurrenceRules sets the value of the recurrenceRules property.
// The recurrence rules for the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) SetRecurrenceRules(value []EKRecurrenceRule) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](e_.ID, objc.Sel("setRecurrenceRules:"), nsArray)
}
// The time zone for the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) TimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone for the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) SetTimeZone(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTimeZone:"), value)
}
// The title for the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) Title() string {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title for the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) SetTitle(value string) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The URL for the calendar item.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL for the calendar item.

//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) SetURL(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setURL:"), value)
}
// The calendar item’s unique identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/uuid
func (e_ EKCalendarItem) UUID() string {
	rv := objc.Send[string](e_.ID, objc.Sel("UUID"))
	return rv
}



