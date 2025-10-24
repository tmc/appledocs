// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionTableRow */


/* debug [class_header]: Header for AXMathExpressionTableRow */
// The class instance for the [AXMathExpressionTableRow] class.
var (
	AXMathExpressionTableRowClass     _AXMathExpressionTableRowClass
	AXMathExpressionTableRowClassOnce sync.Once
)

func getAXMathExpressionTableRowClass() _AXMathExpressionTableRowClass {
	AXMathExpressionTableRowClassOnce.Do(func() {
		AXMathExpressionTableRowClass = _AXMathExpressionTableRowClass{objc.GetClass("AXMathExpressionTableRow")}
	})
	return AXMathExpressionTableRowClass
}

type _AXMathExpressionTableRowClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionTableRow */
// An interface definition for the [AXMathExpressionTableRow] class.
type IAXMathExpressionTableRow interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionTableRow */
	// properties:
	Expressions() []AXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionTableRow */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionTableRow */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableRowClass) Alloc() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionTableRowClass) New() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTableRow) Init() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTableRow) Autorelease() AXMathExpressionTableRow {
	rv := objc.Send[AXMathExpressionTableRow](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableRow creates a new AXMathExpressionTableRow instance.
func NewAXMathExpressionTableRow() AXMathExpressionTableRow {
	return getAXMathExpressionTableRowClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionTableRow */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow
type AXMathExpressionTableRow struct {
	AXMathExpression
}

// AXMathExpressionTableRowFrom constructs a [AXMathExpressionTableRow] from an unsafe.Pointer.
func AXMathExpressionTableRowFrom(ptr unsafe.Pointer) AXMathExpressionTableRow {
	return AXMathExpressionTableRow{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionTableRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/init(expressions:)
func NewAXMathExpressionTableRowWithExpressions(expressions []AXMathExpression) AXMathExpressionTableRow {
	instance := getAXMathExpressionTableRowClass().Alloc()
	rv := objc.Send[AXMathExpressionTableRow](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionTableRowWithExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionTableRow */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionTableRow */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionTableRow */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionTableRow */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableRow/expressions
func (a_ AXMathExpressionTableRow) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}/* debug [instance_properties/getter]: expressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionTableRow */


