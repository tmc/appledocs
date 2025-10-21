// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] class.
var (
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass     _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass() _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass {
	MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass = _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass{objc.GetClass("MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct")}
	})
	return MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass
}

type _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] class.
type IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
type MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructFrom constructs a [MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	return MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass) Alloc() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass) New() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Init() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Autorelease() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct creates a new MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct instance.
func NewMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	return getMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/emergencyBoost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) EmergencyBoost() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("emergencyBoost"))
	return rv
}


// SetEmergencyBoost sets the value of the emergencyBoost property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/emergencyBoost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetEmergencyBoost(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEmergencyBoost:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/oneShot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) OneShot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("oneShot"))
	return rv
}


// SetOneShot sets the value of the oneShot property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/oneShot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetOneShot(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOneShot:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetPercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetPercentage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetPercentage"))
	return rv
}


// SetTargetPercentage sets the value of the targetPercentage property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetPercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetPercentage(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetPercentage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetReheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetReheat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targetReheat"))
	return rv
}


// SetTargetReheat sets the value of the targetReheat property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetReheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetReheat(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetReheat:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/temporarySetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TemporarySetpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("temporarySetpoint"))
	return rv
}


// SetTemporarySetpoint sets the value of the temporarySetpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/temporarySetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTemporarySetpoint(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporarySetpoint:"), value)
}



