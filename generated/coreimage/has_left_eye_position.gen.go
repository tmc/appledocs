// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasLeftEyePosition */


/* debug [class_header]: Header for hasLeftEyePosition */
// The class instance for the [hasLeftEyePosition] class.
var (
	HasLeftEyePositionClass     _hasLeftEyePositionClass
	HasLeftEyePositionClassOnce sync.Once
)

func gethasLeftEyePositionClass() _hasLeftEyePositionClass {
	HasLeftEyePositionClassOnce.Do(func() {
		HasLeftEyePositionClass = _hasLeftEyePositionClass{objc.GetClass("hasLeftEyePosition")}
	})
	return HasLeftEyePositionClass
}

type _hasLeftEyePositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasLeftEyePosition */
// An interface definition for the [hasLeftEyePosition] class.
type IhasLeftEyePosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasLeftEyePosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasLeftEyePosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasLeftEyePosition */
// Alloc allocates a new instance without initialization.
func (hc _hasLeftEyePositionClass) Alloc() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasLeftEyePositionClass) New() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasLeftEyePosition) Init() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasLeftEyePosition) Autorelease() hasLeftEyePosition {
	rv := objc.Send[hasLeftEyePosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasLeftEyePosition creates a new hasLeftEyePosition instance.
func NewhasLeftEyePosition() hasLeftEyePosition {
	return gethasLeftEyePositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasLeftEyePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-c.ivar
type hasLeftEyePosition struct {
	objectivec.Object
}

// hasLeftEyePositionFrom constructs a [hasLeftEyePosition] from an unsafe.Pointer.
func hasLeftEyePositionFrom(ptr unsafe.Pointer) hasLeftEyePosition {
	return hasLeftEyePosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasLeftEyePosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasLeftEyePosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasLeftEyePosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasLeftEyePosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasLeftEyePosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasLeftEyePosition */



