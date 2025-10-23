// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Enum types and constants
// AccessibilityAnnotationPosition - Constants that specify the position where the annotation applies.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition
type AccessibilityAnnotationPosition uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition/start
	AccessibilityAnnotationPositionStart AccessibilityAnnotationPosition = 1
)

// AccessibilityPriorityLevel - A data type for notification priority levels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel
type AccessibilityPriorityLevel uint

const (
	AccessibilityPriorityLow AccessibilityPriorityLevel = 10
	AccessibilityPriorityMedium AccessibilityPriorityLevel = 50
	AccessibilityPriorityHigh AccessibilityPriorityLevel = 90
)

// AnimationEffect - The type for standard system animation effects, which include both display and sound.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type AnimationEffect uint

const (
	AnimationEffectDisappearingItemDefault AnimationEffect = 0
	AnimationEffectPoof AnimationEffect = 10
)

// ApplicationActivationOptions - The following flags are for 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions
type ApplicationActivationOptions uint

const (
	ApplicationActivateAllWindows ApplicationActivationOptions = 1
	ApplicationActivateIgnoringOtherApps ApplicationActivationOptions = 2
)

// ApplicationActivationPolicy - Activation policies (used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum
type ApplicationActivationPolicy uint

const (
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = 0
	ApplicationActivationPolicyAccessory ApplicationActivationPolicy = 1
	ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 2
)

// RequestUserAttentionType - These constants specify the level of severity of a user attention request and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType
type RequestUserAttentionType uint

const (
	CriticalRequest RequestUserAttentionType = 0
	InformationalRequest RequestUserAttentionType = 10
)

// BezierPathElement - Constants that specify basic path element commands.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
type BezierPathElement uint

const (
	// BezierPathElementClosePath - Marks the end of the current subpath at the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/closePath
	BezierPathElementClosePath BezierPathElement = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/cubicCurveTo
	BezierPathElementCubicCurveTo BezierPathElement = 2
	// BezierPathElementCurveTo - Creates a curved line segment from the current point to the specified endpoint using two control points to define the curve.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/curveTo
	BezierPathElementCurveTo BezierPathElement = 5
	// BezierPathElementLineTo - Creates a straight line from the current drawing point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/lineTo
	BezierPathElementLineTo BezierPathElement = 1
	// BezierPathElementMoveTo - Moves the path object’s current drawing point to the specified point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/moveTo
	BezierPathElementMoveTo BezierPathElement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType/quadraticCurveTo
	BezierPathElementQuadraticCurveTo BezierPathElement = 4
)

// LineCapStyle - Constants that specify the shape of endpoints for an open path when it is stroked.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum
type LineCapStyle uint

const (
	// LineCapStyleButt - Specifies a butt line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/butt
	LineCapStyleButt LineCapStyle = 0
	// LineCapStyleRound - Specifies a round line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/round
	LineCapStyleRound LineCapStyle = 1
	// LineCapStyleSquare - Specifies a square line cap style for endpoints for an open path when stroked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineCapStyle-swift.enum/square
	LineCapStyleSquare LineCapStyle = 2
)

// LineJoinStyle - Constants that specify the shape of the joins between connected segments of a stroked path.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum
type LineJoinStyle uint

const (
	// LineJoinStyleBevel - Specifies a bevel line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/bevel
	LineJoinStyleBevel LineJoinStyle = 2
	// LineJoinStyleMiter - Specifies a miter line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/miter
	LineJoinStyleMiter LineJoinStyle = 0
	// LineJoinStyleRound - Specifies a round line shape of the joints between connected segments of a stroked path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/LineJoinStyle-swift.enum/round
	LineJoinStyleRound LineJoinStyle = 1
)

// WindingRule - Constants that specify the winding rule a Bézier path uses.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum
type WindingRule uint

const (
	// WindingRuleEvenOdd - Specifies the even-odd winding rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum/evenOdd
	WindingRuleEvenOdd WindingRule = 1
	// WindingRuleNonZero - Specifies the non-zero winding rule.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath/WindingRule-swift.enum/nonZero
	WindingRuleNonZero WindingRule = 0
)

// TIFFCompression - Constants that represent the supported TIFF data-compression schemes.
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

// BorderType - These constants specify the type of a view’s border.
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
	// NoBorder - No border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/noBorder
	NoBorder BorderType = 0
)

// BoxType - These constants and data type identifies box types, which, in conjunction with a box’s border type, define the appearance of the box.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
type BoxType uint

const (
	BoxPrimary BoxType = 0
	BoxSeparator BoxType = 2
	BoxCustom BoxType = 3
)

// TitlePosition - Specify the location of a box’s title with respect to its border.
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

// CellAttribute - Constants for specifying how a button behaves when pressed and how it displays its state.
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

// CellType - Constants for specifying how a cell represents its data (as text or as an image).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType
type CellType uint

const (
	NullCellType CellType = 0
	TextCellType CellType = 1
	ImageCellType CellType = 2
)

// CellHitResult - Constants used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
type CellHitResult uint

const (
	CellHitNone CellHitResult = 0
	CellHitContentArea CellHitResult = 1
	CellHitEditableTextArea CellHitResult = 2
	CellHitTrackableArea CellHitResult = 4
)

// CellStyleMask - Constants for specifying what happens when a button is pressed or is displaying its alternate state.
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

// CharacterCollection - Values that map character identifiers to glyphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection
type CharacterCollection uint

const (
	// AdobeGB1CharacterCollection - Indicates the Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeGB1CharacterCollection
	AdobeGB1CharacterCollection CharacterCollection = 2
	// AdobeJapan1CharacterCollection - Indicates the Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeJapan1CharacterCollection
	AdobeJapan1CharacterCollection CharacterCollection = 3
	// AdobeJapan2CharacterCollection - Indicates the Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeJapan2CharacterCollection
	AdobeJapan2CharacterCollection CharacterCollection = 4
)

// CollectionViewDropOperation - These constants specify if acceptance of a drop should be at the item it is dropped on or before the item. These constants are used by the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation
type CollectionViewDropOperation uint

const (
	CollectionViewDropOn CollectionViewDropOperation = 0
	CollectionViewDropBefore CollectionViewDropOperation = 1
)

// CollectionViewScrollDirection - Constants indicating the scrolling direction for the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection
type CollectionViewScrollDirection uint

const (
	CollectionViewScrollDirectionVertical CollectionViewScrollDirection = 0
	CollectionViewScrollDirectionHorizontal CollectionViewScrollDirection = 1
)

// CollectionViewScrollPosition - Constants indicating the options for scrolling the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition
type CollectionViewScrollPosition uint

const (
	CollectionViewScrollPositionNone CollectionViewScrollPosition = 0
	CollectionViewScrollPositionTop CollectionViewScrollPosition = 1
	CollectionViewScrollPositionCenteredVertically CollectionViewScrollPosition = 2
	CollectionViewScrollPositionBottom CollectionViewScrollPosition = 4
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 512
	CollectionViewScrollPositionLeft CollectionViewScrollPosition = 8
	CollectionViewScrollPositionCenteredHorizontally CollectionViewScrollPosition = 16
	CollectionViewScrollPositionRight CollectionViewScrollPosition = 32
	CollectionViewScrollPositionLeadingEdge CollectionViewScrollPosition = 64
	CollectionViewScrollPositionTrailingEdge CollectionViewScrollPosition = 128
	CollectionViewScrollPositionNearestVerticalEdge CollectionViewScrollPosition = 256
)

// CollectionUpdateAction - Constants indicating the type of action being performed on an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
type CollectionUpdateAction uint

const (
	CollectionUpdateActionInsert CollectionUpdateAction = 0
	CollectionUpdateActionDelete CollectionUpdateAction = 1
	CollectionUpdateActionReload CollectionUpdateAction = 2
	CollectionUpdateActionMove CollectionUpdateAction = 3
	CollectionUpdateActionNone CollectionUpdateAction = 4
)

// ColorType - Constants that indicate the color’s type, and which methods may be called on the color object.
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

// ColorSystemEffect - Constants for user interactions that change the appearance of a view or control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type ColorSystemEffect uint

const (
	// ColorSystemEffectPressed - The color that indicates the item was pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/pressed
	ColorSystemEffectPressed ColorSystemEffect = 1
	// ColorSystemEffectRollover - The color that indicates the mouse rolled over the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/rollover
	ColorSystemEffectRollover ColorSystemEffect = 4
)

// ColorRenderingIntent - Constants that specify how Cocoa should handle colors that are not located within the destination color space of a graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
type ColorRenderingIntent uint

const (
	ColorRenderingIntentDefault ColorRenderingIntent = 0
	ColorRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
	ColorRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
	ColorRenderingIntentPerceptual ColorRenderingIntent = 3
	ColorRenderingIntentSaturation ColorRenderingIntent = 4
)

// ColorWellStyle - Constants that specify the appearance and interaction modes for a color well.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
type ColorWellStyle uint

const (
	ColorWellStyleDefault ColorWellStyle = 0
	ColorWellStyleMinimal ColorWellStyle = 1
	ColorWellStyleExpanded ColorWellStyle = 2
)

// ComboButtonStyle - Constants that indicate how a combo button presents its menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum
type ComboButtonStyle uint

const (
	// ComboButtonStyleUnified - A style that unifies the button’s title and image with the menu indicator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum/unified
	ComboButtonStyleUnified ComboButtonStyle = 1
)

// CompositingOperation - Constants that describe compositing operators in terms of source and destination images, each having an opaque and transparent region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
type CompositingOperation uint

const (
	// CompositingOperationCopy - The source image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/copy
	CompositingOperationCopy CompositingOperation = 1
	// CompositingOperationSaturation - Uses the saturation value of the source and the hue and luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/saturation
	CompositingOperationSaturation CompositingOperation = 26
	// CompositingOperationSourceOver - The source image wherever it is opaque, and the destination image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceOver
	CompositingOperationSourceOver CompositingOperation = 2
)

// ControlSize - A constant for specifying a cell’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
type ControlSize uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/extraLarge
	ControlSizeExtraLarge ControlSize = 4
)

// ControlTint - Constants for specifying a cell’s tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint
type ControlTint uint

const (
	// BlueControlTint - Aqua control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/blueControlTint
	BlueControlTint ControlTint = 1
	// ClearControlTint - Clear control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/clearControlTint
	ClearControlTint ControlTint = 7
	// DefaultControlTint - The current default tint setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/defaultControlTint
	DefaultControlTint ControlTint = 0
	// GraphiteControlTint - Graphite control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/graphiteControlTint
	GraphiteControlTint ControlTint = 6
)

// CursorFrameResizePosition - The position along the perimeter of a rectangular frame (its edges and corners) from which it’s resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition
type CursorFrameResizePosition uint

const (
	// CursorFrameResizePositionBottom - The bottom edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottom
	CursorFrameResizePositionBottom CursorFrameResizePosition = 4
	// CursorFrameResizePositionBottomLeft - The bottom left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottomLeft
	CursorFrameResizePositionBottomLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionBottomRight - The bottom right corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottomRight
	CursorFrameResizePositionBottomRight CursorFrameResizePosition = 0
	// CursorFrameResizePositionLeft - The left edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/left
	CursorFrameResizePositionLeft CursorFrameResizePosition = 2
	// CursorFrameResizePositionRight - The right edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/right
	CursorFrameResizePositionRight CursorFrameResizePosition = 8
	// CursorFrameResizePositionTop - The top edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/top
	CursorFrameResizePositionTop CursorFrameResizePosition = 1
	// CursorFrameResizePositionTopLeft - The top left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/topLeft
	CursorFrameResizePositionTopLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionTopRight - The top right corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/topRight
	CursorFrameResizePositionTopRight CursorFrameResizePosition = 0
)

// CursorFrameResizeDirections - The directions in which a rectangular frame can be resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections
type CursorFrameResizeDirections uint

const (
	// CursorFrameResizeDirectionsAll - Indicates that the shape can be resized inwards or wards to be either smaller or larger, respectively.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsAll
	CursorFrameResizeDirectionsAll CursorFrameResizeDirections = 0
	// CursorFrameResizeDirectionsInward - Indicates that the shape can be resized inwards to be smaller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsInward
	CursorFrameResizeDirectionsInward CursorFrameResizeDirections = 1
	// CursorFrameResizeDirectionsOutward - Indicates that the shape can be resized outwards to be larger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsOutward
	CursorFrameResizeDirectionsOutward CursorFrameResizeDirections = 2
)

// DatePickerMode - Constants that define whether the picker provides a single date, or a range of dates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode
type DatePickerMode uint

const (
	DatePickerModeSingle DatePickerMode = 0
	DatePickerModeRange DatePickerMode = 1
)

// DisplayGamut enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut
type DisplayGamut uint

const (
	DisplayGamutSRGB DisplayGamut = 1
	DisplayGamutP3 DisplayGamut = 2
)

// SaveOperationType - Constants for specifying the type of document-save operation to perform.
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

// DragOperation - A group of constants that represent which operations the dragging source can perform on dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation
type DragOperation uint

const (
	DragOperationNone DragOperation = 0
	DragOperationCopy DragOperation = 1
	DragOperationLink DragOperation = 2
	DragOperationGeneric DragOperation = 4
	DragOperationPrivate DragOperation = 8
	DragOperationMove DragOperation = 16
	DragOperationDelete DragOperation = 32
	DragOperationAll_Obsolete DragOperation = 33
	DragOperationAll DragOperation = 34
)

// DraggingContext - Constants that specify whether a drag terminates within or outside the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
type DraggingContext uint

const (
	DraggingContextOutsideApplication DraggingContext = 0
	DraggingContextWithinApplication DraggingContext = 1
)

// DraggingFormation - Constants that control the visual format of multiple dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation
type DraggingFormation uint

const (
	DraggingFormationDefault DraggingFormation = 0
	DraggingFormationNone DraggingFormation = 1
	DraggingFormationPile DraggingFormation = 2
	DraggingFormationList DraggingFormation = 3
	DraggingFormationStack DraggingFormation = 4
)

// DraggingItemEnumerationOptions - A group of constants that specify options to use when enumerating dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions
type DraggingItemEnumerationOptions uint

const (
	DraggingItemEnumerationClearNonenumeratedImages DraggingItemEnumerationOptions = 65536
)

// EventButtonMask - Constants you use to identify the activated tablet buttons in an event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct
type EventButtonMask uint

const (
	// EventButtonMaskPenLowerSide - A mask that matches the button on the lower side of the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct/penLowerSide
	EventButtonMaskPenLowerSide EventButtonMask = 2
	// EventButtonMaskPenTip - A mask that matches the pen tip.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct/penTip
	EventButtonMaskPenTip EventButtonMask = 1
	// EventButtonMaskPenUpperSide - A mask that matches the button on the upper side of the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct/penUpperSide
	EventButtonMaskPenUpperSide EventButtonMask = 4
)

// EventSubtype - Subtypes for various types of events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype
type EventSubtype uint

const (
	// EventSubtypeApplicationActivated - An app-activation event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/applicationActivated
	EventSubtypeApplicationActivated EventSubtype = 1
	// EventSubtypeApplicationDeactivated - An app-deactivation event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/applicationDeactivated
	EventSubtypeApplicationDeactivated EventSubtype = 2
	// EventSubtypeMouseEvent - A mouse event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/mouseEvent
	EventSubtypeMouseEvent EventSubtype = 0
	// EventSubtypePowerOff - An event that indicates a system shutdown or restart operation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/powerOff
	EventSubtypePowerOff EventSubtype = 1
	// EventSubtypeScreenChanged - An event that indicates a window changed screens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/screenChanged
	EventSubtypeScreenChanged EventSubtype = 8
	// EventSubtypeTabletPoint - A tablet-pointer event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/tabletPoint
	EventSubtypeTabletPoint EventSubtype = 1
	// EventSubtypeTabletProximity - A tablet-proximity event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/tabletProximity
	EventSubtypeTabletProximity EventSubtype = 2
	// EventSubtypeTouch - A touch event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/touch
	EventSubtypeTouch EventSubtype = 3
	// EventSubtypeWindowExposed - An event that indicates a window’s contents are visible again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/windowExposed
	EventSubtypeWindowExposed EventSubtype = 0
	// EventSubtypeWindowMoved - An event that indicates a window moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/windowMoved
	EventSubtypeWindowMoved EventSubtype = 4
)

// EventType - Constants for the types of events that responder objects can handle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType
type EventType uint

const (
	// EventTypeAppKitDefined - An AppKit-related event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/appKitDefined
	EventTypeAppKitDefined EventType = 13
	// EventTypeKeyUp - The user released a key on the keyboard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/keyUp
	EventTypeKeyUp EventType = 11
	// EventTypeLeftMouseDown - The user pressed the left mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/leftMouseDown
	EventTypeLeftMouseDown EventType = 1
	// EventTypeLeftMouseUp - The user released the left mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/leftMouseUp
	EventTypeLeftMouseUp EventType = 2
	// EventTypeMagnify - The user performed a pinch-open or pinch-close gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/magnify
	EventTypeMagnify EventType = 29
	// EventTypeMouseEntered - The cursor entered a well-defined area, such as a view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseEntered
	EventTypeMouseEntered EventType = 8
	// EventTypePressure - An event that reports a change in pressure on a pressure-sensitive device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/pressure
	EventTypePressure EventType = 36
	// EventTypeSwipe - The user performed a swipe gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/swipe
	EventTypeSwipe EventType = 30
	// EventTypeSystemDefined - A system-related event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/systemDefined
	EventTypeSystemDefined EventType = 14
	// EventTypeTabletPoint - The user touched a point on a tablet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/tabletPoint
	EventTypeTabletPoint EventType = 23
	// EventTypeTabletProximity - A pointing device is near, but not touching, the associated tablet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/tabletProximity
	EventTypeTabletProximity EventType = 24
)

// EventMask - Constants that you use to filter out specific event types from the stream of incoming events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask
type EventMask uint

const (
	// EventMaskAppKitDefined - A mask for AppKit–defined events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/appKitDefined
	EventMaskAppKitDefined EventMask = 0
	// EventMaskApplicationDefined - A mask for app-defined events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/applicationDefined
	EventMaskApplicationDefined EventMask = 0
	// EventMaskBeginGesture - A mask for begin-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/beginGesture
	EventMaskBeginGesture EventMask = 4
	// EventMaskDirectTouch - A mask for touch events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/directTouch
	EventMaskDirectTouch EventMask = 8
	// EventMaskEndGesture - A mask for end-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/endGesture
	EventMaskEndGesture EventMask = 5
	// EventMaskFlagsChanged - A mask for flags-changed events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/flagsChanged
	EventMaskFlagsChanged EventMask = 0
	// EventMaskKeyUp - A mask for key-up events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/keyUp
	EventMaskKeyUp EventMask = 0
	// EventMaskLeftMouseDragged - A mask for left mouse-dragged events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/leftMouseDragged
	EventMaskLeftMouseDragged EventMask = 0
	// EventMaskLeftMouseUp - A mask for left mouse-up events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/leftMouseUp
	EventMaskLeftMouseUp EventMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/mouseCancelled
	EventMaskMouseCancelled EventMask = 10
	// EventMaskMouseEntered - A mask for mouse-entered events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/mouseEntered
	EventMaskMouseEntered EventMask = 0
	// EventMaskMouseExited - A mask for mouse-exited events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/mouseExited
	EventMaskMouseExited EventMask = 0
	// EventMaskMouseMoved - A mask for mouse-moved events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/mouseMoved
	EventMaskMouseMoved EventMask = 0
	// EventMaskOtherMouseDown - A mask for tertiary mouse-down events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/otherMouseDown
	EventMaskOtherMouseDown EventMask = 0
	// EventMaskOtherMouseDragged - A mask for tertiary mouse-dragged events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/otherMouseDragged
	EventMaskOtherMouseDragged EventMask = 0
	// EventMaskOtherMouseUp - A mask for tertiary mouse-up events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/otherMouseUp
	EventMaskOtherMouseUp EventMask = 0
	// EventMaskPeriodic - A mask for periodic events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/periodic
	EventMaskPeriodic EventMask = 0
	// EventMaskPressure - A mask for pressure-change events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/pressure
	EventMaskPressure EventMask = 7
	// EventMaskRightMouseUp - A mask for right mouse-up events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/rightMouseUp
	EventMaskRightMouseUp EventMask = 0
	// EventMaskRotate - A mask for rotate-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/rotate
	EventMaskRotate EventMask = 3
	// EventMaskScrollWheel - A mask for scroll-wheel events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/scrollWheel
	EventMaskScrollWheel EventMask = 0
	// EventMaskSwipe - A mask for swipe-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/swipe
	EventMaskSwipe EventMask = 2
)

// EventGestureAxis - Constants that specify the direction of travel for a gesture.
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
	// EventGestureAxisVertical - The vertical axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/vertical
	EventGestureAxisVertical EventGestureAxis = 2
)

// EventModifierFlags - Flags that represent key states in an event object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type EventModifierFlags uint

const (
	// EventModifierFlagCapsLock - The Caps Lock key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/capsLock
	EventModifierFlagCapsLock EventModifierFlags = 65536
	// EventModifierFlagCommand - The Command key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/command
	EventModifierFlagCommand EventModifierFlags = 1048576
	// EventModifierFlagControl - The Control key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/control
	EventModifierFlagControl EventModifierFlags = 262144
	// EventModifierFlagDeviceIndependentFlagsMask - Device-independent modifier flags are masked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/deviceIndependentFlagsMask
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 4294901760
	// EventModifierFlagFunction - A function key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/function
	EventModifierFlagFunction EventModifierFlags = 8388608
	// EventModifierFlagHelp - The Help key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/help
	EventModifierFlagHelp EventModifierFlags = 4194304
	// EventModifierFlagOption - The Option or Alt key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/option
	EventModifierFlagOption EventModifierFlags = 524288
	// EventModifierFlagShift - The Shift key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/shift
	EventModifierFlagShift EventModifierFlags = 131072
)

// EventPhase - Constants that represent the possible phases during an event phase.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct
type EventPhase uint

const (
	// EventPhaseBegan - An event phase has begun.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/began
	EventPhaseBegan EventPhase = 1
	// EventPhaseCancelled - The system canceled the event phase.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/cancelled
	EventPhaseCancelled EventPhase = 1
	// EventPhaseChanged - An event phase has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/changed
	EventPhaseChanged EventPhase = 1
	// EventPhaseEnded - The event phase ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/ended
	EventPhaseEnded EventPhase = 1
	// EventPhaseMayBegin - The system event phase may begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/mayBegin
	EventPhaseMayBegin EventPhase = 1
	// EventPhaseStationary - An event phase is in progress but hasn’t moved since the previous event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/stationary
	EventPhaseStationary EventPhase = 1
	// EventPhaseNone - The event is not associated with a phase.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEventPhase/NSEventPhaseNone
	EventPhaseNone EventPhase = 0
)

// PointingDeviceType - The pointing-device types for tablet-proximity events or mouse events with a proximity event subtype.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum
type PointingDeviceType uint

const (
	// PointingDeviceTypeCursor - Represents a cursor pointing device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum/cursor
	PointingDeviceTypeCursor PointingDeviceType = 2
	// PointingDeviceTypeEraser - Represents the eraser end of a stylus-like pointing device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum/eraser
	PointingDeviceTypeEraser PointingDeviceType = 3
	// PointingDeviceTypePen - Represents the tip end of a stylus-like pointing device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum/pen
	PointingDeviceTypePen PointingDeviceType = 1
	// PointingDeviceTypeUnknown - Represents an unknown type of pointing device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PointingDeviceType-swift.enum/unknown
	PointingDeviceTypeUnknown PointingDeviceType = 0
)

// PressureBehavior - These constants describe the behavior and progression of a pressure gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum
type PressureBehavior int

const (
	// PressureBehaviorPrimaryAccelerator - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations occur during the mouse-down and mouse-up events when this behavior is configured. This configuration uses specific pressure mappings that are ideal for controlling speed as variable pressure occurs between the mouse-down and mouse-up events. The   class uses this behavior. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryAccelerator
	PressureBehaviorPrimaryAccelerator PressureBehavior = 3
	// PressureBehaviorPrimaryClick - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations (haptic feedback the user feels) occur during mouse-down and mouse-up events when this behavior is configured. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryClick
	PressureBehaviorPrimaryClick PressureBehavior = 1
	// PressureBehaviorPrimaryDeepClick - A pressure gesture’s behavior begins on left mouse-down events. Two stages are supported, and a stage transition animation may occur when moving between stages—from stage 1 to stage 0, stage 1 to stage 2, stage 2 to stage 1, and stage 2 to stage 0. With this behavior type, stage 2 becomes disabled once dragging occurs. When this behavior is configured, actuations occur during the mouse-down and mouse-up events, as well as when force click is activated and released when entering or exiting stage 2. This configuration is ideal for responding to force clicks. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDeepClick
	PressureBehaviorPrimaryDeepClick PressureBehavior = 5
	// PressureBehaviorPrimaryDeepDrag - A pressure gesture’s behavior begins on left mouse-down events. Two stages are supported, and a stage transition animation may occur when moving between stages—from stage 1 to stage 0, stage 1 to stage 2, stage 2 to stage 1, or stage 2 to stage 0. Actuations occur during the mouse-down and mouse-up events, as well as during the transitions up and down between stage 1 and stage 2, when this behavior is configured. This configuration is ideal for responding to force clicks during drag operations. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDeepDrag
	PressureBehaviorPrimaryDeepDrag PressureBehavior = 6
	// PressureBehaviorPrimaryDefault - This is the default behavior when a pressure gesture’s behavior has not been explicitly configured. In OS X 10.10.3, this behavior defaults to the behavior of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDefault
	PressureBehaviorPrimaryDefault PressureBehavior = 0
	// PressureBehaviorPrimaryGeneric - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations occur during the mouse-down and mouse-up events when this behavior is configured. This configuration is ideal for drawing, painting, and general use. Variable pressure occurs throughout the course of the gesture. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryGeneric
	PressureBehaviorPrimaryGeneric PressureBehavior = 2
	// PressureBehaviorUnknown - A pressure gesture’s behavior is not known, perhaps because the input device does not support pressure gestures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/unknown
	PressureBehaviorUnknown PressureBehavior = -1
)

// EventSwipeTrackingOptions - Constants that specify swipe-tracking options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions
type EventSwipeTrackingOptions uint

const (
	// EventSwipeTrackingClampGestureAmount - Don’t allow gestureAmount to go beyond +/-1.0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/clampGestureAmount
	EventSwipeTrackingClampGestureAmount EventSwipeTrackingOptions = 1
	// EventSwipeTrackingLockDirection - Clamp gestureAmount to 0 if the user starts to swipe in the opposite direction than they started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/lockDirection
	EventSwipeTrackingLockDirection EventSwipeTrackingOptions = 1
)

// FocusRingPlacement - Constants that indicate how the system draws the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement
type FocusRingPlacement uint

const (
	// FocusRingAbove - Draw the focus ring over an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/above
	FocusRingAbove FocusRingPlacement = 2
	// FocusRingBelow - Draw the focus ring under text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/below
	FocusRingBelow FocusRingPlacement = 1
	// FocusRingOnly - Draw the focus ring if you don’t have an image or text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/only
	FocusRingOnly FocusRingPlacement = 0
)

// FocusRingType - Constants that describe the style of the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type FocusRingType uint

const (
	// FocusRingTypeDefault - The default focus ring type for a view or cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/default
	FocusRingTypeDefault FocusRingType = 0
	// FocusRingTypeExterior - The standard Aqua focus ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/exterior
	FocusRingTypeExterior FocusRingType = 2
	// FocusRingTypeNone - No focus ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/none
	FocusRingTypeNone FocusRingType = 1
)

// FontAssetRequestOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/Options
type FontAssetRequestOptions uint

const (
	FontAssetRequestOptionUsesStandardUI FontAssetRequestOptions = 1
)

// FontCollectionOptions - Constants that support font collection management.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type FontCollectionOptions uint

const (
	// FontCollectionApplicationOnlyMask - Makes the collection available only to the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions/applicationOnlyMask
	FontCollectionApplicationOnlyMask FontCollectionOptions = 1
)

// FontDescriptorSymbolicTraits - A symbolic description of the stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct
type FontDescriptorSymbolicTraits uint

const (
	// FontDescriptorTraitUIOptimized - The font synthesizes appropriate attributes for user interface rendering, such as in control titles, if necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/UIOptimized
	FontDescriptorTraitUIOptimized FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitBold - The font’s style is boldface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/bold
	FontDescriptorTraitBold FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassClarendonSerifs - The font’s characters include variations of old style and transitional serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classClarendonSerifs
	FontDescriptorClassClarendonSerifs FontDescriptorSymbolicTraits = 4
	// FontDescriptorClassFreeformSerifs - The font’s characters include serifs, and don’t generally fit within other serif design classifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classFreeformSerifs
	FontDescriptorClassFreeformSerifs FontDescriptorSymbolicTraits = 7
	// FontDescriptorClassMask - The font family class mask that you use to access font descriptor values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classMask
	FontDescriptorClassMask FontDescriptorSymbolicTraits = 4026531840
	// FontDescriptorClassModernSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 20th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classModernSerifs
	FontDescriptorClassModernSerifs FontDescriptorSymbolicTraits = 3
	// FontDescriptorClassOldStyleSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 15th to 17th centuries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classOldStyleSerifs
	FontDescriptorClassOldStyleSerifs FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassOrnamentals - The font’s characters use highly decorated or stylized character shapes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classOrnamentals
	FontDescriptorClassOrnamentals FontDescriptorSymbolicTraits = 9
	// FontDescriptorClassSansSerif - The font’s characters don’t have serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSansSerif
	FontDescriptorClassSansSerif FontDescriptorSymbolicTraits = 8
	// FontDescriptorClassScripts - The font’s characters simulate handwriting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classScripts
	FontDescriptorClassScripts FontDescriptorSymbolicTraits = 10
	// FontDescriptorClassSlabSerifs - The font’s characters use square transitions, without brackets, between strokes and serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSlabSerifs
	FontDescriptorClassSlabSerifs FontDescriptorSymbolicTraits = 5
	// FontDescriptorClassSymbolic - The font’s characters consist mainly of symbols rather than letters and numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSymbolic
	FontDescriptorClassSymbolic FontDescriptorSymbolicTraits = 12
	// FontDescriptorClassTransitionalSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 18th to 19th centuries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classTransitionalSerifs
	FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 2
	// FontDescriptorTraitCondensed - The font’s characters have a condensed width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/condensed
	FontDescriptorTraitCondensed FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitExpanded - The font’s characters have an expanded width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/expanded
	FontDescriptorTraitExpanded FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitItalic - The font’s style is italic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/italic
	FontDescriptorTraitItalic FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitLooseLeading - The font uses a leading value that’s greater than the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/looseLeading
	FontDescriptorTraitLooseLeading FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitMonoSpace - The font’s characters all have the same width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/monoSpace
	FontDescriptorTraitMonoSpace FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitTightLeading - The font uses a leading value that’s less than the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/tightLeading
	FontDescriptorTraitTightLeading FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitVertical - The font uses vertical glyph variants and metrics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/vertical
	FontDescriptorTraitVertical FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassUnknown - The font has no design classification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptorSymbolicTraits/NSFontDescriptorClassUnknown
	FontDescriptorClassUnknown FontDescriptorSymbolicTraits = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptorSymbolicTraits/NSFontDescriptorTraitEmphasized
	FontDescriptorTraitEmphasized FontDescriptorSymbolicTraits = 2
)

// FontPanelModeMask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask
type FontPanelModeMask uint

const (
	FontPanelModeMaskFace FontPanelModeMask = 1
	FontPanelModeMaskSize FontPanelModeMask = 2
	FontPanelModeMaskCollection FontPanelModeMask = 4
	FontPanelModeMaskUnderlineEffect FontPanelModeMask = 256
	FontPanelModeMaskStrikethroughEffect FontPanelModeMask = 512
	FontPanelModeMaskTextColorEffect FontPanelModeMask = 1024
	FontPanelModeMaskDocumentColorEffect FontPanelModeMask = 2048
	FontPanelModeMaskShadowEffect FontPanelModeMask = 4096
	FontPanelModeMaskAllEffects FontPanelModeMask = 1048320
	FontPanelModesMaskStandardModes FontPanelModeMask = 65535
	FontPanelModesMaskAllModes FontPanelModeMask = 4294967295
)

// FontRenderingMode - The font rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type FontRenderingMode uint

const (
	// FontAntialiasedRenderingMode - Specifies antialiased, floating-point advancements rendering mode (synonymous with printerFont).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/antialiasedRenderingMode
	FontAntialiasedRenderingMode FontRenderingMode = 1
	// FontDefaultRenderingMode - Determines the actual mode based on the user preference settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/defaultRenderingMode
	FontDefaultRenderingMode FontRenderingMode = 0
	// FontIntegerAdvancementsRenderingMode - Specifies integer advancements rendering mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/integerAdvancementsRenderingMode
	FontIntegerAdvancementsRenderingMode FontRenderingMode = 2
)

// GlassEffectViewStyle enum type
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

// GlyphInscription - Constants that specify how a glyph is laid out relative to the previous glyph.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription
type GlyphInscription uint

const (
	// GlyphInscribeAbove - A glyph is rendered above the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeAbove
	GlyphInscribeAbove GlyphInscription = 2
	// GlyphInscribeBase - A base glyph; a character that the font can represent with a single glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeBase
	GlyphInscribeBase GlyphInscription = 0
	// GlyphInscribeBelow - A glyph is rendered below the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeBelow
	GlyphInscribeBelow GlyphInscription = 1
	// GlyphInscribeOverBelow - A glyph is rendered on top and below the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeOverBelow
	GlyphInscribeOverBelow GlyphInscription = 4
	// GlyphInscribeOverstrike - A glyph is rendered on top of the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeOverstrike
	GlyphInscribeOverstrike GlyphInscription = 3
)

// HorizontalDirections - The absolute directions on the horizontal axis.
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
	HorizontalDirectionsLeft HorizontalDirections = 1
	// HorizontalDirectionsRight - The right direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsRight
	HorizontalDirectionsRight HorizontalDirections = 2
)

// ImageCacheMode - Constants that specify the caching policy on a per-image basis.
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

// ImageDynamicRange - Describes how High Dynamic Range (HDR) image content displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type ImageDynamicRange int

const (
	// ImageDynamicRangeConstrainedHigh - Allows for constrained High Dynamic Range (HDR) image content which is useful for mixing HDR and Standard Dynamic Range (SDR) content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/constrainedHigh
	ImageDynamicRangeConstrainedHigh ImageDynamicRange = 1
	// ImageDynamicRangeHigh - Allows image content to use extended dynamic range if it has dynamic range content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/high
	ImageDynamicRangeHigh ImageDynamicRange = 2
	// ImageDynamicRangeUnspecified - Indicates that the dynamic range treatment of the image is unknown or otherwise unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/unspecified
	ImageDynamicRangeUnspecified ImageDynamicRange = -1
)

// ImageLayoutDirection - Constants that describe the layout direction for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection
type ImageLayoutDirection int

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

// ImageLoadStatus - Status values for incremental image loading.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus
type ImageLoadStatus uint

const (
	// ImageLoadStatusCancelled - Image loading was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/cancelled
	ImageLoadStatusCancelled ImageLoadStatus = 1
	// ImageLoadStatusInvalidData - An error occurred during image decompression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/invalidData
	ImageLoadStatusInvalidData ImageLoadStatus = 2
	// ImageLoadStatusReadError - Not enough data was available for full decompression of the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/readError
	ImageLoadStatusReadError ImageLoadStatus = 4
	// ImageLoadStatusUnexpectedEOF - Not enough data was available to fully decompress the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/unexpectedEOF
	ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
)

// ImageResizingMode - Constants that describe the resizing mode for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum
type ImageResizingMode uint

const (
	// ImageResizingModeStretch - The image stretches when it resizes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum/stretch
	ImageResizingModeStretch ImageResizingMode = 1
)

// ImageSymbolScale - Constants that specify which scale variant of a symbol image to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale
type ImageSymbolScale uint

const (
	ImageSymbolScaleSmall ImageSymbolScale = 1
	ImageSymbolScaleMedium ImageSymbolScale = 2
	ImageSymbolScaleLarge ImageSymbolScale = 3
)

// ImageAlignment - Constants used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment
type ImageAlignment uint

const (
	// ImageAlignCenter - Center the image in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignCenter
	ImageAlignCenter ImageAlignment = 0
)

// ImageInterpolation - Constants that specify the interpolation, or image smoothing, behavior used by the image interpolation property.
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

// ImageScaling - Constants that specify a cell’s image scaling behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling
type ImageScaling uint

const (
	// ScaleNone - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleNone
	ScaleNone ImageScaling = 6
	// ScaleProportionally - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleProportionally
	ScaleProportionally ImageScaling = 4
	// ScaleToFit - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleToFit
	ScaleToFit ImageScaling = 5
	// ImageScaleAxesIndependently - Scale each dimension to exactly fit destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleAxesIndependently
	ImageScaleAxesIndependently ImageScaling = 1
	// ImageScaleNone - Do not scale the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleNone
	ImageScaleNone ImageScaling = 2
	// ImageScaleProportionallyDown - If it is too large for the destination, scale the image down while preserving the aspect ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
	ImageScaleProportionallyDown ImageScaling = 0
	// ImageScaleProportionallyUpOrDown - Scale the image to its maximum possible dimensions while both staying within the destination area and preserving its aspect ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyUpOrDown
	ImageScaleProportionallyUpOrDown ImageScaling = 3
)

// LayoutAttribute - The part of the object’s visual representation that should be used to get the value for the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
type LayoutAttribute uint

const (
	LayoutAttributeLeft LayoutAttribute = 1
	LayoutAttributeRight LayoutAttribute = 2
	LayoutAttributeTop LayoutAttribute = 3
	LayoutAttributeBottom LayoutAttribute = 4
	LayoutAttributeLeading LayoutAttribute = 5
	LayoutAttributeTrailing LayoutAttribute = 6
	LayoutAttributeWidth LayoutAttribute = 7
	LayoutAttributeHeight LayoutAttribute = 8
	LayoutAttributeCenterX LayoutAttribute = 9
	LayoutAttributeCenterY LayoutAttribute = 10
	LayoutAttributeLastBaseline LayoutAttribute = 11
	LayoutAttributeFirstBaseline LayoutAttribute = 12
	LayoutAttributeNotAnAttribute LayoutAttribute = 0
)

// LayoutConstraintOrientation - The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
type LayoutConstraintOrientation uint

const (
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	LayoutConstraintOrientationVertical LayoutConstraintOrientation = 1
)

// ControlCharacterAction - Constants that describe actions for control characters.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction
type ControlCharacterAction uint

const (
	// ControlCharacterActionContainerBreak - An action that triggers a break in layout for the current container.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/containerBreak
	ControlCharacterActionContainerBreak ControlCharacterAction = 32
	// ControlCharacterActionHorizontalTab - An action that inserts a horizontal tab.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/horizontalTab
	ControlCharacterActionHorizontalTab ControlCharacterAction = 4
	// ControlCharacterActionLineBreak - An action that causes a line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/lineBreak
	ControlCharacterActionLineBreak ControlCharacterAction = 8
	// ControlCharacterActionParagraphBreak - An action that causes a paragraph break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/paragraphBreak
	ControlCharacterActionParagraphBreak ControlCharacterAction = 16
	// ControlCharacterActionWhitespace - An action that adds whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/whitespace
	ControlCharacterActionWhitespace ControlCharacterAction = 2
	// ControlCharacterActionZeroAdvancement - An action that removes the glyph from layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/zeroAdvancement
	ControlCharacterActionZeroAdvancement ControlCharacterAction = 1
)

// GlyphProperty - Glyph properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty
type GlyphProperty uint

const (
	// GlyphPropertyControlCharacter - A glyph representing a control character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/controlCharacter
	GlyphPropertyControlCharacter GlyphProperty = 2
	// GlyphPropertyElastic - A glyph with a changeable width, such as a white space character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/elastic
	GlyphPropertyElastic GlyphProperty = 4
	// GlyphPropertyNonBaseCharacter - A glyph that combines several properties.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/nonBaseCharacter
	GlyphPropertyNonBaseCharacter GlyphProperty = 8
	// GlyphPropertyNull - The null glyph, which the layout manager ignores.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/null
	GlyphPropertyNull GlyphProperty = 1
)

// TextLayoutOrientation - Constants that describe the text layout orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation
type TextLayoutOrientation uint

const (
	// TextLayoutOrientationHorizontal - Lines render horizontally, each line following the previous from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/horizontal
	TextLayoutOrientationHorizontal TextLayoutOrientation = 0
	// TextLayoutOrientationVertical - Lines render vertically, each line following the previous from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/vertical
	TextLayoutOrientationVertical TextLayoutOrientation = 1
)

// TypesetterBehavior - Constants that determine the layout manager’s behavior during layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
type TypesetterBehavior uint

const (
	// TypesetterBehavior_10_4 - The typesetter behavior introduced in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/behavior_10_4
	TypesetterBehavior_10_4 TypesetterBehavior = 4
)

// LineBreakMode - Constants that specify what happens when a line is too long for a container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
type LineBreakMode uint

const (
	// LineBreakByClipping - The value that indicates lines don’t extend past the edge of the text container.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byClipping
	LineBreakByClipping LineBreakMode = 2
	// LineBreakByWordWrapping - The value that indicates wrapping occurs at word boundaries, unless the word doesn’t fit on a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
	LineBreakByWordWrapping LineBreakMode = 0
)

// LineMovementDirection - The direction in which a line moves.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection
type LineMovementDirection uint

const (
	// LineDoesntMove - Line has no movement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineDoesntMove
	LineDoesntMove LineMovementDirection = 0
	// LineMovesDown - Lines move from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesDown
	LineMovesDown LineMovementDirection = 3
	// LineMovesLeft - Lines move from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesLeft
	LineMovesLeft LineMovementDirection = 1
	// LineMovesRight - Lines move from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesRight
	LineMovesRight LineMovementDirection = 2
	// LineMovesUp - Lines move from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesUp
	LineMovesUp LineMovementDirection = 4
)

// LineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

const (
	// LineSweepLeft - Characters move from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection/NSLineSweepLeft
	LineSweepLeft LineSweepDirection = 0
)

// MultibyteGlyphPacking - A constant for glyph packing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type MultibyteGlyphPacking uint

const (
	NativeShortGlyphPacking MultibyteGlyphPacking = 0
)

// OpenGLContextParameter - Constants that specify context parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter
type OpenGLContextParameter uint

const (
	OpenGLContextParameterSwapInterval OpenGLContextParameter = 0
	OpenGLContextParameterSurfaceOrder OpenGLContextParameter = 1
	OpenGLContextParameterSurfaceOpacity OpenGLContextParameter = 2
	OpenGLContextParameterSurfaceBackingSize OpenGLContextParameter = 3
	OpenGLContextParameterReclaimResources OpenGLContextParameter = 4
	OpenGLContextParameterCurrentRendererID OpenGLContextParameter = 5
	OpenGLContextParameterGPUVertexProcessing OpenGLContextParameter = 6
	OpenGLContextParameterGPUFragmentProcessing OpenGLContextParameter = 7
	OpenGLContextParameterHasDrawable OpenGLContextParameter = 8
	OpenGLContextParameterMPSwapsInFlight OpenGLContextParameter = 9
	OpenGLContextParameterSwapRectangle OpenGLContextParameter = 10
	OpenGLContextParameterSwapRectangleEnable OpenGLContextParameter = 11
	OpenGLContextParameterRasterizationEnable OpenGLContextParameter = 12
	OpenGLContextParameterStateValidation OpenGLContextParameter = 13
	OpenGLContextParameterSurfaceSurfaceVolatile OpenGLContextParameter = 14
)

// OpenGLGlobalOption - Constants that specify OpenGL options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption
type OpenGLGlobalOption uint

const (
	OpenGLGOFormatCacheSize OpenGLGlobalOption = 0
	OpenGLGOClearFormatCache OpenGLGlobalOption = 1
	OpenGLGORetainRenderers OpenGLGlobalOption = 2
	OpenGLGOUseBuildCache OpenGLGlobalOption = 3
	OpenGLGOResetLibrary OpenGLGlobalOption = 4
)

// PageControllerTransitionStyle - These constants control the transition style of the page controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum
type PageControllerTransitionStyle uint

const (
	PageControllerTransitionStyleStackHistory PageControllerTransitionStyle = 0
	PageControllerTransitionStyleStackBook PageControllerTransitionStyle = 1
	PageControllerTransitionStyleHorizontalStrip PageControllerTransitionStyle = 2
)

// PageLayoutResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result
type PageLayoutResult uint

const (
	PageLayoutResultCancelled PageLayoutResult = 0
	PageLayoutResultChanged PageLayoutResult = 1
)

// LineBreakStrategy - Constants that specify how the text system breaks lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type LineBreakStrategy uint

const (
	// LineBreakStrategyNone - The text system doesn’t use any line-break strategies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakStrategy/NSLineBreakStrategyNone
	LineBreakStrategyNone LineBreakStrategy = 0
	// LineBreakStrategyHangulWordPriority - The text system prohibits breaking between Hangul characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/hangulWordPriority
	LineBreakStrategyHangulWordPriority LineBreakStrategy = 2
	// LineBreakStrategyPushOut - The text system pushes out individual lines to avoid an orphan word on the last line of the paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/pushOut
	LineBreakStrategyPushOut LineBreakStrategy = 1
	// LineBreakStrategyStandard - The text system uses the same configuration of line-break strategies that it uses for standard UI labels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/standard
	LineBreakStrategyStandard LineBreakStrategy = 3
)

// TextTabType - Constants that specify the type of tab stop.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
type TextTabType uint

const (
	// CenterTabStopType - A center-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/centerTabStopType
	CenterTabStopType TextTabType = 2
	// LeftTabStopType - A left-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/leftTabStopType
	LeftTabStopType TextTabType = 0
	// RightTabStopType - A right-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/rightTabStopType
	RightTabStopType TextTabType = 1
)

// PasteboardAccessBehavior - A value indicating pasteboard access behavior.
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

// PasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

const (
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
)

// PasteboardReadingOptions - Options that specify how to interpret data on the pasteboard when initializing pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
type PasteboardReadingOptions uint

const (
	// PasteboardReadingAsData - An option to read data from the pasteboard as-is and return it as a data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asData
	PasteboardReadingAsData PasteboardReadingOptions = 0
	// PasteboardReadingAsKeyedArchive - An option to read data from the pasteboard and use it to initialize the object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asKeyedArchive
	PasteboardReadingAsKeyedArchive PasteboardReadingOptions = 4
	// PasteboardReadingAsPropertyList - An option to read data from the pasteboard and unserialize it as a property list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asPropertyList
	PasteboardReadingAsPropertyList PasteboardReadingOptions = 2
	// PasteboardReadingAsString - An option to read data from the pasteboard and convert it to a string object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asString
	PasteboardReadingAsString PasteboardReadingOptions = 1
)

// PasteboardWritingOptions - Type to specify options for writing to a pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type PasteboardWritingOptions uint

const (
	// PasteboardWritingPromised - Data for a type with this option is promised, not immediately written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions/promised
	PasteboardWritingPromised PasteboardWritingOptions = 512
)

// PickerTouchBarItemControlRepresentation - Constants that specify display styles for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum
type PickerTouchBarItemControlRepresentation uint

const (
	PickerTouchBarItemControlRepresentationAutomatic PickerTouchBarItemControlRepresentation = 0
	PickerTouchBarItemControlRepresentationExpanded PickerTouchBarItemControlRepresentation = 1
	PickerTouchBarItemControlRepresentationCollapsed PickerTouchBarItemControlRepresentation = 2
)

// PickerTouchBarItemSelectionMode - Constants that specify selection modes for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum
type PickerTouchBarItemSelectionMode uint

const (
	PickerTouchBarItemSelectionModeSelectOne PickerTouchBarItemSelectionMode = 0
	PickerTouchBarItemSelectionModeSelectAny PickerTouchBarItemSelectionMode = 1
	PickerTouchBarItemSelectionModeMomentary PickerTouchBarItemSelectionMode = 2
)

// PopoverBehavior - The appearance and disappearance behavior of a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
type PopoverBehavior uint

const (
	PopoverBehaviorApplicationDefined PopoverBehavior = 0
	PopoverBehaviorTransient PopoverBehavior = 1
	PopoverBehaviorSemitransient PopoverBehavior = 2
)

// PrintingOrientation - Constants that specify page orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum
type PrintingOrientation uint

const (
	PortraitOrientation PrintingOrientation = 0
	LandscapeOrientation PrintingOrientation = 1
)

// PrintingPaginationMode - Constants that specify the different ways in which an image is divided into pages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
type PrintingPaginationMode uint

const (
	PrintingPaginationModeAutomatic PrintingPaginationMode = 0
	PrintingPaginationModeFit PrintingPaginationMode = 1
	PrintingPaginationModeClip PrintingPaginationMode = 2
)

// PaperOrientation - Constants that describe the orientation of printing on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation
type PaperOrientation uint

const (
	PaperOrientationPortrait PaperOrientation = 0
	PaperOrientationLandscape PaperOrientation = 1
)

// PrintPanelOptions - Constants that specify options for configuring the contents of the main Print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
type PrintPanelOptions uint

const (
	PrintPanelShowsCopies PrintPanelOptions = 1
	PrintPanelShowsPageRange PrintPanelOptions = 2
	PrintPanelShowsPaperSize PrintPanelOptions = 4
	PrintPanelShowsOrientation PrintPanelOptions = 8
	PrintPanelShowsScaling PrintPanelOptions = 16
	PrintPanelShowsPrintSelection PrintPanelOptions = 17
	PrintPanelShowsPageSetupAccessory PrintPanelOptions = 256
	PrintPanelShowsPreview PrintPanelOptions = 131072
)

// PrintPanelResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result
type PrintPanelResult uint

const (
	PrintPanelResultCancelled PrintPanelResult = 0
	PrintPanelResultPrinted PrintPanelResult = 1
)

// ProgressIndicatorStyle - Constants that specify the progress indicator’s style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
type ProgressIndicatorStyle uint

const (
	ProgressIndicatorStyleBar ProgressIndicatorStyle = 0
	ProgressIndicatorStyleSpinning ProgressIndicatorStyle = 1
)

// RulerOrientation - These constants are defined to specify a ruler’s orientation and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum
type RulerOrientation uint

const (
	HorizontalRuler RulerOrientation = 0
	VerticalRuler RulerOrientation = 1
)

// ScrollElasticity - These constants determine the elasticity behavior for an axis of the scrollview.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
type ScrollElasticity uint

const (
	// ScrollElasticityAutomatic - Automatically determine whether to allow elasticity on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
	ScrollElasticityAutomatic ScrollElasticity = 0
)

// ScrollViewFindBarPosition - These constants define the position of the find bar in relation to the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
type ScrollViewFindBarPosition uint

const (
	ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
	ScrollViewFindBarPositionAboveContent ScrollViewFindBarPosition = 1
	ScrollViewFindBarPositionBelowContent ScrollViewFindBarPosition = 2
)

// ScrollerKnobStyle - Specify different knob styles.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum
type ScrollerKnobStyle uint

const (
	ScrollerKnobStyleDefault ScrollerKnobStyle = 0
	ScrollerKnobStyleDark ScrollerKnobStyle = 1
	ScrollerKnobStyleLight ScrollerKnobStyle = 2
)

// ScrubberAlignment - The specified preferred alignment of items within the scrubber, when they come to rest following a user’s scrolling or paging interaction.
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

// ScrubberMode - The scrolling behavior for a scrubber.
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

// SegmentSwitchTracking - The following constants specify the type of tracking behavior a segmented control exhibits. They are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
type SegmentSwitchTracking uint

const (
	SegmentSwitchTrackingSelectOne SegmentSwitchTracking = 0
	SegmentSwitchTrackingSelectAny SegmentSwitchTracking = 1
	SegmentSwitchTrackingMomentary SegmentSwitchTracking = 2
	SegmentSwitchTrackingMomentaryAccelerator SegmentSwitchTracking = 3
)

// SharingCollaborationMode - Represents the types of sharing (collaborating on an item vs. sending a copy of the item)
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type SharingCollaborationMode uint

const (
	SharingCollaborationModeSendCopy SharingCollaborationMode = 0
	SharingCollaborationModeCollaborate SharingCollaborationMode = 1
)

// CloudKitSharingServiceOptions - Constants that describe how a participant can configure a CloudKit share.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions
type CloudKitSharingServiceOptions uint

const (
	CloudKitSharingServiceStandard CloudKitSharingServiceOptions = 0
	CloudKitSharingServiceAllowPublic CloudKitSharingServiceOptions = 1
	CloudKitSharingServiceAllowPrivate CloudKitSharingServiceOptions = 2
	CloudKitSharingServiceAllowReadOnly CloudKitSharingServiceOptions = 16
	CloudKitSharingServiceAllowReadWrite CloudKitSharingServiceOptions = 32
)

// CorrectionIndicatorType - Constants that allow an app to specify the correction indicator type displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType
type CorrectionIndicatorType uint

const (
	// CorrectionIndicatorTypeDefault - The default indicator that shows a proposed correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType/default
	CorrectionIndicatorTypeDefault CorrectionIndicatorType = 0
	// CorrectionIndicatorTypeGuesses - Shows multiple alternatives from which the user may choose the appropriate spelling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType/guesses
	CorrectionIndicatorTypeGuesses CorrectionIndicatorType = 2
	// CorrectionIndicatorTypeReversion - Provides the option to revert to the original form after a correction has been made.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionIndicatorType/reversion
	CorrectionIndicatorTypeReversion CorrectionIndicatorType = 1
)

// CorrectionResponse - The correction response passed to the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse
type CorrectionResponse uint

const (
	// CorrectionResponseAccepted - The user accepted the correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/accepted
	CorrectionResponseAccepted CorrectionResponse = 1
	// CorrectionResponseEdited - After the correction was accepted, the user edited the corrected word (to something other than its original form.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/edited
	CorrectionResponseEdited CorrectionResponse = 4
	// CorrectionResponseIgnored - The user continued in such a way as to ignore the correction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/ignored
	CorrectionResponseIgnored CorrectionResponse = 3
	// CorrectionResponseNone - No response was received from the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/none
	CorrectionResponseNone CorrectionResponse = 0
	// CorrectionResponseRejected - The user rejected the correction by dismissing the correction indicator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/rejected
	CorrectionResponseRejected CorrectionResponse = 2
	// CorrectionResponseReverted - After the correction was accepted, the user reverted the correction back to the original word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/CorrectionResponse/reverted
	CorrectionResponseReverted CorrectionResponse = 5
)

// SpellingState - Constants for the spelling state attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellingState
type SpellingState int

const (
	SpellingStateSpellingFlag SpellingState = 0
	SpellingStateGrammarFlag SpellingState = 1
)

// SplitViewDividerStyle - Constants that specify the style of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type SplitViewDividerStyle uint

const (
	SplitViewDividerStyleThick SplitViewDividerStyle = 1
	SplitViewDividerStyleThin SplitViewDividerStyle = 2
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
)

// SpringLoadingHighlight - A group of constants that indicate a highlighting style for your app’s user interface to display during a spring-loading operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight
type SpringLoadingHighlight uint

const (
	SpringLoadingHighlightNone SpringLoadingHighlight = 0
	SpringLoadingHighlightStandard SpringLoadingHighlight = 1
	SpringLoadingHighlightEmphasized SpringLoadingHighlight = 2
)

// SpringLoadingOptions - These constants denote the type of spring-loading behavior configured for the destination object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions
type SpringLoadingOptions uint

const (
	SpringLoadingDisabled SpringLoadingOptions = 0
	SpringLoadingEnabled SpringLoadingOptions = 1
	SpringLoadingContinuousActivation SpringLoadingOptions = 2
	SpringLoadingNoHover SpringLoadingOptions = 8
)

// StackViewDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum
type StackViewDistribution int

const (
	StackViewDistributionGravityAreas StackViewDistribution = -1
	StackViewDistributionFill StackViewDistribution = 0
	StackViewDistributionFillEqually StackViewDistribution = 1
	StackViewDistributionFillProportionally StackViewDistribution = 2
	StackViewDistributionEqualSpacing StackViewDistribution = 3
	StackViewDistributionEqualCentering StackViewDistribution = 4
)

// StackViewGravity - The gravity areas available in a stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
type StackViewGravity uint

const (
	StackViewGravityTop StackViewGravity = 1
	StackViewGravityLeading StackViewGravity = 1
	StackViewGravityCenter StackViewGravity = 2
	StackViewGravityBottom StackViewGravity = 3
	StackViewGravityTrailing StackViewGravity = 3
)

// StringDrawingOptions - Constants that specify the rendering options for drawing a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions
type StringDrawingOptions int

const (
	StringDrawingUsesLineFragmentOrigin StringDrawingOptions = 1
	StringDrawingUsesFontLeading StringDrawingOptions = 2
	StringDrawingUsesDeviceMetrics StringDrawingOptions = 8
	StringDrawingTruncatesLastVisibleLine StringDrawingOptions = 9
	StringDrawingOptionsResolvesNaturalAlignmentWithBaseWritingDirection StringDrawingOptions = 10
	StringDrawingDisableScreenFontSubstitution StringDrawingOptions = 11
	StringDrawingOneShot StringDrawingOptions = 12
)

// TabPosition enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum
type TabPosition uint

const (
	TabPositionNone TabPosition = 0
	TabPositionTop TabPosition = 1
	TabPositionLeft TabPosition = 2
	TabPositionBottom TabPosition = 3
	TabPositionRight TabPosition = 4
)

// TabViewType - These constants specify the tab view’s type as used by the 
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

// TabViewBorderType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum
type TabViewBorderType uint

const (
	TabViewBorderTypeNone TabViewBorderType = 0
	TabViewBorderTypeLine TabViewBorderType = 1
	TabViewBorderTypeBezel TabViewBorderType = 2
)

// TableViewRowSizeStyle - The row size style constants define the size of the rows in the table view. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum
type TableViewRowSizeStyle int

const (
	// TableViewRowSizeStyleDefault - The table will use the system default layout size: small, medium or large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/default
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
)

// TableViewStyle - Contains the possible style values for a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum
type TableViewStyle uint

const (
	// TableViewStyleFullWidth - The table view style resolves to a full-width style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/fullWidth
	TableViewStyleFullWidth TableViewStyle = 1
)

// TextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type TextAlignment uint

const (
	// TextAlignmentCenter - Text is center-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/center
	TextAlignmentCenter TextAlignment = 1
	// TextAlignmentJustified - Text is justified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/justified
	TextAlignmentJustified TextAlignment = 3
	// TextAlignmentLeft - Text is left-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/left
	TextAlignmentLeft TextAlignment = 0
	// TextAlignmentRight - Text is right-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/right
	TextAlignmentRight TextAlignment = 2
)

// TextBlockDimension - The following constants specify values used by the methods 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension
type TextBlockDimension uint

const (
	// TextBlockHeight - Height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/height
	TextBlockHeight TextBlockDimension = 4
	// TextBlockMaximumHeight - Maximum height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/maximumHeight
	TextBlockMaximumHeight TextBlockDimension = 6
	// TextBlockMaximumWidth - Maximum width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/maximumWidth
	TextBlockMaximumWidth TextBlockDimension = 2
	// TextBlockMinimumHeight - Minimum height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/minimumHeight
	TextBlockMinimumHeight TextBlockDimension = 5
	// TextBlockMinimumWidth - Minimum width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/minimumWidth
	TextBlockMinimumWidth TextBlockDimension = 1
	// TextBlockWidth - Width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/width
	TextBlockWidth TextBlockDimension = 0
)

// TextCursorAccessoryPlacement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement
type TextCursorAccessoryPlacement uint

const (
	TextCursorAccessoryPlacementUnspecified TextCursorAccessoryPlacement = 0
	TextCursorAccessoryPlacementBackward TextCursorAccessoryPlacement = 1
	TextCursorAccessoryPlacementForward TextCursorAccessoryPlacement = 2
	TextCursorAccessoryPlacementInvisible TextCursorAccessoryPlacement = 3
	TextCursorAccessoryPlacementCenter TextCursorAccessoryPlacement = 4
	TextCursorAccessoryPlacementOffscreenLeft TextCursorAccessoryPlacement = 5
	TextCursorAccessoryPlacementOffscreenTop TextCursorAccessoryPlacement = 6
	TextCursorAccessoryPlacementOffscreenRight TextCursorAccessoryPlacement = 7
	TextCursorAccessoryPlacementOffscreenBottom TextCursorAccessoryPlacement = 8
)

// TextInsertionIndicatorAutomaticModeOptions - Options that affect the automatic display mode.
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
	TextInsertionIndicatorAutomaticModeOptionsShowWhileTracking TextInsertionIndicatorAutomaticModeOptions = 2
)

// TextInsertionIndicatorDisplayMode - Constants that determine how to display the system text cursor in a custom text UI.
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
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/visible
	TextInsertionIndicatorDisplayModeVisible TextInsertionIndicatorDisplayMode = 2
)

// TextMovement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement
type TextMovement uint

const (
	TextMovementReturn TextMovement = 16
	TextMovementTab TextMovement = 17
	TextMovementBacktab TextMovement = 18
	TextMovementLeft TextMovement = 19
	TextMovementRight TextMovement = 20
	TextMovementUp TextMovement = 21
	TextMovementDown TextMovement = 22
	TextMovementCancel TextMovement = 23
	TextMovementOther TextMovement = 0
)

// TextScalingType - Constants that specify the text scaling.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextScalingType
type TextScalingType uint

const (
	TextScalingStandard TextScalingType = 0
	TextScalingiOS TextScalingType = 1
)

// TextTableLayoutAlgorithm - These constants, specifying the type of text table layout algorithm, are used with 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum
type TextTableLayoutAlgorithm uint

const (
	// TextTableAutomaticLayoutAlgorithm - Specifies automatic layout algorithm
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum/automaticLayoutAlgorithm
	TextTableAutomaticLayoutAlgorithm TextTableLayoutAlgorithm = 0
	// TextTableFixedLayoutAlgorithm - Specifies fixed layout algorithm
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum/fixedLayoutAlgorithm
	TextTableFixedLayoutAlgorithm TextTableLayoutAlgorithm = 1
)

// TitlebarSeparatorStyle - Styles that determine the type of separator displayed between the title bar and content of a window.
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

// ToolbarDisplayMode - Constants that indicate whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum
type ToolbarDisplayMode uint

const (
	ToolbarDisplayModeDefault ToolbarDisplayMode = 0
	ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
	ToolbarDisplayModeIconOnly ToolbarDisplayMode = 2
	ToolbarDisplayModeLabelOnly ToolbarDisplayMode = 3
)

// ToolbarSizeMode - Constants that specify toolbar display modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum
type ToolbarSizeMode uint

const (
	ToolbarSizeModeDefault ToolbarSizeMode = 0
	ToolbarSizeModeRegular ToolbarSizeMode = 1
	ToolbarSizeModeSmall ToolbarSizeMode = 2
)

// ToolbarItemGroupControlRepresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum
type ToolbarItemGroupControlRepresentation uint

const (
	ToolbarItemGroupControlRepresentationAutomatic ToolbarItemGroupControlRepresentation = 0
	ToolbarItemGroupControlRepresentationExpanded ToolbarItemGroupControlRepresentation = 1
	ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 2
)

// ToolbarItemGroupSelectionMode - A value that indicates how a grouped toolbar item selects its subitems.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum
type ToolbarItemGroupSelectionMode uint

const (
	ToolbarItemGroupSelectionModeSelectOne ToolbarItemGroupSelectionMode = 0
	ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 1
	ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 2
)

// TouchPhase - The possible phases of a touch.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct
type TouchPhase uint

const (
	TouchPhaseBegan TouchPhase = 1
	TouchPhaseMoved TouchPhase = 2
	TouchPhaseStationary TouchPhase = 4
	TouchPhaseEnded TouchPhase = 8
	TouchPhaseCancelled TouchPhase = 16
)

// TouchType - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType
type TouchType uint

const (
	// TouchTypeDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType/direct
	TouchTypeDirect TouchType = 0
)

// TouchTypeMask - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask
type TouchTypeMask uint

const (
	// TouchTypeMaskDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask/direct
	TouchTypeMaskDirect TouchTypeMask = 0
)

// TrackingAreaOptions - The data type defined for the constants specified in the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct
type TrackingAreaOptions uint

const (
	// TrackingActiveAlways - The owner receives messages regardless of first-responder status, window status, or application status. The   message is   sent when the   option is specified along with this constant. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeAlways
	TrackingActiveAlways TrackingAreaOptions = 128
	// TrackingActiveInActiveApp - The owner receives messages when the application is active. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeInActiveApp
	TrackingActiveInActiveApp TrackingAreaOptions = 64
	// TrackingActiveInKeyWindow - The owner receives messages when the view is in the key window. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeInKeyWindow
	TrackingActiveInKeyWindow TrackingAreaOptions = 32
	// TrackingActiveWhenFirstResponder - The owner receives messages when the view is the first responder. This value specifies when the tracking area defined by an   object is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/activeWhenFirstResponder
	TrackingActiveWhenFirstResponder TrackingAreaOptions = 16
	// TrackingAssumeInside - The first event is generated when the cursor leaves the tracking area, regardless if the cursor is inside the area when the   is added to a view.  If this option is not specified, the first event is generated when the cursor leaves the tracking area if the cursor is initially inside the area, or when the cursor enters the area if the cursor is initially outside it. Generally, you do not want to request this behavior. This value specifies a behavior of the tracking area defined by the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/assumeInside
	TrackingAssumeInside TrackingAreaOptions = 256
	// TrackingCursorUpdate - A tracking option that receives events when the mouse cursor enters and exits the tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/cursorUpdate
	TrackingCursorUpdate TrackingAreaOptions = 4
	// TrackingEnabledDuringMouseDrag - The owner receives   events when the mouse cursor is dragged into the tracking area. If this option is not specified, the owner receives mouse-entered events when the mouse is moved (no buttons pressed) into the tracking area and on   events after a mouse drag.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/enabledDuringMouseDrag
	TrackingEnabledDuringMouseDrag TrackingAreaOptions = 1024
	// TrackingInVisibleRect - Mouse tracking occurs only in the visible rectangle of the view—in other words, that region of the tracking rectangle that is unobscured. Otherwise, the entire tracking area is active regardless of overlapping views. The   object is automatically synchronized with changes in the view’s visible area ( ) and the value returned from   is ignored. This value specifies a behavior of the tracking area defined by the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/inVisibleRect
	TrackingInVisibleRect TrackingAreaOptions = 512
	// TrackingMouseEnteredAndExited - The owner of the tracking area receives   when the mouse cursor enters the area and   events when the mouse leaves the area. This value specifies a type of tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/mouseEnteredAndExited
	TrackingMouseEnteredAndExited TrackingAreaOptions = 1
	// TrackingMouseMoved - The owner of the tracking area receives   messages while the mouse cursor is within the area. This value specifies a type of tracking area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct/mouseMoved
	TrackingMouseMoved TrackingAreaOptions = 2
)

// UnderlineStyle - Constants for the underline style and strikethrough style attribute keys.
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
	UnderlineStylePatternSolid UnderlineStyle = 4
	// UnderlineStyleByWord - Draw the line only beneath or through words, not whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/byWord
	UnderlineStyleByWord UnderlineStyle = 9
	// UnderlineStyleDouble - Draw a double line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/double
	UnderlineStyleDouble UnderlineStyle = 3
	// UnderlineStylePatternDash - Draw a line of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDash
	UnderlineStylePatternDash UnderlineStyle = 6
	// UnderlineStylePatternDashDot - Draw a line of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDot
	UnderlineStylePatternDashDot UnderlineStyle = 7
	// UnderlineStylePatternDashDotDot - Draw a line of alternating dashes and two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDotDot
	UnderlineStylePatternDashDotDot UnderlineStyle = 8
	// UnderlineStylePatternDot - Draw a line of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDot
	UnderlineStylePatternDot UnderlineStyle = 5
	// UnderlineStyleSingle - Draw a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/single
	UnderlineStyleSingle UnderlineStyle = 1
	// UnderlineStyleThick - Draw a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/thick
	UnderlineStyleThick UnderlineStyle = 2
)

// UserInterfaceLayoutDirection - Specifies the directional flow of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
type UserInterfaceLayoutDirection uint

const (
	// UserInterfaceLayoutDirectionRightToLeft - Layout direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/rightToLeft
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

// VerticalDirections - The directions on the vertical axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections
type VerticalDirections uint

const (
	// VerticalDirectionsAll - All vertical directions (up and down).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsAll
	VerticalDirectionsAll VerticalDirections = 0
	// VerticalDirectionsDown - The downward direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsDown
	VerticalDirectionsDown VerticalDirections = 2
	// VerticalDirectionsUp - The upwards direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsUp
	VerticalDirectionsUp VerticalDirections = 1
)

// AutoresizingMaskOptions - Constants that specify the autoresizing behaviors for views.
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

// BackgroundStyle - Background styles to apply to a view’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
type BackgroundStyle uint

const (
	BackgroundStyleNormal BackgroundStyle = 0
	BackgroundStyleEmphasized BackgroundStyle = 1
	BackgroundStyleRaised BackgroundStyle = 2
	BackgroundStyleLowered BackgroundStyle = 3
)

// ViewLayerContentsPlacement - These constants specify the location of the layer content when the content is not rerendered in response to view resizing. For more information, see the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum
type ViewLayerContentsPlacement uint

const (
	// ViewLayerContentsPlacementBottom - The content is horizontally centered at the bottom-edge of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/bottom
	ViewLayerContentsPlacementBottom ViewLayerContentsPlacement = 8
	// ViewLayerContentsPlacementTopLeft - The content is positioned in the top-left corner of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/topLeft
	ViewLayerContentsPlacementTopLeft ViewLayerContentsPlacement = 11
)

// ViewLayerContentsRedrawPolicy - Constants that specify how layer resizing is handled when a view is layer-backed or layer-hosting. For more information, see the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum
type ViewLayerContentsRedrawPolicy uint

const (
	// ViewLayerContentsRedrawOnSetNeedsDisplay - Any of the   methods sent to the view will cause the view redraw the affected layer parts by invoking the view’s  , but neither the layer or the view are marked as needing display when the view’s size changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/onSetNeedsDisplay
	ViewLayerContentsRedrawOnSetNeedsDisplay ViewLayerContentsRedrawPolicy = 1
)

// ViewControllerTransitionOptions - Animation options for view transitions in a view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
type ViewControllerTransitionOptions uint

const (
	// ViewControllerTransitionAllowUserInteraction - A transition animation that allows user interaction during the transition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/allowUserInteraction
	ViewControllerTransitionAllowUserInteraction ViewControllerTransitionOptions = 4096
	// ViewControllerTransitionCrossfade - A transition animation that fades the new view in and simultaneously fades the old view out. You can combine this animation option with any of the “slide” options in this enumeration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/crossfade
	ViewControllerTransitionCrossfade ViewControllerTransitionOptions = 1
	// ViewControllerTransitionSlideBackward - A transition animation that reflects the user interface layout direction ( ) in a “backward” manner, as follows
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideBackward
	ViewControllerTransitionSlideBackward ViewControllerTransitionOptions = 384
	// ViewControllerTransitionSlideRight - A transition animation that slides the old view to the right while the new view slides into view from the left.  In other words, both views slide to the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideRight
	ViewControllerTransitionSlideRight ViewControllerTransitionOptions = 128
	// ViewControllerTransitionSlideUp - A transition animation that slides the old view up while the new view comes into view from the bottom.  In other words, both views slide up.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideUp
	ViewControllerTransitionSlideUp ViewControllerTransitionOptions = 16
)

// ViewLayoutRegionAdaptivityAxis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis
type ViewLayoutRegionAdaptivityAxis int

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

// VisualEffectBlendingMode - Constants that specify whether the visual effect view blends with what’s either behind or within the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum
type VisualEffectBlendingMode uint

const (
	VisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = 0
	VisualEffectBlendingModeWithinWindow VisualEffectBlendingMode = 1
)

// VisualEffectMaterial - Constants to specify the material shown by the visual effect view.
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

// VisualEffectState - Constants to specify how the material appearance should reflect window activity state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum
type VisualEffectState uint

const (
	// VisualEffectStateFollowsWindowActiveState - The backdrop should automatically appear active when the window is active, and inactive when it is not.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum/followsWindowActiveState
	VisualEffectStateFollowsWindowActiveState VisualEffectState = 0
)

// WindowAnimationBehavior - Constants that control the automatic window animation behavior windows use when ordering to the front or out of view.
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

// WindowBackingLocation - The following constants and the related data type represent a window’s possible backing locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum
type WindowBackingLocation uint

const (
	WindowBackingLocationDefault WindowBackingLocation = 0
	WindowBackingLocationVideoMemory WindowBackingLocation = 1
	WindowBackingLocationMainMemory WindowBackingLocation = 2
)

// BackingStoreType - Constants that specify how the window device buffers the drawing done in a window.
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

// WindowButton - Constants that provide a way to access standard title bar buttons.
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

// WindowCollectionBehavior - Window collection behaviors related to Mission Control, Spaces, and Stage Manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type WindowCollectionBehavior uint

const (
	// WindowCollectionBehaviorAuxiliary - The behavior marking this window as auxiliary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/auxiliary
	WindowCollectionBehaviorAuxiliary WindowCollectionBehavior = 14
	// WindowCollectionBehaviorCanJoinAllApplications - The behavior marking this window as one that can join all apps for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllApplications
	WindowCollectionBehaviorCanJoinAllApplications WindowCollectionBehavior = 15
	// WindowCollectionBehaviorCanJoinAllSpaces - The window can appear in all spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllSpaces
	WindowCollectionBehaviorCanJoinAllSpaces WindowCollectionBehavior = 1
	// WindowCollectionBehaviorFullScreenAllowsTiling - The window can be a secondary full screen tile even if it can’t be a full screen window itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAllowsTiling
	WindowCollectionBehaviorFullScreenAllowsTiling WindowCollectionBehavior = 11
	// WindowCollectionBehaviorFullScreenAuxiliary - The window displays on the same space as the full screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAuxiliary
	WindowCollectionBehaviorFullScreenAuxiliary WindowCollectionBehavior = 9
	// WindowCollectionBehaviorFullScreenDisallowsTiling - The window doesn’t support being a full-screen tile window, but may support being a full-screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenDisallowsTiling
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 12
	// WindowCollectionBehaviorFullScreenNone - The window doesn’t support full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenNone
	WindowCollectionBehaviorFullScreenNone WindowCollectionBehavior = 10
	// WindowCollectionBehaviorFullScreenPrimary - The window can enter full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenPrimary
	WindowCollectionBehaviorFullScreenPrimary WindowCollectionBehavior = 8
	// WindowCollectionBehaviorIgnoresCycle - The window isn’t part of the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/ignoresCycle
	WindowCollectionBehaviorIgnoresCycle WindowCollectionBehavior = 7
	// WindowCollectionBehaviorManaged - The window participates in Mission Control and Spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/managed
	WindowCollectionBehaviorManaged WindowCollectionBehavior = 3
	// WindowCollectionBehaviorMoveToActiveSpace - When the window becomes active, move it to the active space instead of switching spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/moveToActiveSpace
	WindowCollectionBehaviorMoveToActiveSpace WindowCollectionBehavior = 2
	// WindowCollectionBehaviorParticipatesInCycle - The window participates in the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/participatesInCycle
	WindowCollectionBehaviorParticipatesInCycle WindowCollectionBehavior = 6
	// WindowCollectionBehaviorPrimary - The behavior marking this window as primary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/primary
	WindowCollectionBehaviorPrimary WindowCollectionBehavior = 13
	// WindowCollectionBehaviorStationary - Mission Control doesn’t affect the window, so it stays visible and stationary, like the desktop window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/stationary
	WindowCollectionBehaviorStationary WindowCollectionBehavior = 5
	// WindowCollectionBehaviorTransient - The window floats in Spaces and hides in Mission Control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/transient
	WindowCollectionBehaviorTransient WindowCollectionBehavior = 4
	// WindowCollectionBehaviorDefault - The window appears in only one space at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowCollectionBehavior/NSWindowCollectionBehaviorDefault
	WindowCollectionBehaviorDefault WindowCollectionBehavior = 0
)

// WindowDepth - A type that represents the depth, or amount of memory, for a single pixel in a window or screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth
type WindowDepth uint

const (
	// WindowDepthOnehundredtwentyeightBitRGB - One hundred and twenty eight bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/onehundredtwentyeightBitRGB
	WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 544
	// WindowDepthSixtyfourBitRGB - Sixty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/sixtyfourBitRGB
	WindowDepthSixtyfourBitRGB WindowDepth = 528
	// WindowDepthTwentyfourBitRGB - Twenty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/twentyfourBitRGB
	WindowDepthTwentyfourBitRGB WindowDepth = 520
)

// WindowNumberListOptions - Options to use when retrieving window numbers from the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
type WindowNumberListOptions uint

const (
	// WindowNumberListAllApplications - The window numbers of windows visible on any space and belonging to any application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allApplications
	WindowNumberListAllApplications WindowNumberListOptions = 1
	// WindowNumberListAllSpaces - The window numbers of windows visible on any space and belonging to the calling application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allSpaces
	WindowNumberListAllSpaces WindowNumberListOptions = 16
)

// WindowOcclusionState - Specifies whether the window is occluded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type WindowOcclusionState uint

const (
	// WindowOcclusionStateVisible - If set, at least part of the window is visible; if not set, the entire window is occluded. A window that has a nonrectangular shape can be entirely occluded onscreen, but if its bounding box falls into a visible region, the window is considered to be visible. Note that a completely transparent window may also be considered visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct/visible
	WindowOcclusionStateVisible WindowOcclusionState = 2
)

// WindowOrderingMode - Constants that let you specify how a window is ordered relative to another window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type WindowOrderingMode int

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

// SelectionDirection - Constants that specify the direction a window is currently using to change the key view.
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

// WindowSharingType - Constants that represent the access levels other processes can have to a window’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
type WindowSharingType uint

const (
	// WindowSharingNone - A legacy constant that macOS no longer uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/none
	WindowSharingNone WindowSharingType = 0
	// WindowSharingReadOnly - The window’s contents can be read but not modified by another process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/readOnly
	WindowSharingReadOnly WindowSharingType = 1
)

// WindowStyleMask - Constants that specify the style of a window, and that you can combine with the C bitwise OR operator.
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
	WindowStyleMaskClosable WindowStyleMask = 2
	// WindowStyleMaskDocModalWindow - The window is a document-modal panel (or  a subclass of  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/docModalWindow
	WindowStyleMaskDocModalWindow WindowStyleMask = 64
	// WindowStyleMaskFullScreen - The window can appear full screen. A fullscreen window does not draw its title bar, and may have special handling for its toolbar. (This mask is automatically toggled when   is called.)
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullScreen
	WindowStyleMaskFullScreen WindowStyleMask = 4097
	// WindowStyleMaskFullSizeContentView - When set, the window’s   consumes the full size of the window. Although you can combine this constant with other window style masks, it is respected only for windows with a title bar. Note that using this mask opts in to layer-backing. Use the   or the   to lay out views underneath the title bar–toolbar area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullSizeContentView
	WindowStyleMaskFullSizeContentView WindowStyleMask = 4098
	// WindowStyleMaskHUDWindow - The window is a HUD panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/hudWindow
	WindowStyleMaskHUDWindow WindowStyleMask = 129
	// WindowStyleMaskMiniaturizable - The window displays a minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/miniaturizable
	WindowStyleMaskMiniaturizable WindowStyleMask = 4
	// WindowStyleMaskNonactivatingPanel - The window is a panel or a subclass of   that does not activate the owning app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/nonactivatingPanel
	WindowStyleMaskNonactivatingPanel WindowStyleMask = 128
	// WindowStyleMaskResizable - The window can be resized by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/resizable
	WindowStyleMaskResizable WindowStyleMask = 8
	// WindowStyleMaskTexturedBackground - The window uses a textured background that darkens when the window is key or main and lightens when it is inactive, and may have a second gradient in the section below the window content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/texturedBackground
	WindowStyleMaskTexturedBackground WindowStyleMask = 9
	// WindowStyleMaskTitled - The window displays a title bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/titled
	WindowStyleMaskTitled WindowStyleMask = 1
	// WindowStyleMaskUnifiedTitleAndToolbar - This constant has no effect, because all windows that include a toolbar use the unified style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/unifiedTitleAndToolbar
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 4096
	// WindowStyleMaskUtilityWindow - The window is a panel or a subclass of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/utilityWindow
	WindowStyleMaskUtilityWindow WindowStyleMask = 16
)

// WindowTabbingMode - The preferred tabbing behavior of a window.
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

// WindowTitleVisibility - Specifies the appearance of the window’s title bar area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type WindowTitleVisibility uint

const (
	// WindowTitleHidden - The window hides the title and moves the toolbar up into the area previously occupied by the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/hidden
	WindowTitleHidden WindowTitleVisibility = 1
	// WindowTitleVisible - The window has the regular window title and title bar buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/visible
	WindowTitleVisible WindowTitleVisibility = 0
)

// WindowToolbarStyle - Styles that determine the appearance and location of the toolbar in relation to the title bar.
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

// WindowUserTabbingPreference - A value that indicates the user’s preference for window tabbing.
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

// WorkspaceAuthorizationType - The types of privileged file operations that can be authorized by the user.
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

// WorkspaceIconCreationOptions - Constants that describe options for creating icons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
type WorkspaceIconCreationOptions uint

const (
	// Exclude10_4ElementsIconCreationOption - An option to suppress generation of the new higher resolution icon representations that are supported in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/exclude10_4ElementsIconCreationOption
	Exclude10_4ElementsIconCreationOption WorkspaceIconCreationOptions = 4
	// ExcludeQuickDrawElementsIconCreationOption - An option to suppress generation of the QuickDraw format icon representations that are used in macOS 10.0 through macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/excludeQuickDrawElementsIconCreationOption
	ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 2
)

// WorkspaceLaunchOptions - Constants specifying how you want to launch an app
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions
type WorkspaceLaunchOptions uint

const (
	// WorkspaceLaunchAllowingClassicStartup - Start up the Classic compatibility environment, if it is required by the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/allowingClassicStartup
	WorkspaceLaunchAllowingClassicStartup WorkspaceLaunchOptions = 10
	// WorkspaceLaunchAndHide - Tell the app to hide itself as soon as it finishes launching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andHide
	WorkspaceLaunchAndHide WorkspaceLaunchOptions = 7
	// WorkspaceLaunchAndHideOthers - Hide all apps except the newly launched one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andHideOthers
	WorkspaceLaunchAndHideOthers WorkspaceLaunchOptions = 8
	// WorkspaceLaunchAndPrint - Print items instead of opening them.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andPrint
	WorkspaceLaunchAndPrint WorkspaceLaunchOptions = 0
	// WorkspaceLaunchAsync - Launch the app and return the results asynchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/async
	WorkspaceLaunchAsync WorkspaceLaunchOptions = 5
	// WorkspaceLaunchDefault - Launch the app asynchronously and launch it in the Classic environment, if required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/default
	WorkspaceLaunchDefault WorkspaceLaunchOptions = 9
	// WorkspaceLaunchInhibitingBackgroundOnly - Causes launch to fail if the target is background-only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/inhibitingBackgroundOnly
	WorkspaceLaunchInhibitingBackgroundOnly WorkspaceLaunchOptions = 2
	// WorkspaceLaunchNewInstance - Create a new instance of the app, even if one is already running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/newInstance
	WorkspaceLaunchNewInstance WorkspaceLaunchOptions = 6
	// WorkspaceLaunchPreferringClassic - Force the app to launch in the Classic compatibility environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/preferringClassic
	WorkspaceLaunchPreferringClassic WorkspaceLaunchOptions = 11
	// WorkspaceLaunchWithErrorPresentation - Display an error panel to the user if a failure occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withErrorPresentation
	WorkspaceLaunchWithErrorPresentation WorkspaceLaunchOptions = 1
	// WorkspaceLaunchWithoutActivation - Launch the app but do not bring it into the foreground.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withoutActivation
	WorkspaceLaunchWithoutActivation WorkspaceLaunchOptions = 4
	// WorkspaceLaunchWithoutAddingToRecents - Do not add the app or documents to the Recents menu.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withoutAddingToRecents
	WorkspaceLaunchWithoutAddingToRecents WorkspaceLaunchOptions = 3
)

// WritingDirection - Constants that specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
type WritingDirection int

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

// WritingDirectionFormatType - Constants for the writing direction attribute key.
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
	WritingDirectionOverride WritingDirectionFormatType = 2
)

// WritingToolsBehavior - Constants that specify the Writing Tools experience for the underlying view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior
type WritingToolsBehavior int

const (
	// WritingToolsBehaviorNone - An option to prevent Writing Tools from modifying the text in the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/none
	WritingToolsBehaviorNone WritingToolsBehavior = -1
)

// WritingToolsCoordinatorContextScope - Options that indicate how much of your content Writing Tools requested.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope
type WritingToolsCoordinatorContextScope uint

const (
	WritingToolsCoordinatorContextScopeUserSelection WritingToolsCoordinatorContextScope = 0
	WritingToolsCoordinatorContextScopeFullDocument WritingToolsCoordinatorContextScope = 1
	WritingToolsCoordinatorContextScopeVisibleArea WritingToolsCoordinatorContextScope = 2
)

// WritingToolsCoordinatorState - The states that indicate the current activity, if any, Writing Tools
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

// WritingToolsCoordinatorTextAnimation - The types of animations that Writing Tools performs during an
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation
type WritingToolsCoordinatorTextAnimation uint

const (
	WritingToolsCoordinatorTextAnimationAnticipate WritingToolsCoordinatorTextAnimation = 0
	WritingToolsCoordinatorTextAnimationRemove WritingToolsCoordinatorTextAnimation = 1
	WritingToolsCoordinatorTextAnimationInsert WritingToolsCoordinatorTextAnimation = 2
	WritingToolsCoordinatorTextAnimationAnticipateInactive WritingToolsCoordinatorTextAnimation = 8
	WritingToolsCoordinatorTextAnimationTranslate WritingToolsCoordinatorTextAnimation = 9
)

// WritingToolsCoordinatorTextReplacementReason - Options that indicate whether Writing Tools is animating changes to
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextReplacementReason
type WritingToolsCoordinatorTextReplacementReason uint

const (
	WritingToolsCoordinatorTextReplacementReasonInteractive WritingToolsCoordinatorTextReplacementReason = 0
	WritingToolsCoordinatorTextReplacementReasonNoninteractive WritingToolsCoordinatorTextReplacementReason = 1
)

// WritingToolsCoordinatorTextUpdateReason - Constants that specify the reason you updated your view’s content
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextUpdateReason
type WritingToolsCoordinatorTextUpdateReason uint

const (
	WritingToolsCoordinatorTextUpdateReasonTyping WritingToolsCoordinatorTextUpdateReason = 0
	WritingToolsCoordinatorTextUpdateReasonUndoRedo WritingToolsCoordinatorTextUpdateReason = 1
)

// WritingToolsResultOptions - Constants to specify what type of content to allow in Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions
type WritingToolsResultOptions uint

const (
	WritingToolsResultDefault WritingToolsResultOptions = 0
	WritingToolsResultPlainText WritingToolsResultOptions = 1
	WritingToolsResultRichText WritingToolsResultOptions = 2
	WritingToolsResultList WritingToolsResultOptions = 4
	WritingToolsResultTable WritingToolsResultOptions = 8
	WritingToolsResultPresentationIntent WritingToolsResultOptions = 9
)


