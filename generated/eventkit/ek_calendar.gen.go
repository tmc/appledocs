// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [EKCalendar] class.
var (
	EKCalendarClass     _EKCalendarClass
	EKCalendarClassOnce sync.Once
)

func getEKCalendarClass() _EKCalendarClass {
	EKCalendarClassOnce.Do(func() {
		EKCalendarClass = _EKCalendarClass{objc.GetClass("EKCalendar")}
	})
	return EKCalendarClass
}

type _EKCalendarClass struct {
	class objc.Class
}

// An interface definition for the [EKCalendar] class.
type IEKCalendar interface {
	IEKObject
	// properties:
	AllowedEntityTypes() unsafe.Pointer
	AllowsContentModifications() bool /* primitive/slice/pointer. */
	CalendarIdentifier() string /* primitive/slice/pointer. */
	CGColor() coregraphics.ColorRef /* not a class type */
	SetCGColor(value coregraphics.ColorRef /* not a class type */)
	Color() objc.IObject /* cross-framework: Color */
	SetColor(value objc.IObject /* cross-framework: Color */)
	Immutable() bool /* primitive/slice/pointer. */
	Subscribed() bool /* primitive/slice/pointer. */
	Source() IEKSource
	SetSource(value IEKSource)
	SupportedEventAvailabilities() EKCalendarEventAvailabilityMask
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	Type() EKCalendarType
	IsImmutable() bool /* primitive/slice/pointer. */
	SetIsImmutable(value bool /* primitive/slice/pointer. */)
	IsSubscribed() bool /* primitive/slice/pointer. */
	SetIsSubscribed(value bool /* primitive/slice/pointer. */)
	// methods:
}

// A class that represents a calendar in EventKit.
//
// Use the properties in this class to get attributes about a calendar, such as its title and type. Use the method to create a calendar object.


// A class that represents a calendar in EventKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar
type EKCalendar struct {
	EKObject
}

// EKCalendarFrom constructs a [EKCalendar] from an unsafe.Pointer.
//
// A class that represents a calendar in EventKit.
func EKCalendarFrom(ptr unsafe.Pointer) EKCalendar {
	return EKCalendar{
		EKObject: EKObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EKCalendarClass) Alloc() EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EKCalendarClass) New() EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKCalendar) Init() EKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKCalendar) Autorelease() EKCalendar {
	rv := objc.Send[EKCalendar](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKCalendar creates a new EKCalendar instance.
func NewEKCalendar() EKCalendar {
	return getEKCalendarClass().New()
}



// Creates a new calendar that can contain the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(for:eventStore:)
func NewEKCalendarForEntityTypeEventStore(entityType EKEntityType, eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(getEKCalendarClass().class), objc.Sel("calendarForEntityType:eventStore:"), entityType, eventStore)
	return rv
}


// Creates and returns a calendar belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(eventStore:)
func NewEKCalendarWithEventStore(eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(getEKCalendarClass().class), objc.Sel("calendarWithEventStore:"), eventStore)
	return rv
}



// Creates and returns a calendar belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(eventStore:)
func (ec _EKCalendarClass) CalendarWithEventStore(eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("calendarWithEventStore:"), eventStore)
	return rv
}


// Creates a new calendar that can contain the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(for:eventStore:)
func (ec _EKCalendarClass) CalendarForEntityTypeEventStore(entityType EKEntityType, eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("calendarForEntityType:eventStore:"), entityType, eventStore)
	return rv
}


// The entity types this calendar can contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/allowedEntityTypes
func (e_ EKCalendar) AllowedEntityTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("allowedEntityTypes"))
	return rv
}


// A Boolean value that indicates whether you can add, edit, and delete items in the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/allowsContentModifications
func (e_ EKCalendar) AllowsContentModifications() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("allowsContentModifications"))
	return rv
}


// A unique identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/calendarIdentifier
func (e_ EKCalendar) CalendarIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("calendarIdentifier"))
	return rv
}


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/cgColor
func (e_ EKCalendar) CGColor() coregraphics.ColorRef /* not a class type */ {
	rv := objc.Send[coregraphics.ColorRef](e_.ID, objc.Sel("CGColor"))
	return rv
}


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/cgColor
func (e_ EKCalendar) SetCGColor(value coregraphics.ColorRef /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCGColor:"), value)
}


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/color
func (e_ EKCalendar) Color() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](e_.ID, objc.Sel("color"))
	return rv
}


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/color
func (e_ EKCalendar) SetColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setColor:"), value)
}


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/isImmutable
func (e_ EKCalendar) Immutable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("immutable"))
	return rv
}


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/isSubscribed
func (e_ EKCalendar) Subscribed() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("subscribed"))
	return rv
}


// The source object representing the account to which this calendar belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/source
func (e_ EKCalendar) Source() IEKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("source"))
	return rv
}


// The source object representing the account to which this calendar belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/source
func (e_ EKCalendar) SetSource(value IEKSource) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSource:"), value)
}


// The event availability settings supported by this calendar, as indicated by a bitmask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/supportedEventAvailabilities
func (e_ EKCalendar) SupportedEventAvailabilities() EKCalendarEventAvailabilityMask {
	rv := objc.Send[EKCalendarEventAvailabilityMask](e_.ID, objc.Sel("supportedEventAvailabilities"))
	return rv
}


// The calendar’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/title
func (e_ EKCalendar) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}


// The calendar’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/title
func (e_ EKCalendar) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The calendar’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/type
func (e_ EKCalendar) Type() EKCalendarType {
	rv := objc.Send[EKCalendarType](e_.ID, objc.Sel("type"))
	return rv
}


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/isimmutable
func (e_ EKCalendar) IsImmutable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isImmutable"))
	return rv
}


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/isimmutable
func (e_ EKCalendar) SetIsImmutable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsImmutable:"), value)
}


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/issubscribed
func (e_ EKCalendar) IsSubscribed() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isSubscribed"))
	return rv
}


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/issubscribed
func (e_ EKCalendar) SetIsSubscribed(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsSubscribed:"), value)
}


