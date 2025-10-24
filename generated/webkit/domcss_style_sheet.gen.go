// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSStyleSheet */

/* debug [class_header]: Header for DOMCSSStyleSheet */
// The class instance for the [DOMCSSStyleSheet] class.
var (
	DOMCSSStyleSheetClass     _DOMCSSStyleSheetClass
	DOMCSSStyleSheetClassOnce sync.Once
)

func getDOMCSSStyleSheetClass() _DOMCSSStyleSheetClass {
	DOMCSSStyleSheetClassOnce.Do(func() {
		DOMCSSStyleSheetClass = _DOMCSSStyleSheetClass{objc.GetClass("DOMCSSStyleSheet")}
	})
	return DOMCSSStyleSheetClass
}

type _DOMCSSStyleSheetClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSStyleSheet */
// An interface definition for the [DOMCSSStyleSheet] class.
type IDOMCSSStyleSheet interface {
	IDOMStyleSheet

	/* debug [class_interface_properties]: Properties for DOMCSSStyleSheet */
	// properties:
	CssRules() IDOMCSSRuleList
	OwnerRule() IDOMCSSRule
	Rules() IDOMCSSRuleList
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSStyleSheet */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSStyleSheet */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSStyleSheetClass) Alloc() DOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSStyleSheetClass) New() DOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSStyleSheet) Init() DOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSStyleSheet) Autorelease() DOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSStyleSheet creates a new DOMCSSStyleSheet instance.
func NewDOMCSSStyleSheet() DOMCSSStyleSheet {
	return getDOMCSSStyleSheetClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSStyleSheet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleSheet
type DOMCSSStyleSheet struct {
	DOMStyleSheet
}

// DOMCSSStyleSheetFrom constructs a [DOMCSSStyleSheet] from an unsafe.Pointer.
func DOMCSSStyleSheetFrom(ptr unsafe.Pointer) DOMCSSStyleSheet {
	return DOMCSSStyleSheet{
		DOMStyleSheet: DOMStyleSheetFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSStyleSheet */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSStyleSheet */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSStyleSheet */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSStyleSheet */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSStyleSheet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleSheet/cssRules
func (d_ DOMCSSStyleSheet) CssRules() IDOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](d_.ID, objc.Sel("cssRules"))
	return rv
} /* debug [instance_properties/getter]: cssRules */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleSheet/ownerRule
func (d_ DOMCSSStyleSheet) OwnerRule() IDOMCSSRule {
	rv := objc.Send[DOMCSSRule](d_.ID, objc.Sel("ownerRule"))
	return rv
} /* debug [instance_properties/getter]: ownerRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSStyleSheet/rules
func (d_ DOMCSSStyleSheet) Rules() IDOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](d_.ID, objc.Sel("rules"))
	return rv
} /* debug [instance_properties/getter]: rules */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSStyleSheet */
