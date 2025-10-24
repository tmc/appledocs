// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMNamedNodeMap */


/* debug [class_header]: Header for DOMNamedNodeMap */
// The class instance for the [DOMNamedNodeMap] class.
var (
	DOMNamedNodeMapClass     _DOMNamedNodeMapClass
	DOMNamedNodeMapClassOnce sync.Once
)

func getDOMNamedNodeMapClass() _DOMNamedNodeMapClass {
	DOMNamedNodeMapClassOnce.Do(func() {
		DOMNamedNodeMapClass = _DOMNamedNodeMapClass{objc.GetClass("DOMNamedNodeMap")}
	})
	return DOMNamedNodeMapClass
}

type _DOMNamedNodeMapClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMNamedNodeMap */
// An interface definition for the [DOMNamedNodeMap] class.
type IDOMNamedNodeMap interface {
	IDOMObject
	
/* debug [class_interface_properties]: Properties for DOMNamedNodeMap */
	// properties:
	Length() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMNamedNodeMap */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMNamedNodeMap */
// Alloc allocates a new instance without initialization.
func (dc _DOMNamedNodeMapClass) Alloc() DOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMNamedNodeMapClass) New() DOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMNamedNodeMap) Init() DOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMNamedNodeMap) Autorelease() DOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMNamedNodeMap creates a new DOMNamedNodeMap instance.
func NewDOMNamedNodeMap() DOMNamedNodeMap {
	return getDOMNamedNodeMapClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMNamedNodeMap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNamedNodeMap
type DOMNamedNodeMap struct {
	DOMObject
}

// DOMNamedNodeMapFrom constructs a [DOMNamedNodeMap] from an unsafe.Pointer.
func DOMNamedNodeMapFrom(ptr unsafe.Pointer) DOMNamedNodeMap {
	return DOMNamedNodeMap{
		DOMObject: DOMObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMNamedNodeMap *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMNamedNodeMap */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMNamedNodeMap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMNamedNodeMap */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMNamedNodeMap */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMNamedNodeMap/length
func (d_ DOMNamedNodeMap) Length() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMNamedNodeMap */



