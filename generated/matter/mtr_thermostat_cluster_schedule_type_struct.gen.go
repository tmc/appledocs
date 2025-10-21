// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThermostatClusterScheduleTypeStruct] class.
var (
	MTRThermostatClusterScheduleTypeStructClass     _MTRThermostatClusterScheduleTypeStructClass
	MTRThermostatClusterScheduleTypeStructClassOnce sync.Once
)

func getMTRThermostatClusterScheduleTypeStructClass() _MTRThermostatClusterScheduleTypeStructClass {
	MTRThermostatClusterScheduleTypeStructClassOnce.Do(func() {
		MTRThermostatClusterScheduleTypeStructClass = _MTRThermostatClusterScheduleTypeStructClass{objc.GetClass("MTRThermostatClusterScheduleTypeStruct")}
	})
	return MTRThermostatClusterScheduleTypeStructClass
}

type _MTRThermostatClusterScheduleTypeStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThermostatClusterScheduleTypeStruct] class.
type IMTRThermostatClusterScheduleTypeStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct
type MTRThermostatClusterScheduleTypeStruct struct {
	objectivec.Object
}

// MTRThermostatClusterScheduleTypeStructFrom constructs a [MTRThermostatClusterScheduleTypeStruct] from an unsafe.Pointer.
func MTRThermostatClusterScheduleTypeStructFrom(ptr unsafe.Pointer) MTRThermostatClusterScheduleTypeStruct {
	return MTRThermostatClusterScheduleTypeStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThermostatClusterScheduleTypeStructClass) Alloc() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThermostatClusterScheduleTypeStructClass) New() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThermostatClusterScheduleTypeStruct) Init() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThermostatClusterScheduleTypeStruct) Autorelease() MTRThermostatClusterScheduleTypeStruct {
	rv := objc.Send[MTRThermostatClusterScheduleTypeStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThermostatClusterScheduleTypeStruct creates a new MTRThermostatClusterScheduleTypeStruct instance.
func NewMTRThermostatClusterScheduleTypeStruct() MTRThermostatClusterScheduleTypeStruct {
	return getMTRThermostatClusterScheduleTypeStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/numberOfSchedules
func (m_ MTRThermostatClusterScheduleTypeStruct) NumberOfSchedules() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("numberOfSchedules"))
	return rv
}


// SetNumberOfSchedules sets the value of the numberOfSchedules property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/numberOfSchedules
func (m_ MTRThermostatClusterScheduleTypeStruct) SetNumberOfSchedules(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfSchedules:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/scheduleTypeFeatures
func (m_ MTRThermostatClusterScheduleTypeStruct) ScheduleTypeFeatures() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("scheduleTypeFeatures"))
	return rv
}


// SetScheduleTypeFeatures sets the value of the scheduleTypeFeatures property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/scheduleTypeFeatures
func (m_ MTRThermostatClusterScheduleTypeStruct) SetScheduleTypeFeatures(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScheduleTypeFeatures:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/systemMode
func (m_ MTRThermostatClusterScheduleTypeStruct) SystemMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("systemMode"))
	return rv
}


// SetSystemMode sets the value of the systemMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThermostatClusterScheduleTypeStruct/systemMode
func (m_ MTRThermostatClusterScheduleTypeStruct) SetSystemMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSystemMode:"), value)
}



