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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/autoplay
func (m_ MTRContentLauncherClusterLaunchContentParams) AutoPlay() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("autoPlay"))
	return rv
}


// SetAutoPlay sets the value of the autoPlay property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/autoplay
func (m_ MTRContentLauncherClusterLaunchContentParams) SetAutoPlay(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAutoPlay:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/data
func (m_ MTRContentLauncherClusterLaunchContentParams) Data() string {
	rv := objc.Send[string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/data
func (m_ MTRContentLauncherClusterLaunchContentParams) SetData(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/search
func (m_ MTRContentLauncherClusterLaunchContentParams) Search() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("search"))
	return rv
}


// SetSearch sets the value of the search property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/search
func (m_ MTRContentLauncherClusterLaunchContentParams) SetSearch(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSearch:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchContentParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/serversideprocessingtimeout
func (m_ MTRContentLauncherClusterLaunchContentParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchContentParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/timedinvoketimeoutms
func (m_ MTRContentLauncherClusterLaunchContentParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/usecurrentcontext
func (m_ MTRContentLauncherClusterLaunchContentParams) UseCurrentContext() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("useCurrentContext"))
	return rv
}


// SetUseCurrentContext sets the value of the useCurrentContext property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterlaunchcontentparams/usecurrentcontext
func (m_ MTRContentLauncherClusterLaunchContentParams) SetUseCurrentContext(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUseCurrentContext:"), value)
}



