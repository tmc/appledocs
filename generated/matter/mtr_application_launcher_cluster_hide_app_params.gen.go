// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationLauncherClusterHideAppParams] class.
var (
	MTRApplicationLauncherClusterHideAppParamsClass     _MTRApplicationLauncherClusterHideAppParamsClass
	MTRApplicationLauncherClusterHideAppParamsClassOnce sync.Once
)

func getMTRApplicationLauncherClusterHideAppParamsClass() _MTRApplicationLauncherClusterHideAppParamsClass {
	MTRApplicationLauncherClusterHideAppParamsClassOnce.Do(func() {
		MTRApplicationLauncherClusterHideAppParamsClass = _MTRApplicationLauncherClusterHideAppParamsClass{objc.GetClass("MTRApplicationLauncherClusterHideAppParams")}
	})
	return MTRApplicationLauncherClusterHideAppParamsClass
}

type _MTRApplicationLauncherClusterHideAppParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterHideAppParams] class.
type IMTRApplicationLauncherClusterHideAppParams interface {
	objectivec.IObject
	Application() MTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterHideAppParams
type MTRApplicationLauncherClusterHideAppParams struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterHideAppParamsFrom constructs a [MTRApplicationLauncherClusterHideAppParams] from an unsafe.Pointer.
func MTRApplicationLauncherClusterHideAppParamsFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterHideAppParams {
	return MTRApplicationLauncherClusterHideAppParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterHideAppParamsClass) Alloc() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterHideAppParamsClass) New() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterHideAppParams) Init() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterHideAppParams) Autorelease() MTRApplicationLauncherClusterHideAppParams {
	rv := objc.Send[MTRApplicationLauncherClusterHideAppParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterHideAppParams creates a new MTRApplicationLauncherClusterHideAppParams instance.
func NewMTRApplicationLauncherClusterHideAppParams() MTRApplicationLauncherClusterHideAppParams {
	return getMTRApplicationLauncherClusterHideAppParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/application
func (m_ MTRApplicationLauncherClusterHideAppParams) Application() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}


// SetApplication sets the value of the application property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/application
func (m_ MTRApplicationLauncherClusterHideAppParams) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterHideAppParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/serversideprocessingtimeout
func (m_ MTRApplicationLauncherClusterHideAppParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterHideAppParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterhideappparams/timedinvoketimeoutms
func (m_ MTRApplicationLauncherClusterHideAppParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



