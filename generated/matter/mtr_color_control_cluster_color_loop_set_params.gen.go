// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterColorLoopSetParams] class.
var (
	MTRColorControlClusterColorLoopSetParamsClass     _MTRColorControlClusterColorLoopSetParamsClass
	MTRColorControlClusterColorLoopSetParamsClassOnce sync.Once
)

func getMTRColorControlClusterColorLoopSetParamsClass() _MTRColorControlClusterColorLoopSetParamsClass {
	MTRColorControlClusterColorLoopSetParamsClassOnce.Do(func() {
		MTRColorControlClusterColorLoopSetParamsClass = _MTRColorControlClusterColorLoopSetParamsClass{objc.GetClass("MTRColorControlClusterColorLoopSetParams")}
	})
	return MTRColorControlClusterColorLoopSetParamsClass
}

type _MTRColorControlClusterColorLoopSetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterColorLoopSetParams] class.
type IMTRColorControlClusterColorLoopSetParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterColorLoopSetParams
type MTRColorControlClusterColorLoopSetParams struct {
	objectivec.Object
}

// MTRColorControlClusterColorLoopSetParamsFrom constructs a [MTRColorControlClusterColorLoopSetParams] from an unsafe.Pointer.
func MTRColorControlClusterColorLoopSetParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterColorLoopSetParams {
	return MTRColorControlClusterColorLoopSetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) Alloc() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterColorLoopSetParamsClass) New() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterColorLoopSetParams) Init() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterColorLoopSetParams) Autorelease() MTRColorControlClusterColorLoopSetParams {
	rv := objc.Send[MTRColorControlClusterColorLoopSetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterColorLoopSetParams creates a new MTRColorControlClusterColorLoopSetParams instance.
func NewMTRColorControlClusterColorLoopSetParams() MTRColorControlClusterColorLoopSetParams {
	return getMTRColorControlClusterColorLoopSetParamsClass().New()
}




