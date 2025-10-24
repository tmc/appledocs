// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionTableCell */


/* debug [class_header]: Header for AXMathExpressionTableCell */
// The class instance for the [AXMathExpressionTableCell] class.
var (
	AXMathExpressionTableCellClass     _AXMathExpressionTableCellClass
	AXMathExpressionTableCellClassOnce sync.Once
)

func getAXMathExpressionTableCellClass() _AXMathExpressionTableCellClass {
	AXMathExpressionTableCellClassOnce.Do(func() {
		AXMathExpressionTableCellClass = _AXMathExpressionTableCellClass{objc.GetClass("AXMathExpressionTableCell")}
	})
	return AXMathExpressionTableCellClass
}

type _AXMathExpressionTableCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionTableCell */
// An interface definition for the [AXMathExpressionTableCell] class.
type IAXMathExpressionTableCell interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionTableCell */
	// properties:
	Expressions() []AXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionTableCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionTableCell */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableCellClass) Alloc() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionTableCellClass) New() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTableCell) Init() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTableCell) Autorelease() AXMathExpressionTableCell {
	rv := objc.Send[AXMathExpressionTableCell](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTableCell creates a new AXMathExpressionTableCell instance.
func NewAXMathExpressionTableCell() AXMathExpressionTableCell {
	return getAXMathExpressionTableCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionTableCell */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell
type AXMathExpressionTableCell struct {
	AXMathExpression
}

// AXMathExpressionTableCellFrom constructs a [AXMathExpressionTableCell] from an unsafe.Pointer.
func AXMathExpressionTableCellFrom(ptr unsafe.Pointer) AXMathExpressionTableCell {
	return AXMathExpressionTableCell{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionTableCell */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/init(expressions:)
func NewAXMathExpressionTableCellWithExpressions(expressions []AXMathExpression) AXMathExpressionTableCell {
	instance := getAXMathExpressionTableCellClass().Alloc()
	rv := objc.Send[AXMathExpressionTableCell](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionTableCellWithExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionTableCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionTableCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionTableCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionTableCell */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTableCell/expressions
func (a_ AXMathExpressionTableCell) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}/* debug [instance_properties/getter]: expressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionTableCell */


