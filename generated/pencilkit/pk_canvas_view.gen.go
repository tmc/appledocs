// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CanvasView] class.
var (
	CanvasViewClass     _CanvasViewClass
	CanvasViewClassOnce sync.Once
)

func getCanvasViewClass() _CanvasViewClass {
	CanvasViewClassOnce.Do(func() {
		CanvasViewClass = _CanvasViewClass{objc.GetClass("PKCanvasView")}
	})
	return CanvasViewClass
}

type _CanvasViewClass struct {
	class objc.Class
}

// An interface definition for the [CanvasView] class.
type ICanvasView interface {
	objectivec.IObject
}

// A view that captures Apple Pencil input and displays the rendered results in an iOS app.
//
// A object captures content drawn using Apple Pencil or the user’s finger and displays it in your app. The canvas view handles all of the touch events and data coming from Apple Pencil, and renders that information using the tool you specify. The canvas stores the captured input in a object. is a scroll view, so you can make the drawable area bigger than the canvas view’s frame rectangle. To do that, set the inherited property to the size you want. The canvas view automatically scales its underlying content to match the size you specify. Users scroll around the canvas using a two-finger pan gesture. (If the property is , users scroll with only one finger.) A canvas view conforms to the protocol, so you can add it as an observer of the window’s tool picker. The tool picker displays a floating palette of tools that the user can choose from. As the user interacts with items in the palette, such as changing ink colors, or line widths, the canvas automatically updates its drawing environment accordingly.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView
type CanvasView struct {
	objectivec.Object
}

// CanvasViewFrom constructs a [CanvasView] from an unsafe.Pointer.
//
// A view that captures Apple Pencil input and displays the rendered results in an iOS app.
func CanvasViewFrom(ptr unsafe.Pointer) CanvasView {
	return CanvasView{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CanvasViewClass) Alloc() CanvasView {
	rv := objc.Send[CanvasView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CanvasViewClass) New() CanvasView {
	rv := objc.Send[CanvasView](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CanvasView) Init() CanvasView {
	rv := objc.Send[CanvasView](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CanvasView) Autorelease() CanvasView {
	rv := objc.Send[CanvasView](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCanvasView creates a new CanvasView instance.
func NewCanvasView() CanvasView {
	return getCanvasViewClass().New()
}


// A Boolean value that indicates whether a ruler view is visible on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isruleractive
func (c_ CanvasView) IsRulerActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRulerActive"))
	return rv
}


// SetIsRulerActive sets the value of the isRulerActive property.
// A Boolean value that indicates whether a ruler view is visible on the canvas.

//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isruleractive
func (c_ CanvasView) SetIsRulerActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRulerActive:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isdrawingenabled
func (c_ CanvasView) IsDrawingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDrawingEnabled"))
	return rv
}


// SetIsDrawingEnabled sets the value of the isDrawingEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isdrawingenabled
func (c_ CanvasView) SetIsDrawingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDrawingEnabled:"), value)
}

// A Boolean value that indicates whether the canvas accepts input from the user’s finger in addition to Apple Pencil.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/allowsFingerDrawing
func (c_ CanvasView) AllowsFingerDrawing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsFingerDrawing"))
	return rv
}


// SetAllowsFingerDrawing sets the value of the allowsFingerDrawing property.
// A Boolean value that indicates whether the canvas accepts input from the user’s finger in addition to Apple Pencil.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/allowsFingerDrawing
func (c_ CanvasView) SetAllowsFingerDrawing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsFingerDrawing:"), value)
}

// The object you use to respond to changes in the drawn content or with the selected tool.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/delegate
func (c_ CanvasView) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object you use to respond to changes in the drawn content or with the selected tool.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/delegate
func (c_ CanvasView) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}

// The data object that the canvas uses to store drawn content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawing
func (c_ CanvasView) Drawing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("drawing"))
	return rv
}


// SetDrawing sets the value of the drawing property.
// The data object that the canvas uses to store drawn content.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawing
func (c_ CanvasView) SetDrawing(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawing:"), value)
}

// The gesture recognizer that the canvas uses to track touch events.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawingGestureRecognizer
func (c_ CanvasView) DrawingGestureRecognizer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("drawingGestureRecognizer"))
	return rv
}

// The policy that controls the types of touches allowed when drawing on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawingPolicy
func (c_ CanvasView) DrawingPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("drawingPolicy"))
	return rv
}


// SetDrawingPolicy sets the value of the drawingPolicy property.
// The policy that controls the types of touches allowed when drawing on the canvas.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/drawingPolicy
func (c_ CanvasView) SetDrawingPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawingPolicy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isDrawingEnabled
func (c_ CanvasView) DrawingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("drawingEnabled"))
	return rv
}


// SetDrawingEnabled sets the value of the drawingEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isDrawingEnabled
func (c_ CanvasView) SetDrawingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDrawingEnabled:"), value)
}

// A Boolean value that indicates whether a ruler view is visible on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isRulerActive
func (c_ CanvasView) RulerActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("rulerActive"))
	return rv
}


// SetRulerActive sets the value of the rulerActive property.
// A Boolean value that indicates whether a ruler view is visible on the canvas.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/isRulerActive
func (c_ CanvasView) SetRulerActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRulerActive:"), value)
}

// The maximum version of PencilKit to support.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/maximumSupportedContentVersion
func (c_ CanvasView) MaximumSupportedContentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maximumSupportedContentVersion"))
	return rv
}


// SetMaximumSupportedContentVersion sets the value of the maximumSupportedContentVersion property.
// The maximum version of PencilKit to support.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/maximumSupportedContentVersion
func (c_ CanvasView) SetMaximumSupportedContentVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaximumSupportedContentVersion:"), value)
}

// The currently selected tool used for drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/tool-6str6
func (c_ CanvasView) Tool() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tool"))
	return rv
}


// SetTool sets the value of the tool property.
// The currently selected tool used for drawing.

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView/tool-6str6
func (c_ CanvasView) SetTool(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTool:"), value)
}



