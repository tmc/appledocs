// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSUnknownRule */

/* debug [class_header]: Header for DOMCSSUnknownRule */
// The class instance for the [DOMCSSUnknownRule] class.
var (
	DOMCSSUnknownRuleClass     _DOMCSSUnknownRuleClass
	DOMCSSUnknownRuleClassOnce sync.Once
)

func getDOMCSSUnknownRuleClass() _DOMCSSUnknownRuleClass {
	DOMCSSUnknownRuleClassOnce.Do(func() {
		DOMCSSUnknownRuleClass = _DOMCSSUnknownRuleClass{objc.GetClass("DOMCSSUnknownRule")}
	})
	return DOMCSSUnknownRuleClass
}

type _DOMCSSUnknownRuleClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSUnknownRule */
// An interface definition for the [DOMCSSUnknownRule] class.
type IDOMCSSUnknownRule interface {
	IDOMCSSRule

	/* debug [class_interface_properties]: Properties for DOMCSSUnknownRule */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSUnknownRule */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSUnknownRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSUnknownRuleClass) Alloc() DOMCSSUnknownRule {
	rv := objc.Send[DOMCSSUnknownRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSUnknownRuleClass) New() DOMCSSUnknownRule {
	rv := objc.Send[DOMCSSUnknownRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSUnknownRule) Init() DOMCSSUnknownRule {
	rv := objc.Send[DOMCSSUnknownRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSUnknownRule) Autorelease() DOMCSSUnknownRule {
	rv := objc.Send[DOMCSSUnknownRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSUnknownRule creates a new DOMCSSUnknownRule instance.
func NewDOMCSSUnknownRule() DOMCSSUnknownRule {
	return getDOMCSSUnknownRuleClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSUnknownRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSUnknownRule
type DOMCSSUnknownRule struct {
	DOMCSSRule
}

// DOMCSSUnknownRuleFrom constructs a [DOMCSSUnknownRule] from an unsafe.Pointer.
func DOMCSSUnknownRuleFrom(ptr unsafe.Pointer) DOMCSSUnknownRule {
	return DOMCSSUnknownRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSUnknownRule */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSUnknownRule */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSUnknownRule */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSUnknownRule */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSUnknownRule */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSUnknownRule */
