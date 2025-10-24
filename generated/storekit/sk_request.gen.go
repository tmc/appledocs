// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKRequest */


/* debug [class_header]: Header for SKRequest */
// The class instance for the [Request] class.
var (
	RequestClass     _RequestClass
	RequestClassOnce sync.Once
)

func getRequestClass() _RequestClass {
	RequestClassOnce.Do(func() {
		RequestClass = _RequestClass{objc.GetClass("SKRequest")}
	})
	return RequestClass
}

type _RequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Request */
// An interface definition for the [Request] class.
type IRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Request */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Request */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Request */
// Alloc allocates a new instance without initialization.
func (rc _RequestClass) Alloc() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RequestClass) New() Request {
	rv := objc.Send[Request](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ Request) Init() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ Request) Autorelease() Request {
	rv := objc.Send[Request](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRequest creates a new Request instance.
func NewRequest() Request {
	return getRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Request */
// An abstract class that represents a request to the App Store.
//
// To make a request, initialize a subclass of —such as or —set the property, and call the method.


// An abstract class that represents a request to the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest
type Request struct {
	objectivec.Object
}

// RequestFrom constructs a [Request] from an unsafe.Pointer.
//
// An abstract class that represents a request to the App Store.
func RequestFrom(ptr unsafe.Pointer) Request {
	return Request{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Request *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Request */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Request */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Request */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Request */

// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/delegate
func (r_ Request) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKRequest/delegate
func (r_ Request) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKRequest */





