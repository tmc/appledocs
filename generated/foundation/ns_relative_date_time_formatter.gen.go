// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRelativeDateTimeFormatter */


/* debug [class_header]: Header for NSRelativeDateTimeFormatter */
// The class instance for the [RelativeDateTimeFormatter] class.
var (
	RelativeDateTimeFormatterClass     _RelativeDateTimeFormatterClass
	RelativeDateTimeFormatterClassOnce sync.Once
)

func getRelativeDateTimeFormatterClass() _RelativeDateTimeFormatterClass {
	RelativeDateTimeFormatterClassOnce.Do(func() {
		RelativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}
	})
	return RelativeDateTimeFormatterClass
}

type _RelativeDateTimeFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RelativeDateTimeFormatter */
// An interface definition for the [RelativeDateTimeFormatter] class.
type IRelativeDateTimeFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for RelativeDateTimeFormatter */
	// properties:
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	DateTimeStyle() RelativeDateTimeFormatterStyle
	SetDateTimeStyle(value RelativeDateTimeFormatterStyle)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	Locale() ILocale
	SetLocale(value ILocale)
	UnitsStyle() RelativeDateTimeFormatterUnitsStyle
	SetUnitsStyle(value RelativeDateTimeFormatterUnitsStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RelativeDateTimeFormatter */
	// methods:
	LocalizedStringForDateRelativeToDate(date IDate, referenceDate IDate) IString
	LocalizedStringFromDateComponents(dateComponents IDateComponents) IString
	LocalizedStringFromTimeInterval(timeInterval float64) IString
	StringForObjectValue(obj objc.IObject) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RelativeDateTimeFormatter */
// Alloc allocates a new instance without initialization.
func (rc _RelativeDateTimeFormatterClass) Alloc() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RelativeDateTimeFormatterClass) New() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelativeDateTimeFormatter) Init() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelativeDateTimeFormatter) Autorelease() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelativeDateTimeFormatter creates a new RelativeDateTimeFormatter instance.
func NewRelativeDateTimeFormatter() RelativeDateTimeFormatter {
	return getRelativeDateTimeFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RelativeDateTimeFormatter */
// A formatter that creates locale-aware string representations of a relative date or time.
//
// Use the strings that the formatter produces, such as “1 hour ago”, “in 2 weeks”, “yesterday”, and “tomorrow” as standalone strings. Embedding them in other strings may not be grammatically correct.


// A formatter that creates locale-aware string representations of a relative date or time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter
type RelativeDateTimeFormatter struct {
	Formatter
}

// RelativeDateTimeFormatterFrom constructs a [RelativeDateTimeFormatter] from an unsafe.Pointer.
//
// A formatter that creates locale-aware string representations of a relative date or time.
func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RelativeDateTimeFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RelativeDateTimeFormatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RelativeDateTimeFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RelativeDateTimeFormatter */

// Formats the date interval from the reference date to the specified date using the formatter’s calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(for:relativeTo:)
func (r_ RelativeDateTimeFormatter) LocalizedStringForDateRelativeToDate(date IDate, referenceDate IDate) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("localizedStringForDate:relativeToDate:"), date, referenceDate)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForDateRelativeToDate */


// Formats a relative time represented by the specified date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(from:)
func (r_ RelativeDateTimeFormatter) LocalizedStringFromDateComponents(dateComponents IDateComponents) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("localizedStringFromDateComponents:"), dateComponents)
	return rv
}/* debug [instance_methods/method]: LocalizedStringFromDateComponents */


// Formats the specified time interval using the formatter’s calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(fromTimeInterval:)
func (r_ RelativeDateTimeFormatter) LocalizedStringFromTimeInterval(timeInterval float64) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("localizedStringFromTimeInterval:"), timeInterval)
	return rv
}/* debug [instance_methods/method]: LocalizedStringFromTimeInterval */


// Creates a formatted string for a date relative to the current date and time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/string(for:)
func (r_ RelativeDateTimeFormatter) StringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](r_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}/* debug [instance_methods/method]: StringForObjectValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RelativeDateTimeFormatter */

// The calendar to use for formatting values that don’t have an inherent calendar of their own.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/calendar
func (r_ RelativeDateTimeFormatter) Calendar() ICalendar {
	rv := objc.Send[Calendar](r_.ID, objc.Sel("calendar"))
	return rv
}/* debug [instance_properties/getter]: calendar */


// The calendar to use for formatting values that don’t have an inherent calendar of their own.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/calendar
func (r_ RelativeDateTimeFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCalendar:"), value)
}/* debug [instance_properties/setter]: calendar */


// The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/dateTimeStyle-swift.property
func (r_ RelativeDateTimeFormatter) DateTimeStyle() RelativeDateTimeFormatterStyle {
	rv := objc.Send[RelativeDateTimeFormatterStyle](r_.ID, objc.Sel("dateTimeStyle"))
	return rv
}/* debug [instance_properties/getter]: dateTimeStyle */


// The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/dateTimeStyle-swift.property
func (r_ RelativeDateTimeFormatter) SetDateTimeStyle(value RelativeDateTimeFormatterStyle) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDateTimeStyle:"), value)
}/* debug [instance_properties/setter]: dateTimeStyle */


// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/formattingContext
func (r_ RelativeDateTimeFormatter) FormattingContext() FormattingContext {
	rv := objc.Send[FormattingContext](r_.ID, objc.Sel("formattingContext"))
	return rv
}/* debug [instance_properties/getter]: formattingContext */


// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/formattingContext
func (r_ RelativeDateTimeFormatter) SetFormattingContext(value FormattingContext) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFormattingContext:"), value)
}/* debug [instance_properties/setter]: formattingContext */


// The locale to use when formatting the date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/locale
func (r_ RelativeDateTimeFormatter) Locale() ILocale {
	rv := objc.Send[Locale](r_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale to use when formatting the date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/locale
func (r_ RelativeDateTimeFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/unitsStyle-swift.property
func (r_ RelativeDateTimeFormatter) UnitsStyle() RelativeDateTimeFormatterUnitsStyle {
	rv := objc.Send[RelativeDateTimeFormatterUnitsStyle](r_.ID, objc.Sel("unitsStyle"))
	return rv
}/* debug [instance_properties/getter]: unitsStyle */


// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/unitsStyle-swift.property
func (r_ RelativeDateTimeFormatter) SetUnitsStyle(value RelativeDateTimeFormatterUnitsStyle) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setUnitsStyle:"), value)
}/* debug [instance_properties/setter]: unitsStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRelativeDateTimeFormatter */



