// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityCustomAction] class.
var (
	AccessibilityCustomActionClass     _AccessibilityCustomActionClass
	AccessibilityCustomActionClassOnce sync.Once
)

func getAccessibilityCustomActionClass() _AccessibilityCustomActionClass {
	AccessibilityCustomActionClassOnce.Do(func() {
		AccessibilityCustomActionClass = _AccessibilityCustomActionClass{objc.GetClass("NSAccessibilityCustomAction")}
	})
	return AccessibilityCustomActionClass
}

type _AccessibilityCustomActionClass struct {
	class objc.Class
}

// An interface definition for the [AccessibilityCustomAction] class.
type IAccessibilityCustomAction interface {
	objectivec.IObject
	// properties:
	Handler() unsafe.Pointer
	SetHandler(value unsafe.Pointer)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Selector() objc.SEL
	SetSelector(value objc.SEL)
	Target() objc.ID
	SetTarget(value objc.ID)
	// methods:
}

// A custom action to perform on an accessible object.
//
// Apps that support custom actions can create instances of this class, specifying the user-readable name of the action, and either a handler closure or the object and selector to use when performing the action. Assistive apps display custom actions in response to specific user cues. For example, VoiceOver lets users access actions quickly using the Actions rotor. After creating an instance of this class, add it to the property of an appropriate accessible object.


// A custom action to perform on an accessible object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction
type AccessibilityCustomAction struct {
	objectivec.Object
}

// AccessibilityCustomActionFrom constructs a [AccessibilityCustomAction] from an unsafe.Pointer.
//
// A custom action to perform on an accessible object.
func AccessibilityCustomActionFrom(ptr unsafe.Pointer) AccessibilityCustomAction {
	return AccessibilityCustomAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomActionClass) Alloc() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccessibilityCustomActionClass) New() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityCustomAction) Init() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityCustomAction) Autorelease() AccessibilityCustomAction {
	rv := objc.Send[AccessibilityCustomAction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityCustomAction creates a new AccessibilityCustomAction instance.
func NewAccessibilityCustomAction() AccessibilityCustomAction {
	return getAccessibilityCustomActionClass().New()
}



// Creates a custom action object with the specified name and handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:handler:)
func NewAccessibilityCustomActionWithNameHandler(name objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) AccessibilityCustomAction {
	instance := getAccessibilityCustomActionClass().Alloc()
	rv := objc.Send[AccessibilityCustomAction](instance.ID, objc.Sel("initWithName:handler:"), name, handler)
	rv.Autorelease()
	return rv
}


// Creates a custom action object with the specified name, target, and selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/init(name:target:selector:)
func NewAccessibilityCustomActionWithNameTargetSelector(name objc.IObject /* cross-framework: NSString */, target objc.IObject, selector objc.SEL) AccessibilityCustomAction {
	instance := getAccessibilityCustomActionClass().Alloc()
	rv := objc.Send[AccessibilityCustomAction](instance.ID, objc.Sel("initWithName:target:selector:"), name, target, selector)
	rv.Autorelease()
	return rv
}



// The closure that handles the execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/handler
func (a_ AccessibilityCustomAction) Handler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("handler"))
	return rv
}


// The closure that handles the execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/handler
func (a_ AccessibilityCustomAction) SetHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHandler:"), value)
}


// A localized name that describes the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/name
func (a_ AccessibilityCustomAction) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}


// A localized name that describes the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/name
func (a_ AccessibilityCustomAction) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}


// The method to call on the target to perform the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/selector
func (a_ AccessibilityCustomAction) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](a_.ID, objc.Sel("selector"))
	return rv
}


// The method to call on the target to perform the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/selector
func (a_ AccessibilityCustomAction) SetSelector(value objc.SEL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSelector:"), value)
}


// The object that performs the action through a selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/target
func (a_ AccessibilityCustomAction) Target() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("target"))
	return rv
}


// The object that performs the action through a selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomAction/target
func (a_ AccessibilityCustomAction) SetTarget(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTarget:"), value)
}


