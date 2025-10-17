// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AccessibilityElement] class.
var AccessibilityElementClass objc.Class

func init() {
	AccessibilityElementClass = objc.GetClass("NSAccessibilityElement")
}

type AccessibilityElement struct {
	objc.ID
}

func AccessibilityElementFrom(ptr unsafe.Pointer) AccessibilityElement {
	return AccessibilityElement{
		ID: objc.ID(ptr),
	}
}


// Instantiates and configures a new accessibility element. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (ac AccessibilityElement) AccessibilityElementWithRoleFrameLabelParent(role unsafe.Pointer, frame unsafe.Pointer, label unsafe.Pointer, parent objc.ID) objc.ID {
	sel := objc.RegisterName("accessibilityElementWithRole:frame:label:parent:")
	ret := objc.ID(AccessibilityElementClass).Send(sel, role, frame, label, parent)
	return objc.ID(ret)
}
// Adds a child to the accessibility element in the accessibility hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement unsafe.Pointer) {
	sel := objc.RegisterName("accessibilityAddChildElement:")
	a_.ID.Send(sel, childElement)
}
// Returns the accessibility element’s frame in its parent’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace()
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() {
	sel := objc.RegisterName("accessibilityFrameInParentSpace")
	a_.ID.Send(sel)
}

