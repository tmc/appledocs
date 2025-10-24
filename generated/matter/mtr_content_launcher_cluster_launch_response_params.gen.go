// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/data
func (m_ MTRContentLauncherClusterLaunchResponseParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/data
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/status
func (m_ MTRContentLauncherClusterLaunchResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/status
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchresponseparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
