// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSCharsetRule */

/* debug [class_header]: Header for DOMCSSCharsetRule */
// The class instance for the [DOMCSSCharsetRule] class.
var (
	DOMCSSCharsetRuleClass     _DOMCSSCharsetRuleClass
	DOMCSSCharsetRuleClassOnce sync.Once
)

func getDOMCSSCharsetRuleClass() _DOMCSSCharsetRuleClass {
	DOMCSSCharsetRuleClassOnce.Do(func() {
		DOMCSSCharsetRuleClass = _DOMCSSCharsetRuleClass{objc.GetClass("DOMCSSCharsetRule")}
	})
	return DOMCSSCharsetRuleClass
}

type _DOMCSSCharsetRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSCharsetRule */
// An interface definition for the [DOMCSSCharsetRule] class.
type IDOMCSSCharsetRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSCharsetRule */
	// properties:
	Encoding() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSCharsetRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSCharsetRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSCharsetRuleClass) Alloc() DOMCSSCharsetRule {
	rv := objc.Send[DOMCSSCharsetRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSCharsetRuleClass) New() DOMCSSCharsetRule {
	rv := objc.Send[DOMCSSCharsetRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSCharsetRule) Init() DOMCSSCharsetRule {
	rv := objc.Send[DOMCSSCharsetRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSCharsetRule) Autorelease() DOMCSSCharsetRule {
	rv := objc.Send[DOMCSSCharsetRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSCharsetRule creates a new DOMCSSCharsetRule instance.
func NewDOMCSSCharsetRule() DOMCSSCharsetRule {
	return getDOMCSSCharsetRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSCharsetRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSCharsetRule
type DOMCSSCharsetRule struct {
	DOMCSSRule
}

// DOMCSSCharsetRuleFrom constructs a [DOMCSSCharsetRule] from an unsafe.Pointer.
func DOMCSSCharsetRuleFrom(ptr unsafe.Pointer) DOMCSSCharsetRule {
	return DOMCSSCharsetRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSCharsetRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSCharsetRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSCharsetRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSCharsetRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSCharsetRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSCharsetRule/encoding
func (d_ DOMCSSCharsetRule) Encoding() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("encoding"))
	return rv
} /* debug [instance_properties/getter]: encoding */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSCharsetRule */
