// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMXPathExpression */

/* debug [class_header]: Header for DOMXPathExpression */
// The class instance for the [DOMXPathExpression] class.
var (
	DOMXPathExpressionClass     _DOMXPathExpressionClass
	DOMXPathExpressionClassOnce sync.Once
)

func getDOMXPathExpressionClass() _DOMXPathExpressionClass {
	DOMXPathExpressionClassOnce.Do(func() {
		DOMXPathExpressionClass = _DOMXPathExpressionClass{objc.GetClass("DOMXPathExpression")}
	})
	return DOMXPathExpressionClass
}

type _DOMXPathExpressionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMXPathExpression */
// An interface definition for the [DOMXPathExpression] class.
type IDOMXPathExpression interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMXPathExpression */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMXPathExpression */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMXPathExpression */
// Alloc allocates a new instance without initialization.
func (dc _DOMXPathExpressionClass) Alloc() DOMXPathExpression {
	rv := objc.Send[DOMXPathExpression](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMXPathExpressionClass) New() DOMXPathExpression {
	rv := objc.Send[DOMXPathExpression](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMXPathExpression) Init() DOMXPathExpression {
	rv := objc.Send[DOMXPathExpression](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMXPathExpression) Autorelease() DOMXPathExpression {
	rv := objc.Send[DOMXPathExpression](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMXPathExpression creates a new DOMXPathExpression instance.
func NewDOMXPathExpression() DOMXPathExpression {
	return getDOMXPathExpressionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMXPathExpression */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMXPathExpression
type DOMXPathExpression struct {
	DOMObject
}

// DOMXPathExpressionFrom constructs a [DOMXPathExpression] from an unsafe.Pointer.
func DOMXPathExpressionFrom(ptr unsafe.Pointer) DOMXPathExpression {
	return DOMXPathExpression{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMXPathExpression */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMXPathExpression */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMXPathExpression */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMXPathExpression */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMXPathExpression */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMXPathExpression */
