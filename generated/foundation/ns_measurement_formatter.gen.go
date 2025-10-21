// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MeasurementFormatter] class.
type IMeasurementFormatter interface {
	IFormatter
}

// A formatter that provides localized representations of units and measurements.
//
// You use the method to create a localized representation of an object, and you use the method to create a localized representation of an object. The formatter takes into account the specified , , and when producing string representations of units and measurements.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The locale of the formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/locale
func (m_ MeasurementFormatter) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale of the formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/locale
func (m_ MeasurementFormatter) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocale:"), value)
}

// The number formatter used to format the quantity of a measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/numberformatter
func (m_ MeasurementFormatter) NumberFormatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("numberFormatter"))
	return rv
}


// SetNumberFormatter sets the value of the numberFormatter property.
// The number formatter used to format the quantity of a measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/numberformatter
func (m_ MeasurementFormatter) SetNumberFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberFormatter:"), value)
}

// The options for how the unit is formatted.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/unitoptions-swift.property
func (m_ MeasurementFormatter) UnitOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("unitOptions"))
	return rv
}


// SetUnitOptions sets the value of the unitOptions property.
// The options for how the unit is formatted.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/unitoptions-swift.property
func (m_ MeasurementFormatter) SetUnitOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitOptions:"), value)
}

// The unit style.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/unitstyle
func (m_ MeasurementFormatter) UnitStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("unitStyle"))
	return rv
}


// SetUnitStyle sets the value of the unitStyle property.
// The unit style.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/measurementformatter/unitstyle
func (m_ MeasurementFormatter) SetUnitStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}



