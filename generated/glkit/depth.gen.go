// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class depth */


/* debug [class_header]: Header for depth */
// The class instance for the [depth] class.
var (
	DepthClass     _depthClass
	DepthClassOnce sync.Once
)

func getdepthClass() _depthClass {
	DepthClassOnce.Do(func() {
		DepthClass = _depthClass{objc.GetClass("depth")}
	})
	return DepthClass
}

type _depthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for depth */
// An interface definition for the [depth] class.
type Idepth interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for depth */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for depth */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for depth */
// Alloc allocates a new instance without initialization.
func (dc _depthClass) Alloc() depth {
	rv := objc.Send[depth](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _depthClass) New() depth {
	rv := objc.Send[depth](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ depth) Init() depth {
	rv := objc.Send[depth](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ depth) Autorelease() depth {
	rv := objc.Send[depth](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdepth creates a new depth instance.
func Newdepth() depth {
	return getdepthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for depth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/depth-c.ivar
type depth struct {
	objectivec.Object
}

// depthFrom constructs a [depth] from an unsafe.Pointer.
func depthFrom(ptr unsafe.Pointer) depth {
	return depth{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for depth *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for depth */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for depth */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for depth */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for depth */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class depth */



