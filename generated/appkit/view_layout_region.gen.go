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


