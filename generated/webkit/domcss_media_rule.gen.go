// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSMediaRule */

/* debug [class_header]: Header for DOMCSSMediaRule */
// The class instance for the [DOMCSSMediaRule] class.
var (
	DOMCSSMediaRuleClass     _DOMCSSMediaRuleClass
	DOMCSSMediaRuleClassOnce sync.Once
)

func getDOMCSSMediaRuleClass() _DOMCSSMediaRuleClass {
	DOMCSSMediaRuleClassOnce.Do(func() {
		DOMCSSMediaRuleClass = _DOMCSSMediaRuleClass{objc.GetClass("DOMCSSMediaRule")}
	})
	return DOMCSSMediaRuleClass
}

type _DOMCSSMediaRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSMediaRule */
// An interface definition for the [DOMCSSMediaRule] class.
type IDOMCSSMediaRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSMediaRule */
	// properties:
	CssRules() IDOMCSSRuleList
	Media() IDOMMediaList
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSMediaRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSMediaRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSMediaRuleClass) Alloc() DOMCSSMediaRule {
	rv := objc.Send[DOMCSSMediaRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSMediaRuleClass) New() DOMCSSMediaRule {
	rv := objc.Send[DOMCSSMediaRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSMediaRule) Init() DOMCSSMediaRule {
	rv := objc.Send[DOMCSSMediaRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSMediaRule) Autorelease() DOMCSSMediaRule {
	rv := objc.Send[DOMCSSMediaRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSMediaRule creates a new DOMCSSMediaRule instance.
func NewDOMCSSMediaRule() DOMCSSMediaRule {
	return getDOMCSSMediaRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSMediaRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSMediaRule
type DOMCSSMediaRule struct {
	DOMCSSRule
}

// DOMCSSMediaRuleFrom constructs a [DOMCSSMediaRule] from an unsafe.Pointer.
func DOMCSSMediaRuleFrom(ptr unsafe.Pointer) DOMCSSMediaRule {
	return DOMCSSMediaRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSMediaRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSMediaRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSMediaRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSMediaRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSMediaRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSMediaRule/cssRules
func (d_ DOMCSSMediaRule) CssRules() IDOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](d_.ID, objc.Sel("cssRules"))
	return rv
} /* debug [instance_properties/getter]: cssRules */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSMediaRule/media
func (d_ DOMCSSMediaRule) Media() IDOMMediaList {
	rv := objc.Send[DOMMediaList](d_.ID, objc.Sel("media"))
	return rv
} /* debug [instance_properties/getter]: media */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSMediaRule */
