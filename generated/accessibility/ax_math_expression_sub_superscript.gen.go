// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionSubSuperscript */


/* debug [class_header]: Header for AXMathExpressionSubSuperscript */
// The class instance for the [AXMathExpressionSubSuperscript] class.
var (
	AXMathExpressionSubSuperscriptClass     _AXMathExpressionSubSuperscriptClass
	AXMathExpressionSubSuperscriptClassOnce sync.Once
)

func getAXMathExpressionSubSuperscriptClass() _AXMathExpressionSubSuperscriptClass {
	AXMathExpressionSubSuperscriptClassOnce.Do(func() {
		AXMathExpressionSubSuperscriptClass = _AXMathExpressionSubSuperscriptClass{objc.GetClass("AXMathExpressionSubSuperscript")}
	})
	return AXMathExpressionSubSuperscriptClass
}

type _AXMathExpressionSubSuperscriptClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionSubSuperscript */
// An interface definition for the [AXMathExpressionSubSuperscript] class.
type IAXMathExpressionSubSuperscript interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionSubSuperscript */
	// properties:
	BaseExpression() IAXMathExpression
	SubscriptExpressions() []AXMathExpression
	SuperscriptExpressions() []AXMathExpression
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionSubSuperscript */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionSubSuperscript */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionSubSuperscriptClass) Alloc() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionSubSuperscriptClass) New() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionSubSuperscript) Init() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionSubSuperscript) Autorelease() AXMathExpressionSubSuperscript {
	rv := objc.Send[AXMathExpressionSubSuperscript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionSubSuperscript creates a new AXMathExpressionSubSuperscript instance.
func NewAXMathExpressionSubSuperscript() AXMathExpressionSubSuperscript {
	return getAXMathExpressionSubSuperscriptClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionSubSuperscript */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript
type AXMathExpressionSubSuperscript struct {
	AXMathExpression
}

// AXMathExpressionSubSuperscriptFrom constructs a [AXMathExpressionSubSuperscript] from an unsafe.Pointer.
func AXMathExpressionSubSuperscriptFrom(ptr unsafe.Pointer) AXMathExpressionSubSuperscript {
	return AXMathExpressionSubSuperscript{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionSubSuperscript */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/init(baseExpression:subscriptExpressions:superscriptExpressions:)
func NewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions(baseExpression []AXMathExpression, subscriptExpressions []AXMathExpression, superscriptExpressions []AXMathExpression) AXMathExpressionSubSuperscript {
	instance := getAXMathExpressionSubSuperscriptClass().Alloc()
	rv := objc.Send[AXMathExpressionSubSuperscript](instance.ID, objc.Sel("initWithBaseExpression:subscriptExpressions:superscriptExpressions:"), baseExpression, subscriptExpressions, superscriptExpressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionSubSuperscriptWithBaseExpressionSubscriptExpressionsSuperscriptExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionSubSuperscript */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionSubSuperscript */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionSubSuperscript */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionSubSuperscript */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/baseExpression
func (a_ AXMathExpressionSubSuperscript) BaseExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}/* debug [instance_properties/getter]: baseExpression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/subscriptExpressions
func (a_ AXMathExpressionSubSuperscript) SubscriptExpressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("subscriptExpressions"))
	return rv
}/* debug [instance_properties/getter]: subscriptExpressions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionSubSuperscript/superscriptExpressions
func (a_ AXMathExpressionSubSuperscript) SuperscriptExpressions() []AXMathExpression {
	rv := objc.Send[[]AXMathExpression](a_.ID, objc.Sel("superscriptExpressions"))
	return rv
}/* debug [instance_properties/getter]: superscriptExpressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionSubSuperscript */


