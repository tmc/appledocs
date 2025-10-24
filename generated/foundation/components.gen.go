// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class components */


/* debug [class_header]: Header for components */
// The class instance for the [components] class.
var (
	ComponentsClass     _componentsClass
	ComponentsClassOnce sync.Once
)

func getcomponentsClass() _componentsClass {
	ComponentsClassOnce.Do(func() {
		ComponentsClass = _componentsClass{objc.GetClass("components")}
	})
	return ComponentsClass
}

type _componentsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for components */
// An interface definition for the [components] class.
type Icomponents interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for components */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for components */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for components */
// Alloc allocates a new instance without initialization.
func (cc _componentsClass) Alloc() components {
	rv := objc.Send[components](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _componentsClass) New() components {
	rv := objc.Send[components](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ components) Init() components {
	rv := objc.Send[components](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ components) Autorelease() components {
	rv := objc.Send[components](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcomponents creates a new components instance.
func Newcomponents() components {
	return getcomponentsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for components */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/components-c.ivar
type components struct {
	objectivec.Object
}

// componentsFrom constructs a [components] from an unsafe.Pointer.
func componentsFrom(ptr unsafe.Pointer) components {
	return components{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for components *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for components */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for components */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for components */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for components */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class components */



