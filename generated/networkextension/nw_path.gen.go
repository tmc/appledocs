// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [NWPath] class.
var (
	NWPathClass     _NWPathClass
	NWPathClassOnce sync.Once
)

func getNWPathClass() _NWPathClass {
	NWPathClassOnce.Do(func() {
		NWPathClass = _NWPathClass{objc.GetClass("NWPath")}
	})
	return NWPathClass
}

type _NWPathClass struct {
	class objc.Class
}

// An interface definition for the [NWPath] class.
type INWPath interface {
	objectivec.IObject
}

// The path made by a network connection, including information about its viability.
//
// For example, if the path status is , then a connection attempt will be made. When attached to a specific connection, a path takes all of the connection parameters into account. For example, if the route for a connection changes or is removed, the path will reflect that change. Note that every path is evaluated within the context of the process it is running in, and may be different across processes. is a static object, and properties of the path will never change. To monitor changing network status, use Key-Value Observing (KVO) to watch a path property on another object. For information about KVO, see .
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NWPath
type NWPath struct {
	objectivec.Object
}

// NWPathFrom constructs a [NWPath] from an unsafe.Pointer.
//
// The path made by a network connection, including information about its viability.
func NWPathFrom(ptr unsafe.Pointer) NWPath {
	return NWPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NWPathClass) Alloc() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NWPathClass) New() NWPath {
	rv := objc.Send[NWPath](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NWPath) Init() NWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NWPath) Autorelease() NWPath {
	rv := objc.Send[NWPath](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWPath creates a new NWPath instance.
func NewNWPath() NWPath {
	return getNWPathClass().New()
}




