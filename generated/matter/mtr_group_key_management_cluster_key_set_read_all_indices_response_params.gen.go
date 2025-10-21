// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams] class.
var (
	MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass     _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass
	MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass() _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass {
	MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass = _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams")}
	})
	return MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass
}

type _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams] class.
type IMTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams
type MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	return MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass) New() MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams) Init() MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams) Autorelease() MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams creates a new MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams instance.
func NewMTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams() MTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParams {
	return getMTRGroupKeyManagementClusterKeySetReadAllIndicesResponseParamsClass().New()
}




