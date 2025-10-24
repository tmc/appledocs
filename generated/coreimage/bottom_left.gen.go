// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class bottomLeft */


/* debug [class_header]: Header for bottomLeft */
// The class instance for the [bottomLeft] class.
var (
	BottomLeftClass     _bottomLeftClass
	BottomLeftClassOnce sync.Once
)

func getbottomLeftClass() _bottomLeftClass {
	BottomLeftClassOnce.Do(func() {
		BottomLeftClass = _bottomLeftClass{objc.GetClass("bottomLeft")}
	})
	return BottomLeftClass
}

type _bottomLeftClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for bottomLeft */
// An interface definition for the [bottomLeft] class.
type IbottomLeft interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for bottomLeft */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for bottomLeft */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for bottomLeft */
// Alloc allocates a new instance without initialization.
func (bc _bottomLeftClass) Alloc() bottomLeft {
	rv := objc.Send[bottomLeft](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _bottomLeftClass) New() bottomLeft {
	rv := objc.Send[bottomLeft](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bottomLeft) Init() bottomLeft {
	rv := objc.Send[bottomLeft](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bottomLeft) Autorelease() bottomLeft {
	rv := objc.Send[bottomLeft](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewbottomLeft creates a new bottomLeft instance.
func NewbottomLeft() bottomLeft {
	return getbottomLeftClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for bottomLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeFeature/bottomLeft-c.ivar
type bottomLeft struct {
	objectivec.Object
}

// bottomLeftFrom constructs a [bottomLeft] from an unsafe.Pointer.
func bottomLeftFrom(ptr unsafe.Pointer) bottomLeft {
	return bottomLeft{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for bottomLeft *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for bottomLeft */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for bottomLeft */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for bottomLeft */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for bottomLeft */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class bottomLeft */



