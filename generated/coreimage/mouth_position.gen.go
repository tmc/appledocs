// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class mouthPosition */


/* debug [class_header]: Header for mouthPosition */
// The class instance for the [mouthPosition] class.
var (
	MouthPositionClass     _mouthPositionClass
	MouthPositionClassOnce sync.Once
)

func getmouthPositionClass() _mouthPositionClass {
	MouthPositionClassOnce.Do(func() {
		MouthPositionClass = _mouthPositionClass{objc.GetClass("mouthPosition")}
	})
	return MouthPositionClass
}

type _mouthPositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for mouthPosition */
// An interface definition for the [mouthPosition] class.
type ImouthPosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for mouthPosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for mouthPosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for mouthPosition */
// Alloc allocates a new instance without initialization.
func (mc _mouthPositionClass) Alloc() mouthPosition {
	rv := objc.Send[mouthPosition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _mouthPositionClass) New() mouthPosition {
	rv := objc.Send[mouthPosition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mouthPosition) Init() mouthPosition {
	rv := objc.Send[mouthPosition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mouthPosition) Autorelease() mouthPosition {
	rv := objc.Send[mouthPosition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmouthPosition creates a new mouthPosition instance.
func NewmouthPosition() mouthPosition {
	return getmouthPositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for mouthPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/mouthPosition-c.ivar
type mouthPosition struct {
	objectivec.Object
}

// mouthPositionFrom constructs a [mouthPosition] from an unsafe.Pointer.
func mouthPositionFrom(ptr unsafe.Pointer) mouthPosition {
	return mouthPosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for mouthPosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for mouthPosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for mouthPosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for mouthPosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for mouthPosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class mouthPosition */



