// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSItemGetAttributesRequest */


/* debug [class_header]: Header for FSItemGetAttributesRequest */
// The class instance for the [FSItemGetAttributesRequest] class.
var (
	FSItemGetAttributesRequestClass     _FSItemGetAttributesRequestClass
	FSItemGetAttributesRequestClassOnce sync.Once
)

func getFSItemGetAttributesRequestClass() _FSItemGetAttributesRequestClass {
	FSItemGetAttributesRequestClassOnce.Do(func() {
		FSItemGetAttributesRequestClass = _FSItemGetAttributesRequestClass{objc.GetClass("FSItemGetAttributesRequest")}
	})
	return FSItemGetAttributesRequestClass
}

type _FSItemGetAttributesRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSItemGetAttributesRequest */
// An interface definition for the [FSItemGetAttributesRequest] class.
type IFSItemGetAttributesRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSItemGetAttributesRequest */
	// properties:
	WantedAttributes() FSItemAttribute
	SetWantedAttributes(value FSItemAttribute)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSItemGetAttributesRequest */
	// methods:
	IsAttributeWanted(attribute FSItemAttribute) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSItemGetAttributesRequest */
// Alloc allocates a new instance without initialization.
func (fc _FSItemGetAttributesRequestClass) Alloc() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSItemGetAttributesRequestClass) New() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSItemGetAttributesRequest) Init() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSItemGetAttributesRequest) Autorelease() FSItemGetAttributesRequest {
	rv := objc.Send[FSItemGetAttributesRequest](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSItemGetAttributesRequest creates a new FSItemGetAttributesRequest instance.
func NewFSItemGetAttributesRequest() FSItemGetAttributesRequest {
	return getFSItemGetAttributesRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSItemGetAttributesRequest */
// A request to get attributes from an item.
//
// Methods that retrieve attributes use this type and inspect the property to determine which attributes to provide. FSKit calls the method to determine whether the request requires a given attribute.


// A request to get attributes from an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest
type FSItemGetAttributesRequest struct {
	objectivec.Object
}

// FSItemGetAttributesRequestFrom constructs a [FSItemGetAttributesRequest] from an unsafe.Pointer.
//
// A request to get attributes from an item.
func FSItemGetAttributesRequestFrom(ptr unsafe.Pointer) FSItemGetAttributesRequest {
	return FSItemGetAttributesRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSItemGetAttributesRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSItemGetAttributesRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSItemGetAttributesRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSItemGetAttributesRequest */

// A method that indicates whether the request wants given attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/isAttributeWanted(_:)
func (f_ FSItemGetAttributesRequest) IsAttributeWanted(attribute FSItemAttribute) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("isAttributeWanted:"), attribute)
	return rv
}/* debug [instance_methods/method]: IsAttributeWanted */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSItemGetAttributesRequest */

// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/wantedAttributes
func (f_ FSItemGetAttributesRequest) WantedAttributes() FSItemAttribute {
	rv := objc.Send[FSItemAttribute](f_.ID, objc.Sel("wantedAttributes"))
	return rv
}/* debug [instance_properties/getter]: wantedAttributes */


// The attributes requested by the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSItem/GetAttributesRequest/wantedAttributes
func (f_ FSItemGetAttributesRequest) SetWantedAttributes(value FSItemAttribute) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWantedAttributes:"), value)
}/* debug [instance_properties/setter]: wantedAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSItemGetAttributesRequest */



