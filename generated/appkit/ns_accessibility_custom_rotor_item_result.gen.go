// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccessibilityCustomRotorItemResult] class.
var (
	AccessibilityCustomRotorItemResultClass     _AccessibilityCustomRotorItemResultClass
	AccessibilityCustomRotorItemResultClassOnce sync.Once
)

func getAccessibilityCustomRotorItemResultClass() _AccessibilityCustomRotorItemResultClass {
	AccessibilityCustomRotorItemResultClassOnce.Do(func() {
		AccessibilityCustomRotorItemResultClass = _AccessibilityCustomRotorItemResultClass{objc.GetClass("NSAccessibilityCustomRotorItemResult")}
	})
	return AccessibilityCustomRotorItemResultClass
}

type _AccessibilityCustomRotorItemResultClass struct {
	class objc.Class
}

// An interface definition for the [AccessibilityCustomRotorItemResult] class.
type IAccessibilityCustomRotorItemResult interface {
	objectivec.IObject
	// properties:
	CustomLabel() objc.IObject /* cross-framework: NSString */
	SetCustomLabel(value objc.IObject /* cross-framework: NSString */)
	ItemLoadingToken() objc.IObject /* cross-framework: AccessibilityLoadingToken */
	TargetElement() objc.ID
	TargetRange() corefoundation.Range
	SetTargetRange(value corefoundation.Range)
	CurrentItem() IAccessibilityCustomRotorItemResult
	SetCurrentItem(value IAccessibilityCustomRotorItemResult)
	// methods:
}

// A target accessibility element that a custom rotor references.


// A target accessibility element that a custom rotor references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult
type AccessibilityCustomRotorItemResult struct {
	objectivec.Object
}

// AccessibilityCustomRotorItemResultFrom constructs a [AccessibilityCustomRotorItemResult] from an unsafe.Pointer.
//
// A target accessibility element that a custom rotor references.
func AccessibilityCustomRotorItemResultFrom(ptr unsafe.Pointer) AccessibilityCustomRotorItemResult {
	return AccessibilityCustomRotorItemResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorItemResultClass) Alloc() AccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccessibilityCustomRotorItemResultClass) New() AccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityCustomRotorItemResult) Init() AccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityCustomRotorItemResult) Autorelease() AccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityCustomRotorItemResult creates a new AccessibilityCustomRotorItemResult instance.
func NewAccessibilityCustomRotorItemResult() AccessibilityCustomRotorItemResult {
	return getAccessibilityCustomRotorItemResultClass().New()
}



// Creates an item result with the specified item load token and custom label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(itemLoadingToken:customLabel:)
func NewAccessibilityCustomRotorItemResultWithItemLoadingTokenCustomLabel(itemLoadingToken objc.IObject /* cross-framework: AccessibilityLoadingToken */, customLabel objc.IObject /* cross-framework: NSString */) AccessibilityCustomRotorItemResult {
	instance := getAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotorItemResult](instance.ID, objc.Sel("initWithItemLoadingToken:customLabel:"), itemLoadingToken, customLabel)
	rv.Autorelease()
	return rv
}


// Creates an item result with the specified target element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(targetElement:)
func NewAccessibilityCustomRotorItemResultWithTargetElement(targetElement objc.IObject) AccessibilityCustomRotorItemResult {
	instance := getAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotorItemResult](instance.ID, objc.Sel("initWithTargetElement:"), targetElement)
	rv.Autorelease()
	return rv
}



// A localized label to use instead of the default item label to describe the item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/customLabel
func (a_ AccessibilityCustomRotorItemResult) CustomLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("customLabel"))
	return rv
}


// A localized label to use instead of the default item label to describe the item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/customLabel
func (a_ AccessibilityCustomRotorItemResult) SetCustomLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomLabel:"), value)
}


// A token to determine which item to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/itemLoadingToken
func (a_ AccessibilityCustomRotorItemResult) ItemLoadingToken() objc.IObject /* cross-framework: AccessibilityLoadingToken */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("itemLoadingToken"))
	return rv
}


// A target element that references an element to message for accessibility properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetElement
func (a_ AccessibilityCustomRotorItemResult) TargetElement() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("targetElement"))
	return rv
}


// A range that specifies the area of interest for text-based elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetRange
func (a_ AccessibilityCustomRotorItemResult) TargetRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](a_.ID, objc.Sel("targetRange"))
	return rv
}


// A range that specifies the area of interest for text-based elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetRange
func (a_ AccessibilityCustomRotorItemResult) SetTargetRange(value corefoundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTargetRange:"), value)
}


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/searchparameters/currentitem
func (a_ AccessibilityCustomRotorItemResult) CurrentItem() IAccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("currentItem"))
	return rv
}


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/searchparameters/currentitem
func (a_ AccessibilityCustomRotorItemResult) SetCurrentItem(value IAccessibilityCustomRotorItemResult) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentItem:"), value)
}


