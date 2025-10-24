// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	NumberOfPresets() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfPresets(value objc.IObject /* cross-framework: NSNumber */)
	PresetScenario() objc.IObject /* cross-framework: NSNumber */
	SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */)
	PresetTypeFeatures() objc.IObject /* cross-framework: NSNumber */
	SetPresetTypeFeatures(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/numberOfPresets
func (m_ MTRThermostatClusterPresetTypeStruct) NumberOfPresets() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfPresets"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/numberOfPresets
func (m_ MTRThermostatClusterPresetTypeStruct) SetNumberOfPresets(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfPresets:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetScenario
func (m_ MTRThermostatClusterPresetTypeStruct) PresetScenario() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetScenario"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetScenario
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) PresetTypeFeatures() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetTypeFeatures"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetTypeStruct/presetTypeFeatures
func (m_ MTRThermostatClusterPresetTypeStruct) SetPresetTypeFeatures(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetTypeFeatures:"), value)
}



