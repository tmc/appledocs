// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChangeChannelParams] class.
var (
	MTRChannelClusterChangeChannelParamsClass     _MTRChannelClusterChangeChannelParamsClass
	MTRChannelClusterChangeChannelParamsClassOnce sync.Once
)

func getMTRChannelClusterChangeChannelParamsClass() _MTRChannelClusterChangeChannelParamsClass {
	MTRChannelClusterChangeChannelParamsClassOnce.Do(func() {
		MTRChannelClusterChangeChannelParamsClass = _MTRChannelClusterChangeChannelParamsClass{objc.GetClass("MTRChannelClusterChangeChannelParams")}
	})
	return MTRChannelClusterChangeChannelParamsClass
}

type _MTRChannelClusterChangeChannelParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChangeChannelParams] class.
type IMTRChannelClusterChangeChannelParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChangeChannelParams
type MTRChannelClusterChangeChannelParams struct {
	objectivec.Object
}

// MTRChannelClusterChangeChannelParamsFrom constructs a [MTRChannelClusterChangeChannelParams] from an unsafe.Pointer.
func MTRChannelClusterChangeChannelParamsFrom(ptr unsafe.Pointer) MTRChannelClusterChangeChannelParams {
	return MTRChannelClusterChangeChannelParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChangeChannelParamsClass) Alloc() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChangeChannelParamsClass) New() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChangeChannelParams) Init() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChangeChannelParams) Autorelease() MTRChannelClusterChangeChannelParams {
	rv := objc.Send[MTRChannelClusterChangeChannelParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChangeChannelParams creates a new MTRChannelClusterChangeChannelParams instance.
func NewMTRChannelClusterChangeChannelParams() MTRChannelClusterChangeChannelParams {
	return getMTRChannelClusterChangeChannelParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/match
func (m_ MTRChannelClusterChangeChannelParams) Match() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("match"))
	return rv
}


// SetMatch sets the value of the match property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/match
func (m_ MTRChannelClusterChangeChannelParams) SetMatch(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMatch:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/serversideprocessingtimeout
func (m_ MTRChannelClusterChangeChannelParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



