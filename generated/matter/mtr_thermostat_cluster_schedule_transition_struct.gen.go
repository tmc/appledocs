// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterScheduleTransitionStruct] class.
var (
	MTRThermostatClusterScheduleTransitionStructClass     _MTRThermostatClusterScheduleTransitionStructClass
	MTRThermostatClusterScheduleTransitionStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleTransitionStructClass() _MTRThermostatClusterScheduleTransitionStructClass {
	MTRThermostatClusterScheduleTransitionStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleTransitionStructClass = _MTRThermostatClusterScheduleTransitionStructClass{objc.GetClass("MTRThermostatClusterScheduleTransitionStruct")}
	})
	return MTRThermostatClusterScheduleTransitionStructClass
}

type _MTRThermostatClusterScheduleTransitionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterScheduleTransitionStruct] class.
type IMTRThermostatClusterScheduleTransitionStruct interface {
	objectivec.IObject
	// properties:
	CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	DayOfWeek() objc.IObject /* cross-framework: NSNumber */
	SetDayOfWeek(value objc.IObject /* cross-framework: NSNumber */)
	HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */
	SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */)
	PresetHandle() objc.IObject /* cross-framework: NSData */
	SetPresetHandle(value objc.IObject /* cross-framework: NSData */)
	SystemMode() objc.IObject /* cross-framework: NSNumber */
	SetSystemMode(value objc.IObject /* cross-framework: NSNumber */)
	TransitionTime() objc.IObject /* cross-framework: NSNumber */
	SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct
type MTRThermostatClusterScheduleTransitionStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleTransitionStructFrom constructs a [MTRThermostatClusterScheduleTransitionStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleTransitionStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleTransitionStruct {
	return MTRThermostatClusterScheduleTransitionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleTransitionStructClass) Alloc() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterScheduleTransitionStructClass) New() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleTransitionStruct) Init() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleTransitionStruct) Autorelease() MTRThermostatClusterScheduleTransitionStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTransitionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleTransitionStruct creates a new MTRThermostatClusterScheduleTransitionStruct instance.
func NewMTRThermostatClusterScheduleTransitionStruct() MTRThermostatClusterScheduleTransitionStruct {
	return getMTRThermostatClusterScheduleTransitionStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/coolingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) CoolingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("coolingSetpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/coolingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetCoolingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCoolingSetpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/dayOfWeek
func (m_ MTRThermostatClusterScheduleTransitionStruct) DayOfWeek() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dayOfWeek"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/dayOfWeek
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetDayOfWeek(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeek:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/heatingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) HeatingSetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("heatingSetpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/heatingSetpoint
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetHeatingSetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHeatingSetpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/presetHandle
func (m_ MTRThermostatClusterScheduleTransitionStruct) PresetHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/presetHandle
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetPresetHandle(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/systemMode
func (m_ MTRThermostatClusterScheduleTransitionStruct) SystemMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("systemMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/systemMode
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetSystemMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/transitionTime
func (m_ MTRThermostatClusterScheduleTransitionStruct) TransitionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("transitionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTransitionStruct/transitionTime
func (m_ MTRThermostatClusterScheduleTransitionStruct) SetTransitionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitionTime:"), value)
}



