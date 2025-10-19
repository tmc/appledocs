// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [EPSImageRep] class.
var (
	ePSImageRepClass     _EPSImageRepClass
	ePSImageRepClassOnce sync.Once
)

func getEPSImageRepClass() _EPSImageRepClass {
	ePSImageRepClassOnce.Do(func() {
		ePSImageRepClass = _EPSImageRepClass{objc.GetClass("NSEPSImageRep")}
	})
	return ePSImageRepClass
}

type _EPSImageRepClass struct {
	class objc.Class
}

// An interface definition for the [EPSImageRep] class.
type IEPSImageRep interface {
	IImageRep
}

// An object that can render an image from encapsulated PostScript (EPS) code. [Full Topic]
//
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




