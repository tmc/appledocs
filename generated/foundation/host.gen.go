// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Host] class.
var (
	hostClass     _HostClass
	hostClassOnce sync.Once
)

func getHostClass() _HostClass {
	hostClassOnce.Do(func() {
		hostClass = _HostClass{objc.GetClass("NSHost")}
	})
	return hostClass
}

type _HostClass struct {
	class objc.Class
}

// An interface definition for the [Host] class.
type IHost interface {
	objectivec.IObject
}

// A representation of an individual host on the network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Host
type Host struct {
	objectivec.Object
}

// HostFrom constructs a [Host] from an unsafe.Pointer.
//
// A representation of an individual host on the network.
func HostFrom(ptr unsafe.Pointer) Host {
	return Host{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HostClass) Alloc() Host {
	rv := objc.Send[Host](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HostClass) New() Host {
	rv := objc.Send[Host](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ Host) Init() Host {
	rv := objc.Send[Host](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ Host) Autorelease() Host {
	rv := objc.Send[Host](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHost creates a new Host instance.
func NewHost() Host {
	return getHostClass().New()
}




