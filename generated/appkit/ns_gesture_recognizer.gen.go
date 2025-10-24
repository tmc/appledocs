// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	AllowedTouchTypes() unsafe.Pointer
	SetAllowedTouchTypes(value unsafe.Pointer)
	DelaysKeyEvents() bool
	SetDelaysKeyEvents(value bool)
	DelaysMagnificationEvents() bool
	SetDelaysMagnificationEvents(value bool)
	DelaysOtherMouseButtonEvents() bool
	SetDelaysOtherMouseButtonEvents(value bool)
	DelaysPrimaryMouseButtonEvents() bool
	SetDelaysPrimaryMouseButtonEvents(value bool)
	DelaysRotationEvents() bool
	SetDelaysRotationEvents(value bool)
	DelaysSecondaryMouseButtonEvents() bool
	SetDelaysSecondaryMouseButtonEvents(value bool)
	Delegate() GestureRecognizerDelegate /* not a class type */
	SetDelegate(value GestureRecognizerDelegate /* not a class type */)
	IsEnabled() bool
	SetIsEnabled(value bool)
	ModifierFlags() unsafe.Pointer
	SetModifierFlags(value unsafe.Pointer)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	PressureConfiguration() objc.IObject /* cross-framework: PressureConfiguration */
	SetPressureConfiguration(value objc.IObject /* cross-framework: PressureConfiguration */)
	State() objc.IObject /* cross-framework: State */
	SetState(value objc.IObject /* cross-framework: State */)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
	View() IView
	SetView(value IView)
	// methods:
}

// An object that monitors events and calls its action method when a predefined sequence of events occur.
//
// A gesture recognizer might recognize a single click, a click and drag, or a sequence of events that imply rotation. You do not create instances of this class directly. This class is an abstract base class that defines the common behavior for all gesture recognizers. When using a gesture recognizer in your app, create an instance of one of the concrete subclasses. The concrete subclasses of are the following: A gesture recognizer operates on events in a specific view (or in any of that view’s subviews). After creating a gesture recognizer, attach it to one of your views using the method. Events received by your app are forwarded automatically to any relevant gesture recognizers before they are sent to the corresponding view. The gesture recognizer can delay the further progression of the events until recognition is complete or allow the events to be delivered normally. A gesture recognizer can detect gestures that are either discrete or continuous in nature. A click gesture is discrete because it involves a mouse-down and mouse-up event without any mouse movements in between. By contrast, a pan or rotation gesture is continuous because it involves tracking mouse movements over a period of time. During the gesture recognition process, a gesture recognizer calls the action method of its associated target object to report the state of the recognition process. For discrete gestures, the action method is typically called only once when the gesture is recognized. For continuous gestures, it may be called multiple times depending on the current state of the gesture recognizer. In that situation, you can use your action method to perform appropriate tasks, such as creating animations for any mouse-related movements, in addition to handling the final results of the gesture recognition process. A gesture recognizer has only one action method and one target object, and the method must conform to one of the following signatures: When your code needs additional information about the particulars of a gesture, define your action method to include the gesture recognizer parameter. You almost always want the gesture recognizer object when handling continuous gestures. For example, for a rotation gesture, you would use the gesture recognizer object to get the updated rotation value. You can also use the gesture recognizer object to get the location of where the gesture occurred.


// An object that monitors events and calls its action method when a predefined sequence of events occur.
//
// [Full Topic]
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



// The action method to call when the gesture is recognized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/action
func (g_ GestureRecognizer) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("action"))
	return rv
}


// The action method to call when the gesture is recognized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/action
func (g_ GestureRecognizer) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (g_ GestureRecognizer) AllowedTouchTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (g_ GestureRecognizer) SetAllowedTouchTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}


// A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayskeyevents
func (g_ GestureRecognizer) DelaysKeyEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysKeyEvents"))
	return rv
}


// A Boolean value that indicates whether key events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayskeyevents
func (g_ GestureRecognizer) SetDelaysKeyEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysKeyEvents:"), value)
}


// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysmagnificationevents
func (g_ GestureRecognizer) DelaysMagnificationEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysMagnificationEvents"))
	return rv
}


// A Boolean value that indicates whether magnification events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysmagnificationevents
func (g_ GestureRecognizer) SetDelaysMagnificationEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysMagnificationEvents:"), value)
}


// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (g_ GestureRecognizer) DelaysOtherMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysOtherMouseButtonEvents"))
	return rv
}


// A Boolean value that indicates whether other mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysothermousebuttonevents
func (g_ GestureRecognizer) SetDelaysOtherMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysOtherMouseButtonEvents:"), value)
}


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (g_ GestureRecognizer) DelaysPrimaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysPrimaryMouseButtonEvents"))
	return rv
}


// A Boolean value that indicates whether primary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysprimarymousebuttonevents
func (g_ GestureRecognizer) SetDelaysPrimaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysPrimaryMouseButtonEvents:"), value)
}


// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysrotationevents
func (g_ GestureRecognizer) DelaysRotationEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysRotationEvents"))
	return rv
}


// A Boolean value that indicates whether rotation events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delaysrotationevents
func (g_ GestureRecognizer) SetDelaysRotationEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysRotationEvents:"), value)
}


// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (g_ GestureRecognizer) DelaysSecondaryMouseButtonEvents() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("delaysSecondaryMouseButtonEvents"))
	return rv
}


// A Boolean value that indicates whether secondary mouse button events are delivered only after gesture recognition fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delayssecondarymousebuttonevents
func (g_ GestureRecognizer) SetDelaysSecondaryMouseButtonEvents(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelaysSecondaryMouseButtonEvents:"), value)
}


// The delegate of the gesture recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delegate
func (g_ GestureRecognizer) Delegate() GestureRecognizerDelegate /* not a class type */ {
	rv := objc.Send[GestureRecognizerDelegate](g_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the gesture recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/delegate
func (g_ GestureRecognizer) SetDelegate(value GestureRecognizerDelegate /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}


// A Boolean value indicating whether the gesture recognizer is able to handle events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/isenabled
func (g_ GestureRecognizer) IsEnabled() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isEnabled"))
	return rv
}


// A Boolean value indicating whether the gesture recognizer is able to handle events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/isenabled
func (g_ GestureRecognizer) SetIsEnabled(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/modifierflags
func (g_ GestureRecognizer) ModifierFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("modifierFlags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/modifierflags
func (g_ GestureRecognizer) SetModifierFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setModifierFlags:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/name
func (g_ GestureRecognizer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/name
func (g_ GestureRecognizer) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setName:"), value)
}


// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/pressureconfiguration
func (g_ GestureRecognizer) PressureConfiguration() objc.IObject /* cross-framework: PressureConfiguration */ {
	rv := objc.Send[PressureConfiguration](g_.ID, objc.Sel("pressureConfiguration"))
	return rv
}


// Configures the behavior and progression of the Force Touch trackpad when responding to recognized pressure gestures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/pressureconfiguration
func (g_ GestureRecognizer) SetPressureConfiguration(value objc.IObject /* cross-framework: PressureConfiguration */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPressureConfiguration:"), value)
}


// The current state of the gesture recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/state-swift.property
func (g_ GestureRecognizer) State() objc.IObject /* cross-framework: State */ {
	rv := objc.Send[coreml.State](g_.ID, objc.Sel("state"))
	return rv
}


// The current state of the gesture recognizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/state-swift.property
func (g_ GestureRecognizer) SetState(value objc.IObject /* cross-framework: State */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setState:"), value)
}


// The object that implements the action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/target
func (g_ GestureRecognizer) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("target"))
	return rv
}


// The object that implements the action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/target
func (g_ GestureRecognizer) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setTarget:"), value)
}


// The view to which the gesture recognizer is attached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/view
func (g_ GestureRecognizer) View() IView {
	rv := objc.Send[View](g_.ID, objc.Sel("view"))
	return rv
}


// The view to which the gesture recognizer is attached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/view
func (g_ GestureRecognizer) SetView(value IView) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setView:"), value)
}



