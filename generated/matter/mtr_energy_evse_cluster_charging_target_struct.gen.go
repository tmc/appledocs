// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEClusterChargingTargetStruct] class.
var (
	MTREnergyEVSEClusterChargingTargetStructClass     _MTREnergyEVSEClusterChargingTargetStructClass
	MTREnergyEVSEClusterChargingTargetStructClassOnce sync.Once
)

func getMTREnergyEVSEClusterChargingTargetStructClass() _MTREnergyEVSEClusterChargingTargetStructClass {
	MTREnergyEVSEClusterChargingTargetStructClassOnce.Do(func() {
		MTREnergyEVSEClusterChargingTargetStructClass = _MTREnergyEVSEClusterChargingTargetStructClass{objc.GetClass("MTREnergyEVSEClusterChargingTargetStruct")}
	})
	return MTREnergyEVSEClusterChargingTargetStructClass
}

type _MTREnergyEVSEClusterChargingTargetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEClusterChargingTargetStruct] class.
type IMTREnergyEVSEClusterChargingTargetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct
type MTREnergyEVSEClusterChargingTargetStruct struct {
	objectivec.Object
}

// MTREnergyEVSEClusterChargingTargetStructFrom constructs a [MTREnergyEVSEClusterChargingTargetStruct] from an unsafe.Pointer.
func MTREnergyEVSEClusterChargingTargetStructFrom(ptr unsafe.Pointer) MTREnergyEVSEClusterChargingTargetStruct {
	return MTREnergyEVSEClusterChargingTargetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEClusterChargingTargetStructClass) Alloc() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEClusterChargingTargetStructClass) New() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEClusterChargingTargetStruct) Init() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEClusterChargingTargetStruct) Autorelease() MTREnergyEVSEClusterChargingTargetStruct {
	rv := objc.Send[MTREnergyEVSEClusterChargingTargetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEClusterChargingTargetStruct creates a new MTREnergyEVSEClusterChargingTargetStruct instance.
func NewMTREnergyEVSEClusterChargingTargetStruct() MTREnergyEVSEClusterChargingTargetStruct {
	return getMTREnergyEVSEClusterChargingTargetStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/addedEnergy
func (m_ MTREnergyEVSEClusterChargingTargetStruct) AddedEnergy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("addedEnergy"))
	return rv
}


// SetAddedEnergy sets the value of the addedEnergy property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/addedEnergy
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetAddedEnergy(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAddedEnergy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/targetSoC
func (m_ MTREnergyEVSEClusterChargingTargetStruct) TargetSoC() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetSoC"))
	return rv
}


// SetTargetSoC sets the value of the targetSoC property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/targetSoC
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetTargetSoC(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetSoC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/targetTimeMinutesPastMidnight
func (m_ MTREnergyEVSEClusterChargingTargetStruct) TargetTimeMinutesPastMidnight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetTimeMinutesPastMidnight"))
	return rv
}


// SetTargetTimeMinutesPastMidnight sets the value of the targetTimeMinutesPastMidnight property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEClusterChargingTargetStruct/targetTimeMinutesPastMidnight
func (m_ MTREnergyEVSEClusterChargingTargetStruct) SetTargetTimeMinutesPastMidnight(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetTimeMinutesPastMidnight:"), value)
}



