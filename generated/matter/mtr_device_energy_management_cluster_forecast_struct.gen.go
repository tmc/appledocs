// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ActiveSlotNumber() objc.IObject /* cross-framework: NSNumber */
	SetActiveSlotNumber(value objc.IObject /* cross-framework: NSNumber */)
	EarliestStartTime() objc.IObject /* cross-framework: NSNumber */
	SetEarliestStartTime(value objc.IObject /* cross-framework: NSNumber */)
	EndTime() objc.IObject /* cross-framework: NSNumber */
	SetEndTime(value objc.IObject /* cross-framework: NSNumber */)
	ForecastID() objc.IObject /* cross-framework: NSNumber */
	SetForecastID(value objc.IObject /* cross-framework: NSNumber */)
	ForecastUpdateReason() objc.IObject /* cross-framework: NSNumber */
	SetForecastUpdateReason(value objc.IObject /* cross-framework: NSNumber */)
	IsPausable() objc.IObject /* cross-framework: NSNumber */
	SetIsPausable(value objc.IObject /* cross-framework: NSNumber */)
	LatestEndTime() objc.IObject /* cross-framework: NSNumber */
	SetLatestEndTime(value objc.IObject /* cross-framework: NSNumber */)
	Slots() objc.IObject /* cross-framework: NSArray */
	SetSlots(value objc.IObject /* cross-framework: NSArray */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/activeSlotNumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ActiveSlotNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeSlotNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/activeSlotNumber
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetActiveSlotNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveSlotNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/earliestStartTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EarliestStartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("earliestStartTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/earliestStartTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEarliestStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEarliestStartTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/endTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) EndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/endTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastID
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastID
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastUpdateReason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) ForecastUpdateReason() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("forecastUpdateReason"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/forecastUpdateReason
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetForecastUpdateReason(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setForecastUpdateReason:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/isPausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) IsPausable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isPausable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/isPausable
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetIsPausable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsPausable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) LatestEndTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("latestEndTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/latestEndTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetLatestEndTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLatestEndTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/slots
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) Slots() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("slots"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/slots
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetSlots(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSlots:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterForecastStruct/startTime
func (m_ MTRDeviceEnergyManagementClusterForecastStruct) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}



