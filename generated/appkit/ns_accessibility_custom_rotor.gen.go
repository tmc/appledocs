// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ItemLoadingDelegate() unsafe.Pointer
	SetItemLoadingDelegate(value unsafe.Pointer)
	ItemSearchDelegate() unsafe.Pointer
	SetItemSearchDelegate(value unsafe.Pointer)
	Label() string
	SetLabel(value string)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
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



// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/itemloadingdelegate
func (a_ AccessibilityCustomRotor) ItemLoadingDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("itemLoadingDelegate"))
	return rv
}


// The delegate for loading item results that don’t have a backing UI element at loading time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/itemloadingdelegate
func (a_ AccessibilityCustomRotor) SetItemLoadingDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemLoadingDelegate:"), value)
}


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/itemsearchdelegate
func (a_ AccessibilityCustomRotor) ItemSearchDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("itemSearchDelegate"))
	return rv
}


// The delegate for finding the next item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/itemsearchdelegate
func (a_ AccessibilityCustomRotor) SetItemSearchDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setItemSearchDelegate:"), value)
}


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/label
func (a_ AccessibilityCustomRotor) Label() string {
	rv := objc.Send[string](a_.ID, objc.Sel("label"))
	return rv
}


// The localized label that assistive apps use to describe the custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/label
func (a_ AccessibilityCustomRotor) SetLabel(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), objc.String(value))
}


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/type
func (a_ AccessibilityCustomRotor) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("type"))
	return rv
}


// The type of content that the rotor represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/type
func (a_ AccessibilityCustomRotor) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setType:"), value)
}



