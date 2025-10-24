// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterPauseParams] class.
var (
	MTRMediaPlaybackClusterPauseParamsClass     _MTRMediaPlaybackClusterPauseParamsClass
	MTRMediaPlaybackClusterPauseParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPauseParamsClass() _MTRMediaPlaybackClusterPauseParamsClass {
	MTRMediaPlaybackClusterPauseParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterPauseParamsClass = _MTRMediaPlaybackClusterPauseParamsClass{objc.GetClass("MTRMediaPlaybackClusterPauseParams")}
	})
	return MTRMediaPlaybackClusterPauseParamsClass
}

type _MTRMediaPlaybackClusterPauseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPauseParams] class.
type IMTRMediaPlaybackClusterPauseParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPauseParams
type MTRMediaPlaybackClusterPauseParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterPauseParamsFrom constructs a [MTRMediaPlaybackClusterPauseParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPauseParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPauseParams {
	return MTRMediaPlaybackClusterPauseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPauseParamsClass) Alloc() MTRMediaPlaybackClusterPauseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPauseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPauseParamsClass) New() MTRMediaPlaybackClusterPauseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPauseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPauseParams) Init() MTRMediaPlaybackClusterPauseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPauseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPauseParams) Autorelease() MTRMediaPlaybackClusterPauseParams {
	rv := objc.Send[MTRMediaPlaybackClusterPauseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPauseParams creates a new MTRMediaPlaybackClusterPauseParams instance.
func NewMTRMediaPlaybackClusterPauseParams() MTRMediaPlaybackClusterPauseParams {
	return getMTRMediaPlaybackClusterPauseParamsClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpauseparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPauseParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpauseparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPauseParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpauseparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPauseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpauseparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPauseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
