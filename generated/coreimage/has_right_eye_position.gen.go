// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasRightEyePosition */


/* debug [class_header]: Header for hasRightEyePosition */
// The class instance for the [hasRightEyePosition] class.
var (
	HasRightEyePositionClass     _hasRightEyePositionClass
	HasRightEyePositionClassOnce sync.Once
)

func gethasRightEyePositionClass() _hasRightEyePositionClass {
	HasRightEyePositionClassOnce.Do(func() {
		HasRightEyePositionClass = _hasRightEyePositionClass{objc.GetClass("hasRightEyePosition")}
	})
	return HasRightEyePositionClass
}

type _hasRightEyePositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasRightEyePosition */
// An interface definition for the [hasRightEyePosition] class.
type IhasRightEyePosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasRightEyePosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasRightEyePosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasRightEyePosition */
// Alloc allocates a new instance without initialization.
func (hc _hasRightEyePositionClass) Alloc() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasRightEyePositionClass) New() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasRightEyePosition) Init() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasRightEyePosition) Autorelease() hasRightEyePosition {
	rv := objc.Send[hasRightEyePosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasRightEyePosition creates a new hasRightEyePosition instance.
func NewhasRightEyePosition() hasRightEyePosition {
	return gethasRightEyePositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasRightEyePosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-c.ivar
type hasRightEyePosition struct {
	objectivec.Object
}

// hasRightEyePositionFrom constructs a [hasRightEyePosition] from an unsafe.Pointer.
func hasRightEyePositionFrom(ptr unsafe.Pointer) hasRightEyePosition {
	return hasRightEyePosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasRightEyePosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasRightEyePosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasRightEyePosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasRightEyePosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasRightEyePosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasRightEyePosition */



