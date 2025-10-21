// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCVirtualController] class.
var (
	GCVirtualControllerClass     _GCVirtualControllerClass
	GCVirtualControllerClassOnce sync.Once
)

func getGCVirtualControllerClass() _GCVirtualControllerClass {
	GCVirtualControllerClassOnce.Do(func() {
		GCVirtualControllerClass = _GCVirtualControllerClass{objc.GetClass("GCVirtualController")}
	})
	return GCVirtualControllerClass
}

type _GCVirtualControllerClass struct {
	class objc.Class
}

// An interface definition for the [GCVirtualController] class.
type IGCVirtualController interface {
	objectivec.IObject
	ConnectWithReplyHandler(reply unsafe.Pointer)
	SetValueForButtonElement(value float64, element string)
}

// A software emulation of a real controller that you configure specifically for your game.
//
// Use a virtual controller to display software controls that you can customize over your game. You create a virtual controller from a configuration where you choose the input elements to display. You can even customize the images for the elements. When you connect the controller to the device, users interact with it similarly to a real controller. To add a virtual controller to your game, create a object containing the elements you want to appear in the controller. Then create the virtual controller by passing the configuration to the method. Use the method to display the virtual controller on the screen. To customize an element in the virtual controller, pass a new object for the element to the method. You process input from a virtual controller similarly to a real controller. Use the property to get the underlying object. You can either poll the elements of the controller object or set the element’s handlers to get callbacks when their input values change.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController
type GCVirtualController struct {
	objectivec.Object
}

// GCVirtualControllerFrom constructs a [GCVirtualController] from an unsafe.Pointer.
//
// A software emulation of a real controller that you configure specifically for your game.
func GCVirtualControllerFrom(ptr unsafe.Pointer) GCVirtualController {
	return GCVirtualController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerClass) Alloc() GCVirtualController {
	rv := objc.Send[GCVirtualController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCVirtualControllerClass) New() GCVirtualController {
	rv := objc.Send[GCVirtualController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCVirtualController) Init() GCVirtualController {
	rv := objc.Send[GCVirtualController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCVirtualController) Autorelease() GCVirtualController {
	rv := objc.Send[GCVirtualController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCVirtualController creates a new GCVirtualController instance.
func NewGCVirtualController() GCVirtualController {
	return getGCVirtualControllerClass().New()
}




// Creates a new virtual controller using the configuration you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/init(configuration:)
func NewGCVirtualControllerWithConfiguration(configuration unsafe.Pointer) GCVirtualController {
	instance := getGCVirtualControllerClass().Alloc()
	rv := objc.Send[GCVirtualController](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}


// Creates a new virtual controller using the configuration you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/virtualControllerWithConfiguration:
func (gc _GCVirtualControllerClass) VirtualControllerWithConfiguration(configuration unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(gc.class), objc.Sel("virtualControllerWithConfiguration:"), configuration)
	return rv
}

// Connects the virtual controller to the device and displays it on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/connect(replyHandler:)
func (g_ GCVirtualController) ConnectWithReplyHandler(reply unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("connectWithReplyHandler:"), reply)
}

// Changes the value of a button element in the virtual controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/setValue(_:forButtonElement:)
func (g_ GCVirtualController) SetValueForButtonElement(value float64, element string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setValue:forButtonElement:"), value, objc.String(element))
}

// The underlying controller object that you use to access input elements.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/controller
func (g_ GCVirtualController) Controller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controller"))
	return rv
}


