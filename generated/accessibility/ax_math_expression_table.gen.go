// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionTable */


/* debug [class_header]: Header for AXMathExpressionTable */
// The class instance for the [AXMathExpressionTable] class.
var (
	AXMathExpressionTableClass     _AXMathExpressionTableClass
	AXMathExpressionTableClassOnce sync.Once
)

func getAXMathExpressionTableClass() _AXMathExpressionTableClass {
	AXMathExpressionTableClassOnce.Do(func() {
		AXMathExpressionTableClass = _AXMathExpressionTableClass{objc.GetClass("AXMathExpressionTable")}
	})
	return AXMathExpressionTableClass
}

type _AXMathExpressionTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionTable */
// An interface definition for the [AXMathExpressionTable] class.
type IAXMathExpressionTable interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionTable */
	// properties:
	Expressions() []AXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionTable */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTableClass) Alloc() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionTableClass) New() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionTable) Init() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionTable) Autorelease() AXMathExpressionTable {
	rv := objc.Send[AXMathExpressionTable](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionTable creates a new AXMathExpressionTable instance.
func NewAXMathExpressionTable() AXMathExpressionTable {
	return getAXMathExpressionTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionTable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable
type AXMathExpressionTable struct {
	AXMathExpression
}

// AXMathExpressionTableFrom constructs a [AXMathExpressionTable] from an unsafe.Pointer.
func AXMathExpressionTableFrom(ptr unsafe.Pointer) AXMathExpressionTable {
	return AXMathExpressionTable{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionTable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/init(expressions:)
func NewAXMathExpressionTableWithExpressions(expressions []AXMathExpression) AXMathExpressionTable {
	instance := getAXMathExpressionTableClass().Alloc()
	rv := objc.Send[AXMathExpressionTable](instance.ID, objc.Sel("initWithExpressions:"), expressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionTableWithExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionTable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionTable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionTable/expressions
func (a_ AXMathExpressionTable) Expressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("expressions"))
	return rv
}/* debug [instance_properties/getter]: expressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionTable */


