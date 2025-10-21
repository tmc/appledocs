// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
}

//
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TrackID() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("trackID"))
	return rv
}


// SetTrackID sets the value of the trackID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTrackID(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}



