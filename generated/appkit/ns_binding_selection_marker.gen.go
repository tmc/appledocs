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
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _BindingSelectionMarkerClass) Alloc() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker
type BindingSelectionMarker struct {
	objectivec.Object
}

// BindingSelectionMarkerFrom constructs a [BindingSelectionMarker] from an unsafe.Pointer.
func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/defaultPlaceholder(for:on:withBinding:)
func (bc _BindingSelectionMarkerClass) DefaultPlaceholderForMarkerOnClassWithBinding(marker IBindingSelectionMarker, objectClass objc.Class, binding string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("defaultPlaceholderForMarker:onClass:withBinding:"), marker, objectClass, objc.String(binding))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/setDefaultPlaceholder(_:for:on:withBinding:)
func (bc _BindingSelectionMarkerClass) SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder objectivec.IObject, marker IBindingSelectionMarker, objectClass objc.Class, binding string) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("setDefaultPlaceholder:forMarker:onClass:withBinding:"), placeholder, marker, objectClass, objc.String(binding))
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/multipleValues
func (bc _BindingSelectionMarkerClass) MultipleValuesSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("multipleValuesSelectionMarker"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/noSelection
func (bc _BindingSelectionMarkerClass) NoSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("noSelectionMarker"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (bc _BindingSelectionMarkerClass) NotApplicableSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("notApplicableSelectionMarker"))
	return rv
}











// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/multipleValues
func (b_ BindingSelectionMarker) MultipleValuesSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("multipleValuesSelectionMarker"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/noSelection
func (b_ BindingSelectionMarker) NoSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("noSelectionMarker"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (b_ BindingSelectionMarker) NotApplicableSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("notApplicableSelectionMarker"))
	return rv
}








