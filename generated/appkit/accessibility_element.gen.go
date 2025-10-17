
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [AccessibilityElement] class.
var AccessibilityElementClass _AccessibilityElementClass

func init() {
	AccessibilityElementClass = _AccessibilityElementClass{objc.GetClass("NSAccessibilityElement")}
}

type _AccessibilityElementClass struct {
	objc.Class
}

// An interface definition for the [AccessibilityElement] class.
type IAccessibilityElement interface {
	ID() objc.ID
	AccessibilityAddChildElement(childElement unsafe.Pointer)
	AccessibilityFrameInParentSpace()
}

type AccessibilityElement struct {
	id objc.ID
}

func AccessibilityElementFrom(ptr unsafe.Pointer) AccessibilityElement {
	return AccessibilityElement{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ AccessibilityElement) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityElementClass) Alloc() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AccessibilityElementClass) New() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAccessibilityElement creates and returns a new initialized instance.
func NewAccessibilityElement() AccessibilityElement {
	return AccessibilityElementClass.New()
}

// Init initializes the instance.
func (a_ AccessibilityElement) Init() AccessibilityElement {
	rv := objc.Send[AccessibilityElement](a_.ID(), selInit)
	return rv
}
// Instantiates and configures a new accessibility element. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func (ac _AccessibilityElementClass) AccessibilityElementWithRoleFrameLabelParent(role unsafe.Pointer, frame foundation.Rect, label string, parent objc.ID) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.Class), objc.RegisterName("accessibilityElementWithRole:frame:label:parent:"), role, frame, label, parent)
	return rv
}

// AccessibilityElement_AccessibilityElementWithRoleFrameLabelParent creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/element(withRole:frame:label:parent:)
func AccessibilityElement_AccessibilityElementWithRoleFrameLabelParent(role unsafe.Pointer, frame foundation.Rect, label string, parent objc.ID) objc.ID {
	return AccessibilityElementClass.AccessibilityElementWithRoleFrameLabelParent(role, frame, label, parent)
}
// Adds a child to the accessibility element in the accessibility hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityAddChildElement(_:)
func (a_ AccessibilityElement) AccessibilityAddChildElement(childElement unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("accessibilityAddChildElement:"), childElement)
}
// Returns the accessibility element’s frame in its parent’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityElement-swift.class/accessibilityFrameInParentSpace()
func (a_ AccessibilityElement) AccessibilityFrameInParentSpace() {
	objc.Send[objc.ID](a_.ID(), objc.RegisterName("accessibilityFrameInParentSpace"))
}
