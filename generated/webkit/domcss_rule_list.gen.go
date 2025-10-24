// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMCSSRuleList */

/* debug [class_header]: Header for DOMCSSRuleList */
// The class instance for the [DOMCSSRuleList] class.
var (
	DOMCSSRuleListClass     _DOMCSSRuleListClass
	DOMCSSRuleListClassOnce sync.Once
)

func getDOMCSSRuleListClass() _DOMCSSRuleListClass {
	DOMCSSRuleListClassOnce.Do(func() {
		DOMCSSRuleListClass = _DOMCSSRuleListClass{objc.GetClass("DOMCSSRuleList")}
	})
	return DOMCSSRuleListClass
}

type _DOMCSSRuleListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMCSSRuleList */
// An interface definition for the [DOMCSSRuleList] class.
type IDOMCSSRuleList interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMCSSRuleList */
	// properties:
	Length() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMCSSRuleList */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMCSSRuleList */
// Alloc allocates a new instance without initialization.
func (dc _DOMCSSRuleListClass) Alloc() DOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMCSSRuleListClass) New() DOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMCSSRuleList) Init() DOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMCSSRuleList) Autorelease() DOMCSSRuleList {
	rv := objc.Send[DOMCSSRuleList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMCSSRuleList creates a new DOMCSSRuleList instance.
func NewDOMCSSRuleList() DOMCSSRuleList {
	return getDOMCSSRuleListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMCSSRuleList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRuleList
type DOMCSSRuleList struct {
	DOMObject
}

// DOMCSSRuleListFrom constructs a [DOMCSSRuleList] from an unsafe.Pointer.
func DOMCSSRuleListFrom(ptr unsafe.Pointer) DOMCSSRuleList {
	return DOMCSSRuleList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMCSSRuleList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMCSSRuleList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMCSSRuleList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMCSSRuleList */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMCSSRuleList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMCSSRuleList/length
func (d_ DOMCSSRuleList) Length() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("length"))
	return rv
} /* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMCSSRuleList */
