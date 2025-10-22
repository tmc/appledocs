// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterNextParams] class.
var (
	MTRMediaPlaybackClusterNextParamsClass     _MTRMediaPlaybackClusterNextParamsClass
	MTRMediaPlaybackClusterNextParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterNextParamsClass() _MTRMediaPlaybackClusterNextParamsClass {
	MTRMediaPlaybackClusterNextParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterNextParamsClass = _MTRMediaPlaybackClusterNextParamsClass{objc.GetClass("MTRMediaPlaybackClusterNextParams")}
	})
	return MTRMediaPlaybackClusterNextParamsClass
}

type _MTRMediaPlaybackClusterNextParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterNextParams] class.
type IMTRMediaPlaybackClusterNextParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterNextParams
type MTRMediaPlaybackClusterNextParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterNextParamsFrom constructs a [MTRMediaPlaybackClusterNextParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterNextParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterNextParams {
	return MTRMediaPlaybackClusterNextParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterNextParamsClass) Alloc() MTRMediaPlaybackClusterNextParams {
	rv := objc.Send[MTRMediaPlaybackClusterNextParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterNextParamsClass) New() MTRMediaPlaybackClusterNextParams {
	rv := objc.Send[MTRMediaPlaybackClusterNextParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterNextParams) Init() MTRMediaPlaybackClusterNextParams {
	rv := objc.Send[MTRMediaPlaybackClusterNextParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterNextParams) Autorelease() MTRMediaPlaybackClusterNextParams {
	rv := objc.Send[MTRMediaPlaybackClusterNextParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterNextParams creates a new MTRMediaPlaybackClusterNextParams instance.
func NewMTRMediaPlaybackClusterNextParams() MTRMediaPlaybackClusterNextParams {
	return getMTRMediaPlaybackClusterNextParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterNextParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterNextParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterNextParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterNextParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



