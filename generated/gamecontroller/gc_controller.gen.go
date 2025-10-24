// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCController */


/* debug [class_header]: Header for GCController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCController */
// An interface definition for the [GCController] class.
type IGCController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCController */
	// properties:
	Battery() IGCDeviceBattery
	ControllerPausedHandler() unsafe.Pointer
	SetControllerPausedHandler(value unsafe.Pointer)
	ExtendedGamepad() IGCExtendedGamepad
	Gamepad() IGCGamepad
	Haptics() IGCDeviceHaptics
	Input() IGCControllerLiveInput
	AttachedToDevice() bool
	Snapshot() bool
	Light() IGCDeviceLight
	MicroGamepad() IGCMicroGamepad
	Motion() IGCMotion
	PhysicalInputProfile() IGCPhysicalInputProfile
	PlayerIndex() GCControllerPlayerIndex
	SetPlayerIndex(value GCControllerPlayerIndex)
	IsAttachedToDevice() bool
	SetIsAttachedToDevice(value bool)
	IsSnapshot() bool
	SetIsSnapshot(value bool)
	LeftThumbstick() IGCControllerDirectionPad
	SetLeftThumbstick(value IGCControllerDirectionPad)
	ValueChangedHandler() unsafe.Pointer
	SetValueChangedHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCController */
	// methods:
	Capture() IGCController
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCController */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerClass) Alloc() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCController */

// Returns the connected controllers for the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/controllers()
func (gc _GCControllerClass) Controllers() []GCController {
	rv := objc.Send[[]GCController](objc.ID(gc.class), objc.Sel("controllers"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Controllers) */


// Starts searching for nearby wireless controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/startWirelessControllerDiscovery(completionHandler:)
func (gc _GCControllerClass) StartWirelessControllerDiscoveryWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("startWirelessControllerDiscoveryWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StartWirelessControllerDiscoveryWithCompletionHandler) */


// Stops searching for nearby wireless controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/stopWirelessControllerDiscovery()
func (gc _GCControllerClass) StopWirelessControllerDiscovery() {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("stopWirelessControllerDiscovery"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StopWirelessControllerDiscovery) */


// Returns a Boolean value that indicates whether the framework supports the specified human interface device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/supportsHIDDevice(_:)
func (gc _GCControllerClass) SupportsHIDDevice(device HIDDeviceRef /* not a class type */) bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("supportsHIDDevice:"), device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsHIDDevice) */


// Returns a snapshot of a newly created controller with an extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/withExtendedGamepad()
func (gc _GCControllerClass) ControllerWithExtendedGamepad() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("controllerWithExtendedGamepad"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ControllerWithExtendedGamepad) */


// Returns a snapshot of a newly created controller with a micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/withMicroGamepad()
func (gc _GCControllerClass) ControllerWithMicroGamepad() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("controllerWithMicroGamepad"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ControllerWithMicroGamepad) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCController */

// The most recently used game controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/current
func (gc _GCControllerClass) Current() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("current"))
	return rv
}/* debug [class_properties_class/property]: current */

// A Boolean value that indicates whether the app needs to respond to controller events when it isn’t the frontmost app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/shouldMonitorBackgroundEvents
func (gc _GCControllerClass) ShouldMonitorBackgroundEvents() bool {
	rv := objc.Send[bool](objc.ID(gc.class), objc.Sel("shouldMonitorBackgroundEvents"))
	return rv
}/* debug [class_properties_class/property]: shouldMonitorBackgroundEvents */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCController */

// Returns a snapshot of the controller with its current element values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/capture()
func (g_ GCController) Capture() GCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("capture"))
	return rv
}/* debug [instance_methods/method]: Capture */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCController */

// The controller’s battery information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/battery
func (g_ GCController) Battery() IGCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g_.ID, objc.Sel("battery"))
	return rv
}/* debug [instance_properties/getter]: battery */


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/controllerPausedHandler
func (g_ GCController) ControllerPausedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controllerPausedHandler"))
	return rv
}/* debug [instance_properties/getter]: controllerPausedHandler */


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/controllerPausedHandler
func (g_ GCController) SetControllerPausedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setControllerPausedHandler:"), value)
}/* debug [instance_properties/setter]: controllerPausedHandler */


// The most recently used game controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/current
func (g_ GCController) Current() IGCController {
	rv := objc.Send[GCController](g_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// The extended gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/extendedGamepad
func (g_ GCController) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}/* debug [instance_properties/getter]: extendedGamepad */


// The gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/gamepad
func (g_ GCController) Gamepad() IGCGamepad {
	rv := objc.Send[GCGamepad](g_.ID, objc.Sel("gamepad"))
	return rv
}/* debug [instance_properties/getter]: gamepad */


// The controller’s haptics information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/haptics
func (g_ GCController) Haptics() IGCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g_.ID, objc.Sel("haptics"))
	return rv
}/* debug [instance_properties/getter]: haptics */


// The input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/input
func (g_ GCController) Input() IGCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g_.ID, objc.Sel("input"))
	return rv
}/* debug [instance_properties/getter]: input */


// A Boolean value that indicates whether the controller closely integrates with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/isAttachedToDevice
func (g_ GCController) AttachedToDevice() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("attachedToDevice"))
	return rv
}/* debug [instance_properties/getter]: attachedToDevice */


// A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/isSnapshot
func (g_ GCController) Snapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("snapshot"))
	return rv
}/* debug [instance_properties/getter]: snapshot */


// The controller’s light settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/light
func (g_ GCController) Light() IGCDeviceLight {
	rv := objc.Send[GCDeviceLight](g_.ID, objc.Sel("light"))
	return rv
}/* debug [instance_properties/getter]: light */


// The micro gamepad profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/microGamepad
func (g_ GCController) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g_.ID, objc.Sel("microGamepad"))
	return rv
}/* debug [instance_properties/getter]: microGamepad */


// The motion input profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/motion
func (g_ GCController) Motion() IGCMotion {
	rv := objc.Send[GCMotion](g_.ID, objc.Sel("motion"))
	return rv
}/* debug [instance_properties/getter]: motion */


// The physical input profile for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/physicalInputProfile
func (g_ GCController) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g_.ID, objc.Sel("physicalInputProfile"))
	return rv
}/* debug [instance_properties/getter]: physicalInputProfile */


// The player index for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) PlayerIndex() GCControllerPlayerIndex {
	rv := objc.Send[GCControllerPlayerIndex](g_.ID, objc.Sel("playerIndex"))
	return rv
}/* debug [instance_properties/getter]: playerIndex */


// The player index for the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) SetPlayerIndex(value GCControllerPlayerIndex) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerIndex:"), value)
}/* debug [instance_properties/setter]: playerIndex */


// A Boolean value that indicates whether the app needs to respond to controller events when it isn’t the frontmost app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/shouldMonitorBackgroundEvents
func (g_ GCController) ShouldMonitorBackgroundEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("shouldMonitorBackgroundEvents"))
	return rv
}/* debug [instance_properties/getter]: shouldMonitorBackgroundEvents */


// A Boolean value that indicates whether the app needs to respond to controller events when it isn’t the frontmost app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/shouldMonitorBackgroundEvents
func (g_ GCController) SetShouldMonitorBackgroundEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setShouldMonitorBackgroundEvents:"), value)
}/* debug [instance_properties/setter]: shouldMonitorBackgroundEvents */


// A Boolean value that indicates whether the controller closely integrates with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/isattachedtodevice
func (g_ GCController) IsAttachedToDevice() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAttachedToDevice"))
	return rv
}/* debug [instance_properties/getter]: isAttachedToDevice */


// A Boolean value that indicates whether the controller closely integrates with the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/isattachedtodevice
func (g_ GCController) SetIsAttachedToDevice(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAttachedToDevice:"), value)
}/* debug [instance_properties/setter]: isAttachedToDevice */


// A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/issnapshot
func (g_ GCController) IsSnapshot() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isSnapshot"))
	return rv
}/* debug [instance_properties/getter]: isSnapshot */


// A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/issnapshot
func (g_ GCController) SetIsSnapshot(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsSnapshot:"), value)
}/* debug [instance_properties/setter]: isSnapshot */


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCController) LeftThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g_.ID, objc.Sel("leftThumbstick"))
	return rv
}/* debug [instance_properties/getter]: leftThumbstick */


// The controller’s left thumbstick element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g_ GCController) SetLeftThumbstick(value IGCControllerDirectionPad) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeftThumbstick:"), value)
}/* debug [instance_properties/setter]: leftThumbstick */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/valuechangedhandler
func (g_ GCController) ValueChangedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("valueChangedHandler"))
	return rv
}/* debug [instance_properties/getter]: valueChangedHandler */


// The block that the profile calls when an element’s value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/valuechangedhandler
func (g_ GCController) SetValueChangedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValueChangedHandler:"), value)
}/* debug [instance_properties/setter]: valueChangedHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCController */



