// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GestureRecognizer] class.
var (
	GestureRecognizerClass     _GestureRecognizerClass
	GestureRecognizerClassOnce sync.Once
)

func getGestureRecognizerClass() _GestureRecognizerClass {
	GestureRecognizerClassOnce.Do(func() {
		GestureRecognizerClass = _GestureRecognizerClass{objc.GetClass("NSGestureRecognizer")}
	})
	return GestureRecognizerClass
}

type _GestureRecognizerClass struct {
	class objc.Class
}

// An interface definition for the [GestureRecognizer] class.
type IGestureRecognizer interface {
	objectivec.IObject
	LocationInView(view IView) coregraphics.CGPoint
	TouchesBeganWithEvent(event IEvent)
	TouchesCancelledWithEvent(event IEvent)
	TouchesEndedWithEvent(event IEvent)
	TouchesMovedWithEvent(event IEvent)
}

// An object that monitors events and calls its action method when a predefined sequence of events occur.
//
// A gesture recognizer might recognize a single click, a click and drag, or a sequence of events that imply rotation. You do not create instances of this class directly. This class is an abstract base class that defines the common behavior for all gesture recognizers. When using a gesture recognizer in your app, create an instance of one of the concrete subclasses. The concrete subclasses of are the following: A gesture recognizer operates on events in a specific view (or in any of that view’s subviews). After creating a gesture recognizer, attach it to one of your views using the method. Events received by your app are forwarded automatically to any relevant gesture recognizers before they are sent to the corresponding view. The gesture recognizer can delay the further progression of the events until recognition is complete or allow the events to be delivered normally. A gesture recognizer can detect gestures that are either discrete or continuous in nature. A click gesture is discrete because it involves a mouse-down and mouse-up event without any mouse movements in between. By contrast, a pan or rotation gesture is continuous because it involves tracking mouse movements over a period of time. During the gesture recognition process, a gesture recognizer calls the action method of its associated target object to report the state of the recognition process. For discrete gestures, the action method is typically called only once when the gesture is recognized. For continuous gestures, it may be called multiple times depending on the current state of the gesture recognizer. In that situation, you can use your action method to perform appropriate tasks, such as creating animations for any mouse-related movements, in addition to handling the final results of the gesture recognition process. A gesture recognizer has only one action method and one target object, and the method must conform to one of the following signatures: When your code needs additional information about the particulars of a gesture, define your action method to include the gesture recognizer parameter. You almost always want the gesture recognizer object when handling continuous gestures. For example, for a rotation gesture, you would use the gesture recognizer object to get the updated rotation value. You can also use the gesture recognizer object to get the location of where the gesture occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer
type GestureRecognizer struct {
	objectivec.Object
}

// GestureRecognizerFrom constructs a [GestureRecognizer] from an unsafe.Pointer.
//
// An object that monitors events and calls its action method when a predefined sequence of events occur.
func GestureRecognizerFrom(ptr unsafe.Pointer) GestureRecognizer {
	return GestureRecognizer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GestureRecognizerClass) Alloc() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GestureRecognizerClass) New() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GestureRecognizer) Init() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GestureRecognizer) Autorelease() GestureRecognizer {
	rv := objc.Send[GestureRecognizer](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGestureRecognizer creates a new GestureRecognizer instance.
func NewGestureRecognizer() GestureRecognizer {
	return getGestureRecognizerClass().New()
}


// Returns the point computed as the location of the gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/location(in:)
func (g_ GestureRecognizer) LocationInView(view IView) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](g_.ID, objc.Sel("locationInView:"), view)
	return rv
}

// Called when one or more fingers first make contact with an instance on the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesBegan(with:)
func (g_ GestureRecognizer) TouchesBeganWithEvent(event IEvent) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesBeganWithEvent:"), event)
}

// Called when a system event, such as a low-memory warning, cancels an in-progress touch event in an object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesCancelled(with:)
func (g_ GestureRecognizer) TouchesCancelledWithEvent(event IEvent) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesCancelledWithEvent:"), event)
}

// Called when one or more fingers are removed from contact with an instance on the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesEnded(with:)
func (g_ GestureRecognizer) TouchesEndedWithEvent(event IEvent) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesEndedWithEvent:"), event)
}

// Called when one or more fingers, associated with an in-progress event, move within an instance on the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/touchesMoved(with:)
func (g_ GestureRecognizer) TouchesMovedWithEvent(event IEvent) {
	objc.Send[objc.ID](g_.ID, objc.Sel("touchesMovedWithEvent:"), event)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/allowedTouchTypes
func (g_ GestureRecognizer) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Send[TouchTypeMask](g_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}


// SetAllowedTouchTypes sets the value of the allowedTouchTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/allowedTouchTypes
func (g_ GestureRecognizer) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}

// A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysKeyEvents
func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysKeyEvents"))
	return rv
}


// SetDelaysKeyEvents sets the value of the delaysKeyEvents property.
// A Boolean value that indicates whether key events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysKeyEvents
func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysKeyEvents:"), value)
}

// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysMagnificationEvents
func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysMagnificationEvents"))
	return rv
}


// SetDelaysMagnificationEvents sets the value of the delaysMagnificationEvents property.
// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysMagnificationEvents
func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysMagnificationEvents:"), value)
}

// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysOtherMouseButtonEvents
func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}


// SetDelaysOtherMouseButtonEvents sets the value of the delaysOtherMouseButtonEvents property.
// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysOtherMouseButtonEvents
func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysPrimaryMouseButtonEvents
func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}


// SetDelaysPrimaryMouseButtonEvents sets the value of the delaysPrimaryMouseButtonEvents property.
// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysPrimaryMouseButtonEvents
func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}

// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysRotationEvents
func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysRotationEvents"))
	return rv
}


// SetDelaysRotationEvents sets the value of the delaysRotationEvents property.
// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysRotationEvents
func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysRotationEvents:"), value)
}

// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysSecondaryMouseButtonEvents
func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}


// SetDelaysSecondaryMouseButtonEvents sets the value of the delaysSecondaryMouseButtonEvents property.
// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/delaysSecondaryMouseButtonEvents
func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}

// The action method to call when the gesture is recognized.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/action
func (g_ GestureRecognizer) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// The action method to call when the gesture is recognized.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/action
func (g_ GestureRecognizer) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAction:"), value)
}

// The delegate of the gesture recognizer.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delegate
func (g_ GestureRecognizer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the gesture recognizer.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delegate
func (g_ GestureRecognizer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether the gesture recognizer is able to handle events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/isenabled
func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value indicating whether the gesture recognizer is able to handle events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/isenabled
func (g_ GestureRecognizer) SetIsEnabled(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/modifierflags
func (g_ GestureRecognizer) ModifierFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("modifierFlags"))
	return rv
}


// SetModifierFlags sets the value of the modifierFlags property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/modifierflags
func (g_ GestureRecognizer) SetModifierFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModifierFlags:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/name
func (g_ GestureRecognizer) Name() string {
	rv := objc.Send[string](g_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/name
func (g_ GestureRecognizer) SetName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setName:"), objc.String(value))
}

// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/pressureconfiguration
func (g_ GestureRecognizer) PressureConfiguration() NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](g_.ID, objc.Sel("pressureConfiguration"))
	return rv
}


// SetPressureConfiguration sets the value of the pressureConfiguration property.
// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/pressureconfiguration
func (g_ GestureRecognizer) SetPressureConfiguration(value IPressureConfiguration) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPressureConfiguration:"), value)
}

// The current state of the gesture recognizer.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/state-swift.property
func (g_ GestureRecognizer) State() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("state"))
	return rv
}


// SetState sets the value of the state property.
// The current state of the gesture recognizer.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/state-swift.property
func (g_ GestureRecognizer) SetState(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setState:"), value)
}

// The object that implements the action method.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/target
func (g_ GestureRecognizer) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
// The object that implements the action method.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/target
func (g_ GestureRecognizer) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTarget:"), value)
}

// The view to which the gesture recognizer is attached.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/view
func (g_ GestureRecognizer) View() NSView {
	rv := objc.Send[NSView](g_.ID, objc.Sel("view"))
	return rv
}


// SetView sets the value of the view property.
// The view to which the gesture recognizer is attached.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/view
func (g_ GestureRecognizer) SetView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setView:"), value)
}



