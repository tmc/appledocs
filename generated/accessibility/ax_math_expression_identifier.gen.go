// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AXMathExpressionIdentifier */


/* debug [class_header]: Header for AXMathExpressionIdentifier */
// The class instance for the [AXMathExpressionIdentifier] class.
var (
	AXMathExpressionIdentifierClass     _AXMathExpressionIdentifierClass
	AXMathExpressionIdentifierClassOnce sync.Once
)

func getAXMathExpressionIdentifierClass() _AXMathExpressionIdentifierClass {
	AXMathExpressionIdentifierClassOnce.Do(func() {
		AXMathExpressionIdentifierClass = _AXMathExpressionIdentifierClass{objc.GetClass("AXMathExpressionIdentifier")}
	})
	return AXMathExpressionIdentifierClass
}

type _AXMathExpressionIdentifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionIdentifier */
// An interface definition for the [AXMathExpressionIdentifier] class.
type IAXMathExpressionIdentifier interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionIdentifier */
	// properties:
	Content() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionIdentifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionIdentifier */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionIdentifierClass) Alloc() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionIdentifierClass) New() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionIdentifier) Init() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionIdentifier) Autorelease() AXMathExpressionIdentifier {
	rv := objc.Send[AXMathExpressionIdentifier](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionIdentifier creates a new AXMathExpressionIdentifier instance.
func NewAXMathExpressionIdentifier() AXMathExpressionIdentifier {
	return getAXMathExpressionIdentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier
type AXMathExpressionIdentifier struct {
	AXMathExpression
}

// AXMathExpressionIdentifierFrom constructs a [AXMathExpressionIdentifier] from an unsafe.Pointer.
func AXMathExpressionIdentifierFrom(ptr unsafe.Pointer) AXMathExpressionIdentifier {
	return AXMathExpressionIdentifier{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionIdentifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/init(content:)
func NewAXMathExpressionIdentifierWithContent(content objc.IObject /* cross-framework: NSString */) AXMathExpressionIdentifier {
	instance := getAXMathExpressionIdentifierClass().Alloc()
	rv := objc.Send[AXMathExpressionIdentifier](instance.ID, objc.Sel("initWithContent:"), content)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionIdentifierWithContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionIdentifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionIdentifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionIdentifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionIdentifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionIdentifier/content
func (a_ AXMathExpressionIdentifier) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionIdentifier */


