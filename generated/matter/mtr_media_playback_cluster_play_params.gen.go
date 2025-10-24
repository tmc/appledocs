// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterPlayParams] class.
var (
	MTRMediaPlaybackClusterPlayParamsClass     _MTRMediaPlaybackClusterPlayParamsClass
	MTRMediaPlaybackClusterPlayParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPlayParamsClass() _MTRMediaPlaybackClusterPlayParamsClass {
	MTRMediaPlaybackClusterPlayParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterPlayParamsClass = _MTRMediaPlaybackClusterPlayParamsClass{objc.GetClass("MTRMediaPlaybackClusterPlayParams")}
	})
	return MTRMediaPlaybackClusterPlayParamsClass
}

type _MTRMediaPlaybackClusterPlayParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPlayParams] class.
type IMTRMediaPlaybackClusterPlayParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPlayParams
type MTRMediaPlaybackClusterPlayParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterPlayParamsFrom constructs a [MTRMediaPlaybackClusterPlayParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPlayParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPlayParams {
	return MTRMediaPlaybackClusterPlayParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPlayParamsClass) Alloc() MTRMediaPlaybackClusterPlayParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlayParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPlayParamsClass) New() MTRMediaPlaybackClusterPlayParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlayParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPlayParams) Init() MTRMediaPlaybackClusterPlayParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlayParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPlayParams) Autorelease() MTRMediaPlaybackClusterPlayParams {
	rv := objc.Send[MTRMediaPlaybackClusterPlayParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPlayParams creates a new MTRMediaPlaybackClusterPlayParams instance.
func NewMTRMediaPlaybackClusterPlayParams() MTRMediaPlaybackClusterPlayParams {
	return getMTRMediaPlaybackClusterPlayParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplayparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPlayParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplayparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPlayParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplayparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPlayParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterplayparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPlayParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
