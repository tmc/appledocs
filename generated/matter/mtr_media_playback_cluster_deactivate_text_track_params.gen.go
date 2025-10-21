// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMediaPlaybackClusterDeactivateTextTrackParams] class.
var (
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClass     _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterDeactivateTextTrackParamsClass() _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass {
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterDeactivateTextTrackParamsClass = _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterDeactivateTextTrackParams")}
	})
	return MTRMediaPlaybackClusterDeactivateTextTrackParamsClass
}

type _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterDeactivateTextTrackParams] class.
type IMTRMediaPlaybackClusterDeactivateTextTrackParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams
type MTRMediaPlaybackClusterDeactivateTextTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterDeactivateTextTrackParamsFrom constructs a [MTRMediaPlaybackClusterDeactivateTextTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterDeactivateTextTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterDeactivateTextTrackParams {
	return MTRMediaPlaybackClusterDeactivateTextTrackParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass) Alloc() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass) New() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) Init() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) Autorelease() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterDeactivateTextTrackParams creates a new MTRMediaPlaybackClusterDeactivateTextTrackParams instance.
func NewMTRMediaPlaybackClusterDeactivateTextTrackParams() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	return getMTRMediaPlaybackClusterDeactivateTextTrackParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


