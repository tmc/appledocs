// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PQCPlugInOutputImageProvider is the QCPlugInOutputImageProvider protocol interface.
//
// The   protocol eliminates the need to use explicit image types for the image output ports on a custom patch. The methods in this protocol are called by the Quartz Composer engine when the output image is needed. If your custom patch has an image output port, you need to implement the appropriate methods for rendering image data and to supply information about the rendering destination and the image bounds.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// See: doc://com.apple.quartz/documentation/Quartz/QCPlugInOutputImageProvider
type PQCPlugInOutputImageProvider interface {
	// Required methods
	ImageBounds() Rect/* debug [protocol_interface/required_method]: ImageBounds */
	ImageColorSpace() ColorSpaceRef/* debug [protocol_interface/required_method]: ImageColorSpace */
	// Optional methods
	CanRenderWithCGLContext(cgl_ctx LContextObj /* not a class type */) bool
	HasCanRenderWithCGLContext() bool
	CopyRenderedTextureForCGLContextPixelFormatBoundsIsFlipped(cgl_ctx LContextObj /* not a class type */, format objc.IObject /* cross-framework: NSString */, bounds Rect /* not a class type */, flipped objectivec.IObject) objectivec.IObject
	HasCopyRenderedTextureForCGLContextPixelFormatBoundsIsFlipped() bool
	ReleaseRenderedTextureForCGLContext(name objectivec.IObject, cgl_ctx LContextObj /* not a class type */)
	HasReleaseRenderedTextureForCGLContext() bool
	RenderToBufferWithBytesPerRowPixelFormatForBounds(baseAddress objectivec.IObject, rowBytes uint, format objc.IObject /* cross-framework: NSString */, bounds Rect /* not a class type */) bool
	HasRenderToBufferWithBytesPerRowPixelFormatForBounds() bool
	RenderWithCGLContextForBounds(cgl_ctx LContextObj /* not a class type */, bounds Rect /* not a class type */) bool
	HasRenderWithCGLContextForBounds() bool
	ShouldColorMatch() bool
	HasShouldColorMatch() bool
	SupportedBufferPixelFormats() foundation.Array
	HasSupportedBufferPixelFormats() bool
	SupportedRenderedTexturePixelFormats() foundation.Array
	HasSupportedRenderedTexturePixelFormats() bool
}
