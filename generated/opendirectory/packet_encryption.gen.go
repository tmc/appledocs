// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [packetEncryption] class.
type IpacketEncryption interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-c.ivar
type packetEncryption struct {
	objectivec.Object
}

// packetEncryptionFrom constructs a [packetEncryption] from an unsafe.Pointer.
func packetEncryptionFrom(ptr unsafe.Pointer) packetEncryption {
	return packetEncryption{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _packetEncryptionClass) Alloc() packetEncryption {
	rv := objc.Send[packetEncryption](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




