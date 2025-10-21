// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentLauncherClusterLauncherResponseParams] class.
var (
	MTRContentLauncherClusterLauncherResponseParamsClass     _MTRContentLauncherClusterLauncherResponseParamsClass
	MTRContentLauncherClusterLauncherResponseParamsClassOnce sync.Once
)

func getMTRContentLauncherClusterLauncherResponseParamsClass() _MTRContentLauncherClusterLauncherResponseParamsClass {
	MTRContentLauncherClusterLauncherResponseParamsClassOnce.Do(func() {
		MTRContentLauncherClusterLauncherResponseParamsClass = _MTRContentLauncherClusterLauncherResponseParamsClass{objc.GetClass("MTRContentLauncherClusterLauncherResponseParams")}
	})
	return MTRContentLauncherClusterLauncherResponseParamsClass
}

type _MTRContentLauncherClusterLauncherResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterLauncherResponseParams] class.
type IMTRContentLauncherClusterLauncherResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterLauncherResponseParams
type MTRContentLauncherClusterLauncherResponseParams struct {
	objectivec.Object
}

// MTRContentLauncherClusterLauncherResponseParamsFrom constructs a [MTRContentLauncherClusterLauncherResponseParams] from an unsafe.Pointer.
func MTRContentLauncherClusterLauncherResponseParamsFrom(ptr unsafe.Pointer) MTRContentLauncherClusterLauncherResponseParams {
	return MTRContentLauncherClusterLauncherResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterLauncherResponseParamsClass) Alloc() MTRContentLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterLauncherResponseParamsClass) New() MTRContentLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLauncherResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterLauncherResponseParams) Init() MTRContentLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterLauncherResponseParams) Autorelease() MTRContentLauncherClusterLauncherResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLauncherResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterLauncherResponseParams creates a new MTRContentLauncherClusterLauncherResponseParams instance.
func NewMTRContentLauncherClusterLauncherResponseParams() MTRContentLauncherClusterLauncherResponseParams {
	return getMTRContentLauncherClusterLauncherResponseParamsClass().New()
}




