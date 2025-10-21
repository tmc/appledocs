// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChangeChannelByNumberParams] class.
var (
	MTRChannelClusterChangeChannelByNumberParamsClass     _MTRChannelClusterChangeChannelByNumberParamsClass
	MTRChannelClusterChangeChannelByNumberParamsClassOnce sync.Once
)

func getMTRChannelClusterChangeChannelByNumberParamsClass() _MTRChannelClusterChangeChannelByNumberParamsClass {
	MTRChannelClusterChangeChannelByNumberParamsClassOnce.Do(func() {
		MTRChannelClusterChangeChannelByNumberParamsClass = _MTRChannelClusterChangeChannelByNumberParamsClass{objc.GetClass("MTRChannelClusterChangeChannelByNumberParams")}
	})
	return MTRChannelClusterChangeChannelByNumberParamsClass
}

type _MTRChannelClusterChangeChannelByNumberParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChangeChannelByNumberParams] class.
type IMTRChannelClusterChangeChannelByNumberParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChangeChannelByNumberParams
type MTRChannelClusterChangeChannelByNumberParams struct {
	objectivec.Object
}

// MTRChannelClusterChangeChannelByNumberParamsFrom constructs a [MTRChannelClusterChangeChannelByNumberParams] from an unsafe.Pointer.
func MTRChannelClusterChangeChannelByNumberParamsFrom(ptr unsafe.Pointer) MTRChannelClusterChangeChannelByNumberParams {
	return MTRChannelClusterChangeChannelByNumberParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChangeChannelByNumberParamsClass) Alloc() MTRChannelClusterChangeChannelByNumberParams {
	rv := objc.Send[MTRChannelClusterChangeChannelByNumberParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChangeChannelByNumberParamsClass) New() MTRChannelClusterChangeChannelByNumberParams {
	rv := objc.Send[MTRChannelClusterChangeChannelByNumberParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChangeChannelByNumberParams) Init() MTRChannelClusterChangeChannelByNumberParams {
	rv := objc.Send[MTRChannelClusterChangeChannelByNumberParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChangeChannelByNumberParams) Autorelease() MTRChannelClusterChangeChannelByNumberParams {
	rv := objc.Send[MTRChannelClusterChangeChannelByNumberParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChangeChannelByNumberParams creates a new MTRChannelClusterChangeChannelByNumberParams instance.
func NewMTRChannelClusterChangeChannelByNumberParams() MTRChannelClusterChangeChannelByNumberParams {
	return getMTRChannelClusterChangeChannelByNumberParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/majornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) MajorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("majorNumber"))
	return rv
}


// SetMajorNumber sets the value of the majorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/majornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetMajorNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/minornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) MinorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minorNumber"))
	return rv
}


// SetMinorNumber sets the value of the minorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/minornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetMinorNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelByNumberParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelByNumberParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



