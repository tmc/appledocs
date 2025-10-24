// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class AXMathExpressionText */


/* debug [class_header]: Header for AXMathExpressionText */
// The class instance for the [AXMathExpressionText] class.
var (
	AXMathExpressionTextClass     _AXMathExpressionTextClass
	AXMathExpressionTextClassOnce sync.Once
)

func getAXMathExpressionTextClass() _AXMathExpressionTextClass {
	AXMathExpressionTextClassOnce.Do(func() {
		AXMathExpressionTextClass = _AXMathExpressionTextClass{objc.GetClass("AXMathExpressionText")}
	})
	return AXMathExpressionTextClass
}

type _AXMathExpressionTextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXMathExpressionText */
// An interface definition for the [AXMathExpressionText] class.
type IAXMathExpressionText interface {
	IAXMathExpression
	
/* debug [class_interface_properties]: Properties for AXMathExpressionText */
	// properties:
	Content() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXMathExpressionText */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXMathExpressionText */
// Alloc allocates a new instance without initialization.
func (ac _AXMathExpressionTextClass) Alloc() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXMathExpressionTextClass) New() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXMathExpressionText) Init() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXMathExpressionText) Autorelease() AXMathExpressionText {
	rv := objc.Send[AXMathExpressionText](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXMathExpressionText creates a new AXMathExpressionText instance.
func NewAXMathExpressionText() AXMathExpressionText {
	return getAXMathExpressionTextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXMathExpressionText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText
type AXMathExpressionText struct {
	AXMathExpression
}

// AXMathExpressionTextFrom constructs a [AXMathExpressionText] from an unsafe.Pointer.
func AXMathExpressionTextFrom(ptr unsafe.Pointer) AXMathExpressionText {
	return AXMathExpressionText{
		AXMathExpression: AXMathExpressionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXMathExpressionText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/init(content:)
func NewAXMathExpressionTextWithContent(content objc.IObject /* cross-framework: NSString */) AXMathExpressionText {
	instance := getAXMathExpressionTextClass().Alloc()
	rv := objc.Send[AXMathExpressionText](instance.ID, objc.Sel("initWithContent:"), content)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXMathExpressionTextWithContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXMathExpressionText */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXMathExpressionText */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXMathExpressionText */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXMathExpressionText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXMathExpressionText/content
func (a_ AXMathExpressionText) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXMathExpressionText */


