// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionRoot */


/* debug [class_header]: Header for AXMathExpressionRoot */
// The class instance for the [AXMathExpressionRoot] class.
var (
	AXMathExpressionRootClass     _AXMathExpressionRootClass
	AXMathExpressionRootClassOnce sync.Once
)

func getAXMathExpressionRootClass() _AXMathExpressionRootClass {
	AXMathExpressionRootClassOnce.Do(func() {
		AXMathExpressionRootClass = _AXMathExpressionRootClass{objc.GetClass("AXMathExpressionRoot")}
	})
	return AXMathExpressionRootClass
}

type _AXMathExpressionRootClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionRoot */
// An interface definition for the [AXMathExpressionRoot] class.
type IAXMathExpressionRoot interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionRoot */
	// properties:
	RadicandExpressions() []AXMathExpression
	RootIndexExpression() IAXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionRoot */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionRoot */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionRootClass) Alloc() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionRootClass) New() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionRoot) Init() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionRoot) Autorelease() AXMathExpressionRoot {
	rv := objc.Send[AXMathExpressionRoot](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionRoot creates a new AXMathExpressionRoot instance.
func NewAXMathExpressionRoot() AXMathExpressionRoot {
	return getAXMathExpressionRootClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionRoot */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot
type AXMathExpressionRoot struct {
	AXMathExpression
}

// AXMathExpressionRootFrom constructs a [AXMathExpressionRoot] from an unsafe.Pointer.
func AXMathExpressionRootFrom(ptr unsafe.Pointer) AXMathExpressionRoot {
	return AXMathExpressionRoot{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionRoot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/init(radicandExpressions:rootIndexExpression:)
func NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression(radicandExpressions []AXMathExpression, rootIndexExpression IAXMathExpression) AXMathExpressionRoot {
	instance := getAXMathExpressionRootClass().Alloc()
	rv := objc.Send[AXMathExpressionRoot](instance.ID, objc.Sel("initWithRadicandExpressions:rootIndexExpression:"), radicandExpressions, rootIndexExpression)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionRootWithRadicandExpressionsRootIndexExpression */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionRoot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionRoot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionRoot */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionRoot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/radicandExpressions
func (a_ AXMathExpressionRoot) RadicandExpressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("radicandExpressions"))
	return rv
}/* debug [instance_properties/getter]: radicandExpressions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionRoot/rootIndexExpression
func (a_ AXMathExpressionRoot) RootIndexExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("rootIndexExpression"))
	return rv
}/* debug [instance_properties/getter]: rootIndexExpression */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionRoot */


