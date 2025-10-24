// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [packetSigning] class.
type IpacketSigning interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-c.ivar
type packetSigning struct {
	objectivec.Object
}

// packetSigningFrom constructs a [packetSigning] from an unsafe.Pointer.
func packetSigningFrom(ptr unsafe.Pointer) packetSigning {
	return packetSigning{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _packetSigningClass) Alloc() packetSigning {
	rv := objc.Send[packetSigning](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




