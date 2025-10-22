// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureSession] class.
var (
	CaptureSessionClass     _CaptureSessionClass
	CaptureSessionClassOnce sync.Once
)

func getCaptureSessionClass() _CaptureSessionClass {
	CaptureSessionClassOnce.Do(func() {
		CaptureSessionClass = _CaptureSessionClass{objc.GetClass("AVCaptureSession")}
	})
	return CaptureSessionClass
}

type _CaptureSessionClass struct {
	class objc.Class
}

// An interface definition for the [CaptureSession] class.
type ICaptureSession interface {
	objectivec.IObject
	AddConnection(connection IAVCaptureConnection)
	AddControl(control unsafe.Pointer)
	AddInput(input IAVCaptureInput)
	AddInputWithNoConnections(input IAVCaptureInput)
	AddOutput(output IAVCaptureOutput)
	AddOutputWithNoConnections(output IAVCaptureOutput)
	BeginConfiguration()
	CanAddConnection(connection IAVCaptureConnection) bool
	CanAddControl(control unsafe.Pointer) bool
	CanAddInput(input IAVCaptureInput) bool
	CanAddOutput(output IAVCaptureOutput) bool
	CanSetSessionPreset(preset ICaptureSessionPreset) bool
	CommitConfiguration()
	RemoveConnection(connection IAVCaptureConnection)
	RemoveControl(control unsafe.Pointer)
	RemoveInput(input IAVCaptureInput)
	RemoveOutput(output IAVCaptureOutput)
	RunDeferredStartWhenNeeded()
	SetControlsDelegateQueue(controlsDelegate objectivec.IObject, controlsDelegateCallbackQueue unsafe.Pointer)
	SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate objectivec.IObject, deferredStartDelegateCallbackQueue unsafe.Pointer)
	StartRunning()
	StopRunning()
	AutomaticallyConfiguresApplicationAudioSession() bool
	SetAutomaticallyConfiguresApplicationAudioSession(value bool)
	AutomaticallyConfiguresCaptureDeviceForWideColor() bool
	SetAutomaticallyConfiguresCaptureDeviceForWideColor(value bool)
	AutomaticallyRunsDeferredStart() bool
	SetAutomaticallyRunsDeferredStart(value bool)
	ConfiguresApplicationAudioSessionForBluetoothHighQualityRecording() bool
	SetConfiguresApplicationAudioSessionForBluetoothHighQualityRecording(value bool)
	ConfiguresApplicationAudioSessionToMixWithOthers() bool
	SetConfiguresApplicationAudioSessionToMixWithOthers(value bool)
	Connections() []CaptureConnection
	Controls() []unsafe.Pointer
	ControlsDelegate() objc.ID
	ControlsDelegateCallbackQueue() unsafe.Pointer
	DeferredStartDelegate() objc.ID
	DeferredStartDelegateCallbackQueue() unsafe.Pointer
	HardwareCost() float32
	Inputs() []CaptureInput
	Interrupted() bool
	ManualDeferredStartSupported() bool
	MultitaskingCameraAccessEnabled() bool
	SetMultitaskingCameraAccessEnabled(value bool)
	MultitaskingCameraAccessSupported() bool
	Running() bool
	MasterClock() unsafe.Pointer
	MaxControlsCount() int
	Outputs() []CaptureOutput
	SessionPreset() CaptureSessionPreset
	SetSessionPreset(value ICaptureSessionPreset)
	SupportsControls() bool
	SynchronizationClock() unsafe.Pointer
	UsesApplicationAudioSession() bool
	SetUsesApplicationAudioSession(value bool)
	IsInterrupted() bool
	SetIsInterrupted(value bool)
	IsManualDeferredStartSupported() bool
	SetIsManualDeferredStartSupported(value bool)
	IsMultitaskingCameraAccessEnabled() bool
	SetIsMultitaskingCameraAccessEnabled(value bool)
	IsMultitaskingCameraAccessSupported() bool
	SetIsMultitaskingCameraAccessSupported(value bool)
	IsRunning() bool
	SetIsRunning(value bool)
}

// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
//
// To perform real-time capture, you instantiate a capture session and add appropriate inputs and outputs. The following code fragment illustrates how to configure a capture device to record audio. Call the method to start the flow of data from the inputs to the outputs, and call the method to stop the flow. You use the property to customize the quality level, bitrate, or other settings for the output. Most common capture configurations are available through session presets; however, some specialized options (such as high frame rate) require directly setting a capture format on an instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession
type CaptureSession struct {
	objectivec.Object
}

// CaptureSessionFrom constructs a [CaptureSession] from an unsafe.Pointer.
//
// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
func CaptureSessionFrom(ptr unsafe.Pointer) CaptureSession {
	return CaptureSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureSessionClass) Alloc() CaptureSession {
	rv := objc.Send[CaptureSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureSessionClass) New() CaptureSession {
	rv := objc.Send[CaptureSession](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureSession) Init() CaptureSession {
	rv := objc.Send[CaptureSession](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureSession) Autorelease() CaptureSession {
	rv := objc.Send[CaptureSession](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureSession creates a new CaptureSession instance.
func NewCaptureSession() CaptureSession {
	return getCaptureSessionClass().New()
}


// Adds a connection to the capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addConnection(_:)
func (c_ CaptureSession) AddConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addConnection:"), connection)
}

// Adds a control to a capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addControl(_:)
func (c_ CaptureSession) AddControl(control unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addControl:"), control)
}

// Adds a capture input to the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInput(_:)
func (c_ CaptureSession) AddInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addInput:"), input)
}

// Adds a capture input to a session without forming any connections.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInputWithNoConnections(_:)
func (c_ CaptureSession) AddInputWithNoConnections(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addInputWithNoConnections:"), input)
}

// Adds an output to the capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutput(_:)
func (c_ CaptureSession) AddOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOutput:"), output)
}

// Adds a capture output to the session without forming any connections.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutputWithNoConnections(_:)
func (c_ CaptureSession) AddOutputWithNoConnections(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOutputWithNoConnections:"), output)
}

// Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/beginConfiguration()
func (c_ CaptureSession) BeginConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("beginConfiguration"))
}

// Determines whether a you can add a connection to a capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddConnection(_:)
func (c_ CaptureSession) CanAddConnection(connection IAVCaptureConnection) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddConnection:"), connection)
	return rv
}

// Returns a Boolean value that indicates whether a capture session add the specified control.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddControl(_:)
func (c_ CaptureSession) CanAddControl(control unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddControl:"), control)
	return rv
}

// Determines whether you can add an input to a session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddInput(_:)
func (c_ CaptureSession) CanAddInput(input IAVCaptureInput) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddInput:"), input)
	return rv
}

// Determines whether you can add an output to a session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddOutput(_:)
func (c_ CaptureSession) CanAddOutput(output IAVCaptureOutput) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}

// Determines whether you can configure a capture session with the specified preset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canSetSessionPreset(_:)
func (c_ CaptureSession) CanSetSessionPreset(preset ICaptureSessionPreset) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canSetSessionPreset:"), preset)
	return rv
}

// Commits one or more changes to a running capture session’s configuration in a single atomic update.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/commitConfiguration()
func (c_ CaptureSession) CommitConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("commitConfiguration"))
}

// Removes a capture connection from the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeConnection(_:)
func (c_ CaptureSession) RemoveConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeConnection:"), connection)
}

// Removes a control from a capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeControl(_:)
func (c_ CaptureSession) RemoveControl(control unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeControl:"), control)
}

// Removes an input from the session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeInput(_:)
func (c_ CaptureSession) RemoveInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeInput:"), input)
}

// Removes an output from a capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeOutput(_:)
func (c_ CaptureSession) RemoveOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeOutput:"), output)
}

// Tells the session to run deferred start when appropriate.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runDeferredStartWhenNeeded()
func (c_ CaptureSession) RunDeferredStartWhenNeeded() {
	objc.Send[objc.ID](c_.ID, objc.Sel("runDeferredStartWhenNeeded"))
}

// Sets a delegate object for the system to call when it activates and presents controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setControlsDelegate(_:queue:)
func (c_ CaptureSession) SetControlsDelegateQueue(controlsDelegate objectivec.IObject, controlsDelegateCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlsDelegate:queue:"), controlsDelegate, controlsDelegateCallbackQueue)
}

// Sets a delegate object for the session to call when performing deferred start.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setDeferredStartDelegate(_:deferredStartDelegateCallbackQueue:)
func (c_ CaptureSession) SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate objectivec.IObject, deferredStartDelegateCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeferredStartDelegate:deferredStartDelegateCallbackQueue:"), deferredStartDelegate, deferredStartDelegateCallbackQueue)
}

// Starts the flow of data through the capture pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/startRunning()
func (c_ CaptureSession) StartRunning() {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRunning"))
}

// Stops the flow of data through the capture pipeline.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/stopRunning()
func (c_ CaptureSession) StopRunning() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopRunning"))
}

// A Boolean value that indicates whether the capture session automatically changes settings in the app’s shared audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresApplicationAudioSession
func (c_ CaptureSession) AutomaticallyConfiguresApplicationAudioSession() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresApplicationAudioSession"))
	return rv
}


// SetAutomaticallyConfiguresApplicationAudioSession sets the value of the automaticallyConfiguresApplicationAudioSession property.
// A Boolean value that indicates whether the capture session automatically changes settings in the app’s shared audio session.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresApplicationAudioSession
func (c_ CaptureSession) SetAutomaticallyConfiguresApplicationAudioSession(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyConfiguresApplicationAudioSession:"), value)
}

// A Boolean value that specifies whether the session should automatically use wide-gamut color where available.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresCaptureDeviceForWideColor
func (c_ CaptureSession) AutomaticallyConfiguresCaptureDeviceForWideColor() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyConfiguresCaptureDeviceForWideColor"))
	return rv
}


// SetAutomaticallyConfiguresCaptureDeviceForWideColor sets the value of the automaticallyConfiguresCaptureDeviceForWideColor property.
// A Boolean value that specifies whether the session should automatically use wide-gamut color where available.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresCaptureDeviceForWideColor
func (c_ CaptureSession) SetAutomaticallyConfiguresCaptureDeviceForWideColor(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyConfiguresCaptureDeviceForWideColor:"), value)
}

// A Boolean value that indicates whether deferred start runs automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyRunsDeferredStart
func (c_ CaptureSession) AutomaticallyRunsDeferredStart() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyRunsDeferredStart"))
	return rv
}


// SetAutomaticallyRunsDeferredStart sets the value of the automaticallyRunsDeferredStart property.
// A Boolean value that indicates whether deferred start runs automatically.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyRunsDeferredStart
func (c_ CaptureSession) SetAutomaticallyRunsDeferredStart(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyRunsDeferredStart:"), value)
}

// A Boolean value that indicates whether the capture session configures the app’s audio session for bluetooth high-quality recording.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionForBluetoothHighQualityRecording
func (c_ CaptureSession) ConfiguresApplicationAudioSessionForBluetoothHighQualityRecording() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("configuresApplicationAudioSessionForBluetoothHighQualityRecording"))
	return rv
}


// SetConfiguresApplicationAudioSessionForBluetoothHighQualityRecording sets the value of the configuresApplicationAudioSessionForBluetoothHighQualityRecording property.
// A Boolean value that indicates whether the capture session configures the app’s audio session for bluetooth high-quality recording.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionForBluetoothHighQualityRecording
func (c_ CaptureSession) SetConfiguresApplicationAudioSessionForBluetoothHighQualityRecording(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguresApplicationAudioSessionForBluetoothHighQualityRecording:"), value)
}

// A Boolean value that Indicates whether the capture session configures the app’s audio session to mix with others.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionToMixWithOthers
func (c_ CaptureSession) ConfiguresApplicationAudioSessionToMixWithOthers() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("configuresApplicationAudioSessionToMixWithOthers"))
	return rv
}


// SetConfiguresApplicationAudioSessionToMixWithOthers sets the value of the configuresApplicationAudioSessionToMixWithOthers property.
// A Boolean value that Indicates whether the capture session configures the app’s audio session to mix with others.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionToMixWithOthers
func (c_ CaptureSession) SetConfiguresApplicationAudioSessionToMixWithOthers(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguresApplicationAudioSessionToMixWithOthers:"), value)
}

// The connections between inputs and outputs that a capture session contains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/connections
func (c_ CaptureSession) Connections() []CaptureConnection {
	rv := objc.Send[[]CaptureConnection](c_.ID, objc.Sel("connections"))
	return rv
}

// The controls that allow configuring the camera system from device hardware.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controls
func (c_ CaptureSession) Controls() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](c_.ID, objc.Sel("controls"))
	return rv
}

// A delegate object that observes changes to the state of capture controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegate
func (c_ CaptureSession) ControlsDelegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("controlsDelegate"))
	return rv
}

// The dispatch queue on which the system calls controls delegate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegateCallbackQueue
func (c_ CaptureSession) ControlsDelegateCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlsDelegateCallbackQueue"))
	return rv
}

// A delegate object that observes events about deferred start.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegate
func (c_ CaptureSession) DeferredStartDelegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("deferredStartDelegate"))
	return rv
}

// The dispatch queue on which the session calls deferred start delegate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegateCallbackQueue
func (c_ CaptureSession) DeferredStartDelegateCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deferredStartDelegateCallbackQueue"))
	return rv
}

// A value that indicates the percentage of the session’s available hardware budget in use.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/hardwareCost
func (c_ CaptureSession) HardwareCost() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("hardwareCost"))
	return rv
}

// The inputs that provide media data to a capture session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/inputs
func (c_ CaptureSession) Inputs() []CaptureInput {
	rv := objc.Send[[]CaptureInput](c_.ID, objc.Sel("inputs"))
	return rv
}

// A Boolean value that indicates whether the capture session is in an interrupted state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isInterrupted
func (c_ CaptureSession) Interrupted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("interrupted"))
	return rv
}

// A value that indicates whether the session supports manually running deferred start.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isManualDeferredStartSupported
func (c_ CaptureSession) ManualDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("manualDeferredStartSupported"))
	return rv
}

// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessEnabled
func (c_ CaptureSession) MultitaskingCameraAccessEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multitaskingCameraAccessEnabled"))
	return rv
}


// SetMultitaskingCameraAccessEnabled sets the value of the multitaskingCameraAccessEnabled property.
// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessEnabled
func (c_ CaptureSession) SetMultitaskingCameraAccessEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMultitaskingCameraAccessEnabled:"), value)
}

// A Boolean value that indicates whether the capture session supports using the camera while multitasking.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessSupported
func (c_ CaptureSession) MultitaskingCameraAccessSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("multitaskingCameraAccessSupported"))
	return rv
}

// A Boolean value that indicates whether the capture session is in a running state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isRunning
func (c_ CaptureSession) Running() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("running"))
	return rv
}

// A clock object used for output synchronization.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/masterClock
func (c_ CaptureSession) MasterClock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("masterClock"))
	return rv
}

// The maximum number of controls a capture session supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/maxControlsCount
func (c_ CaptureSession) MaxControlsCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxControlsCount"))
	return rv
}

// The output destinations to which a captures session sends its data.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/outputs
func (c_ CaptureSession) Outputs() []CaptureOutput {
	rv := objc.Send[[]CaptureOutput](c_.ID, objc.Sel("outputs"))
	return rv
}

// A preset value that indicates the quality level or bit rate of the output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/sessionPreset
func (c_ CaptureSession) SessionPreset() CaptureSessionPreset {
	rv := objc.Send[CaptureSessionPreset](c_.ID, objc.Sel("sessionPreset"))
	return rv
}


// SetSessionPreset sets the value of the sessionPreset property.
// A preset value that indicates the quality level or bit rate of the output.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/sessionPreset
func (c_ CaptureSession) SetSessionPreset(value ICaptureSessionPreset) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSessionPreset:"), value)
}

// A Boolean value that indicates whether a capture session supports controls.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/supportsControls
func (c_ CaptureSession) SupportsControls() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsControls"))
	return rv
}

// A clock to use for output synchronization.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/synchronizationClock
func (c_ CaptureSession) SynchronizationClock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("synchronizationClock"))
	return rv
}

// A Boolean value that indicates whether the capture session uses the app’s shared audio session.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/usesApplicationAudioSession
func (c_ CaptureSession) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}


// SetUsesApplicationAudioSession sets the value of the usesApplicationAudioSession property.
// A Boolean value that indicates whether the capture session uses the app’s shared audio session.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/usesApplicationAudioSession
func (c_ CaptureSession) SetUsesApplicationAudioSession(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUsesApplicationAudioSession:"), value)
}

// A Boolean value that indicates whether the capture session is in an interrupted state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isinterrupted
func (c_ CaptureSession) IsInterrupted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isInterrupted"))
	return rv
}


// SetIsInterrupted sets the value of the isInterrupted property.
// A Boolean value that indicates whether the capture session is in an interrupted state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isinterrupted
func (c_ CaptureSession) SetIsInterrupted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsInterrupted:"), value)
}

// A
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismanualdeferredstartsupported
func (c_ CaptureSession) IsManualDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isManualDeferredStartSupported"))
	return rv
}


// SetIsManualDeferredStartSupported sets the value of the isManualDeferredStartSupported property.
// A

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismanualdeferredstartsupported
func (c_ CaptureSession) SetIsManualDeferredStartSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsManualDeferredStartSupported:"), value)
}

// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccessenabled
func (c_ CaptureSession) IsMultitaskingCameraAccessEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultitaskingCameraAccessEnabled"))
	return rv
}


// SetIsMultitaskingCameraAccessEnabled sets the value of the isMultitaskingCameraAccessEnabled property.
// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccessenabled
func (c_ CaptureSession) SetIsMultitaskingCameraAccessEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultitaskingCameraAccessEnabled:"), value)
}

// A Boolean value that indicates whether the capture session supports using the camera while multitasking.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccesssupported
func (c_ CaptureSession) IsMultitaskingCameraAccessSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultitaskingCameraAccessSupported"))
	return rv
}


// SetIsMultitaskingCameraAccessSupported sets the value of the isMultitaskingCameraAccessSupported property.
// A Boolean value that indicates whether the capture session supports using the camera while multitasking.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccesssupported
func (c_ CaptureSession) SetIsMultitaskingCameraAccessSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultitaskingCameraAccessSupported:"), value)
}

// A Boolean value that indicates whether the capture session is in a running state.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isrunning
func (c_ CaptureSession) IsRunning() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRunning"))
	return rv
}


// SetIsRunning sets the value of the isRunning property.
// A Boolean value that indicates whether the capture session is in a running state.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isrunning
func (c_ CaptureSession) SetIsRunning(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRunning:"), value)
}



