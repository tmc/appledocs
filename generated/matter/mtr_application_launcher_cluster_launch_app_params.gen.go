// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRApplicationLauncherClusterLaunchAppParams] class.
var (
	MTRApplicationLauncherClusterLaunchAppParamsClass     _MTRApplicationLauncherClusterLaunchAppParamsClass
	MTRApplicationLauncherClusterLaunchAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterLaunchAppParamsClass() _MTRApplicationLauncherClusterLaunchAppParamsClass {
	MTRApplicationLauncherClusterLaunchAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterLaunchAppParamsClass = _MTRApplicationLauncherClusterLaunchAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterLaunchAppParams")}
	})
	return MTRApplicationLauncherClusterLaunchAppParamsClass
}

type _MTRApplicationLauncherClusterLaunchAppParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterLaunchAppParams] class.
type IMTRApplicationLauncherClusterLaunchAppParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterLaunchAppParams
type MTRApplicationLauncherClusterLaunchAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterLaunchAppParamsFrom constructs a [MTRApplicationLauncherClusterLaunchAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterLaunchAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterLaunchAppParams {
	return MTRApplicationLauncherClusterLaunchAppParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterLaunchAppParamsClass) Alloc() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterLaunchAppParamsClass) New() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Init() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Autorelease() MTRApplicationLauncherClusterLaunchAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterLaunchAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterLaunchAppParams creates a new MTRApplicationLauncherClusterLaunchAppParams instance.
func NewMTRApplicationLauncherClusterLaunchAppParams() MTRApplicationLauncherClusterLaunchAppParams {
	return getMTRApplicationLauncherClusterLaunchAppParamsClass().New()
}




