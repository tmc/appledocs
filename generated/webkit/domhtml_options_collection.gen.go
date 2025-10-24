// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMHTMLOptionsCollection */


/* debug [class_header]: Header for DOMHTMLOptionsCollection */
// The class instance for the [DOMHTMLOptionsCollection] class.
var (
	DOMHTMLOptionsCollectionClass     _DOMHTMLOptionsCollectionClass
	DOMHTMLOptionsCollectionClassOnce sync.Once
)

func getDOMHTMLOptionsCollectionClass() _DOMHTMLOptionsCollectionClass {
	DOMHTMLOptionsCollectionClassOnce.Do(func() {
		DOMHTMLOptionsCollectionClass = _DOMHTMLOptionsCollectionClass{objc.GetClass("DOMHTMLOptionsCollection")}
	})
	return DOMHTMLOptionsCollectionClass
}

type _DOMHTMLOptionsCollectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLOptionsCollection */
// An interface definition for the [DOMHTMLOptionsCollection] class.
type IDOMHTMLOptionsCollection interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMHTMLOptionsCollection */
	// properties:
	Length() objectivec.IObject
	SetLength(value objectivec.IObject)
	SelectedIndex() int
	SetSelectedIndex(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLOptionsCollection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLOptionsCollection */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLOptionsCollectionClass) Alloc() DOMHTMLOptionsCollection {
	rv := objc.Send[DOMHTMLOptionsCollection](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLOptionsCollectionClass) New() DOMHTMLOptionsCollection {
	rv := objc.Send[DOMHTMLOptionsCollection](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLOptionsCollection) Init() DOMHTMLOptionsCollection {
	rv := objc.Send[DOMHTMLOptionsCollection](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLOptionsCollection) Autorelease() DOMHTMLOptionsCollection {
	rv := objc.Send[DOMHTMLOptionsCollection](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLOptionsCollection creates a new DOMHTMLOptionsCollection instance.
func NewDOMHTMLOptionsCollection() DOMHTMLOptionsCollection {
	return getDOMHTMLOptionsCollectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLOptionsCollection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionsCollection
type DOMHTMLOptionsCollection struct {
	DOMObject
}

// DOMHTMLOptionsCollectionFrom constructs a [DOMHTMLOptionsCollection] from an unsafe.Pointer.
func DOMHTMLOptionsCollectionFrom(ptr unsafe.Pointer) DOMHTMLOptionsCollection {
	return DOMHTMLOptionsCollection{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLOptionsCollection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLOptionsCollection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLOptionsCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLOptionsCollection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLOptionsCollection */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionsCollection/length
func (d_ DOMHTMLOptionsCollection) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionsCollection/length
func (d_ DOMHTMLOptionsCollection) SetLength(value objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLength:"), value)
}/* debug [instance_properties/setter]: length */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionsCollection/selectedIndex
func (d_ DOMHTMLOptionsCollection) SelectedIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLOptionsCollection/selectedIndex
func (d_ DOMHTMLOptionsCollection) SetSelectedIndex(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLOptionsCollection */



