// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCVirtualController */


/* debug [class_header]: Header for GCVirtualController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCVirtualController */
// An interface definition for the [GCVirtualController] class.
type IGCVirtualController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCVirtualController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCVirtualController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCVirtualController */
// Alloc allocates a new instance without initialization.
func (gc _GCVirtualControllerClass) Alloc() GCVirtualController {
	rv := objc.Send[GCVirtualController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCVirtualController */
// A software emulation of a real controller that you configure specifically for your game.
//
// Use a virtual controller to display software controls that you can customize over your game. You create a virtual controller from a configuration where you choose the input elements to display. You can even customize the images for the elements. When you connect the controller to the device, users interact with it similarly to a real controller. To add a virtual controller to your game, create a object containing the elements you want to appear in the controller. Then create the virtual controller by passing the configuration to the method. Use the method to display the virtual controller on the screen. To customize an element in the virtual controller, pass a new object for the element to the method. You process input from a virtual controller similarly to a real controller. Use the property to get the underlying object. You can either poll the elements of the controller object or set the element’s handlers to get callbacks when their input values change.


// A software emulation of a real controller that you configure specifically for your game.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCVirtualController */

// Creates a new virtual controller using the configuration you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/init(configuration:)
func NewGCVirtualControllerWithConfiguration(configuration IGCVirtualControllerConfiguration) GCVirtualController {
	instance := getGCVirtualControllerClass().Alloc()
	rv := objc.Send[GCVirtualController](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCVirtualControllerWithConfiguration */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCVirtualController */

// Creates a new virtual controller using the configuration you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCVirtualController/virtualControllerWithConfiguration:
func (gc _GCVirtualControllerClass) VirtualControllerWithConfiguration(configuration IGCVirtualControllerConfiguration) GCVirtualController {
	rv := objc.Send[GCVirtualController](objc.ID(gc.class), objc.Sel("virtualControllerWithConfiguration:"), configuration)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VirtualControllerWithConfiguration) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCVirtualController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCVirtualController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCVirtualController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCVirtualController */


