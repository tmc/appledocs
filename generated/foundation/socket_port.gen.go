// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SocketPort] class.
var (
	socketPortClass     _SocketPortClass
	socketPortClassOnce sync.Once
)

func getSocketPortClass() _SocketPortClass {
	socketPortClassOnce.Do(func() {
		socketPortClass = _SocketPortClass{objc.GetClass("NSSocketPort")}
	})
	return socketPortClass
}

type _SocketPortClass struct {
	class objc.Class
}

// An interface definition for the [SocketPort] class.
type ISocketPort interface {
	IPort
}

// A port that represents a BSD socket.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/SocketPort
type SocketPort struct {
	Port
}

// SocketPortFrom constructs a [SocketPort] from an unsafe.Pointer.
//
// A port that represents a BSD socket.
func SocketPortFrom(ptr unsafe.Pointer) SocketPort {
	return SocketPort{
		Port: PortFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SocketPortClass) Alloc() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SocketPortClass) New() SocketPort {
	rv := objc.Send[SocketPort](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SocketPort) Init() SocketPort {
	rv := objc.Send[SocketPort](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SocketPort) Autorelease() SocketPort {
	rv := objc.Send[SocketPort](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSocketPort creates a new SocketPort instance.
func NewSocketPort() SocketPort {
	return getSocketPortClass().New()
}




