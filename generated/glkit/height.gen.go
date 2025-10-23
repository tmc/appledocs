// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [height] class.
var (
	HeightClass     _heightClass
	HeightClassOnce sync.Once
)

func getheightClass() _heightClass {
	HeightClassOnce.Do(func() {
		HeightClass = _heightClass{objc.GetClass("height")}
	})
	return HeightClass
}

type _heightClass struct {
	class objc.Class
}

// An interface definition for the [height] class.
type Iheight interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/height-c.ivar
type height struct {
	objectivec.Object
}

// heightFrom constructs a [height] from an unsafe.Pointer.
func heightFrom(ptr unsafe.Pointer) height {
	return height{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _heightClass) Alloc() height {
	rv := objc.Send[height](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _heightClass) New() height {
	rv := objc.Send[height](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ height) Init() height {
	rv := objc.Send[height](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ height) Autorelease() height {
	rv := objc.Send[height](h_.ID, objc.Sel("autorelease"))
	return rv
}

// Newheight creates a new height instance.
func Newheight() height {
	return getheightClass().New()
}




