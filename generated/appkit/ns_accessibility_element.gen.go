// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

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

// An interface definition for the [AccessibilityElement] class.
type IAccessibilityElement interface {
	objectivec.IObject
	AccessibilityAddChildElement(childElement unsafe.Pointer)
}

// The basic infrastructure necessary for interacting with an assistive app.
//
// Create subclasses of the class to represent any of your user interface elements that don’t inherit from or from one of the standard AppKit controls. This class represents your user interface element in the accessibility hierarchy and manages the details necessary for working with assistive apps. To support accessibility features for a custom user interface element: Create your subclass by using . You can also set these values using , and . Call the parent’s method to add your subclass. You can also add the subclass to its parent’s array using . In your subclass, call . This ensures that your control moves with its superview. In your subclass, adopt a role-specific protocol, customize the role, and post notifications just as you would handle any other accessible control. See . In your subclass, implement any additional properties and methods you may need to use to further customize your user interface element’s accessibility behavior. See .
//
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

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityElementClass) Alloc() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Instantiates and configures a new accessibility element.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (ac _AccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role unsafe.Pointer, frame coregraphics.CGRect, label string, parent objc.ID) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("accessibilityElementWithRole:frame:label:parent:"), role, frame, objc.String(label), parent)
	return rv
}

// Adds a child to the accessibility element in the accessibility hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("accessibilityAddChildElement:"), childElement)
}

// The accessibility element’s frame in its parent’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](a_.ID, objc.Sel("accessibilityFrameInParentSpace"))
	return rv
}


// SetAccessibilityFrameInParentSpace sets the value of the accessibilityFrameInParentSpace property.
// The accessibility element’s frame in its parent’s coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace
func (a_ AccessibilityElement) SetAccessibilityFrameInParentSpace(value coregraphics.CGRect) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccessibilityFrameInParentSpace:"), value)
}



