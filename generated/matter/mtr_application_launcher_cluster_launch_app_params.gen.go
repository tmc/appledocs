// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/application
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Application() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("application"))
	return rv
}


// SetApplication sets the value of the application property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/application
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetApplication(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/data
func (m_ MTRApplicationLauncherClusterLaunchAppParams) Data() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/data
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetData(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterLaunchAppParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterLaunchAppParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterlaunchappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterLaunchAppParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



