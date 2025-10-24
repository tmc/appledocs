// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [Protocol] class.
type IProtocol interface {
	IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/Protocol
type Protocol struct {
	Object
}

// ProtocolFrom constructs a [Protocol] from an unsafe.Pointer.
func ProtocolFrom(ptr unsafe.Pointer) Protocol {
	return Protocol{Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProtocolClass) Alloc() Protocol {
	rv := objc.Send[Protocol](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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





