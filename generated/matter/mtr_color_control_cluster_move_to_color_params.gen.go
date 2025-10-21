// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveToColorParams] class.
var (
	MTRColorControlClusterMoveToColorParamsClass     _MTRColorControlClusterMoveToColorParamsClass
	MTRColorControlClusterMoveToColorParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToColorParamsClass() _MTRColorControlClusterMoveToColorParamsClass {
	MTRColorControlClusterMoveToColorParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToColorParamsClass = _MTRColorControlClusterMoveToColorParamsClass{objc.GetClass("MTRColorControlClusterMoveToColorParams")}
	})
	return MTRColorControlClusterMoveToColorParamsClass
}

type _MTRColorControlClusterMoveToColorParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToColorParams] class.
type IMTRColorControlClusterMoveToColorParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToColorParams
type MTRColorControlClusterMoveToColorParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToColorParamsFrom constructs a [MTRColorControlClusterMoveToColorParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToColorParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToColorParams {
	return MTRColorControlClusterMoveToColorParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToColorParamsClass) Alloc() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToColorParamsClass) New() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToColorParams) Init() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToColorParams) Autorelease() MTRColorControlClusterMoveToColorParams {
	rv := objc.Send[MTRColorControlClusterMoveToColorParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToColorParams creates a new MTRColorControlClusterMoveToColorParams instance.
func NewMTRColorControlClusterMoveToColorParams() MTRColorControlClusterMoveToColorParams {
	return getMTRColorControlClusterMoveToColorParamsClass().New()
}




