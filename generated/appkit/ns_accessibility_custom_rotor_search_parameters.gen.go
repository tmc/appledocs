// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAccessibilityCustomRotorSearchParameters */


/* debug [class_header]: Header for NSAccessibilityCustomRotorSearchParameters */
// The class instance for the [AccessibilityCustomRotorSearchParameters] class.
var (
	AccessibilityCustomRotorSearchParametersClass     _AccessibilityCustomRotorSearchParametersClass
	AccessibilityCustomRotorSearchParametersClassOnce sync.Once
)

func getAccessibilityCustomRotorSearchParametersClass() _AccessibilityCustomRotorSearchParametersClass {
	AccessibilityCustomRotorSearchParametersClassOnce.Do(func() {
		AccessibilityCustomRotorSearchParametersClass = _AccessibilityCustomRotorSearchParametersClass{objc.GetClass("NSAccessibilityCustomRotorSearchParameters")}
	})
	return AccessibilityCustomRotorSearchParametersClass
}

type _AccessibilityCustomRotorSearchParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccessibilityCustomRotorSearchParameters */
// An interface definition for the [AccessibilityCustomRotorSearchParameters] class.
type IAccessibilityCustomRotorSearchParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccessibilityCustomRotorSearchParameters */
	// properties:
	CurrentItem() IAccessibilityCustomRotorItemResult
	SetCurrentItem(value IAccessibilityCustomRotorItemResult)
	FilterString() objc.IObject /* cross-framework: NSString */
	SetFilterString(value objc.IObject /* cross-framework: NSString */)
	SearchDirection() AccessibilityCustomRotorSearchDirection
	SetSearchDirection(value AccessibilityCustomRotorSearchDirection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccessibilityCustomRotorSearchParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccessibilityCustomRotorSearchParameters */
// Alloc allocates a new instance without initialization.
func (ac _AccessibilityCustomRotorSearchParametersClass) Alloc() AccessibilityCustomRotorSearchParameters {
	rv := objc.Send[AccessibilityCustomRotorSearchParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccessibilityCustomRotorSearchParametersClass) New() AccessibilityCustomRotorSearchParameters {
	rv := objc.Send[AccessibilityCustomRotorSearchParameters](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccessibilityCustomRotorSearchParameters) Init() AccessibilityCustomRotorSearchParameters {
	rv := objc.Send[AccessibilityCustomRotorSearchParameters](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccessibilityCustomRotorSearchParameters) Autorelease() AccessibilityCustomRotorSearchParameters {
	rv := objc.Send[AccessibilityCustomRotorSearchParameters](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccessibilityCustomRotorSearchParameters creates a new AccessibilityCustomRotorSearchParameters instance.
func NewAccessibilityCustomRotorSearchParameters() AccessibilityCustomRotorSearchParameters {
	return getAccessibilityCustomRotorSearchParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccessibilityCustomRotorSearchParameters */
// Search parameters for a custom rotor.
//
// Use these parameters to determine the next matching .


// Search parameters for a custom rotor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters
type AccessibilityCustomRotorSearchParameters struct {
	objectivec.Object
}

// AccessibilityCustomRotorSearchParametersFrom constructs a [AccessibilityCustomRotorSearchParameters] from an unsafe.Pointer.
//
// Search parameters for a custom rotor.
func AccessibilityCustomRotorSearchParametersFrom(ptr unsafe.Pointer) AccessibilityCustomRotorSearchParameters {
	return AccessibilityCustomRotorSearchParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccessibilityCustomRotorSearchParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccessibilityCustomRotorSearchParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccessibilityCustomRotorSearchParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccessibilityCustomRotorSearchParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccessibilityCustomRotorSearchParameters */

// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/currentItem
func (a_ AccessibilityCustomRotorSearchParameters) CurrentItem() IAccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("currentItem"))
	return rv
}/* debug [instance_properties/getter]: currentItem */


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/currentItem
func (a_ AccessibilityCustomRotorSearchParameters) SetCurrentItem(value IAccessibilityCustomRotorItemResult) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentItem:"), value)
}/* debug [instance_properties/setter]: currentItem */


// A string of text to filter the results against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/filterString
func (a_ AccessibilityCustomRotorSearchParameters) FilterString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("filterString"))
	return rv
}/* debug [instance_properties/getter]: filterString */


// A string of text to filter the results against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/filterString
func (a_ AccessibilityCustomRotorSearchParameters) SetFilterString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterString:"), value)
}/* debug [instance_properties/setter]: filterString */


// The direction to search for an item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/searchDirection
func (a_ AccessibilityCustomRotorSearchParameters) SearchDirection() AccessibilityCustomRotorSearchDirection {
	rv := objc.Send[AccessibilityCustomRotorSearchDirection](a_.ID, objc.Sel("searchDirection"))
	return rv
}/* debug [instance_properties/getter]: searchDirection */


// The direction to search for an item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/searchDirection
func (a_ AccessibilityCustomRotorSearchParameters) SetSearchDirection(value AccessibilityCustomRotorSearchDirection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSearchDirection:"), value)
}/* debug [instance_properties/setter]: searchDirection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAccessibilityCustomRotorSearchParameters */



