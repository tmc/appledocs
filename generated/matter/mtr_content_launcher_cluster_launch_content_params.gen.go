// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AutoPlay() objc.IObject /* cross-framework: NSNumber */
	SetAutoPlay(value objc.IObject /* cross-framework: NSNumber */)
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	Search() IMTRContentLauncherClusterContentSearchStruct
	SetSearch(value IMTRContentLauncherClusterContentSearchStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UseCurrentContext() objc.IObject /* cross-framework: NSNumber */
	SetUseCurrentContext(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/autoplay
func (m_ MTRContentLauncherClusterLaunchContentParams) AutoPlay() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("autoPlay"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/autoplay
func (m_ MTRContentLauncherClusterLaunchContentParams) SetAutoPlay(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoPlay:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/data
func (m_ MTRContentLauncherClusterLaunchContentParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/data
func (m_ MTRContentLauncherClusterLaunchContentParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/search
func (m_ MTRContentLauncherClusterLaunchContentParams) Search() IMTRContentLauncherClusterContentSearchStruct {
	rv := objc.Send[MTRContentLauncherClusterContentSearchStruct](m_.ID, objc.Sel("search"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/search
func (m_ MTRContentLauncherClusterLaunchContentParams) SetSearch(value IMTRContentLauncherClusterContentSearchStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearch:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchContentParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchContentParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchContentParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchContentParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/usecurrentcontext
func (m_ MTRContentLauncherClusterLaunchContentParams) UseCurrentContext() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("useCurrentContext"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/usecurrentcontext
func (m_ MTRContentLauncherClusterLaunchContentParams) SetUseCurrentContext(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUseCurrentContext:"), value)
}



