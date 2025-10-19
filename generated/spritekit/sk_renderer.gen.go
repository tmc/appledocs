// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKRenderer] class.
var sKRendererClass = _SKRendererClass{objc.GetClass("SKRenderer")}

type _SKRendererClass struct {
	class objc.Class
}

// An interface definition for the [SKRenderer] class.
type ISKRenderer interface {
	objectivec.IObject
	RenderWithViewportCommandBufferRenderPassDescriptor(viewport unsafe.Pointer, commandBuffer unsafe.Pointer, renderPassDescriptor unsafe.Pointer)
	RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue(viewport unsafe.Pointer, renderCommandEncoder unsafe.Pointer, renderPassDescriptor unsafe.Pointer, commandQueue unsafe.Pointer)
	UpdateAtTime(currentTime TimeInterval)
}

// An object that renders a scene into a custom Metal rendering pipeline and drives the scene update cycle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer

type SKRenderer struct {
	objectivec.Object
}

// SKRendererFrom constructs a [SKRenderer] from an unsafe.Pointer.
//
// An object that renders a scene into a custom Metal rendering pipeline and drives the scene update cycle.
func SKRendererFrom(ptr unsafe.Pointer) SKRenderer {
	return SKRenderer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKRendererClass) Alloc() SKRenderer {
	rv := objc.Send[SKRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKRendererClass) New() SKRenderer {
	rv := objc.Send[SKRenderer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKRenderer) Init() SKRenderer {
	rv := objc.Send[SKRenderer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKRenderer) Autorelease() SKRenderer {
	rv := objc.Send[SKRenderer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRenderer creates a new SKRenderer instance.
func NewSKRenderer() SKRenderer {
	return sKRendererClass.New()
}


// Initializes with a specific GPU to render into. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer/init(device:)
func NewRendererWithDevice(device unsafe.Pointer) SKRenderer {
	rv := objc.Send[SKRenderer](objc.ID(sKRendererClass.class), objc.Sel("rendererWithDevice:"), device)
	rv.Autorelease()
	return rv
}


// Initializes with a specific GPU to render into. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer/init(device:)
func (sc _SKRendererClass) RendererWithDevice(device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rendererWithDevice:"), device)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer/render(withViewport:commandBuffer:renderPassDescriptor:)
func (s_ SKRenderer) RenderWithViewportCommandBufferRenderPassDescriptor(viewport unsafe.Pointer, commandBuffer unsafe.Pointer, renderPassDescriptor unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("renderWithViewport:commandBuffer:renderPassDescriptor:"), viewport, commandBuffer, renderPassDescriptor)
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer/render(withViewport:renderCommandEncoder:renderPassDescriptor:commandQueue:)
func (s_ SKRenderer) RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue(viewport unsafe.Pointer, renderCommandEncoder unsafe.Pointer, renderPassDescriptor unsafe.Pointer, commandQueue unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("renderWithViewport:renderCommandEncoder:renderPassDescriptor:commandQueue:"), viewport, renderCommandEncoder, renderPassDescriptor, commandQueue)
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRenderer/update(atTime:)
func (s_ SKRenderer) UpdateAtTime(currentTime TimeInterval) {
	objc.Send[objc.ID](s_.ID, objc.Sel("updateAtTime:"), currentTime)
}

