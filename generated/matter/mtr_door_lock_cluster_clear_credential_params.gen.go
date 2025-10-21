// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDoorLockClusterClearCredentialParams] class.
var (
	MTRDoorLockClusterClearCredentialParamsClass     _MTRDoorLockClusterClearCredentialParamsClass
	MTRDoorLockClusterClearCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterClearCredentialParamsClass() _MTRDoorLockClusterClearCredentialParamsClass {
	MTRDoorLockClusterClearCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterClearCredentialParamsClass = _MTRDoorLockClusterClearCredentialParamsClass{objc.GetClass("MTRDoorLockClusterClearCredentialParams")}
	})
	return MTRDoorLockClusterClearCredentialParamsClass
}

type _MTRDoorLockClusterClearCredentialParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterClearCredentialParams] class.
type IMTRDoorLockClusterClearCredentialParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterClearCredentialParams
type MTRDoorLockClusterClearCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterClearCredentialParamsFrom constructs a [MTRDoorLockClusterClearCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterClearCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterClearCredentialParams {
	return MTRDoorLockClusterClearCredentialParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterClearCredentialParamsClass) Alloc() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterClearCredentialParamsClass) New() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterClearCredentialParams) Init() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterClearCredentialParams) Autorelease() MTRDoorLockClusterClearCredentialParams {
	rv := objc.Send[MTRDoorLockClusterClearCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterClearCredentialParams creates a new MTRDoorLockClusterClearCredentialParams instance.
func NewMTRDoorLockClusterClearCredentialParams() MTRDoorLockClusterClearCredentialParams {
	return getMTRDoorLockClusterClearCredentialParamsClass().New()
}




