// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSStyleRule */

/* debug [class_header]: Header for DOMCSSStyleRule */
// The class instance for the [DOMCSSStyleRule] class.
var (
	DOMCSSStyleRuleClass     _DOMCSSStyleRuleClass
	DOMCSSStyleRuleClassOnce sync.Once
)

func getDOMCSSStyleRuleClass() _DOMCSSStyleRuleClass {
	DOMCSSStyleRuleClassOnce.Do(func() {
		DOMCSSStyleRuleClass = _DOMCSSStyleRuleClass{objc.GetClass("DOMCSSStyleRule")}
	})
	return DOMCSSStyleRuleClass
}

type _DOMCSSStyleRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSStyleRule */
// An interface definition for the [DOMCSSStyleRule] class.
type IDOMCSSStyleRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSStyleRule */
	// properties:
	SelectorText() objc.IObject /* cross-framework: NSString */
	SetSelectorText(value objc.IObject /* cross-framework: NSString */)
	Style() IDOMCSSStyleDeclaration
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSStyleRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSStyleRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSStyleRuleClass) Alloc() DOMCSSStyleRule {
	rv := objc.Send[DOMCSSStyleRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSStyleRuleClass) New() DOMCSSStyleRule {
	rv := objc.Send[DOMCSSStyleRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSStyleRule) Init() DOMCSSStyleRule {
	rv := objc.Send[DOMCSSStyleRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSStyleRule) Autorelease() DOMCSSStyleRule {
	rv := objc.Send[DOMCSSStyleRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSStyleRule creates a new DOMCSSStyleRule instance.
func NewDOMCSSStyleRule() DOMCSSStyleRule {
	return getDOMCSSStyleRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSStyleRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleRule
type DOMCSSStyleRule struct {
	DOMCSSRule
}

// DOMCSSStyleRuleFrom constructs a [DOMCSSStyleRule] from an unsafe.Pointer.
func DOMCSSStyleRuleFrom(ptr unsafe.Pointer) DOMCSSStyleRule {
	return DOMCSSStyleRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSStyleRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSStyleRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSStyleRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSStyleRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSStyleRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleRule/selectorText
func (d_ DOMCSSStyleRule) SelectorText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("selectorText"))
	return rv
} /* debug [instance_properties/getter]: selectorText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleRule/selectorText
func (d_ DOMCSSStyleRule) SetSelectorText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectorText:"), value)
} /* debug [instance_properties/setter]: selectorText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleRule/style
func (d_ DOMCSSStyleRule) Style() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("style"))
	return rv
} /* debug [instance_properties/getter]: style */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSStyleRule */
