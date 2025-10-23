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
	// properties:
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	Calendar() foundation.objc.IObject /* cross-framework: Calendar */
	SetCalendar(value foundation.objc.IObject /* cross-framework: Calendar */)
	DatePickerElements() unsafe.Pointer
	SetDatePickerElements(value unsafe.Pointer)
	DatePickerMode() unsafe.Pointer
	SetDatePickerMode(value unsafe.Pointer)
	DatePickerStyle() unsafe.Pointer
	SetDatePickerStyle(value unsafe.Pointer)
	DateValue() foundation.objc.IObject /* cross-framework: Date */
	SetDateValue(value foundation.objc.IObject /* cross-framework: Date */)
	Delegate() DatePickerCellDelegate /* not a class type */
	SetDelegate(value DatePickerCellDelegate /* not a class type */)
	DrawsBackground() bool /* primitive/slice/pointer. */
	SetDrawsBackground(value bool /* primitive/slice/pointer. */)
	Locale() foundation.objc.IObject /* cross-framework: Locale */
	SetLocale(value foundation.objc.IObject /* cross-framework: Locale */)
	MaxDate() foundation.objc.IObject /* cross-framework: Date */
	SetMaxDate(value foundation.objc.IObject /* cross-framework: Date */)
	MinDate() foundation.objc.IObject /* cross-framework: Date */
	SetMinDate(value foundation.objc.IObject /* cross-framework: Date */)
	TextColor() IColor
	SetTextColor(value IColor)
	TimeInterval() unsafe.Pointer
	SetTimeInterval(value unsafe.Pointer)
	TimeZone() foundation.objc.IObject /* cross-framework: TimeZone */
	SetTimeZone(value foundation.objc.IObject /* cross-framework: TimeZone */)
	// methods:
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



// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/backgroundcolor
func (d_ DatePickerCell) BackgroundColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("backgroundColor"))
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
func (d_ DatePickerCell) Calendar() foundation.objc.IObject /* cross-framework: Calendar */ {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/calendar
func (d_ DatePickerCell) SetCalendar(value foundation.objc.IObject /* cross-framework: Calendar */) {
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
func (d_ DatePickerCell) DateValue() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("dateValue"))
	return rv
}


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/datevalue
func (d_ DatePickerCell) SetDateValue(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/delegate
func (d_ DatePickerCell) Delegate() DatePickerCellDelegate /* not a class type */ {
	rv := objc.Send[DatePickerCellDelegate](d_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/delegate
func (d_ DatePickerCell) SetDelegate(value DatePickerCellDelegate /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/drawsbackground
func (d_ DatePickerCell) DrawsBackground() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/drawsbackground
func (d_ DatePickerCell) SetDrawsBackground(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/locale
func (d_ DatePickerCell) Locale() foundation.objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/locale
func (d_ DatePickerCell) SetLocale(value foundation.objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/maxdate
func (d_ DatePickerCell) MaxDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("maxDate"))
	return rv
}


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/maxdate
func (d_ DatePickerCell) SetMaxDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/mindate
func (d_ DatePickerCell) MinDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("minDate"))
	return rv
}


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/mindate
func (d_ DatePickerCell) SetMinDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/textcolor
func (d_ DatePickerCell) TextColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("textColor"))
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
func (d_ DatePickerCell) TimeZone() foundation.objc.IObject /* cross-framework: TimeZone */ {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercell/timezone
func (d_ DatePickerCell) SetTimeZone(value foundation.objc.IObject /* cross-framework: TimeZone */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}



