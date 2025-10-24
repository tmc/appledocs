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
	Alarms() []IEKAlarm
	SetAlarms(value []IEKAlarm)
	Attendees() []IEKParticipant
	Calendar() IEKCalendar
	SetCalendar(value IEKCalendar)
	CalendarItemExternalIdentifier() objc.IObject /* cross-framework: NSString */
	CalendarItemIdentifier() objc.IObject /* cross-framework: NSString */
	CreationDate() objc.IObject /* cross-framework: NSDate */
	HasAlarms() bool
	HasAttendees() bool
	HasNotes() bool
	HasRecurrenceRules() bool
	LastModifiedDate() objc.IObject /* cross-framework: NSDate */
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	Notes() objc.IObject /* cross-framework: NSString */
	SetNotes(value objc.IObject /* cross-framework: NSString */)
	RecurrenceRules() []IEKRecurrenceRule
	SetRecurrenceRules(value []IEKRecurrenceRule)
	TimeZone() objc.IObject /* cross-framework: TimeZone */
	SetTimeZone(value objc.IObject /* cross-framework: TimeZone */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
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
func (e_ EKCalendarItem) Alarms() []IEKAlarm {
	rv := objc.Send[[]EKAlarm](e_.ID, objc.Sel("alarms"))
	return rv
}


// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) SetAlarms(value []IEKAlarm) {
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
func (e_ EKCalendarItem) Attendees() []IEKParticipant {
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
func (e_ EKCalendarItem) CalendarItemExternalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarItemExternalIdentifier"))
	return rv
}


// The calendar item’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemIdentifier
func (e_ EKCalendarItem) CalendarItemIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarItemIdentifier"))
	return rv
}


// The date that this calendar item was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/creationDate
func (e_ EKCalendarItem) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("creationDate"))
	return rv
}


// A Boolean value that indicates whether the calendar item has alarms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAlarms
func (e_ EKCalendarItem) HasAlarms() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAlarms"))
	return rv
}


// A Boolean value that indicates whether the calendar item has attendees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAttendees
func (e_ EKCalendarItem) HasAttendees() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAttendees"))
	return rv
}


// A Boolean value that indicates whether the calendar item has notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasNotes
func (e_ EKCalendarItem) HasNotes() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasNotes"))
	return rv
}


// A Boolean value that indicates whether the calendar item has recurrence rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasRecurrenceRules
func (e_ EKCalendarItem) HasRecurrenceRules() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasRecurrenceRules"))
	return rv
}


// The date that the calendar item was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/lastModifiedDate
func (e_ EKCalendarItem) LastModifiedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("lastModifiedDate"))
	return rv
}


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("location"))
	return rv
}


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocation:"), value)
}


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) Notes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("notes"))
	return rv
}


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) SetNotes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNotes:"), value)
}


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) RecurrenceRules() []IEKRecurrenceRule {
	rv := objc.Send[[]EKRecurrenceRule](e_.ID, objc.Sel("recurrenceRules"))
	return rv
}


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) SetRecurrenceRules(value []IEKRecurrenceRule) {
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
	rv := objc.Send[foundation.TimeZone](e_.ID, objc.Sel("timeZone"))
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
func (e_ EKCalendarItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}


// The title for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), value)
}


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](e_.ID, objc.Sel("URL"))
	return rv
}


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setURL:"), value)
}


