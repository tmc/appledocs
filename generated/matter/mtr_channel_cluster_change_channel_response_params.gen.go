// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChangeChannelResponseParams] class.
var (
	MTRChannelClusterChangeChannelResponseParamsClass     _MTRChannelClusterChangeChannelResponseParamsClass
	MTRChannelClusterChangeChannelResponseParamsClassOnce sync.Once
)

func getMTRChannelClusterChangeChannelResponseParamsClass() _MTRChannelClusterChangeChannelResponseParamsClass {
	MTRChannelClusterChangeChannelResponseParamsClassOnce.Do(func() {
		MTRChannelClusterChangeChannelResponseParamsClass = _MTRChannelClusterChangeChannelResponseParamsClass{objc.GetClass("MTRChannelClusterChangeChannelResponseParams")}
	})
	return MTRChannelClusterChangeChannelResponseParamsClass
}

type _MTRChannelClusterChangeChannelResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChangeChannelResponseParams] class.
type IMTRChannelClusterChangeChannelResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChangeChannelResponseParams
type MTRChannelClusterChangeChannelResponseParams struct {
	objectivec.Object
}

// MTRChannelClusterChangeChannelResponseParamsFrom constructs a [MTRChannelClusterChangeChannelResponseParams] from an unsafe.Pointer.
func MTRChannelClusterChangeChannelResponseParamsFrom(ptr unsafe.Pointer) MTRChannelClusterChangeChannelResponseParams {
	return MTRChannelClusterChangeChannelResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChangeChannelResponseParamsClass) Alloc() MTRChannelClusterChangeChannelResponseParams {
	rv := objc.Send[MTRChannelClusterChangeChannelResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChangeChannelResponseParamsClass) New() MTRChannelClusterChangeChannelResponseParams {
	rv := objc.Send[MTRChannelClusterChangeChannelResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChangeChannelResponseParams) Init() MTRChannelClusterChangeChannelResponseParams {
	rv := objc.Send[MTRChannelClusterChangeChannelResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChangeChannelResponseParams) Autorelease() MTRChannelClusterChangeChannelResponseParams {
	rv := objc.Send[MTRChannelClusterChangeChannelResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChangeChannelResponseParams creates a new MTRChannelClusterChangeChannelResponseParams instance.
func NewMTRChannelClusterChangeChannelResponseParams() MTRChannelClusterChangeChannelResponseParams {
	return getMTRChannelClusterChangeChannelResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/data
func (m_ MTRChannelClusterChangeChannelResponseParams) Data() string {
	rv := objc.Send[string](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/data
func (m_ MTRChannelClusterChangeChannelResponseParams) SetData(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/status
func (m_ MTRChannelClusterChangeChannelResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/status
func (m_ MTRChannelClusterChangeChannelResponseParams) SetStatus(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



