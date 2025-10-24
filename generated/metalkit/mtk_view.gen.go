// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/metal"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTKView */


/* debug [class_header]: Header for MTKView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for View */
// An interface definition for the [View] class.
type IView interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for View */
	// properties:
	AutoResizeDrawable() bool
	SetAutoResizeDrawable(value bool)
	ClearColor() objc.IObject /* cross-framework: ClearColor */
	SetClearColor(value objc.IObject /* cross-framework: ClearColor */)
	ClearDepth() float64
	SetClearDepth(value float64)
	ClearStencil() uint32 /* not a class type */
	SetClearStencil(value uint32 /* not a class type */)
	ColorPixelFormat() PixelFormat /* not a class type */
	SetColorPixelFormat(value PixelFormat /* not a class type */)
	Colorspace() ColorSpaceRef /* not a class type */
	SetColorspace(value ColorSpaceRef /* not a class type */)
	CurrentDrawable() unsafe.Pointer
	CurrentMTL4RenderPassDescriptor() metal.MTL4RenderPassDescriptor
	CurrentRenderPassDescriptor() metal.RenderPassDescriptor
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DepthStencilAttachmentTextureUsage() TextureUsage /* not a class type */
	SetDepthStencilAttachmentTextureUsage(value TextureUsage /* not a class type */)
	DepthStencilPixelFormat() PixelFormat /* not a class type */
	SetDepthStencilPixelFormat(value PixelFormat /* not a class type */)
	DepthStencilStorageMode() StorageMode /* not a class type */
	SetDepthStencilStorageMode(value StorageMode /* not a class type */)
	DepthStencilTexture() unsafe.Pointer
	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	DrawableSize() corefoundation.CGSize
	SetDrawableSize(value corefoundation.CGSize)
	EnableSetNeedsDisplay() bool
	SetEnableSetNeedsDisplay(value bool)
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	Paused() bool
	SetPaused(value bool)
	MultisampleColorAttachmentTextureUsage() TextureUsage /* not a class type */
	SetMultisampleColorAttachmentTextureUsage(value TextureUsage /* not a class type */)
	MultisampleColorTexture() unsafe.Pointer
	PreferredDevice() unsafe.Pointer
	PreferredDrawableSize() corefoundation.CGSize
	PreferredFramesPerSecond() int
	SetPreferredFramesPerSecond(value int)
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)
	SampleCount() uint
	SetSampleCount(value uint)
	IsPaused() bool
	SetIsPaused(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for View */
	// methods:
	Draw()
	ReleaseDrawables()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for View */
// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for View */
// A specialized view that creates, configures, and displays Metal objects.
//
// The class provides a default implementation of a Metal-aware view that you can use to render graphics using Metal and display them onscreen. When asked, the view provides a object that points at a texture for you to render new contents into. Optionally, an can create depth and stencil textures for you and any intermediate textures needed for antialiasing. The view uses a to manage the Metal drawable objects. The view requires a object to manage the Metal objects it creates for you. You must set the property and, optionally, modify the view’s drawable properties before drawing.


// A specialized view that creates, configures, and displays Metal objects.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for View */

// Initializes a view from data in a given unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(coder:)
func NewViewWithCoder(coder foundation.Coder) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewWithCoder */


// Initializes a view with the specified frame rectangle and Metal device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(frame:device:)
func NewViewWithFrameDevice(frameRect corefoundation.CGRect, device unsafe.Pointer) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithFrame:device:"), frameRect, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewWithFrameDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for View */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for View */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for View */

// Redraws the view’s contents immediately.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/draw()
func (v_ View) Draw() {
	objc.Send[objc.ID](v_.ID, objc.Sel("draw"))
}/* debug [instance_methods/method]: Draw */


// Releases the and objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/releaseDrawables()
func (v_ View) ReleaseDrawables() {
	objc.Send[objc.ID](v_.ID, objc.Sel("releaseDrawables"))
}/* debug [instance_methods/method]: ReleaseDrawables */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for View */

// A Boolean value that controls whether to resize the drawable as the view changes size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/autoResizeDrawable
func (v_ View) AutoResizeDrawable() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoResizeDrawable"))
	return rv
}/* debug [instance_properties/getter]: autoResizeDrawable */


// A Boolean value that controls whether to resize the drawable as the view changes size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/autoResizeDrawable
func (v_ View) SetAutoResizeDrawable(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoResizeDrawable:"), value)
}/* debug [instance_properties/setter]: autoResizeDrawable */


// The color to use to clear the color target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearColor
func (v_ View) ClearColor() objc.IObject /* cross-framework: ClearColor */ {
	rv := objc.Send[metal.ClearColor](v_.ID, objc.Sel("clearColor"))
	return rv
}/* debug [instance_properties/getter]: clearColor */


// The color to use to clear the color target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearColor
func (v_ View) SetClearColor(value objc.IObject /* cross-framework: ClearColor */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearColor:"), value)
}/* debug [instance_properties/setter]: clearColor */


// The depth value to use to clear the depth target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearDepth
func (v_ View) ClearDepth() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("clearDepth"))
	return rv
}/* debug [instance_properties/getter]: clearDepth */


// The depth value to use to clear the depth target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearDepth
func (v_ View) SetClearDepth(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearDepth:"), value)
}/* debug [instance_properties/setter]: clearDepth */


// The stencil value to use to clear the stencil target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearStencil
func (v_ View) ClearStencil() uint32 /* not a class type */ {
	rv := objc.Send[uint32](v_.ID, objc.Sel("clearStencil"))
	return rv
}/* debug [instance_properties/getter]: clearStencil */


// The stencil value to use to clear the stencil target when creating a render pass descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/clearStencil
func (v_ View) SetClearStencil(value uint32 /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClearStencil:"), value)
}/* debug [instance_properties/setter]: clearStencil */


// The color pixel format for the current drawable’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorPixelFormat
func (v_ View) ColorPixelFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](v_.ID, objc.Sel("colorPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: colorPixelFormat */


// The color pixel format for the current drawable’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorPixelFormat
func (v_ View) SetColorPixelFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorPixelFormat:"), value)
}/* debug [instance_properties/setter]: colorPixelFormat */


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorspace
func (v_ View) Colorspace() ColorSpaceRef /* not a class type */ {
	rv := objc.Send[ColorSpaceRef](v_.ID, objc.Sel("colorspace"))
	return rv
}/* debug [instance_properties/getter]: colorspace */


// The color space of the rendered content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/colorspace
func (v_ View) SetColorspace(value ColorSpaceRef /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setColorspace:"), value)
}/* debug [instance_properties/setter]: colorspace */


// The drawable to use for the current frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentDrawable
func (v_ View) CurrentDrawable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("currentDrawable"))
	return rv
}/* debug [instance_properties/getter]: currentDrawable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentMTL4RenderPassDescriptor
func (v_ View) CurrentMTL4RenderPassDescriptor() metal.MTL4RenderPassDescriptor {
	rv := objc.Send[metal.MTL4RenderPassDescriptor](v_.ID, objc.Sel("currentMTL4RenderPassDescriptor"))
	return rv
}/* debug [instance_properties/getter]: currentMTL4RenderPassDescriptor */


// A render pass descriptor to draw into the current drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/currentRenderPassDescriptor
func (v_ View) CurrentRenderPassDescriptor() metal.RenderPassDescriptor {
	rv := objc.Send[metal.RenderPassDescriptor](v_.ID, objc.Sel("currentRenderPassDescriptor"))
	return rv
}/* debug [instance_properties/getter]: currentRenderPassDescriptor */


// The view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/delegate
func (v_ View) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/delegate
func (v_ View) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The texture usage characteristics that the view uses when creating the depth and stencil textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilAttachmentTextureUsage
func (v_ View) DepthStencilAttachmentTextureUsage() TextureUsage /* not a class type */ {
	rv := objc.Send[TextureUsage](v_.ID, objc.Sel("depthStencilAttachmentTextureUsage"))
	return rv
}/* debug [instance_properties/getter]: depthStencilAttachmentTextureUsage */


// The texture usage characteristics that the view uses when creating the depth and stencil textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilAttachmentTextureUsage
func (v_ View) SetDepthStencilAttachmentTextureUsage(value TextureUsage /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilAttachmentTextureUsage:"), value)
}/* debug [instance_properties/setter]: depthStencilAttachmentTextureUsage */


// The format used to generate the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilPixelFormat
func (v_ View) DepthStencilPixelFormat() PixelFormat /* not a class type */ {
	rv := objc.Send[PixelFormat](v_.ID, objc.Sel("depthStencilPixelFormat"))
	return rv
}/* debug [instance_properties/getter]: depthStencilPixelFormat */


// The format used to generate the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilPixelFormat
func (v_ View) SetDepthStencilPixelFormat(value PixelFormat /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilPixelFormat:"), value)
}/* debug [instance_properties/setter]: depthStencilPixelFormat */


// The storage mode that the packed depth and stencil texture use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilStorageMode
func (v_ View) DepthStencilStorageMode() StorageMode /* not a class type */ {
	rv := objc.Send[StorageMode](v_.ID, objc.Sel("depthStencilStorageMode"))
	return rv
}/* debug [instance_properties/getter]: depthStencilStorageMode */


// The storage mode that the packed depth and stencil texture use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilStorageMode
func (v_ View) SetDepthStencilStorageMode(value StorageMode /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDepthStencilStorageMode:"), value)
}/* debug [instance_properties/setter]: depthStencilStorageMode */


// A packed depth and stencil texture associated with the current drawable object’s texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilTexture
func (v_ View) DepthStencilTexture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("depthStencilTexture"))
	return rv
}/* debug [instance_properties/getter]: depthStencilTexture */


// The device object the view uses to create its Metal objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/device
func (v_ View) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device object the view uses to create its Metal objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/device
func (v_ View) SetDevice(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// The current size of drawable textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/drawableSize
func (v_ View) DrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("drawableSize"))
	return rv
}/* debug [instance_properties/getter]: drawableSize */


// The current size of drawable textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/drawableSize
func (v_ View) SetDrawableSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDrawableSize:"), value)
}/* debug [instance_properties/setter]: drawableSize */


// A Boolean value that indicates whether the view responds to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/enableSetNeedsDisplay
func (v_ View) EnableSetNeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enableSetNeedsDisplay"))
	return rv
}/* debug [instance_properties/getter]: enableSetNeedsDisplay */


// A Boolean value that indicates whether the view responds to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/enableSetNeedsDisplay
func (v_ View) SetEnableSetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setEnableSetNeedsDisplay:"), value)
}/* debug [instance_properties/setter]: enableSetNeedsDisplay */


// A Boolean value that determines whether the drawable’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/framebufferOnly
func (v_ View) FramebufferOnly() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("framebufferOnly"))
	return rv
}/* debug [instance_properties/getter]: framebufferOnly */


// A Boolean value that determines whether the drawable’s textures are used only for rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/framebufferOnly
func (v_ View) SetFramebufferOnly(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFramebufferOnly:"), value)
}/* debug [instance_properties/setter]: framebufferOnly */


// A Boolean value that indicates whether the draw loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/isPaused
func (v_ View) Paused() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("paused"))
	return rv
}/* debug [instance_properties/getter]: paused */


// A Boolean value that indicates whether the draw loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/isPaused
func (v_ View) SetPaused(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPaused:"), value)
}/* debug [instance_properties/setter]: paused */


// The texture usage characteristics that the view uses when creating multisample textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorAttachmentTextureUsage
func (v_ View) MultisampleColorAttachmentTextureUsage() TextureUsage /* not a class type */ {
	rv := objc.Send[TextureUsage](v_.ID, objc.Sel("multisampleColorAttachmentTextureUsage"))
	return rv
}/* debug [instance_properties/getter]: multisampleColorAttachmentTextureUsage */


// The texture usage characteristics that the view uses when creating multisample textures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorAttachmentTextureUsage
func (v_ View) SetMultisampleColorAttachmentTextureUsage(value TextureUsage /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setMultisampleColorAttachmentTextureUsage:"), value)
}/* debug [instance_properties/setter]: multisampleColorAttachmentTextureUsage */


// The multisample color sample texture to render into.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorTexture
func (v_ View) MultisampleColorTexture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("multisampleColorTexture"))
	return rv
}/* debug [instance_properties/getter]: multisampleColorTexture */


// The device object that the system recommends using for this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDevice
func (v_ View) PreferredDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("preferredDevice"))
	return rv
}/* debug [instance_properties/getter]: preferredDevice */


// The recommended dimensions of the drawable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDrawableSize
func (v_ View) PreferredDrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v_.ID, objc.Sel("preferredDrawableSize"))
	return rv
}/* debug [instance_properties/getter]: preferredDrawableSize */


// The rate at which the view redraws its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredFramesPerSecond
func (v_ View) PreferredFramesPerSecond() int {
	rv := objc.Send[int](v_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}/* debug [instance_properties/getter]: preferredFramesPerSecond */


// The rate at which the view redraws its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/preferredFramesPerSecond
func (v_ View) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}/* debug [instance_properties/setter]: preferredFramesPerSecond */


// A Boolean value that determines whether the view presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/presentsWithTransaction
func (v_ View) PresentsWithTransaction() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("presentsWithTransaction"))
	return rv
}/* debug [instance_properties/getter]: presentsWithTransaction */


// A Boolean value that determines whether the view presents its content using a Core Animation transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/presentsWithTransaction
func (v_ View) SetPresentsWithTransaction(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPresentsWithTransaction:"), value)
}/* debug [instance_properties/setter]: presentsWithTransaction */


// The sample count used to generate the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
func (v_ View) SampleCount() uint {
	rv := objc.Send[uint](v_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The sample count used to generate the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
func (v_ View) SetSampleCount(value uint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setSampleCount:"), value)
}/* debug [instance_properties/setter]: sampleCount */


// A Boolean value that indicates whether the draw loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalkit/mtkview/ispaused
func (v_ View) IsPaused() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isPaused"))
	return rv
}/* debug [instance_properties/getter]: isPaused */


// A Boolean value that indicates whether the draw loop is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalkit/mtkview/ispaused
func (v_ View) SetIsPaused(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIsPaused:"), value)
}/* debug [instance_properties/setter]: isPaused */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTKView */


