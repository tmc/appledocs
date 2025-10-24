// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

/* debug [enums.gen.go]: Generating 5 enums for PencilKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum PKContentVersion (5 cases) */
// PKContentVersion - Constants that represent versions of PencilKit for backward compatibility.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion
type PKContentVersion uint

const (
	// PKContentVersionLatest - A property that returns latest version of PencilKit, which supports all currently available inks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/latest
	PKContentVersionLatest PKContentVersion = 0
	// PKContentVersion1 - The PencilKit version that supports inks from iPadOS 14 and earlier, including marker, pen, and pencil.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version1
	PKContentVersion1 PKContentVersion = 0
	// PKContentVersion2 - The PencilKit version that supports inks from iPadOS 17 and earlier, including marker, pen, pencil, monoline, fountain pen, watercolor, and crayon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version2
	PKContentVersion2 PKContentVersion = 0
	// PKContentVersion3 - The PencilKit version that supports barrel-roll angle data in inks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version3
	PKContentVersion3 PKContentVersion = 0
	// PKContentVersion4 - New Reed Pen
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKContentVersion/version4
	PKContentVersion4 PKContentVersion = 0
)

/* debug [enums.gen.go]: Processing enum PKToolPickerCustomItemControlOptions (3 cases) */
// PKToolPickerCustomItemControlOptions - Options for which controls to present.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions
type PKToolPickerCustomItemControlOptions uint

const (
	// PKToolPickerCustomItemControlOpacity - Present an opacity control if color adjustment is supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions/opacity
	PKToolPickerCustomItemControlOpacity PKToolPickerCustomItemControlOptions = 0
	// PKToolPickerCustomItemControlWidth - Present a width control if width adjustment is supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItem/ControlOptions/width
	PKToolPickerCustomItemControlWidth PKToolPickerCustomItemControlOptions = 0
	// PKToolPickerCustomItemControlNone - Present neither a width nor opacity control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerCustomItemControlOptions/PKToolPickerCustomItemControlNone
	PKToolPickerCustomItemControlNone PKToolPickerCustomItemControlOptions = 0
)

/* debug [enums.gen.go]: Processing enum PKToolPickerVisibility (4 cases) */
// PKToolPickerVisibility - The visibility state of a tool picker.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility
type PKToolPickerVisibility uint

const (
	// PKToolPickerVisibilityHidden - Tool picker is active but offscreen, and can appear temporarily in response to user actions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/hidden
	PKToolPickerVisibilityHidden PKToolPickerVisibility = 0
	// PKToolPickerVisibilityInactive - No active tool picker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/inactive
	PKToolPickerVisibilityInactive PKToolPickerVisibility = 0
	// PKToolPickerVisibilityInherited - Inherit the tool picker visibility from the next responder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/PKToolPickerVisibilityInherited
	PKToolPickerVisibilityInherited PKToolPickerVisibility = 0
	// PKToolPickerVisibilityVisible - Tool picker is active and onscreen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKToolPickerVisibility/visible
	PKToolPickerVisibilityVisible PKToolPickerVisibility = 0
)

/* debug [enums.gen.go]: Processing enum PKCanvasViewDrawingPolicy (3 cases) */
// PKCanvasViewDrawingPolicy - Constants that you use to specify the type of drawing gestures your app permits while the user draws on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy
type PKCanvasViewDrawingPolicy uint

const (
	// PKCanvasViewDrawingPolicyAnyInput - Allows drawing on the canvas from any input source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/anyInput
	PKCanvasViewDrawingPolicyAnyInput PKCanvasViewDrawingPolicy = 0
	// PKCanvasViewDrawingPolicyDefault - The default input type to use for drawing on a canvas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/default
	PKCanvasViewDrawingPolicyDefault PKCanvasViewDrawingPolicy = 0
	// PKCanvasViewDrawingPolicyPencilOnly - Pencil touches are the only input that draw on the canvas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasViewDrawingPolicy/pencilOnly
	PKCanvasViewDrawingPolicyPencilOnly PKCanvasViewDrawingPolicy = 0
)

/* debug [enums.gen.go]: Processing enum PKEraserType (3 cases) */
// PKEraserType - Constants that indicate the behavior of the eraser.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType
type PKEraserType int

const (
	// PKEraserTypeBitmap - An eraser that removes only those portions of the drawing it touches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeBitmap
	PKEraserTypeBitmap PKEraserType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeFixedWidthBitmap
	PKEraserTypeFixedWidthBitmap PKEraserType = 0
	// PKEraserTypeVector - An eraser that removes an entire drawn line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserType/PKEraserTypeVector
	PKEraserTypeVector PKEraserType = 0
)


