// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEClusterChargingTargetScheduleStruct] class.
var (
	MTREnergyEVSEClusterChargingTargetScheduleStructClass     _MTREnergyEVSEClusterChargingTargetScheduleStructClass
	MTREnergyEVSEClusterChargingTargetScheduleStructClassOnce sync.Once
)

func getMTREnergyEVSEClusterChargingTargetScheduleStructClass() _MTREnergyEVSEClusterChargingTargetScheduleStructClass {
	MTREnergyEVSEClusterChargingTargetScheduleStructClassOnce.Do(func() {
		MTREnergyEVSEClusterChargingTargetScheduleStructClass = _MTREnergyEVSEClusterChargingTargetScheduleStructClass{objc.GetClass("MTREnergyEVSEClusterChargingTargetScheduleStruct")}
	})
	return MTREnergyEVSEClusterChargingTargetScheduleStructClass
}

type _MTREnergyEVSEClusterChargingTargetScheduleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterChargingTargetScheduleStruct] class.
type IMTREnergyEVSEClusterChargingTargetScheduleStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetScheduleStruct
type MTREnergyEVSEClusterChargingTargetScheduleStruct struct {
	objectivec.Object
}

// MTREnergyEVSEClusterChargingTargetScheduleStructFrom constructs a [MTREnergyEVSEClusterChargingTargetScheduleStruct] from an unsafe.Pointer.
func MTREnergyEVSEClusterChargingTargetScheduleStructFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterChargingTargetScheduleStruct {
	return MTREnergyEVSEClusterChargingTargetScheduleStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterChargingTargetScheduleStructClass) Alloc() MTREnergyEVSEClusterChargingTargetScheduleStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetScheduleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterChargingTargetScheduleStructClass) New() MTREnergyEVSEClusterChargingTargetScheduleStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetScheduleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) Init() MTREnergyEVSEClusterChargingTargetScheduleStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetScheduleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) Autorelease() MTREnergyEVSEClusterChargingTargetScheduleStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetScheduleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterChargingTargetScheduleStruct creates a new MTREnergyEVSEClusterChargingTargetScheduleStruct instance.
func NewMTREnergyEVSEClusterChargingTargetScheduleStruct() MTREnergyEVSEClusterChargingTargetScheduleStruct {
	return getMTREnergyEVSEClusterChargingTargetScheduleStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetScheduleStruct/chargingTargets
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) ChargingTargets() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("chargingTargets"))
	return rv
}


// SetChargingTargets sets the value of the chargingTargets property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetScheduleStruct/chargingTargets
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) SetChargingTargets(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChargingTargets:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetScheduleStruct/dayOfWeekForSequence
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) DayOfWeekForSequence() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("dayOfWeekForSequence"))
	return rv
}


// SetDayOfWeekForSequence sets the value of the dayOfWeekForSequence property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetScheduleStruct/dayOfWeekForSequence
func (m_ MTREnergyEVSEClusterChargingTargetScheduleStruct) SetDayOfWeekForSequence(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDayOfWeekForSequence:"), value)
}



