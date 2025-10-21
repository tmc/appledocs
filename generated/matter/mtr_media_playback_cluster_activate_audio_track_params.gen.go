// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMediaPlaybackClusterActivateAudioTrackParams] class.
var (
	MTRMediaPlaybackClusterActivateAudioTrackParamsClass     _MTRMediaPlaybackClusterActivateAudioTrackParamsClass
	MTRMediaPlaybackClusterActivateAudioTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterActivateAudioTrackParamsClass() _MTRMediaPlaybackClusterActivateAudioTrackParamsClass {
	MTRMediaPlaybackClusterActivateAudioTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterActivateAudioTrackParamsClass = _MTRMediaPlaybackClusterActivateAudioTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterActivateAudioTrackParams")}
	})
	return MTRMediaPlaybackClusterActivateAudioTrackParamsClass
}

type _MTRMediaPlaybackClusterActivateAudioTrackParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMediaPlaybackClusterActivateAudioTrackParams] class.
type IMTRMediaPlaybackClusterActivateAudioTrackParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams
type MTRMediaPlaybackClusterActivateAudioTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterActivateAudioTrackParamsFrom constructs a [MTRMediaPlaybackClusterActivateAudioTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterActivateAudioTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterActivateAudioTrackParams {
	return MTRMediaPlaybackClusterActivateAudioTrackParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterActivateAudioTrackParamsClass) Alloc() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMediaPlaybackClusterActivateAudioTrackParamsClass) New() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) Init() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) Autorelease() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterActivateAudioTrackParams creates a new MTRMediaPlaybackClusterActivateAudioTrackParams instance.
func NewMTRMediaPlaybackClusterActivateAudioTrackParams() MTRMediaPlaybackClusterActivateAudioTrackParams {
	return getMTRMediaPlaybackClusterActivateAudioTrackParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/audioOutputIndex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) AudioOutputIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("audioOutputIndex"))
	return rv
}


// SetAudioOutputIndex sets the value of the audioOutputIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/audioOutputIndex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetAudioOutputIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioOutputIndex:"), value)
}
// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TrackID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("trackID"))
	return rv
}


// SetTrackID sets the value of the trackID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTrackID(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), objc.String(value))
}


