// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class height */


/* debug [class_header]: Header for height */
// The class instance for the [height] class.
var (
	HeightClass     _heightClass
	HeightClassOnce sync.Once
)

func getheightClass() _heightClass {
	HeightClassOnce.Do(func() {
		HeightClass = _heightClass{objc.GetClass("height")}
	})
	return HeightClass
}

type _heightClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for height */
// An interface definition for the [height] class.
type Iheight interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for height */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for height */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for height */
// Alloc allocates a new instance without initialization.
func (hc _heightClass) Alloc() height {
	rv := objc.Send[height](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _heightClass) New() height {
	rv := objc.Send[height](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ height) Init() height {
	rv := objc.Send[height](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ height) Autorelease() height {
	rv := objc.Send[height](h_.ID, objc.Sel("autorelease"))
	return rv
}

// Newheight creates a new height instance.
func Newheight() height {
	return getheightClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for height */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/height-c.ivar
type height struct {
	objectivec.Object
}

// heightFrom constructs a [height] from an unsafe.Pointer.
func heightFrom(ptr unsafe.Pointer) height {
	return height{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for height *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for height */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for height */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for height */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for height */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class height */



