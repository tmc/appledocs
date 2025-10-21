// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetUserParams] class.
var (
	MTRDoorLockClusterSetUserParamsClass     _MTRDoorLockClusterSetUserParamsClass
	MTRDoorLockClusterSetUserParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetUserParamsClass() _MTRDoorLockClusterSetUserParamsClass {
	MTRDoorLockClusterSetUserParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetUserParamsClass = _MTRDoorLockClusterSetUserParamsClass{objc.GetClass("MTRDoorLockClusterSetUserParams")}
	})
	return MTRDoorLockClusterSetUserParamsClass
}

type _MTRDoorLockClusterSetUserParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetUserParams] class.
type IMTRDoorLockClusterSetUserParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetUserParams
type MTRDoorLockClusterSetUserParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetUserParamsFrom constructs a [MTRDoorLockClusterSetUserParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetUserParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetUserParams {
	return MTRDoorLockClusterSetUserParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetUserParamsClass) Alloc() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetUserParamsClass) New() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetUserParams) Init() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetUserParams) Autorelease() MTRDoorLockClusterSetUserParams {
	rv := objc.Send[MTRDoorLockClusterSetUserParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetUserParams creates a new MTRDoorLockClusterSetUserParams instance.
func NewMTRDoorLockClusterSetUserParams() MTRDoorLockClusterSetUserParams {
	return getMTRDoorLockClusterSetUserParamsClass().New()
}




