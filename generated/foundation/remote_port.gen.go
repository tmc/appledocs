// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [remotePort] class.
var (
	RemotePortClass     _remotePortClass
	RemotePortClassOnce sync.Once
)

func getremotePortClass() _remotePortClass {
	RemotePortClassOnce.Do(func() {
		RemotePortClass = _remotePortClass{objc.GetClass("remotePort")}
	})
	return RemotePortClass
}

type _remotePortClass struct {
	class objc.Class
}

// An interface definition for the [remotePort] class.
type IremotePort interface {
	objectivec.IObject
}

//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/remotePort
type remotePort struct {
	objectivec.Object
}

// remotePortFrom constructs a [remotePort] from an unsafe.Pointer.
func remotePortFrom(ptr unsafe.Pointer) remotePort {
	return remotePort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _remotePortClass) Alloc() remotePort {
	rv := objc.Send[remotePort](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _remotePortClass) New() remotePort {
	rv := objc.Send[remotePort](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ remotePort) Init() remotePort {
	rv := objc.Send[remotePort](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ remotePort) Autorelease() remotePort {
	rv := objc.Send[remotePort](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewremotePort creates a new remotePort instance.
func NewremotePort() remotePort {
	return getremotePortClass().New()
}




