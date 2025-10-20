// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MeasurementFormatter] class.
var (
	measurementFormatterClass     _MeasurementFormatterClass
	measurementFormatterClassOnce sync.Once
)

func getMeasurementFormatterClass() _MeasurementFormatterClass {
	measurementFormatterClassOnce.Do(func() {
		measurementFormatterClass = _MeasurementFormatterClass{objc.GetClass("NSMeasurementFormatter")}
	})
	return measurementFormatterClass
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




