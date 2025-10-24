// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasTrackingFrameCount */


/* debug [class_header]: Header for hasTrackingFrameCount */
// The class instance for the [hasTrackingFrameCount] class.
var (
	HasTrackingFrameCountClass     _hasTrackingFrameCountClass
	HasTrackingFrameCountClassOnce sync.Once
)

func gethasTrackingFrameCountClass() _hasTrackingFrameCountClass {
	HasTrackingFrameCountClassOnce.Do(func() {
		HasTrackingFrameCountClass = _hasTrackingFrameCountClass{objc.GetClass("hasTrackingFrameCount")}
	})
	return HasTrackingFrameCountClass
}

type _hasTrackingFrameCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasTrackingFrameCount */
// An interface definition for the [hasTrackingFrameCount] class.
type IhasTrackingFrameCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasTrackingFrameCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasTrackingFrameCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasTrackingFrameCount */
// Alloc allocates a new instance without initialization.
func (hc _hasTrackingFrameCountClass) Alloc() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasTrackingFrameCountClass) New() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasTrackingFrameCount) Init() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasTrackingFrameCount) Autorelease() hasTrackingFrameCount {
	rv := objc.Send[hasTrackingFrameCount](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasTrackingFrameCount creates a new hasTrackingFrameCount instance.
func NewhasTrackingFrameCount() hasTrackingFrameCount {
	return gethasTrackingFrameCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasTrackingFrameCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-c.ivar
type hasTrackingFrameCount struct {
	objectivec.Object
}

// hasTrackingFrameCountFrom constructs a [hasTrackingFrameCount] from an unsafe.Pointer.
func hasTrackingFrameCountFrom(ptr unsafe.Pointer) hasTrackingFrameCount {
	return hasTrackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasTrackingFrameCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasTrackingFrameCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasTrackingFrameCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasTrackingFrameCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasTrackingFrameCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasTrackingFrameCount */



