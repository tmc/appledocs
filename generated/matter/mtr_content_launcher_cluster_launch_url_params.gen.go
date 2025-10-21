// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterLaunchURLParams] class.
var (
	MTRContentLauncherClusterLaunchURLParamsClass     _MTRContentLauncherClusterLaunchURLParamsClass
	MTRContentLauncherClusterLaunchURLParamsClassOnce sync.Once
)

func getMTRContentLauncherClusterLaunchURLParamsClass() _MTRContentLauncherClusterLaunchURLParamsClass {
	MTRContentLauncherClusterLaunchURLParamsClassOnce.Do(func() {
		MTRContentLauncherClusterLaunchURLParamsClass = _MTRContentLauncherClusterLaunchURLParamsClass{objc.GetClass("MTRContentLauncherClusterLaunchURLParams")}
	})
	return MTRContentLauncherClusterLaunchURLParamsClass
}

type _MTRContentLauncherClusterLaunchURLParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterLaunchURLParams] class.
type IMTRContentLauncherClusterLaunchURLParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterLaunchURLParams
type MTRContentLauncherClusterLaunchURLParams struct {
	objectivec.Object
}

// MTRContentLauncherClusterLaunchURLParamsFrom constructs a [MTRContentLauncherClusterLaunchURLParams] from an unsafe.Pointer.
func MTRContentLauncherClusterLaunchURLParamsFrom(ptr unsafe.Pointer) MTRContentLauncherClusterLaunchURLParams {
	return MTRContentLauncherClusterLaunchURLParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterLaunchURLParamsClass) Alloc() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterLaunchURLParamsClass) New() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterLaunchURLParams) Init() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterLaunchURLParams) Autorelease() MTRContentLauncherClusterLaunchURLParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchURLParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterLaunchURLParams creates a new MTRContentLauncherClusterLaunchURLParams instance.
func NewMTRContentLauncherClusterLaunchURLParams() MTRContentLauncherClusterLaunchURLParams {
	return getMTRContentLauncherClusterLaunchURLParamsClass().New()
}




