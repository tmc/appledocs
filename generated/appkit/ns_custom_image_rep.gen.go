// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
	

	// properties:
	DrawSelector() objc.SEL
	DrawingHandler() func(unsafe.Pointer) unsafe.Pointer


	

	// methods:


}





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






// Returns a representation of an image initialized with the specified delegate information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(draw:delegate:)
func NewCustomImageRepWithDrawSelectorDelegate(selector objc.SEL, delegate objectivec.IObject) CustomImageRep {
	instance := getCustomImageRepClass().Alloc()
	rv := objc.Send[CustomImageRep](instance.ID, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	rv.Autorelease()
	return rv
}


// Initializes a representation of an image of the specified size and flipped status, using a block to draw its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(size:flipped:drawingHandler:)
func NewCustomImageRepWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler unsafe.Pointer) CustomImageRep {
	instance := getCustomImageRepClass().Alloc()
	rv := objc.Send[CustomImageRep](instance.ID, objc.Sel("initWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	rv.Autorelease()
	return rv
}






















// The selector for the delegate’s drawing method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawSelector
func (c_ CustomImageRep) DrawSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c_.ID, objc.Sel("drawSelector"))
	return rv
}


// The destination rectangle of the drawing handler block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawingHandler
func (c_ CustomImageRep) DrawingHandler() func(unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[func(unsafe.Pointer) unsafe.Pointer](c_.ID, objc.Sel("drawingHandler"))
	return rv
}







