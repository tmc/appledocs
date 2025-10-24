// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AXMathExpressionFenced */


/* debug [class_header]: Header for AXMathExpressionFenced */
// The class instance for the [AXMathExpressionFenced] class.
var (
	AXMathExpressionFencedClass     _AXMathExpressionFencedClass
	AXMathExpressionFencedClassOnce sync.Once
)

func getAXMathExpressionFencedClass() _AXMathExpressionFencedClass {
	AXMathExpressionFencedClassOnce.Do(func() {
		AXMathExpressionFencedClass = _AXMathExpressionFencedClass{objc.GetClass("AXMathExpressionFenced")}
	})
	return AXMathExpressionFencedClass
}

type _AXMathExpressionFencedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionFenced */
// An interface definition for the [AXMathExpressionFenced] class.
type IAXMathExpressionFenced interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionFenced */
	// properties:
	CloseString() objc.IObject /* cross-framework: NSString */
	Expressions() []AXMathExpression
	OpenString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionFenced */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionFenced */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionFencedClass) Alloc() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionFencedClass) New() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionFenced) Init() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionFenced) Autorelease() AXMathExpressionFenced {
	rv := objc.Send[AXMathExpressionFenced](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFenced creates a new AXMathExpressionFenced instance.
func NewAXMathExpressionFenced() AXMathExpressionFenced {
	return getAXMathExpressionFencedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionFenced */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced
type AXMathExpressionFenced struct {
	AXMathExpression
}

// AXMathExpressionFencedFrom constructs a [AXMathExpressionFenced] from an unsafe.Pointer.
func AXMathExpressionFencedFrom(ptr unsafe.Pointer) AXMathExpressionFenced {
	return AXMathExpressionFenced{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionFenced */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/init(expressions:open:close:)
func NewAXMathExpressionFencedWithExpressionsOpenStringCloseString(expressions []AXMathExpression, openString objc.IObject /* cross-framework: NSString */, closeString objc.IObject /* cross-framework: NSString */) AXMathExpressionFenced {
	instance := getAXMathExpressionFencedClass().Alloc()
	rv := objc.Send[AXMathExpressionFenced](instance.ID, objc.Sel("initWithExpressions:openString:closeString:"), expressions, openString, closeString)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionFencedWithExpressionsOpenStringCloseString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionFenced */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionFenced */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionFenced */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionFenced */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/closeString
func (a_ AXMathExpressionFenced) CloseString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("closeString"))
	return rv
}/* debug [instance_properties/getter]: closeString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/expressions
func (a_ AXMathExpressionFenced) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}/* debug [instance_properties/getter]: expressions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFenced/openString
func (a_ AXMathExpressionFenced) OpenString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("openString"))
	return rv
}/* debug [instance_properties/getter]: openString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionFenced */


