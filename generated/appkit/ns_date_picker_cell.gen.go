// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DatePickerCell] class.
var (
	DatePickerCellClass     _DatePickerCellClass
	DatePickerCellClassOnce sync.Once
)

func getDatePickerCellClass() _DatePickerCellClass {
	DatePickerCellClassOnce.Do(func() {
		DatePickerCellClass = _DatePickerCellClass{objc.GetClass("NSDatePickerCell")}
	})
	return DatePickerCellClass
}

type _DatePickerCellClass struct {
	class objc.Class
}

// An interface definition for the [DatePickerCell] class.
type IDatePickerCell interface {
	IActionCell
	MaxDate() foundation.NSDate
	SetMaxDate(value foundation.IDate)
	BackgroundColor() NSColor
	SetBackgroundColor(value IColor)
	Calendar() foundation.Calendar
	SetCalendar(value foundation.ICalendar)
	DatePickerElements() unsafe.Pointer
	SetDatePickerElements(value unsafe.Pointer)
	DatePickerMode() unsafe.Pointer
	SetDatePickerMode(value unsafe.Pointer)
	DatePickerStyle() unsafe.Pointer
	SetDatePickerStyle(value unsafe.Pointer)
	DateValue() foundation.Date
	SetDateValue(value foundation.IDate)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Locale() foundation.Locale
	SetLocale(value foundation.ILocale)
	MinDate() foundation.Date
	SetMinDate(value foundation.IDate)
	TextColor() NSColor
	SetTextColor(value IColor)
	TimeInterval() unsafe.Pointer
	SetTimeInterval(value unsafe.Pointer)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.ITimeZone)
}

// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix.


// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell

type DatePickerCell struct {
	ActionCell
}

// DatePickerCellFrom constructs a [DatePickerCell] from an unsafe.Pointer.
//
// An object that controls the behavior of a date picker, or of a single date picker cell in a matrix.
func DatePickerCellFrom(ptr unsafe.Pointer) DatePickerCell {
	return DatePickerCell{
		ActionCell: ActionCellFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.Send[DatePickerCell](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DatePickerCellClass) New() DatePickerCell {
	rv := objc.Send[DatePickerCell](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DatePickerCell) Init() DatePickerCell {
	rv := objc.Send[DatePickerCell](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DatePickerCell) Autorelease() DatePickerCell {
	rv := objc.Send[DatePickerCell](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDatePickerCell creates a new DatePickerCell instance.
func NewDatePickerCell() DatePickerCell {
	return getDatePickerCellClass().New()
}



// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate

func (d_ DatePickerCell) MaxDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("maxDate"))
	return rv
}


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate

func (d_ DatePickerCell) SetMaxDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}


// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/backgroundcolor

func (d_ DatePickerCell) BackgroundColor() NSColor {
	rv := objc.Send[NSColor](d_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/backgroundcolor

func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/calendar

func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/calendar

func (d_ DatePickerCell) SetCalendar(value foundation.ICalendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickerelements

func (d_ DatePickerCell) DatePickerElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerElements"))
	return rv
}


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickerelements

func (d_ DatePickerCell) SetDatePickerElements(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
}


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickermode

func (d_ DatePickerCell) DatePickerMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerMode"))
	return rv
}


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickermode

func (d_ DatePickerCell) SetDatePickerMode(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerMode:"), value)
}


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickerstyle

func (d_ DatePickerCell) DatePickerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datepickerstyle

func (d_ DatePickerCell) SetDatePickerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datevalue

func (d_ DatePickerCell) DateValue() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("dateValue"))
	return rv
}


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datevalue

func (d_ DatePickerCell) SetDateValue(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/delegate

func (d_ DatePickerCell) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/delegate

func (d_ DatePickerCell) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/drawsbackground

func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/drawsbackground

func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/locale

func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/locale

func (d_ DatePickerCell) SetLocale(value foundation.ILocale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/mindate

func (d_ DatePickerCell) MinDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("minDate"))
	return rv
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/mindate

func (d_ DatePickerCell) SetMinDate(value foundation.IDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/textcolor

func (d_ DatePickerCell) TextColor() NSColor {
	rv := objc.Send[NSColor](d_.ID, objc.Sel("textColor"))
	return rv
}


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/textcolor

func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextColor:"), value)
}


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/timeinterval

func (d_ DatePickerCell) TimeInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("timeInterval"))
	return rv
}


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/timeinterval

func (d_ DatePickerCell) SetTimeInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/timezone

func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/timezone

func (d_ DatePickerCell) SetTimeZone(value foundation.ITimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}



