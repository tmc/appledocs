// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Renderer] class.
var rendererClass = _RendererClass{objc.GetClass("CARenderer")}

type _RendererClass struct {
	class objc.Class
}

// A layer that allows an application to render a layer tree into a Core OpenGL context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer

type Renderer struct {
	objectivec.Object
}

// RendererFrom constructs a [Renderer] from an unsafe.Pointer.
//
// A layer that allows an application to render a layer tree into a Core OpenGL context.
func RendererFrom(ptr unsafe.Pointer) Renderer {
	return Renderer{objectivec.Object{objc.ID(ptr)}}
}

// Render the update region of the current frame to the target context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CARenderer/render()
func (r_ Renderer) Render() {
	objc.Send[objc.ID](r_.ID, objc.Sel("render"))
}


