// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class bottomRight */


/* debug [class_header]: Header for bottomRight */
// The class instance for the [bottomRight] class.
var (
	BottomRightClass     _bottomRightClass
	BottomRightClassOnce sync.Once
)

func getbottomRightClass() _bottomRightClass {
	BottomRightClassOnce.Do(func() {
		BottomRightClass = _bottomRightClass{objc.GetClass("bottomRight")}
	})
	return BottomRightClass
}

type _bottomRightClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for bottomRight */
// An interface definition for the [bottomRight] class.
type IbottomRight interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for bottomRight */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for bottomRight */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for bottomRight */
// Alloc allocates a new instance without initialization.
func (bc _bottomRightClass) Alloc() bottomRight {
	rv := objc.Send[bottomRight](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _bottomRightClass) New() bottomRight {
	rv := objc.Send[bottomRight](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bottomRight) Init() bottomRight {
	rv := objc.Send[bottomRight](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bottomRight) Autorelease() bottomRight {
	rv := objc.Send[bottomRight](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewbottomRight creates a new bottomRight instance.
func NewbottomRight() bottomRight {
	return getbottomRightClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for bottomRight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomRight-c.ivar
type bottomRight struct {
	objectivec.Object
}

// bottomRightFrom constructs a [bottomRight] from an unsafe.Pointer.
func bottomRightFrom(ptr unsafe.Pointer) bottomRight {
	return bottomRight{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for bottomRight *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for bottomRight */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for bottomRight */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for bottomRight */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for bottomRight */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class bottomRight */



