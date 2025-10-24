// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MassFormatter] class.
var (
	MassFormatterClass     _MassFormatterClass
	MassFormatterClassOnce sync.Once
)

func getMassFormatterClass() _MassFormatterClass {
	MassFormatterClassOnce.Do(func() {
		MassFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}
	})
	return MassFormatterClass
}

type _MassFormatterClass struct {
	class objc.Class
}

// An interface definition for the [MassFormatter] class.
type IMassFormatter interface {
	IFormatter
	// properties:
	ForPersonMassUse() bool
	SetForPersonMassUse(value bool)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	IsForPersonMassUse() bool
	SetIsForPersonMassUse(value bool)
	// methods:
	GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool
	StringFromKilograms(numberInKilograms float64) IString
	StringFromValueUnit(value float64, unit MassFormatterUnit) IString
	UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp MassFormatterUnit) IString
	UnitStringFromValueUnit(value float64, unit MassFormatterUnit) IString
}

// A formatter that provides localized descriptions of mass and weight values.


// A formatter that provides localized descriptions of mass and weight values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter
type MassFormatter struct {
	Formatter
}

// MassFormatterFrom constructs a [MassFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of mass and weight values.
func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MassFormatterClass) Alloc() MassFormatter {
	rv := objc.Send[MassFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MassFormatterClass) New() MassFormatter {
	rv := objc.Send[MassFormatter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MassFormatter) Init() MassFormatter {
	rv := objc.Send[MassFormatter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MassFormatter) Autorelease() MassFormatter {
	rv := objc.Send[MassFormatter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMassFormatter creates a new MassFormatter instance.
func NewMassFormatter() MassFormatter {
	return getMassFormatterClass().New()
}



// This method is not supported for the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/getObjectValue(_:for:errorDescription:)
func (m_ MassFormatter) GetObjectValueForStringErrorDescription(obj unsafe.Pointer, string_ IString, error_ IString) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}


// Returns a mass string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/string(fromKilograms:)
func (m_ MassFormatter) StringFromKilograms(numberInKilograms float64) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("stringFromKilograms:"), numberInKilograms)
	return rv
}


// Returns a properly formatted mass string for the given value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/string(fromValue:unit:)
func (m_ MassFormatter) StringFromValueUnit(value float64, unit MassFormatterUnit) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}


// Returns the unit string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitString(fromKilograms:usedUnit:)
func (m_ MassFormatter) UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp MassFormatterUnit) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("unitStringFromKilograms:usedUnit:"), numberInKilograms, unitp)
	return rv
}


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitString(fromValue:unit:)
func (m_ MassFormatter) UnitStringFromValueUnit(value float64, unit MassFormatterUnit) IString {
	rv := objc.Send[String](m_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/isForPersonMassUse
func (m_ MassFormatter) ForPersonMassUse() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("forPersonMassUse"))
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/isForPersonMassUse
func (m_ MassFormatter) SetForPersonMassUse(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForPersonMassUse:"), value)
}


// The number formatter used to format the numbers in a mass strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/numberFormatter
func (m_ MassFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](m_.ID, objc.Sel("numberFormatter"))
	return rv
}


// The number formatter used to format the numbers in a mass strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/numberFormatter
func (m_ MassFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberFormatter:"), value)
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitStyle
func (m_ MassFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](m_.ID, objc.Sel("unitStyle"))
	return rv
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitStyle
func (m_ MassFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) IsForPersonMassUse() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isForPersonMassUse"))
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) SetIsForPersonMassUse(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForPersonMassUse:"), value)
}



