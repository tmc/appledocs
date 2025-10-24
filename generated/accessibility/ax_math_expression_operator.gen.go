// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AXMathExpressionOperator */


/* debug [class_header]: Header for AXMathExpressionOperator */
// The class instance for the [AXMathExpressionOperator] class.
var (
	AXMathExpressionOperatorClass     _AXMathExpressionOperatorClass
	AXMathExpressionOperatorClassOnce sync.Once
)

func getAXMathExpressionOperatorClass() _AXMathExpressionOperatorClass {
	AXMathExpressionOperatorClassOnce.Do(func() {
		AXMathExpressionOperatorClass = _AXMathExpressionOperatorClass{objc.GetClass("AXMathExpressionOperator")}
	})
	return AXMathExpressionOperatorClass
}

type _AXMathExpressionOperatorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionOperator */
// An interface definition for the [AXMathExpressionOperator] class.
type IAXMathExpressionOperator interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionOperator */
	// properties:
	Content() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionOperator */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionOperator */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionOperatorClass) Alloc() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionOperatorClass) New() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionOperator) Init() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionOperator) Autorelease() AXMathExpressionOperator {
	rv := objc.Send[AXMathExpressionOperator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionOperator creates a new AXMathExpressionOperator instance.
func NewAXMathExpressionOperator() AXMathExpressionOperator {
	return getAXMathExpressionOperatorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionOperator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator
type AXMathExpressionOperator struct {
	AXMathExpression
}

// AXMathExpressionOperatorFrom constructs a [AXMathExpressionOperator] from an unsafe.Pointer.
func AXMathExpressionOperatorFrom(ptr unsafe.Pointer) AXMathExpressionOperator {
	return AXMathExpressionOperator{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionOperator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/init(content:)
func NewAXMathExpressionOperatorWithContent(content objc.IObject /* cross-framework: NSString */) AXMathExpressionOperator {
	instance := getAXMathExpressionOperatorClass().Alloc()
	rv := objc.Send[AXMathExpressionOperator](instance.ID, objc.Sel("initWithContent:"), content)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionOperatorWithContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionOperator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionOperator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionOperator */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionOperator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionOperator/content
func (a_ AXMathExpressionOperator) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionOperator */


