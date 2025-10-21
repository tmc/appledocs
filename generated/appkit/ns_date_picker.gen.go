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
}

// A display of a calendar date with controls for editing the date value.
//
// uses an to implement much of the control’s functionality. provides cover methods for most of methods, which invoke the corresponding cell method.
//
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
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/backgroundcolor
func (d_ DatePicker) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](d_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The date picker’s background color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/backgroundcolor
func (d_ DatePicker) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), value)
}

// The calendar used by the date picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/calendar
func (d_ DatePicker) Calendar() foundation.Calendar {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// SetCalendar sets the value of the calendar property.
// The calendar used by the date picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/calendar
func (d_ DatePicker) SetCalendar(value foundation.ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}

// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerelements
func (d_ DatePicker) DatePickerElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerElements"))
	return rv
}


// SetDatePickerElements sets the value of the datePickerElements property.
// A bitmask that indicates which visual elements of the date picker are currently shown, and which won’t be usable because they are hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerelements
func (d_ DatePicker) SetDatePickerElements(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
}

// The date picker’s mode.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickermode
func (d_ DatePicker) DatePickerMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerMode"))
	return rv
}


// SetDatePickerMode sets the value of the datePickerMode property.
// The date picker’s mode.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickermode
func (d_ DatePicker) SetDatePickerMode(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerMode:"), value)
}

// The date picker’s style.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerstyle
func (d_ DatePicker) DatePickerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}


// SetDatePickerStyle sets the value of the datePickerStyle property.
// The date picker’s style.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datepickerstyle
func (d_ DatePicker) SetDatePickerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}

// The date selected by the date picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datevalue
func (d_ DatePicker) DateValue() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("dateValue"))
	return rv
}


// SetDateValue sets the value of the dateValue property.
// The date selected by the date picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/datevalue
func (d_ DatePicker) SetDateValue(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}

// A delegate for the date picker’s cell
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/delegate
func (d_ DatePicker) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate for the date picker’s cell

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/delegate
func (d_ DatePicker) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the date picker draws the background.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/drawsbackground
func (d_ DatePicker) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean value that indicates whether the date picker draws the background.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/drawsbackground
func (d_ DatePicker) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}

// A Boolean value that indicates whether the date picker draws a bezeled border.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbezeled
func (d_ DatePicker) IsBezeled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isBezeled"))
	return rv
}


// SetIsBezeled sets the value of the isBezeled property.
// A Boolean value that indicates whether the date picker draws a bezeled border.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbezeled
func (d_ DatePicker) SetIsBezeled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsBezeled:"), value)
}

// A Boolean value that indicates whether the date picker has a plain border.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbordered
func (d_ DatePicker) IsBordered() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isBordered"))
	return rv
}


// SetIsBordered sets the value of the isBordered property.
// A Boolean value that indicates whether the date picker has a plain border.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/isbordered
func (d_ DatePicker) SetIsBordered(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsBordered:"), value)
}

// The date picker’s locale.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/locale
func (d_ DatePicker) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The date picker’s locale.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/locale
func (d_ DatePicker) SetLocale(value foundation.ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}

// The date picker’s maximum date value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/maxdate
func (d_ DatePicker) MaxDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("maxDate"))
	return rv
}


// SetMaxDate sets the value of the maxDate property.
// The date picker’s maximum date value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/maxdate
func (d_ DatePicker) SetMaxDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}

// The date picker’s minimum date value.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/mindate
func (d_ DatePicker) MinDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("minDate"))
	return rv
}


// SetMinDate sets the value of the minDate property.
// The date picker’s minimum date value.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/mindate
func (d_ DatePicker) SetMinDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}

// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/presentscalendaroverlay
func (d_ DatePicker) PresentsCalendarOverlay() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentsCalendarOverlay"))
	return rv
}


// SetPresentsCalendarOverlay sets the value of the presentsCalendarOverlay property.
// A Boolean value that indicates whether to present a graphical calendar overlay when editing a calendar element within a text-field style date picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/presentscalendaroverlay
func (d_ DatePicker) SetPresentsCalendarOverlay(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPresentsCalendarOverlay:"), value)
}

// The date picker’s text color.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/textcolor
func (d_ DatePicker) TextColor() NSColor {
	rv := objc.Send[NSColor](d_.ID, objc.Sel("textColor"))
	return rv
}


// SetTextColor sets the value of the textColor property.
// The date picker’s text color.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/textcolor
func (d_ DatePicker) SetTextColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextColor:"), value)
}

// The time interval selected by the date picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timeinterval
func (d_ DatePicker) TimeInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timeInterval"))
	return rv
}


// SetTimeInterval sets the value of the timeInterval property.
// The time interval selected by the date picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timeinterval
func (d_ DatePicker) SetTimeInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}

// The time zone for the date picker.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timezone
func (d_ DatePicker) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// SetTimeZone sets the value of the timeZone property.
// The time zone for the date picker.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepicker/timezone
func (d_ DatePicker) SetTimeZone(value foundation.ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}



