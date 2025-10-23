// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Enum types and constants
// NSAccessibilityAnnotationPosition - Constants that specify the position where the annotation applies.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition
type AccessibilityAnnotationPosition uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition/start
	AccessibilityAnnotationPositionStart AccessibilityAnnotationPosition = 1
)

// NSAccessibilityOrientation - Values that indicate the orientation of accessibility elements, such as scroll bars and split views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation
type AccessibilityOrientation uint

const (
	// AccessibilityOrientationHorizontal - The element is oriented horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/horizontal
	AccessibilityOrientationHorizontal AccessibilityOrientation = 2
	// AccessibilityOrientationUnknown - The element has unknown orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/unknown
	AccessibilityOrientationUnknown AccessibilityOrientation = 0
)

// NSAccessibilityRulerMarkerType - Values that indicate the marker type of an accessibility element.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType
type AccessibilityRulerMarkerType uint

const (
	// AccessibilityRulerMarkerTypeIndentHead - Head indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentHead
	AccessibilityRulerMarkerTypeIndentHead AccessibilityRulerMarkerType = 5
	// AccessibilityRulerMarkerTypeIndentTail - Tail indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentTail
	AccessibilityRulerMarkerTypeIndentTail AccessibilityRulerMarkerType = 6
	// AccessibilityRulerMarkerTypeTabStopCenter - Center tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopCenter
	AccessibilityRulerMarkerTypeTabStopCenter AccessibilityRulerMarkerType = 3
	// AccessibilityRulerMarkerTypeTabStopDecimal - Decimal tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopDecimal
	AccessibilityRulerMarkerTypeTabStopDecimal AccessibilityRulerMarkerType = 4
)

// NSAccessibilitySortDirection - Values that indicate the sort direction of a column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
type AccessibilitySortDirection uint

const (
	AccessibilitySortDirectionUnknown AccessibilitySortDirection = 0
	AccessibilitySortDirectionAscending AccessibilitySortDirection = 1
	AccessibilitySortDirectionDescending AccessibilitySortDirection = 2
	AccessibilityRulerMarkerTypeUnknown AccessibilitySortDirection = 0
	AccessibilityRulerMarkerTypeTabStopLeft AccessibilitySortDirection = 1
	AccessibilityRulerMarkerTypeTabStopRight AccessibilitySortDirection = 2
	AccessibilityRulerMarkerTypeTabStopCenter AccessibilitySortDirection = 3
	AccessibilityRulerMarkerTypeTabStopDecimal AccessibilitySortDirection = 4
	AccessibilityRulerMarkerTypeIndentHead AccessibilitySortDirection = 5
	AccessibilityRulerMarkerTypeIndentTail AccessibilitySortDirection = 6
	AccessibilityRulerMarkerTypeIndentFirstLine AccessibilitySortDirection = 7
	AccessibilityUnitsUnknown AccessibilitySortDirection = 0
	AccessibilityUnitsInches AccessibilitySortDirection = 1
	AccessibilityUnitsCentimeters AccessibilitySortDirection = 2
	AccessibilityUnitsPoints AccessibilitySortDirection = 3
	AccessibilityUnitsPicas AccessibilitySortDirection = 4
	AccessibilityPriorityLow AccessibilitySortDirection = 10
	AccessibilityPriorityMedium AccessibilitySortDirection = 50
	AccessibilityPriorityHigh AccessibilitySortDirection = 90
	AccessibilityHourMinuteDateTimeComponentsFlag AccessibilitySortDirection = 0
	AccessibilityHourMinuteSecondDateTimeComponentsFlag AccessibilitySortDirection = 0
	AccessibilityYearMonthDateTimeComponentsFlag AccessibilitySortDirection = 0
	AccessibilityYearMonthDayDateTimeComponentsFlag AccessibilitySortDirection = 0
)

// NSAccessibilityUnits - Values that indicate the unit values of a ruler or layout area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
type AccessibilityUnits uint

const (
	// AccessibilityUnitsPicas - The units are picas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/picas
	AccessibilityUnitsPicas AccessibilityUnits = 4
)

// NSAnimationCurve - These constants describe the curve of an animation—that is, the relative speed of an animation from start to finish.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve
type AnimationCurve uint

const (
	AnimationEaseInOut AnimationCurve = 0
	AnimationEaseIn AnimationCurve = 1
	AnimationEaseOut AnimationCurve = 2
	AnimationLinear AnimationCurve = 3
)

// NSAnimationEffect - The type for standard system animation effects, which include both display and sound.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type AnimationEffect uint

const (
	AnimationEffectDisappearingItemDefault AnimationEffect = 0
	AnimationEffectPoof AnimationEffect = 10
)

// NSApplicationActivationPolicy - Activation policies (used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum
type ApplicationActivationPolicy uint

const (
	// ApplicationActivationPolicyAccessory - The application doesn’t appear in the Dock and doesn’t have a menu bar, but it may be activated programmatically or by clicking on one of its windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/accessory
	ApplicationActivationPolicyAccessory ApplicationActivationPolicy = 1
	// ApplicationActivationPolicyProhibited - The application doesn’t appear in the Dock and may not create windows or be activated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/prohibited
	ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 2
	// ApplicationActivationPolicyRegular - The application is an ordinary app that appears in the Dock and may have a user interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/regular
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = 0
)

// NSApplicationPrintReply - Constants that indicate the outcome of a print request.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply
type ApplicationPrintReply uint

const (
	PrintingCancelled ApplicationPrintReply = 0
	PrintingSuccess ApplicationPrintReply = 1
	PrintingReplyLater ApplicationPrintReply = 2
	PrintingFailure ApplicationPrintReply = 3
)

// NSRemoteNotificationType - These constants determine whether apps launched by remote notifications display a badge.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
type RemoteNotificationType uint

const (
	// RemoteNotificationTypeSound - The app should play a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType/sound
	RemoteNotificationTypeSound RemoteNotificationType = 2
)

// NSBezierPathElement - Constants that specify basic path element commands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
type BezierPathElement uint

const (
	BezierPathElementMoveTo BezierPathElement = 0
	BezierPathElementLineTo BezierPathElement = 1
	BezierPathElementCubicCurveTo BezierPathElement = 2
	BezierPathElementClosePath BezierPathElement = 3
	BezierPathElementQuadraticCurveTo BezierPathElement = 4
	BezierPathElementCurveTo BezierPathElement = 5
)

// NSLineCapStyle - Constants that specify the shape of endpoints for an open path when it is stroked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum
type LineCapStyle uint

const (
	LineCapStyleButt LineCapStyle = 0
	LineCapStyleRound LineCapStyle = 1
	LineCapStyleSquare LineCapStyle = 2
)

// NSLineJoinStyle - Constants that specify the shape of the joins between connected segments of a stroked path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum
type LineJoinStyle uint

const (
	LineJoinStyleMiter LineJoinStyle = 0
	LineJoinStyleRound LineJoinStyle = 1
	LineJoinStyleBevel LineJoinStyle = 2
)

// NSWindingRule - Constants that specify the winding rule a Bézier path uses.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum
type WindingRule uint

const (
	WindingRuleNonZero WindingRule = 0
	WindingRuleEvenOdd WindingRule = 1
)

// NSTIFFCompression - Constants that represent the supported TIFF data-compression schemes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
type TIFFCompression uint

const (
	TIFFCompressionNone TIFFCompression = 1
	TIFFCompressionCCITTFAX3 TIFFCompression = 3
	TIFFCompressionCCITTFAX4 TIFFCompression = 4
	TIFFCompressionLZW TIFFCompression = 5
	TIFFCompressionJPEG TIFFCompression = 6
	TIFFCompressionNEXT TIFFCompression = 32766
	TIFFCompressionPackBits TIFFCompression = 32773
	TIFFCompressionOldJPEG TIFFCompression = 32865
)

// NSBorderType - These constants specify the type of a view’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType
type BorderType uint

const (
	// BezelBorder - A concave border that makes the view look sunken.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/bezelBorder
	BezelBorder BorderType = 2
	// GrooveBorder - A thin border that looks etched around the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/grooveBorder
	GrooveBorder BorderType = 3
	// LineBorder - A black line border around the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
	LineBorder BorderType = 1
)

// NSBoxType - These constants and data type identifies box types, which, in conjunction with a box’s border type, define the appearance of the box.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
type BoxType uint

const (
	BoxPrimary BoxType = 0
	BoxSeparator BoxType = 2
	BoxCustom BoxType = 3
)

// NSTitlePosition - Specify the location of a box’s title with respect to its border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum
type TitlePosition uint

const (
	NoTitle TitlePosition = 0
	AboveTop TitlePosition = 1
	AtTop TitlePosition = 2
	BelowTop TitlePosition = 3
	AboveBottom TitlePosition = 4
	AtBottom TitlePosition = 5
	BelowBottom TitlePosition = 6
)

// NSBezelStyle - The set of bezel styles to style buttons in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
type BezelStyle uint

const (
	// BezelStyleAccessoryBarAction - A button style that you use for extra actions in an accessory toolbar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/accessoryBarAction
	BezelStyleAccessoryBarAction BezelStyle = 12
	// BezelStyleAutomatic - The default button style based on the button’s contents and position within the window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/automatic
	BezelStyleAutomatic BezelStyle = 0
	// BezelStyleGlass - A bezel style with a glass effect
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/glass
	BezelStyleGlass BezelStyle = 16
	// BezelStyleToolbar - A button style that’s appropriate for a toolbar item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/toolbar
	BezelStyleToolbar BezelStyle = 11
)

// NSButtonType - Button types that you can specify using 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType
type ButtonType uint

const (
	// ButtonTypeAccelerator - A button that sends repeating actions as pressure changes occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/accelerator
	ButtonTypeAccelerator ButtonType = 8
	// ButtonTypeMultiLevelAccelerator - A button that allows for a configurable number of stepped pressure levels and provides tactile feedback as the user reaches each step.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/multiLevelAccelerator
	ButtonTypeMultiLevelAccelerator ButtonType = 9
)

// NSGradientType - Specify the gradients used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType
type GradientType uint

const (
	GradientNone GradientType = 0
	GradientConcaveWeak GradientType = 1
	GradientConcaveStrong GradientType = 2
	GradientConvexWeak GradientType = 3
	GradientConvexStrong GradientType = 4
	TintProminenceAutomatic GradientType = 0
	TintProminenceNone GradientType = 1
	TintProminencePrimary GradientType = 2
	TintProminenceSecondary GradientType = 3
	PopoverAppearanceMinimal GradientType = 4
	PopoverAppearanceHUD GradientType = 5
	PopoverBehaviorApplicationDefined GradientType = 0
	PopoverBehaviorTransient GradientType = 1
	PopoverBehaviorSemitransient GradientType = 2
)

// NSCellAttribute - Constants for specifying how a button behaves when pressed and how it displays its state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute
type CellAttribute uint

const (
	CellDisabled CellAttribute = 0
	CellState CellAttribute = 1
	PushInCell CellAttribute = 2
	CellEditable CellAttribute = 3
	ChangeGrayCell CellAttribute = 4
	CellHighlighted CellAttribute = 5
	CellLightsByContents CellAttribute = 6
	CellLightsByGray CellAttribute = 7
	ChangeBackgroundCell CellAttribute = 8
	CellLightsByBackground CellAttribute = 9
	CellIsBordered CellAttribute = 10
	CellHasOverlappingImage CellAttribute = 11
	CellHasImageHorizontal CellAttribute = 12
	CellHasImageOnLeftOrBottom CellAttribute = 13
	CellChangesContents CellAttribute = 14
	CellIsInsetButton CellAttribute = 15
	CellAllowsMixedState CellAttribute = 16
)

// NSCellType - Constants for specifying how a cell represents its data (as text or as an image).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType
type CellType uint

const (
	// TextCellType - Cell displays text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/textCellType
	TextCellType CellType = 1
)

// NSCellHitResult - Constants used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
type CellHitResult uint

const (
	CellHitNone CellHitResult = 0
	CellHitContentArea CellHitResult = 1
	CellHitEditableTextArea CellHitResult = 1
	CellHitTrackableArea CellHitResult = 1
	BackgroundStyleNormal CellHitResult = 0
	BackgroundStyleEmphasized CellHitResult = 1
	BackgroundStyleRaised CellHitResult = 2
	BackgroundStyleLowered CellHitResult = 3
	AnyType CellHitResult = 4
	IntType CellHitResult = 5
	PositiveIntType CellHitResult = 6
	FloatType CellHitResult = 7
	PositiveFloatType CellHitResult = 8
	DoubleType CellHitResult = 9
	PositiveDoubleType CellHitResult = 10
)

// NSCellStyleMask - Constants for specifying what happens when a button is pressed or is displaying its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask
type CellStyleMask uint

const (
	NoCellMask CellStyleMask = 0
	ContentsCellMask CellStyleMask = 1
	PushInCellMask CellStyleMask = 2
	ChangeGrayCellMask CellStyleMask = 4
	ChangeBackgroundCellMask CellStyleMask = 8
)

// NSCollectionViewDropOperation - These constants specify if acceptance of a drop should be at the item it is dropped on or before the item. These constants are used by the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation
type CollectionViewDropOperation uint

const (
	// CollectionViewDropBefore - The drop occurs before the collection view item to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/before
	CollectionViewDropBefore CollectionViewDropOperation = 1
)

// NSCollectionViewScrollDirection - Constants indicating the scrolling direction for the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection
type CollectionViewScrollDirection uint

const (
	CollectionViewScrollDirectionVertical CollectionViewScrollDirection = 0
	CollectionViewScrollDirectionHorizontal CollectionViewScrollDirection = 1
	DirectionalRectEdgeNone CollectionViewScrollDirection = 0
	DirectionalRectEdgeTop CollectionViewScrollDirection = 1
	DirectionalRectEdgeLeading CollectionViewScrollDirection = 1
	DirectionalRectEdgeBottom CollectionViewScrollDirection = 1
	DirectionalRectEdgeTrailing CollectionViewScrollDirection = 1
	RectAlignmentNone CollectionViewScrollDirection = 0
	RectAlignmentTop CollectionViewScrollDirection = 1
	RectAlignmentTopLeading CollectionViewScrollDirection = 2
	RectAlignmentLeading CollectionViewScrollDirection = 3
	RectAlignmentBottomLeading CollectionViewScrollDirection = 4
	RectAlignmentBottom CollectionViewScrollDirection = 5
	RectAlignmentBottomTrailing CollectionViewScrollDirection = 6
	RectAlignmentTrailing CollectionViewScrollDirection = 7
	RectAlignmentTopTrailing CollectionViewScrollDirection = 8
)

// NSCollectionViewScrollPosition - Constants indicating the options for scrolling the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition
type CollectionViewScrollPosition uint

const (
	CollectionViewScrollPositionNone CollectionViewScrollPosition = 0
	CollectionViewScrollPositionTop CollectionViewScrollPosition = 1
	CollectionViewScrollPositionCenteredVertically CollectionViewScrollPosition = 1
	CollectionViewScrollPositionBottom CollectionViewScrollPosition = 1
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 1
	CollectionViewScrollPositionLeft CollectionViewScrollPosition = 1
	CollectionViewScrollPositionCenteredHorizontally CollectionViewScrollPosition = 1
	CollectionViewScrollPositionRight CollectionViewScrollPosition = 1
	CollectionViewScrollPositionLeadingEdge CollectionViewScrollPosition = 1
	CollectionViewScrollPositionTrailingEdge CollectionViewScrollPosition = 1
	CollectionViewScrollPositionNearestVerticalEdge CollectionViewScrollPosition = 1
)

// NSCollectionUpdateAction - Constants indicating the type of action being performed on an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
type CollectionUpdateAction uint

const (
	CollectionUpdateActionInsert CollectionUpdateAction = 0
	CollectionUpdateActionDelete CollectionUpdateAction = 1
	CollectionUpdateActionReload CollectionUpdateAction = 2
	CollectionUpdateActionMove CollectionUpdateAction = 3
	CollectionUpdateActionNone CollectionUpdateAction = 4
	CollectionViewScrollDirectionVertical CollectionUpdateAction = 5
	CollectionViewScrollDirectionHorizontal CollectionUpdateAction = 6
	DirectionalRectEdgeNone CollectionUpdateAction = 0
	DirectionalRectEdgeTop CollectionUpdateAction = 1
	DirectionalRectEdgeLeading CollectionUpdateAction = 1
	DirectionalRectEdgeBottom CollectionUpdateAction = 1
	DirectionalRectEdgeTrailing CollectionUpdateAction = 1
	RectAlignmentNone CollectionUpdateAction = 0
	RectAlignmentTop CollectionUpdateAction = 1
	RectAlignmentTopLeading CollectionUpdateAction = 2
	RectAlignmentLeading CollectionUpdateAction = 3
	RectAlignmentBottomLeading CollectionUpdateAction = 4
	RectAlignmentBottom CollectionUpdateAction = 5
	RectAlignmentBottomTrailing CollectionUpdateAction = 6
	RectAlignmentTrailing CollectionUpdateAction = 7
	RectAlignmentTopTrailing CollectionUpdateAction = 8
)

// NSCollectionViewItemHighlightState - Constants indicating the type of highlight applied to an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum
type CollectionViewItemHighlightState uint

const (
	CollectionViewItemHighlightNone CollectionViewItemHighlightState = 0
	CollectionViewItemHighlightForSelection CollectionViewItemHighlightState = 1
	CollectionViewItemHighlightForDeselection CollectionViewItemHighlightState = 2
	CollectionViewItemHighlightAsDropTarget CollectionViewItemHighlightState = 3
	CollectionViewScrollPositionNone CollectionViewItemHighlightState = 0
	CollectionViewScrollPositionTop CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionCenteredVertically CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionBottom CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionLeft CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionCenteredHorizontally CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionRight CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionLeadingEdge CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionTrailingEdge CollectionViewItemHighlightState = 1
	CollectionViewScrollPositionNearestVerticalEdge CollectionViewItemHighlightState = 1
)

// NSColorType - Constants that indicate the color’s type, and which methods may be called on the color object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType
type ColorType uint

const (
	// ColorTypeCatalog - Colors that are retrieved from an asset catalog.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/catalog
	ColorTypeCatalog ColorType = 2
	// ColorTypeComponentBased - Colors that include floating-point color components and a color space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/componentBased
	ColorTypeComponentBased ColorType = 0
	// ColorTypePattern - Colors that include an image to be used as a pattern.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/pattern
	ColorTypePattern ColorType = 1
)

// NSColorSystemEffect - Constants for user interactions that change the appearance of a view or control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type ColorSystemEffect uint

const (
	// ColorSystemEffectDeepPressed - The color that indicates the item received a deep press.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/deepPressed
	ColorSystemEffectDeepPressed ColorSystemEffect = 2
	// ColorSystemEffectDisabled - The color that indicates the item is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/disabled
	ColorSystemEffectDisabled ColorSystemEffect = 3
	// ColorSystemEffectPressed - The color that indicates the item was pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/pressed
	ColorSystemEffectPressed ColorSystemEffect = 1
)

// NSColorPanelMode - A type defined for the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum
type ColorPanelMode uint

const (
	ColorPanelModeNone ColorPanelMode = 0
	ColorPanelModeGray ColorPanelMode = 0
	ColorPanelModeRGB ColorPanelMode = 1
	ColorPanelModeCMYK ColorPanelMode = 2
	ColorPanelModeHSB ColorPanelMode = 3
	ColorPanelModeCustomPalette ColorPanelMode = 4
	ColorPanelModeColorList ColorPanelMode = 5
	ColorPanelModeWheel ColorPanelMode = 6
	ColorPanelModeCrayon ColorPanelMode = 7
)

// NSColorPanelOptions - The color modes that are enabled for a color panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Options
type ColorPanelOptions uint

const (
	ColorPanelGrayModeMask ColorPanelOptions = 0
	ColorPanelRGBModeMask ColorPanelOptions = 0
	ColorPanelCMYKModeMask ColorPanelOptions = 0
	ColorPanelHSBModeMask ColorPanelOptions = 0
	ColorPanelCustomPaletteModeMask ColorPanelOptions = 0
	ColorPanelColorListModeMask ColorPanelOptions = 0
	ColorPanelWheelModeMask ColorPanelOptions = 0
	ColorPanelCrayonModeMask ColorPanelOptions = 0
	ColorPanelAllModesMask ColorPanelOptions = 0
)

// NSColorRenderingIntent - Constants that specify how Cocoa should handle colors that are not located within the destination color space of a graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
type ColorRenderingIntent uint

const (
	// ColorRenderingIntentAbsoluteColorimetric - Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/absoluteColorimetric
	ColorRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
	// ColorRenderingIntentDefault - Use the default rendering intent for the graphics context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/default
	ColorRenderingIntentDefault ColorRenderingIntent = 0
	// ColorRenderingIntentPerceptual - Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/perceptual
	ColorRenderingIntentPerceptual ColorRenderingIntent = 3
	// ColorRenderingIntentRelativeColorimetric - Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/relativeColorimetric
	ColorRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
	// ColorRenderingIntentSaturation - Preserve the relative saturation value of the colors when converting into the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/saturation
	ColorRenderingIntentSaturation ColorRenderingIntent = 4
)

// NSColorWellStyle - Constants that specify the appearance and interaction modes for a color well.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
type ColorWellStyle uint

const (
	// ColorWellStyleExpanded - A style that supports a color picker popover for fast interactions, and adds a dedicated button to display the color panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
	ColorWellStyleExpanded ColorWellStyle = 2
	// ColorWellStyleMinimal - A style that adds minimal adornments to the color well.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/minimal
	ColorWellStyleMinimal ColorWellStyle = 1
)

// NSCompositingOperation - Constants that describe compositing operators in terms of source and destination images, each having an opaque and transparent region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
type CompositingOperation uint

const (
	// CompositingOperationClear - Transparency everywhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/clear
	CompositingOperationClear CompositingOperation = 0
	// CompositingOperationSaturation - Uses the saturation value of the source and the hue and luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/saturation
	CompositingOperationSaturation CompositingOperation = 26
	// CompositingOperationSourceOver - The source image wherever it is opaque, and the destination image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceOver
	CompositingOperationSourceOver CompositingOperation = 2
)

// NSControlBorderShape enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape
type ControlBorderShape uint

const (
	// ControlBorderShapeAutomatic - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/automatic
	ControlBorderShapeAutomatic ControlBorderShape = 0
	// ControlBorderShapeCapsule - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/capsule
	ControlBorderShapeCapsule ControlBorderShape = 0
	// ControlBorderShapeCircle - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/circle
	ControlBorderShapeCircle ControlBorderShape = 0
	// ControlBorderShapeRoundedRectangle - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/roundedRectangle
	ControlBorderShapeRoundedRectangle ControlBorderShape = 0
)

// NSControlSize - A constant for specifying a cell’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
type ControlSize uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/extraLarge
	ControlSizeExtraLarge ControlSize = 4
	// ControlSizeLarge - A size larger than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/large
	ControlSizeLarge ControlSize = 3
	// ControlSizeMini - The smallest control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/mini
	ControlSizeMini ControlSize = 2
	// ControlSizeRegular - The default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/regular
	ControlSizeRegular ControlSize = 0
	// ControlSizeSmall - A size smaller than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/small
	ControlSizeSmall ControlSize = 1
)

// NSCellImagePosition - A constant for specifying the position of a button’s image relative to its title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition
type CellImagePosition uint

const (
	NoImage CellImagePosition = 0
	ImageOnly CellImagePosition = 1
	ImageLeft CellImagePosition = 2
	ImageRight CellImagePosition = 3
	ImageBelow CellImagePosition = 4
	ImageAbove CellImagePosition = 5
	ImageOverlaps CellImagePosition = 6
	ImageLeading CellImagePosition = 7
	ImageTrailing CellImagePosition = 8
)

// NSControlTint - Constants for specifying a cell’s tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint
type ControlTint uint

const (
	DefaultControlTint ControlTint = 0
	BlueControlTint ControlTint = 1
	GraphiteControlTint ControlTint = 6
	ClearControlTint ControlTint = 7
)

// NSCursorFrameResizePosition - The position along the perimeter of a rectangular frame (its edges and corners) from which it’s resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition
type CursorFrameResizePosition uint

const (
	// CursorFrameResizePositionBottom - The bottom edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottom
	CursorFrameResizePositionBottom CursorFrameResizePosition = 0
	// CursorFrameResizePositionBottomLeft - The bottom left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottomLeft
	CursorFrameResizePositionBottomLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionLeft - The left edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/left
	CursorFrameResizePositionLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionTop - The top edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/top
	CursorFrameResizePositionTop CursorFrameResizePosition = 0
	// CursorFrameResizePositionTopLeft - The top left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/topLeft
	CursorFrameResizePositionTopLeft CursorFrameResizePosition = 0
)

// NSCursorFrameResizeDirections - The directions in which a rectangular frame can be resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections
type CursorFrameResizeDirections uint

const (
	// CursorFrameResizeDirectionsOutward - Indicates that the shape can be resized outwards to be larger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsOutward
	CursorFrameResizeDirectionsOutward CursorFrameResizeDirections = 0
)

// NSDatePickerMode - Constants that define whether the picker provides a single date, or a range of dates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode
type DatePickerMode uint

const (
	DatePickerModeSingle DatePickerMode = 0
	DatePickerModeRange DatePickerMode = 1
)

// NSDirectionalRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge
type DirectionalRectEdge uint

const (
	DirectionalRectEdgeNone DirectionalRectEdge = 0
	DirectionalRectEdgeTop DirectionalRectEdge = 1
	DirectionalRectEdgeLeading DirectionalRectEdge = 1
	DirectionalRectEdgeBottom DirectionalRectEdge = 1
	DirectionalRectEdgeTrailing DirectionalRectEdge = 1
	RectAlignmentNone DirectionalRectEdge = 0
	RectAlignmentTop DirectionalRectEdge = 1
	RectAlignmentTopLeading DirectionalRectEdge = 2
	RectAlignmentLeading DirectionalRectEdge = 3
	RectAlignmentBottomLeading DirectionalRectEdge = 4
	RectAlignmentBottom DirectionalRectEdge = 5
	RectAlignmentBottomTrailing DirectionalRectEdge = 6
	RectAlignmentTrailing DirectionalRectEdge = 7
	RectAlignmentTopTrailing DirectionalRectEdge = 8
)

// NSDisplayGamut enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut
type DisplayGamut uint

const (
	DisplayGamutSRGB DisplayGamut = 1
	DisplayGamutP3 DisplayGamut = 2
	AnimationEffectDisappearingItemDefault DisplayGamut = 0
	AnimationEffectPoof DisplayGamut = 10
)

// NSDocumentChangeType - Values that indicate a document’s edit status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType
type DocumentChangeType uint

const (
	ChangeDone DocumentChangeType = 0
	ChangeUndone DocumentChangeType = 1
	ChangeRedone DocumentChangeType = 2
	ChangeCleared DocumentChangeType = 2
	ChangeReadOtherContents DocumentChangeType = 3
	ChangeAutosaved DocumentChangeType = 4
	ChangeDiscardable DocumentChangeType = 5
)

// NSSaveOperationType - Constants for specifying the type of document-save operation to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
type SaveOperationType uint

const (
	// AutosaveAsOperation - An operation that writes a document’s contents to a new file or file package even though the user has not explicitly requested it, then changes the document’s current location to point to the just-written file or file package.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveAsOperation
	AutosaveAsOperation SaveOperationType = 5
	// AutosaveElsewhereOperation - An operation that writes an autosave version of the file to a different location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveElsewhereOperation
	AutosaveElsewhereOperation SaveOperationType = 4
	// AutosaveInPlaceOperation - An operation that overwrites the document’s current contents with autosave data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveInPlaceOperation
	AutosaveInPlaceOperation SaveOperationType = 3
	// SaveAsOperation - An operation that writes the document’s contents to a new location and updates the document to point to that location
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveAsOperation
	SaveAsOperation SaveOperationType = 1
	// SaveOperation - An operation that overwrites a document’s file or file package with the document’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
	SaveOperation SaveOperationType = 0
	// SaveToOperation - An operation that writes a copy of the document’s contents to the specified location, without changing the original document’s location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveToOperation
	SaveToOperation SaveOperationType = 2
	// AutosaveOperation - Old name for the   operation type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSaveOperationType/NSAutosaveOperation
	AutosaveOperation SaveOperationType = 6
)

// NSDragOperation - A group of constants that represent which operations the dragging source can perform on dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation
type DragOperation uint

const (
	// DragOperationAll_Obsolete - The   constant is deprecated. Use   instead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/all_Obsolete
	DragOperationAll_Obsolete DragOperation = 33
	// DragOperationDelete - A constant that indicates the drag can delete the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/delete
	DragOperationDelete DragOperation = 32
	// DragOperationEvery - A constant that indicates that drag can perform all of the drag operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/every
	DragOperationEvery DragOperation = 0
	// DragOperationMove - A constant that indicates the drag can move the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/move
	DragOperationMove DragOperation = 16
)

// NSEventButtonMask - Constants you use to identify the activated tablet buttons in an event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct
type EventButtonMask uint

const (
	// EventButtonMaskPenLowerSide - A mask that matches the button on the lower side of the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct/penLowerSide
	EventButtonMaskPenLowerSide EventButtonMask = 0
)

// NSEventType - Constants for the types of events that responder objects can handle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType
type EventType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseCancelled
	EventTypeMouseCancelled EventType = 39
)

// NSEventMask - Constants that you use to filter out specific event types from the stream of incoming events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask
type EventMask uint

const (
	// EventMaskPressure - A mask for pressure-change events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/pressure
	EventMaskPressure EventMask = 0
)

// NSEventGestureAxis - Constants that specify the direction of travel for a gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis
type EventGestureAxis uint

const (
	// EventGestureAxisHorizontal - The horizontal axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/horizontal
	EventGestureAxisHorizontal EventGestureAxis = 1
	// EventGestureAxisNone - No specific axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/none
	EventGestureAxisNone EventGestureAxis = 0
)

// NSEventModifierFlags - Flags that represent key states in an event object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type EventModifierFlags uint

const (
	EventModifierFlagCapsLock EventModifierFlags = 1
	EventModifierFlagShift EventModifierFlags = 1
	EventModifierFlagControl EventModifierFlags = 1
	EventModifierFlagOption EventModifierFlags = 1
	EventModifierFlagCommand EventModifierFlags = 1
	EventModifierFlagNumericPad EventModifierFlags = 1
	EventModifierFlagHelp EventModifierFlags = 1
	EventModifierFlagFunction EventModifierFlags = 1
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 0
)

// NSEventPhase - Constants that represent the possible phases during an event phase.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct
type EventPhase uint

const (
	// EventPhaseEnded - The event phase ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/ended
	EventPhaseEnded EventPhase = 0
	// EventPhaseMayBegin - The system event phase may begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/mayBegin
	EventPhaseMayBegin EventPhase = 0
)

// NSEventSwipeTrackingOptions - Constants that specify swipe-tracking options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions
type EventSwipeTrackingOptions uint

const (
	// EventSwipeTrackingClampGestureAmount - Don’t allow gestureAmount to go beyond +/-1.0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/clampGestureAmount
	EventSwipeTrackingClampGestureAmount EventSwipeTrackingOptions = 0
	// EventSwipeTrackingLockDirection - Clamp gestureAmount to 0 if the user starts to swipe in the opposite direction than they started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/lockDirection
	EventSwipeTrackingLockDirection EventSwipeTrackingOptions = 0
)

// NSFocusRingPlacement - Constants that indicate how the system draws the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement
type FocusRingPlacement uint

const (
	FocusRingOnly FocusRingPlacement = 0
	FocusRingBelow FocusRingPlacement = 1
	FocusRingAbove FocusRingPlacement = 2
)

// NSFocusRingType - Constants that describe the style of the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type FocusRingType uint

const (
	FocusRingTypeDefault FocusRingType = 0
	FocusRingTypeNone FocusRingType = 1
	FocusRingTypeExterior FocusRingType = 2
)

// NSFontCollectionOptions - Constants that support font collection management.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type FontCollectionOptions uint

const (
	FontCollectionApplicationOnlyMask FontCollectionOptions = 1
)

// NSFontDescriptorSymbolicTraits - A symbolic description of the stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct
type FontDescriptorSymbolicTraits uint

const (
	FontDescriptorTraitItalic FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitBold FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitExpanded FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitCondensed FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitMonoSpace FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitVertical FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitUIOptimized FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitTightLeading FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitLooseLeading FontDescriptorSymbolicTraits = 1
	FontDescriptorTraitEmphasized FontDescriptorSymbolicTraits = 2
	FontDescriptorClassMask FontDescriptorSymbolicTraits = 0
	FontDescriptorClassUnknown FontDescriptorSymbolicTraits = 0
	FontDescriptorClassOldStyleSerifs FontDescriptorSymbolicTraits = 1
	FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 2
	FontDescriptorClassModernSerifs FontDescriptorSymbolicTraits = 3
	FontDescriptorClassClarendonSerifs FontDescriptorSymbolicTraits = 4
	FontDescriptorClassSlabSerifs FontDescriptorSymbolicTraits = 5
	FontDescriptorClassFreeformSerifs FontDescriptorSymbolicTraits = 7
	FontDescriptorClassSansSerif FontDescriptorSymbolicTraits = 8
	FontDescriptorClassOrnamentals FontDescriptorSymbolicTraits = 9
	FontDescriptorClassScripts FontDescriptorSymbolicTraits = 10
	FontDescriptorClassSymbolic FontDescriptorSymbolicTraits = 12
)

// NSFontPanelModeMask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask
type FontPanelModeMask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask/collection
	FontPanelModeMaskCollection FontPanelModeMask = 1
)

// NSFontRenderingMode - The font rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type FontRenderingMode uint

const (
	FontDefaultRenderingMode FontRenderingMode = 0
	FontAntialiasedRenderingMode FontRenderingMode = 1
	FontIntegerAdvancementsRenderingMode FontRenderingMode = 2
	FontAntialiasedIntegerAdvancementsRenderingMode FontRenderingMode = 3
)

// NSFontTraitMask - Constants for isolating specific traits of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask
type FontTraitMask uint

const (
	// CondensedFontMask - A mask that specifies a condensed font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/condensedFontMask
	CondensedFontMask FontTraitMask = 0
	// ExpandedFontMask - A mask that specifies an expanded font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/expandedFontMask
	ExpandedFontMask FontTraitMask = 0
	// NonStandardCharacterSetFontMask - A mask that specifies a font containing a non-standard character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/nonStandardCharacterSetFontMask
	NonStandardCharacterSetFontMask FontTraitMask = 0
)

// NSGlassEffectViewStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum
type GlassEffectViewStyle uint

const (
	// GlassEffectViewStyleClear - Clear glass effect style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum/clear
	GlassEffectViewStyleClear GlassEffectViewStyle = 0
	// GlassEffectViewStyleRegular - Standard glass effect style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum/regular
	GlassEffectViewStyleRegular GlassEffectViewStyle = 0
)

// NSGlyphInscription - Constants that specify how a glyph is laid out relative to the previous glyph.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription
type GlyphInscription uint

const (
	GlyphInscribeBase GlyphInscription = 0
	GlyphInscribeBelow GlyphInscription = 1
	GlyphInscribeAbove GlyphInscription = 2
	GlyphInscribeOverstrike GlyphInscription = 3
	GlyphInscribeOverBelow GlyphInscription = 4
	LineSweepLeft GlyphInscription = 0
	LineSweepRight GlyphInscription = 1
	LineSweepDown GlyphInscription = 2
	LineSweepUp GlyphInscription = 3
)

// NSGradientDrawingOptions - Constants that specify gradient drawing options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient/DrawingOptions
type GradientDrawingOptions uint

// NSHapticFeedbackPattern - A pattern of haptic feedback to be provided to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern
type HapticFeedbackPattern uint

const (
	// HapticFeedbackPatternAlignment - A haptic feedback pattern to be used in response to the alignment of an object the user is dragging around. For example, this pattern of feedback could be used in a drawing app when the user drags a shape into alignment with another shape. Other scenarios where this type of feedback could be used might include scaling an object to fit within specific dimensions, positioning an object at a preferred location, or reaching the beginning/minimum or end/maximum of something, such as a track view in an audio/video app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/alignment
	HapticFeedbackPatternAlignment HapticFeedbackPattern = 1
	// HapticFeedbackPatternGeneric - A general haptic feedback pattern. Use this when no other feedback patterns apply.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/generic
	HapticFeedbackPatternGeneric HapticFeedbackPattern = 0
	// HapticFeedbackPatternLevelChange - A haptic feedback pattern to be used as the user moves between discrete levels of pressure. This pattern of feedback is used by multilevel accelerator buttons (class  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/levelChange
	HapticFeedbackPatternLevelChange HapticFeedbackPattern = 2
)

// NSHapticFeedbackPerformanceTime - A time at which to provide haptic feedback to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime
type HapticFeedbackPerformanceTime uint

const (
	HapticFeedbackPerformanceTimeDefault HapticFeedbackPerformanceTime = 0
	HapticFeedbackPerformanceTimeNow HapticFeedbackPerformanceTime = 1
	HapticFeedbackPerformanceTimeDrawCompleted HapticFeedbackPerformanceTime = 2
	PickerTouchBarItemSelectionModeSelectOne HapticFeedbackPerformanceTime = 0
	PickerTouchBarItemSelectionModeSelectAny HapticFeedbackPerformanceTime = 1
	PickerTouchBarItemSelectionModeMomentary HapticFeedbackPerformanceTime = 2
	PickerTouchBarItemControlRepresentationAutomatic HapticFeedbackPerformanceTime = 0
	PickerTouchBarItemControlRepresentationExpanded HapticFeedbackPerformanceTime = 1
	PickerTouchBarItemControlRepresentationCollapsed HapticFeedbackPerformanceTime = 2
	TextSelectionGranularityCharacter HapticFeedbackPerformanceTime = 3
	TextSelectionGranularityWord HapticFeedbackPerformanceTime = 4
	TextSelectionGranularityParagraph HapticFeedbackPerformanceTime = 5
	TextSelectionGranularityLine HapticFeedbackPerformanceTime = 6
	TextSelectionGranularitySentence HapticFeedbackPerformanceTime = 7
	TextSelectionAffinityUpstream HapticFeedbackPerformanceTime = 0
	TextSelectionAffinityDownstream HapticFeedbackPerformanceTime = 1
	TextSelectionNavigationDirectionForward HapticFeedbackPerformanceTime = 2
	TextSelectionNavigationDirectionBackward HapticFeedbackPerformanceTime = 3
	TextSelectionNavigationDirectionRight HapticFeedbackPerformanceTime = 4
	TextSelectionNavigationDirectionLeft HapticFeedbackPerformanceTime = 5
	TextSelectionNavigationDirectionUp HapticFeedbackPerformanceTime = 6
	TextSelectionNavigationDirectionDown HapticFeedbackPerformanceTime = 7
	TextSelectionNavigationDestinationCharacter HapticFeedbackPerformanceTime = 8
	TextSelectionNavigationDestinationWord HapticFeedbackPerformanceTime = 9
	TextSelectionNavigationDestinationLine HapticFeedbackPerformanceTime = 10
	TextSelectionNavigationDestinationSentence HapticFeedbackPerformanceTime = 11
	TextSelectionNavigationDestinationParagraph HapticFeedbackPerformanceTime = 12
	TextSelectionNavigationDestinationContainer HapticFeedbackPerformanceTime = 13
	TextSelectionNavigationDestinationDocument HapticFeedbackPerformanceTime = 14
	TextSelectionNavigationWritingDirectionLeftToRight HapticFeedbackPerformanceTime = 0
	TextSelectionNavigationWritingDirectionRightToLeft HapticFeedbackPerformanceTime = 1
	TextSelectionNavigationLayoutOrientationHorizontal HapticFeedbackPerformanceTime = 0
	TextSelectionNavigationLayoutOrientationVertical HapticFeedbackPerformanceTime = 1
	TextContentManagerEnumerationOptionsNone HapticFeedbackPerformanceTime = 0
	TextLayoutFragmentEnumerationOptionsNone HapticFeedbackPerformanceTime = 0
	TextLayoutFragmentStateNone HapticFeedbackPerformanceTime = 0
	TextLayoutFragmentStateEstimatedUsageBounds HapticFeedbackPerformanceTime = 1
	TextLayoutFragmentStateCalculatedUsageBounds HapticFeedbackPerformanceTime = 2
	TextLayoutFragmentStateLayoutAvailable HapticFeedbackPerformanceTime = 3
	TextLayoutManagerSegmentTypeStandard HapticFeedbackPerformanceTime = 0
	TextLayoutManagerSegmentTypeSelection HapticFeedbackPerformanceTime = 1
	TextLayoutManagerSegmentTypeHighlight HapticFeedbackPerformanceTime = 2
	TextLayoutManagerSegmentOptionsNone HapticFeedbackPerformanceTime = 0
	WritingToolsCoordinatorTextUpdateReasonTyping HapticFeedbackPerformanceTime = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo HapticFeedbackPerformanceTime = 2
	WritingToolsCoordinatorStateInactive HapticFeedbackPerformanceTime = 3
	WritingToolsCoordinatorStateNoninteractive HapticFeedbackPerformanceTime = 4
	WritingToolsCoordinatorStateInteractiveResting HapticFeedbackPerformanceTime = 5
	WritingToolsCoordinatorStateInteractiveStreaming HapticFeedbackPerformanceTime = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive HapticFeedbackPerformanceTime = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive HapticFeedbackPerformanceTime = 8
	WritingToolsCoordinatorContextScopeUserSelection HapticFeedbackPerformanceTime = 9
	WritingToolsCoordinatorContextScopeFullDocument HapticFeedbackPerformanceTime = 10
	WritingToolsCoordinatorContextScopeVisibleArea HapticFeedbackPerformanceTime = 11
	WritingToolsCoordinatorTextAnimationAnticipate HapticFeedbackPerformanceTime = 12
	WritingToolsCoordinatorTextAnimationRemove HapticFeedbackPerformanceTime = 13
	WritingToolsCoordinatorTextAnimationInsert HapticFeedbackPerformanceTime = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive HapticFeedbackPerformanceTime = 8
	WritingToolsCoordinatorTextAnimationTranslate HapticFeedbackPerformanceTime = 9
)

// NSHorizontalDirections - The absolute directions on the horizontal axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections
type HorizontalDirections uint

const (
	// HorizontalDirectionsAll - All horizontal directions (left and right).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsAll
	HorizontalDirectionsAll HorizontalDirections = 0
	// HorizontalDirectionsLeft - The left direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsLeft
	HorizontalDirectionsLeft HorizontalDirections = 0
)

// NSImageCacheMode - Constants that specify the caching policy on a per-image basis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum
type ImageCacheMode uint

const (
	// ImageCacheAlways - Always generate a cache when drawing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/always
	ImageCacheAlways ImageCacheMode = 1
	// ImageCacheBySize - Cache if the cache size is smaller than the original data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/bySize
	ImageCacheBySize ImageCacheMode = 2
	// ImageCacheDefault - Caching is unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/default
	ImageCacheDefault ImageCacheMode = 0
	// ImageCacheNever - Never cache; always draw direct.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/never
	ImageCacheNever ImageCacheMode = 3
)

// NSImageDynamicRange - Describes how High Dynamic Range (HDR) image content displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type ImageDynamicRange uint

const (
	// ImageDynamicRangeHigh - Allows image content to use extended dynamic range if it has dynamic range content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/high
	ImageDynamicRangeHigh ImageDynamicRange = 2
)

// NSImageLayoutDirection - Constants that describe the layout direction for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection
type ImageLayoutDirection uint

const (
	// ImageLayoutDirectionLeftToRight - A left-to-right layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection/leftToRight
	ImageLayoutDirectionLeftToRight ImageLayoutDirection = 2
	// ImageLayoutDirectionUnspecified - An unspecified layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection/unspecified
	ImageLayoutDirectionUnspecified ImageLayoutDirection = -1
)

// NSImageLoadStatus - Status values for incremental image loading.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus
type ImageLoadStatus uint

const (
	ImageLoadStatusCompleted ImageLoadStatus = 0
	ImageLoadStatusCancelled ImageLoadStatus = 1
	ImageLoadStatusInvalidData ImageLoadStatus = 2
	ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
	ImageLoadStatusReadError ImageLoadStatus = 4
)

// NSImageResizingMode - Constants that describe the resizing mode for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum
type ImageResizingMode uint

const (
	// ImageResizingModeStretch - The image stretches when it resizes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum/stretch
	ImageResizingModeStretch ImageResizingMode = 1
)

// NSImageSymbolColorRenderingMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode
type ImageSymbolColorRenderingMode uint

const (
	ImageSymbolColorRenderingModeAutomatic ImageSymbolColorRenderingMode = 0
	ImageSymbolColorRenderingModeFlat ImageSymbolColorRenderingMode = 1
	ImageSymbolColorRenderingModeGradient ImageSymbolColorRenderingMode = 2
	ImageAlignCenter ImageSymbolColorRenderingMode = 0
	ImageAlignTop ImageSymbolColorRenderingMode = 1
	ImageAlignTopLeft ImageSymbolColorRenderingMode = 2
	ImageAlignTopRight ImageSymbolColorRenderingMode = 3
	ImageAlignLeft ImageSymbolColorRenderingMode = 4
	ImageAlignBottom ImageSymbolColorRenderingMode = 5
	ImageAlignBottomLeft ImageSymbolColorRenderingMode = 6
	ImageAlignBottomRight ImageSymbolColorRenderingMode = 7
	ImageAlignRight ImageSymbolColorRenderingMode = 8
)

// NSImageSymbolScale - Constants that specify which scale variant of a symbol image to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale
type ImageSymbolScale uint

const (
	ImageSymbolScaleSmall ImageSymbolScale = 1
	ImageSymbolScaleMedium ImageSymbolScale = 2
	ImageSymbolScaleLarge ImageSymbolScale = 3
	ImageSymbolVariableValueModeAutomatic ImageSymbolScale = 0
	ImageSymbolVariableValueModeColor ImageSymbolScale = 1
	ImageSymbolVariableValueModeDraw ImageSymbolScale = 2
	ImageSymbolColorRenderingModeAutomatic ImageSymbolScale = 0
	ImageSymbolColorRenderingModeFlat ImageSymbolScale = 1
	ImageSymbolColorRenderingModeGradient ImageSymbolScale = 2
	ImageAlignCenter ImageSymbolScale = 0
	ImageAlignTop ImageSymbolScale = 1
	ImageAlignTopLeft ImageSymbolScale = 2
	ImageAlignTopRight ImageSymbolScale = 3
	ImageAlignLeft ImageSymbolScale = 4
	ImageAlignBottom ImageSymbolScale = 5
	ImageAlignBottomLeft ImageSymbolScale = 6
	ImageAlignBottomRight ImageSymbolScale = 7
	ImageAlignRight ImageSymbolScale = 8
)

// NSImageSymbolVariableValueMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode
type ImageSymbolVariableValueMode uint

const (
	ImageSymbolVariableValueModeAutomatic ImageSymbolVariableValueMode = 0
	ImageSymbolVariableValueModeColor ImageSymbolVariableValueMode = 1
	ImageSymbolVariableValueModeDraw ImageSymbolVariableValueMode = 2
	ImageSymbolColorRenderingModeAutomatic ImageSymbolVariableValueMode = 0
	ImageSymbolColorRenderingModeFlat ImageSymbolVariableValueMode = 1
	ImageSymbolColorRenderingModeGradient ImageSymbolVariableValueMode = 2
	ImageAlignCenter ImageSymbolVariableValueMode = 0
	ImageAlignTop ImageSymbolVariableValueMode = 1
	ImageAlignTopLeft ImageSymbolVariableValueMode = 2
	ImageAlignTopRight ImageSymbolVariableValueMode = 3
	ImageAlignLeft ImageSymbolVariableValueMode = 4
	ImageAlignBottom ImageSymbolVariableValueMode = 5
	ImageAlignBottomLeft ImageSymbolVariableValueMode = 6
	ImageAlignBottomRight ImageSymbolVariableValueMode = 7
	ImageAlignRight ImageSymbolVariableValueMode = 8
)

// NSImageInterpolation - Constants that specify the interpolation, or image smoothing, behavior used by the image interpolation property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation
type ImageInterpolation uint

const (
	ImageInterpolationDefault ImageInterpolation = 0
	ImageInterpolationNone ImageInterpolation = 1
	ImageInterpolationLow ImageInterpolation = 2
	ImageInterpolationMedium ImageInterpolation = 3
	ImageInterpolationHigh ImageInterpolation = 3
)

// NSImageScaling - Constants that specify a cell’s image scaling behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling
type ImageScaling uint

const (
	ImageScaleProportionallyDown ImageScaling = 0
	ImageScaleAxesIndependently ImageScaling = 1
	ImageScaleNone ImageScaling = 2
	ImageScaleProportionallyUpOrDown ImageScaling = 3
	ScaleProportionally ImageScaling = 4
	ScaleToFit ImageScaling = 5
	ScaleNone ImageScaling = 6
	NoCellMask ImageScaling = 0
	ContentsCellMask ImageScaling = 1
	PushInCellMask ImageScaling = 2
	ChangeGrayCellMask ImageScaling = 4
	ChangeBackgroundCellMask ImageScaling = 8
)

// NSLayoutAttribute - The part of the object’s visual representation that should be used to get the value for the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
type LayoutAttribute uint

const (
	// LayoutAttributeLastBaseline - The object’s baseline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/lastBaseline
	LayoutAttributeLastBaseline LayoutAttribute = 11
	// LayoutAttributeLeading - The leading edge of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/leading
	LayoutAttributeLeading LayoutAttribute = 5
	// LayoutAttributeWidth - The width of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/width
	LayoutAttributeWidth LayoutAttribute = 7
)

// NSLayoutConstraintOrientation - The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
type LayoutConstraintOrientation uint

const (
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	LayoutConstraintOrientationVertical LayoutConstraintOrientation = 1
)

// NSLevelIndicatorPlaceholderVisibility enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum
type LevelIndicatorPlaceholderVisibility uint

const (
	LevelIndicatorPlaceholderVisibilityAutomatic LevelIndicatorPlaceholderVisibility = 0
	LevelIndicatorPlaceholderVisibilityAlways LevelIndicatorPlaceholderVisibility = 1
	LevelIndicatorPlaceholderVisibilityWhileEditing LevelIndicatorPlaceholderVisibility = 2
)

// NSLevelIndicatorStyle - Constants that specify a level indicator’s appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style
type LevelIndicatorStyle uint

const (
	LevelIndicatorStyleRelevancy LevelIndicatorStyle = 0
	LevelIndicatorStyleContinuousCapacity LevelIndicatorStyle = 1
	LevelIndicatorStyleDiscreteCapacity LevelIndicatorStyle = 2
	LevelIndicatorStyleRating LevelIndicatorStyle = 3
)

// NSLineBreakMode - Constants that specify what happens when a line is too long for a container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
type LineBreakMode uint

const (
	// LineBreakByWordWrapping - The value that indicates wrapping occurs at word boundaries, unless the word doesn’t fit on a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
	LineBreakByWordWrapping LineBreakMode = 0
)

// NSLineMovementDirection - The direction in which a line moves.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection
type LineMovementDirection uint

const (
	LineDoesntMove LineMovementDirection = 0
	LineMovesLeft LineMovementDirection = 1
	LineMovesRight LineMovementDirection = 2
	LineMovesDown LineMovementDirection = 3
	LineMovesUp LineMovementDirection = 4
)

// NSLineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

const (
	LineSweepLeft LineSweepDirection = 0
	LineSweepRight LineSweepDirection = 1
	LineSweepDown LineSweepDirection = 2
	LineSweepUp LineSweepDirection = 3
)

// NSMatrixMode - These constants determine how 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
type MatrixMode uint

const (
	RadioModeMatrix MatrixMode = 0
	HighlightModeMatrix MatrixMode = 1
	ListModeMatrix MatrixMode = 2
	TrackModeMatrix MatrixMode = 3
)

// NSMediaLibrary - These constants are masks used to configure a Media Library Browser to display specific types of media. Combined masks are not yet supported.  In other words, only one nonzero mask value is supported at a time.  If masks are combined, the lowest mask value is used.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library
type MediaLibrary uint

const (
	MediaLibraryAudio MediaLibrary = 1
	MediaLibraryImage MediaLibrary = 1
	MediaLibraryMovie MediaLibrary = 1
	NoScrollerParts MediaLibrary = 0
	OnlyScrollerArrows MediaLibrary = 1
	AllScrollerParts MediaLibrary = 2
)

// NSMenuPresentationStyle - Specifies the style of a menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/PresentationStyle-swift.enum
type MenuPresentationStyle uint

const (
	MenuPresentationStyleRegular MenuPresentationStyle = 0
	MenuPresentationStylePalette MenuPresentationStyle = 1
	MenuSelectionModeAutomatic MenuPresentationStyle = 0
	MenuSelectionModeSelectOne MenuPresentationStyle = 1
	MenuSelectionModeSelectAny MenuPresentationStyle = 2
	MenuPropertyItemTitle MenuPresentationStyle = 1
	MenuPropertyItemAttributedTitle MenuPresentationStyle = 1
	MenuPropertyItemKeyEquivalent MenuPresentationStyle = 1
	MenuPropertyItemImage MenuPresentationStyle = 1
	MenuPropertyItemEnabled MenuPresentationStyle = 1
	MenuPropertyItemAccessibilityDescription MenuPresentationStyle = 1
)

// NSMenuProperties - These constants are used as a bitmask for specifying a set of menu or menu item properties, and are contained by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/Properties
type MenuProperties uint

const (
	MenuPropertyItemTitle MenuProperties = 1
	MenuPropertyItemAttributedTitle MenuProperties = 1
	MenuPropertyItemKeyEquivalent MenuProperties = 1
	MenuPropertyItemImage MenuProperties = 1
	MenuPropertyItemEnabled MenuProperties = 1
	MenuPropertyItemAccessibilityDescription MenuProperties = 1
)

// NSMenuSelectionMode - Describes how the menu manages selection states of the menu items that belong to the same selection group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenu/SelectionMode-swift.enum
type MenuSelectionMode uint

const (
	MenuSelectionModeAutomatic MenuSelectionMode = 0
	MenuSelectionModeSelectOne MenuSelectionMode = 1
	MenuSelectionModeSelectAny MenuSelectionMode = 2
	MenuPropertyItemTitle MenuSelectionMode = 1
	MenuPropertyItemAttributedTitle MenuSelectionMode = 1
	MenuPropertyItemKeyEquivalent MenuSelectionMode = 1
	MenuPropertyItemImage MenuSelectionMode = 1
	MenuPropertyItemEnabled MenuSelectionMode = 1
	MenuPropertyItemAccessibilityDescription MenuSelectionMode = 1
)

// NSMenuItemBadgeType - Constants that define types of badges for display.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType
type MenuItemBadgeType uint

const (
	// MenuItemBadgeTypeNone - A badge with no string portion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMenuItemBadge/BadgeType/none
	MenuItemBadgeTypeNone MenuItemBadgeType = 0
)

// NSMultibyteGlyphPacking - A constant for glyph packing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type MultibyteGlyphPacking uint

const (
	NativeShortGlyphPacking MultibyteGlyphPacking = 0
	FontAssetRequestOptionUsesStandardUI MultibyteGlyphPacking = 1
)

// NSOpenGLGlobalOption - Constants that specify OpenGL options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption
type OpenGLGlobalOption uint

const (
	// OpenGLGOUseBuildCache - Whether to enable the function compilation block cache. This is off by default. It must be enabled at startup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/useBuildCache
	OpenGLGOUseBuildCache OpenGLGlobalOption = 3
)

// NSPDFPanelOptions - Constants used to configure the contents of a PDF panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/Options-swift.struct
type PDFPanelOptions uint

const (
	PDFPanelShowsPaperSize PDFPanelOptions = 1
	PDFPanelShowsOrientation PDFPanelOptions = 1
	PDFPanelRequestsParentDirectory PDFPanelOptions = 1
	MediaLibraryAudio PDFPanelOptions = 1
	MediaLibraryImage PDFPanelOptions = 1
	MediaLibraryMovie PDFPanelOptions = 1
	NoScrollerParts PDFPanelOptions = 0
	OnlyScrollerArrows PDFPanelOptions = 1
	AllScrollerParts PDFPanelOptions = 2
)

// NSPageControllerTransitionStyle - These constants control the transition style of the page controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum
type PageControllerTransitionStyle uint

const (
	PageControllerTransitionStyleStackHistory PageControllerTransitionStyle = 0
	PageControllerTransitionStyleStackBook PageControllerTransitionStyle = 1
	PageControllerTransitionStyleHorizontalStrip PageControllerTransitionStyle = 2
	VisualEffectMaterialTitlebar PageControllerTransitionStyle = 3
	VisualEffectMaterialSelection PageControllerTransitionStyle = 4
	VisualEffectMaterialMenu PageControllerTransitionStyle = 5
	VisualEffectMaterialPopover PageControllerTransitionStyle = 6
	VisualEffectMaterialSidebar PageControllerTransitionStyle = 7
	VisualEffectMaterialHeaderView PageControllerTransitionStyle = 8
	VisualEffectMaterialSheet PageControllerTransitionStyle = 9
	VisualEffectMaterialWindowBackground PageControllerTransitionStyle = 10
	VisualEffectMaterialHUDWindow PageControllerTransitionStyle = 11
	VisualEffectMaterialFullScreenUI PageControllerTransitionStyle = 12
	VisualEffectMaterialToolTip PageControllerTransitionStyle = 13
	VisualEffectMaterialContentBackground PageControllerTransitionStyle = 14
	VisualEffectMaterialUnderWindowBackground PageControllerTransitionStyle = 15
	VisualEffectMaterialUnderPageBackground PageControllerTransitionStyle = 16
	VisualEffectMaterialAppearanceBased PageControllerTransitionStyle = 17
	VisualEffectMaterialLight PageControllerTransitionStyle = 18
	VisualEffectMaterialDark PageControllerTransitionStyle = 19
	VisualEffectMaterialMediumLight PageControllerTransitionStyle = 20
	VisualEffectMaterialUltraDark PageControllerTransitionStyle = 21
	VisualEffectBlendingModeBehindWindow PageControllerTransitionStyle = 22
	VisualEffectBlendingModeWithinWindow PageControllerTransitionStyle = 23
	VisualEffectStateFollowsWindowActiveState PageControllerTransitionStyle = 24
	VisualEffectStateActive PageControllerTransitionStyle = 25
	VisualEffectStateInactive PageControllerTransitionStyle = 26
	HapticFeedbackPatternGeneric PageControllerTransitionStyle = 0
	HapticFeedbackPatternAlignment PageControllerTransitionStyle = 1
	HapticFeedbackPatternLevelChange PageControllerTransitionStyle = 2
	HapticFeedbackPerformanceTimeDefault PageControllerTransitionStyle = 0
	HapticFeedbackPerformanceTimeNow PageControllerTransitionStyle = 1
	HapticFeedbackPerformanceTimeDrawCompleted PageControllerTransitionStyle = 2
	PickerTouchBarItemSelectionModeSelectOne PageControllerTransitionStyle = 0
	PickerTouchBarItemSelectionModeSelectAny PageControllerTransitionStyle = 1
	PickerTouchBarItemSelectionModeMomentary PageControllerTransitionStyle = 2
	PickerTouchBarItemControlRepresentationAutomatic PageControllerTransitionStyle = 0
	PickerTouchBarItemControlRepresentationExpanded PageControllerTransitionStyle = 1
	PickerTouchBarItemControlRepresentationCollapsed PageControllerTransitionStyle = 2
	TextSelectionGranularityCharacter PageControllerTransitionStyle = 3
	TextSelectionGranularityWord PageControllerTransitionStyle = 4
	TextSelectionGranularityParagraph PageControllerTransitionStyle = 5
	TextSelectionGranularityLine PageControllerTransitionStyle = 6
	TextSelectionGranularitySentence PageControllerTransitionStyle = 7
	TextSelectionAffinityUpstream PageControllerTransitionStyle = 0
	TextSelectionAffinityDownstream PageControllerTransitionStyle = 1
	TextSelectionNavigationDirectionForward PageControllerTransitionStyle = 2
	TextSelectionNavigationDirectionBackward PageControllerTransitionStyle = 3
	TextSelectionNavigationDirectionRight PageControllerTransitionStyle = 4
	TextSelectionNavigationDirectionLeft PageControllerTransitionStyle = 5
	TextSelectionNavigationDirectionUp PageControllerTransitionStyle = 6
	TextSelectionNavigationDirectionDown PageControllerTransitionStyle = 7
	TextSelectionNavigationDestinationCharacter PageControllerTransitionStyle = 8
	TextSelectionNavigationDestinationWord PageControllerTransitionStyle = 9
	TextSelectionNavigationDestinationLine PageControllerTransitionStyle = 10
	TextSelectionNavigationDestinationSentence PageControllerTransitionStyle = 11
	TextSelectionNavigationDestinationParagraph PageControllerTransitionStyle = 12
	TextSelectionNavigationDestinationContainer PageControllerTransitionStyle = 13
	TextSelectionNavigationDestinationDocument PageControllerTransitionStyle = 14
	TextSelectionNavigationWritingDirectionLeftToRight PageControllerTransitionStyle = 0
	TextSelectionNavigationWritingDirectionRightToLeft PageControllerTransitionStyle = 1
	TextSelectionNavigationLayoutOrientationHorizontal PageControllerTransitionStyle = 0
	TextSelectionNavigationLayoutOrientationVertical PageControllerTransitionStyle = 1
	TextContentManagerEnumerationOptionsNone PageControllerTransitionStyle = 0
	TextLayoutFragmentEnumerationOptionsNone PageControllerTransitionStyle = 0
	TextLayoutFragmentStateNone PageControllerTransitionStyle = 0
	TextLayoutFragmentStateEstimatedUsageBounds PageControllerTransitionStyle = 1
	TextLayoutFragmentStateCalculatedUsageBounds PageControllerTransitionStyle = 2
	TextLayoutFragmentStateLayoutAvailable PageControllerTransitionStyle = 3
	TextLayoutManagerSegmentTypeStandard PageControllerTransitionStyle = 0
	TextLayoutManagerSegmentTypeSelection PageControllerTransitionStyle = 1
	TextLayoutManagerSegmentTypeHighlight PageControllerTransitionStyle = 2
	TextLayoutManagerSegmentOptionsNone PageControllerTransitionStyle = 0
	WritingToolsCoordinatorTextUpdateReasonTyping PageControllerTransitionStyle = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo PageControllerTransitionStyle = 2
	WritingToolsCoordinatorStateInactive PageControllerTransitionStyle = 3
	WritingToolsCoordinatorStateNoninteractive PageControllerTransitionStyle = 4
	WritingToolsCoordinatorStateInteractiveResting PageControllerTransitionStyle = 5
	WritingToolsCoordinatorStateInteractiveStreaming PageControllerTransitionStyle = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive PageControllerTransitionStyle = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive PageControllerTransitionStyle = 8
	WritingToolsCoordinatorContextScopeUserSelection PageControllerTransitionStyle = 9
	WritingToolsCoordinatorContextScopeFullDocument PageControllerTransitionStyle = 10
	WritingToolsCoordinatorContextScopeVisibleArea PageControllerTransitionStyle = 11
	WritingToolsCoordinatorTextAnimationAnticipate PageControllerTransitionStyle = 12
	WritingToolsCoordinatorTextAnimationRemove PageControllerTransitionStyle = 13
	WritingToolsCoordinatorTextAnimationInsert PageControllerTransitionStyle = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive PageControllerTransitionStyle = 8
	WritingToolsCoordinatorTextAnimationTranslate PageControllerTransitionStyle = 9
)

// NSPageLayoutResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result
type PageLayoutResult uint

const (
	PageLayoutResultCancelled PageLayoutResult = 0
	PageLayoutResultChanged PageLayoutResult = 1
	PopUpNoArrow PageLayoutResult = 0
	PopUpArrowAtCenter PageLayoutResult = 1
	PopUpArrowAtBottom PageLayoutResult = 2
)

// NSLineBreakStrategy - Constants that specify how the text system breaks lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type LineBreakStrategy uint

const (
	LineBreakStrategyNone LineBreakStrategy = 0
	LineBreakStrategyPushOut LineBreakStrategy = 1
	LineBreakStrategyHangulWordPriority LineBreakStrategy = 2
	LineBreakStrategyStandard LineBreakStrategy = 3
	LeftTabStopType LineBreakStrategy = 0
	RightTabStopType LineBreakStrategy = 1
	CenterTabStopType LineBreakStrategy = 2
	DecimalTabStopType LineBreakStrategy = 3
)

// NSPasteboardAccessBehavior - A value indicating pasteboard access behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum
type PasteboardAccessBehavior uint

const (
	// PasteboardAccessBehaviorAlwaysAllow - The system will automatically allow all pasteboard access, without notifying the user.  The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysAllow
	PasteboardAccessBehaviorAlwaysAllow PasteboardAccessBehavior = 2
	// PasteboardAccessBehaviorAlwaysDeny - The system will automatically deny all pasteboard access, without notifying the user. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysDeny
	PasteboardAccessBehaviorAlwaysDeny PasteboardAccessBehavior = 3
	// PasteboardAccessBehaviorAsk - The system will notify the user and ask for permission before granting pasteboard access. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/ask
	PasteboardAccessBehaviorAsk PasteboardAccessBehavior = 1
	// PasteboardAccessBehaviorDefault - The default behavior for the General pasteboard is to ask upon programmatic access. All other pasteboards default to always allow access.   If an app has never triggered a pasteboard access alert, its General pasteboard will report   behavior. Such an app is not shown in the corresponding System Settings pane.   Once programmatic pasteboard access triggers the first pasteboard access alert, the state automatically changes to  . At this point the app starts being shown in System Settings, where the user can toggle the behavior between  ,  , and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/default
	PasteboardAccessBehaviorDefault PasteboardAccessBehavior = 0
)

// NSPasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

const (
	// PasteboardContentsCurrentHostOnly - The pasteboard contents are available only on the current device, and not on any other devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions/currentHostOnly
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
)

// NSPasteboardReadingOptions - Options that specify how to interpret data on the pasteboard when initializing pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
type PasteboardReadingOptions uint

const (
	PasteboardReadingAsData PasteboardReadingOptions = 0
	PasteboardReadingAsString PasteboardReadingOptions = 1
	PasteboardReadingAsPropertyList PasteboardReadingOptions = 1
	PasteboardReadingAsKeyedArchive PasteboardReadingOptions = 1
	ApplicationActivateAllWindows PasteboardReadingOptions = 1
	ApplicationActivateIgnoringOtherApps PasteboardReadingOptions = 2
)

// NSPasteboardWritingOptions - Type to specify options for writing to a pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type PasteboardWritingOptions uint

const (
	// PasteboardWritingPromised - Data for a type with this option is promised, not immediately written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions/promised
	PasteboardWritingPromised PasteboardWritingOptions = 1
)

// NSPickerTouchBarItemControlRepresentation - Constants that specify display styles for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum
type PickerTouchBarItemControlRepresentation uint

const (
	// PickerTouchBarItemControlRepresentationCollapsed - The system displays the control’s options through a popover.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum/collapsed
	PickerTouchBarItemControlRepresentationCollapsed PickerTouchBarItemControlRepresentation = 2
)

// NSPickerTouchBarItemSelectionMode - Constants that specify selection modes for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum
type PickerTouchBarItemSelectionMode uint

const (
	// PickerTouchBarItemSelectionModeSelectAny - A mode in which a person can select one or more options in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectAny
	PickerTouchBarItemSelectionModeSelectAny PickerTouchBarItemSelectionMode = 1
	// PickerTouchBarItemSelectionModeSelectOne - A mode in which a person can only select one option in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectOne
	PickerTouchBarItemSelectionModeSelectOne PickerTouchBarItemSelectionMode = 0
)

// NSPopoverAppearance - The set of predefined appearances for a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/Appearance-swift.enum
type PopoverAppearance uint

const (
	PopoverAppearanceMinimal PopoverAppearance = 0
	PopoverAppearanceHUD PopoverAppearance = 1
	PopoverBehaviorApplicationDefined PopoverAppearance = 0
	PopoverBehaviorTransient PopoverAppearance = 1
	PopoverBehaviorSemitransient PopoverAppearance = 2
)

// NSPopoverBehavior - The appearance and disappearance behavior of a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
type PopoverBehavior uint

const (
	// PopoverBehaviorApplicationDefined - Your application assumes responsibility for closing the popover.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum/applicationDefined
	PopoverBehaviorApplicationDefined PopoverBehavior = 0
)

// NSPrintingOrientation - Constants that specify page orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum
type PrintingOrientation uint

const (
	PortraitOrientation PrintingOrientation = 0
	LandscapeOrientation PrintingOrientation = 1
	ConfinementConcurrencyType PrintingOrientation = 2
	PrivateQueueConcurrencyType PrintingOrientation = 0
	MainQueueConcurrencyType PrintingOrientation = 0
	ChangeDone PrintingOrientation = 0
	ChangeUndone PrintingOrientation = 1
	ChangeRedone PrintingOrientation = 2
	ChangeCleared PrintingOrientation = 2
	ChangeReadOtherContents PrintingOrientation = 3
	ChangeAutosaved PrintingOrientation = 4
	ChangeDiscardable PrintingOrientation = 5
)

// NSPrintingPaginationMode - Constants that specify the different ways in which an image is divided into pages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
type PrintingPaginationMode uint

const (
	PrintingPaginationModeAutomatic PrintingPaginationMode = 0
	PrintingPaginationModeFit PrintingPaginationMode = 1
	PrintingPaginationModeClip PrintingPaginationMode = 2
)

// NSPaperOrientation - Constants that describe the orientation of printing on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation
type PaperOrientation uint

const (
	PaperOrientationPortrait PaperOrientation = 0
	PaperOrientationLandscape PaperOrientation = 1
	PrintingPaginationModeAutomatic PaperOrientation = 0
	PrintingPaginationModeFit PaperOrientation = 1
	PrintingPaginationModeClip PaperOrientation = 2
)

// NSPrintRenderingQuality - Constants that specify the print quality in use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality
type PrintRenderingQuality uint

const (
	// PrintRenderingQualityBest - Renders the printing at the best possible quality, regardless of speed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality/best
	PrintRenderingQualityBest PrintRenderingQuality = 0
	// PrintRenderingQualityResponsive - Sacrifices the least possible amount of rendering quality for speed to maintain a responsive user interface. This option should be used only after establishing that best quality rendering does indeed make the user interface unresponsive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality/responsive
	PrintRenderingQualityResponsive PrintRenderingQuality = 1
)

// NSPrintPanelOptions - Constants that specify options for configuring the contents of the main Print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
type PrintPanelOptions uint

const (
	PrintPanelShowsCopies PrintPanelOptions = 1
	PrintPanelShowsPageRange PrintPanelOptions = 1
	PrintPanelShowsPaperSize PrintPanelOptions = 1
	PrintPanelShowsOrientation PrintPanelOptions = 1
	PrintPanelShowsScaling PrintPanelOptions = 1
	PrintPanelShowsPrintSelection PrintPanelOptions = 2
	PrintPanelShowsPageSetupAccessory PrintPanelOptions = 1
	PrintPanelShowsPreview PrintPanelOptions = 1
	PDFPanelShowsPaperSize PrintPanelOptions = 1
	PDFPanelShowsOrientation PrintPanelOptions = 1
	PDFPanelRequestsParentDirectory PrintPanelOptions = 1
	MediaLibraryAudio PrintPanelOptions = 1
	MediaLibraryImage PrintPanelOptions = 1
	MediaLibraryMovie PrintPanelOptions = 1
	NoScrollerParts PrintPanelOptions = 0
	OnlyScrollerArrows PrintPanelOptions = 1
	AllScrollerParts PrintPanelOptions = 2
)

// NSPrintPanelResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result
type PrintPanelResult uint

const (
	PrintPanelResultCancelled PrintPanelResult = 0
	PrintPanelResultPrinted PrintPanelResult = 1
	PrintPanelShowsCopies PrintPanelResult = 1
	PrintPanelShowsPageRange PrintPanelResult = 1
	PrintPanelShowsPaperSize PrintPanelResult = 1
	PrintPanelShowsOrientation PrintPanelResult = 1
	PrintPanelShowsScaling PrintPanelResult = 1
	PrintPanelShowsPrintSelection PrintPanelResult = 2
	PrintPanelShowsPageSetupAccessory PrintPanelResult = 1
	PrintPanelShowsPreview PrintPanelResult = 1
	PDFPanelShowsPaperSize PrintPanelResult = 1
	PDFPanelShowsOrientation PrintPanelResult = 1
	PDFPanelRequestsParentDirectory PrintPanelResult = 1
	MediaLibraryAudio PrintPanelResult = 1
	MediaLibraryImage PrintPanelResult = 1
	MediaLibraryMovie PrintPanelResult = 1
	NoScrollerParts PrintPanelResult = 0
	OnlyScrollerArrows PrintPanelResult = 1
	AllScrollerParts PrintPanelResult = 2
)

// NSProgressIndicatorStyle - Constants that specify the progress indicator’s style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
type ProgressIndicatorStyle uint

const (
	ProgressIndicatorStyleBar ProgressIndicatorStyle = 0
	ProgressIndicatorStyleSpinning ProgressIndicatorStyle = 1
)

// NSProgressIndicatorThickness - Specify the height of a progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness
type ProgressIndicatorThickness uint

const (
	ProgressIndicatorPreferredThickness ProgressIndicatorThickness = 0
	ProgressIndicatorPreferredSmallThickness ProgressIndicatorThickness = 1
	ProgressIndicatorPreferredLargeThickness ProgressIndicatorThickness = 2
	ProgressIndicatorPreferredAquaThickness ProgressIndicatorThickness = 3
	TopTabsBezelBorder ProgressIndicatorThickness = 0
	LeftTabsBezelBorder ProgressIndicatorThickness = 1
	BottomTabsBezelBorder ProgressIndicatorThickness = 2
	RightTabsBezelBorder ProgressIndicatorThickness = 3
	NoTabsBezelBorder ProgressIndicatorThickness = 4
	NoTabsLineBorder ProgressIndicatorThickness = 5
	NoTabsNoBorder ProgressIndicatorThickness = 6
)

// NSRectAlignment - Constants that specify alignment to an edge or a set of edges depending on the user interface layout direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment
type RectAlignment uint

const (
	RectAlignmentNone RectAlignment = 0
	RectAlignmentTop RectAlignment = 1
	RectAlignmentTopLeading RectAlignment = 2
	RectAlignmentLeading RectAlignment = 3
	RectAlignmentBottomLeading RectAlignment = 4
	RectAlignmentBottom RectAlignment = 5
	RectAlignmentBottomTrailing RectAlignment = 6
	RectAlignmentTrailing RectAlignment = 7
	RectAlignmentTopTrailing RectAlignment = 8
)

// NSRuleEditorNestingMode - Specifies a type for nesting modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum
type RuleEditorNestingMode uint

const (
	RuleEditorNestingModeSingle RuleEditorNestingMode = 0
	RuleEditorNestingModeList RuleEditorNestingMode = 1
	RuleEditorNestingModeCompound RuleEditorNestingMode = 2
	RuleEditorNestingModeSimple RuleEditorNestingMode = 3
)

// NSRuleEditorRowType - Specifies a type for row types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/RowType
type RuleEditorRowType uint

const (
	RuleEditorRowTypeSimple RuleEditorRowType = 0
	RuleEditorRowTypeCompound RuleEditorRowType = 1
)

// NSRulerOrientation - These constants are defined to specify a ruler’s orientation and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum
type RulerOrientation uint

const (
	HorizontalRuler RulerOrientation = 0
	VerticalRuler RulerOrientation = 1
)

// NSScrollElasticity - These constants determine the elasticity behavior for an axis of the scrollview.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
type ScrollElasticity uint

const (
	// ScrollElasticityAllowed - Allow content to be scrolled past its bounds on this axis in an elastic fashion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/allowed
	ScrollElasticityAllowed ScrollElasticity = 2
	// ScrollElasticityAutomatic - Automatically determine whether to allow elasticity on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
	ScrollElasticityAutomatic ScrollElasticity = 0
	// ScrollElasticityNone - Disallow scrolling beyond document bounds on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/none
	ScrollElasticityNone ScrollElasticity = 1
)

// NSScrollViewFindBarPosition - These constants define the position of the find bar in relation to the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
type ScrollViewFindBarPosition uint

const (
	ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
	ScrollViewFindBarPositionAboveContent ScrollViewFindBarPosition = 1
	ScrollViewFindBarPositionBelowContent ScrollViewFindBarPosition = 2
	SegmentSwitchTrackingSelectOne ScrollViewFindBarPosition = 0
	SegmentSwitchTrackingSelectAny ScrollViewFindBarPosition = 1
	SegmentSwitchTrackingMomentary ScrollViewFindBarPosition = 2
	SegmentSwitchTrackingMomentaryAccelerator ScrollViewFindBarPosition = 3
)

// NSScrubberAlignment - The specified preferred alignment of items within the scrubber, when they come to rest following a user’s scrolling or paging interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
type ScrubberAlignment uint

const (
	// ScrubberAlignmentCenter - Center alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/center
	ScrubberAlignmentCenter ScrubberAlignment = 3
	// ScrubberAlignmentLeading - Leading alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/leading
	ScrubberAlignmentLeading ScrubberAlignment = 1
	// ScrubberAlignmentNone - No preference for item alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
	ScrubberAlignmentNone ScrubberAlignment = 0
	// ScrubberAlignmentTrailing - Trailing alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/trailing
	ScrubberAlignmentTrailing ScrubberAlignment = 2
)

// NSScrubberMode - The scrolling behavior for a scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum
type ScrubberMode uint

const (
	// ScrubberModeFixed - A scrolling mode in which scrubber items remain fixed in place, and the item under the user’s finger is highlighted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/fixed
	ScrubberModeFixed ScrubberMode = 0
	// ScrubberModeFree - A scrolling mode in which the scrubber scrolls as the user swipes horizontally across the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/free
	ScrubberModeFree ScrubberMode = 1
)

// NSSegmentSwitchTracking - The following constants specify the type of tracking behavior a segmented control exhibits. They are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
type SegmentSwitchTracking uint

const (
	// SegmentSwitchTrackingMomentary - A segment is selected only when the user is pressing the mouse down within the bounds of the segment. When the mouse is no longer down within the segment, the segment is automatically deselected. A momentary segmented control sends an action when the user clicks a segment, and another action when the user releases the segment. If configured as continuous (see  ), the control also sends actions at repeating intervals until the user releases the segment, at which point the control sends its final action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentary
	SegmentSwitchTrackingMomentary SegmentSwitchTracking = 2
)

// NSSharingCollaborationMode - Represents the types of sharing (collaborating on an item vs. sending a copy of the item)
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type SharingCollaborationMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/collaborate
	SharingCollaborationModeCollaborate SharingCollaborationMode = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/sendCopy
	SharingCollaborationModeSendCopy SharingCollaborationMode = 0
)

// NSCloudKitSharingServiceOptions - Constants that describe how a participant can configure a CloudKit share.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions
type CloudKitSharingServiceOptions uint

const (
	CloudKitSharingServiceStandard CloudKitSharingServiceOptions = 0
	CloudKitSharingServiceAllowPublic CloudKitSharingServiceOptions = 1
	CloudKitSharingServiceAllowPrivate CloudKitSharingServiceOptions = 1
	CloudKitSharingServiceAllowReadOnly CloudKitSharingServiceOptions = 1
	CloudKitSharingServiceAllowReadWrite CloudKitSharingServiceOptions = 1
	SpeechImmediateBoundary CloudKitSharingServiceOptions = 0
	SpeechWordBoundary CloudKitSharingServiceOptions = 1
	SpeechSentenceBoundary CloudKitSharingServiceOptions = 2
	CorrectionResponseNone CloudKitSharingServiceOptions = 3
	CorrectionResponseAccepted CloudKitSharingServiceOptions = 4
	CorrectionResponseRejected CloudKitSharingServiceOptions = 5
	CorrectionResponseIgnored CloudKitSharingServiceOptions = 6
	CorrectionResponseEdited CloudKitSharingServiceOptions = 7
	CorrectionResponseReverted CloudKitSharingServiceOptions = 8
)

// NSSliderType - The types of sliders, used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
type SliderType uint

const (
	SliderTypeLinear SliderType = 0
	SliderTypeCircular SliderType = 1
)

// NSSpeechBoundary - These constants are used to indicate where speech should be stopped and paused. See 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary
type SpeechBoundary uint

const (
	SpeechImmediateBoundary SpeechBoundary = 0
	SpeechWordBoundary SpeechBoundary = 1
	SpeechSentenceBoundary SpeechBoundary = 2
	CorrectionResponseNone SpeechBoundary = 3
	CorrectionResponseAccepted SpeechBoundary = 4
	CorrectionResponseRejected SpeechBoundary = 5
	CorrectionResponseIgnored SpeechBoundary = 6
	CorrectionResponseEdited SpeechBoundary = 7
	CorrectionResponseReverted SpeechBoundary = 8
)

// NSSplitViewDividerStyle - Constants that specify the style of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type SplitViewDividerStyle uint

const (
	SplitViewDividerStyleThick SplitViewDividerStyle = 1
	SplitViewDividerStyleThin SplitViewDividerStyle = 2
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
	SplitViewItemBehaviorDefault SplitViewDividerStyle = 4
	SplitViewItemBehaviorSidebar SplitViewDividerStyle = 5
	SplitViewItemBehaviorContentList SplitViewDividerStyle = 6
	SplitViewItemBehaviorInspector SplitViewDividerStyle = 7
	SplitViewItemCollapseBehaviorDefault SplitViewDividerStyle = 8
	SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings SplitViewDividerStyle = 9
	SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView SplitViewDividerStyle = 10
	SplitViewItemCollapseBehaviorUseConstraints SplitViewDividerStyle = 11
	FileHandlingPanelCancelButton SplitViewDividerStyle = 12
	FileHandlingPanelOKButton SplitViewDividerStyle = 13
)

// NSSplitViewItemBehavior - Constants that describe the behavior of the split view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum
type SplitViewItemBehavior uint

const (
	SplitViewItemBehaviorDefault SplitViewItemBehavior = 0
	SplitViewItemBehaviorSidebar SplitViewItemBehavior = 1
	SplitViewItemBehaviorContentList SplitViewItemBehavior = 2
	SplitViewItemBehaviorInspector SplitViewItemBehavior = 3
	SplitViewItemCollapseBehaviorDefault SplitViewItemBehavior = 4
	SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings SplitViewItemBehavior = 5
	SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView SplitViewItemBehavior = 6
	SplitViewItemCollapseBehaviorUseConstraints SplitViewItemBehavior = 7
	FileHandlingPanelCancelButton SplitViewItemBehavior = 8
	FileHandlingPanelOKButton SplitViewItemBehavior = 9
)

// NSSplitViewItemCollapseBehavior - Constants that describe the split view item’s collapsing behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum
type SplitViewItemCollapseBehavior uint

const (
	SplitViewItemCollapseBehaviorDefault SplitViewItemCollapseBehavior = 0
	SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings SplitViewItemCollapseBehavior = 1
	SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView SplitViewItemCollapseBehavior = 2
	SplitViewItemCollapseBehaviorUseConstraints SplitViewItemCollapseBehavior = 3
	FileHandlingPanelCancelButton SplitViewItemCollapseBehavior = 4
	FileHandlingPanelOKButton SplitViewItemCollapseBehavior = 5
)

// NSSpringLoadingHighlight - A group of constants that indicate a highlighting style for your app’s user interface to display during a spring-loading operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight
type SpringLoadingHighlight uint

const (
	SpringLoadingHighlightNone SpringLoadingHighlight = 0
	SpringLoadingHighlightStandard SpringLoadingHighlight = 1
	SpringLoadingHighlightEmphasized SpringLoadingHighlight = 2
	SpringLoadingDisabled SpringLoadingHighlight = 0
	SpringLoadingEnabled SpringLoadingHighlight = 1
	SpringLoadingContinuousActivation SpringLoadingHighlight = 1
	SpringLoadingNoHover SpringLoadingHighlight = 1
	UserInterfaceLayoutDirectionLeftToRight SpringLoadingHighlight = 0
	UserInterfaceLayoutDirectionRightToLeft SpringLoadingHighlight = 1
)

// NSStackViewDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum
type StackViewDistribution uint

const (
	StackViewDistributionGravityAreas StackViewDistribution = -1
	StackViewDistributionFill StackViewDistribution = 0
	StackViewDistributionFillEqually StackViewDistribution = 1
	StackViewDistributionFillProportionally StackViewDistribution = 2
	StackViewDistributionEqualSpacing StackViewDistribution = 3
	StackViewDistributionEqualCentering StackViewDistribution = 4
	GridCellPlacementInherited StackViewDistribution = 0
	GridCellPlacementNone StackViewDistribution = 1
	GridCellPlacementLeading StackViewDistribution = 2
	GridCellPlacementTrailing StackViewDistribution = 3
	GridCellPlacementCenter StackViewDistribution = 4
	GridCellPlacementFill StackViewDistribution = 5
	GridRowAlignmentInherited StackViewDistribution = 0
	GridRowAlignmentNone StackViewDistribution = 1
	GridRowAlignmentFirstBaseline StackViewDistribution = 2
	GridRowAlignmentLastBaseline StackViewDistribution = 3
	TextCursorAccessoryPlacementUnspecified StackViewDistribution = 4
	TextCursorAccessoryPlacementBackward StackViewDistribution = 5
	TextCursorAccessoryPlacementForward StackViewDistribution = 6
	TextCursorAccessoryPlacementInvisible StackViewDistribution = 7
	TextCursorAccessoryPlacementCenter StackViewDistribution = 8
	TextCursorAccessoryPlacementOffscreenLeft StackViewDistribution = 9
	TextCursorAccessoryPlacementOffscreenTop StackViewDistribution = 10
	TextCursorAccessoryPlacementOffscreenRight StackViewDistribution = 11
	TextCursorAccessoryPlacementOffscreenBottom StackViewDistribution = 12
	TextInputTraitTypeDefault StackViewDistribution = 13
	TextInputTraitTypeNo StackViewDistribution = 14
	TextInputTraitTypeYes StackViewDistribution = 15
)

// NSStackViewGravity - The gravity areas available in a stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
type StackViewGravity uint

const (
	StackViewGravityTop StackViewGravity = 1
	StackViewGravityLeading StackViewGravity = 1
	StackViewGravityCenter StackViewGravity = 2
	StackViewGravityBottom StackViewGravity = 3
	StackViewGravityTrailing StackViewGravity = 3
	StackViewDistributionGravityAreas StackViewGravity = -1
	StackViewDistributionFill StackViewGravity = 0
	StackViewDistributionFillEqually StackViewGravity = 1
	StackViewDistributionFillProportionally StackViewGravity = 2
	StackViewDistributionEqualSpacing StackViewGravity = 3
	StackViewDistributionEqualCentering StackViewGravity = 4
	GridCellPlacementInherited StackViewGravity = 0
	GridCellPlacementNone StackViewGravity = 1
	GridCellPlacementLeading StackViewGravity = 2
	GridCellPlacementTrailing StackViewGravity = 3
	GridCellPlacementCenter StackViewGravity = 4
	GridCellPlacementFill StackViewGravity = 5
	GridRowAlignmentInherited StackViewGravity = 0
	GridRowAlignmentNone StackViewGravity = 1
	GridRowAlignmentFirstBaseline StackViewGravity = 2
	GridRowAlignmentLastBaseline StackViewGravity = 3
	TextCursorAccessoryPlacementUnspecified StackViewGravity = 4
	TextCursorAccessoryPlacementBackward StackViewGravity = 5
	TextCursorAccessoryPlacementForward StackViewGravity = 6
	TextCursorAccessoryPlacementInvisible StackViewGravity = 7
	TextCursorAccessoryPlacementCenter StackViewGravity = 8
	TextCursorAccessoryPlacementOffscreenLeft StackViewGravity = 9
	TextCursorAccessoryPlacementOffscreenTop StackViewGravity = 10
	TextCursorAccessoryPlacementOffscreenRight StackViewGravity = 11
	TextCursorAccessoryPlacementOffscreenBottom StackViewGravity = 12
	TextInputTraitTypeDefault StackViewGravity = 13
	TextInputTraitTypeNo StackViewGravity = 14
	TextInputTraitTypeYes StackViewGravity = 15
)

// NSStatusItemBehavior - A set of optional status item behaviors.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStatusItem/Behavior-swift.struct
type StatusItemBehavior uint

const (
	DrawerClosedState StatusItemBehavior = 0
	DrawerOpeningState StatusItemBehavior = 1
	DrawerOpenState StatusItemBehavior = 2
	DrawerClosingState StatusItemBehavior = 3
)

// NSStringDrawingOptions - Constants that specify the rendering options for drawing a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions
type StringDrawingOptions uint

const (
	// StringDrawingDisableScreenFontSubstitution - An option that disables screen font substitution.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingDisableScreenFontSubstitution
	StringDrawingDisableScreenFontSubstitution StringDrawingOptions = 4
	// StringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection - Specifies the behavior for resolving   to the visual alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection
	StringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection StringDrawingOptions = 3
	// StringDrawingTruncatesLastVisibleLine - Truncates and adds the ellipsis character to the last visible line if the text doesn’t fit into the specified bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingTruncatesLastVisibleLine
	StringDrawingTruncatesLastVisibleLine StringDrawingOptions = 2
	// StringDrawingUsesDeviceMetrics - Uses image glyph bounds instead of typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingUsesDeviceMetrics
	StringDrawingUsesDeviceMetrics StringDrawingOptions = 1
	// StringDrawingUsesFontLeading - Uses the font leading for calculating line heights.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingUsesFontLeading
	StringDrawingUsesFontLeading StringDrawingOptions = 1
	// StringDrawingUsesLineFragmentOrigin - Uses the line fragment origin instead of the baseline origin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingUsesLineFragmentOrigin
	StringDrawingUsesLineFragmentOrigin StringDrawingOptions = 1
)

// NSTabPosition enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum
type TabPosition uint

const (
	TabPositionNone TabPosition = 0
	TabPositionTop TabPosition = 1
	TabPositionLeft TabPosition = 2
	TabPositionBottom TabPosition = 3
	TabPositionRight TabPosition = 4
	TabViewBorderTypeNone TabPosition = 0
	TabViewBorderTypeLine TabPosition = 1
	TabViewBorderTypeBezel TabPosition = 2
	TabViewControllerTabStyleSegmentedControlOnTop TabPosition = 0
	TabViewControllerTabStyleSegmentedControlOnBottom TabPosition = 1
	TabViewControllerTabStyleToolbar TabPosition = 2
	TabViewControllerTabStyleUnspecified TabPosition = -1
	SelectedTab TabPosition = 0
	BackgroundTab TabPosition = 1
	PressedTab TabPosition = 2
)

// NSTabViewType - These constants specify the tab view’s type as used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType
type TabViewType uint

const (
	TopTabsBezelBorder TabViewType = 0
	LeftTabsBezelBorder TabViewType = 1
	BottomTabsBezelBorder TabViewType = 2
	RightTabsBezelBorder TabViewType = 3
	NoTabsBezelBorder TabViewType = 4
	NoTabsLineBorder TabViewType = 5
	NoTabsNoBorder TabViewType = 6
)

// NSTabViewBorderType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum
type TabViewBorderType uint

const (
	TabViewBorderTypeNone TabViewBorderType = 0
	TabViewBorderTypeLine TabViewBorderType = 1
	TabViewBorderTypeBezel TabViewBorderType = 2
	TabViewControllerTabStyleSegmentedControlOnTop TabViewBorderType = 0
	TabViewControllerTabStyleSegmentedControlOnBottom TabViewBorderType = 1
	TabViewControllerTabStyleToolbar TabViewBorderType = 2
	TabViewControllerTabStyleUnspecified TabViewBorderType = -1
	SelectedTab TabViewBorderType = 0
	BackgroundTab TabViewBorderType = 1
	PressedTab TabViewBorderType = 2
)

// NSTabViewControllerTabStyle - Tab control style options for a tab view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum
type TabViewControllerTabStyle uint

const (
	TabViewControllerTabStyleSegmentedControlOnTop TabViewControllerTabStyle = 0
	TabViewControllerTabStyleSegmentedControlOnBottom TabViewControllerTabStyle = 1
	TabViewControllerTabStyleToolbar TabViewControllerTabStyle = 2
	TabViewControllerTabStyleUnspecified TabViewControllerTabStyle = -1
	SelectedTab TabViewControllerTabStyle = 0
	BackgroundTab TabViewControllerTabStyle = 1
	PressedTab TabViewControllerTabStyle = 2
)

// NSTabState - These constants describe the current display state of a tab:
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State
type TabState uint

const (
	SelectedTab TabState = 0
	BackgroundTab TabState = 1
	PressedTab TabState = 2
)

// NSTableColumnResizingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions
type TableColumnResizingOptions uint

const (
	TableColumnNoResizing TableColumnResizingOptions = 0
)

// NSTableViewAnimationOptions - Specifies the animation effects to apply when inserting or removing rows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
type TableViewAnimationOptions uint

const (
	TableViewAnimationEffectNone TableViewAnimationOptions = 0
	TableViewAnimationEffectFade TableViewAnimationOptions = 0
	TableViewAnimationEffectGap TableViewAnimationOptions = 0
	TableViewAnimationSlideUp TableViewAnimationOptions = 0
	TableViewAnimationSlideDown TableViewAnimationOptions = 0
	TableViewAnimationSlideLeft TableViewAnimationOptions = 0
	TableViewAnimationSlideRight TableViewAnimationOptions = 0
	TableColumnNoResizing TableViewAnimationOptions = 0
)

// NSTableViewColumnAutoresizingStyle - The following constants specify the autoresizing styles. These constants are used by the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum
type TableViewColumnAutoresizingStyle uint

const (
	TableViewNoColumnAutoresizing TableViewColumnAutoresizingStyle = 0
	TableViewUniformColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 1
	TableViewSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 2
	TableViewReverseSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 3
	TableViewLastColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 4
	TableViewFirstColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 5
)

// NSTableViewDraggingDestinationFeedbackStyle - These constants specify the drag styles displayed by the table view. They’re used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum
type TableViewDraggingDestinationFeedbackStyle uint

const (
	// TableViewDraggingDestinationFeedbackStyleSourceList - Draws an outline on drop target rows, and an insertion marker between rows. This style will automatically be set for source lists when the table’s   is set to  . This is the standard look for Source Lists, but may be used in other areas as needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/sourceList
	TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 1
)

// NSTableViewDropOperation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation
type TableViewDropOperation uint

const (
	// TableViewDropAbove - Specifies that the drop should occur above the specified row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation/above
	TableViewDropAbove TableViewDropOperation = 1
	// TableViewDropOn - Specifies that the drop should occur on the specified row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation/on
	TableViewDropOn TableViewDropOperation = 0
)

// NSTableViewGridLineStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle
type TableViewGridLineStyle uint

const (
	// TableViewGridNone - Specifies that no grid lines should be displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewGridLineStyle/NSTableViewGridNone
	TableViewGridNone TableViewGridLineStyle = 0
)

// NSTableRowActionEdge - These constants define table row edges on which row actions are attached. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge
type TableRowActionEdge uint

const (
	TableRowActionEdgeLeading TableRowActionEdge = 0
	TableRowActionEdgeTrailing TableRowActionEdge = 1
	TableViewAnimationEffectNone TableRowActionEdge = 0
	TableViewAnimationEffectFade TableRowActionEdge = 0
	TableViewAnimationEffectGap TableRowActionEdge = 0
	TableViewAnimationSlideUp TableRowActionEdge = 0
	TableViewAnimationSlideDown TableRowActionEdge = 0
	TableViewAnimationSlideLeft TableRowActionEdge = 0
	TableViewAnimationSlideRight TableRowActionEdge = 0
	TableColumnNoResizing TableRowActionEdge = 0
)

// NSTableViewRowSizeStyle - The row size style constants define the size of the rows in the table view. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum
type TableViewRowSizeStyle uint

const (
	// TableViewRowSizeStyleCustom - The table will use the   or invoke the delegate method  , if implemented. The cell layout is not changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/custom
	TableViewRowSizeStyleCustom TableViewRowSizeStyle = 0
	// TableViewRowSizeStyleDefault - The table will use the system default layout size: small, medium or large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/default
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
	// TableViewRowSizeStyleLarge - The table will use a row height specified for a large table. It is required that the size be fully tested and supported if   is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/large
	TableViewRowSizeStyleLarge TableViewRowSizeStyle = 3
)

// NSTableViewStyle - Contains the possible style values for a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum
type TableViewStyle uint

const (
	// TableViewStyleSourceList - The table view style resolves to a source-list style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/sourceList
	TableViewStyleSourceList TableViewStyle = 3
)

// NSTextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type TextAlignment uint

const (
	TextAlignmentLeft TextAlignment = 0
	TextAlignmentCenter TextAlignment = 1
	TextAlignmentRight TextAlignment = 2
	TextAlignmentJustified TextAlignment = 3
	TextAlignmentNatural TextAlignment = 4
	EnterCharacter TextAlignment = 0
	BackspaceCharacter TextAlignment = 0
	TabCharacter TextAlignment = 0
	NewlineCharacter TextAlignment = 0
	FormFeedCharacter TextAlignment = 0
	CarriageReturnCharacter TextAlignment = 0
	BackTabCharacter TextAlignment = 0
	DeleteCharacter TextAlignment = 0
	LineSeparatorCharacter TextAlignment = 0
	ParagraphSeparatorCharacter TextAlignment = 0
)

// NSTextCursorAccessoryPlacement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement
type TextCursorAccessoryPlacement uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/backward
	TextCursorAccessoryPlacementBackward TextCursorAccessoryPlacement = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/center
	TextCursorAccessoryPlacementCenter TextCursorAccessoryPlacement = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenLeft
	TextCursorAccessoryPlacementOffscreenLeft TextCursorAccessoryPlacement = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenTop
	TextCursorAccessoryPlacementOffscreenTop TextCursorAccessoryPlacement = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/unspecified
	TextCursorAccessoryPlacementUnspecified TextCursorAccessoryPlacement = 0
)

// NSTextInputTraitType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputTraitType
type TextInputTraitType uint

const (
	TextInputTraitTypeDefault TextInputTraitType = 0
	TextInputTraitTypeNo TextInputTraitType = 1
	TextInputTraitTypeYes TextInputTraitType = 2
)

// NSTextInsertionIndicatorAutomaticModeOptions - Options that affect the automatic display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct
type TextInsertionIndicatorAutomaticModeOptions uint

const (
	// TextInsertionIndicatorAutomaticModeOptionsShowEffectsView - Specifies whether a trailing glow displays during dictation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct/showEffectsView
	TextInsertionIndicatorAutomaticModeOptionsShowEffectsView TextInsertionIndicatorAutomaticModeOptions = 1
	// TextInsertionIndicatorAutomaticModeOptionsShowWhileTracking - Specifies whether the insertion indicator shows during a tracking loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct/showWhileTracking
	TextInsertionIndicatorAutomaticModeOptionsShowWhileTracking TextInsertionIndicatorAutomaticModeOptions = 1
)

// NSTextInsertionIndicatorDisplayMode - Constants that determine how to display the system text cursor in a custom text UI.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum
type TextInsertionIndicatorDisplayMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/automatic
	TextInsertionIndicatorDisplayModeAutomatic TextInsertionIndicatorDisplayMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/hidden
	TextInsertionIndicatorDisplayModeHidden TextInsertionIndicatorDisplayMode = 1
)

// NSTextLayoutManagerSegmentOptions - Values that describe where and how the framework extends segments of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions
type TextLayoutManagerSegmentOptions uint

const (
	TextLayoutManagerSegmentOptionsNone TextLayoutManagerSegmentOptions = 0
	WritingToolsCoordinatorTextUpdateReasonTyping TextLayoutManagerSegmentOptions = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo TextLayoutManagerSegmentOptions = 2
	WritingToolsCoordinatorStateInactive TextLayoutManagerSegmentOptions = 3
	WritingToolsCoordinatorStateNoninteractive TextLayoutManagerSegmentOptions = 4
	WritingToolsCoordinatorStateInteractiveResting TextLayoutManagerSegmentOptions = 5
	WritingToolsCoordinatorStateInteractiveStreaming TextLayoutManagerSegmentOptions = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive TextLayoutManagerSegmentOptions = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive TextLayoutManagerSegmentOptions = 8
	WritingToolsCoordinatorContextScopeUserSelection TextLayoutManagerSegmentOptions = 9
	WritingToolsCoordinatorContextScopeFullDocument TextLayoutManagerSegmentOptions = 10
	WritingToolsCoordinatorContextScopeVisibleArea TextLayoutManagerSegmentOptions = 11
	WritingToolsCoordinatorTextAnimationAnticipate TextLayoutManagerSegmentOptions = 12
	WritingToolsCoordinatorTextAnimationRemove TextLayoutManagerSegmentOptions = 13
	WritingToolsCoordinatorTextAnimationInsert TextLayoutManagerSegmentOptions = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive TextLayoutManagerSegmentOptions = 8
	WritingToolsCoordinatorTextAnimationTranslate TextLayoutManagerSegmentOptions = 9
)

// NSTextLayoutManagerSegmentType - Values that describe the rendering of selection boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType
type TextLayoutManagerSegmentType uint

const (
	TextLayoutManagerSegmentTypeStandard TextLayoutManagerSegmentType = 0
	TextLayoutManagerSegmentTypeSelection TextLayoutManagerSegmentType = 1
	TextLayoutManagerSegmentTypeHighlight TextLayoutManagerSegmentType = 2
	TextLayoutManagerSegmentOptionsNone TextLayoutManagerSegmentType = 0
	WritingToolsCoordinatorTextUpdateReasonTyping TextLayoutManagerSegmentType = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo TextLayoutManagerSegmentType = 2
	WritingToolsCoordinatorStateInactive TextLayoutManagerSegmentType = 3
	WritingToolsCoordinatorStateNoninteractive TextLayoutManagerSegmentType = 4
	WritingToolsCoordinatorStateInteractiveResting TextLayoutManagerSegmentType = 5
	WritingToolsCoordinatorStateInteractiveStreaming TextLayoutManagerSegmentType = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive TextLayoutManagerSegmentType = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive TextLayoutManagerSegmentType = 8
	WritingToolsCoordinatorContextScopeUserSelection TextLayoutManagerSegmentType = 9
	WritingToolsCoordinatorContextScopeFullDocument TextLayoutManagerSegmentType = 10
	WritingToolsCoordinatorContextScopeVisibleArea TextLayoutManagerSegmentType = 11
	WritingToolsCoordinatorTextAnimationAnticipate TextLayoutManagerSegmentType = 12
	WritingToolsCoordinatorTextAnimationRemove TextLayoutManagerSegmentType = 13
	WritingToolsCoordinatorTextAnimationInsert TextLayoutManagerSegmentType = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive TextLayoutManagerSegmentType = 8
	WritingToolsCoordinatorTextAnimationTranslate TextLayoutManagerSegmentType = 9
)

// NSTextMovement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement
type TextMovement uint

const (
	TextMovementReturn TextMovement = 0
	TextMovementTab TextMovement = 0
	TextMovementBacktab TextMovement = 0
	TextMovementLeft TextMovement = 0
	TextMovementRight TextMovement = 0
	TextMovementUp TextMovement = 0
	TextMovementDown TextMovement = 0
	TextMovementCancel TextMovement = 0
	TextMovementOther TextMovement = 0
)

// NSTintProminence - Controls how strongly the tint color applies in a view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence
type TintProminence uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/automatic
	TintProminenceAutomatic TintProminence = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/none
	TintProminenceNone TintProminence = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/primary
	TintProminencePrimary TintProminence = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/secondary
	TintProminenceSecondary TintProminence = 3
)

// NSTitlebarSeparatorStyle - Styles that determine the type of separator displayed between the title bar and content of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle
type TitlebarSeparatorStyle uint

const (
	// TitlebarSeparatorStyleAutomatic - A style indicating that the system determines the type of separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/automatic
	TitlebarSeparatorStyleAutomatic TitlebarSeparatorStyle = 0
	// TitlebarSeparatorStyleLine - A style indicating that the title bar separator is a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/line
	TitlebarSeparatorStyleLine TitlebarSeparatorStyle = 2
	// TitlebarSeparatorStyleNone - A style indicating that there’s no title bar separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/none
	TitlebarSeparatorStyleNone TitlebarSeparatorStyle = 1
	// TitlebarSeparatorStyleShadow - A style indicating that the title bar separator is a shadow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/shadow
	TitlebarSeparatorStyleShadow TitlebarSeparatorStyle = 3
)

// NSTokenStyle - The NSTokenStyle constants define how tokens are displayed and editable in the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
type TokenStyle uint

const (
	TokenStyleDefault TokenStyle = 0
	TokenStyleNone TokenStyle = 1
	TokenStyleRounded TokenStyle = 2
	TokenStyleSquared TokenStyle = 3
	TokenStylePlainSquared TokenStyle = 4
)

// NSToolbarDisplayMode - Constants that indicate whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum
type ToolbarDisplayMode uint

const (
	// ToolbarDisplayModeDefault - The default display mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/default
	ToolbarDisplayModeDefault ToolbarDisplayMode = 0
	// ToolbarDisplayModeIconAndLabel - The toolbar displays an icon and label for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/iconAndLabel
	ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
	// ToolbarDisplayModeIconOnly - The toolbar displays only an icon for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/iconOnly
	ToolbarDisplayModeIconOnly ToolbarDisplayMode = 2
	// ToolbarDisplayModeLabelOnly - The toolbar displays only a label for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/labelOnly
	ToolbarDisplayModeLabelOnly ToolbarDisplayMode = 3
)

// NSToolbarSizeMode - Constants that specify toolbar display modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum
type ToolbarSizeMode uint

const (
	// ToolbarSizeModeDefault - The toolbar uses the system-defined default size, which is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/default
	ToolbarSizeModeDefault ToolbarSizeMode = 0
	// ToolbarSizeModeRegular - The toolbar uses regular-sized controls and 32 by 32 pixel icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/regular
	ToolbarSizeModeRegular ToolbarSizeMode = 1
	// ToolbarSizeModeSmall - The toolbar uses small-sized controls and 24 by 24 pixel icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/small
	ToolbarSizeModeSmall ToolbarSizeMode = 2
)

// NSToolbarItemStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum
type ToolbarItemStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum/plain
	ToolbarItemStylePlain ToolbarItemStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum/prominent
	ToolbarItemStyleProminent ToolbarItemStyle = 1
)

// NSToolbarItemGroupControlRepresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum
type ToolbarItemGroupControlRepresentation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum/collapsed
	ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 2
)

// NSToolbarItemGroupSelectionMode - A value that indicates how a grouped toolbar item selects its subitems.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum
type ToolbarItemGroupSelectionMode uint

const (
	// ToolbarItemGroupSelectionModeMomentary - The system temporarily highlights the select group item when the user selects the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/momentary
	ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 2
	// ToolbarItemGroupSelectionModeSelectAny - The system toggles a highlight on any item selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectAny
	ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 1
	// ToolbarItemGroupSelectionModeSelectOne - The system displays a highlighted mode on the most recent item selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectOne
	ToolbarItemGroupSelectionModeSelectOne ToolbarItemGroupSelectionMode = 0
)

// NSTouchType - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType
type TouchType uint

const (
	// TouchTypeDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType/direct
	TouchTypeDirect TouchType = 0
)

// NSTouchTypeMask - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask
type TouchTypeMask uint

const (
	// TouchTypeMaskDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask/direct
	TouchTypeMaskDirect TouchTypeMask = 0
)

// NSTrackingAreaOptions - The data type defined for the constants specified in the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct
type TrackingAreaOptions uint

const (
	// TrackingActiveAlways - The owner receives messages regardless of first-responder status, window status, or application status. The   message is   sent when the   option is specified along with this constant. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeAlways
	TrackingActiveAlways TrackingAreaOptions = 0
	// TrackingActiveInActiveApp - The owner receives messages when the application is active. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeInActiveApp
	TrackingActiveInActiveApp TrackingAreaOptions = 0
	// TrackingActiveInKeyWindow - The owner receives messages when the view is in the key window. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeInKeyWindow
	TrackingActiveInKeyWindow TrackingAreaOptions = 0
	// TrackingActiveWhenFirstResponder - The owner receives messages when the view is the first responder. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeWhenFirstResponder
	TrackingActiveWhenFirstResponder TrackingAreaOptions = 0
	// TrackingAssumeInside - The first event is generated when the cursor leaves the tracking area, regardless if the cursor is inside the area when the   is added to a view.  If this option is not specified, the first event is generated when the cursor leaves the tracking area if the cursor is initially inside the area, or when the cursor enters the area if the cursor is initially outside it. Generally, you do not want to request this behavior. This value specifies a behavior of the tracking area defined by the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/assumeInside
	TrackingAssumeInside TrackingAreaOptions = 0
	// TrackingCursorUpdate - A tracking option that receives events when the mouse cursor enters and exits the tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/cursorUpdate
	TrackingCursorUpdate TrackingAreaOptions = 0
	// TrackingEnabledDuringMouseDrag - The owner receives   events when the mouse cursor is dragged into the tracking area. If this option is not specified, the owner receives mouse-entered events when the mouse is moved (no buttons pressed) into the tracking area and on   events after a mouse drag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/enabledDuringMouseDrag
	TrackingEnabledDuringMouseDrag TrackingAreaOptions = 0
	// TrackingInVisibleRect - Mouse tracking occurs only in the visible rectangle of the view—in other words, that region of the tracking rectangle that is unobscured. Otherwise, the entire tracking area is active regardless of overlapping views. The   object is automatically synchronized with changes in the view’s visible area ( ) and the value returned from   is ignored. This value specifies a behavior of the tracking area defined by the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/inVisibleRect
	TrackingInVisibleRect TrackingAreaOptions = 0
	// TrackingMouseEnteredAndExited - The owner of the tracking area receives   when the mouse cursor enters the area and   events when the mouse leaves the area. This value specifies a type of tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/mouseEnteredAndExited
	TrackingMouseEnteredAndExited TrackingAreaOptions = 0
	// TrackingMouseMoved - The owner of the tracking area receives   messages while the mouse cursor is within the area. This value specifies a type of tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/mouseMoved
	TrackingMouseMoved TrackingAreaOptions = 0
)

// NSUnderlineStyle - Constants for the underline style and strikethrough style attribute keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle
type UnderlineStyle uint

const (
	// UnderlineStyleNone - Don’t draw a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/NSUnderlineStyleNone
	UnderlineStyleNone UnderlineStyle = 0
	// UnderlineStylePatternSolid - Draw a solid line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/NSUnderlineStylePatternSolid
	UnderlineStylePatternSolid UnderlineStyle = 3
	// UnderlineStyleByWord - Draw the line only beneath or through words, not whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/byWord
	UnderlineStyleByWord UnderlineStyle = 8
	// UnderlineStyleDouble - Draw a double line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/double
	UnderlineStyleDouble UnderlineStyle = 2
	// UnderlineStylePatternDash - Draw a line of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDash
	UnderlineStylePatternDash UnderlineStyle = 5
	// UnderlineStylePatternDashDot - Draw a line of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDot
	UnderlineStylePatternDashDot UnderlineStyle = 6
	// UnderlineStylePatternDashDotDot - Draw a line of alternating dashes and two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDotDot
	UnderlineStylePatternDashDotDot UnderlineStyle = 7
	// UnderlineStylePatternDot - Draw a line of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDot
	UnderlineStylePatternDot UnderlineStyle = 4
	// UnderlineStyleSingle - Draw a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/single
	UnderlineStyleSingle UnderlineStyle = 0
	// UnderlineStyleThick - Draw a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/thick
	UnderlineStyleThick UnderlineStyle = 1
)

// NSUserInterfaceLayoutDirection - Specifies the directional flow of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
type UserInterfaceLayoutDirection uint

const (
	// UserInterfaceLayoutDirectionRightToLeft - Layout direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/rightToLeft
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

// NSUserInterfaceLayoutOrientation - The stack view layout directions, and user interface axes for hugging priority and clipping resistance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
type UserInterfaceLayoutOrientation uint

const (
	// UserInterfaceLayoutOrientationVertical - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation/vertical
	UserInterfaceLayoutOrientationVertical UserInterfaceLayoutOrientation = 1
)

// NSVerticalDirections - The directions on the vertical axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections
type VerticalDirections uint

const (
	// VerticalDirectionsAll - All vertical directions (up and down).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsAll
	VerticalDirectionsAll VerticalDirections = 0
)

// NSAutoresizingMaskOptions - Constants that specify the autoresizing behaviors for views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct
type AutoresizingMaskOptions uint

const (
	ViewNotSizable AutoresizingMaskOptions = 0
	ViewMinXMargin AutoresizingMaskOptions = 1
	ViewWidthSizable AutoresizingMaskOptions = 2
	ViewMaxXMargin AutoresizingMaskOptions = 4
	ViewMinYMargin AutoresizingMaskOptions = 8
	ViewHeightSizable AutoresizingMaskOptions = 16
	ViewMaxYMargin AutoresizingMaskOptions = 32
)

// NSBackgroundStyle - Background styles to apply to a view’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
type BackgroundStyle uint

const (
	BackgroundStyleNormal BackgroundStyle = 0
	BackgroundStyleEmphasized BackgroundStyle = 1
	BackgroundStyleRaised BackgroundStyle = 2
	BackgroundStyleLowered BackgroundStyle = 3
	AnyType BackgroundStyle = 4
	IntType BackgroundStyle = 5
	PositiveIntType BackgroundStyle = 6
	FloatType BackgroundStyle = 7
	PositiveFloatType BackgroundStyle = 8
	DoubleType BackgroundStyle = 9
	PositiveDoubleType BackgroundStyle = 10
)

// NSViewLayerContentsPlacement - These constants specify the location of the layer content when the content is not rerendered in response to view resizing. For more information, see the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum
type ViewLayerContentsPlacement uint

const (
	ViewLayerContentsPlacementScaleAxesIndependently ViewLayerContentsPlacement = 0
	ViewLayerContentsPlacementScaleProportionallyToFit ViewLayerContentsPlacement = 1
	ViewLayerContentsPlacementScaleProportionallyToFill ViewLayerContentsPlacement = 2
	ViewLayerContentsPlacementCenter ViewLayerContentsPlacement = 3
	ViewLayerContentsPlacementTop ViewLayerContentsPlacement = 4
	ViewLayerContentsPlacementTopRight ViewLayerContentsPlacement = 5
	ViewLayerContentsPlacementRight ViewLayerContentsPlacement = 6
	ViewLayerContentsPlacementBottomRight ViewLayerContentsPlacement = 7
	ViewLayerContentsPlacementBottom ViewLayerContentsPlacement = 8
	ViewLayerContentsPlacementBottomLeft ViewLayerContentsPlacement = 9
	ViewLayerContentsPlacementLeft ViewLayerContentsPlacement = 10
	ViewLayerContentsPlacementTopLeft ViewLayerContentsPlacement = 11
	WritingDirectionNatural ViewLayerContentsPlacement = -1
	WritingDirectionLeftToRight ViewLayerContentsPlacement = 0
	WritingDirectionRightToLeft ViewLayerContentsPlacement = 1
	TextAlignmentLeft ViewLayerContentsPlacement = 0
	TextAlignmentCenter ViewLayerContentsPlacement = 1
	TextAlignmentRight ViewLayerContentsPlacement = 2
	TextAlignmentJustified ViewLayerContentsPlacement = 3
	TextAlignmentNatural ViewLayerContentsPlacement = 4
	EnterCharacter ViewLayerContentsPlacement = 0
	BackspaceCharacter ViewLayerContentsPlacement = 0
	TabCharacter ViewLayerContentsPlacement = 0
	NewlineCharacter ViewLayerContentsPlacement = 0
	FormFeedCharacter ViewLayerContentsPlacement = 0
	CarriageReturnCharacter ViewLayerContentsPlacement = 0
	BackTabCharacter ViewLayerContentsPlacement = 0
	DeleteCharacter ViewLayerContentsPlacement = 0
	LineSeparatorCharacter ViewLayerContentsPlacement = 0
	ParagraphSeparatorCharacter ViewLayerContentsPlacement = 0
)

// NSViewLayerContentsRedrawPolicy - Constants that specify how layer resizing is handled when a view is layer-backed or layer-hosting. For more information, see the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum
type ViewLayerContentsRedrawPolicy uint

const (
	ViewLayerContentsRedrawNever ViewLayerContentsRedrawPolicy = 0
	ViewLayerContentsRedrawOnSetNeedsDisplay ViewLayerContentsRedrawPolicy = 1
	ViewLayerContentsRedrawDuringViewResize ViewLayerContentsRedrawPolicy = 2
	ViewLayerContentsRedrawBeforeViewResize ViewLayerContentsRedrawPolicy = 3
	ViewLayerContentsRedrawCrossfade ViewLayerContentsRedrawPolicy = 4
	ViewLayerContentsPlacementScaleAxesIndependently ViewLayerContentsRedrawPolicy = 0
	ViewLayerContentsPlacementScaleProportionallyToFit ViewLayerContentsRedrawPolicy = 1
	ViewLayerContentsPlacementScaleProportionallyToFill ViewLayerContentsRedrawPolicy = 2
	ViewLayerContentsPlacementCenter ViewLayerContentsRedrawPolicy = 3
	ViewLayerContentsPlacementTop ViewLayerContentsRedrawPolicy = 4
	ViewLayerContentsPlacementTopRight ViewLayerContentsRedrawPolicy = 5
	ViewLayerContentsPlacementRight ViewLayerContentsRedrawPolicy = 6
	ViewLayerContentsPlacementBottomRight ViewLayerContentsRedrawPolicy = 7
	ViewLayerContentsPlacementBottom ViewLayerContentsRedrawPolicy = 8
	ViewLayerContentsPlacementBottomLeft ViewLayerContentsRedrawPolicy = 9
	ViewLayerContentsPlacementLeft ViewLayerContentsRedrawPolicy = 10
	ViewLayerContentsPlacementTopLeft ViewLayerContentsRedrawPolicy = 11
	WritingDirectionNatural ViewLayerContentsRedrawPolicy = -1
	WritingDirectionLeftToRight ViewLayerContentsRedrawPolicy = 0
	WritingDirectionRightToLeft ViewLayerContentsRedrawPolicy = 1
	TextAlignmentLeft ViewLayerContentsRedrawPolicy = 0
	TextAlignmentCenter ViewLayerContentsRedrawPolicy = 1
	TextAlignmentRight ViewLayerContentsRedrawPolicy = 2
	TextAlignmentJustified ViewLayerContentsRedrawPolicy = 3
	TextAlignmentNatural ViewLayerContentsRedrawPolicy = 4
	EnterCharacter ViewLayerContentsRedrawPolicy = 0
	BackspaceCharacter ViewLayerContentsRedrawPolicy = 0
	TabCharacter ViewLayerContentsRedrawPolicy = 0
	NewlineCharacter ViewLayerContentsRedrawPolicy = 0
	FormFeedCharacter ViewLayerContentsRedrawPolicy = 0
	CarriageReturnCharacter ViewLayerContentsRedrawPolicy = 0
	BackTabCharacter ViewLayerContentsRedrawPolicy = 0
	DeleteCharacter ViewLayerContentsRedrawPolicy = 0
	LineSeparatorCharacter ViewLayerContentsRedrawPolicy = 0
	ParagraphSeparatorCharacter ViewLayerContentsRedrawPolicy = 0
)

// NSViewControllerTransitionOptions - Animation options for view transitions in a view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
type ViewControllerTransitionOptions uint

const (
	// ViewControllerTransitionAllowUserInteraction - A transition animation that allows user interaction during the transition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/allowUserInteraction
	ViewControllerTransitionAllowUserInteraction ViewControllerTransitionOptions = 0
	// ViewControllerTransitionCrossfade - A transition animation that fades the new view in and simultaneously fades the old view out. You can combine this animation option with any of the “slide” options in this enumeration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/crossfade
	ViewControllerTransitionCrossfade ViewControllerTransitionOptions = 0
	// ViewControllerTransitionSlideBackward - A transition animation that reflects the user interface layout direction ( ) in a “backward” manner, as follows
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideBackward
	ViewControllerTransitionSlideBackward ViewControllerTransitionOptions = 0
	// ViewControllerTransitionSlideRight - A transition animation that slides the old view to the right while the new view slides into view from the left.  In other words, both views slide to the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideRight
	ViewControllerTransitionSlideRight ViewControllerTransitionOptions = 0
	// ViewControllerTransitionSlideUp - A transition animation that slides the old view up while the new view comes into view from the bottom.  In other words, both views slide up.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideUp
	ViewControllerTransitionSlideUp ViewControllerTransitionOptions = 0
)

// NSViewLayoutRegionAdaptivityAxis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis
type ViewLayoutRegionAdaptivityAxis uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisHorizontal
	ViewLayoutRegionAdaptivityAxisHorizontal ViewLayoutRegionAdaptivityAxis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisNone
	ViewLayoutRegionAdaptivityAxisNone ViewLayoutRegionAdaptivityAxis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisVertical
	ViewLayoutRegionAdaptivityAxisVertical ViewLayoutRegionAdaptivityAxis = 0
)

// NSVisualEffectBlendingMode - Constants that specify whether the visual effect view blends with what’s either behind or within the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum
type VisualEffectBlendingMode uint

const (
	VisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = 0
	VisualEffectBlendingModeWithinWindow VisualEffectBlendingMode = 1
	VisualEffectStateFollowsWindowActiveState VisualEffectBlendingMode = 2
	VisualEffectStateActive VisualEffectBlendingMode = 3
	VisualEffectStateInactive VisualEffectBlendingMode = 4
	HapticFeedbackPatternGeneric VisualEffectBlendingMode = 0
	HapticFeedbackPatternAlignment VisualEffectBlendingMode = 1
	HapticFeedbackPatternLevelChange VisualEffectBlendingMode = 2
	HapticFeedbackPerformanceTimeDefault VisualEffectBlendingMode = 0
	HapticFeedbackPerformanceTimeNow VisualEffectBlendingMode = 1
	HapticFeedbackPerformanceTimeDrawCompleted VisualEffectBlendingMode = 2
	PickerTouchBarItemSelectionModeSelectOne VisualEffectBlendingMode = 0
	PickerTouchBarItemSelectionModeSelectAny VisualEffectBlendingMode = 1
	PickerTouchBarItemSelectionModeMomentary VisualEffectBlendingMode = 2
	PickerTouchBarItemControlRepresentationAutomatic VisualEffectBlendingMode = 0
	PickerTouchBarItemControlRepresentationExpanded VisualEffectBlendingMode = 1
	PickerTouchBarItemControlRepresentationCollapsed VisualEffectBlendingMode = 2
	TextSelectionGranularityCharacter VisualEffectBlendingMode = 3
	TextSelectionGranularityWord VisualEffectBlendingMode = 4
	TextSelectionGranularityParagraph VisualEffectBlendingMode = 5
	TextSelectionGranularityLine VisualEffectBlendingMode = 6
	TextSelectionGranularitySentence VisualEffectBlendingMode = 7
	TextSelectionAffinityUpstream VisualEffectBlendingMode = 0
	TextSelectionAffinityDownstream VisualEffectBlendingMode = 1
	TextSelectionNavigationDirectionForward VisualEffectBlendingMode = 2
	TextSelectionNavigationDirectionBackward VisualEffectBlendingMode = 3
	TextSelectionNavigationDirectionRight VisualEffectBlendingMode = 4
	TextSelectionNavigationDirectionLeft VisualEffectBlendingMode = 5
	TextSelectionNavigationDirectionUp VisualEffectBlendingMode = 6
	TextSelectionNavigationDirectionDown VisualEffectBlendingMode = 7
	TextSelectionNavigationDestinationCharacter VisualEffectBlendingMode = 8
	TextSelectionNavigationDestinationWord VisualEffectBlendingMode = 9
	TextSelectionNavigationDestinationLine VisualEffectBlendingMode = 10
	TextSelectionNavigationDestinationSentence VisualEffectBlendingMode = 11
	TextSelectionNavigationDestinationParagraph VisualEffectBlendingMode = 12
	TextSelectionNavigationDestinationContainer VisualEffectBlendingMode = 13
	TextSelectionNavigationDestinationDocument VisualEffectBlendingMode = 14
	TextSelectionNavigationWritingDirectionLeftToRight VisualEffectBlendingMode = 0
	TextSelectionNavigationWritingDirectionRightToLeft VisualEffectBlendingMode = 1
	TextSelectionNavigationLayoutOrientationHorizontal VisualEffectBlendingMode = 0
	TextSelectionNavigationLayoutOrientationVertical VisualEffectBlendingMode = 1
	TextContentManagerEnumerationOptionsNone VisualEffectBlendingMode = 0
	TextLayoutFragmentEnumerationOptionsNone VisualEffectBlendingMode = 0
	TextLayoutFragmentStateNone VisualEffectBlendingMode = 0
	TextLayoutFragmentStateEstimatedUsageBounds VisualEffectBlendingMode = 1
	TextLayoutFragmentStateCalculatedUsageBounds VisualEffectBlendingMode = 2
	TextLayoutFragmentStateLayoutAvailable VisualEffectBlendingMode = 3
	TextLayoutManagerSegmentTypeStandard VisualEffectBlendingMode = 0
	TextLayoutManagerSegmentTypeSelection VisualEffectBlendingMode = 1
	TextLayoutManagerSegmentTypeHighlight VisualEffectBlendingMode = 2
	TextLayoutManagerSegmentOptionsNone VisualEffectBlendingMode = 0
	WritingToolsCoordinatorTextUpdateReasonTyping VisualEffectBlendingMode = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo VisualEffectBlendingMode = 2
	WritingToolsCoordinatorStateInactive VisualEffectBlendingMode = 3
	WritingToolsCoordinatorStateNoninteractive VisualEffectBlendingMode = 4
	WritingToolsCoordinatorStateInteractiveResting VisualEffectBlendingMode = 5
	WritingToolsCoordinatorStateInteractiveStreaming VisualEffectBlendingMode = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive VisualEffectBlendingMode = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive VisualEffectBlendingMode = 8
	WritingToolsCoordinatorContextScopeUserSelection VisualEffectBlendingMode = 9
	WritingToolsCoordinatorContextScopeFullDocument VisualEffectBlendingMode = 10
	WritingToolsCoordinatorContextScopeVisibleArea VisualEffectBlendingMode = 11
	WritingToolsCoordinatorTextAnimationAnticipate VisualEffectBlendingMode = 12
	WritingToolsCoordinatorTextAnimationRemove VisualEffectBlendingMode = 13
	WritingToolsCoordinatorTextAnimationInsert VisualEffectBlendingMode = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive VisualEffectBlendingMode = 8
	WritingToolsCoordinatorTextAnimationTranslate VisualEffectBlendingMode = 9
)

// NSVisualEffectMaterial - Constants to specify the material shown by the visual effect view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum
type VisualEffectMaterial uint

const (
	// VisualEffectMaterialLight - A material with a light effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/light
	VisualEffectMaterialLight VisualEffectMaterial = 18
	// VisualEffectMaterialSidebar - The material for the background of window sidebars.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/sidebar
	VisualEffectMaterialSidebar VisualEffectMaterial = 7
)

// NSVisualEffectState - Constants to specify how the material appearance should reflect window activity state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum
type VisualEffectState uint

const (
	VisualEffectStateFollowsWindowActiveState VisualEffectState = 0
	VisualEffectStateActive VisualEffectState = 1
	VisualEffectStateInactive VisualEffectState = 2
	HapticFeedbackPatternGeneric VisualEffectState = 0
	HapticFeedbackPatternAlignment VisualEffectState = 1
	HapticFeedbackPatternLevelChange VisualEffectState = 2
	HapticFeedbackPerformanceTimeDefault VisualEffectState = 0
	HapticFeedbackPerformanceTimeNow VisualEffectState = 1
	HapticFeedbackPerformanceTimeDrawCompleted VisualEffectState = 2
	PickerTouchBarItemSelectionModeSelectOne VisualEffectState = 0
	PickerTouchBarItemSelectionModeSelectAny VisualEffectState = 1
	PickerTouchBarItemSelectionModeMomentary VisualEffectState = 2
	PickerTouchBarItemControlRepresentationAutomatic VisualEffectState = 0
	PickerTouchBarItemControlRepresentationExpanded VisualEffectState = 1
	PickerTouchBarItemControlRepresentationCollapsed VisualEffectState = 2
	TextSelectionGranularityCharacter VisualEffectState = 3
	TextSelectionGranularityWord VisualEffectState = 4
	TextSelectionGranularityParagraph VisualEffectState = 5
	TextSelectionGranularityLine VisualEffectState = 6
	TextSelectionGranularitySentence VisualEffectState = 7
	TextSelectionAffinityUpstream VisualEffectState = 0
	TextSelectionAffinityDownstream VisualEffectState = 1
	TextSelectionNavigationDirectionForward VisualEffectState = 2
	TextSelectionNavigationDirectionBackward VisualEffectState = 3
	TextSelectionNavigationDirectionRight VisualEffectState = 4
	TextSelectionNavigationDirectionLeft VisualEffectState = 5
	TextSelectionNavigationDirectionUp VisualEffectState = 6
	TextSelectionNavigationDirectionDown VisualEffectState = 7
	TextSelectionNavigationDestinationCharacter VisualEffectState = 8
	TextSelectionNavigationDestinationWord VisualEffectState = 9
	TextSelectionNavigationDestinationLine VisualEffectState = 10
	TextSelectionNavigationDestinationSentence VisualEffectState = 11
	TextSelectionNavigationDestinationParagraph VisualEffectState = 12
	TextSelectionNavigationDestinationContainer VisualEffectState = 13
	TextSelectionNavigationDestinationDocument VisualEffectState = 14
	TextSelectionNavigationWritingDirectionLeftToRight VisualEffectState = 0
	TextSelectionNavigationWritingDirectionRightToLeft VisualEffectState = 1
	TextSelectionNavigationLayoutOrientationHorizontal VisualEffectState = 0
	TextSelectionNavigationLayoutOrientationVertical VisualEffectState = 1
	TextContentManagerEnumerationOptionsNone VisualEffectState = 0
	TextLayoutFragmentEnumerationOptionsNone VisualEffectState = 0
	TextLayoutFragmentStateNone VisualEffectState = 0
	TextLayoutFragmentStateEstimatedUsageBounds VisualEffectState = 1
	TextLayoutFragmentStateCalculatedUsageBounds VisualEffectState = 2
	TextLayoutFragmentStateLayoutAvailable VisualEffectState = 3
	TextLayoutManagerSegmentTypeStandard VisualEffectState = 0
	TextLayoutManagerSegmentTypeSelection VisualEffectState = 1
	TextLayoutManagerSegmentTypeHighlight VisualEffectState = 2
	TextLayoutManagerSegmentOptionsNone VisualEffectState = 0
	WritingToolsCoordinatorTextUpdateReasonTyping VisualEffectState = 1
	WritingToolsCoordinatorTextUpdateReasonUndoRedo VisualEffectState = 2
	WritingToolsCoordinatorStateInactive VisualEffectState = 3
	WritingToolsCoordinatorStateNoninteractive VisualEffectState = 4
	WritingToolsCoordinatorStateInteractiveResting VisualEffectState = 5
	WritingToolsCoordinatorStateInteractiveStreaming VisualEffectState = 6
	WritingToolsCoordinatorTextReplacementReasonInteractive VisualEffectState = 7
	WritingToolsCoordinatorTextReplacementReasonNoninteractive VisualEffectState = 8
	WritingToolsCoordinatorContextScopeUserSelection VisualEffectState = 9
	WritingToolsCoordinatorContextScopeFullDocument VisualEffectState = 10
	WritingToolsCoordinatorContextScopeVisibleArea VisualEffectState = 11
	WritingToolsCoordinatorTextAnimationAnticipate VisualEffectState = 12
	WritingToolsCoordinatorTextAnimationRemove VisualEffectState = 13
	WritingToolsCoordinatorTextAnimationInsert VisualEffectState = 14
	WritingToolsCoordinatorTextAnimationAnticipateInactive VisualEffectState = 8
	WritingToolsCoordinatorTextAnimationTranslate VisualEffectState = 9
)

// NSWindowAnimationBehavior - Constants that control the automatic window animation behavior windows use when ordering to the front or out of view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
type WindowAnimationBehavior uint

const (
	// WindowAnimationBehaviorAlertPanel - The animation behavior that’s appropriate to the alert window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/alertPanel
	WindowAnimationBehaviorAlertPanel WindowAnimationBehavior = 5
	// WindowAnimationBehaviorDefault - The automatic animation that’s appropriate to the window type. This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/default
	WindowAnimationBehaviorDefault WindowAnimationBehavior = 0
	// WindowAnimationBehaviorDocumentWindow - The animation behavior that’s appropriate to the document window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/documentWindow
	WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 3
	// WindowAnimationBehaviorNone - No automatic animation used. This may be useful when you perform your own window animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/none
	WindowAnimationBehaviorNone WindowAnimationBehavior = 2
	// WindowAnimationBehaviorUtilityWindow - The animation behavior that’s appropriate to the utility window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/utilityWindow
	WindowAnimationBehaviorUtilityWindow WindowAnimationBehavior = 4
)

// NSWindowBackingLocation - The following constants and the related data type represent a window’s possible backing locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum
type WindowBackingLocation uint

const (
	WindowBackingLocationDefault WindowBackingLocation = 0
	WindowBackingLocationVideoMemory WindowBackingLocation = 1
	WindowBackingLocationMainMemory WindowBackingLocation = 2
	AlertDefaultReturn WindowBackingLocation = 3
	AlertAlternateReturn WindowBackingLocation = 4
	AlertOtherReturn WindowBackingLocation = 5
	AlertErrorReturn WindowBackingLocation = 6
)

// NSBackingStoreType - Constants that specify how the window device buffers the drawing done in a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
type BackingStoreType uint

const (
	// BackingStoreBuffered - The window renders all drawing into a display buffer and then flushes it to the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/buffered
	BackingStoreBuffered BackingStoreType = 2
	// BackingStoreNonretained - The window draws directly to the screen without using any buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/nonretained
	BackingStoreNonretained BackingStoreType = 1
	// BackingStoreRetained - The window uses a buffer, but draws directly to the screen where possible and to the buffer for obscured portions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/retained
	BackingStoreRetained BackingStoreType = 0
)

// NSWindowButton - Constants that provide a way to access standard title bar buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType
type WindowButton uint

const (
	// WindowCloseButton - The close button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/closeButton
	WindowCloseButton WindowButton = 0
	// WindowDocumentIconButton - The document icon button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/documentIconButton
	WindowDocumentIconButton WindowButton = 4
	// WindowDocumentVersionsButton - The document versions button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/documentVersionsButton
	WindowDocumentVersionsButton WindowButton = 5
	// WindowMiniaturizeButton - The minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/miniaturizeButton
	WindowMiniaturizeButton WindowButton = 1
	// WindowToolbarButton - The toolbar button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/toolbarButton
	WindowToolbarButton WindowButton = 3
	// WindowZoomButton - The zoom button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/zoomButton
	WindowZoomButton WindowButton = 2
)

// NSWindowCollectionBehavior - Window collection behaviors related to Mission Control, Spaces, and Stage Manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type WindowCollectionBehavior uint

const (
	// WindowCollectionBehaviorAuxiliary - The behavior marking this window as auxiliary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/auxiliary
	WindowCollectionBehaviorAuxiliary WindowCollectionBehavior = 13
	// WindowCollectionBehaviorCanJoinAllApplications - The behavior marking this window as one that can join all apps for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllApplications
	WindowCollectionBehaviorCanJoinAllApplications WindowCollectionBehavior = 14
	// WindowCollectionBehaviorCanJoinAllSpaces - The window can appear in all spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllSpaces
	WindowCollectionBehaviorCanJoinAllSpaces WindowCollectionBehavior = 1
	// WindowCollectionBehaviorFullScreenAllowsTiling - The window can be a secondary full screen tile even if it can’t be a full screen window itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAllowsTiling
	WindowCollectionBehaviorFullScreenAllowsTiling WindowCollectionBehavior = 10
	// WindowCollectionBehaviorFullScreenAuxiliary - The window displays on the same space as the full screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAuxiliary
	WindowCollectionBehaviorFullScreenAuxiliary WindowCollectionBehavior = 8
	// WindowCollectionBehaviorFullScreenDisallowsTiling - The window doesn’t support being a full-screen tile window, but may support being a full-screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenDisallowsTiling
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 11
	// WindowCollectionBehaviorFullScreenNone - The window doesn’t support full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenNone
	WindowCollectionBehaviorFullScreenNone WindowCollectionBehavior = 9
	// WindowCollectionBehaviorFullScreenPrimary - The window can enter full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenPrimary
	WindowCollectionBehaviorFullScreenPrimary WindowCollectionBehavior = 7
	// WindowCollectionBehaviorIgnoresCycle - The window isn’t part of the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/ignoresCycle
	WindowCollectionBehaviorIgnoresCycle WindowCollectionBehavior = 6
	// WindowCollectionBehaviorManaged - The window participates in Mission Control and Spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/managed
	WindowCollectionBehaviorManaged WindowCollectionBehavior = 2
	// WindowCollectionBehaviorMoveToActiveSpace - When the window becomes active, move it to the active space instead of switching spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/moveToActiveSpace
	WindowCollectionBehaviorMoveToActiveSpace WindowCollectionBehavior = 1
	// WindowCollectionBehaviorParticipatesInCycle - The window participates in the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/participatesInCycle
	WindowCollectionBehaviorParticipatesInCycle WindowCollectionBehavior = 5
	// WindowCollectionBehaviorPrimary - The behavior marking this window as primary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/primary
	WindowCollectionBehaviorPrimary WindowCollectionBehavior = 12
	// WindowCollectionBehaviorStationary - Mission Control doesn’t affect the window, so it stays visible and stationary, like the desktop window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/stationary
	WindowCollectionBehaviorStationary WindowCollectionBehavior = 4
	// WindowCollectionBehaviorTransient - The window floats in Spaces and hides in Mission Control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/transient
	WindowCollectionBehaviorTransient WindowCollectionBehavior = 3
	// WindowCollectionBehaviorDefault - The window appears in only one space at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowCollectionBehavior/NSWindowCollectionBehaviorDefault
	WindowCollectionBehaviorDefault WindowCollectionBehavior = 0
)

// NSWindowDepth - A type that represents the depth, or amount of memory, for a single pixel in a window or screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth
type WindowDepth uint

const (
	// WindowDepthOnehundredtwentyeightBitRGB - One hundred and twenty eight bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/onehundredtwentyeightBitRGB
	WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 0
	// WindowDepthSixtyfourBitRGB - Sixty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/sixtyfourBitRGB
	WindowDepthSixtyfourBitRGB WindowDepth = 0
	// WindowDepthTwentyfourBitRGB - Twenty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/twentyfourBitRGB
	WindowDepthTwentyfourBitRGB WindowDepth = 0
)

// NSWindowNumberListOptions - Options to use when retrieving window numbers from the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
type WindowNumberListOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allApplications
	WindowNumberListAllApplications WindowNumberListOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allSpaces
	WindowNumberListAllSpaces WindowNumberListOptions = 1
)

// NSWindowOcclusionState - Specifies whether the window is occluded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type WindowOcclusionState uint

const (
	// WindowOcclusionStateVisible - If set, at least part of the window is visible; if not set, the entire window is occluded. A window that has a nonrectangular shape can be entirely occluded onscreen, but if its bounding box falls into a visible region, the window is considered to be visible. Note that a completely transparent window may also be considered visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct/visible
	WindowOcclusionStateVisible WindowOcclusionState = 1
)

// NSWindowOrderingMode - Constants that let you specify how a window is ordered relative to another window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type WindowOrderingMode uint

const (
	// WindowAbove - Moves the window above the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/above
	WindowAbove WindowOrderingMode = 1
	// WindowBelow - Moves the window below the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/below
	WindowBelow WindowOrderingMode = -1
	// WindowOut - Moves the window off the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/out
	WindowOut WindowOrderingMode = 0
)

// NSSelectionDirection - Constants that specify the direction a window is currently using to change the key view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
type SelectionDirection uint

const (
	// DirectSelection - The window isn’t traversing the key view loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/directSelection
	DirectSelection SelectionDirection = 0
	// SelectingNext - The window is proceeding to the next valid key view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/selectingNext
	SelectingNext SelectionDirection = 1
	// SelectingPrevious - The window is proceeding to the previous valid key view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/selectingPrevious
	SelectingPrevious SelectionDirection = 2
)

// NSWindowSharingType - Constants that represent the access levels other processes can have to a window’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
type WindowSharingType uint

const (
	// WindowSharingNone - A legacy constant that macOS no longer uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/none
	WindowSharingNone WindowSharingType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/readOnly
	WindowSharingReadOnly WindowSharingType = 1
)

// NSWindowStyleMask - Constants that specify the style of a window, and that you can combine with the C bitwise OR operator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
type WindowStyleMask uint

const (
	// WindowStyleMaskBorderless - The window displays none of the usual peripheral elements. Useful only for display or caching purposes. A window that uses   can’t become key or main, unless the value of   or   is  . Note that you can set a window’s or panel’s style mask to   in Interface Builder by deselecting Title Bar in the Appearance section of the Attributes inspector.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/borderless
	WindowStyleMaskBorderless WindowStyleMask = 0
	// WindowStyleMaskClosable - The window displays a close button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/closable
	WindowStyleMaskClosable WindowStyleMask = 1
	// WindowStyleMaskDocModalWindow - The window is a document-modal panel (or  a subclass of  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/docModalWindow
	WindowStyleMaskDocModalWindow WindowStyleMask = 1
	// WindowStyleMaskFullScreen - The window can appear full screen. A fullscreen window does not draw its title bar, and may have special handling for its toolbar. (This mask is automatically toggled when   is called.)
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullScreen
	WindowStyleMaskFullScreen WindowStyleMask = 2
	// WindowStyleMaskFullSizeContentView - When set, the window’s   consumes the full size of the window. Although you can combine this constant with other window style masks, it is respected only for windows with a title bar. Note that using this mask opts in to layer-backing. Use the   or the   to lay out views underneath the title bar–toolbar area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullSizeContentView
	WindowStyleMaskFullSizeContentView WindowStyleMask = 3
	// WindowStyleMaskHUDWindow - The window is a HUD panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/hudWindow
	WindowStyleMaskHUDWindow WindowStyleMask = 2
	// WindowStyleMaskMiniaturizable - The window displays a minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/miniaturizable
	WindowStyleMaskMiniaturizable WindowStyleMask = 1
	// WindowStyleMaskNonactivatingPanel - The window is a panel or a subclass of   that does not activate the owning app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/nonactivatingPanel
	WindowStyleMaskNonactivatingPanel WindowStyleMask = 1
	// WindowStyleMaskResizable - The window can be resized by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/resizable
	WindowStyleMaskResizable WindowStyleMask = 1
	// WindowStyleMaskTexturedBackground - The window uses a textured background that darkens when the window is key or main and lightens when it is inactive, and may have a second gradient in the section below the window content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/texturedBackground
	WindowStyleMaskTexturedBackground WindowStyleMask = 2
	// WindowStyleMaskTitled - The window displays a title bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/titled
	WindowStyleMaskTitled WindowStyleMask = 1
	// WindowStyleMaskUnifiedTitleAndToolbar - This constant has no effect, because all windows that include a toolbar use the unified style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/unifiedTitleAndToolbar
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 1
	// WindowStyleMaskUtilityWindow - The window is a panel or a subclass of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/utilityWindow
	WindowStyleMaskUtilityWindow WindowStyleMask = 1
)

// NSWindowTabbingMode - The preferred tabbing behavior of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
type WindowTabbingMode uint

const (
	// WindowTabbingModeAutomatic - A window that automatically tabs together based on the user’s tabbing preference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/automatic
	WindowTabbingModeAutomatic WindowTabbingMode = 0
	// WindowTabbingModeDisallowed - A window that explicitly does not prefer to tab together with other windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/disallowed
	WindowTabbingModeDisallowed WindowTabbingMode = 2
	// WindowTabbingModePreferred - A window that explicitly prefers to tab together with other windows with the same tabbing identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/preferred
	WindowTabbingModePreferred WindowTabbingMode = 1
)

// NSWindowTitleVisibility - Specifies the appearance of the window’s title bar area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type WindowTitleVisibility uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/hidden
	WindowTitleHidden WindowTitleVisibility = 1
	// WindowTitleVisible - The window has the regular window title and title bar buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/visible
	WindowTitleVisible WindowTitleVisibility = 0
)

// NSWindowToolbarStyle - Styles that determine the appearance and location of the toolbar in relation to the title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum
type WindowToolbarStyle uint

const (
	// WindowToolbarStyleAutomatic - A style indicating that the system determines the toolbar’s appearance and location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/automatic
	WindowToolbarStyleAutomatic WindowToolbarStyle = 0
	// WindowToolbarStyleExpanded - A style indicating that the toolbar appears below the window title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/expanded
	WindowToolbarStyleExpanded WindowToolbarStyle = 1
	// WindowToolbarStylePreference - A style indicating that the toolbar appears below the window title with toolbar items centered in the toolbar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/preference
	WindowToolbarStylePreference WindowToolbarStyle = 2
	// WindowToolbarStyleUnified - A style indicating that the toolbar appears next to the window title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/unified
	WindowToolbarStyleUnified WindowToolbarStyle = 3
	// WindowToolbarStyleUnifiedCompact - A style indicating that the toolbar appears next to the window title and with reduced margins to allow more focus on the window’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/unifiedCompact
	WindowToolbarStyleUnifiedCompact WindowToolbarStyle = 4
)

// NSWindowUserTabbingPreference - A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
type WindowUserTabbingPreference uint

const (
	// WindowUserTabbingPreferenceAlways - A value that indicates a window should always display as tabs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/always
	WindowUserTabbingPreferenceAlways WindowUserTabbingPreference = 1
	// WindowUserTabbingPreferenceInFullScreen - A value that indicates a window should only display as tabs when in full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/inFullScreen
	WindowUserTabbingPreferenceInFullScreen WindowUserTabbingPreference = 2
	// WindowUserTabbingPreferenceManual - A value that indicates a window should display as tabs according to the window’s tabbing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/manual
	WindowUserTabbingPreferenceManual WindowUserTabbingPreference = 0
)

// NSWorkspaceAuthorizationType - The types of privileged file operations that can be authorized by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType
type WorkspaceAuthorizationType uint

const (
	// WorkspaceAuthorizationTypeCreateSymbolicLink - Authorization for the app to create a symbolic link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/createSymbolicLink
	WorkspaceAuthorizationTypeCreateSymbolicLink WorkspaceAuthorizationType = 0
	// WorkspaceAuthorizationTypeReplaceFile - Authorization for the app to perform an atomic file write without changing the target file’s permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/replaceFile
	WorkspaceAuthorizationTypeReplaceFile WorkspaceAuthorizationType = 2
	// WorkspaceAuthorizationTypeSetAttributes - Authorization for the app to change specific file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/setAttributes
	WorkspaceAuthorizationTypeSetAttributes WorkspaceAuthorizationType = 1
)

// NSWorkspaceIconCreationOptions - Constants that describe options for creating icons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
type WorkspaceIconCreationOptions uint

const (
	// Exclude10_4ElementsIconCreationOption - An option to suppress generation of the new higher resolution icon representations that are supported in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/exclude10_4ElementsIconCreationOption
	Exclude10_4ElementsIconCreationOption WorkspaceIconCreationOptions = 1
	// ExcludeQuickDrawElementsIconCreationOption - An option to suppress generation of the QuickDraw format icon representations that are used in macOS 10.0 through macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/excludeQuickDrawElementsIconCreationOption
	ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 1
)

// NSWorkspaceLaunchOptions - Constants specifying how you want to launch an app
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions
type WorkspaceLaunchOptions uint

const (
	WorkspaceLaunchAndPrint WorkspaceLaunchOptions = 0
	WorkspaceLaunchWithErrorPresentation WorkspaceLaunchOptions = 1
	WorkspaceLaunchInhibitingBackgroundOnly WorkspaceLaunchOptions = 2
	WorkspaceLaunchWithoutAddingToRecents WorkspaceLaunchOptions = 3
	WorkspaceLaunchWithoutActivation WorkspaceLaunchOptions = 4
	WorkspaceLaunchAsync WorkspaceLaunchOptions = 5
	WorkspaceLaunchNewInstance WorkspaceLaunchOptions = 6
	WorkspaceLaunchAndHide WorkspaceLaunchOptions = 7
	WorkspaceLaunchAndHideOthers WorkspaceLaunchOptions = 8
	WorkspaceLaunchDefault WorkspaceLaunchOptions = 9
	WorkspaceLaunchAllowingClassicStartup WorkspaceLaunchOptions = 10
	WorkspaceLaunchPreferringClassic WorkspaceLaunchOptions = 11
)

// NSWritingDirection - Constants that specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
type WritingDirection uint

const (
	// WritingDirectionLeftToRight - The writing direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/leftToRight
	WritingDirectionLeftToRight WritingDirection = 0
	// WritingDirectionNatural - The writing direction of the current script that the system determines using the Unicode Bidi Algorithm rules P2 and P3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/natural
	WritingDirectionNatural WritingDirection = -1
	// WritingDirectionRightToLeft - The writing direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/rightToLeft
	WritingDirectionRightToLeft WritingDirection = 1
)

// NSWritingDirectionFormatType - Constants for the writing direction attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType
type WritingDirectionFormatType uint

const (
	// WritingDirectionEmbedding - Text is embedded in text with another writing direction. For example, an English quotation in the middle of an Arabic sentence could be marked as being embedded left-to-right text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType/embedding
	WritingDirectionEmbedding WritingDirectionFormatType = 0
	// WritingDirectionOverride - Enables character types with inherent directionality to be overridden when required for special cases, such as for part numbers made of mixed English, digits, and Hebrew letters to be written from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType/override
	WritingDirectionOverride WritingDirectionFormatType = 0
)

// NSWritingToolsBehavior - Constants that specify the Writing Tools experience for the underlying view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior
type WritingToolsBehavior uint

const (
	// WritingToolsBehaviorComplete - An option to provide the complete Writing Tools experience for the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/complete
	WritingToolsBehaviorComplete WritingToolsBehavior = 1
	// WritingToolsBehaviorNone - An option to prevent Writing Tools from modifying the text in the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/none
	WritingToolsBehaviorNone WritingToolsBehavior = -1
)

// NSWritingToolsCoordinatorState - The states that indicate the current activity, if any, Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum
type WritingToolsCoordinatorState uint

const (
	// WritingToolsCoordinatorStateInactive - A state that indicates Writing Tools isn’t currently performing   any work on your view’s content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/inactive
	WritingToolsCoordinatorStateInactive WritingToolsCoordinatorState = 0
	// WritingToolsCoordinatorStateInteractiveResting - A state that indicates Writing Tools is in the resting state   for an inline editing experience.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/interactiveResting
	WritingToolsCoordinatorStateInteractiveResting WritingToolsCoordinatorState = 2
	// WritingToolsCoordinatorStateInteractiveStreaming - A state that indicates Writing Tools is processing a request and   incorporating changes interactively into your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/interactiveStreaming
	WritingToolsCoordinatorStateInteractiveStreaming WritingToolsCoordinatorState = 3
	// WritingToolsCoordinatorStateNoninteractive - A state that indicates Writing Tools is handling interactions in   the system UI, instead of in your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/noninteractive
	WritingToolsCoordinatorStateNoninteractive WritingToolsCoordinatorState = 1
)

// NSWritingToolsResultOptions - Constants to specify what type of content to allow in Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions
type WritingToolsResultOptions uint

const (
	WritingToolsResultDefault WritingToolsResultOptions = 0
	WritingToolsResultPlainText WritingToolsResultOptions = 1
	WritingToolsResultRichText WritingToolsResultOptions = 1
	WritingToolsResultList WritingToolsResultOptions = 1
	WritingToolsResultTable WritingToolsResultOptions = 1
	WritingToolsResultPresentationIntent WritingToolsResultOptions = 2
	TextFieldSquareBezel WritingToolsResultOptions = 0
	TextFieldRoundedBezel WritingToolsResultOptions = 1
)


