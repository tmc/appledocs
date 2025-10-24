// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

/* debug [class.gen.go]: Generating class PKCanvasView */


/* debug [class_header]: Header for PKCanvasView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CanvasView */
// An interface definition for the [CanvasView] class.
type ICanvasView interface {
	IScrollView
	
/* debug [class_interface_properties]: Properties for CanvasView */
	// properties:
	IsDrawingEnabled() bool
	SetIsDrawingEnabled(value bool)
	IsRulerActive() bool
	SetIsRulerActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CanvasView */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CanvasView */
// Alloc allocates a new instance without initialization.
func (cc _CanvasViewClass) Alloc() CanvasView {
	rv := objc.Send[CanvasView](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CanvasView */
// A view that captures Apple Pencil input and displays the rendered results in an iOS app.
//
// A object captures content drawn using Apple Pencil or the user’s finger and displays it in your app. The canvas view handles all of the touch events and data coming from Apple Pencil, and renders that information using the tool you specify. The canvas stores the captured input in a object. is a scroll view, so you can make the drawable area bigger than the canvas view’s frame rectangle. To do that, set the inherited property to the size you want. The canvas view automatically scales its underlying content to match the size you specify. Users scroll around the canvas using a two-finger pan gesture. (If the property is , users scroll with only one finger.) A canvas view conforms to the protocol, so you can add it as an observer of the window’s tool picker. The tool picker displays a floating palette of tools that the user can choose from. As the user interacts with items in the palette, such as changing ink colors, or line widths, the canvas automatically updates its drawing environment accordingly.


// A view that captures Apple Pencil input and displays the rendered results in an iOS app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKCanvasView
type CanvasView struct {
	ScrollView
}

// CanvasViewFrom constructs a [CanvasView] from an unsafe.Pointer.
//
// A view that captures Apple Pencil input and displays the rendered results in an iOS app.
func CanvasViewFrom(ptr unsafe.Pointer) CanvasView {
	return CanvasView{
		ScrollView: ScrollViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CanvasView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CanvasView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CanvasView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CanvasView */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CanvasView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isdrawingenabled
func (c_ CanvasView) IsDrawingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isDrawingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDrawingEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isdrawingenabled
func (c_ CanvasView) SetIsDrawingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsDrawingEnabled:"), value)
}/* debug [instance_properties/setter]: isDrawingEnabled */


// A Boolean value that indicates whether a ruler view is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isruleractive
func (c_ CanvasView) IsRulerActive() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRulerActive"))
	return rv
}/* debug [instance_properties/getter]: isRulerActive */


// A Boolean value that indicates whether a ruler view is visible on the canvas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pencilkit/pkcanvasview/isruleractive
func (c_ CanvasView) SetIsRulerActive(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRulerActive:"), value)
}/* debug [instance_properties/setter]: isRulerActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PKCanvasView */


