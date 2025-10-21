// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
var (
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClass     _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass() _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass {
	MTRLevelControlClusterMoveToLevelWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterMoveToLevelWithOnOffParamsClass = _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterMoveToLevelWithOnOffParams")}
	})
	return MTRLevelControlClusterMoveToLevelWithOnOffParamsClass
}

type _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterMoveToLevelWithOnOffParams] class.
type IMTRLevelControlClusterMoveToLevelWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterMoveToLevelWithOnOffParams
type MTRLevelControlClusterMoveToLevelWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom constructs a [MTRLevelControlClusterMoveToLevelWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterMoveToLevelWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return MTRLevelControlClusterMoveToLevelWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) Alloc() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterMoveToLevelWithOnOffParamsClass) New() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Init() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterMoveToLevelWithOnOffParams) Autorelease() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterMoveToLevelWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterMoveToLevelWithOnOffParams creates a new MTRLevelControlClusterMoveToLevelWithOnOffParams instance.
func NewMTRLevelControlClusterMoveToLevelWithOnOffParams() MTRLevelControlClusterMoveToLevelWithOnOffParams {
	return getMTRLevelControlClusterMoveToLevelWithOnOffParamsClass().New()
}




