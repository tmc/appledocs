// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEURLFilter] class.
var (
	nEURLFilterClass     _NEURLFilterClass
	nEURLFilterClassOnce sync.Once
)

func getNEURLFilterClass() _NEURLFilterClass {
	nEURLFilterClassOnce.Do(func() {
		nEURLFilterClass = _NEURLFilterClass{objc.GetClass("NEURLFilter")}
	})
	return nEURLFilterClass
}

type _NEURLFilterClass struct {
	class objc.Class
}

// An interface definition for the [NEURLFilter] class.
type INEURLFilter interface {
	objectivec.IObject
}

// A class used to voluntarily validate URLs for apps that don’t use WebKit or the URL session API.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEURLFilter
type NEURLFilter struct {
	objectivec.Object
}

// NEURLFilterFrom constructs a [NEURLFilter] from an unsafe.Pointer.
//
// A class used to voluntarily validate URLs for apps that don’t use WebKit or the URL session API.
func NEURLFilterFrom(ptr unsafe.Pointer) NEURLFilter {
	return NEURLFilter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEURLFilterClass) Alloc() NEURLFilter {
	rv := objc.Send[NEURLFilter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEURLFilterClass) New() NEURLFilter {
	rv := objc.Send[NEURLFilter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEURLFilter) Init() NEURLFilter {
	rv := objc.Send[NEURLFilter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEURLFilter) Autorelease() NEURLFilter {
	rv := objc.Send[NEURLFilter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEURLFilter creates a new NEURLFilter instance.
func NewNEURLFilter() NEURLFilter {
	return getNEURLFilterClass().New()
}




