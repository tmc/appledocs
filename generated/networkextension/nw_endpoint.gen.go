// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NWEndpoint */


/* debug [class_header]: Header for NWEndpoint */
// The class instance for the [NWEndpoint] class.
var (
	NWEndpointClass     _NWEndpointClass
	NWEndpointClassOnce sync.Once
)

func getNWEndpointClass() _NWEndpointClass {
	NWEndpointClassOnce.Do(func() {
		NWEndpointClass = _NWEndpointClass{objc.GetClass("NWEndpoint")}
	})
	return NWEndpointClass
}

type _NWEndpointClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NWEndpoint */
// An interface definition for the [NWEndpoint] class.
type INWEndpoint interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NWEndpoint */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NWEndpoint */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NWEndpoint */
// Alloc allocates a new instance without initialization.
func (nc _NWEndpointClass) Alloc() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NWEndpointClass) New() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWEndpoint) Init() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWEndpoint) Autorelease() NWEndpoint {
	rv := objc.Send[NWEndpoint](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWEndpoint creates a new NWEndpoint instance.
func NewNWEndpoint() NWEndpoint {
	return getNWEndpointClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NWEndpoint */
// An abstract base class, shared by or , that represents the source or destination of a network connection.
//
// All endpoint objects are static collections of parameters that describe a network resource. They do not directly provide any resolution services, but instead must be used with other classes to be resolved and create connections.


// An abstract base class, shared by or , that represents the source or destination of a network connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWEndpoint
type NWEndpoint struct {
	objectivec.Object
}

// NWEndpointFrom constructs a [NWEndpoint] from an unsafe.Pointer.
//
// An abstract base class, shared by or , that represents the source or destination of a network connection.
func NWEndpointFrom(ptr unsafe.Pointer) NWEndpoint {
	return NWEndpoint{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NWEndpoint *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NWEndpoint */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NWEndpoint */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NWEndpoint */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NWEndpoint */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NWEndpoint */



