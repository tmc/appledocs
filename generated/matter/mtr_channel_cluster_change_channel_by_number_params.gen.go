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
	// properties:
	MajorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */)
	MinorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/majornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) MajorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("majorNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/majornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/minornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) MinorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minorNumber"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/minornumber
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelByNumberParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelByNumberParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelbynumberparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelByNumberParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



