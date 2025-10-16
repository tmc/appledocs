
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewLayoutRegion] class.
var ViewLayoutRegionClass _ViewLayoutRegionClass

func init() {
	ViewLayoutRegionClass = _ViewLayoutRegionClass{objc.GetClass("NSViewLayoutRegion")}
}

type _ViewLayoutRegionClass struct {
	objc.Class
}

// An interface definition for the [ViewLayoutRegion] class.
type IViewLayoutRegion interface {
	ID() objc.ID
}

type ViewLayoutRegion struct {
	id objc.ID
}

func ViewLayoutRegionFrom(ptr unsafe.Pointer) ViewLayoutRegion {
	return ViewLayoutRegion{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ ViewLayoutRegion) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _ViewLayoutRegionClass) Alloc() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _ViewLayoutRegionClass) New() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewViewLayoutRegion creates and returns a new initialized instance.
func NewViewLayoutRegion() ViewLayoutRegion {
	return ViewLayoutRegionClass.New()
}

// Init initializes the instance.
func (v_ ViewLayoutRegion) Init() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](v_.ID(), selInit)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewLayoutRegion/marginsLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) MarginsLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.Class), objc.RegisterName("marginsLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSViewLayoutRegion/safeAreaLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) SafeAreaLayoutRegionWithCornerAdaptation(adaptivityAxis unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(vc.Class), objc.RegisterName("safeAreaLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}
