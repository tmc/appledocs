// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/mfgCode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/mfgCode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/value
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct/value
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



