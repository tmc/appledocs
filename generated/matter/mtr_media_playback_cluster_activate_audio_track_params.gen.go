// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	AudioOutputIndex() objc.IObject /* cross-framework: NSNumber */
	SetAudioOutputIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TrackID() objc.IObject /* cross-framework: NSString */
	SetTrackID(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/audioOutputIndex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) AudioOutputIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("audioOutputIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/audioOutputIndex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetAudioOutputIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioOutputIndex:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TrackID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("trackID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/trackID
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTrackID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}



