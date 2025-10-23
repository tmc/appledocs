// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EnergyFormatter] class.
var (
	EnergyFormatterClass     _EnergyFormatterClass
	EnergyFormatterClassOnce sync.Once
)

func getEnergyFormatterClass() _EnergyFormatterClass {
	EnergyFormatterClassOnce.Do(func() {
		EnergyFormatterClass = _EnergyFormatterClass{objc.GetClass("NSEnergyFormatter")}
	})
	return EnergyFormatterClass
}

type _EnergyFormatterClass struct {
	class objc.Class
}

// An interface definition for the [EnergyFormatter] class.
type IEnergyFormatter interface {
	IFormatter
	GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ string, error_ string) bool
	StringFromJoules(numberInJoules float64) IString
	StringFromValueUnit(value float64, unit NSEnergyFormatterUnit) IString
	UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp NSEnergyFormatterUnit) IString
	UnitStringFromValueUnit(value float64, unit NSEnergyFormatterUnit) IString
	ForFoodEnergyUse() bool
	SetForFoodEnergyUse(value bool)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() NSFormattingUnitStyle
	SetUnitStyle(value NSFormattingUnitStyle)
	IsForFoodEnergyUse() bool
	SetIsForFoodEnergyUse(value bool)
}

// A formatter that provides localized descriptions of energy values.


// A formatter that provides localized descriptions of energy values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter
type EnergyFormatter struct {
	Formatter
}

// EnergyFormatterFrom constructs a [EnergyFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of energy values.
func EnergyFormatterFrom(ptr unsafe.Pointer) EnergyFormatter {
	return EnergyFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EnergyFormatterClass) Alloc() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EnergyFormatterClass) New() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EnergyFormatter) Init() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EnergyFormatter) Autorelease() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnergyFormatter creates a new EnergyFormatter instance.
func NewEnergyFormatter() EnergyFormatter {
	return getEnergyFormatterClass().New()
}



// This method is not supported for the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/getObjectValue(_:for:errorDescription:)
func (e_ EnergyFormatter) GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ string, error_ string) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, objc.String(string_), objc.String(error_))
	return rv
}


// Returns an energy string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromJoules:)
func (e_ EnergyFormatter) StringFromJoules(numberInJoules float64) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("stringFromJoules:"), numberInJoules)
	return rv
}


// Returns a properly formatted energy string for the given value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromValue:unit:)
func (e_ EnergyFormatter) StringFromValueUnit(value float64, unit NSEnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}


// Returns the unit string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromJoules:usedUnit:)
func (e_ EnergyFormatter) UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp NSEnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("unitStringFromJoules:usedUnit:"), numberInJoules, unitp)
	return rv
}


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromValue:unit:)
func (e_ EnergyFormatter) UnitStringFromValueUnit(value float64, unit NSEnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/isForFoodEnergyUse
func (e_ EnergyFormatter) ForFoodEnergyUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("forFoodEnergyUse"))
	return rv
}


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/isForFoodEnergyUse
func (e_ EnergyFormatter) SetForFoodEnergyUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setForFoodEnergyUse:"), value)
}


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/numberFormatter
func (e_ EnergyFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NSNumberFormatter](e_.ID, objc.Sel("numberFormatter"))
	return rv
}


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/numberFormatter
func (e_ EnergyFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNumberFormatter:"), value)
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitStyle
func (e_ EnergyFormatter) UnitStyle() NSFormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](e_.ID, objc.Sel("unitStyle"))
	return rv
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitStyle
func (e_ EnergyFormatter) SetUnitStyle(value NSFormattingUnitStyle) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUnitStyle:"), value)
}


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/isforfoodenergyuse
func (e_ EnergyFormatter) IsForFoodEnergyUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isForFoodEnergyUse"))
	return rv
}


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/isforfoodenergyuse
func (e_ EnergyFormatter) SetIsForFoodEnergyUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsForFoodEnergyUse:"), value)
}



