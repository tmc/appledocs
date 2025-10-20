// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [View] class.
var (
	ViewClass     _ViewClass
	ViewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	ViewClassOnce.Do(func() {
		ViewClass = _ViewClass{objc.GetClass("MTKView")}
	})
	return ViewClass
}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	objectivec.IObject
	Draw()
	ReleaseDrawables()
}

// A specialized view that creates, configures, and displays Metal objects.
//
// The class provides a default implementation of a Metal-aware view that you can use to render graphics using Metal and display them onscreen. When asked, the view provides a object that points at a texture for you to render new contents into. Optionally, an can create depth and stencil textures for you and any intermediate textures needed for antialiasing. The view uses a to manage the Metal drawable objects. The view requires a object to manage the Metal objects it creates for you. You must set the property and, optionally, modify the view’s drawable properties before drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView
type View struct {
	objectivec.Object
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// A specialized view that creates, configures, and displays Metal objects.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return getViewClass().New()
}


// Initializes a view from data in a given unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(coder:)
func NewViewWithCoder(coder unsafe.Pointer) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Initializes a view with the specified frame rectangle and Metal device.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(frame:device:)
func NewViewWithFrameDevice(frameRect coregraphics.CGRect, device objc.ID) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithFrame:device:"), frameRect, device)
	rv.Autorelease()
	return rv
}


// Redraws the view’s contents immediately.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/draw()
func (v_ View) Draw() {
	objc.Send[objc.ID](v_.ID, objc.Sel("draw"))
}

// Releases the and objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/releaseDrawables()
func (v_ View) ReleaseDrawables() {
	objc.Send[objc.ID](v_.ID, objc.Sel("releaseDrawables"))
}

// A Boolean value that controls whether to resize the drawable as the view changes size.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/autoResizeDrawable
func (v_ View) AutoResizeDrawable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoResizeDrawable"))
	return rv
}


// SetAutoResizeDrawable sets the value of the autoResizeDrawable property.
// A Boolean value that controls whether to resize the drawable as the view changes size.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/autoResizeDrawable
func (v_ View) SetAutoResizeDrawable(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoResizeDrawable:"), value)
}
// The color to use to clear the color target when creating a render pass descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearColor
func (v_ View) ClearColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("clearColor"))
	return rv
}


// SetClearColor sets the value of the clearColor property.
// The color to use to clear the color target when creating a render pass descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearColor
func (v_ View) SetClearColor(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearColor:"), value)
}
// The depth value to use to clear the depth target when creating a render pass descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearDepth
func (v_ View) ClearDepth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("clearDepth"))
	return rv
}


// SetClearDepth sets the value of the clearDepth property.
// The depth value to use to clear the depth target when creating a render pass descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearDepth
func (v_ View) SetClearDepth(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearDepth:"), value)
}
// The stencil value to use to clear the stencil target when creating a render pass descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearStencil
func (v_ View) ClearStencil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("clearStencil"))
	return rv
}


// SetClearStencil sets the value of the clearStencil property.
// The stencil value to use to clear the stencil target when creating a render pass descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearStencil
func (v_ View) SetClearStencil(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearStencil:"), value)
}
// The color pixel format for the current drawable’s texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorPixelFormat
func (v_ View) ColorPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("colorPixelFormat"))
	return rv
}


// SetColorPixelFormat sets the value of the colorPixelFormat property.
// The color pixel format for the current drawable’s texture.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorPixelFormat
func (v_ View) SetColorPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorPixelFormat:"), value)
}
// The color space of the rendered content.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorspace
func (v_ View) Colorspace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](v_.ID, objc.Sel("colorspace"))
	return rv
}


// SetColorspace sets the value of the colorspace property.
// The color space of the rendered content.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorspace
func (v_ View) SetColorspace(value coregraphics.CGColorSpaceRef) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorspace:"), value)
}
// The drawable to use for the current frame.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentDrawable
func (v_ View) CurrentDrawable() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("currentDrawable"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentMTL4RenderPassDescriptor
func (v_ View) CurrentMTL4RenderPassDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("currentMTL4RenderPassDescriptor"))
	return rv
}

// A render pass descriptor to draw into the current drawable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentRenderPassDescriptor
func (v_ View) CurrentRenderPassDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("currentRenderPassDescriptor"))
	return rv
}

// The view’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/delegate
func (v_ View) Delegate() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The view’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/delegate
func (v_ View) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}
// The texture usage characteristics that the view uses when creating the depth and stencil textures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilAttachmentTextureUsage
func (v_ View) DepthStencilAttachmentTextureUsage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("depthStencilAttachmentTextureUsage"))
	return rv
}


// SetDepthStencilAttachmentTextureUsage sets the value of the depthStencilAttachmentTextureUsage property.
// The texture usage characteristics that the view uses when creating the depth and stencil textures.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilAttachmentTextureUsage
func (v_ View) SetDepthStencilAttachmentTextureUsage(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilAttachmentTextureUsage:"), value)
}
// The format used to generate the object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilPixelFormat
func (v_ View) DepthStencilPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("depthStencilPixelFormat"))
	return rv
}


// SetDepthStencilPixelFormat sets the value of the depthStencilPixelFormat property.
// The format used to generate the object.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilPixelFormat
func (v_ View) SetDepthStencilPixelFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilPixelFormat:"), value)
}
// The storage mode that the packed depth and stencil texture use.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilStorageMode
func (v_ View) DepthStencilStorageMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("depthStencilStorageMode"))
	return rv
}


// SetDepthStencilStorageMode sets the value of the depthStencilStorageMode property.
// The storage mode that the packed depth and stencil texture use.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilStorageMode
func (v_ View) SetDepthStencilStorageMode(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilStorageMode:"), value)
}
// A packed depth and stencil texture associated with the current drawable object’s texture.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilTexture
func (v_ View) DepthStencilTexture() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("depthStencilTexture"))
	return rv
}

// The device object the view uses to create its Metal objects.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/device
func (v_ View) Device() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("device"))
	return rv
}


// SetDevice sets the value of the device property.
// The device object the view uses to create its Metal objects.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/device
func (v_ View) SetDevice(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDevice:"), value)
}
// The current size of drawable textures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/drawableSize
func (v_ View) DrawableSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("drawableSize"))
	return rv
}


// SetDrawableSize sets the value of the drawableSize property.
// The current size of drawable textures.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/drawableSize
func (v_ View) SetDrawableSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDrawableSize:"), value)
}
// A Boolean value that indicates whether the view responds to .
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/enableSetNeedsDisplay
func (v_ View) EnableSetNeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enableSetNeedsDisplay"))
	return rv
}


// SetEnableSetNeedsDisplay sets the value of the enableSetNeedsDisplay property.
// A Boolean value that indicates whether the view responds to .

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/enableSetNeedsDisplay
func (v_ View) SetEnableSetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEnableSetNeedsDisplay:"), value)
}
// A Boolean value that determines whether the drawable’s textures are used only for rendering.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/framebufferOnly
func (v_ View) FramebufferOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("framebufferOnly"))
	return rv
}


// SetFramebufferOnly sets the value of the framebufferOnly property.
// A Boolean value that determines whether the drawable’s textures are used only for rendering.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/framebufferOnly
func (v_ View) SetFramebufferOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFramebufferOnly:"), value)
}
// A Boolean value that indicates whether the draw loop is paused.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/isPaused
func (v_ View) Paused() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("paused"))
	return rv
}


// SetPaused sets the value of the paused property.
// A Boolean value that indicates whether the draw loop is paused.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/isPaused
func (v_ View) SetPaused(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPaused:"), value)
}
// The texture usage characteristics that the view uses when creating multisample textures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorAttachmentTextureUsage
func (v_ View) MultisampleColorAttachmentTextureUsage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("multisampleColorAttachmentTextureUsage"))
	return rv
}


// SetMultisampleColorAttachmentTextureUsage sets the value of the multisampleColorAttachmentTextureUsage property.
// The texture usage characteristics that the view uses when creating multisample textures.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorAttachmentTextureUsage
func (v_ View) SetMultisampleColorAttachmentTextureUsage(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMultisampleColorAttachmentTextureUsage:"), value)
}
// The multisample color sample texture to render into.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorTexture
func (v_ View) MultisampleColorTexture() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("multisampleColorTexture"))
	return rv
}

// The device object that the system recommends using for this view.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDevice
func (v_ View) PreferredDevice() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("preferredDevice"))
	return rv
}

// The recommended dimensions of the drawable.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDrawableSize
func (v_ View) PreferredDrawableSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("preferredDrawableSize"))
	return rv
}

// The rate at which the view redraws its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredFramesPerSecond
func (v_ View) PreferredFramesPerSecond() int {
	rv := objc.Send[int](v_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}


// SetPreferredFramesPerSecond sets the value of the preferredFramesPerSecond property.
// The rate at which the view redraws its contents.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredFramesPerSecond
func (v_ View) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}
// A Boolean value that determines whether the view presents its content using a Core Animation transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/presentsWithTransaction
func (v_ View) PresentsWithTransaction() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}


// SetPresentsWithTransaction sets the value of the presentsWithTransaction property.
// A Boolean value that determines whether the view presents its content using a Core Animation transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/presentsWithTransaction
func (v_ View) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}
// The sample count used to generate the object.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
func (v_ View) SampleCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("sampleCount"))
	return rv
}


// SetSampleCount sets the value of the sampleCount property.
// The sample count used to generate the object.

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
func (v_ View) SetSampleCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSampleCount:"), value)
}

