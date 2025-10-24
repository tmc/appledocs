//go:build darwin && ios

// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for CanvasView


// iOS-only properties

// A Boolean value that indicates whether the canvas accepts input from the user’s finger in addition to Apple Pencil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/allowsFingerDrawing
func (c_ CanvasView) AllowsFingerDrawing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsFingerDrawing"))
	return rv
}
func (c_ CanvasView) SetAllowsFingerDrawing(value bool) {
	c_.ID.Send(objc.RegisterName("setAllowsFingerDrawing:"), value)
}

// The object you use to respond to changes in the drawn content or with the selected tool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/delegate
func (c_ CanvasView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}
func (c_ CanvasView) SetDelegate(value unsafe.Pointer) {
	c_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The data object that the canvas uses to store drawn content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawing
func (c_ CanvasView) Drawing() IPKDrawing {
	rv := objc.Send[Drawing](c_.ID, objc.Sel("drawing"))
	return rv
}
func (c_ CanvasView) SetDrawing(value IPKDrawing) {
	c_.ID.Send(objc.RegisterName("setDrawing:"), value)
}

// The gesture recognizer that the canvas uses to track touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawingGestureRecognizer
func (c_ CanvasView) DrawingGestureRecognizer() appkit.GestureRecognizer {
	rv := objc.Send[appkit.GestureRecognizer](c_.ID, objc.Sel("drawingGestureRecognizer"))
	return rv
}

// The policy that controls the types of touches allowed when drawing on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawingPolicy
func (c_ CanvasView) DrawingPolicy() CanvasViewDrawingPolicy {
	rv := objc.Send[CanvasViewDrawingPolicy](c_.ID, objc.Sel("drawingPolicy"))
	return rv
}
func (c_ CanvasView) SetDrawingPolicy(value CanvasViewDrawingPolicy) {
	c_.ID.Send(objc.RegisterName("setDrawingPolicy:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isDrawingEnabled
func (c_ CanvasView) DrawingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("drawingEnabled"))
	return rv
}
func (c_ CanvasView) SetDrawingEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setDrawingEnabled:"), value)
}

// A Boolean value that indicates whether a ruler view is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isRulerActive
func (c_ CanvasView) RulerActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rulerActive"))
	return rv
}
func (c_ CanvasView) SetRulerActive(value bool) {
	c_.ID.Send(objc.RegisterName("setRulerActive:"), value)
}

// The maximum version of PencilKit to support.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/maximumSupportedContentVersion
func (c_ CanvasView) MaximumSupportedContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](c_.ID, objc.Sel("maximumSupportedContentVersion"))
	return rv
}
func (c_ CanvasView) SetMaximumSupportedContentVersion(value ContentVersion) {
	c_.ID.Send(objc.RegisterName("setMaximumSupportedContentVersion:"), value)
}

// The currently selected tool used for drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/tool-6str6
func (c_ CanvasView) Tool() IPKTool {
	rv := objc.Send[Tool](c_.ID, objc.Sel("tool"))
	return rv
}
func (c_ CanvasView) SetTool(value IPKTool) {
	c_.ID.Send(objc.RegisterName("setTool:"), value)
}





