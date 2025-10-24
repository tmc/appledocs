// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class packetSigning */


/* debug [class_header]: Header for packetSigning */
// The class instance for the [packetSigning] class.
var (
	PacketSigningClass     _packetSigningClass
	PacketSigningClassOnce sync.Once
)

func getpacketSigningClass() _packetSigningClass {
	PacketSigningClassOnce.Do(func() {
		PacketSigningClass = _packetSigningClass{objc.GetClass("packetSigning")}
	})
	return PacketSigningClass
}

type _packetSigningClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for packetSigning */
// An interface definition for the [packetSigning] class.
type IpacketSigning interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for packetSigning */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for packetSigning */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for packetSigning */
// Alloc allocates a new instance without initialization.
func (pc _packetSigningClass) Alloc() packetSigning {
	rv := objc.Send[packetSigning](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _packetSigningClass) New() packetSigning {
	rv := objc.Send[packetSigning](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ packetSigning) Init() packetSigning {
	rv := objc.Send[packetSigning](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ packetSigning) Autorelease() packetSigning {
	rv := objc.Send[packetSigning](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpacketSigning creates a new packetSigning instance.
func NewpacketSigning() packetSigning {
	return getpacketSigningClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for packetSigning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-c.ivar
type packetSigning struct {
	objectivec.Object
}

// packetSigningFrom constructs a [packetSigning] from an unsafe.Pointer.
func packetSigningFrom(ptr unsafe.Pointer) packetSigning {
	return packetSigning{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for packetSigning *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for packetSigning */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for packetSigning */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for packetSigning */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for packetSigning */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class packetSigning */



