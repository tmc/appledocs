// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SVGFDefaultTextureAllocator] class.
var (
	SVGFDefaultTextureAllocatorClass     _SVGFDefaultTextureAllocatorClass
	SVGFDefaultTextureAllocatorClassOnce sync.Once
)

func getSVGFDefaultTextureAllocatorClass() _SVGFDefaultTextureAllocatorClass {
	SVGFDefaultTextureAllocatorClassOnce.Do(func() {
		SVGFDefaultTextureAllocatorClass = _SVGFDefaultTextureAllocatorClass{objc.GetClass("MPSSVGFDefaultTextureAllocator")}
	})
	return SVGFDefaultTextureAllocatorClass
}

type _SVGFDefaultTextureAllocatorClass struct {
	class objc.Class
}

// An interface definition for the [SVGFDefaultTextureAllocator] class.
type ISVGFDefaultTextureAllocator interface {
	objectivec.IObject
	ReturnTexture(texture objc.ID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator
type SVGFDefaultTextureAllocator struct {
	objectivec.Object
}

// SVGFDefaultTextureAllocatorFrom constructs a [SVGFDefaultTextureAllocator] from an unsafe.Pointer.
func SVGFDefaultTextureAllocatorFrom(ptr unsafe.Pointer) SVGFDefaultTextureAllocator {
	return SVGFDefaultTextureAllocator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SVGFDefaultTextureAllocatorClass) Alloc() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SVGFDefaultTextureAllocatorClass) New() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SVGFDefaultTextureAllocator) Init() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SVGFDefaultTextureAllocator) Autorelease() SVGFDefaultTextureAllocator {
	rv := objc.Send[SVGFDefaultTextureAllocator](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSVGFDefaultTextureAllocator creates a new SVGFDefaultTextureAllocator instance.
func NewSVGFDefaultTextureAllocator() SVGFDefaultTextureAllocator {
	return getSVGFDefaultTextureAllocatorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSSVGFDefaultTextureAllocator/return(_:)
func (s_ SVGFDefaultTextureAllocator) ReturnTexture(texture objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("returnTexture:"), texture)
}



