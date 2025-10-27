// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	MaxDate() foundation.foundation.INSDate
	SetMaxDate(value foundation.foundation.INSDate)
	MinDate() foundation.foundation.INSDate
	SetMinDate(value foundation.foundation.INSDate)
	TextColor() IColor
	SetTextColor(value IColor)
	TimeInterval() float64
	SetTimeInterval(value float64)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DatePickerCellClass) Alloc() DatePickerCell {
	rv := objc.Send[DatePickerCell](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(textCell:)
func NewDatePickerCellTextCell(string_ foundation.foundation.INSString) DatePickerCell {
	instance := getDatePickerCellClass().Alloc()
	rv := objc.Send[DatePickerCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(coder:)
func NewDatePickerCellWithCoder(coder foundation.foundation.INSCoder) DatePickerCell {
	instance := getDatePickerCellClass().Alloc()
	rv := objc.Send[DatePickerCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}






















// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/backgroundColor
func (d_ DatePickerCell) BackgroundColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("backgroundColor"))
	return rv
}


// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/backgroundColor
func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), value)
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/calendar
func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/calendar
func (d_ DatePickerCell) SetCalendar(value foundation.Calendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerElements
func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := objc.Send[DatePickerElementFlags](d_.ID, objc.Sel("datePickerElements"))
	return rv
}


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerElements
func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
}


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerMode
func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := objc.Send[DatePickerMode](d_.ID, objc.Sel("datePickerMode"))
	return rv
}


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerMode
func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerMode:"), value)
}


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerStyle
func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := objc.Send[DatePickerStyle](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerStyle
func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/dateValue
func (d_ DatePickerCell) DateValue() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("dateValue"))
	return rv
}


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/dateValue
func (d_ DatePickerCell) SetDateValue(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/drawsBackground
func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/drawsBackground
func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/locale
func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/locale
func (d_ DatePickerCell) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate
func (d_ DatePickerCell) MaxDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("maxDate"))
	return rv
}


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate
func (d_ DatePickerCell) SetMaxDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/minDate
func (d_ DatePickerCell) MinDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("minDate"))
	return rv
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/minDate
func (d_ DatePickerCell) SetMinDate(value foundation.foundation.INSDate) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/textColor
func (d_ DatePickerCell) TextColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("textColor"))
	return rv
}


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/textColor
func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextColor:"), value)
}


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeInterval
func (d_ DatePickerCell) TimeInterval() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("timeInterval"))
	return rv
}


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeInterval
func (d_ DatePickerCell) SetTimeInterval(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeZone
func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeZone
func (d_ DatePickerCell) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}







