// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Enum types and constants
// AnimationEffect - The type for standard system animation effects, which include both display and sound.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type AnimationEffect uint

const (
	AnimationEffectDisappearingItemDefault AnimationEffect = 0
	AnimationEffectPoof AnimationEffect = 10
)

// CursorFrameResizePosition - The position along the perimeter of a rectangular frame (its edges and corners) from which it’s resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition
type CursorFrameResizePosition uint

const (
	CursorFrameResizePositionTop CursorFrameResizePosition = 1
	CursorFrameResizePositionLeft CursorFrameResizePosition = 2
	CursorFrameResizePositionBottom CursorFrameResizePosition = 4
	CursorFrameResizePositionRight CursorFrameResizePosition = 8
)

// CursorFrameResizeDirections - The directions in which a rectangular frame can be resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections
type CursorFrameResizeDirections uint

const (
	CursorFrameResizeDirectionsInward CursorFrameResizeDirections = 1
	CursorFrameResizeDirectionsOutward CursorFrameResizeDirections = 2
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
	HorizontalDirectionsLeft HorizontalDirections = 1
	HorizontalDirectionsRight HorizontalDirections = 2
)

// ImageDynamicRange - Describes how High Dynamic Range (HDR) image content displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type ImageDynamicRange int

const (
	ImageDynamicRangeUnspecified ImageDynamicRange = -1
	ImageDynamicRangeStandard ImageDynamicRange = 0
	ImageDynamicRangeConstrainedHigh ImageDynamicRange = 1
	ImageDynamicRangeHigh ImageDynamicRange = 2
)

// LineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

const (
	LineSweepLeft LineSweepDirection = 0
	LineSweepRight LineSweepDirection = 1
	LineSweepDown LineSweepDirection = 2
	LineSweepUp LineSweepDirection = 3
)

// MultibyteGlyphPacking - A constant for glyph packing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type MultibyteGlyphPacking uint

const (
	NativeShortGlyphPacking MultibyteGlyphPacking = 0
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

// PasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

const (
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
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

// SharingCollaborationMode - Represents the types of sharing (collaborating on an item vs. sending a copy of the item)
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type SharingCollaborationMode uint

const (
	SharingCollaborationModeSendCopy SharingCollaborationMode = 0
	SharingCollaborationModeCollaborate SharingCollaborationMode = 1
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

// VerticalDirections - The directions on the vertical axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections
type VerticalDirections uint

const (
	VerticalDirectionsUp VerticalDirections = 1
	VerticalDirectionsDown VerticalDirections = 2
)

// ViewLayoutRegionAdaptivityAxis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis
type ViewLayoutRegionAdaptivityAxis int

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


