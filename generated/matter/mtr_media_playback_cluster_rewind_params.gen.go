// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterRewindParams] class.
var (
	MTRMediaPlaybackClusterRewindParamsClass     _MTRMediaPlaybackClusterRewindParamsClass
	MTRMediaPlaybackClusterRewindParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterRewindParamsClass() _MTRMediaPlaybackClusterRewindParamsClass {
	MTRMediaPlaybackClusterRewindParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterRewindParamsClass = _MTRMediaPlaybackClusterRewindParamsClass{objc.GetClass("MTRMediaPlaybackClusterRewindParams")}
	})
	return MTRMediaPlaybackClusterRewindParamsClass
}

type _MTRMediaPlaybackClusterRewindParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterRewindParams] class.
type IMTRMediaPlaybackClusterRewindParams interface {
	objectivec.IObject
	AudioAdvanceUnmuted() foundation.Number
	SetAudioAdvanceUnmuted(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterRewindParams
type MTRMediaPlaybackClusterRewindParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterRewindParamsFrom constructs a [MTRMediaPlaybackClusterRewindParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterRewindParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterRewindParams {
	return MTRMediaPlaybackClusterRewindParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterRewindParamsClass) Alloc() MTRMediaPlaybackClusterRewindParams {
	rv := objc.Send[MTRMediaPlaybackClusterRewindParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterRewindParamsClass) New() MTRMediaPlaybackClusterRewindParams {
	rv := objc.Send[MTRMediaPlaybackClusterRewindParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterRewindParams) Init() MTRMediaPlaybackClusterRewindParams {
	rv := objc.Send[MTRMediaPlaybackClusterRewindParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterRewindParams) Autorelease() MTRMediaPlaybackClusterRewindParams {
	rv := objc.Send[MTRMediaPlaybackClusterRewindParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterRewindParams creates a new MTRMediaPlaybackClusterRewindParams instance.
func NewMTRMediaPlaybackClusterRewindParams() MTRMediaPlaybackClusterRewindParams {
	return getMTRMediaPlaybackClusterRewindParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/audioadvanceunmuted
func (m_ MTRMediaPlaybackClusterRewindParams) AudioAdvanceUnmuted() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("audioAdvanceUnmuted"))
	return rv
}


// SetAudioAdvanceUnmuted sets the value of the audioAdvanceUnmuted property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/audioadvanceunmuted
func (m_ MTRMediaPlaybackClusterRewindParams) SetAudioAdvanceUnmuted(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioAdvanceUnmuted:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterRewindParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterRewindParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterRewindParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterrewindparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterRewindParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



