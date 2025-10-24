// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMCSSRule */


/* debug [class_header]: Header for DOMCSSRule */
// The class instance for the [DOMCSSRule] class.
var (
	DOMCSSRuleClass     _DOMCSSRuleClass
	DOMCSSRuleClassOnce sync.Once
)

func getDOMCSSRuleClass() _DOMCSSRuleClass {
	DOMCSSRuleClassOnce.Do(func() {
		DOMCSSRuleClass = _DOMCSSRuleClass{objc.GetClass("DOMCSSRule")}
	})
	return DOMCSSRuleClass
}

type _DOMCSSRuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMCSSRule */
// An interface definition for the [DOMCSSRule] class.
type IDOMCSSRule interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMCSSRule */
	// properties:
	CssText() objc.IObject /* cross-framework: NSString */
	SetCssText(value objc.IObject /* cross-framework: NSString */)
	ParentRule() IDOMCSSRule
	ParentStyleSheet() IDOMCSSStyleSheet
	Type() objectivec.IObject
	Parent() IDOMCSSRule
	SetParent(value IDOMCSSRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMCSSRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMCSSRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSRuleClass) Alloc() DOMCSSRule {
	rv := objc.Send[DOMCSSRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSRuleClass) New() DOMCSSRule {
	rv := objc.Send[DOMCSSRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSRule) Init() DOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSRule) Autorelease() DOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSRule creates a new DOMCSSRule instance.
func NewDOMCSSRule() DOMCSSRule {
	return getDOMCSSRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMCSSRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule
type DOMCSSRule struct {
	DOMObject
}

// DOMCSSRuleFrom constructs a [DOMCSSRule] from an unsafe.Pointer.
func DOMCSSRuleFrom(ptr unsafe.Pointer) DOMCSSRule {
	return DOMCSSRule{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMCSSRule *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMCSSRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMCSSRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMCSSRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMCSSRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule/cssText
func (d_ DOMCSSRule) CssText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("cssText"))
	return rv
}/* debug [instance_properties/getter]: cssText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule/cssText
func (d_ DOMCSSRule) SetCssText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCssText:"), value)
}/* debug [instance_properties/setter]: cssText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule/parent
func (d_ DOMCSSRule) ParentRule() IDOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("parentRule"))
	return rv
}/* debug [instance_properties/getter]: parentRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule/parentStyleSheet
func (d_ DOMCSSRule) ParentStyleSheet() IDOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](d_.ID, objc.Sel("parentStyleSheet"))
	return rv
}/* debug [instance_properties/getter]: parentStyleSheet */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRule/type
func (d_ DOMCSSRule) Type() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssrule/parent
func (d_ DOMCSSRule) Parent() IDOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/domcssrule/parent
func (d_ DOMCSSRule) SetParent(value IDOMCSSRule) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMCSSRule */



