// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class packetEncryption */


/* debug [class_header]: Header for packetEncryption */
// The class instance for the [packetEncryption] class.
var (
	PacketEncryptionClass     _packetEncryptionClass
	PacketEncryptionClassOnce sync.Once
)

func getpacketEncryptionClass() _packetEncryptionClass {
	PacketEncryptionClassOnce.Do(func() {
		PacketEncryptionClass = _packetEncryptionClass{objc.GetClass("packetEncryption")}
	})
	return PacketEncryptionClass
}

type _packetEncryptionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for packetEncryption */
// An interface definition for the [packetEncryption] class.
type IpacketEncryption interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for packetEncryption */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for packetEncryption */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for packetEncryption */
// Alloc allocates a new instance without initialization.
func (pc _packetEncryptionClass) Alloc() packetEncryption {
	rv := objc.Send[packetEncryption](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _packetEncryptionClass) New() packetEncryption {
	rv := objc.Send[packetEncryption](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ packetEncryption) Init() packetEncryption {
	rv := objc.Send[packetEncryption](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ packetEncryption) Autorelease() packetEncryption {
	rv := objc.Send[packetEncryption](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpacketEncryption creates a new packetEncryption instance.
func NewpacketEncryption() packetEncryption {
	return getpacketEncryptionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for packetEncryption */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-c.ivar
type packetEncryption struct {
	objectivec.Object
}

// packetEncryptionFrom constructs a [packetEncryption] from an unsafe.Pointer.
func packetEncryptionFrom(ptr unsafe.Pointer) packetEncryption {
	return packetEncryption{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for packetEncryption *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for packetEncryption */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for packetEncryption */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for packetEncryption */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for packetEncryption */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class packetEncryption */



