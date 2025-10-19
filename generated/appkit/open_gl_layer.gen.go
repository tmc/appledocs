// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OpenGLLayer] class.
var (
	openGLLayerClass     _OpenGLLayerClass
	openGLLayerClassOnce sync.Once
)

func getOpenGLLayerClass() _OpenGLLayerClass {
	openGLLayerClassOnce.Do(func() {
		openGLLayerClass = _OpenGLLayerClass{objc.GetClass("NSOpenGLLayer")}
	})
	return openGLLayerClass
}

type _OpenGLLayerClass struct {
	class objc.Class
}

// An interface definition for the [OpenGLLayer] class.
type IOpenGLLayer interface {
	objectivec.IObject
}

// A subclass of that is suitable for rendering OpenGL into layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLLayer

type OpenGLLayer struct {
	objectivec.Object
}

// OpenGLLayerFrom constructs a [OpenGLLayer] from an unsafe.Pointer.
//
// A subclass of that is suitable for rendering OpenGL into layers.
func OpenGLLayerFrom(ptr unsafe.Pointer) OpenGLLayer {
	return OpenGLLayer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (oc _OpenGLLayerClass) Alloc() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OpenGLLayerClass) New() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenGLLayer) Init() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenGLLayer) Autorelease() OpenGLLayer {
	rv := objc.Send[OpenGLLayer](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenGLLayer creates a new OpenGLLayer instance.
func NewOpenGLLayer() OpenGLLayer {
	return getOpenGLLayerClass().New()
}




