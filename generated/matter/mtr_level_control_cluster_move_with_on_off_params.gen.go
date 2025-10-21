// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLevelControlClusterMoveWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveWithOnOffParamsClass     _MTRLevelControlClusterMoveWithOnOffParamsClass
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveWithOnOffParamsClass() _MTRLevelControlClusterMoveWithOnOffParamsClass {
	MTRLevelControlClusterMoveWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveWithOnOffParamsClass = _MTRLevelControlClusterMoveWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveWithOnOffParams] class.
type IMTRLevelControlClusterMoveWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveWithOnOffParams
type MTRLevelControlClusterMoveWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveWithOnOffParams {
	return MTRLevelControlClusterMoveWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveWithOnOffParamsClass) New() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Init() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveWithOnOffParams) Autorelease() MTRLevelControlClusterMoveWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveWithOnOffParams creates a new MTRLevelControlClusterMoveWithOnOffParams instance.
func NewMTRLevelControlClusterMoveWithOnOffParams() MTRLevelControlClusterMoveWithOnOffParams {
	return getMTRLevelControlClusterMoveWithOnOffParamsClass().New()
}




