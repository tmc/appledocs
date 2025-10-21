// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

// Enum types and constants
// PKCanvasViewDrawingPolicy - Constants that you use to specify the type of drawing gestures your app permits while the user draws on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy
type CanvasViewDrawingPolicy uint

const (
	// CanvasViewDrawingPolicyAnyInput - Allows drawing on the canvas from any input source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/anyInput
	CanvasViewDrawingPolicyAnyInput CanvasViewDrawingPolicy = 0
	// CanvasViewDrawingPolicyDefault - The default input type to use for drawing on a canvas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/default
	CanvasViewDrawingPolicyDefault CanvasViewDrawingPolicy = 0
	// CanvasViewDrawingPolicyPencilOnly - Pencil touches are the only input that draw on the canvas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/pencilOnly
	CanvasViewDrawingPolicyPencilOnly CanvasViewDrawingPolicy = 0
)

// PKContentVersion - Constants that represent versions of PencilKit for backward compatibility.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion
type ContentVersion uint

const (
	// ContentVersionLatest - A property that returns latest version of PencilKit, which supports all currently available inks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/latest
	ContentVersionLatest ContentVersion = 0
	// ContentVersion1 - The PencilKit version that supports inks from iPadOS 14 and earlier, including marker, pen, and pencil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version1
	ContentVersion1 ContentVersion = 0
	// ContentVersion2 - The PencilKit version that supports inks from iPadOS 17 and earlier, including marker, pen, pencil, monoline, fountain pen, watercolor, and crayon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version2
	ContentVersion2 ContentVersion = 0
	// ContentVersion3 - The PencilKit version that supports barrel-roll angle data in inks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version3
	ContentVersion3 ContentVersion = 0
	// ContentVersion4 - New Reed Pen
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version4
	ContentVersion4 ContentVersion = 0
)

// PKEraserType - Constants that indicate the behavior of the eraser.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType
type EraserType uint

const (
	// EraserTypeBitmap - An eraser that removes only those portions of the drawing it touches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeBitmap
	EraserTypeBitmap EraserType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeFixedWidthBitmap
	EraserTypeFixedWidthBitmap EraserType = 0
	// EraserTypeVector - An eraser that removes an entire drawn line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeVector
	EraserTypeVector EraserType = 0
)

// PKToolPickerCustomItemControlOptions - Options for which controls to present.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions
type ToolPickerCustomItemControlOptions uint

const (
	// ToolPickerCustomItemControlOpacity - Present an opacity control if color adjustment is supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions/opacity
	ToolPickerCustomItemControlOpacity ToolPickerCustomItemControlOptions = 0
	// ToolPickerCustomItemControlWidth - Present a width control if width adjustment is supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions/width
	ToolPickerCustomItemControlWidth ToolPickerCustomItemControlOptions = 0
	// ToolPickerCustomItemControlNone - Present neither a width nor opacity control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemControlOptions/PKToolPickerCustomItemControlNone
	ToolPickerCustomItemControlNone ToolPickerCustomItemControlOptions = 0
)

// PKToolPickerVisibility - The visibility state of a tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility
type ToolPickerVisibility uint

const (
	// ToolPickerVisibilityInherited - Inherit the tool picker visibility from the next responder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/PKToolPickerVisibilityInherited
	ToolPickerVisibilityInherited ToolPickerVisibility = 0
	// ToolPickerVisibilityHidden - Tool picker is active but offscreen, and can appear temporarily in response to user actions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/hidden
	ToolPickerVisibilityHidden ToolPickerVisibility = 0
	// ToolPickerVisibilityInactive - No active tool picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/inactive
	ToolPickerVisibilityInactive ToolPickerVisibility = 0
	// ToolPickerVisibilityVisible - Tool picker is active and onscreen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/visible
	ToolPickerVisibilityVisible ToolPickerVisibility = 0
)


