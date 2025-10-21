// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLevelControlClusterStopWithOnOffParams] class.
var (
	MTRLevelControlClusterStopWithOnOffParamsClass     _MTRLevelControlClusterStopWithOnOffParamsClass
	MTRLevelControlClusterStopWithOnOffParamsClassOnce sync.Once
)

func getMTRLevelControlClusterStopWithOnOffParamsClass() _MTRLevelControlClusterStopWithOnOffParamsClass {
	MTRLevelControlClusterStopWithOnOffParamsClassOnce.Do(func() {
		MTRLevelControlClusterStopWithOnOffParamsClass = _MTRLevelControlClusterStopWithOnOffParamsClass{objc.GetClass("MTRLevelControlClusterStopWithOnOffParams")}
	})
	return MTRLevelControlClusterStopWithOnOffParamsClass
}

type _MTRLevelControlClusterStopWithOnOffParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRLevelControlClusterStopWithOnOffParams] class.
type IMTRLevelControlClusterStopWithOnOffParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLevelControlClusterStopWithOnOffParams
type MTRLevelControlClusterStopWithOnOffParams struct {
	objectivec.Object
}

// MTRLevelControlClusterStopWithOnOffParamsFrom constructs a [MTRLevelControlClusterStopWithOnOffParams] from an unsafe.Pointer.
func MTRLevelControlClusterStopWithOnOffParamsFrom(ptr unsafe.Pointer) MTRLevelControlClusterStopWithOnOffParams {
	return MTRLevelControlClusterStopWithOnOffParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) Alloc() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLevelControlClusterStopWithOnOffParamsClass) New() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Init() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLevelControlClusterStopWithOnOffParams) Autorelease() MTRLevelControlClusterStopWithOnOffParams {
	rv := objc.Send[MTRLevelControlClusterStopWithOnOffParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLevelControlClusterStopWithOnOffParams creates a new MTRLevelControlClusterStopWithOnOffParams instance.
func NewMTRLevelControlClusterStopWithOnOffParams() MTRLevelControlClusterStopWithOnOffParams {
	return getMTRLevelControlClusterStopWithOnOffParamsClass().New()
}




