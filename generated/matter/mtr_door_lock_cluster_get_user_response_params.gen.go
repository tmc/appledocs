// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterGetUserResponseParams] class.
var (
	MTRDoorLockClusterGetUserResponseParamsClass     _MTRDoorLockClusterGetUserResponseParamsClass
	MTRDoorLockClusterGetUserResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetUserResponseParamsClass() _MTRDoorLockClusterGetUserResponseParamsClass {
	MTRDoorLockClusterGetUserResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetUserResponseParamsClass = _MTRDoorLockClusterGetUserResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetUserResponseParams")}
	})
	return MTRDoorLockClusterGetUserResponseParamsClass
}

type _MTRDoorLockClusterGetUserResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetUserResponseParams] class.
type IMTRDoorLockClusterGetUserResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetUserResponseParams
type MTRDoorLockClusterGetUserResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetUserResponseParamsFrom constructs a [MTRDoorLockClusterGetUserResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetUserResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetUserResponseParams {
	return MTRDoorLockClusterGetUserResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) Alloc() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetUserResponseParamsClass) New() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetUserResponseParams) Init() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetUserResponseParams) Autorelease() MTRDoorLockClusterGetUserResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetUserResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetUserResponseParams creates a new MTRDoorLockClusterGetUserResponseParams instance.
func NewMTRDoorLockClusterGetUserResponseParams() MTRDoorLockClusterGetUserResponseParams {
	return getMTRDoorLockClusterGetUserResponseParamsClass().New()
}




