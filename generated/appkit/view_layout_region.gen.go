// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ViewLayoutRegion] class.
var viewLayoutRegionClass = _ViewLayoutRegionClass{objc.GetClass("NSViewLayoutRegion")}

type _ViewLayoutRegionClass struct {
	class objc.Class
}

// An interface definition for the [ViewLayoutRegion] class.
type IViewLayoutRegion interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion

type ViewLayoutRegion struct {
	objectivec.Object
}

// ViewLayoutRegionFrom constructs a [ViewLayoutRegion] from an unsafe.Pointer.
func ViewLayoutRegionFrom(ptr unsafe.Pointer) ViewLayoutRegion {
	return ViewLayoutRegion{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (vc _ViewLayoutRegionClass) Alloc() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (vc _ViewLayoutRegionClass) New() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ViewLayoutRegion) Init() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ViewLayoutRegion) Autorelease() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewViewLayoutRegion creates a new ViewLayoutRegion instance.
func NewViewLayoutRegion() ViewLayoutRegion {
	return viewLayoutRegionClass.New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/marginsLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) MarginsLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("marginsLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/safeAreaLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) SafeAreaLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.class), objc.Sel("safeAreaLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}


