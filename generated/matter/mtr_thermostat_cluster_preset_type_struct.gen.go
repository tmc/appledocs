// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThermostatClusterPresetTypeStruct] class.
var (
	MTRThermostatClusterPresetTypeStructClass     _MTRThermostatClusterPresetTypeStructClass
	MTRThermostatClusterPresetTypeStructClassOnce sync.Once
)

func getMTRThermostatClusterPresetTypeStructClass() _MTRThermostatClusterPresetTypeStructClass {
	MTRThermostatClusterPresetTypeStructClassOnce.Do(func() {
		MTRThermostatClusterPresetTypeStructClass = _MTRThermostatClusterPresetTypeStructClass{objc.GetClass("MTRThermostatClusterPresetTypeStruct")}
	})
	return MTRThermostatClusterPresetTypeStructClass
}

type _MTRThermostatClusterPresetTypeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterPresetTypeStruct] class.
type IMTRThermostatClusterPresetTypeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct
type MTRThermostatClusterPresetTypeStruct struct {
	objectivec.Object
}

// MTRThermostatClusterPresetTypeStructFrom constructs a [MTRThermostatClusterPresetTypeStruct] from an unsafe.Pointer.
func MTRThermostatClusterPresetTypeStructFrom(ptr unsafe.Pointer) MTRThermostatClusterPresetTypeStruct {
	return MTRThermostatClusterPresetTypeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterPresetTypeStructClass) Alloc() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterPresetTypeStructClass) New() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterPresetTypeStruct) Init() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterPresetTypeStruct) Autorelease() MTRThermostatClusterPresetTypeStruct {
	rv := objc.Send[MTRThermostatClusterPresetTypeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterPresetTypeStruct creates a new MTRThermostatClusterPresetTypeStruct instance.
func NewMTRThermostatClusterPresetTypeStruct() MTRThermostatClusterPresetTypeStruct {
	return getMTRThermostatClusterPresetTypeStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/numberOfPresets
func (m_ MTRThermostatClusterPresetTypeStruct) NumberOfPresets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("numberOfPresets"))
	return rv
}


// SetNumberOfPresets sets the value of the numberOfPresets property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/numberOfPresets
func (m_ MTRThermostatClusterPresetTypeStruct) SetNumberOfPresets(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfPresets:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetScenario
func (m_ MTRThermostatClusterPresetTypeStruct) PresetScenario() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presetScenario"))
	return rv
}


// SetPresetScenario sets the value of the presetScenario property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetScenario
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetScenario(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) PresetTypeFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presetTypeFeatures"))
	return rv
}


// SetPresetTypeFeatures sets the value of the presetTypeFeatures property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetTypeFeatures(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetTypeFeatures:"), value)
}


