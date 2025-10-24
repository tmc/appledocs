// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMNodeList */

/* debug [class_header]: Header for DOMNodeList */
// The class instance for the [DOMNodeList] class.
var (
	DOMNodeListClass     _DOMNodeListClass
	DOMNodeListClassOnce sync.Once
)

func getDOMNodeListClass() _DOMNodeListClass {
	DOMNodeListClassOnce.Do(func() {
		DOMNodeListClass = _DOMNodeListClass{objc.GetClass("DOMNodeList")}
	})
	return DOMNodeListClass
}

type _DOMNodeListClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DOMNodeList */
// An interface definition for the [DOMNodeList] class.
type IDOMNodeList interface {
	IDOMObject

	/* debug [class_interface_properties]: Properties for DOMNodeList */
	// properties:
	Length() unsafe.Pointer
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DOMNodeList */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DOMNodeList */
// Alloc allocates a new instance without initialization.
func (dc _DOMNodeListClass) Alloc() DOMNodeList {
	rv := objc.Send[DOMNodeList](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMNodeListClass) New() DOMNodeList {
	rv := objc.Send[DOMNodeList](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMNodeList) Init() DOMNodeList {
	rv := objc.Send[DOMNodeList](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMNodeList) Autorelease() DOMNodeList {
	rv := objc.Send[DOMNodeList](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMNodeList creates a new DOMNodeList instance.
func NewDOMNodeList() DOMNodeList {
	return getDOMNodeListClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DOMNodeList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeList
type DOMNodeList struct {
	DOMObject
}

// DOMNodeListFrom constructs a [DOMNodeList] from an unsafe.Pointer.
func DOMNodeListFrom(ptr unsafe.Pointer) DOMNodeList {
	return DOMNodeList{
		DOMObject: DOMObjectFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DOMNodeList */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DOMNodeList */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DOMNodeList */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DOMNodeList */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DOMNodeList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNodeList/length
func (d_ DOMNodeList) Length() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("length"))
	return rv
} /* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class DOMNodeList */
