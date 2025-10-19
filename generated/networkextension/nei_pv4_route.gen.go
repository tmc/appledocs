// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEIPv4Route] class.
var (
	nEIPv4RouteClass     _NEIPv4RouteClass
	nEIPv4RouteClassOnce sync.Once
)

func getNEIPv4RouteClass() _NEIPv4RouteClass {
	nEIPv4RouteClassOnce.Do(func() {
		nEIPv4RouteClass = _NEIPv4RouteClass{objc.GetClass("NEIPv4Route")}
	})
	return nEIPv4RouteClass
}

type _NEIPv4RouteClass struct {
	class objc.Class
}

// An interface definition for the [NEIPv4Route] class.
type INEIPv4Route interface {
	objectivec.IObject
}

// The settings for an IPv4 route. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEIPv4Route
type NEIPv4Route struct {
	objectivec.Object
}

// NEIPv4RouteFrom constructs a [NEIPv4Route] from an unsafe.Pointer.
//
// The settings for an IPv4 route.
func NEIPv4RouteFrom(ptr unsafe.Pointer) NEIPv4Route {
	return NEIPv4Route{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEIPv4RouteClass) Alloc() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEIPv4RouteClass) New() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEIPv4Route) Init() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEIPv4Route) Autorelease() NEIPv4Route {
	rv := objc.Send[NEIPv4Route](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEIPv4Route creates a new NEIPv4Route instance.
func NewNEIPv4Route() NEIPv4Route {
	return getNEIPv4RouteClass().New()
}




