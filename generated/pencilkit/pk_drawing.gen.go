// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PKDrawing */


/* debug [class_header]: Header for PKDrawing */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Drawing */
// An interface definition for the [Drawing] class.
type IDrawing interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Drawing */
	// properties:
	Bounds() corefoundation.CGRect
	RequiredContentVersion() ContentVersion
	Strokes() []Stroke
	PKAppleDrawingTypeIdentifier() foundation.String
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Drawing */
	// methods:
	DrawingByAppendingDrawing(drawing IPKDrawing) IDrawing
	DrawingByAppendingStrokes(strokes []Stroke) IDrawing
	DrawingByApplyingTransform(transform corefoundation.CGAffineTransform) IDrawing
	DataRepresentation() foundation.Data
	ImageFromRectScale(rect corefoundation.CGRect, scale float64) appkit.Image
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Drawing */
// Alloc allocates a new instance without initialization.
func (dc _DrawingClass) Alloc() Drawing {
	rv := objc.Send[Drawing](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Drawing */
// A data structure that contains the drawing information captured by a canvas view.
//
// A object stores the user-drawn content from a object. You use drawing objects to store the data associated with your user’s drawings. You can save this data with the rest of your app’s content, and you can use that saved data to create a new drawing object later. You can also generate an image based on the drawn content.


// A data structure that contains the drawing information captured by a canvas view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Drawing */

// Creates a drawing object and populates it with previously drawn content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/init(data:)
func NewDrawingWithDataError(data objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) Drawing {
	instance := getDrawingClass().Alloc()
	rv := objc.Send[Drawing](instance.ID, objc.Sel("initWithData:error:"), data, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDrawingWithDataError */


// Creates a drawing object with the strokes you supply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/init(strokes:)
func NewDrawingWithStrokes(strokes []Stroke) Drawing {
	instance := getDrawingClass().Alloc()
	rv := objc.Send[Drawing](instance.ID, objc.Sel("initWithStrokes:"), strokes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDrawingWithStrokes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Drawing */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Drawing */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Drawing */

// Returns a new drawing created by appending the current drawing with another drawing you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/appending(_:)
func (d_ Drawing) DrawingByAppendingDrawing(drawing IPKDrawing) IDrawing {
	rv := objc.Send[Drawing](d_.ID, objc.Sel("drawingByAppendingDrawing:"), drawing)
	return rv
}/* debug [instance_methods/method]: DrawingByAppendingDrawing */


// Returns a copy of the current drawing with the strokes you provide appended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/appendingStrokes(_:)
func (d_ Drawing) DrawingByAppendingStrokes(strokes []Stroke) IDrawing {
	rv := objc.Send[Drawing](d_.ID, objc.Sel("drawingByAppendingStrokes:"), strokes)
	return rv
}/* debug [instance_methods/method]: DrawingByAppendingStrokes */


// Returns a new drawing object by applying the specified transform to a copy of the current object’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/applying(_:)
func (d_ Drawing) DrawingByApplyingTransform(transform corefoundation.CGAffineTransform) IDrawing {
	rv := objc.Send[Drawing](d_.ID, objc.Sel("drawingByApplyingTransform:"), transform)
	return rv
}/* debug [instance_methods/method]: DrawingByApplyingTransform */


// Returns a representation of the rendered content as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/dataRepresentation()
func (d_ Drawing) DataRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](d_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_methods/method]: DataRepresentation */


// Returns an image object that contains the specified portion of the drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/image(from:scale:)
func (d_ Drawing) ImageFromRectScale(rect corefoundation.CGRect, scale float64) appkit.Image {
	rv := objc.Send[appkit.Image](d_.ID, objc.Sel("imageFromRect:scale:"), rect, scale)
	return rv
}/* debug [instance_methods/method]: ImageFromRectScale */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Drawing */

// The smallest rectangle used to represent the content’s bounds, taking into account line widths of that content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/bounds
func (d_ Drawing) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](d_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The version of PencilKit necessary to use the drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/requiredContentVersion
func (d_ Drawing) RequiredContentVersion() ContentVersion {
	rv := objc.Send[ContentVersion](d_.ID, objc.Sel("requiredContentVersion"))
	return rv
}/* debug [instance_properties/getter]: requiredContentVersion */


// An array of strokes used to represent the drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKDrawingReference/strokes
func (d_ Drawing) Strokes() []Stroke {
	rv := objc.Send[[]Stroke](d_.ID, objc.Sel("strokes"))
	return rv
}/* debug [instance_properties/getter]: strokes */


// The uniform type identifier for data associated with a drawing object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkappledrawingtypeidentifier
func (d_ Drawing) PKAppleDrawingTypeIdentifier() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("PKAppleDrawingTypeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: PKAppleDrawingTypeIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKDrawing */


