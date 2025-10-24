// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSBindingSelectionMarker */


/* debug [class_header]: Header for NSBindingSelectionMarker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BindingSelectionMarker */
// An interface definition for the [BindingSelectionMarker] class.
type IBindingSelectionMarker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BindingSelectionMarker */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BindingSelectionMarker */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BindingSelectionMarker */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BindingSelectionMarker */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker
type BindingSelectionMarker struct {
	objectivec.Object
}

// BindingSelectionMarkerFrom constructs a [BindingSelectionMarker] from an unsafe.Pointer.
func BindingSelectionMarkerFrom(ptr unsafe.Pointer) BindingSelectionMarker {
	return BindingSelectionMarker{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BindingSelectionMarker *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BindingSelectionMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/defaultPlaceholder(for:on:withBinding:)
func (bc _BindingSelectionMarkerClass) DefaultPlaceholderForMarkerOnClassWithBinding(marker IBindingSelectionMarker, objectClass objc.Class, binding string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("defaultPlaceholderForMarker:onClass:withBinding:"), marker, objectClass, objc.String(binding))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultPlaceholderForMarkerOnClassWithBinding) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/setDefaultPlaceholder(_:for:on:withBinding:)
func (bc _BindingSelectionMarkerClass) SetDefaultPlaceholderForMarkerOnClassWithBinding(placeholder objc.IObject, marker IBindingSelectionMarker, objectClass objc.Class, binding string) {
	objc.Send[objc.ID](objc.ID(bc.class), objc.Sel("setDefaultPlaceholder:forMarker:onClass:withBinding:"), placeholder, marker, objectClass, objc.String(binding))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetDefaultPlaceholderForMarkerOnClassWithBinding) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BindingSelectionMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/multipleValues
func (bc _BindingSelectionMarkerClass) MultipleValuesSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("multipleValuesSelectionMarker"))
	return rv
}/* debug [class_properties_class/property]: multipleValuesSelectionMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/noSelection
func (bc _BindingSelectionMarkerClass) NoSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("noSelectionMarker"))
	return rv
}/* debug [class_properties_class/property]: noSelectionMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (bc _BindingSelectionMarkerClass) NotApplicableSelectionMarker() BindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](objc.ID(bc.class), objc.Sel("notApplicableSelectionMarker"))
	return rv
}/* debug [class_properties_class/property]: notApplicableSelectionMarker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BindingSelectionMarker */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BindingSelectionMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/multipleValues
func (b_ BindingSelectionMarker) MultipleValuesSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("multipleValuesSelectionMarker"))
	return rv
}/* debug [instance_properties/getter]: multipleValuesSelectionMarker */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/noSelection
func (b_ BindingSelectionMarker) NoSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("noSelectionMarker"))
	return rv
}/* debug [instance_properties/getter]: noSelectionMarker */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBindingSelectionMarker/notApplicable
func (b_ BindingSelectionMarker) NotApplicableSelectionMarker() IBindingSelectionMarker {
	rv := objc.Send[BindingSelectionMarker](b_.ID, objc.Sel("notApplicableSelectionMarker"))
	return rv
}/* debug [instance_properties/getter]: notApplicableSelectionMarker */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBindingSelectionMarker */



