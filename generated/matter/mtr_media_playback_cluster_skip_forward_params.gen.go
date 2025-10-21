// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterSkipForwardParams] class.
var (
	MTRMediaPlaybackClusterSkipForwardParamsClass     _MTRMediaPlaybackClusterSkipForwardParamsClass
	MTRMediaPlaybackClusterSkipForwardParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterSkipForwardParamsClass() _MTRMediaPlaybackClusterSkipForwardParamsClass {
	MTRMediaPlaybackClusterSkipForwardParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterSkipForwardParamsClass = _MTRMediaPlaybackClusterSkipForwardParamsClass{objc.GetClass("MTRMediaPlaybackClusterSkipForwardParams")}
	})
	return MTRMediaPlaybackClusterSkipForwardParamsClass
}

type _MTRMediaPlaybackClusterSkipForwardParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterSkipForwardParams] class.
type IMTRMediaPlaybackClusterSkipForwardParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterSkipForwardParams
type MTRMediaPlaybackClusterSkipForwardParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterSkipForwardParamsFrom constructs a [MTRMediaPlaybackClusterSkipForwardParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterSkipForwardParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterSkipForwardParams {
	return MTRMediaPlaybackClusterSkipForwardParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterSkipForwardParamsClass) Alloc() MTRMediaPlaybackClusterSkipForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipForwardParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterSkipForwardParamsClass) New() MTRMediaPlaybackClusterSkipForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipForwardParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterSkipForwardParams) Init() MTRMediaPlaybackClusterSkipForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipForwardParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterSkipForwardParams) Autorelease() MTRMediaPlaybackClusterSkipForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterSkipForwardParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterSkipForwardParams creates a new MTRMediaPlaybackClusterSkipForwardParams instance.
func NewMTRMediaPlaybackClusterSkipForwardParams() MTRMediaPlaybackClusterSkipForwardParams {
	return getMTRMediaPlaybackClusterSkipForwardParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/deltapositionmilliseconds
func (m_ MTRMediaPlaybackClusterSkipForwardParams) DeltaPositionMilliseconds() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("deltaPositionMilliseconds"))
	return rv
}


// SetDeltaPositionMilliseconds sets the value of the deltaPositionMilliseconds property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/deltapositionmilliseconds
func (m_ MTRMediaPlaybackClusterSkipForwardParams) SetDeltaPositionMilliseconds(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeltaPositionMilliseconds:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSkipForwardParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSkipForwardParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSkipForwardParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterskipforwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSkipForwardParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



