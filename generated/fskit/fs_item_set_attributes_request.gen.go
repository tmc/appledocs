// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class FSItemSetAttributesRequest */


/* debug [class_header]: Header for FSItemSetAttributesRequest */
// The class instance for the [FSItemSetAttributesRequest] class.
var (
	FSItemSetAttributesRequestClass     _FSItemSetAttributesRequestClass
	FSItemSetAttributesRequestClassOnce sync.Once
)

func getFSItemSetAttributesRequestClass() _FSItemSetAttributesRequestClass {
	FSItemSetAttributesRequestClassOnce.Do(func() {
		FSItemSetAttributesRequestClass = _FSItemSetAttributesRequestClass{objc.GetClass("FSItemSetAttributesRequest")}
	})
	return FSItemSetAttributesRequestClass
}

type _FSItemSetAttributesRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSItemSetAttributesRequest */
// An interface definition for the [FSItemSetAttributesRequest] class.
type IFSItemSetAttributesRequest interface {
	IFSItemAttributes
	
/* debug [class_interface_properties]: Properties for FSItemSetAttributesRequest */
	// properties:
	ConsumedAttributes() FSItemAttribute
	SetConsumedAttributes(value FSItemAttribute)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSItemSetAttributesRequest */
	// methods:
	WasAttributeConsumed(attribute FSItemAttribute) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSItemSetAttributesRequest */
// Alloc allocates a new instance without initialization.
func (fc _FSItemSetAttributesRequestClass) Alloc() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSItemSetAttributesRequestClass) New() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItemSetAttributesRequest) Init() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItemSetAttributesRequest) Autorelease() FSItemSetAttributesRequest {
	rv := objc.Send[FSItemSetAttributesRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemSetAttributesRequest creates a new FSItemSetAttributesRequest instance.
func NewFSItemSetAttributesRequest() FSItemSetAttributesRequest {
	return getFSItemSetAttributesRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSItemSetAttributesRequest */
// A request to set attributes on an item.
//
// Methods that take attributes use this type to receive attribute values and to indicate which attributes they support. The various members of the parent type, , contain the values of the attributes to set. Modify the property to indicate which attributes your file system successfully used. FSKit calls the method to determine whether the file system successfully used a given attribute. Only set the attributes that your file system supports.


// A request to set attributes on an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest
type FSItemSetAttributesRequest struct {
	FSItemAttributes
}

// FSItemSetAttributesRequestFrom constructs a [FSItemSetAttributesRequest] from an unsafe.Pointer.
//
// A request to set attributes on an item.
func FSItemSetAttributesRequestFrom(ptr unsafe.Pointer) FSItemSetAttributesRequest {
	return FSItemSetAttributesRequest{
		FSItemAttributes: FSItemAttributesFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSItemSetAttributesRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSItemSetAttributesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSItemSetAttributesRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSItemSetAttributesRequest */

// A method that indicates whether the file system used the given attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/wasAttributeConsumed(_:)
func (f_ FSItemSetAttributesRequest) WasAttributeConsumed(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("wasAttributeConsumed:"), attribute)
	return rv
}/* debug [instance_methods/method]: WasAttributeConsumed */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSItemSetAttributesRequest */

// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/consumedAttributes
func (f_ FSItemSetAttributesRequest) ConsumedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](f_.ID, objc.Sel("consumedAttributes"))
	return rv
}/* debug [instance_properties/getter]: consumedAttributes */


// The attributes successfully used by the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/SetAttributesRequest/consumedAttributes
func (f_ FSItemSetAttributesRequest) SetConsumedAttributes(value FSItemAttribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConsumedAttributes:"), value)
}/* debug [instance_properties/setter]: consumedAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSItemSetAttributesRequest */



