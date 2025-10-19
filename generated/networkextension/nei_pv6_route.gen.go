// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEIPv6Route] class.
var (
	nEIPv6RouteClass     _NEIPv6RouteClass
	nEIPv6RouteClassOnce sync.Once
)

func getNEIPv6RouteClass() _NEIPv6RouteClass {
	nEIPv6RouteClassOnce.Do(func() {
		nEIPv6RouteClass = _NEIPv6RouteClass{objc.GetClass("NEIPv6Route")}
	})
	return nEIPv6RouteClass
}

type _NEIPv6RouteClass struct {
	class objc.Class
}

// An interface definition for the [NEIPv6Route] class.
type INEIPv6Route interface {
	objectivec.IObject
}

// The settings for an IPv6 route. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv6Route
type NEIPv6Route struct {
	objectivec.Object
}

// NEIPv6RouteFrom constructs a [NEIPv6Route] from an unsafe.Pointer.
//
// The settings for an IPv6 route.
func NEIPv6RouteFrom(ptr unsafe.Pointer) NEIPv6Route {
	return NEIPv6Route{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEIPv6RouteClass) Alloc() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEIPv6RouteClass) New() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv6Route) Init() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv6Route) Autorelease() NEIPv6Route {
	rv := objc.Send[NEIPv6Route](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv6Route creates a new NEIPv6Route instance.
func NewNEIPv6Route() NEIPv6Route {
	return getNEIPv6RouteClass().New()
}




