// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaInputClusterHideInputStatusParams] class.
var (
	MTRMediaInputClusterHideInputStatusParamsClass     _MTRMediaInputClusterHideInputStatusParamsClass
	MTRMediaInputClusterHideInputStatusParamsClassOnce sync.Once
)

func getMTRMediaInputClusterHideInputStatusParamsClass() _MTRMediaInputClusterHideInputStatusParamsClass {
	MTRMediaInputClusterHideInputStatusParamsClassOnce.Do(func() {
		MTRMediaInputClusterHideInputStatusParamsClass = _MTRMediaInputClusterHideInputStatusParamsClass{objc.GetClass("MTRMediaInputClusterHideInputStatusParams")}
	})
	return MTRMediaInputClusterHideInputStatusParamsClass
}

type _MTRMediaInputClusterHideInputStatusParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaInputClusterHideInputStatusParams] class.
type IMTRMediaInputClusterHideInputStatusParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaInputClusterHideInputStatusParams
type MTRMediaInputClusterHideInputStatusParams struct {
	objectivec.Object
}

// MTRMediaInputClusterHideInputStatusParamsFrom constructs a [MTRMediaInputClusterHideInputStatusParams] from an unsafe.Pointer.
func MTRMediaInputClusterHideInputStatusParamsFrom(ptr unsafe.Pointer) MTRMediaInputClusterHideInputStatusParams {
	return MTRMediaInputClusterHideInputStatusParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaInputClusterHideInputStatusParamsClass) Alloc() MTRMediaInputClusterHideInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterHideInputStatusParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaInputClusterHideInputStatusParamsClass) New() MTRMediaInputClusterHideInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterHideInputStatusParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaInputClusterHideInputStatusParams) Init() MTRMediaInputClusterHideInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterHideInputStatusParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaInputClusterHideInputStatusParams) Autorelease() MTRMediaInputClusterHideInputStatusParams {
	rv := objc.Send[MTRMediaInputClusterHideInputStatusParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaInputClusterHideInputStatusParams creates a new MTRMediaInputClusterHideInputStatusParams instance.
func NewMTRMediaInputClusterHideInputStatusParams() MTRMediaInputClusterHideInputStatusParams {
	return getMTRMediaInputClusterHideInputStatusParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterhideinputstatusparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterHideInputStatusParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterhideinputstatusparams/serversideprocessingtimeout
func (m_ MTRMediaInputClusterHideInputStatusParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterhideinputstatusparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterHideInputStatusParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediainputclusterhideinputstatusparams/timedinvoketimeoutms
func (m_ MTRMediaInputClusterHideInputStatusParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
