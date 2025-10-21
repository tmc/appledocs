// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterLaunchContentParams] class.
var (
	MTRContentLauncherClusterLaunchContentParamsClass     _MTRContentLauncherClusterLaunchContentParamsClass
	MTRContentLauncherClusterLaunchContentParamsClassOnce sync.Once
)

func getMTRContentLauncherClusterLaunchContentParamsClass() _MTRContentLauncherClusterLaunchContentParamsClass {
	MTRContentLauncherClusterLaunchContentParamsClassOnce.Do(func() {
		MTRContentLauncherClusterLaunchContentParamsClass = _MTRContentLauncherClusterLaunchContentParamsClass{objc.GetClass("MTRContentLauncherClusterLaunchContentParams")}
	})
	return MTRContentLauncherClusterLaunchContentParamsClass
}

type _MTRContentLauncherClusterLaunchContentParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterLaunchContentParams] class.
type IMTRContentLauncherClusterLaunchContentParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterLaunchContentParams
type MTRContentLauncherClusterLaunchContentParams struct {
	objectivec.Object
}

// MTRContentLauncherClusterLaunchContentParamsFrom constructs a [MTRContentLauncherClusterLaunchContentParams] from an unsafe.Pointer.
func MTRContentLauncherClusterLaunchContentParamsFrom(ptr unsafe.Pointer) MTRContentLauncherClusterLaunchContentParams {
	return MTRContentLauncherClusterLaunchContentParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterLaunchContentParamsClass) Alloc() MTRContentLauncherClusterLaunchContentParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchContentParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterLaunchContentParamsClass) New() MTRContentLauncherClusterLaunchContentParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchContentParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterLaunchContentParams) Init() MTRContentLauncherClusterLaunchContentParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchContentParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterLaunchContentParams) Autorelease() MTRContentLauncherClusterLaunchContentParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchContentParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterLaunchContentParams creates a new MTRContentLauncherClusterLaunchContentParams instance.
func NewMTRContentLauncherClusterLaunchContentParams() MTRContentLauncherClusterLaunchContentParams {
	return getMTRContentLauncherClusterLaunchContentParamsClass().New()
}




