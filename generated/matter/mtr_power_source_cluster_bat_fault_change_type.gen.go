// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRPowerSourceClusterBatFaultChangeType] class.
var (
	MTRPowerSourceClusterBatFaultChangeTypeClass     _MTRPowerSourceClusterBatFaultChangeTypeClass
	MTRPowerSourceClusterBatFaultChangeTypeClassOnce sync.Once
)

func getMTRPowerSourceClusterBatFaultChangeTypeClass() _MTRPowerSourceClusterBatFaultChangeTypeClass {
	MTRPowerSourceClusterBatFaultChangeTypeClassOnce.Do(func() {
		MTRPowerSourceClusterBatFaultChangeTypeClass = _MTRPowerSourceClusterBatFaultChangeTypeClass{objc.GetClass("MTRPowerSourceClusterBatFaultChangeType")}
	})
	return MTRPowerSourceClusterBatFaultChangeTypeClass
}

type _MTRPowerSourceClusterBatFaultChangeTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRPowerSourceClusterBatFaultChangeType] class.
type IMTRPowerSourceClusterBatFaultChangeType interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeType
type MTRPowerSourceClusterBatFaultChangeType struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatFaultChangeTypeFrom constructs a [MTRPowerSourceClusterBatFaultChangeType] from an unsafe.Pointer.
func MTRPowerSourceClusterBatFaultChangeTypeFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatFaultChangeType {
	return MTRPowerSourceClusterBatFaultChangeType{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatFaultChangeTypeClass) Alloc() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRPowerSourceClusterBatFaultChangeTypeClass) New() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatFaultChangeType) Init() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatFaultChangeType) Autorelease() MTRPowerSourceClusterBatFaultChangeType {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatFaultChangeType creates a new MTRPowerSourceClusterBatFaultChangeType instance.
func NewMTRPowerSourceClusterBatFaultChangeType() MTRPowerSourceClusterBatFaultChangeType {
	return getMTRPowerSourceClusterBatFaultChangeTypeClass().New()
}




