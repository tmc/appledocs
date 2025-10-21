// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterScheduleStruct] class.
var (
	MTRThermostatClusterScheduleStructClass     _MTRThermostatClusterScheduleStructClass
	MTRThermostatClusterScheduleStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleStructClass() _MTRThermostatClusterScheduleStructClass {
	MTRThermostatClusterScheduleStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleStructClass = _MTRThermostatClusterScheduleStructClass{objc.GetClass("MTRThermostatClusterScheduleStruct")}
	})
	return MTRThermostatClusterScheduleStructClass
}

type _MTRThermostatClusterScheduleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterScheduleStruct] class.
type IMTRThermostatClusterScheduleStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct
type MTRThermostatClusterScheduleStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleStructFrom constructs a [MTRThermostatClusterScheduleStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleStruct {
	return MTRThermostatClusterScheduleStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleStructClass) Alloc() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterScheduleStructClass) New() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleStruct) Init() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleStruct) Autorelease() MTRThermostatClusterScheduleStruct {
	rv := objc.Send[MTRThermostatClusterScheduleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleStruct creates a new MTRThermostatClusterScheduleStruct instance.
func NewMTRThermostatClusterScheduleStruct() MTRThermostatClusterScheduleStruct {
	return getMTRThermostatClusterScheduleStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/builtIn
func (m_ MTRThermostatClusterScheduleStruct) BuiltIn() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("builtIn"))
	return rv
}


// SetBuiltIn sets the value of the builtIn property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/builtIn
func (m_ MTRThermostatClusterScheduleStruct) SetBuiltIn(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBuiltIn:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/name
func (m_ MTRThermostatClusterScheduleStruct) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/presetHandle
func (m_ MTRThermostatClusterScheduleStruct) PresetHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("presetHandle"))
	return rv
}


// SetPresetHandle sets the value of the presetHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/presetHandle
func (m_ MTRThermostatClusterScheduleStruct) SetPresetHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPresetHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/scheduleHandle
func (m_ MTRThermostatClusterScheduleStruct) ScheduleHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("scheduleHandle"))
	return rv
}


// SetScheduleHandle sets the value of the scheduleHandle property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/scheduleHandle
func (m_ MTRThermostatClusterScheduleStruct) SetScheduleHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/systemMode
func (m_ MTRThermostatClusterScheduleStruct) SystemMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("systemMode"))
	return rv
}


// SetSystemMode sets the value of the systemMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/systemMode
func (m_ MTRThermostatClusterScheduleStruct) SetSystemMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/transitions
func (m_ MTRThermostatClusterScheduleStruct) Transitions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("transitions"))
	return rv
}


// SetTransitions sets the value of the transitions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleStruct/transitions
func (m_ MTRThermostatClusterScheduleStruct) SetTransitions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTransitions:"), value)
}



