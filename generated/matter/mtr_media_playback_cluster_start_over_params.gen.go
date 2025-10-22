// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterStartOverParams] class.
var (
	MTRMediaPlaybackClusterStartOverParamsClass     _MTRMediaPlaybackClusterStartOverParamsClass
	MTRMediaPlaybackClusterStartOverParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterStartOverParamsClass() _MTRMediaPlaybackClusterStartOverParamsClass {
	MTRMediaPlaybackClusterStartOverParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterStartOverParamsClass = _MTRMediaPlaybackClusterStartOverParamsClass{objc.GetClass("MTRMediaPlaybackClusterStartOverParams")}
	})
	return MTRMediaPlaybackClusterStartOverParamsClass
}

type _MTRMediaPlaybackClusterStartOverParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterStartOverParams] class.
type IMTRMediaPlaybackClusterStartOverParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterStartOverParams
type MTRMediaPlaybackClusterStartOverParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterStartOverParamsFrom constructs a [MTRMediaPlaybackClusterStartOverParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterStartOverParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterStartOverParams {
	return MTRMediaPlaybackClusterStartOverParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterStartOverParamsClass) Alloc() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterStartOverParamsClass) New() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterStartOverParams) Init() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterStartOverParams) Autorelease() MTRMediaPlaybackClusterStartOverParams {
	rv := objc.Send[MTRMediaPlaybackClusterStartOverParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterStartOverParams creates a new MTRMediaPlaybackClusterStartOverParams instance.
func NewMTRMediaPlaybackClusterStartOverParams() MTRMediaPlaybackClusterStartOverParams {
	return getMTRMediaPlaybackClusterStartOverParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstartoverparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterStartOverParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstartoverparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterStartOverParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstartoverparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterStartOverParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterstartoverparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterStartOverParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



