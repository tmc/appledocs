// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
var (
	MTRApplicationLauncherClusterLauncherResponseParamsClass     _MTRApplicationLauncherClusterLauncherResponseParamsClass
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterLauncherResponseParamsClass() _MTRApplicationLauncherClusterLauncherResponseParamsClass {
	MTRApplicationLauncherClusterLauncherResponseParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterLauncherResponseParamsClass = _MTRApplicationLauncherClusterLauncherResponseParamsClass{objc.GetClass("MTRApplicationLauncherClusterLauncherResponseParams")}
	})
	return MTRApplicationLauncherClusterLauncherResponseParamsClass
}

type _MTRApplicationLauncherClusterLauncherResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterLauncherResponseParams] class.
type IMTRApplicationLauncherClusterLauncherResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLauncherResponseParams
type MTRApplicationLauncherClusterLauncherResponseParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterLauncherResponseParamsFrom constructs a [MTRApplicationLauncherClusterLauncherResponseParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterLauncherResponseParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterLauncherResponseParams {
	return MTRApplicationLauncherClusterLauncherResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) Alloc() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterLauncherResponseParamsClass) New() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Init() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterLauncherResponseParams) Autorelease() MTRApplicationLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRApplicationLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterLauncherResponseParams creates a new MTRApplicationLauncherClusterLauncherResponseParams instance.
func NewMTRApplicationLauncherClusterLauncherResponseParams() MTRApplicationLauncherClusterLauncherResponseParams {
	return getMTRApplicationLauncherClusterLauncherResponseParamsClass().New()
}




