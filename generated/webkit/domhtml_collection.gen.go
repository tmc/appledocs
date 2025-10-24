// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMHTMLCollection */


/* debug [class_header]: Header for DOMHTMLCollection */
// The class instance for the [DOMHTMLCollection] class.
var (
	DOMHTMLCollectionClass     _DOMHTMLCollectionClass
	DOMHTMLCollectionClassOnce sync.Once
)

func getDOMHTMLCollectionClass() _DOMHTMLCollectionClass {
	DOMHTMLCollectionClassOnce.Do(func() {
		DOMHTMLCollectionClass = _DOMHTMLCollectionClass{objc.GetClass("DOMHTMLCollection")}
	})
	return DOMHTMLCollectionClass
}

type _DOMHTMLCollectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLCollection */
// An interface definition for the [DOMHTMLCollection] class.
type IDOMHTMLCollection interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMHTMLCollection */
	// properties:
	Length() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLCollection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLCollection */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLCollectionClass) Alloc() DOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLCollectionClass) New() DOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLCollection) Init() DOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLCollection) Autorelease() DOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLCollection creates a new DOMHTMLCollection instance.
func NewDOMHTMLCollection() DOMHTMLCollection {
	return getDOMHTMLCollectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLCollection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLCollection
type DOMHTMLCollection struct {
	DOMObject
}

// DOMHTMLCollectionFrom constructs a [DOMHTMLCollection] from an unsafe.Pointer.
func DOMHTMLCollectionFrom(ptr unsafe.Pointer) DOMHTMLCollection {
	return DOMHTMLCollection{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLCollection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLCollection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLCollection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLCollection */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLCollection/length
func (d_ DOMHTMLCollection) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLCollection */



