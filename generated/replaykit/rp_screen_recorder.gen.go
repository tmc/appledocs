// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class RPScreenRecorder */


/* debug [class_header]: Header for RPScreenRecorder */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPScreenRecorder */
// An interface definition for the [RPScreenRecorder] class.
type IRPScreenRecorder interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RPScreenRecorder */
	// properties:
	CameraPosition() RPCameraPosition
	SetCameraPosition(value RPCameraPosition)
	CameraPreviewView() objc.IObject /* cross-framework: View */
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Available() bool
	CameraEnabled() bool
	SetCameraEnabled(value bool)
	MicrophoneEnabled() bool
	SetMicrophoneEnabled(value bool)
	Recording() bool
	IsAvailable() bool
	SetIsAvailable(value bool)
	IsCameraEnabled() bool
	SetIsCameraEnabled(value bool)
	IsMicrophoneEnabled() bool
	SetIsMicrophoneEnabled(value bool)
	IsRecording() bool
	SetIsRecording(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPScreenRecorder */
	// methods:
	DiscardRecordingWithHandler(handler unsafe.Pointer)
	ExportClipToURLDurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, duration float64, completionHandler unsafe.Pointer)
	StartCaptureWithHandlerCompletionHandler(captureHandler unsafe.Pointer, completionHandler unsafe.Pointer)
	StartClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer)
	StartRecordingWithHandler(handler unsafe.Pointer)
	StopCaptureWithHandler(handler unsafe.Pointer)
	StopClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer)
	StopRecordingWithHandler(handler unsafe.Pointer)
	StopRecordingWithOutputURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPScreenRecorder */
// Alloc allocates a new instance without initialization.
func (rc _RPScreenRecorderClass) Alloc() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPScreenRecorder */
// The shared recorder object that provides the ability to record audio and video of your app.
//
// Apps on a user’s device can share the recording function, with each app having its own instance of . Your app can record the audio and video inside of the app, along with user commentary through the microphone. You get a reference to the recorder through the function and use it to implement start-and-stop recording functionality. You can present a user interface (view controller) where a user can trim and preview recordings, and share them with other users. Only one app at a time can use the recorder on the user’s device. Your app can’t record video from .


// The shared recorder object that provides the ability to record audio and video of your app.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPScreenRecorder */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPScreenRecorder */

// Returns an app’s instance of the shared screen recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/shared()
func (rc _RPScreenRecorderClass) SharedRecorder() RPScreenRecorder {
	rv := objc.Send[RPScreenRecorder](objc.ID(rc.class), objc.Sel("sharedRecorder"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedRecorder) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPScreenRecorder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPScreenRecorder */

// Discards the current recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/discardRecording(handler:)
func (r_ RPScreenRecorder) DiscardRecordingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("discardRecordingWithHandler:"), handler)
}/* debug [instance_methods/method]: DiscardRecordingWithHandler */


// Exports a clip recording to a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/exportClip(to:duration:completionHandler:)
func (r_ RPScreenRecorder) ExportClipToURLDurationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, duration float64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("exportClipToURL:duration:completionHandler:"), url, duration, completionHandler)
}/* debug [instance_methods/method]: ExportClipToURLDurationCompletionHandler */


// Starts screen and audio capture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startCapture(handler:completionHandler:)
func (r_ RPScreenRecorder) StartCaptureWithHandlerCompletionHandler(captureHandler unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startCaptureWithHandler:completionHandler:"), captureHandler, completionHandler)
}/* debug [instance_methods/method]: StartCaptureWithHandlerCompletionHandler */


// Starts buffering a clip recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startClipBuffering(completionHandler:)
func (r_ RPScreenRecorder) StartClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startClipBufferingWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StartClipBufferingWithCompletionHandler */


// Starts recording the app display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/startRecording(handler:)
func (r_ RPScreenRecorder) StartRecordingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("startRecordingWithHandler:"), handler)
}/* debug [instance_methods/method]: StartRecordingWithHandler */


// Stops screen capture
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopCapture(handler:)
func (r_ RPScreenRecorder) StopCaptureWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopCaptureWithHandler:"), handler)
}/* debug [instance_methods/method]: StopCaptureWithHandler */


// Stops buffering a clip recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopClipBuffering(completionHandler:)
func (r_ RPScreenRecorder) StopClipBufferingWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopClipBufferingWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: StopClipBufferingWithCompletionHandler */


// Stops the current recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopRecording(handler:)
func (r_ RPScreenRecorder) StopRecordingWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopRecordingWithHandler:"), handler)
}/* debug [instance_methods/method]: StopRecordingWithHandler */


// Stops the current recording and writes the movie to the specified output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/stopRecording(withOutput:completionHandler:)
func (r_ RPScreenRecorder) StopRecordingWithOutputURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("stopRecordingWithOutputURL:completionHandler:"), url, completionHandler)
}/* debug [instance_methods/method]: StopRecordingWithOutputURLCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPScreenRecorder */

// The camera position to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPosition
func (r_ RPScreenRecorder) CameraPosition() RPCameraPosition {
	rv := objc.Send[RPCameraPosition](r_.ID, objc.Sel("cameraPosition"))
	return rv
}/* debug [instance_properties/getter]: cameraPosition */


// The camera position to use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPosition
func (r_ RPScreenRecorder) SetCameraPosition(value RPCameraPosition) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCameraPosition:"), value)
}/* debug [instance_properties/setter]: cameraPosition */


// A view containing the contents of the front-facing camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/cameraPreviewView
func (r_ RPScreenRecorder) CameraPreviewView() objc.IObject /* cross-framework: View */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("cameraPreviewView"))
	return rv
}/* debug [instance_properties/getter]: cameraPreviewView */


// The delegate for the screen recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/delegate
func (r_ RPScreenRecorder) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the screen recorder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/delegate
func (r_ RPScreenRecorder) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean value that indicates whether the screen recorder is available for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isAvailable
func (r_ RPScreenRecorder) Available() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("available"))
	return rv
}/* debug [instance_properties/getter]: available */


// A Boolean value that indicates whether the camera is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isCameraEnabled
func (r_ RPScreenRecorder) CameraEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("cameraEnabled"))
	return rv
}/* debug [instance_properties/getter]: cameraEnabled */


// A Boolean value that indicates whether the camera is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isCameraEnabled
func (r_ RPScreenRecorder) SetCameraEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCameraEnabled:"), value)
}/* debug [instance_properties/setter]: cameraEnabled */


// A Boolean value that indicates whether the microphone is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isMicrophoneEnabled
func (r_ RPScreenRecorder) MicrophoneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("microphoneEnabled"))
	return rv
}/* debug [instance_properties/getter]: microphoneEnabled */


// A Boolean value that indicates whether the microphone is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isMicrophoneEnabled
func (r_ RPScreenRecorder) SetMicrophoneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMicrophoneEnabled:"), value)
}/* debug [instance_properties/setter]: microphoneEnabled */


// A Boolean value that indicates whether the app is currently recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPScreenRecorder/isRecording
func (r_ RPScreenRecorder) Recording() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("recording"))
	return rv
}/* debug [instance_properties/getter]: recording */


// A Boolean value that indicates whether the screen recorder is available for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/isavailable
func (r_ RPScreenRecorder) IsAvailable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isAvailable"))
	return rv
}/* debug [instance_properties/getter]: isAvailable */


// A Boolean value that indicates whether the screen recorder is available for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/isavailable
func (r_ RPScreenRecorder) SetIsAvailable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsAvailable:"), value)
}/* debug [instance_properties/setter]: isAvailable */


// A Boolean value that indicates whether the camera is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/iscameraenabled
func (r_ RPScreenRecorder) IsCameraEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isCameraEnabled"))
	return rv
}/* debug [instance_properties/getter]: isCameraEnabled */


// A Boolean value that indicates whether the camera is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/iscameraenabled
func (r_ RPScreenRecorder) SetIsCameraEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsCameraEnabled:"), value)
}/* debug [instance_properties/setter]: isCameraEnabled */


// A Boolean value that indicates whether the microphone is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/ismicrophoneenabled
func (r_ RPScreenRecorder) IsMicrophoneEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isMicrophoneEnabled"))
	return rv
}/* debug [instance_properties/getter]: isMicrophoneEnabled */


// A Boolean value that indicates whether the microphone is currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/ismicrophoneenabled
func (r_ RPScreenRecorder) SetIsMicrophoneEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsMicrophoneEnabled:"), value)
}/* debug [instance_properties/setter]: isMicrophoneEnabled */


// A Boolean value that indicates whether the app is currently recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/isrecording
func (r_ RPScreenRecorder) IsRecording() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRecording"))
	return rv
}/* debug [instance_properties/getter]: isRecording */


// A Boolean value that indicates whether the app is currently recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpscreenrecorder/isrecording
func (r_ RPScreenRecorder) SetIsRecording(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRecording:"), value)
}/* debug [instance_properties/setter]: isRecording */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPScreenRecorder */


