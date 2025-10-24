// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterSkipBackwardParams] class.
var (
	MTRMediaPlaybackClusterSkipBackwardParamsClass     _MTRMediaPlaybackClusterSkipBackwardParamsClass
	MTRMediaPlaybackClusterSkipBackwardParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterSkipBackwardParamsClass() _MTRMediaPlaybackClusterSkipBackwardParamsClass {
	MTRMediaPlaybackClusterSkipBackwardParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterSkipBackwardParamsClass = _MTRMediaPlaybackClusterSkipBackwardParamsClass{objc.GetClass("MTRMediaPlaybackClusterSkipBackwardParams")}
	})
	return MTRMediaPlaybackClusterSkipBackwardParamsClass
}

type _MTRMediaPlaybackClusterSkipBackwardParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterSkipBackwardParams] class.
type IMTRMediaPlaybackClusterSkipBackwardParams interface {
	objectivec.IObject
	// properties:
	DeltaPositionMilliseconds() objc.IObject /* cross-framework: NSNumber */
	SetDeltaPositionMilliseconds(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterSkipBackwardParams
type MTRMediaPlaybackClusterSkipBackwardParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterSkipBackwardParamsFrom constructs a [MTRMediaPlaybackClusterSkipBackwardParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterSkipBackwardParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterSkipBackwardParams {
	return MTRMediaPlaybackClusterSkipBackwardParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterSkipBackwardParamsClass) Alloc() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterSkipBackwardParamsClass) New() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) Init() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) Autorelease() MTRMediaPlaybackClusterSkipBackwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipBackwardParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterSkipBackwardParams creates a new MTRMediaPlaybackClusterSkipBackwardParams instance.
func NewMTRMediaPlaybackClusterSkipBackwardParams() MTRMediaPlaybackClusterSkipBackwardParams {
	return getMTRMediaPlaybackClusterSkipBackwardParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/deltapositionmilliseconds
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) DeltaPositionMilliseconds() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("deltaPositionMilliseconds"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/deltapositionmilliseconds
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) SetDeltaPositionMilliseconds(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeltaPositionMilliseconds:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipbackwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSkipBackwardParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



