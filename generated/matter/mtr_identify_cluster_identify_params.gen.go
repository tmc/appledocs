// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRIdentifyClusterIdentifyParams] class.
var (
	MTRIdentifyClusterIdentifyParamsClass     _MTRIdentifyClusterIdentifyParamsClass
	MTRIdentifyClusterIdentifyParamsClassOnce sync.Once
)

func getMTRIdentifyClusterIdentifyParamsClass() _MTRIdentifyClusterIdentifyParamsClass {
	MTRIdentifyClusterIdentifyParamsClassOnce.Do(func() {
		MTRIdentifyClusterIdentifyParamsClass = _MTRIdentifyClusterIdentifyParamsClass{objc.GetClass("MTRIdentifyClusterIdentifyParams")}
	})
	return MTRIdentifyClusterIdentifyParamsClass
}

type _MTRIdentifyClusterIdentifyParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRIdentifyClusterIdentifyParams] class.
type IMTRIdentifyClusterIdentifyParams interface {
	objectivec.IObject
	// properties:
	IdentifyTime() objc.IObject /* cross-framework: NSNumber */
	SetIdentifyTime(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRIdentifyClusterIdentifyParams
type MTRIdentifyClusterIdentifyParams struct {
	objectivec.Object
}

// MTRIdentifyClusterIdentifyParamsFrom constructs a [MTRIdentifyClusterIdentifyParams] from an unsafe.Pointer.
func MTRIdentifyClusterIdentifyParamsFrom(ptr unsafe.Pointer) MTRIdentifyClusterIdentifyParams {
	return MTRIdentifyClusterIdentifyParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRIdentifyClusterIdentifyParamsClass) Alloc() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRIdentifyClusterIdentifyParamsClass) New() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRIdentifyClusterIdentifyParams) Init() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRIdentifyClusterIdentifyParams) Autorelease() MTRIdentifyClusterIdentifyParams {
	rv := objc.Send[MTRIdentifyClusterIdentifyParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRIdentifyClusterIdentifyParams creates a new MTRIdentifyClusterIdentifyParams instance.
func NewMTRIdentifyClusterIdentifyParams() MTRIdentifyClusterIdentifyParams {
	return getMTRIdentifyClusterIdentifyParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/identifytime
func (m_ MTRIdentifyClusterIdentifyParams) IdentifyTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("identifyTime"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/identifytime
func (m_ MTRIdentifyClusterIdentifyParams) SetIdentifyTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifyTime:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/serversideprocessingtimeout
func (m_ MTRIdentifyClusterIdentifyParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/serversideprocessingtimeout
func (m_ MTRIdentifyClusterIdentifyParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/timedinvoketimeoutms
func (m_ MTRIdentifyClusterIdentifyParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtridentifyclusteridentifyparams/timedinvoketimeoutms
func (m_ MTRIdentifyClusterIdentifyParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
