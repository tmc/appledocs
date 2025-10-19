// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Port] class.
var (
	portClass     _PortClass
	portClassOnce sync.Once
)

func getPortClass() _PortClass {
	portClassOnce.Do(func() {
		portClass = _PortClass{objc.GetClass("NSPort")}
	})
	return portClass
}

type _PortClass struct {
	class objc.Class
}

// An interface definition for the [Port] class.
type IPort interface {
	objectivec.IObject
}

// An abstract class that represents a communication channel.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Port
type Port struct {
	objectivec.Object
}

// PortFrom constructs a [Port] from an unsafe.Pointer.
//
// An abstract class that represents a communication channel.
func PortFrom(ptr unsafe.Pointer) Port {
	return Port{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PortClass) Alloc() Port {
	rv := objc.Send[Port](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PortClass) New() Port {
	rv := objc.Send[Port](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Port) Init() Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Port) Autorelease() Port {
	rv := objc.Send[Port](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPort creates a new Port instance.
func NewPort() Port {
	return getPortClass().New()
}




