// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalStateClusterStartParams] class.
var (
	MTROperationalStateClusterStartParamsClass     _MTROperationalStateClusterStartParamsClass
	MTROperationalStateClusterStartParamsClassOnce sync.Once
)

func getMTROperationalStateClusterStartParamsClass() _MTROperationalStateClusterStartParamsClass {
	MTROperationalStateClusterStartParamsClassOnce.Do(func() {
		MTROperationalStateClusterStartParamsClass = _MTROperationalStateClusterStartParamsClass{objc.GetClass("MTROperationalStateClusterStartParams")}
	})
	return MTROperationalStateClusterStartParamsClass
}

type _MTROperationalStateClusterStartParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterStartParams] class.
type IMTROperationalStateClusterStartParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterStartParams
type MTROperationalStateClusterStartParams struct {
	objectivec.Object
}

// MTROperationalStateClusterStartParamsFrom constructs a [MTROperationalStateClusterStartParams] from an unsafe.Pointer.
func MTROperationalStateClusterStartParamsFrom(ptr unsafe.Pointer) MTROperationalStateClusterStartParams {
	return MTROperationalStateClusterStartParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterStartParamsClass) Alloc() MTROperationalStateClusterStartParams {
	rv := objc.Send[MTROperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterStartParamsClass) New() MTROperationalStateClusterStartParams {
	rv := objc.Send[MTROperationalStateClusterStartParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterStartParams) Init() MTROperationalStateClusterStartParams {
	rv := objc.Send[MTROperationalStateClusterStartParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterStartParams) Autorelease() MTROperationalStateClusterStartParams {
	rv := objc.Send[MTROperationalStateClusterStartParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterStartParams creates a new MTROperationalStateClusterStartParams instance.
func NewMTROperationalStateClusterStartParams() MTROperationalStateClusterStartParams {
	return getMTROperationalStateClusterStartParamsClass().New()
}




