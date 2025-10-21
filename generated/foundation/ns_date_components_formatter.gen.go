// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	StringForObjectValue(obj objc.ID) string
}

// A formatter that creates string representations of quantities of time.
//
// An object takes quantities of time and formats them as a user-readable string. Use a date components formatter to create strings for your app’s interface. The formatter object has many options for creating both abbreviated and expanded strings. The formatter takes the current user’s locale and language into account when generating strings. To use this class, create an instance, configure its properties, and call one of its methods to generate an appropriate string. The properties of this class let you configure the calendar and specify the date and time units you want displayed in the resulting string. The listing below shows how to configure a formatter to create the string “About 5 minutes remaining”. The methods of this class may be called safely from any thread of your app. It is also safe to share a single instance of this class from multiple threads, with the caveat that you should not change the configuration of the object while another thread is using it to generate a string.
//
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


// Returns a formatted string based on the date information in the specified object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(for:)
func (d_ DateComponentsFormatter) StringForObjectValue(obj objc.ID) string {
	rv := objc.Send[string](d_.ID, objc.Sel("stringForObjectValue:"), obj)
	return rv
}

// The maximum number of time units to include in the output string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/maximumunitcount
func (d_ DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumUnitCount"))
	return rv
}


// SetMaximumUnitCount sets the value of the maximumUnitCount property.
// The maximum number of time units to include in the output string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/maximumunitcount
func (d_ DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumUnitCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/referencedate
func (d_ DateComponentsFormatter) ReferenceDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("referenceDate"))
	return rv
}


// SetReferenceDate sets the value of the referenceDate property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/referencedate
func (d_ DateComponentsFormatter) SetReferenceDate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setReferenceDate:"), value)
}

// A Boolean value indicating whether output strings reflect the amount of time remaining.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/includestimeremainingphrase
func (d_ DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("includesTimeRemainingPhrase"))
	return rv
}


// SetIncludesTimeRemainingPhrase sets the value of the includesTimeRemainingPhrase property.
// A Boolean value indicating whether output strings reflect the amount of time remaining.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/includestimeremainingphrase
func (d_ DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIncludesTimeRemainingPhrase:"), value)
}

// The bitmask of calendrical units such as day and month to include in the output string.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/allowedunits
func (d_ DateComponentsFormatter) AllowedUnits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("allowedUnits"))
	return rv
}


// SetAllowedUnits sets the value of the allowedUnits property.
// The bitmask of calendrical units such as day and month to include in the output string.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/allowedunits
func (d_ DateComponentsFormatter) SetAllowedUnits(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowedUnits:"), value)
}

// A Boolean value indicating whether the resulting phrase reflects an inexact time value.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/includesapproximationphrase
func (d_ DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("includesApproximationPhrase"))
	return rv
}


// SetIncludesApproximationPhrase sets the value of the includesApproximationPhrase property.
// A Boolean value indicating whether the resulting phrase reflects an inexact time value.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/includesapproximationphrase
func (d_ DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIncludesApproximationPhrase:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/formattingcontext
func (d_ DateComponentsFormatter) FormattingContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("formattingContext"))
	return rv
}


// SetFormattingContext sets the value of the formattingContext property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/formattingcontext
func (d_ DateComponentsFormatter) SetFormattingContext(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFormattingContext:"), value)
}

// A Boolean indicating whether non-integer units may be used for values.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/allowsfractionalunits
func (d_ DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsFractionalUnits"))
	return rv
}


// SetAllowsFractionalUnits sets the value of the allowsFractionalUnits property.
// A Boolean indicating whether non-integer units may be used for values.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/allowsfractionalunits
func (d_ DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowsFractionalUnits:"), value)
}

// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/collapseslargestunit
func (d_ DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("collapsesLargestUnit"))
	return rv
}


// SetCollapsesLargestUnit sets the value of the collapsesLargestUnit property.
// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/collapseslargestunit
func (d_ DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCollapsesLargestUnit:"), value)
}

// The formatting style for unit names.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/unitsstyle-swift.property
func (d_ DateComponentsFormatter) UnitsStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("unitsStyle"))
	return rv
}


// SetUnitsStyle sets the value of the unitsStyle property.
// The formatting style for unit names.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/unitsstyle-swift.property
func (d_ DateComponentsFormatter) SetUnitsStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUnitsStyle:"), value)
}

// The default calendar to use when formatting date components.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/calendar
func (d_ DateComponentsFormatter) Calendar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("calendar"))
	return rv
}


// SetCalendar sets the value of the calendar property.
// The default calendar to use when formatting date components.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/datecomponentsformatter/calendar
func (d_ DateComponentsFormatter) SetCalendar(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}

// The formatting style for units whose value is 0.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) ZeroFormattingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("zeroFormattingBehavior"))
	return rv
}


// SetZeroFormattingBehavior sets the value of the zeroFormattingBehavior property.
// The formatting style for units whose value is 0.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d_ DateComponentsFormatter) SetZeroFormattingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setZeroFormattingBehavior:"), value)
}



