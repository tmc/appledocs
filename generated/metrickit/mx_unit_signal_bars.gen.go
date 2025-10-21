// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MXUnitSignalBars] class.
var (
	MXUnitSignalBarsClass     _MXUnitSignalBarsClass
	MXUnitSignalBarsClassOnce sync.Once
)

func getMXUnitSignalBarsClass() _MXUnitSignalBarsClass {
	MXUnitSignalBarsClassOnce.Do(func() {
		MXUnitSignalBarsClass = _MXUnitSignalBarsClass{objc.GetClass("MXUnitSignalBars")}
	})
	return MXUnitSignalBarsClass
}

type _MXUnitSignalBarsClass struct {
	class objc.Class
}

// An interface definition for the [MXUnitSignalBars] class.
type IMXUnitSignalBars interface {
	objectivec.IObject
}

// A unit of measure for the number of bars of cellular network connectivity.
//
// Cellular connectivity measures the relative strength of the device’s signal reception in decibels, which usually falls in a range of 0 to -110. defines the base unit as bars corresponding to the bars in the cellular connection status icon at the top of a device screen.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars
type MXUnitSignalBars struct {
	objectivec.Object
}

// MXUnitSignalBarsFrom constructs a [MXUnitSignalBars] from an unsafe.Pointer.
//
// A unit of measure for the number of bars of cellular network connectivity.
func MXUnitSignalBarsFrom(ptr unsafe.Pointer) MXUnitSignalBars {
	return MXUnitSignalBars{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXUnitSignalBarsClass) Alloc() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXUnitSignalBarsClass) New() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXUnitSignalBars) Init() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXUnitSignalBars) Autorelease() MXUnitSignalBars {
	rv := objc.Send[MXUnitSignalBars](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXUnitSignalBars creates a new MXUnitSignalBars instance.
func NewMXUnitSignalBars() MXUnitSignalBars {
	return getMXUnitSignalBarsClass().New()
}


// The number of bars of connectivity to the cellular network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars/bars
func (mc _MXUnitSignalBarsClass) Bars() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("bars"))
	return rv
}
// The number of bars of connectivity to the cellular network.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXUnitSignalBars/bars
func (m_ MXUnitSignalBars) Bars() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("bars"))
	return rv
}




