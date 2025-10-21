// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

// Enum types and constants
// CGBitmapInfo - Component information for a bitmap image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo
type BitmapInfo uint

const (
// kCGBitmapByteOrder16Big - 16-bit, big endian format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder16Big
kCGBitmapByteOrder16Big BitmapInfo = 0
// kCGBitmapByteOrder16Little - 16-bit, little endian format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder16Little
kCGBitmapByteOrder16Little BitmapInfo = 0
// kCGBitmapByteOrder32Big - 32-bit, big endian format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder32Big
kCGBitmapByteOrder32Big BitmapInfo = 0
// kCGBitmapByteOrder32Little - 32-bit, little endian format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder32Little
kCGBitmapByteOrder32Little BitmapInfo = 0
// kCGBitmapByteOrderDefault - The default byte order.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrderDefault
kCGBitmapByteOrderDefault BitmapInfo = 0
// kCGBitmapFloatComponents - The components of a bitmap are floating-point values.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/floatComponents
kCGBitmapFloatComponents BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapAlphaInfoMask
kCGBitmapAlphaInfoMask BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapByteOrderInfoMask
kCGBitmapByteOrderInfoMask BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapByteOrderMask
kCGBitmapByteOrderMask BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapComponentInfoMask
kCGBitmapComponentInfoMask BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapFloatInfoMask
kCGBitmapFloatInfoMask BitmapInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapPixelFormatInfoMask
kCGBitmapPixelFormatInfoMask BitmapInfo = 0
)

// CGBitmapLayout enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout
type BitmapLayout uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/abgr
kCGBitmapLayoutABGR BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/alphaOnly
kCGBitmapLayoutAlphaOnly BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/argb
kCGBitmapLayoutARGB BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/bgra
kCGBitmapLayoutBGRA BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/bgrx
kCGBitmapLayoutBGRX BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/cmyk
kCGBitmapLayoutCMYK BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/gray
kCGBitmapLayoutGray BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/grayAlpha
kCGBitmapLayoutGrayAlpha BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/rgba
kCGBitmapLayoutRGBA BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/rgbx
kCGBitmapLayoutRGBX BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/xbgr
kCGBitmapLayoutXBGR BitmapLayout = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/xrgb
kCGBitmapLayoutXRGB BitmapLayout = 0
)

// CGBlendMode - Compositing operations for images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode
type BlendMode uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/clear
kCGBlendModeClear BlendMode = 0
// kCGBlendModeColor - Uses the luminance values of the background with the hue and saturation values of the source image. This mode preserves the gray levels in the image. You can use this mode to color monochrome images or to tint color images.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/color
kCGBlendModeColor BlendMode = 0
// kCGBlendModeColorDodge - Brightens the background image samples to reflect the source image samples. Source image sample values that specify black do not produce a change.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/colorDodge
kCGBlendModeColorDodge BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/copy
kCGBlendModeCopy BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/darken
kCGBlendModeDarken BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationAtop
kCGBlendModeDestinationAtop BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationIn
kCGBlendModeDestinationIn BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationOut
kCGBlendModeDestinationOut BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationOver
kCGBlendModeDestinationOver BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/difference
kCGBlendModeDifference BlendMode = 0
// kCGBlendModeExclusion - Produces an effect similar to that produced by  , but with lower contrast. Source image sample values that are black don’t produce a change; white inverts the background color values.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/exclusion
kCGBlendModeExclusion BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/hardLight
kCGBlendModeHardLight BlendMode = 0
// kCGBlendModeHue - Uses the luminance and saturation values of the background with the hue of the source image.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/hue
kCGBlendModeHue BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/lighten
kCGBlendModeLighten BlendMode = 0
// kCGBlendModeLuminosity - Uses the hue and saturation of the background with the luminance of the source image. This mode creates an effect that is inverse to the effect created by  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/luminosity
kCGBlendModeLuminosity BlendMode = 0
// kCGBlendModeMultiply - Multiplies the source image samples with the background image samples. This results in colors that are at least as dark as either of the two contributing sample colors.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/multiply
kCGBlendModeMultiply BlendMode = 0
// kCGBlendModeNormal - Paints the source image samples over the background image samples.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/normal
kCGBlendModeNormal BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/overlay
kCGBlendModeOverlay BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/plusDarker
kCGBlendModePlusDarker BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/plusLighter
kCGBlendModePlusLighter BlendMode = 0
// kCGBlendModeSaturation - Uses the luminance and hue values of the background with the saturation of the source image. Areas of the background that have no saturation (that is, pure gray areas) don’t produce a change.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/saturation
kCGBlendModeSaturation BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/softLight
kCGBlendModeSoftLight BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceAtop
kCGBlendModeSourceAtop BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceIn
kCGBlendModeSourceIn BlendMode = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceOut
kCGBlendModeSourceOut BlendMode = 0
// kCGBlendModeXOR - . This XOR mode is only nominally related to the classical bitmap XOR operation, which is not supported by Core Graphics
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/xor
kCGBlendModeXOR BlendMode = 0
)

// CGCaptureOptions - Configuration parameters that are used when capturing displays.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions
type CaptureOptions uint

const (
// kCGCaptureNoOptions - The system should use the default fill behavior, which is fill with black.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions/kCGCaptureNoOptions
kCGCaptureNoOptions CaptureOptions = 0
// kCGCaptureNoFill - Disables fill with black.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions/noFill
kCGCaptureNoFill CaptureOptions = 0
)

// CGColorConversionInfoTransformType - Constants describing how a color conversion uses color spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType
type ColorConversionInfoTransformType uint

const (
// kCGColorConversionTransformApplySpace - Specifies a color conversion between one color profile and another.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformApplySpace
kCGColorConversionTransformApplySpace ColorConversionInfoTransformType = 0
// kCGColorConversionTransformFromSpace - Specifies a color conversion from a device color space to a color profile.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformFromSpace
kCGColorConversionTransformFromSpace ColorConversionInfoTransformType = 0
// kCGColorConversionTransformToSpace - Specifies a color conversion from a color profile to a device color space.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformToSpace
kCGColorConversionTransformToSpace ColorConversionInfoTransformType = 0
)

// CGColorModel enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel
type ColorModel uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/cmyk
kCGColorModelCMYK ColorModel = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/deviceN
kCGColorModelDeviceN ColorModel = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/gray
kCGColorModelGray ColorModel = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/kCGColorModelNoColorant
kCGColorModelNoColorant ColorModel = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/lab
kCGColorModelLab ColorModel = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/rgb
kCGColorModelRGB ColorModel = 0
)

// CGColorRenderingIntent - Handling options for colors that are not located within the destination color space of a graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent
type ColorRenderingIntent uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/absoluteColorimetric
kCGRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 0
// kCGRenderingIntentDefault - The default rendering intent for the graphics context.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/defaultIntent
kCGRenderingIntentDefault ColorRenderingIntent = 0
// kCGRenderingIntentPerceptual - Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device. Perceptual intent is good for photographs and other complex, detailed images.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/perceptual
kCGRenderingIntentPerceptual ColorRenderingIntent = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/relativeColorimetric
kCGRenderingIntentRelativeColorimetric ColorRenderingIntent = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/saturation
kCGRenderingIntentSaturation ColorRenderingIntent = 0
)

// CGColorSpaceModel - Models for color spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel
type ColorSpaceModel uint

const (
// kCGColorSpaceModelXYZ - An XYZ color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/XYZ
kCGColorSpaceModelXYZ ColorSpaceModel = 0
// kCGColorSpaceModelCMYK - A CMYK color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/cmyk
kCGColorSpaceModelCMYK ColorSpaceModel = 0
// kCGColorSpaceModelDeviceN - A DeviceN color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/deviceN
kCGColorSpaceModelDeviceN ColorSpaceModel = 0
// kCGColorSpaceModelIndexed - An indexed color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/indexed
kCGColorSpaceModelIndexed ColorSpaceModel = 0
// kCGColorSpaceModelLab - A Lab color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/lab
kCGColorSpaceModelLab ColorSpaceModel = 0
// kCGColorSpaceModelMonochrome - A monochrome color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
kCGColorSpaceModelMonochrome ColorSpaceModel = 0
// kCGColorSpaceModelPattern - A pattern color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/pattern
kCGColorSpaceModelPattern ColorSpaceModel = 0
// kCGColorSpaceModelRGB - An RGB color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
kCGColorSpaceModelRGB ColorSpaceModel = 0
// kCGColorSpaceModelUnknown - An unknown color space model.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/unknown
kCGColorSpaceModelUnknown ColorSpaceModel = 0
)

// CGComponent enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent
type Component uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/float16Bit
kCGComponentFloat16Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/float32Bit
kCGComponentFloat32Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer10Bit
kCGComponentInteger10Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer16Bit
kCGComponentInteger16Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer32Bit
kCGComponentInteger32Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer8Bit
kCGComponentInteger8Bit Component = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/unknown
kCGComponentUnknown Component = 0
)

// CGConfigureOption - The scope of the changes in a display configuration transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption
type ConfigureOption uint

const (
// kCGConfigureForAppOnly - Changes persist for the lifetime of the current application. After the application terminates, the display configuration settings revert to the current login session.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/forAppOnly
kCGConfigureForAppOnly ConfigureOption = 0
// kCGConfigureForSession - Changes persist for the lifetime of the current login session. After the current session terminates, the displays revert to the last saved permanent configuration.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/forSession
kCGConfigureForSession ConfigureOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/permanently
kCGConfigurePermanently ConfigureOption = 0
)

// CGDisplayChangeSummaryFlags - The configuration parameters that are passed to a display reconfiguration callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags
type DisplayChangeSummaryFlags uint

const (
// kCGDisplayAddFlag - The display has been added to the active display list.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/addFlag
kCGDisplayAddFlag DisplayChangeSummaryFlags = 0
// kCGDisplayBeginConfigurationFlag - The display configuration is about to change.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/beginConfigurationFlag
kCGDisplayBeginConfigurationFlag DisplayChangeSummaryFlags = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/desktopShapeChangedFlag
kCGDisplayDesktopShapeChangedFlag DisplayChangeSummaryFlags = 0
// kCGDisplayDisabledFlag - The display has been disabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/disabledFlag
kCGDisplayDisabledFlag DisplayChangeSummaryFlags = 0
// kCGDisplayEnabledFlag - The display has been enabled.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/enabledFlag
kCGDisplayEnabledFlag DisplayChangeSummaryFlags = 0
// kCGDisplayMirrorFlag - The display is now mirroring another display.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/mirrorFlag
kCGDisplayMirrorFlag DisplayChangeSummaryFlags = 0
// kCGDisplayMovedFlag - The location of the upper-left corner of the display in the global display coordinate space has changed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/movedFlag
kCGDisplayMovedFlag DisplayChangeSummaryFlags = 0
// kCGDisplayRemoveFlag - The display has been removed from the active display list.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/removeFlag
kCGDisplayRemoveFlag DisplayChangeSummaryFlags = 0
// kCGDisplaySetMainFlag - The display is now the main display.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/setMainFlag
kCGDisplaySetMainFlag DisplayChangeSummaryFlags = 0
// kCGDisplaySetModeFlag - The display mode has changed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/setModeFlag
kCGDisplaySetModeFlag DisplayChangeSummaryFlags = 0
// kCGDisplayUnMirrorFlag - The display is no longer mirroring another display.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/unMirrorFlag
kCGDisplayUnMirrorFlag DisplayChangeSummaryFlags = 0
)

// CGDisplayStreamFrameStatus - Describes a frame update event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus
type DisplayStreamFrameStatus uint

const (
// kCGDisplayStreamFrameStatusFrameBlank - A new frame was not generated because the display has gone blank.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameBlank
kCGDisplayStreamFrameStatusFrameBlank DisplayStreamFrameStatus = 0
// kCGDisplayStreamFrameStatusFrameComplete - A new frame was generated.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameComplete
kCGDisplayStreamFrameStatusFrameComplete DisplayStreamFrameStatus = 0
// kCGDisplayStreamFrameStatusFrameIdle - A new frame was not generated because the display did not change.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameIdle
kCGDisplayStreamFrameStatusFrameIdle DisplayStreamFrameStatus = 0
// kCGDisplayStreamFrameStatusStopped - The display stream was stopped.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/stopped
kCGDisplayStreamFrameStatusStopped DisplayStreamFrameStatus = 0
)

// CGDisplayStreamUpdateRectType - Use these constants to determine which rectangles your app is interested in.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType
type DisplayStreamUpdateRectType uint

const (
// kCGDisplayStreamUpdateDirtyRects - The union of both rectangles that were redrawn and rectangles that were moved.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/dirtyRects
kCGDisplayStreamUpdateDirtyRects DisplayStreamUpdateRectType = 0
// kCGDisplayStreamUpdateMovedRects - The rectangles for the portions of the display that were simply moved from one part of the display to another.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/movedRects
kCGDisplayStreamUpdateMovedRects DisplayStreamUpdateRectType = 0
// kCGDisplayStreamUpdateReducedDirtyRects - The union is calculated and then simplified. This reduces the number of rectangles returned to your app, but it may report some pixels that were not actually changed.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/reducedDirtyRects
kCGDisplayStreamUpdateReducedDirtyRects DisplayStreamUpdateRectType = 0
// kCGDisplayStreamUpdateRefreshedRects - The rectangles for the portions of the display that were redrawn.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/refreshedRects
kCGDisplayStreamUpdateRefreshedRects DisplayStreamUpdateRectType = 0
)

// CGError - A uniform type for result codes returned by functions in Core Graphics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError
type Error uint

const (
// kCGErrorCannotComplete - The requested operation is inappropriate for the parameters passed in, or the current system state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/cannotComplete
kCGErrorCannotComplete Error = 0
// kCGErrorFailure - A general failure occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/failure
kCGErrorFailure Error = 0
// kCGErrorIllegalArgument - One or more of the parameters passed to a function are invalid. Check for   pointers.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/illegalArgument
kCGErrorIllegalArgument Error = 0
// kCGErrorInvalidConnection - The parameter representing a connection to the window server is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidConnection
kCGErrorInvalidConnection Error = 0
// kCGErrorInvalidContext - The   or context identifier parameter is not valid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidContext
kCGErrorInvalidContext Error = 0
// kCGErrorInvalidOperation - The requested operation is not valid for the parameters passed in, or the current system state.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidOperation
kCGErrorInvalidOperation Error = 0
// kCGErrorNoneAvailable - The requested operation could not be completed as the indicated resources were not found.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/noneAvailable
kCGErrorNoneAvailable Error = 0
// kCGErrorNotImplemented - Return value from obsolete function stubs present for binary compatibility, but not typically called.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/notImplemented
kCGErrorNotImplemented Error = 0
// kCGErrorRangeCheck - A parameter passed in has a value that is inappropriate, or which does not map to a useful operation or value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/rangeCheck
kCGErrorRangeCheck Error = 0
// kCGErrorSuccess - The requested operation was completed successfully.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/success
kCGErrorSuccess Error = 0
// kCGErrorTypeCheck - A data type or token was encountered that did not match the expected type or token.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/typeCheck
kCGErrorTypeCheck Error = 0
)

// CGEventField - Constants used as keys to access specialized fields in low-level events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField
type EventField uint

const (
// kCGEventSourceGroupID - Key to access a field that contains the event source Unix effective GID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceGroupID
kCGEventSourceGroupID EventField = 0
// kCGEventSourceStateID - Key to access a field that contains the event source state ID used to create this event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceStateID
kCGEventSourceStateID EventField = 0
// kCGEventSourceUnixProcessID - Key to access a field that contains the event source Unix process ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUnixProcessID
kCGEventSourceUnixProcessID EventField = 0
// kCGEventSourceUserData - Key to access a field that contains the event source user-supplied data, up to 64 bits.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUserData
kCGEventSourceUserData EventField = 0
// kCGEventSourceUserID - Key to access a field that contains the event source Unix effective UID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUserID
kCGEventSourceUserID EventField = 0
// kCGEventTargetProcessSerialNumber - Key to access a field that contains the event target process serial number. The value is a 64-bit long word.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventTargetProcessSerialNumber
kCGEventTargetProcessSerialNumber EventField = 0
// kCGEventTargetUnixProcessID - Key to access a field that contains the event target Unix process ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventTargetUnixProcessID
kCGEventTargetUnixProcessID EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventUnacceleratedPointerMovementX
kCGEventUnacceleratedPointerMovementX EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventUnacceleratedPointerMovementY
kCGEventUnacceleratedPointerMovementY EventField = 0
// kCGKeyboardEventAutorepeat - Key to access an integer field, non-zero when this is an autorepeat of a key-down, and zero otherwise.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventAutorepeat
kCGKeyboardEventAutorepeat EventField = 0
// kCGKeyboardEventKeyboardType - Key to access an integer field that contains the keyboard type identifier.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventKeyboardType
kCGKeyboardEventKeyboardType EventField = 0
// kCGKeyboardEventKeycode - Key to access an integer field that contains the virtual keycode of the key-down or key-up event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventKeycode
kCGKeyboardEventKeycode EventField = 0
// kCGMouseEventButtonNumber - Key to access an integer field that contains the mouse button number. For information about the possible values, see  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventButtonNumber
kCGMouseEventButtonNumber EventField = 0
// kCGMouseEventClickState - Key to access an integer field that contains the mouse button click state. A click state of 1 represents a single click. A click state of 2 represents a double-click. A click state of 3 represents a triple-click.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventClickState
kCGMouseEventClickState EventField = 0
// kCGMouseEventDeltaX - Key to access an integer field that contains the horizontal mouse delta since the last mouse movement event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventDeltaX
kCGMouseEventDeltaX EventField = 0
// kCGMouseEventDeltaY - Key to access an integer field that contains the vertical mouse delta since the last mouse movement event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventDeltaY
kCGMouseEventDeltaY EventField = 0
// kCGMouseEventInstantMouser - Key to access an integer field. The value is non-zero if the event should be ignored by the Inkwell subsystem.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventInstantMouser
kCGMouseEventInstantMouser EventField = 0
// kCGMouseEventNumber - Key to access an integer field that contains the mouse button event number. Matching mouse-down and mouse-up events will have the same event number.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventNumber
kCGMouseEventNumber EventField = 0
// kCGMouseEventPressure - Key to access a double field that contains the mouse button pressure. The pressure value may range from 0 to 1, with 0 representing the mouse being up. This value is commonly set by tablet pens mimicking a mouse.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventPressure
kCGMouseEventPressure EventField = 0
// kCGMouseEventSubtype - Key to access an integer field that encodes the mouse event subtype as a  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventSubtype
kCGMouseEventSubtype EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventWindowUnderMousePointer
kCGMouseEventWindowUnderMousePointer EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventWindowUnderMousePointerThatCanHandleThisEvent
kCGMouseEventWindowUnderMousePointerThatCanHandleThisEvent EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventAcceleratedDeltaAxis1
kCGScrollWheelEventAcceleratedDeltaAxis1 EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventAcceleratedDeltaAxis2
kCGScrollWheelEventAcceleratedDeltaAxis2 EventField = 0
// kCGScrollWheelEventDeltaAxis1 - Key to access an integer field that contains scrolling data. This field typically contains the change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis1
kCGScrollWheelEventDeltaAxis1 EventField = 0
// kCGScrollWheelEventDeltaAxis2 - Key to access an integer field that contains scrolling data. This field typically contains the change in horizontal position since the last scrolling event from a Mighty Mouse scroller.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis2
kCGScrollWheelEventDeltaAxis2 EventField = 0
// kCGScrollWheelEventDeltaAxis3 - This field is not used.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis3
kCGScrollWheelEventDeltaAxis3 EventField = 0
// kCGScrollWheelEventFixedPtDeltaAxis1 - Key to access a field that contains scrolling data. The scrolling data represents a line-based or pixel-based change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller. The scrolling data uses a fixed-point 16.16 signed integer format. For example, if the field contains a value of 1.0, the integer 0x00010000 is returned by  . If this key is passed to  , the fixed-point value is converted to a double value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis1
kCGScrollWheelEventFixedPtDeltaAxis1 EventField = 0
// kCGScrollWheelEventFixedPtDeltaAxis2 - Key to access a field that contains scrolling data. The scrolling data represents a line-based or pixel-based change in horizontal position since the last scrolling event from a Mighty Mouse scroller. The scrolling data uses a fixed-point 16.16 signed integer format. For example, if the field contains a value of 1.0, the integer 0x00010000 is returned by  . If this key is passed to  , the fixed-point value is converted to a double value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis2
kCGScrollWheelEventFixedPtDeltaAxis2 EventField = 0
// kCGScrollWheelEventFixedPtDeltaAxis3 - This field is not used.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis3
kCGScrollWheelEventFixedPtDeltaAxis3 EventField = 0
// kCGScrollWheelEventInstantMouser - Key to access an integer field that indicates whether the event should be ignored by the Inkwell subsystem. If the value is non-zero, the event should be ignored.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventInstantMouser
kCGScrollWheelEventInstantMouser EventField = 0
// kCGScrollWheelEventIsContinuous - Key to access an integer field that indicates whether a scrolling event contains continuous, pixel-based scrolling data. The value is non-zero when the scrolling data is pixel-based and zero when the scrolling data is line-based.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventIsContinuous
kCGScrollWheelEventIsContinuous EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventMomentumOptionPhase
kCGScrollWheelEventMomentumOptionPhase EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventMomentumPhase
kCGScrollWheelEventMomentumPhase EventField = 0
// kCGScrollWheelEventPointDeltaAxis1 - Key to access an integer field that contains pixel-based scrolling data. The scrolling data represents the change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis1
kCGScrollWheelEventPointDeltaAxis1 EventField = 0
// kCGScrollWheelEventPointDeltaAxis2 - Key to access an integer field that contains pixel-based scrolling data. The scrolling data represents the change in horizontal position since the last scrolling event from a Mighty Mouse scroller.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis2
kCGScrollWheelEventPointDeltaAxis2 EventField = 0
// kCGScrollWheelEventPointDeltaAxis3 - This field is not used.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis3
kCGScrollWheelEventPointDeltaAxis3 EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventRawDeltaAxis1
kCGScrollWheelEventRawDeltaAxis1 EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventRawDeltaAxis2
kCGScrollWheelEventRawDeltaAxis2 EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventScrollCount
kCGScrollWheelEventScrollCount EventField = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventScrollPhase
kCGScrollWheelEventScrollPhase EventField = 0
// kCGTabletEventDeviceID - Key to access an integer field that contains the system-assigned unique device ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventDeviceID
kCGTabletEventDeviceID EventField = 0
// kCGTabletEventPointButtons - Key to access an integer field that contains the tablet button state. Bit 0 is the first button, and a set bit represents a closed or pressed button. Up to 16 buttons are supported.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointButtons
kCGTabletEventPointButtons EventField = 0
// kCGTabletEventPointPressure - Key to access a double field that contains the tablet pen pressure. A value of 0.0 represents no pressure, and 1.0 represents maximum pressure.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointPressure
kCGTabletEventPointPressure EventField = 0
// kCGTabletEventPointX - Key to access an integer field that contains the absolute X coordinate in tablet space at full tablet resolution.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointX
kCGTabletEventPointX EventField = 0
// kCGTabletEventPointY - Key to access an integer field that contains the absolute Y coordinate in tablet space at full tablet resolution.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointY
kCGTabletEventPointY EventField = 0
// kCGTabletEventPointZ - Key to access an integer field that contains the absolute Z coordinate in tablet space at full tablet resolution.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointZ
kCGTabletEventPointZ EventField = 0
// kCGTabletEventRotation - Key to access a double field that contains the tablet pen rotation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventRotation
kCGTabletEventRotation EventField = 0
// kCGTabletEventTangentialPressure - Key to access a double field that contains the tangential pressure on the device. A value of 0.0 represents no pressure, and 1.0 represents maximum pressure.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTangentialPressure
kCGTabletEventTangentialPressure EventField = 0
// kCGTabletEventTiltX - Key to access a double field that contains the horizontal tablet pen tilt. A value of 0.0 represents no tilt, and 1.0 represents maximum tilt.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTiltX
kCGTabletEventTiltX EventField = 0
// kCGTabletEventTiltY - Key to access a double field that contains the vertical tablet pen tilt. A value of 0.0 represents no tilt, and 1.0 represents maximum tilt.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTiltY
kCGTabletEventTiltY EventField = 0
// kCGTabletEventVendor1 - Key to access an integer field that contains a vendor-specified value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor1
kCGTabletEventVendor1 EventField = 0
// kCGTabletEventVendor2 - Key to access an integer field that contains a vendor-specified value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor2
kCGTabletEventVendor2 EventField = 0
// kCGTabletEventVendor3 - Key to access an integer field that contains a vendor-specified value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor3
kCGTabletEventVendor3 EventField = 0
// kCGTabletProximityEventCapabilityMask - Key to access an integer field that contains the device capabilities mask.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventCapabilityMask
kCGTabletProximityEventCapabilityMask EventField = 0
// kCGTabletProximityEventDeviceID - Key to access an integer field that contains the system-assigned device ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventDeviceID
kCGTabletProximityEventDeviceID EventField = 0
// kCGTabletProximityEventEnterProximity - Key to access an integer field that indicates whether the pen is in proximity to the tablet. The value is non-zero if the pen is in proximity to the tablet and zero when leaving the tablet.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventEnterProximity
kCGTabletProximityEventEnterProximity EventField = 0
// kCGTabletProximityEventPointerID - Key to access an integer field that contains the vendor-defined ID of the pointing device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventPointerID
kCGTabletProximityEventPointerID EventField = 0
// kCGTabletProximityEventPointerType - Key to access an integer field that contains the pointer type.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventPointerType
kCGTabletProximityEventPointerType EventField = 0
// kCGTabletProximityEventSystemTabletID - Key to access an integer field that contains the system-assigned unique tablet ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventSystemTabletID
kCGTabletProximityEventSystemTabletID EventField = 0
// kCGTabletProximityEventTabletID - Key to access an integer field that contains the vendor-defined tablet ID, typically the USB product ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventTabletID
kCGTabletProximityEventTabletID EventField = 0
// kCGTabletProximityEventVendorID - Key to access an integer field that contains the vendor-defined ID, typically the USB vendor ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorID
kCGTabletProximityEventVendorID EventField = 0
// kCGTabletProximityEventVendorPointerSerialNumber - Key to access an integer field that contains the vendor-defined pointer serial number.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorPointerSerialNumber
kCGTabletProximityEventVendorPointerSerialNumber EventField = 0
// kCGTabletProximityEventVendorPointerType - Key to access an integer field that contains the vendor-assigned pointer type.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorPointerType
kCGTabletProximityEventVendorPointerType EventField = 0
// kCGTabletProximityEventVendorUniqueID - Key to access an integer field that contains the vendor-defined unique ID.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorUniqueID
kCGTabletProximityEventVendorUniqueID EventField = 0
)

// CGEventFilterMask - Specify masks for classes of low-level events that can be filtered during event suppression states.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask
type EventFilterMask uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitLocalKeyboardEvents
kCGEventFilterMaskPermitLocalKeyboardEvents EventFilterMask = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitLocalMouseEvents
kCGEventFilterMaskPermitLocalMouseEvents EventFilterMask = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitSystemDefinedEvents
kCGEventFilterMaskPermitSystemDefinedEvents EventFilterMask = 0
)

// CGEventFlags - Constants that indicate the modifier key state at the time an event is created, as well as other event-related states.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags
type EventFlags uint

const (
// kCGEventFlagMaskAlphaShift - Indicates that the Caps Lock key is down for a keyboard, mouse, or flag-changed event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskAlphaShift
kCGEventFlagMaskAlphaShift EventFlags = 0
// kCGEventFlagMaskAlternate - Indicates that the Alt or Option key is down for a keyboard, mouse, or flag-changed event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskAlternate
kCGEventFlagMaskAlternate EventFlags = 0
// kCGEventFlagMaskCommand - Indicates that the Command key is down for a keyboard, mouse, or flag-changed event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskCommand
kCGEventFlagMaskCommand EventFlags = 0
// kCGEventFlagMaskControl - Indicates that the Control key is down for a keyboard, mouse, or flag-changed event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskControl
kCGEventFlagMaskControl EventFlags = 0
// kCGEventFlagMaskHelp - Indicates that the Help modifier key is down for a keyboard, mouse, or flag-changed event. This key is not present on most keyboards, and is different than the Help key found in the same row as Home and Page Up.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskHelp
kCGEventFlagMaskHelp EventFlags = 0
// kCGEventFlagMaskNonCoalesced - Indicates that mouse and pen movement events are not being coalesced.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskNonCoalesced
kCGEventFlagMaskNonCoalesced EventFlags = 0
// kCGEventFlagMaskNumericPad - Identifies key events from the numeric keypad area on extended keyboards.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskNumericPad
kCGEventFlagMaskNumericPad EventFlags = 0
// kCGEventFlagMaskSecondaryFn - Indicates that the Fn (Function) key is down for a keyboard, mouse, or flag-changed event. This key is found primarily on laptop keyboards.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskSecondaryFn
kCGEventFlagMaskSecondaryFn EventFlags = 0
// kCGEventFlagMaskShift - Indicates that the Shift key is down for a keyboard, mouse, or flag-changed event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskShift
kCGEventFlagMaskShift EventFlags = 0
)

// CGEventMouseSubtype - Constants used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype
type EventMouseSubtype uint

const (
// kCGEventMouseSubtypeDefault - Specifies that the event is an ordinary mouse event, and does not contain additional tablet device information.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/defaultType
kCGEventMouseSubtypeDefault EventMouseSubtype = 0
// kCGEventMouseSubtypeTabletPoint - Specifies that the mouse event originated from a tablet device, and that the various   field selectors may be used to obtain tablet-specific data from the mouse event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/tabletPoint
kCGEventMouseSubtypeTabletPoint EventMouseSubtype = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/tabletProximity
kCGEventMouseSubtypeTabletProximity EventMouseSubtype = 0
)

// CGEventSourceStateID - Constants that specify the possible source states of an event source.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID
type EventSourceStateID uint

const (
// kCGEventSourceStateCombinedSessionState - Specifies that an event source should use the event state table that reflects the combined state of all event sources posting to the current user login session.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/combinedSessionState
kCGEventSourceStateCombinedSessionState EventSourceStateID = 0
// kCGEventSourceStateHIDSystemState - Specifies that an event source should use the event state table that reflects the combined state of all hardware event sources posting from the HID system.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/hidSystemState
kCGEventSourceStateHIDSystemState EventSourceStateID = 0
// kCGEventSourceStatePrivate - Specifies that an event source should use a private event state table.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/privateState
kCGEventSourceStatePrivate EventSourceStateID = 0
)

// CGEventSuppressionState - Specify the event suppression states that can occur after posting an event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState
type EventSuppressionState uint

const (
// kCGEventSuppressionStateRemoteMouseDrag - Specifies that certain local hardware events may be suppressed during a mouse drag operation (mouse movement with the left or only mouse button down).
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/eventSuppressionStateRemoteMouseDrag
kCGEventSuppressionStateRemoteMouseDrag EventSuppressionState = 0
// kCGEventSuppressionStateSuppressionInterval - Specifies that certain local hardware events may be suppressed for a short interval after posting an event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/eventSuppressionStateSuppressionInterval
kCGEventSuppressionStateSuppressionInterval EventSuppressionState = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/numberOfEventSuppressionStates
kCGNumberOfEventSuppressionStates EventSuppressionState = 0
)

// CGEventTapLocation - Constants that specify possible tapping points for events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation
type EventTapLocation uint

const (
// kCGAnnotatedSessionEventTap - Specifies that an event tap is placed at the point where session events have been annotated to flow to an application.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cgAnnotatedSessionEventTap
kCGAnnotatedSessionEventTap EventTapLocation = 0
// kCGSessionEventTap - Specifies that an event tap is placed at the point where HID system and remote control events enter a login session.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cgSessionEventTap
kCGSessionEventTap EventTapLocation = 0
// kCGHIDEventTap - Specifies that an event tap is placed at the point where HID system events enter the window server.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cghidEventTap
kCGHIDEventTap EventTapLocation = 0
)

// CGEventTapOptions - Constants that specify whether a new event tap is an active filter or a passive listener.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions
type EventTapOptions uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions/defaultTap
kCGEventTapOptionDefault EventTapOptions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions/listenOnly
kCGEventTapOptionListenOnly EventTapOptions = 0
)

// CGEventTapPlacement - Constants that specify where a new event tap is inserted into the list of active event taps.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement
type EventTapPlacement uint

const (
// kCGHeadInsertEventTap - Specifies that a new event tap should be inserted before any pre-existing event taps at the same location.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement/headInsertEventTap
kCGHeadInsertEventTap EventTapPlacement = 0
// kCGTailAppendEventTap - Specifies that a new event tap should be inserted after any pre-existing event taps at the same location.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement/tailAppendEventTap
kCGTailAppendEventTap EventTapPlacement = 0
)

// CGEventType - Constants that specify the different types of input events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType
type EventType uint

const (
// kCGEventFlagsChanged - Specifies a key changed event for a modifier or status key.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/flagsChanged
kCGEventFlagsChanged EventType = 0
// kCGEventKeyDown - Specifies a key down event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/keyDown
kCGEventKeyDown EventType = 0
// kCGEventKeyUp - Specifies a key up event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/keyUp
kCGEventKeyUp EventType = 0
// kCGEventLeftMouseDown - Specifies a mouse down event with the left button.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseDown
kCGEventLeftMouseDown EventType = 0
// kCGEventLeftMouseDragged - Specifies a mouse drag event with the left button down.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseDragged
kCGEventLeftMouseDragged EventType = 0
// kCGEventLeftMouseUp - Specifies a mouse up event with the left button.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseUp
kCGEventLeftMouseUp EventType = 0
// kCGEventMouseMoved - Specifies a mouse moved event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/mouseMoved
kCGEventMouseMoved EventType = 0
// kCGEventNull - Specifies a null event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/null
kCGEventNull EventType = 0
// kCGEventOtherMouseDown - Specifies a mouse down event with one of buttons 2-31.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseDown
kCGEventOtherMouseDown EventType = 0
// kCGEventOtherMouseDragged - Specifies a mouse drag event with one of buttons 2-31 down.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseDragged
kCGEventOtherMouseDragged EventType = 0
// kCGEventOtherMouseUp - Specifies a mouse up event with one of buttons 2-31.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseUp
kCGEventOtherMouseUp EventType = 0
// kCGEventRightMouseDown - Specifies a mouse down event with the right button.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseDown
kCGEventRightMouseDown EventType = 0
// kCGEventRightMouseDragged - Specifies a mouse drag event with the right button down.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseDragged
kCGEventRightMouseDragged EventType = 0
// kCGEventRightMouseUp - Specifies a mouse up event with the right button.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseUp
kCGEventRightMouseUp EventType = 0
// kCGEventScrollWheel - Specifies a scroll wheel moved event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/scrollWheel
kCGEventScrollWheel EventType = 0
// kCGEventTabletPointer - Specifies a tablet pointer event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tabletPointer
kCGEventTabletPointer EventType = 0
// kCGEventTabletProximity - Specifies a tablet proximity event.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tabletProximity
kCGEventTabletProximity EventType = 0
// kCGEventTapDisabledByTimeout - Specifies an event indicating the event tap is disabled because of timeout.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tapDisabledByTimeout
kCGEventTapDisabledByTimeout EventType = 0
// kCGEventTapDisabledByUserInput - Specifies an event indicating the event tap is disabled because of user input.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tapDisabledByUserInput
kCGEventTapDisabledByUserInput EventType = 0
)

// CGFontPostScriptFormat - Possible formats for a PostScript font subset.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat
type FontPostScriptFormat uint

const (
// kCGFontPostScriptFormatType1 - A Type 1 font format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type1
kCGFontPostScriptFormatType1 FontPostScriptFormat = 0
// kCGFontPostScriptFormatType3 - A Type 3 PostScript format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type3
kCGFontPostScriptFormatType3 FontPostScriptFormat = 0
// kCGFontPostScriptFormatType42 - A constant representing a Type 42 font format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type42
kCGFontPostScriptFormatType42 FontPostScriptFormat = 0
)

// CGGesturePhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase
type GesturePhase uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/began
kCGGesturePhaseBegan GesturePhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/cancelled
kCGGesturePhaseCancelled GesturePhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/changed
kCGGesturePhaseChanged GesturePhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/ended
kCGGesturePhaseEnded GesturePhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/mayBegin
kCGGesturePhaseMayBegin GesturePhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/none
kCGGesturePhaseNone GesturePhase = 0
)

// CGGlyphDeprecatedEnum enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum
type GlyphDeprecatedEnum uint

const (
// GlyphMax - Maximum font index value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum/max
GlyphMax GlyphDeprecatedEnum = 0
// GlyphMin - Minimum font index value.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum/min
GlyphMin GlyphDeprecatedEnum = 0
)

// CGGradientDrawingOptions - Drawing locations for gradients.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions
type GradientDrawingOptions uint

const (
// kCGGradientDrawsAfterEndLocation - The fill should extend beyond the ending location. The color that extends beyond the ending point is the solid color defined by the   object to be at location 1.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions/drawsAfterEndLocation
kCGGradientDrawsAfterEndLocation GradientDrawingOptions = 0
// kCGGradientDrawsBeforeStartLocation - The fill should extend beyond the starting location. The color that extends beyond the starting point is the solid color defined by the   object to be at location 0.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions/drawsBeforeStartLocation
kCGGradientDrawsBeforeStartLocation GradientDrawingOptions = 0
)

// CGImageAlphaInfo - Storage options for alpha component data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo
type ImageAlphaInfo uint

const (
// kCGImageAlphaOnly - There is no color data, only an alpha channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/alphaOnly
kCGImageAlphaOnly ImageAlphaInfo = 0
// kCGImageAlphaFirst - The alpha component is stored in the most significant bits of each pixel. For example, non-premultiplied ARGB.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/first
kCGImageAlphaFirst ImageAlphaInfo = 0
// kCGImageAlphaLast - The alpha component is stored in the least significant bits of each pixel. For example, non-premultiplied RGBA.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/last
kCGImageAlphaLast ImageAlphaInfo = 0
// kCGImageAlphaNone - There is no alpha channel.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/none
kCGImageAlphaNone ImageAlphaInfo = 0
)

// CGImageByteOrderInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo
type ImageByteOrderInfo uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Big
kCGImageByteOrder16Big ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Host
kCGImageByteOrder16Host ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Little
kCGImageByteOrder16Little ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Big
kCGImageByteOrder32Big ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Host
kCGImageByteOrder32Host ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Little
kCGImageByteOrder32Little ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/orderDefault
kCGImageByteOrderDefault ImageByteOrderInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/orderMask
kCGImageByteOrderMask ImageByteOrderInfo = 0
)

// CGImageComponentInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo
type ImageComponentInfo uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo/float
kCGImageComponentFloat ImageComponentInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo/integer
kCGImageComponentInteger ImageComponentInfo = 0
)

// CGImagePixelFormatInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo
type ImagePixelFormatInfo uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB101010
kCGImagePixelFormatRGB101010 ImagePixelFormatInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB555
kCGImagePixelFormatRGB555 ImagePixelFormatInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB565
kCGImagePixelFormatRGB565 ImagePixelFormatInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGBCIF10
kCGImagePixelFormatRGBCIF10 ImagePixelFormatInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/mask
kCGImagePixelFormatMask ImagePixelFormatInfo = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/packed
kCGImagePixelFormatPacked ImagePixelFormatInfo = 0
)

// CGInterpolationQuality - Levels of interpolation quality for rendering an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality
type InterpolationQuality uint

const (
// kCGInterpolationDefault - The default level of quality.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/default
kCGInterpolationDefault InterpolationQuality = 0
// kCGInterpolationHigh - A high level of interpolation quality. This setting may slow down image rendering.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/high
kCGInterpolationHigh InterpolationQuality = 0
// kCGInterpolationLow - A low level of interpolation quality. This setting may speed up image rendering.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/low
kCGInterpolationLow InterpolationQuality = 0
// kCGInterpolationMedium - A medium level of interpolation quality. This setting is slower than the low setting but faster than the high setting.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/medium
kCGInterpolationMedium InterpolationQuality = 0
// kCGInterpolationNone - No interpolation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/none
kCGInterpolationNone InterpolationQuality = 0
)

// CGLineCap - Styles for rendering the endpoint of a stroked line.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap
type LineCap uint

const (
// kCGLineCapButt - A line with a squared-off end. Core Graphics draws the line to extend only to the exact endpoint of the path. This is the default.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/butt
kCGLineCapButt LineCap = 0
// kCGLineCapRound - A line with a rounded end. Core Graphics draws the line to extend beyond the endpoint of the path. The line ends with a semicircular arc with a radius of 1/2 the line’s width, centered on the endpoint.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/round
kCGLineCapRound LineCap = 0
// kCGLineCapSquare - A line with a squared-off end. Core Graphics extends the line beyond the endpoint of the path for a distance equal to half the line width.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/square
kCGLineCapSquare LineCap = 0
)

// CGLineJoin - Junction types for stroked lines.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin
type LineJoin uint

const (
// kCGLineJoinBevel - A join with a squared-off end. Core Graphics draws the line to extend beyond the endpoint of the path, for a distance of 1/2 the line’s width.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/bevel
kCGLineJoinBevel LineJoin = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/miter
kCGLineJoinMiter LineJoin = 0
// kCGLineJoinRound - A join with a rounded end. Core Graphics draws the line to extend beyond the endpoint of the path. The line ends with a semicircular arc with a radius of 1/2 the line’s width, centered on the endpoint.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/round
kCGLineJoinRound LineJoin = 0
)

// CGMomentumScrollPhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase
type MomentumScrollPhase uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/begin
kCGMomentumScrollPhaseBegin MomentumScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/continuous
kCGMomentumScrollPhaseContinue MomentumScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/end
kCGMomentumScrollPhaseEnd MomentumScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/none
kCGMomentumScrollPhaseNone MomentumScrollPhase = 0
)

// CGMouseButton - Constants that specify buttons on a one, two, or three-button mouse.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton
type MouseButton uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/center
kCGMouseButtonCenter MouseButton = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/left
kCGMouseButtonLeft MouseButton = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/right
kCGMouseButtonRight MouseButton = 0
)

// CGPDFAccessPermissions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions
type PDFAccessPermissions uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsCommenting
kCGPDFAllowsCommenting PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsContentAccessibility
kCGPDFAllowsContentAccessibility PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsContentCopying
kCGPDFAllowsContentCopying PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsDocumentAssembly
kCGPDFAllowsDocumentAssembly PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsDocumentChanges
kCGPDFAllowsDocumentChanges PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsFormFieldEntry
kCGPDFAllowsFormFieldEntry PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsHighQualityPrinting
kCGPDFAllowsHighQualityPrinting PDFAccessPermissions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsLowQualityPrinting
kCGPDFAllowsLowQualityPrinting PDFAccessPermissions = 0
)

// CGPDFBox - Box types for a PDF page.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox
type PDFBox uint

const (
// kCGPDFArtBox - The page art box—a rectangle, expressed in default user space units, defining the extent of the page’s meaningful content (including potential white space) as intended by the page’s creator.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/artBox
kCGPDFArtBox PDFBox = 0
// kCGPDFBleedBox - The page bleed box—a rectangle, expressed in default user space units, that defines the region to which the contents of the page should be clipped when output in a production environment.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/bleedBox
kCGPDFBleedBox PDFBox = 0
// kCGPDFCropBox - The page crop box—a rectangle, expressed in default user space units, that defines the visible region of default user space. When the page is displayed or printed, its contents are to be clipped to this rectangle.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/cropBox
kCGPDFCropBox PDFBox = 0
// kCGPDFMediaBox - The page media box—a rectangle, expressed in default user space units, that defines the boundaries of the physical medium on which the page is intended to be displayed or printed
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/mediaBox
kCGPDFMediaBox PDFBox = 0
// kCGPDFTrimBox - The page trim box—a rectangle, expressed in default user space units, that defines the intended dimensions of the finished page after trimming.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/trimBox
kCGPDFTrimBox PDFBox = 0
)

// CGPDFDataFormat - The encoding format of PDF data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat
type PDFDataFormat uint

const (
// PDFDataFormatJPEG2000 - The data stream is encoded in JPEG-2000 format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/JPEG2000
PDFDataFormatJPEG2000 PDFDataFormat = 0
// PDFDataFormatJPEGEncoded - The data stream is encoded in JPEG format.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/jpegEncoded
PDFDataFormatJPEGEncoded PDFDataFormat = 0
// PDFDataFormatRaw - The data stream is not encoded.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/raw
PDFDataFormatRaw PDFDataFormat = 0
)

// CGPDFObjectType - Types of PDF object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType
type PDFObjectType uint

const (
// kCGPDFObjectTypeArray - Type for a PDF array.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/array
kCGPDFObjectTypeArray PDFObjectType = 0
// kCGPDFObjectTypeBoolean - The type for a PDF Boolean.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/boolean
kCGPDFObjectTypeBoolean PDFObjectType = 0
// kCGPDFObjectTypeDictionary - The type for a PDF dictionary.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/dictionary
kCGPDFObjectTypeDictionary PDFObjectType = 0
// kCGPDFObjectTypeInteger - The type for a PDF integer.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/integer
kCGPDFObjectTypeInteger PDFObjectType = 0
// kCGPDFObjectTypeName - Type for a PDF name.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/name
kCGPDFObjectTypeName PDFObjectType = 0
// kCGPDFObjectTypeNull - The type for a PDF null.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/null
kCGPDFObjectTypeNull PDFObjectType = 0
// kCGPDFObjectTypeReal - The type for a PDF real.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/real
kCGPDFObjectTypeReal PDFObjectType = 0
// kCGPDFObjectTypeStream - The type for a PDF stream.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/stream
kCGPDFObjectTypeStream PDFObjectType = 0
// kCGPDFObjectTypeString - The type for a PDF string.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/string
kCGPDFObjectTypeString PDFObjectType = 0
)

// CGPDFTagType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType
type PDFTagType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/TOC
PDFTagTypeTOC PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/TOCI
PDFTagTypeTOCI PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/annotation
PDFTagTypeAnnotation PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/art
PDFTagTypeArt PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/bibliography
PDFTagTypeBibliography PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/blockQuote
PDFTagTypeBlockQuote PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/caption
PDFTagTypeCaption PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/code
PDFTagTypeCode PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/div
PDFTagTypeDiv PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/document
PDFTagTypeDocument PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/figure
PDFTagTypeFigure PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/form
PDFTagTypeForm PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/formula
PDFTagTypeFormula PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header
PDFTagTypeHeader PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header1
PDFTagTypeHeader1 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header2
PDFTagTypeHeader2 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header3
PDFTagTypeHeader3 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header4
PDFTagTypeHeader4 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header5
PDFTagTypeHeader5 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header6
PDFTagTypeHeader6 PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/index
PDFTagTypeIndex PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/label
PDFTagTypeLabel PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/link
PDFTagTypeLink PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/list
PDFTagTypeList PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/listBody
PDFTagTypeListBody PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/listItem
PDFTagTypeListItem PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/nonStructure
PDFTagTypeNonStructure PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/note
PDFTagTypeNote PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/object
PDFTagTypeObject PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/paragraph
PDFTagTypeParagraph PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/part
PDFTagTypePart PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/private
PDFTagTypePrivate PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/quote
PDFTagTypeQuote PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/reference
PDFTagTypeReference PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/ruby
PDFTagTypeRuby PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyAnnotationText
PDFTagTypeRubyAnnotationText PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyBaseText
PDFTagTypeRubyBaseText PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyPunctuation
PDFTagTypeRubyPunctuation PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/section
PDFTagTypeSection PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/span
PDFTagTypeSpan PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/table
PDFTagTypeTable PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableBody
PDFTagTypeTableBody PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableDataCell
PDFTagTypeTableDataCell PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableFooter
PDFTagTypeTableFooter PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableHeader
PDFTagTypeTableHeader PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableHeaderCell
PDFTagTypeTableHeaderCell PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableRow
PDFTagTypeTableRow PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichu
PDFTagTypeWarichu PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichuPunctiation
PDFTagTypeWarichuPunctiation PDFTagType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichuText
PDFTagTypeWarichuText PDFTagType = 0
)

// CGPathDrawingMode - Options for rendering a path.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode
type PathDrawingMode uint

const (
// kCGPathEOFill - Render the area within the path using the even-odd rule.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/eoFill
kCGPathEOFill PathDrawingMode = 0
// kCGPathEOFillStroke - First fill and then stroke the path, using the even-odd rule.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/eoFillStroke
kCGPathEOFillStroke PathDrawingMode = 0
// kCGPathFill - Render the area contained within the path using the non-zero winding number rule.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/fill
kCGPathFill PathDrawingMode = 0
// kCGPathFillStroke - First fill and then stroke the path, using the nonzero winding number rule.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/fillStroke
kCGPathFillStroke PathDrawingMode = 0
// kCGPathStroke - Render a line along the path.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/stroke
kCGPathStroke PathDrawingMode = 0
)

// CGPathElementType - The type of element found in a path.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType
type PathElementType uint

const (
// kCGPathElementAddCurveToPoint - The path element that adds a cubic curve from the current point to the specified point.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addCurveToPoint
kCGPathElementAddCurveToPoint PathElementType = 0
// kCGPathElementAddLineToPoint - The path element that adds a line from the current point to a new point.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addLineToPoint
kCGPathElementAddLineToPoint PathElementType = 0
// kCGPathElementAddQuadCurveToPoint - The path element that adds a quadratic curve from the current point to the specified point.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addQuadCurveToPoint
kCGPathElementAddQuadCurveToPoint PathElementType = 0
// kCGPathElementCloseSubpath - The path element that closes and completes a subpath. The element does not contain any points. See the function  .
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/closeSubpath
kCGPathElementCloseSubpath PathElementType = 0
// kCGPathElementMoveToPoint - The path element that starts a new subpath.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/moveToPoint
kCGPathElementMoveToPoint PathElementType = 0
)

// CGPatternTiling - Different methods for rendering a tiled pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling
type PatternTiling uint

const (
// kCGPatternTilingConstantSpacing - Pattern cells are spaced consistently, as with  .The pattern cell may be distorted additionally to permit a moreefficient implementation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/constantSpacing
kCGPatternTilingConstantSpacing PatternTiling = 0
// kCGPatternTilingConstantSpacingMinimalDistortion - Pattern cells are spaced consistently. Thepattern cell may be distorted by as much as 1 device pixel whenthe pattern is painted.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/constantSpacingMinimalDistortion
kCGPatternTilingConstantSpacingMinimalDistortion PatternTiling = 0
// kCGPatternTilingNoDistortion - The pattern cell is not distorted when painted.The spacing between pattern cells may vary by as much as 1 devicepixel.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/noDistortion
kCGPatternTilingNoDistortion PatternTiling = 0
)

// CGScreenUpdateOperation - Types of screen-update operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation
type ScreenUpdateOperation uint

const (
// kCGScreenUpdateOperationMove - A screen-move operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/move
kCGScreenUpdateOperationMove ScreenUpdateOperation = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/reducedDirtyRectangleCount
kCGScreenUpdateOperationReducedDirtyRectangleCount ScreenUpdateOperation = 0
// kCGScreenUpdateOperationRefresh - A screen-refresh operation.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/refresh
kCGScreenUpdateOperationRefresh ScreenUpdateOperation = 0
)

// CGScrollEventUnit - Constants that specify the unit of measurement for a scrolling event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit
type ScrollEventUnit uint

const (
// kCGScrollEventUnitLine - Specifies that the unit of measurement is lines.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit/line
kCGScrollEventUnitLine ScrollEventUnit = 0
// kCGScrollEventUnitPixel - Specifies that the unit of measurement is pixels.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit/pixel
kCGScrollEventUnitPixel ScrollEventUnit = 0
)

// CGScrollPhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase
type ScrollPhase uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/began
kCGScrollPhaseBegan ScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/cancelled
kCGScrollPhaseCancelled ScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/changed
kCGScrollPhaseChanged ScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/ended
kCGScrollPhaseEnded ScrollPhase = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/mayBegin
kCGScrollPhaseMayBegin ScrollPhase = 0
)

// CGTextDrawingMode - Modes for rendering text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode
type TextDrawingMode uint

const (
// kCGTextClip - Specifies to intersect the text with the current clipping path. This mode does not paint the text.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/clip
kCGTextClip TextDrawingMode = 0
// kCGTextFill - Perform a fill operation on the text.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fill
kCGTextFill TextDrawingMode = 0
// kCGTextFillClip - Perform a fill operation, then intersect the text with the current clipping path.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillClip
kCGTextFillClip TextDrawingMode = 0
// kCGTextFillStroke - Perform fill, then stroke operations on the text.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillStroke
kCGTextFillStroke TextDrawingMode = 0
// kCGTextFillStrokeClip - Perform fill then stroke operations, then intersect the text with the current clipping path.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillStrokeClip
kCGTextFillStrokeClip TextDrawingMode = 0
// kCGTextInvisible - Do not draw the text, but do update the text position.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/invisible
kCGTextInvisible TextDrawingMode = 0
// kCGTextStroke - Perform a stroke operation on the text.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/stroke
kCGTextStroke TextDrawingMode = 0
// kCGTextStrokeClip - Perform a stroke operation, then intersect the text with the current clipping path.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/strokeClip
kCGTextStrokeClip TextDrawingMode = 0
)

// CGTextEncoding - Text encodings for fonts.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding
type TextEncoding uint

const (
// kCGEncodingFontSpecific - The built-in encoding of the font.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding/encodingFontSpecific
kCGEncodingFontSpecific TextEncoding = 0
// kCGEncodingMacRoman - The MacRoman encoding. MacRoman is an ASCII variant originally created for use in the Mac OS, in which characters 127 and lower are ASCII, and characters 128 and higher are non-English characters and symbols.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding/encodingMacRoman
kCGEncodingMacRoman TextEncoding = 0
)

// CGToneMapping enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping
type ToneMapping uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/default
kCGToneMappingDefault ToneMapping = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/exrGamma
kCGToneMappingEXRGamma ToneMapping = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/imageSpecificLumaScaling
kCGToneMappingImageSpecificLumaScaling ToneMapping = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/ituRecommended
kCGToneMappingITURecommended ToneMapping = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/none
kCGToneMappingNone ToneMapping = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/referenceWhiteBased
kCGToneMappingReferenceWhiteBased ToneMapping = 0
)

// CGWindowBackingType - The data type used to specify the backing option for a given window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType
type WindowBackingType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreBuffered
kCGBackingStoreBuffered WindowBackingType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreNonretained
kCGBackingStoreNonretained WindowBackingType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreRetained
kCGBackingStoreRetained WindowBackingType = 0
)

// CGWindowImageOption - The data type to use to specify the type of image to be generated for a window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption
type WindowImageOption uint

const (
// kCGWindowImageBestResolution - When capturing the window, return the best image resolution. The returned image size may be different than the screen size.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/bestResolution
kCGWindowImageBestResolution WindowImageOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/boundsIgnoreFraming
kCGWindowImageBoundsIgnoreFraming WindowImageOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/kCGWindowImageDefault
kCGWindowImageDefault WindowImageOption = 0
// kCGWindowImageNominalResolution - When capturing the window, return the nominal image resolution. The returned image size is the same as the screen size.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/nominalResolution
kCGWindowImageNominalResolution WindowImageOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/onlyShadows
kCGWindowImageOnlyShadows WindowImageOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/shouldBeOpaque
kCGWindowImageShouldBeOpaque WindowImageOption = 0
)

// CGWindowLevelKey - Keys that represent the standard window levels in macOS. Quartz includes these keys to support application frameworks like Cocoa. Applications do not need to use them directly.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey
type WindowLevelKey uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/assistiveTechHighWindow
kCGAssistiveTechHighWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/backstopMenu
kCGBackstopMenuLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/baseWindow
kCGBaseWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/cursorWindow
kCGCursorWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/desktopIconWindow
kCGDesktopIconWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/desktopWindow
kCGDesktopWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/dockWindow
kCGDockWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/draggingWindow
kCGDraggingWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/floatingWindow
kCGFloatingWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/helpWindow
kCGHelpWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/mainMenuWindow
kCGMainMenuWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/maximumWindow
kCGMaximumWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/minimumWindow
kCGMinimumWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/modalPanelWindow
kCGModalPanelWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/normalWindow
kCGNormalWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/numberOfWindowLevelKeys
kCGNumberOfWindowLevelKeys WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/overlayWindow
kCGOverlayWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/popUpMenuWindow
kCGPopUpMenuWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/screenSaverWindow
kCGScreenSaverWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/statusWindow
kCGStatusWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/tornOffMenuWindow
kCGTornOffMenuWindowLevelKey WindowLevelKey = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/utilityWindow
kCGUtilityWindowLevelKey WindowLevelKey = 0
)

// CGWindowListOption - The data type used to specify the options for gathering a list of windows.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption
type WindowListOption uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/excludeDesktopElements
kCGWindowListExcludeDesktopElements WindowListOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionAll
kCGWindowListOptionAll WindowListOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionIncludingWindow
kCGWindowListOptionIncludingWindow WindowListOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenAboveWindow
kCGWindowListOptionOnScreenAboveWindow WindowListOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenBelowWindow
kCGWindowListOptionOnScreenBelowWindow WindowListOption = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenOnly
kCGWindowListOptionOnScreenOnly WindowListOption = 0
)

// CGWindowSharingType - The data type used to specify the sharing mode used by a window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType
type WindowSharingType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/none
kCGWindowSharingNone WindowSharingType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/readOnly
kCGWindowSharingReadOnly WindowSharingType = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/readWrite
kCGWindowSharingReadWrite WindowSharingType = 0
)


