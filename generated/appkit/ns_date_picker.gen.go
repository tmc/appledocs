// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Calendar() foundation.Calendar
	SetCalendar(value foundation.Calendar)
	DatePickerElements() DatePickerElementFlags
	SetDatePickerElements(value DatePickerElementFlags)
	DatePickerMode() DatePickerMode
	SetDatePickerMode(value DatePickerMode)
	DatePickerStyle() DatePickerStyle
	SetDatePickerStyle(value DatePickerStyle)
	DateValue() foundation.foundation.INSDate
	SetDateValue(value foundation.foundation.INSDate)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Bezeled() bool
	SetBezeled(value bool)
	Bordered() bool
	SetBordered(value bool)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	MaxDate() foundation.foundation.INSDate
	SetMaxDate(value foundation.foundation.INSDate)
	MinDate() foundation.foundation.INSDate
	SetMinDate(value foundation.foundation.INSDate)
	PresentsCalendarOverlay() bool
	SetPresentsCalendarOverlay(value bool)
	TextColor() IColor
	SetTextColor(value IColor)
	TimeInterval() float64
	SetTimeInterval(value float64)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
	IsBezeled() bool
	SetIsBezeled(value bool)
	IsBordered() bool
	SetIsBordered(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DatePickerClass) Alloc() DatePicker {
	rv := objc.Send[DatePicker](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/calendar
func (d_ DatePicker) Calendar() foundation.Calendar {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/calendar
func (d_ DatePicker) SetCalendar(value foundation.Calendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerElements
func (d_ DatePicker) DatePickerElements() DatePickerElementFlags {
	rv := objc.Send[DatePickerElementFlags](d_.ID, objc.Sel("datePickerElements"))
	return rv
}


// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerElements
func (d_ DatePicker) SetDatePickerElements(value DatePickerElementFlags) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
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


// The date picker’s style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerStyle
func (d_ DatePicker) DatePickerStyle() DatePickerStyle {
	rv := objc.Send[DatePickerStyle](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}


// The date picker’s style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/datePickerStyle
func (d_ DatePicker) SetDatePickerStyle(value DatePickerStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}


// The date selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/dateValue
func (d_ DatePicker) DateValue() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("dateValue"))
	return rv
}


// The date selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/dateValue
func (d_ DatePicker) SetDateValue(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}


// A Boolean value that indicates whether the date picker draws the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/drawsBackground
func (d_ DatePicker) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value that indicates whether the date picker draws the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/drawsBackground
func (d_ DatePicker) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
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
func (d_ DatePicker) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The date picker’s locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/locale
func (d_ DatePicker) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The date picker’s maximum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/maxDate
func (d_ DatePicker) MaxDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("maxDate"))
	return rv
}


// The date picker’s maximum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/maxDate
func (d_ DatePicker) SetMaxDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}


// The date picker’s minimum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/minDate
func (d_ DatePicker) MinDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("minDate"))
	return rv
}


// The date picker’s minimum date value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/minDate
func (d_ DatePicker) SetMinDate(value foundation.foundation.INSDate) {
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
	rv := objc.Send[float64](d_.ID, objc.Sel("timeInterval"))
	return rv
}


// The time interval selected by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeInterval
func (d_ DatePicker) SetTimeInterval(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}


// The time zone for the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeZone
func (d_ DatePicker) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone for the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/timeZone
func (d_ DatePicker) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
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








