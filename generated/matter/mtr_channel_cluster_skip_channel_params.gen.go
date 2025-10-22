// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterSkipChannelParams] class.
var (
	MTRChannelClusterSkipChannelParamsClass     _MTRChannelClusterSkipChannelParamsClass
	MTRChannelClusterSkipChannelParamsClassOnce sync.Once
)

func getMTRChannelClusterSkipChannelParamsClass() _MTRChannelClusterSkipChannelParamsClass {
	MTRChannelClusterSkipChannelParamsClassOnce.Do(func() {
		MTRChannelClusterSkipChannelParamsClass = _MTRChannelClusterSkipChannelParamsClass{objc.GetClass("MTRChannelClusterSkipChannelParams")}
	})
	return MTRChannelClusterSkipChannelParamsClass
}

type _MTRChannelClusterSkipChannelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterSkipChannelParams] class.
type IMTRChannelClusterSkipChannelParams interface {
	objectivec.IObject
	Count() foundation.Number
	SetCount(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterSkipChannelParams
type MTRChannelClusterSkipChannelParams struct {
	objectivec.Object
}

// MTRChannelClusterSkipChannelParamsFrom constructs a [MTRChannelClusterSkipChannelParams] from an unsafe.Pointer.
func MTRChannelClusterSkipChannelParamsFrom(ptr unsafe.Pointer) MTRChannelClusterSkipChannelParams {
	return MTRChannelClusterSkipChannelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterSkipChannelParamsClass) Alloc() MTRChannelClusterSkipChannelParams {
	rv := objc.Send[MTRChannelClusterSkipChannelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterSkipChannelParamsClass) New() MTRChannelClusterSkipChannelParams {
	rv := objc.Send[MTRChannelClusterSkipChannelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterSkipChannelParams) Init() MTRChannelClusterSkipChannelParams {
	rv := objc.Send[MTRChannelClusterSkipChannelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterSkipChannelParams) Autorelease() MTRChannelClusterSkipChannelParams {
	rv := objc.Send[MTRChannelClusterSkipChannelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterSkipChannelParams creates a new MTRChannelClusterSkipChannelParams instance.
func NewMTRChannelClusterSkipChannelParams() MTRChannelClusterSkipChannelParams {
	return getMTRChannelClusterSkipChannelParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/count
func (m_ MTRChannelClusterSkipChannelParams) Count() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("count"))
	return rv
}


// SetCount sets the value of the count property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/count
func (m_ MTRChannelClusterSkipChannelParams) SetCount(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCount:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/serversideprocessingtimeout
func (m_ MTRChannelClusterSkipChannelParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/serversideprocessingtimeout
func (m_ MTRChannelClusterSkipChannelParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/timedinvoketimeoutms
func (m_ MTRChannelClusterSkipChannelParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterskipchannelparams/timedinvoketimeoutms
func (m_ MTRChannelClusterSkipChannelParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



