// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationLauncherClusterStopAppParams] class.
var (
	MTRApplicationLauncherClusterStopAppParamsClass     _MTRApplicationLauncherClusterStopAppParamsClass
	MTRApplicationLauncherClusterStopAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterStopAppParamsClass() _MTRApplicationLauncherClusterStopAppParamsClass {
	MTRApplicationLauncherClusterStopAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterStopAppParamsClass = _MTRApplicationLauncherClusterStopAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterStopAppParams")}
	})
	return MTRApplicationLauncherClusterStopAppParamsClass
}

type _MTRApplicationLauncherClusterStopAppParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterStopAppParams] class.
type IMTRApplicationLauncherClusterStopAppParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterStopAppParams
type MTRApplicationLauncherClusterStopAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterStopAppParamsFrom constructs a [MTRApplicationLauncherClusterStopAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterStopAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterStopAppParams {
	return MTRApplicationLauncherClusterStopAppParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterStopAppParamsClass) Alloc() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterStopAppParamsClass) New() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterStopAppParams) Init() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterStopAppParams) Autorelease() MTRApplicationLauncherClusterStopAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterStopAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterStopAppParams creates a new MTRApplicationLauncherClusterStopAppParams instance.
func NewMTRApplicationLauncherClusterStopAppParams() MTRApplicationLauncherClusterStopAppParams {
	return getMTRApplicationLauncherClusterStopAppParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/application
func (m_ MTRApplicationLauncherClusterStopAppParams) Application() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("application"))
	return rv
}


// SetApplication sets the value of the application property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/application
func (m_ MTRApplicationLauncherClusterStopAppParams) SetApplication(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterStopAppParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterStopAppParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterStopAppParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterstopappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterStopAppParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



