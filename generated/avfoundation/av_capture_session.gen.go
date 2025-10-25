// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureSession */


/* debug [class_header]: Header for AVCaptureSession */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureSession */
// An interface definition for the [CaptureSession] class.
type ICaptureSession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureSession */
	// properties:
	AutomaticallyRunsDeferredStart() bool
	SetAutomaticallyRunsDeferredStart(value bool)
	Connections() []CaptureConnection
	Controls() []CaptureControl
	ControlsDelegate() unsafe.Pointer
	ControlsDelegateCallbackQueue() objectivec.IObject
	DeferredStartDelegate() unsafe.Pointer
	DeferredStartDelegateCallbackQueue() objectivec.IObject
	Inputs() []CaptureInput
	ManualDeferredStartSupported() bool
	Running() bool
	MasterClock() ClockRef /* not a class type */
	MaxControlsCount() int
	Outputs() []CaptureOutput
	SessionPreset() CaptureSessionPreset /* typedef */
	SetSessionPreset(value CaptureSessionPreset /* typedef */)
	SupportsControls() bool
	SynchronizationClock() ClockRef /* not a class type */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureSession */
	// methods:
	AddConnection(connection IAVCaptureConnection)
	AddControl(control IAVCaptureControl)
	AddInput(input IAVCaptureInput)
	AddInputWithNoConnections(input IAVCaptureInput)
	AddOutput(output IAVCaptureOutput)
	AddOutputWithNoConnections(output IAVCaptureOutput)
	BeginConfiguration()
	CanAddConnection(connection IAVCaptureConnection) bool
	CanAddControl(control IAVCaptureControl) bool
	CanAddInput(input IAVCaptureInput) bool
	CanAddOutput(output IAVCaptureOutput) bool
	CanSetSessionPreset(preset CaptureSessionPreset /* typedef */) bool
	CommitConfiguration()
	RemoveConnection(connection IAVCaptureConnection)
	RemoveControl(control IAVCaptureControl)
	RemoveInput(input IAVCaptureInput)
	RemoveOutput(output IAVCaptureOutput)
	RunDeferredStartWhenNeeded()
	SetControlsDelegateQueue(controlsDelegate unsafe.Pointer, controlsDelegateCallbackQueue objectivec.IObject)
	SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate unsafe.Pointer, deferredStartDelegateCallbackQueue objectivec.IObject)
	StartRunning()
	StopRunning()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureSession */
// Alloc allocates a new instance without initialization.
func (cc _CaptureSessionClass) Alloc() CaptureSession {
	rv := objc.Send[CaptureSession](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureSession */
// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
//
// To perform real-time capture, you instantiate a capture session and add appropriate inputs and outputs. The following code fragment illustrates how to configure a capture device to record audio. Call the method to start the flow of data from the inputs to the outputs, and call the method to stop the flow. You use the property to customize the quality level, bitrate, or other settings for the output. Most common capture configurations are available through session presets; however, some specialized options (such as high frame rate) require directly setting a capture format on an instance.


// An object that configures capture behavior and coordinates the flow of data from input devices to capture outputs.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureSession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureSession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureSession */

// Adds a connection to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addConnection(_:)
func (c_ CaptureSession) AddConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addConnection:"), connection)
}/* debug [instance_methods/method]: AddConnection */


// Adds a control to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addControl(_:)
func (c_ CaptureSession) AddControl(control IAVCaptureControl) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addControl:"), control)
}/* debug [instance_methods/method]: AddControl */


// Adds a capture input to the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInput(_:)
func (c_ CaptureSession) AddInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addInput:"), input)
}/* debug [instance_methods/method]: AddInput */


// Adds a capture input to a session without forming any connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addInputWithNoConnections(_:)
func (c_ CaptureSession) AddInputWithNoConnections(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addInputWithNoConnections:"), input)
}/* debug [instance_methods/method]: AddInputWithNoConnections */


// Adds an output to the capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutput(_:)
func (c_ CaptureSession) AddOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOutput:"), output)
}/* debug [instance_methods/method]: AddOutput */


// Adds a capture output to the session without forming any connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/addOutputWithNoConnections(_:)
func (c_ CaptureSession) AddOutputWithNoConnections(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addOutputWithNoConnections:"), output)
}/* debug [instance_methods/method]: AddOutputWithNoConnections */


// Marks the beginning of changes to a running capture session’s configuration to perform in a single atomic update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/beginConfiguration()
func (c_ CaptureSession) BeginConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("beginConfiguration"))
}/* debug [instance_methods/method]: BeginConfiguration */


// Determines whether a you can add a connection to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddConnection(_:)
func (c_ CaptureSession) CanAddConnection(connection IAVCaptureConnection) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddConnection:"), connection)
	return rv
}/* debug [instance_methods/method]: CanAddConnection */


// Returns a Boolean value that indicates whether a capture session add the specified control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddControl(_:)
func (c_ CaptureSession) CanAddControl(control IAVCaptureControl) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddControl:"), control)
	return rv
}/* debug [instance_methods/method]: CanAddControl */


// Determines whether you can add an input to a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddInput(_:)
func (c_ CaptureSession) CanAddInput(input IAVCaptureInput) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddInput:"), input)
	return rv
}/* debug [instance_methods/method]: CanAddInput */


// Determines whether you can add an output to a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canAddOutput(_:)
func (c_ CaptureSession) CanAddOutput(output IAVCaptureOutput) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canAddOutput:"), output)
	return rv
}/* debug [instance_methods/method]: CanAddOutput */


// Determines whether you can configure a capture session with the specified preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/canSetSessionPreset(_:)
func (c_ CaptureSession) CanSetSessionPreset(preset CaptureSessionPreset /* typedef */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canSetSessionPreset:"), preset)
	return rv
}/* debug [instance_methods/method]: CanSetSessionPreset */


// Commits one or more changes to a running capture session’s configuration in a single atomic update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/commitConfiguration()
func (c_ CaptureSession) CommitConfiguration() {
	objc.Send[objc.ID](c_.ID, objc.Sel("commitConfiguration"))
}/* debug [instance_methods/method]: CommitConfiguration */


// Removes a capture connection from the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeConnection(_:)
func (c_ CaptureSession) RemoveConnection(connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeConnection:"), connection)
}/* debug [instance_methods/method]: RemoveConnection */


// Removes a control from a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeControl(_:)
func (c_ CaptureSession) RemoveControl(control IAVCaptureControl) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeControl:"), control)
}/* debug [instance_methods/method]: RemoveControl */


// Removes an input from the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeInput(_:)
func (c_ CaptureSession) RemoveInput(input IAVCaptureInput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeInput:"), input)
}/* debug [instance_methods/method]: RemoveInput */


// Removes an output from a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/removeOutput(_:)
func (c_ CaptureSession) RemoveOutput(output IAVCaptureOutput) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeOutput:"), output)
}/* debug [instance_methods/method]: RemoveOutput */


// Tells the session to run deferred start when appropriate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runDeferredStartWhenNeeded()
func (c_ CaptureSession) RunDeferredStartWhenNeeded() {
	objc.Send[objc.ID](c_.ID, objc.Sel("runDeferredStartWhenNeeded"))
}/* debug [instance_methods/method]: RunDeferredStartWhenNeeded */


// Sets a delegate object for the system to call when it activates and presents controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setControlsDelegate(_:queue:)
func (c_ CaptureSession) SetControlsDelegateQueue(controlsDelegate unsafe.Pointer, controlsDelegateCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setControlsDelegate:queue:"), controlsDelegate, controlsDelegateCallbackQueue)
}/* debug [instance_methods/method]: SetControlsDelegateQueue */


// Sets a delegate object for the session to call when performing deferred start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/setDeferredStartDelegate(_:deferredStartDelegateCallbackQueue:)
func (c_ CaptureSession) SetDeferredStartDelegateDeferredStartDelegateCallbackQueue(deferredStartDelegate unsafe.Pointer, deferredStartDelegateCallbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeferredStartDelegate:deferredStartDelegateCallbackQueue:"), deferredStartDelegate, deferredStartDelegateCallbackQueue)
}/* debug [instance_methods/method]: SetDeferredStartDelegateDeferredStartDelegateCallbackQueue */


// Starts the flow of data through the capture pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/startRunning()
func (c_ CaptureSession) StartRunning() {
	objc.Send[objc.ID](c_.ID, objc.Sel("startRunning"))
}/* debug [instance_methods/method]: StartRunning */


// Stops the flow of data through the capture pipeline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/stopRunning()
func (c_ CaptureSession) StopRunning() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopRunning"))
}/* debug [instance_methods/method]: StopRunning */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureSession */

// A Boolean value that indicates whether deferred start runs automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyRunsDeferredStart
func (c_ CaptureSession) AutomaticallyRunsDeferredStart() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("automaticallyRunsDeferredStart"))
	return rv
}/* debug [instance_properties/getter]: automaticallyRunsDeferredStart */


// A Boolean value that indicates whether deferred start runs automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyRunsDeferredStart
func (c_ CaptureSession) SetAutomaticallyRunsDeferredStart(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAutomaticallyRunsDeferredStart:"), value)
}/* debug [instance_properties/setter]: automaticallyRunsDeferredStart */


// The connections between inputs and outputs that a capture session contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/connections
func (c_ CaptureSession) Connections() []CaptureConnection {
	rv := objc.Send[[]CaptureConnection](c_.ID, objc.Sel("connections"))
	return rv
}/* debug [instance_properties/getter]: connections */


// The controls that allow configuring the camera system from device hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controls
func (c_ CaptureSession) Controls() []CaptureControl {
	rv := objc.Send[[]CaptureControl](c_.ID, objc.Sel("controls"))
	return rv
}/* debug [instance_properties/getter]: controls */


// A delegate object that observes changes to the state of capture controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegate
func (c_ CaptureSession) ControlsDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlsDelegate"))
	return rv
}/* debug [instance_properties/getter]: controlsDelegate */


// The dispatch queue on which the system calls controls delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/controlsDelegateCallbackQueue
func (c_ CaptureSession) ControlsDelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("controlsDelegateCallbackQueue"))
	return rv
}/* debug [instance_properties/getter]: controlsDelegateCallbackQueue */


// A delegate object that observes events about deferred start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegate
func (c_ CaptureSession) DeferredStartDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deferredStartDelegate"))
	return rv
}/* debug [instance_properties/getter]: deferredStartDelegate */


// The dispatch queue on which the session calls deferred start delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/deferredStartDelegateCallbackQueue
func (c_ CaptureSession) DeferredStartDelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("deferredStartDelegateCallbackQueue"))
	return rv
}/* debug [instance_properties/getter]: deferredStartDelegateCallbackQueue */


// The inputs that provide media data to a capture session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/inputs
func (c_ CaptureSession) Inputs() []CaptureInput {
	rv := objc.Send[[]CaptureInput](c_.ID, objc.Sel("inputs"))
	return rv
}/* debug [instance_properties/getter]: inputs */


// A value that indicates whether the session supports manually running deferred start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isManualDeferredStartSupported
func (c_ CaptureSession) ManualDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("manualDeferredStartSupported"))
	return rv
}/* debug [instance_properties/getter]: manualDeferredStartSupported */


// A Boolean value that indicates whether the capture session is in a running state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isRunning
func (c_ CaptureSession) Running() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// A clock object used for output synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/masterClock
func (c_ CaptureSession) MasterClock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](c_.ID, objc.Sel("masterClock"))
	return rv
}/* debug [instance_properties/getter]: masterClock */


// The maximum number of controls a capture session supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/maxControlsCount
func (c_ CaptureSession) MaxControlsCount() int {
	rv := objc.Send[int](c_.ID, objc.Sel("maxControlsCount"))
	return rv
}/* debug [instance_properties/getter]: maxControlsCount */


// The output destinations to which a captures session sends its data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/outputs
func (c_ CaptureSession) Outputs() []CaptureOutput {
	rv := objc.Send[[]CaptureOutput](c_.ID, objc.Sel("outputs"))
	return rv
}/* debug [instance_properties/getter]: outputs */


// A preset value that indicates the quality level or bit rate of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/sessionPreset
func (c_ CaptureSession) SessionPreset() CaptureSessionPreset /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("sessionPreset"))
	return rv
}/* debug [instance_properties/getter]: sessionPreset */


// A preset value that indicates the quality level or bit rate of the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/sessionPreset
func (c_ CaptureSession) SetSessionPreset(value CaptureSessionPreset /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSessionPreset:"), value)
}/* debug [instance_properties/setter]: sessionPreset */


// A Boolean value that indicates whether a capture session supports controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/supportsControls
func (c_ CaptureSession) SupportsControls() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsControls"))
	return rv
}/* debug [instance_properties/getter]: supportsControls */


// A clock to use for output synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/synchronizationClock
func (c_ CaptureSession) SynchronizationClock() ClockRef /* not a class type */ {
	rv := objc.Send[ClockRef](c_.ID, objc.Sel("synchronizationClock"))
	return rv
}/* debug [instance_properties/getter]: synchronizationClock */


// A Boolean value that indicates whether the capture session is in an interrupted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isinterrupted
func (c_ CaptureSession) IsInterrupted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isInterrupted"))
	return rv
}/* debug [instance_properties/getter]: isInterrupted */


// A Boolean value that indicates whether the capture session is in an interrupted state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isinterrupted
func (c_ CaptureSession) SetIsInterrupted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsInterrupted:"), value)
}/* debug [instance_properties/setter]: isInterrupted */


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismanualdeferredstartsupported
func (c_ CaptureSession) IsManualDeferredStartSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isManualDeferredStartSupported"))
	return rv
}/* debug [instance_properties/getter]: isManualDeferredStartSupported */


// A
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismanualdeferredstartsupported
func (c_ CaptureSession) SetIsManualDeferredStartSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsManualDeferredStartSupported:"), value)
}/* debug [instance_properties/setter]: isManualDeferredStartSupported */


// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccessenabled
func (c_ CaptureSession) IsMultitaskingCameraAccessEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultitaskingCameraAccessEnabled"))
	return rv
}/* debug [instance_properties/getter]: isMultitaskingCameraAccessEnabled */


// A Boolean value that indicates whether the capture session enables access to the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccessenabled
func (c_ CaptureSession) SetIsMultitaskingCameraAccessEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultitaskingCameraAccessEnabled:"), value)
}/* debug [instance_properties/setter]: isMultitaskingCameraAccessEnabled */


// A Boolean value that indicates whether the capture session supports using the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccesssupported
func (c_ CaptureSession) IsMultitaskingCameraAccessSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMultitaskingCameraAccessSupported"))
	return rv
}/* debug [instance_properties/getter]: isMultitaskingCameraAccessSupported */


// A Boolean value that indicates whether the capture session supports using the camera while multitasking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/ismultitaskingcameraaccesssupported
func (c_ CaptureSession) SetIsMultitaskingCameraAccessSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMultitaskingCameraAccessSupported:"), value)
}/* debug [instance_properties/setter]: isMultitaskingCameraAccessSupported */


// A Boolean value that indicates whether the capture session is in a running state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isrunning
func (c_ CaptureSession) IsRunning() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// A Boolean value that indicates whether the capture session is in a running state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturesession/isrunning
func (c_ CaptureSession) SetIsRunning(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureSession */


