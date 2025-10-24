// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	EmergencyBoost() objc.IObject /* cross-framework: NSNumber */
	SetEmergencyBoost(value objc.IObject /* cross-framework: NSNumber */)
	OneShot() objc.IObject /* cross-framework: NSNumber */
	SetOneShot(value objc.IObject /* cross-framework: NSNumber */)
	TargetPercentage() objc.IObject /* cross-framework: NSNumber */
	SetTargetPercentage(value objc.IObject /* cross-framework: NSNumber */)
	TargetReheat() objc.IObject /* cross-framework: NSNumber */
	SetTargetReheat(value objc.IObject /* cross-framework: NSNumber */)
	TemporarySetpoint() objc.IObject /* cross-framework: NSNumber */
	SetTemporarySetpoint(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("duration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/duration
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDuration:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/emergencyBoost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) EmergencyBoost() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("emergencyBoost"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/emergencyBoost
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetEmergencyBoost(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEmergencyBoost:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/oneShot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) OneShot() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("oneShot"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/oneShot
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetOneShot(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOneShot:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetPercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetPercentage() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetPercentage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetPercentage
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetPercentage(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetPercentage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetReheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TargetReheat() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetReheat"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/targetReheat
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTargetReheat(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetReheat:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/temporarySetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) TemporarySetpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("temporarySetpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct/temporarySetpoint
func (m_ MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) SetTemporarySetpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporarySetpoint:"), value)
}



