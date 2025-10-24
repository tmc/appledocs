// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AXMathExpressionMultiscript */


/* debug [class_header]: Header for AXMathExpressionMultiscript */
// The class instance for the [AXMathExpressionMultiscript] class.
var (
	AXMathExpressionMultiscriptClass     _AXMathExpressionMultiscriptClass
	AXMathExpressionMultiscriptClassOnce sync.Once
)

func getAXMathExpressionMultiscriptClass() _AXMathExpressionMultiscriptClass {
	AXMathExpressionMultiscriptClassOnce.Do(func() {
		AXMathExpressionMultiscriptClass = _AXMathExpressionMultiscriptClass{objc.GetClass("AXMathExpressionMultiscript")}
	})
	return AXMathExpressionMultiscriptClass
}

type _AXMathExpressionMultiscriptClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionMultiscript */
// An interface definition for the [AXMathExpressionMultiscript] class.
type IAXMathExpressionMultiscript interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionMultiscript */
	// properties:
	BaseExpression() IAXMathExpression
	PostscriptExpressions() []AXMathExpressionSubSuperscript
	PrescriptExpressions() []AXMathExpressionSubSuperscript
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionMultiscript */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionMultiscript */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionMultiscriptClass) Alloc() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionMultiscriptClass) New() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionMultiscript) Init() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionMultiscript) Autorelease() AXMathExpressionMultiscript {
	rv := objc.Send[AXMathExpressionMultiscript](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionMultiscript creates a new AXMathExpressionMultiscript instance.
func NewAXMathExpressionMultiscript() AXMathExpressionMultiscript {
	return getAXMathExpressionMultiscriptClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionMultiscript */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript
type AXMathExpressionMultiscript struct {
	AXMathExpression
}

// AXMathExpressionMultiscriptFrom constructs a [AXMathExpressionMultiscript] from an unsafe.Pointer.
func AXMathExpressionMultiscriptFrom(ptr unsafe.Pointer) AXMathExpressionMultiscript {
	return AXMathExpressionMultiscript{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionMultiscript */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/init(baseExpression:prescriptExpressions:postscriptExpressions:)
func NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions(baseExpression IAXMathExpression, prescriptExpressions []AXMathExpressionSubSuperscript, postscriptExpressions []AXMathExpressionSubSuperscript) AXMathExpressionMultiscript {
	instance := getAXMathExpressionMultiscriptClass().Alloc()
	rv := objc.Send[AXMathExpressionMultiscript](instance.ID, objc.Sel("initWithBaseExpression:prescriptExpressions:postscriptExpressions:"), baseExpression, prescriptExpressions, postscriptExpressions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionMultiscriptWithBaseExpressionPrescriptExpressionsPostscriptExpressions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionMultiscript */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionMultiscript */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionMultiscript */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionMultiscript */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/baseExpression
func (a_ AXMathExpressionMultiscript) BaseExpression() IAXMathExpression {
	rv := objc.Send[AXMathExpression](a_.ID, objc.Sel("baseExpression"))
	return rv
}/* debug [instance_properties/getter]: baseExpression */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/postscriptExpressions
func (a_ AXMathExpressionMultiscript) PostscriptExpressions() []AXMathExpressionSubSuperscript {
	rv := objc.Send[[]AXMathExpressionSubSuperscript](a_.ID, objc.Sel("postscriptExpressions"))
	return rv
}/* debug [instance_properties/getter]: postscriptExpressions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionMultiscript/prescriptExpressions
func (a_ AXMathExpressionMultiscript) PrescriptExpressions() []AXMathExpressionSubSuperscript {
	rv := objc.Send[[]AXMathExpressionSubSuperscript](a_.ID, objc.Sel("prescriptExpressions"))
	return rv
}/* debug [instance_properties/getter]: prescriptExpressions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionMultiscript */


