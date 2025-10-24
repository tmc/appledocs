// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterActivateTextTrackParams] class.
var (
	MTRMediaPlaybackClusterActivateTextTrackParamsClass     _MTRMediaPlaybackClusterActivateTextTrackParamsClass
	MTRMediaPlaybackClusterActivateTextTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterActivateTextTrackParamsClass() _MTRMediaPlaybackClusterActivateTextTrackParamsClass {
	MTRMediaPlaybackClusterActivateTextTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterActivateTextTrackParamsClass = _MTRMediaPlaybackClusterActivateTextTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterActivateTextTrackParams")}
	})
	return MTRMediaPlaybackClusterActivateTextTrackParamsClass
}

type _MTRMediaPlaybackClusterActivateTextTrackParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterActivateTextTrackParams] class.
type IMTRMediaPlaybackClusterActivateTextTrackParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TrackID() objc.IObject /* cross-framework: NSString */
	SetTrackID(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams
type MTRMediaPlaybackClusterActivateTextTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterActivateTextTrackParamsFrom constructs a [MTRMediaPlaybackClusterActivateTextTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterActivateTextTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterActivateTextTrackParams {
	return MTRMediaPlaybackClusterActivateTextTrackParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterActivateTextTrackParamsClass) Alloc() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterActivateTextTrackParamsClass) New() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) Init() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) Autorelease() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterActivateTextTrackParams creates a new MTRMediaPlaybackClusterActivateTextTrackParams instance.
func NewMTRMediaPlaybackClusterActivateTextTrackParams() MTRMediaPlaybackClusterActivateTextTrackParams {
	return getMTRMediaPlaybackClusterActivateTextTrackParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TrackID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("trackID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTrackID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}



