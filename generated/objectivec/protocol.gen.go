// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class Protocol */


/* debug [class_header]: Header for Protocol */
// The class instance for the [Protocol] class.
var (
	ProtocolClass     _ProtocolClass
	ProtocolClassOnce sync.Once
)

func getProtocolClass() _ProtocolClass {
	ProtocolClassOnce.Do(func() {
		ProtocolClass = _ProtocolClass{objc.GetClass("Protocol")}
	})
	return ProtocolClass
}

type _ProtocolClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Protocol */
// An interface definition for the [Protocol] class.
type IProtocol interface {
	IObject
	
/* debug [class_interface_properties]: Properties for Protocol */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Protocol */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Protocol */
// Alloc allocates a new instance without initialization.
func (pc _ProtocolClass) Alloc() Protocol {
	rv := objc.Send[Protocol](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ProtocolClass) New() Protocol {
	rv := objc.Send[Protocol](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Protocol) Init() Protocol {
	rv := objc.Send[Protocol](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Protocol) Autorelease() Protocol {
	rv := objc.Send[Protocol](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProtocol creates a new Protocol instance.
func NewProtocol() Protocol {
	return getProtocolClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Protocol */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Protocol
type Protocol struct {
	Object
}

// ProtocolFrom constructs a [Protocol] from an unsafe.Pointer.
func ProtocolFrom(ptr unsafe.Pointer) Protocol {
	return Protocol{Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Protocol *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Protocol */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Protocol */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Protocol */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Protocol */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class Protocol */





