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
	AccessibilityAnnotationPositionStart AccessibilityAnnotationPosition = 0
)

// NSAccessibilityOrientation - Values that indicate the orientation of accessibility elements, such as scroll bars and split views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation
type AccessibilityOrientation uint

const (
	// AccessibilityOrientationHorizontal - The element is oriented horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/horizontal
	AccessibilityOrientationHorizontal AccessibilityOrientation = 0
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
	AccessibilityRulerMarkerTypeIndentHead AccessibilityRulerMarkerType = 0
	// AccessibilityRulerMarkerTypeIndentTail - Tail indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentTail
	AccessibilityRulerMarkerTypeIndentTail AccessibilityRulerMarkerType = 0
	// AccessibilityRulerMarkerTypeTabStopCenter - Center tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopCenter
	AccessibilityRulerMarkerTypeTabStopCenter AccessibilityRulerMarkerType = 0
	// AccessibilityRulerMarkerTypeTabStopDecimal - Decimal tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopDecimal
	AccessibilityRulerMarkerTypeTabStopDecimal AccessibilityRulerMarkerType = 0
)

// NSAccessibilitySortDirection - Values that indicate the sort direction of a column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
type AccessibilitySortDirection uint

// NSAccessibilityUnits - Values that indicate the unit values of a ruler or layout area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
type AccessibilityUnits uint

const (
	// AccessibilityUnitsPicas - The units are picas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/picas
	AccessibilityUnitsPicas AccessibilityUnits = 0
)

// NSAnimationCurve - These constants describe the curve of an animation—that is, the relative speed of an animation from start to finish.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve
type AnimationCurve uint

// NSAnimationEffect - The type for standard system animation effects, which include both display and sound.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type AnimationEffect uint

// NSApplicationActivationPolicy - Activation policies (used by
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum
type ApplicationActivationPolicy uint

const (
	// ApplicationActivationPolicyAccessory - The application doesn’t appear in the Dock and doesn’t have a menu bar, but it may be activated programmatically or by clicking on one of its windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/accessory
	ApplicationActivationPolicyAccessory ApplicationActivationPolicy = 0
	// ApplicationActivationPolicyProhibited - The application doesn’t appear in the Dock and may not create windows or be activated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/prohibited
	ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 0
	// ApplicationActivationPolicyRegular - The application is an ordinary app that appears in the Dock and may have a user interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/regular
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = 0
)

// NSApplicationPrintReply - Constants that indicate the outcome of a print request.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply
type ApplicationPrintReply uint

// NSRemoteNotificationType - These constants determine whether apps launched by remote notifications display a badge.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
type RemoteNotificationType uint

const (
	// RemoteNotificationTypeSound - The app should play a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType/sound
	RemoteNotificationTypeSound RemoteNotificationType = 0
)

// NSBorderType - These constants specify the type of a view’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType
type BorderType uint

const (
	// BezelBorder - A concave border that makes the view look sunken.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/bezelBorder
	BezelBorder BorderType = 0
	// GrooveBorder - A thin border that looks etched around the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/grooveBorder
	GrooveBorder BorderType = 0
	// LineBorder - A black line border around the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
	LineBorder BorderType = 0
)

// NSBoxType - These constants and data type identifies box types, which, in conjunction with a box’s border type, define the appearance of the box.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
type BoxType uint

// NSTitlePosition - Specify the location of a box’s title with respect to its border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum
type TitlePosition uint

// NSBezelStyle - The set of bezel styles to style buttons in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
type BezelStyle uint

const (
	// BezelStyleAccessoryBarAction - A button style that you use for extra actions in an accessory toolbar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/accessoryBarAction
	BezelStyleAccessoryBarAction BezelStyle = 0
	// BezelStyleAutomatic - The default button style based on the button’s contents and position within the window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/automatic
	BezelStyleAutomatic BezelStyle = 0
	// BezelStyleGlass - A bezel style with a glass effect
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/glass
	BezelStyleGlass BezelStyle = 0
	// BezelStyleToolbar - A button style that’s appropriate for a toolbar item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/toolbar
	BezelStyleToolbar BezelStyle = 0
)

// NSGradientType - Specify the gradients used by the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType
type GradientType uint

// NSCellAttribute - Constants for specifying how a button behaves when pressed and how it displays its state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute
type CellAttribute uint

// NSCellType - Constants for specifying how a cell represents its data (as text or as an image).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType
type CellType uint

const (
	// TextCellType - Cell displays text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/textCellType
	TextCellType CellType = 0
)

// NSCellHitResult - Constants used by the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
type CellHitResult uint

// NSCellStyleMask - Constants for specifying what happens when a button is pressed or is displaying its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask
type CellStyleMask uint

// NSCollectionViewDropOperation - These constants specify if acceptance of a drop should be at the item it is dropped on or before the item. These constants are used by the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation
type CollectionViewDropOperation uint

const (
	// CollectionViewDropBefore - The drop occurs before the collection view item to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/before
	CollectionViewDropBefore CollectionViewDropOperation = 0
)

// NSCollectionViewScrollDirection - Constants indicating the scrolling direction for the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection
type CollectionViewScrollDirection uint

// NSCollectionViewScrollPosition - Constants indicating the options for scrolling the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition
type CollectionViewScrollPosition uint

// NSCollectionUpdateAction - Constants indicating the type of action being performed on an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
type CollectionUpdateAction uint

// NSCollectionViewItemHighlightState - Constants indicating the type of highlight applied to an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum
type CollectionViewItemHighlightState uint

// NSColorType - Constants that indicate the color’s type, and which methods may be called on the color object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType
type ColorType uint

const (
	// ColorTypeCatalog - Colors that are retrieved from an asset catalog.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/catalog
	ColorTypeCatalog ColorType = 0
	// ColorTypeComponentBased - Colors that include floating-point color components and a color space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/componentBased
	ColorTypeComponentBased ColorType = 0
	// ColorTypePattern - Colors that include an image to be used as a pattern.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/pattern
	ColorTypePattern ColorType = 0
)

// NSColorSystemEffect - Constants for user interactions that change the appearance of a view or control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type ColorSystemEffect uint

const (
	// ColorSystemEffectDeepPressed - The color that indicates the item received a deep press.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/deepPressed
	ColorSystemEffectDeepPressed ColorSystemEffect = 0
	// ColorSystemEffectDisabled - The color that indicates the item is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/disabled
	ColorSystemEffectDisabled ColorSystemEffect = 0
	// ColorSystemEffectPressed - The color that indicates the item was pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/pressed
	ColorSystemEffectPressed ColorSystemEffect = 0
)

// NSColorPanelMode - A type defined for the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum
type ColorPanelMode uint

// NSColorPanelOptions - The color modes that are enabled for a color panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Options
type ColorPanelOptions uint

// NSColorRenderingIntent - Constants that specify how Cocoa should handle colors that are not located within the destination color space of a graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
type ColorRenderingIntent uint

const (
	// ColorRenderingIntentAbsoluteColorimetric - Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/absoluteColorimetric
	ColorRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 0
	// ColorRenderingIntentDefault - Use the default rendering intent for the graphics context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/default
	ColorRenderingIntentDefault ColorRenderingIntent = 0
	// ColorRenderingIntentPerceptual - Preserve the visual relationship between colors by compressing the gamut of the graphics context to fit inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/perceptual
	ColorRenderingIntentPerceptual ColorRenderingIntent = 0
	// ColorRenderingIntentRelativeColorimetric - Map colors outside of the gamut of the output device to the closest possible match inside the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/relativeColorimetric
	ColorRenderingIntentRelativeColorimetric ColorRenderingIntent = 0
	// ColorRenderingIntentSaturation - Preserve the relative saturation value of the colors when converting into the gamut of the output device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent/saturation
	ColorRenderingIntentSaturation ColorRenderingIntent = 0
)

// NSColorWellStyle - Constants that specify the appearance and interaction modes for a color well.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
type ColorWellStyle uint

const (
	// ColorWellStyleExpanded - A style that supports a color picker popover for fast interactions, and adds a dedicated button to display the color panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
	ColorWellStyleExpanded ColorWellStyle = 0
	// ColorWellStyleMinimal - A style that adds minimal adornments to the color well.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/minimal
	ColorWellStyleMinimal ColorWellStyle = 0
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
)

// NSControlSize - A constant for specifying a cell’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
type ControlSize uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/extraLarge
	ControlSizeExtraLarge ControlSize = 0
	// ControlSizeLarge - A size larger than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/large
	ControlSizeLarge ControlSize = 0
	// ControlSizeMini - The smallest control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/mini
	ControlSizeMini ControlSize = 0
	// ControlSizeRegular - The default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/regular
	ControlSizeRegular ControlSize = 0
	// ControlSizeSmall - A size smaller than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/small
	ControlSizeSmall ControlSize = 0
)

// NSCellImagePosition - A constant for specifying the position of a button’s image relative to its title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition
type CellImagePosition uint

// NSControlTint - Constants for specifying a cell’s tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint
type ControlTint uint

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

// NSDirectionalRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge
type DirectionalRectEdge uint

// NSDisplayGamut enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut
type DisplayGamut uint

// NSDocumentChangeType - Values that indicate a document’s edit status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType
type DocumentChangeType uint

// NSSaveOperationType - Constants for specifying the type of document-save operation to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
type SaveOperationType uint

const (
	// SaveAsOperation - An operation that writes the document’s contents to a new location and updates the document to point to that location
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveAsOperation
	SaveAsOperation SaveOperationType = 0
	// SaveOperation - An operation that overwrites a document’s file or file package with the document’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
	SaveOperation SaveOperationType = 0
)

// NSDragOperation - A group of constants that represent which operations the dragging source can perform on dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation
type DragOperation uint

const (
	// DragOperationAll_Obsolete - The   constant is deprecated. Use   instead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/all_Obsolete
	DragOperationAll_Obsolete DragOperation = 0
	// DragOperationDelete - A constant that indicates the drag can delete the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/delete
	DragOperationDelete DragOperation = 0
	// DragOperationEvery - A constant that indicates that drag can perform all of the drag operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/every
	DragOperationEvery DragOperation = 0
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
	EventTypeMouseCancelled EventType = 0
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
	EventGestureAxisHorizontal EventGestureAxis = 0
	// EventGestureAxisNone - No specific axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/none
	EventGestureAxisNone EventGestureAxis = 0
)

// NSEventModifierFlags - Flags that represent key states in an event object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type EventModifierFlags uint

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

// NSFontCollectionOptions - Constants that support font collection management.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type FontCollectionOptions uint

// NSFontDescriptorSymbolicTraits - A symbolic description of the stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct
type FontDescriptorSymbolicTraits uint

// NSFontPanelModeMask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask
type FontPanelModeMask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask/collection
	FontPanelModeMaskCollection FontPanelModeMask = 0
)

// NSFontRenderingMode - The font rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type FontRenderingMode uint

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

// NSHapticFeedbackPattern - A pattern of haptic feedback to be provided to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern
type HapticFeedbackPattern uint

const (
	// HapticFeedbackPatternAlignment - A haptic feedback pattern to be used in response to the alignment of an object the user is dragging around. For example, this pattern of feedback could be used in a drawing app when the user drags a shape into alignment with another shape. Other scenarios where this type of feedback could be used might include scaling an object to fit within specific dimensions, positioning an object at a preferred location, or reaching the beginning/minimum or end/maximum of something, such as a track view in an audio/video app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/alignment
	HapticFeedbackPatternAlignment HapticFeedbackPattern = 0
	// HapticFeedbackPatternGeneric - A general haptic feedback pattern. Use this when no other feedback patterns apply.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/generic
	HapticFeedbackPatternGeneric HapticFeedbackPattern = 0
	// HapticFeedbackPatternLevelChange - A haptic feedback pattern to be used as the user moves between discrete levels of pressure. This pattern of feedback is used by multilevel accelerator buttons (class  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/levelChange
	HapticFeedbackPatternLevelChange HapticFeedbackPattern = 0
)

// NSHapticFeedbackPerformanceTime - A time at which to provide haptic feedback to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime
type HapticFeedbackPerformanceTime uint

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

// NSImageDynamicRange - Describes how High Dynamic Range (HDR) image content displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type ImageDynamicRange uint

// NSImageResizingMode - Constants that describe the resizing mode for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum
type ImageResizingMode uint

// NSImageSymbolColorRenderingMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode
type ImageSymbolColorRenderingMode uint

// NSImageSymbolScale - Constants that specify which scale variant of a symbol image to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale
type ImageSymbolScale uint

// NSImageSymbolVariableValueMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode
type ImageSymbolVariableValueMode uint

// NSImageScaling - Constants that specify a cell’s image scaling behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling
type ImageScaling uint

// NSLayoutAttribute - The part of the object’s visual representation that should be used to get the value for the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
type LayoutAttribute uint

const (
	// LayoutAttributeLastBaseline - The object’s baseline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/lastBaseline
	LayoutAttributeLastBaseline LayoutAttribute = 0
	// LayoutAttributeLeading - The leading edge of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/leading
	LayoutAttributeLeading LayoutAttribute = 0
	// LayoutAttributeWidth - The width of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/width
	LayoutAttributeWidth LayoutAttribute = 0
)

// NSLayoutConstraintOrientation - The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
type LayoutConstraintOrientation uint

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

// NSLineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

// NSMultibyteGlyphPacking - A constant for glyph packing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type MultibyteGlyphPacking uint

// NSOpenGLGlobalOption - Constants that specify OpenGL options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption
type OpenGLGlobalOption uint

const (
	// OpenGLGOUseBuildCache - Whether to enable the function compilation block cache. This is off by default. It must be enabled at startup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/useBuildCache
	OpenGLGOUseBuildCache OpenGLGlobalOption = 0
)

// NSPDFPanelOptions - Constants used to configure the contents of a PDF panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFPanel/Options-swift.struct
type PDFPanelOptions uint

// NSPageLayoutResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result
type PageLayoutResult uint

// NSLineBreakStrategy - Constants that specify how the text system breaks lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type LineBreakStrategy uint

// NSPasteboardAccessBehavior - A value indicating pasteboard access behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum
type PasteboardAccessBehavior uint

const (
	// PasteboardAccessBehaviorAlwaysAllow - The system will automatically allow all pasteboard access, without notifying the user.  The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysAllow
	PasteboardAccessBehaviorAlwaysAllow PasteboardAccessBehavior = 0
	// PasteboardAccessBehaviorAlwaysDeny - The system will automatically deny all pasteboard access, without notifying the user. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysDeny
	PasteboardAccessBehaviorAlwaysDeny PasteboardAccessBehavior = 0
	// PasteboardAccessBehaviorAsk - The system will notify the user and ask for permission before granting pasteboard access. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/ask
	PasteboardAccessBehaviorAsk PasteboardAccessBehavior = 0
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
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 0
)

// NSPickerTouchBarItemControlRepresentation - Constants that specify display styles for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum
type PickerTouchBarItemControlRepresentation uint

const (
	// PickerTouchBarItemControlRepresentationCollapsed - The system displays the control’s options through a popover.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum/collapsed
	PickerTouchBarItemControlRepresentationCollapsed PickerTouchBarItemControlRepresentation = 0
)

// NSPickerTouchBarItemSelectionMode - Constants that specify selection modes for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum
type PickerTouchBarItemSelectionMode uint

const (
	// PickerTouchBarItemSelectionModeSelectAny - A mode in which a person can select one or more options in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectAny
	PickerTouchBarItemSelectionModeSelectAny PickerTouchBarItemSelectionMode = 0
	// PickerTouchBarItemSelectionModeSelectOne - A mode in which a person can only select one option in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectOne
	PickerTouchBarItemSelectionModeSelectOne PickerTouchBarItemSelectionMode = 0
)

// NSPopoverAppearance - The set of predefined appearances for a popover.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopover/Appearance-swift.enum
type PopoverAppearance uint

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

// NSPaperOrientation - Constants that describe the orientation of printing on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation
type PaperOrientation uint

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
	PrintRenderingQualityResponsive PrintRenderingQuality = 0
)

// NSPrintPanelOptions - Constants that specify options for configuring the contents of the main Print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
type PrintPanelOptions uint

// NSProgressIndicatorStyle - Constants that specify the progress indicator’s style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
type ProgressIndicatorStyle uint

// NSProgressIndicatorThickness - Specify the height of a progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness
type ProgressIndicatorThickness uint

// NSRectAlignment - Constants that specify alignment to an edge or a set of edges depending on the user interface layout direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment
type RectAlignment uint

// NSRulerOrientation - These constants are defined to specify a ruler’s orientation and are used by
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum
type RulerOrientation uint

// NSScrollElasticity - These constants determine the elasticity behavior for an axis of the scrollview.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
type ScrollElasticity uint

const (
	// ScrollElasticityAllowed - Allow content to be scrolled past its bounds on this axis in an elastic fashion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/allowed
	ScrollElasticityAllowed ScrollElasticity = 0
	// ScrollElasticityAutomatic - Automatically determine whether to allow elasticity on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
	ScrollElasticityAutomatic ScrollElasticity = 0
	// ScrollElasticityNone - Disallow scrolling beyond document bounds on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/none
	ScrollElasticityNone ScrollElasticity = 0
)

// NSScrollViewFindBarPosition - These constants define the position of the find bar in relation to the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
type ScrollViewFindBarPosition uint

// NSScrubberAlignment - The specified preferred alignment of items within the scrubber, when they come to rest following a user’s scrolling or paging interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
type ScrubberAlignment uint

const (
	// ScrubberAlignmentCenter - Center alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/center
	ScrubberAlignmentCenter ScrubberAlignment = 0
	// ScrubberAlignmentLeading - Leading alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/leading
	ScrubberAlignmentLeading ScrubberAlignment = 0
	// ScrubberAlignmentNone - No preference for item alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
	ScrubberAlignmentNone ScrubberAlignment = 0
	// ScrubberAlignmentTrailing - Trailing alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/trailing
	ScrubberAlignmentTrailing ScrubberAlignment = 0
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
	ScrubberModeFree ScrubberMode = 0
)

// NSSegmentSwitchTracking - The following constants specify the type of tracking behavior a segmented control exhibits. They are used by
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
type SegmentSwitchTracking uint

const (
	// SegmentSwitchTrackingMomentary - A segment is selected only when the user is pressing the mouse down within the bounds of the segment. When the mouse is no longer down within the segment, the segment is automatically deselected. A momentary segmented control sends an action when the user clicks a segment, and another action when the user releases the segment. If configured as continuous (see  ), the control also sends actions at repeating intervals until the user releases the segment, at which point the control sends its final action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentary
	SegmentSwitchTrackingMomentary SegmentSwitchTracking = 0
)

// NSSharingCollaborationMode - Represents the types of sharing (collaborating on an item vs. sending a copy of the item)
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type SharingCollaborationMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/collaborate
	SharingCollaborationModeCollaborate SharingCollaborationMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/sendCopy
	SharingCollaborationModeSendCopy SharingCollaborationMode = 0
)

// NSSliderType - The types of sliders, used by
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
type SliderType uint

// NSSplitViewDividerStyle - Constants that specify the style of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type SplitViewDividerStyle uint

// NSSpringLoadingHighlight - A group of constants that indicate a highlighting style for your app’s user interface to display during a spring-loading operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight
type SpringLoadingHighlight uint

// NSStackViewDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum
type StackViewDistribution uint

// NSStackViewGravity - The gravity areas available in a stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
type StackViewGravity uint

// NSStringDrawingOptions - Constants that specify the rendering options for drawing a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions
type StringDrawingOptions uint

const (
	// StringDrawingTruncatesLastVisibleLine - Truncates and adds the ellipsis character to the last visible line if the text doesn’t fit into the specified bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingTruncatesLastVisibleLine
	StringDrawingTruncatesLastVisibleLine StringDrawingOptions = 0
	// StringDrawingUsesDeviceMetrics - Uses image glyph bounds instead of typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingUsesDeviceMetrics
	StringDrawingUsesDeviceMetrics StringDrawingOptions = 0
)

// NSTabPosition enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum
type TabPosition uint

// NSTabViewType - These constants specify the tab view’s type as used by the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType
type TabViewType uint

// NSTabViewBorderType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum
type TabViewBorderType uint

// NSTableColumnResizingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions
type TableColumnResizingOptions uint

// NSTableViewAnimationOptions - Specifies the animation effects to apply when inserting or removing rows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
type TableViewAnimationOptions uint

// NSTableViewColumnAutoresizingStyle - The following constants specify the autoresizing styles. These constants are used by the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum
type TableViewColumnAutoresizingStyle uint

// NSTableViewDraggingDestinationFeedbackStyle - These constants specify the drag styles displayed by the table view. They’re used by
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum
type TableViewDraggingDestinationFeedbackStyle uint

const (
	// TableViewDraggingDestinationFeedbackStyleSourceList - Draws an outline on drop target rows, and an insertion marker between rows. This style will automatically be set for source lists when the table’s   is set to  . This is the standard look for Source Lists, but may be used in other areas as needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/sourceList
	TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 0
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
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = 0
	// TableViewRowSizeStyleLarge - The table will use a row height specified for a large table. It is required that the size be fully tested and supported if   is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/large
	TableViewRowSizeStyleLarge TableViewRowSizeStyle = 0
)

// NSTableViewStyle - Contains the possible style values for a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum
type TableViewStyle uint

const (
	// TableViewStyleSourceList - The table view style resolves to a source-list style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/sourceList
	TableViewStyleSourceList TableViewStyle = 0
)

// NSTextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type TextAlignment uint

// NSTextCursorAccessoryPlacement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement
type TextCursorAccessoryPlacement uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/backward
	TextCursorAccessoryPlacementBackward TextCursorAccessoryPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/center
	TextCursorAccessoryPlacementCenter TextCursorAccessoryPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenLeft
	TextCursorAccessoryPlacementOffscreenLeft TextCursorAccessoryPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenTop
	TextCursorAccessoryPlacementOffscreenTop TextCursorAccessoryPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/unspecified
	TextCursorAccessoryPlacementUnspecified TextCursorAccessoryPlacement = 0
)

// NSTextInputTraitType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputTraitType
type TextInputTraitType uint

// NSTextInsertionIndicatorAutomaticModeOptions - Options that affect the automatic display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct
type TextInsertionIndicatorAutomaticModeOptions uint

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
	TextInsertionIndicatorDisplayModeHidden TextInsertionIndicatorDisplayMode = 0
)

// NSTextLayoutManagerSegmentOptions - Values that describe where and how the framework extends segments of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions
type TextLayoutManagerSegmentOptions uint

// NSTextLayoutManagerSegmentType - Values that describe the rendering of selection boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType
type TextLayoutManagerSegmentType uint

// NSTextMovement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement
type TextMovement uint

// NSTintProminence - Controls how strongly the tint color applies in a view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence
type TintProminence uint

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
	TitlebarSeparatorStyleLine TitlebarSeparatorStyle = 0
	// TitlebarSeparatorStyleNone - A style indicating that there’s no title bar separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/none
	TitlebarSeparatorStyleNone TitlebarSeparatorStyle = 0
	// TitlebarSeparatorStyleShadow - A style indicating that the title bar separator is a shadow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/shadow
	TitlebarSeparatorStyleShadow TitlebarSeparatorStyle = 0
)

// NSTokenStyle - The NSTokenStyle constants define how tokens are displayed and editable in the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
type TokenStyle uint

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
	ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 0
	// ToolbarDisplayModeIconOnly - The toolbar displays only an icon for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/iconOnly
	ToolbarDisplayModeIconOnly ToolbarDisplayMode = 0
	// ToolbarDisplayModeLabelOnly - The toolbar displays only a label for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/labelOnly
	ToolbarDisplayModeLabelOnly ToolbarDisplayMode = 0
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
	ToolbarSizeModeRegular ToolbarSizeMode = 0
	// ToolbarSizeModeSmall - The toolbar uses small-sized controls and 24 by 24 pixel icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/small
	ToolbarSizeModeSmall ToolbarSizeMode = 0
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
	ToolbarItemStyleProminent ToolbarItemStyle = 0
)

// NSToolbarItemGroupControlRepresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum
type ToolbarItemGroupControlRepresentation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum/collapsed
	ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 0
)

// NSToolbarItemGroupSelectionMode - A value that indicates how a grouped toolbar item selects its subitems.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum
type ToolbarItemGroupSelectionMode uint

const (
	// ToolbarItemGroupSelectionModeMomentary - The system temporarily highlights the select group item when the user selects the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/momentary
	ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 0
	// ToolbarItemGroupSelectionModeSelectAny - The system toggles a highlight on any item selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectAny
	ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 0
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

// NSUserInterfaceLayoutDirection - Specifies the directional flow of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
type UserInterfaceLayoutDirection uint

const (
	// UserInterfaceLayoutDirectionRightToLeft - Layout direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/rightToLeft
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 0
)

// NSUserInterfaceLayoutOrientation - The stack view layout directions, and user interface axes for hugging priority and clipping resistance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
type UserInterfaceLayoutOrientation uint

const (
	// UserInterfaceLayoutOrientationVertical - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation/vertical
	UserInterfaceLayoutOrientationVertical UserInterfaceLayoutOrientation = 0
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

// NSBackgroundStyle - Background styles to apply to a view’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
type BackgroundStyle uint

// NSViewLayerContentsPlacement - These constants specify the location of the layer content when the content is not rerendered in response to view resizing. For more information, see the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum
type ViewLayerContentsPlacement uint

// NSViewLayerContentsRedrawPolicy - Constants that specify how layer resizing is handled when a view is layer-backed or layer-hosting. For more information, see the
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum
type ViewLayerContentsRedrawPolicy uint

// NSViewControllerTransitionOptions - Animation options for view transitions in a view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
type ViewControllerTransitionOptions uint

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

// NSVisualEffectMaterial - Constants to specify the material shown by the visual effect view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum
type VisualEffectMaterial uint

const (
	// VisualEffectMaterialSidebar - The material for the background of window sidebars.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/sidebar
	VisualEffectMaterialSidebar VisualEffectMaterial = 0
)

// NSVisualEffectState - Constants to specify how the material appearance should reflect window activity state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum
type VisualEffectState uint

// NSWindowAnimationBehavior - Constants that control the automatic window animation behavior windows use when ordering to the front or out of view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
type WindowAnimationBehavior uint

const (
	// WindowAnimationBehaviorAlertPanel - The animation behavior that’s appropriate to the alert window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/alertPanel
	WindowAnimationBehaviorAlertPanel WindowAnimationBehavior = 0
	// WindowAnimationBehaviorDefault - The automatic animation that’s appropriate to the window type. This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/default
	WindowAnimationBehaviorDefault WindowAnimationBehavior = 0
	// WindowAnimationBehaviorDocumentWindow - The animation behavior that’s appropriate to the document window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/documentWindow
	WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 0
	// WindowAnimationBehaviorNone - No automatic animation used. This may be useful when you perform your own window animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/none
	WindowAnimationBehaviorNone WindowAnimationBehavior = 0
	// WindowAnimationBehaviorUtilityWindow - The animation behavior that’s appropriate to the utility window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/utilityWindow
	WindowAnimationBehaviorUtilityWindow WindowAnimationBehavior = 0
)

// NSWindowBackingLocation - The following constants and the related data type represent a window’s possible backing locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum
type WindowBackingLocation uint

// NSBackingStoreType - Constants that specify how the window device buffers the drawing done in a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
type BackingStoreType uint

const (
	// BackingStoreBuffered - The window renders all drawing into a display buffer and then flushes it to the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/buffered
	BackingStoreBuffered BackingStoreType = 0
	// BackingStoreNonretained - The window draws directly to the screen without using any buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/nonretained
	BackingStoreNonretained BackingStoreType = 0
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
	WindowDocumentIconButton WindowButton = 0
	// WindowDocumentVersionsButton - The document versions button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/documentVersionsButton
	WindowDocumentVersionsButton WindowButton = 0
	// WindowMiniaturizeButton - The minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/miniaturizeButton
	WindowMiniaturizeButton WindowButton = 0
	// WindowToolbarButton - The toolbar button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/toolbarButton
	WindowToolbarButton WindowButton = 0
	// WindowZoomButton - The zoom button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/zoomButton
	WindowZoomButton WindowButton = 0
)

// NSWindowCollectionBehavior - Window collection behaviors related to Mission Control, Spaces, and Stage Manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type WindowCollectionBehavior uint

const (
	// WindowCollectionBehaviorAuxiliary - The behavior marking this window as auxiliary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/auxiliary
	WindowCollectionBehaviorAuxiliary WindowCollectionBehavior = 0
	// WindowCollectionBehaviorCanJoinAllApplications - The behavior marking this window as one that can join all apps for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllApplications
	WindowCollectionBehaviorCanJoinAllApplications WindowCollectionBehavior = 0
	// WindowCollectionBehaviorCanJoinAllSpaces - The window can appear in all spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllSpaces
	WindowCollectionBehaviorCanJoinAllSpaces WindowCollectionBehavior = 0
	// WindowCollectionBehaviorFullScreenAllowsTiling - The window can be a secondary full screen tile even if it can’t be a full screen window itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAllowsTiling
	WindowCollectionBehaviorFullScreenAllowsTiling WindowCollectionBehavior = 0
	// WindowCollectionBehaviorFullScreenAuxiliary - The window displays on the same space as the full screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAuxiliary
	WindowCollectionBehaviorFullScreenAuxiliary WindowCollectionBehavior = 0
	// WindowCollectionBehaviorFullScreenDisallowsTiling - The window doesn’t support being a full-screen tile window, but may support being a full-screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenDisallowsTiling
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 0
	// WindowCollectionBehaviorFullScreenNone - The window doesn’t support full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenNone
	WindowCollectionBehaviorFullScreenNone WindowCollectionBehavior = 0
	// WindowCollectionBehaviorFullScreenPrimary - The window can enter full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenPrimary
	WindowCollectionBehaviorFullScreenPrimary WindowCollectionBehavior = 0
	// WindowCollectionBehaviorIgnoresCycle - The window isn’t part of the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/ignoresCycle
	WindowCollectionBehaviorIgnoresCycle WindowCollectionBehavior = 0
	// WindowCollectionBehaviorManaged - The window participates in Mission Control and Spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/managed
	WindowCollectionBehaviorManaged WindowCollectionBehavior = 0
	// WindowCollectionBehaviorMoveToActiveSpace - When the window becomes active, move it to the active space instead of switching spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/moveToActiveSpace
	WindowCollectionBehaviorMoveToActiveSpace WindowCollectionBehavior = 0
	// WindowCollectionBehaviorParticipatesInCycle - The window participates in the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/participatesInCycle
	WindowCollectionBehaviorParticipatesInCycle WindowCollectionBehavior = 0
	// WindowCollectionBehaviorPrimary - The behavior marking this window as primary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/primary
	WindowCollectionBehaviorPrimary WindowCollectionBehavior = 0
	// WindowCollectionBehaviorStationary - Mission Control doesn’t affect the window, so it stays visible and stationary, like the desktop window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/stationary
	WindowCollectionBehaviorStationary WindowCollectionBehavior = 0
	// WindowCollectionBehaviorTransient - The window floats in Spaces and hides in Mission Control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/transient
	WindowCollectionBehaviorTransient WindowCollectionBehavior = 0
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
	WindowNumberListAllApplications WindowNumberListOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allSpaces
	WindowNumberListAllSpaces WindowNumberListOptions = 0
)

// NSWindowOcclusionState - Specifies whether the window is occluded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type WindowOcclusionState uint

const (
	// WindowOcclusionStateVisible - If set, at least part of the window is visible; if not set, the entire window is occluded. A window that has a nonrectangular shape can be entirely occluded onscreen, but if its bounding box falls into a visible region, the window is considered to be visible. Note that a completely transparent window may also be considered visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct/visible
	WindowOcclusionStateVisible WindowOcclusionState = 0
)

// NSWindowOrderingMode - Constants that let you specify how a window is ordered relative to another window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type WindowOrderingMode uint

const (
	// WindowAbove - Moves the window above the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/above
	WindowAbove WindowOrderingMode = 0
	// WindowBelow - Moves the window below the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/below
	WindowBelow WindowOrderingMode = 0
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
	SelectingNext SelectionDirection = 0
	// SelectingPrevious - The window is proceeding to the previous valid key view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/selectingPrevious
	SelectingPrevious SelectionDirection = 0
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
	WindowSharingReadOnly WindowSharingType = 0
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
	WindowStyleMaskClosable WindowStyleMask = 0
	// WindowStyleMaskDocModalWindow - The window is a document-modal panel (or  a subclass of  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/docModalWindow
	WindowStyleMaskDocModalWindow WindowStyleMask = 0
	// WindowStyleMaskFullScreen - The window can appear full screen. A fullscreen window does not draw its title bar, and may have special handling for its toolbar. (This mask is automatically toggled when   is called.)
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullScreen
	WindowStyleMaskFullScreen WindowStyleMask = 0
	// WindowStyleMaskFullSizeContentView - When set, the window’s   consumes the full size of the window. Although you can combine this constant with other window style masks, it is respected only for windows with a title bar. Note that using this mask opts in to layer-backing. Use the   or the   to lay out views underneath the title bar–toolbar area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullSizeContentView
	WindowStyleMaskFullSizeContentView WindowStyleMask = 0
	// WindowStyleMaskHUDWindow - The window is a HUD panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/hudWindow
	WindowStyleMaskHUDWindow WindowStyleMask = 0
	// WindowStyleMaskMiniaturizable - The window displays a minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/miniaturizable
	WindowStyleMaskMiniaturizable WindowStyleMask = 0
	// WindowStyleMaskNonactivatingPanel - The window is a panel or a subclass of   that does not activate the owning app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/nonactivatingPanel
	WindowStyleMaskNonactivatingPanel WindowStyleMask = 0
	// WindowStyleMaskResizable - The window can be resized by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/resizable
	WindowStyleMaskResizable WindowStyleMask = 0
	// WindowStyleMaskTexturedBackground - The window uses a textured background that darkens when the window is key or main and lightens when it is inactive, and may have a second gradient in the section below the window content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/texturedBackground
	WindowStyleMaskTexturedBackground WindowStyleMask = 0
	// WindowStyleMaskTitled - The window displays a title bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/titled
	WindowStyleMaskTitled WindowStyleMask = 0
	// WindowStyleMaskUnifiedTitleAndToolbar - This constant has no effect, because all windows that include a toolbar use the unified style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/unifiedTitleAndToolbar
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 0
	// WindowStyleMaskUtilityWindow - The window is a panel or a subclass of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/utilityWindow
	WindowStyleMaskUtilityWindow WindowStyleMask = 0
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
	WindowTabbingModeDisallowed WindowTabbingMode = 0
	// WindowTabbingModePreferred - A window that explicitly prefers to tab together with other windows with the same tabbing identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/preferred
	WindowTabbingModePreferred WindowTabbingMode = 0
)

// NSWindowTitleVisibility - Specifies the appearance of the window’s title bar area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type WindowTitleVisibility uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/hidden
	WindowTitleHidden WindowTitleVisibility = 0
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
	WindowToolbarStyleExpanded WindowToolbarStyle = 0
	// WindowToolbarStylePreference - A style indicating that the toolbar appears below the window title with toolbar items centered in the toolbar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/preference
	WindowToolbarStylePreference WindowToolbarStyle = 0
	// WindowToolbarStyleUnified - A style indicating that the toolbar appears next to the window title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/unified
	WindowToolbarStyleUnified WindowToolbarStyle = 0
	// WindowToolbarStyleUnifiedCompact - A style indicating that the toolbar appears next to the window title and with reduced margins to allow more focus on the window’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ToolbarStyle-swift.enum/unifiedCompact
	WindowToolbarStyleUnifiedCompact WindowToolbarStyle = 0
)

// NSWindowUserTabbingPreference - A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
type WindowUserTabbingPreference uint

const (
	// WindowUserTabbingPreferenceAlways - A value that indicates a window should always display as tabs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/always
	WindowUserTabbingPreferenceAlways WindowUserTabbingPreference = 0
	// WindowUserTabbingPreferenceInFullScreen - A value that indicates a window should only display as tabs when in full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/inFullScreen
	WindowUserTabbingPreferenceInFullScreen WindowUserTabbingPreference = 0
	// WindowUserTabbingPreferenceManual - A value that indicates a window should display as tabs according to the window’s tabbing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/manual
	WindowUserTabbingPreferenceManual WindowUserTabbingPreference = 0
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
	WritingDirectionNatural WritingDirection = 0
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
	WritingToolsBehaviorComplete WritingToolsBehavior = 0
	// WritingToolsBehaviorNone - An option to prevent Writing Tools from modifying the text in the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/none
	WritingToolsBehaviorNone WritingToolsBehavior = 0
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
	WritingToolsCoordinatorStateInteractiveResting WritingToolsCoordinatorState = 0
	// WritingToolsCoordinatorStateInteractiveStreaming - A state that indicates Writing Tools is processing a request and   incorporating changes interactively into your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/interactiveStreaming
	WritingToolsCoordinatorStateInteractiveStreaming WritingToolsCoordinatorState = 0
	// WritingToolsCoordinatorStateNoninteractive - A state that indicates Writing Tools is handling interactions in   the system UI, instead of in your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/noninteractive
	WritingToolsCoordinatorStateNoninteractive WritingToolsCoordinatorState = 0
)

// NSWritingToolsResultOptions - Constants to specify what type of content to allow in Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions
type WritingToolsResultOptions uint
