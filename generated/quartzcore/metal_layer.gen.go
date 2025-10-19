// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MetalLayer] class.
var (
	metalLayerClass     _MetalLayerClass
	metalLayerClassOnce sync.Once
)

func getMetalLayerClass() _MetalLayerClass {
	metalLayerClassOnce.Do(func() {
		metalLayerClass = _MetalLayerClass{objc.GetClass("CAMetalLayer")}
	})
	return metalLayerClass
}

type _MetalLayerClass struct {
	class objc.Class
}

// An interface definition for the [MetalLayer] class.
type IMetalLayer interface {
	ILayer
}

// A Core Animation layer that Metal can render into, typically displayed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
type MetalLayer struct {
	Layer
}

// MetalLayerFrom constructs a [MetalLayer] from an unsafe.Pointer.
//
// A Core Animation layer that Metal can render into, typically displayed onscreen.
func MetalLayerFrom(ptr unsafe.Pointer) MetalLayer {
	return MetalLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MetalLayerClass) Alloc() MetalLayer {
	rv := objc.Send[MetalLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetalLayerClass) New() MetalLayer {
	rv := objc.Send[MetalLayer](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetalLayer) Init() MetalLayer {
	rv := objc.Send[MetalLayer](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetalLayer) Autorelease() MetalLayer {
	rv := objc.Send[MetalLayer](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalLayer creates a new MetalLayer instance.
func NewMetalLayer() MetalLayer {
	return getMetalLayerClass().New()
}




