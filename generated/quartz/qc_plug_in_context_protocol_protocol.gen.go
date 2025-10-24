// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PQCPlugInContext is the QCPlugInContext protocol interface.
//
// The   protocol defines methods that you use only from within the execution method ( ) of a   object.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// See: doc://com.apple.quartz/documentation/Quartz/QCPlugInContext
type PQCPlugInContext interface {
	// Required methods
	Bounds() Rect/* debug [protocol_interface/required_method]: Bounds */
	CGLContextObj() LContextObj/* debug [protocol_interface/required_method]: CGLContextObj */
	ColorSpace() ColorSpaceRef/* debug [protocol_interface/required_method]: ColorSpace */
	CompositionURL() foundation.URL/* debug [protocol_interface/required_method]: CompositionURL */
	LogMessage(format objc.IObject /* cross-framework: NSString */)/* debug [protocol_interface/required_method]: LogMessage */
	OutputImageProviderFromBufferWithPixelFormatPixelsWidePixelsHighBaseAddressBytesPerRowReleaseCallbackReleaseContextColorSpaceShouldColorMatch(format objc.IObject /* cross-framework: NSString */, width uint, height uint, baseAddress objectivec.IObject, rowBytes uint, callback QCPlugInBufferReleaseCallback /* typedef */, context objectivec.IObject, colorSpace ColorSpaceRef /* not a class type */, colorMatch bool) objc.ID/* debug [protocol_interface/required_method]: OutputImageProviderFromBufferWithPixelFormatPixelsWidePixelsHighBaseAddressBytesPerRowReleaseCallbackReleaseContextColorSpaceShouldColorMatch */
	OutputImageProviderFromTextureWithPixelFormatPixelsWidePixelsHighNameFlippedReleaseCallbackReleaseContextColorSpaceShouldColorMatch(format objc.IObject /* cross-framework: NSString */, width uint, height uint, name objectivec.IObject, flipped bool, callback QCPlugInTextureReleaseCallback /* typedef */, context objectivec.IObject, colorSpace ColorSpaceRef /* not a class type */, colorMatch bool) objc.ID/* debug [protocol_interface/required_method]: OutputImageProviderFromTextureWithPixelFormatPixelsWidePixelsHighNameFlippedReleaseCallbackReleaseContextColorSpaceShouldColorMatch */
	UserInfo() foundation.MutableDictionary/* debug [protocol_interface/required_method]: UserInfo */
}
