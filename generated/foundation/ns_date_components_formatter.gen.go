// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DateComponentsFormatter] class.
var (
	DateComponentsFormatterClass     _DateComponentsFormatterClass
	DateComponentsFormatterClassOnce sync.Once
)

func getDateComponentsFormatterClass() _DateComponentsFormatterClass {
	DateComponentsFormatterClassOnce.Do(func() {
		DateComponentsFormatterClass = _DateComponentsFormatterClass{objc.GetClass("NSDateComponentsFormatter")}
	})
	return DateComponentsFormatterClass
}

type _DateComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [DateComponentsFormatter] class.
type IDateComponentsFormatter interface {
	IFormatter
	// properties:
	AllowedUnits() CalendarUnit
	SetAllowedUnits(value CalendarUnit)
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	Calendar() ICalendar
	SetCalendar(value ICalendar)
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	ReferenceDate() IDate
	SetReferenceDate(value IDate)
	UnitsStyle() DateComponentsFormatterUnitsStyle
	SetUnitsStyle(value DateComponentsFormatterUnitsStyle)
	ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior)
	// methods:
	GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool
	StringForObjectValue(obj objc.IObject) IString
	StringFromTimeInterval(ti float64) IString
	StringFromDateComponents(components IDateComponents) IString
	StringFromDateToDate(startDate IDate, endDate IDate) IString
}

// A formatter that creates string representations of quantities of time.
//
// An object takes quantities of time and formats them as a user-readable string. Use a date components formatter to create strings for your app’s interface. The formatter object has many options for creating both abbreviated and expanded strings. The formatter takes the current user’s locale and language into account when generating strings. To use this class, create an instance, configure its properties, and call one of its methods to generate an appropriate string. The properties of this class let you configure the calendar and specify the date and time units you want displayed in the resulting string. The listing below shows how to configure a formatter to create the string “About 5 minutes remaining”. The methods of this class may be called safely from any thread of your app. It is also safe to share a single instance of this class from multiple threads, with the caveat that you should not change the configuration of the object while another thread is using it to generate a string.


// A formatter that creates string representations of quantities of time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter
type DateComponentsFormatter struct {
	Formatter
}

// DateComponentsFormatterFrom constructs a [DateComponentsFormatter] from an unsafe.Pointer.
//
// A formatter that creates string representations of quantities of time.
func DateComponentsFormatterFrom(ptr unsafe.Pointer) DateComponentsFormatter {
	return DateComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DateComponentsFormatterClass) New() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DateComponentsFormatter) Autorelease() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateComponentsFormatter creates a new DateComponentsFormatter instance.
func NewDateComponentsFormatter() DateComponentsFormatter {
	return getDateComponentsFormatterClass().New()
}



// Returns a localized string based on the specified date components and style option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/localizedString(from:unitsStyle:)
func (dc _DateComponentsFormatterClass) LocalizedStringFromDateComponentsUnitsStyle(components IDateComponents, unitsStyle DateComponentsFormatterUnitsStyle) IString {
	rv := objc.Send[String](objc.ID(dc.class), objc.Sel("localizedStringFromDateComponents:unitsStyle:"), components, unitsStyle)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/getObjectValue(_:for:errorDescription:)
func (d_ DateComponentsFormatter) GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}


// Returns a formatted string based on the date information in the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(for:)
func (d_ DateComponentsFormatter) StringForObjectValue(obj objc.IObject) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}


// Returns a formatted string based on the specified number of seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:)-7sj4j
func (d_ DateComponentsFormatter) StringFromTimeInterval(ti float64) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("stringFromTimeInterval:"), ti)
	return rv
}


// Returns a formatted string based on the specified date component information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:)-9exxn
func (d_ DateComponentsFormatter) StringFromDateComponents(components IDateComponents) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("stringFromDateComponents:"), components)
	return rv
}


// Returns a formatted string based on the time difference between two dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:to:)
func (d_ DateComponentsFormatter) StringFromDateToDate(startDate IDate, endDate IDate) IString {
	rv := objc.Send[String](d_.ID, objc.Sel("stringFromDate:toDate:"), startDate, endDate)
	return rv
}


// The bitmask of calendrical units such as day and month to include in the output string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowedUnits
func (d_ DateComponentsFormatter) AllowedUnits() CalendarUnit {
	rv := objc.Send[CalendarUnit](d_.ID, objc.Sel("allowedUnits"))
	return rv
}


// The bitmask of calendrical units such as day and month to include in the output string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowedUnits
func (d_ DateComponentsFormatter) SetAllowedUnits(value CalendarUnit) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowedUnits:"), value)
}


// A Boolean indicating whether non-integer units may be used for values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowsFractionalUnits
func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsFractionalUnits"))
	return rv
}


// A Boolean indicating whether non-integer units may be used for values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowsFractionalUnits
func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowsFractionalUnits:"), value)
}


// The default calendar to use when formatting date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/calendar
func (d_ DateComponentsFormatter) Calendar() ICalendar {
	rv := objc.Send[Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The default calendar to use when formatting date components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/calendar
func (d_ DateComponentsFormatter) SetCalendar(value ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/collapsesLargestUnit
func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("collapsesLargestUnit"))
	return rv
}


// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/collapsesLargestUnit
func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCollapsesLargestUnit:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/formattingContext
func (d_ DateComponentsFormatter) FormattingContext() FormattingContext {
	rv := objc.Send[FormattingContext](d_.ID, objc.Sel("formattingContext"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/formattingContext
func (d_ DateComponentsFormatter) SetFormattingContext(value FormattingContext) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}


// A Boolean value indicating whether the resulting phrase reflects an inexact time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesApproximationPhrase
func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("includesApproximationPhrase"))
	return rv
}


// A Boolean value indicating whether the resulting phrase reflects an inexact time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesApproximationPhrase
func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIncludesApproximationPhrase:"), value)
}


// A Boolean value indicating whether output strings reflect the amount of time remaining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesTimeRemainingPhrase
func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("includesTimeRemainingPhrase"))
	return rv
}


// A Boolean value indicating whether output strings reflect the amount of time remaining.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesTimeRemainingPhrase
func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIncludesTimeRemainingPhrase:"), value)
}


// The maximum number of time units to include in the output string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/maximumUnitCount
func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumUnitCount"))
	return rv
}


// The maximum number of time units to include in the output string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/maximumUnitCount
func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumUnitCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/referenceDate
func (d_ DateComponentsFormatter) ReferenceDate() IDate {
	rv := objc.Send[Date](d_.ID, objc.Sel("referenceDate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/referenceDate
func (d_ DateComponentsFormatter) SetReferenceDate(value IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReferenceDate:"), value)
}


// The formatting style for unit names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/unitsStyle-swift.property
func (d_ DateComponentsFormatter) UnitsStyle() DateComponentsFormatterUnitsStyle {
	rv := objc.Send[DateComponentsFormatterUnitsStyle](d_.ID, objc.Sel("unitsStyle"))
	return rv
}


// The formatting style for unit names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/unitsStyle-swift.property
func (d_ DateComponentsFormatter) SetUnitsStyle(value DateComponentsFormatterUnitsStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUnitsStyle:"), value)
}


// The formatting style for units whose value is 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) ZeroFormattingBehavior() DateComponentsFormatterZeroFormattingBehavior {
	rv := objc.Send[DateComponentsFormatterZeroFormattingBehavior](d_.ID, objc.Sel("zeroFormattingBehavior"))
	return rv
}


// The formatting style for units whose value is 0.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value DateComponentsFormatterZeroFormattingBehavior) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setZeroFormattingBehavior:"), value)
}



