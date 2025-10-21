// Code generated from Apple documentation for DataDetection. DO NOT EDIT.

package datadetection

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DDMatchCalendarEvent] class.
var (
	DDMatchCalendarEventClass     _DDMatchCalendarEventClass
	DDMatchCalendarEventClassOnce sync.Once
)

func getDDMatchCalendarEventClass() _DDMatchCalendarEventClass {
	DDMatchCalendarEventClassOnce.Do(func() {
		DDMatchCalendarEventClass = _DDMatchCalendarEventClass{objc.GetClass("DDMatchCalendarEvent")}
	})
	return DDMatchCalendarEventClass
}

type _DDMatchCalendarEventClass struct {
	class objc.Class
}

// An interface definition for the [DDMatchCalendarEvent] class.
type IDDMatchCalendarEvent interface {
	IDDMatch
}

// An object that represents a calendar date or date range that the data detection system matches.
//
// The DataDetection framework returns a calendar event match in a object, which has only a beginning date, only an end date, or both a beginning date and an end date.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent
type DDMatchCalendarEvent struct {
	DDMatch
}

// DDMatchCalendarEventFrom constructs a [DDMatchCalendarEvent] from an unsafe.Pointer.
//
// An object that represents a calendar date or date range that the data detection system matches.
func DDMatchCalendarEventFrom(ptr unsafe.Pointer) DDMatchCalendarEvent {
	return DDMatchCalendarEvent{
		DDMatch: DDMatchFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DDMatchCalendarEventClass) Alloc() DDMatchCalendarEvent {
	rv := objc.Send[DDMatchCalendarEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DDMatchCalendarEventClass) New() DDMatchCalendarEvent {
	rv := objc.Send[DDMatchCalendarEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DDMatchCalendarEvent) Init() DDMatchCalendarEvent {
	rv := objc.Send[DDMatchCalendarEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DDMatchCalendarEvent) Autorelease() DDMatchCalendarEvent {
	rv := objc.Send[DDMatchCalendarEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDDMatchCalendarEvent creates a new DDMatchCalendarEvent instance.
func NewDDMatchCalendarEvent() DDMatchCalendarEvent {
	return getDDMatchCalendarEventClass().New()
}


// A date that represents the end of the event.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent/endDate
func (d_ DDMatchCalendarEvent) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("endDate"))
	return rv
}

// The time zone for the event’s end date.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent/endTimeZone
func (d_ DDMatchCalendarEvent) EndTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("endTimeZone"))
	return rv
}

// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent/isAllDay
func (d_ DDMatchCalendarEvent) AllDay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allDay"))
	return rv
}

// A date that represents the start of the event.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent/startDate
func (d_ DDMatchCalendarEvent) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("startDate"))
	return rv
}

// The time zone for the event’s start date.
//
// [Full Topic]: https://developer.apple.com/documentation/DataDetection/DDMatchCalendarEvent/startTimeZone
func (d_ DDMatchCalendarEvent) StartTimeZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("startTimeZone"))
	return rv
}

// A Boolean value that indicates whether the event is an all-day event.
//
// [Full Topic]: https://developer.apple.com/documentation/datadetection/ddmatchcalendarevent/isallday
func (d_ DDMatchCalendarEvent) IsAllDay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isAllDay"))
	return rv
}


// SetIsAllDay sets the value of the isAllDay property.
// A Boolean value that indicates whether the event is an all-day event.

//
// [Full Topic]: https://developer.apple.com/documentation/datadetection/ddmatchcalendarevent/isallday
func (d_ DDMatchCalendarEvent) SetIsAllDay(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsAllDay:"), value)
}



