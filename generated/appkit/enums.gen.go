// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Enum types and constants
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
	ColorTypeComponentBased ColorType = 0
	ColorTypePattern ColorType = 1
	ColorTypeCatalog ColorType = 2
)

// ColorSystemEffect - Constants for user interactions that change the appearance of a view or control.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect
type ColorSystemEffect uint

const (
	ColorSystemEffectNone ColorSystemEffect = 0
	ColorSystemEffectPressed ColorSystemEffect = 1
	ColorSystemEffectDeepPressed ColorSystemEffect = 2
	ColorSystemEffectDisabled ColorSystemEffect = 3
	ColorSystemEffectRollover ColorSystemEffect = 4
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

// CompositingOperation - Constants that describe compositing operators in terms of source and destination images, each having an opaque and transparent region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
type CompositingOperation uint

const (
	// CompositingOperationSaturation - Uses the saturation value of the source and the hue and luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/saturation
	CompositingOperationSaturation CompositingOperation = 26
	// CompositingOperationSourceOver - The source image wherever it is opaque, and the destination image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceOver
	CompositingOperationSourceOver CompositingOperation = 2
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
	EventSubtypeWindowExposed EventSubtype = 0
	EventSubtypeApplicationActivated EventSubtype = 1
	EventSubtypeApplicationDeactivated EventSubtype = 2
	EventSubtypeWindowMoved EventSubtype = 4
	EventSubtypeScreenChanged EventSubtype = 8
	EventSubtypePowerOff EventSubtype = 1
	EventSubtypeMouseEvent EventSubtype = 0
	EventSubtypeTabletPoint EventSubtype = 1
	EventSubtypeTabletProximity EventSubtype = 2
	EventSubtypeTouch EventSubtype = 3
)

// EventType - Constants for the types of events that responder objects can handle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType
type EventType uint

const (
	// EventTypeMagnify - The user performed a pinch-open or pinch-close gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/magnify
	EventTypeMagnify EventType = 29
	// EventTypePressure - An event that reports a change in pressure on a pressure-sensitive device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/pressure
	EventTypePressure EventType = 36
	// EventTypeSwipe - The user performed a swipe gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/swipe
	EventTypeSwipe EventType = 30
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
	// EventMaskDirectTouch - A mask for touch events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/directTouch
	EventMaskDirectTouch EventMask = 8
	// EventMaskLeftMouseDragged - A mask for left mouse-dragged events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/leftMouseDragged
	EventMaskLeftMouseDragged EventMask = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/mouseCancelled
	EventMaskMouseCancelled EventMask = 10
	// EventMaskOtherMouseDown - A mask for tertiary mouse-down events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/otherMouseDown
	EventMaskOtherMouseDown EventMask = 0
	// EventMaskPressure - A mask for pressure-change events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/pressure
	EventMaskPressure EventMask = 7
	// EventMaskScrollWheel - A mask for scroll-wheel events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/scrollWheel
	EventMaskScrollWheel EventMask = 0
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
)

// EventModifierFlags - Flags that represent key states in an event object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type EventModifierFlags uint

const (
	// EventModifierFlagControl - The Control key has been pressed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/control
	EventModifierFlagControl EventModifierFlags = 262144
	// EventModifierFlagDeviceIndependentFlagsMask - Device-independent modifier flags are masked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct/deviceIndependentFlagsMask
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 4294901760
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

// FocusRingType - Constants that describe the style of the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type FocusRingType uint

const (
	// FocusRingTypeNone - No focus ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/none
	FocusRingTypeNone FocusRingType = 1
)

// FontDescriptorSymbolicTraits - A symbolic description of the stylistic aspects of a font.
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
	FontDescriptorClassMask FontDescriptorSymbolicTraits = 4026531840
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

// HorizontalDirections - The absolute directions on the horizontal axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections
type HorizontalDirections uint

const (
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
type ImageDynamicRange uint

const (
	// ImageDynamicRangeHigh - Allows image content to use extended dynamic range if it has dynamic range content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/high
	ImageDynamicRangeHigh ImageDynamicRange = 2
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
	ImageLoadStatusCompleted ImageLoadStatus = 0
	ImageLoadStatusCancelled ImageLoadStatus = 1
	ImageLoadStatusInvalidData ImageLoadStatus = 2
	ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
	ImageLoadStatusReadError ImageLoadStatus = 4
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

// PasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

const (
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
)

// PasteboardWritingOptions - Type to specify options for writing to a pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type PasteboardWritingOptions uint

const (
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

// SplitViewDividerStyle - Constants that specify the style of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type SplitViewDividerStyle uint

const (
	SplitViewDividerStyleThick SplitViewDividerStyle = 1
	SplitViewDividerStyleThin SplitViewDividerStyle = 2
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
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
	TextInsertionIndicatorDisplayModeAutomatic TextInsertionIndicatorDisplayMode = 0
	TextInsertionIndicatorDisplayModeHidden TextInsertionIndicatorDisplayMode = 1
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
	WritingToolsBehaviorNone WritingToolsBehavior = -1
	WritingToolsBehaviorDefault WritingToolsBehavior = 0
	WritingToolsBehaviorComplete WritingToolsBehavior = 1
	WritingToolsBehaviorLimited WritingToolsBehavior = 2
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


