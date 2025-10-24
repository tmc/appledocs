// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXBrailleTranslationResult */


/* debug [class_header]: Header for AXBrailleTranslationResult */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXBrailleTranslationResult */
// An interface definition for the [AXBrailleTranslationResult] class.
type IAXBrailleTranslationResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXBrailleTranslationResult */
	// properties:
	LocationMap() []foundation.Number
	ResultString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXBrailleTranslationResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXBrailleTranslationResult */
// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTranslationResultClass) Alloc() AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXBrailleTranslationResult */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXBrailleTranslationResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXBrailleTranslationResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXBrailleTranslationResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXBrailleTranslationResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXBrailleTranslationResult */

// An array of integers that has the same length as the resultString. locationMap[i]-th character in the input string corresponds to resultString[i].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/locationMap
func (a_ AXBrailleTranslationResult) LocationMap() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("locationMap"))
	return rv
}/* debug [instance_properties/getter]: locationMap */


// The resulting string after translation or back-translation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslationResult/resultString
func (a_ AXBrailleTranslationResult) ResultString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("resultString"))
	return rv
}/* debug [instance_properties/getter]: resultString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXBrailleTranslationResult */



