// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceEnergyManagementClusterForecastStruct] class.
var (
	MTRDeviceEnergyManagementClusterForecastStructClass     _MTRDeviceEnergyManagementClusterForecastStructClass
	MTRDeviceEnergyManagementClusterForecastStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterForecastStructClass() _MTRDeviceEnergyManagementClusterForecastStructClass {
	MTRDeviceEnergyManagementClusterForecastStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterForecastStructClass = _MTRDeviceEnergyManagementClusterForecastStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterForecastStruct")}
	})
	return MTRDeviceEnergyManagementClusterForecastStructClass
}

type _MTRDeviceEnergyManagementClusterForecastStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterForecastStruct] class.
type IMTRDeviceEnergyManagementClusterForecastStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct
type MTRDeviceEnergyManagementClusterForecastStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterForecastStructFrom constructs a [MTRDeviceEnergyManagementClusterForecastStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterForecastStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterForecastStruct {
	return MTRDeviceEnergyManagementClusterForecastStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterForecastStructClass) Alloc() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterForecastStructClass) New() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Init() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Autorelease() MTRDeviceEnergyManagementClusterForecastStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterForecastStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterForecastStruct creates a new MTRDeviceEnergyManagementClusterForecastStruct instance.
func NewMTRDeviceEnergyManagementClusterForecastStruct() MTRDeviceEnergyManagementClusterForecastStruct {
	return getMTRDeviceEnergyManagementClusterForecastStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/activeSlotNumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ActiveSlotNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("activeSlotNumber"))
	return rv
}


// SetActiveSlotNumber sets the value of the activeSlotNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/activeSlotNumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetActiveSlotNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveSlotNumber:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/earliestStartTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EarliestStartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("earliestStartTime"))
	return rv
}


// SetEarliestStartTime sets the value of the earliestStartTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/earliestStartTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEarliestStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEarliestStartTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/endTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endTime"))
	return rv
}


// SetEndTime sets the value of the endTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/endTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastID
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("forecastID"))
	return rv
}


// SetForecastID sets the value of the forecastID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastID
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastID:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastUpdateReason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastUpdateReason() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("forecastUpdateReason"))
	return rv
}


// SetForecastUpdateReason sets the value of the forecastUpdateReason property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastUpdateReason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastUpdateReason(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastUpdateReason:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/isPausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) IsPausable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("isPausable"))
	return rv
}


// SetIsPausable sets the value of the isPausable property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/isPausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetIsPausable(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPausable:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) LatestEndTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("latestEndTime"))
	return rv
}


// SetLatestEndTime sets the value of the latestEndTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetLatestEndTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestEndTime:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/slots
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Slots() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("slots"))
	return rv
}


// SetSlots sets the value of the slots property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/slots
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetSlots(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlots:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) StartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startTime"))
	return rv
}


// SetStartTime sets the value of the startTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}


