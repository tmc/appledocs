// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMDocumentFragment */


/* debug [class_header]: Header for DOMDocumentFragment */
// The class instance for the [DOMDocumentFragment] class.
var (
	DOMDocumentFragmentClass     _DOMDocumentFragmentClass
	DOMDocumentFragmentClassOnce sync.Once
)

func getDOMDocumentFragmentClass() _DOMDocumentFragmentClass {
	DOMDocumentFragmentClassOnce.Do(func() {
		DOMDocumentFragmentClass = _DOMDocumentFragmentClass{objc.GetClass("DOMDocumentFragment")}
	})
	return DOMDocumentFragmentClass
}

type _DOMDocumentFragmentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMDocumentFragment */
// An interface definition for the [DOMDocumentFragment] class.
type IDOMDocumentFragment interface {
	IDOMNode
	
/* debug [class_interface_properties]: Properties for DOMDocumentFragment */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMDocumentFragment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMDocumentFragment */
// Alloc allocates a new instance without initialization.
func (dc _DOMDocumentFragmentClass) Alloc() DOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMDocumentFragmentClass) New() DOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMDocumentFragment) Init() DOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMDocumentFragment) Autorelease() DOMDocumentFragment {
	rv := objc.Send[DOMDocumentFragment](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMDocumentFragment creates a new DOMDocumentFragment instance.
func NewDOMDocumentFragment() DOMDocumentFragment {
	return getDOMDocumentFragmentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMDocumentFragment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentFragment
type DOMDocumentFragment struct {
	DOMNode
}

// DOMDocumentFragmentFrom constructs a [DOMDocumentFragment] from an unsafe.Pointer.
func DOMDocumentFragmentFrom(ptr unsafe.Pointer) DOMDocumentFragment {
	return DOMDocumentFragment{
		DOMNode: DOMNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMDocumentFragment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMDocumentFragment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMDocumentFragment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMDocumentFragment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMDocumentFragment */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMDocumentFragment */



