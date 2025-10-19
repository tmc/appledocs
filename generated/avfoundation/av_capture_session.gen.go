// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureSession] class.
var aVCaptureSessionClass = _AVCaptureSessionClass{objc.GetClass("AVCaptureSession")}

type _AVCaptureSessionClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureSession] class.
type IAVCaptureSession interface {
	objectivec.IObject
	AddConnection(connection unsafe.Pointer)
	AddControl(control unsafe.Pointer)
	AddInput(input unsafe.Pointer)
	AddInputWithNoConnections(input unsafe.Pointer)
	AddOutput(output unsafe.Pointer)
	AddOutputWithNoConnections(output unsafe.Pointer)
	BeginConfiguration()
	CanAddConnection(connection unsafe.Pointer) bool
	CanAddControl(control unsafe.Pointer) bool
	CanAddInput(input unsafe.Pointer) bool
	CanAddOutput(output unsafe.Pointer) bool
	CanSetSessionPreset(preset unsafe.Pointer) bool
	CommitConfiguration()
	RemoveConnection(connection unsafe.Pointer)
	RemoveControl(control unsafe.Pointer)
	RemoveInput(input unsafe.Pointer)
	RemoveOutput(output unsafe.Pointer)
	RunDeferredStartWhenNeeded()
	SetControlsDelegateQueue(controlsDelegate unsafe.Pointer, controlsDelegateCallbackQueue unsafe.Pointer)
	SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate unsafe.Pointer, deferredStartDelegateCallbackQueue unsafe.Pointer)
	StartRunning()
	StopRunning()
}

// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession

type AVCaptureSession struct {
	objectivec.Object
}

// AVCaptureSessionFrom constructs a [AVCaptureSession] from an unsafe.Pointer.
//
// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
func AVCaptureSessionFrom(ptr unsafe.Pointer) AVCaptureSession {
	return AVCaptureSession{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureSessionClass) Alloc() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureSessionClass) New() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureSession) Init() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureSession) Autorelease() AVCaptureSession {
	rv := objc.Send[AVCaptureSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureSession creates a new AVCaptureSession instance.
func NewAVCaptureSession() AVCaptureSession {
	return aVCaptureSessionClass.New()
}


// Adds a connection to the capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addConnection(_:)
func (a_ AVCaptureSession) AddConnection(connection unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addConnection:"), connection)
}
// Adds a control to a capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addControl(_:)
func (a_ AVCaptureSession) AddControl(control unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addControl:"), control)
}
// Adds a capture input to the session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInput(_:)
func (a_ AVCaptureSession) AddInput(input unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addInput:"), input)
}
// Adds a capture input to a session without forming any connections. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInputWithNoConnections(_:)
func (a_ AVCaptureSession) AddInputWithNoConnections(input unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addInputWithNoConnections:"), input)
}
// Adds an output to the capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutput(_:)
func (a_ AVCaptureSession) AddOutput(output unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addOutput:"), output)
}
// Adds a capture output to the session without forming any connections. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutputWithNoConnections(_:)
func (a_ AVCaptureSession) AddOutputWithNoConnections(output unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addOutputWithNoConnections:"), output)
}
// Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/beginConfiguration()
func (a_ AVCaptureSession) BeginConfiguration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("beginConfiguration"))
}
// Determines whether a you can add a connection to a capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddConnection(_:)
func (a_ AVCaptureSession) CanAddConnection(connection unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddConnection:"), connection)
	return rv
}
// Returns a Boolean value that indicates whether a capture session add the specified control. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddControl(_:)
func (a_ AVCaptureSession) CanAddControl(control unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddControl:"), control)
	return rv
}
// Determines whether you can add an input to a session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddInput(_:)
func (a_ AVCaptureSession) CanAddInput(input unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInput:"), input)
	return rv
}
// Determines whether you can add an output to a session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddOutput(_:)
func (a_ AVCaptureSession) CanAddOutput(output unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}
// Determines whether you can configure a capture session with the specified preset. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canSetSessionPreset(_:)
func (a_ AVCaptureSession) CanSetSessionPreset(preset unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canSetSessionPreset:"), preset)
	return rv
}
// Commits one or more changes to a running capture session’s configuration in a single atomic update. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/commitConfiguration()
func (a_ AVCaptureSession) CommitConfiguration() {
	objc.Send[objc.ID](a_.ID, objc.Sel("commitConfiguration"))
}
// Removes a capture connection from the session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeConnection(_:)
func (a_ AVCaptureSession) RemoveConnection(connection unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeConnection:"), connection)
}
// Removes a control from a capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeControl(_:)
func (a_ AVCaptureSession) RemoveControl(control unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeControl:"), control)
}
// Removes an input from the session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeInput(_:)
func (a_ AVCaptureSession) RemoveInput(input unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeInput:"), input)
}
// Removes an output from a capture session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeOutput(_:)
func (a_ AVCaptureSession) RemoveOutput(output unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeOutput:"), output)
}
// Tells the session to run deferred start when appropriate. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runDeferredStartWhenNeeded()
func (a_ AVCaptureSession) RunDeferredStartWhenNeeded() {
	objc.Send[objc.ID](a_.ID, objc.Sel("runDeferredStartWhenNeeded"))
}
// Sets a delegate object for the system to call when it activates and presents controls. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setControlsDelegate(_:queue:)
func (a_ AVCaptureSession) SetControlsDelegateQueue(controlsDelegate unsafe.Pointer, controlsDelegateCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setControlsDelegate:queue:"), controlsDelegate, controlsDelegateCallbackQueue)
}
// Sets a delegate object for the session to call when performing deferred start. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setDeferredStartDelegate(_:deferredStartDelegateCallbackQueue:)
func (a_ AVCaptureSession) SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate unsafe.Pointer, deferredStartDelegateCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDeferredStartDelegate:deferredStartDelegateCallbackQueue:"), deferredStartDelegate, deferredStartDelegateCallbackQueue)
}
// Starts the flow of data through the capture pipeline. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/startRunning()
func (a_ AVCaptureSession) StartRunning() {
	objc.Send[objc.ID](a_.ID, objc.Sel("startRunning"))
}
// Stops the flow of data through the capture pipeline. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/stopRunning()
func (a_ AVCaptureSession) StopRunning() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopRunning"))
}


