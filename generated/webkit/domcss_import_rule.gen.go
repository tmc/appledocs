// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMCSSImportRule */


/* debug [class_header]: Header for DOMCSSImportRule */
// The class instance for the [DOMCSSImportRule] class.
var (
	DOMCSSImportRuleClass     _DOMCSSImportRuleClass
	DOMCSSImportRuleClassOnce sync.Once
)

func getDOMCSSImportRuleClass() _DOMCSSImportRuleClass {
	DOMCSSImportRuleClassOnce.Do(func() {
		DOMCSSImportRuleClass = _DOMCSSImportRuleClass{objc.GetClass("DOMCSSImportRule")}
	})
	return DOMCSSImportRuleClass
}

type _DOMCSSImportRuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMCSSImportRule */
// An interface definition for the [DOMCSSImportRule] class.
type IDOMCSSImportRule interface {
	IDOMCSSRule
	
/* debug [class_interface_properties]: Properties for DOMCSSImportRule */
	// properties:
	Href() objc.IObject /* cross-framework: NSString */
	Media() IDOMMediaList
	StyleSheet() IDOMCSSStyleSheet
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMCSSImportRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMCSSImportRule */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSImportRuleClass) Alloc() DOMCSSImportRule {
	rv := objc.Send[DOMCSSImportRule](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSImportRuleClass) New() DOMCSSImportRule {
	rv := objc.Send[DOMCSSImportRule](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSImportRule) Init() DOMCSSImportRule {
	rv := objc.Send[DOMCSSImportRule](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSImportRule) Autorelease() DOMCSSImportRule {
	rv := objc.Send[DOMCSSImportRule](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSImportRule creates a new DOMCSSImportRule instance.
func NewDOMCSSImportRule() DOMCSSImportRule {
	return getDOMCSSImportRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMCSSImportRule */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSImportRule
type DOMCSSImportRule struct {
	DOMCSSRule
}

// DOMCSSImportRuleFrom constructs a [DOMCSSImportRule] from an unsafe.Pointer.
func DOMCSSImportRuleFrom(ptr unsafe.Pointer) DOMCSSImportRule {
	return DOMCSSImportRule{
		DOMCSSRule: DOMCSSRuleFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMCSSImportRule *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMCSSImportRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMCSSImportRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMCSSImportRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMCSSImportRule */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSImportRule/href
func (d_ DOMCSSImportRule) Href() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("href"))
	return rv
}/* debug [instance_properties/getter]: href */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSImportRule/media
func (d_ DOMCSSImportRule) Media() IDOMMediaList {
	rv := objc.Send[DOMMediaList](d_.ID, objc.Sel("media"))
	return rv
}/* debug [instance_properties/getter]: media */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSImportRule/styleSheet
func (d_ DOMCSSImportRule) StyleSheet() IDOMCSSStyleSheet {
	rv := objc.Send[DOMCSSStyleSheet](d_.ID, objc.Sel("styleSheet"))
	return rv
}/* debug [instance_properties/getter]: styleSheet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMCSSImportRule */



