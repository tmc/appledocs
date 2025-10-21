// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RPScreenRecorder] class.
var (
	RPScreenRecorderClass     _RPScreenRecorderClass
	RPScreenRecorderClassOnce sync.Once
)

func getRPScreenRecorderClass() _RPScreenRecorderClass {
	RPScreenRecorderClassOnce.Do(func() {
		RPScreenRecorderClass = _RPScreenRecorderClass{objc.GetClass("RPScreenRecorder")}
	})
	return RPScreenRecorderClass
}

type _RPScreenRecorderClass struct {
	class objc.Class
}

// An interface definition for the [RPScreenRecorder] class.
type IRPScreenRecorder interface {
	objectivec.IObject
	DiscardRecordingWithHandler(handler unsafe.Pointer)
	ExportClipToURLDurationCompletionHandler(url unsafe.Pointer, duration TimeInterval, completionHandler func(error objc.ID))
	StartCaptureWithHandlerCompletionHandler(captureHandler unsafe.Pointer, completionHandler func(error objc.ID))
	StartClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer)
	StartRecordingWithHandler(handler func(error objc.ID))
	StartRecordingWithMicrophoneEnabledHandler(microphoneEnabled bool, handler func(error objc.ID))
	StopCaptureWithHandler(handler func(error objc.ID))
	StopClipBufferingWithCompletionHandler(completionHandler func(error objc.ID))
	StopRecordingWithHandler(handler unsafe.Pointer)
	StopRecordingWithOutputURLCompletionHandler(url unsafe.Pointer, completionHandler func(error objc.ID))
}

// The shared recorder object that provides the ability to record audio and video of your app.
//
// Apps on a user’s device can share the recording function, with each app having its own instance of . Your app can record the audio and video inside of the app, along with user commentary through the microphone. You get a reference to the recorder through the function and use it to implement start-and-stop recording functionality. You can present a user interface (view controller) where a user can trim and preview recordings, and share them with other users. Only one app at a time can use the recorder on the user’s device. Your app can’t record video from .
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder
type RPScreenRecorder struct {
	objectivec.Object
}

// RPScreenRecorderFrom constructs a [RPScreenRecorder] from an unsafe.Pointer.
//
// The shared recorder object that provides the ability to record audio and video of your app.
func RPScreenRecorderFrom(ptr unsafe.Pointer) RPScreenRecorder {
	return RPScreenRecorder{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RPScreenRecorderClass) Alloc() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPScreenRecorderClass) New() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPScreenRecorder) Init() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPScreenRecorder) Autorelease() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPScreenRecorder creates a new RPScreenRecorder instance.
func NewRPScreenRecorder() RPScreenRecorder {
	return getRPScreenRecorderClass().New()
}



// Returns an app’s instance of the shared screen recorder.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/shared()
func (rc _RPScreenRecorderClass) SharedRecorder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("sharedRecorder"))
	return rv
}

// Discards the current recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/discardRecording(handler:)
func (r_ RPScreenRecorder) DiscardRecordingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("discardRecordingWithHandler:"), handler)
}

// Exports a clip recording to a file.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/exportClip(to:duration:completionHandler:)
func (r_ RPScreenRecorder) ExportClipToURLDurationCompletionHandler(url unsafe.Pointer, duration TimeInterval, completionHandler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("exportClipToURL:duration:completionHandler:"), url, duration, completionHandler)
}

// Starts screen and audio capture.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startCapture(handler:completionHandler:)
func (r_ RPScreenRecorder) StartCaptureWithHandlerCompletionHandler(captureHandler unsafe.Pointer, completionHandler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startCaptureWithHandler:completionHandler:"), captureHandler, completionHandler)
}

// Starts buffering a clip recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startClipBuffering(completionHandler:)
func (r_ RPScreenRecorder) StartClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startClipBufferingWithCompletionHandler:"), completionHandler)
}

// Starts recording the app display.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startRecording(handler:)
func (r_ RPScreenRecorder) StartRecordingWithHandler(handler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startRecordingWithHandler:"), handler)
}

// Starts recording the app’s audio and video.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startRecording(withMicrophoneEnabled:handler:)
func (r_ RPScreenRecorder) StartRecordingWithMicrophoneEnabledHandler(microphoneEnabled bool, handler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startRecordingWithMicrophoneEnabled:handler:"), microphoneEnabled, handler)
}

// Stops screen capture
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopCapture(handler:)
func (r_ RPScreenRecorder) StopCaptureWithHandler(handler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopCaptureWithHandler:"), handler)
}

// Stops buffering a clip recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopClipBuffering(completionHandler:)
func (r_ RPScreenRecorder) StopClipBufferingWithCompletionHandler(completionHandler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopClipBufferingWithCompletionHandler:"), completionHandler)
}

// Stops the current recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopRecording(handler:)
func (r_ RPScreenRecorder) StopRecordingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopRecordingWithHandler:"), handler)
}

// Stops the current recording and writes the movie to the specified output URL.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopRecording(withOutput:completionHandler:)
func (r_ RPScreenRecorder) StopRecordingWithOutputURLCompletionHandler(url unsafe.Pointer, completionHandler func(error objc.ID)) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopRecordingWithOutputURL:completionHandler:"), url, completionHandler)
}

// The camera position to use.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPosition
func (r_ RPScreenRecorder) CameraPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("cameraPosition"))
	return rv
}


// SetCameraPosition sets the value of the cameraPosition property.
// The camera position to use.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPosition
func (r_ RPScreenRecorder) SetCameraPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCameraPosition:"), value)
}
// A view containing the contents of the front-facing camera.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPreviewView
func (r_ RPScreenRecorder) CameraPreviewView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("cameraPreviewView"))
	return rv
}

// The delegate for the screen recorder.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/delegate
func (r_ RPScreenRecorder) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the screen recorder.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/delegate
func (r_ RPScreenRecorder) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that indicates whether the screen recorder is available for recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isAvailable
func (r_ RPScreenRecorder) Available() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("available"))
	return rv
}

// A Boolean value that indicates whether the camera is currently enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isCameraEnabled
func (r_ RPScreenRecorder) CameraEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("cameraEnabled"))
	return rv
}


// SetCameraEnabled sets the value of the cameraEnabled property.
// A Boolean value that indicates whether the camera is currently enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isCameraEnabled
func (r_ RPScreenRecorder) SetCameraEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCameraEnabled:"), value)
}
// A Boolean value that indicates whether the microphone is currently enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isMicrophoneEnabled
func (r_ RPScreenRecorder) MicrophoneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("microphoneEnabled"))
	return rv
}


// SetMicrophoneEnabled sets the value of the microphoneEnabled property.
// A Boolean value that indicates whether the microphone is currently enabled.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isMicrophoneEnabled
func (r_ RPScreenRecorder) SetMicrophoneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMicrophoneEnabled:"), value)
}
// A Boolean value that indicates whether the app is currently recording.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isRecording
func (r_ RPScreenRecorder) Recording() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("recording"))
	return rv
}


