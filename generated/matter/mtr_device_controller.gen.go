// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceController */


/* debug [class_header]: Header for MTRDeviceController */
// The class instance for the [MTRDeviceController] class.
var (
	MTRDeviceControllerClass     _MTRDeviceControllerClass
	MTRDeviceControllerClassOnce sync.Once
)

func getMTRDeviceControllerClass() _MTRDeviceControllerClass {
	MTRDeviceControllerClassOnce.Do(func() {
		MTRDeviceControllerClass = _MTRDeviceControllerClass{objc.GetClass("MTRDeviceController")}
	})
	return MTRDeviceControllerClass
}

type _MTRDeviceControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceController */
// An interface definition for the [MTRDeviceController] class.
type IMTRDeviceController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceController */
	// properties:
	ControllerNodeId() objc.IObject /* cross-framework: NSNumber */
	ControllerNodeID() objc.IObject /* cross-framework: NSNumber */
	Devices() []MTRDevice
	Running() bool
	Suspended() bool
	NodesWithStoredData() []foundation.Number
	UniqueIdentifier() foundation.UUID
	IsRunning() bool
	SetIsRunning(value bool)
	IsSuspended() bool
	SetIsSuspended(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceController */
	// methods:
	AddServerEndpoint(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */) bool
	AddDeviceControllerDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer)
	AttestationChallengeForDeviceID(deviceID objc.IObject /* cross-framework: NSNumber */) foundation.Data
	CancelCommissioningForNodeIDError(nodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool
	CommissionNodeWithIDCommissioningParamsError(nodeID objc.IObject /* cross-framework: NSNumber */, commissioningParams IMTRCommissioningParameters, error_ unsafe.Pointer) bool
	ContinueCommissioningDeviceIgnoreAttestationFailureError(opaqueDeviceHandle unsafe.Pointer, ignoreAttestationFailure bool, error_ unsafe.Pointer) bool
	DeviceBeingCommissionedWithNodeIDError(nodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) IMTRBaseDevice
	ForgetDeviceWithNodeID(nodeID objc.IObject /* cross-framework: NSNumber */)
	RemoveServerEndpoint(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */)
	RemoveDeviceControllerDelegate(delegate unsafe.Pointer)
	RemoveServerEndpointQueueCompletion(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */, queue unsafe.Pointer, completion unsafe.Pointer)
	Resume()
	SetDeviceControllerDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer)
	SetupCommissioningSessionWithPayloadNewNodeIDError(payload objc.IObject /* cross-framework: MTRSetupPayload */, newNodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool
	SetupCommissioningSessionWithDiscoveredDevicePayloadNewNodeIDError(discoveredDevice objc.IObject /* cross-framework: MTRCommissionableBrowserResult */, payload objc.IObject /* cross-framework: MTRSetupPayload */, newNodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool
	Shutdown()
	StartBrowseForCommissionablesQueue(delegate unsafe.Pointer, queue unsafe.Pointer) bool
	StopBrowseForCommissionables() bool
	Suspend()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceController */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerClass) Alloc() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceControllerClass) New() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceController) Init() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceController) Autorelease() MTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceController creates a new MTRDeviceController instance.
func NewMTRDeviceController() MTRDeviceController {
	return getMTRDeviceControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController
type MTRDeviceController struct {
	objectivec.Object
}

// MTRDeviceControllerFrom constructs a [MTRDeviceController] from an unsafe.Pointer.
func MTRDeviceControllerFrom(ptr unsafe.Pointer) MTRDeviceController {
	return MTRDeviceController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/init(parameters:)
func NewMTRDeviceControllerWithParametersError(parameters objc.IObject /* cross-framework: MTRDeviceControllerAbstractParameters */, error_ unsafe.Pointer) MTRDeviceController {
	instance := getMTRDeviceControllerClass().Alloc()
	rv := objc.Send[MTRDeviceController](instance.ID, objc.Sel("initWithParameters:error:"), parameters, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerWithParametersError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/computePASEVerifier(forSetupPasscode:iterations:salt:)
func (mc _MTRDeviceControllerClass) ComputePASEVerifierForSetupPasscodeIterationsSaltError(setupPasscode objc.IObject /* cross-framework: NSNumber */, iterations objc.IObject /* cross-framework: NSNumber */, salt objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](objc.ID(mc.class), objc.Sel("computePASEVerifierForSetupPasscode:iterations:salt:error:"), setupPasscode, iterations, salt, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComputePASEVerifierForSetupPasscodeIterationsSaltError) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/decodeXPCReadParams(_:)
func (mc _MTRDeviceControllerClass) DecodeXPCReadParams(params foundation.IDictionary) IMTRReadParams {
	rv := objc.Send[MTRReadParams](objc.ID(mc.class), objc.Sel("decodeXPCReadParams:"), params)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecodeXPCReadParams) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/decodeXPCResponseValues(_:)
func (mc _MTRDeviceControllerClass) DecodeXPCResponseValues(values foundation.IDictionary) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(mc.class), objc.Sel("decodeXPCResponseValues:"), values)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecodeXPCResponseValues) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/decodeXPCSubscribeParams(_:)
func (mc _MTRDeviceControllerClass) DecodeXPCSubscribeParams(params foundation.IDictionary) IMTRSubscribeParams {
	rv := objc.Send[MTRSubscribeParams](objc.ID(mc.class), objc.Sel("decodeXPCSubscribeParams:"), params)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecodeXPCSubscribeParams) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/encodeXPCReadParams(_:)
func (mc _MTRDeviceControllerClass) EncodeXPCReadParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(mc.class), objc.Sel("encodeXPCReadParams:"), params)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EncodeXPCReadParams) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/encodeXPCResponseValues(_:)
func (mc _MTRDeviceControllerClass) EncodeXPCResponseValues(values foundation.IDictionary) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(mc.class), objc.Sel("encodeXPCResponseValues:"), values)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EncodeXPCResponseValues) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/encodeXPCSubscribeParams(_:)
func (mc _MTRDeviceControllerClass) EncodeXPCSubscribeParams(params IMTRSubscribeParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](objc.ID(mc.class), objc.Sel("encodeXPCSubscribeParams:"), params)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EncodeXPCSubscribeParams) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/sharedController(withID:xpcConnect:)-5yhq4
func (mc _MTRDeviceControllerClass) SharedControllerWithIDXpcConnectBlock(controllerID unsafe.Pointer, xpcConnectBlock unsafe.Pointer) MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("sharedControllerWithID:xpcConnectBlock:"), controllerID, xpcConnectBlock)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedControllerWithIDXpcConnectBlock) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/sharedController(withId:xpcConnect:)-6rg64
func (mc _MTRDeviceControllerClass) SharedControllerWithIdXpcConnectBlock(controllerID unsafe.Pointer, xpcConnectBlock unsafe.Pointer) MTRDeviceController {
	rv := objc.Send[MTRDeviceController](objc.ID(mc.class), objc.Sel("sharedControllerWithId:xpcConnectBlock:"), controllerID, xpcConnectBlock)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedControllerWithIdXpcConnectBlock) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/xpcInterfaceForClientProtocol()
func (mc _MTRDeviceControllerClass) XpcInterfaceForClientProtocol() foundation.XPCInterface {
	rv := objc.Send[foundation.XPCInterface](objc.ID(mc.class), objc.Sel("xpcInterfaceForClientProtocol"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=XpcInterfaceForClientProtocol) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/xpcInterfaceForServerProtocol()
func (mc _MTRDeviceControllerClass) XpcInterfaceForServerProtocol() foundation.XPCInterface {
	rv := objc.Send[foundation.XPCInterface](objc.ID(mc.class), objc.Sel("xpcInterfaceForServerProtocol"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=XpcInterfaceForServerProtocol) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/add(_:)
func (m_ MTRDeviceController) AddServerEndpoint(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("addServerEndpoint:"), endpoint)
	return rv
}/* debug [instance_methods/method]: AddServerEndpoint */


// Adds a Delegate to the device controller as well as the Queue on which the Delegate callbacks will be triggered
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/add(_:queue:)
func (m_ MTRDeviceController) AddDeviceControllerDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addDeviceControllerDelegate:queue:"), delegate, queue)
}/* debug [instance_methods/method]: AddDeviceControllerDelegateQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/attestationChallenge(forDeviceID:)
func (m_ MTRDeviceController) AttestationChallengeForDeviceID(deviceID objc.IObject /* cross-framework: NSNumber */) foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("attestationChallengeForDeviceID:"), deviceID)
	return rv
}/* debug [instance_methods/method]: AttestationChallengeForDeviceID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/cancelCommissioning(forNodeID:)
func (m_ MTRDeviceController) CancelCommissioningForNodeIDError(nodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("cancelCommissioningForNodeID:error:"), nodeID, error_)
	return rv
}/* debug [instance_methods/method]: CancelCommissioningForNodeIDError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/commissionNode(withID:commissioningParams:)
func (m_ MTRDeviceController) CommissionNodeWithIDCommissioningParamsError(nodeID objc.IObject /* cross-framework: NSNumber */, commissioningParams IMTRCommissioningParameters, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("commissionNodeWithID:commissioningParams:error:"), nodeID, commissioningParams, error_)
	return rv
}/* debug [instance_methods/method]: CommissionNodeWithIDCommissioningParamsError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/continueCommissioningDevice(_:ignoreAttestationFailure:)
func (m_ MTRDeviceController) ContinueCommissioningDeviceIgnoreAttestationFailureError(opaqueDeviceHandle unsafe.Pointer, ignoreAttestationFailure bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("continueCommissioningDevice:ignoreAttestationFailure:error:"), opaqueDeviceHandle, ignoreAttestationFailure, error_)
	return rv
}/* debug [instance_methods/method]: ContinueCommissioningDeviceIgnoreAttestationFailureError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/deviceBeingCommissioned(withNodeID:)
func (m_ MTRDeviceController) DeviceBeingCommissionedWithNodeIDError(nodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) IMTRBaseDevice {
	rv := objc.Send[MTRBaseDevice](m_.ID, objc.Sel("deviceBeingCommissionedWithNodeID:error:"), nodeID, error_)
	return rv
}/* debug [instance_methods/method]: DeviceBeingCommissionedWithNodeIDError */


// Forget any information we have about the device with the given node ID. That includes clearing any information we have stored about it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/forgetDevice(withNodeID:)
func (m_ MTRDeviceController) ForgetDeviceWithNodeID(nodeID objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("forgetDeviceWithNodeID:"), nodeID)
}/* debug [instance_methods/method]: ForgetDeviceWithNodeID */


// Remove the given server endpoint without being notified when the removal completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/remove(_:)-2i5l5
func (m_ MTRDeviceController) RemoveServerEndpoint(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeServerEndpoint:"), endpoint)
}/* debug [instance_methods/method]: RemoveServerEndpoint */


// Removes a Delegate from the device controller
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/remove(_:)-8pxve
func (m_ MTRDeviceController) RemoveDeviceControllerDelegate(delegate unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeDeviceControllerDelegate:"), delegate)
}/* debug [instance_methods/method]: RemoveDeviceControllerDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/remove(_:queue:completion:)
func (m_ MTRDeviceController) RemoveServerEndpointQueueCompletion(endpoint objc.IObject /* cross-framework: MTRServerEndpoint */, queue unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeServerEndpoint:queue:completion:"), endpoint, queue, completion)
}/* debug [instance_methods/method]: RemoveServerEndpointQueueCompletion */


// Resume the controller. This has no effect if the controller is not suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/resume()
func (m_ MTRDeviceController) Resume() {
	objc.Send[objc.ID](m_.ID, objc.Sel("resume"))
}/* debug [instance_methods/method]: Resume */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setDeviceControllerDelegate(_:queue:)
func (m_ MTRDeviceController) SetDeviceControllerDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDeviceControllerDelegate:queue:"), delegate, queue)
}/* debug [instance_methods/method]: SetDeviceControllerDelegateQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setupCommissioningSession(with:newNodeID:)
func (m_ MTRDeviceController) SetupCommissioningSessionWithPayloadNewNodeIDError(payload objc.IObject /* cross-framework: MTRSetupPayload */, newNodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setupCommissioningSessionWithPayload:newNodeID:error:"), payload, newNodeID, error_)
	return rv
}/* debug [instance_methods/method]: SetupCommissioningSessionWithPayloadNewNodeIDError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/setupCommissioningSession(withDiscoveredDevice:payload:newNodeID:)
func (m_ MTRDeviceController) SetupCommissioningSessionWithDiscoveredDevicePayloadNewNodeIDError(discoveredDevice objc.IObject /* cross-framework: MTRCommissionableBrowserResult */, payload objc.IObject /* cross-framework: MTRSetupPayload */, newNodeID objc.IObject /* cross-framework: NSNumber */, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setupCommissioningSessionWithDiscoveredDevice:payload:newNodeID:error:"), discoveredDevice, payload, newNodeID, error_)
	return rv
}/* debug [instance_methods/method]: SetupCommissioningSessionWithDiscoveredDevicePayloadNewNodeIDError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/shutdown()
func (m_ MTRDeviceController) Shutdown() {
	objc.Send[objc.ID](m_.ID, objc.Sel("shutdown"))
}/* debug [instance_methods/method]: Shutdown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/startBrowse(forCommissionables:queue:)
func (m_ MTRDeviceController) StartBrowseForCommissionablesQueue(delegate unsafe.Pointer, queue unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startBrowseForCommissionables:queue:"), delegate, queue)
	return rv
}/* debug [instance_methods/method]: StartBrowseForCommissionablesQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/stopBrowseForCommissionables()
func (m_ MTRDeviceController) StopBrowseForCommissionables() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("stopBrowseForCommissionables"))
	return rv
}/* debug [instance_methods/method]: StopBrowseForCommissionables */


// Suspend the controller. This will attempt to stop all network traffic associated with the controller. The controller will remain suspended until it is resumed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/suspend()
func (m_ MTRDeviceController) Suspend() {
	objc.Send[objc.ID](m_.ID, objc.Sel("suspend"))
}/* debug [instance_methods/method]: Suspend */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/controllerNodeId-6a03y
func (m_ MTRDeviceController) ControllerNodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("controllerNodeId"))
	return rv
}/* debug [instance_properties/getter]: controllerNodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/controllerNodeID-6a04u
func (m_ MTRDeviceController) ControllerNodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("controllerNodeID"))
	return rv
}/* debug [instance_properties/getter]: controllerNodeID */


// Returns the list of MTRDevice instances that this controller has loaded into memory. Returns an empty array if no devices are in memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/devices
func (m_ MTRDeviceController) Devices() []MTRDevice {
	rv := objc.Send[[]MTRDevice](m_.ID, objc.Sel("devices"))
	return rv
}/* debug [instance_properties/getter]: devices */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/isRunning
func (m_ MTRDeviceController) Running() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// If true, the controller has been suspended via and not resumed yet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/isSuspended
func (m_ MTRDeviceController) Suspended() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("suspended"))
	return rv
}/* debug [instance_properties/getter]: suspended */


// Returns the list of node IDs for which this controller has stored information. Returns empty list if the controller does not have any information stored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/nodesWithStoredData
func (m_ MTRDeviceController) NodesWithStoredData() []foundation.Number {
	rv := objc.Send[[]foundation.Number](m_.ID, objc.Sel("nodesWithStoredData"))
	return rv
}/* debug [instance_properties/getter]: nodesWithStoredData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceController/uniqueIdentifier
func (m_ MTRDeviceController) UniqueIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](m_.ID, objc.Sel("uniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: uniqueIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/isrunning
func (m_ MTRDeviceController) IsRunning() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/isrunning
func (m_ MTRDeviceController) SetIsRunning(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */


// If true, the controller has been suspended via
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/issuspended
func (m_ MTRDeviceController) IsSuspended() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isSuspended"))
	return rv
}/* debug [instance_properties/getter]: isSuspended */


// If true, the controller has been suspended via
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontroller/issuspended
func (m_ MTRDeviceController) SetIsSuspended(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsSuspended:"), value)
}/* debug [instance_properties/setter]: isSuspended */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceController */


