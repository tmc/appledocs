// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [depth] class.
var (
	DepthClass     _depthClass
	DepthClassOnce sync.Once
)

func getdepthClass() _depthClass {
	DepthClassOnce.Do(func() {
		DepthClass = _depthClass{objc.GetClass("depth")}
	})
	return DepthClass
}

type _depthClass struct {
	class objc.Class
}

// An interface definition for the [depth] class.
type Idepth interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/depth-c.ivar
type depth struct {
	objectivec.Object
}

// depthFrom constructs a [depth] from an unsafe.Pointer.
func depthFrom(ptr unsafe.Pointer) depth {
	return depth{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _depthClass) Alloc() depth {
	rv := objc.Send[depth](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _depthClass) New() depth {
	rv := objc.Send[depth](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ depth) Init() depth {
	rv := objc.Send[depth](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ depth) Autorelease() depth {
	rv := objc.Send[depth](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdepth creates a new depth instance.
func Newdepth() depth {
	return getdepthClass().New()
}




