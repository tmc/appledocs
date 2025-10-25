// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDistantObjectRequest */


/* debug [class_header]: Header for NSDistantObjectRequest */
// The class instance for the [DistantObjectRequest] class.
var (
	DistantObjectRequestClass     _DistantObjectRequestClass
	DistantObjectRequestClassOnce sync.Once
)

func getDistantObjectRequestClass() _DistantObjectRequestClass {
	DistantObjectRequestClassOnce.Do(func() {
		DistantObjectRequestClass = _DistantObjectRequestClass{objc.GetClass("NSDistantObjectRequest")}
	})
	return DistantObjectRequestClass
}

type _DistantObjectRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DistantObjectRequest */
// An interface definition for the [DistantObjectRequest] class.
type IDistantObjectRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DistantObjectRequest */
	// properties:
	Connection() IConnection
	Conversation() objc.ID
	Invocation() IInvocation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DistantObjectRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DistantObjectRequest */
// Alloc allocates a new instance without initialization.
func (dc _DistantObjectRequestClass) Alloc() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DistantObjectRequestClass) New() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DistantObjectRequest) Init() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DistantObjectRequest) Autorelease() DistantObjectRequest {
	rv := objc.Send[DistantObjectRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistantObjectRequest creates a new DistantObjectRequest instance.
func NewDistantObjectRequest() DistantObjectRequest {
	return getDistantObjectRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DistantObjectRequest */
// An object used by the distributed objects system to help handle invocations between different processes.
//
// Do not create objects directly. Unless you are getting involved with the low-level details of distributed objects, there should never be a need to access an . To intercept and possibly process requests yourself, implement the delegate method .


// An object used by the distributed objects system to help handle invocations between different processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest
type DistantObjectRequest struct {
	objectivec.Object
}

// DistantObjectRequestFrom constructs a [DistantObjectRequest] from an unsafe.Pointer.
//
// An object used by the distributed objects system to help handle invocations between different processes.
func DistantObjectRequestFrom(ptr unsafe.Pointer) DistantObjectRequest {
	return DistantObjectRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DistantObjectRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DistantObjectRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DistantObjectRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DistantObjectRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DistantObjectRequest */

// Returns the object involved in the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/connection
func (d_ DistantObjectRequest) Connection() IConnection {
	rv := objc.Send[Connection](d_.ID, objc.Sel("connection"))
	return rv
}/* debug [instance_properties/getter]: connection */


// Returns the token object representing the conversation in which the receiver was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/conversation
func (d_ DistantObjectRequest) Conversation() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("conversation"))
	return rv
}/* debug [instance_properties/getter]: conversation */


// Returns the object for the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDistantObjectRequest/invocation
func (d_ DistantObjectRequest) Invocation() IInvocation {
	rv := objc.Send[Invocation](d_.ID, objc.Sel("invocation"))
	return rv
}/* debug [instance_properties/getter]: invocation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDistantObjectRequest */



