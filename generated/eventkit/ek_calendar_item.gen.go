// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKCalendarItem */


/* debug [class_header]: Header for EKCalendarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKCalendarItem */
// An interface definition for the [EKCalendarItem] class.
type IEKCalendarItem interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKCalendarItem */
	// properties:
	Alarms() []EKAlarm
	SetAlarms(value []EKAlarm)
	Attendees() []EKParticipant
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
	RecurrenceRules() []EKRecurrenceRule
	SetRecurrenceRules(value []EKRecurrenceRule)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKCalendarItem */
	// methods:
	AddAlarm(alarm IEKAlarm)
	AddRecurrenceRule(rule IEKRecurrenceRule)
	RemoveAlarm(alarm IEKAlarm)
	RemoveRecurrenceRule(rule IEKRecurrenceRule)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKCalendarItem */
// Alloc allocates a new instance without initialization.
func (ec _EKCalendarItemClass) Alloc() EKCalendarItem {
	rv := objc.Send[EKCalendarItem](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKCalendarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKCalendarItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKCalendarItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKCalendarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKCalendarItem */

// Adds an alarm to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addAlarm(_:)
func (e_ EKCalendarItem) AddAlarm(alarm IEKAlarm) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addAlarm:"), alarm)
}/* debug [instance_methods/method]: AddAlarm */


// Adds a recurrence rule to the recurrence rule array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/addRecurrenceRule(_:)
func (e_ EKCalendarItem) AddRecurrenceRule(rule IEKRecurrenceRule) {
	objc.Send[objc.ID](e_.ID, objc.Sel("addRecurrenceRule:"), rule)
}/* debug [instance_methods/method]: AddRecurrenceRule */


// Removes an alarm from the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeAlarm(_:)
func (e_ EKCalendarItem) RemoveAlarm(alarm IEKAlarm) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeAlarm:"), alarm)
}/* debug [instance_methods/method]: RemoveAlarm */


// Removes a recurrence rule from the recurrence rule array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/removeRecurrenceRule(_:)
func (e_ EKCalendarItem) RemoveRecurrenceRule(rule IEKRecurrenceRule) {
	objc.Send[objc.ID](e_.ID, objc.Sel("removeRecurrenceRule:"), rule)
}/* debug [instance_methods/method]: RemoveRecurrenceRule */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKCalendarItem */

// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) Alarms() []EKAlarm {
	rv := objc.Send[[]EKAlarm](e_.ID, objc.Sel("alarms"))
	return rv
}/* debug [instance_properties/getter]: alarms */


// The alarms associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/alarms
func (e_ EKCalendarItem) SetAlarms(value []EKAlarm) {
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
}/* debug [instance_properties/setter]: alarms */


// The attendees associated with the calendar item, as an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/attendees
func (e_ EKCalendarItem) Attendees() []EKParticipant {
	rv := objc.Send[[]EKParticipant](e_.ID, objc.Sel("attendees"))
	return rv
}/* debug [instance_properties/getter]: attendees */


// The calendar for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) Calendar() IEKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("calendar"))
	return rv
}/* debug [instance_properties/getter]: calendar */


// The calendar for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendar
func (e_ EKCalendarItem) SetCalendar(value IEKCalendar) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCalendar:"), value)
}/* debug [instance_properties/setter]: calendar */


// The calendar item’s external identifier as provided by the calendar server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemExternalIdentifier
func (e_ EKCalendarItem) CalendarItemExternalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarItemExternalIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarItemExternalIdentifier */


// The calendar item’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/calendarItemIdentifier
func (e_ EKCalendarItem) CalendarItemIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarItemIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarItemIdentifier */


// The date that this calendar item was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/creationDate
func (e_ EKCalendarItem) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// A Boolean value that indicates whether the calendar item has alarms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAlarms
func (e_ EKCalendarItem) HasAlarms() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAlarms"))
	return rv
}/* debug [instance_properties/getter]: hasAlarms */


// A Boolean value that indicates whether the calendar item has attendees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasAttendees
func (e_ EKCalendarItem) HasAttendees() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasAttendees"))
	return rv
}/* debug [instance_properties/getter]: hasAttendees */


// A Boolean value that indicates whether the calendar item has notes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasNotes
func (e_ EKCalendarItem) HasNotes() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasNotes"))
	return rv
}/* debug [instance_properties/getter]: hasNotes */


// A Boolean value that indicates whether the calendar item has recurrence rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/hasRecurrenceRules
func (e_ EKCalendarItem) HasRecurrenceRules() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("hasRecurrenceRules"))
	return rv
}/* debug [instance_properties/getter]: hasRecurrenceRules */


// The date that the calendar item was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/lastModifiedDate
func (e_ EKCalendarItem) LastModifiedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](e_.ID, objc.Sel("lastModifiedDate"))
	return rv
}/* debug [instance_properties/getter]: lastModifiedDate */


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// The location associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/location
func (e_ EKCalendarItem) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setLocation:"), value)
}/* debug [instance_properties/setter]: location */


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) Notes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("notes"))
	return rv
}/* debug [instance_properties/getter]: notes */


// The notes associated with the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/notes
func (e_ EKCalendarItem) SetNotes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNotes:"), value)
}/* debug [instance_properties/setter]: notes */


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) RecurrenceRules() []EKRecurrenceRule {
	rv := objc.Send[[]EKRecurrenceRule](e_.ID, objc.Sel("recurrenceRules"))
	return rv
}/* debug [instance_properties/getter]: recurrenceRules */


// The recurrence rules for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/recurrenceRules
func (e_ EKCalendarItem) SetRecurrenceRules(value []EKRecurrenceRule) {
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
}/* debug [instance_properties/setter]: recurrenceRules */


// The time zone for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](e_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/timeZone
func (e_ EKCalendarItem) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */


// The title for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/title
func (e_ EKCalendarItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](e_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The URL for the calendar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendarItem/url
func (e_ EKCalendarItem) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKCalendarItem */


