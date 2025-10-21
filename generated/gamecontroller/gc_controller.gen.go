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
}

// A representation of a real game controller, a virtual controller, or a snapshot of a controller.
//
// This class represents a real or virtual controller that a user interacts with during a game. A is a physical controller that connects directly or wirelessly to the device. A real controller can be formfitting or can attach closely to a device so players can use controls on both simultaneously. A is a software emulation of a real controller. You discover controllers, and then you process the input from those controllers during gameplay. Use the method to get the currently connected controllers. If necessary, use the method to connect with wireless controllers. This framework supports multiple connected game controllers. To identify which player is using a controller in a multiplayer game, check the property and set it, if necessary. For single-player games, use the property to get the controller that the player is actively using. A controller’s profile encapsulates the details about a controller’s buttons, pads, axis, and other input elements. Get the controller’s profile using one of the profile properties, such as , and then process the input from its elements. You can either get the values of input elements on each iteration of your game loop, or set handlers to receive callbacks when those values change. For example, use the property of the profile to get the thumbstick state. Use the property to set a handler that you implement to process any input values that change in the profile. Alternatively, you can create a snapshot of a real or virtual controller using the method. A is a copy of a controller at a moment in time with its current element values. Creating a snapshot may impact performance, and over time a snapshot doesn’t stay current. Unlike other types of controllers, you can set the values of elements in a snapshot.
//
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
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/controllers()
func (gc _GCControllerClass) Controllers() []GCController {
	rv := objc.Send[[]GCController](objc.ID(gc.class), objc.Sel("controllers"))
	return rv
}

// Starts searching for nearby wireless controllers.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/startWirelessControllerDiscovery(completionHandler:)
func (gc _GCControllerClass) StartWirelessControllerDiscoveryWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(gc.class), objc.Sel("startWirelessControllerDiscoveryWithCompletionHandler:"), completionHandler)
}

// The extended gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/extendedGamepad
func (g_ GCController) ExtendedGamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("extendedGamepad"))
	return rv
}

// The gamepad profile.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/gamepad
func (g_ GCController) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gamepad"))
	return rv
}

// The player index for the controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) PlayerIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("playerIndex"))
	return rv
}


// SetPlayerIndex sets the value of the playerIndex property.
// The player index for the controller.

//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g_ GCController) SetPlayerIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPlayerIndex:"), value)
}



