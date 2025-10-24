// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PQCPlugInInputImageSource is the QCPlugInInputImageSource protocol interface.
//
// The   protocol eliminates the need to use explicit image types for the image input ports on your custom patch. Not only does using the protocol avoid restrictions of a specific image type, but it avoids impedance mismatches, and provides better performance by deferring pixel computation until it is needed. When you need to access the pixels in an image, you simply convert the image to a representation (texture or buffer) using one of the methods defined by the   protocol. Use a texture representation when you want to use input images on the GPU. Use a buffer representation when you want to use input images on the CPU.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// See: doc://com.apple.quartz/documentation/Quartz/QCPlugInInputImageSource
type PQCPlugInInputImageSource interface {
	// Required methods
	BindTextureRepresentationToCGLContextTextureUnitNormalizeCoordinates(cgl_ctx LContextObj /* not a class type */, unit objectivec.IObject, flag bool)/* debug [protocol_interface/required_method]: BindTextureRepresentationToCGLContextTextureUnitNormalizeCoordinates */
	BufferBaseAddress() /* debug [protocol_interface/required_method]: BufferBaseAddress */
	BufferBytesPerRow() uint/* debug [protocol_interface/required_method]: BufferBytesPerRow */
	BufferColorSpace() ColorSpaceRef/* debug [protocol_interface/required_method]: BufferColorSpace */
	BufferPixelFormat() foundation.String/* debug [protocol_interface/required_method]: BufferPixelFormat */
	BufferPixelsHigh() uint/* debug [protocol_interface/required_method]: BufferPixelsHigh */
	BufferPixelsWide() uint/* debug [protocol_interface/required_method]: BufferPixelsWide */
	ImageBounds() Rect/* debug [protocol_interface/required_method]: ImageBounds */
	ImageColorSpace() ColorSpaceRef/* debug [protocol_interface/required_method]: ImageColorSpace */
	LockBufferRepresentationWithPixelFormatColorSpaceForBounds(format objc.IObject /* cross-framework: NSString */, colorSpace ColorSpaceRef /* not a class type */, bounds Rect /* not a class type */) bool/* debug [protocol_interface/required_method]: LockBufferRepresentationWithPixelFormatColorSpaceForBounds */
	LockTextureRepresentationWithColorSpaceForBounds(colorSpace ColorSpaceRef /* not a class type */, bounds Rect /* not a class type */) bool/* debug [protocol_interface/required_method]: LockTextureRepresentationWithColorSpaceForBounds */
	ShouldColorMatch() bool/* debug [protocol_interface/required_method]: ShouldColorMatch */
	TextureColorSpace() ColorSpaceRef/* debug [protocol_interface/required_method]: TextureColorSpace */
	TextureFlipped() bool/* debug [protocol_interface/required_method]: TextureFlipped */
	TextureMatrix() objectivec.IObject/* debug [protocol_interface/required_method]: TextureMatrix */
	TextureName() objectivec.IObject/* debug [protocol_interface/required_method]: TextureName */
	TexturePixelsHigh() uint/* debug [protocol_interface/required_method]: TexturePixelsHigh */
	TexturePixelsWide() uint/* debug [protocol_interface/required_method]: TexturePixelsWide */
	TextureTarget() objectivec.IObject/* debug [protocol_interface/required_method]: TextureTarget */
	UnbindTextureRepresentationFromCGLContextTextureUnit(cgl_ctx LContextObj /* not a class type */, unit objectivec.IObject)/* debug [protocol_interface/required_method]: UnbindTextureRepresentationFromCGLContextTextureUnit */
	UnlockBufferRepresentation()/* debug [protocol_interface/required_method]: UnlockBufferRepresentation */
	UnlockTextureRepresentation()/* debug [protocol_interface/required_method]: UnlockTextureRepresentation */
}
