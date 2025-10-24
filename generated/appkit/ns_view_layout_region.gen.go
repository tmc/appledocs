// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSViewLayoutRegion */


/* debug [class_header]: Header for NSViewLayoutRegion */
// The class instance for the [ViewLayoutRegion] class.
var (
	ViewLayoutRegionClass     _ViewLayoutRegionClass
	ViewLayoutRegionClassOnce sync.Once
)

func getViewLayoutRegionClass() _ViewLayoutRegionClass {
	ViewLayoutRegionClassOnce.Do(func() {
		ViewLayoutRegionClass = _ViewLayoutRegionClass{objc.GetClass("NSViewLayoutRegion")}
	})
	return ViewLayoutRegionClass
}

type _ViewLayoutRegionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ViewLayoutRegion */
// An interface definition for the [ViewLayoutRegion] class.
type IViewLayoutRegion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ViewLayoutRegion */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ViewLayoutRegion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ViewLayoutRegion */
// Alloc allocates a new instance without initialization.
func (vc _ViewLayoutRegionClass) Alloc() ViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getViewLayoutRegionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ViewLayoutRegion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion
type ViewLayoutRegion struct {
	objectivec.Object
}

// ViewLayoutRegionFrom constructs a [ViewLayoutRegion] from an unsafe.Pointer.
func ViewLayoutRegionFrom(ptr unsafe.Pointer) ViewLayoutRegion {
	return ViewLayoutRegion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ViewLayoutRegion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ViewLayoutRegion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/marginsLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) MarginsLayoutRegionWithCornerAdaptation(adaptivityAxis ViewLayoutRegionAdaptivityAxis) IViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.class), objc.Sel("marginsLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MarginsLayoutRegionWithCornerAdaptation) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/safeAreaLayoutRegionWithCornerAdaptation:
func (vc _ViewLayoutRegionClass) SafeAreaLayoutRegionWithCornerAdaptation(adaptivityAxis ViewLayoutRegionAdaptivityAxis) IViewLayoutRegion {
	rv := objc.Send[ViewLayoutRegion](objc.ID(vc.class), objc.Sel("safeAreaLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SafeAreaLayoutRegionWithCornerAdaptation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ViewLayoutRegion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ViewLayoutRegion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ViewLayoutRegion */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSViewLayoutRegion */



