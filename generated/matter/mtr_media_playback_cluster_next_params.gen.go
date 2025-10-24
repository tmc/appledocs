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
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterNextParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterNextParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterNextParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusternextparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterNextParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



