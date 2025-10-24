// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class DOMBlob */


/* debug [class_header]: Header for DOMBlob */
// The class instance for the [DOMBlob] class.
var (
	DOMBlobClass     _DOMBlobClass
	DOMBlobClassOnce sync.Once
)

func getDOMBlobClass() _DOMBlobClass {
	DOMBlobClassOnce.Do(func() {
		DOMBlobClass = _DOMBlobClass{objc.GetClass("DOMBlob")}
	})
	return DOMBlobClass
}

type _DOMBlobClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMBlob */
// An interface definition for the [DOMBlob] class.
type IDOMBlob interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMBlob */
	// properties:
	Size() uint64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMBlob */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMBlob */
// Alloc allocates a new instance without initialization.
func (dc _DOMBlobClass) Alloc() DOMBlob {
	rv := objc.Send[DOMBlob](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMBlobClass) New() DOMBlob {
	rv := objc.Send[DOMBlob](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMBlob) Init() DOMBlob {
	rv := objc.Send[DOMBlob](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMBlob) Autorelease() DOMBlob {
	rv := objc.Send[DOMBlob](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMBlob creates a new DOMBlob instance.
func NewDOMBlob() DOMBlob {
	return getDOMBlobClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMBlob */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMBlob
type DOMBlob struct {
	DOMObject
}

// DOMBlobFrom constructs a [DOMBlob] from an unsafe.Pointer.
func DOMBlobFrom(ptr unsafe.Pointer) DOMBlob {
	return DOMBlob{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMBlob *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMBlob */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMBlob */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMBlob */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMBlob */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMBlob/size
func (d_ DOMBlob) Size() uint64 {
	rv := objc.Send[uint64](d_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMBlob */



