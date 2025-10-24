// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class layerCount */


/* debug [class_header]: Header for layerCount */
// The class instance for the [layerCount] class.
var (
	LayerCountClass     _layerCountClass
	LayerCountClassOnce sync.Once
)

func getlayerCountClass() _layerCountClass {
	LayerCountClassOnce.Do(func() {
		LayerCountClass = _layerCountClass{objc.GetClass("layerCount")}
	})
	return LayerCountClass
}

type _layerCountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for layerCount */
// An interface definition for the [layerCount] class.
type IlayerCount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for layerCount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for layerCount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for layerCount */
// Alloc allocates a new instance without initialization.
func (lc _layerCountClass) Alloc() layerCount {
	rv := objc.Send[layerCount](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _layerCountClass) New() layerCount {
	rv := objc.Send[layerCount](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ layerCount) Init() layerCount {
	rv := objc.Send[layerCount](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ layerCount) Autorelease() layerCount {
	rv := objc.Send[layerCount](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlayerCount creates a new layerCount instance.
func NewlayerCount() layerCount {
	return getlayerCountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for layerCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar
type layerCount struct {
	objectivec.Object
}

// layerCountFrom constructs a [layerCount] from an unsafe.Pointer.
func layerCountFrom(ptr unsafe.Pointer) layerCount {
	return layerCount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for layerCount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for layerCount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for layerCount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for layerCount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for layerCount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class layerCount */



