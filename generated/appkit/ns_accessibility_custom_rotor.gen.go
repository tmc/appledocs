// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAccessibilityCustomRotor */


/* debug [class_header]: Header for NSAccessibilityCustomRotor */
// The class instance for the [AccessibilityCustomRotor] class.
var (
	AccessibilityCustomRotorClass     _AccessibilityCustomRotorClass
	AccessibilityCustomRotorClassOnce sync.Once
)

func getAccessibilityCustomRotorClass() _AccessibilityCustomRotorClass {
	AccessibilityCustomRotorClassOnce.Do(func() {
		AccessibilityCustomRotorClass = _AccessibilityCustomRotorClass{objc.GetClass("NSAccessibilityCustomRotor")}
	})
	return AccessibilityCustomRotorClass
}

type _AccessibilityCustomRotorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccessibilityCustomRotor */
// An interface definition for the [AccessibilityCustomRotor] class.
type IAccessibilityCustomRotor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccessibilityCustomRotor */
	// properties:
	ItemLoadingDelegate() unsafe.Pointer
	SetItemLoadingDelegate(value unsafe.Pointer)
	ItemSearchDelegate() unsafe.Pointer
	SetItemSearchDelegate(value unsafe.Pointer)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Type() AccessibilityCustomRotorType
	SetType(value AccessibilityCustomRotorType)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccessibilityCustomRotor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccessibilityCustomRotor */
// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorClass) Alloc() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccessibilityCustomRotorClass) New() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityCustomRotor) Init() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityCustomRotor) Autorelease() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityCustomRotor creates a new AccessibilityCustomRotor instance.
func NewAccessibilityCustomRotor() AccessibilityCustomRotor {
	return getAccessibilityCustomRotorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccessibilityCustomRotor */
// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element.
//
// Assistive apps, like VoiceOver, provide interfaces to quickly search apps for content of a specific type. For example, in a web browser, a user can quickly explore a list of navigational links or buttons using VoiceOver’s content menus. provides a way for apps to vend their own content menus. For example, Pages can create a custom rotor that allows assistive apps to search the Pages document for all headings.


// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor
type AccessibilityCustomRotor struct {
	objectivec.Object
}

// AccessibilityCustomRotorFrom constructs a [AccessibilityCustomRotor] from an unsafe.Pointer.
//
// A context-sensitive function that helps VoiceOver users find the next instance of a related accessibility element.
func AccessibilityCustomRotorFrom(ptr unsafe.Pointer) AccessibilityCustomRotor {
	return AccessibilityCustomRotor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccessibilityCustomRotor */

// Creates a custom rotor with the specified label and item search delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(label:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithLabelItemSearchDelegate(label objc.IObject /* cross-framework: NSString */, itemSearchDelegate unsafe.Pointer) AccessibilityCustomRotor {
	instance := getAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotor](instance.ID, objc.Sel("initWithLabel:itemSearchDelegate:"), label, itemSearchDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccessibilityCustomRotorWithLabelItemSearchDelegate */


// Creates a custom rotor with the specified rotor type and item search delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(rotorType:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithRotorTypeItemSearchDelegate(rotorType AccessibilityCustomRotorType, itemSearchDelegate unsafe.Pointer) AccessibilityCustomRotor {
	instance := getAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotor](instance.ID, objc.Sel("initWithRotorType:itemSearchDelegate:"), rotorType, itemSearchDelegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccessibilityCustomRotorWithRotorTypeItemSearchDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccessibilityCustomRotor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccessibilityCustomRotor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccessibilityCustomRotor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccessibilityCustomRotor */

// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemLoadingDelegate
func (a_ AccessibilityCustomRotor) ItemLoadingDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("itemLoadingDelegate"))
	return rv
}/* debug [instance_properties/getter]: itemLoadingDelegate */


// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemLoadingDelegate
func (a_ AccessibilityCustomRotor) SetItemLoadingDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemLoadingDelegate:"), value)
}/* debug [instance_properties/setter]: itemLoadingDelegate */


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemSearchDelegate
func (a_ AccessibilityCustomRotor) ItemSearchDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("itemSearchDelegate"))
	return rv
}/* debug [instance_properties/getter]: itemSearchDelegate */


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemSearchDelegate
func (a_ AccessibilityCustomRotor) SetItemSearchDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemSearchDelegate:"), value)
}/* debug [instance_properties/setter]: itemSearchDelegate */


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/label
func (a_ AccessibilityCustomRotor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/label
func (a_ AccessibilityCustomRotor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/type
func (a_ AccessibilityCustomRotor) Type() AccessibilityCustomRotorType {
	rv := objc.Send[AccessibilityCustomRotorType](a_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/type
func (a_ AccessibilityCustomRotor) SetType(value AccessibilityCustomRotorType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), value)
}/* debug [instance_properties/setter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAccessibilityCustomRotor */


