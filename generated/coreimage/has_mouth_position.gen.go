// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hasMouthPosition */


/* debug [class_header]: Header for hasMouthPosition */
// The class instance for the [hasMouthPosition] class.
var (
	HasMouthPositionClass     _hasMouthPositionClass
	HasMouthPositionClassOnce sync.Once
)

func gethasMouthPositionClass() _hasMouthPositionClass {
	HasMouthPositionClassOnce.Do(func() {
		HasMouthPositionClass = _hasMouthPositionClass{objc.GetClass("hasMouthPosition")}
	})
	return HasMouthPositionClass
}

type _hasMouthPositionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hasMouthPosition */
// An interface definition for the [hasMouthPosition] class.
type IhasMouthPosition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hasMouthPosition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hasMouthPosition */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hasMouthPosition */
// Alloc allocates a new instance without initialization.
func (hc _hasMouthPositionClass) Alloc() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hasMouthPositionClass) New() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hasMouthPosition) Init() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hasMouthPosition) Autorelease() hasMouthPosition {
	rv := objc.Send[hasMouthPosition](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhasMouthPosition creates a new hasMouthPosition instance.
func NewhasMouthPosition() hasMouthPosition {
	return gethasMouthPositionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hasMouthPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasMouthPosition-c.ivar
type hasMouthPosition struct {
	objectivec.Object
}

// hasMouthPositionFrom constructs a [hasMouthPosition] from an unsafe.Pointer.
func hasMouthPositionFrom(ptr unsafe.Pointer) hasMouthPosition {
	return hasMouthPosition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hasMouthPosition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hasMouthPosition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hasMouthPosition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hasMouthPosition */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hasMouthPosition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hasMouthPosition */



