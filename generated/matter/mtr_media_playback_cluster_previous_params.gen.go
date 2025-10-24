// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterPreviousParams] class.
var (
	MTRMediaPlaybackClusterPreviousParamsClass     _MTRMediaPlaybackClusterPreviousParamsClass
	MTRMediaPlaybackClusterPreviousParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterPreviousParamsClass() _MTRMediaPlaybackClusterPreviousParamsClass {
	MTRMediaPlaybackClusterPreviousParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterPreviousParamsClass = _MTRMediaPlaybackClusterPreviousParamsClass{objc.GetClass("MTRMediaPlaybackClusterPreviousParams")}
	})
	return MTRMediaPlaybackClusterPreviousParamsClass
}

type _MTRMediaPlaybackClusterPreviousParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterPreviousParams] class.
type IMTRMediaPlaybackClusterPreviousParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterPreviousParams
type MTRMediaPlaybackClusterPreviousParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterPreviousParamsFrom constructs a [MTRMediaPlaybackClusterPreviousParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterPreviousParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterPreviousParams {
	return MTRMediaPlaybackClusterPreviousParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterPreviousParamsClass) Alloc() MTRMediaPlaybackClusterPreviousParams {
	rv := objc.Send[MTRMediaPlaybackClusterPreviousParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterPreviousParamsClass) New() MTRMediaPlaybackClusterPreviousParams {
	rv := objc.Send[MTRMediaPlaybackClusterPreviousParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterPreviousParams) Init() MTRMediaPlaybackClusterPreviousParams {
	rv := objc.Send[MTRMediaPlaybackClusterPreviousParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterPreviousParams) Autorelease() MTRMediaPlaybackClusterPreviousParams {
	rv := objc.Send[MTRMediaPlaybackClusterPreviousParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterPreviousParams creates a new MTRMediaPlaybackClusterPreviousParams instance.
func NewMTRMediaPlaybackClusterPreviousParams() MTRMediaPlaybackClusterPreviousParams {
	return getMTRMediaPlaybackClusterPreviousParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpreviousparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPreviousParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpreviousparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterPreviousParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpreviousparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPreviousParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterpreviousparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterPreviousParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



