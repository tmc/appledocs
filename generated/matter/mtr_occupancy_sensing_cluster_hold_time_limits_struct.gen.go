// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROccupancySensingClusterHoldTimeLimitsStruct] class.
var (
	MTROccupancySensingClusterHoldTimeLimitsStructClass     _MTROccupancySensingClusterHoldTimeLimitsStructClass
	MTROccupancySensingClusterHoldTimeLimitsStructClassOnce sync.Once
)

func getMTROccupancySensingClusterHoldTimeLimitsStructClass() _MTROccupancySensingClusterHoldTimeLimitsStructClass {
	MTROccupancySensingClusterHoldTimeLimitsStructClassOnce.Do(func() {
		MTROccupancySensingClusterHoldTimeLimitsStructClass = _MTROccupancySensingClusterHoldTimeLimitsStructClass{objc.GetClass("MTROccupancySensingClusterHoldTimeLimitsStruct")}
	})
	return MTROccupancySensingClusterHoldTimeLimitsStructClass
}

type _MTROccupancySensingClusterHoldTimeLimitsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROccupancySensingClusterHoldTimeLimitsStruct] class.
type IMTROccupancySensingClusterHoldTimeLimitsStruct interface {
	objectivec.IObject
	HoldTimeDefault() foundation.Number
	SetHoldTimeDefault(value foundation.INumber)
	HoldTimeMax() foundation.Number
	SetHoldTimeMax(value foundation.INumber)
	HoldTimeMin() foundation.Number
	SetHoldTimeMin(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct
type MTROccupancySensingClusterHoldTimeLimitsStruct struct {
	objectivec.Object
}

// MTROccupancySensingClusterHoldTimeLimitsStructFrom constructs a [MTROccupancySensingClusterHoldTimeLimitsStruct] from an unsafe.Pointer.
func MTROccupancySensingClusterHoldTimeLimitsStructFrom(ptr unsafe.Pointer) MTROccupancySensingClusterHoldTimeLimitsStruct {
	return MTROccupancySensingClusterHoldTimeLimitsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROccupancySensingClusterHoldTimeLimitsStructClass) Alloc() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROccupancySensingClusterHoldTimeLimitsStructClass) New() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) Init() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) Autorelease() MTROccupancySensingClusterHoldTimeLimitsStruct {
	rv := objc.Send[MTROccupancySensingClusterHoldTimeLimitsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROccupancySensingClusterHoldTimeLimitsStruct creates a new MTROccupancySensingClusterHoldTimeLimitsStruct instance.
func NewMTROccupancySensingClusterHoldTimeLimitsStruct() MTROccupancySensingClusterHoldTimeLimitsStruct {
	return getMTROccupancySensingClusterHoldTimeLimitsStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeDefault
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeDefault() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holdTimeDefault"))
	return rv
}


// SetHoldTimeDefault sets the value of the holdTimeDefault property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeDefault
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeDefault(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeDefault:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMax
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeMax() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holdTimeMax"))
	return rv
}


// SetHoldTimeMax sets the value of the holdTimeMax property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMax
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeMax(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeMax:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMin
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) HoldTimeMin() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("holdTimeMin"))
	return rv
}


// SetHoldTimeMin sets the value of the holdTimeMin property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterHoldTimeLimitsStruct/holdTimeMin
func (m_ MTROccupancySensingClusterHoldTimeLimitsStruct) SetHoldTimeMin(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHoldTimeMin:"), value)
}



