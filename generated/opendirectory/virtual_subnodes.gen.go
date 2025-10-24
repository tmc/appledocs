// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [virtualSubnodes] class.
var (
	VirtualSubnodesClass     _virtualSubnodesClass
	VirtualSubnodesClassOnce sync.Once
)

func getvirtualSubnodesClass() _virtualSubnodesClass {
	VirtualSubnodesClassOnce.Do(func() {
		VirtualSubnodesClass = _virtualSubnodesClass{objc.GetClass("virtualSubnodes")}
	})
	return VirtualSubnodesClass
}

type _virtualSubnodesClass struct {
	class objc.Class
}

// An interface definition for the [virtualSubnodes] class.
type IvirtualSubnodes interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-c.ivar
type virtualSubnodes struct {
	objectivec.Object
}

// virtualSubnodesFrom constructs a [virtualSubnodes] from an unsafe.Pointer.
func virtualSubnodesFrom(ptr unsafe.Pointer) virtualSubnodes {
	return virtualSubnodes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _virtualSubnodesClass) Alloc() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _virtualSubnodesClass) New() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ virtualSubnodes) Init() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ virtualSubnodes) Autorelease() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewvirtualSubnodes creates a new virtualSubnodes instance.
func NewvirtualSubnodes() virtualSubnodes {
	return getvirtualSubnodesClass().New()
}




