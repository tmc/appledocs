// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AXMathExpressionNumber */


/* debug [class_header]: Header for AXMathExpressionNumber */
// The class instance for the [AXMathExpressionNumber] class.
var (
	AXMathExpressionNumberClass     _AXMathExpressionNumberClass
	AXMathExpressionNumberClassOnce sync.Once
)

func getAXMathExpressionNumberClass() _AXMathExpressionNumberClass {
	AXMathExpressionNumberClassOnce.Do(func() {
		AXMathExpressionNumberClass = _AXMathExpressionNumberClass{objc.GetClass("AXMathExpressionNumber")}
	})
	return AXMathExpressionNumberClass
}

type _AXMathExpressionNumberClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionNumber */
// An interface definition for the [AXMathExpressionNumber] class.
type IAXMathExpressionNumber interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionNumber */
	// properties:
	Content() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionNumber */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionNumber */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionNumberClass) Alloc() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionNumberClass) New() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionNumber) Init() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionNumber) Autorelease() AXMathExpressionNumber {
	rv := objc.Send[AXMathExpressionNumber](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionNumber creates a new AXMathExpressionNumber instance.
func NewAXMathExpressionNumber() AXMathExpressionNumber {
	return getAXMathExpressionNumberClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber
type AXMathExpressionNumber struct {
	AXMathExpression
}

// AXMathExpressionNumberFrom constructs a [AXMathExpressionNumber] from an unsafe.Pointer.
func AXMathExpressionNumberFrom(ptr unsafe.Pointer) AXMathExpressionNumber {
	return AXMathExpressionNumber{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionNumber */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/init(content:)
func NewAXMathExpressionNumberWithContent(content objc.IObject /* cross-framework: NSString */) AXMathExpressionNumber {
	instance := getAXMathExpressionNumberClass().Alloc()
	rv := objc.Send[AXMathExpressionNumber](instance.ID, objc.Sel("initWithContent:"), content)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionNumberWithContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionNumber */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionNumber */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionNumber */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionNumber */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionNumber/content
func (a_ AXMathExpressionNumber) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionNumber */


