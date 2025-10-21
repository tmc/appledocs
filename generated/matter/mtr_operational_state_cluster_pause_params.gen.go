// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalStateClusterPauseParams] class.
var (
	MTROperationalStateClusterPauseParamsClass     _MTROperationalStateClusterPauseParamsClass
	MTROperationalStateClusterPauseParamsClassOnce sync.Once
)

func getMTROperationalStateClusterPauseParamsClass() _MTROperationalStateClusterPauseParamsClass {
	MTROperationalStateClusterPauseParamsClassOnce.Do(func() {
		MTROperationalStateClusterPauseParamsClass = _MTROperationalStateClusterPauseParamsClass{objc.GetClass("MTROperationalStateClusterPauseParams")}
	})
	return MTROperationalStateClusterPauseParamsClass
}

type _MTROperationalStateClusterPauseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterPauseParams] class.
type IMTROperationalStateClusterPauseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterPauseParams
type MTROperationalStateClusterPauseParams struct {
	objectivec.Object
}

// MTROperationalStateClusterPauseParamsFrom constructs a [MTROperationalStateClusterPauseParams] from an unsafe.Pointer.
func MTROperationalStateClusterPauseParamsFrom(ptr unsafe.Pointer) MTROperationalStateClusterPauseParams {
	return MTROperationalStateClusterPauseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterPauseParamsClass) Alloc() MTROperationalStateClusterPauseParams {
	rv := objc.Send[MTROperationalStateClusterPauseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterPauseParamsClass) New() MTROperationalStateClusterPauseParams {
	rv := objc.Send[MTROperationalStateClusterPauseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterPauseParams) Init() MTROperationalStateClusterPauseParams {
	rv := objc.Send[MTROperationalStateClusterPauseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterPauseParams) Autorelease() MTROperationalStateClusterPauseParams {
	rv := objc.Send[MTROperationalStateClusterPauseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterPauseParams creates a new MTROperationalStateClusterPauseParams instance.
func NewMTROperationalStateClusterPauseParams() MTROperationalStateClusterPauseParams {
	return getMTROperationalStateClusterPauseParamsClass().New()
}




