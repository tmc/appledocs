// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRSmokeCOAlarmClusterSelfTestRequestParams] class.
var (
	MTRSmokeCOAlarmClusterSelfTestRequestParamsClass     _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass
	MTRSmokeCOAlarmClusterSelfTestRequestParamsClassOnce sync.Once
)

func getMTRSmokeCOAlarmClusterSelfTestRequestParamsClass() _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass {
	MTRSmokeCOAlarmClusterSelfTestRequestParamsClassOnce.Do(func() {
		MTRSmokeCOAlarmClusterSelfTestRequestParamsClass = _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass{objc.GetClass("MTRSmokeCOAlarmClusterSelfTestRequestParams")}
	})
	return MTRSmokeCOAlarmClusterSelfTestRequestParamsClass
}

type _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRSmokeCOAlarmClusterSelfTestRequestParams] class.
type IMTRSmokeCOAlarmClusterSelfTestRequestParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSmokeCOAlarmClusterSelfTestRequestParams
type MTRSmokeCOAlarmClusterSelfTestRequestParams struct {
	objectivec.Object
}

// MTRSmokeCOAlarmClusterSelfTestRequestParamsFrom constructs a [MTRSmokeCOAlarmClusterSelfTestRequestParams] from an unsafe.Pointer.
func MTRSmokeCOAlarmClusterSelfTestRequestParamsFrom(ptr unsafe.Pointer) MTRSmokeCOAlarmClusterSelfTestRequestParams {
	return MTRSmokeCOAlarmClusterSelfTestRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass) Alloc() MTRSmokeCOAlarmClusterSelfTestRequestParams {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRSmokeCOAlarmClusterSelfTestRequestParamsClass) New() MTRSmokeCOAlarmClusterSelfTestRequestParams {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) Init() MTRSmokeCOAlarmClusterSelfTestRequestParams {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) Autorelease() MTRSmokeCOAlarmClusterSelfTestRequestParams {
	rv := objc.Send[MTRSmokeCOAlarmClusterSelfTestRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSmokeCOAlarmClusterSelfTestRequestParams creates a new MTRSmokeCOAlarmClusterSelfTestRequestParams instance.
func NewMTRSmokeCOAlarmClusterSelfTestRequestParams() MTRSmokeCOAlarmClusterSelfTestRequestParams {
	return getMTRSmokeCOAlarmClusterSelfTestRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterselftestrequestparams/serversideprocessingtimeout
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterselftestrequestparams/serversideprocessingtimeout
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterselftestrequestparams/timedinvoketimeoutms
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrsmokecoalarmclusterselftestrequestparams/timedinvoketimeoutms
func (m_ MTRSmokeCOAlarmClusterSelfTestRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



