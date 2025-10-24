// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSFontFaceRule */

/* debug [class_header]: Header for DOMCSSFontFaceRule */
// The class instance for the [DOMCSSFontFaceRule] class.
var (
	DOMCSSFontFaceRuleClass     _DOMCSSFontFaceRuleClass
	DOMCSSFontFaceRuleClassOnce sync.Once
)

func getDOMCSSFontFaceRuleClass() _DOMCSSFontFaceRuleClass {
	DOMCSSFontFaceRuleClassOnce.Do(func() {
		DOMCSSFontFaceRuleClass = _DOMCSSFontFaceRuleClass{objc.GetClass("DOMCSSFontFaceRule")}
	})
	return DOMCSSFontFaceRuleClass
}

type _DOMCSSFontFaceRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSFontFaceRule */
// An interface definition for the [DOMCSSFontFaceRule] class.
type IDOMCSSFontFaceRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSFontFaceRule */
	// properties:
	Style() IDOMCSSStyleDeclaration
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSFontFaceRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSFontFaceRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSFontFaceRuleClass) Alloc() DOMCSSFontFaceRule {
	rv := objc.Send[DOMCSSFontFaceRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSFontFaceRuleClass) New() DOMCSSFontFaceRule {
	rv := objc.Send[DOMCSSFontFaceRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSFontFaceRule) Init() DOMCSSFontFaceRule {
	rv := objc.Send[DOMCSSFontFaceRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSFontFaceRule) Autorelease() DOMCSSFontFaceRule {
	rv := objc.Send[DOMCSSFontFaceRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSFontFaceRule creates a new DOMCSSFontFaceRule instance.
func NewDOMCSSFontFaceRule() DOMCSSFontFaceRule {
	return getDOMCSSFontFaceRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSFontFaceRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSFontFaceRule
type DOMCSSFontFaceRule struct {
	DOMCSSRule
}

// DOMCSSFontFaceRuleFrom constructs a [DOMCSSFontFaceRule] from an unsafe.Pointer.
func DOMCSSFontFaceRuleFrom(ptr unsafe.Pointer) DOMCSSFontFaceRule {
	return DOMCSSFontFaceRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSFontFaceRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSFontFaceRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSFontFaceRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSFontFaceRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSFontFaceRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSFontFaceRule/style
func (d_ DOMCSSFontFaceRule) Style() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("style"))
	return rv
} /* debug [instance_properties/getter]: style */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSFontFaceRule */
