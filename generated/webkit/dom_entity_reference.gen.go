// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMEntityReference */

/* debug [class_header]: Header for DOMEntityReference */
// The class instance for the [DOMEntityReference] class.
var (
	DOMEntityReferenceClass     _DOMEntityReferenceClass
	DOMEntityReferenceClassOnce sync.Once
)

func getDOMEntityReferenceClass() _DOMEntityReferenceClass {
	DOMEntityReferenceClassOnce.Do(func() {
		DOMEntityReferenceClass = _DOMEntityReferenceClass{objc.GetClass("DOMEntityReference")}
	})
	return DOMEntityReferenceClass
}

type _DOMEntityReferenceClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMEntityReference */
// An interface definition for the [DOMEntityReference] class.
type IDOMEntityReference interface {
	IDOMNode

	/* debug [class_interface_properties]: Properties for DOMEntityReference */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMEntityReference */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMEntityReference */
// Alloc allocates a new instance without initialization.
func (dc _DOMEntityReferenceClass) Alloc() DOMEntityReference {
	rv := objc.Send[DOMEntityReference](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMEntityReferenceClass) New() DOMEntityReference {
	rv := objc.Send[DOMEntityReference](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMEntityReference) Init() DOMEntityReference {
	rv := objc.Send[DOMEntityReference](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMEntityReference) Autorelease() DOMEntityReference {
	rv := objc.Send[DOMEntityReference](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMEntityReference creates a new DOMEntityReference instance.
func NewDOMEntityReference() DOMEntityReference {
	return getDOMEntityReferenceClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMEntityReference */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEntityReference
type DOMEntityReference struct {
	DOMNode
}

// DOMEntityReferenceFrom constructs a [DOMEntityReference] from an unsafe.Pointer.
func DOMEntityReferenceFrom(ptr unsafe.Pointer) DOMEntityReference {
	return DOMEntityReference{
		DOMNode: DOMNodeFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMEntityReference */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMEntityReference */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMEntityReference */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMEntityReference */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMEntityReference */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMEntityReference */
