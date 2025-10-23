// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [containsMipmaps] class.
var (
	ContainsMipmapsClass     _containsMipmapsClass
	ContainsMipmapsClassOnce sync.Once
)

func getcontainsMipmapsClass() _containsMipmapsClass {
	ContainsMipmapsClassOnce.Do(func() {
		ContainsMipmapsClass = _containsMipmapsClass{objc.GetClass("containsMipmaps")}
	})
	return ContainsMipmapsClass
}

type _containsMipmapsClass struct {
	class objc.Class
}

// An interface definition for the [containsMipmaps] class.
type IcontainsMipmaps interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/containsMipmaps-c.ivar
type containsMipmaps struct {
	objectivec.Object
}

// containsMipmapsFrom constructs a [containsMipmaps] from an unsafe.Pointer.
func containsMipmapsFrom(ptr unsafe.Pointer) containsMipmaps {
	return containsMipmaps{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _containsMipmapsClass) Alloc() containsMipmaps {
	rv := objc.Send[containsMipmaps](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _containsMipmapsClass) New() containsMipmaps {
	rv := objc.Send[containsMipmaps](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ containsMipmaps) Init() containsMipmaps {
	rv := objc.Send[containsMipmaps](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ containsMipmaps) Autorelease() containsMipmaps {
	rv := objc.Send[containsMipmaps](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcontainsMipmaps creates a new containsMipmaps instance.
func NewcontainsMipmaps() containsMipmaps {
	return getcontainsMipmapsClass().New()
}




