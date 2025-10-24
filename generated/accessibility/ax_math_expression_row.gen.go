// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionRow */


/* debug [class_header]: Header for AXMathExpressionRow */
// The class instance for the [AXMathExpressionRow] class.
var (
	AXMathExpressionRowClass     _AXMathExpressionRowClass
	AXMathExpressionRowClassOnce sync.Once
)

func getAXMathExpressionRowClass() _AXMathExpressionRowClass {
	AXMathExpressionRowClassOnce.Do(func() {
		AXMathExpressionRowClass = _AXMathExpressionRowClass{objc.GetClass("AXMathExpressionRow")}
	})
	return AXMathExpressionRowClass
}

type _AXMathExpressionRowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionRow */
// An interface definition for the [AXMathExpressionRow] class.
type IAXMathExpressionRow interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionRow */
	// properties:
	Expressions() []AXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionRow */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionRow */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionRowClass) Alloc() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionRowClass) New() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionRow) Init() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionRow) Autorelease() AXMathExpressionRow {
	rv := objc.Send[AXMathExpressionRow](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRow creates a new AXMathExpressionRow instance.
func NewAXMathExpressionRow() AXMathExpressionRow {
	return getAXMathExpressionRowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow
type AXMathExpressionRow struct {
	AXMathExpression
}

// AXMathExpressionRowFrom constructs a [AXMathExpressionRow] from an unsafe.Pointer.
func AXMathExpressionRowFrom(ptr unsafe.Pointer) AXMathExpressionRow {
	return AXMathExpressionRow{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/init(expressions:)
func NewAXMathExpressionRowWithExpressions(expressions []AXMathExpression) AXMathExpressionRow {
	instance := getAXMathExpressionRowClass().Alloc()
	rv := objc.Send[AXMathExpressionRow](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionRowWithExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionRow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionRow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionRow */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRow/expressions
func (a_ AXMathExpressionRow) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}/* debug [instance_properties/getter]: expressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionRow */


