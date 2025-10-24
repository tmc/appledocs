// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMEntity */


/* debug [class_header]: Header for DOMEntity */
// The class instance for the [DOMEntity] class.
var (
	DOMEntityClass     _DOMEntityClass
	DOMEntityClassOnce sync.Once
)

func getDOMEntityClass() _DOMEntityClass {
	DOMEntityClassOnce.Do(func() {
		DOMEntityClass = _DOMEntityClass{objc.GetClass("DOMEntity")}
	})
	return DOMEntityClass
}

type _DOMEntityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMEntity */
// An interface definition for the [DOMEntity] class.
type IDOMEntity interface {
	IDOMNode
	
/* debug [class_interface_properties]: Properties for DOMEntity */
	// properties:
	NotationName() objc.IObject /* cross-framework: NSString */
	PublicId() objc.IObject /* cross-framework: NSString */
	SystemId() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMEntity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMEntity */
// Alloc allocates a new instance without initialization.
func (dc _DOMEntityClass) Alloc() DOMEntity {
	rv := objc.Send[DOMEntity](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMEntityClass) New() DOMEntity {
	rv := objc.Send[DOMEntity](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMEntity) Init() DOMEntity {
	rv := objc.Send[DOMEntity](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMEntity) Autorelease() DOMEntity {
	rv := objc.Send[DOMEntity](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMEntity creates a new DOMEntity instance.
func NewDOMEntity() DOMEntity {
	return getDOMEntityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMEntity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEntity
type DOMEntity struct {
	DOMNode
}

// DOMEntityFrom constructs a [DOMEntity] from an unsafe.Pointer.
func DOMEntityFrom(ptr unsafe.Pointer) DOMEntity {
	return DOMEntity{
		DOMNode: DOMNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMEntity *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMEntity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMEntity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMEntity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMEntity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEntity/notationName
func (d_ DOMEntity) NotationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("notationName"))
	return rv
}/* debug [instance_properties/getter]: notationName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEntity/publicId
func (d_ DOMEntity) PublicId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("publicId"))
	return rv
}/* debug [instance_properties/getter]: publicId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMEntity/systemId
func (d_ DOMEntity) SystemId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("systemId"))
	return rv
}/* debug [instance_properties/getter]: systemId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMEntity */



