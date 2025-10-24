// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/data
func (m_ MTRChannelClusterChangeChannelResponseParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/data
func (m_ MTRChannelClusterChangeChannelResponseParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/status
func (m_ MTRChannelClusterChangeChannelResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/status
func (m_ MTRChannelClusterChangeChannelResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchangechannelresponseparams/timedinvoketimeoutms
func (m_ MTRChannelClusterChangeChannelResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
