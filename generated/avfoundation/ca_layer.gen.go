// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CALayer */


/* debug [class_header]: Header for CALayer */
// The class instance for the [Layer] class.
var (
	LayerClass     _LayerClass
	LayerClassOnce sync.Once
)

func getLayerClass() _LayerClass {
	LayerClassOnce.Do(func() {
		LayerClass = _LayerClass{objc.GetClass("CALayer")}
	})
	return LayerClass
}

type _LayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Layer */
// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Layer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Layer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Layer */
// Alloc allocates a new instance without initialization.
func (lc _LayerClass) Alloc() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayerClass) New() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Layer) Init() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Layer) Autorelease() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayer creates a new Layer instance.
func NewLayer() Layer {
	return getLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Layer */
// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type Layer struct {
	objectivec.Object
}

// LayerFrom constructs a [Layer] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func LayerFrom(ptr unsafe.Pointer) Layer {
	return Layer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Layer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Layer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Layer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Layer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Layer */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CALayer */



