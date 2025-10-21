// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRPowerSourceClusterWiredFaultChangeType] class.
var (
	MTRPowerSourceClusterWiredFaultChangeTypeClass     _MTRPowerSourceClusterWiredFaultChangeTypeClass
	MTRPowerSourceClusterWiredFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterWiredFaultChangeTypeClass() _MTRPowerSourceClusterWiredFaultChangeTypeClass {
	MTRPowerSourceClusterWiredFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterWiredFaultChangeTypeClass = _MTRPowerSourceClusterWiredFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterWiredFaultChangeType")}
	})
	return MTRPowerSourceClusterWiredFaultChangeTypeClass
}

type _MTRPowerSourceClusterWiredFaultChangeTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterWiredFaultChangeType] class.
type IMTRPowerSourceClusterWiredFaultChangeType interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeType
type MTRPowerSourceClusterWiredFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterWiredFaultChangeTypeFrom constructs a [MTRPowerSourceClusterWiredFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterWiredFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterWiredFaultChangeType {
	return MTRPowerSourceClusterWiredFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterWiredFaultChangeTypeClass) Alloc() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterWiredFaultChangeTypeClass) New() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Init() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Autorelease() MTRPowerSourceClusterWiredFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterWiredFaultChangeType creates a new MTRPowerSourceClusterWiredFaultChangeType instance.
func NewMTRPowerSourceClusterWiredFaultChangeType() MTRPowerSourceClusterWiredFaultChangeType {
	return getMTRPowerSourceClusterWiredFaultChangeTypeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangetype/current
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Current() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("current"))
	return rv
}


// SetCurrent sets the value of the current property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangetype/current
func (m_ MTRPowerSourceClusterWiredFaultChangeType) SetCurrent(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangetype/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeType) Previous() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("previous"))
	return rv
}


// SetPrevious sets the value of the previous property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrpowersourceclusterwiredfaultchangetype/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeType) SetPrevious(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}



