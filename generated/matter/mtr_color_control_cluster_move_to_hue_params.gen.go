// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRColorControlClusterMoveToHueParams] class.
var (
	MTRColorControlClusterMoveToHueParamsClass     _MTRColorControlClusterMoveToHueParamsClass
	MTRColorControlClusterMoveToHueParamsClassOnce sync.Once
)

func getMTRColorControlClusterMoveToHueParamsClass() _MTRColorControlClusterMoveToHueParamsClass {
	MTRColorControlClusterMoveToHueParamsClassOnce.Do(func() {
		MTRColorControlClusterMoveToHueParamsClass = _MTRColorControlClusterMoveToHueParamsClass{objc.GetClass("MTRColorControlClusterMoveToHueParams")}
	})
	return MTRColorControlClusterMoveToHueParamsClass
}

type _MTRColorControlClusterMoveToHueParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRColorControlClusterMoveToHueParams] class.
type IMTRColorControlClusterMoveToHueParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRColorControlClusterMoveToHueParams
type MTRColorControlClusterMoveToHueParams struct {
	objectivec.Object
}

// MTRColorControlClusterMoveToHueParamsFrom constructs a [MTRColorControlClusterMoveToHueParams] from an unsafe.Pointer.
func MTRColorControlClusterMoveToHueParamsFrom(ptr unsafe.Pointer) MTRColorControlClusterMoveToHueParams {
	return MTRColorControlClusterMoveToHueParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRColorControlClusterMoveToHueParamsClass) Alloc() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRColorControlClusterMoveToHueParamsClass) New() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRColorControlClusterMoveToHueParams) Init() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRColorControlClusterMoveToHueParams) Autorelease() MTRColorControlClusterMoveToHueParams {
	rv := objc.Send[MTRColorControlClusterMoveToHueParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRColorControlClusterMoveToHueParams creates a new MTRColorControlClusterMoveToHueParams instance.
func NewMTRColorControlClusterMoveToHueParams() MTRColorControlClusterMoveToHueParams {
	return getMTRColorControlClusterMoveToHueParamsClass().New()
}




