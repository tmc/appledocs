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
	EPSRepresentation() objc.IObject /* cross-framework: NSData */
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



// Returns a representation of an image initialized with the specified EPS data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/init(data:)
func NewEPSImageRepWithData(epsData objc.IObject /* cross-framework: NSData */) EPSImageRep {
	instance := getEPSImageRepClass().Alloc()
	rv := objc.Send[EPSImageRep](instance.ID, objc.Sel("initWithData:"), epsData)
	rv.Autorelease()
	return rv
}



// Creates and returns a representation of an image initialized with the specified EPS data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/imageRepWithData:
func (ec _EPSImageRepClass) ImageRepWithData(epsData objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ec.class), objc.Sel("imageRepWithData:"), epsData)
	return rv
}


// The rectangle that bounds the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/boundingBox
func (e_ EPSImageRep) BoundingBox() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](e_.ID, objc.Sel("boundingBox"))
	return rv
}


// The EPS representation of the image representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEPSImageRep/epsRepresentation
func (e_ EPSImageRep) EPSRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](e_.ID, objc.Sel("EPSRepresentation"))
	return rv
}


