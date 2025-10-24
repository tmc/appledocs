// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCEventInteraction */


/* debug [class_header]: Header for GCEventInteraction */
// The class instance for the [GCEventInteraction] class.
var (
	GCEventInteractionClass     _GCEventInteractionClass
	GCEventInteractionClassOnce sync.Once
)

func getGCEventInteractionClass() _GCEventInteractionClass {
	GCEventInteractionClassOnce.Do(func() {
		GCEventInteractionClass = _GCEventInteractionClass{objc.GetClass("GCEventInteraction")}
	})
	return GCEventInteractionClass
}

type _GCEventInteractionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCEventInteraction */
// An interface definition for the [GCEventInteraction] class.
type IGCEventInteraction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCEventInteraction */
	// properties:
	ControllerPausedHandler() unsafe.Pointer
	SetControllerPausedHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCEventInteraction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCEventInteraction */
// Alloc allocates a new instance without initialization.
func (gc _GCEventInteractionClass) Alloc() GCEventInteraction {
	rv := objc.Send[GCEventInteraction](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCEventInteractionClass) New() GCEventInteraction {
	rv := objc.Send[GCEventInteraction](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCEventInteraction) Init() GCEventInteraction {
	rv := objc.Send[GCEventInteraction](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCEventInteraction) Autorelease() GCEventInteraction {
	rv := objc.Send[GCEventInteraction](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCEventInteraction creates a new GCEventInteraction instance.
func NewGCEventInteraction() GCEventInteraction {
	return getGCEventInteractionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCEventInteraction */
// An interaction that indicates the view’s intent to receive game controller events through the Game Controller framework.
//
// On visionOS, users can interact with your app using a game controller. By default, the system converts game controller actions into pinch events and sends them to the view the user is gazing at, its gesture recognizers, and then up the responder chain. If you use the Game Controller framework to handle game controller events for part of your user interface, add an instance of to the root of that part of your app’s view hierarchy. For example, if you are writing a game using Metal, add this interaction to the view that hosts your game’s .


// An interaction that indicates the view’s intent to receive game controller events through the Game Controller framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventInteraction
type GCEventInteraction struct {
	objectivec.Object
}

// GCEventInteractionFrom constructs a [GCEventInteraction] from an unsafe.Pointer.
//
// An interaction that indicates the view’s intent to receive game controller events through the Game Controller framework.
func GCEventInteractionFrom(ptr unsafe.Pointer) GCEventInteraction {
	return GCEventInteraction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCEventInteraction */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCEventInteraction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCEventInteraction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCEventInteraction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCEventInteraction */

// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCEventInteraction) ControllerPausedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controllerPausedHandler"))
	return rv
}/* debug [instance_properties/getter]: controllerPausedHandler */


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCEventInteraction) SetControllerPausedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setControllerPausedHandler:"), value)
}/* debug [instance_properties/setter]: controllerPausedHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCEventInteraction */


