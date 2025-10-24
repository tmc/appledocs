// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AccessibilityCustomRotor] class.
type IAccessibilityCustomRotor interface {
	objectivec.IObject
	// properties:
	ItemLoadingDelegate() objc.ID
	SetItemLoadingDelegate(value objc.ID)
	ItemSearchDelegate() objc.ID
	SetItemSearchDelegate(value objc.ID)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Type() AccessibilityCustomRotorType
	SetType(value AccessibilityCustomRotorType)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorClass) Alloc() AccessibilityCustomRotor {
	rv := objc.Send[AccessibilityCustomRotor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a custom rotor with the specified label and item search delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(label:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithLabelItemSearchDelegate(label objc.IObject /* cross-framework: NSString */, itemSearchDelegate objc.IObject) AccessibilityCustomRotor {
	instance := getAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotor](instance.ID, objc.Sel("initWithLabel:itemSearchDelegate:"), label, itemSearchDelegate)
	rv.Autorelease()
	return rv
}


// Creates a custom rotor with the specified rotor type and item search delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/init(rotorType:itemSearchDelegate:)
func NewAccessibilityCustomRotorWithRotorTypeItemSearchDelegate(rotorType AccessibilityCustomRotorType, itemSearchDelegate objc.IObject) AccessibilityCustomRotor {
	instance := getAccessibilityCustomRotorClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotor](instance.ID, objc.Sel("initWithRotorType:itemSearchDelegate:"), rotorType, itemSearchDelegate)
	rv.Autorelease()
	return rv
}



// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemLoadingDelegate
func (a_ AccessibilityCustomRotor) ItemLoadingDelegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("itemLoadingDelegate"))
	return rv
}


// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemLoadingDelegate
func (a_ AccessibilityCustomRotor) SetItemLoadingDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemLoadingDelegate:"), value)
}


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemSearchDelegate
func (a_ AccessibilityCustomRotor) ItemSearchDelegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("itemSearchDelegate"))
	return rv
}


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/itemSearchDelegate
func (a_ AccessibilityCustomRotor) SetItemSearchDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemSearchDelegate:"), value)
}


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/label
func (a_ AccessibilityCustomRotor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/label
func (a_ AccessibilityCustomRotor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/type
func (a_ AccessibilityCustomRotor) Type() AccessibilityCustomRotorType {
	rv := objc.Send[AccessibilityCustomRotorType](a_.ID, objc.Sel("type"))
	return rv
}


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/type
func (a_ AccessibilityCustomRotor) SetType(value AccessibilityCustomRotorType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), value)
}


