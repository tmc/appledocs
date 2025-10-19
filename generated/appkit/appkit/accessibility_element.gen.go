// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityElement] class.
var (
	accessibilityElementClass     _AccessibilityElementClass
	accessibilityElementClassOnce sync.Once
)

func getAccessibilityElementClass() _AccessibilityElementClass {
	accessibilityElementClassOnce.Do(func() {
		accessibilityElementClass = _AccessibilityElementClass{objc.GetClass("NSAccessibilityElement")}
	})
	return accessibilityElementClass
}

type _AccessibilityElementClass struct {
	class objc.Class
}

// An interface definition for the [AccessibilityElement] class.
type IAccessibilityElement interface {
	objectivec.IObject
	AccessibilityAddChildElement(childElement unsafe.Pointer)
	AccessibilityFrameInParentSpace()
}

// The basic infrastructure necessary for interacting with an assistive app. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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


// Instantiates and configures a new accessibility element. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (ac _AccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role unsafe.Pointer, frame unsafe.Pointer, label string, parent objc.ID) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("accessibilityElementWithRole:frame:label:parent:"), role, frame, objc.String(label), parent)
	return rv
}
// Adds a child to the accessibility element in the accessibility hierarchy. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("accessibilityAddChildElement:"), childElement)
}
// Returns the accessibility element’s frame in its parent’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace()
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() {
	objc.Send[objc.ID](a_.ID, objc.Sel("accessibilityFrameInParentSpace"))
}


