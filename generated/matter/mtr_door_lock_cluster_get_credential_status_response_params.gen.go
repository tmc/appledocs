// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
var (
	MTRDoorLockClusterGetCredentialStatusResponseParamsClass     _MTRDoorLockClusterGetCredentialStatusResponseParamsClass
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce sync.Once
)

func getMTRDoorLockClusterGetCredentialStatusResponseParamsClass() _MTRDoorLockClusterGetCredentialStatusResponseParamsClass {
	MTRDoorLockClusterGetCredentialStatusResponseParamsClassOnce.Do(func() {
		MTRDoorLockClusterGetCredentialStatusResponseParamsClass = _MTRDoorLockClusterGetCredentialStatusResponseParamsClass{objc.GetClass("MTRDoorLockClusterGetCredentialStatusResponseParams")}
	})
	return MTRDoorLockClusterGetCredentialStatusResponseParamsClass
}

type _MTRDoorLockClusterGetCredentialStatusResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDoorLockClusterGetCredentialStatusResponseParams] class.
type IMTRDoorLockClusterGetCredentialStatusResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterGetCredentialStatusResponseParams
type MTRDoorLockClusterGetCredentialStatusResponseParams struct {
	objectivec.Object
}

// MTRDoorLockClusterGetCredentialStatusResponseParamsFrom constructs a [MTRDoorLockClusterGetCredentialStatusResponseParams] from an unsafe.Pointer.
func MTRDoorLockClusterGetCredentialStatusResponseParamsFrom(ptr unsafe.Pointer) MTRDoorLockClusterGetCredentialStatusResponseParams {
	return MTRDoorLockClusterGetCredentialStatusResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) Alloc() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDoorLockClusterGetCredentialStatusResponseParamsClass) New() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Init() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterGetCredentialStatusResponseParams) Autorelease() MTRDoorLockClusterGetCredentialStatusResponseParams {
	rv := objc.Send[MTRDoorLockClusterGetCredentialStatusResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterGetCredentialStatusResponseParams creates a new MTRDoorLockClusterGetCredentialStatusResponseParams instance.
func NewMTRDoorLockClusterGetCredentialStatusResponseParams() MTRDoorLockClusterGetCredentialStatusResponseParams {
	return getMTRDoorLockClusterGetCredentialStatusResponseParamsClass().New()
}




