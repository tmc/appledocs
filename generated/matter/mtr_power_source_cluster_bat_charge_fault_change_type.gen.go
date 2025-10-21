// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPowerSourceClusterBatChargeFaultChangeType] class.
var (
	MTRPowerSourceClusterBatChargeFaultChangeTypeClass     _MTRPowerSourceClusterBatChargeFaultChangeTypeClass
	MTRPowerSourceClusterBatChargeFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterBatChargeFaultChangeTypeClass() _MTRPowerSourceClusterBatChargeFaultChangeTypeClass {
	MTRPowerSourceClusterBatChargeFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterBatChargeFaultChangeTypeClass = _MTRPowerSourceClusterBatChargeFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterBatChargeFaultChangeType")}
	})
	return MTRPowerSourceClusterBatChargeFaultChangeTypeClass
}

type _MTRPowerSourceClusterBatChargeFaultChangeTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterBatChargeFaultChangeType] class.
type IMTRPowerSourceClusterBatChargeFaultChangeType interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeType
type MTRPowerSourceClusterBatChargeFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatChargeFaultChangeTypeFrom constructs a [MTRPowerSourceClusterBatChargeFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterBatChargeFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatChargeFaultChangeType {
	return MTRPowerSourceClusterBatChargeFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeTypeClass) Alloc() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeTypeClass) New() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Init() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Autorelease() MTRPowerSourceClusterBatChargeFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatChargeFaultChangeType creates a new MTRPowerSourceClusterBatChargeFaultChangeType instance.
func NewMTRPowerSourceClusterBatChargeFaultChangeType() MTRPowerSourceClusterBatChargeFaultChangeType {
	return getMTRPowerSourceClusterBatChargeFaultChangeTypeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatchargefaultchangetype/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatchargefaultchangetype/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatchargefaultchangetype/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterbatchargefaultchangetype/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeType) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}



