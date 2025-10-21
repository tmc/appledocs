// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetRemoveParams] class.
var (
	MTRGroupKeyManagementClusterKeySetRemoveParamsClass     _MTRGroupKeyManagementClusterKeySetRemoveParamsClass
	MTRGroupKeyManagementClusterKeySetRemoveParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetRemoveParamsClass() _MTRGroupKeyManagementClusterKeySetRemoveParamsClass {
	MTRGroupKeyManagementClusterKeySetRemoveParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetRemoveParamsClass = _MTRGroupKeyManagementClusterKeySetRemoveParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetRemoveParams")}
	})
	return MTRGroupKeyManagementClusterKeySetRemoveParamsClass
}

type _MTRGroupKeyManagementClusterKeySetRemoveParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetRemoveParams] class.
type IMTRGroupKeyManagementClusterKeySetRemoveParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetRemoveParams
type MTRGroupKeyManagementClusterKeySetRemoveParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetRemoveParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetRemoveParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetRemoveParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetRemoveParams {
	return MTRGroupKeyManagementClusterKeySetRemoveParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetRemoveParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetRemoveParamsClass) New() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) Init() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetRemoveParams) Autorelease() MTRGroupKeyManagementClusterKeySetRemoveParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetRemoveParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetRemoveParams creates a new MTRGroupKeyManagementClusterKeySetRemoveParams instance.
func NewMTRGroupKeyManagementClusterKeySetRemoveParams() MTRGroupKeyManagementClusterKeySetRemoveParams {
	return getMTRGroupKeyManagementClusterKeySetRemoveParamsClass().New()
}




