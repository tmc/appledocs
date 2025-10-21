// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EraserTool] class.
var (
	EraserToolClass     _EraserToolClass
	EraserToolClassOnce sync.Once
)

func getEraserToolClass() _EraserToolClass {
	EraserToolClassOnce.Do(func() {
		EraserToolClass = _EraserToolClass{objc.GetClass("PKEraserTool")}
	})
	return EraserToolClass
}

type _EraserToolClass struct {
	class objc.Class
}

// An interface definition for the [EraserTool] class.
type IEraserTool interface {
	ITool
}

// A tool for erasing previously drawn content in a canvas view.
//
// A object supports the deletion of content from a object. The eraser tool’s type determines whether the canvas removes entire items or just the portion of an item that it touches. Create an eraser tool programmatically or display a object and let the user select the eraser. Assign the resulting object to the property of your object. The canvas uses any subsequent touch sequences to erase content on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference
type EraserTool struct {
	Tool
}

// EraserToolFrom constructs a [EraserTool] from an unsafe.Pointer.
//
// A tool for erasing previously drawn content in a canvas view.
func EraserToolFrom(ptr unsafe.Pointer) EraserTool {
	return EraserTool{
		Tool: ToolFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EraserToolClass) Alloc() EraserTool {
	rv := objc.Send[EraserTool](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EraserToolClass) New() EraserTool {
	rv := objc.Send[EraserTool](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EraserTool) Init() EraserTool {
	rv := objc.Send[EraserTool](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EraserTool) Autorelease() EraserTool {
	rv := objc.Send[EraserTool](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEraserTool creates a new EraserTool instance.
func NewEraserTool() EraserTool {
	return getEraserToolClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/init(eraserType:width:)
func NewEraserToolWithEraserTypeWidth(eraserType unsafe.Pointer, width float64) EraserTool {
	instance := getEraserToolClass().Alloc()
	rv := objc.Send[EraserTool](instance.ID, objc.Sel("initWithEraserType:width:"), eraserType, width)
	rv.Autorelease()
	return rv
}

// Creates an eraser tool object that removes objects wholly or partially from a canvas view.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/init(eraserType:)
func NewEraserToolWithEraserType(eraserType unsafe.Pointer) EraserTool {
	instance := getEraserToolClass().Alloc()
	rv := objc.Send[EraserTool](instance.ID, objc.Sel("initWithEraserType:"), eraserType)
	rv.Autorelease()
	return rv
}


// The default width for the specified eraser type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/defaultWidth(for:)
func (ec _EraserToolClass) DefaultWidthForEraserType(eraserType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("defaultWidthForEraserType:"), eraserType)
	return rv
}

// The maximum width for the specified eraser type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/maximumWidth(for:)
func (ec _EraserToolClass) MaximumWidthForEraserType(eraserType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("maximumWidthForEraserType:"), eraserType)
	return rv
}

// The minimum width for the specified eraser type.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/minimumWidth(for:)
func (ec _EraserToolClass) MinimumWidthForEraserType(eraserType unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("minimumWidthForEraserType:"), eraserType)
	return rv
}

// The behavior adopted by the eraser when deleting content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/eraserType
func (e_ EraserTool) EraserType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("eraserType"))
	return rv
}

// The width of the eraser.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/width
func (e_ EraserTool) Width() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("width"))
	return rv
}


