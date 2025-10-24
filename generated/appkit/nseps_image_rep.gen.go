// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [EPSImageRep] class.
var (
	EPSImageRepClass     _EPSImageRepClass
	EPSImageRepClassOnce sync.Once
)

func getEPSImageRepClass() _EPSImageRepClass {
	EPSImageRepClassOnce.Do(func() {
		EPSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}
	})
	return EPSImageRepClass
}

type _EPSImageRepClass struct {
	class objc.Class
}

// An interface definition for the [EPSImageRep] class.
type IEPSImageRep interface {
	IImageRep
	// properties:
	BoundingBox() objc.IObject /* cross-framework: Rect */
	SetBoundingBox(value objc.IObject /* cross-framework: Rect */)
	EpsRepresentation() objc.IObject /* cross-framework: Data */
	SetEpsRepresentation(value objc.IObject /* cross-framework: Data */)
	// methods:
}

// An object that can render an image from encapsulated PostScript (EPS) code.


// An object that can render an image from encapsulated PostScript (EPS) code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep
type EPSImageRep struct {
	ImageRep
}

// EPSImageRepFrom constructs a [EPSImageRep] from an unsafe.Pointer.
//
// An object that can render an image from encapsulated PostScript (EPS) code.
func EPSImageRepFrom(ptr unsafe.Pointer) EPSImageRep {
	return EPSImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ec _EPSImageRepClass) Alloc() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EPSImageRepClass) New() EPSImageRep {
	rv := objc.Send[EPSImageRep](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EPSImageRep) Init() EPSImageRep {
	rv := objc.Send[EPSImageRep](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EPSImageRep) Autorelease() EPSImageRep {
	rv := objc.Send[EPSImageRep](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEPSImageRep creates a new EPSImageRep instance.
func NewEPSImageRep() EPSImageRep {
	return getEPSImageRepClass().New()
}



// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsepsimagerep/boundingbox
func (e_ EPSImageRep) BoundingBox() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](e_.ID, objc.Sel("boundingBox"))
	return rv
}


// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsepsimagerep/boundingbox
func (e_ EPSImageRep) SetBoundingBox(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setBoundingBox:"), value)
}


// The EPS representation of the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsepsimagerep/epsrepresentation
func (e_ EPSImageRep) EpsRepresentation() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](e_.ID, objc.Sel("epsRepresentation"))
	return rv
}


// The EPS representation of the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsepsimagerep/epsrepresentation
func (e_ EPSImageRep) SetEpsRepresentation(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setEpsRepresentation:"), value)
}



