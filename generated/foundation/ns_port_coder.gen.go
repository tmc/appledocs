// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PortCoder] class.
var (
	PortCoderClass     _PortCoderClass
	PortCoderClassOnce sync.Once
)

func getPortCoderClass() _PortCoderClass {
	PortCoderClassOnce.Do(func() {
		PortCoderClass = _PortCoderClass{objc.GetClass("NSPortCoder")}
	})
	return PortCoderClass
}

type _PortCoderClass struct {
	class objc.Class
}

// An interface definition for the [PortCoder] class.
type IPortCoder interface {
	ICoder
	// properties:
	// methods:
}

// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
//
// is a concrete subclass of used in the distributed objects system to transmit object proxies (and sometimes objects themselves) between objects. An instance is always created and used by an object; you should never need to explicitly create or use one directly yourself.


// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder
type PortCoder struct {
	Coder
}

// PortCoderFrom constructs a [PortCoder] from an unsafe.Pointer.
//
// A coder used to transmit object proxies (and sometimes objects themselves) between connections.
func PortCoderFrom(ptr unsafe.Pointer) PortCoder {
	return PortCoder{
		Coder: CoderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PortCoderClass) Alloc() PortCoder {
	rv := objc.Send[PortCoder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PortCoderClass) New() PortCoder {
	rv := objc.Send[PortCoder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PortCoder) Init() PortCoder {
	rv := objc.Send[PortCoder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PortCoder) Autorelease() PortCoder {
	rv := objc.Send[PortCoder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortCoder creates a new PortCoder instance.
func NewPortCoder() PortCoder {
	return getPortCoderClass().New()
}



// Initializes and returns an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder/initWithReceivePort:sendPort:components:
func NewPortCoderWithReceivePortSendPortComponents(rcvPort IPort, sndPort IPort, comps objectivec.IObject) PortCoder {
	instance := getPortCoderClass().Alloc()
	rv := objc.Send[PortCoder](instance.ID, objc.Sel("initWithReceivePort:sendPort:components:"), rcvPort, sndPort, comps)
	rv.Autorelease()
	return rv
}



// Creates and returns a new object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortCoder/portCoderWithReceivePort:sendPort:components:
func (pc _PortCoderClass) PortCoderWithReceivePortSendPortComponents(rcvPort IPort, sndPort IPort, comps objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("portCoderWithReceivePort:sendPort:components:"), rcvPort, sndPort, comps)
	return rv
}


