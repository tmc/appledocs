// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [AccessibilityCustomRotorSearchParameters] class.
type IAccessibilityCustomRotorSearchParameters interface {
	objectivec.IObject
	

	// properties:
	CurrentItem() IAccessibilityCustomRotorItemResult
	SetCurrentItem(value IAccessibilityCustomRotorItemResult)
	FilterString() foundation.foundation.INSString
	SetFilterString(value foundation.foundation.INSString)
	SearchDirection() AccessibilityCustomRotorSearchDirection
	SetSearchDirection(value AccessibilityCustomRotorSearchDirection)


	

	// methods:


}





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

























// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/currentItem
func (a_ AccessibilityCustomRotorSearchParameters) CurrentItem() IAccessibilityCustomRotorItemResult {
	rv := objc.Send[AccessibilityCustomRotorItemResult](a_.ID, objc.Sel("currentItem"))
	return rv
}


// The current item that determines where the search starts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/currentItem
func (a_ AccessibilityCustomRotorSearchParameters) SetCurrentItem(value IAccessibilityCustomRotorItemResult) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentItem:"), value)
}


// A string of text to filter the results against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/filterString
func (a_ AccessibilityCustomRotorSearchParameters) FilterString() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("filterString"))
	return rv
}


// A string of text to filter the results against.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/filterString
func (a_ AccessibilityCustomRotorSearchParameters) SetFilterString(value foundation.foundation.INSString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFilterString:"), value)
}


// The direction to search for an item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/searchDirection
func (a_ AccessibilityCustomRotorSearchParameters) SearchDirection() AccessibilityCustomRotorSearchDirection {
	rv := objc.Send[AccessibilityCustomRotorSearchDirection](a_.ID, objc.Sel("searchDirection"))
	return rv
}


// The direction to search for an item result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchParameters/searchDirection
func (a_ AccessibilityCustomRotorSearchParameters) SetSearchDirection(value AccessibilityCustomRotorSearchDirection) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSearchDirection:"), value)
}








