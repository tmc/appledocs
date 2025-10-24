// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DatePicker] class.
var (
	DatePickerClass     _DatePickerClass
	DatePickerClassOnce sync.Once
)

func getDatePickerClass() _DatePickerClass {
	DatePickerClassOnce.Do(func() {
		DatePickerClass = _DatePickerClass{objc.GetClass("NSDatePicker")}
	})
	return DatePickerClass
}

type _DatePickerClass struct {
	class objc.Class
}

// An interface definition for the [DatePicker] class.
type IDatePicker interface {
	IControl
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	DateValue() objc.IObject /* cross-framework: NSDate */
	SetDateValue(value objc.IObject /* cross-framework: NSDate */)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Bezeled() bool
	SetBezeled(value bool)
	Bordered() bool
	SetBordered(value bool)
	Locale() objc.IObject /* cross-framework: Locale */
	SetLocale(value objc.IObject /* cross-framework: Locale */)
	MaxDate() objc.IObject /* cross-framework: NSDate */
	SetMaxDate(value objc.IObject /* cross-framework: NSDate */)
	MinDate() objc.IObject /* cross-framework: NSDate */
	SetMinDate(value objc.IObject /* cross-framework: NSDate */)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	TextColor() IColor
	SetTextColor(value IColor)
	TimeInterval() float64
	SetTimeInterval(value float64)
	Calendar() objc.IObject /* cross-framework: Calendar */
	SetCalendar(value objc.IObject /* cross-framework: Calendar */)
	DatePickerElements() unsafe.Pointer
	SetDatePickerElements(value unsafe.Pointer)
	DatePickerStyle() unsafe.Pointer
	SetDatePickerStyle(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsBordered() bool
	SetIsBordered(value bool)
	TimeZone() objc.IObject /* cross-framework: TimeZone */
	SetTimeZone(value objc.IObject /* cross-framework: TimeZone */)
	// methods:
}

// A display of a calendar date with controls for editing the date value.
//
// uses an to implement much of the control’s functionality. provides cover methods for most of methods, which invoke the corresponding cell method.


// A display of a calendar date with controls for editing the date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker
type DatePicker struct {
	Control
}

// DatePickerFrom constructs a [DatePicker] from an unsafe.Pointer.
//
// A display of a calendar date with controls for editing the date value.
func DatePickerFrom(ptr unsafe.Pointer) DatePicker {
	return DatePicker{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DatePickerClass) New() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DatePicker) Init() DatePicker {
	rv := objc.Send[DatePicker](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DatePicker) Autorelease() DatePicker {
	rv := objc.Send[DatePicker](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDatePicker creates a new DatePicker instance.
func NewDatePicker() DatePicker {
	return getDatePickerClass().New()
}



// The date picker’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/backgroundColor
func (d_ DatePicker) BackgroundColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The date picker’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/backgroundColor
func (d_ DatePicker) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The date picker’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerMode
func (d_ DatePicker) DatePickerMode() DatePickerMode {
	rv := objc.Send[DatePickerMode](d_.ID, objc.Sel("datePickerMode"))
	return rv
}


// The date picker’s mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerMode
func (d_ DatePicker) SetDatePickerMode(value DatePickerMode) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerMode:"), value)
}


// The date selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/dateValue
func (d_ DatePicker) DateValue() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("dateValue"))
	return rv
}


// The date selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/dateValue
func (d_ DatePicker) SetDateValue(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}


// A delegate for the date picker’s cell
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/delegate
func (d_ DatePicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate for the date picker’s cell
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/delegate
func (d_ DatePicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value that indicates whether the date picker draws a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBezeled
func (d_ DatePicker) Bezeled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("bezeled"))
	return rv
}


// A Boolean value that indicates whether the date picker draws a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBezeled
func (d_ DatePicker) SetBezeled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBezeled:"), value)
}


// A Boolean value that indicates whether the date picker has a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBordered
func (d_ DatePicker) Bordered() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("bordered"))
	return rv
}


// A Boolean value that indicates whether the date picker has a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/isBordered
func (d_ DatePicker) SetBordered(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBordered:"), value)
}


// The date picker’s locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/locale
func (d_ DatePicker) Locale() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The date picker’s locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/locale
func (d_ DatePicker) SetLocale(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The date picker’s maximum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/maxDate
func (d_ DatePicker) MaxDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("maxDate"))
	return rv
}


// The date picker’s maximum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/maxDate
func (d_ DatePicker) SetMaxDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}


// The date picker’s minimum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/minDate
func (d_ DatePicker) MinDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("minDate"))
	return rv
}


// The date picker’s minimum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/minDate
func (d_ DatePicker) SetMinDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}


// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/presentsCalendarOverlay
func (d_ DatePicker) PresentsCalendarOverlay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentsCalendarOverlay"))
	return rv
}


// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/presentsCalendarOverlay
func (d_ DatePicker) SetPresentsCalendarOverlay(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPresentsCalendarOverlay:"), value)
}


// The date picker’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/textColor
func (d_ DatePicker) TextColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("textColor"))
	return rv
}


// The date picker’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/textColor
func (d_ DatePicker) SetTextColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextColor:"), value)
}


// The time interval selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeInterval
func (d_ DatePicker) TimeInterval() float64 {
	rv := objc.Send[TimeInterval](d_.ID, objc.Sel("timeInterval"))
	return rv
}


// The time interval selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeInterval
func (d_ DatePicker) SetTimeInterval(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/calendar
func (d_ DatePicker) Calendar() objc.IObject /* cross-framework: Calendar */ {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/calendar
func (d_ DatePicker) SetCalendar(value objc.IObject /* cross-framework: Calendar */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerelements
func (d_ DatePicker) DatePickerElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerElements"))
	return rv
}


// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerelements
func (d_ DatePicker) SetDatePickerElements(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
}


// The date picker’s style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerstyle
func (d_ DatePicker) DatePickerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}


// The date picker’s style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerstyle
func (d_ DatePicker) SetDatePickerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}


// A Boolean value that indicates whether the date picker draws the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/drawsbackground
func (d_ DatePicker) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the date picker draws the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/drawsbackground
func (d_ DatePicker) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}


// A Boolean value that indicates whether the date picker draws a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbezeled
func (d_ DatePicker) IsBezeled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isBezeled"))
	return rv
}


// A Boolean value that indicates whether the date picker draws a bezeled border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbezeled
func (d_ DatePicker) SetIsBezeled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsBezeled:"), value)
}


// A Boolean value that indicates whether the date picker has a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbordered
func (d_ DatePicker) IsBordered() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isBordered"))
	return rv
}


// A Boolean value that indicates whether the date picker has a plain border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbordered
func (d_ DatePicker) SetIsBordered(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsBordered:"), value)
}


// The time zone for the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timezone
func (d_ DatePicker) TimeZone() objc.IObject /* cross-framework: TimeZone */ {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timezone
func (d_ DatePicker) SetTimeZone(value objc.IObject /* cross-framework: TimeZone */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}



