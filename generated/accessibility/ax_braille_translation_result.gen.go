// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXBrailleTranslationResult] class.
var (
	AXBrailleTranslationResultClass     _AXBrailleTranslationResultClass
	AXBrailleTranslationResultClassOnce sync.Once
)

func getAXBrailleTranslationResultClass() _AXBrailleTranslationResultClass {
	AXBrailleTranslationResultClassOnce.Do(func() {
		AXBrailleTranslationResultClass = _AXBrailleTranslationResultClass{objc.GetClass("AXBrailleTranslationResult")}
	})
	return AXBrailleTranslationResultClass
}

type _AXBrailleTranslationResultClass struct {
	class objc.Class
}

// An interface definition for the [AXBrailleTranslationResult] class.
type IAXBrailleTranslationResult interface {
	objectivec.IObject
	LocationMap() []foundation.Number
	ResultString() string
}

// The result of translation or back-translation.


// The result of translation or back-translation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult
type AXBrailleTranslationResult struct {
	objectivec.Object
}

// AXBrailleTranslationResultFrom constructs a [AXBrailleTranslationResult] from an unsafe.Pointer.
//
// The result of translation or back-translation.
func AXBrailleTranslationResultFrom(ptr unsafe.Pointer) AXBrailleTranslationResult {
	return AXBrailleTranslationResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTranslationResultClass) Alloc() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXBrailleTranslationResultClass) New() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXBrailleTranslationResult) Init() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXBrailleTranslationResult) Autorelease() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTranslationResult creates a new AXBrailleTranslationResult instance.
func NewAXBrailleTranslationResult() AXBrailleTranslationResult {
	return getAXBrailleTranslationResultClass().New()
}



// An array of integers that has the same length as the resultString. locationMap[i]-th character in the input string corresponds to resultString[i].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/locationMap
func (a_ AXBrailleTranslationResult) LocationMap() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("locationMap"))
	return rv
}


// The resulting string after translation or back-translation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/resultString
func (a_ AXBrailleTranslationResult) ResultString() string {
	rv := objc.Send[string](a_.ID, objc.Sel("resultString"))
	return rv
}



