// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterPresetStruct] class.
var (
	MTRThermostatClusterPresetStructClass     _MTRThermostatClusterPresetStructClass
	MTRThermostatClusterPresetStructClassOnce sync.Once
)

func getMTRThermostatClusterPresetStructClass() _MTRThermostatClusterPresetStructClass {
	MTRThermostatClusterPresetStructClassOnce.Do(func() {
		MTRThermostatClusterPresetStructClass = _MTRThermostatClusterPresetStructClass{objc.GetClass("MTRThermostatClusterPresetStruct")}
	})
	return MTRThermostatClusterPresetStructClass
}

type _MTRThermostatClusterPresetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterPresetStruct] class.
type IMTRThermostatClusterPresetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct
type MTRThermostatClusterPresetStruct struct {
	objectivec.Object
}

// MTRThermostatClusterPresetStructFrom constructs a [MTRThermostatClusterPresetStruct] from an unsafe.Pointer.
func MTRThermostatClusterPresetStructFrom(ptr unsafe.Pointer) MTRThermostatClusterPresetStruct {
	return MTRThermostatClusterPresetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterPresetStructClass) Alloc() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterPresetStructClass) New() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterPresetStruct) Init() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterPresetStruct) Autorelease() MTRThermostatClusterPresetStruct {
	rv := objc.Send[MTRThermostatClusterPresetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterPresetStruct creates a new MTRThermostatClusterPresetStruct instance.
func NewMTRThermostatClusterPresetStruct() MTRThermostatClusterPresetStruct {
	return getMTRThermostatClusterPresetStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) BuiltIn() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("builtIn"))
	return rv
}


// SetBuiltIn sets the value of the builtIn property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) SetBuiltIn(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/coolingSetpoint
func (m_ MTRThermostatClusterPresetStruct) CoolingSetpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("coolingSetpoint"))
	return rv
}


// SetCoolingSetpoint sets the value of the coolingSetpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/coolingSetpoint
func (m_ MTRThermostatClusterPresetStruct) SetCoolingSetpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolingSetpoint:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/heatingSetpoint
func (m_ MTRThermostatClusterPresetStruct) HeatingSetpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("heatingSetpoint"))
	return rv
}


// SetHeatingSetpoint sets the value of the heatingSetpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/heatingSetpoint
func (m_ MTRThermostatClusterPresetStruct) SetHeatingSetpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatingSetpoint:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/name
func (m_ MTRThermostatClusterPresetStruct) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/name
func (m_ MTRThermostatClusterPresetStruct) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetHandle
func (m_ MTRThermostatClusterPresetStruct) PresetHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// SetPresetHandle sets the value of the presetHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetHandle
func (m_ MTRThermostatClusterPresetStruct) SetPresetHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetScenario
func (m_ MTRThermostatClusterPresetStruct) PresetScenario() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presetScenario"))
	return rv
}


// SetPresetScenario sets the value of the presetScenario property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetScenario
func (m_ MTRThermostatClusterPresetStruct) SetPresetScenario(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}


