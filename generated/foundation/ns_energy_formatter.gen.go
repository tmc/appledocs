// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSEnergyFormatter */


/* debug [class_header]: Header for NSEnergyFormatter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EnergyFormatter */
// An interface definition for the [EnergyFormatter] class.
type IEnergyFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for EnergyFormatter */
	// properties:
	ForFoodEnergyUse() bool
	SetForFoodEnergyUse(value bool)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	IsForFoodEnergyUse() bool
	SetIsForFoodEnergyUse(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EnergyFormatter */
	// methods:
	GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool
	StringFromJoules(numberInJoules float64) IString
	StringFromValueUnit(value float64, unit EnergyFormatterUnit) IString
	UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp EnergyFormatterUnit) IString
	UnitStringFromValueUnit(value float64, unit EnergyFormatterUnit) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EnergyFormatter */
// Alloc allocates a new instance without initialization.
func (ec _EnergyFormatterClass) Alloc() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EnergyFormatter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EnergyFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EnergyFormatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EnergyFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EnergyFormatter */

// This method is not supported for the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/getObjectValue(_:for:errorDescription:)
func (e_ EnergyFormatter) GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}/* debug [instance_methods/method]: GetObjectValueForStringErrorDescription */


// Returns an energy string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromJoules:)
func (e_ EnergyFormatter) StringFromJoules(numberInJoules float64) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("stringFromJoules:"), numberInJoules)
	return rv
}/* debug [instance_methods/method]: StringFromJoules */


// Returns a properly formatted energy string for the given value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromValue:unit:)
func (e_ EnergyFormatter) StringFromValueUnit(value float64, unit EnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}/* debug [instance_methods/method]: StringFromValueUnit */


// Returns the unit string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromJoules:usedUnit:)
func (e_ EnergyFormatter) UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp EnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("unitStringFromJoules:usedUnit:"), numberInJoules, unitp)
	return rv
}/* debug [instance_methods/method]: UnitStringFromJoulesUsedUnit */


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromValue:unit:)
func (e_ EnergyFormatter) UnitStringFromValueUnit(value float64, unit EnergyFormatterUnit) IString {
	rv := objc.Send[String](e_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}/* debug [instance_methods/method]: UnitStringFromValueUnit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EnergyFormatter */

// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/isForFoodEnergyUse
func (e_ EnergyFormatter) ForFoodEnergyUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("forFoodEnergyUse"))
	return rv
}/* debug [instance_properties/getter]: forFoodEnergyUse */


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/isForFoodEnergyUse
func (e_ EnergyFormatter) SetForFoodEnergyUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setForFoodEnergyUse:"), value)
}/* debug [instance_properties/setter]: forFoodEnergyUse */


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/numberFormatter
func (e_ EnergyFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](e_.ID, objc.Sel("numberFormatter"))
	return rv
}/* debug [instance_properties/getter]: numberFormatter */


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/numberFormatter
func (e_ EnergyFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNumberFormatter:"), value)
}/* debug [instance_properties/setter]: numberFormatter */


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitStyle
func (e_ EnergyFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](e_.ID, objc.Sel("unitStyle"))
	return rv
}/* debug [instance_properties/getter]: unitStyle */


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitStyle
func (e_ EnergyFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUnitStyle:"), value)
}/* debug [instance_properties/setter]: unitStyle */


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/isforfoodenergyuse
func (e_ EnergyFormatter) IsForFoodEnergyUse() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isForFoodEnergyUse"))
	return rv
}/* debug [instance_properties/getter]: isForFoodEnergyUse */


// A Boolean value that indicates whether the energy value is used to measure food energy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/isforfoodenergyuse
func (e_ EnergyFormatter) SetIsForFoodEnergyUse(value bool) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIsForFoodEnergyUse:"), value)
}/* debug [instance_properties/setter]: isForFoodEnergyUse */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSEnergyFormatter */



