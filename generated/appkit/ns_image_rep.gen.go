// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ImageRep] class.
var (
	ImageRepClass     _ImageRepClass
	ImageRepClassOnce sync.Once
)

func getImageRepClass() _ImageRepClass {
	ImageRepClassOnce.Do(func() {
		ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}
	})
	return ImageRepClass
}

type _ImageRepClass struct {
	class objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	objectivec.IObject
	Draw() bool
	DrawAtPoint(point coregraphics.CGPoint) bool
	DrawInRect(rect coregraphics.CGRect) bool
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect coregraphics.CGRect, srcSpacePortionRect coregraphics.CGRect, op unsafe.Pointer, requestedAlpha float64, respectContextIsFlipped bool, hints unsafe.Pointer) bool
}

// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
//
// The class is called “semiabstract” because it has some instance variables and implementation of its own, in addition to defining subclasses. Although an subclass can be used directly, it is typically accessed through an object, which manages a group of image representations, choosing the best one for the current output device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep
type ImageRep struct {
	objectivec.Object
}

// ImageRepFrom constructs a [ImageRep] from an unsafe.Pointer.
//
// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data.
func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _ImageRepClass) New() ImageRep {
	rv := objc.Send[ImageRep](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ImageRep) Init() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ImageRep) Autorelease() ImageRep {
	rv := objc.Send[ImageRep](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewImageRep creates a new ImageRep instance.
func NewImageRep() ImageRep {
	return getImageRepClass().New()
}


// Creates and returns an image representation object from data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewImageRepWithCoder(coder unsafe.Pointer) ImageRep {
	instance := getImageRepClass().Alloc()
	rv := objc.Send[ImageRep](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Creates and returns an image representation object using the contents of the specified pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func NewImageRepWithPasteboard(pasteboard unsafe.Pointer) ImageRep {
	rv := objc.Send[ImageRep](objc.ID(getImageRepClass().class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return rv
}


// Returns the image representation subclass that handles the specified type of data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(for:)
func (ic _ImageRepClass) ImageRepClassForData(data unsafe.Pointer) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(ic.class), objc.Sel("imageRepClassForData:"), data)
	return rv
}

// Creates and returns an image representation object using the contents of the specified pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ic.class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return rv
}

// Implemented by subclasses to draw the image in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw()
func (i_ ImageRep) Draw() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("draw"))
	return rv
}

// Draws the image representation’s image data at the specified point in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(at:)
func (i_ ImageRep) DrawAtPoint(point coregraphics.CGPoint) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawAtPoint:"), point)
	return rv
}

// Draws the image, scaling it (as needed) to fit the specified rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:)
func (i_ ImageRep) DrawInRect(rect coregraphics.CGRect) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:"), rect)
	return rv
}

// Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:from:operation:fraction:respectFlipped:hints:)
func (i_ ImageRep) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect coregraphics.CGRect, srcSpacePortionRect coregraphics.CGRect, op unsafe.Pointer, requestedAlpha float64, respectContextIsFlipped bool, hints unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}


