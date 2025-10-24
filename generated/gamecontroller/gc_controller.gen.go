// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCController] class.
var (
	GCControllerClass     _GCControllerClass
	GCControllerClassOnce sync.Once
)

func getGCControllerClass() _GCControllerClass {
	GCControllerClassOnce.Do(func() {
		GCControllerClass = _GCControllerClass{objc.GetClass("GCController")}
	})
	return GCControllerClass
}

type _GCControllerClass struct {
	class objc.Class
}

// An interface definition for the [GCController] class.
type IGCController interface {
	objectivec.IObject
	// properties:
	ExtendedGamepad() IGCExtendedGamepad
	PhysicalInputProfile() IGCPhysicalInputProfile
	PlayerIndex() unsafe.Pointer
	SetPlayerIndex(value unsafe.Pointer)
	Battery() IGCDeviceBattery
	SetBattery(value IGCDeviceBattery)
	ControllerPausedHandler() unsafe.Pointer
	SetControllerPausedHandler(value unsafe.Pointer)
	Gamepad() unsafe.Pointer
	SetGamepad(value unsafe.Pointer)
	Haptics() IGCDeviceHaptics
	SetHaptics(value IGCDeviceHaptics)
	Input() objc.IObject /* cross-framework: GCControllerLiveInput */
	SetInput(value objc.IObject /* cross-framework: GCControllerLiveInput */)
	IsAttachedToDevice() bool
	SetIsAttachedToDevice(value bool)
	IsSnapshot() bool
	SetIsSnapshot(value bool)
	Light() IGCDeviceLight
	SetLight(value IGCDeviceLight)
	MicroGamepad() objc.IObject /* cross-framework: GCMicroGamepad */
	SetMicroGamepad(value objc.IObject /* cross-framework: GCMicroGamepad */)
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	LeftThumbstick() objc.IObject /* cross-framework: GCControllerDirectionPad */
	SetLeftThumbstick(value objc.IObject /* cross-framework: GCControllerDirectionPad */)
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
	// methods:
	Capture() IGCController
}

// A representation of a real game controller, a virtual controller, or a snapshot of a controller.
//
// This class represents a real or virtual controller that a user interacts with during a game. A is a physical controller that connects directly or wirelessly to the device. A real controller can be formfitting or can attach closely to a device so players can use controls on both simultaneously. A is a software emulation of a real controller. You discover controllers, and then you process the input from those controllers during gameplay. Use the method to get the currently connected controllers. If necessary, use the method to connect with wireless controllers. This framework supports multiple connected game controllers. To identify which player is using a controller in a multiplayer game, check the property and set it, if necessary. For single-player games, use the property to get the controller that the player is actively using. A controller’s profile encapsulates the details about a controller’s buttons, pads, axis, and other input elements. Get the controller’s profile using one of the profile properties, such as , and then process the input from its elements. You can either get the values of input elements on each iteration of your game loop, or set handlers to receive callbacks when those values change. For example, use the property of the profile to get the thumbstick state. Use the property to set a handler that you implement to process any input values that change in the profile. Alternatively, you can create a snapshot of a real or virtual controller using the method. A is a copy of a controller at a moment in time with its current element values. Creating a snapshot may impact performance, and over time a snapshot doesn’t stay current. Unlike other types of controllers, you can set the values of elements in a snapshot.


// A representation of a real game controller, a virtual controller, or a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController
type GCController struct {
	objectivec.Object
}

// GCControllerFrom constructs a [GCController] from an unsafe.Pointer.
//
// A representation of a real game controller, a virtual controller, or a snapshot of a controller.
func GCControllerFrom(ptr unsafe.Pointer) GCController {
	return GCController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerClass) Alloc() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerClass) New() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCController) Init() GCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCController) Autorelease() GCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCController creates a new GCController instance.
func NewGCController() GCController {
	return getGCControllerClass().New()
}



// Returns the connected controllers for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/controllers()
func (gc _GCControllerClass) Controllers() []IGCController {
	rv := objc.Send[[]GCController](objc.ID(gc.class), objc.Sel("controllers"))
	return rv
}


// Starts searching for nearby wireless controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/startWirelessControllerDiscovery(completionHandler:)
func (gc _GCControllerClass) StartWirelessControllerDiscoveryWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("startWirelessControllerDiscoveryWithCompletionHandler:"), completionHandler)
}


// Stops searching for nearby wireless controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/stopWirelessControllerDiscovery()
func (gc _GCControllerClass) StopWirelessControllerDiscovery() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("stopWirelessControllerDiscovery"))
}


// The most recently used game controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/current
func (gc _GCControllerClass) Current() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("current"))
	return rv
}

// Returns a snapshot of the controller with its current element values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/capture()
func (g_ GCController) Capture() GCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("capture"))
	return rv
}


// The most recently used game controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/current
func (g_ GCController) Current() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("current"))
	return rv
}


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/extendedGamepad
func (g_ GCController) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/physicalInputProfile
func (g_ GCController) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}


// The player index for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) PlayerIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playerIndex"))
	return rv
}


// The player index for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) SetPlayerIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerIndex:"), value)
}


// The controller’s battery information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/battery
func (g_ GCController) Battery() IGCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g_.ID, objc.Sel("battery"))
	return rv
}


// The controller’s battery information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/battery
func (g_ GCController) SetBattery(value IGCDeviceBattery) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setBattery:"), value)
}


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCController) ControllerPausedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controllerPausedHandler"))
	return rv
}


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCController) SetControllerPausedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setControllerPausedHandler:"), value)
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCController) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g_ GCController) SetGamepad(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGamepad:"), value)
}


// The controller’s haptics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/haptics
func (g_ GCController) Haptics() IGCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("haptics"))
	return rv
}


// The controller’s haptics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/haptics
func (g_ GCController) SetHaptics(value IGCDeviceHaptics) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHaptics:"), value)
}


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCController) Input() objc.IObject /* cross-framework: GCControllerLiveInput */ {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("input"))
	return rv
}


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g_ GCController) SetInput(value objc.IObject /* cross-framework: GCControllerLiveInput */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setInput:"), value)
}


// A Boolean value that indicates whether the controller closely integrates with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/isattachedtodevice
func (g_ GCController) IsAttachedToDevice() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAttachedToDevice"))
	return rv
}


// A Boolean value that indicates whether the controller closely integrates with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/isattachedtodevice
func (g_ GCController) SetIsAttachedToDevice(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAttachedToDevice:"), value)
}


// A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/issnapshot
func (g_ GCController) IsSnapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isSnapshot"))
	return rv
}


// A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/issnapshot
func (g_ GCController) SetIsSnapshot(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsSnapshot:"), value)
}


// The controller’s light settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/light
func (g_ GCController) Light() IGCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("light"))
	return rv
}


// The controller’s light settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/light
func (g_ GCController) SetLight(value IGCDeviceLight) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLight:"), value)
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCController) MicroGamepad() objc.IObject /* cross-framework: GCMicroGamepad */ {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g_ GCController) SetMicroGamepad(value objc.IObject /* cross-framework: GCMicroGamepad */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMicroGamepad:"), value)
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCController) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g_ GCController) SetMotion(value IGCMotion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMotion:"), value)
}


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCController) LeftThumbstick() objc.IObject /* cross-framework: GCControllerDirectionPad */ {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("leftThumbstick"))
	return rv
}


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCController) SetLeftThumbstick(value objc.IObject /* cross-framework: GCControllerDirectionPad */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftThumbstick:"), value)
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/valuechangedhandler
func (g_ GCController) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/valuechangedhandler
func (g_ GCController) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}



