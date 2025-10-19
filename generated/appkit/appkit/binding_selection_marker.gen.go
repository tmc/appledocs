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
	bindingSelectionMarkerClass     _BindingSelectionMarkerClass
	bindingSelectionMarkerClassOnce sync.Once
)

func getBindingSelectionMarkerClass() _BindingSelectionMarkerClass {
	bindingSelectionMarkerClassOnce.Do(func() {
		bindingSelectionMarkerClass = _BindingSelectionMarkerClass{objc.GetClass("NSBindingSelectionMarker")}
	})
	return bindingSelectionMarkerClass
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

// New creates and returns a new instance with a +1 retain count.
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




