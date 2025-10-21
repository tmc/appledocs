// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRColorControlClusterMoveHueParams] class.
var (
	MTRColorControlClusterMoveHueParamsClass     _MTRColorControlClusterMoveHueParamsClass
	MTRColorControlClusterMoveHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveHueParamsClass() _MTRColorControlClusterMoveHueParamsClass {
	MTRColorControlClusterMoveHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveHueParamsClass = _MTRColorControlClusterMoveHueParamsClass{objc.GetClass("MTRColorControlClusterMoveHueParams")}
	})
	return MTRColorControlClusterMoveHueParamsClass
}

type _MTRColorControlClusterMoveHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveHueParams] class.
type IMTRColorControlClusterMoveHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveHueParams
type MTRColorControlClusterMoveHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveHueParamsFrom constructs a [MTRColorControlClusterMoveHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveHueParams {
	return MTRColorControlClusterMoveHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveHueParamsClass) Alloc() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveHueParamsClass) New() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveHueParams) Init() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveHueParams) Autorelease() MTRColorControlClusterMoveHueParams {
	rv := objc.Send[MTRColorControlClusterMoveHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveHueParams creates a new MTRColorControlClusterMoveHueParams instance.
func NewMTRColorControlClusterMoveHueParams() MTRColorControlClusterMoveHueParams {
	return getMTRColorControlClusterMoveHueParamsClass().New()
}




