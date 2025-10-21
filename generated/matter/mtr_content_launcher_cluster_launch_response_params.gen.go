// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRContentLauncherClusterLaunchResponseParams] class.
var (
	MTRContentLauncherClusterLaunchResponseParamsClass     _MTRContentLauncherClusterLaunchResponseParamsClass
	MTRContentLauncherClusterLaunchResponseParamsClassOnce sync.Once
)

func getMTRContentLauncherClusterLaunchResponseParamsClass() _MTRContentLauncherClusterLaunchResponseParamsClass {
	MTRContentLauncherClusterLaunchResponseParamsClassOnce.Do(func() {
		MTRContentLauncherClusterLaunchResponseParamsClass = _MTRContentLauncherClusterLaunchResponseParamsClass{objc.GetClass("MTRContentLauncherClusterLaunchResponseParams")}
	})
	return MTRContentLauncherClusterLaunchResponseParamsClass
}

type _MTRContentLauncherClusterLaunchResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterLaunchResponseParams] class.
type IMTRContentLauncherClusterLaunchResponseParams interface {
	IMTRContentLauncherClusterLauncherResponseParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterLaunchResponseParams
type MTRContentLauncherClusterLaunchResponseParams struct {
	MTRContentLauncherClusterLauncherResponseParams
}

// MTRContentLauncherClusterLaunchResponseParamsFrom constructs a [MTRContentLauncherClusterLaunchResponseParams] from an unsafe.Pointer.
func MTRContentLauncherClusterLaunchResponseParamsFrom(ptr unsafe.Pointer) MTRContentLauncherClusterLaunchResponseParams {
	return MTRContentLauncherClusterLaunchResponseParams{
		MTRContentLauncherClusterLauncherResponseParams: MTRContentLauncherClusterLauncherResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterLaunchResponseParamsClass) Alloc() MTRContentLauncherClusterLaunchResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterLaunchResponseParamsClass) New() MTRContentLauncherClusterLaunchResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterLaunchResponseParams) Init() MTRContentLauncherClusterLaunchResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterLaunchResponseParams) Autorelease() MTRContentLauncherClusterLaunchResponseParams {
	rv := objc.Send[MTRContentLauncherClusterLaunchResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterLaunchResponseParams creates a new MTRContentLauncherClusterLaunchResponseParams instance.
func NewMTRContentLauncherClusterLaunchResponseParams() MTRContentLauncherClusterLaunchResponseParams {
	return getMTRContentLauncherClusterLaunchResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/data
func (m_ MTRContentLauncherClusterLaunchResponseParams) Data() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/data
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetData(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/status
func (m_ MTRContentLauncherClusterLaunchResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/status
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



