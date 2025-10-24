// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCustomImageRep */


/* debug [class_header]: Header for NSCustomImageRep */
// The class instance for the [CustomImageRep] class.
var (
	CustomImageRepClass     _CustomImageRepClass
	CustomImageRepClassOnce sync.Once
)

func getCustomImageRepClass() _CustomImageRepClass {
	CustomImageRepClassOnce.Do(func() {
		CustomImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}
	})
	return CustomImageRepClass
}

type _CustomImageRepClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CustomImageRep */
// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
	
/* debug [class_interface_properties]: Properties for CustomImageRep */
	// properties:
	Delegate() objc.ID
	DrawSelector() objc.SEL
	DrawingHandler() func(unsafe.Pointer) unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CustomImageRep */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CustomImageRep */
// Alloc allocates a new instance without initialization.
func (cc _CustomImageRepClass) Alloc() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CustomImageRepClass) New() CustomImageRep {
	rv := objc.Send[CustomImageRep](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomImageRep) Init() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomImageRep) Autorelease() CustomImageRep {
	rv := objc.Send[CustomImageRep](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomImageRep creates a new CustomImageRep instance.
func NewCustomImageRep() CustomImageRep {
	return getCustomImageRepClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CustomImageRep */
// An object that uses a delegate object to render an image from a custom format.
//
// When called upon to produce an image, an sends a message to its delegate to do the actual drawing. You can use this class to support custom image formats without going to the trouble of subclassing directly.


// An object that uses a delegate object to render an image from a custom format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep
type CustomImageRep struct {
	ImageRep
}

// CustomImageRepFrom constructs a [CustomImageRep] from an unsafe.Pointer.
//
// An object that uses a delegate object to render an image from a custom format.
func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CustomImageRep */

// Returns a representation of an image initialized with the specified delegate information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(draw:delegate:)
func NewCustomImageRepWithDrawSelectorDelegate(selector objc.SEL, delegate objc.IObject) CustomImageRep {
	instance := getCustomImageRepClass().Alloc()
	rv := objc.Send[CustomImageRep](instance.ID, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCustomImageRepWithDrawSelectorDelegate */


// Initializes a representation of an image of the specified size and flipped status, using a block to draw its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(size:flipped:drawingHandler:)
func NewCustomImageRepWithSizeFlippedDrawingHandler(size Size /* not a class type */, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) CustomImageRep {
	instance := getCustomImageRepClass().Alloc()
	rv := objc.Send[CustomImageRep](instance.ID, objc.Sel("initWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCustomImageRepWithSizeFlippedDrawingHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CustomImageRep */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CustomImageRep */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CustomImageRep */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CustomImageRep */

// The delegate object that renders the image for the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/delegate
func (c_ CustomImageRep) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The selector for the delegate’s drawing method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawSelector
func (c_ CustomImageRep) DrawSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("drawSelector"))
	return rv
}/* debug [instance_properties/getter]: drawSelector */


// The destination rectangle of the drawing handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawingHandler
func (c_ CustomImageRep) DrawingHandler() func(unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[func(unsafe.Pointer) unsafe.Pointer](c_.ID, objc.Sel("drawingHandler"))
	return rv
}/* debug [instance_properties/getter]: drawingHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCustomImageRep */


