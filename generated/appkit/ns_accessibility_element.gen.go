// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAccessibilityElement */


/* debug [class_header]: Header for NSAccessibilityElement */
// The class instance for the [AccessibilityElement] class.
var (
	AccessibilityElementClass     _AccessibilityElementClass
	AccessibilityElementClassOnce sync.Once
)

func getAccessibilityElementClass() _AccessibilityElementClass {
	AccessibilityElementClassOnce.Do(func() {
		AccessibilityElementClass = _AccessibilityElementClass{objc.GetClass("NSAccessibilityElement")}
	})
	return AccessibilityElementClass
}

type _AccessibilityElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccessibilityElement */
// An interface definition for the [AccessibilityElement] class.
type IAccessibilityElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccessibilityElement */
	// properties:
	AccessibilityFrameInParentSpace() Rect /* not a class type */
	SetAccessibilityFrameInParentSpace(value Rect /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccessibilityElement */
	// methods:
	AccessibilityAddChildElement(childElement IAccessibilityElement)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccessibilityElement */
// Alloc allocates a new instance without initialization.
func (ac _AccessibilityElementClass) Alloc() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccessibilityElementClass) New() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityElement) Init() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityElement) Autorelease() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityElement creates a new AccessibilityElement instance.
func NewAccessibilityElement() AccessibilityElement {
	return getAccessibilityElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccessibilityElement */
// The basic infrastructure necessary for interacting with an assistive app.
//
// Create subclasses of the class to represent any of your user interface elements that don’t inherit from or from one of the standard AppKit controls. This class represents your user interface element in the accessibility hierarchy and manages the details necessary for working with assistive apps. To support accessibility features for a custom user interface element: Create your subclass by using . You can also set these values using , and . Call the parent’s method to add your subclass. You can also add the subclass to its parent’s array using . In your subclass, call . This ensures that your control moves with its superview. In your subclass, adopt a role-specific protocol, customize the role, and post notifications just as you would handle any other accessible control. See . In your subclass, implement any additional properties and methods you may need to use to further customize your user interface element’s accessibility behavior. See .


// The basic infrastructure necessary for interacting with an assistive app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class
type AccessibilityElement struct {
	objectivec.Object
}

// AccessibilityElementFrom constructs a [AccessibilityElement] from an unsafe.Pointer.
//
// The basic infrastructure necessary for interacting with an assistive app.
func AccessibilityElementFrom(ptr unsafe.Pointer) AccessibilityElement {
	return AccessibilityElement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccessibilityElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccessibilityElement */

// Instantiates and configures a new accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (ac _AccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role AccessibilityRole /* typedef */, frame Rect /* not a class type */, label objc.IObject /* cross-framework: NSString */, parent objc.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("accessibilityElementWithRole:frame:label:parent:"), role, frame, label, parent)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AccessibilityElementWithRoleFrameLabelParent) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccessibilityElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccessibilityElement */

// Adds a child to the accessibility element in the accessibility hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement IAccessibilityElement) {
	objc.Send[objc.ID](a_.ID, objc.Sel("accessibilityAddChildElement:"), childElement)
}/* debug [instance_methods/method]: AccessibilityAddChildElement */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccessibilityElement */

// The accessibility element’s frame in its parent’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() Rect /* not a class type */ {
	rv := objc.Send[Rect](a_.ID, objc.Sel("accessibilityFrameInParentSpace"))
	return rv
}/* debug [instance_properties/getter]: accessibilityFrameInParentSpace */


// The accessibility element’s frame in its parent’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
func (a_ AccessibilityElement) SetAccessibilityFrameInParentSpace(value Rect /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessibilityFrameInParentSpace:"), value)
}/* debug [instance_properties/setter]: accessibilityFrameInParentSpace */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAccessibilityElement */



