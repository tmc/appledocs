// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PKEraserTool */


/* debug [class_header]: Header for PKEraserTool */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EraserTool */
// An interface definition for the [EraserTool] class.
type IEraserTool interface {
	ITool
	
/* debug [class_interface_properties]: Properties for EraserTool */
	// properties:
	EraserType() EraserType
	Width() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EraserTool */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EraserTool */
// Alloc allocates a new instance without initialization.
func (ec _EraserToolClass) Alloc() EraserTool {
	rv := objc.Send[EraserTool](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EraserTool */
// A tool for erasing previously drawn content in a canvas view.
//
// A object supports the deletion of content from a object. The eraser tool’s type determines whether the canvas removes entire items or just the portion of an item that it touches. Create an eraser tool programmatically or display a object and let the user select the eraser. Assign the resulting object to the property of your object. The canvas uses any subsequent touch sequences to erase content on the canvas.


// A tool for erasing previously drawn content in a canvas view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EraserTool */

// Creates an eraser tool object that removes objects wholly or partially from a canvas view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/init(eraserType:)
func NewEraserToolWithEraserType(eraserType EraserType) EraserTool {
	instance := getEraserToolClass().Alloc()
	rv := objc.Send[EraserTool](instance.ID, objc.Sel("initWithEraserType:"), eraserType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEraserToolWithEraserType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/init(eraserType:width:)
func NewEraserToolWithEraserTypeWidth(eraserType EraserType, width float64) EraserTool {
	instance := getEraserToolClass().Alloc()
	rv := objc.Send[EraserTool](instance.ID, objc.Sel("initWithEraserType:width:"), eraserType, width)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewEraserToolWithEraserTypeWidth */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EraserTool */

// The default width for the specified eraser type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/defaultWidth(for:)
func (ec _EraserToolClass) DefaultWidthForEraserType(eraserType EraserType) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("defaultWidthForEraserType:"), eraserType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultWidthForEraserType) */


// The maximum width for the specified eraser type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/maximumWidth(for:)
func (ec _EraserToolClass) MaximumWidthForEraserType(eraserType EraserType) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("maximumWidthForEraserType:"), eraserType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaximumWidthForEraserType) */


// The minimum width for the specified eraser type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/minimumWidth(for:)
func (ec _EraserToolClass) MinimumWidthForEraserType(eraserType EraserType) float64 {
	rv := objc.Send[float64](objc.ID(ec.class), objc.Sel("minimumWidthForEraserType:"), eraserType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MinimumWidthForEraserType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EraserTool */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EraserTool */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EraserTool */

// The behavior adopted by the eraser when deleting content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/eraserType
func (e_ EraserTool) EraserType() EraserType {
	rv := objc.Send[EraserType](e_.ID, objc.Sel("eraserType"))
	return rv
}/* debug [instance_properties/getter]: eraserType */


// The width of the eraser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKEraserToolReference/width
func (e_ EraserTool) Width() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("width"))
	return rv
}/* debug [instance_properties/getter]: width */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKEraserTool */


