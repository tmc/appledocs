// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewLayoutRegion] class.
var ViewLayoutRegionClass objc.Class

func init() {
	ViewLayoutRegionClass = objc.GetClass("NSViewLayoutRegion")
}

type ViewLayoutRegion struct {
	objc.ID
}

func ViewLayoutRegionFrom(ptr unsafe.Pointer) ViewLayoutRegion {
	return ViewLayoutRegion{
		ID: objc.ID(ptr),
	}
}


//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewLayoutRegion/marginsLayoutRegionWithCornerAdaptation:
func (vc ViewLayoutRegion) MarginsLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("marginsLayoutRegionWithCornerAdaptation:")
	ret := objc.ID(ViewLayoutRegionClass).Send(sel, adaptivityAxis)
	return unsafe.Pointer(ret)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewLayoutRegion/safeAreaLayoutRegionWithCornerAdaptation:
func (vc ViewLayoutRegion) SafeAreaLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("safeAreaLayoutRegionWithCornerAdaptation:")
	ret := objc.ID(ViewLayoutRegionClass).Send(sel, adaptivityAxis)
	return unsafe.Pointer(ret)
}

