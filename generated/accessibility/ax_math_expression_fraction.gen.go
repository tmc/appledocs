// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionFraction */


/* debug [class_header]: Header for AXMathExpressionFraction */
// The class instance for the [AXMathExpressionFraction] class.
var (
	AXMathExpressionFractionClass     _AXMathExpressionFractionClass
	AXMathExpressionFractionClassOnce sync.Once
)

func getAXMathExpressionFractionClass() _AXMathExpressionFractionClass {
	AXMathExpressionFractionClassOnce.Do(func() {
		AXMathExpressionFractionClass = _AXMathExpressionFractionClass{objc.GetClass("AXMathExpressionFraction")}
	})
	return AXMathExpressionFractionClass
}

type _AXMathExpressionFractionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionFraction */
// An interface definition for the [AXMathExpressionFraction] class.
type IAXMathExpressionFraction interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionFraction */
	// properties:
	DenimonatorExpression() IAXMathExpression
	NumeratorExpression() IAXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionFraction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionFraction */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionFractionClass) Alloc() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionFractionClass) New() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionFraction) Init() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionFraction) Autorelease() AXMathExpressionFraction {
	rv := objc.Send[AXMathExpressionFraction](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionFraction creates a new AXMathExpressionFraction instance.
func NewAXMathExpressionFraction() AXMathExpressionFraction {
	return getAXMathExpressionFractionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionFraction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction
type AXMathExpressionFraction struct {
	AXMathExpression
}

// AXMathExpressionFractionFrom constructs a [AXMathExpressionFraction] from an unsafe.Pointer.
func AXMathExpressionFractionFrom(ptr unsafe.Pointer) AXMathExpressionFraction {
	return AXMathExpressionFraction{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionFraction */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/init(numeratorExpression:denimonatorExpression:)
func NewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression(numeratorExpression IAXMathExpression, denimonatorExpression IAXMathExpression) AXMathExpressionFraction {
	instance := getAXMathExpressionFractionClass().Alloc()
	rv := objc.Send[AXMathExpressionFraction](instance.ID, objc.Sel("initWithNumeratorExpression:denimonatorExpression:"), numeratorExpression, denimonatorExpression)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionFractionWithNumeratorExpressionDenimonatorExpression */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionFraction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionFraction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionFraction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionFraction */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/denimonatorExpression
func (a_ AXMathExpressionFraction) DenimonatorExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("denimonatorExpression"))
	return rv
}/* debug [instance_properties/getter]: denimonatorExpression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionFraction/numeratorExpression
func (a_ AXMathExpressionFraction) NumeratorExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("numeratorExpression"))
	return rv
}/* debug [instance_properties/getter]: numeratorExpression */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionFraction */


