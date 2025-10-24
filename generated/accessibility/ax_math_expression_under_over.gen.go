// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionUnderOver */


/* debug [class_header]: Header for AXMathExpressionUnderOver */
// The class instance for the [AXMathExpressionUnderOver] class.
var (
	AXMathExpressionUnderOverClass     _AXMathExpressionUnderOverClass
	AXMathExpressionUnderOverClassOnce sync.Once
)

func getAXMathExpressionUnderOverClass() _AXMathExpressionUnderOverClass {
	AXMathExpressionUnderOverClassOnce.Do(func() {
		AXMathExpressionUnderOverClass = _AXMathExpressionUnderOverClass{objc.GetClass("AXMathExpressionUnderOver")}
	})
	return AXMathExpressionUnderOverClass
}

type _AXMathExpressionUnderOverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionUnderOver */
// An interface definition for the [AXMathExpressionUnderOver] class.
type IAXMathExpressionUnderOver interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionUnderOver */
	// properties:
	BaseExpression() IAXMathExpression
	OverExpression() IAXMathExpression
	UnderExpression() IAXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionUnderOver */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionUnderOver */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionUnderOverClass) Alloc() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionUnderOverClass) New() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionUnderOver) Init() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionUnderOver) Autorelease() AXMathExpressionUnderOver {
	rv := objc.Send[AXMathExpressionUnderOver](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionUnderOver creates a new AXMathExpressionUnderOver instance.
func NewAXMathExpressionUnderOver() AXMathExpressionUnderOver {
	return getAXMathExpressionUnderOverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionUnderOver */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver
type AXMathExpressionUnderOver struct {
	AXMathExpression
}

// AXMathExpressionUnderOverFrom constructs a [AXMathExpressionUnderOver] from an unsafe.Pointer.
func AXMathExpressionUnderOverFrom(ptr unsafe.Pointer) AXMathExpressionUnderOver {
	return AXMathExpressionUnderOver{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionUnderOver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/init(baseExpression:underExpression:overExpression:)
func NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression(baseExpression IAXMathExpression, underExpression IAXMathExpression, overExpression IAXMathExpression) AXMathExpressionUnderOver {
	instance := getAXMathExpressionUnderOverClass().Alloc()
	rv := objc.Send[AXMathExpressionUnderOver](instance.ID, objc.Sel("initWithBaseExpression:underExpression:overExpression:"), baseExpression, underExpression, overExpression)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionUnderOverWithBaseExpressionUnderExpressionOverExpression */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionUnderOver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionUnderOver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionUnderOver */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionUnderOver */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/baseExpression
func (a_ AXMathExpressionUnderOver) BaseExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}/* debug [instance_properties/getter]: baseExpression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/overExpression
func (a_ AXMathExpressionUnderOver) OverExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("overExpression"))
	return rv
}/* debug [instance_properties/getter]: overExpression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionUnderOver/underExpression
func (a_ AXMathExpressionUnderOver) UnderExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("underExpression"))
	return rv
}/* debug [instance_properties/getter]: underExpression */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionUnderOver */


