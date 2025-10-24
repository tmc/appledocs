// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	BuiltIn() objc.IObject /* cross-framework: NSNumber */
	SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */)
	CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PresetHandle() objc.IObject /* cross-framework: NSData */
	SetPresetHandle(value objc.IObject /* cross-framework: NSData */)
	PresetScenario() objc.IObject /* cross-framework: NSNumber */
	SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) BuiltIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("builtIn"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/builtIn
func (m_ MTRThermostatClusterPresetStruct) SetBuiltIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/coolingSetpoint
func (m_ MTRThermostatClusterPresetStruct) CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("coolingSetpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/coolingSetpoint
func (m_ MTRThermostatClusterPresetStruct) SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolingSetpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/heatingSetpoint
func (m_ MTRThermostatClusterPresetStruct) HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("heatingSetpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/heatingSetpoint
func (m_ MTRThermostatClusterPresetStruct) SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatingSetpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/name
func (m_ MTRThermostatClusterPresetStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/name
func (m_ MTRThermostatClusterPresetStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetHandle
func (m_ MTRThermostatClusterPresetStruct) PresetHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetHandle
func (m_ MTRThermostatClusterPresetStruct) SetPresetHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetScenario
func (m_ MTRThermostatClusterPresetStruct) PresetScenario() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("presetScenario"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterPresetStruct/presetScenario
func (m_ MTRThermostatClusterPresetStruct) SetPresetScenario(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetScenario:"), value)
}



