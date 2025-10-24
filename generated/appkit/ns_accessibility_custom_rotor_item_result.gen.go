// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAccessibilityCustomRotorItemResult */


/* debug [class_header]: Header for NSAccessibilityCustomRotorItemResult */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccessibilityCustomRotorItemResult */
// An interface definition for the [AccessibilityCustomRotorItemResult] class.
type IAccessibilityCustomRotorItemResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccessibilityCustomRotorItemResult */
	// properties:
	CustomLabel() objc.IObject /* cross-framework: NSString */
	SetCustomLabel(value objc.IObject /* cross-framework: NSString */)
	ItemLoadingToken() AccessibilityLoadingToken /* typedef */
	TargetElement() unsafe.Pointer
	TargetRange() corefoundation.Range
	SetTargetRange(value corefoundation.Range)
	CurrentItem() IAccessibilityCustomRotorItemResult
	SetCurrentItem(value IAccessibilityCustomRotorItemResult)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccessibilityCustomRotorItemResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccessibilityCustomRotorItemResult */
// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorItemResultClass) Alloc() AccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccessibilityCustomRotorItemResult */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccessibilityCustomRotorItemResult */

// Creates an item result with the specified item load token and custom label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(itemLoadingToken:customLabel:)
func NewAccessibilityCustomRotorItemResultWithItemLoadingTokenCustomLabel(itemLoadingToken AccessibilityLoadingToken /* typedef */, customLabel objc.IObject /* cross-framework: NSString */) AccessibilityCustomRotorItemResult {
	instance := getAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotorItemResult](instance.ID, objc.Sel("initWithItemLoadingToken:customLabel:"), itemLoadingToken, customLabel)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccessibilityCustomRotorItemResultWithItemLoadingTokenCustomLabel */


// Creates an item result with the specified target element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/init(targetElement:)
func NewAccessibilityCustomRotorItemResultWithTargetElement(targetElement unsafe.Pointer) AccessibilityCustomRotorItemResult {
	instance := getAccessibilityCustomRotorItemResultClass().Alloc()
	rv := objc.Send[AccessibilityCustomRotorItemResult](instance.ID, objc.Sel("initWithTargetElement:"), targetElement)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccessibilityCustomRotorItemResultWithTargetElement */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccessibilityCustomRotorItemResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccessibilityCustomRotorItemResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccessibilityCustomRotorItemResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccessibilityCustomRotorItemResult */

// A localized label to use instead of the default item label to describe the item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/customLabel
func (a_ AccessibilityCustomRotorItemResult) CustomLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("customLabel"))
	return rv
}/* debug [instance_properties/getter]: customLabel */


// A localized label to use instead of the default item label to describe the item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/customLabel
func (a_ AccessibilityCustomRotorItemResult) SetCustomLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCustomLabel:"), value)
}/* debug [instance_properties/setter]: customLabel */


// A token to determine which item to return.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/itemLoadingToken
func (a_ AccessibilityCustomRotorItemResult) ItemLoadingToken() AccessibilityLoadingToken /* typedef */ {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("itemLoadingToken"))
	return rv
}/* debug [instance_properties/getter]: itemLoadingToken */


// A target element that references an element to message for accessibility properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetElement
func (a_ AccessibilityCustomRotorItemResult) TargetElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("targetElement"))
	return rv
}/* debug [instance_properties/getter]: targetElement */


// A range that specifies the area of interest for text-based elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetRange
func (a_ AccessibilityCustomRotorItemResult) TargetRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](a_.ID, objc.Sel("targetRange"))
	return rv
}/* debug [instance_properties/getter]: targetRange */


// A range that specifies the area of interest for text-based elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/ItemResult/targetRange
func (a_ AccessibilityCustomRotorItemResult) SetTargetRange(value corefoundation.Range) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTargetRange:"), value)
}/* debug [instance_properties/setter]: targetRange */


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/searchparameters/currentitem
func (a_ AccessibilityCustomRotorItemResult) CurrentItem() IAccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("currentItem"))
	return rv
}/* debug [instance_properties/getter]: currentItem */


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilitycustomrotor/searchparameters/currentitem
func (a_ AccessibilityCustomRotorItemResult) SetCurrentItem(value IAccessibilityCustomRotorItemResult) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentItem:"), value)
}/* debug [instance_properties/setter]: currentItem */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAccessibilityCustomRotorItemResult */


