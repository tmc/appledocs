// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class containsMipmaps */


/* debug [class_header]: Header for containsMipmaps */
// The class instance for the [containsMipmaps] class.
var (
	ContainsMipmapsClass     _containsMipmapsClass
	ContainsMipmapsClassOnce sync.Once
)

func getcontainsMipmapsClass() _containsMipmapsClass {
	ContainsMipmapsClassOnce.Do(func() {
		ContainsMipmapsClass = _containsMipmapsClass{objc.GetClass("containsMipmaps")}
	})
	return ContainsMipmapsClass
}

type _containsMipmapsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for containsMipmaps */
// An interface definition for the [containsMipmaps] class.
type IcontainsMipmaps interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for containsMipmaps */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for containsMipmaps */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for containsMipmaps */
// Alloc allocates a new instance without initialization.
func (cc _containsMipmapsClass) Alloc() containsMipmaps {
	rv := objc.Send[containsMipmaps](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _containsMipmapsClass) New() containsMipmaps {
	rv := objc.Send[containsMipmaps](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ containsMipmaps) Init() containsMipmaps {
	rv := objc.Send[containsMipmaps](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ containsMipmaps) Autorelease() containsMipmaps {
	rv := objc.Send[containsMipmaps](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewcontainsMipmaps creates a new containsMipmaps instance.
func NewcontainsMipmaps() containsMipmaps {
	return getcontainsMipmapsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for containsMipmaps */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/containsMipmaps-c.ivar
type containsMipmaps struct {
	objectivec.Object
}

// containsMipmapsFrom constructs a [containsMipmaps] from an unsafe.Pointer.
func containsMipmapsFrom(ptr unsafe.Pointer) containsMipmaps {
	return containsMipmaps{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for containsMipmaps *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for containsMipmaps */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for containsMipmaps */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for containsMipmaps */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for containsMipmaps */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class containsMipmaps */



