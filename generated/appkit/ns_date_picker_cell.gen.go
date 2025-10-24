// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDatePickerCell */


/* debug [class_header]: Header for NSDatePickerCell */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DatePickerCell */
// An interface definition for the [DatePickerCell] class.
type IDatePickerCell interface {
	IActionCell
	
/* debug [class_interface_properties]: Properties for DatePickerCell */
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
	DateValue() objc.IObject /* cross-framework: NSDate */
	SetDateValue(value objc.IObject /* cross-framework: NSDate */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	Locale() foundation.Locale
	SetLocale(value foundation.Locale)
	MaxDate() objc.IObject /* cross-framework: NSDate */
	SetMaxDate(value objc.IObject /* cross-framework: NSDate */)
	MinDate() objc.IObject /* cross-framework: NSDate */
	SetMinDate(value objc.IObject /* cross-framework: NSDate */)
	TextColor() IColor
	SetTextColor(value IColor)
	TimeInterval() float64
	SetTimeInterval(value float64)
	TimeZone() foundation.TimeZone
	SetTimeZone(value foundation.TimeZone)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DatePickerCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DatePickerCell */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DatePickerCell */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DatePickerCell */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(textCell:)
func NewDatePickerCellTextCell(string_ objc.IObject /* cross-framework: NSString */) DatePickerCell {
	instance := getDatePickerCellClass().Alloc()
	rv := objc.Send[DatePickerCell](instance.ID, objc.Sel("initTextCell:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDatePickerCellTextCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/init(coder:)
func NewDatePickerCellWithCoder(coder foundation.Coder) DatePickerCell {
	instance := getDatePickerCellClass().Alloc()
	rv := objc.Send[DatePickerCell](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDatePickerCellWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DatePickerCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DatePickerCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DatePickerCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DatePickerCell */

// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/backgroundColor
func (d_ DatePickerCell) BackgroundColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The cell’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/backgroundColor
func (d_ DatePickerCell) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/calendar
func (d_ DatePickerCell) Calendar() foundation.Calendar {
	rv := objc.Send[foundation.Calendar](d_.ID, objc.Sel("calendar"))
	return rv
}/* debug [instance_properties/getter]: calendar */


// The calendar used by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/calendar
func (d_ DatePickerCell) SetCalendar(value foundation.Calendar) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCalendar:"), value)
}/* debug [instance_properties/setter]: calendar */


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerElements
func (d_ DatePickerCell) DatePickerElements() DatePickerElementFlags {
	rv := objc.Send[DatePickerElementFlags](d_.ID, objc.Sel("datePickerElements"))
	return rv
}/* debug [instance_properties/getter]: datePickerElements */


// A bitmask that indicates which visual elements are shown by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerElements
func (d_ DatePickerCell) SetDatePickerElements(value DatePickerElementFlags) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerElements:"), value)
}/* debug [instance_properties/setter]: datePickerElements */


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerMode
func (d_ DatePickerCell) DatePickerMode() DatePickerMode {
	rv := objc.Send[DatePickerMode](d_.ID, objc.Sel("datePickerMode"))
	return rv
}/* debug [instance_properties/getter]: datePickerMode */


// The mode in use by the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerMode
func (d_ DatePickerCell) SetDatePickerMode(value DatePickerMode) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerMode:"), value)
}/* debug [instance_properties/setter]: datePickerMode */


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerStyle
func (d_ DatePickerCell) DatePickerStyle() DatePickerStyle {
	rv := objc.Send[DatePickerStyle](d_.ID, objc.Sel("datePickerStyle"))
	return rv
}/* debug [instance_properties/getter]: datePickerStyle */


// The date picker style to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/datePickerStyle
func (d_ DatePickerCell) SetDatePickerStyle(value DatePickerStyle) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDatePickerStyle:"), value)
}/* debug [instance_properties/setter]: datePickerStyle */


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/dateValue
func (d_ DatePickerCell) DateValue() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("dateValue"))
	return rv
}/* debug [instance_properties/getter]: dateValue */


// The date currently specified in the picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/dateValue
func (d_ DatePickerCell) SetDateValue(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDateValue:"), value)
}/* debug [instance_properties/setter]: dateValue */


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/delegate
func (d_ DatePickerCell) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate associated with the date picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/delegate
func (d_ DatePickerCell) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/drawsBackground
func (d_ DatePickerCell) DrawsBackground() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean value indicating whether the cell draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/drawsBackground
func (d_ DatePickerCell) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/locale
func (d_ DatePickerCell) Locale() foundation.Locale {
	rv := objc.Send[foundation.Locale](d_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale used to display dates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/locale
func (d_ DatePickerCell) SetLocale(value foundation.Locale) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate
func (d_ DatePickerCell) MaxDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("maxDate"))
	return rv
}/* debug [instance_properties/getter]: maxDate */


// The maximum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/maxDate
func (d_ DatePickerCell) SetMaxDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaxDate:"), value)
}/* debug [instance_properties/setter]: maxDate */


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/minDate
func (d_ DatePickerCell) MinDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("minDate"))
	return rv
}/* debug [instance_properties/getter]: minDate */


// The minimum date that the picker allows as input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/minDate
func (d_ DatePickerCell) SetMinDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMinDate:"), value)
}/* debug [instance_properties/setter]: minDate */


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/textColor
func (d_ DatePickerCell) TextColor() IColor {
	rv := objc.Send[Color](d_.ID, objc.Sel("textColor"))
	return rv
}/* debug [instance_properties/getter]: textColor */


// The cell’s text color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/textColor
func (d_ DatePickerCell) SetTextColor(value IColor) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTextColor:"), value)
}/* debug [instance_properties/setter]: textColor */


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeInterval
func (d_ DatePickerCell) TimeInterval() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("timeInterval"))
	return rv
}/* debug [instance_properties/getter]: timeInterval */


// The time interval that represents the date range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeInterval
func (d_ DatePickerCell) SetTimeInterval(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeInterval:"), value)
}/* debug [instance_properties/setter]: timeInterval */


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeZone
func (d_ DatePickerCell) TimeZone() foundation.TimeZone {
	rv := objc.Send[foundation.TimeZone](d_.ID, objc.Sel("timeZone"))
	return rv
}/* debug [instance_properties/getter]: timeZone */


// The time zone used to display time-related values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePickerCell/timeZone
func (d_ DatePickerCell) SetTimeZone(value foundation.TimeZone) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTimeZone:"), value)
}/* debug [instance_properties/setter]: timeZone */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDatePickerCell */


