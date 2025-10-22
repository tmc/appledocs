// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AmbientPressureData] class.
var (
	AmbientPressureDataClass     _AmbientPressureDataClass
	AmbientPressureDataClassOnce sync.Once
)

func getAmbientPressureDataClass() _AmbientPressureDataClass {
	AmbientPressureDataClassOnce.Do(func() {
		AmbientPressureDataClass = _AmbientPressureDataClass{objc.GetClass("CMAmbientPressureData")}
	})
	return AmbientPressureDataClass
}

type _AmbientPressureDataClass struct {
	class objc.Class
}

// An interface definition for the [AmbientPressureData] class.
type IAmbientPressureData interface {
	ILogItem
	Pressure() unsafe.Pointer
	Temperature() unsafe.Pointer
}

// A measurement of the ambient pressure and temperature.


// A measurement of the ambient pressure and temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData

type AmbientPressureData struct {
	LogItem
}

// AmbientPressureDataFrom constructs a [AmbientPressureData] from an unsafe.Pointer.
//
// A measurement of the ambient pressure and temperature.
func AmbientPressureDataFrom(ptr unsafe.Pointer) AmbientPressureData {
	return AmbientPressureData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AmbientPressureDataClass) Alloc() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AmbientPressureDataClass) New() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AmbientPressureData) Init() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AmbientPressureData) Autorelease() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAmbientPressureData creates a new AmbientPressureData instance.
func NewAmbientPressureData() AmbientPressureData {
	return getAmbientPressureDataClass().New()
}



// The ambient pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData/pressure

func (a_ AmbientPressureData) Pressure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("pressure"))
	return rv
}


// The temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData/temperature

func (a_ AmbientPressureData) Temperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("temperature"))
	return rv
}



