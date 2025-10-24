// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSPageRule */

/* debug [class_header]: Header for DOMCSSPageRule */
// The class instance for the [DOMCSSPageRule] class.
var (
	DOMCSSPageRuleClass     _DOMCSSPageRuleClass
	DOMCSSPageRuleClassOnce sync.Once
)

func getDOMCSSPageRuleClass() _DOMCSSPageRuleClass {
	DOMCSSPageRuleClassOnce.Do(func() {
		DOMCSSPageRuleClass = _DOMCSSPageRuleClass{objc.GetClass("DOMCSSPageRule")}
	})
	return DOMCSSPageRuleClass
}

type _DOMCSSPageRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSPageRule */
// An interface definition for the [DOMCSSPageRule] class.
type IDOMCSSPageRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSPageRule */
	// properties:
	SelectorText() objc.IObject /* cross-framework: NSString */
	SetSelectorText(value objc.IObject /* cross-framework: NSString */)
	Style() IDOMCSSStyleDeclaration
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSPageRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSPageRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSPageRuleClass) Alloc() DOMCSSPageRule {
	rv := objc.Send[DOMCSSPageRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSPageRuleClass) New() DOMCSSPageRule {
	rv := objc.Send[DOMCSSPageRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSPageRule) Init() DOMCSSPageRule {
	rv := objc.Send[DOMCSSPageRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSPageRule) Autorelease() DOMCSSPageRule {
	rv := objc.Send[DOMCSSPageRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSPageRule creates a new DOMCSSPageRule instance.
func NewDOMCSSPageRule() DOMCSSPageRule {
	return getDOMCSSPageRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSPageRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPageRule
type DOMCSSPageRule struct {
	DOMCSSRule
}

// DOMCSSPageRuleFrom constructs a [DOMCSSPageRule] from an unsafe.Pointer.
func DOMCSSPageRuleFrom(ptr unsafe.Pointer) DOMCSSPageRule {
	return DOMCSSPageRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSPageRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSPageRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSPageRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSPageRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSPageRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPageRule/selectorText
func (d_ DOMCSSPageRule) SelectorText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("selectorText"))
	return rv
} /* debug [instance_properties/getter]: selectorText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPageRule/selectorText
func (d_ DOMCSSPageRule) SetSelectorText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectorText:"), value)
} /* debug [instance_properties/setter]: selectorText */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSPageRule/style
func (d_ DOMCSSPageRule) Style() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("style"))
	return rv
} /* debug [instance_properties/getter]: style */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSPageRule */
