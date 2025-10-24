// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	IsForFoodEnergyUse() bool
	SetIsForFoodEnergyUse(value bool)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() unsafe.Pointer
	SetUnitStyle(value unsafe.Pointer)
	// methods:
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


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/numberformatter
func (e_ EnergyFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](e_.ID, objc.Sel("numberFormatter"))
	return rv
}


// The number formatter used to format the numbers in energy strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/numberformatter
func (e_ EnergyFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setNumberFormatter:"), value)
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/unitstyle
func (e_ EnergyFormatter) UnitStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("unitStyle"))
	return rv
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/energyformatter/unitstyle
func (e_ EnergyFormatter) SetUnitStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setUnitStyle:"), value)
}



