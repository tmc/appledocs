// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextViewportLayoutController */


/* debug [class_header]: Header for NSTextViewportLayoutController */
// The class instance for the [TextViewportLayoutController] class.
var (
	TextViewportLayoutControllerClass     _TextViewportLayoutControllerClass
	TextViewportLayoutControllerClassOnce sync.Once
)

func getTextViewportLayoutControllerClass() _TextViewportLayoutControllerClass {
	TextViewportLayoutControllerClassOnce.Do(func() {
		TextViewportLayoutControllerClass = _TextViewportLayoutControllerClass{objc.GetClass("NSTextViewportLayoutController")}
	})
	return TextViewportLayoutControllerClass
}

type _TextViewportLayoutControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextViewportLayoutController */
// An interface definition for the [TextViewportLayoutController] class.
type ITextViewportLayoutController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextViewportLayoutController */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	TextLayoutManager() ITextLayoutManager
	ViewportBounds() corefoundation.CGRect
	ViewportRange() ITextRange
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextViewportLayoutController */
	// methods:
	AdjustViewportByVerticalOffset(verticalOffset float64)
	LayoutViewport()
	RelocateViewportToTextLocation(textLocation unsafe.Pointer) float64
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextViewportLayoutController */
// Alloc allocates a new instance without initialization.
func (tc _TextViewportLayoutControllerClass) Alloc() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextViewportLayoutControllerClass) New() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextViewportLayoutController) Init() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextViewportLayoutController) Autorelease() TextViewportLayoutController {
	rv := objc.Send[TextViewportLayoutController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextViewportLayoutController creates a new TextViewportLayoutController instance.
func NewTextViewportLayoutController() TextViewportLayoutController {
	return getTextViewportLayoutControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextViewportLayoutController */
// Manages the layout process inside the viewport interacting with its delegate.
//
// A viewport is a rectangular area within a flipped coordinate system expanding along the y-axis. With text contents, lines advance expanding the view in the current writing direction. The viewport defines the active area where the framework lays out text fragments. In most cases, the area corresponds to the user visible area with an additional over-scroll region.


// Manages the layout process inside the viewport interacting with its delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController
type TextViewportLayoutController struct {
	objectivec.Object
}

// TextViewportLayoutControllerFrom constructs a [TextViewportLayoutController] from an unsafe.Pointer.
//
// Manages the layout process inside the viewport interacting with its delegate.
func TextViewportLayoutControllerFrom(ptr unsafe.Pointer) TextViewportLayoutController {
	return TextViewportLayoutController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextViewportLayoutController */

// Creates a new instance with the text layout manager you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/init(textLayoutManager:)
func NewTextViewportLayoutControllerWithTextLayoutManager(textLayoutManager ITextLayoutManager) TextViewportLayoutController {
	instance := getTextViewportLayoutControllerClass().Alloc()
	rv := objc.Send[TextViewportLayoutController](instance.ID, objc.Sel("initWithTextLayoutManager:"), textLayoutManager)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextViewportLayoutControllerWithTextLayoutManager */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextViewportLayoutController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextViewportLayoutController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextViewportLayoutController */

// Adjusts the viewport rect by the specified offset if needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/adjustViewport(byVerticalOffset:)
func (t_ TextViewportLayoutController) AdjustViewportByVerticalOffset(verticalOffset float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("adjustViewportByVerticalOffset:"), verticalOffset)
}/* debug [instance_methods/method]: AdjustViewportByVerticalOffset */


// Performs layout in the viewport.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/layoutViewport()
func (t_ TextViewportLayoutController) LayoutViewport() {
	objc.Send[objc.ID](t_.ID, objc.Sel("layoutViewport"))
}/* debug [instance_methods/method]: LayoutViewport */


// Relocates the viewport to the location you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/relocateViewport(to:)
func (t_ TextViewportLayoutController) RelocateViewportToTextLocation(textLocation unsafe.Pointer) float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("relocateViewportToTextLocation:"), textLocation)
	return rv
}/* debug [instance_methods/method]: RelocateViewportToTextLocation */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextViewportLayoutController */

// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/delegate
func (t_ TextViewportLayoutController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the text layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/delegate
func (t_ TextViewportLayoutController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Returns the text layout manager for this viewport layout controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/textLayoutManager
func (t_ TextViewportLayoutController) TextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("textLayoutManager"))
	return rv
}/* debug [instance_properties/getter]: textLayoutManager */


// Returns the visible bounds of the view, plus the overdraw area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/viewportBounds
func (t_ TextViewportLayoutController) ViewportBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t_.ID, objc.Sel("viewportBounds"))
	return rv
}/* debug [instance_properties/getter]: viewportBounds */


// Returns the text range of the current viewport layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/viewportRange
func (t_ TextViewportLayoutController) ViewportRange() ITextRange {
	rv := objc.Send[TextRange](t_.ID, objc.Sel("viewportRange"))
	return rv
}/* debug [instance_properties/getter]: viewportRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextViewportLayoutController */


