// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Drawing] class.
var (
	DrawingClass     _DrawingClass
	DrawingClassOnce sync.Once
)

func getDrawingClass() _DrawingClass {
	DrawingClassOnce.Do(func() {
		DrawingClass = _DrawingClass{objc.GetClass("PKDrawing")}
	})
	return DrawingClass
}

type _DrawingClass struct {
	class objc.Class
}

// An interface definition for the [Drawing] class.
type IDrawing interface {
	objectivec.IObject
	DrawingByAppendingDrawing(drawing unsafe.Pointer) unsafe.Pointer
	DrawingByAppendingStrokes(strokes unsafe.Pointer) unsafe.Pointer
	DrawingByApplyingTransform(transform coregraphics.CGAffineTransform) unsafe.Pointer
	DataRepresentation() unsafe.Pointer
	ImageFromRectScale(rect coregraphics.CGRect, scale float64) unsafe.Pointer
}

// A data structure that contains the drawing information captured by a canvas view.
//
// A object stores the user-drawn content from a object. You use drawing objects to store the data associated with your user’s drawings. You can save this data with the rest of your app’s content, and you can use that saved data to create a new drawing object later. You can also generate an image based on the drawn content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference
type Drawing struct {
	objectivec.Object
}

// DrawingFrom constructs a [Drawing] from an unsafe.Pointer.
//
// A data structure that contains the drawing information captured by a canvas view.
func DrawingFrom(ptr unsafe.Pointer) Drawing {
	return Drawing{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DrawingClass) Alloc() Drawing {
	rv := objc.Send[Drawing](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DrawingClass) New() Drawing {
	rv := objc.Send[Drawing](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Drawing) Init() Drawing {
	rv := objc.Send[Drawing](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Drawing) Autorelease() Drawing {
	rv := objc.Send[Drawing](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDrawing creates a new Drawing instance.
func NewDrawing() Drawing {
	return getDrawingClass().New()
}




// Creates a drawing object and populates it with previously drawn content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/init(data:)
func NewDrawingWithDataError(data unsafe.Pointer, error_ unsafe.Pointer) Drawing {
	instance := getDrawingClass().Alloc()
	rv := objc.Send[Drawing](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}



// Creates a drawing object with the strokes you supply.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/init(strokes:)
func NewDrawingWithStrokes(strokes unsafe.Pointer) Drawing {
	instance := getDrawingClass().Alloc()
	rv := objc.Send[Drawing](instance.ID, objc.Sel("initWithStrokes:"), strokes)
	rv.Autorelease()
	return rv
}


// Returns a new drawing created by appending the current drawing with another drawing you provide.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/appending(_:)
func (d_ Drawing) DrawingByAppendingDrawing(drawing unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("drawingByAppendingDrawing:"), drawing)
	return rv
}

// Returns a copy of the current drawing with the strokes you provide appended.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/appendingStrokes(_:)
func (d_ Drawing) DrawingByAppendingStrokes(strokes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("drawingByAppendingStrokes:"), strokes)
	return rv
}

// Returns a new drawing object by applying the specified transform to a copy of the current object’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/applying(_:)
func (d_ Drawing) DrawingByApplyingTransform(transform coregraphics.CGAffineTransform) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("drawingByApplyingTransform:"), transform)
	return rv
}

// Returns a representation of the rendered content as data.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/dataRepresentation()
func (d_ Drawing) DataRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dataRepresentation"))
	return rv
}

// Returns an image object that contains the specified portion of the drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/image(from:scale:)
func (d_ Drawing) ImageFromRectScale(rect coregraphics.CGRect, scale float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("imageFromRect:scale:"), rect, scale)
	return rv
}

// The smallest rectangle used to represent the content’s bounds, taking into account line widths of that content.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/bounds
func (d_ Drawing) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](d_.ID, objc.Sel("bounds"))
	return rv
}

// The version of PencilKit necessary to use the drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/requiredContentVersion
func (d_ Drawing) RequiredContentVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("requiredContentVersion"))
	return rv
}

// An array of strokes used to represent the drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/strokes
func (d_ Drawing) Strokes() []Stroke {
	rv := objc.Send[[]Stroke](d_.ID, objc.Sel("strokes"))
	return rv
}


