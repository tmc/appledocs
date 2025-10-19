// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EAGLLayer] class.
var (
	eAGLLayerClass     _EAGLLayerClass
	eAGLLayerClassOnce sync.Once
)

func getEAGLLayerClass() _EAGLLayerClass {
	eAGLLayerClassOnce.Do(func() {
		eAGLLayerClass = _EAGLLayerClass{objc.GetClass("CAEAGLLayer")}
	})
	return eAGLLayerClass
}

type _EAGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [EAGLLayer] class.
type IEAGLLayer interface {
	ILayer
}

// A layer that supports drawing OpenGL content in iOS and tvOS applications.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEAGLLayer
type EAGLLayer struct {
	Layer
}

// EAGLLayerFrom constructs a [EAGLLayer] from an unsafe.Pointer.
//
// A layer that supports drawing OpenGL content in iOS and tvOS applications.
func EAGLLayerFrom(ptr unsafe.Pointer) EAGLLayer {
	return EAGLLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EAGLLayerClass) Alloc() EAGLLayer {
	rv := objc.Send[EAGLLayer](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAGLLayerClass) New() EAGLLayer {
	rv := objc.Send[EAGLLayer](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAGLLayer) Init() EAGLLayer {
	rv := objc.Send[EAGLLayer](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAGLLayer) Autorelease() EAGLLayer {
	rv := objc.Send[EAGLLayer](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAGLLayer creates a new EAGLLayer instance.
func NewEAGLLayer() EAGLLayer {
	return getEAGLLayerClass().New()
}




