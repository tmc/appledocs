// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLevelControlClusterStopParams] class.
var (
	MTRLevelControlClusterStopParamsClass     _MTRLevelControlClusterStopParamsClass
	MTRLevelControlClusterStopParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStopParamsClass() _MTRLevelControlClusterStopParamsClass {
	MTRLevelControlClusterStopParamsClassOnce.Do(func() {
		MTRLevelControlClusterStopParamsClass = _MTRLevelControlClusterStopParamsClass{objc.GetClass("MTRLevelControlClusterStopParams")}
	})
	return MTRLevelControlClusterStopParamsClass
}

type _MTRLevelControlClusterStopParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStopParams] class.
type IMTRLevelControlClusterStopParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopParams
type MTRLevelControlClusterStopParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStopParamsFrom constructs a [MTRLevelControlClusterStopParams] from an unsafe.Pointer.
func MTRLevelControlClusterStopParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStopParams {
	return MTRLevelControlClusterStopParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStopParamsClass) Alloc() MTRLevelControlClusterStopParams {
	rv := objc.Send[MTRLevelControlClusterStopParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStopParamsClass) New() MTRLevelControlClusterStopParams {
	rv := objc.Send[MTRLevelControlClusterStopParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStopParams) Init() MTRLevelControlClusterStopParams {
	rv := objc.Send[MTRLevelControlClusterStopParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStopParams) Autorelease() MTRLevelControlClusterStopParams {
	rv := objc.Send[MTRLevelControlClusterStopParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStopParams creates a new MTRLevelControlClusterStopParams instance.
func NewMTRLevelControlClusterStopParams() MTRLevelControlClusterStopParams {
	return getMTRLevelControlClusterStopParamsClass().New()
}




