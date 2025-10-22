// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterSeekParams] class.
var (
	MTRMediaPlaybackClusterSeekParamsClass     _MTRMediaPlaybackClusterSeekParamsClass
	MTRMediaPlaybackClusterSeekParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterSeekParamsClass() _MTRMediaPlaybackClusterSeekParamsClass {
	MTRMediaPlaybackClusterSeekParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterSeekParamsClass = _MTRMediaPlaybackClusterSeekParamsClass{objc.GetClass("MTRMediaPlaybackClusterSeekParams")}
	})
	return MTRMediaPlaybackClusterSeekParamsClass
}

type _MTRMediaPlaybackClusterSeekParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterSeekParams] class.
type IMTRMediaPlaybackClusterSeekParams interface {
	objectivec.IObject
	Position() foundation.Number
	SetPosition(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterSeekParams
type MTRMediaPlaybackClusterSeekParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterSeekParamsFrom constructs a [MTRMediaPlaybackClusterSeekParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterSeekParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterSeekParams {
	return MTRMediaPlaybackClusterSeekParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterSeekParamsClass) Alloc() MTRMediaPlaybackClusterSeekParams {
	rv := objc.Send[MTRMediaPlaybackClusterSeekParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterSeekParamsClass) New() MTRMediaPlaybackClusterSeekParams {
	rv := objc.Send[MTRMediaPlaybackClusterSeekParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterSeekParams) Init() MTRMediaPlaybackClusterSeekParams {
	rv := objc.Send[MTRMediaPlaybackClusterSeekParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterSeekParams) Autorelease() MTRMediaPlaybackClusterSeekParams {
	rv := objc.Send[MTRMediaPlaybackClusterSeekParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterSeekParams creates a new MTRMediaPlaybackClusterSeekParams instance.
func NewMTRMediaPlaybackClusterSeekParams() MTRMediaPlaybackClusterSeekParams {
	return getMTRMediaPlaybackClusterSeekParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/position
func (m_ MTRMediaPlaybackClusterSeekParams) Position() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/position
func (m_ MTRMediaPlaybackClusterSeekParams) SetPosition(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPosition:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSeekParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterSeekParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSeekParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterseekparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterSeekParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



