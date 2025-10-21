// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RPBroadcastSampleHandler] class.
var (
	RPBroadcastSampleHandlerClass     _RPBroadcastSampleHandlerClass
	RPBroadcastSampleHandlerClassOnce sync.Once
)

func getRPBroadcastSampleHandlerClass() _RPBroadcastSampleHandlerClass {
	RPBroadcastSampleHandlerClassOnce.Do(func() {
		RPBroadcastSampleHandlerClass = _RPBroadcastSampleHandlerClass{objc.GetClass("RPBroadcastSampleHandler")}
	})
	return RPBroadcastSampleHandlerClass
}

type _RPBroadcastSampleHandlerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastSampleHandler] class.
type IRPBroadcastSampleHandler interface {
	IRPBroadcastHandler
	BroadcastAnnotatedWithApplicationInfo(applicationInfo objectivec.IObject)
	BroadcastFinished()
	BroadcastPaused()
	BroadcastResumed()
	BroadcastStartedWithSetupInfo(setupInfo unsafe.Pointer)
	FinishBroadcastWithError(error_ foundation.IError)
	ProcessSampleBufferWithType(sampleBuffer unsafe.Pointer, sampleBufferType RPSampleBufferType)
}

// An object that processes buffer objects as received from ReplayKit.
//
// To handle objects as captured by ReplayKit, you subclass . You enable this mode of handling by setting in the extension’s file to . In your subclass, implement the method to handle video and audio buffers, as well as the , , , and methods to handle starting and stopping the broadcast. ReplayKit invokes methods in your subclass in a serial fashion. After invoking one method, ReplayKit won’t invoke another method until the first method returns. That means it’s safe for your implementations to update their stored state without the use of locks or synchronization to provide thread safety.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler
type RPBroadcastSampleHandler struct {
	RPBroadcastHandler
}

// RPBroadcastSampleHandlerFrom constructs a [RPBroadcastSampleHandler] from an unsafe.Pointer.
//
// An object that processes buffer objects as received from ReplayKit.
func RPBroadcastSampleHandlerFrom(ptr unsafe.Pointer) RPBroadcastSampleHandler {
	return RPBroadcastSampleHandler{
		RPBroadcastHandler: RPBroadcastHandlerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastSampleHandlerClass) Alloc() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastSampleHandlerClass) New() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastSampleHandler) Init() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastSampleHandler) Autorelease() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastSampleHandler creates a new RPBroadcastSampleHandler instance.
func NewRPBroadcastSampleHandler() RPBroadcastSampleHandler {
	return getRPBroadcastSampleHandlerClass().New()
}


// Perform any required actions after starting a live broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastAnnotated(withApplicationInfo:)
func (r_ RPBroadcastSampleHandler) BroadcastAnnotatedWithApplicationInfo(applicationInfo objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastAnnotatedWithApplicationInfo:"), applicationInfo)
}

// Perform any required actions after a live broadcast is finished.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastFinished()
func (r_ RPBroadcastSampleHandler) BroadcastFinished() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastFinished"))
}

// Perform any required actions after a live broadcast is paused.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastPaused()
func (r_ RPBroadcastSampleHandler) BroadcastPaused() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastPaused"))
}

// Perform any required actions after a live broadcast is resumed.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastResumed()
func (r_ RPBroadcastSampleHandler) BroadcastResumed() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastResumed"))
}

// Perform any required actions after starting a live broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastStarted(withSetupInfo:)
func (r_ RPBroadcastSampleHandler) BroadcastStartedWithSetupInfo(setupInfo unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastStartedWithSetupInfo:"), setupInfo)
}

// Stops the broadcast and passes an error back to the broadcasting app.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/finishBroadcastWithError(_:)
func (r_ RPBroadcastSampleHandler) FinishBroadcastWithError(error_ foundation.IError) {
	objc.Send[objc.ID](r_.ID, objc.Sel("finishBroadcastWithError:"), error_)
}

// Processes video and audio data as it becomes available during a live broadcast.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/processSampleBuffer(_:with:)
func (r_ RPBroadcastSampleHandler) ProcessSampleBufferWithType(sampleBuffer unsafe.Pointer, sampleBufferType RPSampleBufferType) {
	objc.Send[objc.ID](r_.ID, objc.Sel("processSampleBuffer:withType:"), sampleBuffer, sampleBufferType)
}

// The key to retrieve the app’s bundle identifier from the user-information dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpapplicationinfobundleidentifierkey
func (r_ RPBroadcastSampleHandler) RPApplicationInfoBundleIdentifierKey() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("RPApplicationInfoBundleIdentifierKey"))
	return rv
}

// The sample attachment key that describes the video orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpvideosampleorientationkey
func (r_ RPBroadcastSampleHandler) RPVideoSampleOrientationKey() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("RPVideoSampleOrientationKey"))
	return rv
}



