// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveColorParams] class.
var (
	MTRColorControlClusterMoveColorParamsClass     _MTRColorControlClusterMoveColorParamsClass
	MTRColorControlClusterMoveColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveColorParamsClass() _MTRColorControlClusterMoveColorParamsClass {
	MTRColorControlClusterMoveColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveColorParamsClass = _MTRColorControlClusterMoveColorParamsClass{objc.GetClass("MTRColorControlClusterMoveColorParams")}
	})
	return MTRColorControlClusterMoveColorParamsClass
}

type _MTRColorControlClusterMoveColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveColorParams] class.
type IMTRColorControlClusterMoveColorParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveColorParams
type MTRColorControlClusterMoveColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveColorParamsFrom constructs a [MTRColorControlClusterMoveColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveColorParams {
	return MTRColorControlClusterMoveColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveColorParamsClass) Alloc() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveColorParamsClass) New() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveColorParams) Init() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveColorParams) Autorelease() MTRColorControlClusterMoveColorParams {
	rv := objc.Send[MTRColorControlClusterMoveColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveColorParams creates a new MTRColorControlClusterMoveColorParams instance.
func NewMTRColorControlClusterMoveColorParams() MTRColorControlClusterMoveColorParams {
	return getMTRColorControlClusterMoveColorParamsClass().New()
}




