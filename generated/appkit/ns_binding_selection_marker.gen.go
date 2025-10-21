// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BindingSelectionMarker] class.
var (
	BindingSelectionMarkerClass     _BindingSelectionMarkerClass
	BindingSelectionMarkerClassOnce sync.Once
)

func getBindingSelectionMarkerClass() _BindingSelectionMarkerClass {
	BindingSelectionMarkerClassOnce.Do(func() {
		BindingSelectionMarkerClass = _BindingSelectionMarkerClass{objc.GetClass("NSBindingSelectionMarker")}
	})
	return BindingSelectionMarkerClass
}

type _BindingSelectionMarkerClass struct {
	class objc.Class
}

// An interface definition for the [BindingSelectionMarker] class.
type IBindingSelectionMarker interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker
type BindingSelectionMarker struct {
	objectivec.Object
}

// BindingSelectionMarkerFrom constructs a [BindingSelectionMarker] from an unsafe.Pointer.
func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BindingSelectionMarkerClass) Alloc() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BindingSelectionMarkerClass) New() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BindingSelectionMarker) Init() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BindingSelectionMarker) Autorelease() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBindingSelectionMarker creates a new BindingSelectionMarker instance.
func NewBindingSelectionMarker() BindingSelectionMarker {
	return getBindingSelectionMarkerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (bc _BindingSelectionMarkerClass) NotApplicableSelectionMarker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("notApplicableSelectionMarker"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (b_ BindingSelectionMarker) NotApplicableSelectionMarker() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("notApplicableSelectionMarker"))
	return rv
}



