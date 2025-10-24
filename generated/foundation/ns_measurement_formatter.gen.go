// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMeasurementFormatter */


/* debug [class_header]: Header for NSMeasurementFormatter */
// The class instance for the [MeasurementFormatter] class.
var (
	MeasurementFormatterClass     _MeasurementFormatterClass
	MeasurementFormatterClassOnce sync.Once
)

func getMeasurementFormatterClass() _MeasurementFormatterClass {
	MeasurementFormatterClassOnce.Do(func() {
		MeasurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}
	})
	return MeasurementFormatterClass
}

type _MeasurementFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MeasurementFormatter */
// An interface definition for the [MeasurementFormatter] class.
type IMeasurementFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for MeasurementFormatter */
	// properties:
	Locale() ILocale
	SetLocale(value ILocale)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitOptions() MeasurementFormatterUnitOptions
	SetUnitOptions(value MeasurementFormatterUnitOptions)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MeasurementFormatter */
	// methods:
	StringFromUnit(unit IUnit) IString
	StringFromMeasurement(measurement IMeasurement) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MeasurementFormatter */
// Alloc allocates a new instance without initialization.
func (mc _MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MeasurementFormatterClass) New() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MeasurementFormatter) Init() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MeasurementFormatter) Autorelease() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeasurementFormatter creates a new MeasurementFormatter instance.
func NewMeasurementFormatter() MeasurementFormatter {
	return getMeasurementFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MeasurementFormatter */
// A formatter that provides localized representations of units and measurements.
//
// You use the method to create a localized representation of an object, and you use the method to create a localized representation of an object. The formatter takes into account the specified , , and when producing string representations of units and measurements.


// A formatter that provides localized representations of units and measurements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter
type MeasurementFormatter struct {
	Formatter
}

// MeasurementFormatterFrom constructs a [MeasurementFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized representations of units and measurements.
func MeasurementFormatterFrom(ptr unsafe.Pointer) MeasurementFormatter {
	return MeasurementFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MeasurementFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MeasurementFormatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MeasurementFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MeasurementFormatter */

// Creates and returns a localized string representation of the provided unit of measure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/string(from:)-4hwjz
func (m_ MeasurementFormatter) StringFromUnit(unit IUnit) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("stringFromUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: StringFromUnit */


// Creates and returns a localized string representation of the provided measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/string(from:)-wt9y
func (m_ MeasurementFormatter) StringFromMeasurement(measurement IMeasurement) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("stringFromMeasurement:"), measurement)
	return rv
}/* debug [instance_methods/method]: StringFromMeasurement */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MeasurementFormatter */

// The locale of the formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/locale
func (m_ MeasurementFormatter) Locale() ILocale {
	rv := objc.Send[Locale](m_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale of the formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/locale
func (m_ MeasurementFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// The number formatter used to format the quantity of a measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/numberFormatter
func (m_ MeasurementFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](m_.ID, objc.Sel("numberFormatter"))
	return rv
}/* debug [instance_properties/getter]: numberFormatter */


// The number formatter used to format the quantity of a measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/numberFormatter
func (m_ MeasurementFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberFormatter:"), value)
}/* debug [instance_properties/setter]: numberFormatter */


// The options for how the unit is formatted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitOptions-swift.property
func (m_ MeasurementFormatter) UnitOptions() MeasurementFormatterUnitOptions {
	rv := objc.Send[MeasurementFormatterUnitOptions](m_.ID, objc.Sel("unitOptions"))
	return rv
}/* debug [instance_properties/getter]: unitOptions */


// The options for how the unit is formatted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitOptions-swift.property
func (m_ MeasurementFormatter) SetUnitOptions(value MeasurementFormatterUnitOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitOptions:"), value)
}/* debug [instance_properties/setter]: unitOptions */


// The unit style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitStyle
func (m_ MeasurementFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](m_.ID, objc.Sel("unitStyle"))
	return rv
}/* debug [instance_properties/getter]: unitStyle */


// The unit style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitStyle
func (m_ MeasurementFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}/* debug [instance_properties/setter]: unitStyle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMeasurementFormatter */



