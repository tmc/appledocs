// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EnergyFormatter] class.
var (
	energyFormatterClass     _EnergyFormatterClass
	energyFormatterClassOnce sync.Once
)

func getEnergyFormatterClass() _EnergyFormatterClass {
	energyFormatterClassOnce.Do(func() {
		energyFormatterClass = _EnergyFormatterClass{objc.GetClass("NSEnergyFormatter")}
	})
	return energyFormatterClass
}

type _EnergyFormatterClass struct {
	class objc.Class
}

// An interface definition for the [EnergyFormatter] class.
type IEnergyFormatter interface {
	IFormatter
}

// A formatter that provides localized descriptions of energy values.
//
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




