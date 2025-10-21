// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterUnlockDoorParams] class.
var (
	MTRDoorLockClusterUnlockDoorParamsClass     _MTRDoorLockClusterUnlockDoorParamsClass
	MTRDoorLockClusterUnlockDoorParamsClassOnce sync.Once
)

func getMTRDoorLockClusterUnlockDoorParamsClass() _MTRDoorLockClusterUnlockDoorParamsClass {
	MTRDoorLockClusterUnlockDoorParamsClassOnce.Do(func() {
		MTRDoorLockClusterUnlockDoorParamsClass = _MTRDoorLockClusterUnlockDoorParamsClass{objc.GetClass("MTRDoorLockClusterUnlockDoorParams")}
	})
	return MTRDoorLockClusterUnlockDoorParamsClass
}

type _MTRDoorLockClusterUnlockDoorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterUnlockDoorParams] class.
type IMTRDoorLockClusterUnlockDoorParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterUnlockDoorParams
type MTRDoorLockClusterUnlockDoorParams struct {
	objectivec.Object
}

// MTRDoorLockClusterUnlockDoorParamsFrom constructs a [MTRDoorLockClusterUnlockDoorParams] from an unsafe.Pointer.
func MTRDoorLockClusterUnlockDoorParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterUnlockDoorParams {
	return MTRDoorLockClusterUnlockDoorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) Alloc() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterUnlockDoorParamsClass) New() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterUnlockDoorParams) Init() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterUnlockDoorParams) Autorelease() MTRDoorLockClusterUnlockDoorParams {
	rv := objc.Send[MTRDoorLockClusterUnlockDoorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterUnlockDoorParams creates a new MTRDoorLockClusterUnlockDoorParams instance.
func NewMTRDoorLockClusterUnlockDoorParams() MTRDoorLockClusterUnlockDoorParams {
	return getMTRDoorLockClusterUnlockDoorParamsClass().New()
}




