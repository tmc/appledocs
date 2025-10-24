// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKCalendar */


/* debug [class_header]: Header for EKCalendar */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKCalendar */
// An interface definition for the [EKCalendar] class.
type IEKCalendar interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKCalendar */
	// properties:
	AllowedEntityTypes() EKEntityMask
	AllowsContentModifications() bool
	CalendarIdentifier() objc.IObject /* cross-framework: NSString */
	CGColor() ColorRef /* not a class type */
	SetCGColor(value ColorRef /* not a class type */)
	Color() appkit.Color
	SetColor(value appkit.Color)
	Immutable() bool
	Subscribed() bool
	Source() IEKSource
	SetSource(value IEKSource)
	SupportedEventAvailabilities() EKCalendarEventAvailabilityMask
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	Type() EKCalendarType
	IsImmutable() bool
	SetIsImmutable(value bool)
	IsSubscribed() bool
	SetIsSubscribed(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKCalendar */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKCalendar */
// Alloc allocates a new instance without initialization.
func (ec _EKCalendarClass) Alloc() EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKCalendar */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKCalendar */

// Creates a new calendar that can contain the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(for:eventStore:)
func NewEKCalendarForEntityTypeEventStore(entityType EKEntityType, eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(getEKCalendarClass().class), objc.Sel("calendarForEntityType:eventStore:"), entityType, eventStore)
	return rv
}/* debug [class_init_methods/constructor]: NewEKCalendarForEntityTypeEventStore */


// Creates and returns a calendar belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(eventStore:)
func NewEKCalendarWithEventStore(eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(getEKCalendarClass().class), objc.Sel("calendarWithEventStore:"), eventStore)
	return rv
}/* debug [class_init_methods/constructor]: NewEKCalendarWithEventStore */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKCalendar */

// Creates and returns a calendar belonging to a specified event store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(eventStore:)
func (ec _EKCalendarClass) CalendarWithEventStore(eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("calendarWithEventStore:"), eventStore)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalendarWithEventStore) */


// Creates a new calendar that can contain the given entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/init(for:eventStore:)
func (ec _EKCalendarClass) CalendarForEntityTypeEventStore(entityType EKEntityType, eventStore IEKEventStore) EKCalendar {
	rv := objc.Send[EKCalendar](objc.ID(ec.class), objc.Sel("calendarForEntityType:eventStore:"), entityType, eventStore)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalendarForEntityTypeEventStore) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKCalendar */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKCalendar */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKCalendar */

// The entity types this calendar can contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/allowedEntityTypes
func (e_ EKCalendar) AllowedEntityTypes() EKEntityMask {
	rv := objc.Send[EKEntityMask](e_.ID, objc.Sel("allowedEntityTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedEntityTypes */


// A Boolean value that indicates whether you can add, edit, and delete items in the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/allowsContentModifications
func (e_ EKCalendar) AllowsContentModifications() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("allowsContentModifications"))
	return rv
}/* debug [instance_properties/getter]: allowsContentModifications */


// A unique identifier for the calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/calendarIdentifier
func (e_ EKCalendar) CalendarIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("calendarIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarIdentifier */


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/cgColor
func (e_ EKCalendar) CGColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](e_.ID, objc.Sel("CGColor"))
	return rv
}/* debug [instance_properties/getter]: CGColor */


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/cgColor
func (e_ EKCalendar) SetCGColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setCGColor:"), value)
}/* debug [instance_properties/setter]: CGColor */


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/color
func (e_ EKCalendar) Color() appkit.Color {
	rv := objc.Send[appkit.Color](e_.ID, objc.Sel("color"))
	return rv
}/* debug [instance_properties/getter]: color */


// The calendar’s color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/color
func (e_ EKCalendar) SetColor(value appkit.Color) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setColor:"), value)
}/* debug [instance_properties/setter]: color */


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/isImmutable
func (e_ EKCalendar) Immutable() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("immutable"))
	return rv
}/* debug [instance_properties/getter]: immutable */


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/isSubscribed
func (e_ EKCalendar) Subscribed() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("subscribed"))
	return rv
}/* debug [instance_properties/getter]: subscribed */


// The source object representing the account to which this calendar belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/source
func (e_ EKCalendar) Source() IEKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// The source object representing the account to which this calendar belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/source
func (e_ EKCalendar) SetSource(value IEKSource) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setSource:"), value)
}/* debug [instance_properties/setter]: source */


// The event availability settings supported by this calendar, as indicated by a bitmask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/supportedEventAvailabilities
func (e_ EKCalendar) SupportedEventAvailabilities() EKCalendarEventAvailabilityMask {
	rv := objc.Send[EKCalendarEventAvailabilityMask](e_.ID, objc.Sel("supportedEventAvailabilities"))
	return rv
}/* debug [instance_properties/getter]: supportedEventAvailabilities */


// The calendar’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/title
func (e_ EKCalendar) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The calendar’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/title
func (e_ EKCalendar) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The calendar’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKCalendar/type
func (e_ EKCalendar) Type() EKCalendarType {
	rv := objc.Send[EKCalendarType](e_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/isimmutable
func (e_ EKCalendar) IsImmutable() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isImmutable"))
	return rv
}/* debug [instance_properties/getter]: isImmutable */


// A Boolean value indicating whether the calendar’s properties can be edited or deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/isimmutable
func (e_ EKCalendar) SetIsImmutable(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsImmutable:"), value)
}/* debug [instance_properties/setter]: isImmutable */


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/issubscribed
func (e_ EKCalendar) IsSubscribed() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isSubscribed"))
	return rv
}/* debug [instance_properties/getter]: isSubscribed */


// A Boolean value indicating whether the calendar is a subscribed calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/eventkit/ekcalendar/issubscribed
func (e_ EKCalendar) SetIsSubscribed(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsSubscribed:"), value)
}/* debug [instance_properties/setter]: isSubscribed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKCalendar */


