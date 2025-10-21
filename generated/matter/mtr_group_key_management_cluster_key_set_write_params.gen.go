// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterKeySetWriteParams] class.
var (
	MTRGroupKeyManagementClusterKeySetWriteParamsClass     _MTRGroupKeyManagementClusterKeySetWriteParamsClass
	MTRGroupKeyManagementClusterKeySetWriteParamsClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterKeySetWriteParamsClass() _MTRGroupKeyManagementClusterKeySetWriteParamsClass {
	MTRGroupKeyManagementClusterKeySetWriteParamsClassOnce.Do(func() {
		MTRGroupKeyManagementClusterKeySetWriteParamsClass = _MTRGroupKeyManagementClusterKeySetWriteParamsClass{objc.GetClass("MTRGroupKeyManagementClusterKeySetWriteParams")}
	})
	return MTRGroupKeyManagementClusterKeySetWriteParamsClass
}

type _MTRGroupKeyManagementClusterKeySetWriteParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterKeySetWriteParams] class.
type IMTRGroupKeyManagementClusterKeySetWriteParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterKeySetWriteParams
type MTRGroupKeyManagementClusterKeySetWriteParams struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterKeySetWriteParamsFrom constructs a [MTRGroupKeyManagementClusterKeySetWriteParams] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterKeySetWriteParamsFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterKeySetWriteParams {
	return MTRGroupKeyManagementClusterKeySetWriteParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterKeySetWriteParamsClass) Alloc() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterKeySetWriteParamsClass) New() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) Init() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterKeySetWriteParams) Autorelease() MTRGroupKeyManagementClusterKeySetWriteParams {
	rv := objc.Send[MTRGroupKeyManagementClusterKeySetWriteParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterKeySetWriteParams creates a new MTRGroupKeyManagementClusterKeySetWriteParams instance.
func NewMTRGroupKeyManagementClusterKeySetWriteParams() MTRGroupKeyManagementClusterKeySetWriteParams {
	return getMTRGroupKeyManagementClusterKeySetWriteParamsClass().New()
}




