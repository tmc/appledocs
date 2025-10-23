// Code generated from Apple documentation for CoreGraphics. DO NOT EDIT.

package coregraphics

// Enum types and constants
// CGBitmapLayout enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout
type CGBitmapLayout uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/abgr
	kCGBitmapLayoutABGR CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/alphaOnly
	kCGBitmapLayoutAlphaOnly CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/argb
	kCGBitmapLayoutARGB CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/bgra
	kCGBitmapLayoutBGRA CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/bgrx
	kCGBitmapLayoutBGRX CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/cmyk
	kCGBitmapLayoutCMYK CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/gray
	kCGBitmapLayoutGray CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/grayAlpha
	kCGBitmapLayoutGrayAlpha CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/rgba
	kCGBitmapLayoutRGBA CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/rgbx
	kCGBitmapLayoutRGBX CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/xbgr
	kCGBitmapLayoutXBGR CGBitmapLayout = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapLayout/xrgb
	kCGBitmapLayoutXRGB CGBitmapLayout = 0
)

// CGColorModel enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel
type CGColorModel uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/cmyk
	kCGColorModelCMYK CGColorModel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/deviceN
	kCGColorModelDeviceN CGColorModel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/gray
	kCGColorModelGray CGColorModel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/kCGColorModelNoColorant
	kCGColorModelNoColorant CGColorModel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/lab
	kCGColorModelLab CGColorModel = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorModel/rgb
	kCGColorModelRGB CGColorModel = 0
)

// CGComponent enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent
type CGComponent uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/float16Bit
	kCGComponentFloat16Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/float32Bit
	kCGComponentFloat32Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer10Bit
	kCGComponentInteger10Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer16Bit
	kCGComponentInteger16Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer32Bit
	kCGComponentInteger32Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/integer8Bit
	kCGComponentInteger8Bit CGComponent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGComponent/unknown
	kCGComponentUnknown CGComponent = 0
)

// CGImageComponentInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo
type CGImageComponentInfo uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo/float
	kCGImageComponentFloat CGImageComponentInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageComponentInfo/integer
	kCGImageComponentInteger CGImageComponentInfo = 0
)

// CGBitmapInfo - Component information for a bitmap image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo
type CGBitmapInfo uint

const (
	// kCGBitmapByteOrder16Big - 16-bit, big endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder16Big
	kCGBitmapByteOrder16Big CGBitmapInfo = 0
	// kCGBitmapByteOrder16Little - 16-bit, little endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder16Little
	kCGBitmapByteOrder16Little CGBitmapInfo = 0
	// kCGBitmapByteOrder32Big - 32-bit, big endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder32Big
	kCGBitmapByteOrder32Big CGBitmapInfo = 0
	// kCGBitmapByteOrder32Little - 32-bit, little endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrder32Little
	kCGBitmapByteOrder32Little CGBitmapInfo = 0
	// kCGBitmapByteOrderDefault - The default byte order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/byteOrderDefault
	kCGBitmapByteOrderDefault CGBitmapInfo = 0
	// kCGBitmapFloatComponents - The components of a bitmap are floating-point values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/floatComponents
	kCGBitmapFloatComponents CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapAlphaInfoMask
	kCGBitmapAlphaInfoMask CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapByteOrderInfoMask
	kCGBitmapByteOrderInfoMask CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapByteOrderMask
	kCGBitmapByteOrderMask CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapComponentInfoMask
	kCGBitmapComponentInfoMask CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapFloatInfoMask
	kCGBitmapFloatInfoMask CGBitmapInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBitmapInfo/kCGBitmapPixelFormatInfoMask
	kCGBitmapPixelFormatInfoMask CGBitmapInfo = 0
)

// CGBlendMode - Compositing operations for images.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode
type CGBlendMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/clear
	kCGBlendModeClear CGBlendMode = 0
	// kCGBlendModeColor - Uses the luminance values of the background with the hue and saturation values of the source image. This mode preserves the gray levels in the image. You can use this mode to color monochrome images or to tint color images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/color
	kCGBlendModeColor CGBlendMode = 0
	// kCGBlendModeColorBurn - Darkens the background image samples to reflect the source image samples. Source image sample values that specify white do not produce a change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/colorBurn
	kCGBlendModeColorBurn CGBlendMode = 0
	// kCGBlendModeColorDodge - Brightens the background image samples to reflect the source image samples. Source image sample values that specify black do not produce a change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/colorDodge
	kCGBlendModeColorDodge CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/copy
	kCGBlendModeCopy CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/darken
	kCGBlendModeDarken CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationAtop
	kCGBlendModeDestinationAtop CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationIn
	kCGBlendModeDestinationIn CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationOut
	kCGBlendModeDestinationOut CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/destinationOver
	kCGBlendModeDestinationOver CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/difference
	kCGBlendModeDifference CGBlendMode = 0
	// kCGBlendModeExclusion - Produces an effect similar to that produced by  , but with lower contrast. Source image sample values that are black don’t produce a change; white inverts the background color values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/exclusion
	kCGBlendModeExclusion CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/hardLight
	kCGBlendModeHardLight CGBlendMode = 0
	// kCGBlendModeHue - Uses the luminance and saturation values of the background with the hue of the source image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/hue
	kCGBlendModeHue CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/lighten
	kCGBlendModeLighten CGBlendMode = 0
	// kCGBlendModeLuminosity - Uses the hue and saturation of the background with the luminance of the source image. This mode creates an effect that is inverse to the effect created by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/luminosity
	kCGBlendModeLuminosity CGBlendMode = 0
	// kCGBlendModeMultiply - Multiplies the source image samples with the background image samples. This results in colors that are at least as dark as either of the two contributing sample colors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/multiply
	kCGBlendModeMultiply CGBlendMode = 0
	// kCGBlendModeNormal - Paints the source image samples over the background image samples.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/normal
	kCGBlendModeNormal CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/overlay
	kCGBlendModeOverlay CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/plusDarker
	kCGBlendModePlusDarker CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/plusLighter
	kCGBlendModePlusLighter CGBlendMode = 0
	// kCGBlendModeSaturation - Uses the luminance and hue values of the background with the saturation of the source image. Areas of the background that have no saturation (that is, pure gray areas) don’t produce a change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/saturation
	kCGBlendModeSaturation CGBlendMode = 0
	// kCGBlendModeScreen - Multiplies the inverse of the source image samples with the inverse of the background image samples, resulting in colors that are at least as light as either of the two contributing sample colors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/screen
	kCGBlendModeScreen CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/softLight
	kCGBlendModeSoftLight CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceAtop
	kCGBlendModeSourceAtop CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceIn
	kCGBlendModeSourceIn CGBlendMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/sourceOut
	kCGBlendModeSourceOut CGBlendMode = 0
	// kCGBlendModeXOR - . This XOR mode is only nominally related to the classical bitmap XOR operation, which is not supported by Core Graphics
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGBlendMode/xor
	kCGBlendModeXOR CGBlendMode = 0
)

// CGCaptureOptions - Configuration parameters that are used when capturing displays.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions
type CGCaptureOptions uint

const (
	// kCGCaptureNoOptions - The system should use the default fill behavior, which is fill with black.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions/kCGCaptureNoOptions
	kCGCaptureNoOptions CGCaptureOptions = 0
	// kCGCaptureNoFill - Disables fill with black.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGCaptureOptions/noFill
	kCGCaptureNoFill CGCaptureOptions = 0
)

// CGColorConversionInfoTransformType - Constants describing how a color conversion uses color spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType
type CGColorConversionInfoTransformType uint

const (
	// kCGColorConversionTransformApplySpace - Specifies a color conversion between one color profile and another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformApplySpace
	kCGColorConversionTransformApplySpace CGColorConversionInfoTransformType = 0
	// kCGColorConversionTransformFromSpace - Specifies a color conversion from a device color space to a color profile.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformFromSpace
	kCGColorConversionTransformFromSpace CGColorConversionInfoTransformType = 0
	// kCGColorConversionTransformToSpace - Specifies a color conversion from a color profile to a device color space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorConversionInfoTransformType/transformToSpace
	kCGColorConversionTransformToSpace CGColorConversionInfoTransformType = 0
)

// CGColorRenderingIntent - Handling options for colors that are not located within the destination color space of a graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent
type CGColorRenderingIntent uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/absoluteColorimetric
	kCGRenderingIntentAbsoluteColorimetric CGColorRenderingIntent = 0
	// kCGRenderingIntentDefault - The default rendering intent for the graphics context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/defaultIntent
	kCGRenderingIntentDefault CGColorRenderingIntent = 0
	// kCGRenderingIntentPerceptual - Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device. Perceptual intent is good for photographs and other complex, detailed images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/perceptual
	kCGRenderingIntentPerceptual CGColorRenderingIntent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/relativeColorimetric
	kCGRenderingIntentRelativeColorimetric CGColorRenderingIntent = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorRenderingIntent/saturation
	kCGRenderingIntentSaturation CGColorRenderingIntent = 0
)

// CGColorSpaceModel - Models for color spaces.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel
type CGColorSpaceModel uint

const (
	// kCGColorSpaceModelCMYK - A CMYK color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/cmyk
	kCGColorSpaceModelCMYK CGColorSpaceModel = 0
	// kCGColorSpaceModelDeviceN - A DeviceN color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/deviceN
	kCGColorSpaceModelDeviceN CGColorSpaceModel = 0
	// kCGColorSpaceModelIndexed - An indexed color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/indexed
	kCGColorSpaceModelIndexed CGColorSpaceModel = 0
	// kCGColorSpaceModelLab - A Lab color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/lab
	kCGColorSpaceModelLab CGColorSpaceModel = 0
	// kCGColorSpaceModelMonochrome - A monochrome color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/monochrome
	kCGColorSpaceModelMonochrome CGColorSpaceModel = 0
	// kCGColorSpaceModelPattern - A pattern color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/pattern
	kCGColorSpaceModelPattern CGColorSpaceModel = 0
	// kCGColorSpaceModelRGB - An RGB color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/rgb
	kCGColorSpaceModelRGB CGColorSpaceModel = 0
	// kCGColorSpaceModelUnknown - An unknown color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/unknown
	kCGColorSpaceModelUnknown CGColorSpaceModel = 0
	// kCGColorSpaceModelXYZ - An XYZ color space model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpaceModel/XYZ
	kCGColorSpaceModelXYZ CGColorSpaceModel = 0
)

// CGConfigureOption - The scope of the changes in a display configuration transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption
type CGConfigureOption uint

const (
	// kCGConfigureForAppOnly - Changes persist for the lifetime of the current application. After the application terminates, the display configuration settings revert to the current login session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/forAppOnly
	kCGConfigureForAppOnly CGConfigureOption = 0
	// kCGConfigureForSession - Changes persist for the lifetime of the current login session. After the current session terminates, the displays revert to the last saved permanent configuration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/forSession
	kCGConfigureForSession CGConfigureOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGConfigureOption/permanently
	kCGConfigurePermanently CGConfigureOption = 0
)

// CGDisplayChangeSummaryFlags - The configuration parameters that are passed to a display reconfiguration callback function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags
type CGDisplayChangeSummaryFlags uint

const (
	// kCGDisplayAddFlag - The display has been added to the active display list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/addFlag
	kCGDisplayAddFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayBeginConfigurationFlag - The display configuration is about to change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/beginConfigurationFlag
	kCGDisplayBeginConfigurationFlag CGDisplayChangeSummaryFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/desktopShapeChangedFlag
	kCGDisplayDesktopShapeChangedFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayDisabledFlag - The display has been disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/disabledFlag
	kCGDisplayDisabledFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayEnabledFlag - The display has been enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/enabledFlag
	kCGDisplayEnabledFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayMirrorFlag - The display is now mirroring another display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/mirrorFlag
	kCGDisplayMirrorFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayMovedFlag - The location of the upper-left corner of the display in the global display coordinate space has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/movedFlag
	kCGDisplayMovedFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayRemoveFlag - The display has been removed from the active display list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/removeFlag
	kCGDisplayRemoveFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplaySetMainFlag - The display is now the main display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/setMainFlag
	kCGDisplaySetMainFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplaySetModeFlag - The display mode has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/setModeFlag
	kCGDisplaySetModeFlag CGDisplayChangeSummaryFlags = 0
	// kCGDisplayUnMirrorFlag - The display is no longer mirroring another display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayChangeSummaryFlags/unMirrorFlag
	kCGDisplayUnMirrorFlag CGDisplayChangeSummaryFlags = 0
)

// CGDisplayStreamFrameStatus - Describes a frame update event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus
type CGDisplayStreamFrameStatus uint

const (
	// kCGDisplayStreamFrameStatusFrameBlank - A new frame was not generated because the display has gone blank.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameBlank
	kCGDisplayStreamFrameStatusFrameBlank CGDisplayStreamFrameStatus = 0
	// kCGDisplayStreamFrameStatusFrameComplete - A new frame was generated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameComplete
	kCGDisplayStreamFrameStatusFrameComplete CGDisplayStreamFrameStatus = 0
	// kCGDisplayStreamFrameStatusFrameIdle - A new frame was not generated because the display did not change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/frameIdle
	kCGDisplayStreamFrameStatusFrameIdle CGDisplayStreamFrameStatus = 0
	// kCGDisplayStreamFrameStatusStopped - The display stream was stopped.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamFrameStatus/stopped
	kCGDisplayStreamFrameStatusStopped CGDisplayStreamFrameStatus = 0
)

// CGDisplayStreamUpdateRectType - Use these constants to determine which rectangles your app is interested in.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType
type CGDisplayStreamUpdateRectType uint

const (
	// kCGDisplayStreamUpdateDirtyRects - The union of both rectangles that were redrawn and rectangles that were moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/dirtyRects
	kCGDisplayStreamUpdateDirtyRects CGDisplayStreamUpdateRectType = 0
	// kCGDisplayStreamUpdateMovedRects - The rectangles for the portions of the display that were simply moved from one part of the display to another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/movedRects
	kCGDisplayStreamUpdateMovedRects CGDisplayStreamUpdateRectType = 0
	// kCGDisplayStreamUpdateReducedDirtyRects - The union is calculated and then simplified. This reduces the number of rectangles returned to your app, but it may report some pixels that were not actually changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/reducedDirtyRects
	kCGDisplayStreamUpdateReducedDirtyRects CGDisplayStreamUpdateRectType = 0
	// kCGDisplayStreamUpdateRefreshedRects - The rectangles for the portions of the display that were redrawn.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGDisplayStreamUpdateRectType/refreshedRects
	kCGDisplayStreamUpdateRefreshedRects CGDisplayStreamUpdateRectType = 0
)

// CGError - A uniform type for result codes returned by functions in Core Graphics.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError
type CGError uint

const (
	// kCGErrorCannotComplete - The requested operation is inappropriate for the parameters passed in, or the current system state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/cannotComplete
	kCGErrorCannotComplete CGError = 0
	// kCGErrorFailure - A general failure occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/failure
	kCGErrorFailure CGError = 0
	// kCGErrorIllegalArgument - One or more of the parameters passed to a function are invalid. Check for   pointers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/illegalArgument
	kCGErrorIllegalArgument CGError = 0
	// kCGErrorInvalidConnection - The parameter representing a connection to the window server is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidConnection
	kCGErrorInvalidConnection CGError = 0
	// kCGErrorInvalidContext - The   or context identifier parameter is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidContext
	kCGErrorInvalidContext CGError = 0
	// kCGErrorInvalidOperation - The requested operation is not valid for the parameters passed in, or the current system state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/invalidOperation
	kCGErrorInvalidOperation CGError = 0
	// kCGErrorNoneAvailable - The requested operation could not be completed as the indicated resources were not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/noneAvailable
	kCGErrorNoneAvailable CGError = 0
	// kCGErrorNotImplemented - Return value from obsolete function stubs present for binary compatibility, but not typically called.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/notImplemented
	kCGErrorNotImplemented CGError = 0
	// kCGErrorRangeCheck - A parameter passed in has a value that is inappropriate, or which does not map to a useful operation or value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/rangeCheck
	kCGErrorRangeCheck CGError = 0
	// kCGErrorSuccess - The requested operation was completed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/success
	kCGErrorSuccess CGError = 0
	// kCGErrorTypeCheck - A data type or token was encountered that did not match the expected type or token.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGError/typeCheck
	kCGErrorTypeCheck CGError = 0
)

// CGEventField - Constants used as keys to access specialized fields in low-level events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField
type CGEventField uint

const (
	// kCGEventSourceGroupID - Key to access a field that contains the event source Unix effective GID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceGroupID
	kCGEventSourceGroupID CGEventField = 0
	// kCGEventSourceStateID - Key to access a field that contains the event source state ID used to create this event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceStateID
	kCGEventSourceStateID CGEventField = 0
	// kCGEventSourceUnixProcessID - Key to access a field that contains the event source Unix process ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUnixProcessID
	kCGEventSourceUnixProcessID CGEventField = 0
	// kCGEventSourceUserData - Key to access a field that contains the event source user-supplied data, up to 64 bits.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUserData
	kCGEventSourceUserData CGEventField = 0
	// kCGEventSourceUserID - Key to access a field that contains the event source Unix effective UID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventSourceUserID
	kCGEventSourceUserID CGEventField = 0
	// kCGEventTargetProcessSerialNumber - Key to access a field that contains the event target process serial number. The value is a 64-bit long word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventTargetProcessSerialNumber
	kCGEventTargetProcessSerialNumber CGEventField = 0
	// kCGEventTargetUnixProcessID - Key to access a field that contains the event target Unix process ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventTargetUnixProcessID
	kCGEventTargetUnixProcessID CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventUnacceleratedPointerMovementX
	kCGEventUnacceleratedPointerMovementX CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/eventUnacceleratedPointerMovementY
	kCGEventUnacceleratedPointerMovementY CGEventField = 0
	// kCGKeyboardEventAutorepeat - Key to access an integer field, non-zero when this is an autorepeat of a key-down, and zero otherwise.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventAutorepeat
	kCGKeyboardEventAutorepeat CGEventField = 0
	// kCGKeyboardEventKeyboardType - Key to access an integer field that contains the keyboard type identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventKeyboardType
	kCGKeyboardEventKeyboardType CGEventField = 0
	// kCGKeyboardEventKeycode - Key to access an integer field that contains the virtual keycode of the key-down or key-up event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/keyboardEventKeycode
	kCGKeyboardEventKeycode CGEventField = 0
	// kCGMouseEventButtonNumber - Key to access an integer field that contains the mouse button number. For information about the possible values, see  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventButtonNumber
	kCGMouseEventButtonNumber CGEventField = 0
	// kCGMouseEventClickState - Key to access an integer field that contains the mouse button click state. A click state of 1 represents a single click. A click state of 2 represents a double-click. A click state of 3 represents a triple-click.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventClickState
	kCGMouseEventClickState CGEventField = 0
	// kCGMouseEventDeltaX - Key to access an integer field that contains the horizontal mouse delta since the last mouse movement event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventDeltaX
	kCGMouseEventDeltaX CGEventField = 0
	// kCGMouseEventDeltaY - Key to access an integer field that contains the vertical mouse delta since the last mouse movement event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventDeltaY
	kCGMouseEventDeltaY CGEventField = 0
	// kCGMouseEventInstantMouser - Key to access an integer field. The value is non-zero if the event should be ignored by the Inkwell subsystem.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventInstantMouser
	kCGMouseEventInstantMouser CGEventField = 0
	// kCGMouseEventNumber - Key to access an integer field that contains the mouse button event number. Matching mouse-down and mouse-up events will have the same event number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventNumber
	kCGMouseEventNumber CGEventField = 0
	// kCGMouseEventPressure - Key to access a double field that contains the mouse button pressure. The pressure value may range from 0 to 1, with 0 representing the mouse being up. This value is commonly set by tablet pens mimicking a mouse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventPressure
	kCGMouseEventPressure CGEventField = 0
	// kCGMouseEventSubtype - Key to access an integer field that encodes the mouse event subtype as a  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventSubtype
	kCGMouseEventSubtype CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventWindowUnderMousePointer
	kCGMouseEventWindowUnderMousePointer CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/mouseEventWindowUnderMousePointerThatCanHandleThisEvent
	kCGMouseEventWindowUnderMousePointerThatCanHandleThisEvent CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventAcceleratedDeltaAxis1
	kCGScrollWheelEventAcceleratedDeltaAxis1 CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventAcceleratedDeltaAxis2
	kCGScrollWheelEventAcceleratedDeltaAxis2 CGEventField = 0
	// kCGScrollWheelEventDeltaAxis1 - Key to access an integer field that contains scrolling data. This field typically contains the change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis1
	kCGScrollWheelEventDeltaAxis1 CGEventField = 0
	// kCGScrollWheelEventDeltaAxis2 - Key to access an integer field that contains scrolling data. This field typically contains the change in horizontal position since the last scrolling event from a Mighty Mouse scroller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis2
	kCGScrollWheelEventDeltaAxis2 CGEventField = 0
	// kCGScrollWheelEventDeltaAxis3 - This field is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventDeltaAxis3
	kCGScrollWheelEventDeltaAxis3 CGEventField = 0
	// kCGScrollWheelEventFixedPtDeltaAxis1 - Key to access a field that contains scrolling data. The scrolling data represents a line-based or pixel-based change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller. The scrolling data uses a fixed-point 16.16 signed integer format. For example, if the field contains a value of 1.0, the integer 0x00010000 is returned by  . If this key is passed to  , the fixed-point value is converted to a double value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis1
	kCGScrollWheelEventFixedPtDeltaAxis1 CGEventField = 0
	// kCGScrollWheelEventFixedPtDeltaAxis2 - Key to access a field that contains scrolling data. The scrolling data represents a line-based or pixel-based change in horizontal position since the last scrolling event from a Mighty Mouse scroller. The scrolling data uses a fixed-point 16.16 signed integer format. For example, if the field contains a value of 1.0, the integer 0x00010000 is returned by  . If this key is passed to  , the fixed-point value is converted to a double value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis2
	kCGScrollWheelEventFixedPtDeltaAxis2 CGEventField = 0
	// kCGScrollWheelEventFixedPtDeltaAxis3 - This field is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventFixedPtDeltaAxis3
	kCGScrollWheelEventFixedPtDeltaAxis3 CGEventField = 0
	// kCGScrollWheelEventInstantMouser - Key to access an integer field that indicates whether the event should be ignored by the Inkwell subsystem. If the value is non-zero, the event should be ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventInstantMouser
	kCGScrollWheelEventInstantMouser CGEventField = 0
	// kCGScrollWheelEventIsContinuous - Key to access an integer field that indicates whether a scrolling event contains continuous, pixel-based scrolling data. The value is non-zero when the scrolling data is pixel-based and zero when the scrolling data is line-based.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventIsContinuous
	kCGScrollWheelEventIsContinuous CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventMomentumOptionPhase
	kCGScrollWheelEventMomentumOptionPhase CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventMomentumPhase
	kCGScrollWheelEventMomentumPhase CGEventField = 0
	// kCGScrollWheelEventPointDeltaAxis1 - Key to access an integer field that contains pixel-based scrolling data. The scrolling data represents the change in vertical position since the last scrolling event from a Mighty Mouse scroller or a single-wheel mouse scroller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis1
	kCGScrollWheelEventPointDeltaAxis1 CGEventField = 0
	// kCGScrollWheelEventPointDeltaAxis2 - Key to access an integer field that contains pixel-based scrolling data. The scrolling data represents the change in horizontal position since the last scrolling event from a Mighty Mouse scroller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis2
	kCGScrollWheelEventPointDeltaAxis2 CGEventField = 0
	// kCGScrollWheelEventPointDeltaAxis3 - This field is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventPointDeltaAxis3
	kCGScrollWheelEventPointDeltaAxis3 CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventRawDeltaAxis1
	kCGScrollWheelEventRawDeltaAxis1 CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventRawDeltaAxis2
	kCGScrollWheelEventRawDeltaAxis2 CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventScrollCount
	kCGScrollWheelEventScrollCount CGEventField = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/scrollWheelEventScrollPhase
	kCGScrollWheelEventScrollPhase CGEventField = 0
	// kCGTabletEventDeviceID - Key to access an integer field that contains the system-assigned unique device ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventDeviceID
	kCGTabletEventDeviceID CGEventField = 0
	// kCGTabletEventPointButtons - Key to access an integer field that contains the tablet button state. Bit 0 is the first button, and a set bit represents a closed or pressed button. Up to 16 buttons are supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointButtons
	kCGTabletEventPointButtons CGEventField = 0
	// kCGTabletEventPointPressure - Key to access a double field that contains the tablet pen pressure. A value of 0.0 represents no pressure, and 1.0 represents maximum pressure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointPressure
	kCGTabletEventPointPressure CGEventField = 0
	// kCGTabletEventPointX - Key to access an integer field that contains the absolute X coordinate in tablet space at full tablet resolution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointX
	kCGTabletEventPointX CGEventField = 0
	// kCGTabletEventPointY - Key to access an integer field that contains the absolute Y coordinate in tablet space at full tablet resolution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointY
	kCGTabletEventPointY CGEventField = 0
	// kCGTabletEventPointZ - Key to access an integer field that contains the absolute Z coordinate in tablet space at full tablet resolution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventPointZ
	kCGTabletEventPointZ CGEventField = 0
	// kCGTabletEventRotation - Key to access a double field that contains the tablet pen rotation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventRotation
	kCGTabletEventRotation CGEventField = 0
	// kCGTabletEventTangentialPressure - Key to access a double field that contains the tangential pressure on the device. A value of 0.0 represents no pressure, and 1.0 represents maximum pressure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTangentialPressure
	kCGTabletEventTangentialPressure CGEventField = 0
	// kCGTabletEventTiltX - Key to access a double field that contains the horizontal tablet pen tilt. A value of 0.0 represents no tilt, and 1.0 represents maximum tilt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTiltX
	kCGTabletEventTiltX CGEventField = 0
	// kCGTabletEventTiltY - Key to access a double field that contains the vertical tablet pen tilt. A value of 0.0 represents no tilt, and 1.0 represents maximum tilt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventTiltY
	kCGTabletEventTiltY CGEventField = 0
	// kCGTabletEventVendor1 - Key to access an integer field that contains a vendor-specified value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor1
	kCGTabletEventVendor1 CGEventField = 0
	// kCGTabletEventVendor2 - Key to access an integer field that contains a vendor-specified value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor2
	kCGTabletEventVendor2 CGEventField = 0
	// kCGTabletEventVendor3 - Key to access an integer field that contains a vendor-specified value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletEventVendor3
	kCGTabletEventVendor3 CGEventField = 0
	// kCGTabletProximityEventCapabilityMask - Key to access an integer field that contains the device capabilities mask.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventCapabilityMask
	kCGTabletProximityEventCapabilityMask CGEventField = 0
	// kCGTabletProximityEventDeviceID - Key to access an integer field that contains the system-assigned device ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventDeviceID
	kCGTabletProximityEventDeviceID CGEventField = 0
	// kCGTabletProximityEventEnterProximity - Key to access an integer field that indicates whether the pen is in proximity to the tablet. The value is non-zero if the pen is in proximity to the tablet and zero when leaving the tablet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventEnterProximity
	kCGTabletProximityEventEnterProximity CGEventField = 0
	// kCGTabletProximityEventPointerID - Key to access an integer field that contains the vendor-defined ID of the pointing device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventPointerID
	kCGTabletProximityEventPointerID CGEventField = 0
	// kCGTabletProximityEventPointerType - Key to access an integer field that contains the pointer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventPointerType
	kCGTabletProximityEventPointerType CGEventField = 0
	// kCGTabletProximityEventSystemTabletID - Key to access an integer field that contains the system-assigned unique tablet ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventSystemTabletID
	kCGTabletProximityEventSystemTabletID CGEventField = 0
	// kCGTabletProximityEventTabletID - Key to access an integer field that contains the vendor-defined tablet ID, typically the USB product ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventTabletID
	kCGTabletProximityEventTabletID CGEventField = 0
	// kCGTabletProximityEventVendorID - Key to access an integer field that contains the vendor-defined ID, typically the USB vendor ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorID
	kCGTabletProximityEventVendorID CGEventField = 0
	// kCGTabletProximityEventVendorPointerSerialNumber - Key to access an integer field that contains the vendor-defined pointer serial number.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorPointerSerialNumber
	kCGTabletProximityEventVendorPointerSerialNumber CGEventField = 0
	// kCGTabletProximityEventVendorPointerType - Key to access an integer field that contains the vendor-assigned pointer type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorPointerType
	kCGTabletProximityEventVendorPointerType CGEventField = 0
	// kCGTabletProximityEventVendorUniqueID - Key to access an integer field that contains the vendor-defined unique ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventField/tabletProximityEventVendorUniqueID
	kCGTabletProximityEventVendorUniqueID CGEventField = 0
)

// CGEventFilterMask - Specify masks for classes of low-level events that can be filtered during event suppression states.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask
type CGEventFilterMask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitLocalKeyboardEvents
	kCGEventFilterMaskPermitLocalKeyboardEvents CGEventFilterMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitLocalMouseEvents
	kCGEventFilterMaskPermitLocalMouseEvents CGEventFilterMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFilterMask/permitSystemDefinedEvents
	kCGEventFilterMaskPermitSystemDefinedEvents CGEventFilterMask = 0
)

// CGEventFlags - Constants that indicate the modifier key state at the time an event is created, as well as other event-related states.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags
type CGEventFlags uint

const (
	// kCGEventFlagMaskAlphaShift - Indicates that the Caps Lock key is down for a keyboard, mouse, or flag-changed event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskAlphaShift
	kCGEventFlagMaskAlphaShift CGEventFlags = 0
	// kCGEventFlagMaskAlternate - Indicates that the Alt or Option key is down for a keyboard, mouse, or flag-changed event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskAlternate
	kCGEventFlagMaskAlternate CGEventFlags = 0
	// kCGEventFlagMaskCommand - Indicates that the Command key is down for a keyboard, mouse, or flag-changed event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskCommand
	kCGEventFlagMaskCommand CGEventFlags = 0
	// kCGEventFlagMaskControl - Indicates that the Control key is down for a keyboard, mouse, or flag-changed event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskControl
	kCGEventFlagMaskControl CGEventFlags = 0
	// kCGEventFlagMaskHelp - Indicates that the Help modifier key is down for a keyboard, mouse, or flag-changed event. This key is not present on most keyboards, and is different than the Help key found in the same row as Home and Page Up.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskHelp
	kCGEventFlagMaskHelp CGEventFlags = 0
	// kCGEventFlagMaskNonCoalesced - Indicates that mouse and pen movement events are not being coalesced.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskNonCoalesced
	kCGEventFlagMaskNonCoalesced CGEventFlags = 0
	// kCGEventFlagMaskNumericPad - Identifies key events from the numeric keypad area on extended keyboards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskNumericPad
	kCGEventFlagMaskNumericPad CGEventFlags = 0
	// kCGEventFlagMaskSecondaryFn - Indicates that the Fn (Function) key is down for a keyboard, mouse, or flag-changed event. This key is found primarily on laptop keyboards.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskSecondaryFn
	kCGEventFlagMaskSecondaryFn CGEventFlags = 0
	// kCGEventFlagMaskShift - Indicates that the Shift key is down for a keyboard, mouse, or flag-changed event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventFlags/maskShift
	kCGEventFlagMaskShift CGEventFlags = 0
)

// CGEventMouseSubtype - Constants used with the 
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype
type CGEventMouseSubtype uint

const (
	// kCGEventMouseSubtypeDefault - Specifies that the event is an ordinary mouse event, and does not contain additional tablet device information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/defaultType
	kCGEventMouseSubtypeDefault CGEventMouseSubtype = 0
	// kCGEventMouseSubtypeTabletPoint - Specifies that the mouse event originated from a tablet device, and that the various   field selectors may be used to obtain tablet-specific data from the mouse event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/tabletPoint
	kCGEventMouseSubtypeTabletPoint CGEventMouseSubtype = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventMouseSubtype/tabletProximity
	kCGEventMouseSubtypeTabletProximity CGEventMouseSubtype = 0
)

// CGEventSourceStateID - Constants that specify the possible source states of an event source.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID
type CGEventSourceStateID uint

const (
	// kCGEventSourceStateCombinedSessionState - Specifies that an event source should use the event state table that reflects the combined state of all event sources posting to the current user login session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/combinedSessionState
	kCGEventSourceStateCombinedSessionState CGEventSourceStateID = 0
	// kCGEventSourceStateHIDSystemState - Specifies that an event source should use the event state table that reflects the combined state of all hardware event sources posting from the HID system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/hidSystemState
	kCGEventSourceStateHIDSystemState CGEventSourceStateID = 0
	// kCGEventSourceStatePrivate - Specifies that an event source should use a private event state table.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSourceStateID/privateState
	kCGEventSourceStatePrivate CGEventSourceStateID = 0
)

// CGEventSuppressionState - Specify the event suppression states that can occur after posting an event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState
type CGEventSuppressionState uint

const (
	// kCGEventSuppressionStateRemoteMouseDrag - Specifies that certain local hardware events may be suppressed during a mouse drag operation (mouse movement with the left or only mouse button down).
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/eventSuppressionStateRemoteMouseDrag
	kCGEventSuppressionStateRemoteMouseDrag CGEventSuppressionState = 0
	// kCGEventSuppressionStateSuppressionInterval - Specifies that certain local hardware events may be suppressed for a short interval after posting an event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/eventSuppressionStateSuppressionInterval
	kCGEventSuppressionStateSuppressionInterval CGEventSuppressionState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventSuppressionState/numberOfEventSuppressionStates
	kCGNumberOfEventSuppressionStates CGEventSuppressionState = 0
)

// CGEventTapLocation - Constants that specify possible tapping points for events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation
type CGEventTapLocation uint

const (
	// kCGAnnotatedSessionEventTap - Specifies that an event tap is placed at the point where session events have been annotated to flow to an application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cgAnnotatedSessionEventTap
	kCGAnnotatedSessionEventTap CGEventTapLocation = 0
	// kCGHIDEventTap - Specifies that an event tap is placed at the point where HID system events enter the window server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cghidEventTap
	kCGHIDEventTap CGEventTapLocation = 0
	// kCGSessionEventTap - Specifies that an event tap is placed at the point where HID system and remote control events enter a login session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapLocation/cgSessionEventTap
	kCGSessionEventTap CGEventTapLocation = 0
)

// CGEventTapOptions - Constants that specify whether a new event tap is an active filter or a passive listener.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions
type CGEventTapOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions/defaultTap
	kCGEventTapOptionDefault CGEventTapOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapOptions/listenOnly
	kCGEventTapOptionListenOnly CGEventTapOptions = 0
)

// CGEventTapPlacement - Constants that specify where a new event tap is inserted into the list of active event taps.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement
type CGEventTapPlacement uint

const (
	// kCGHeadInsertEventTap - Specifies that a new event tap should be inserted before any pre-existing event taps at the same location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement/headInsertEventTap
	kCGHeadInsertEventTap CGEventTapPlacement = 0
	// kCGTailAppendEventTap - Specifies that a new event tap should be inserted after any pre-existing event taps at the same location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventTapPlacement/tailAppendEventTap
	kCGTailAppendEventTap CGEventTapPlacement = 0
)

// CGEventType - Constants that specify the different types of input events.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType
type CGEventType uint

const (
	// kCGEventFlagsChanged - Specifies a key changed event for a modifier or status key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/flagsChanged
	kCGEventFlagsChanged CGEventType = 0
	// kCGEventKeyDown - Specifies a key down event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/keyDown
	kCGEventKeyDown CGEventType = 0
	// kCGEventKeyUp - Specifies a key up event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/keyUp
	kCGEventKeyUp CGEventType = 0
	// kCGEventLeftMouseDown - Specifies a mouse down event with the left button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseDown
	kCGEventLeftMouseDown CGEventType = 0
	// kCGEventLeftMouseDragged - Specifies a mouse drag event with the left button down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseDragged
	kCGEventLeftMouseDragged CGEventType = 0
	// kCGEventLeftMouseUp - Specifies a mouse up event with the left button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/leftMouseUp
	kCGEventLeftMouseUp CGEventType = 0
	// kCGEventMouseMoved - Specifies a mouse moved event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/mouseMoved
	kCGEventMouseMoved CGEventType = 0
	// kCGEventNull - Specifies a null event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/null
	kCGEventNull CGEventType = 0
	// kCGEventOtherMouseDown - Specifies a mouse down event with one of buttons 2-31.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseDown
	kCGEventOtherMouseDown CGEventType = 0
	// kCGEventOtherMouseDragged - Specifies a mouse drag event with one of buttons 2-31 down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseDragged
	kCGEventOtherMouseDragged CGEventType = 0
	// kCGEventOtherMouseUp - Specifies a mouse up event with one of buttons 2-31.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/otherMouseUp
	kCGEventOtherMouseUp CGEventType = 0
	// kCGEventRightMouseDown - Specifies a mouse down event with the right button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseDown
	kCGEventRightMouseDown CGEventType = 0
	// kCGEventRightMouseDragged - Specifies a mouse drag event with the right button down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseDragged
	kCGEventRightMouseDragged CGEventType = 0
	// kCGEventRightMouseUp - Specifies a mouse up event with the right button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/rightMouseUp
	kCGEventRightMouseUp CGEventType = 0
	// kCGEventScrollWheel - Specifies a scroll wheel moved event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/scrollWheel
	kCGEventScrollWheel CGEventType = 0
	// kCGEventTabletPointer - Specifies a tablet pointer event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tabletPointer
	kCGEventTabletPointer CGEventType = 0
	// kCGEventTabletProximity - Specifies a tablet proximity event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tabletProximity
	kCGEventTabletProximity CGEventType = 0
	// kCGEventTapDisabledByTimeout - Specifies an event indicating the event tap is disabled because of timeout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tapDisabledByTimeout
	kCGEventTapDisabledByTimeout CGEventType = 0
	// kCGEventTapDisabledByUserInput - Specifies an event indicating the event tap is disabled because of user input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGEventType/tapDisabledByUserInput
	kCGEventTapDisabledByUserInput CGEventType = 0
)

// CGFontPostScriptFormat - Possible formats for a PostScript font subset.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat
type CGFontPostScriptFormat uint

const (
	// kCGFontPostScriptFormatType1 - A Type 1 font format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type1
	kCGFontPostScriptFormatType1 CGFontPostScriptFormat = 0
	// kCGFontPostScriptFormatType3 - A Type 3 PostScript format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type3
	kCGFontPostScriptFormatType3 CGFontPostScriptFormat = 0
	// kCGFontPostScriptFormatType42 - A constant representing a Type 42 font format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGFontPostScriptFormat/type42
	kCGFontPostScriptFormatType42 CGFontPostScriptFormat = 0
)

// CGGesturePhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase
type CGGesturePhase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/began
	kCGGesturePhaseBegan CGGesturePhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/cancelled
	kCGGesturePhaseCancelled CGGesturePhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/changed
	kCGGesturePhaseChanged CGGesturePhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/ended
	kCGGesturePhaseEnded CGGesturePhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/mayBegin
	kCGGesturePhaseMayBegin CGGesturePhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGesturePhase/none
	kCGGesturePhaseNone CGGesturePhase = 0
)

// CGGlyphDeprecatedEnum enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum
type CGGlyphDeprecatedEnum uint

const (
	// CGGlyphMax - Maximum font index value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum/max
	CGGlyphMax CGGlyphDeprecatedEnum = 0
	// CGGlyphMin - Minimum font index value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGlyphDeprecatedEnum/min
	CGGlyphMin CGGlyphDeprecatedEnum = 0
)

// CGGradientDrawingOptions - Drawing locations for gradients.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions
type CGGradientDrawingOptions uint

const (
	// kCGGradientDrawsAfterEndLocation - The fill should extend beyond the ending location. The color that extends beyond the ending point is the solid color defined by the   object to be at location 1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions/drawsAfterEndLocation
	kCGGradientDrawsAfterEndLocation CGGradientDrawingOptions = 0
	// kCGGradientDrawsBeforeStartLocation - The fill should extend beyond the starting location. The color that extends beyond the starting point is the solid color defined by the   object to be at location 0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGGradientDrawingOptions/drawsBeforeStartLocation
	kCGGradientDrawsBeforeStartLocation CGGradientDrawingOptions = 0
)

// CGImageAlphaInfo - Storage options for alpha component data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo
type CGImageAlphaInfo uint

const (
	// kCGImageAlphaOnly - There is no color data, only an alpha channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/alphaOnly
	kCGImageAlphaOnly CGImageAlphaInfo = 0
	// kCGImageAlphaFirst - The alpha component is stored in the most significant bits of each pixel. For example, non-premultiplied ARGB.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/first
	kCGImageAlphaFirst CGImageAlphaInfo = 0
	// kCGImageAlphaLast - The alpha component is stored in the least significant bits of each pixel. For example, non-premultiplied RGBA.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/last
	kCGImageAlphaLast CGImageAlphaInfo = 0
	// kCGImageAlphaNone - There is no alpha channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/none
	kCGImageAlphaNone CGImageAlphaInfo = 0
	// kCGImageAlphaNoneSkipFirst - There is no alpha channel. If the total size of the pixel is greater than the space required for the number of color components in the color space, the most significant bits are ignored.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/noneSkipFirst
	kCGImageAlphaNoneSkipFirst CGImageAlphaInfo = 0
	// kCGImageAlphaNoneSkipLast - There is no alpha channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/noneSkipLast
	kCGImageAlphaNoneSkipLast CGImageAlphaInfo = 0
	// kCGImageAlphaPremultipliedFirst - The alpha component is stored in the most significant bits of each pixel and the color components have already been multiplied by this alpha value. For example, premultiplied ARGB.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/premultipliedFirst
	kCGImageAlphaPremultipliedFirst CGImageAlphaInfo = 0
	// kCGImageAlphaPremultipliedLast - The alpha component is stored in the least significant bits of each pixel and the color components have already been multiplied by this alpha value. For example, premultiplied RGBA.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageAlphaInfo/premultipliedLast
	kCGImageAlphaPremultipliedLast CGImageAlphaInfo = 0
)

// CGImageByteOrderInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo
type CGImageByteOrderInfo uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Big
	kCGImageByteOrder16Big CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Host
	kCGImageByteOrder16Host CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order16Little
	kCGImageByteOrder16Little CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Big
	kCGImageByteOrder32Big CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Host
	kCGImageByteOrder32Host CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/order32Little
	kCGImageByteOrder32Little CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/orderDefault
	kCGImageByteOrderDefault CGImageByteOrderInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImageByteOrderInfo/orderMask
	kCGImageByteOrderMask CGImageByteOrderInfo = 0
)

// CGImagePixelFormatInfo enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo
type CGImagePixelFormatInfo uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/mask
	kCGImagePixelFormatMask CGImagePixelFormatInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/packed
	kCGImagePixelFormatPacked CGImagePixelFormatInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB101010
	kCGImagePixelFormatRGB101010 CGImagePixelFormatInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB555
	kCGImagePixelFormatRGB555 CGImagePixelFormatInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGB565
	kCGImagePixelFormatRGB565 CGImagePixelFormatInfo = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImagePixelFormatInfo/RGBCIF10
	kCGImagePixelFormatRGBCIF10 CGImagePixelFormatInfo = 0
)

// CGInterpolationQuality - Levels of interpolation quality for rendering an image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality
type CGInterpolationQuality uint

const (
	// kCGInterpolationDefault - The default level of quality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/default
	kCGInterpolationDefault CGInterpolationQuality = 0
	// kCGInterpolationHigh - A high level of interpolation quality. This setting may slow down image rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/high
	kCGInterpolationHigh CGInterpolationQuality = 0
	// kCGInterpolationLow - A low level of interpolation quality. This setting may speed up image rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/low
	kCGInterpolationLow CGInterpolationQuality = 0
	// kCGInterpolationMedium - A medium level of interpolation quality. This setting is slower than the low setting but faster than the high setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/medium
	kCGInterpolationMedium CGInterpolationQuality = 0
	// kCGInterpolationNone - No interpolation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGInterpolationQuality/none
	kCGInterpolationNone CGInterpolationQuality = 0
)

// CGLineCap - Styles for rendering the endpoint of a stroked line.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap
type CGLineCap uint

const (
	// kCGLineCapButt - A line with a squared-off end. Core Graphics draws the line to extend only to the exact endpoint of the path. This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/butt
	kCGLineCapButt CGLineCap = 0
	// kCGLineCapRound - A line with a rounded end. Core Graphics draws the line to extend beyond the endpoint of the path. The line ends with a semicircular arc with a radius of 1/2 the line’s width, centered on the endpoint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/round
	kCGLineCapRound CGLineCap = 0
	// kCGLineCapSquare - A line with a squared-off end. Core Graphics extends the line beyond the endpoint of the path for a distance equal to half the line width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/square
	kCGLineCapSquare CGLineCap = 0
)

// CGLineJoin - Junction types for stroked lines.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin
type CGLineJoin uint

const (
	// kCGLineJoinBevel - A join with a squared-off end. Core Graphics draws the line to extend beyond the endpoint of the path, for a distance of 1/2 the line’s width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/bevel
	kCGLineJoinBevel CGLineJoin = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/miter
	kCGLineJoinMiter CGLineJoin = 0
	// kCGLineJoinRound - A join with a rounded end. Core Graphics draws the line to extend beyond the endpoint of the path. The line ends with a semicircular arc with a radius of 1/2 the line’s width, centered on the endpoint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/round
	kCGLineJoinRound CGLineJoin = 0
)

// CGMomentumScrollPhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase
type CGMomentumScrollPhase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/begin
	kCGMomentumScrollPhaseBegin CGMomentumScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/continuous
	kCGMomentumScrollPhaseContinue CGMomentumScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/end
	kCGMomentumScrollPhaseEnd CGMomentumScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMomentumScrollPhase/none
	kCGMomentumScrollPhaseNone CGMomentumScrollPhase = 0
)

// CGMouseButton - Constants that specify buttons on a one, two, or three-button mouse.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton
type CGMouseButton uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/center
	kCGMouseButtonCenter CGMouseButton = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/left
	kCGMouseButtonLeft CGMouseButton = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGMouseButton/right
	kCGMouseButtonRight CGMouseButton = 0
)

// CGPathDrawingMode - Options for rendering a path.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode
type CGPathDrawingMode uint

const (
	// kCGPathEOFill - Render the area within the path using the even-odd rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/eoFill
	kCGPathEOFill CGPathDrawingMode = 0
	// kCGPathEOFillStroke - First fill and then stroke the path, using the even-odd rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/eoFillStroke
	kCGPathEOFillStroke CGPathDrawingMode = 0
	// kCGPathFill - Render the area contained within the path using the non-zero winding number rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/fill
	kCGPathFill CGPathDrawingMode = 0
	// kCGPathFillStroke - First fill and then stroke the path, using the nonzero winding number rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/fillStroke
	kCGPathFillStroke CGPathDrawingMode = 0
	// kCGPathStroke - Render a line along the path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathDrawingMode/stroke
	kCGPathStroke CGPathDrawingMode = 0
)

// CGPathElementType - The type of element found in a path.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType
type CGPathElementType uint

const (
	// kCGPathElementAddCurveToPoint - The path element that adds a cubic curve from the current point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addCurveToPoint
	kCGPathElementAddCurveToPoint CGPathElementType = 0
	// kCGPathElementAddLineToPoint - The path element that adds a line from the current point to a new point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addLineToPoint
	kCGPathElementAddLineToPoint CGPathElementType = 0
	// kCGPathElementAddQuadCurveToPoint - The path element that adds a quadratic curve from the current point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/addQuadCurveToPoint
	kCGPathElementAddQuadCurveToPoint CGPathElementType = 0
	// kCGPathElementCloseSubpath - The path element that closes and completes a subpath. The element does not contain any points. See the function  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/closeSubpath
	kCGPathElementCloseSubpath CGPathElementType = 0
	// kCGPathElementMoveToPoint - The path element that starts a new subpath.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPathElementType/moveToPoint
	kCGPathElementMoveToPoint CGPathElementType = 0
)

// CGPatternTiling - Different methods for rendering a tiled pattern.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling
type CGPatternTiling uint

const (
	// kCGPatternTilingConstantSpacing - Pattern cells are spaced consistently, as with  .The pattern cell may be distorted additionally to permit a moreefficient implementation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/constantSpacing
	kCGPatternTilingConstantSpacing CGPatternTiling = 0
	// kCGPatternTilingConstantSpacingMinimalDistortion - Pattern cells are spaced consistently. Thepattern cell may be distorted by as much as 1 device pixel whenthe pattern is painted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/constantSpacingMinimalDistortion
	kCGPatternTilingConstantSpacingMinimalDistortion CGPatternTiling = 0
	// kCGPatternTilingNoDistortion - The pattern cell is not distorted when painted.The spacing between pattern cells may vary by as much as 1 devicepixel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPatternTiling/noDistortion
	kCGPatternTilingNoDistortion CGPatternTiling = 0
)

// CGPDFAccessPermissions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions
type CGPDFAccessPermissions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsCommenting
	kCGPDFAllowsCommenting CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsContentAccessibility
	kCGPDFAllowsContentAccessibility CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsContentCopying
	kCGPDFAllowsContentCopying CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsDocumentAssembly
	kCGPDFAllowsDocumentAssembly CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsDocumentChanges
	kCGPDFAllowsDocumentChanges CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsFormFieldEntry
	kCGPDFAllowsFormFieldEntry CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsHighQualityPrinting
	kCGPDFAllowsHighQualityPrinting CGPDFAccessPermissions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFAccessPermissions/allowsLowQualityPrinting
	kCGPDFAllowsLowQualityPrinting CGPDFAccessPermissions = 0
)

// CGPDFBox - Box types for a PDF page.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox
type CGPDFBox uint

const (
	// kCGPDFArtBox - The page art box—a rectangle, expressed in default user space units, defining the extent of the page’s meaningful content (including potential white space) as intended by the page’s creator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/artBox
	kCGPDFArtBox CGPDFBox = 0
	// kCGPDFBleedBox - The page bleed box—a rectangle, expressed in default user space units, that defines the region to which the contents of the page should be clipped when output in a production environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/bleedBox
	kCGPDFBleedBox CGPDFBox = 0
	// kCGPDFCropBox - The page crop box—a rectangle, expressed in default user space units, that defines the visible region of default user space. When the page is displayed or printed, its contents are to be clipped to this rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/cropBox
	kCGPDFCropBox CGPDFBox = 0
	// kCGPDFMediaBox - The page media box—a rectangle, expressed in default user space units, that defines the boundaries of the physical medium on which the page is intended to be displayed or printed
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/mediaBox
	kCGPDFMediaBox CGPDFBox = 0
	// kCGPDFTrimBox - The page trim box—a rectangle, expressed in default user space units, that defines the intended dimensions of the finished page after trimming.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFBox/trimBox
	kCGPDFTrimBox CGPDFBox = 0
)

// CGPDFDataFormat - The encoding format of PDF data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat
type CGPDFDataFormat uint

const (
	// CGPDFDataFormatJPEG2000 - The data stream is encoded in JPEG-2000 format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/JPEG2000
	CGPDFDataFormatJPEG2000 CGPDFDataFormat = 0
	// CGPDFDataFormatJPEGEncoded - The data stream is encoded in JPEG format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/jpegEncoded
	CGPDFDataFormatJPEGEncoded CGPDFDataFormat = 0
	// CGPDFDataFormatRaw - The data stream is not encoded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFDataFormat/raw
	CGPDFDataFormatRaw CGPDFDataFormat = 0
)

// CGPDFObjectType - Types of PDF object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType
type CGPDFObjectType uint

const (
	// kCGPDFObjectTypeArray - Type for a PDF array.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/array
	kCGPDFObjectTypeArray CGPDFObjectType = 0
	// kCGPDFObjectTypeBoolean - The type for a PDF Boolean.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/boolean
	kCGPDFObjectTypeBoolean CGPDFObjectType = 0
	// kCGPDFObjectTypeDictionary - The type for a PDF dictionary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/dictionary
	kCGPDFObjectTypeDictionary CGPDFObjectType = 0
	// kCGPDFObjectTypeInteger - The type for a PDF integer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/integer
	kCGPDFObjectTypeInteger CGPDFObjectType = 0
	// kCGPDFObjectTypeName - Type for a PDF name.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/name
	kCGPDFObjectTypeName CGPDFObjectType = 0
	// kCGPDFObjectTypeNull - The type for a PDF null.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/null
	kCGPDFObjectTypeNull CGPDFObjectType = 0
	// kCGPDFObjectTypeReal - The type for a PDF real.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/real
	kCGPDFObjectTypeReal CGPDFObjectType = 0
	// kCGPDFObjectTypeStream - The type for a PDF stream.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/stream
	kCGPDFObjectTypeStream CGPDFObjectType = 0
	// kCGPDFObjectTypeString - The type for a PDF string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFObjectType/string
	kCGPDFObjectTypeString CGPDFObjectType = 0
)

// CGPDFTagType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType
type CGPDFTagType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/annotation
	CGPDFTagTypeAnnotation CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/art
	CGPDFTagTypeArt CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/bibliography
	CGPDFTagTypeBibliography CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/blockQuote
	CGPDFTagTypeBlockQuote CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/caption
	CGPDFTagTypeCaption CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/code
	CGPDFTagTypeCode CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/div
	CGPDFTagTypeDiv CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/document
	CGPDFTagTypeDocument CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/figure
	CGPDFTagTypeFigure CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/form
	CGPDFTagTypeForm CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/formula
	CGPDFTagTypeFormula CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header
	CGPDFTagTypeHeader CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header1
	CGPDFTagTypeHeader1 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header2
	CGPDFTagTypeHeader2 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header3
	CGPDFTagTypeHeader3 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header4
	CGPDFTagTypeHeader4 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header5
	CGPDFTagTypeHeader5 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/header6
	CGPDFTagTypeHeader6 CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/index
	CGPDFTagTypeIndex CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/label
	CGPDFTagTypeLabel CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/link
	CGPDFTagTypeLink CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/list
	CGPDFTagTypeList CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/listBody
	CGPDFTagTypeListBody CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/listItem
	CGPDFTagTypeListItem CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/nonStructure
	CGPDFTagTypeNonStructure CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/note
	CGPDFTagTypeNote CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/object
	CGPDFTagTypeObject CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/paragraph
	CGPDFTagTypeParagraph CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/part
	CGPDFTagTypePart CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/private
	CGPDFTagTypePrivate CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/quote
	CGPDFTagTypeQuote CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/reference
	CGPDFTagTypeReference CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/ruby
	CGPDFTagTypeRuby CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyAnnotationText
	CGPDFTagTypeRubyAnnotationText CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyBaseText
	CGPDFTagTypeRubyBaseText CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/rubyPunctuation
	CGPDFTagTypeRubyPunctuation CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/section
	CGPDFTagTypeSection CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/span
	CGPDFTagTypeSpan CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/table
	CGPDFTagTypeTable CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableBody
	CGPDFTagTypeTableBody CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableDataCell
	CGPDFTagTypeTableDataCell CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableFooter
	CGPDFTagTypeTableFooter CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableHeader
	CGPDFTagTypeTableHeader CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableHeaderCell
	CGPDFTagTypeTableHeaderCell CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/tableRow
	CGPDFTagTypeTableRow CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/TOC
	CGPDFTagTypeTOC CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/TOCI
	CGPDFTagTypeTOCI CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichu
	CGPDFTagTypeWarichu CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichuPunctiation
	CGPDFTagTypeWarichuPunctiation CGPDFTagType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGPDFTagType/warichuText
	CGPDFTagTypeWarichuText CGPDFTagType = 0
)

// CGScreenUpdateOperation - Types of screen-update operations.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation
type CGScreenUpdateOperation uint

const (
	// kCGScreenUpdateOperationMove - A screen-move operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/move
	kCGScreenUpdateOperationMove CGScreenUpdateOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/reducedDirtyRectangleCount
	kCGScreenUpdateOperationReducedDirtyRectangleCount CGScreenUpdateOperation = 0
	// kCGScreenUpdateOperationRefresh - A screen-refresh operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScreenUpdateOperation/refresh
	kCGScreenUpdateOperationRefresh CGScreenUpdateOperation = 0
)

// CGScrollEventUnit - Constants that specify the unit of measurement for a scrolling event.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit
type CGScrollEventUnit uint

const (
	// kCGScrollEventUnitLine - Specifies that the unit of measurement is lines.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit/line
	kCGScrollEventUnitLine CGScrollEventUnit = 0
	// kCGScrollEventUnitPixel - Specifies that the unit of measurement is pixels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollEventUnit/pixel
	kCGScrollEventUnitPixel CGScrollEventUnit = 0
)

// CGScrollPhase enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase
type CGScrollPhase uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/began
	kCGScrollPhaseBegan CGScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/cancelled
	kCGScrollPhaseCancelled CGScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/changed
	kCGScrollPhaseChanged CGScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/ended
	kCGScrollPhaseEnded CGScrollPhase = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGScrollPhase/mayBegin
	kCGScrollPhaseMayBegin CGScrollPhase = 0
)

// CGTextDrawingMode - Modes for rendering text.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode
type CGTextDrawingMode uint

const (
	// kCGTextClip - Specifies to intersect the text with the current clipping path. This mode does not paint the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/clip
	kCGTextClip CGTextDrawingMode = 0
	// kCGTextFill - Perform a fill operation on the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fill
	kCGTextFill CGTextDrawingMode = 0
	// kCGTextFillClip - Perform a fill operation, then intersect the text with the current clipping path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillClip
	kCGTextFillClip CGTextDrawingMode = 0
	// kCGTextFillStroke - Perform fill, then stroke operations on the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillStroke
	kCGTextFillStroke CGTextDrawingMode = 0
	// kCGTextFillStrokeClip - Perform fill then stroke operations, then intersect the text with the current clipping path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/fillStrokeClip
	kCGTextFillStrokeClip CGTextDrawingMode = 0
	// kCGTextInvisible - Do not draw the text, but do update the text position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/invisible
	kCGTextInvisible CGTextDrawingMode = 0
	// kCGTextStroke - Perform a stroke operation on the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/stroke
	kCGTextStroke CGTextDrawingMode = 0
	// kCGTextStrokeClip - Perform a stroke operation, then intersect the text with the current clipping path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextDrawingMode/strokeClip
	kCGTextStrokeClip CGTextDrawingMode = 0
)

// CGTextEncoding - Text encodings for fonts.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding
type CGTextEncoding uint

const (
	// kCGEncodingFontSpecific - The built-in encoding of the font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding/encodingFontSpecific
	kCGEncodingFontSpecific CGTextEncoding = 0
	// kCGEncodingMacRoman - The MacRoman encoding. MacRoman is an ASCII variant originally created for use in the Mac OS, in which characters 127 and lower are ASCII, and characters 128 and higher are non-English characters and symbols.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGTextEncoding/encodingMacRoman
	kCGEncodingMacRoman CGTextEncoding = 0
)

// CGToneMapping enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping
type CGToneMapping uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/default
	kCGToneMappingDefault CGToneMapping = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/exrGamma
	kCGToneMappingEXRGamma CGToneMapping = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/imageSpecificLumaScaling
	kCGToneMappingImageSpecificLumaScaling CGToneMapping = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/ituRecommended
	kCGToneMappingITURecommended CGToneMapping = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/none
	kCGToneMappingNone CGToneMapping = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGToneMapping/referenceWhiteBased
	kCGToneMappingReferenceWhiteBased CGToneMapping = 0
)

// CGWindowBackingType - The data type used to specify the backing option for a given window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType
type CGWindowBackingType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreBuffered
	kCGBackingStoreBuffered CGWindowBackingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreNonretained
	kCGBackingStoreNonretained CGWindowBackingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowBackingType/backingStoreRetained
	kCGBackingStoreRetained CGWindowBackingType = 0
)

// CGWindowImageOption - The data type to use to specify the type of image to be generated for a window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption
type CGWindowImageOption uint

const (
	// kCGWindowImageBestResolution - When capturing the window, return the best image resolution. The returned image size may be different than the screen size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/bestResolution
	kCGWindowImageBestResolution CGWindowImageOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/boundsIgnoreFraming
	kCGWindowImageBoundsIgnoreFraming CGWindowImageOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/kCGWindowImageDefault
	kCGWindowImageDefault CGWindowImageOption = 0
	// kCGWindowImageNominalResolution - When capturing the window, return the nominal image resolution. The returned image size is the same as the screen size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/nominalResolution
	kCGWindowImageNominalResolution CGWindowImageOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/onlyShadows
	kCGWindowImageOnlyShadows CGWindowImageOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowImageOption/shouldBeOpaque
	kCGWindowImageShouldBeOpaque CGWindowImageOption = 0
)

// CGWindowLevelKey - Keys that represent the standard window levels in macOS. Quartz includes these keys to support application frameworks like Cocoa. Applications do not need to use them directly.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey
type CGWindowLevelKey uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/assistiveTechHighWindow
	kCGAssistiveTechHighWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/backstopMenu
	kCGBackstopMenuLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/baseWindow
	kCGBaseWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/cursorWindow
	kCGCursorWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/desktopIconWindow
	kCGDesktopIconWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/desktopWindow
	kCGDesktopWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/dockWindow
	kCGDockWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/draggingWindow
	kCGDraggingWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/floatingWindow
	kCGFloatingWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/helpWindow
	kCGHelpWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/mainMenuWindow
	kCGMainMenuWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/maximumWindow
	kCGMaximumWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/minimumWindow
	kCGMinimumWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/modalPanelWindow
	kCGModalPanelWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/normalWindow
	kCGNormalWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/numberOfWindowLevelKeys
	kCGNumberOfWindowLevelKeys CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/overlayWindow
	kCGOverlayWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/popUpMenuWindow
	kCGPopUpMenuWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/screenSaverWindow
	kCGScreenSaverWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/statusWindow
	kCGStatusWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/tornOffMenuWindow
	kCGTornOffMenuWindowLevelKey CGWindowLevelKey = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowLevelKey/utilityWindow
	kCGUtilityWindowLevelKey CGWindowLevelKey = 0
)

// CGWindowListOption - The data type used to specify the options for gathering a list of windows.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption
type CGWindowListOption uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/excludeDesktopElements
	kCGWindowListExcludeDesktopElements CGWindowListOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionAll
	kCGWindowListOptionAll CGWindowListOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionIncludingWindow
	kCGWindowListOptionIncludingWindow CGWindowListOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenAboveWindow
	kCGWindowListOptionOnScreenAboveWindow CGWindowListOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenBelowWindow
	kCGWindowListOptionOnScreenBelowWindow CGWindowListOption = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowListOption/optionOnScreenOnly
	kCGWindowListOptionOnScreenOnly CGWindowListOption = 0
)

// CGWindowSharingType - The data type used to specify the sharing mode used by a window.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType
type CGWindowSharingType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/none
	kCGWindowSharingNone CGWindowSharingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/readOnly
	kCGWindowSharingReadOnly CGWindowSharingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGWindowSharingType/readWrite
	kCGWindowSharingReadWrite CGWindowSharingType = 0
)


