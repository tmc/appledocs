// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Alarms() []EKAlarm /* primitive/slice/pointer. */
	SetAlarms(value []EKAlarm /* primitive/slice/pointer. */)
	Attendees() []EKParticipant /* primitive/slice/pointer. */
	Calendar() IEKCalendar
	SetCalendar(value IEKCalendar)
	CalendarItemExternalIdentifier() string /* primitive/slice/pointer. */
	CalendarItemIdentifier() string /* primitive/slice/pointer. */
	CreationDate() foundation.objc.IObject /* cross-framework: NSDate */
	HasAlarms() bool /* primitive/slice/pointer. */
	HasAttendees() bool /* primitive/slice/pointer. */
	HasNotes() bool /* primitive/slice/pointer. */
	HasRecurrenceRules() bool /* primitive/slice/pointer. */
	LastModifiedDate() foundation.objc.IObject /* cross-framework: NSDate */
	Location() string /* primitive/slice/pointer. */
	SetLocation(value string /* primitive/slice/pointer. */)
	Notes() string /* primitive/slice/pointer. */
	SetNotes(value string /* primitive/slice/pointer. */)
	RecurrenceRules() []EKRecurrenceRule /* primitive/slice/pointer. */
	SetRecurrenceRules(value []EKRecurrenceRule /* primitive/slice/pointer. */)
	TimeZone() objc.IObject /* cross-framework: TimeZone */
	SetTimeZone(value objc.IObject /* cross-framework: TimeZone */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	URL() foundation.objc.IObject /* cross-framework: URL */
	SetURL(value foundation.objc.IObject /* cross-framework: URL */)
	UUID() string /* primitive/slice/pointer. */
	// methods:
	AddAlarm(alarm IEKAlarm)
	AddRecurrenceRule(rule IEKRecurrenceRule)
	RemoveAlarm(alarm IEKAlarm)
	RemoveRecurrenceRule(rule IEKRecurrenceRule)
}

// An abstract superclass for calendar events and reminders.
//
// The is a an abstract superclass for calendar events and reminders. This class provides common properties and methods for accessing properties of calendar items such as the ability to set the calendar, title, and location, as well as support for attaching notes, displaying attendees, setting multiple alarms, and specifying recurrence rules.


// An abstract superclass for calendar events and reminders.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addAlarm(_:)
func (e_ EKCalendarItem) AddAlarm(alarm IEKAlarm) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addAlarm:"), alarm)
}


// Adds a recurrence rule to the recurrence rule array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addRecurrenceRule(_:)
func (e_ EKCalendarItem) AddRecurrenceRule(rule IEKRecurrenceRule) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addRecurrenceRule:"), rule)
}


// Removes an alarm from the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeAlarm(_:)
func (e_ EKCalendarItem) RemoveAlarm(alarm IEKAlarm) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeAlarm:"), alarm)
}


// Removes a recurrence rule from the recurrence rule array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeRecurrenceRule(_:)
func (e_ EKCalendarItem) RemoveRecurrenceRule(rule IEKRecurrenceRule) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeRecurrenceRule:"), rule)
}


// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) Alarms() []EKAlarm /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKAlarm](e_.ID, objc.Sel("alarms"))
	return rv
}


// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) SetAlarms(value []EKAlarm /* primitive/slice/pointer. */) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/attendees
func (e_ EKCalendarItem) Attendees() []EKParticipant /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKParticipant](e_.ID, objc.Sel("attendees"))
	return rv
}


// The calendar for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) Calendar() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) SetCalendar(value IEKCalendar) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCalendar:"), value)
}


// The calendar item’s external identifier as provided by the calendar server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemExternalIdentifier
func (e_ EKCalendarItem) CalendarItemExternalIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarItemExternalIdentifier"))
	return rv
}


// The calendar item’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemIdentifier
func (e_ EKCalendarItem) CalendarItemIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarItemIdentifier"))
	return rv
}


// The date that this calendar item was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/creationDate
func (e_ EKCalendarItem) CreationDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("creationDate"))
	return rv
}


// A Boolean value that indicates whether the calendar item has alarms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAlarms
func (e_ EKCalendarItem) HasAlarms() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAlarms"))
	return rv
}


// A Boolean value that indicates whether the calendar item has attendees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAttendees
func (e_ EKCalendarItem) HasAttendees() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAttendees"))
	return rv
}


// A Boolean value that indicates whether the calendar item has notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasNotes
func (e_ EKCalendarItem) HasNotes() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasNotes"))
	return rv
}


// A Boolean value that indicates whether the calendar item has recurrence rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasRecurrenceRules
func (e_ EKCalendarItem) HasRecurrenceRules() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasRecurrenceRules"))
	return rv
}


// The date that the calendar item was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/lastModifiedDate
func (e_ EKCalendarItem) LastModifiedDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("lastModifiedDate"))
	return rv
}


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) Location() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("location"))
	return rv
}


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) SetLocation(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocation:"), objc.String(value))
}


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) Notes() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("notes"))
	return rv
}


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) SetNotes(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNotes:"), objc.String(value))
}


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) RecurrenceRules() []EKRecurrenceRule /* primitive/slice/pointer. */ {
	rv := objc.Send[[]EKRecurrenceRule](e_.ID, objc.Sel("recurrenceRules"))
	return rv
}


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) SetRecurrenceRules(value []EKRecurrenceRule /* primitive/slice/pointer. */) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) TimeZone() objc.IObject /* cross-framework: TimeZone */ {
	rv := objc.Send[TimeZone](e_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) SetTimeZone(value objc.IObject /* cross-framework: TimeZone */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTimeZone:"), value)
}


// The title for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}


// The title for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) URL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](e_.ID, objc.Sel("URL"))
	return rv
}


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) SetURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setURL:"), value)
}


// The calendar item’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/uuid
func (e_ EKCalendarItem) UUID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("UUID"))
	return rv
}



