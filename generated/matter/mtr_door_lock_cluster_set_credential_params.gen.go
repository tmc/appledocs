// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterSetCredentialParams] class.
var (
	MTRDoorLockClusterSetCredentialParamsClass     _MTRDoorLockClusterSetCredentialParamsClass
	MTRDoorLockClusterSetCredentialParamsClassOnce sync.Once
)

func getMTRDoorLockClusterSetCredentialParamsClass() _MTRDoorLockClusterSetCredentialParamsClass {
	MTRDoorLockClusterSetCredentialParamsClassOnce.Do(func() {
		MTRDoorLockClusterSetCredentialParamsClass = _MTRDoorLockClusterSetCredentialParamsClass{objc.GetClass("MTRDoorLockClusterSetCredentialParams")}
	})
	return MTRDoorLockClusterSetCredentialParamsClass
}

type _MTRDoorLockClusterSetCredentialParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterSetCredentialParams] class.
type IMTRDoorLockClusterSetCredentialParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterSetCredentialParams
type MTRDoorLockClusterSetCredentialParams struct {
	objectivec.Object
}

// MTRDoorLockClusterSetCredentialParamsFrom constructs a [MTRDoorLockClusterSetCredentialParams] from an unsafe.Pointer.
func MTRDoorLockClusterSetCredentialParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterSetCredentialParams {
	return MTRDoorLockClusterSetCredentialParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterSetCredentialParamsClass) Alloc() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterSetCredentialParamsClass) New() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterSetCredentialParams) Init() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterSetCredentialParams) Autorelease() MTRDoorLockClusterSetCredentialParams {
	rv := objc.Send[MTRDoorLockClusterSetCredentialParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterSetCredentialParams creates a new MTRDoorLockClusterSetCredentialParams instance.
func NewMTRDoorLockClusterSetCredentialParams() MTRDoorLockClusterSetCredentialParams {
	return getMTRDoorLockClusterSetCredentialParamsClass().New()
}




