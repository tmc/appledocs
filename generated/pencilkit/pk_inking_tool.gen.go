// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InkingTool] class.
var (
	InkingToolClass     _InkingToolClass
	InkingToolClassOnce sync.Once
)

func getInkingToolClass() _InkingToolClass {
	InkingToolClassOnce.Do(func() {
		InkingToolClass = _InkingToolClass{objc.GetClass("PKInkingTool")}
	})
	return InkingToolClass
}

type _InkingToolClass struct {
	class objc.Class
}

// An interface definition for the [InkingTool] class.
type IInkingTool interface {
	ITool
}

// An object that defines the drawing characteristics (width, color, pen style) to use when drawing lines on a canvas view.
//
// A object supports the creation of new content on a . With an inking tool, the canvas turns touch input from the user into a continuously rendered stroke. The value in the property determines the base width of that stroke; however, that base value also depends on input from Apple Pencil, including force, azimuth, and angle data. Create an inking tool programmatically, or display a object and from which a user can select a tool. Assign the resulting object to the property of your object. The canvas uses any subsequent touch sequences to draw new content on the canvas. Assigning a new inking tool doesn’t change the characteristics for any previously drawn strokes.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference
type InkingTool struct {
	Tool
}

// InkingToolFrom constructs a [InkingTool] from an unsafe.Pointer.
//
// An object that defines the drawing characteristics (width, color, pen style) to use when drawing lines on a canvas view.
func InkingToolFrom(ptr unsafe.Pointer) InkingTool {
	return InkingTool{
		Tool: ToolFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _InkingToolClass) Alloc() InkingTool {
	rv := objc.Send[InkingTool](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _InkingToolClass) New() InkingTool {
	rv := objc.Send[InkingTool](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InkingTool) Init() InkingTool {
	rv := objc.Send[InkingTool](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InkingTool) Autorelease() InkingTool {
	rv := objc.Send[InkingTool](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInkingTool creates a new InkingTool instance.
func NewInkingTool() InkingTool {
	return getInkingToolClass().New()
}


// Create an inking tool with the specified ink and width.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/init(ink:width:)
func NewInkingToolWithInkWidth(ink unsafe.Pointer, width float64) InkingTool {
	instance := getInkingToolClass().Alloc()
	rv := objc.Send[InkingTool](instance.ID, objc.Sel("initWithInk:width:"), ink, width)
	rv.Autorelease()
	return rv
}

// Creates an ink tool object with the default line width and the specified color.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/init(inkType:color:)
func NewInkingToolWithInkTypeColor(type_ unsafe.Pointer, color unsafe.Pointer) InkingTool {
	instance := getInkingToolClass().Alloc()
	rv := objc.Send[InkingTool](instance.ID, objc.Sel("initWithInkType:color:"), type_, color)
	rv.Autorelease()
	return rv
}

// Creates an ink tool object with the specified color and line width values.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/init(inkType:color:width:)
func NewInkingToolWithInkTypeColorWidth(type_ unsafe.Pointer, color unsafe.Pointer, width float64) InkingTool {
	instance := getInkingToolClass().Alloc()
	rv := objc.Send[InkingTool](instance.ID, objc.Sel("initWithInkType:color:width:"), type_, color, width)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/init(inkType:color:width:azimuth:)
func NewInkingToolWithInkTypeColorWidthAzimuth(type_ unsafe.Pointer, color unsafe.Pointer, width float64, angle float64) InkingTool {
	instance := getInkingToolClass().Alloc()
	rv := objc.Send[InkingTool](instance.ID, objc.Sel("initWithInkType:color:width:azimuth:"), type_, color, width, angle)
	rv.Autorelease()
	return rv
}


// Converts a color from one user interface style to another.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/convert(_:from:to:)
func (ic _InkingToolClass) ConvertColorFromUserInterfaceStyleTo(color unsafe.Pointer, fromUserInterfaceStyle unsafe.Pointer, toUserInterfaceStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("convertColor:fromUserInterfaceStyle:to:"), color, fromUserInterfaceStyle, toUserInterfaceStyle)
	return rv
}

// Returns the default line width for the specified tool type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/defaultWidth(forInkType:)
func (ic _InkingToolClass) DefaultWidthForInkType(inkType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ic.class), objc.Sel("defaultWidthForInkType:"), inkType)
	return rv
}

// Converts a color from light to dark appearance or vice versa.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/invertColor(_:)
func (ic _InkingToolClass) InvertColor(color CGColorRef) CGColorRef {
	rv := objc.Send[CGColorRef](objc.ID(ic.class), objc.Sel("invertColor:"), color)
	return rv
}

// Returns the maximum allowed line width for the specified tool type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/maximumWidth(forInkType:)
func (ic _InkingToolClass) MaximumWidthForInkType(inkType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ic.class), objc.Sel("maximumWidthForInkType:"), inkType)
	return rv
}

// Returns the minimum allowed line width for the specified tool type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/minimumWidth(forInkType:)
func (ic _InkingToolClass) MinimumWidthForInkType(inkType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ic.class), objc.Sel("minimumWidthForInkType:"), inkType)
	return rv
}

// The base angle of the ink.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/azimuth
func (i_ InkingTool) Azimuth() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("azimuth"))
	return rv
}

// The color of the ink.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/color
func (i_ InkingTool) Color() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("color"))
	return rv
}

// The ink that this tool creates strokes with.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/ink
func (i_ InkingTool) Ink() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("ink"))
	return rv
}

// The tool type that determines the shape of the rendered content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/inkType
func (i_ InkingTool) InkType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("inkType"))
	return rv
}

// The version of PencilKit necessary to use the inking tool.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/requiredContentVersion
func (i_ InkingTool) RequiredContentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("requiredContentVersion"))
	return rv
}

// The base line width for new content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKInkingToolReference/width
func (i_ InkingTool) Width() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("width"))
	return rv
}


