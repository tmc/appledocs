// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GCEventInteraction] class.
type IGCEventInteraction interface {
	objectivec.IObject
	HandledEventTypes() unsafe.Pointer
	SetHandledEventTypes(value unsafe.Pointer)
	ControllerPausedHandler() unsafe.Pointer
	SetControllerPausedHandler(value unsafe.Pointer)
	ReceivesEventsInView() bool
	SetReceivesEventsInView(value bool)
}

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

// Alloc allocates a new instance without initialization.
func (gc _GCEventInteractionClass) Alloc() GCEventInteraction {
	rv := objc.Send[GCEventInteraction](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The types of game controller events that should be delivered through the Game Controller framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventInteraction/handledEventTypes
func (g_ GCEventInteraction) HandledEventTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("handledEventTypes"))
	return rv
}


// The types of game controller events that should be delivered through the Game Controller framework.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEventInteraction/handledEventTypes
func (g_ GCEventInteraction) SetHandledEventTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setHandledEventTypes:"), value)
}


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCEventInteraction) ControllerPausedHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("controllerPausedHandler"))
	return rv
}


// The block that the framework calls when the user presses the pause button on the controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontroller/controllerpausedhandler
func (g_ GCEventInteraction) SetControllerPausedHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setControllerPausedHandler:"), value)
}


// A Boolean value that determines whether events are delivered exclusively
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gceventinteraction/receiveseventsinview
func (g_ GCEventInteraction) ReceivesEventsInView() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("receivesEventsInView"))
	return rv
}


// A Boolean value that determines whether events are delivered exclusively
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gceventinteraction/receiveseventsinview
func (g_ GCEventInteraction) SetReceivesEventsInView(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setReceivesEventsInView:"), value)
}



