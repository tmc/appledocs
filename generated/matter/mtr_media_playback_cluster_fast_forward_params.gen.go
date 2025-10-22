// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterFastForwardParams] class.
var (
	MTRMediaPlaybackClusterFastForwardParamsClass     _MTRMediaPlaybackClusterFastForwardParamsClass
	MTRMediaPlaybackClusterFastForwardParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterFastForwardParamsClass() _MTRMediaPlaybackClusterFastForwardParamsClass {
	MTRMediaPlaybackClusterFastForwardParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterFastForwardParamsClass = _MTRMediaPlaybackClusterFastForwardParamsClass{objc.GetClass("MTRMediaPlaybackClusterFastForwardParams")}
	})
	return MTRMediaPlaybackClusterFastForwardParamsClass
}

type _MTRMediaPlaybackClusterFastForwardParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterFastForwardParams] class.
type IMTRMediaPlaybackClusterFastForwardParams interface {
	objectivec.IObject
	AudioAdvanceUnmuted() foundation.Number
	SetAudioAdvanceUnmuted(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterFastForwardParams
type MTRMediaPlaybackClusterFastForwardParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterFastForwardParamsFrom constructs a [MTRMediaPlaybackClusterFastForwardParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterFastForwardParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterFastForwardParams {
	return MTRMediaPlaybackClusterFastForwardParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterFastForwardParamsClass) Alloc() MTRMediaPlaybackClusterFastForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterFastForwardParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterFastForwardParamsClass) New() MTRMediaPlaybackClusterFastForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterFastForwardParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterFastForwardParams) Init() MTRMediaPlaybackClusterFastForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterFastForwardParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterFastForwardParams) Autorelease() MTRMediaPlaybackClusterFastForwardParams {
	rv := objc.Send[MTRMediaPlaybackClusterFastForwardParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterFastForwardParams creates a new MTRMediaPlaybackClusterFastForwardParams instance.
func NewMTRMediaPlaybackClusterFastForwardParams() MTRMediaPlaybackClusterFastForwardParams {
	return getMTRMediaPlaybackClusterFastForwardParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/audioadvanceunmuted
func (m_ MTRMediaPlaybackClusterFastForwardParams) AudioAdvanceUnmuted() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("audioAdvanceUnmuted"))
	return rv
}


// SetAudioAdvanceUnmuted sets the value of the audioAdvanceUnmuted property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/audioadvanceunmuted
func (m_ MTRMediaPlaybackClusterFastForwardParams) SetAudioAdvanceUnmuted(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioAdvanceUnmuted:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterFastForwardParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterFastForwardParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterFastForwardParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterfastforwardparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterFastForwardParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



