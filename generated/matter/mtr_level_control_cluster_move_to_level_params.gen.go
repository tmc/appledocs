// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelParams] class.
var (
	MTRLevelControlClusterMoveToLevelParamsClass     _MTRLevelControlClusterMoveToLevelParamsClass
	MTRLevelControlClusterMoveToLevelParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelParamsClass() _MTRLevelControlClusterMoveToLevelParamsClass {
	MTRLevelControlClusterMoveToLevelParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelParamsClass = _MTRLevelControlClusterMoveToLevelParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelParams")}
	})
	return MTRLevelControlClusterMoveToLevelParamsClass
}

type _MTRLevelControlClusterMoveToLevelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelParams] class.
type IMTRLevelControlClusterMoveToLevelParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelParams
type MTRLevelControlClusterMoveToLevelParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelParamsFrom constructs a [MTRLevelControlClusterMoveToLevelParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelParams {
	return MTRLevelControlClusterMoveToLevelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) Alloc() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelParamsClass) New() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelParams) Init() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelParams) Autorelease() MTRLevelControlClusterMoveToLevelParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelParams creates a new MTRLevelControlClusterMoveToLevelParams instance.
func NewMTRLevelControlClusterMoveToLevelParams() MTRLevelControlClusterMoveToLevelParams {
	return getMTRLevelControlClusterMoveToLevelParamsClass().New()
}




